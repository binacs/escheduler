package commands

import (
	"github.com/spf13/cobra"

	"github.com/binacsgo/inject"

	"github.com/BinacsLee/escheduler/core"
	"github.com/BinacsLee/escheduler/core/strategy"
	"github.com/BinacsLee/escheduler/plugins/decision"
	"github.com/BinacsLee/escheduler/plugins/prepare"
	"github.com/BinacsLee/escheduler/plugins/process"
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
			return sched.Run()
		},
	}
)

func initService() (core.EScheduler, error) {
	sched := &core.ESchedulerImpl{}

	inject.Regist("Logger", logger)
	inject.Regist("EScheduler", sched)

	// Plugins
	inject.Regist("PluginDefaultPrepare", &prepare.DefaultPrepare{})
	inject.Regist("PluginDefaultProcess", &process.DefaultProcess{})
	inject.Regist("PluginDefaultDecision", &decision.DefaultDecision{})

	// Strategys
	inject.Regist("DefaultStrategy", &strategy.DefaultStrategy{})

	inject.DoInject()
	return sched, nil
}
