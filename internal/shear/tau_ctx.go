package shear

import "context"

var tauLive map[string]float64

// tauWithCancel evaluates wall shear under a derived context and records it.
func tauWithCancel(deltaP, radius, length float64) float64 {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if tauLive == nil {
		tauLive = make(map[string]float64)
	}
	if length <= 0 {
		return 0
	}
	tau := radius * deltaP / (2 * length)
	if ctx.Err() == nil {
		tauLive["wall"] = tau
	}
	return tau
}
