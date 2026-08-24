package model

// Bounds keeps every quantity inside a defensive operating envelope.
type Bounds struct {
	MinPositive   float64
	MaxRadius     float64
	MaxLength     float64
	MaxDeltaP     float64
	MaxMu         float64
	MaxRho        float64
	MaxTargetFlow float64
}

// DefaultBounds returns the standard defensive limits.
func DefaultBounds() Bounds {
	return Bounds{
		MinPositive:   1e-12,
		MaxRadius:     10,
		MaxLength:     1e5,
		MaxDeltaP:     1e10,
		MaxMu:         1e6,
		MaxRho:        1e7,
		MaxTargetFlow: 1e6,
	}
}

// Check returns the first bound violation, if any.
func (b Bounds) Check(in Input) error {
	if in.Radius > b.MaxRadius {
		return NewError(CodeOutOfRange, "radius exceeds the supported operating bound").
			WithField("radius", in.Radius)
	}
	if in.Length > b.MaxLength {
		return NewError(CodeOutOfRange, "length exceeds the supported operating bound").
			WithField("length", in.Length)
	}
	if in.Mu > b.MaxMu {
		return NewError(CodeOutOfRange, "viscosity exceeds the supported operating bound").
			WithField("mu", in.Mu)
	}
	if in.Rho > b.MaxRho {
		return NewError(CodeOutOfRange, "density exceeds the supported operating bound").
			WithField("rho", in.Rho)
	}
	if in.DeltaP > b.MaxDeltaP {
		return NewError(CodeOutOfRange, "pressure difference exceeds the supported operating bound").
			WithField("delta_p", in.DeltaP)
	}
	return nil
}

// CheckTargetFlow validates the requested flow against the envelope.
func (b Bounds) CheckTargetFlow(targetQ float64) error {
	abs := targetQ
	if abs < 0 {
		abs = -abs
	}
	if abs > b.MaxTargetFlow {
		return NewError(CodeInvalidTargetFlow, "target flow exceeds the supported operating bound").
			WithField("target_q", targetQ)
	}
	return nil
}
