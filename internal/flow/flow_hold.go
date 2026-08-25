package flow

// liveHoldAlias hands back a two-slot summary scratch. BuildFlow keeps
// volumetric flow in slot 0 and Reynolds number in slot 1.
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
	buf[1] = re
	return buf[0]
}
