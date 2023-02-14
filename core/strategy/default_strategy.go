package strategy

import (
	"context"
	"fmt"

	"github.com/binacs/escheduler/framework"
	"github.com/binacs/escheduler/util/hashid"
	"github.com/binacsgo/log"
)

type DefaultStrategy struct {
	Logger   log.Logger         `inject-name:"Logger"`
	Prepare  framework.Prepare  `inject-name:"PluginDefaultPrepare"`
	Process  framework.Process  `inject-name:"PluginDefaultProcess"`
	Decision framework.Decision `inject-name:"PluginDefaultDecision"`
}

var (
	_ framework.Strategy = &DefaultStrategy{}
)

func (s *DefaultStrategy) AfterInject() error {
	return nil
}

func (s *DefaultStrategy) Schedule(ctx context.Context, relations []framework.Relation) ([]framework.Relations, error) {
	s.Logger.Info("Schedule start", "relations", len(relations))
	nodes, edges := make([]framework.Node, 0), make([]framework.Edge, 0)
	nodesID, edgesID := hashid.NewHashID(), hashid.NewHashID()
	{
		// TODO: make the following progress parallel.
		// nodeDSU := dsu.NewDSU(int64(2 * len(relations)))
		for _, relation := range relations {
			nodeNameFrom, nodeNameTo, podName := relation.GetFrom(), relation.GetTo(), relation.GetName()
			if nodeNameFrom == nodeNameTo {
				// ATTENTION: Skip this situation.
				continue
			}
			s.Logger.Info("Schedule parse relation", "from", nodeNameFrom, "to", nodeNameTo, "podName", podName)
			existFrom, existTo := nodesID.Exist(nodeNameFrom), nodesID.Exist(nodeNameTo)
			nodeFromID, nodeToID, podID := nodesID.Get(nodeNameFrom), nodesID.Get(nodeNameTo), edgesID.Get(podName)
			if !existFrom {
				node := framework.MakeNode().Name(nodeNameFrom).ID(nodeFromID).Obj()
				nodes = append(nodes, node)
				s.Logger.Debug("Schedule add node", "node", node)
			}
			if !existTo {
				node := framework.MakeNode().Name(nodeNameTo).ID(nodeToID).Obj()
				nodes = append(nodes, node)
				s.Logger.Debug("Schedule add node", "node", node)
			}
			edge := framework.MakeEdge().Name(podName).ID(podID).From(nodeFromID).To(nodeToID).Obj()
			edges = append(edges, edge)
			s.Logger.Debug("Schedule add edge", "edge", edge)
			// nodeDSU.Merge(nodeFromID, nodeToID)
		}
		// TODO: Partition ...
		// nodeDSU.Squash()
	}
	s.Logger.Info("Schedule ready", "nodes", len(nodes), "edges", len(edges))
	for i := range nodes {
		s.Logger.Debug("Schedule DataSet", "node", nodes[i])
	}
	for i := range edges {
		s.Logger.Debug("Schedule DataSet", "edge", edges[i])
	}

	g, err := s.Prepare.GenerateGraph(ctx, nodes, edges)
	if err != nil {
		return nil, err
	}
	chart, err := s.Process.ProcessGraph(ctx, g)
	if err != nil {
		return nil, err
	}

	relationSet := make([]framework.Relations, 0)
	for {
		es, empty := s.Decision.SelectEdges(ctx, g, chart,
			// Test CheckFunc, can not select more than four edges.
			func(es framework.EdgeSet, e ...framework.Edge) bool {
				if es.Len()+len(e) > 4 {
					return false
				}
				return true
			})
		if empty {
			if chart.Len() > 0 {
				return nil, fmt.Errorf("invalid graph, cann't pass the CheckFunc")
			}
			return relationSet, nil
		}
		relations := make(framework.Relations, 0)
		for _, obj := range es.List() {
			edge := obj.(framework.Edge)
			from, to, name := nodesID.Lookup(edge.GetFrom()), nodesID.Lookup(edge.GetTo()), edgesID.Lookup(edge.GetID())
			relations = append(relations, framework.MakeRelatione().From(from).To(to).Name(name).Obj())
		}
		s.Logger.Info("DefaultStrategy.Schedule", "SelectRelations", relations)
		relationSet = append(relationSet, relations)
	}
}
