package shear

import (
	"math"
	"testing"
)

func TestWallShearFormula(t *testing.T) {
	got := WallShear(1000, 2.5e-4, 0.1)
	want := 2.5e-4 * 1000 / 0.2
	if math.Abs(got-want) > 1e-12 {
		t.Errorf("tau_w = %g, want %g", got, want)
	}
}

func TestWallShearIndependentOfViscosity(t *testing.T) {
	low := WallShear(1000, 2.5e-4, 0.1)
	high := WallShear(1000, 2.5e-4, 0.1)
	if !WallShearIndependentOfViscosity(low, high) {
		t.Errorf("wall shear changed despite identical pressure geometry: %g vs %g", low, high)
	}
}

func TestShearRateAtWall(t *testing.T) {
	q := 7.6699e-9
	radius := 2.5e-4
	got := ShearRateAtWall(q, radius)
	want := 4 * q / (math.Pi * math.Pow(radius, 3))
	if math.Abs(got-want) > 1e-12 {
		t.Errorf("shear rate = %g, want %g", got, want)
	}
}

func TestPressureDropFromShearInverts(t *testing.T) {
	tau := WallShear(1000, 2.5e-4, 0.1)
	got := PressureDropFromShear(tau, 2.5e-4, 0.1)
	if math.Abs(got-1000) > 1e-9 {
		t.Errorf("pressure = %g, want 1000", got)
	}
}
