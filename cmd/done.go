package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var Done = &cobra.Command{
	Use:   "done [Task id]",
	Short: "Mark a task as done",
	// Example: AppName + " done 1",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// get the id as integer
		ID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("invalid id provided: %s\n", args[0])
		}
		changeTaskDone(true, ID)
	},
}
