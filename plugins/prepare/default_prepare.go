package prepare

import (
	"context"

	"github.com/BinacsLee/escheduler/framework"
	"github.com/BinacsLee/escheduler/util/names"
)

const (
	Name = names.DefaultPrepare
)

type DefaultPrepare struct {
}

var (
	_ framework.Prepare = &DefaultPrepare{}
)

func (p *DefaultPrepare) AfterInject() error {
	return nil
}

func (p *DefaultPrepare) Name() string {
	return Name
}

func (p *DefaultPrepare) GenerateGraph(ctx context.Context, nodes []framework.GraphNode, edges []framework.GraphEdge) (framework.Graph, error) {
	// TODO: improve this int64
	g := framework.NewGraph(int64(len(nodes)), int64(len(edges)))
	for i := range edges {
		e := edges[i]
		g.AddEdge(e.GetFrom(), e.GetTo(), e.GetID())
	}
	return g, nil
}
