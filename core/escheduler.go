package core

import (
	"context"

	"github.com/BinacsLee/escheduler/core/testdata"
	"github.com/BinacsLee/escheduler/framework"
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
	}
}

func (es *ESchedulerImpl) run() bool {
	testData := testdata.SelectRandomTestData()
	log.Info("Got TestData", "testData", testData)
	es.DefaultStrategy.Schedule(context.TODO(), testData.Relations)
	return true
}
