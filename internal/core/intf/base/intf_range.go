// Interfície general de Range.
// CreatedAt: 2025/01/04 ds. JIQ

package base

// INTERFÍCIE GLOBAL PER A QUALSEVOL INSTÀNCIA DE 'Range'
type RangeIntf interface {
	TypeConversionsIntf
	ClonableIntf
	ComparableIntf

	// Blocs
	Is64() bool    // Cert només si la instància és de 64 bits.
	Is32() bool    // Cert només si la instància és de 32 bits.
	IsError() bool // Cert només si la instància correspon a un error.

	// Grups
	IsGroupA() bool // Cert només si la instància pertany al grup A.
	IsGroupB() bool // Cert només si la instància pertany al grup B.
	IsGroupC() bool // Cert només si la instància pertany al grup C.
	IsGroupD() bool // Cert només si la instància pertany al grup D.
	IsGroupE() bool // Cert només si la instància pertany al grup E (només per a 64bits).
}
