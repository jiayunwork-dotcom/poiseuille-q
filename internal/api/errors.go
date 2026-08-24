package api

import (
	"encoding/json"
	"net/http"

	"poiseuille-q/internal/model"
)

// ErrorBody is the stable error shape returned to clients.
type ErrorBody struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail carries the machine code and human message.
type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// writeError serializes a structured error as JSON.
func writeError(w http.ResponseWriter, err error) {
	status := http.StatusBadRequest
	if model.IsCode(err, model.CodeNotLaminar) {
		status = http.StatusUnprocessableEntity
	}
	body := ErrorBody{
		Error: ErrorDetail{
			Code:    string(model.Code(err)),
			Message: err.Error(),
		},
	}
	writeJSON(w, status, body)
}

// writeJSON writes a response with a status code.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

// methodNotAllowed responds with a JSON method error.
func methodNotAllowed(w http.ResponseWriter) {
	writeError(w, model.NewError(model.CodeUnknownCommand, "method not allowed"))
}

// notFound responds with a JSON route error.
func notFound(w http.ResponseWriter) {
	writeError(w, model.NewError(model.CodeUnknownCommand, "route not found"))
}
