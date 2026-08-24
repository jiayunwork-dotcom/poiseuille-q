package units

import "math"

const (
	MillimetersPerMeter               = 1000.0
	CentimetersPerMeter               = 100.0
	PascalsPerKiloPascal              = 1000.0
	PascalsPerBar                     = 1e5
	SecondsPerMinute                  = 60.0
	MilliPascalsSecondPerPascalSecond = 1000.0
	MetersPerMillimeter               = 1 / MillimetersPerMeter
	MetersPerCentimeter               = 1 / CentimetersPerMeter
)

// ToMeters converts a length in millimeters to SI meters.
func ToMeters(mm float64) float64 {
	return mm * MetersPerMillimeter
}

// ToPascals converts kilo-pascals to pascals.
func ToPascals(kPa float64) float64 {
	return kPa * PascalsPerKiloPascal
}

// ToPascalSeconds converts milli-pascal-seconds to SI viscosity.
func ToPascalSeconds(mPa_s float64) float64 {
	return mPa_s / MilliPascalsSecondPerPascalSecond
}

// FromMeters converts SI meters to millimeters.
func FromMeters(m float64) float64 {
	return m * MillimetersPerMeter
}

// FromPascals converts pascals to kilo-pascals.
func FromPascals(pa float64) float64 {
	return pa / PascalsPerKiloPascal
}

// FromPascalSeconds converts SI viscosity to milli-pascal-seconds.
func FromPascalSeconds(pa_s float64) float64 {
	return pa_s * MilliPascalsSecondPerPascalSecond
}

// PressureHead converts a pressure difference to head using rho*g.
func PressureHead(deltaP, rho, g float64) float64 {
	if rho <= 0 || g == 0 {
		return 0
	}
	return deltaP / (rho * g)
}

// DynamicFromKinematic converts kinematic viscosity to dynamic.
func DynamicFromKinematic(nu, rho float64) float64 {
	return nu * rho
}

// KinematicFromDynamic converts dynamic viscosity to kinematic.
func KinematicFromDynamic(mu, rho float64) float64 {
	if rho == 0 {
		return math.Inf(1)
	}
	return mu / rho
}
