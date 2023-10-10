package cmd

import (
	"github.com/bersen66/cbrapp/pkg/ntp"
	"github.com/spf13/cobra"
)

const (
	defaultNtpServer = "0.beevik-ntp.pool.ntp.org"
)

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Get time using ntp",
	Long:  `Initiating ntp query to server with aim to get current time`,
	Run: func(cmd *cobra.Command, args []string) {
		server, _ := cmd.Flags().GetString("server")
		ntp.RunClient(server)
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)
	timeCmd.PersistentFlags().String("server", defaultNtpServer, "specify ntp-server")
}
