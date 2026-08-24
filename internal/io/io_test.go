package io

import (
	"encoding/json"
	"strings"
	"testing"

	"poiseuille-q/internal/model"
)

func TestLoadFlowFileExample(t *testing.T) {
	in, err := LoadFlowFile("../../example/capillary.json")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if in.Radius <= 0 || in.Length <= 0 || in.DeltaP <= 0 || in.Mu <= 0 || in.Rho <= 0 {
		t.Errorf("example input not populated: %+v", in)
	}
}

func TestDecodeFlowRequestMissingField(t *testing.T) {
	body := `{"radius":0.001,"length":0.1,"delta_p":1000,"mu":0.002}`
	_, err := DecodeFlowRequest(strings.NewReader(body))
	if err == nil {
		t.Fatalf("expected missing-field error, got nil")
	}
	if !model.IsCode(err, model.CodeMissingField) {
		t.Errorf("expected code %q, got %q", model.CodeMissingField, model.Code(err))
	}
}

func TestDecodeDeltaRequest(t *testing.T) {
	body := `{"target_q":1e-8,"radius":2.5e-4,"length":0.1,"mu":0.002,"rho":1000}`
	in, targetQ, err := DecodeDeltaRequest(strings.NewReader(body))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if in.Radius != 2.5e-4 {
		t.Errorf("radius = %g, want 0.00025", in.Radius)
	}
	if targetQ != 1e-8 {
		t.Errorf("target_q = %g, want 1e-8", targetQ)
	}
}

func TestWriteJSONRoundTrip(t *testing.T) {
	var buf strings.Builder
	payload := map[string]float64{"q": 1e-8}
	if err := WriteJSON(&buf, payload); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var decoded map[string]float64
	if err := json.Unmarshal([]byte(buf.String()), &decoded); err != nil {
		t.Fatalf("unexpected decode error: %v", err)
	}
	if decoded["q"] != 1e-8 {
		t.Errorf("decoded q = %g, want 1e-8", decoded["q"])
	}
}
