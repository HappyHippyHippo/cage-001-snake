package cage

import (
	"strings"
)

// ITraversable @todo doc
type ITraversable interface {
	Self() ITraversable
	ID() string

	GetRoot() ITraversable
	GetParent() ITraversable
	SetParent(parent ITraversable)

	ListChildren() []string
	HasChild(id string) bool
	GetChild(id string) ITraversable
	AddChild(child ITraversable) error
	RemoveChild(id string) error

	Path(path string) ITraversable
}

// Traversable @todo doc
type Traversable struct {
	self     ITraversable
	id       string
	parent   ITraversable
	children []ITraversable
}

var _ ITraversable = &Traversable{}

// Init @todo doc
func (t *Traversable) Init(id string, self ITraversable) ITraversable {
	t.self = self
	t.id = id
	t.parent = nil
	t.children = []ITraversable{}

	return self
}

// Self @todo doc
func (t *Traversable) Self() ITraversable {
	return t.self
}

// ID @todo doc
func (t *Traversable) ID() string {
	return t.id
}

// GetRoot @todo doc
func (t *Traversable) GetRoot() ITraversable {
	if t.parent != nil {
		return t.parent.GetRoot()
	}
	return t.self
}

// GetParent @todo doc
func (t *Traversable) GetParent() ITraversable {
	return t.parent
}

// SetParent @todo doc
func (t *Traversable) SetParent(parent ITraversable) {
	t.parent = parent
}

// ListChildren @todo doc
func (t *Traversable) ListChildren() []string {
	var res []string
	for _, child := range t.children {
		res = append(res, child.ID())
	}
	return res
}

// HasChild @todo doc
func (t *Traversable) HasChild(id string) bool {
	for _, child := range t.children {
		if child.ID() == id {
			return true
		}
	}
	return false
}

// GetChild @todo doc
func (t *Traversable) GetChild(id string) ITraversable {
	for _, child := range t.children {
		if child.ID() == id {
			return child
		}
	}
	return nil
}

// AddChild @todo doc
func (t *Traversable) AddChild(child ITraversable) error {
	if child == nil {
		return errNilPointer("child")
	}

	id := child.ID()
	if t.self.HasChild(id) {
		return errDuplicateChild(t.id, id)
	}

	t.children = append(t.children, child)
	child.SetParent(t.self)
	return nil
}

// RemoveChild @todo doc
func (t *Traversable) RemoveChild(id string) error {
	for idx, child := range t.children {
		if child.ID() == id {
			child.SetParent(nil)
			t.children = append(t.children[:idx], t.children[idx+1:]...)
			return nil
		}
	}
	return errChildNotFound(t.id, id)
}

// Path @todo doc
func (t *Traversable) Path(path string) ITraversable {
	switch {
	case path == "" || path == "." || path == "./":
		return t.self
	case path == ".." || path == "../":
		return t.parent
	case strings.HasPrefix(path, "/"):
		return t.self.GetRoot().Path(path[1:])
	case strings.HasPrefix(path, "./"):
		return t.self.Path(path[2:])
	case strings.HasPrefix(path, "../"):
		if t.parent == nil {
			return nil
		}
		return t.parent.Path(path[3:])
	}

	parts := strings.Split(path, "/")
	child := t.self.GetChild(parts[0])
	if child == nil {
		return nil
	}
	if len(parts) == 1 {
		return child
	}

	return child.Path(path[len(parts[0])+1:])
}
