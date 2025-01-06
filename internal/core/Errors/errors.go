// Identificació i codificació dels errors que es poden produïr.
// CreatedAt: 2024/12/11 dc. JIQ[GPT]

package errors

import (
	cs "github.com/jibort/ld_mcac/internal/core/consts"
)

func E_Unknown(pArgs uint64) Error {
	return NewError(true, cs.ErrCode_UnknownError, pArgs)
}

func E_InvalidArguments(pArgs uint64) Error {
	return NewError(true, cs.ErrCode_InvalidArguments, pArgs)
}

func E_InvalidPercentage(pArgs uint64) Error {
	return NewError(true, cs.ErrCode_InvalidPercentage, pArgs)
}

func E_OutOfRangeSymbol(pArgs uint64) Error {
	return NewError(true, cs.ErrCode_OutOfRangeSymbol, pArgs)
}
