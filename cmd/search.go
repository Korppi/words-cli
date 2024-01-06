/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/Korppi/words"
	"github.com/spf13/cobra"
)

var (
	shortest int
	longest  int
	format   string
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use: "search",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			cmd.Help()
			os.Exit(0)
		}
		return nil
	},
	Short: "Search words for given letters",
	Long: `Search words for given letters and optionally given size constraints and/or format.
Example (search words, that can be written with letters 'act', has min and max size of 3, and has letter 'a' in the middle. Returns: cat):
  words search act -s 3 -l 3 -f .a.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if format != "" && (len(format) < shortest || len(format) > longest) {
			return errors.New("format is not matching with shortest or longest")
		}
		result := words.SearchPossibleWords(args[0])
		if len(result) > 0 {
			for _, word := range result {
				r, _ := regexp.Compile(format)
				if len(word) >= shortest && len(word) <= longest && r.MatchString(word) {
					fmt.Println(word)
				}
			}
		} else {
			fmt.Println("No words found.")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	searchCmd.Flags().IntVarP(&shortest, "shortest", "s", 0, "Minimum lenght of word")
	searchCmd.Flags().IntVarP(&longest, "longest", "l", 100, "Maximum lenght of word")
	searchCmd.Flags().StringVarP(&format, "format", "f", "", "Format for words")
}
