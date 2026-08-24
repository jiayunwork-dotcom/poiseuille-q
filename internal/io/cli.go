package io

import (
	"fmt"
	"io"

	"poiseuille-q/internal/delta"
	"poiseuille-q/internal/flow"
	"poiseuille-q/internal/model"
	"poiseuille-q/internal/report"
)

// RunFlowFile computes a flow case from a JSON file.
func RunFlowFile(path string, table bool, out io.Writer) error {
	in, err := LoadFlowFile(path)
	if err != nil {
		return err
	}
	res, err := flow.Compute(in)
	if err != nil {
		return err
	}
	summary := report.BuildFlow(res)
	if table {
		return WriteFlowTable(out, summary)
	}
	return WriteFlowOutput(out, summary)
}

// RunDelta computes the pressure required for a target flow.
func RunDelta(in modelInput, targetQ float64, table bool, out io.Writer) error {
	res, err := delta.ResolveTargetPressure(in, targetQ)
	if err != nil {
		return err
	}
	summary := report.BuildDelta(res)
	if table {
		return WriteDeltaTable(out, summary)
	}
	return WriteDeltaOutput(out, summary)
}

// BuildInputFromValues constructs a model.Input from CLI flag values.
func BuildInputFromValues(radius, length, mu, rho float64, allowReverse bool) (modelInput, error) {
	in := modelInput{
		Radius:       radius,
		Length:       length,
		DeltaP:       0,
		Mu:           mu,
		Rho:          rho,
		AllowReverse: allowReverse,
	}
	if err := validateInput(in); err != nil {
		return modelInput{}, err
	}
	return in, nil
}

// DescribeInput prints the five input values.
func DescribeInput(in modelInput) string {
	return fmt.Sprintf("R=%g m L=%g m dP=%g Pa mu=%g Pa.s rho=%g kg/m3",
		in.Radius, in.Length, in.DeltaP, in.Mu, in.Rho)
}

// Alias types keep the CLI layer decoupled from model names in call sites.
type modelInput = model.Input

func validateInput(in modelInput) error {
	return model.Validate(in)
}
