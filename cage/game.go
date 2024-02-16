// Package cage @todo doc
package cage

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

// IGame @todo doc
type IGame interface {
	ITraversable

	Close() error
	Run() error
}

type ebitenAdapter struct {
	window IWindow
	fps    ICounter
	ups    ICounter
	scenes ISceneManager
}

func (ea ebitenAdapter) Update() error {
	return ea.scenes.Update(UpdateArgs{Delta: ea.ups.Step()})
}

func (ea ebitenAdapter) Draw(target *ebiten.Image) {
	_ = ea.scenes.Render(RenderArgs{Delta: ea.fps.Step(), Target: target, Geom: &ebiten.GeoM{}})
}

func (ea ebitenAdapter) Layout(_, _ int) (_, _ int) {
	size := ea.window.GetViewportSize()
	return int(size[0]), int(size[1])
}

// Game @todo doc
type Game struct {
	Traversable
	Runnable

	logger Logger
	window Window
	fps    FPS
	ups    UPS
	res    ResLoader
	scenes SceneManager

	ea ebitenAdapter
}

var _ IGame = &Game{}

// Init @todo doc
func (g *Game) Init() IGame {
	g.Traversable.Init("/", g)
	g.Runnable.Init(g)

	_ = g.AddChild(g.logger.Init())
	_ = g.AddChild(g.window.Init())
	_ = g.AddChild(g.fps.Init())
	_ = g.AddChild(g.ups.Init())
	_ = g.AddChild(g.res.Init(&g.logger))
	_ = g.AddChild(g.scenes.Init(&g.logger))

	g.ea.window = &g.window
	g.ea.fps = &g.fps
	g.ea.ups = &g.ups
	g.ea.scenes = &g.scenes
	return g
}

// Close @todo doc
func (g *Game) Close() error {
	_ = g.scenes.Close()
	_ = g.res.Close()
	_ = g.logger.Close()
	return nil
}

// SetParent @todo doc
func (g *Game) SetParent(ITraversable) {
	// no-op
}

// Run @todo doc
func (g *Game) Run() error {
	_ = g.Runnable.Run()
	defer func() { _ = g.Runnable.Pause() }()
	if e := ebiten.RunGame(g.ea); e != nil {
		if !errors.Is(e, ErrBreak) {
			return e
		}
	}
	return nil
}
