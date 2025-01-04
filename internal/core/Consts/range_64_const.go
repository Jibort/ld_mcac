// Constants associades al tipus Range64.
// CreatedAt: 2024/12/28 ds. GPT

package consts

import (
	"math"
	"unsafe"
)

// Constant d'ús general
const (
	IEEE754ZeroBits uint64 = 0x0000000000000000 // 0.0 segons l'estàndard IEEE 754
)

// Màscares principals de bits
const (
	Sign64Mask          uint64  = 0x8000000000000000 // Bit de signe
	Value64Mask         uint64  = 0x7FFFFFFFFFFFFFFF // Bits de valor
	Meta64Mask          uint64  = 0x0000FFF000000000 // Bits de metadades
	Meta64Shift         int     = 36                 // Desplaçament de metadades
	MaxValue64TwoPi     float64 = TwoPi64            // Valor màxim representable
	MinValue64TwoPi     float64 = -TwoPi64           // Valor mínim representable
	Range64NullMask     uint64  = 0x000000000000000F // Exemple d'identificador per NUL
	Saturation64Mask    uint64  = 0x0000800000000000 // Exemple: Bit que indica saturació
	NullFlag64Mask      uint64  = 0x0000400000000000 // Exemple: Bit per valors nuls
	UnitFlag64Mask      uint64  = 0x0000200000000000 // Exemple: Bit per valors ±1.0
	SequenceTypeShift64 int     = 52                 // Desplaçament per al tipus de seqüència
	ElementTypeShift64  int     = 48                 // Desplaçament per al tipus d'element
	Group64Mask         uint64  = 0xF000000000000000 // Màscara general per identificar el grup
	Subgroup64Mask      uint64  = 0x0F00000000000000 // Màscara per identificar subgrups
)

var (
	Epsilon64 float64 = float64(math.Pow(10, -Precision64)) // Precisió per a comparacions
	Factor64  float64 = math.Pow(10, Precision64)           // Augmenta la precisió en l'arrodoniment
)

type RangeLF64 struct {
	Min float64
	Max float64
}
type RangeLU64 struct {
	Min uint64
	Max uint64
}

var Range64Configs = struct {
	OneF64   RangeLF64
	TwoPiF64 RangeLF64
	OneU64   RangeLU64
	TwoPiU64 RangeLU64

	Groups struct {
		A, B, C, D, E uint64
	}
}{
	OneF64:   RangeLF64{-1.0, +1.0},
	TwoPiF64: RangeLF64{MinValue64TwoPi, MaxValue64TwoPi},
	OneU64:   RangeLU64{f64ToU64(-1.0), f64ToU64(+1.0)},
	TwoPiU64: RangeLU64{f64ToU64(MinValue64TwoPi), f64ToU64(MaxValue64TwoPi)},
	Groups: struct {
		A, B, C, D, E uint64
	}{
		A: 0x0000000000000000,
		B: 0x2000000000000000,
		C: 0x4000000000000000,
		D: 0x6000000000000000,
		E: 0x8000000000000000,
	},
}

func f64ToU64(pF64 float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&pF64))
}
