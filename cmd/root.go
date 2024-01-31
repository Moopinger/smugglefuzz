/*
2024 Moopinger
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "smugglefuzz",
	Short: "Tool for identifying HTTP Smuggling vulnerabilities that arise when downgrading to HTTP/1.",
	Long:  `Tool for identifying HTTP Smuggling vulnerabilities that arise when downgrading to HTTP/1.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {

		os.Exit(1)
	}
}

func init() {

	rootCmd.SetHelpCommand(&cobra.Command{Use: "no-help", Run: func(cmd *cobra.Command, args []string) {}})

	rootCmd.CompletionOptions.DisableDefaultCmd = true

}
