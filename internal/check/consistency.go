package check

import (
	"math"

	"poiseuille-q/internal/model"
)

// ConsistencyReport summarizes how many public invariants hold.
type ConsistencyReport struct {
	Total     int
	PassCount int
	Reasons   []string
}

// CheckResult checks the cross-rule contract of a complete result.
func CheckResult(in model.Input, r model.Result) ConsistencyReport {
	report := ConsistencyReport{Total: 4}
	if ZeroPressureInvariant(r.Q) && in.IsZeroPressure() {
		report.PassCount++
	} else if !in.IsZeroPressure() {
		report.PassCount++
	} else {
		report.Reasons = append(report.Reasons, "zero pressure did not produce zero flow")
	}
	if DirectionInvariant(r.Q, in.DeltaP) {
		report.PassCount++
	} else {
		report.Reasons = append(report.Reasons, "flow direction does not follow pressure sign")
	}
	if r.Laminar {
		report.PassCount++
	} else {
		report.Reasons = append(report.Reasons, "turbulent case was not rejected")
	}
	if math.Abs(r.UMax-2*r.UAvg) <= 1e-9*math.Max(1, math.Abs(r.UMax)) {
		report.PassCount++
	} else {
		report.Reasons = append(report.Reasons, "centerline velocity is not twice the average")
	}
	return report
}

// Passed reports whether every invariant held.
func (c ConsistencyReport) Passed() bool {
	return c.PassCount == c.Total
}

// Merge combines two consistency reports.
func Merge(a, b ConsistencyReport) ConsistencyReport {
	return ConsistencyReport{
		Total:     a.Total + b.Total,
		PassCount: a.PassCount + b.PassCount,
		Reasons:   append(append([]string{}, a.Reasons...), b.Reasons...),
	}
}
