// Ampliació del rang del Grup A a [-2π, +2π]
// CreatedAt: 2024/12/08 dg. JIQ

package core

import (
	"fmt"
	"math"
)

)

// Ampliació de les funcions constructors del Grup A

// Crea un valor RangeF64 dins del rang [-2π, +2π]
func NewRangeF64TwoPi(pF64 float64) RangeF64 {
	if pF64 < RangeNegTwoPi || pF64 > RangePosTwoPi {
		panic("NewRangeF64TwoPi: valor fora del rang [-2π, +2π]")
	}

	return RangeF64{value: pF64}
}

// Crea un valor RangeF64 a partir d'un uint64 amb suport per a [-2π, +2π]
func NewRangeF64TwoPiFromU64(pU64 uint64) RangeF64 {
	value := U64ToF64(pU64)
	if value < RangeNegTwoPi || value > RangePosTwoPi {
		panic("NewRangeF64TwoPiFromU64: valor fora del rang [-2π, +2π]")
	}

	return RangeF64{value: value}
}

// Validació si un valor pertany al rang ampliat de [-2π, +2π]
func (sF64 RangeF64) IsTwoPiRange() bool {
	value := sF64.GetF64Value()
	return value >= RangeNegTwoPi && value <= RangePosTwoPi
}

// Implementació de les operacions matemàtiques dins del nou rang

// Suma dos valors Range mantenint el resultat dins [-2π, +2π]
func (sF64 RangeF64) AddTwoPi(pOther RangeIntf) RangeIntf {
	result := sF64.value + pOther.ValueF64()

	// Si el resultat supera els límits, retornem valor saturat
	if result > RangePosTwoPi {
		return NewRangeF64(RangePosTwoPi)
	}
	if result < RangeNegTwoPi {
		return NewRangeF64(RangeNegTwoPi)
	}

	return NewRangeF64(result)
}

// Resta dos valors Range mantenint el resultat dins [-2π, +2π]
func (sF64 RangeF64) SubTwoPi(pOther RangeIntf) RangeIntf {
	result := sF64.value - pOther.ValueF64()

	// Si el resultat supera els límits, retornem valor saturat
	if result > RangePosTwoPi {
		return NewRangeF64(RangePosTwoPi)
	}
	if result < RangeNegTwoPi {
		return NewRangeF64(RangeNegTwoPi)
	}

	return NewRangeF64(result)
}

// Multiplica dos valors Range mantenint el resultat dins [-2π, +2π]
func (sF64 RangeF64) MulTwoPi(pOther RangeIntf) RangeIntf {
	result := sF64.value * pOther.ValueF64()

	// Si el resultat supera els límits, retornem valor saturat
	if result > RangePosTwoPi {
		return NewRangeF64(RangePosTwoPi)
	}
	if result < RangeNegTwoPi {
		return NewRangeF64(RangeNegTwoPi)
	}

	return NewRangeF64(result)
}

// Divideix dos valors Range mantenint el resultat dins [-2π, +2π]
func (sF64 RangeF64) DivTwoPi(pOther RangeIntf) (RangeIntf, error) {
	// Comprovem divisió per zero
	if math.Abs(pOther.ValueF64()) < SmallThreshold64 {
		return sF64, fmt.Errorf("divisió per zero")
	}

	result := sF64.value / pOther.ValueF64()

	// Si el resultat supera els límits, retornem valor saturat
	if result > RangePosTwoPi {
		return NewRangeF64(RangePosTwoPi), nil
	}
	if result < RangeNegTwoPi {
		return NewRangeF64(RangeNegTwoPi), nil
	}

	return NewRangeF64(result), nil
}
