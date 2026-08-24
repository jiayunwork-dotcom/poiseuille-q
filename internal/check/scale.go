package check

// ScaleCase pairs a scale multiplier with the expected flow multiplier.
type ScaleCase struct {
	InputMultiplier  float64
	OutputMultiplier float64
}

// RadiusFourthCases returns the R-scale cases used by public checks.
func RadiusFourthCases() []ScaleCase {
	return []ScaleCase{
		{InputMultiplier: 1, OutputMultiplier: 1},
		{InputMultiplier: 2, OutputMultiplier: 16},
		{InputMultiplier: 0.5, OutputMultiplier: 0.0625},
	}
}

// ViscosityCases returns the mu-scale cases used by public checks.
func ViscosityCases() []ScaleCase {
	return []ScaleCase{
		{InputMultiplier: 1, OutputMultiplier: 1},
		{InputMultiplier: 2, OutputMultiplier: 0.5},
		{InputMultiplier: 0.5, OutputMultiplier: 2},
	}
}

// LengthCases returns the L-scale cases used by public checks.
func LengthCases() []ScaleCase {
	return []ScaleCase{
		{InputMultiplier: 1, OutputMultiplier: 1},
		{InputMultiplier: 2, OutputMultiplier: 0.5},
		{InputMultiplier: 0.5, OutputMultiplier: 2},
	}
}

// MatchScale verifies a measured output multiplier against the rule.
func MatchScale(c ScaleCase, actual float64) bool {
	return CloseEnough(c.OutputMultiplier, actual)
}

// RelativeError returns the unsigned relative difference.
func RelativeError(actual, expected float64) float64 {
	if expected == 0 {
		if actual == 0 {
			return 0
		}
		return 1
	}
	diff := actual - expected
	if diff < 0 {
		diff = -diff
	}
	return diff / abs(expected)
}

func abs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}

// WithinRelative reports whether a relative error is acceptable.
func WithinRelative(actual, expected, maxErr float64) bool {
	return RelativeError(actual, expected) <= maxErr
}
