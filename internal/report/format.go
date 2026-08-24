package report

import (
	"fmt"
	"math"
	"strings"
)

// FormatNumber prints a float with enough digits for engineering work.
func FormatNumber(v float64) string {
	if math.Abs(v) >= 1e5 || (math.Abs(v) < 1e-4 && v != 0) {
		return fmt.Sprintf("%.6e", v)
	}
	return fmt.Sprintf("%.8g", v)
}

// FormatCompact trims trailing zeros while keeping a useful precision.
func FormatCompact(v float64) string {
	s := fmt.Sprintf("%.6g", v)
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}

// FormatSigned always includes the direction sign.
func FormatSigned(v float64) string {
	if v >= 0 {
		return "+" + FormatCompact(v)
	}
	return FormatCompact(v)
}

// FormatPercent prints a percentage with one decimal.
func FormatPercent(v float64) string {
	return fmt.Sprintf("%.1f%%", v)
}

// FormatCount prints a dimensionless count with four significant digits.
func FormatCount(v float64) string {
	return fmt.Sprintf("%.4g", v)
}

// FormatRatio prints a ratio with four significant digits.
func FormatRatio(v float64) string {
	return fmt.Sprintf("%.4g", v)
}

// FormatScientific prints a compact scientific notation.
func FormatScientific(v float64) string {
	return fmt.Sprintf("%.4e", v)
}

// PadRight right-pads a label to a fixed width.
func PadRight(s string, width int) string {
	if len(s) >= width {
		return s
	}
	return s + strings.Repeat(" ", width-len(s))
}
