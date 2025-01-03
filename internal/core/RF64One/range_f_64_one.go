// Implementació del tipus RangeF64One utilitzant constants consolidades
// CreatedAt: 2024-12-27 dc. GPT(JIQ)

package RF64One

import (
	"fmt"
	"math"

	cs "github.com/jibort/ld_mcac/internal/core/Consts"
	intf "github.com/jibort/ld_mcac/internal/core/Intf"
	tools "github.com/jibort/ld_mcac/internal/core/tools"
)

// Tipus RangeF64One representa el rang [-1.0, +1.0].
type RangeF64One struct {
	intf.RangeF64OneIntf
	value float64
}

// NewRangeF64One crea una nova instància de RangeF64One amb validació del rang.
func NewRangeF64One(value float64) *RangeF64One {
	if math.IsNaN(value) || math.IsInf(value, 0) || (value >= cs.Range64Configs.OneF64.Max && value <= cs.Range64Configs.OneF64.Max) {
		return &RangeF64One{value: value}
	}
	panic(fmt.Sprintf("valor fora del rang [-1.0, +1.0]: %f", value))
}

func (sR64One RangeF64One) Equals(pOther intf.RangeIntf) bool {
	switch other := pOther.(type) {
	case RangeF64One:
		return tools.Equals64(sR64One.value, other.value, &cs.Epsilon64)
	default:
		return false
	}
}

func (sR64One RangeF64One) Is32() bool     { return false }
func (sR64One RangeF64One) Is64() bool     { return true }
func (sR64One RangeF64One) IsGroupB() bool { return false }
func (sR64One RangeF64One) IsGroupC() bool { return false }
func (sR64One RangeF64One) IsGroupD() bool { return false }
func (sR64One RangeF64One) SetFloat64(pF64 float64) {
	sR64One.value = pF64
}

// IsGroupA verifica si el valor pertany al Grup A.
func (sR64One RangeF64One) IsGroupA() bool {
	_, exponent, mantissa := tools.DecomposeF64(sR64One.value)

	// Valors normals
	if exponent >= 1 && exponent <= 2046 {
		return sR64One.value >= cs.Range64Configs.OneF64.Min && sR64One.value <= cs.Range64Configs.OneF64.Max
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

func (sR64One RangeF64One) AsFloat32() float32 {
	return float32(sR64One.value)
}

func (sR64One RangeF64One) AsFloat64() float64 {
	return sR64One.value
}

func (sT64One RangeF64One) AsUint32() uint32 {
	panic("Encara no implementat!: func (sT64One RangeF64One) AsUint32() uint32")
}

func (sT64One RangeF64One) AsUint64() uint64 {
	return tools.F64ToU64(sT64One.value)
}

func (sT64One RangeF64One) SetUint32(value uint32) {
}

func (sT64One RangeF64One) SetUint64(value uint64) {
}

func (sT64One RangeF64One) Clone() intf.RangeIntf {
	return NewRangeF64One(sT64One.value)
}

// Add suma dos RangeF64One i retorna un nou valor saturat dins del rang.
func (sR64One RangeF64One) Add(other RangeF64One) RangeF64One {
	result := sR64One.value + other.value
	if result > cs.Range64Configs.OneF64.Max {
		result = cs.Range64Configs.OneF64.Max
	} else if result < cs.Range64Configs.OneF64.Min {
		result = cs.Range64Configs.OneF64.Min
	}
	return RangeF64One{value: result}
}

// Funcions dummies per completar la interfície RangeIntf

// LessThan comprova si el valor actual és menor que un altre.
func (sR64One RangeF64One) LessThan(other RangeF64One) bool {
	return sR64One.value < other.value
}

// LessOrEqualThan comprova si el valor actual és menor o igual que un altre.
func (sR64One RangeF64One) LessOrEqualThan(other RangeF64One) bool {
	return sR64One.value <= other.value
}

// GreaterThan comprova si el valor actual és major que un altre.
func (sR64One RangeF64One) GreaterThan(other RangeF64One) bool {
	return sR64One.value > other.value
}

// GreaterOrEqualThan comprova si el valor actual és major o igual que un altre.
func (sR64One RangeF64One) GreaterOrEqualThan(other RangeF64One) bool {
	return sR64One.value >= other.value
}

// IsInfinitePos comprova si el valor és +Inf.
func (sR64One RangeF64One) IsInfinitePos() bool {
	return math.IsInf(sR64One.value, 1)
}

// IsInfiniteNeg comprova si el valor és -Inf.
func (sR64One RangeF64One) IsInfiniteNeg() bool {
	return math.IsInf(sR64One.value, -1)
}

// IsInfinite comprova si el valor és infinit (positiu o negatiu).
func (sR64One RangeF64One) IsInfinite() bool {
	return math.IsInf(sR64One.value, 0)
}

// GetF64Value retorna el valor float64.
func (sR64One RangeF64One) GetF64Value() float64 {
	return sR64One.value
}

// SetF64Value modifica el valor float64 i retorna l'instància actualitzada.
func (sR64One *RangeF64One) SetF64Value(value float64) RangeF64One {
	sR64One.value = value
	return *sR64One
}

// // Fitxer per al tipus RangeF64One.
// // CreatedAt: 2024/12/27 dv. GPT(JIQ)

// package core

// import (
// 	"fmt"
// 	"math"
// )

// // RangeF64One representa valors al rang [-1.0, +1.0]
// type RangeF64One struct {
// 	RangeF64
// }

// // CONSTRUCTORS -----------------------
// // Constructor per a RangeF64One
// func NewRangeF64One(value float64) RangeF64One {
// 	// Valida valors normals dins del rang [-1.0, +1.0]
// 	if math.IsInf(value, 0) || math.IsNaN(value) || (value >= -1.0 && value <= 1.0) {
// 		return RangeF64One{RangeF64{value: value}}
// 	}

// 	panic(fmt.Sprintf("Valor fora del rang [-1.0, +1.0]: %v", value))
// }

// // FUNCIONS IS... ---------------------
// // Validació per al rang [-1.0, +1.0]
// func (sR64One RangeF64One) IsOneRange() bool {
// 	return r.GetF64Value() >= MinValue && r.GetF64Value() <= MaxValue
// }

// func (sR64One RangeF64One) IsGroupA() bool {
// 	// bits := F64ToU64(r.GetF64Value())
// 	_, exponent, mantissa := DecomposeF64(r.GetF64Value())

// 	// Grup A.1: Valors normalitzats amb exponent entre 1 i 2046
// 	if exponent >= 1 && exponent <= 2046 {
// 		return true
// 	}

// 	// Grup A.2: Valors subnormalitzats (exponent = 0, mantissa != 0)
// 	if exponent == 0 && mantissa != 0 {
// 		return true
// 	}

// 	// Grup A.3: ±Inf i NaN (exponent = 2047)
// 	if exponent == 2047 {
// 		return true
// 	}

// 	return false
// }

// // Implementació d'exemple per a Add dins RangeF64One
// func (sR64One RangeF64One) Add(pOther RangeIntf) RangeIntf {
// 	result := r.GetF64Value() + pOther.ValueF64()
// 	if result > MaxValue {
// 		return NewRangeF64One(MaxValue)
// 	}
// 	if result < MinValue {
// 		return NewRangeF64One(MinValue)
// 	}
// 	return NewRangeF64One(result)
// }
