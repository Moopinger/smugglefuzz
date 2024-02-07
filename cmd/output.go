/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// outputCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// outputCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
