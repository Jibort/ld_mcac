// Constants associades al tipus Range32.
// CreatedAt: 2025/01/02 dj. JIQ

package consts

import (
	"math"
)

// Màscares principals de bits
const (
	Sign32Mask        uint32  = 0x80000000         // Bit de signe
	Value32Mask       uint32  = 0x7FFFFFFF         // Bits de valor
	Meta32Mask        uint64  = 0x0000FFF000000000 // Bits de metadades
	Meta32Shift       int     = 36                 // Desplaçament de metadades
	MaxValue32One     float32 = 1.0                // Valor màxim representable
	MinValue32One     float32 = -1.0               // Valor mínim representable
	MaxValue32TwoPi   float32 = TwoPi32            // Valor màxim representable
	MinValue32TwoPi   float32 = -TwoPi32           // Valor mínim representable
	Range32NullMask   uint32  = 0x0000000F         // Exemple d'identificador per NUL
	SaturationMask    uint64  = 0x0000800000000000 // Exemple: Bit que indica saturació
	NullFlagMask      uint64  = 0x0000400000000000 // Exemple: Bit per valors nuls
	UnitFlagMask      uint64  = 0x0000200000000000 // Exemple: Bit per valors ±1.0
	SequenceTypeShift int     = 52                 // Desplaçament per al tipus de seqüència
	ElementTypeShift  int     = 48                 // Desplaçament per al tipus d'element
	GroupMask         uint64  = 0xF000000000000000 // Màscara general per identificar el grup
	SubgroupMask      uint64  = 0x0F00000000000000 // Màscara per identificar subgrups
)

// Pendent de la codificació de 32b
var (
	Epsilon32 = float32(math.Pow(10, -Precision32))
	Factor32  = float32(math.Pow(10, Precision32))
)

// Constants per Range64
const ()

type RangeL32 struct {
	Min float32
	Max float32
}

var Range32Configs = struct {
	One    RangeL32
	TwoPi  RangeL32
	Groups struct {
		A, B, C, D, E uint32
	}
}{
	One:   RangeL32{-1.0, +1.0},
	TwoPi: RangeL32{MinValue32TwoPi, MaxValue32TwoPi},
	Groups: struct {
		A, B, C, D, E uint32
	}{
		A: 0x00000000,
		B: 0x20000000,
		C: 0x40000000,
		D: 0x60000000,
		E: 0x80000000,
	},
}
