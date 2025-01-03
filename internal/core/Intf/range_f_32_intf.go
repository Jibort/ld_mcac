// Interfícies de Range per a float32.
// CreatedAt: 2025/01/03 dv. JIQ

package Intf

// Interfície per a valors de 32 bits decimals.
type RangeF32Intf interface {
	Range32Intf
	MathOperationsIntf

	ToIntf() RangeIntf   // Retorna la instància com a RangeIntf
	ToF64() RangeF64Intf // Retorna la instància com a RangeF64Intf
}

// Interfície per a valors de 32 bits (float32) del Grup A.
type RangeF32GroupAIntf interface {
	RangeF32Intf // Hereta les funcions generals per a float32
	GroupAIntf   // Hereta les funcions generals per a Grup A

	Exponent() uint8  // Retorna l'exponent com a uint8
	Mantissa() uint32 // Retorna la mantissa com a uint32
}
