package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"poiseuille-q/internal/api"
	"poiseuille-q/internal/flow"
	"poiseuille-q/internal/io"
	"poiseuille-q/internal/model"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	if len(args) == 0 {
		return serve([]string{})
	}
	switch args[0] {
	case "serve":
		return serve(args[1:])
	case "flow":
		return flowCommand(args[1:])
	case "delta-p":
		return deltaCommand(args[1:])
	case "example":
		return exampleCommand(args[1:])
	case "check":
		return checkCommand(args[1:])
	case "help", "-h", "--help":
		usage()
		return 0
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n", args[0])
		usage()
		return 2
	}
}

func serve(args []string) int {
	fs := flag.NewFlagSet("serve", flag.ContinueOnError)
	addr := fs.String("http", ":8080", "HTTP listen address")
	if err := fs.Parse(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	if err := api.Listen(*addr); err != nil {
		fmt.Fprintf(os.Stderr, "serve: %v\n", err)
		return 1
	}
	return 0
}

func flowCommand(args []string) int {
	path, table, err := parseFlowArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	if path == "" {
		fmt.Fprintln(os.Stderr, "usage: poiseuille-q flow <input.json> [--table]")
		return 2
	}
	if err := io.RunFlowFile(path, table, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "flow: %v\n", err)
		return 1
	}
	return 0
}

func parseFlowArgs(args []string) (string, bool, error) {
	path := ""
	table := false
	for _, a := range args {
		switch {
		case a == "--table" || a == "-table":
			table = true
		case strings.HasPrefix(a, "-"):
			return "", false, fmt.Errorf("unknown flag %q", a)
		case path == "":
			path = a
		default:
			return "", false, fmt.Errorf("unexpected argument %q", a)
		}
	}
	return path, table, nil
}

func deltaCommand(args []string) int {
	fs := flag.NewFlagSet("delta-p", flag.ContinueOnError)
	targetQ := fs.Float64("q", 0, "target volumetric flow in m3/s")
	radius := fs.Float64("radius", 0, "tube radius in m")
	length := fs.Float64("length", 0, "tube length in m")
	mu := fs.Float64("mu", 0, "dynamic viscosity in Pa.s")
	rho := fs.Float64("rho", 0, "density in kg/m3")
	allowReverse := fs.Bool("allow-reverse", false, "accept signed target flow")
	table := fs.Bool("table", false, "print a human-readable table")
	if err := fs.Parse(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	in, err := io.BuildInputFromValues(*radius, *length, *mu, *rho, *allowReverse)
	if err != nil {
		fmt.Fprintf(os.Stderr, "delta-p: %v\n", err)
		return 1
	}
	if err := model.ValidateFlowTarget(*targetQ, *allowReverse); err != nil {
		fmt.Fprintf(os.Stderr, "delta-p: %v\n", err)
		return 1
	}
	if err := io.RunDelta(in, *targetQ, *table, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "delta-p: %v\n", err)
		return 1
	}
	return 0
}

func exampleCommand(args []string) int {
	path := io.ExamplePath()
	table := len(args) > 0 && args[0] == "--table"
	if err := io.RunFlowFile(path, table, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "example: %v\n", err)
		return 1
	}
	return 0
}

func checkCommand(args []string) int {
	fs := flag.NewFlagSet("check", flag.ContinueOnError)
	if err := fs.Parse(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 2
	}
	path := fs.Arg(0)
	if path == "" {
		path = io.ExamplePath()
	}
	in, err := io.LoadFlowFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "check: %v\n", err)
		return 1
	}
	res, err := flow.Compute(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "check: %v\n", err)
		return 1
	}
	fmt.Fprintf(os.Stdout, "case %q: Q=%g m3/s, Re=%.1f, laminar=%v\n",
		label(in), res.Q, res.Re, res.Laminar)
	return 0
}

func usage() {
	fmt.Fprintln(os.Stdout, `poiseuille-q: laminar circular pipe flow calculator

Commands:
  serve [--http :8080]       start the HTTP service (default command)
  flow <input.json> [--table]
  delta-p --q ... --radius ... --length ... --mu ... --rho ...
  example [--table]          run example/capillary.json
  check [input.json]         load a case and print the laminar verdict`)
}

func label(in model.Input) string {
	if in.Label != "" {
		return in.Label
	}
	return "untitled"
}
