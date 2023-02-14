package framework

import (
	"context"
	"fmt"

	"github.com/binacs/escheduler/util/set"
)

type (
	NodeSet   set.Set
	EdgeSet   set.Set
	Relations []Relation
	CheckFunc func(EdgeSet, ...Edge) bool
)

func (s Relations) String() string {
	var str string
	for i := range s {
		str += fmt.Sprintf("{%v},", s[i])
	}
	return str
}

type Strategy interface {
	Schedule(context.Context, []Relation) ([]Relations, error)
}

type Prepare interface {
	GenerateGraph(context.Context, []Node, []Edge) (Graph, error)
}

type Process interface {
	ProcessGraph(context.Context, Graph) (DepthChart, error)
}

type Decision interface {
	SelectEdges(context.Context, Graph, DepthChart, CheckFunc) (EdgeSet, bool)
}
