package cmd

import (
	"dedup/find"
	"fmt"
	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use: "find",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			panic(err)
		}

		dupes, err := find.FindDuplicates(path)
		if err != nil {
			panic(err)
		}

		for _, dupe := range dupes {
			fmt.Println(dupe)
		}
	},
}

func init() {
	findCmd.Flags().StringP("path", "p", ".", "Path to search for duplicates")
	rootCmd.AddCommand(findCmd)
}
