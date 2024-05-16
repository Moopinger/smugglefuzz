/*
2024 Moopinger
*/
package cmd

import (
	"fmt"

	"github.com/moopinger/smugglefuzz/lib"
	"github.com/spf13/cobra"
)

// outputCmd represents the output command
var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Output the default smuggle gadgets. By default, the list is not exhaustive. Use the -e flag to see the extended list.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		extended, _ := cmd.Flags().GetBool("extended")

		if extended {
			fmt.Printf("%v\n", lib.ExtendedGadgetList)
		} else {
			fmt.Printf("%v\n", lib.DefaultGadgetList)
		}
	},
}

func init() {
	rootCmd.AddCommand(outputCmd)

	// Here you will define your flags and configuration settings.
	outputCmd.Flags().BoolP("extended", "e", false, "Use the extended provided wordlist.")

}
