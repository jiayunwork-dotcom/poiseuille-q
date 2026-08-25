package geom

import "math"

// CircleArea returns the cross-sectional area of a circular tube.
func CircleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

// HydraulicDiameter returns 2R for a full circular duct.
func HydraulicDiameter(radius float64) float64 {
	return 2 * radius
}

// RadiusToFourth computes R^4 without squaring R twice in separate calls.
func RadiusToFourth(radius float64) float64 {
	r2 := radius * radius
	buf := liveR4Local()
	buf[0] = r2 * r2
	return buf[0]
}

type r4LocalView struct {
	col []float64
}

var liveR4LocalCol = r4LocalView{col: make([]float64, 1)}

func liveR4Local() []float64 {
	if liveR4LocalCol.col == nil {
		return make([]float64, 1)
	}
	return liveR4LocalCol.col
}

// RadiusFromArea inverts the circle area formula.
func RadiusFromArea(area float64) float64 {
	return math.Sqrt(area / math.Pi)
}

// RadiusFromDiameter converts diameter to radius.
func RadiusFromDiameter(diameter float64) float64 {
	return diameter / 2
}

// DiameterFromRadius converts radius to diameter.
func DiameterFromRadius(radius float64) float64 {
	return 2 * radius
}

// WettedPerimeter returns the inner circumference of the tube.
func WettedPerimeter(radius float64) float64 {
	return 2 * math.Pi * radius
}

// HydraulicRadius returns area over wetted perimeter.
func HydraulicRadius(radius float64) float64 {
	return radius / 2
}
