package strategy

import (
	"context"

	"github.com/BinacsLee/escheduler/framework"
	"github.com/BinacsLee/escheduler/testing"
	"github.com/BinacsLee/escheduler/util/hashid"
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

func (s *DefaultStrategy) Schedule(ctx context.Context, relations []framework.Relation) ([]framework.EdgeSet, error) {
	s.Logger.Info("Schedule start", "relations", len(relations))
	nodes := make([]framework.GraphNode, 0)
	edges := make([]framework.GraphEdge, 0)
	{
		nodesID, edgesID := hashid.NewHashID(), hashid.NewHashID()
		for _, relation := range relations {
			nodeNameFrom, nodeNameTo, podName := relation.GetFrom(), relation.GetTo(), relation.GetName()
			if nodeNameFrom == nodeNameTo {
				// ATTENTION
				continue
			}
			existFrom, existTo := nodesID.Exist(nodeNameFrom), nodesID.Exist(nodeNameTo)
			nodeID1, nodeID2, podID := nodesID.Get(nodeNameFrom), nodesID.Get(nodeNameTo), edgesID.Get(podName)
			// TODO: use framework.Structure instead of testing.
			if !existFrom {
				nodes = append(nodes, testing.MakeGraphNode().Name(nodeNameFrom).ID(nodeID1).Obj())
			}
			if !existTo {
				nodes = append(nodes, testing.MakeGraphNode().Name(nodeNameTo).ID(nodeID2).Obj())
			}
			edges = append(edges, testing.MakeGraphEdge().Name(podName).ID(podID).From(nodeID1).To(nodeID2).Obj())
		}
	}
	s.Logger.Info("Schedule ready", "nodes", len(nodes), "edges", len(edges))
	// TODO: support DEBUG-MODE
	for i := range nodes {
		s.Logger.Info("DEBUG", "node", nodes[i])
	}
	for i := range edges {
		s.Logger.Info("DEBUG", "edge", edges[i])
	}

	g, err := s.Prepare.GenerateGraph(ctx, nodes, edges)
	if err != nil {
		return nil, err
	}
	g, err = s.Process.ProcessGraph(ctx, g)
	if err != nil {
		return nil, err
	}
	return nil, nil
	edgeSet := make([]framework.EdgeSet, 0)
	for {
		es, empty := s.Decision.SelectEdges(ctx, g)
		if empty {
			return edgeSet, nil
		}
		log.Info("DefaultStrategy.Schedule", "SelectEdges", es)
		edgeSet = append(edgeSet, es)
	}
}
