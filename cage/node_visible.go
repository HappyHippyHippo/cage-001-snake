package cage

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// RenderArgs @todo doc
type RenderArgs struct {
	Delta  time.Duration
	Target *ebiten.Image
	Geom   *ebiten.GeoM
}

// IVisible @todo doc
type IVisible interface {
	IsVisible() bool
	Show() error
	Hide() error
	Render(args RenderArgs) error
}

// Visible @todo doc
type Visible struct {
	isVisible bool
}

var _ IVisible = &Visible{}

// Init @todo doc
func (v *Visible) Init(self ...IVisible) IVisible {
	var ref IVisible = v
	if len(self) > 0 {
		ref = self[0]
	}

	v.isVisible = false

	if ps, ok := ref.(IPubSub); ok {
		_ = ps.Subscribe(SignalShow, "cage", func(interface{}) error { return ref.Show() })
		_ = ps.Subscribe(SignalHide, "cage", func(interface{}) error { return ref.Hide() })
		_ = ps.Subscribe(SignalRender, "cage", func(data interface{}) error { return ref.Render(data.(RenderArgs)) })
	}

	return v
}

// IsVisible @todo doc
func (v *Visible) IsVisible() bool {
	return v.isVisible
}

// Show @todo doc
func (v *Visible) Show() error {
	v.isVisible = true
	return nil
}

// Hide @todo doc
func (v *Visible) Hide() error {
	v.isVisible = false
	return nil
}

// Render @todo doc
func (v *Visible) Render(RenderArgs) error {
	return nil
}
