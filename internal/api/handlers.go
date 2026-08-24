package api

import (
	"net/http"

	"poiseuille-q/internal/delta"
	"poiseuille-q/internal/flow"
	"poiseuille-q/internal/report"
)

// HandleFlow computes the laminar result for a POST /api/flow request.
func HandleFlow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}
	in, err := DecodeFlowBody(LimitBody(r.Body, MaxBodySize))
	if err != nil {
		writeError(w, err)
		return
	}
	res, err := flow.Compute(in)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, report.BuildFlow(res))
}

// HandleDeltaP computes required pressure for a POST /api/delta-p request.
func HandleDeltaP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		methodNotAllowed(w)
		return
	}
	in, targetQ, err := DecodeDeltaBody(LimitBody(r.Body, MaxBodySize))
	if err != nil {
		writeError(w, err)
		return
	}
	res, err := delta.ResolveTargetPressure(in, targetQ)
	if err != nil {
		writeError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, report.BuildDelta(res))
}

// HandleHealth reports that the service is running.
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
		"name":   "poiseuille-q",
	})
}

// HandleRoot describes the available laminar-flow endpoints.
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"service": "poiseuille-q",
		"endpoints": []string{
			"POST /api/flow",
			"POST /api/delta-p",
		},
	})
}

// NewRoutes registers every public route.
func NewRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", HandleRoot)
	mux.HandleFunc("/health", HandleHealth)
	mux.HandleFunc("/api/flow", HandleFlow)
	mux.HandleFunc("/api/delta-p", HandleDeltaP)
}
