# poiseuille-q

poiseuille-q computes the laminar Hagen-Poiseuille flow of a Newtonian
fluid in a circular tube. Given radius, length, pressure difference,
dynamic viscosity, and density, it returns volumetric flow, average and
centerline velocity, wall shear stress, and Reynolds number. A second
calculation inverts the problem: given a target volumetric flow, it finds
the required pressure difference.

The solver enforces the laminar boundary. A case with Reynolds number at
or above 2300 is rejected with a `not_laminar` error instead of returning
an invalid laminar answer. Negative pressure differences are rejected by
default; set `allow_reverse` to accept signed flow.

## Run a bundled example

```bash
go run . flow --table example/capillary.json
```

The capillary example is a 0.5 mm tube, 0.1 m long, driven by 1000 Pa.
Its Reynolds number is far below 2300.

## CLI

```bash
go run . flow example/capillary.json
go run . delta-p --q 7.6699e-9 --radius 0.00025 --length 0.1 --mu 0.002 --rho 1000
```

Invalid input is written to standard error and exits nonzero.

## HTTP service

The default command starts a service on port 8080:

```bash
go run . serve --http :8080
```

Endpoints:

```text
POST /api/flow
POST /api/delta-p
GET  /health
```

Example request:

```bash
curl -s -X POST http://127.0.0.1:8080/api/flow \
  -H 'Content-Type: application/json' \
  -d '{"radius":0.00025,"length":0.1,"delta_p":1000,"mu":0.002,"rho":1000}'
```

Errors are returned as JSON with a machine-readable `code` and a human
message, for example `not_laminar` or `missing_field`.

## Build and test

```bash
go build ./...
go test ./...
```

The calculation kernel lives under `internal/`; `main.go` only wires the
CLI and HTTP entry points.
