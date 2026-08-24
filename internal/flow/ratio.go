package flow

import "math"

// RadiusFourthRatio returns the flow multiplier when radius changes.
func RadiusFourthRatio(newRadius, oldRadius float64) float64 {
	return math.Pow(newRadius/oldRadius, 4)
}

// ViscosityRatio returns the flow multiplier when viscosity changes.
func ViscosityRatio(newMu, oldMu float64) float64 {
	return oldMu / newMu
}

// LengthRatio returns the flow multiplier when length changes.
func LengthRatio(newLength, oldLength float64) float64 {
	return oldLength / newLength
}

// DeltaPRatio returns the flow multiplier when pressure changes.
func DeltaPRatio(newDeltaP, oldDeltaP float64) float64 {
	return newDeltaP / oldDeltaP
}

// CombinedScaling multiplies the four independent geometric ratios.
func CombinedScaling(radiusRatio, muRatio, lengthRatio, deltaRatio float64) float64 {
	return radiusRatio * muRatio * lengthRatio * deltaRatio
}

// FlowScale computes Q2/Q1 from two complete inputs.
func FlowScale(newInput, oldInput interface {
	Radius() float64
	Length() float64
	DeltaP() float64
	Mu() float64
}) float64 {
	r := math.Pow(newInput.Radius()/oldInput.Radius(), 4)
	l := oldInput.Length() / newInput.Length()
	d := newInput.DeltaP() / oldInput.DeltaP()
	m := oldInput.Mu() / newInput.Mu()
	return r * l * d * m
}

// ScaleTo keeps the same driving pressure and multiplies only R and mu.
func ScaleTo(baseRadius, baseMu, radiusMultiplier, viscosityMultiplier float64) (float64, float64) {
	return baseRadius * radiusMultiplier, baseMu * viscosityMultiplier
}
