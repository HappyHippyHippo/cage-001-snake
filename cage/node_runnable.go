package cage

import (
	"time"
)

// UpdateArgs @todo doc
type UpdateArgs struct {
	Delta time.Duration
}

// IRunnable @todo doc
type IRunnable interface {
	IsRunning() bool
	Run() error
	Pause() error
	Update(args UpdateArgs) error
}

// Runnable @todo doc
type Runnable struct {
	isRunning bool
}

var _ IRunnable = &Runnable{}

// Init @todo doc
func (r *Runnable) Init(self ...IRunnable) IRunnable {
	var ref IRunnable = r
	if len(self) > 0 {
		ref = self[0]
	}

	r.isRunning = false

	if ps, ok := ref.(IPubSub); ok {
		_ = ps.Subscribe(SignalRun, "cage", func(interface{}) error { return ref.Run() })
		_ = ps.Subscribe(SignalPause, "cage", func(interface{}) error { return ref.Pause() })
		_ = ps.Subscribe(SignalUpdate, "cage", func(data interface{}) error { return ref.Update(data.(UpdateArgs)) })
	}

	return r
}

// IsRunning @todo doc
func (r *Runnable) IsRunning() bool {
	return r.isRunning
}

// Run @todo doc
func (r *Runnable) Run() error {
	r.isRunning = true
	return nil
}

// Pause @todo doc
func (r *Runnable) Pause() error {
	r.isRunning = false
	return nil
}

// Update @todo doc
func (r *Runnable) Update(UpdateArgs) error {
	return nil
}
