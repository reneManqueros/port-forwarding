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
		for _, portForward := range settings.PortForwards {
			log.Printf(`Type: %s, From: %s, To: %s`, portForward.Network, portForward.Source, portForward.Destination)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags()
}
