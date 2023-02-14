package framework

import "github.com/binacs/escheduler/util/set"

type DepthChart interface {
	AddNode(node Node, depth int64)
	DeleteNode(node Node)
	GetNodes(depth int64) NodeSet
	GetMaxDepth() int64
	Len() int64
}

type DepthChartImpl struct {
	chart    map[int64]NodeSet
	hash     map[Node]int64
	maxDepth int64
}

func NewDepthChart() DepthChart {
	return &DepthChartImpl{
		chart:    make(map[int64]NodeSet),
		hash:     make(map[Node]int64),
		maxDepth: 0,
	}
}

func (dc *DepthChartImpl) AddNode(node Node, depth int64) {
	if _, ok := dc.hash[node]; ok {
		return
	}
	if dc.chart[depth] == nil {
		dc.chart[depth] = set.NewSet()
	}
	dc.chart[depth].Insert(node)
	if depth > dc.maxDepth {
		dc.maxDepth = depth
	}
	dc.hash[node] = depth
}

func (dc *DepthChartImpl) DeleteNode(node Node) {
	if _, ok := dc.hash[node]; !ok {
		return
	}
	depth := dc.hash[node]
	if dc.chart[depth] == nil {
		return
	}
	dc.chart[depth].Delete(node)
	if dc.chart[depth].Len() == 0 {
		delete(dc.chart, depth)
		if depth == dc.maxDepth {
			dc.maxDepth--
		}
	}
	delete(dc.hash, node)
}

func (dc *DepthChartImpl) GetNodes(depth int64) NodeSet {
	return dc.chart[depth]
}

func (dc *DepthChartImpl) GetMaxDepth() int64 {
	return dc.maxDepth
}

func (dc *DepthChartImpl) Len() int64 {
	return int64(len(dc.hash))
}
