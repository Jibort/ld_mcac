// Implementació dels valors Range en float64.
// CreatedAt: 2024/12/08 dg. JIQ

package core

import (
	"fmt"
	"math"
)

// Tipus pels Range float64.
type RangeF64 struct {
	Range64Intf
	value float64
}

// GETTERS/SETTERS --------------------
func (sSrc RangeF64) GetValue() float64 { return sSrc.value }

// CONSTRUCTORS -----------------------
// Constructor amb grup.
func NewRangeF64WithGroup(value float64, group int) RangeF64 {
	bits := math.Float64bits(value)
	exponent := uint64(0)

	switch group {
	case 1: // Grup A
		exponent = 1023 // o 1022
	case 2: // Grup B
		exponent = 1024
	case 3: // Grup C
		exponent = 1025
	case 4: // Grup D
		exponent = 1026
	default:
		// Manejar error o assignar un valor per defecte
	}

	// Ajusta els bits de l'exponent en la representació binària
	bits = (bits &^ (0x7FF << 52)) | (exponent << 52)

	return RangeF64{value: math.Float64frombits(bits)}
}

func NewRangeF64(pVal float64) (rF64 RangeF64) {
	return NewRangeF64WithGroup(pVal, 1)
}

// Constructor general a partir d'un float64.
func NewRangeF64__(pVal float64) (rF64 RangeF64) {
	// Primer comprovem els límits
	absVal := math.Abs(pVal)
	if absVal >= 1.0 || absVal > (1.0-SmallThreshold64) {
		return NewRangeF64Saturated(pVal >= 0)
	}

	bits := F64ToU64(pVal)
	if (bits & GroupMask) != GroupAMask {
		return RangeF64{value: pVal}
	}

	if absVal < SmallThreshold64 {
		return NewRangeF64Zero()
	}

	return RangeF64{value: Quantize64(pVal)}
}

// Constructor de RangeF64 per a 0.000_000_000_000_000
func NewRangeF64Zero() RangeF64 {
	return NewRangeF64WithGroup(float64(0.0), 1)
}

// Constructors de valors especials
func NewRangeF64Saturated(positive bool) RangeF64 {
	var bits uint64
	if positive {
		bits = GroupDMask | SaturationMask
	} else {
		bits = SignMask | GroupDMask | SaturationMask
	}
	return RangeF64{value: U64ToF64(bits)}
}

func NewRangeF64Infinite(positive bool) RangeF64 {
	bits := GroupBMask | SubgroupInfMask
	if !positive {
		bits |= SignMask
	}
	return RangeF64{value: U64ToF64(bits)} // Retornem directament sense usar NewRangeF64
}

func NewRangeF64Null() RangeF64 {
	bits := GroupBMask | SubgroupNullMask
	return RangeF64{value: U64ToF64(bits)}
}

// INTERFÍCIE 'RangeIntf' -------------
// Retorna cert només si la diferència entre els valors és inferior a Precisio64.
func (sSrc RangeF64) Equals(pOther RangeIntf) bool {
	if sSrc.IsSaturated() || pOther.IsSaturated() {
		return sSrc.IsSaturatedPos() == pOther.IsSaturatedPos() &&
			sSrc.IsSaturatedNeg() == pOther.IsSaturatedNeg()
	}
	return math.Abs(sSrc.ValueF64()-pOther.ValueF64()) < Epsilon64
}

// Retorna cert només en cas que el valor de la instància sigui inferior al de pOther.
func (sSrc RangeF64) LessThan(pOther RangeIntf) bool {
	if sSrc.IsSaturated() || pOther.IsSaturated() {
		return sSrc.IsSaturatedNeg() && !pOther.IsSaturatedNeg()
	}
	return (pOther.ValueF64() - sSrc.ValueF64()) > Epsilon64
}

func (sSrc RangeF64) LessOrEqualThan(pOther RangeIntf) bool {
	if sSrc.IsSaturated() || pOther.IsSaturated() {
		return sSrc.IsSaturatedNeg() || sSrc.Equals(pOther)
	}
	return sSrc.LessThan(pOther) || sSrc.Equals(pOther)
}

// Retorna cert només si el valor de la instància es superior al de pOther.
func (sSrc RangeF64) GreaterThan(pOther RangeIntf) bool {
	if sSrc.IsSaturated() || pOther.IsSaturated() {
		return sSrc.IsSaturatedPos() && !pOther.IsSaturatedPos()
	}
	return (sSrc.ValueF64() - pOther.ValueF64()) > Epsilon64
}

// Retorna cert només si el valor de la instància és superior o igual al de pOther.
func (sSrc RangeF64) GreaterOrEqualThan(pOther RangeIntf) bool {
	if sSrc.IsSaturated() || pOther.IsSaturated() {
		return sSrc.IsSaturatedPos() || sSrc.Equals(pOther)
	}
	return sSrc.GreaterThan(pOther) || sSrc.Equals(pOther)
}

// Valors especials.
func (sSrc RangeF64) IsNullValue() bool {
	bits := F64ToU64(sSrc.value)
	return (bits & (MetaMask | SubgroupMask)) == (GroupBMask | SubgroupNullMask)
}

func (sSrc RangeF64) IsInfinitePos() bool {
	bits := F64ToU64(sSrc.value)
	expectedBits := GroupBMask | SubgroupInfMask
	return bits == expectedBits && (bits&SignMask) == 0
}

func (sSrc RangeF64) IsInfiniteNeg() bool {
	bits := F64ToU64(sSrc.value)
	expectedBits := SignMask | GroupBMask | SubgroupInfMask
	return bits == expectedBits
}

// Grups.
func (sSrc RangeF64) IsGroupA() bool {
	bits := F64ToU64(sSrc.value)
	exponent := int((bits >> 52) & 0x7FF)
	return (exponent == 1022) || (exponent == 1023)
}

func (sSrc RangeF64) IsGroupB() bool {
	bits := F64ToU64(sSrc.value)
	exponent := int((bits >> 52) & 0x7FF)
	return (exponent == 1024)

}

func (sSrc RangeF64) IsGroupA_() bool {
	expBits := sSrc.ExtractExponent()
	return expBits == 1023 // 0b00000000_00000000 // Comprovar si els bits de l'exponent són tots zeros
}

func (sSrc RangeF64) IsGroupB_() bool {
	expBits := sSrc.ExtractExponent()
	return expBits == 1024 // 0b00000000_00000001
}

func (sSrc RangeF64) IsGroupC_() bool {
	bits := F64ToU64(sSrc.value)
	return ((bits<<1)>>1)&GroupMask == GroupCMask
}

func (sSrc RangeF64) IsGroupD_() bool {
	bits := F64ToU64(sSrc.value)
	return ((bits<<1)>>1)&GroupMask == GroupDMask
}

func (sSrc RangeF64) IsGroupC() bool {
	bits := F64ToU64(sSrc.value)
	exponent := int((bits >> 52) & 0x7FF)
	return (exponent == 1025)
}

func (sSrc RangeF64) IsGroupD() bool {
	bits := F64ToU64(sSrc.value)
	exponent := int((bits >> 52) & 0x7FF)
	return (exponent == 1026)
}

func (sSrc RangeF64) ClassifyBlock() int {
	bits := F64ToU64(sSrc.value)
	exponent := int((bits >> 52) & 0x7FF)
	exponentReal := exponent - 1023

	switch {
	case exponentReal >= 1:
		return 1
	case exponentReal == 0:
		return 2
	case exponentReal == -1:
		return 3
	case exponentReal <= -2:
		return 4
	default:
		return 0 // Cas inesperat
	}
}

// TODO: Per a debugar.
func (sSrc RangeF64) DebugBits() string {
	bits := F64ToU64(sSrc.value)
	return fmt.Sprintf("Bits: %064b\nGroup: %02b", bits, (bits&GroupMask)>>62)
}

// Operacions
// Suma dos valors Range mantenint el resultat dins [-1.0, +1.0]
func (sSrc RangeF64) Add(pOther RangeIntf) RangeIntf {
	result := sSrc.value + pOther.ValueF64()

	// Si el resultat supera els límits, retornem valor saturat
	if result > 1.0 {
		return NewRangeF64Saturated(true)
	}
	if result < -1.0 {
		return NewRangeF64Saturated(false)
	}

	return NewRangeF64(result)
}

// Resta dos valors Range mantenint el resultat dins [-1.0, +1.0]
func (sSrc RangeF64) Sub(pOther RangeIntf) RangeIntf {
	result := sSrc.value - pOther.ValueF64()

	// Si el resultat supera els límits, retornem valor saturat
	if result > 1.0 {
		return NewRangeF64Saturated(true)
	}
	if result < -1.0 {
		return NewRangeF64Saturated(false)
	}

	return NewRangeF64(result)
}

// Multiplica dos valors Range mantenint el resultat dins [-1.0, +1.0]
func (sSrc RangeF64) Mul(pOther RangeIntf) RangeIntf {
	result := sSrc.value * pOther.ValueF64()

	// Si el resultat supera els límits, retornem valor saturat
	if result > 1.0 {
		return NewRangeF64Saturated(true)
	}
	if result < -1.0 {
		return NewRangeF64Saturated(false)
	}

	return NewRangeF64(result)
}

// Divideix dos valors Range mantenint el resultat dins [-1.0, +1.0]
func (sSrc RangeF64) Div(pOther RangeIntf) (RangeIntf, error) {
	// Comprovem divisió per zero
	if math.Abs(pOther.ValueF64()) < SmallThreshold64 {
		return sSrc, fmt.Errorf("divisió per zero")
	}

	result := sSrc.value / pOther.ValueF64()

	// Si el resultat supera els límits, retornem valor saturat
	if result > 1.0 {
		return NewRangeF64Saturated(true), nil
	}
	if result < -1.0 {
		return NewRangeF64Saturated(false), nil
	}

	return NewRangeF64(result), nil
}

// ... i altres operacions comunes necessàries.

// Funcions de desencapçulament.
func (sSrc RangeF64) ValueF64() float64 {
	return F64ToF64(sSrc.value)
}

func (sSrc RangeF64) ValueF32() float32 {
	return F64ToF32(sSrc.value)
}

func (sSrc RangeF64) ValueI64() int64 {
	return F64ToI64(sSrc.value)
}

func (sSrc RangeF64) ValueI32() int32 {
	return F64ToI32(sSrc.value)
}

func (sSrc RangeF64) ValueU64() uint64 {
	return F64ToU64(sSrc.value)
}

func (sSrc RangeF64) ValueU32() uint32 {
	return F64ToU32(sSrc.value)
}

// Aquestes funcions poden ser útils en un futur.
func (sSrc RangeF64) AsF64() RangeF64 { return sSrc }

// func (sSrc RangeF64) AsF32() RangeF32 { return sSrc }
// func (sSrc RangeF64) AsI64() RangeI64 { return sSrc }
// func (sSrc RangeF64) AsI32() RangeI32 { return sSrc }
// func (sSrc RangeF64) AsU64() RangeU64 { return sSrc }
// func (sSrc RangeF64) AsU32() RangeU32 { return sSrc }

// Funcions de saturació
func (sSrc RangeF64) IsSaturated() bool {
	return sSrc.IsSaturatedPos() || sSrc.IsSaturatedNeg()
}

func (sSrc RangeF64) IsSaturatedPos() bool {
	bits := F64ToU64(sSrc.value)
	return (bits & (SignMask | GroupMask)) == GroupDMask // No signe + Group D
}

func (sSrc RangeF64) IsSaturatedNeg() bool {
	bits := F64ToU64(sSrc.value)
	return (bits & (SignMask | GroupMask)) == (SignMask | GroupDMask) // Amb signe + Group D
}
