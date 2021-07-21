package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var UnDone = &cobra.Command{
	Use:   "undone [Task id]",
	Short: "Mark a task as undone",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// get the id as integer
		ID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("invalid id provided: %s\n", args[0])
		}
		changeTaskDone(false, ID)
	},
}
