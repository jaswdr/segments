package segments

import (
    "math"
    "time"
    "testing"
)

func TestSegmentStart(t *testing.T) {
    s := &Segment{}
    s.Start()
    if s.From.Unix() <= 0 {
        t.Errorf("Expected 'From' attribute to be initialized, found %v", s.From)
    }
}

func TestSegmentStop(t *testing.T) {
    s := &Segment{}
    s.Stop()
    if s.To.Unix() <= 0 {
        t.Errorf("Expected 'To' attribute to be initialized, found %v", s.To)
    }
}

func TestDiff(t *testing.T) {
    s := NewSegment()
    s.Start()
    time.Sleep(time.Second)
    s.Stop()
    time.Sleep(time.Second)
    diff := int(math.Round(s.Diff().Seconds()))
    if diff != 1 {
        t.Errorf("Expecting 1 second of difference, found %d", diff)
    }
}

func TestWrap(t *testing.T) {
    s := Wrap(func() {
        time.Sleep(time.Second)
    })
    diff := int(math.Round(s.Diff().Seconds()))
    if diff != 1 {
        t.Errorf("Expecting 1 second of difference, found %d", diff)
    }
}
