package testing

import "github.com/BinacsLee/escheduler/framework"

type GraphNodeWrapper struct{ framework.GraphNode }

func MakeGraphNode() *GraphNodeWrapper {
	return &GraphNodeWrapper{&framework.GraphNodeImpl{}}
}

func (gn *GraphNodeWrapper) Obj() framework.GraphNode {
	return gn.GraphNode
}

func (gn *GraphNodeWrapper) Name(name string) *GraphNodeWrapper {
	gn.GraphNode.SetName(name)
	return gn
}

func (gn *GraphNodeWrapper) ID(id int64) *GraphNodeWrapper {
	gn.GraphNode.SetID(id)
	return gn
}

type GraphEdgeWrapper struct{ framework.GraphEdge }

func MakeGraphEdge() *GraphEdgeWrapper {
	return &GraphEdgeWrapper{&framework.GraphEdgeImpl{}}
}

func (ge *GraphEdgeWrapper) Obj() framework.GraphEdge {
	return ge.GraphEdge
}

func (ge *GraphEdgeWrapper) Name(name string) *GraphEdgeWrapper {
	ge.GraphEdge.SetName(name)
	return ge
}

func (ge *GraphEdgeWrapper) ID(id int64) *GraphEdgeWrapper {
	ge.GraphEdge.SetID(id)
	return ge
}

func (ge *GraphEdgeWrapper) From(from int64) *GraphEdgeWrapper {
	ge.GraphEdge.SetFrom(from)
	return ge
}

func (ge *GraphEdgeWrapper) To(to int64) *GraphEdgeWrapper {
	ge.GraphEdge.SetTo(to)
	return ge
}

type RelationWrapper struct{ framework.Relation }

func MakeRelatione() *RelationWrapper {
	return &RelationWrapper{&framework.RelationImpl{}}
}

func (ge *RelationWrapper) Obj() framework.Relation {
	return ge.Relation
}

func (ge *RelationWrapper) Name(name string) *RelationWrapper {
	ge.Relation.SetName(name)
	return ge
}

func (ge *RelationWrapper) From(from string) *RelationWrapper {
	ge.Relation.SetFrom(from)
	return ge
}

func (ge *RelationWrapper) To(to string) *RelationWrapper {
	ge.Relation.SetTo(to)
	return ge
}
