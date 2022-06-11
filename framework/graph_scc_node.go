package framework

import (
	"fmt"

	"github.com/BinacsLee/escheduler/util/set"
)

type SccNode interface {
	Node
	GetNodes() []Node
	AddNodes(...Node)
	RemoveNodes(...Node)
	NumNodes() int64
	GetEdges() []Edge
	AddEdges(...Edge)
	RemoveEdges(...Edge)
	NumEdges() int64
}

type SccNodeImpl struct {
	name string
	id   int64
	// ...
	selfCircle EdgeSet
	nodeSet    NodeSet
}

var (
	_ Node    = &SccNodeImpl{}
	_ SccNode = &SccNodeImpl{}
)

func NewSccNode() SccNode {
	return &SccNodeImpl{
		selfCircle: set.NewSet(),
		nodeSet:    set.NewSet(),
	}
}

func (n *SccNodeImpl) GetName() string {
	return n.name
}

func (n *SccNodeImpl) SetName(name string) {
	n.name = name
}

func (n *SccNodeImpl) GetID() int64 {
	return n.id
}

func (n *SccNodeImpl) SetID(id int64) {
	n.id = id
}

func (n *SccNodeImpl) String() string {
	var nodes string
	for _, obj := range n.nodeSet.List() {
		node := obj.(Node)
		nodes += fmt.Sprintf("{%v}", node.GetName())
	}
	return fmt.Sprintf("{Name:%v{%v},ID:%v}", n.name, nodes, n.id)
}

func (n *SccNodeImpl) AddEdges(edges ...Edge) {
	for i := range edges {
		n.selfCircle.Insert(edges[i])
	}
}

func (n *SccNodeImpl) GetEdges() []Edge {
	objs := n.selfCircle.List()
	ret := make([]Edge, len(objs))
	for i := range objs {
		ret[i] = objs[i].(Edge)
	}
	return ret
}

func (n *SccNodeImpl) RemoveEdges(edges ...Edge) {
	for i := range edges {
		n.selfCircle.Delete(edges[i])
	}
}

func (n *SccNodeImpl) NumEdges() int64 {
	return int64(n.selfCircle.Len())
}

func (n *SccNodeImpl) AddNodes(nodes ...Node) {
	for i := range nodes {
		n.nodeSet.Insert(nodes[i])
	}
}

func (n *SccNodeImpl) GetNodes() []Node {
	objs := n.nodeSet.List()
	ret := make([]Node, len(objs))
	for i := range objs {
		ret[i] = objs[i].(Node)
	}
	return ret
}

func (n *SccNodeImpl) RemoveNodes(nodes ...Node) {
	for i := range nodes {
		n.nodeSet.Delete(nodes[i])
	}
}

func (n *SccNodeImpl) NumNodes() int64 {
	return int64(n.nodeSet.Len())
}
