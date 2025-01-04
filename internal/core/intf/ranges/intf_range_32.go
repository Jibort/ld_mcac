// Interfícies de Range per a 32 bits.
// CreatedAt: 2025/01/03 dv. JIQ

package ranges

import (
	base "github.com/jibort/ld_mcac/internal/core/intf/base"
)

// Per tots els Range32 ---------------
// Interfície per a valors fixed o floating points de 32bits.
type X32RangeIntf interface {
	base.RangeIntf          // Hereta les funcions generals per a tots els Range32
	base.MathOperationsIntf // Hereta les funcions matemàtiques per a tots els Range32

	As64() X64RangeIntf // Conversió a 64 bits
}

// PER TIPUS FLOAT32 I UINT32 ---------
// Interfície per a valors floating points de 32bits.
type F32RangeIntf interface {
	X32RangeIntf // Hereta les funcions generals per a tots els Range32

	AsF64() F64RangeIntf // Conversió a 64bits
}

// Interfície per a valors sencers positius de 32bits.
type U32RangeIntf interface {
	X32RangeIntf

	AsU64() U64RangeIntf // Conversió a uint64 bits
}

// TIPUS ONE I TWOPI (32 bits) --------
// Interfície per a valors floating points de 32bits.
type F32RangeOneIntf interface {
	F32RangeIntf // Hereta les funcions generals per a tots els Range32

	AsF32TwoPi() F32RangeTwoPiIntf // Conversió a [-2·π, +2·π] bits
}

// Interfície per a valors floating points de 32bits.
type F32RangeTwoPiIntf interface {
	F32RangeIntf // Hereta les funcions generals per a tots els RangeF32

	AsF64One() F64RangeOneIntf // Conversió a rang [-1.0,+1.0] bits
}

// Interfície per a valors sencers positius de 32bits.
type U32RangeOneIntf interface {
	U32RangeIntf // Hereta les funcions generals per a tots els RangeU32

	AsU32TwoPi() U32RangeTwoPiIntf // Conversió a [-2·π, +2·π] bits
}

// Interfície per a valors sencers positius de 32bits.
type U32RangeTwoPiIntf interface {
	U32RangeIntf // Hereta les funcions generals per a tots els RangeU32

	AsU32One() U32RangeOneIntf // Conversió a rang [-1.0,+1.0] bits
}
