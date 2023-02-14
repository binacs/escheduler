package framework

import "github.com/binacs/escheduler/util/set"

type NodeWrapper struct{ Node }

func MakeNode() *NodeWrapper {
	return &NodeWrapper{&NodeImpl{}}
}

func (gn *NodeWrapper) Obj() Node {
	return gn.Node
}

func (gn *NodeWrapper) Name(name string) *NodeWrapper {
	gn.Node.SetName(name)
	return gn
}

func (gn *NodeWrapper) ID(id int64) *NodeWrapper {
	gn.Node.SetID(id)
	return gn
}

type SccNodeWrapper struct{ SccNode }

func MakeSccNode() *SccNodeWrapper {
	return &SccNodeWrapper{&SccNodeImpl{selfCircle: set.NewSet(), nodeSet: set.NewSet()}}
}

func (gns *SccNodeWrapper) Obj() SccNode {
	return gns.SccNode
}

func (gns *SccNodeWrapper) Name(name string) *SccNodeWrapper {
	gns.SccNode.SetName(name)
	return gns
}

func (gns *SccNodeWrapper) ID(id int64) *SccNodeWrapper {
	gns.SccNode.SetID(id)
	return gns
}

func (gns *SccNodeWrapper) Edge(selfCircle []Edge) *SccNodeWrapper {
	gns.SccNode.AddEdges(selfCircle...)
	return gns
}

type EdgeWrapper struct{ Edge }

func MakeEdge() *EdgeWrapper {
	return &EdgeWrapper{&EdgeImpl{}}
}

func (ge *EdgeWrapper) Obj() Edge {
	return ge.Edge
}

func (ge *EdgeWrapper) Name(name string) *EdgeWrapper {
	ge.Edge.SetName(name)
	return ge
}

func (ge *EdgeWrapper) ID(id int64) *EdgeWrapper {
	ge.Edge.SetID(id)
	return ge
}

func (ge *EdgeWrapper) From(from int64) *EdgeWrapper {
	ge.Edge.SetFrom(from)
	return ge
}

func (ge *EdgeWrapper) To(to int64) *EdgeWrapper {
	ge.Edge.SetTo(to)
	return ge
}

type SccEdgeWrapper struct{ SccEdge }

func MakeSccEdge() *SccEdgeWrapper {
	return &SccEdgeWrapper{&SccEdgeImpl{}}
}

func (ge *SccEdgeWrapper) Obj() SccEdge {
	return ge.SccEdge
}

func (ge *SccEdgeWrapper) Name(name string) *SccEdgeWrapper {
	ge.SccEdge.SetName(name)
	return ge
}

func (ge *SccEdgeWrapper) ID(id int64) *SccEdgeWrapper {
	ge.SccEdge.SetID(id)
	return ge
}

func (ge *SccEdgeWrapper) From(from int64) *SccEdgeWrapper {
	ge.SccEdge.SetFrom(from)
	return ge
}

func (ge *SccEdgeWrapper) To(to int64) *SccEdgeWrapper {
	ge.SccEdge.SetTo(to)
	return ge
}

func (ge *SccEdgeWrapper) Edge(edge Edge) *SccEdgeWrapper {
	ge.SccEdge.SetEdge(edge)
	return ge
}

type RelationWrapper struct{ Relation }

func MakeRelatione() *RelationWrapper {
	return &RelationWrapper{&RelationImpl{}}
}

func (ge *RelationWrapper) Obj() Relation {
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
