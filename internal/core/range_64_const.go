// Constants associades al tipus Range64.
// CreatedAt: 2024/12/10 dt GPT

package core

// Constants i pseudoconstants per Range64
// Màscares principals de bits
const (
	SignMask  = uint64(0b10000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000) // bit 63
	MetaMask  = uint64(0b01111111_11111111_00000000_00000000_00000000_00000000_00000000_00000000) // bits 62-49
	ValueMask = uint64(0b00000000_11111111_11111111_11111111_11111111_11111111_11111111_11111111) // bits de valor
	// SaturationMask = uint64(0b01111111_11111111_11111111_11111111_11111111_11111111_11111111_11111111) // saturació (Group D)
)

// Màscares de grup (bits 62-61)
const (
	// Màscara principal per seleccionar només els bits del grup (bits 62 i 61)
	GroupMask = uint64(0b01100000_00000000_00000000_00000000_00000000_00000000_00000000_00000000)

	// Definició de grups
	GroupAMask = uint64(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000) // Grup A: 00
	GroupBMask = uint64(0b00100000_00000000_00000000_00000000_00000000_00000000_00000000_00000000) // Grup B: 01
	GroupCMask = uint64(0b01000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000) // Grup C: 10
	GroupDMask = uint64(0b01100000_00000000_00000000_00000000_00000000_00000000_00000000_00000000) // Grup D: 11
)

// Màscares de subgrup (4 bits després del grup)
const (
	SubgroupMask     = uint64(0b00011110_00000000_00000000_00000000_00000000_00000000_00000000_00000000) // màscara genèrica de subgrup
	SubGroupC3Mask   = uint64(0b00001100_00000000_00000000_00000000_00000000_00000000_00000000_00000000) // Subgrup C3
	SubgroupNullMask = uint64(0b00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000000) // Subgrup B (null)
	SubgroupInfMask  = uint64(0b00000110_00000000_00000000_00000000_00000000_00000000_00000000_00000000) // Subgrup B (infinit)
)

const (
	SequenceTypeMask  = uint64(0b00001100_00000000_00000000_00000000_00000000_00000000_00000000_00000000) // bits 60-59
	ElementTypeMask   = uint64(0b00000011_10000000_00000000_00000000_00000000_00000000_00000000_00000000) // bits 58-56
	SequenceTypeShift = 60
	ElementTypeShift  = 57
	ElementIDMask     = uint64(0x000FFFFFFFFFFFFF) // bits inferiors (52 bits)

)

// Màscares de padding.
const (
	PaddingStartMask = uint64(0b01000100_00000000_00000000_00000000_00000000_00000000_11111111_00001010) // rune(0xFFFA): 0xFF0A, // Paddind d'inici
	PaddingEndMask   = uint64(0b11000100_00000000_00000000_00000000_00000000_00000000_11111111_00001011) // rune(0xFFFB): 0xFF0B, // Paddind de final
	PaddingMask      = uint64(0b01000100_00000000_00000000_00000000_00000000_00000000_00000000_00000000) // rune(0xFFFB): 0xFF0B, // Paddind de final
)

// Desplaçaments de bits
const (
	MetaShift = 48 // Bits a desplaçar per accedir a les metadades
)

// Valors especials
const (
	MaxValue               = float64(1.0)
	MinValue               = float64(-1.0)
	IEEE754ZeroBits uint64 = 0x0000000000000000 // 0.0 segons l'estàndard IEEE 754
)

// Tokens i valors especials
const (
	GroupCTokenIDMask = uint64(0b00111111_11111111_11111111_11111111_11111111_11111111_11111111_11111111) // Identificadors de tokens (Grup C)
)

// Valors precompilats
var (
	R64_NULL     = NewRangeF64(U64ToF64(GroupBMask | SubgroupNullMask))
	R64_INF_POS  = NewRangeF64(U64ToF64(GroupBMask | SubgroupInfMask))
	R64_INF_NEG  = NewRangeF64(U64ToF64(SignMask | GroupBMask | SubgroupInfMask))
	R64_SAT_1POS = NewRangeF64Saturated(1.0, false)
	R64_SAT_1NEG = NewRangeF64Saturated(-1.0, false)
	R64_ZERO     = NewRangeF64Zero()
)

var (
	RangeF64PositiveSaturated = NewRangeF64Saturated(1, false)  // +SAT
	RangeF64NegativeSaturated = NewRangeF64Saturated(-1, false) // -SAT
)

const (
	SaturationMask    uint64 = 0x8000000000000000 // Bit de saturació (S1)
	SignBitMask       uint64 = 0x4000000000000000 // Bit de signe
	NullFlagMask      uint64 = 0x2000000000000000 // Flag 'x' per valors nuls
	UnitFlagMask      uint64 = 0x1000000000000000 // Flag 'u' per +1.0* / -1.0*
	SubnormalFlagMask uint64 = 0x0800000000000000 // Flag 'n' per subnormalitzats
)

// Constants necessàries per a l'ampliació del Grup A
const (
	// Exponents per a [-2π, +2π]
	Exponent1024 = 1024
	Exponent1025 = 1025

	// Nous rangs del Grup A
	RangeNegTwoPi = -2 * math.Pi // -2π
	RangePosTwoPi = 2 * math.Pi  // 2π
