package flow

import (
	"math"
	"testing"

	"poiseuille-q/internal/model"
)

func TestHagenFlowExactFormula(t *testing.T) {
	radius := 2.5e-4
	length := 0.1
	dp := 1000.0
	mu := 0.002
	got := VolumetricFlow(radius, length, dp, mu)
	want := math.Pi * math.Pow(radius, 4) * dp / (8 * mu * length)
	if math.Abs(got-want) > 1e-18 {
		t.Errorf("Q = %.12g, want %.12g", got, want)
	}
}

func TestZeroDeltaPZeroFlow(t *testing.T) {
	in := model.NewInput(2.5e-4, 0.1, 0, 0.002, 1000)
	res, err := Compute(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Q != 0 {
		t.Errorf("Q = %g, want 0", res.Q)
	}
	if res.UAvg != 0 || res.UMax != 0 {
		t.Errorf("velocities = %g/%g, want 0/0", res.UAvg, res.UMax)
	}
}

func TestViscosityDoublingHalvesFlow(t *testing.T) {
	base := model.NewInput(2.5e-4, 0.1, 1000, 0.002, 1000)
	double := base
	double.Mu = 0.004
	baseRes, err := Compute(base)
	if err != nil {
		t.Fatalf("base compute failed: %v", err)
	}
	doubleRes, err := Compute(double)
	if err != nil {
		t.Fatalf("double compute failed: %v", err)
	}
	want := baseRes.Q / 2
	if math.Abs(doubleRes.Q-want) > 1e-15*math.Max(1, math.Abs(want)) {
		t.Errorf("Q = %g, want %g", doubleRes.Q, want)
	}
}

func TestRadiusDoublingScalesFlowSixteen(t *testing.T) {
	base := model.NewInput(2.5e-4, 0.1, 1000, 0.002, 1000)
	larger := base
	larger.Radius = 5e-4
	baseRes, err := Compute(base)
	if err != nil {
		t.Fatalf("base compute failed: %v", err)
	}
	largerRes, err := Compute(larger)
	if err != nil {
		t.Fatalf("larger compute failed: %v", err)
	}
	want := baseRes.Q * 16
	if math.Abs(largerRes.Q-want) > 1e-12*math.Max(1, math.Abs(want)) {
		t.Errorf("Q = %g, want %g (R^4 rule)", largerRes.Q, want)
	}
}

func TestComputeRejectsTurbulent(t *testing.T) {
	in := model.NewInput(0.01, 1, 1000, 0.001, 1000)
	_, err := Compute(in)
	if err == nil {
		t.Fatalf("expected not_laminar error, got nil")
	}
	if !model.IsCode(err, model.CodeNotLaminar) {
		t.Errorf("expected code %q, got %q", model.CodeNotLaminar, model.Code(err))
	}
}

func TestComputeReverseDirection(t *testing.T) {
	in := model.NewInput(2.5e-4, 0.1, -1000, 0.002, 1000).WithReverse(true)
	res, err := Compute(in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Q >= 0 {
		t.Errorf("Q = %g, want negative for reverse flow", res.Q)
	}
	if res.Direction != "reverse" {
		t.Errorf("direction = %q, want reverse", res.Direction)
	}
}
