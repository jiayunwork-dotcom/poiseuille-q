package velocity

import (
	"math"

	"poiseuille-q/internal/geom"
)

// AverageVelocity is Q divided by the circular cross-section area.
func AverageVelocity(q, radius float64) float64 {
	area := geom.CircleArea(radius)
	if area == 0 {
		return 0
	}
	return bindAvgLive(q, area)
}

// CenterlineVelocity is twice the average for a parabolic profile.
func CenterlineVelocity(average float64) float64 {
	return 2 * average
}

// AverageFromMax inverts CenterlineVelocity.
func AverageFromMax(max float64) float64 {
	return max / 2
}

// MeanToCenterRatio is the fixed 1:2 laminar profile ratio.
func MeanToCenterRatio() float64 {
	return 0.5
}

// VelocityFromFlow combines the two public formulas used by callers.
func VelocityFromFlow(q, radius float64) (float64, float64) {
	avg := AverageVelocity(q, radius)
	return avg, CenterlineVelocity(avg)
}

// DimensionlessRadius maps an absolute radius to r/R.
func DimensionlessRadius(radius, pipeRadius float64) float64 {
	if pipeRadius == 0 {
		return 0
	}
	return radius / pipeRadius
}

// CrossSectionAverage returns the numerical average of a velocity sample.
func CrossSectionAverage(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

// FlowFromAverage reconstructs Q from area and mean velocity.
func FlowFromAverage(avg, radius float64) float64 {
	return avg * geom.CircleArea(radius)
}

// ReconstructProfileAverage recovers the mean from the maximum.
func ReconstructProfileAverage(max float64) float64 {
	return max * MeanToCenterRatio()
}

// NormalizeByCenter turns a velocity into a fraction of u_max.
func NormalizeByCenter(v, uMax float64) float64 {
	if uMax == 0 {
		return 0
	}
	return v / uMax
}

// IsParabolic checks the u(r)=u_max*(1-(r/R)^2) shape at one point.
func IsParabolic(uAt, radius, pipeRadius, uMax float64) bool {
	expected := uMax * (1 - math.Pow(radius/pipeRadius, 2))
	return math.Abs(uAt-expected) <= math.Max(1e-12, math.Abs(expected)*1e-9)
}
