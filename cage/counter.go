package cage

import (
	"time"
)

// ICounter @todo doc
type ICounter interface {
	ITraversable

	Step() time.Duration
	GetCount() int
	GetTarget() int
	SetTarget(target int)
}

// Counter @todo doc
type Counter struct {
	Traversable

	count   int
	steps   int
	target  int
	process func() time.Duration
}

var (
	_ ITraversable = &Counter{}
	_ ICounter     = &Counter{}
)

// Init @todo doc
func (pc *Counter) Init(id string, self ICounter) ICounter {
	pc.Traversable.Init(id, self)

	pc.count = 0
	pc.steps = 0
	pc.target = 60
	pc.process = func() func() time.Duration {
		lastUpdate := time.Now()
		dt := time.Since(lastUpdate)
		return func() time.Duration {
			dt = time.Since(lastUpdate)
			now := time.Now()
			if lastUpdate.Second() != now.Second() {
				pc.count = pc.steps
				pc.steps = 0
			} else {
				pc.steps++
			}
			lastUpdate = now
			return dt
		}
	}()

	return self
}

// Step @todo doc
func (pc *Counter) Step() time.Duration {
	return pc.process()
}

// GetCount @todo doc
func (pc *Counter) GetCount() int {
	return pc.count
}

// GetTarget @todo doc
func (pc *Counter) GetTarget() int {
	return pc.target
}

// SetTarget @todo doc
func (pc *Counter) SetTarget(target int) {
	pc.target = target
}
