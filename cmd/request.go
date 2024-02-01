/*
2024 Moopinger
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/moopinger/smugglefuzz/lib"
	"github.com/spf13/cobra"
)

var requestCmd = &cobra.Command{
	Use:   "request",
	Short: "Submit a single request to the target server.",
	Long:  `Discovered a potential HTTP Smuggling vulnerability? Tamper the request parameters here to confirm the vulnerability.`,

	Example: `	smugglefuzz request -u https://www.example.com/ -a "content-length\t; 13"
	smugglefuzz request -u https://www.example.com/ -a "content-length\t; 13" -x PUT
	smugglefuzz request -u https://www.example.com/ -a "content-length\t; 13" -H "Cookie: date=...; session=...;" `,

	Run: func(cmd *cobra.Command, args []string) {

		if cmd.Flag("url").Value.String() == "" {
			fmt.Println("You must provide a target URL (-u)")
			cmd.Help()
			return
		}

		if cmd.Flag("attack").Value.String() == "" {
			fmt.Println("You must provide an attack header (-a)")
			cmd.Help()
			return
		}

		fmt.Print(lib.Banner)

		var scanJob lib.ScanJob

		//print the method flag
		//will make the switch to Flags().StringVARP at some point
		method, _ := cmd.Flags().GetString("method")
		singleTarget, _ := cmd.Flags().GetString("url")
		attackHeader, _ := cmd.Flags().GetString("attack")

		detectionTimout, _ := cmd.Flags().GetInt("interval")
		colorDisabled, _ := cmd.Flags().GetBool("dc")
		additionalHeader, _ := cmd.Flags().GetString("header")

		if method == "" {
			method = "POST"
		}

		target, err := lib.NewTarget(singleTarget)
		if err != nil {
			fmt.Println("Error parsing URL:", err)
			return
		}

		attackHeaderPayload, err := lib.ImportSinglePayloads(attackHeader)

		if err != nil {
			fmt.Println("Error parsing attack header:", err)
			return
		}

		connection, err := target.GetConnection()

		if err != nil {
			fmt.Println("Error getting connection:", err)
			return
		}

		err = lib.EstablishH2Connection(connection)

		if err != nil {
			fmt.Println("Error establishing connection:", err)
			return
		}

		scanJob = *lib.NewScanJob(target, connection, []lib.Payload{attackHeaderPayload})

		fmt.Printf("[+]Starting Response Handler for: %s\n", target.URL.Hostname())

		streamChan := make(chan string)

		go lib.HandleConnection(&scanJob, &streamChan)

		var responseInfo string

		targetUrl := scanJob.Target.URL.Path

		queryString := scanJob.Target.URL.RawQuery
		if queryString != "" {
			targetUrl += "?" + queryString
		}

		getRequest, err := lib.GenerateRequest(scanJob.Target.URL.Hostname(), targetUrl, attackHeaderPayload.HeaderName, attackHeaderPayload.HeaderValue, byte(scanJob.StreamId), method, false, additionalHeader, scanJob.Target.H2CEnabled)
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

			fmt.Print(lib.OutputParser(attackHeaderPayload.Name, responseInfo, colorDisabled, ""))

		case <-time.After(time.Duration(detectionTimout) * time.Second):

			responseInfo = "*TIMEOUT"

			fmt.Print(lib.OutputParser(attackHeaderPayload.Name, responseInfo, colorDisabled, ""))

		}

		scanJob.Conn.Close()

	},
}

func init() {
	rootCmd.AddCommand(requestCmd)

	requestCmd.Flags().StringP("url", "u", "", "The target URL to submit the request to.")
	requestCmd.Flags().BoolP("dc", "", false, "Disable colour in the output.")
	requestCmd.Flags().StringP("method", "x", "POST", "The method to use.")
	requestCmd.Flags().StringP("header", "H", "", "Insert custom header. eg \"Cookie: values\"")
	requestCmd.Flags().IntP("interval", "i", 5, "Detection timeout interval in seconds.")
	requestCmd.Flags().StringP("attack", "a", "", "Attack Header, separated by (; ) like the wordlist in 'scan' mode.")
}
