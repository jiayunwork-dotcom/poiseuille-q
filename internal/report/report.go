package report

import (
	"fmt"
	"sort"

	"poiseuille-q/internal/model"
)

// FlowSummary is the user-facing numeric summary of one case.
type FlowSummary struct {
	Q         float64 `json:"q_m3_s"`
	UAvg      float64 `json:"u_avg_m_s"`
	UMax      float64 `json:"u_max_m_s"`
	Re        float64 `json:"re"`
	TauW      float64 `json:"tau_w_pa"`
	DeltaP    float64 `json:"delta_p_pa"`
	Direction string  `json:"direction"`
	Laminar   bool    `json:"laminar"`
}

// DeltaSummary is the user-facing output of the delta-p endpoint.
type DeltaSummary struct {
	RequiredDeltaP float64 `json:"required_delta_p_pa"`
	ResultingQ     float64 `json:"resulting_q_m3_s"`
	Re             float64 `json:"re"`
	TauW           float64 `json:"tau_w_pa"`
	UAvg           float64 `json:"u_avg_m_s"`
	UMax           float64 `json:"u_max_m_s"`
	Laminar        bool    `json:"laminar"`
}

// BuildFlow converts a model.Result into an API-safe summary.
func BuildFlow(r model.Result) FlowSummary {
	return FlowSummary{
		Q:         r.Q,
		UAvg:      r.UAvg,
		UMax:      r.UMax,
		Re:        r.Re,
		TauW:      r.TauW,
		DeltaP:    r.DeltaP,
		Direction: r.Direction,
		Laminar:   r.Laminar,
	}
}

// BuildDelta converts a delta model result into a summary.
func BuildDelta(d model.DeltaResult) DeltaSummary {
	return DeltaSummary{
		RequiredDeltaP: d.RequiredDeltaP,
		ResultingQ:     d.ResultingQ,
		Re:             d.Re,
		TauW:           d.TauW,
		UAvg:           d.UAvg,
		UMax:           d.UMax,
		Laminar:        d.Laminar,
	}
}

// TextTable renders a stable key-value report for CLI output.
func TextTable(rows map[string]string) string {
	keys := make([]string, 0, len(rows))
	for k := range rows {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := ""
	for _, k := range keys {
		out += fmt.Sprintf("%-28s %s\n", k, rows[k])
	}
	return out
}

// Labeled returns the label with a fallback when empty.
func Labeled(label string) string {
	if label == "" {
		return "untitled case"
	}
	return label
}
