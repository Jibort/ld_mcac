// Constants associades al tipus Range64.
// CreatedAt: 2024/12/08 dg CLD

package core

// Constants i pseudoconstants per Range64
const (
	// Màscares principals de bits
	SignMask       = uint64(0b1000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000) // bit 63
	MetaMask       = uint64(0b0111_1111_1111_1111_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000) // bits 62-49
	ValueMask      = uint64(0b00_11111111111111_111111111111111111111111111111111111111111111111)
	SaturationMask = uint64(0b0111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111)

	// Màscares de grup (bits 62-61)
	GroupMask  = uint64(0b0110_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000)
	GroupAMask = uint64(0b0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000)
	GroupBMask = uint64(0b0010_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000)
	GroupCMask = uint64(0b0100_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000)
	GroupDMask = uint64(0b0110_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000)

	// Màscara de subgrups (4 bits després del grup)
	SubgroupMask = uint64(0b0011_1100_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000)

	// Subgrups del Grup B
	SubgroupNullMask = uint64(0b0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000)
	SubgroupInfMask  = uint64(0b0000_0110_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000)

	// Desplaçaments de bits
	MetaShift = 48 // bits a desplaçar per accedir a les metadades

	// Valors especials
	MaxValue = float64(1.0)
	MinValue = float64(-1.0)

	// Valors saturats (Group D)
	R64_SAT_POS_BITS = GroupDMask | SaturationMask
	R64_SAT_NEG_BITS = SignMask | GroupDMask | SaturationMask

	// Tokens (Group C)
	TokenIDMask = uint64(0b0011_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111_1111)
)

// Valors especials precompilats
var (
	R64_NULL    = NewRangeF64(U64ToF64(GroupBMask | SubgroupNullMask))
	R64_INF_POS = NewRangeF64(U64ToF64(GroupBMask | SubgroupInfMask))
	R64_INF_NEG = NewRangeF64(U64ToF64(SignMask | GroupBMask | SubgroupInfMask))
	R64_SAT_POS = NewRangeF64Saturated(true)
	R64_SAT_NEG = NewRangeF64Saturated(false)
	R64_ZERO    = NewRangeF64Zero()
)
