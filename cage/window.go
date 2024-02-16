package cage

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// ResizeMode @todo doc
type ResizeMode string

const (
	// ResizeDisabled @todo doc
	ResizeDisabled ResizeMode = "disabled"
	
  // ResizeFullscreen @todo doc
	ResizeFullscreen = "fullscreen"
	
  // ResizeEnabled @todo doc
	ResizeEnabled = "enabled"
)

// IWindow @todo doc
type IWindow interface {
	ITraversable
	IPubSub

	GetTitle() string
	SetTitle(title string)
	GetWindowSize() VectorI32
	SetWindowSize(size VectorI32)
	GetViewportSize() VectorI32
	SetViewportSize(size VectorI32)
	GetResizeMode() ResizeMode
	SetResizeMode(mode ResizeMode)
}

// Window @todo doc
type Window struct {
	Traversable
	PubSub

	title        string
	windowSize   VectorI32
	viewportSize VectorI32
	resizeMode   ResizeMode
}

var _ ITraversable = &Window{}
var _ IWindow = &Window{}

// Init @todo doc
func (w *Window) Init(self ...IWindow) IWindow {
	var ref IWindow = w
	if len(self) > 0 {
		ref = self[0]
	}

	w.Traversable.Init("window", ref)
	w.PubSub.Init(ref)
	w.SetTitle("cage")
	w.SetWindowSize(VectorI32{320, 200})
	w.SetViewportSize(VectorI32{320, 200})
	w.SetResizeMode(ResizeDisabled)

	return ref
}

// GetTitle @todo doc
func (w *Window) GetTitle() string {
	return w.title
}

// SetTitle @todo doc
func (w *Window) SetTitle(title string) {
	w.title = title
	ebiten.SetWindowTitle(title)
}

// GetWindowSize @todo doc
func (w *Window) GetWindowSize() VectorI32 {
	return w.windowSize
}

// SetWindowSize @todo doc
func (w *Window) SetWindowSize(size VectorI32) {
	w.windowSize = size
	ebiten.SetWindowSize(int(w.windowSize[0]), int(w.windowSize[1]))
}

// GetViewportSize @todo doc
func (w *Window) GetViewportSize() VectorI32 {
	return w.viewportSize
}

// SetViewportSize @todo doc
func (w *Window) SetViewportSize(size VectorI32) {
	w.viewportSize = size
}

// GetResizeMode @todo doc
func (w *Window) GetResizeMode() ResizeMode {
	return w.resizeMode
}

// SetResizeMode @todo doc
func (w *Window) SetResizeMode(mode ResizeMode) {
	w.resizeMode = mode
	ebiten.SetWindowResizingMode(w.parseResizeFlag(w.resizeMode))
}

func (*Window) parseResizeFlag(mode ResizeMode) ebiten.WindowResizingModeType {
	switch mode {
	case ResizeEnabled:
		return ebiten.WindowResizingModeEnabled
	case ResizeFullscreen:
		return ebiten.WindowResizingModeOnlyFullscreenEnabled
	case ResizeDisabled:
		return ebiten.WindowResizingModeDisabled
	default:
		return ebiten.WindowResizingModeDisabled
	}
}
