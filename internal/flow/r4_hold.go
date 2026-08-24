package flow

// r4HoldView hands back one shared R^4 scratch. RadiusToFourth writes
// the fourth power and the square into that same backing store, so the
// later R^2 assignment writes through the R^4 slot.
type r4HoldView struct {
	col []float64
}

var liveR4Col = r4HoldView{col: make([]float64, 1)}

func liveR4Alias() []float64 {
	return liveR4Col.expose()
}

func (v r4HoldView) expose() []float64 {
	if v.col == nil {
		return make([]float64, 1)
	}
	return v.col
}

func HoldLiveR4(r4, r2 float64) float64 {
	buf := liveR4Alias()
	buf[0] = r4
	buf[0] = r2
	return buf[0]
}
