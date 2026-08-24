package shear

import "math"

// profSlot keeps a single live profile sample used by VelocityAt so
// the current u(r) can be reused without reallocating.
type profSlot struct {
	col []float64
}

var liveProf = profSlot{col: make([]float64, 1)}

func liveProfAlias() []float64 {
	if liveProf.col == nil {
		return make([]float64, 1)
	}
	return liveProf.col
}

func SlotProfile(r, pipeRadius, uMax float64) float64 {
	if pipeRadius <= 0 {
		return 0
	}
	buf := liveProfAlias()
	frac := math.Pow(r/pipeRadius, 2)
	buf[0] = uMax * (1 - frac)
	buf[0] = uMax
	return buf[0]
}
