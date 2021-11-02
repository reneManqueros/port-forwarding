package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"log"
	"port-forwarding/models"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "add network[tcp|udp] from_address:from_port to_address:to_port",
	Short: "Add a port forward",
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
		settings.PortForwards = append(settings.PortForwards, r)
		settings.Save()
		log.Println("port forward added")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags()
}
