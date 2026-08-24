package check

import (
	"math"

	"poiseuille-q/internal/model"
)

const tolerance = 1e-9

// RadiusFourthInvariant asserts Q scales with R^4.
func RadiusFourthInvariant(qBase, qNew float64, radiusMultiplier float64) bool {
	expected := qBase * math.Pow(radiusMultiplier, 4)
	return closeEnough(qNew, expected)
}

// ViscosityInvariant asserts Q is inversely proportional to mu.
func ViscosityInvariant(qBase, qNew float64, muMultiplier float64) bool {
	expected := qBase / muMultiplier
	return closeEnough(qNew, expected)
}

// ZeroPressureInvariant asserts zero dP gives zero flow.
func ZeroPressureInvariant(q float64) bool {
	return math.Abs(q) <= 1e-12
}

// ShearInvariant asserts wall shear does not depend on mu.
func ShearInvariant(tauBase, tauNew float64) bool {
	return closeEnough(tauBase, tauNew)
}

// DirectionInvariant verifies the sign of flow matches the sign of dP.
func DirectionInvariant(q, deltaP float64) bool {
	if deltaP == 0 {
		return q == 0
	}
	return (q > 0 && deltaP > 0) || (q < 0 && deltaP < 0)
}

// LaminarInvariant verifies the gate boundary is respected.
func LaminarInvariant(re, limit float64) bool {
	return re < limit
}

// CloseEnough is exported for cross-package checks.
func CloseEnough(a, b float64) bool {
	return closeEnough(a, b)
}

func closeEnough(a, b float64) bool {
	scale := math.Max(math.Abs(a), math.Abs(b))
	return math.Abs(a-b) <= math.Max(1e-12, scale*tolerance)
}

// NonNegative reports that a quantity cannot physically go below zero.
func NonNegative(v float64) bool {
	return v >= 0
}

// StrictlyPositive reports a strictly positive physical quantity.
func StrictlyPositive(v float64) bool {
	return v > 0
}

// InputValid wraps the model validation for check callers.
func InputValid(in model.Input) bool {
	return model.Validate(in) == nil
}
