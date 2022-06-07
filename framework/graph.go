package framework

type GraphObj interface {
	GetName() string
	SetName(string)
	GetID() int64
	SetID(int64)
	String() string
}

type Graph interface {
	AddEdge(a, b, c int64)
	DeleteEdge(ids ...int64)
	NumNode() int64
	NumEdge() int64

	Head(int64) int64
	Edge(idx int64) int64
	Next(idx int64) int64
}

type GraphImpl struct {
	h, e, w, ne, din   []int64
	fr, pre            []int64
	n, m, tot, deleted int64
}

func NewGraph(n, m int64) Graph {
	g := &GraphImpl{
		h:       make([]int64, n, n),
		e:       make([]int64, m, m),
		w:       make([]int64, m, m),
		ne:      make([]int64, m, m),
		din:     make([]int64, n, n),
		fr:      make([]int64, m, m),
		pre:     make([]int64, m, m),
		n:       n,
		m:       m,
		tot:     0,
		deleted: 0,
	}
	for i := range g.h {
		g.h[i] = -1
	}
	for i := range g.e {
		g.pre[i] = -1
	}
	return g
}

func (g *GraphImpl) AddEdge(a, b, c int64) {
	if g.h[a] >= 0 {
		g.pre[g.h[a]] = g.tot
	}
	g.e[g.tot], g.w[g.tot], g.ne[g.tot], g.h[a] = b, c, g.h[a], g.tot
	g.din[b]++
	g.tot++
}

func (g *GraphImpl) DeleteEdge(ids ...int64) {
	for _, id := range ids {
		if id < 0 || id >= g.tot {
			continue
		}
		if g.pre[id] >= 0 {
			g.ne[g.pre[id]] = g.ne[id]
		} else {
			// This means the current id is head.
			g.h[g.fr[id]] = g.ne[id]
		}
		if g.ne[id] >= 0 {
			g.pre[g.ne[id]] = g.pre[id]
		} else {
			// Do nothing.
		}
		g.deleted++
	}
}

func (g *GraphImpl) NumNode() int64 {
	// TODO: FIXME
	return g.n
}

func (g *GraphImpl) NumEdge() int64 {
	return g.tot - g.deleted
}

func (g *GraphImpl) Head(u int64) int64 {
	return g.h[u]
}

func (g *GraphImpl) Edge(idx int64) int64 {
	return g.e[idx]
}

func (g *GraphImpl) Next(idx int64) int64 {
	return g.ne[idx]
}
