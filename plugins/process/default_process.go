package process

import (
	"context"
	"fmt"

	"github.com/BinacsLee/escheduler/framework"
	"github.com/BinacsLee/escheduler/util/names"
	"github.com/BinacsLee/escheduler/util/tarjan"
)

const (
	Name = names.DefaultProcess
)

type DefaultProcess struct {
}

var (
	_ framework.Process = &DefaultProcess{}
)

func (p *DefaultProcess) AfterInject() error {
	return nil
}

func (p *DefaultProcess) Name() string {
	return Name
}

func (p *DefaultProcess) ProcessGraph(ctx context.Context, g framework.Graph) (framework.Graph, error) {
	// TODO: Tarjan SCC and rebuild graph.
	scc := tarjan.SCCTarjan(g)
	fmt.Printf("SCC = %v\n", scc)
	return nil, nil
}
