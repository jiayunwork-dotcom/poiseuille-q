package geom

import (
	"math"
	"testing"
)

func TestCircleAreaKnownValue(t *testing.T) {
	got := CircleArea(0.001)
	want := math.Pi * 1e-6
	if math.Abs(got-want) > 1e-15 {
		t.Errorf("area = %.12g, want %.12g", got, want)
	}
}

func TestRadiusToFourthMatchesPower(t *testing.T) {
	got := RadiusToFourth(2.5e-4)
	want := math.Pow(2.5e-4, 4)
	if math.Abs(got-want) > 1e-24 {
		t.Errorf("R^4 = %.12g, want %.12g", got, want)
	}
}

func TestHydraulicDiameter(t *testing.T) {
	got := HydraulicDiameter(0.003)
	if got != 0.006 {
		t.Errorf("hydraulic diameter = %g, want 0.006", got)
	}
}

func TestTubeVolume(t *testing.T) {
	got := TubeVolume(0.001, 0.2)
	want := math.Pi * 1e-6 * 0.2
	if math.Abs(got-want) > 1e-12 {
		t.Errorf("volume = %g, want %g", got, want)
	}
}
