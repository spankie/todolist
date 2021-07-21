package cmd

import (
	"github.com/spf13/cobra"
)

var Cleanup = &cobra.Command{
	Use:   "cleanup",
	Short: "Clean up done tasks",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		readAllTasksFromFileByDone(DataFile, false)
		persistTaskStore()
	},
}
