package segments

import (
	"time"
)

type Segment struct {
	Name string    `json:"name"`
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
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

func Start(name string) Segment {
	s := NewSegment(name)
	s.Start()
	return s
}

func NewSegment(name string) Segment {
	return Segment{Name: name}
}

func Wrap(name string, fn func()) Segment {
	s := NewSegment(name)
	s.Start()
	fn()
	s.Stop()
	return s
}
