package framework

import "fmt"

type GraphNode interface {
	GraphObj
}

type GraphNodeImpl struct {
	name string
	id   int64
	// ...
}

var (
	_ GraphNode = &GraphNodeImpl{}
)

func (n *GraphNodeImpl) GetName() string {
	return n.name
}

func (n *GraphNodeImpl) SetName(name string) {
	n.name = name
}

func (n *GraphNodeImpl) GetID() int64 {
	return n.id
}

func (n *GraphNodeImpl) SetID(id int64) {
	n.id = id
}

func (n *GraphNodeImpl) String() string {
	return fmt.Sprintf("{Name:%v,ID:%v}", n.name, n.id)
}
