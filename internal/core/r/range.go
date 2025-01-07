// Definició i implementació de funcions generals per X64Range
// CreatedAt: 2024-12-27 dc. GPT(JIQ)

package r

import (
	base "github.com/jibort/ld_mcac/internal/core/intf/base"
)

// ESTRUCUTURES -----------------------
// X64Range representa un tipus especialitzat de 64 bits.
type Range struct {
	base.RangeIntf
}

// Assegura que Range implementa RangeIntf
var _ base.RangeIntf = (*Range)(nil)

func (m Range) AsFloat64() float64 {
	panic("Range.AsFloat64: not implemented") // TODO: Implement
}

func (m *Range) SetFloat64(_ float64) {
	panic("Range.SetFloat64: not implemented") // TODO: Implement
}

func (m Range) AsUint64() uint64 {
	panic("Range.AsUint64: not implemented") // TODO: Implement
}

func (m *Range) SetUint64(_ uint64) {
	panic("Range.SetUint64: not implemented") // TODO: Implement
}

func (m Range) AsFloat32() float32 {
	panic("Range.AsFloat32: not implemented") // TODO: Implement
}

func (m *Range) SetFloat32(_ float32) {
	panic("Range.SetFloat32: not implemented") // TODO: Implement
}

func (m Range) AsUint32() uint32 {
	panic("Range.AsUint32: not implemented") // TODO: Implement
}

func (m *Range) SetUint32(_ uint32) {
	panic("Range.SetUint32: not implemented") // TODO: Implement
}

func (m Range) Clone() base.RangeIntf {
	panic("Range.Clone: not implemented") // TODO: Implement
}

func (m Range) Equals(pOther base.RangeIntf) bool {
	panic("Range.Equals: not implemented") // TODO: Implement
}

func (m Range) LessThan(pOther base.RangeIntf) bool {
	panic("Range.LessThan: not implemented") // TODO: Implement
}

func (m Range) LessOrEqualThan(pOther base.RangeIntf) bool {
	panic("Range.LessOrEqualThan: not implemented") // TODO: Implement
}

func (m Range) GreaterThan(pOther base.RangeIntf) bool {
	panic("Range.GreaterThan: not implemented") // TODO: Implement
}

func (m Range) GreaterOrEqualThan(pOther base.RangeIntf) bool {
	panic("Range.GreaterOrEqualThan: not implemented") // TODO: Implement
}

// Blocs
func (m Range) Is64() bool {
	panic("Range.Is64: not implemented") // TODO: Implement
}

func (m Range) Is32() bool {
	panic("Range.Is32: not implemented") // TODO: Implement
}

func (m Range) IsError() bool {
	panic("Range.IsError: not implemented") // TODO: Implement
}

// Grups
func (m Range) IsGroupA() bool {
	panic("Range.IsGroupA: not implemented") // TODO: Implement
}

func (m Range) IsGroupB() bool {
	panic("Range.IsGroupB: not implemented") // TODO: Implement
}

func (m Range) IsGroupC() bool {
	panic("Range.IsGroupC: not implemented") // TODO: Implement
}

func (m Range) IsGroupD() bool {
	panic("Range.IsGroupD: not implemented") // TODO: Implement
}

func (m Range) IsGroupE() bool {
	panic("Range.IsGroupE: not implemented") // TODO: Implement
}
