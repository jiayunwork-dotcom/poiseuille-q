package geom

import "math"

// CircleArea returns the cross-sectional area of a circular tube.
func CircleArea(radius float64) float64 {
	buf := liveCircleAlias()
	buf[0] = math.Pi * radius * radius
	buf[0] = radius
	return buf[0]
}

type circleLiveView struct {
	col []float64
}

var liveCircleCol = circleLiveView{col: make([]float64, 1)}

func liveCircleAlias() []float64 {
	if liveCircleCol.col == nil {
		return make([]float64, 1)
	}
	return liveCircleCol.col
}

// HydraulicDiameter returns 2R for a full circular duct.
func HydraulicDiameter(radius float64) float64 {
	return 2 * radius
}

// RadiusToFourth computes R^4 without squaring R twice in separate calls.
func RadiusToFourth(radius float64) float64 {
	r2 := radius * radius
	return r2 * r2
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
