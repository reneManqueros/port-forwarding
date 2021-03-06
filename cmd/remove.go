package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"log"
	"port-forwarding/models"
	"strings"
)

var removeCmd = &cobra.Command{
	Use:   "remove network[tcp|udp] from_address:from_port to_address:to_port",
	Short: "Remove a port forward",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 {
			return errors.New("not enough arguments")
		}
		network := strings.ToLower(args[0])
		if network != "tcp" && network != "udp" {
			return errors.New("network must be tcp or udp")
		}

		// ToDo: Validate from|to addresses

		var settings models.Settings
		settings.Load()
		r := models.PortForward{
			Source:      args[1],
			Destination: args[2],
			Network:     network,
		}

		var filteredPortForwards []models.PortForward

		for _, value := range settings.PortForwards {
			if value != r {
				filteredPortForwards = append(filteredPortForwards, value)
			}
		}

		settings.PortForwards = filteredPortForwards
		settings.Save()
		log.Println("port forward removed")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags()
}
