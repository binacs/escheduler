package commands

import (
	"github.com/binacsgo/log"
	"github.com/spf13/cobra"
)

var (
	configFile string
	logger     log.Logger

	RootCmd = &cobra.Command{
		Use:   "root",
		Short: "Root Command",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logger = log.Sugar()
			logger.Info("Init finished")
		},
	}
)
