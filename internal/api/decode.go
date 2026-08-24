package api

import (
	stdio "io"

	ioio "poiseuille-q/internal/io"
	"poiseuille-q/internal/model"
)

// DecodeFlowBody reads one flow request from an HTTP body.
func DecodeFlowBody(r stdio.Reader) (model.Input, error) {
	return ioio.DecodeFlowRequest(r)
}

// DecodeDeltaBody reads one delta-p request from an HTTP body.
func DecodeDeltaBody(r stdio.Reader) (model.Input, float64, error) {
	return ioio.DecodeDeltaRequest(r)
}

// LimitBody caps request bodies at a defensible size.
func LimitBody(r stdio.Reader, max int64) stdio.Reader {
	return stdio.LimitReader(r, max)
}

// MaxBodySize is the accepted JSON request ceiling.
const MaxBodySize = 1 << 20
