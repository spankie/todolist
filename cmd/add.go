package cmd

import (
	"github.com/spf13/cobra"
)

var Add = &cobra.Command{
	Use:     "add [Task name]",
	Short:   "Add task to the list",
	Example: `todolist add "take zalando test"`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// save the task to the list
		task := Task{
			Name: args[0],
			Done: false,
		}
		persistTaskToTaskStore(task)
	},
}
