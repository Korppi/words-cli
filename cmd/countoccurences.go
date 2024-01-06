/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Korppi/words"
	"github.com/spf13/cobra"
)

// countoccurencesCmd represents the countoccurences command
var countoccurencesCmd = &cobra.Command{
	Use:   "countoccurences",
	Short: "Count occurences of text a in text b",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			cmd.Help()
			os.Exit(0)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		result := words.CountOccurences(args[0], args[1])
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(countoccurencesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// countoccurencesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// countoccurencesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
