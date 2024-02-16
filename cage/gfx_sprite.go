package cage

import (
  "image"

  "github.com/hajimehoshi/ebiten/v2"
)

// ISprite @todo doc
type ISprite interface {
  IGraphic

  GetSourceRect() RectI32
  SetSourceRect(rect RectI32) ISprite
  SetSourceRectXY(x, y, w, h int32) ISprite
}

// Sprite @todo doc
type Sprite struct {
  Graphic

  imageSrcID   string
  imageSrcRect RectI32
  imageSrc     *ebiten.Image
  image        *ebiten.Image
}

var _ ISprite = &Sprite{}

// Init @todo doc
func (s *Sprite) Init(id, imageID string, self ...ISprite) ISprite {
  s.self = s
  if len(self) > 0 {
    s.self = self[0]
  }
  s.Graphic.Init(id, s.self)
  s.imageSrcID = imageID
  return s
}

// Awake @todo doc
func (s *Sprite) Awake() error {
  imageSrc, e := s.Images().Get(s.imageSrcID)
  if e != nil {
    return e
  }
  s.imageSrc = imageSrc
  s.subImage()
  return nil
}

// Render @todo doc
func (s *Sprite) Render(args RenderArgs) error {
  if s.image == nil {
    return nil
  }

  ops := ebiten.DrawImageOptions{}
  ops.GeoM = s.CalcGlobalGeom(args.Geom)
  args.Target.DrawImage(s.image, &ops)

  return nil
}

// GetSourceRect @todo doc
func (s *Sprite) GetSourceRect() RectI32 {
  return s.imageSrcRect
}

// SetSourceRect @todo doc
func (s *Sprite) SetSourceRect(rect RectI32) ISprite {
  s.imageSrcRect = rect
  return s.subImage()
}

// SetSourceRectXY @todo doc
func (s *Sprite) SetSourceRectXY(x, y, w, h int32) ISprite {
  return s.SetSourceRect(RectI32{X: x, Y: y, Width: w, Height: h})
}

func (s *Sprite) subImage() ISprite {
  if s.imageSrc == nil {
    return s
  }
  if s.image != nil {
    s.image.Clear()
  }
  if s.imageSrcRect.Width == 0 || s.imageSrcRect.Height == 0 {
    s.imageSrcRect.X = 0
    s.imageSrcRect.Y = 0
    s.imageSrcRect.Width = int32(s.imageSrc.Bounds().Dx())
    s.imageSrcRect.Height = int32(s.imageSrc.Bounds().Dy())
  }
  s.image = s.imageSrc.SubImage(image.Rect(
    int(s.imageSrcRect.X),
    int(s.imageSrcRect.Y),
    int(s.imageSrcRect.X + s.imageSrcRect.Width),
    int(s.imageSrcRect.Y + s.imageSrcRect.Height),
  )).(*ebiten.Image)
  s.size[0] = float64(s.imageSrcRect.Width)
  s.size[1] = float64(s.imageSrcRect.Height)
  s.CalcLocalGeom()
  return s
}
