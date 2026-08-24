package reynolds

import (
	"math"
	"testing"

	"poiseuille-q/internal/model"
)

func TestReynoldsKnownValue(t *testing.T) {
	q := 7.6699e-9
	rho := 1000.0
	radius := 2.5e-4
	mu := 0.002
	got := ReynoldsNumber(rho, q, radius, mu)
	want := 2 * rho * q / (math.Pi * radius * mu)
	if math.Abs(got-want) > 1e-9 {
		t.Errorf("Re = %g, want %g", got, want)
	}
}

func TestGateRejectsTurbulent(t *testing.T) {
	ok, err := Gate(2500)
	if ok {
		t.Errorf("laminar = true, want false for Re 2500")
	}
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if !model.IsCode(err, model.CodeNotLaminar) {
		t.Errorf("expected code %q, got %q", model.CodeNotLaminar, model.Code(err))
	}
}

func TestGateAcceptsSubcritical(t *testing.T) {
	ok, err := Gate(2000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !ok {
		t.Errorf("laminar = false, want true")
	}
}

func TestCriticalFlowAtLimit(t *testing.T) {
	q := CriticalVolumetricFlow(1000, 0.01, 0.001)
	re := ReynoldsNumber(1000, q, 0.01, 0.001)
	if math.Abs(re-LaminarLimit) > 1e-9 {
		t.Errorf("critical Re = %g, want %g", re, LaminarLimit)
	}
}
