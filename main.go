package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var (
		original string // The orginal DB structure (path to sql file)
		current  string // The current DB structure (path to sql file)
	)

	var cmdDiff = &cobra.Command{
		Use:   "diff [string to print]",
		Short: "Create the diff and print the changes to screen",
		Run: func(cmd *cobra.Command, args []string) {
			diff(original, current)
		},
	}

	cmdDiff.Flags().StringVarP(&original, "original", "o", "original.sql", "The path to the original DB structure (sql file)")
	cmdDiff.Flags().StringVarP(&current, "current", "c", "current.sql", "The path to the current DB structure (sql file)")

	var rootCmd = &cobra.Command{Use: "dbdiff"}
	rootCmd.AddCommand(cmdDiff)
	rootCmd.Execute()
}
