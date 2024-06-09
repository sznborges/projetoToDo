package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "to-do-project",
	Short: "To Do Project",
	Long:  "To Do Project",
}

func Execute() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error: ", r)
		}
	}()
	err := rootCmd.Execute()
	if err != nil {
		panic("error while executing root command")
	}
}
