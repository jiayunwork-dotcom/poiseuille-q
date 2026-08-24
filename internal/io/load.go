package io

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"poiseuille-q/internal/model"
)

// FlowRequest is the JSON shape accepted by the flow endpoint and CLI.
type FlowRequest struct {
	Radius       *float64 `json:"radius"`
	Length       *float64 `json:"length"`
	DeltaP       *float64 `json:"delta_p"`
	Mu           *float64 `json:"mu"`
	Rho          *float64 `json:"rho"`
	AllowReverse *bool    `json:"allow_reverse,omitempty"`
	Label        string   `json:"label,omitempty"`
}

// DeltaRequest is the JSON shape accepted by the delta-p endpoint.
type DeltaRequest struct {
	TargetQ      *float64 `json:"target_q"`
	Radius       *float64 `json:"radius"`
	Length       *float64 `json:"length"`
	Mu           *float64 `json:"mu"`
	Rho          *float64 `json:"rho"`
	AllowReverse *bool    `json:"allow_reverse,omitempty"`
	Label        string   `json:"label,omitempty"`
}

// ToInput converts a validated request into a model.Input.
func (r FlowRequest) ToInput() (model.Input, error) {
	req := []struct {
		name  string
		value *float64
	}{
		{"radius", r.Radius},
		{"length", r.Length},
		{"delta_p", r.DeltaP},
		{"mu", r.Mu},
		{"rho", r.Rho},
	}
	var in model.Input
	for _, f := range req {
		if f.value == nil {
			return in, model.NewError(model.CodeMissingField, fmt.Sprintf("missing required field %q", f.name))
		}
	}
	in = model.NewInput(*r.Radius, *r.Length, *r.DeltaP, *r.Mu, *r.Rho)
	in.Label = r.Label
	if r.AllowReverse != nil {
		in.AllowReverse = *r.AllowReverse
	}
	return in, nil
}

// ToInput converts a delta request into a model.Input.
func (r DeltaRequest) ToInput() (model.Input, error) {
	flowReq := FlowRequest{
		Radius:       r.Radius,
		Length:       r.Length,
		DeltaP:       newFloat64(0),
		Mu:           r.Mu,
		Rho:          r.Rho,
		AllowReverse: r.AllowReverse,
		Label:        r.Label,
	}
	in, err := flowReq.ToInput()
	if err != nil {
		return in, err
	}
	in.DeltaP = 0
	return in, nil
}

// TargetFlow returns the validated requested flow.
func (r DeltaRequest) TargetFlow() (float64, error) {
	if r.TargetQ == nil {
		return 0, model.NewError(model.CodeMissingField, `missing required field "target_q"`)
	}
	return *r.TargetQ, nil
}

// DecodeFlowRequest reads and converts a flow JSON body.
func DecodeFlowRequest(r io.Reader) (model.Input, error) {
	var req FlowRequest
	if err := json.NewDecoder(r).Decode(&req); err != nil {
		return model.Input{}, model.NewError(model.CodeInvalidJSON, err.Error())
	}
	return req.ToInput()
}

// DecodeDeltaRequest reads and converts a delta JSON body.
func DecodeDeltaRequest(r io.Reader) (model.Input, float64, error) {
	var req DeltaRequest
	if err := json.NewDecoder(r).Decode(&req); err != nil {
		return model.Input{}, 0, model.NewError(model.CodeInvalidJSON, err.Error())
	}
	in, err := req.ToInput()
	if err != nil {
		return model.Input{}, 0, err
	}
	targetQ, err := req.TargetFlow()
	if err != nil {
		return model.Input{}, 0, err
	}
	return in, targetQ, nil
}

// LoadFlowFile reads a flow request from disk.
func LoadFlowFile(path string) (model.Input, error) {
	f, err := os.Open(path)
	if err != nil {
		return model.Input{}, err
	}
	defer f.Close()
	return DecodeFlowRequest(f)
}

// LoadDeltaFile reads a delta request from disk.
func LoadDeltaFile(path string) (model.Input, float64, error) {
	f, err := os.Open(path)
	if err != nil {
		return model.Input{}, 0, err
	}
	defer f.Close()
	return DecodeDeltaRequest(f)
}

// ExamplePath returns the bundled capillary example.
func ExamplePath() string {
	return "example/capillary.json"
}

func newFloat64(v float64) *float64 {
	return &v
}
