// range_f_64.go
// Implementació dels valors Range en float64.
// CreatedAt: 2024/12/08 dg. JIQ

package RF64

import (
	"fmt"

	cs "github.com/jibort/ld_mcac/internal/core/Consts"
	base "github.com/jibort/ld_mcac/internal/core/intf/base"
	i64 "github.com/jibort/ld_mcac/internal/core/intf/ranges"
	tools "github.com/jibort/ld_mcac/internal/core/tools"
)

// ESTRUCTURES ------------------------
// Tipus pels Range float64.
type F64Range struct {
	i64.F64RangeIntf
	value float64
}

// CONSTRUCTORS =======================
// Constructor general a partir d'un valor float64 (64bits).
func NewF64Range(pF64 float64) F64Range {
	return F64Range{value: pF64}
}

// Constructor a partir d'un uint64 (64bits).
func NewF64RangeFromU64(pU64 uint64) F64Range {
	return F64Range{value: tools.U64ToF64(pU64)}
}

// Constructor de F64Range per a 0.0.
func NewF64RangeZero() F64Range {
	return NewF64RangeFromU64(cs.IEEE754ZeroBits)
}

// INTERFÍCIES ========================

func (sRF64 F64Range) SetFloat64(pF64 float64) {
	sRF64.value = pF64
}

func (sRF64 F64Range) AsUint64() uint64 {
	return tools.F64ToU64(sRF64.AsFloat64())
}

func (sRF64 F64Range) SetUint64(pU64 uint64) {
	sRF64.SetFloat64(tools.U64ToF64(pU64))
}

func (sRF64 F64Range) AsFloat32() float32 {
	panic("F64Range.AsFloat32: Not implemented") // TODO: Implement
}

func (sRF64 F64Range) SetFloat32(_ float32) {
	panic("F64Range.SetFloat32: Not implemented") // TODO: Implement
}

func (sRF64 F64Range) AsUint32() uint32 {
	panic("F64Range.AsUint32: Not implemented") // TODO: Implement
}

func (sRF64 F64Range) SetUint32(_ uint32) {
	panic("F64Range.SetUint32: Not implemented") // TODO: Implement
}

// INTERFÍCIE 'ClonableIntf' ----------
func (sRF64 F64Range) Clone() base.ClonableIntf {
	return NewF64Range(sRF64.value)
}

// INTERFÍCIE 'ComparableIntf' --------
func (sRF64 F64Range) Equals(pOther base.ComparableIntf) bool {
	switch other := pOther.(type) {
	case F64Range:
		return tools.Equals64(sRF64.value, other.value, &cs.Epsilon64)
	default:
		panic(fmt.Sprintf("F64Range.Equals: Invalid type %T", pOther))
	}
}

func (sRF64 F64Range) LessThan(pOther base.ComparableIntf) bool {
	switch other := pOther.(type) {
	case F64Range:
		return sRF64.value < other.value
	default:
		panic(fmt.Sprintf("F64Range.LessThan: Invalid type %T", pOther))
	}
}

func (sRF64 F64Range) LessOrEqualThan(pOther base.ComparableIntf) bool {
	return sRF64.LessThan(pOther) || sRF64.Equals(pOther)
}

func (sRF64 F64Range) GreaterThan(pOther base.ComparableIntf) bool {
	switch other := pOther.(type) {
	case F64Range:
		return sRF64.value > other.value
	default:
		panic(fmt.Sprintf("F64Range.GreaterThan: Invalid type %T", pOther))
	}
}

func (sRF64 F64Range) GreaterOrEqualThan(pOther base.ComparableIntf) bool {
	return sRF64.LessThan(pOther) || sRF64.Equals(pOther)
}

// INTERFÍCIE 'RangeIntf' -------------
func (sRF64 F64Range) Is64() bool {
	panic("F64Range.Is64: Not implemented") // TODO: Implement
}

func (sRF64 F64Range) Is32() bool {
	panic("F64Range.Is32: Not implemented") // TODO: Implement
}

// Retorna cert només si el F64Range representa un error.
func (sRF64 F64Range) IsError() bool {
	return false
}

func (sRF64 F64Range) IsGroupA() bool {
	panic("F64Range.IsGroupA: Not implemented") // TODO: Implement
}

func (sRF64 F64Range) IsGroupB() bool {
	panic("F64Range.IsGroupB: Not implemented") // TODO: Implement
}

func (sRF64 F64Range) IsGroupC() bool {
	panic("F64Range.IsGroupC: Not implemented") // TODO: Implement
}

func (sRF64 F64Range) IsGroupD() bool {
	panic("F64Range.IsGroupD: Not implemented") // TODO: Implement
}

func (sRF64 F64Range) IsGroupE() bool {
	panic("F64Range.IsGroupE: Not implemented") // TODO: Implement
}

// func NewF64RangeNull() F64Range {
// 	// Basat en la documentació del Grup B, Subgrup ±Saturació i Nul
// 	// Assignem el patró fixat per a valors nul·ls
// 	bits := cs.Range64Configs.Groups.B | cs.Range64NullMask

// 	return NewF64RangeFromU64(bits)
// }

// // Crea valors saturats
// func NewF64RangeSaturated(value float64, isNull bool) F64Range {
// 	var u64 uint64

// 	if value < 0 {
// 		u64 |= cs.Saturation64Mask | cs.Sign64Mask
// 	} else {
// 		u64 |= cs.Saturation64Mask
// 	}

// 	if isNull {
// 		u64 |= cs.NullFlag64Mask
// 	} else if value == 1.0 || value == -1.0 {
// 		u64 |= cs.UnitFlag64Mask
// 	}

// 	mantissa := ExtractFloat64tMantissa(value)
// 	u64 |= mantissa

// 	return NewF64RangeFromU64(u64)
// }

// // Crea ±infinit
// func NewF64RangeInfinite(pIsPositive bool) F64Range {
// 	bits := cs.Range64Configs.Groups.B | 0x0000000000000001
// 	if !pIsPositive {
// 		bits |= cs.Sign64Mask
// 	}
// 	return NewF64RangeFromU64(bits)
// }

// // Crea un símbol
// func NewF64RangeFromSymbol(pSym rune) F64Range {
// 	sym := uint64(pSym)
// 	u64 := cs.Range64Configs.Groups.C | sym

// 	return NewF64RangeFromU64(u64)
// }

// func NewF64RangePercentage(pF64 float64) intf.Range64Intf {
// 	if pF64 < 0.0 || pF64 > 1.0 {
// 		return errs.E_InvalidPercentage()
// 	}

// 	valueBits := uint64(math.Round(pF64 * (1 << 52)))
// 	valueBits &= cs.Value64Mask

// 	finalBits := cs.Range64Configs.Groups.C | 0x0000000000000001 | valueBits

// 	return NewF64RangeFromU64(finalBits)
// }

// func NewF64RangeIdentifier(sequenceType uint8, elementType uint8, elementID uint64) F64Range {
// 	if sequenceType > 3 {
// 		panic("sequenceType ha de ser entre 0 i 3")
// 	}
// 	if elementType > 7 {
// 		panic("elementType ha de ser entre 0 i 7")
// 	}
// 	if elementID >= (1 << 52) {
// 		panic("elementID excedeix els 52 bits disponibles")
// 	}

// 	value := cs.Range64Configs.Groups.D |
// 		uint64(sequenceType)<<cs.SequenceTypeShift64 |
// 		uint64(elementType)<<cs.ElementTypeShift64 |
// 		(elementID & cs.Value64Mask)
// 	return NewF64RangeFromU64(value)
// }

// // GETTERS/SETTERS --------------------
// func (sF64 F64Range) GetF64Value() float64 { return sF64.value }
// func (sF64 F64Range) SetF64Value(pVal float64) intf.RangeIntf {
// 	sF64.value = pVal
// 	return sF64
// }
// func (sF64 F64Range) GetU64Value() uint64 { return tools.F64ToU64(sF64.value) }
// func (sF64 F64Range) SetU64Value(pVal uint64) intf.RangeIntf {
// 	sF64.value = tools.U64ToF64(pVal)
// 	return sF64
// }

// // Funcions Is....()
// func (sF64 F64Range) IsNullValue() bool {
// 	bits := tools.F64ToU64(sF64.value)
// 	return (bits & (cs.Group64Mask | cs.Subgroup64Mask)) == cs.Range64Configs.Groups.B
// }

// // Retorna el codi d'error a partir dels bits definits.
// func (sF64 F64Range) ErrorCode() int {
// 	bits := tools.F64ToU64(sF64.value)
// 	if !sF64.IsError() {
// 		return 0
// 	}
// 	return int((bits >> 48) & 0xFFFF)
// }

// // Retorna la resta de bits per a poder incorporar un paràmetre a l'error.
// func (sF64 F64Range) ErroneousValue() float64 {
// 	bits := tools.F64ToU64(sF64.value)
// 	if !sF64.IsError() {
// 		return 0
// 	}
// 	valueBits := bits & cs.Value64Mask
// 	return tools.U64ToF64(valueBits)
// }

// func ExtractFloat64tMantissa(value float64) uint64 {
// 	return math.Float64bits(value) & 0x000FFFFFFFFFFFFF // Últims 52 bits
// }

// func (sF64 F64Range) ExtractMantissa() uint64 {
// 	return ExtractFloat64tMantissa(sF64.value)
// }

// func ExtractFloat64Exponent(value float64) int {
// 	return int((math.Float64bits(value) >> 52) & 0x7FF) // Bits 52-62
// }
// func (sF64 F64Range) ExtractExponent() int {
// 	return ExtractFloat64Exponent(sF64.value)
// }

// func (sF64 F64Range) Add(pOther base.RangeIntf) base.RangeIntf { return sF64 }
// func (sF64 F64Range) Sub(pOther base.RangeIntf) base.RangeIntf { return sF64 }
// func (sF64 F64Range) Mul(pOther base.RangeIntf) base.RangeIntf { return sF64 }
// func (sF64 F64Range) Div(pOther base.RangeIntf) base.RangeIntf { return sF64 }
