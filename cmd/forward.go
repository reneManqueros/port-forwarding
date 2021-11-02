package cmd

import (
	"github.com/spf13/cobra"
	"port-forwarding/models"
)

var forwardCmd = &cobra.Command{
	Use:   "forward",
	Short: "Run port forwarding",
	Run: func(cmd *cobra.Command, args []string) {
		var settings models.Settings
		settings.Load()

		for _, portForward := range settings.PortForwards {
			go func(pf models.PortForward) {
				pf.Listen()
			}(portForward)
		}
		waitChan := make(chan int)
		<-waitChan
	},
}

func init() {
	rootCmd.AddCommand(forwardCmd)
	forwardCmd.Flags()
}
