// range_f_64.go
// Implementació dels valors Range en float64.
// CreatedAt: 2024/12/08 dg. JIQ

package RangeF64

import (
	"math"

	cs "github.com/jibort/ld_mcac/internal/core/Consts"
	errs "github.com/jibort/ld_mcac/internal/core/Errors"
	intf "github.com/jibort/ld_mcac/internal/core/Intf"
	rF32 "github.com/jibort/ld_mcac/internal/core/RF32"
	rF64One "github.com/jibort/ld_mcac/internal/core/RF64One"
	rF64TwoPi "github.com/jibort/ld_mcac/internal/core/RF64TwoPi"
	tools "github.com/jibort/ld_mcac/internal/core/tools"
)

// ESTRUCTURES ------------------------
// Tipus pels Range float64.
type RangeF64 struct {
	intf.Range64Intf
	value float64
}

// CONSTRUCTORS -----------------------
// Constructor general a partir d'un valor float64 (64bits).
func NewRangeF64(pF64 float64) RangeF64 {
	return RangeF64{value: pF64}
}

// Constructor a partir d'un uint64 (64bits).
func NewRangeF64FromU64(pU64 uint64) RangeF64 {
	return RangeF64{value: tools.U64ToF64(pU64)}
}

// Constructor de RangeF64 per a 0.0.
func NewRangeF64Zero() RangeF64 {
	return NewRangeF64FromU64(cs.IEEE754ZeroBits)
}

func NewRangeF64Null() RangeF64 {
	// Basat en la documentació del Grup B, Subgrup ±Saturació i Nul
	// Assignem el patró fixat per a valors nul·ls
	bits := cs.Range64Configs.Groups.B | cs.Range64NullMask

	return NewRangeF64FromU64(bits)
}

// Crea valors saturats
func NewRangeF64Saturated(value float64, isNull bool) RangeF64 {
	var u64 uint64

	if value < 0 {
		u64 |= cs.Saturation64Mask | cs.Sign64Mask
	} else {
		u64 |= cs.Saturation64Mask
	}

	if isNull {
		u64 |= cs.NullFlag64Mask
	} else if value == 1.0 || value == -1.0 {
		u64 |= cs.UnitFlag64Mask
	}

	mantissa := ExtractFloat64tMantissa(value)
	u64 |= mantissa

	return NewRangeF64FromU64(u64)
}

// Crea ±infinit
func NewRangeF64Infinite(pIsPositive bool) RangeF64 {
	bits := cs.Range64Configs.Groups.B | 0x0000000000000001
	if !pIsPositive {
		bits |= cs.Sign64Mask
	}
	return NewRangeF64FromU64(bits)
}

// Crea un símbol
func NewRangeF64FromSymbol(pSym rune) RangeF64 {
	sym := uint64(pSym)
	u64 := cs.Range64Configs.Groups.C | sym

	return NewRangeF64FromU64(u64)
}

func NewRangeF64Percentage(pF64 float64) intf.Range64Intf {
	if pF64 < 0.0 || pF64 > 1.0 {
		return errs.NewRange64Error(true, cs.ErrCode_InvalidPercentage, 0)
	}

	valueBits := uint64(math.Round(pF64 * (1 << 52)))
	valueBits &= cs.Value64Mask

	finalBits := cs.Range64Configs.Groups.C | 0x0000000000000001 | valueBits

	return NewRangeF64FromU64(finalBits)
}

func NewRangeF64Identifier(sequenceType uint8, elementType uint8, elementID uint64) RangeF64 {
	if sequenceType > 3 {
		panic("sequenceType ha de ser entre 0 i 3")
	}
	if elementType > 7 {
		panic("elementType ha de ser entre 0 i 7")
	}
	if elementID >= (1 << 52) {
		panic("elementID excedeix els 52 bits disponibles")
	}

	value := cs.Range64Configs.Groups.D |
		uint64(sequenceType)<<cs.SequenceTypeShift64 |
		uint64(elementType)<<cs.ElementTypeShift64 |
		(elementID & cs.Value64Mask)
	return NewRangeF64FromU64(value)
}

// GETTERS/SETTERS --------------------
func (sF64 RangeF64) GetF64Value() float64 { return sF64.value }
func (sF64 RangeF64) SetF64Value(pVal float64) intf.RangeIntf {
	sF64.value = pVal
	return sF64
}
func (sF64 RangeF64) GetU64Value() uint64 { return tools.F64ToU64(sF64.value) }
func (sF64 RangeF64) SetU64Value(pVal uint64) intf.RangeIntf {
	sF64.value = tools.U64ToF64(pVal)
	return sF64
}

// Funcions Is....()
func (sF64 RangeF64) IsNullValue() bool {
	bits := tools.F64ToU64(sF64.value)
	return (bits & (cs.Group64Mask | cs.Subgroup64Mask)) == cs.Range64Configs.Groups.B
}

// Retorna cert només si el RangeF64 representa un error.
func (sF64 RangeF64) IsError() bool {
	bits := tools.F64ToU64(sF64.value)
	return (bits & cs.Group64Mask) == cs.Range64Configs.Groups.B
}

// Retorna el codi d'error a partir dels bits definits.
func (sF64 RangeF64) ErrorCode() int {
	bits := tools.F64ToU64(sF64.value)
	if !sF64.IsError() {
		return 0
	}
	return int((bits >> 48) & 0xFFFF)
}

// Retorna la resta de bits per a poder incorporar un paràmetre a l'error.
func (sF64 RangeF64) ErroneousValue() float64 {
	bits := tools.F64ToU64(sF64.value)
	if !sF64.IsError() {
		return 0
	}
	valueBits := bits & cs.Value64Mask
	return tools.U64ToF64(valueBits)
}

func ExtractFloat64tMantissa(value float64) uint64 {
	return math.Float64bits(value) & 0x000FFFFFFFFFFFFF // Últims 52 bits
}

func (sF64 RangeF64) ExtractMantissa() uint64 {
	return ExtractFloat64tMantissa(sF64.value)
}

func ExtractFloat64Exponent(value float64) int {
	return int((math.Float64bits(value) >> 52) & 0x7FF) // Bits 52-62
}
func (sF64 RangeF64) ExtractExponent() int {
	return ExtractFloat64Exponent(sF64.value)
}

// INTERFÍCIES ------------------------
// Cert només si la instància és menor o igual que pOther.
// TODO: Cal implementar la comparació amb 32b.
func (sF64 RangeF64) LessOrEqualThan(pOther intf.RangeIntf) bool {
	var fo64 float64

	switch other := pOther.(type) {
	case RangeF64:
		fo64 = pOther.(RangeF64).value
		break
	case rF64One.RangeF64One:
		fo64 = other.AsFloat64()
		break
	case rF64TwoPi.RangeF64TwoPi:
		fo64 = other.AsFloat64()
		break
	// case r32.Range32:
	// 	return false
	case rF32.RangeF32:
		return false
	default:
		return false
	}

	return sF64.AsFloat64()-fo64 <= cs.Epsilon64
}

// Cert només si la instància és major que pOther.
func (sF64 RangeF64) GreaterThan(pOther intf.RangeIntf) bool {
	var fo64 float64

	switch other := pOther.(type) {
	case RangeF64:
		fo64 = pOther.(RangeF64).value
		break
	case rF64One.RangeF64One:
		fo64 = other.AsFloat64()
		break
	case rF64TwoPi.RangeF64TwoPi:
		fo64 = other.AsFloat64()
		break
	// case r32.Range32:
	// 	return false
	case rF32.RangeF32:
		return false
	default:
		return false
	}

	return sF64.AsFloat64() > fo64
}

// Cert només si la instància és major o igual que pOther.
func (sF64 RangeF64) GreaterOrEqualThan(pOther intf.RangeIntf) bool {
	var fo64 float64

	switch other := pOther.(type) {
	case RangeF64:
		fo64 = pOther.(RangeF64).value
		break
	case rF64One.RangeF64One:
		fo64 = other.AsFloat64()
		break
	case rF64TwoPi.RangeF64TwoPi:
		fo64 = other.AsFloat64()
		break
	// case r32.Range32:
	// 	return false
	case rF32.RangeF32:
		return false
	default:
		return false
	}

	return sF64.AsFloat64() >= fo64
}

func (sF64 RangeF64) Add(pOther intf.RangeIntf) intf.RangeIntf { return sF64 }
func (sF64 RangeF64) Sub(pOther intf.RangeIntf) intf.RangeIntf { return sF64 }
func (sF64 RangeF64) Mul(pOther intf.RangeIntf) intf.RangeIntf { return sF64 }
func (sF64 RangeF64) Div(pOther intf.RangeIntf) intf.RangeIntf { return sF64 }
