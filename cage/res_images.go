package cage

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// ImageLoaderOkCallback @todo doc
type ImageLoaderOkCallback func(id string, image *ebiten.Image) error

// ImageLoaderErrorCallback @todo doc
type ImageLoaderErrorCallback func(id string, e error) error

// IImageManager @todo doc
type IImageManager interface {
	ITraversable
	IPubSub

	Queue(id, url string, onOkCallback ImageLoaderOkCallback, onErrorCallback ImageLoaderErrorCallback) error
	Load() error
	Get(id string) (*ebiten.Image, error)
}

// ImageManager @todo doc
type ImageManager struct {
	Traversable
	PubSub

	self   IImageManager
	loader IResLoader
	pool   map[string]*ebiten.Image
}

var _ ITraversable = &ImageManager{}
var _ IImageManager = &ImageManager{}

// Init @todo doc
func (im *ImageManager) Init(loader IResLoader, self ...IImageManager) IImageManager {
	var ref IImageManager = im
	if len(self) > 0 {
		ref = self[0]
	}

	im.Traversable.Init("images", ref)
	im.PubSub.Init(ref)
	im.loader = loader
	im.pool = map[string]*ebiten.Image{}

	return ref
}

// Queue @todo doc
func (im *ImageManager) Queue(id, url string, onOkCallback ImageLoaderOkCallback, onErrorCallback ImageLoaderErrorCallback) error {
	return im.loader.Queue(
		id,
		url,
		func(id, url string, b []byte) error {
			img, _, e := ebitenutil.NewImageFromReader(bytes.NewReader(b))
			if e != nil {
				return e
			}
			im.pool[id] = img
			if onOkCallback != nil {
				return onOkCallback(id, img)
			}
			return nil
		},
		func(id string, e error) error {
			if onErrorCallback != nil {
				return onErrorCallback(id, e)
			}
			return nil
		})
}

// Load @todo doc
func (im *ImageManager) Load() error {
	return im.loader.Load()
}

// Get @todo doc
func (im *ImageManager) Get(id string) (*ebiten.Image, error) {
	img, ok := im.pool[id]
	if !ok {
		return nil, errImageNotFound(id)
	}
	return img, nil
}
