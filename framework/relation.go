package framework

import "fmt"

type RelationAttr interface{}

type Relation interface {
	GetFrom() string
	SetFrom(string)
	GetTo() string
	SetTo(string)
	GetName() string
	SetName(string)
	String() string
}

type RelationImpl struct {
	from, to  string
	name      string
	attribute RelationAttr
}

var (
	_ Relation = &RelationImpl{}
)

func (r *RelationImpl) GetFrom() string {
	return r.from
}

func (r *RelationImpl) SetFrom(from string) {
	r.from = from
}

func (r *RelationImpl) GetTo() string {
	return r.to
}

func (r *RelationImpl) SetTo(to string) {
	r.to = to
}

func (r *RelationImpl) GetName() string {
	return r.name
}

func (r *RelationImpl) SetName(name string) {
	r.name = name
}

func (r *RelationImpl) String() string {
	return fmt.Sprintf("{Name:%s,From:%s,To:%s}", r.name, r.from, r.to)
}
