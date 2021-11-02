package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"port-forwarding/models"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all port forwards",
	RunE: func(cmd *cobra.Command, args []string) error {
		var settings models.Settings
		settings.Load()
		for _, redirection := range settings.Redirections {
			log.Printf(`Type: %s, From: %s, To: %s`, redirection.Network, redirection.Source, redirection.Destination)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags()
}
