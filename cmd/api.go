package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sznborges/projetoToDo/infra/server"
)

var (
	apiCommand = &cobra.Command{
		Use:  "api",
		Short: "Start the API server",
		Long:  "Start the running API server",
		RunE:  apiExecute,
	}
)

func init() {
	rootCmd.AddCommand(apiCommand)
}

func apiExecute(cmd *cobra.Command, args []string) error {
	server.StartHTTP()
	return nil
}