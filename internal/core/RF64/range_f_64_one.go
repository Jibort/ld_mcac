// Implementació del tipus RangeF64One utilitzant constants consolidades
// CreatedAt: 2024-12-27 dc. GPT(JIQ)

package RF64

import (
	"fmt"
	"math"

	cs "github.com/jibort/ld_mcac/internal/core/Consts"
	intf "github.com/jibort/ld_mcac/internal/core/intf/ranges"
	tools "github.com/jibort/ld_mcac/internal/core/tools"
)

// Tipus F64RangeOne representa el rang [-1.0, +1.0].
type F64RangeOne struct {
	intf.F64RangeOneIntf
}

// CONSTRUCTORS ------------------------
// NewRangeF64One crea una nova instància de RangeF64One amb validació del rang.
func NewRangeF64One(value float64) F64RangeOne {
	if math.IsNaN(value) || math.IsInf(value, 0) || (value >= cs.Range64Configs.OneF64.Max && value <= cs.Range64Configs.OneF64.Max) {
		return F64RangeOne{v: RangeF64{value: value}}
	}
	panic(fmt.Sprintf("NewRangeF64One: valor fora del rang [-1.0, +1.0]: %f", value))
}

// INTERFÍCIE '
// INTERFÍCIE 'ClonableIntf' ----------
// Clone retorna una còpia de la instància.
func (sR64One F64RangeOne) Clone() intf.ClonableIntf {
	return NewRangeF64TwoPi(sR64One.v.value)
}

// INTERFÍCIE 'ComparableIntf' --------
// Equals comprova si dos valors són iguals.
func (sR64One F64RangeOne) Equals(pOther intf.ComparableIntf) bool {
	switch other := pOther.(type) {
	case F64RangeOne:
		return tools.Equals64(sR64One.AsFloat64(), other.AsFloat64(), &cs.Epsilon64)
	case RangeF64TwoPi:
		return tools.Equals64(sR64One.AsFloat64(), other.AsFloat64(), &cs.Epsilon64)
	default:
		panic(fmt.Sprintf("RangeF64One.Equals: tipus no vàlid: %T", pOther))
	}
}

// Cert només si la instància és menor que pOther.
func (sR64One F64RangeOne) LessThan(pOther intf.ComparableIntf) bool {
	switch other := pOther.(type) {
	case F64RangeOne:
		return sR64One.AsFloat64() < other.AsFloat64()
	case RangeF64TwoPi:
		return sR64One.AsFloat64() < other.AsFloat64()
	default:
		panic(fmt.Sprintf("RangeF64One.LessThan: tipus no vàlid: %T", pOther))
	}
}

// Cert només si la instància és menor o igual que pOther.
func (sR64One F64RangeOne) LessOrEqualThan(pOther intf.ComparableIntf) bool {
	switch other := pOther.(type) {
	case F64RangeOne:
		return sR64One.LessThan(other) || sR64One.Equals(other)
	case RangeF64TwoPi:
		return sR64One.LessThan(other) || sR64One.Equals(other)
	default:
		panic(fmt.Sprintf("RangeF64One.LessOrEqualThan: tipus no vàlid: %T", pOther))
	}
}

// Cert només si la instància és major que pOther.
func (sR64One F64RangeOne) GreaterThan(pOther intf.ComparableIntf) bool {
	switch other := pOther.(type) {
	case F64RangeOne:
		return sR64One.AsFloat64() > other.AsFloat64()
	case RangeF64TwoPi:
		return sR64One.AsFloat64() > other.AsFloat64()
	default:
		panic(fmt.Sprintf("RangeF64One.GreaterThan: tipus no vàlid: %T", pOther))
	}
}

// Cert només si la instància és major o igual que pOther.
func (sR64One F64RangeOne) GreaterOrEqualThan(pOther intf.ComparableIntf) bool {
	switch other := pOther.(type) {
	case F64RangeOne:
		return sR64One.LessThan(other) || sR64One.Equals(other)
	case RangeF64TwoPi:
		return sR64One.LessThan(other) || sR64One.Equals(other)
	default:
		panic(fmt.Sprintf("RangeF64One.GreaterOrEqualThan: tipus no vàlid: %T", pOther))
	}
}

// INTERFÍCIE 'TypeConversionsIntf' ---
func (sR64One F64RangeOne) Is32() bool     { return false }
func (sR64One F64RangeOne) Is64() bool     { return true }
func (sR64One F64RangeOne) IsGroupB() bool { return false }
func (sR64One F64RangeOne) IsGroupC() bool { return false }
func (sR64One F64RangeOne) IsGroupD() bool { return false }
func (sR64One F64RangeOne) SetFloat64(pF64 float64) {
	sR64One.v.value = pF64
}

// INTERFÍCIE 'RangeIntf' -------------
// IsGroupA verifica si el valor pertany al Grup A.
func (sR64One F64RangeOne) IsGroupA() bool {
	_, exponent, mantissa := tools.DecomposeF64(sR64One.v.value)

	// Valors normals
	if exponent >= 1 && exponent <= 2046 {
		return sR64One.v.value >= cs.Range64Configs.OneF64.Min && sR64One.v.value <= cs.Range64Configs.OneF64.Max
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

func (sR64One F64RangeOne) AsFloat32() float32 {
	return float32(sR64One.v.value)
}

func (sR64One F64RangeOne) AsFloat64() float64 {
	return sR64One.v.value
}

func (sT64One F64RangeOne) AsUint32() uint32 {
	panic("RangeF64One) AsUint32(): TODO: Encara no implementada!")
}

func (sT64One F64RangeOne) AsUint64() uint64 {
	return tools.F64ToU64(sT64One.v.value)
}

func (sT64One F64RangeOne) SetUint32(value uint32) {
	sT64One.SetF64Value(float64(value))
}

func (sT64One F64RangeOne) SetUint64(value uint64) {
	sT64One.SetF64Value(tools.U64ToF64(value))
}

// INTERFÍCIE 'MathOperationsIntf' ----
// Add suma dos RangeF64One i retorna un nou valor saturat dins del rang.
func (sR64One F64RangeOne) Add(other F64RangeOne) F64RangeOne {
	result := sR64One.v.value + other.v.value
	if result > cs.Range64Configs.OneF64.Max {
		result = cs.Range64Configs.OneF64.Max
	} else if result < cs.Range64Configs.OneF64.Min {
		result = cs.Range64Configs.OneF64.Min
	}
	return NewRangeF64One(result)
}

// INTERFÍCIE 'RangeF64OneIntf' -------
// Cert només si el valor és negatiu.
func (sR64One F64RangeOne) Sign() bool {
	return sR64One.v.value < 0.0
}

// IsInfinitePos comprova si el valor és +Inf.
func (sR64One F64RangeOne) IsInfinitePos() bool {
	return math.IsInf(sR64One.v.value, 1)
}

// IsInfiniteNeg comprova si el valor és -Inf.
func (sR64One F64RangeOne) IsInfiniteNeg() bool {
	return math.IsInf(sR64One.v.value, -1)
}

// IsInfinite comprova si el valor és infinit (positiu o negatiu).
func (sR64One F64RangeOne) IsInfinite() bool {
	return math.IsInf(sR64One.v.value, 0)
}

// SetF64Value modifica el valor float64 i retorna l'instància actualitzada.
func (sR64One *F64RangeOne) SetF64Value(value float64) F64RangeOne {
	sR64One.v.value = value
	return *sR64One
}
