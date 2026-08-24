package shear

import (
	"math"

	"poiseuille-q/internal/geom"
)

// WallShear is the pressure-based wall shear stress R*dP/(2L).
func WallShear(deltaP, radius, length float64) float64 {
	if length <= 0 {
		return 0
	}
	return radius * deltaP / (2 * length)
}

// ShearRateAtWall converts Q to the wall shear rate for a Newtonian fluid.
func ShearRateAtWall(q, radius float64) float64 {
	if radius <= 0 {
		return 0
	}
	return 4 * q / (math.Pi * math.Pow(radius, 3))
}

// ViscousShearStress returns mu*du/dr at the wall using the same formula set.
func ViscousShearStress(q, radius, mu float64) float64 {
	return mu * ShearRateAtWall(q, radius)
}

// FrictionVelocity computes u_tau from wall shear and density.
func FrictionVelocity(tau, rho float64) float64 {
	if rho <= 0 {
		return 0
	}
	return math.Sqrt(math.Abs(tau) / rho)
}

// FrictionFactorLaminar returns 64/Re for a circular laminar tube.
func FrictionFactorLaminar(re float64) float64 {
	if re == 0 {
		return 0
	}
	return 64 / re
}

// PressureDropFromShear inverts the wall-shear formula.
func PressureDropFromShear(tau, radius, length float64) float64 {
	if radius <= 0 {
		return 0
	}
	return 2 * tau * length / radius
}

// WallShearFromFlow derives tau from Q using the Newtonian velocity gradient.
func WallShearFromFlow(q, radius, mu float64) float64 {
	return ViscousShearStress(q, radius, mu)
}

// ReynoldsFromFriction combines 64/Re with the usual definitions.
func ReynoldsFromFriction(frictionFactor float64) float64 {
	if frictionFactor == 0 {
		return 0
	}
	return 64 / frictionFactor
}

// TubeAreaKept is a small helper for dimensional audits.
func TubeAreaKept(radius float64) float64 {
	return geom.CircleArea(radius)
}
