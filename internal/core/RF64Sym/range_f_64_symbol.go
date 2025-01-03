//
// CreatedAt: 2025/01/02 dj. JIQ

package RF64Sym

import (
	cs "github.com/jibort/ld_mcac/internal/core/Consts"
	errs "github.com/jibort/ld_mcac/internal/core/Errors"
	rF64 "github.com/jibort/ld_mcac/internal/core/RF64"
	intf "github.com/jibort/ld_mcac/internal/core/intf"
	tools "github.com/jibort/ld_mcac/internal/core/tools"
)

type RangeF64Symbol struct {
	intf.RangeF64SymbolIntf

	inst rF64.RangeF64
}

// Constructor per crear un RangeF64Symbol des d'un rune (UTF-32)
func NewRangeF64Symbol(symbol rune) intf.Range64Intf {
	// Verifiquem si el símbol està fora del rang UTF-32
	if symbol > 0x10FFFF {
		// Retornem un RangeF64Error amb un codi d'error específic
		err := errs.NewRange64Error(false, cs.ErrCode_OutOfRangeSymbol, uint64(symbol))
		return err
	}

	// Codifiquem el símbol directament dins de RangeF64
	return RangeF64Symbol{rF64.NewRangeF64(EncodeSymbolValue(symbol))}
}

// Codifica un symbol com a 'float64'.
func EncodeSymbolValue(symbol rune) (rF64 float64) {
	bits := cs.Mask_Subgrup_B1 | uint64(symbol)
	rF64 = tools.U64ToF64(bits)

	return
}

// Decodificar un RangeF64Symbol a un rune
func (sR64Sym RangeF64Symbol) Decode() rune {
	bits := tools.F64ToU64(sR64Sym.AsFloat64())
	if (bits & cs.Mask_Subgrup_B1) != cs.Mask_Subgrup_B1 {
		panic("Error de codificació de RangeF64Symbol")
	}
	bits &= ^cs.Mask_Subgrup_B1
	return rune(bits)
}
