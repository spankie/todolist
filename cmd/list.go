package cmd

import (
	"github.com/spf13/cobra"
)

var List = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks still to do",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		listTasksByDone(false)
	},
}
