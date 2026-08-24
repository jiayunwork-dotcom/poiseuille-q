package model

// Input carries the physical quantities of one circular pipe case.
type Input struct {
	Radius       float64
	Length       float64
	DeltaP       float64
	Mu           float64
	Rho          float64
	AllowReverse bool
	Label        string
}

// Result is the complete set of laminar pipe-flow outputs.
type Result struct {
	Q         float64
	UAvg      float64
	UMax      float64
	Re        float64
	TauW      float64
	DeltaP    float64
	Radius    float64
	Length    float64
	Mu        float64
	Rho       float64
	Direction string
	Laminar   bool
}

// DeltaResult carries the outputs of a target-flow inversion.
type DeltaResult struct {
	RequiredDeltaP float64
	ResultingQ     float64
	Re             float64
	TauW           float64
	UAvg           float64
	UMax           float64
	Laminar        bool
}

// NewInput builds an Input with the five required physical fields.
func NewInput(radius, length, deltaP, mu, rho float64) Input {
	return Input{
		Radius:       radius,
		Length:       length,
		DeltaP:       deltaP,
		Mu:           mu,
		Rho:          rho,
		AllowReverse: false,
	}
}

// WithReverse enables signed pressure differences and returns a copy.
func (i Input) WithReverse(enabled bool) Input {
	i.AllowReverse = enabled
	return i
}

// WithLabel attaches a human-readable case name.
func (i Input) WithLabel(label string) Input {
	i.Label = label
	return i
}

// IsForward reports whether the pressure difference drives flow forward.
func (i Input) IsForward() bool {
	return i.DeltaP >= 0
}

// DirectionName returns the direction implied by the pressure sign.
func (i Input) DirectionName() string {
	if i.DeltaP < 0 {
		return "reverse"
	}
	return "forward"
}

// IsZeroPressure reports the no-flow boundary case.
func (i Input) IsZeroPressure() bool {
	return i.DeltaP == 0
}

// Copy returns a shallow copy safe for downstream mutation.
func (i Input) Copy() Input {
	return i
}
