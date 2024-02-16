// Package present @todo doc
package present

import (
	"image/color"

	"github.com/happyhippyhippo/cage"
)

// Scene @todo doc
type Scene struct {
	cage.Node

	sprite cage.Sprite
}

// Init @todo doc
func (s *Scene) Init() *Scene {
	s.Node.Init("present", s)

	_ = s.AddChild(s.sprite.Init("image", "person"))
	s.sprite.SetAnchorXY(0.5, 0.5)

	return s
}

// Awake @todo doc
func (s *Scene) Awake() error {
	wsize := s.Window().GetViewportSize()
	s.sprite.SetPositionXY(float64(wsize[0]/2), float64(wsize[1]/2))
	s.sprite.SetSourceRectXY(0, 0, 32, 32)
	return nil
}

// Render @todo doc
func (s *Scene) Render(args cage.RenderArgs) error {
	args.Target.Fill(color.RGBA{R: 75, G: 0, B: 138, A: 0})
	return nil
}
