package io

import (
	"encoding/json"
	"io"

	"poiseuille-q/internal/model"
	"poiseuille-q/internal/report"
)

// WriteJSON encodes a value with indentation and a trailing newline.
func WriteJSON(w io.Writer, v any) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

// WriteFlowOutput writes the flow summary as JSON.
func WriteFlowOutput(w io.Writer, s report.FlowSummary) error {
	return WriteJSON(w, s)
}

// WriteDeltaOutput writes the delta-p summary as JSON.
func WriteDeltaOutput(w io.Writer, s report.DeltaSummary) error {
	return WriteJSON(w, s)
}

// WriteModelError writes a structured error body.
func WriteModelError(w io.Writer, err error) error {
	body := map[string]any{
		"error": map[string]any{
			"code":    model.Code(err),
			"message": err.Error(),
		},
	}
	return WriteJSON(w, body)
}

// WriteFlowTable writes a human-readable table.
func WriteFlowTable(w io.Writer, s report.FlowSummary) error {
	t := &report.Table{}
	t.AddFloat("volumetric flow (m3/s)", s.Q)
	t.AddFloat("average velocity (m/s)", s.UAvg)
	t.AddFloat("centerline velocity (m/s)", s.UMax)
	t.AddFloat("Reynolds number", s.Re)
	t.AddFloat("wall shear stress (Pa)", s.TauW)
	t.AddFloat("pressure difference (Pa)", s.DeltaP)
	t.Add("direction", s.Direction)
	t.Add("laminar", boolWord(s.Laminar))
	return t.Render(w)
}

// WriteDeltaTable writes a human-readable delta-p result.
func WriteDeltaTable(w io.Writer, s report.DeltaSummary) error {
	t := &report.Table{}
	t.AddFloat("required pressure (Pa)", s.RequiredDeltaP)
	t.AddFloat("resulting flow (m3/s)", s.ResultingQ)
	t.AddFloat("Reynolds number", s.Re)
	t.AddFloat("wall shear stress (Pa)", s.TauW)
	t.AddFloat("average velocity (m/s)", s.UAvg)
	t.AddFloat("centerline velocity (m/s)", s.UMax)
	t.Add("laminar", boolWord(s.Laminar))
	return t.Render(w)
}

func boolWord(v bool) string {
	if v {
		return "yes"
	}
	return "no"
}
