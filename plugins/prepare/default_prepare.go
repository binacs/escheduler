package prepare

import (
	"context"
	"fmt"

	"github.com/BinacsLee/escheduler/framework"
	"github.com/BinacsLee/escheduler/internal/tarjan"
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

func (p *DefaultPrepare) GenerateGraph(ctx context.Context, nodes []framework.Node, edges []framework.Edge) (framework.Graph, error) {
	g := framework.NewGraph(nodes, edges)
	// TODO: Tarjan SCC and rebuild graph.
	scc := tarjan.SCCTarjan(g)
	count := scc.GetSccCnt()
	sccNodes := make([]framework.Node, count)
	for i := int64(0); i < count; i++ {
		// Make a SccNode instead of a Node
		name := fmt.Sprintf("SccSccNode-%v", i)
		sccNodes[i] = framework.MakeSccNode().Name(name).ID(int64(i)).Obj()
	}
	sccEdges := make([]framework.Edge, 0)
	for i := int64(0); i < g.NumNode(); i++ {
		sccNodes[scc.GetID(i)].(framework.SccNode).AddNodes(g.GetNode(i))
		for j := g.Head(i); j != -1; j = g.Next(j) {
			k := g.Edge(j)
			a, b := scc.GetID(i), scc.GetID(k)
			if a == b {
				sccNodes[scc.GetID(i)].(framework.SccNode).AddEdges(g.GetEdge(j))
			} else {
				// ATTENTION: make a reverse graph (b->a).
				nodeEdge := g.GetEdge(j)
				name := fmt.Sprintf("SccEdge-%v", j)
				sccEdge := framework.MakeSccEdge().Name(name).ID(j).From(b).To(a).Edge(nodeEdge).Obj()
				sccEdges = append(sccEdges, sccEdge)
			}
		}
	}
	// TODO: Check nodeSet and edges in advance.

	// Return a DAG.
	return framework.NewGraph(sccNodes, sccEdges), nil
}
