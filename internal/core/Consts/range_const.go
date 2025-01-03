package Consts

import "math"

// Constant d'ús general

var (
	c_true  = true
	c_false = false
)

func TruePoint() *bool    { return &c_true }
func FalsePointer() *bool { return &c_false }

const (
	TwoPi64 = 2 * math.Pi      // 2π en float64
	TwoPi32 = float32(TwoPi64) // 2π en float32

	Precision32 = 6  // Pendent de la codificació dels 32b
	Precision64 = 13 // Pendent de les proves anb 64b

	SmallThreshold32 float32 = 1e-8 // Augmenta la sensibilitat per valors petits
	SmallThreshold64 float64 = 1e-8 // Augmenta la sensibilitat per valors petits

	IEEE75464ZeroBits uint64 = 0x0000000000000000 // 0.0 segons l'estàndard IEEE 754
	IEEE75432ZeroBits uint32 = 0x00000000         // 0.0 segons l'estàndard IEEE 754
)

// Màscares principals de bits
const (
	SequenceType64Shift int = 52 // Desplaçament per al tipus de seqüència
	ElementType64Shift  int = 48 // Desplaçament per al tipus d'element

	Saturation32Mask    uint64 = 0x0000800000000000 // Exemple: Bit que indica saturació
	NullFlag32Mask      uint64 = 0x0000400000000000 // Exemple: Bit per valors nuls
	UnitFlag32Mask      uint64 = 0x0000200000000000 // Exemple: Bit per valors ±1.0
	SequenceType32Shift int    = 52                 // Desplaçament per al tipus de seqüència
	ElementType32Shift  int    = 48                 // Desplaçament per al tipus d'element
	Group32Mask         uint64 = 0xF000000000000000 // Màscara general per identificar el grup
	Subgroup32Mask      uint64 = 0x0F00000000000000 // Màscara per identificar subgrups
)
