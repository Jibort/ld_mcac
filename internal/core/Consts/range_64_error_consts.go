// Llistat dels errors codificats dins RangeF64Error.
// CreatedAt: 2024/12/31 dg. GPT(JIQ)

package consts

// MÀSCARES ---------------------------
const Mask_Subgrup_B1 uint64 = 0b01000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000
const Mask_Subgrup_B4 uint64 = 0b01001100_00000000_00000000_00000000_00000000_00000000_00000000_00000000
const Mask_B64_Critical uint64 = 0b10000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000
const Mask_B4_ErrorCode uint64 = 0b00000011_11111111_00000000_00000000_00000000_00000000_00000000_00000000
const Mask_B4_Arguments uint64 = 0b00000000_00000000_11111111_11111111_11111111_11111111_11111111_11111111

// CODIS D'ERROR ----------------------
// '0000000000' (0): Codi d'error desconegut!
// 		a: '-00000000-00000000-00000000-00000000-00000000-000000cc-cccccccc', on:
//			c: és el codi d'error desconegut.
const ErrCode_UnknownError uint16 = 0b00000000_00000000

// '0000000001' (1): Arguments d'error incorrectes!
// 		a: '-aaaaaaaa-aaaaaaaa-aaaaaaaa-aaaaaaaa-aaaaaaaa-aaaaaaaa-aaaaaaaa', on:
//			a: és el conjunt de bits que no es reconeixen com a arguments.
const ErrCode_InvalidArguments uint16 = 0b00000000_00000001

// '0000000010' (2): Valor de percentatge invàlid
// 		a: L'error no fa servir paràmetres
const ErrCode_InvalidPercentage uint16 = 0b00000000_00000010

// '0000000011' (3): Símbol fora de rang
// 		a: L'error no fa servir paràmetres
const ErrCode_OutOfRangeSymbol uint16 = 0b00000000_00000011
