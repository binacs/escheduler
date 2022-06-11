package decision

import (
	"context"

	"github.com/BinacsLee/escheduler/framework"
	"github.com/BinacsLee/escheduler/util/names"
	"github.com/BinacsLee/escheduler/util/set"
)

const (
	Name = names.DefaultDecision
)

type DefaultDecision struct {
}

var (
	_ framework.Decision = &DefaultDecision{}
)

func (d *DefaultDecision) AfterInject() error {
	return nil
}

func (d *DefaultDecision) Name() string {
	return Name
}

func (d *DefaultDecision) SelectEdges(
	ctx context.Context,
	g framework.Graph,
	chart framework.DepthChart,
	checkFunc framework.CheckFunc,
) (framework.EdgeSet, bool) {
	selectedEdges := set.NewSet()
	for dep := chart.GetMaxDepth(); dep >= 0; dep-- {
		sccNodeSet, changed := chart.GetNodes(dep), false

		// Select self edge firstly.
		for _, obj := range sccNodeSet.List() {
			sccNode := obj.(framework.SccNode)
			selfEdges := sccNode.GetEdges()
			// TODO: check function.
			if checkFunc(selectedEdges, selfEdges...) {
				for i := range selfEdges {
					selectedEdges.Insert(selfEdges[i])
				}
				sccNode.RemoveEdges(selfEdges...)
				changed = true
			}
		}

		for _, obj := range sccNodeSet.List() {
			sccNode := obj.(framework.SccNode)
			if sccNode.NumEdges() > 0 {
				continue
			}
			deleteSet := set.NewSet()
			i := sccNode.GetID()
			for j := g.Head(i); j != -1; j = g.Next(j) {
				e := g.GetEdge(j).(framework.SccEdge).GetEdge()
				if checkFunc(selectedEdges, e) {
					selectedEdges.Insert(e)
					deleteSet.Insert(j)
					changed = true
				}
			}
			if deleteSet.Len() > 0 {
				for _, obj := range deleteSet.List() {
					i := obj.(int64)
					g.DeleteEdge(i)
				}
			}
		}
		for _, obj := range sccNodeSet.List() {
			sccNode := obj.(framework.SccNode)
			if sccNode.NumEdges() == 0 && g.Head(sccNode.GetID()) == -1 {
				chart.DeleteNode(sccNode)
			}
		}

		if !changed {
			break
		}

		for _, obj := range selectedEdges.List() {
			selectedEdges.Insert(obj)
		}

		if chart.GetNodes(dep) != nil && chart.GetNodes(dep).Len() > 0 {
			break
		}
	}

	if selectedEdges.Len() == 0 {
		return nil, true
	}

	return selectedEdges, false
}
