package velocity

import (
	"math"
	"testing"
)

func TestCenterlineTwiceAverage(t *testing.T) {
	avg := AverageVelocity(7.6699e-9, 2.5e-4)
	max := CenterlineVelocity(avg)
	if math.Abs(max-2*avg) > 1e-18 {
		t.Errorf("u_max = %g, want %g", max, 2*avg)
	}
}

func TestAverageVelocityFromArea(t *testing.T) {
	got := AverageVelocity(7.6699e-9, 2.5e-4)
	want := 7.6699e-9 / (math.Pi * 6.25e-8)
	if math.Abs(got-want) > 1e-12 {
		t.Errorf("u_avg = %g, want %g", got, want)
	}
}

func TestParabolicProfileAreaWeightedMean(t *testing.T) {
	uMax := 0.1
	mean := AreaWeightedAverage(0.001, uMax, 2000)
	want := uMax / 2
	if math.Abs(mean-want) > 1e-4 {
		t.Errorf("area-weighted mean = %g, want %g", mean, want)
	}
}

func TestVelocityAtWallZero(t *testing.T) {
	got := VelocityAt(0.001, 0.001, 0.1)
	if got != 0 {
		t.Errorf("wall velocity = %g, want 0", got)
	}
}
