// Inicialització del paquet 'core'.
// CreatedAt: 2024/12/10 dt. GPT

package core

// Inicialització del paquet 'core'.
func init() {
	IdToSymbol = make(map[uint64]rune)
	for symbol, id := range SymbolToID {
		IdToSymbol[id] = symbol
	}
}

var (
	R64Space    = NewRangeU64(0x4000000000000001) // ' ' (Espai)
	R64Tab      = NewRangeU64(0x4000000000000002) // '\t' (Tabulador)
	R64Newline  = NewRangeU64(0x4000000000000003) // '\n' (Salt de línia)
	R64Carriage = NewRangeU64(0x4000000000000004) // '\r' (Retorn de carro)

	R64Error   = NewRangeU64(0x4000000000000FFF) // Representa 'error'
	R64Unknown = NewRangeU64(0x4000000000000FFE) // Representa 'desconegut'
	R64Any     = NewRangeU64(0x4000000000000FFD) // Representa 'qualsevol'

	R64A           = NewRangeU64(0x4000000000000010) // 'A'
	R64a           = NewRangeU64(0x4000000000000060) // 'a'
	R64Zero        = NewRangeU64(0x40000000000000B0) // '0'
	R64OpenParen   = NewRangeU64(0x4000000000000200) // '('
	R64CloseParen  = NewRangeU64(0x4000000000000300) // ')'
	R64Period      = NewRangeU64(0x4000000000000800) // '.'
	R64Exclamation = NewRangeU64(0x4000000000000900) // '!'
	R64Dash        = NewRangeU64(0x4000000000000A00) // '-'
)

// Codificació de cada símbol com a RangeF64
func EncodeSymbol(symbol rune) RangeF64 {
	if id, exists := SymbolToID[symbol]; exists {
		// Combinem els camps de grup, subgrup i identificador
		value := GroupCMask | SubGroupC3Mask | id

		return NewRangeU64(value)
	}

	return R64Unknown
}

// Decodificació de cada RangeF64 símbol en un rune
func DecodeSymbol(r RangeF64) rune {
	bits := F64ToU64(r.GetValue())
	id := bits & 0xFFF // Extreu només els últims 12 bits (identificador)

	if symbol, exists := IdToSymbol[id]; exists {
		return symbol
	}

	return rune(0xFFFE) // Retorna 'unknown' si no es troba
}
