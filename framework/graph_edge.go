package framework

import "fmt"

type GraphEdge interface {
	GraphObj
	GetFrom() int64
	SetFrom(int64)
	GetTo() int64
	SetTo(int64)
}

type GraphEdgeImpl struct {
	name string
	id   int64
	// ...
	from, to int64
}

var (
	_ GraphEdge = &GraphEdgeImpl{}
)

func (e *GraphEdgeImpl) GetName() string {
	return e.name
}

func (e *GraphEdgeImpl) SetName(name string) {
	e.name = name
}

func (e *GraphEdgeImpl) GetID() int64 {
	return e.id
}

func (e *GraphEdgeImpl) SetID(id int64) {
	e.id = id
}

func (e *GraphEdgeImpl) GetFrom() int64 {
	return e.from
}

func (e *GraphEdgeImpl) SetFrom(from int64) {
	e.from = from
}

func (e *GraphEdgeImpl) GetTo() int64 {
	return e.to
}

func (e *GraphEdgeImpl) SetTo(to int64) {
	e.to = to
}

func (e *GraphEdgeImpl) String() string {
	return fmt.Sprintf("{Name:%v,ID:%v,From:%v,To:%v}", e.name, e.id, e.from, e.to)
}
