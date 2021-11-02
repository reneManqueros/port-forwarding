package cmd

import (
	"github.com/spf13/cobra"
	"port-forwarding/models"
)

var forwardCmd = &cobra.Command{
	Use:   "forward",
	Short: "Run forward service",
	Run: func(cmd *cobra.Command, args []string) {
		var settings models.Settings
		settings.Load()

		for key, _ := range settings.Redirections {
			thisRedir := settings.Redirections[key]
			go func(r *models.Redirection) {
				r.Listen()
			}(&thisRedir)
		}
		waitChan := make(chan int)
		<-waitChan
	},
}

func init() {
	rootCmd.AddCommand(forwardCmd)
	forwardCmd.Flags()
}
