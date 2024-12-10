// Interfícies generals dels elements individuals de càlcul pel projecte.
// CreatedAt: 2024/12/08 dg. JIQ

package core

import "math"

const (
	Precision32      = 6
	Precision64      = 15
	SmallThreshold64 = 1e-6 // Nou llindar per valors petits
)

var (
	Epsilon32 = float32(math.Pow(10, -Precision32))
	Factor32  = math.Pow(10, Precision32)
	Epsilon64 = math.Pow(10, -Precision64)
	Factor64  = math.Pow(10, Precision64)
)

// Interfície global
type RangeIntf interface {
	// Comparacions
	Equals(pOther RangeIntf) bool
	LessThan(pOther RangeIntf) bool
	LessOrEqualThan(pOther RangeIntf) bool
	GreaterThan(pOther RangeIntf) bool
	GreaterOrEqualThan(pOther RangeIntf) bool

	// Valors especials.
	IsInfinitePos() bool
	IsInfiniteNeg() bool
	IsNullValue() bool

	// Grups.
	IsGroupA() bool
	IsGroupB() bool
	IsGroupC() bool
	IsGroupD() bool

	// Operacions
	Add(pOther RangeIntf) RangeIntf
	Sub(pOther RangeIntf) RangeIntf
	Mul(pOther RangeIntf) RangeIntf
	Div(pOther RangeIntf) (RangeIntf, error)
	// ... i altres operacions comunes necessàries.

	// Funcions de desencapçulament.
	ValueF64() float64
	ValueF32() float32
	ValueI64() int64
	ValueI32() int32
	ValueU64() uint64
	ValueU32() uint32

	// Funcions de saturació.
	IsSaturated() bool
	IsSaturatedPos() bool
	IsSaturatedNeg() bool

	// Aquestes funcions poden ser útils en un futur.
	AsF64() RangeF64
	// AsF32() RangeF32
	// AsI64() RangeI64
	// AsI32() RangeI32
	// AsU64() RangeU64
	// AsU32() RangeU32
}

// Interfície per a valors fixed o floating points de 64bits.
type Range64Intf interface {
	RangeIntf
	As32() Range32Intf // conversió a 32 bits
}

// Interfície per a valors fixed o floating points de 32bits.
type Range32Intf interface {
	RangeIntf
	As64() Range64Intf // conversió a 64 bits
}
