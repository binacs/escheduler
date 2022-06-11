package dsu

import "sync"

// github.com/binacsgo/datastructure/dsu

type DSU interface {
	Find(x int64) int64
	Merge(x, y int64) bool
	SameSet(x, y int64) bool
	SetSize(x int64) int64
	SCC() int64
	Squash()
}

type DSUImpl struct {
	n, scc int64
	p      []int64
	size   []int64
	mu     sync.RWMutex
}

var _ DSU = &DSUImpl{}

func NewDSU(n int64) DSU {
	p, size := make([]int64, n+1), make([]int64, n+1)
	for i := int64(0); i <= n; i++ {
		p[i], size[i] = i, 1
	}
	return &DSUImpl{
		n:    n,
		scc:  n,
		p:    p,
		size: size,
	}
}

func (u *DSUImpl) Find(x int64) int64 {
	u.mu.Lock()
	defer u.mu.Unlock()
	return u.find(x)
}

func (u *DSUImpl) Merge(x, y int64) bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	return u.merge(x, y)
}

func (u *DSUImpl) SameSet(x, y int64) bool {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.find(x) == u.find(y)
}

func (u *DSUImpl) SetSize(x int64) int64 {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.size[u.find(x)]
}

func (u *DSUImpl) SCC() int64 {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.scc
}

func (u *DSUImpl) Squash() {
	u.mu.Lock()
	defer u.mu.Unlock()
	for i := int64(0); i <= u.n; i++ {
		u.p[i] = u.find(i)
	}
}

func (u *DSUImpl) find(x int64) int64 {
	if u.p[x] != x {
		u.p[x] = u.find(u.p[x])
	}
	return u.p[x]
}

func (u *DSUImpl) merge(x, y int64) bool {
	fx, fy := u.find(x), u.find(y)
	if fx == fy {
		return false
	}
	if u.size[fx] > u.size[fy] {
		fx, fy = fy, fx
	}
	u.p[fx] = fy
	u.size[fy] += u.size[fx]
	u.scc--
	return true
}
