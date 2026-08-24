package flow

import (
	"math"

	"poiseuille-q/internal/model"
)

const flowTolerance = 1e-9

// CheckZeroDeltaP verifies the no-pressure boundary result.
func CheckZeroDeltaP(in model.Input) bool {
	if in.DeltaP != 0 {
		return false
	}
	return math.Abs(VolumetricFlow(in.Radius, in.Length, 0, in.Mu)) <= flowTolerance
}

// CheckRadiusFourthScaling verifies the R^4 rule for a two-radius pair.
func CheckRadiusFourthScaling(baseRadius, newRadius, baseQ, newQ float64) bool {
	expected := baseQ * RadiusFourthRatio(newRadius, baseRadius)
	return math.Abs(newQ-expected) <= math.Max(1e-12, math.Abs(expected)*flowTolerance)
}

// CheckViscosityHalving verifies Q halves when mu doubles.
func CheckViscosityHalving(baseMu, newMu, baseQ, newQ float64) bool {
	expected := baseQ * ViscosityRatio(newMu, baseMu)
	return math.Abs(newQ-expected) <= math.Max(1e-12, math.Abs(expected)*flowTolerance)
}

// CheckShearInvariant verifies wall shear does not depend on viscosity.
func CheckShearInvariant(tauBase, tauNew float64) bool {
	return math.Abs(tauBase-tauNew) <= math.Max(1e-12, math.Abs(tauBase)*flowTolerance)
}

// SameDirection compares two direction labels.
func SameDirection(a, b string) bool {
	return a == b
}
