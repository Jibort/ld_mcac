// Interfície d'instàncies de 64 o 32 bits.
// CreatedAt: 2025/01/04 ds. JIQ

package base

type RangebleIntf interface {
	Is64() bool // Cert només si la instància és de 64 bits.
	Is32() bool // Cert només si la instància és de 32 bits.
}
