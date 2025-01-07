// Interfícies de Range per a 64 bits.
// CreatedAt: 2025/01/03 dv. JIQ

package ranges

import (
	base "github.com/jibort/ld_mcac/internal/core/intf/base"
)

// Per tots els Range64 ---------------
// Interfície per a valors fixed o floating points de 64bits.
type X64RangeIntf interface {
	base.RangeIntf          // Hereta les funcions generals per a tots els Range64.
	base.MathOperationsIntf // Hereta les funcions matemàtiques per a tots els Range64.

	As32() X32RangeIntf // Conversió a 32bits
}

// PER TIPUS FLOAT64 I UINT64 ---------
// Interfície per a valors floating points de 64bits.
type F64RangeIntf interface {
	X64RangeIntf // Hereta les funcions generals per a tots els Range64

	AsF32() F32RangeIntf // Conversió a float32 bits
}

// Interfície per a valors sencers positius de 64bits.
type U64RangeIntf interface {
	X64RangeIntf

	AsU32() U32RangeIntf // Conversió a uint32 bits
}

// TIPUS ONE I TWOPI (64 bits) --------
// Interfície per a valors floating points de 64bits.
type F64RangeOneIntf interface {
	F64RangeIntf // Hereta les funcions generals per a tots els Range64

	AsF64TwoPi() F64RangeTwoPiIntf // Conversió a rang [-2·π, +2·π] en float64
}

// Interfície per a valors floating points de 64bits.
type F64RangeTwoPiIntf interface {
	F64RangeIntf // Hereta les funcions generals per a tots els Range64

	AsF64One() F64RangeOneIntf // Conversió a rang [-1.0,+1.0] bits
}

// TIPUS ONE I TWOPI (64 bits) --------
// Interfície per a valors floating points de 64bits.
type U64RangeOneIntf interface {
	F64RangeIntf // Hereta les funcions generals per a tots els Range64

	AsU64TwoPi() U64RangeTwoPiIntf // Conversió a [-2π, +2π] bits
}

// Interfície per a valors floating points de 64bits.
type U64RangeTwoPiIntf interface {
	F64RangeIntf // Hereta les funcions generals per a tots els Range64

	AsU64One() U64RangeOneIntf // Conversió a rang [-1.0,+1.0] bits
}
