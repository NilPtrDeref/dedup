package cmd

import (
	"dedup/find"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use: "delete",
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
			err = os.Remove(dupe)
			if err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	deleteCmd.Flags().StringP("path", "p", ".", "Path to delete duplicates from")
	rootCmd.AddCommand(deleteCmd)
}
