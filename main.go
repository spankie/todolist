package main

import (
	"log"

	"github.com/spankie/todolist/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "todolist"} // cmd.AppName

	rootCmd.AddCommand(cmd.Add)
	rootCmd.AddCommand(cmd.Done)
	rootCmd.AddCommand(cmd.UnDone)
	rootCmd.AddCommand(cmd.List)
	rootCmd.AddCommand(cmd.Cleanup)

	// save the database to file and close the data file
	defer func() {
		err := cmd.DataFile.Close()
		if err != nil {
			log.Fatal("could not close the database file properly")
		}
	}()

	rootCmd.Execute() // Don't change this
}
