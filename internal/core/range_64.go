// Package core proporciona els tipus i operacions fonamentals per la implementació MCAC.
// CreatedAt: 2024/12/08 dg CLD

package core

import (
	"fmt"
	"math"
)

// Range64 representa un tipus especialitzat de 64 bits que combina:
// - Un bit de signe (63)
// - Bits de metadades (62-49)
// - Un valor decimal en el rang [-1.0, +1.0] (48-0)
type Range64 uint64

// CONSTRUCTORS -----------------------------
// NewRange64 crea un nou Range64 a partir d'un valor float64
func NewRange64(pVal float64) (rRange Range64) {
	// Limita el valor a [-1.0, 1.0]
	pVal = math.Max(MinValue, math.Min(MaxValue, pVal))

	// Converteix a uint64
	var result uint64

	if pVal < 0 {
		result |= SignMask
		pVal = -pVal
	}

	// Escala el valor per utilitzar la precisió completa de 48 bits
	scaledVal := pVal * float64(ValueMask & ^SignMask)
	result |= uint64(scaledVal) & (ValueMask & ^SignMask)

	return Range64(result)
}

// FromBits crea un Range64 directament des d'una representació uint64
func FromBits(pBits uint64) (rRange Range64) {
	return Range64(pBits)
}

// Zero retorna un Range64 amb valor 0.0
func Zero() (rRange Range64) {
	return Range64(0)
}

// One retorna un Range64 amb valor 1.0
func One() (rRange Range64) {
	return NewRange64(1.0)
}

// CONVERSIONS -----------------------------
// ToFloat64 converteix un Range64 a float64
func (sRange Range64) ToFloat64() (rVal float64) {
	return float64(int64(uint64(sRange)&ValueMask)) / float64(ValueMask & ^SignMask)
}

// ToBits retorna la representació uint64 d'un Range64
func (sRange Range64) ToBits() (rBits uint64) {
	return uint64(sRange)
}

// String retorna una representació en text del Range64
func (sRange Range64) String() (rStr string) {
	return fmt.Sprintf("Range64{valor: %v, meta: %014b}", sRange.ToFloat64(), sRange.GetMeta())
}

// OPERACIONS DE METADADES ----------------
// SetMeta estableix els bits de metadades
func (sRange Range64) SetMeta(pMeta uint64) (rRange Range64) {
	clearedMeta := uint64(sRange) & ^MetaMask
	return Range64(clearedMeta | ((pMeta << MetaShift) & MetaMask))
}

// GetMeta obté els bits de metadades
func (sRange Range64) GetMeta() (rMeta uint64) {
	return (uint64(sRange) & MetaMask) >> MetaShift
}

// OPERACIONS DE VALOR -------------------
// IsNegative comprova si el valor és negatiu
func (sRange Range64) IsNegative() (rIsNeg bool) {
	return (uint64(sRange) & SignMask) != 0
}

// Abs retorna el valor absolut
func (sRange Range64) Abs() (rRange Range64) {
	return Range64(uint64(sRange) & ^SignMask)
}

// Neg retorna la negació del valor
func (sRange Range64) Neg() (rRange Range64) {
	return Range64(uint64(sRange) ^ SignMask)
}

// OPERACIONS ARITMÈTIQUES ---------------
// Add suma dos valors Range64
func (sRange Range64) Add(pOther Range64) (rRange Range64) {
	return NewRange64(sRange.ToFloat64() + pOther.ToFloat64())
}

// Sub resta dos valors Range64
func (sRange Range64) Sub(pOther Range64) (rRange Range64) {
	return NewRange64(sRange.ToFloat64() - pOther.ToFloat64())
}

// Mul multiplica dos valors Range64
func (sRange Range64) Mul(pOther Range64) (rRange Range64) {
	return NewRange64(sRange.ToFloat64() * pOther.ToFloat64())
}

// Div divideix dos valors Range64
func (sRange Range64) Div(pOther Range64) (rRange Range64) {
	if math.Abs(pOther.ToFloat64()) < Epsilon64 {
		if sRange.IsNegative() != pOther.IsNegative() {
			return NewRange64(MinValue)
		}
		return NewRange64(MaxValue)
	}
	return NewRange64(sRange.ToFloat64() / pOther.ToFloat64())
}

// OPERACIONS AVANÇADES ------------------
// Lerp realitza una interpolació lineal entre dos valors Range64
func (sRange Range64) Lerp(pOther Range64, pT Range64) (rRange Range64) {
	t64 := pT.ToFloat64()
	// Limita t a [0,1]
	t64 = math.Max(0, math.Min(1, t64))
	return NewRange64(sRange.ToFloat64()*(1-t64) + pOther.ToFloat64()*t64)
}

// Distance calcula la diferència absoluta entre dos valors Range64
func (sRange Range64) Distance(pOther Range64) (rRange Range64) {
	return NewRange64(math.Abs(sRange.ToFloat64() - pOther.ToFloat64()))
}

// Equals comprova si dos valors Range64 són iguals dins d'Epsilon
func (sRange Range64) Equals(pOther Range64) (rEquals bool) {
	return math.Abs(sRange.ToFloat64()-pOther.ToFloat64()) < Epsilon64
}

// FUNCIONS D'ACTIVACIÓ ------------------
// Sigmoid aplica la funció sigmoide al valor Range64
func (sRange Range64) Sigmoid() (rRange Range64) {
	return NewRange64(1.0 / (1.0 + math.Exp(-sRange.ToFloat64())))
}
