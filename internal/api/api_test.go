package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIFlowSuccess(t *testing.T) {
	mux := NewMux()
	body := `{"radius":2.5e-4,"length":0.1,"delta_p":1000,"mu":0.002,"rho":1000}`
	req := httptest.NewRequest(http.MethodPost, "/api/flow", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200; body=%s", rec.Code, rec.Body.String())
	}
	var out map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("unexpected JSON error: %v", err)
	}
	if out["q_m3_s"] == nil || out["re"] == nil {
		t.Errorf("response missing q_m3_s or re: %v", out)
	}
	if out["laminar"] != true {
		t.Errorf("laminar = %v, want true", out["laminar"])
	}
}

func TestAPIRejectsTurbulentWithErrorBody(t *testing.T) {
	mux := NewMux()
	body := `{"radius":0.01,"length":1,"delta_p":1000,"mu":0.001,"rho":1000}`
	req := httptest.NewRequest(http.MethodPost, "/api/flow", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusUnprocessableEntity {
		t.Fatalf("status = %d, want 422; body=%s", rec.Code, rec.Body.String())
	}
	var out ErrorBody
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("unexpected JSON error: %v", err)
	}
	if out.Error.Code != "not_laminar" {
		t.Errorf("error code = %q, want not_laminar", out.Error.Code)
	}
}

func TestAPIDeltaPSuccess(t *testing.T) {
	mux := NewMux()
	body := `{"target_q":7.6699e-9,"radius":2.5e-4,"length":0.1,"mu":0.002,"rho":1000}`
	req := httptest.NewRequest(http.MethodPost, "/api/delta-p", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200; body=%s", rec.Code, rec.Body.String())
	}
	var out map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("unexpected JSON error: %v", err)
	}
	if out["required_delta_p_pa"] == nil {
		t.Errorf("response missing required_delta_p_pa: %v", out)
	}
}

func TestAPIMissingFieldError(t *testing.T) {
	mux := NewMux()
	body := `{"radius":2.5e-4,"length":0.1,"delta_p":1000,"mu":0.002}`
	req := httptest.NewRequest(http.MethodPost, "/api/flow", bytes.NewBufferString(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400; body=%s", rec.Code, rec.Body.String())
	}
	var out ErrorBody
	if err := json.Unmarshal(rec.Body.Bytes(), &out); err != nil {
		t.Fatalf("unexpected JSON error: %v", err)
	}
	if out.Error.Code != "missing_field" {
		t.Errorf("error code = %q, want missing_field", out.Error.Code)
	}
}

func TestAPIHealth(t *testing.T) {
	mux := NewMux()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
}
