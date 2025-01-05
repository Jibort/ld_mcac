// Implementació dels valors Range en float32.
// CreatedAt: 2024/12/08 dg. JIQ

package RF32

import (
	"math"

	cs "github.com/jibort/ld_mcac/internal/core/Consts"
	base "github.com/jibort/ld_mcac/internal/core/intf/base"
	"github.com/jibort/ld_mcac/internal/core/intf/ranges"
	intf "github.com/jibort/ld_mcac/internal/core/intf/ranges"
	tools "github.com/jibort/ld_mcac/internal/core/tools"
)

// Tipus pels Range float64.
type F32Range struct {
	intf.X32RangeIntf
	value float32
}

// CONSTRUCTORS -----------------------
func F32NewRange(pVal float32) F32Range {
	val := tools.Quantize32One(pVal)
	if val < -1.0 {
		val = -1.0

	} else if val > +1.0 {
		val = 1.0
	}

	return F32Range{value: tools.Quantize32One(val)}
}

// INTERFÍCIES ========================
// INTERFÍCIE 'ComparableIntf' --------
// Retorna cert només si la diferència entre els valors és inferior a epsilon32.
func (sRF32 F32Range) Equals(pOther base.RangeIntf) bool {
	var src, tgt float32

	switch other := pOther.(type) {
	case *F32Range:
		src = tools.Quantize32One(sRF32.value)
		tgt = tools.Quantize32One(other.AsFloat32())
	default:
		return false
	}

	return (src - tgt) < cs.Epsilon32
}

// Retorna cert només en cas que el valor de la instància sigui inferior al de pOther.
func (sRF32 F32Range) LessThan(pOther base.RangeIntf) bool {
	src := tools.Quantize32One(float32(math.Abs(float64(sRF32.value))))
	tgt := tools.Quantize32One(pOther.AsFloat32())
	return src < tgt
}

// Retorna cert només si el valor de la instància és inferior o igual al de pOtheer.
func (sRF32 F32Range) LessOrEqualThan(pOther base.RangeIntf) bool {
	return sRF32.LessThan(pOther) || sRF32.Equals(pOther)
}

func (sRF32 F32Range) GreaterThan(pOther base.RangeIntf) bool {
	src := tools.Quantize32One(float32(math.Abs(float64(sRF32.value))))
	tgt := tools.Quantize32One(pOther.AsFloat32())
	return src > tgt
}

func (sRF32 F32Range) GreaterOrEqualThan(pOther base.RangeIntf) bool {
	return sRF32.GreaterThan(pOther) || sRF32.Equals(pOther)
}

// INTERFÍCIE 'CloneableIntf' ---------
func (sRF32 F32Range) Clone() base.RangeIntf {
	res := F32NewRange(sRF32.value)
	return &res
}

// INTERFÍCIE 'RangeIntf' -------------
func (sRF32 F32Range) Is64() bool {
	return false
}

func (sRF32 F32Range) Is32() bool {
	return true
}

func (sRF32 F32Range) IsError() bool {
	return false
}

// Grups
func (sRF32 F32Range) IsGroupA() bool {
	panic("F32Range.IsGroupA : not implemented") // TODO: Implement
}

func (sRF32 F32Range) IsGroupB() bool {
	panic("F32Range.IsGroupB : not implemented") // TODO: Implement
}

func (sRF32 F32Range) IsGroupC() bool {
	panic("F32Range.IsGroupC : not implemented") // TODO: Implement
}

func (sRF32 F32Range) IsGroupD() bool {
	panic("F32Range.IsGroupD : not implemented") // TODO: Implement
}

func (sRF32 F32Range) IsGroupE() bool {
	return false
}

// INTERFÍCIE 'MathOperationsIntf'  ---
func (sRF32 F32Range) Add(pOther base.RangeIntf) base.RangeIntf {
	panic("F32Range.Add : not implemented") // TODO: Implement
}

func (sRF32 F32Range) Sub(pOther base.RangeIntf) base.RangeIntf {
	panic("F32Range.Sub : not implemented") // TODO: Implement
}

func (sRF32 F32Range) Mul(pOther base.RangeIntf) base.RangeIntf {
	panic("F32Range.Mul : not implemented") // TODO: Implement
}

func (sRF32 F32Range) Div(pOther base.RangeIntf) base.RangeIntf {
	panic("F32Range.Div : not implemented") // TODO: Implement
}

func (sRF32 F32Range) IsF32() bool {
	return true
}

func (sRF32 F32Range) IsU32() bool {
	return false
}

func (sRF32 F32Range) IsF64() bool {
	return false
}

func (sRF32 F32Range) IsU64() bool {
	return false
}

func (sRF32 F32Range) As64() ranges.X64RangeIntf {
	panic("F32Range.As64 : not implemented") // TODO: Implement
}

func (sRF32 F32Range) AsF64() ranges.F64RangeIntf {
	panic("F32Range.AsF64() : not implemented") // TODO: Implement
}

// INTERFÍCIE 'TypeConversionsIntf' ---
func (sRF32 F32Range) AsFloat64() float64 {
	return float64(tools.F32ToF64(sRF32.value))
}

func (sRF32 *F32Range) SetFloat64(pF64 float64) {
	r32 := tools.F64ToF32(pF64)
	sRF32.value = r32
}

func (sRF32 F32Range) AsUint64() uint64 {
	panic("F32Range.AsUint64() : not implemented") // TODO: Implement
}

func (sRF32 *F32Range) SetUint64(uint64) {
	panic("F32Range.SetUint64() : not implemented") // TODO: Implement
}

func (sRF32 F32Range) AsFloat32() float32 {
	return sRF32.value
}

func (sRF32 *F32Range) SetFloat32(pF32 float32) {
	sRF32.value = tools.Quantize32One(pF32)
}

func (sRF32 F32Range) AsUint32() uint32 {
	return tools.F32ToU32(sRF32.value)
}

func (sRF32 *F32Range) SetUint32(pF32 uint32) {
	sRF32.value = tools.U32ToF32(pF32)
}
