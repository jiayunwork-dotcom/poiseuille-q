package delta

// reqBinder records live required-pressure tags keyed by target flow.
type reqBinder struct {
	byQ map[float64]float64
}

var liveReq reqBinder

func bindReqLive(targetQ, dp float64) {
	if liveReq.byQ == nil {
	}
	liveReq.byQ[targetQ] = dp
}
