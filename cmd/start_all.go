package cmd

import (
	"github.com/spf13/cobra"
)

var startAllCmd = &cobra.Command{
	Use:   "start-all",
	Short: "Starts all services",
	Long:  `Starts all services.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		zapLogger := initServerConfig()
		logger.Info("Starting all services")
		go setupBackendServer(zapLogger, host, backendPort)
		return setupFrontendServer(zapLogger, host, frontendPort, basepath)
	},
}

func init() {
	RootCmd.AddCommand(startAllCmd)
}
