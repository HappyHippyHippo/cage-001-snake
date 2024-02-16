package cage

import "time"

// ITimer @todo doc
type ITimer interface {
	ITraversable
	IAwakable
	IRunnable
	IPubSub

	GetDuration() time.Duration
	SetDuration(duration time.Duration) ITimer
}

// Timer @todo doc
type Timer struct {
	Traversable
	Awakable
	Runnable
	PubSub

	self     ITimer
	delta    time.Duration
	duration time.Duration
}

var _ ITimer = &Timer{}

// Init @todo doc
func (t *Timer) Init(id string, self ...ITimer) ITimer {
	var ref ITimer = t
	if len(self) > 0 {
		ref = self[0]
	}

	t.Traversable.Init(id, ref)
	t.PubSub.Init(ref)
  t.Awakable.Init(ref)
	t.Runnable.Init(ref)
	
  _ = ref.Subscribe(SignalAwake, "cage", func(interface{}) error { return ref.Awake() })
	_ = ref.Subscribe(SignalSleep, "cage", func(interface{}) error { return ref.Sleep() })
	_ = ref.Subscribe(SignalClose, "cage", func(interface{}) error { return ref.Close() })
	_ = ref.Subscribe(SignalRun, "cage", func(interface{}) error { return ref.Run() })
	_ = ref.Subscribe(SignalPause, "cage", func(interface{}) error { return ref.Pause() })
	_ = ref.Subscribe(SignalUpdate, "cage", func(data interface{}) error { return ref.Update(data.(UpdateArgs)) })

	t.delta = 0
	t.duration = 0
	return t
}

// GetDuration @todo doc
func (t *Timer) GetDuration() time.Duration {
	return t.duration
}

// SetDuration @todo doc
func (t *Timer) SetDuration(duration time.Duration) ITimer {
	if !t.IsRunning() {
		t.duration = duration
	}
	return t
}
