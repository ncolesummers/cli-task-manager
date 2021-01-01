package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		for i, arg := range args {
			fmt.Println(i, ":", arg)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
