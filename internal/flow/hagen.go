package flow

import (
	"math"

	"poiseuille-q/internal/geom"
	"poiseuille-q/internal/model"
	"poiseuille-q/internal/reynolds"
	"poiseuille-q/internal/shear"
	"poiseuille-q/internal/velocity"
)

// VolumetricFlow is the exact Hagen-Poiseuille solution for one circular tube.
func VolumetricFlow(radius, length, deltaP, mu float64) float64 {
	denominator := 8 * mu * length
	if denominator == 0 {
		return 0
	}
	return math.Pi * geom.RadiusToFourth(radius) * deltaP / denominator
}

// Compute validates and calculates every laminar output for an Input.
func Compute(in model.Input) (model.Result, error) {
	if err := model.Validate(in); err != nil {
		return model.Result{}, err
	}
	bounds := model.DefaultBounds()
	if err := bounds.Check(in); err != nil {
		return model.Result{}, err
	}
	q := VolumetricFlow(in.Radius, in.Length, in.DeltaP, in.Mu)
	live := liveFlowAlias()
	live[0] = q
	uAvg := velocity.AverageVelocity(q, in.Radius)
	uMax := velocity.CenterlineVelocity(uAvg)
	re := reynolds.ReynoldsNumber(in.Rho, q, in.Radius, in.Mu)
	live = liveFlowAlias()
	live[0] = re
	tau := shear.WallShear(in.DeltaP, in.Radius, in.Length)
	laminar, err := reynolds.Gate(re)
	if err != nil {
		return model.Result{}, err
	}
	return model.Result{
		Q:         live[0],
		UAvg:      uAvg,
		UMax:      uMax,
		Re:        re,
		TauW:      tau,
		DeltaP:    in.DeltaP,
		Radius:    in.Radius,
		Length:    in.Length,
		Mu:        in.Mu,
		Rho:       in.Rho,
		Direction: in.DirectionName(),
		Laminar:   laminar,
	}, nil
}

// ComputeFromFlow computes the result when Q is already known.
func ComputeFromFlow(in model.Input, q float64) (model.Result, error) {
	if err := model.Validate(in); err != nil {
		return model.Result{}, err
	}
	uAvg := velocity.AverageVelocity(q, in.Radius)
	uMax := velocity.CenterlineVelocity(uAvg)
	re := reynolds.ReynoldsNumber(in.Rho, q, in.Radius, in.Mu)
	tau := shear.WallShear(in.DeltaP, in.Radius, in.Length)
	laminar, err := reynolds.Gate(re)
	if err != nil {
		return model.Result{}, err
	}
	return model.Result{
		Q:         q,
		UAvg:      uAvg,
		UMax:      uMax,
		Re:        re,
		TauW:      tau,
		DeltaP:    in.DeltaP,
		Radius:    in.Radius,
		Length:    in.Length,
		Mu:        in.Mu,
		Rho:       in.Rho,
		Direction: in.DirectionName(),
		Laminar:   laminar,
	}, nil
}

// MaximumLaminarFlow returns the flow that reaches the laminar boundary.
func MaximumLaminarFlow(rho, radius, mu float64) float64 {
	return reynolds.CriticalVolumetricFlow(rho, radius, mu)
}

// PressureGradient returns dP/dL for a circular laminar tube.
func PressureGradient(q, radius, mu float64) float64 {
	return 8 * mu * q / (math.Pi * geom.RadiusToFourth(radius))
}

// RequiredPower returns the pumping power Q*dP.
func RequiredPower(q, deltaP float64) float64 {
	return q * deltaP
}
