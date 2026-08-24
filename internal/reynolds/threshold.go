package reynolds

import (
	"math"

	"poiseuille-q/internal/geom"
)

// CriticalAverageVelocity is the mean velocity at Re=LaminarLimit.
func CriticalAverageVelocity(rho, radius, mu float64) float64 {
	return LaminarLimit * mu / (2 * rho * radius)
}

// CriticalVolumetricFlow is Q at Re=LaminarLimit.
func CriticalVolumetricFlow(rho, radius, mu float64) float64 {
	return CriticalAverageVelocity(rho, radius, mu) * geom.CircleArea(radius)
}

// CriticalMassFlow multiplies the critical volumetric flow by density.
func CriticalMassFlow(rho, radius, mu float64) float64 {
	return rho * CriticalVolumetricFlow(rho, radius, mu)
}

// CriticalShearRate returns the wall shear rate at the boundary.
func CriticalShearRate(rho, radius, mu float64) float64 {
	q := CriticalVolumetricFlow(rho, radius, mu)
	return 4 * q / (math.Pi * math.Pow(radius, 3))
}

// RemainingHeadroom reports Q/Q_critical as a percentage.
func RemainingHeadroom(re float64) float64 {
	if re >= LaminarLimit {
		return 0
	}
	return (LaminarLimit - re) / LaminarLimit * 100
}

// FractionOfLimit normalizes the current Re by the laminar limit.
func FractionOfLimit(re float64) float64 {
	return re / LaminarLimit
}
