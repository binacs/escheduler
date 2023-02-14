package tarjan

import (
	"fmt"

	"github.com/binacs/escheduler/framework"
)

type SCC interface {
	GetID(x int64) int64
	GetSccCnt() int64
}

type SCCImpl struct {
	dfn, low           []int64
	id, size           []int64
	timestamp, scc_cnt int64
}

var _ SCC = &SCCImpl{}

func newSCC(g framework.Graph) *SCCImpl {
	return &SCCImpl{
		dfn:       make([]int64, g.NumNode()),
		low:       make([]int64, g.NumNode()),
		id:        make([]int64, g.NumNode()),
		size:      make([]int64, g.NumNode()),
		timestamp: 0,
		scc_cnt:   0,
	}
}

func (scc *SCCImpl) GetID(x int64) int64 {
	return scc.id[x]
}

func (scc *SCCImpl) GetSccCnt() int64 {
	return scc.scc_cnt
}

func (scc *SCCImpl) String() string {
	var nodes, sccsize string
	for i := 0; i < len(scc.dfn); i++ {
		// nodes += fmt.Sprintf("{idx:%v,dfn:%v,low:%v,id:%v}", i, scc.dfn[i], scc.low[i], scc.id[i])
		nodes += fmt.Sprintf("{idx:%v,id:%v},", i, scc.id[i])
	}
	for i := 0; i < int(scc.scc_cnt); i++ {
		sccsize += fmt.Sprintf("{sccid:%v,size:%v},", i, scc.size[i])
	}
	return fmt.Sprintf("[%v,%v]", nodes, sccsize)
}

func SCCTarjan(g framework.Graph) SCC {
	scc := newSCC(g)

	stk := make([]int64, g.NumNode()+1)
	in_stk := make([]bool, g.NumNode()+1)
	top := 0

	var tarjan func(u int64)
	tarjan = func(u int64) {
		{
			scc.timestamp++
			scc.dfn[u], scc.low[u] = scc.timestamp, scc.timestamp
		}
		{
			top++
			stk[top], in_stk[u] = u, true
		}
		for i := g.Head(u); i != -1; i = g.Next(i) {
			j := g.Edge(i)
			if scc.dfn[j] == 0 {
				tarjan(j)
				if scc.low[j] < scc.low[u] {
					scc.low[u] = scc.low[j]
				}
			} else if in_stk[j] {
				if scc.dfn[j] < scc.low[u] {
					scc.low[u] = scc.dfn[j]
				}
			}
		}
		if scc.dfn[u] == scc.low[u] {
			for {
				y := stk[top]
				top--
				in_stk[y] = false
				scc.id[y] = scc.scc_cnt
				scc.size[scc.scc_cnt]++
				if y == u {
					break
				}
			}
			scc.scc_cnt++
		}
	}

	for i := int64(0); i < g.NumNode(); i++ {
		if scc.dfn[i] == 0 {
			tarjan(i)
		}
	}
	return scc
}
