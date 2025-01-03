// Interfícies generals de Range.
// CreatedAt: 2024/12/08 dg. JIQ

package intf

// Interfície 'Clonable'
type ClonableIntf interface {
	Clone() ClonableIntf
}

// Interície 'Comparable'
type ComparableIntf interface {
	Equals(pOther RangeIntf) bool             // Cert només si pOther correspon al mateix tipus i contingut.
	LessThan(pOther RangeIntf) bool           // Cert només si la instància és menor que pOther.
	LessOrEqualThan(pOther RangeIntf) bool    // Cert només si la instància és menor o igual que pOther.
	GreaterThan(pOther RangeIntf) bool        // Cert només si la instància és major que pOther.
	GreaterOrEqualThan(pOther RangeIntf) bool // Cert només si la instància és major o igual que pOther.

}

// Interfície 'MathOperationsIntf'
type MathOperationsIntf interface {
	ComparableIntf
	Add(pOther RangeIntf) RangeIntf
	Sub(pOther RangeIntf) RangeIntf
	Mul(pOther RangeIntf) RangeIntf
	Div(pOther RangeIntf) RangeIntf
}

// Interfície 'TypeConversions'
type TypeConversionsIntf interface {
	AsFloat64() float64
	SetFloat64(float64)
	AsUint64() uint64
	SetUint64(uint64)

	AsFloat32() float32
	SetFloat32(float32)
	AsUint32() uint32
	SetUint32(uint32)
}

// INTERFÍCIE GLOBAL PER A QUALSEVOL INSTÀNCIA DE 'Range'
type RangeIntf interface {
	TypeConversionsIntf
	ClonableIntf
	ComparableIntf

	// Blocs
	Is64() bool    // Cert només si la instància és de 64 bits.
	Is32() bool    // Cert només si la instància és de 32 bits.
	IsError() bool // Cert només si la instància correspon a un error.

	// Grups
	IsGroupA() bool // Cert només si la instància pertany al grup A.
	IsGroupB() bool // Cert només si la instància pertany al grup B.
	IsGroupC() bool // Cert només si la instància pertany al grup C.
	IsGroupD() bool // Cert només si la instància pertany al grup D.
}

type GroupAIntf interface {
	Sign() bool          // Cert només si el valor és negatiu.
	IsSubnormal() bool   // Cert només si la instància pertany al subgrup A1.
	IsNormalNear0() bool // Cert només si la instància pertany al subgrup A2.
	IsNormalFar0() bool  // Cert només si la instància pertany al subgrup A3.
	IsInfinite() bool    // Cert només si la instància és un valor infinit.
	IsInfinitePos() bool // Cert només si la instància és un valor infinit positiu.
	IsInfiniteNeg() bool // Cert només si la instància és un valor infinit negatiu.
}

type GroupBIntf interface {
	IsSymbolType() bool          // Cert només si la instància és un símbol.
	IsPaddingType() bool         // Cert només si la instància és un valor de padding.
	IsNullType() bool            // Cert només si la instància és un valor nul.
	IsCoordinatesType() bool     // Cert només si la instància és un valor de coordenades.
	IsPercentageRangeType() bool // Cert només si la instància és un rang de percentatges.
	IsPercentageType() bool      // Cert només si la instància és un valor de percentatge.
	IsKeyOrMouseType() bool      // Cert només si la instància és un valor de teclat o ratolí.
	IsSaturatedType() bool       // Cert només si la instància és un valor saturat.
	IsErrorType() bool           // Cert només si la instància és un error.
}

type GroupCIntf interface {
	RangeIntf
	Category() int           // Retorna la categoria
	Fiability() int          // Retorna la fiabilitat
	RelativeWeight() float64 // Retorna el pes relatiu
	Token() uint32           // Retorna el token
}
