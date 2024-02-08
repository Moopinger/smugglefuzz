/*
2024 Moopinger
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/moopinger/smugglefuzz/lib"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Run a scan using multiple different gadgets.",
	Long:  `Scan for downgrade attacks. Use the default provided wordlist(gadgetlist) or your own.`,
	Example: `	smugglefuzz scan -u https://example.com/ --confirm 
	smugglefuzz scan -u https://example.com/ --filter 200 --confirm
	smugglefuzz scan -u https://example.com/ -w wordlist.txt -t 10 --confirm
	smugglefuzz scan --dc -u https://example.com/ -w wordlist.txt -x PUT --confirm
	
	//Multiple targets? just use -f instead of -u and provide a file with the targets in it:
	
	smugglefuzz scan -f multiple_targets.txt --confirm -t 10
	smugglefuzz scan -f multiple_targets.txt -w wordlist.txt --confirm -s ./save-success.txt 
	smugglefuzz scan -f multiple_targets.txt -w wordlist.txt -H "Cookie: date=...; session=...;" --confirm -s ./save-success.txt -x PUT`,

	Run: func(cmd *cobra.Command, args []string) {

		//if the user has not provided a target or a wordlist, then we should return the help message
		if cmd.Flag("url").Value.String() == "" && cmd.Flag("file").Value.String() == "" {
			fmt.Println("You must provide either a url (-u) or file of urls (-f) or a file containing multiple targets")
			cmd.Help()
			return
		}

		fmt.Print(lib.Banner)

		var targets []string

		//will make the switch to Flags().StringVARP at some point
		method, _ := cmd.Flags().GetString("method")
		targetFile, _ := cmd.Flags().GetString("file")
		singleTarget, _ := cmd.Flags().GetString("url")
		gadgetList, _ := cmd.Flags().GetString("wordlist")
		saveSuccessfulRequests, _ := cmd.Flags().GetString("save-success")
		enableConfirmation, _ := cmd.Flags().GetBool("confirm")
		routineCount, _ := cmd.Flags().GetInt("threads")
		detectionTimout, _ := cmd.Flags().GetInt("interval")
		colorDisabled, _ := cmd.Flags().GetBool("dc")
		stringFilter, _ := cmd.Flags().GetString("filter")
		additionalHeader, _ := cmd.Flags().GetString("header")
		userDataFrame, _ := cmd.Flags().GetString("data")

		userDataFrame = strings.ReplaceAll(userDataFrame, "\\r", "\r")
		userDataFrame = strings.ReplaceAll(userDataFrame, "\\n", "\n")
		userDataFrame = strings.ReplaceAll(userDataFrame, "\\t", "\t")

		if method == "" {
			method = "POST"
		}

		if targetFile == "" {
			targets = append(targets, singleTarget)
		}

		if targetFile != "" {

			fileTargets, err := lib.ReadTargetFile(targetFile)

			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}

			targets = append(targets, fileTargets...)

			if len(targets) == 0 {
				fmt.Println("No targets found in file")
				return
			}

			if targets[len(targets)-1] == "" {
				targets = targets[:len(targets)-1]
			}
		}

		for _, host := range targets {

			var scanJobs []lib.ScanJob
			var wg sync.WaitGroup

			target, err := lib.NewTarget(host)
			if err != nil {
				fmt.Println("Error parsing URL:", err)
				return
			}

			fmt.Printf("[+]Processing requests...%80s\n", "\t\t")

			var headerValues []string

			if gadgetList != "" {

				headerValues, err = lib.ReadPayloadsFile(gadgetList)

				if err != nil {
					fmt.Println("Error reading file:", err)
					return
				}

				//remove trailing newline
				if headerValues[len(headerValues)-1] == "" {
					headerValues = headerValues[:len(headerValues)-1]
				}

			} else {
				headerValues = strings.Split(lib.DefaultGadgetList, "\n")
			}

			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}

			payloads, err := lib.BulkImportPayloads(headerValues, target.URL.Hostname())

			if err != nil {
				fmt.Println("Error importing payloads:", err)
				return
			}

			//for limit threads for smaller wordlists
			if routineCount > len(payloads) {
				routineCount = len(payloads)
			}

			payloadChunks := lib.ChunkPayloads(payloads, routineCount)

			for _, payloadChunk := range payloadChunks {

				conn, err := target.GetConnection()
				if err != nil {
					fmt.Println("Error connecting to target:", err)

					//prevent other threads from starting if the connection fails
					break

				}
				//defer conn.Close()
				//fmt.Println("[+]Thread Ready. HTTP2 Connected to target:", target.URL.Hostname())

				err = lib.EstablishH2Connection(conn)

				if err != nil {
					fmt.Println("Error establishing HTTP2 connection:", err)
					continue
				}
				scanJob := lib.NewScanJob(target, conn, payloadChunk)
				scanJobs = append(scanJobs, *scanJob)
			}

			fmt.Printf("[+]Starting Response Handler for: %s\n", target.URL.Hostname())

			//routine for each scanjob
			for _, scanJob := range scanJobs {

				wg.Add(1)

				go func(scanJob lib.ScanJob) {

					streamChan := make(chan string)

					defer wg.Done()
					defer scanJob.Conn.Close()

					go lib.HandleConnection(&scanJob, &streamChan)

					for _, payload := range scanJob.Payloads {

						var responseInfo string

						targetUrl := scanJob.Target.URL.Path

						queryString := scanJob.Target.URL.RawQuery
						if queryString != "" {
							targetUrl += "?" + queryString
						}

						//I REALLY need to make a builder for these
						getRequest, err := lib.GenerateRequest(scanJob.Target.URL.Hostname(), targetUrl, payload.HeaderName, payload.HeaderValue, byte(scanJob.StreamId), method, additionalHeader, userDataFrame)
						if err != nil {
							fmt.Println("Error generating request:", err)

							os.Exit(1)

						}
						err = lib.SendCustomFrame(getRequest, scanJob.Conn)
						if err != nil {
							fmt.Println("Error sending request frame:", err)
							return
						}

						select {
						case responseInfo = <-streamChan:

							fmt.Print(lib.OutputParser(payload.Name, responseInfo, colorDisabled, stringFilter))

							if responseInfo == "GOAWAY" {

								newconn, err := target.GetConnection()

								if err != nil {
									fmt.Println("Error connecting to target:", err)
									return
								}

								scanJob.Conn.Close()

								newconn, err = target.GetConnection()
								err = lib.EstablishH2Connection(newconn)
								scanJob.SetConn(newconn)

								go lib.HandleConnection(&scanJob, &streamChan)

							}

						case <-time.After(time.Duration(detectionTimout) * time.Second):

							responseInfo = "*TIMEOUT"

							scanJob.Conn.Close()

							fmt.Print(lib.OutputParser(payload.Name, responseInfo, colorDisabled, stringFilter))

							newconn, err := target.GetConnection()
							if err != nil {
								fmt.Println("Error connecting to target:", err)
								return
							}

							err = lib.EstablishH2Connection(newconn)

							scanJob.SetConn(newconn)

							go lib.HandleConnection(&scanJob, &streamChan)

							if enableConfirmation {
								scanJob.StreamId += 2

								//send a confirmation frame
								fmt.Print(lib.OutputParser("", "*Sending a confirmation request... ", colorDisabled, stringFilter))
								confirmationRequest, err := lib.GenerateRequest(scanJob.Target.URL.Hostname(), scanJob.Target.URL.Path, payload.HeaderName, payload.HeaderValue, byte(scanJob.StreamId), method, additionalHeader, "3\r\nABC\r\n0\r\n\r\n")

								if err != nil {
									fmt.Println("Error generating request:", err)
									return
								}

								err = lib.SendCustomFrame(confirmationRequest, scanJob.Conn)
								if err != nil {
									fmt.Println("Error sending request frame:", err)
									return
								}

								select {
								case responseInfo = <-streamChan:
									fmt.Print(lib.OutputParser(payload.Name, "[+] Confirmation was a success. This might indicate a vulnerable endpoint: "+responseInfo, colorDisabled, stringFilter))

									if saveSuccessfulRequests != "" {
										err = lib.WritePayloadsToFile([]lib.Payload{payload}, scanJob.Target.URL.Hostname(), saveSuccessfulRequests)
										if err != nil {
											fmt.Println("Error writing to file:", err)
										}
									}

								case <-time.After(time.Duration(detectionTimout) * time.Second):

									responseInfo = "*TIMEOUT"

									fmt.Print(lib.OutputParser(payload.Name, responseInfo, colorDisabled, stringFilter))

									newconn, err := target.GetConnection()
									if err != nil {
										fmt.Println("Error connecting to target:", err)
										return
									}

									err = lib.EstablishH2Connection(newconn)

									scanJob.SetConn(newconn)

									if err != nil {
										fmt.Println("Error establishing HTTP2 connection:", err)
										return
									}

									go lib.HandleConnection(&scanJob, &streamChan)
								}
							}
						}
						scanJob.StreamId += 2
					}

				}(scanJob)

			}
			wg.Wait()

		}

	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().BoolP("confirm", "c", false, "Enable this flag to send a confirmation to the target when a timeout is encountered. Helps confirm if the target is vulnerable.")
	scanCmd.Flags().BoolP("dc", "", false, "Disable colour in the output. This is useful when you want to save the output to a file.")
	scanCmd.Flags().StringP("url", "u", "", "The target URL to be scanned.")
	scanCmd.Flags().StringP("wordlist", "w", "", "Provide a custom list of gadgets to use. If not provided, the default list will be used.")
	scanCmd.Flags().StringP("file", "f", "", "A file containing multiple targets in url format. One target per line.")
	scanCmd.Flags().StringP("save-success", "s", "", "If a request is confirmed to be successful (via the --confirm flag), it will be saved to a file. This is useful when dealing with lots of targets.")
	scanCmd.Flags().StringP("method", "x", "POST", "The HTTP request method to be used.")
	scanCmd.Flags().StringP("header", "H", "", "Insert a custom header. It should be provided in the regular header format: \"Cookie: date=...; session=...;\"")
	scanCmd.Flags().IntP("threads", "t", 4, "The number of threads to run. Smugglefuzz can go fast, so set the desired number. However, too many may upset any WAFs.")
	scanCmd.Flags().IntP("interval", "i", 5, "The timeout interval in seconds.")
	scanCmd.Flags().StringP("filter", "", "", "Filter responses by string or frame type, etc. For example: 405, 200, 502, TIMEOUT, RST, GOAWAY, etc.")
	scanCmd.Flags().StringP("data", "d", "99\r\n", "HTTP/2 Data frame to send. eg: 99\\r\\n")
}
