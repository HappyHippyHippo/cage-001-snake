package cage

// Signed @todo doc
type Signed interface {
	int | int8 | int16 | int32 | int64
}

// Unsigned @todo doc
type Unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64
}

// Float @todo doc
type Float interface {
	float32 | float64
}

// Number @todo doc
type Number interface {
	Signed | Unsigned | Float
}

type (
	// Vector @todo doc
	Vector[T Number] [2]T
	// VectorI32 @todo doc
	VectorI32 Vector[int32]
	// VectorI64 @todo doc
	VectorI64 Vector[int64]
	// VectorF32 @todo doc
	VectorF32 Vector[float32]
	// VectorF64 @todo doc
	VectorF64 Vector[float64]
)

type (
	// Matrix @todo doc
	Matrix[T Number] [3][3]T
	// MatrixI32 @todo doc
	MatrixI32 Matrix[int32]
	// MatrixI64 @todo doc
	MatrixI64 Matrix[int64]
	// MatrixF32 @todo doc
	MatrixF32 Matrix[float32]
	// MatrixF64 @todo doc
	MatrixF64 Matrix[float64]
)

// Rect @todo doc
type Rect[T Number] struct {
	X      T
	Y      T
	Width  T
	Height T
}

type (
	// RectI32 @todo doc
	RectI32 Rect[int32]
	// RectI64 @todo doc
	RectI64 Rect[int64]
	// RectF32 @todo doc
	RectF32 Rect[float32]
	// RectF64 @todo doc
	RectF64 Rect[float64]
)
