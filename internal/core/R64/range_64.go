// Definició i implementació de funcions generals per X64Range
// CreatedAt: 2024-12-27 dc. GPT(JIQ)

package r64

import (
	iF64 "github.com/jibort/ld_mcac/internal/core/intf/ranges"
	r "github.com/jibort/ld_mcac/internal/core/r"
)

// ESTRUCUTURES -----------------------
// X64Range representa un tipus especialitzat de 64 bits.
type X64Range struct {
	iF64.F64RangeIntf
	r.Range
}

func (m X64Range) Is64() bool {
	return true
}

func (m X64Range) Is32() bool {
	return false
}

func (m X64Range) IsError() bool {
	return false
}

// Grups
func (m X64Range) IsGroupA() bool {
	return false
}

func (m X64Range) IsGroupB() bool {
	return false
}

func (m X64Range) IsGroupC() bool {
	return false
}

func (m X64Range) IsGroupD() bool {
	return false
}

func (m X64Range) IsGroupE() bool {
	return false
}
