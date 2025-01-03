// Identificació i codificació dels errors que es poden produïr.
// CreatedAt: 2024/12/11 dc. JIQ[GPT]

package Errors

import (
	cs "github.com/jibort/ld_mcac/internal/core/Consts"
)

func E_Unknown(pArgs uint64) RangeF64Error {
	return NewRange64Error(true, cs.ErrCode_UnknownError, pArgs)
}

func E_InvalidArguments(pArgs uint64) RangeF64Error {
	return NewRange64Error(true, cs.ErrCode_InvalidArguments, pArgs)
}

func E_InvalidPercentage(pArgs uint64) RangeF64Error {
	return NewRange64Error(true, cs.ErrCode_InvalidPercentage, pArgs)
}

func E_OutOfRangeSymbol(pArgs uint64) RangeF64Error {
	return NewRange64Error(true, cs.ErrCode_OutOfRangeSymbol, pArgs)
}
