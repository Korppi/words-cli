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

// pangramCmd represents the pangram command
var pangramCmd = &cobra.Command{
	Use: "pangram",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			cmd.Help()
			os.Exit(0)
		}
		return nil
	},
	Short: "Check if text is pangram",
	Long:  `Check if text is pangram. Text is pangram if it contains all letters from alphabe atleast once`,
	Run: func(cmd *cobra.Command, args []string) {
		result := words.IsPangram(strings.Join(args[:], " "))
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(pangramCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pangramCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pangramCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
