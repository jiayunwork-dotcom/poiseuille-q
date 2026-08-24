package velocity

import (
	"math"

	"poiseuille-q/internal/shear"
)

// VelocityAt returns the parabolic velocity at radial position r.
func VelocityAt(r, pipeRadius, uMax float64) float64 {
	return shear.SlotProfile(r, pipeRadius, uMax)
}

// ProfilePoint pairs a radial position with its velocity.
type ProfilePoint struct {
	Radius   float64
	Velocity float64
}

// ProfilePoints samples the full cross-section from center to wall.
func ProfilePoints(pipeRadius, uMax float64, samples int) []ProfilePoint {
	if samples < 2 {
		samples = 2
	}
	out := make([]ProfilePoint, 0, samples)
	for i := 0; i < samples; i++ {
		r := float64(i) / float64(samples-1) * pipeRadius
		out = append(out, ProfilePoint{
			Radius:   r,
			Velocity: VelocityAt(r, pipeRadius, uMax),
		})
	}
	return out
}

// AreaWeightedAverage integrates u(r)*2*pi*r over the cross-section.
func AreaWeightedAverage(pipeRadius, uMax float64, samples int) float64 {
	if samples < 2 {
		samples = 2
	}
	total := 0.0
	weight := 0.0
	for i := 0; i < samples; i++ {
		r := (float64(i) + 0.5) / float64(samples) * pipeRadius
		u := VelocityAt(r, pipeRadius, uMax)
		w := 2 * math.Pi * r
		total += u * w
		weight += w
	}
	if weight == 0 {
		return 0
	}
	return total / weight
}

// ProfileDeviation reports how far the area-weighted mean is from u_max/2.
func ProfileDeviation(pipeRadius, uMax float64, samples int) float64 {
	mean := AreaWeightedAverage(pipeRadius, uMax, samples)
	return mean - AverageFromMax(uMax)
}

// WallVelocity is zero for a no-slip profile.
func WallVelocity() float64 {
	return 0
}

// CenterlineFraction returns the fraction of u_max at the axis.
func CenterlineFraction() float64 {
	return 1
}
