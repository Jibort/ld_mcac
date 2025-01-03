// Interfícies de Range per a F64 bits.
// CreatedAt: 2025/01/03 dv. JIQ

package intf

// Interfície per a valors fixed o floating points de 64bits.
type RangeF64Intf interface {
	Range64Intf

	IsPositive() bool    // Cert només si el valor es positiu
	IsNegative() bool    // Cert només si el valor es negatiu o zero
	ToF32() RangeF32Intf // Retorna la instància com a RangeF32Intf
}
