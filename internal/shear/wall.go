package shear

// TauAtRadius returns the local shear at radial position r.
func TauAtRadius(r, radius, deltaP, length float64) float64 {
	if length <= 0 {
		return 0
	}
	return r * deltaP / (2 * length)
}

// TauAtCenter is zero by symmetry of the laminar profile.
func TauAtCenter() float64 {
	return 0
}

// TauRatioToWall normalizes local shear by the wall value.
func TauRatioToWall(r, radius float64) float64 {
	if radius <= 0 {
		return 0
	}
	return r / radius
}

// WallShearIndependentOfViscosity compares two wall-shear values.
func WallShearIndependentOfViscosity(tauA, tauB float64) bool {
	diff := tauA - tauB
	scale := tauA
	if scale < 0 {
		scale = -scale
	}
	return diff == 0 || diff/scale <= 1e-9
}

// MaximumShear is the wall value for a straight tube.
func MaximumShear(deltaP, radius, length float64) float64 {
	return WallShear(deltaP, radius, length)
}

// MinimumShear is the centerline value.
func MinimumShear() float64 {
	return TauAtCenter()
}

// ShearGradient returns tau per unit radius for the linear profile.
func ShearGradient(deltaP, length float64) float64 {
	if length <= 0 {
		return 0
	}
	return deltaP / (2 * length)
}

// ProfileRatio returns tau(r)/tau_w for the linear shear profile.
func ProfileRatio(r, radius float64) float64 {
	return TauRatioToWall(r, radius)
}
