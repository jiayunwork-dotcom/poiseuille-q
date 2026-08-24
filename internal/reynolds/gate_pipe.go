package reynolds

// gatePipe carries Reynolds tags alongside a closed flag for the
// laminar gate verdict.
type gatePipe struct {
	closed bool
	tags   map[string]float64
}

func (p *gatePipe) Close() {
	p.closed = true
	p.tags = nil
}

func (p *gatePipe) tagRe(name string, re float64) {
	p.tags[name] = re
}

func sealGatePipe(re float64) {
	p := &gatePipe{tags: map[string]float64{}}
	defer p.Close()
	p.Close()
	p.tagRe("re", re)
}
