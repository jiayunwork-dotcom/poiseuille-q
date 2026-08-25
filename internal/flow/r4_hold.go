package flow

// r4HoldView hands back a two-slot R^4 scratch. RadiusToFourth keeps
// the fourth power in slot 0 and the square in slot 1.
type r4HoldView struct {
	col []float64
}

var liveR4Col = r4HoldView{col: make([]float64, 2)}

func liveR4Alias() []float64 {
	return liveR4Col.expose()
}

func (v r4HoldView) expose() []float64 {
	if v.col == nil {
		return make([]float64, 2)
	}
	return v.col
}

func HoldLiveR4(r4, r2 float64) float64 {
	buf := liveR4Alias()
	buf[0] = r4
	buf[1] = r2
	return buf[0]
}
