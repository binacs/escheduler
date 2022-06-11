package framework

import "fmt"

type Node interface {
	GraphObj
}

type NodeImpl struct {
	name string
	id   int64
	// ...
}

var (
	_ Node = &NodeImpl{}
)

func NewNode() Node {
	return &NodeImpl{}
}

func (n *NodeImpl) GetName() string {
	return n.name
}

func (n *NodeImpl) SetName(name string) {
	n.name = name
}

func (n *NodeImpl) GetID() int64 {
	return n.id
}

func (n *NodeImpl) SetID(id int64) {
	n.id = id
}

func (n *NodeImpl) String() string {
	return fmt.Sprintf("{Name:%v,ID:%v}", n.name, n.id)
}
