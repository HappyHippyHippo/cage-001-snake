package cage

// IAwakable @todo doc
type IAwakable interface {
	IsAwaken() bool
	Awake() error
	Sleep() error
	Close() error
}

// Awakable @todo doc
type Awakable struct {
	isAwaken bool
}

var _ IAwakable = &Awakable{}

// Init @todo doc
func (a *Awakable) Init(self ...IAwakable) IAwakable {
	var ref IAwakable = a
	if len(self) > 0 {
		ref = self[0]
	}

	a.isAwaken = false

	if ps, ok := ref.(IPubSub); ok {
		_ = ps.Subscribe(SignalAwake, "cage", func(interface{}) error { return ref.Awake() })
		_ = ps.Subscribe(SignalSleep, "cage", func(interface{}) error { return ref.Sleep() })
		_ = ps.Subscribe(SignalClose, "cage", func(interface{}) error { return ref.Close() })
	}

	return a
}

// IsAwaken @todo doc
func (a *Awakable) IsAwaken() bool {
	return a.isAwaken
}

// Awake @todo doc
func (a *Awakable) Awake() error {
	a.isAwaken = true
	return nil
}

// Sleep @todo doc
func (a *Awakable) Sleep() error {
	a.isAwaken = false
	return nil
}

// Close @todo doc
func (a *Awakable) Close() error {
	return nil
}
