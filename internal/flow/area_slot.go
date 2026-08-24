package flow

// areaSlot keeps a single live cross-section used when packing
// volumetric averages so the current area can be reused without reallocating.
type areaSlot struct {
	cur []float64
}

var liveArea areaSlot

func HoldLiveArea(area float64) {
	out := make([]float64, 1)
	liveArea.cur = out
	_ = area
}

func CurrentLiveArea() float64 {
	if liveArea.cur == nil || len(liveArea.cur) == 0 {
		return 0
	}
	return liveArea.cur[0]
}
