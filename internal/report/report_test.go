package report

import (
	"math"
	"strconv"
	"testing"

	"poiseuille-q/internal/model"
)

func TestBuildFlowSummaryFields(t *testing.T) {
	res := model.Result{
		Q:         1e-8,
		UAvg:      0.05,
		UMax:      0.1,
		Re:        500,
		TauW:      1.25,
		DeltaP:    1000,
		Direction: "forward",
		Laminar:   true,
	}
	s := BuildFlow(res)
	if s.Q != 1e-8 || s.UAvg != 0.05 || s.UMax != 0.1 {
		t.Errorf("summary fields wrong: %+v", s)
	}
	if s.Re != 500 || s.TauW != 1.25 || !s.Laminar {
		t.Errorf("summary numeric fields wrong: %+v", s)
	}
}

func TestBuildDeltaSummaryFields(t *testing.T) {
	d := model.DeltaResult{
		RequiredDeltaP: 1000,
		ResultingQ:     1e-8,
		Re:             500,
		TauW:           1.25,
		UAvg:           0.05,
		UMax:           0.1,
		Laminar:        true,
	}
	s := BuildDelta(d)
	if s.RequiredDeltaP != 1000 || s.ResultingQ != 1e-8 {
		t.Errorf("delta summary fields wrong: %+v", s)
	}
	if s.Re != 500 || s.TauW != 1.25 {
		t.Errorf("delta summary numeric fields wrong: %+v", s)
	}
}

func TestFormatNumber(t *testing.T) {
	if got := FormatNumber(1.0 / 3); math.Abs(gotDiff(got)-1.0/3) > 1e-8 {
		t.Errorf("format %q did not preserve value", got)
	}
}

func gotDiff(s string) float64 {
	var v float64
	v, _ = strconv.ParseFloat(s, 64)
	return v
}
