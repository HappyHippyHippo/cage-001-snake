package cage

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// IGraphic @todo doc
type IGraphic interface {
	INode

	GetPosition() VectorF64
	GetPositionXY() (x, y float64)
	SetPosition(position VectorF64) IGraphic
	SetPositionXY(x, y float64) IGraphic
	Translate(delta VectorF64)
	TranslateXY(x, y float64)
	GetAnchor() VectorF64
	GetAnchorXY() (x, y float64)
	SetAnchor(anchor VectorF64) IGraphic
	SetAnchorXY(x, y float64) IGraphic
	GetSize() VectorF64
	GetSizeXY() (x, y float64)
	GetScaledSize() VectorF64
	GetScaledSizeXY() (x, y float64)
	GetScale() VectorF64
	GetScaleXY() (x, y float64)
	SetScale(scale VectorF64) IGraphic
	SetScaleXY(x, y float64) IGraphic
	GetRotation() float64
	SetRotation(theta float64) IGraphic
	Rotate(theta float64) IGraphic
	GetSkew() VectorF64
	GetSkewXY() (x, y float64)
	SetSkew(skew VectorF64) IGraphic
	SetSkewXY(x, y float64) IGraphic
	GetGeometry() ebiten.GeoM
	CalcLocalGeom() IGraphic
	CalcGlobalGeom(global *ebiten.GeoM) ebiten.GeoM
}

// Graphic @todo doc
type Graphic struct {
	Node

	self     IGraphic
	position VectorF64
	anchor   VectorF64
	size     VectorF64
	scale    VectorF64
	rotation float64
	skew     VectorF64
	geometry ebiten.GeoM
}

var _ IGraphic = &Graphic{}

// Init @todo doc
func (g *Graphic) Init(id string, self ...IGraphic) IGraphic {
	g.self = g
	if len(self) > 0 {
		g.self = self[0]
	}
	g.Node.Init(id, g.self)
	g.position = VectorF64{0, 0}
	g.anchor = VectorF64{0.5, 0.5}
	g.size = VectorF64{0, 0}
	g.scale = VectorF64{1, 1}
	g.rotation = 0
	g.skew = VectorF64{0, 0}
	g.geometry = ebiten.GeoM{}
	return g.self.CalcLocalGeom()
}

// GetPosition @todo doc
func (g *Graphic) GetPosition() VectorF64 {
	return g.position
}

// GetPositionXY @todo doc
func (g *Graphic) GetPositionXY() (_, _ float64) {
	return g.position[0], g.position[1]
}

// SetPosition @todo doc
func (g *Graphic) SetPosition(position VectorF64) IGraphic {
	return g.SetPositionXY(position[0], position[1])
}

// SetPositionXY @todo doc
func (g *Graphic) SetPositionXY(x, y float64) IGraphic {
	g.position[0] = x
	g.position[1] = y
	return g.CalcLocalGeom()
}

// Translate @todo doc
func (g *Graphic) Translate(delta VectorF64) {
	g.TranslateXY(g.position[0]+delta[0], g.position[1]+delta[1])
}

// TranslateXY @todo doc
func (g *Graphic) TranslateXY(x, y float64) {
	g.Translate(VectorF64{x, y})
}

// GetAnchor @todo doc
func (g *Graphic) GetAnchor() VectorF64 {
	return g.anchor
}

// GetAnchorXY @todo doc
func (g *Graphic) GetAnchorXY() (_, _ float64) {
	return g.anchor[0], g.anchor[1]
}

// SetAnchor @todo doc
func (g *Graphic) SetAnchor(anchor VectorF64) IGraphic {
	return g.SetAnchorXY(anchor[0], anchor[1])
}

// SetAnchorXY @todo doc
func (g *Graphic) SetAnchorXY(x, y float64) IGraphic {
	g.anchor[0] = min(max(x, 0.0), 1.0)
	g.anchor[1] = min(max(y, 0.0), 1.0)
	return g.CalcLocalGeom()
}

// GetSize @todo doc
func (g *Graphic) GetSize() VectorF64 {
	return g.size
}

// GetSizeXY @todo doc
func (g *Graphic) GetSizeXY() (_, _ float64) {
	s := g.GetSize()
	return s[0], s[1]
}

// GetScaledSize @todo doc
func (g *Graphic) GetScaledSize() VectorF64 {
	return VectorF64{g.size[0] * g.scale[0], g.size[1] * g.scale[1]}
}

// GetScaledSizeXY @todo doc
func (g *Graphic) GetScaledSizeXY() (_, _ float64) {
	s := g.GetScaledSize()
	return s[0], s[1]
}

// GetScale @todo doc
func (g *Graphic) GetScale() VectorF64 {
	return g.scale
}

// GetScaleXY @todo doc
func (g *Graphic) GetScaleXY() (_, _ float64) {
	return g.scale[0], g.scale[1]
}

// SetScale @todo doc
func (g *Graphic) SetScale(scale VectorF64) IGraphic {
	return g.SetScaleXY(scale[0], scale[1])
}

// SetScaleXY @todo doc
func (g *Graphic) SetScaleXY(x, y float64) IGraphic {
	g.scale[0] = x
	g.scale[1] = y
	return g.CalcLocalGeom()
}

// GetRotation @todo doc
func (g *Graphic) GetRotation() float64 {
	return g.rotation
}

// SetRotation @todo doc
func (g *Graphic) SetRotation(theta float64) IGraphic {
	g.rotation = math.Mod(theta+math.Pi*2, math.Pi*2)
	return g.CalcLocalGeom()
}

// Rotate @todo doc
func (g *Graphic) Rotate(theta float64) IGraphic {
	g.rotation = math.Mod(g.rotation+theta+math.Pi*2, math.Pi*2)
	return g.CalcLocalGeom()
}

// GetSkew @todo doc
func (g *Graphic) GetSkew() VectorF64 {
	return g.skew
}

// GetSkewXY @todo doc
func (g *Graphic) GetSkewXY() (_, _ float64) {
	return g.skew[0], g.skew[1]
}

// SetSkew @todo doc
func (g *Graphic) SetSkew(skew VectorF64) IGraphic {
	return g.SetSkewXY(skew[0], skew[1])
}

// SetSkewXY @todo doc
func (g *Graphic) SetSkewXY(x, y float64) IGraphic {
	g.skew[0] = x
	g.skew[1] = y
	return g.CalcLocalGeom()
}

// GetGeometry @todo doc
func (g *Graphic) GetGeometry() ebiten.GeoM {
	return g.geometry
}

// CalcLocalGeom @todo doc
func (g *Graphic) CalcLocalGeom() IGraphic {
	g.geometry.Reset()
	g.geometry.Scale(g.scale[0], g.scale[1])
	g.geometry.Skew(g.skew[0], g.skew[1])
	g.geometry.Rotate(g.rotation)
	g.geometry.Translate(g.position[0], g.position[1])
	g.geometry.Translate(-1*(g.size[0]*g.scale[0]*g.anchor[0]), -1*(g.size[1]*g.scale[1]*g.anchor[1]))
	return g.self
}

// CalcGlobalGeom @todo doc
func (g *Graphic) CalcGlobalGeom(global *ebiten.GeoM) ebiten.GeoM {
	geometry := g.geometry
	geometry.Concat(*global)
	return geometry
}
