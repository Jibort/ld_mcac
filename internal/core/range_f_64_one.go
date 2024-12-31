// Implementació del tipus RangeF64One utilitzant constants consolidades
// CreatedAt: 2024-12-27 dc. GPT(JIQ)

package core

import (
	"fmt"
	"math"
)

// Tipus RangeF64One representa el rang [-1.0, +1.0].
type RangeF64One struct {
	value float64
}

// NewRangeF64One crea una nova instància de RangeF64One amb validació del rang.
func NewRangeF64One(value float64) (*RangeF64One, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) || (value >= RangeF64Configs.One.Min && value <= RangeF64Configs.One.Max) {
		return &RangeF64One{value: value}, nil
	}
	return nil, fmt.Errorf("valor fora del rang [-1.0, +1.0]: %f", value)
}

// IsGroupA verifica si el valor pertany al Grup A.
func (r RangeF64One) IsGroupA() bool {
	_, exponent, mantissa := DecomposeF64(r.value)

	// Valors normals
	if exponent >= 1 && exponent <= 2046 {
		return r.value >= RangeF64Configs.One.Min && r.value <= RangeF64Configs.One.Max
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

// Add suma dos RangeF64One i retorna un nou valor saturat dins del rang.
func (r RangeF64One) Add(other RangeF64One) RangeF64One {
	result := r.value + other.value
	if result > RangeF64Configs.One.Max {
		result = RangeF64Configs.One.Max
	} else if result < RangeF64Configs.One.Min {
		result = RangeF64Configs.One.Min
	}
	return RangeF64One{value: result}
}

// Funcions dummies per completar la interfície RangeIntf

// Equals comprova si dos valors són iguals.
func (r RangeF64One) Equals(other RangeF64One) bool {
	return r.value == other.value
}

// LessThan comprova si el valor actual és menor que un altre.
func (r RangeF64One) LessThan(other RangeF64One) bool {
	return r.value < other.value
}

// LessOrEqualThan comprova si el valor actual és menor o igual que un altre.
func (r RangeF64One) LessOrEqualThan(other RangeF64One) bool {
	return r.value <= other.value
}

// GreaterThan comprova si el valor actual és major que un altre.
func (r RangeF64One) GreaterThan(other RangeF64One) bool {
	return r.value > other.value
}

// GreaterOrEqualThan comprova si el valor actual és major o igual que un altre.
func (r RangeF64One) GreaterOrEqualThan(other RangeF64One) bool {
	return r.value >= other.value
}

// IsInfinitePos comprova si el valor és +Inf.
func (r RangeF64One) IsInfinitePos() bool {
	return math.IsInf(r.value, 1)
}

// IsInfiniteNeg comprova si el valor és -Inf.
func (r RangeF64One) IsInfiniteNeg() bool {
	return math.IsInf(r.value, -1)
}

// IsInfinite comprova si el valor és infinit (positiu o negatiu).
func (r RangeF64One) IsInfinite() bool {
	return math.IsInf(r.value, 0)
}

// GetF64Value retorna el valor float64.
func (r RangeF64One) GetF64Value() float64 {
	return r.value
}

// SetF64Value modifica el valor float64 i retorna l'instància actualitzada.
func (r *RangeF64One) SetF64Value(value float64) RangeF64One {
	r.value = value
	return *r
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
// func (r RangeF64One) IsOneRange() bool {
// 	return r.GetF64Value() >= MinValue && r.GetF64Value() <= MaxValue
// }

// func (r RangeF64One) IsGroupA() bool {
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
// func (r RangeF64One) Add(pOther RangeIntf) RangeIntf {
// 	result := r.GetF64Value() + pOther.ValueF64()
// 	if result > MaxValue {
// 		return NewRangeF64One(MaxValue)
// 	}
// 	if result < MinValue {
// 		return NewRangeF64One(MinValue)
// 	}
// 	return NewRangeF64One(result)
// }
