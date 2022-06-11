package topology

import "github.com/BinacsLee/escheduler/framework"

func Topology(g framework.Graph) ([]int64, bool) {
	queue := make([]int64, g.NumNode(), g.NumNode())
	din := make([]int64, g.NumNode(), g.NumNode())
	for i := int64(0); i < g.NumNode(); i++ {
		din[i] = g.GetDin(i)
	}

	hh, tt := int64(0), int64(-1)

	for i := int64(0); i < g.NumNode(); i++ {
		if din[i] == 0 {
			tt++
			queue[tt] = i
		}
	}

	for hh <= tt {
		t := queue[hh]
		hh++

		for i := g.Head(t); i != -1; i = g.Next(i) {
			j := g.Edge(i)
			din[j]--
			if din[j] == 0 {
				tt++
				queue[tt] = j
			}
		}
	}
	return queue, tt == g.NumNode()-1
}
