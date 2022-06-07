package decision

import (
	"context"

	"github.com/BinacsLee/escheduler/framework"
	"github.com/BinacsLee/escheduler/util/names"
)

const (
	Name = names.DefaultDecision
)

type DefaultDecision struct {
}

var (
	_ framework.Decision = &DefaultDecision{}
)

func (d *DefaultDecision) AfterInject() error {
	return nil
}

func (d *DefaultDecision) Name() string {
	return Name
}

func (d *DefaultDecision) SelectEdges(ctx context.Context, g framework.Graph) ([]framework.GraphEdge, bool) {
	return nil, false
}
