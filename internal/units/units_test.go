package units

import (
	"math"
	"testing"

	"poiseuille-q/internal/model"
)

func TestToSIFromLab(t *testing.T) {
	in := model.NewInput(500, 250, 100, 2, 1000)
	converted := ConvertInput(in, Laboratory)
	if math.Abs(converted.Radius-0.5) > 1e-12 {
		t.Errorf("radius = %g, want 0.5", converted.Radius)
	}
	if math.Abs(converted.Length-0.25) > 1e-12 {
		t.Errorf("length = %g, want 0.25", converted.Length)
	}
	if math.Abs(converted.DeltaP-1e5) > 1e-6 {
		t.Errorf("delta_p = %g, want 100000", converted.DeltaP)
	}
	if math.Abs(converted.Mu-0.002) > 1e-12 {
		t.Errorf("mu = %g, want 0.002", converted.Mu)
	}
}

func TestDynamicFromKinematic(t *testing.T) {
	got := DynamicFromKinematic(1e-6, 1000)
	if math.Abs(got-0.001) > 1e-12 {
		t.Errorf("dynamic viscosity = %g, want 0.001", got)
	}
}

func TestPressureHead(t *testing.T) {
	got := PressureHead(98066.5, 1000, 9.80665)
	if math.Abs(got-10) > 1e-9 {
		t.Errorf("head = %g, want 10", got)
	}
}

func TestParseUnitSystem(t *testing.T) {
	if _, err := ParseUnitSystem("lab"); err != nil {
		t.Errorf("unexpected error for lab: %v", err)
	}
	if _, err := ParseUnitSystem("bogus"); err == nil {
		t.Errorf("expected error for bogus unit system")
	}
}
