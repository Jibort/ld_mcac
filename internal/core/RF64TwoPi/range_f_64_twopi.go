// Implementació del tipus RangeF64TwoPi utilitzant constants consolidades
// CreatedAt: 2024-12-27 dc. GPT(JIQ)

package RF64One

import (
	"fmt"
	"math"

	cs "github.com/jibort/ld_mcac/internal/core/Consts"
	intf "github.com/jibort/ld_mcac/internal/core/Intf"
	rF64 "github.com/jibort/ld_mcac/internal/core/RF64"
	tools "github.com/jibort/ld_mcac/internal/core/tools"
)

// Tipus RangeF64TwoPi representa el rang [-2·π, +2·π].
type RangeF64TwoPi struct {
	intf.RangeF64TwoPiIntf
	rF64.RangeF64
	value float64
}

// NewRangeF64TwoPi crea una nova instància de RangeF64TwoPi amb validació del rang.
func NewRangeF64TwoPi(value float64) *RangeF64TwoPi {
	if math.IsNaN(value) || math.IsInf(value, 0) || (value >= cs.Range64Configs.TwoPiF64.Max && value <= cs.Range64Configs.TwoPiF64.Max) {
		return &RangeF64TwoPi{value: value}
	}
	panic(fmt.Sprintf("valor fora del rang [-2·π, +2·π]: %f", value))
}

// Equals comprova si dos valors són iguals.
func (sR642Pi RangeF64TwoPi) Equals(pOther intf.RangeIntf) bool {
	switch other := pOther.(type) {
	case RangeF64TwoPi:
		return tools.Equals64(sR642Pi.value, other.value, &cs.Epsilon64)
	default:
		return false
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
	sR642Pi.value = pF64
}

// IsGroupA verifica si el valor pertany al Grup A.
func (sR642Pi RangeF64TwoPi) IsGroupA() bool {
	_, exponent, mantissa := tools.DecomposeF64(sR642Pi.value)

	// Valors normals
	if exponent >= 1 && exponent <= 2046 {
		return sR642Pi.value >= cs.Range64Configs.TwoPiF64.Min && sR642Pi.value <= cs.Range64Configs.TwoPiF64.Max
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
	return float32(sR642Pi.value)
}

func (sR642Pi RangeF64TwoPi) AsFloat64() float64 {
	return sR642Pi.value
}

func (sR642Pi RangeF64TwoPi) AsUint32() uint32 {
	panic("Encara no implementat!: func (sR642Pi RangeF64TwoPi) AsUint32() uint32")
}

func (sR642Pi RangeF64TwoPi) AsUint64() uint64 {
	return tools.F64ToU64(sR642Pi.value)
}

func (sR642Pi RangeF64TwoPi) SetUint32(value uint32) {
	panic("Encara no implementat!: func (sR642Pi RangeF64TwoPi) SetUint32(value uint32)")
}

func (sR642Pi RangeF64TwoPi) SetUint64(value uint64) {
	sR642Pi.value = tools.U64ToF64(value)
}

func (sR642Pi RangeF64TwoPi) Clone() intf.RangeIntf {
	return *NewRangeF64TwoPi(sR642Pi.value)
}

// Add suma dos RangeF64TwoPi i retorna un nou valor saturat dins del rang.
func (sR64One RangeF64TwoPi) Add(other RangeF64TwoPi) RangeF64TwoPi {
	result := sR64One.value + other.value
	if result > cs.Range64Configs.TwoPiF64.Max {
		result = cs.Range64Configs.TwoPiF64.Max
	} else if result < cs.Range64Configs.TwoPiF64.Min {
		result = cs.Range64Configs.TwoPiF64.Min
	}
	return RangeF64TwoPi{value: result}
}

// IsInfinitePos comprova si el valor és +Inf.
func (sR64One RangeF64TwoPi) IsInfinitePos() bool {
	return math.IsInf(sR64One.value, 1)
}

// IsInfiniteNeg comprova si el valor és -Inf.
func (sR64One RangeF64TwoPi) IsInfiniteNeg() bool {
	return math.IsInf(sR64One.value, -1)
}

// IsInfinite comprova si el valor és infinit (positiu o negatiu).
func (sR64One RangeF64TwoPi) IsInfinite() bool {
	return math.IsInf(sR64One.value, 0)
}

// GetF64Value retorna el valor float64.
func (sR64One RangeF64TwoPi) GetF64Value() float64 {
	return sR64One.value
}

// SetF64Value modifica el valor float64 i retorna l'instància actualitzada.
func (sR64One *RangeF64TwoPi) SetF64Value(value float64) RangeF64TwoPi {
	sR64One.value = value
	return *sR64One
}
