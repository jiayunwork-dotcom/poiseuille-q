package units

import (
	"fmt"

	"poiseuille-q/internal/model"
)

// UnitSystem names an accepted input convention.
type UnitSystem string

const (
	SI         UnitSystem = "si"
	Laboratory UnitSystem = "lab"
)

// ConvertInput converts a laboratory input into SI model.Input.
func ConvertInput(in model.Input, system UnitSystem) model.Input {
	switch system {
	case Laboratory:
		in.Radius = ToMeters(in.Radius)
		in.Length = ToMeters(in.Length)
		in.DeltaP = ToPascals(in.DeltaP)
		in.Mu = ToPascalSeconds(in.Mu)
	}
	return in
}

// ParseUnitSystem accepts "si" or "lab".
func ParseUnitSystem(s string) (UnitSystem, error) {
	switch s {
	case "si", "SI", "lab", "LAB":
		return UnitSystem(s), nil
	default:
		return "", fmt.Errorf("unknown unit system %q", s)
	}
}

// NormalizedSystem returns a canonical UnitSystem value.
func NormalizedSystem(s string) UnitSystem {
	if s == "lab" || s == "LAB" {
		return Laboratory
	}
	return SI
}

// InputLabels lists the five physical field labels for validation.
func InputLabels() []string {
	return []string{"radius", "length", "delta_p", "mu", "rho"}
}

// HasValidSystem reports whether a system string is recognized.
func HasValidSystem(s string) bool {
	_, err := ParseUnitSystem(s)
	return err == nil
}

// SafeDivide returns zero for a zero denominator.
func SafeDivide(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

// BoundsMessage produces a human-readable operating envelope.
func BoundsMessage(b model.Bounds) string {
	return fmt.Sprintf("R<=%g L<=%g dP<=%g mu<=%g rho<=%g",
		b.MaxRadius, b.MaxLength, b.MaxDeltaP, b.MaxMu, b.MaxRho)
}
