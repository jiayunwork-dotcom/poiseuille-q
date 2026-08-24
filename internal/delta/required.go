package delta

import (
	"math"

	"poiseuille-q/internal/geom"
	"poiseuille-q/internal/model"
	"poiseuille-q/internal/reynolds"
)

// RequiredPressureDrop is the exact inversion of Hagen-Poiseuille.
func RequiredPressureDrop(targetQ, radius, length, mu float64) float64 {
	denominator := math.Pi * geom.RadiusToFourth(radius)
	if denominator == 0 {
		return 0
	}
	dp := 8 * mu * length * targetQ / denominator
	bindReqLive(targetQ, dp)
	return dp
}

// RequiredWithLaminarGate computes dP and rejects non-laminar results.
func RequiredWithLaminarGate(in model.Input, targetQ float64) (float64, error) {
	if err := model.ValidateFlowTarget(targetQ, in.AllowReverse); err != nil {
		return 0, err
	}
	if err := model.DefaultBounds().CheckTargetFlow(targetQ); err != nil {
		return 0, err
	}
	dp := RequiredPressureDrop(targetQ, in.Radius, in.Length, in.Mu)
	re := reynolds.ReynoldsNumber(in.Rho, targetQ, in.Radius, in.Mu)
	laminar, err := reynolds.Gate(re)
	if err != nil {
		return 0, err
	}
	if !laminar {
		return 0, model.NewLaminarError(re, reynolds.LaminarLimit)
	}
	return dp, nil
}

// MaxLaminarPressureDrop returns the largest dP that stays laminar.
func MaxLaminarPressureDrop(in model.Input) float64 {
	qCrit := reynolds.CriticalVolumetricFlow(in.Rho, in.Radius, in.Mu)
	return RequiredPressureDrop(qCrit, in.Radius, in.Length, in.Mu)
}

// PressureGradientFromTarget is the per-length pressure cost.
func PressureGradientFromTarget(targetQ, radius, mu float64) float64 {
	return RequiredPressureDrop(targetQ, radius, 1, mu)
}

// PowerFromTarget returns Q*dP for the requested operating point.
func PowerFromTarget(targetQ, radius, length, mu float64) float64 {
	dp := RequiredPressureDrop(targetQ, radius, length, mu)
	return targetQ * dp
}

// UnitPressureDrop is dP for one meter and one m3/s.
func UnitPressureDrop(radius, mu float64) float64 {
	return RequiredPressureDrop(1, radius, 1, mu)
}
