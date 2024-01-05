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

// lengthCmd represents the length command
var lengthCmd = &cobra.Command{
	Use:   "length",
	Short: "Return lenght of given text",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			cmd.Help()
			os.Exit(0)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		result := words.WordLength(strings.Join(args[:], " "))
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(lengthCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lengthCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lengthCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
