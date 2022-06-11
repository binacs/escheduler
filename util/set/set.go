package set

import "fmt"

type Set interface {
	Exist(item interface{}) bool
	Insert(item ...interface{})
	Delete(item ...interface{})
	Len() int
	List() []interface{}
}

type SetImpl struct {
	data map[interface{}]struct{}
}

func NewSet() Set {
	return &SetImpl{
		data: make(map[interface{}]struct{}),
	}
}

func (s *SetImpl) Exist(item interface{}) bool {
	if s == nil {
		return false
	}
	_, ok := s.data[item]
	return ok
}

func (s *SetImpl) Insert(items ...interface{}) {
	if s == nil {
		return
	}
	for _, item := range items {
		if _, ok := s.data[item]; ok {
			continue
		}
		s.data[item] = struct{}{}
	}
}

func (s *SetImpl) Delete(items ...interface{}) {
	if s == nil {
		return
	}
	for _, item := range items {
		if _, ok := s.data[item]; !ok {
			continue
		}
		delete(s.data, item)
	}
}

func (s *SetImpl) Len() int {
	if s == nil {
		return 0
	}
	return len(s.data)
}

func (s *SetImpl) List() []interface{} {
	if s == nil {
		return nil
	}
	ret := make([]interface{}, 0)
	for k := range s.data {
		ret = append(ret, k)
	}
	return ret
}

func (s *SetImpl) String() string {
	if s == nil {
		return "{}"
	}
	var str string
	for k := range s.data {
		str += fmt.Sprintf("{%v},", s.data[k])
	}
	return str
}
