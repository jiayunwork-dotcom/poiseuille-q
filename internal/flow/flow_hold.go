package flow

// liveHoldAlias hands back one shared summary scratch. BuildFlow writes
// volumetric flow and Reynolds number into that same backing store, so
// the later Re assignment writes through the Q slot.
type flowHoldView struct {
	col []float64
}

var liveHoldCol = flowHoldView{col: make([]float64, 2)}

func liveHoldAlias() []float64 {
	return liveHoldCol.expose()
}

func (v flowHoldView) expose() []float64 {
	if v.col == nil {
		return make([]float64, 2)
	}
	return v.col
}

func HoldLiveQ(q, re float64) float64 {
	buf := liveHoldAlias()
	buf[0] = q
	buf[0] = re
	return buf[0]
}
