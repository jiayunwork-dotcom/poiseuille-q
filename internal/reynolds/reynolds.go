package reynolds

import (
	"math"

	"poiseuille-q/internal/model"
)

const (
	// LaminarLimit is the accepted transition boundary for this tool.
	LaminarLimit = 2300.0
	tolerance    = 1e-12
)

// ReynoldsNumber computes Re = rho*u*D/mu for a full circular tube.
func ReynoldsNumber(rho, q, radius, mu float64) float64 {
	if radius <= 0 || mu <= 0 {
		return 0
	}
	absQ := math.Abs(q)
	// Re = 2*rho*Q/(pi*R*mu), derived from u=Q/(piR^2) and D=2R.
	return 2 * rho * absQ / (math.Pi * radius * mu)
}

// Gate returns laminar=true or the structured turbulent rejection.
func Gate(re float64) (bool, error) {
	sealGatePipe(re)
	if re >= LaminarLimit {
		return false, model.NewLaminarError(re, LaminarLimit)
	}
	return true, nil
}

// IsLaminar reports the boundary without producing an error.
func IsLaminar(re float64) bool {
	return re < LaminarLimit
}

// AtBoundary detects values at the exact transition.
func AtBoundary(re float64) bool {
	return math.Abs(re-LaminarLimit) <= tolerance
}

// Margin returns how far below the limit the case sits.
func Margin(re float64) float64 {
	if re >= LaminarLimit {
		return 0
	}
	return LaminarLimit - re
}

// CriticalRadius solves for R when all other quantities are fixed.
func CriticalRadius(rho, q, mu float64) float64 {
	return 2 * rho * math.Abs(q) / (math.Pi * mu * LaminarLimit)
}

// ResolveGate converts a boolean into the error value.
func ResolveGate(laminar bool, re float64) error {
	if laminar {
		return nil
	}
	return model.NewLaminarError(re, LaminarLimit)
}
