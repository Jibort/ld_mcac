// Interfície d'error general
// CreatedAt: 2025/01/03 dv. JIQ

package Intf

// Interfície per a valors de 64 bits (float64) del Subgrup B.4 (errors)
type Error64Intf interface {
	RangeIntf

	Critical() bool      // Cert només si l'error és crític
	Code() uint16        // Retorna el codi d'error
	Arguments() []uint64 // Retorna els arguments de l'error
}
