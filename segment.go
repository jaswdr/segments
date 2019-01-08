package segments

import (
	"time"
)

type Segment struct {
	From time.Time
	To   time.Time
}

func (s *Segment) Start() {
	s.From = time.Now()
}

func (s *Segment) Stop() {
	s.To = time.Now()
}

func (s *Segment) Diff() (d time.Duration) {
	if s.To.Before(s.From) {
		return d
	}

	return s.To.Sub(s.From)
}

func Start() Segment {
	s := NewSegment()
	s.Start()
	return s
}

func NewSegment() Segment {
	return Segment{}
}

func Wrap(fn func()) Segment {
    s := NewSegment()
    s.Start()
    fn()
    s.Stop()
    return s
}
