package model

import "fmt"

// ErrorCode identifies the machine-readable category of a rejection.
type ErrorCode string

const (
	CodeInvalidRadius     ErrorCode = "invalid_radius"
	CodeInvalidLength     ErrorCode = "invalid_length"
	CodeInvalidViscosity  ErrorCode = "invalid_viscosity"
	CodeInvalidDensity    ErrorCode = "invalid_density"
	CodeNegativeDeltaP    ErrorCode = "negative_delta_p"
	CodeNotLaminar        ErrorCode = "not_laminar"
	CodeInvalidTargetFlow ErrorCode = "invalid_target_flow"
	CodeInvalidJSON       ErrorCode = "invalid_json"
	CodeMissingField      ErrorCode = "missing_field"
	CodeOutOfRange        ErrorCode = "out_of_range"
	CodeUnknownCommand    ErrorCode = "unknown_command"
)

// Error is a structured validation or calculation error.
type Error struct {
	Code    ErrorCode
	Field   string
	Value   float64
	Message string
}

// Error implements the error interface.
func (e *Error) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("%s: %s (field %s=%g)", e.Code, e.Message, e.Field, e.Value)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// NewError builds a structured error with a public message.
func NewError(code ErrorCode, message string) *Error {
	return &Error{Code: code, Message: message}
}

// WithField attaches the invalid field name and value.
func (e *Error) WithField(field string, value float64) *Error {
	e.Field = field
	e.Value = value
	return e
}

// IsCode checks whether an error is a model.Error with the given code.
func IsCode(err error, code ErrorCode) bool {
	if err == nil {
		return false
	}
	me, ok := err.(*Error)
	if !ok {
		return false
	}
	return me.Code == code
}

// Code extracts a stable error code from any error.
func Code(err error) ErrorCode {
	if err == nil {
		return ""
	}
	if me, ok := err.(*Error); ok {
		return me.Code
	}
	return CodeUnknownCommand
}

// LaminarError is the specific turbulent rejection carrying the measured Re.
type LaminarError struct {
	Re    float64
	Limit float64
}

// Error formats the laminar boundary message for API and CLI users.
func (e *LaminarError) Error() string {
	return fmt.Sprintf("Reynolds number %.2f exceeds laminar limit %.0f", e.Re, e.Limit)
}

// NewLaminarError builds a not_laminar model.Error with detail.
func NewLaminarError(re, limit float64) *Error {
	return &Error{
		Code:    CodeNotLaminar,
		Message: fmt.Sprintf("Reynolds number %.2f exceeds laminar limit %.0f", re, limit),
	}
}
