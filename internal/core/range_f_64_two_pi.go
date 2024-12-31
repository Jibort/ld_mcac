// Implementació del tipus RangeF64TwoPi utilitzant constants consolidades
// CreatedAt: 2024-12-27 dc. GPT(JIQ)

package core

import (
	"fmt"
	"math"
)

// Tipus RangeF64TwoPi representa el rang [-2π, +2π].
type RangeF64TwoPi struct {
	value float64
}

// NewRangeF64TwoPi crea una nova instància de RangeF64TwoPi amb validació del rang.
func NewRangeF64TwoPi(value float64) (*RangeF64TwoPi, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) || (value >= RangeF64Configs.TwoPi.Min && value <= RangeF64Configs.TwoPi.Max) {
		return &RangeF64TwoPi{value: value}, nil
	}
	return nil, fmt.Errorf("valor fora del rang [-2π, +2π]: %f", value)
}

// IsGroupA verifica si el valor pertany al Grup A.
func (r RangeF64TwoPi) IsGroupA() bool {
	_, exponent, mantissa := DecomposeF64(r.value)

	// Valors normals
	if exponent >= 1 && exponent <= 2046 {
		return r.value >= RangeF64Configs.TwoPi.Min && r.value <= RangeF64Configs.TwoPi.Max
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

// Add suma dos RangeF64TwoPi i retorna un nou valor saturat dins del rang.
func (r RangeF64TwoPi) Add(other RangeF64TwoPi) RangeF64TwoPi {
	result := r.value + other.value
	if result > RangeF64Configs.TwoPi.Max {
		result = RangeF64Configs.TwoPi.Max
	} else if result < RangeF64Configs.TwoPi.Min {
		result = RangeF64Configs.TwoPi.Min
	}
	return RangeF64TwoPi{value: result}
}

// Funcions dummies per completar la interfície RangeIntf

// Equals comprova si dos valors són iguals.
func (r RangeF64TwoPi) Equals(other RangeF64TwoPi) bool {
	return r.value == other.value
}

// LessThan comprova si el valor actual és menor que un altre.
func (r RangeF64TwoPi) LessThan(other RangeF64TwoPi) bool {
	return r.value < other.value
}

// LessOrEqualThan comprova si el valor actual és menor o igual que un altre.
func (r RangeF64TwoPi) LessOrEqualThan(other RangeF64TwoPi) bool {
	return r.value <= other.value
}

// GreaterThan comprova si el valor actual és major que un altre.
func (r RangeF64TwoPi) GreaterThan(other RangeF64TwoPi) bool {
	return r.value > other.value
}

// GreaterOrEqualThan comprova si el valor actual és major o igual que un altre.
func (r RangeF64TwoPi) GreaterOrEqualThan(other RangeF64TwoPi) bool {
	return r.value >= other.value
}

// IsInfinitePos comprova si el valor és +Inf.
func (r RangeF64TwoPi) IsInfinitePos() bool {
	return math.IsInf(r.value, 1)
}

// IsInfiniteNeg comprova si el valor és -Inf.
func (r RangeF64TwoPi) IsInfiniteNeg() bool {
	return math.IsInf(r.value, -1)
}

// IsInfinite comprova si el valor és infinit (positiu o negatiu).
func (r RangeF64TwoPi) IsInfinite() bool {
	return math.IsInf(r.value, 0)
}

// GetF64Value retorna el valor float64.
func (r RangeF64TwoPi) GetF64Value() float64 {
	return r.value
}

// SetF64Value modifica el valor float64 i retorna l'instància actualitzada.
func (r *RangeF64TwoPi) SetF64Value(value float64) RangeF64TwoPi {
	r.value = value
	return *r
}

// // Fitxer per al tipus RangeF64TwoPi.
// // CreatedAt: 2024/12/27 dv. GPT(JIQ)

// package core

// import (
// 	"fmt"
// 	"math"
// )

// const (
// 	// TwoPi    = 2 * math.Pi // 2π
// 	NegTwoPi = -TwoPi // -2π
// 	PosTwoPi = TwoPi  // +2π
// )

// // RangeF64TwoPi representa valors al rang [-2π, +2π]
// type RangeF64TwoPi struct {
// 	RangeF64
// }

// // Constructor per a RangeF64TwoPi
// func NewRangeF64TwoPi(value float64) RangeF64TwoPi {
// 	// Permet valors normals, subnormals, ±Inf i NaN
// 	if math.IsInf(value, 0) || math.IsNaN(value) || (value >= NegTwoPi && value <= PosTwoPi) {
// 		return RangeF64TwoPi{RangeF64{value: value}}
// 	}

// 	panic(fmt.Sprintf("Valor fora del rang [-2π, +2π]: %v", value))
// }

// // FUNCIONS IS... ---------------------
// // Validació per al rang [-2π, +2π]
// func (r RangeF64TwoPi) IsTwoPiRange() bool {
// 	return r.GetF64Value() >= RangeNegTwoPi && r.GetF64Value() <= RangePosTwoPi
// }

// func (r RangeF64TwoPi) IsGroupA() bool {
// 	_, exponent, mantissa := DecomposeF64(r.GetF64Value())

// 	// Subgrup A.1: Valors normals dins [-2π, +2π]
// 	if exponent > 0 && exponent < 2047 {
// 		return r.GetF64Value() >= NegTwoPi && r.GetF64Value() <= PosTwoPi
// 	}

// 	// Subgrup A.2: Valors subnormals (exponent = 0, mantissa != 0)
// 	if exponent == 0 && mantissa != 0 {
// 		return true
// 	}

// 	// Subgrup A.3: ±Inf i NaN (exponent = 2047)
// 	if exponent == 2047 {
// 		return true
// 	}

// 	return false
// }

// // Implementació d'exemple per a Add dins RangeF64TwoPi
// func (r RangeF64TwoPi) Add(pOther RangeIntf) RangeIntf {
// 	result := r.GetF64Value() + pOther.ValueF64()
// 	if result > RangePosTwoPi {
// 		return NewRangeF64TwoPi(RangePosTwoPi)
// 	}
// 	if result < RangeNegTwoPi {
// 		return NewRangeF64TwoPi(RangeNegTwoPi)
// 	}
// 	return NewRangeF64TwoPi(result)
// }
