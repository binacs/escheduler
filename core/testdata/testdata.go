package testdata

import (
	"math/rand"

	"github.com/BinacsLee/escheduler/framework"
	"github.com/BinacsLee/escheduler/testing"
)

type TestGraph struct {
	Nodes []framework.GraphNode
	Edges []framework.GraphEdge
}

var TestGraphSet = []TestGraph{
	{
		[]framework.GraphNode{
			testing.MakeGraphNode().Name("0").ID(0).Obj(),
			testing.MakeGraphNode().Name("1").ID(1).Obj(),
			testing.MakeGraphNode().Name("2").ID(2).Obj(),
		},
		[]framework.GraphEdge{
			testing.MakeGraphEdge().Name("e0_0->1").ID(0).From(0).To(1).Obj(),
			testing.MakeGraphEdge().Name("e1_1->2").ID(1).From(1).To(2).Obj(),
		},
	},
	{
		[]framework.GraphNode{
			testing.MakeGraphNode().Name("0").ID(0).Obj(),
			testing.MakeGraphNode().Name("1").ID(1).Obj(),
			testing.MakeGraphNode().Name("2").ID(2).Obj(),
			testing.MakeGraphNode().Name("3").ID(3).Obj(),
			testing.MakeGraphNode().Name("4").ID(4).Obj(),
			testing.MakeGraphNode().Name("5").ID(5).Obj(),
			testing.MakeGraphNode().Name("6").ID(6).Obj(),
			testing.MakeGraphNode().Name("7").ID(7).Obj(),
			testing.MakeGraphNode().Name("8").ID(8).Obj(),
		},
		[]framework.GraphEdge{
			testing.MakeGraphEdge().Name("e0_0->1").ID(0).From(0).To(1).Obj(),
			testing.MakeGraphEdge().Name("e1_1->2").ID(1).From(1).To(2).Obj(),
			testing.MakeGraphEdge().Name("e2_2->3").ID(2).From(2).To(3).Obj(),
			testing.MakeGraphEdge().Name("e3_3->4").ID(3).From(3).To(4).Obj(),
			testing.MakeGraphEdge().Name("e4_4->1").ID(4).From(4).To(1).Obj(),
			testing.MakeGraphEdge().Name("e5_1->5").ID(5).From(1).To(5).Obj(),
			testing.MakeGraphEdge().Name("e6_2->5").ID(6).From(2).To(5).Obj(),
			testing.MakeGraphEdge().Name("e7_5->6").ID(7).From(5).To(6).Obj(),
			testing.MakeGraphEdge().Name("e8_5->7").ID(8).From(5).To(7).Obj(),
			testing.MakeGraphEdge().Name("e9_6->7").ID(9).From(6).To(7).Obj(),
			testing.MakeGraphEdge().Name("e10_6->8").ID(10).From(6).To(8).Obj(),
		},
	},
	{
		[]framework.GraphNode{
			// NodeSet 1
			testing.MakeGraphNode().Name("0").ID(0).Obj(),
			testing.MakeGraphNode().Name("1").ID(1).Obj(),
			testing.MakeGraphNode().Name("2").ID(2).Obj(),
			testing.MakeGraphNode().Name("3").ID(3).Obj(),
			// NodeSet 2
			testing.MakeGraphNode().Name("4").ID(4).Obj(),
			testing.MakeGraphNode().Name("5").ID(5).Obj(),
			testing.MakeGraphNode().Name("6").ID(6).Obj(),
			testing.MakeGraphNode().Name("7").ID(7).Obj(),
			testing.MakeGraphNode().Name("8").ID(8).Obj(),
		},
		[]framework.GraphEdge{
			// NodeSet 1
			testing.MakeGraphEdge().Name("e0_0->1").ID(0).From(0).To(2).Obj(),
			testing.MakeGraphEdge().Name("e1_1->2").ID(1).From(1).To(2).Obj(),
			testing.MakeGraphEdge().Name("e2_2->3").ID(2).From(2).To(3).Obj(),
			testing.MakeGraphEdge().Name("e3_0->3").ID(3).From(0).To(3).Obj(),
			// NodeSet 2
			testing.MakeGraphEdge().Name("e4_4->6").ID(4).From(4).To(6).Obj(),
			testing.MakeGraphEdge().Name("e5_5->6").ID(5).From(5).To(6).Obj(),
			testing.MakeGraphEdge().Name("e6_5->7").ID(6).From(5).To(7).Obj(),
			testing.MakeGraphEdge().Name("e7_6->7").ID(7).From(6).To(7).Obj(),
			testing.MakeGraphEdge().Name("e8_7->8").ID(8).From(7).To(8).Obj(),
		},
	},
}

type TestData struct {
	Relations []framework.Relation
}

var TestDataSet = []TestData{
	{
		[]framework.Relation{
			testing.MakeRelatione().Name("1").From("A").To("B").Obj(),
			testing.MakeRelatione().Name("2").From("B").To("C").Obj(),
		},
	},
	{
		[]framework.Relation{
			testing.MakeRelatione().Name("1").From("A").To("B").Obj(),
			testing.MakeRelatione().Name("2").From("B").To("C").Obj(),
			testing.MakeRelatione().Name("3").From("C").To("D").Obj(),
			testing.MakeRelatione().Name("4").From("D").To("E").Obj(),
			testing.MakeRelatione().Name("5").From("E").To("B").Obj(),
			testing.MakeRelatione().Name("6").From("B").To("F").Obj(),
			testing.MakeRelatione().Name("7").From("C").To("F").Obj(),
			testing.MakeRelatione().Name("8").From("F").To("G").Obj(),
			testing.MakeRelatione().Name("9").From("F").To("H").Obj(),
			testing.MakeRelatione().Name("10").From("G").To("H").Obj(),
			testing.MakeRelatione().Name("11").From("6").To("I").Obj(),
		},
	},
	{
		[]framework.Relation{
			// NodeSet 1
			testing.MakeRelatione().Name("1").From("A").To("B").Obj(),
			testing.MakeRelatione().Name("2").From("B").To("C").Obj(),
			testing.MakeRelatione().Name("3").From("C").To("D").Obj(),
			testing.MakeRelatione().Name("4").From("A").To("D").Obj(),
			// NodeSet 2
			testing.MakeRelatione().Name("5").From("E").To("G").Obj(),
			testing.MakeRelatione().Name("6").From("F").To("G").Obj(),
			testing.MakeRelatione().Name("7").From("F").To("H").Obj(),
			testing.MakeRelatione().Name("8").From("G").To("H").Obj(),
			testing.MakeRelatione().Name("9").From("H").To("I").Obj(),
		},
	},
}

func SelectRandomTestData() TestData {
	return TestDataSet[rand.Int()%len(TestDataSet)]
}
