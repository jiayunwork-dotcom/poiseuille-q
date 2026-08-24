package delta

import (
	"math"
	"testing"

	"poiseuille-q/internal/model"
)

func TestRequiredPressureDropExact(t *testing.T) {
	radius := 2.5e-4
	length := 0.1
	mu := 0.002
	q := 7.6699e-9
	got := RequiredPressureDrop(q, radius, length, mu)
	want := 8 * mu * length * q / (math.Pi * math.Pow(radius, 4))
	if math.Abs(got-want) > 1e-12 {
		t.Errorf("dP = %g, want %g", got, want)
	}
}

func TestRequiredPressureDropKeepsLaminarGate(t *testing.T) {
	in := model.NewInput(2.5e-4, 0.1, 0, 0.002, 1000)
	res, err := ResolveTargetPressure(in, 7.6699e-9)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !res.Laminar {
		t.Errorf("laminar = false, want true")
	}
	if res.RequiredDeltaP <= 0 {
		t.Errorf("dP = %g, want positive", res.RequiredDeltaP)
	}
}

func TestRequiredPressureDropRejectsTurbulentTarget(t *testing.T) {
	in := model.NewInput(0.01, 1, 0, 0.001, 1000)
	_, err := ResolveTargetPressure(in, 0.01)
	if err == nil {
		t.Fatalf("expected not_laminar error, got nil")
	}
	if !model.IsCode(err, model.CodeNotLaminar) {
		t.Errorf("expected code %q, got %q", model.CodeNotLaminar, model.Code(err))
	}
}

func TestBisectionMatchesAnalytic(t *testing.T) {
	in := model.NewInput(2.5e-4, 0.1, 0, 0.002, 1000)
	target := 7.6699e-9
	analytic := RequiredPressureDrop(target, in.Radius, in.Length, in.Mu)
	numeric, err := SolveDeltaPByBisection(in, target, 1e-12, 120)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if math.Abs(numeric-analytic) > 1e-6*math.Max(1, math.Abs(analytic)) {
		t.Errorf("numeric dP = %g, want %g", numeric, analytic)
	}
}

func TestZeroTargetFlowZeroPressure(t *testing.T) {
	in := model.NewInput(2.5e-4, 0.1, 0, 0.002, 1000)
	res, err := ResolveTargetPressure(in, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.RequiredDeltaP != 0 {
		t.Errorf("dP = %g, want 0", res.RequiredDeltaP)
	}
}
