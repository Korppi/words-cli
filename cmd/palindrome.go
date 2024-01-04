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

var CaseSensitiveEnabled bool

// palindromeCmd represents the palindrome command
var palindromeCmd = &cobra.Command{
	Use: "palindrome",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			cmd.Help()
			os.Exit(0)
		}
		return nil
	},
	Short: "Return true, if text is palindrome",
	Long: `Return true, if text is palindrome. 
Text is palindrome if it is same text after reversing text.

NOTE: By default this is case-insensitive meaning text 'Aa' will be considered as palindrome.
Use flag -c or --case-sensitive to change to case-sensitive. 
	
Example of palindrome: rotator`,
	Run: func(cmd *cobra.Command, args []string) {
		var result bool
		if CaseSensitiveEnabled {
			result = words.IsPalindromeCaseSensitive(strings.Join(args[:], " "))
		} else {
			result = words.IsPalindrome(strings.Join(args[:], " "))
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(palindromeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// palindromeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// palindromeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	palindromeCmd.Flags().BoolVarP(&CaseSensitiveEnabled, "case-sensitive", "c", false, "Run functions as case-sensitive.")
}
