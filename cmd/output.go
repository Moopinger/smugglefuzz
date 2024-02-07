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
	Short: "Output the default smuggle gadgets",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%v\n", lib.DefaultGadgetList)
	},
}

func init() {
	rootCmd.AddCommand(outputCmd)

}
