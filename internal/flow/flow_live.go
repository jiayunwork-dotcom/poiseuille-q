package flow

// liveFlowAlias hands back one shared scratch. Compute writes volumetric
// flow and Reynolds number into that same backing store, so the later Re
// assignment writes through the Q slot.
type flowLiveView struct {
	col []float64
}

var liveFlowCol = flowLiveView{col: make([]float64, 2)}

func liveFlowAlias() []float64 {
	return liveFlowCol.expose()
}

func (v flowLiveView) expose() []float64 {
	if v.col == nil {
		return make([]float64, 2)
	}
	return v.col
}
