package cage

import (
	"errors"
)

// PublishHandler @todo doc
type PublishHandler func(interface{}) error

// PublishType @todo doc
type PublishType int

const (
	// DontPropagate @todo doc
	DontPropagate PublishType = iota

	// PropagatePrePublish @todo doc
	PropagatePrePublish

	// PropagatePostPublish @todo doc
	PropagatePostPublish
)

// IPubSub @todo doc
type IPubSub interface {
	Subscribe(signal int, id string, handler PublishHandler) error
	Unsubscribe(signal int, id string) error
	Publish(signal int, data interface{}, propagate ...PublishType) error
}

// PubSub @todo doc
type PubSub struct {
	traversable ITraversable
	signals     map[int]map[string]PublishHandler
}

var _ IPubSub = &PubSub{}

// Init @todo doc
func (ps *PubSub) Init(traversable ITraversable) IPubSub {
	ps.traversable = traversable
	ps.signals = map[int]map[string]PublishHandler{}

	return ps
}

// Subscribe @todo doc
func (ps *PubSub) Subscribe(signal int, id string, handler PublishHandler) error {
	if handler == nil {
		return errNilPointer("handler")
	}

	if _, exists := ps.signals[signal]; !exists {
		ps.signals[signal] = map[string]PublishHandler{}
	}

	ps.signals[signal][id] = handler
	return nil
}

// Unsubscribe @todo doc
func (ps *PubSub) Unsubscribe(signal int, id string) error {
	if _, exists := ps.signals[signal]; exists {
		delete(ps.signals[signal], id)
	}
	return nil
}

// Publish @todo doc
func (ps *PubSub) Publish(signal int, data interface{}, propagate ...PublishType) error {
	publishType := DontPropagate
	if len(propagate) > 0 {
		publishType = propagate[0]
	}

	if publishType == PropagatePrePublish {
		if e := ps.propagate(signal, data, PropagatePrePublish); e != nil {
			return e
		}
	}

	if handlers, exists := ps.signals[signal]; exists {
		for _, handler := range handlers {
			if e := handler(data); e != nil {
				if !errors.Is(e, ErrBreak) {
					return e
				}
				return nil
			}
		}
	}

	if publishType == PropagatePostPublish {
		if e := ps.propagate(signal, data, PropagatePostPublish); e != nil {
			return e
		}
	}
	return nil
}

func (ps *PubSub) propagate(signal int, data interface{}, propagate PublishType) error {
	for _, id := range ps.traversable.ListChildren() {
		if pub, ok := ps.traversable.GetChild(id).(IPubSub); ok {
			if e := pub.Publish(signal, data, propagate); e != nil {
				if !errors.Is(e, ErrBreak) {
					return e
				}
				return nil
			}
		}
	}
	return nil
}
