package tarjan

import (
	"reflect"
	"testing"

	"github.com/BinacsLee/escheduler/framework"
)

func TestSCCTarjan(t *testing.T) {
	type args struct {
		n, m int64
		es   [][]int64
	}
	tests := []struct {
		name string
		args args
		want *SCC
	}{
		{
			name: "normal case 1",
			args: args{
				n: 4,
				m: 5,
				es: [][]int64{
					{0, 1},
					{0, 3},
					{1, 2},
					{2, 3},
					{3, 1},
				},
			},
			want: &SCC{
				dfn:       []int64{1, 3, 4, 2},
				low:       []int64{1, 2, 2, 2},
				id:        []int64{2, 1, 1, 1},
				size:      []int64{0, 3, 1, 0, 0},
				timestamp: 4,
				scc_cnt:   2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := framework.NewGraph(tt.args.n, tt.args.m)
			for i := range tt.args.es {
				e := tt.args.es[i]
				g.AddEdge(e[0], e[1], 1)
			}
			if got := SCCTarjan(g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SCCTarjan() = %v, want %v", got, tt.want)
			}
		})
	}
}
