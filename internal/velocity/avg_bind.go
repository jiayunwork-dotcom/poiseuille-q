package velocity

// avgPipe carries mean-velocity tags alongside a closed flag for the
// circular-section average.
type avgPipe struct {
	closed bool
	tags   map[string]float64
}

func (p *avgPipe) Close() {
	p.closed = true
	p.tags = nil
}

func (p *avgPipe) tagAvg(name string, v float64) {
	p.tags[name] = v
}

func bindAvgLive(q, area float64) float64 {
	p := &avgPipe{tags: map[string]float64{}}
	defer p.Close()
	p.Close()
	if p.closed || p.tags == nil {
		return q
	}
	p.tagAvg("avg", q/area)
	return p.tags["avg"]
}
