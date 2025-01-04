//
// CreatedAt: 2025/01/02 dj. JIQ

package RF64Sym

import (
	cs "github.com/jibort/ld_mcac/internal/core/Consts"
	errs "github.com/jibort/ld_mcac/internal/core/Errors"
	rF64 "github.com/jibort/ld_mcac/internal/core/RF64"
	syms "github.com/jibort/ld_mcac/internal/core/intf/symbols"
	tools "github.com/jibort/ld_mcac/internal/core/tools"
)

type F64RangeSymbol struct {
	syms.F64SymbolIntf

	inst rF64.F64Range
}

// Constructor per crear un F64RangeSymbol des d'un rune (UTF-32)
func NewF64RangeSymbol(symbol rune) intf.Range64Intf {
	// Verifiquem si el símbol està fora del rang UTF-32
	if symbol > 0x10FFFF {
		// Retornem un F64RangeError amb un codi d'error específic
		err := errs.NewRange64Error(false, cs.ErrCode_OutOfRangeSymbol, uint64(symbol))
		return err
	}

	// Codifiquem el símbol directament dins de F64Range
	return F64RangeSymbol{rF64.NewF64Range(EncodeSymbolValue(symbol))}
}

// Codifica un symbol com a 'float64'.
func EncodeSymbolValue(symbol rune) (rF64 float64) {
	bits := cs.Mask_Subgrup_B1 | uint64(symbol)
	rF64 = tools.U64ToF64(bits)

	return
}

// Decodificar un F64RangeSymbol a un rune
func (sR64Sym F64RangeSymbol) Decode() rune {
	bits := tools.F64ToU64(sR64Sym.AsFloat64())
	if (bits & cs.Mask_Subgrup_B1) != cs.Mask_Subgrup_B1 {
		panic("Error de codificació de F64RangeSymbol")
	}
	bits &= ^cs.Mask_Subgrup_B1
	return rune(bits)
}
