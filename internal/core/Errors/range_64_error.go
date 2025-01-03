// Representa codis d'error i arguments associats sobre float64 (subgrup B.4)
// CreatedAt: 2024/12/31 dg. GPT(JIQ)

// Format estés sobre float64 (cas subgrup B.4):
// B7       B6       B5       B4       B3       B2       B1       B0
// c10011ee-eeeeeeee-aaaaaaaa-aaaaaaaa-aaaaaaaa-aaaaaaaa-aaaaaaaa-aaaaaaaa, on:
//		?   : Bits reservat per a extensions com:
//		c   : Flag de criticitat (0: error gestionable, 1: error que provoca pànic).
//		100 : Identificador del Grup B.
// 		11  : Identificador del Subgrup B.4.
// 		e   : 10b per al codi d'error (disponibles 1024 errors diferents).
//	    a   : 48b per a possibles arguments associats (el format d'aquests arguments variarà segons de quin error es tracti).

package Errors

import (
	"fmt"
	"strings"

	cs "github.com/jibort/ld_mcac/internal/core/Consts"
	intf "github.com/jibort/ld_mcac/internal/core/Intf"
	tools "github.com/jibort/ld_mcac/internal/core/tools"
)

type RangeF64Error struct {
	intf.ErrorF64Intf
	Critical bool
	Code     uint16
	Args     uint64
}

// Crea un valor RangeF64 per a errors, codificant el codi d'error i el valor erroni.
func NewRange64Error(pCritic bool, pCode uint16, pArgs uint64) RangeF64Error {
	// Validar el codi d'error i els arguments
	if pCritic || !isValidErrorCode(pCode) {
		panic(fmt.Sprintf("Codi d'error crític o desconegut: %d", pCode))
	}
	if pCritic || !isValidErrorArgs(pCode, pArgs) {
		panic(fmt.Sprintf("Arguments d'error invàlids: %d", pArgs))
	}

	// Codificar el valor segons l'especificació
	return RangeF64Error{Critical: pCritic, Code: pCode, Args: pArgs}
}

// Validar si un codi d'error és vàlid
func isValidErrorCode(pCode uint16) bool {
	// Suposem que els codis d'error vàlids són de 0 a 1023 (10 bits disponibles)
	return pCode < 1024
}

// Validar si els arguments són vàlids
func isValidErrorArgs(pCode uint16, pArgs uint64) bool {
	// Aquesta funció s'ha d'anar recodificant per a cada error que identifiquem.
	switch pCode {
	case cs.ErrCode_UnknownError:
		return true
	case cs.ErrCode_InvalidArguments:
		return true
	default:
		break
	}

	// Aquest return és un simple retorn dummy.
	return pArgs >= 0 && pArgs <= 65535
}

// Codificar el valor d'un error i els corresponents arguments dins d'un float64
func EncodeErrorValue(pCritic bool, pCode uint16, pArgs uint64) (rF64 float64) {
	var bits uint64 = cs.Mask_Subgrup_B4

	// Codificar el codi d'error i els arguments segons els bits assignats al PDF
	if pCritic {
		bits |= cs.Mask_B64_Flag_Critical_Error
	}

	bits |= uint64(pCode) << 48                  // Els primers 10 bits (Codi d'error).
	bits |= uint64(pArgs) & cs.Mask_B4_Arguments // Els 48b d'arguments.
	rF64 = tools.U64ToF64(bits)

	return
}

func FormatUint64AsBytes(value uint64) string {
	binary := fmt.Sprintf("%064b", value) // Generar la representació binària completa
	var grouped []string

	// Separar en grups de 8 bits (bytes)
	for i := 0; i < len(binary); i += 8 {
		grouped = append(grouped, binary[i:i+8])
	}

	// Unir els grups amb un espai
	return strings.Join(grouped, " ")
}

// Descodificar un RangeF64Error
func (sErr RangeF64Error) Decode() (rCritic bool, rCode int, rArgs []any) {
	// Decodifica el valor float64 per obtenir el codi d'error i els arguments
	bits := tools.F64ToU64(sErr.AsFloat64())
	rCritic = (bits & cs.Mask_B64_Flag_Critical_Error) != 0
	rCode = int(bits&cs.Mask_B4_ErrorCode) >> 48

	// Cal decodificar els arguments per tipus d'error
	args := (bits & cs.Mask_B4_Arguments) // Últims 48 bits pels arguments

	numArgs := 1 // S'haurà de calcular dins el switch
	switch rCode {
	// Avaluar 'args' i 'numArgs' per tipus d'error
	default:
		break
	}
	rArgs = make([]any, numArgs) // Creem una llista d'arguments dummy.
	rArgs[0] = args              // Només per a evitar l'error de compilació

	return
}

// Retorna cert sempre perquè totes les instàncies de RangeF64Error són errors.
func (sErr RangeF64Error) IsError() bool {
	return true
}
