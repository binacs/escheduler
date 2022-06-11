package process

import (
	"context"
	"fmt"

	"github.com/BinacsLee/escheduler/framework"
	"github.com/BinacsLee/escheduler/internal/topology"
	"github.com/BinacsLee/escheduler/util/names"
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

func (p *DefaultProcess) ProcessGraph(ctx context.Context, g framework.Graph) (framework.DepthChart, error) {
	topo, isTopo := topology.Topology(g)
	if !isTopo {
		return nil, fmt.Errorf("not topology graph")
	}

	depth := make([]int64, g.NumNode())
	for i := g.NumNode() - 1; i >= 0; i-- {
		u := topo[i]
		for j := g.Head(u); j != -1; j = g.Next(j) {
			k := g.Edge(j)
			if depth[u] < depth[k]+1 {
				depth[u] = depth[k] + 1
			}
		}
	}

	depthChart := framework.NewDepthChart()
	for i := int64(0); i < g.NumNode(); i++ {
		u := topo[i]
		depthChart.AddNode(g.GetNode(u), depth[u])
	}
	return depthChart, nil
}
