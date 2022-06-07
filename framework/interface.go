package framework

import (
	"context"
)

type (
	NodeSet   []GraphNode
	EdgeSet   []GraphEdge
	Relations []Relation
)

type Strategy interface {
	Schedule(context.Context, []Relation) ([]EdgeSet, error)
}

type Prepare interface {
	GenerateGraph(context.Context, []GraphNode, []GraphEdge) (Graph, error)
}

type Process interface {
	ProcessGraph(context.Context, Graph) (Graph, error)
}

type Decision interface {
	SelectEdges(context.Context, Graph) ([]GraphEdge, bool)
}
