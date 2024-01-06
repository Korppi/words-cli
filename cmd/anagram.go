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

// anagramCmd represents the anagram command
var anagramCmd = &cobra.Command{
	Use: "anagram",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			cmd.Help()
			os.Exit(0)
		}
		return nil
	},
	Example: "\twords anagram \"I am Lord Voldemort\" \"Tom Marvolo Riddle\"",
	Short:   "Check if 2 given words are anagrams",
	Long: `Check if 2 given words are anagrams. Text are anagrams if they contains same letters in different or same order.
Nice known example of anagrams is 'I am Lord Voldemort' which is anagram of 'Tom Marvolo Riddle'`,
	Run: func(cmd *cobra.Command, args []string) {
		result := words.IsAnagram(args[0], args[1])
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(anagramCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// anagramCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// anagramCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
