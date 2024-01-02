/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Korppi/words"
	"github.com/spf13/cobra"
)

// reverseCmd represents the reverse command
var reverseCmd = &cobra.Command{
	Use: "reverse",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			cmd.Help()
			os.Exit(0)
		}
		return nil
	},
	Short: "Reverse text",
	Long: `Reverse given text.
	
Example input: words reverse abc
output: cba`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(words.Reverse(strings.Join(args[:], " ")))
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reverseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reverseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
