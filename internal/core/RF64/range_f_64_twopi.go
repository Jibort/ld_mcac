// Implementació del tipus F64RangeTwoPi (rang [-2·π, +2·π]) utilitzant constants consolidades
// CreatedAt: 2024-12-27 dc. GPT(JIQ)

package rf64

import (
	"fmt"
	"math"

	cs "github.com/jibort/ld_mcac/internal/core/consts"
	base "github.com/jibort/ld_mcac/internal/core/intf/base"
	rngs "github.com/jibort/ld_mcac/internal/core/intf/ranges"
	tools "github.com/jibort/ld_mcac/internal/core/tools"
)

// Tipus F64RangeTwoPi representa el rang [-2π, +2π].
type F64RangeTwoPi struct {
	rngs.F64RangeTwoPiIntf

	v F64Range
}

// NewF64RangeTwoPi crea una nova instància de F64RangeTwoPi amb validació del rang.
func NewF64RangeTwoPi(value float64) F64RangeTwoPi {
	if math.IsNaN(value) || math.IsInf(value, 0) || (value >= cs.Range64Configs.TwoPiF64.Max && value <= cs.Range64Configs.TwoPiF64.Max) {
		return F64RangeTwoPi{v: F64Range{Value: value}}
	}
	panic(fmt.Sprintf("valor fora del rang [-2·π, +2·π]: %f", value))
}

// INTERFÍCIE 'ClonableIntf' ----------
// Clone retorna una còpia de la instància.
func (sR642Pi F64RangeTwoPi) Clone() base.RangeIntf {
	res := NewF64RangeTwoPi(sR642Pi.v.Value)
	return &res
}

// INTERFÍCIE 'ComparableIntf' --------
// Equals comprova si dos valors són iguals.
func (sR642Pi F64RangeTwoPi) Equals(pOther base.RangeIntf) bool {
	switch other := pOther.(type) {
	case *F64RangeOne:
		return tools.Equals64(sR642Pi.AsFloat64(), other.AsFloat64(), &cs.Epsilon64)
	case *F64RangeTwoPi:
		return tools.Equals64(sR642Pi.AsFloat64(), other.AsFloat64(), &cs.Epsilon64)
	default:
		panic(fmt.Sprintf("F64RangeTwoPi.Equals: tipus no vàlid: %T", pOther))

	}
}

// Cert només si la instància és menor que pOther.
func (sR642Pi F64RangeTwoPi) LessThan(pOther base.RangeIntf) bool {
	switch other := pOther.(type) {
	case *F64RangeOne:
		return sR642Pi.AsFloat64() < other.AsFloat64()
	case *F64RangeTwoPi:
		return sR642Pi.AsFloat64() < other.AsFloat64()
	default:
		panic(fmt.Sprintf("F64RangeOne.LessThan: tipus no vàlid: %T", pOther))
	}
}

// Cert només si la instància és menor o igual que pOther.
func (sR642Pi F64RangeTwoPi) LessOrEqualThan(pOther base.RangeIntf) bool {
	switch other := pOther.(type) {
	case *F64RangeOne:
		return sR642Pi.LessThan(other) || sR642Pi.Equals(other)
	case *F64RangeTwoPi:
		return sR642Pi.LessThan(other) || sR642Pi.Equals(other)
	default:
		panic(fmt.Sprintf("F64RangeOne.LessOrEqualThan: tipus no vàlid: %T", pOther))
	}
}

// Cert només si la instància és major que pOther.
func (sR642Pi F64RangeTwoPi) GreaterThan(pOther base.RangeIntf) bool {
	switch other := pOther.(type) {
	case *F64RangeOne:
		return sR642Pi.AsFloat64() > other.AsFloat64()
	case *F64RangeTwoPi:
		return sR642Pi.AsFloat64() > other.AsFloat64()
	default:
		panic(fmt.Sprintf("F64RangeOne.GreaterThan: tipus no vàlid: %T", pOther))
	}
}

// Cert només si la instància és major o igual que pOther.
func (sR642Pi F64RangeTwoPi) GreaterOrEqualThan(pOther base.RangeIntf) bool {
	switch other := pOther.(type) {
	case *F64RangeOne:
		return sR642Pi.LessThan(other) || sR642Pi.Equals(other)
	case *F64RangeTwoPi:
		return sR642Pi.LessThan(other) || sR642Pi.Equals(other)
	default:
		panic(fmt.Sprintf("F64RangeOne.GreaterOrEqualThan: tipus no vàlid: %T", pOther))
	}
}

func (sR642Pi F64RangeTwoPi) Is32() bool     { return false }
func (sR642Pi F64RangeTwoPi) Is64() bool     { return true }
func (sR642Pi F64RangeTwoPi) IsGroupB() bool { return false }
func (sR642Pi F64RangeTwoPi) IsGroupC() bool { return false }
func (sR642Pi F64RangeTwoPi) IsGroupD() bool { return false }
func (sR642Pi F64RangeTwoPi) IsOne() bool    { return false }
func (sR642Pi F64RangeTwoPi) IsTwoPit() bool { return true }

func (sR642Pi *F64RangeTwoPi) SetFloat64(pF64 float64) {
	sR642Pi.v.Value = pF64
}

// IsGroupA verifica si el valor pertany al Grup A.
func (sR642Pi F64RangeTwoPi) IsGroupA() bool {
	_, exponent, mantissa := tools.DecomposeF64(sR642Pi.v.Value)

	// Valors normals
	if exponent >= 1 && exponent <= 2046 {
		return sR642Pi.v.Value >= cs.Range64Configs.TwoPiF64.Min && sR642Pi.v.Value <= cs.Range64Configs.TwoPiF64.Max
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

func (sR642Pi F64RangeTwoPi) AsFloat32() float32 {
	return float32(sR642Pi.v.Value)
}

// func (sR642Pi F64RangeTwoPi) AsFloat64() float64 {
// 	return sR642Pi.v.value
// }

func (sR642Pi F64RangeTwoPi) AsUint32() uint32 {
	panic("Encara no implementat!: func (sR642Pi F64RangeTwoPi) AsUint32() uint32")
}

func (sR642Pi F64RangeTwoPi) AsUint64() uint64 {
	return tools.F64ToU64(sR642Pi.v.Value)
}

func (sR642Pi *F64RangeTwoPi) SetUint32(value uint32) {
	panic("Encara no implementat!: func (sR642Pi F64RangeTwoPi) SetUint32(value uint32)")
}

func (sR642Pi *F64RangeTwoPi) SetUint64(value uint64) {
	sR642Pi.v.Value = tools.U64ToF64(value)
}

// Add suma dos F64RangeTwoPi i retorna un nou valor saturat dins del rang.
func (sR64One F64RangeTwoPi) Add(other F64RangeTwoPi) F64RangeTwoPi {
	result := sR64One.AsFloat64() + sR64One.AsFloat64()
	if result > cs.Range64Configs.TwoPiF64.Max {
		result = cs.Range64Configs.TwoPiF64.Max
	} else if result < cs.Range64Configs.TwoPiF64.Min {
		result = cs.Range64Configs.TwoPiF64.Min
	}
	return NewF64RangeTwoPi(result)
}

// IsInfinitePos comprova si el valor és +Inf.
func (sR64One F64RangeTwoPi) IsInfinitePos() bool {
	return math.IsInf(sR64One.AsFloat64(), 1)
}

// IsInfiniteNeg comprova si el valor és -Inf.
func (sR64One F64RangeTwoPi) IsInfiniteNeg() bool {
	return math.IsInf(sR64One.AsFloat64(), -1)
}

// IsInfinite comprova si el valor és infinit (positiu o negatiu).
func (sR64One F64RangeTwoPi) IsInfinite() bool {
	return math.IsInf(sR64One.AsFloat64(), 0)
}

// GetF64Value retorna el valor float64.
func (sR64One F64RangeTwoPi) GetF64Value() float64 {
	return sR64One.AsFloat64()
}

// SetF64Value modifica el valor float64 i retorna l'instància actualitzada.
func (sR64One *F64RangeTwoPi) SetF64Value(pF64 float64) F64RangeTwoPi {
	sR64One.v.Value = tools.Quantize64One(pF64)
	return *sR64One
}
