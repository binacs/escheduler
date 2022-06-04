package commands

import (
	"github.com/spf13/cobra"

	"github.com/binacsgo/inject"

	"github.com/BinacsLee/escheduler/core"
)

func init() {
	StartCmd.PersistentFlags().StringVar(&configFile, "configFile", "config.toml", "config file (default is ./config.toml)")
}

var (
	// StartCmd the start command
	StartCmd = &cobra.Command{
		Use:   "start",
		Short: "Start Command",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			sched, err := initService()
			if err != nil {
				return err
			}
			sched.OnStart()
			return nil
		},
	}
)

func initService() (core.EScheduler, error) {
	sched := &core.ESchedulerImpl{}

	inject.Regist("Logger", logger)
	inject.Regist("EScheduler", sched)

	inject.DoInject()

	return sched, nil
}
