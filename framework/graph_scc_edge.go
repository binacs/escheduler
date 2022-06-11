package framework

import "fmt"

type SccEdge interface {
	Edge
	SetEdge(Edge)
	GetEdge() Edge
}

type SccEdgeImpl struct {
	name string
	id   int64
	// ...
	from, to int64
	nodeEdge Edge
}

var (
	_ SccEdge = &SccEdgeImpl{}
)

func (e *SccEdgeImpl) GetName() string {
	return e.name
}

func (e *SccEdgeImpl) SetName(name string) {
	e.name = name
}

func (e *SccEdgeImpl) GetID() int64 {
	return e.id
}

func (e *SccEdgeImpl) SetID(id int64) {
	e.id = id
}

func (e *SccEdgeImpl) GetFrom() int64 {
	return e.from
}

func (e *SccEdgeImpl) SetFrom(from int64) {
	e.from = from
}

func (e *SccEdgeImpl) GetTo() int64 {
	return e.to
}

func (e *SccEdgeImpl) SetTo(to int64) {
	e.to = to
}

func (e *SccEdgeImpl) GetEdge() Edge {
	return e.nodeEdge
}

func (e *SccEdgeImpl) SetEdge(edge Edge) {
	e.nodeEdge = edge
}

func (e *SccEdgeImpl) String() string {
	return fmt.Sprintf("{Name:%v,ID:%v,From:%v,To:%v}", e.name, e.id, e.from, e.to)
}
