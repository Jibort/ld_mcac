// Implementació del tipus RangeF64TwoPi utilitzant constants consolidades
// CreatedAt: 2024-12-27 dc. GPT(JIQ)

package RF64

import (
	"fmt"
	"math"

	cs "github.com/jibort/ld_mcac/internal/core/Consts"
	rngs "github.com/jibort/ld_mcac/internal/core/intf/ranges"
	tools "github.com/jibort/ld_mcac/internal/core/tools"
)

// Tipus RangeF64TwoPi representa el rang [-2·π, +2·π].
type RangeF64TwoPi struct {
	rngs.F64RangeTwoPiIntf

	v RangeF64
}

// NewRangeF64TwoPi crea una nova instància de RangeF64TwoPi amb validació del rang.
func NewRangeF64TwoPi(value float64) RangeF64TwoPi {
	if math.IsNaN(value) || math.IsInf(value, 0) || (value >= cs.Range64Configs.TwoPiF64.Max && value <= cs.Range64Configs.TwoPiF64.Max) {
		return RangeF64TwoPi{v: RangeF64{value: value}}
	}
	panic(fmt.Sprintf("valor fora del rang [-2·π, +2·π]: %f", value))
}

// INTERFÍCIE 'ClonableIntf' ----------
// Clone retorna una còpia de la instància.
func (sR642Pi RangeF64TwoPi) Clone() rngs.ClonableIntf {
	return NewRangeF64TwoPi(sR642Pi.v.value)
}

// INTERFÍCIE 'ComparableIntf' --------
// Equals comprova si dos valors són iguals.
func (sR642Pi RangeF64TwoPi) Equals(pOther rngs.ComparableIntf) bool {
	switch other := pOther.(type) {
	case F64RangeOne:
		return tools.Equals64(sR642Pi.AsFloat64(), other.AsFloat64(), &cs.Epsilon64)
	case RangeF64TwoPi:
		return tools.Equals64(sR642Pi.AsFloat64(), other.AsFloat64(), &cs.Epsilon64)
	default:
		panic(fmt.Sprintf("RangeF64TwoPi.Equals: tipus no vàlid: %T", pOther))

	}
}

// Cert només si la instància és menor que pOther.
func (sR642Pi RangeF64TwoPi) LessThan(pOther rngs.ComparableIntf) bool {
	switch other := pOther.(type) {
	case F64RangeOne:
		return sR642Pi.AsFloat64() < other.AsFloat64()
	case RangeF64TwoPi:
		return sR642Pi.AsFloat64() < other.AsFloat64()
	default:
		panic(fmt.Sprintf("RangeF64One.LessThan: tipus no vàlid: %T", pOther))
	}
}

// Cert només si la instància és menor o igual que pOther.
func (sR642Pi RangeF64TwoPi) LessOrEqualThan(pOther rngs.ComparableIntf) bool {
	switch other := pOther.(type) {
	case F64RangeOne:
		return sR642Pi.LessThan(other) || sR642Pi.Equals(other)
	case RangeF64TwoPi:
		return sR642Pi.LessThan(other) || sR642Pi.Equals(other)
	default:
		panic(fmt.Sprintf("RangeF64One.LessOrEqualThan: tipus no vàlid: %T", pOther))
	}
}

// Cert només si la instància és major que pOther.
func (sR642Pi RangeF64TwoPi) GreaterThan(pOther rngs.ComparableIntf) bool {
	switch other := pOther.(type) {
	case F64RangeOne:
		return sR642Pi.AsFloat64() > other.AsFloat64()
	case RangeF64TwoPi:
		return sR642Pi.AsFloat64() > other.AsFloat64()
	default:
		panic(fmt.Sprintf("RangeF64One.GreaterThan: tipus no vàlid: %T", pOther))
	}
}

// Cert només si la instància és major o igual que pOther.
func (sR642Pi RangeF64TwoPi) GreaterOrEqualThan(pOther rngs.ComparableIntf) bool {
	switch other := pOther.(type) {
	case F64RangeOne:
		return sR642Pi.LessThan(other) || sR642Pi.Equals(other)
	case RangeF64TwoPi:
		return sR642Pi.LessThan(other) || sR642Pi.Equals(other)
	default:
		panic(fmt.Sprintf("RangeF64One.GreaterOrEqualThan: tipus no vàlid: %T", pOther))
	}
}

func (sR642Pi RangeF64TwoPi) Is32() bool     { return false }
func (sR642Pi RangeF64TwoPi) Is64() bool     { return true }
func (sR642Pi RangeF64TwoPi) IsGroupB() bool { return false }
func (sR642Pi RangeF64TwoPi) IsGroupC() bool { return false }
func (sR642Pi RangeF64TwoPi) IsGroupD() bool { return false }
func (sR642Pi RangeF64TwoPi) IsOne() bool    { return false }
func (sR642Pi RangeF64TwoPi) IsTwoPit() bool { return true }

func (sR642Pi RangeF64TwoPi) SetFloat64(pF64 float64) {
	sR642Pi.v.value = pF64
}

// IsGroupA verifica si el valor pertany al Grup A.
func (sR642Pi RangeF64TwoPi) IsGroupA() bool {
	_, exponent, mantissa := tools.DecomposeF64(sR642Pi.v.value)

	// Valors normals
	if exponent >= 1 && exponent <= 2046 {
		return sR642Pi.v.value >= cs.Range64Configs.TwoPiF64.Min && sR642Pi.v.value <= cs.Range64Configs.TwoPiF64.Max
	}

	// Valors subnormals
	if exponent == 0 && mantissa != 0 {
		return true
	}

	// ±Inf i NaN
	if exponent == 2047 {
		return true
	}

	return false
}

func (sR642Pi RangeF64TwoPi) AsFloat32() float32 {
	return float32(sR642Pi.v.value)
}

func (sR642Pi RangeF64TwoPi) AsFloat64() float64 {
	return sR642Pi.v.value
}

func (sR642Pi RangeF64TwoPi) AsUint32() uint32 {
	panic("Encara no implementat!: func (sR642Pi RangeF64TwoPi) AsUint32() uint32")
}

func (sR642Pi RangeF64TwoPi) AsUint64() uint64 {
	return tools.F64ToU64(sR642Pi.v.value)
}

func (sR642Pi RangeF64TwoPi) SetUint32(value uint32) {
	panic("Encara no implementat!: func (sR642Pi RangeF64TwoPi) SetUint32(value uint32)")
}

func (sR642Pi RangeF64TwoPi) SetUint64(value uint64) {
	sR642Pi.v.value = tools.U64ToF64(value)
}

func (sR642Pi RangeF64TwoPi) Clone() rngs.RangeIntf {
	return *NewRangeF64TwoPi(sR642Pi.v.value)
}

// Add suma dos RangeF64TwoPi i retorna un nou valor saturat dins del rang.
func (sR64One RangeF64TwoPi) Add(other RangeF64TwoPi) RangeF64TwoPi {
	result := sR642Pi.v.value + other.value
	if result > cs.Range64Configs.TwoPiF64.Max {
		result = cs.Range64Configs.TwoPiF64.Max
	} else if result < cs.Range64Configs.TwoPiF64.Min {
		result = cs.Range64Configs.TwoPiF64.Min
	}
	return RangeF64TwoPi{value: result}
}

// IsInfinitePos comprova si el valor és +Inf.
func (sR64One RangeF64TwoPi) IsInfinitePos() bool {
	return math.IsInf(sR642Pi.v.value, 1)
}

// IsInfiniteNeg comprova si el valor és -Inf.
func (sR64One RangeF64TwoPi) IsInfiniteNeg() bool {
	return math.IsInf(sR642Pi.v.value, -1)
}

// IsInfinite comprova si el valor és infinit (positiu o negatiu).
func (sR64One RangeF64TwoPi) IsInfinite() bool {
	return math.IsInf(sR642Pi.v.value, 0)
}

// GetF64Value retorna el valor float64.
func (sR64One RangeF64TwoPi) GetF64Value() float64 {
	return sR642Pi.v.value
}

// SetF64Value modifica el valor float64 i retorna l'instància actualitzada.
func (sR64One *RangeF64TwoPi) SetF64Value(value float64) RangeF64TwoPi {
	sR642Pi.v.value = value
	return *sR64One
}
