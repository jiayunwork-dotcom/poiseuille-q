package delta

import (
	"math"

	"poiseuille-q/internal/model"
	"poiseuille-q/internal/reynolds"
	"poiseuille-q/internal/shear"
	"poiseuille-q/internal/velocity"
)

// ResolveTargetPressure validates and gates a requested operating point.
func ResolveTargetPressure(in model.Input, targetQ float64) (model.DeltaResult, error) {
	dp, err := RequiredWithLaminarGate(in, targetQ)
	if err != nil {
		return model.DeltaResult{}, err
	}
	re := reynolds.ReynoldsNumber(in.Rho, targetQ, in.Radius, in.Mu)
	laminar, gateErr := reynolds.Gate(re)
	if gateErr != nil {
		return model.DeltaResult{}, gateErr
	}
	uAvg := velocity.AverageVelocity(targetQ, in.Radius)
	uMax := velocity.CenterlineVelocity(uAvg)
	tau := shear.WallShear(dp, in.Radius, in.Length)
	return model.DeltaResult{
		RequiredDeltaP: dp,
		ResultingQ:     targetQ,
		Re:             re,
		TauW:           tau,
		UAvg:           uAvg,
		UMax:           uMax,
		Laminar:        laminar,
	}, nil
}

// CriticalHeadroom returns the fraction of the laminar limit consumed.
func CriticalHeadroom(re float64) float64 {
	if re >= reynolds.LaminarLimit {
		return 1
	}
	return re / reynolds.LaminarLimit
}

// IsInsideLimit returns true when Re is strictly below the threshold.
func IsInsideLimit(re float64) bool {
	return re < reynolds.LaminarLimit
}

// CapTargetFlow lowers an impossible request to the critical value.
func CapTargetFlow(rho, radius, mu float64) float64 {
	return reynolds.CriticalVolumetricFlow(rho, radius, mu)
}

// SignOfTarget normalizes a target flow sign.
func SignOfTarget(targetQ float64) float64 {
	if targetQ < 0 {
		return -1
	}
	if targetQ > 0 {
		return 1
	}
	return 0
}

// WithinTolerance compares two floats with a relative scale.
func WithinTolerance(a, b float64) bool {
	diff := math.Abs(a - b)
	scale := math.Max(math.Abs(a), math.Abs(b))
	return diff <= math.Max(1e-12, scale*1e-9)
}

// ModelError wraps a gate failure into a structured error.
func ModelError(code model.ErrorCode, message string) error {
	return model.NewError(code, message)
}
