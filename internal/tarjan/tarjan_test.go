package tarjan

import (
	"reflect"
	"testing"

	"github.com/binacs/escheduler/framework"
)

func TestSCCTarjan(t *testing.T) {
	type args struct {
		n, m  int64
		nodes []framework.Node
		edges []framework.Edge
	}
	tests := []struct {
		name string
		args args
		want SCC
	}{
		{
			name: "normal case 1",
			args: args{
				nodes: []framework.Node{
					framework.MakeNode().ID(0),
					framework.MakeNode().ID(1),
					framework.MakeNode().ID(2),
					framework.MakeNode().ID(3),
				},
				edges: []framework.Edge{
					framework.MakeEdge().From(0).To(1).Obj(),
					framework.MakeEdge().From(0).To(3).Obj(),
					framework.MakeEdge().From(1).To(2).Obj(),
					framework.MakeEdge().From(2).To(3).Obj(),
					framework.MakeEdge().From(3).To(1).Obj(),
				},
			},
			want: &SCCImpl{
				dfn:       []int64{1, 3, 4, 2},
				low:       []int64{1, 2, 2, 2},
				id:        []int64{1, 0, 0, 0},
				size:      []int64{3, 1, 0, 0},
				timestamp: 4,
				scc_cnt:   2,
			},
		},
		{
			name: "duplicate edges",
			args: args{
				nodes: []framework.Node{
					framework.MakeNode().ID(0),
					framework.MakeNode().ID(1),
				},
				edges: []framework.Edge{
					framework.MakeEdge().From(0).To(1).Obj(),
					framework.MakeEdge().From(0).To(1).Obj(),
					framework.MakeEdge().From(0).To(1).Obj(),
				},
			},
			want: &SCCImpl{
				dfn:       []int64{1, 2},
				low:       []int64{1, 2},
				id:        []int64{1, 0},
				size:      []int64{1, 1},
				timestamp: 2,
				scc_cnt:   2,
			},
		},
		{
			name: "duplicate edges and self circle",
			args: args{
				nodes: []framework.Node{
					framework.MakeNode().ID(0),
					framework.MakeNode().ID(1),
				},
				edges: []framework.Edge{
					framework.MakeEdge().From(0).To(1).Obj(),
					framework.MakeEdge().From(0).To(1).Obj(),
					framework.MakeEdge().From(0).To(1).Obj(),
					framework.MakeEdge().From(1).To(0).Obj(),
					framework.MakeEdge().From(1).To(0).Obj(),
				},
			},
			want: &SCCImpl{
				dfn:       []int64{1, 2},
				low:       []int64{1, 1},
				id:        []int64{0, 0},
				size:      []int64{2, 0},
				timestamp: 2,
				scc_cnt:   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := framework.NewGraph(tt.args.nodes, tt.args.edges)
			if got := SCCTarjan(g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SCCTarjan() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
