package core

import (
	"github.com/binacsgo/log"
)

type ESchedulerImpl struct {
	Logger log.Logger `inject-name:"Logger"`
}

func (es *ESchedulerImpl) OnStart() {
	es.Logger.Info("run")
}
