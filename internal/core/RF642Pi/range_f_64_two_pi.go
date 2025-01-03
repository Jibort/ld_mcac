// Implementació del tipus RangeF64TwoPi utilitzant constants consolidades
// CreatedAt: 2024-12-27 dc. GPT(JIQ)

package RF642Pi

import (
	"fmt"
	"math"

	cs "github.com/jibort/ld_mcac/internal/core/Consts"
	tools "github.com/jibort/ld_mcac/internal/core/tools"
)

// Tipus RangeF64TwoPi representa el rang [-2π, +2π].
type RangeF64TwoPi struct {
	value float64
}

// NewRangeF64TwoPi crea una nova instància de RangeF64TwoPi amb validació del rang.
func NewRangeF64TwoPi(value float64) (*RangeF64TwoPi, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) || (value >= cs.Range64Configs.TwoPiF64.Min && value <= cs.Range64Configs.TwoPiF64.Max) {
		return &RangeF64TwoPi{value: value}, nil
	}
	return nil, fmt.Errorf("valor fora del rang [-2π, +2π]: %f", value)
}

// IsGroupA verifica si el valor pertany al Grup A.
func (sF64 RangeF64TwoPi) IsGroupA() bool {
	_, exponent, mantissa := tools.DecomposeF64(sF64.value)

	// Valors normals
	if exponent >= 1 && exponent <= 2046 {
		return sF64.value >= cs.Range64Configs.TwoPiF64.Min && sF64.value <= cs.Range64Configs.TwoPiF64.Max
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
	if result > cs.Range64Configs.TwoPiF64.Max {
		result = cs.Range64Configs.TwoPiF64.Max
	} else if result < cs.Range64Configs.TwoPiF64.Min {
		result = cs.Range64Configs.TwoPiF64.Min
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
