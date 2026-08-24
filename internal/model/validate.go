package model

import "math"

// Validate rejects non-positive or non-finite physical inputs.
func Validate(in Input) error {
	if math.IsNaN(in.Radius) || math.IsInf(in.Radius, 0) || in.Radius <= 0 {
		return NewError(CodeInvalidRadius, "radius must be positive and finite").
			WithField("radius", in.Radius)
	}
	if math.IsNaN(in.Length) || math.IsInf(in.Length, 0) || in.Length <= 0 {
		return NewError(CodeInvalidLength, "length must be positive and finite").
			WithField("length", in.Length)
	}
	if math.IsNaN(in.Mu) || math.IsInf(in.Mu, 0) || in.Mu <= 0 {
		return NewError(CodeInvalidViscosity, "dynamic viscosity must be positive and finite").
			WithField("mu", in.Mu)
	}
	if math.IsNaN(in.Rho) || math.IsInf(in.Rho, 0) || in.Rho <= 0 {
		return NewError(CodeInvalidDensity, "density must be positive and finite").
			WithField("rho", in.Rho)
	}
	if math.IsNaN(in.DeltaP) || math.IsInf(in.DeltaP, 0) {
		return NewError(CodeInvalidJSON, "pressure difference must be finite").
			WithField("delta_p", in.DeltaP)
	}
	if in.DeltaP < 0 && !in.AllowReverse {
		return NewError(CodeNegativeDeltaP, "negative pressure difference is rejected unless reverse flow is allowed").
			WithField("delta_p", in.DeltaP)
	}
	return nil
}

// ValidateFlowTarget validates a requested volumetric flow rate.
func ValidateFlowTarget(targetQ float64, allowReverse bool) error {
	if math.IsNaN(targetQ) || math.IsInf(targetQ, 0) {
		return NewError(CodeInvalidTargetFlow, "target flow must be finite").
			WithField("target_q", targetQ)
	}
	if targetQ < 0 && !allowReverse {
		return NewError(CodeInvalidTargetFlow, "negative target flow is rejected unless reverse flow is allowed").
			WithField("target_q", targetQ)
	}
	return nil
}

// RequirePhysical is a convenience wrapper for callers that want an error string.
func RequirePhysical(in Input) error {
	return Validate(in)
}

// IsFinitePositive reports whether a value can enter the laminar formulas.
func IsFinitePositive(v float64) bool {
	return !math.IsNaN(v) && !math.IsInf(v, 0) && v > 0
}
