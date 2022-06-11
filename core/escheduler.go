package core

import (
	"context"
	"time"

	"github.com/BinacsLee/escheduler/framework"
	"github.com/BinacsLee/escheduler/testing/testdata"
	"github.com/binacsgo/log"
)

type ESchedulerImpl struct {
	Logger          log.Logger         `inject-name:"Logger"`
	DefaultStrategy framework.Strategy `inject-name:"DefaultStrategy"`
}

var (
	_ EScheduler = &ESchedulerImpl{}
)

func (es *ESchedulerImpl) Run() error {
	es.Logger.Info("EScheduler Start")
	defer es.Logger.Info("EScheduler End")
	for {
		if quit := es.run(); quit {
			return nil
		}
		time.Sleep(5 * time.Second)
	}
}

func (es *ESchedulerImpl) run() bool {
	testData := testdata.SelectRandomTestData()
	results, err := es.DefaultStrategy.Schedule(context.TODO(), testData.Relations)
	if err != nil {
		es.Logger.Info("Schedule Failed", "err", err)
		return false
	}
	es.Logger.Info("Schedule Success", "results", results)
	return false
}
