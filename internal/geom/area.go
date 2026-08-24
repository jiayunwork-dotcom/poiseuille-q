package geom

// TubeVolume returns the internal volume of a pipe segment.
func TubeVolume(radius, length float64) float64 {
	return CircleArea(radius) * length
}

// TubeWallArea returns the cylindrical inner surface area.
func TubeWallArea(radius, length float64) float64 {
	return WettedPerimeter(radius) * length
}

// CrossSectionPerimeter is an alias for WettedPerimeter.
func CrossSectionPerimeter(radius float64) float64 {
	return WettedPerimeter(radius)
}

// AnnulusArea returns the area between an outer and inner radius.
func AnnulusArea(outer, inner float64) float64 {
	if inner >= outer {
		return 0
	}
	return CircleArea(outer) - CircleArea(inner)
}

// DiameterSquared computes D^2, useful when comparing geometric scales.
func DiameterSquared(radius float64) float64 {
	return 4 * radius * radius
}

// CharacteristicLength returns the hydraulic diameter for a full tube.
func CharacteristicLength(radius float64) float64 {
	return HydraulicDiameter(radius)
}

// VolumeFlowArea keeps the cross-section that carries flow.
func VolumeFlowArea(radius float64) float64 {
	return CircleArea(radius)
}
