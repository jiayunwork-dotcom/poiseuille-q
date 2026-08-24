package model

import "testing"

func TestValidateRejectsZeroRadius(t *testing.T) {
	in := NewInput(0, 1, 100, 0.001, 1000)
	err := Validate(in)
	if err == nil {
		t.Fatalf("expected validation error, got nil")
	}
	if !IsCode(err, CodeInvalidRadius) {
		t.Errorf("expected code %q, got %q", CodeInvalidRadius, Code(err))
	}
}

func TestValidateRejectsNegativeDeltaPWithoutReverse(t *testing.T) {
	in := NewInput(0.001, 1, -50, 0.001, 1000)
	err := Validate(in)
	if err == nil {
		t.Fatalf("expected validation error, got nil")
	}
	if !IsCode(err, CodeNegativeDeltaP) {
		t.Errorf("expected code %q, got %q", CodeNegativeDeltaP, Code(err))
	}
}

func TestValidateAllowsReverseFlow(t *testing.T) {
	in := NewInput(0.001, 1, -50, 0.001, 1000).WithReverse(true)
	if err := Validate(in); err != nil {
		t.Fatalf("expected reverse case to pass, got %v", err)
	}
	if got := in.DirectionName(); got != "reverse" {
		t.Errorf("expected reverse direction, got %q", got)
	}
}

func TestDefaultBoundsRejectsHugeRadius(t *testing.T) {
	in := NewInput(100, 1, 100, 0.001, 1000)
	err := DefaultBounds().Check(in)
	if err == nil {
		t.Fatalf("expected bounds error, got nil")
	}
	if !IsCode(err, CodeOutOfRange) {
		t.Errorf("expected code %q, got %q", CodeOutOfRange, Code(err))
	}
}

func TestErrorCodeStable(t *testing.T) {
	err := NewLaminarError(2500, 2300)
	if !IsCode(err, CodeNotLaminar) {
		t.Errorf("expected not_laminar code, got %q", Code(err))
	}
}
