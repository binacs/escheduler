package testdata

import (
	"math/rand"

	"github.com/binacs/escheduler/framework"
)

type TestGraph struct {
	Nodes []framework.Node
	Edges []framework.Edge
}

var TestGraphSet = []TestGraph{
	{
		[]framework.Node{
			framework.MakeNode().Name("0").ID(0).Obj(),
			framework.MakeNode().Name("1").ID(1).Obj(),
			framework.MakeNode().Name("2").ID(2).Obj(),
		},
		[]framework.Edge{
			framework.MakeEdge().Name("e0_0->1").ID(0).From(0).To(1).Obj(),
			framework.MakeEdge().Name("e1_1->2").ID(1).From(1).To(2).Obj(),
		},
	},
	{
		[]framework.Node{
			framework.MakeNode().Name("0").ID(0).Obj(),
			framework.MakeNode().Name("1").ID(1).Obj(),
			framework.MakeNode().Name("2").ID(2).Obj(),
			framework.MakeNode().Name("3").ID(3).Obj(),
			framework.MakeNode().Name("4").ID(4).Obj(),
			framework.MakeNode().Name("5").ID(5).Obj(),
			framework.MakeNode().Name("6").ID(6).Obj(),
			framework.MakeNode().Name("7").ID(7).Obj(),
			framework.MakeNode().Name("8").ID(8).Obj(),
		},
		[]framework.Edge{
			framework.MakeEdge().Name("e0_0->1").ID(0).From(0).To(1).Obj(),
			framework.MakeEdge().Name("e1_1->2").ID(1).From(1).To(2).Obj(),
			framework.MakeEdge().Name("e2_2->3").ID(2).From(2).To(3).Obj(),
			framework.MakeEdge().Name("e3_3->4").ID(3).From(3).To(4).Obj(),
			framework.MakeEdge().Name("e4_4->1").ID(4).From(4).To(1).Obj(),
			framework.MakeEdge().Name("e5_1->5").ID(5).From(1).To(5).Obj(),
			framework.MakeEdge().Name("e6_2->5").ID(6).From(2).To(5).Obj(),
			framework.MakeEdge().Name("e7_5->6").ID(7).From(5).To(6).Obj(),
			framework.MakeEdge().Name("e8_5->7").ID(8).From(5).To(7).Obj(),
			framework.MakeEdge().Name("e9_6->7").ID(9).From(6).To(7).Obj(),
			framework.MakeEdge().Name("e10_6->8").ID(10).From(6).To(8).Obj(),
		},
	},
	{
		[]framework.Node{
			// SccNode 1
			framework.MakeNode().Name("0").ID(0).Obj(),
			framework.MakeNode().Name("1").ID(1).Obj(),
			framework.MakeNode().Name("2").ID(2).Obj(),
			framework.MakeNode().Name("3").ID(3).Obj(),
			// SccNode 2
			framework.MakeNode().Name("4").ID(4).Obj(),
			framework.MakeNode().Name("5").ID(5).Obj(),
			framework.MakeNode().Name("6").ID(6).Obj(),
			framework.MakeNode().Name("7").ID(7).Obj(),
			framework.MakeNode().Name("8").ID(8).Obj(),
		},
		[]framework.Edge{
			// SccNode 1
			framework.MakeEdge().Name("e0_0->1").ID(0).From(0).To(2).Obj(),
			framework.MakeEdge().Name("e1_1->2").ID(1).From(1).To(2).Obj(),
			framework.MakeEdge().Name("e2_2->3").ID(2).From(2).To(3).Obj(),
			framework.MakeEdge().Name("e3_0->3").ID(3).From(0).To(3).Obj(),
			// SccNode 2
			framework.MakeEdge().Name("e4_4->6").ID(4).From(4).To(6).Obj(),
			framework.MakeEdge().Name("e5_5->6").ID(5).From(5).To(6).Obj(),
			framework.MakeEdge().Name("e6_5->7").ID(6).From(5).To(7).Obj(),
			framework.MakeEdge().Name("e7_6->7").ID(7).From(6).To(7).Obj(),
			framework.MakeEdge().Name("e8_7->8").ID(8).From(7).To(8).Obj(),
		},
	},
}

type TestData struct {
	Relations []framework.Relation
}

var TestDataSet = []TestData{
	{
		[]framework.Relation{
			framework.MakeRelatione().Name("1").From("A").To("B").Obj(),
			framework.MakeRelatione().Name("2").From("B").To("C").Obj(),
		},
	},
	{
		[]framework.Relation{
			framework.MakeRelatione().Name("1").From("A").To("B").Obj(),
			framework.MakeRelatione().Name("2").From("B").To("C").Obj(),
			framework.MakeRelatione().Name("3").From("C").To("D").Obj(),
			framework.MakeRelatione().Name("4").From("D").To("E").Obj(),
			framework.MakeRelatione().Name("5").From("E").To("B").Obj(),
			framework.MakeRelatione().Name("6").From("B").To("F").Obj(),
			framework.MakeRelatione().Name("7").From("C").To("F").Obj(),
			framework.MakeRelatione().Name("8").From("F").To("G").Obj(),
			framework.MakeRelatione().Name("9").From("F").To("H").Obj(),
			framework.MakeRelatione().Name("10").From("G").To("H").Obj(),
			framework.MakeRelatione().Name("11").From("G").To("I").Obj(),
		},
	},
	{
		// {idx:0,id:4},{idx:1,id:3},{idx:2,id:2},{idx:3,id:1}
		// {idx:4,id:8},{idx:5,id:7},{idx:6,id:9},{idx:7,id:6},{idx:8,id:5},
		[]framework.Relation{
			// SccNode 1
			framework.MakeRelatione().Name("1").From("A").To("B").Obj(),
			framework.MakeRelatione().Name("2").From("B").To("C").Obj(),
			framework.MakeRelatione().Name("3").From("C").To("D").Obj(),
			framework.MakeRelatione().Name("4").From("A").To("D").Obj(),
			// SccNode 2
			framework.MakeRelatione().Name("5").From("E").To("G").Obj(),
			framework.MakeRelatione().Name("6").From("F").To("G").Obj(),
			framework.MakeRelatione().Name("7").From("F").To("H").Obj(),
			framework.MakeRelatione().Name("8").From("G").To("H").Obj(),
			framework.MakeRelatione().Name("9").From("H").To("I").Obj(),
		},
	},
}

func SelectRandomTestData() TestData {
	return TestDataSet[rand.Int()%len(TestDataSet)]
}

func SelectTestData(i int) TestData {
	return TestDataSet[i]
}
