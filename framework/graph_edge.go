package framework

import "fmt"

type Edge interface {
	GraphObj
	GetFrom() int64
	SetFrom(int64)
	GetTo() int64
	SetTo(int64)
}

type EdgeImpl struct {
	name string
	id   int64
	// ...
	from, to int64
}

var (
	_ Edge = &EdgeImpl{}
)

func (e *EdgeImpl) GetName() string {
	return e.name
}

func (e *EdgeImpl) SetName(name string) {
	e.name = name
}

func (e *EdgeImpl) GetID() int64 {
	return e.id
}

func (e *EdgeImpl) SetID(id int64) {
	e.id = id
}

func (e *EdgeImpl) GetFrom() int64 {
	return e.from
}

func (e *EdgeImpl) SetFrom(from int64) {
	e.from = from
}

func (e *EdgeImpl) GetTo() int64 {
	return e.to
}

func (e *EdgeImpl) SetTo(to int64) {
	e.to = to
}

func (e *EdgeImpl) String() string {
	return fmt.Sprintf("{Name:%v,ID:%v,From:%v,To:%v}", e.name, e.id, e.from, e.to)
}
