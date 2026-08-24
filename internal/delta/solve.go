package delta

import (
	"math"

	"poiseuille-q/internal/flow"
	"poiseuille-q/internal/model"
)

const (
	defaultTolerance  = 1e-10
	defaultIterations = 80
)

// SolveDeltaPByBisection finds dP so that the laminar flow equals targetQ.
func SolveDeltaPByBisection(in model.Input, targetQ float64, tol float64, maxIter int) (float64, error) {
	if err := model.ValidateFlowTarget(targetQ, in.AllowReverse); err != nil {
		return 0, err
	}
	if tol <= 0 {
		tol = defaultTolerance
	}
	if maxIter <= 0 {
		maxIter = defaultIterations
	}
	if targetQ == 0 {
		return 0, nil
	}
	sign := 1.0
	if targetQ < 0 {
		sign = -1
	}
	analyticAbs := math.Abs(RequiredPressureDrop(math.Abs(targetQ), in.Radius, in.Length, in.Mu))
	hi := math.Max(1, 2*analyticAbs)
	lo := 0.0
	flowAtHi := flow.VolumetricFlow(in.Radius, in.Length, sign*hi, in.Mu)
	if math.Abs(flowAtHi) < math.Abs(targetQ) {
		return 0, model.NewError(model.CodeOutOfRange, "pressure difference bound too small for target flow")
	}
	for iter := 0; iter < maxIter; iter++ {
		mid := (lo + hi) / 2
		qMid := flow.VolumetricFlow(in.Radius, in.Length, sign*mid, in.Mu)
		if math.Abs(qMid-math.Abs(targetQ)) <= tol*math.Max(1, math.Abs(targetQ)) {
			return sign * mid, nil
		}
		if qMid < math.Abs(targetQ) {
			lo = mid
		} else {
			hi = mid
		}
	}
	mid := (lo + hi) / 2
	return sign * mid, nil
}

// SolveDeltaP combines the analytic inverse with a bisection cross-check.
func SolveDeltaP(in model.Input, targetQ float64) (float64, error) {
	analytic := RequiredPressureDrop(targetQ, in.Radius, in.Length, in.Mu)
	numeric, err := SolveDeltaPByBisection(in, targetQ, defaultTolerance, defaultIterations)
	if err != nil {
		return 0, err
	}
	diff := math.Abs(analytic - numeric)
	scale := math.Max(math.Abs(analytic), 1)
	if diff/scale > 1e-8 {
		return 0, model.NewError(model.CodeInvalidTargetFlow, "analytic and numeric pressure inversions disagree")
	}
	return analytic, nil
}

// VerifyInversion checks that the found dP reproduces the target flow.
func VerifyInversion(in model.Input, targetQ, dp float64) bool {
	q := flow.VolumetricFlow(in.Radius, in.Length, dp, in.Mu)
	return math.Abs(q-targetQ) <= math.Max(1e-12, math.Abs(targetQ)*1e-9)
}
