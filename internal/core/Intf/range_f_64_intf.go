// Interfícies de Range per a F64 bits.
// CreatedAt: 2025/01/03 dv. JIQ

package Intf

// Interfície per a valors fixed o floating points de 64bits.
type RangeF64Intf interface {
	Range64Intf

	ToIntf() RangeIntf   // Retorna la instància com a RangeIntf
	ToF32() RangeF32Intf // Retorna la instància com a RangeF32Intf
}
