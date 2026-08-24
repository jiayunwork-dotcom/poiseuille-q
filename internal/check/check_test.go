package check

import (
	"math"
	"testing"

	"poiseuille-q/internal/model"
)

func TestRadiusFourthInvariant(t *testing.T) {
	if !RadiusFourthInvariant(1e-8, 1.6e-7, 2) {
		t.Errorf("R^4 invariant failed for doubling radius")
	}
	if RadiusFourthInvariant(1e-8, 2e-8, 2) {
		t.Errorf("R^4 invariant accepted an R^2 result")
	}
}

func TestViscosityInvariant(t *testing.T) {
	if !ViscosityInvariant(1e-8, 5e-9, 2) {
		t.Errorf("viscosity invariant failed for doubling mu")
	}
}

func TestZeroPressureInvariant(t *testing.T) {
	if !ZeroPressureInvariant(0) {
		t.Errorf("zero pressure invariant rejected zero flow")
	}
	if ZeroPressureInvariant(1e-9) {
		t.Errorf("zero pressure invariant accepted nonzero flow")
	}
}

func TestShearInvariant(t *testing.T) {
	if !ShearInvariant(1.25, 1.25) {
		t.Errorf("shear invariant rejected equal values")
	}
	if ShearInvariant(1.25, 2.5) {
		t.Errorf("shear invariant accepted viscosity-dependent change")
	}
}

func TestDirectionInvariant(t *testing.T) {
	if !DirectionInvariant(-1e-8, -1000) {
		t.Errorf("direction invariant rejected negative pressure with negative flow")
	}
	if DirectionInvariant(-1e-8, 1000) {
		t.Errorf("direction invariant accepted sign mismatch")
	}
}

func TestConsistencyReport(t *testing.T) {
	in := model.NewInput(2.5e-4, 0.1, 1000, 0.002, 1000)
	res := model.Result{
		Q:         7.6699e-9,
		UAvg:      0.0391,
		UMax:      0.0782,
		Re:        9.78,
		TauW:      1.25,
		DeltaP:    1000,
		Direction: "forward",
		Laminar:   true,
	}
	rep := CheckResult(in, res)
	if !rep.Passed() {
		t.Errorf("consistency report failed: %+v", rep)
	}
	if rep.Total != 4 || rep.PassCount != 4 {
		t.Errorf("consistency counts = %d/%d, want 4/4", rep.PassCount, rep.Total)
	}
}

func TestRelativeError(t *testing.T) {
	if math.Abs(RelativeError(0.5, 1)-0.5) > 1e-12 {
		t.Errorf("relative error wrong")
	}
}
