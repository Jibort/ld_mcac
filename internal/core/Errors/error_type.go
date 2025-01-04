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

package errors

import (
	"fmt"
	"strings"

	cs "github.com/jibort/ld_mcac/internal/core/consts"
	ierr "github.com/jibort/ld_mcac/internal/core/intf/errors"
	tools "github.com/jibort/ld_mcac/internal/core/tools"
)

type Error struct {
	ierr.ErrorIntf
	value float64
}

// Crea un error, codificant el codi i el valor erroni.
func NewError(pCritic bool, pCode uint16, pArgs uint64) Error {
	// Validar el codi d'error i els arguments
	if pCritic || !isValidErrorCode(pCode) {
		panic(fmt.Sprintf("Codi d'error crític o desconegut: %d", pCode))
	}
	if pCritic || !isValidErrorArgs(pCode, pArgs) {
		panic(fmt.Sprintf("Arguments d'error invàlids: %d", pArgs))
	}

	// Codificar el valor segons l'especificació
	coded := Encode(pCritic, pCode, pArgs)
	return Error{value: coded}
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
func Encode(pCritic bool, pCode uint16, pArgs uint64) (rF64 float64) {
	var bits uint64 = cs.Mask_Subgrup_B4

	// Codificar el codi d'error i els arguments segons els bits assignats al PDF
	if pCritic {
		bits |= cs.Mask_B64_Critical
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

// INTERFÍCIE 'ErrorIntf' -------------
// Cert només si l'error és crític
func (sErr Error) IsCritical() bool {
	u64 := tools.F64ToU64(sErr.value)
	return (u64 & cs.Mask_B64_Critical) != 0
}

// Re<torna el codi d'error
func (sErr Error) Code() uint16 {
	u64 := tools.F64ToU64(sErr.value)
	return uint16(u64 & cs.Mask_B4_ErrorCode >> 48)
}

// Retorna els arguments de l'error
func (sErr Error) Arguments() []uint64 {
	u64 := tools.F64ToU64(sErr.value)
	return []uint64{(u64 & cs.Mask_B4_Arguments >> 48)}
}

// Descodificar un Error
func (sErr Error) Decode() (rCritic bool, rCode int, rArgs []any) {
	// Decodifica el valor float64 per obtenir el codi d'error i els arguments
	bits := sErr.AsUint64()
	rCritic = (bits & cs.Mask_B64_Critical) != 0
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

// Retorna cert sempre perquè totes les instàncies d'Error són errors.
func (sErr Error) IsError() bool {
	return true
}
