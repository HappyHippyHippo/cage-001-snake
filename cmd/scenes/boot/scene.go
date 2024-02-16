// Package boot @todo doc
package boot

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/happyhippyhippo/cage"
)

// Scene @todo doc
type Scene struct {
	cage.Node

	count  int
	scenes cage.ISceneManager
}

// Init @todo doc
func (s *Scene) Init() *Scene {
	s.Node.Init("boot", s)
	return s
}

// Awake @todo doc
func (s *Scene) Awake() error {
	s.count = 1
	s.scenes = s.Path("/scenes").(cage.ISceneManager)
	images := s.Path("/res/images").(cage.IImageManager)
	_ = images.Queue("person", "res/person.png", s.imageLoaded, s.imageError)
	_ = images.Load()
	return nil
}

// Render @todo doc
func (s *Scene) Render(args cage.RenderArgs) error {
	args.Target.Fill(color.RGBA{R: 75, G: 0, B: 138, A: 0})
	return nil
}

func (s *Scene) imageLoaded(_ string, _ *ebiten.Image) error {
	s.count--
	if s.count == 0 {
		_ = s.scenes.QueueScene("present")
	}
	return nil
}

func (s *Scene) imageError(_ string, _ error) error {
	_ = s.scenes.Close()
	return nil
}
