package cmd

import (
	"github.com/spf13/cobra/cobra"
)

var RootCmd = &cobra.Command{
	Use: "task",
	Short: "Task is a CLI task manager",
}