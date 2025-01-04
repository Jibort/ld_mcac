// Interfície per a la clonació d'instàncies dels tipus Rang.
// CreatedAt: 2025/01/04 ds. JIQ

package base

// Interfície 'MathOperationsIntf'
type MathOperationsIntf interface {
	ComparableIntf

	Add(pOther RangeIntf) RangeIntf
	Sub(pOther RangeIntf) RangeIntf
	Mul(pOther RangeIntf) RangeIntf
	Div(pOther RangeIntf) RangeIntf

	IsF32() bool // Cert només si la instància és de 64 bits en punt flotant.
	IsU32() bool // Cert només si la instància és de 64 bits en enter sense signe.
	IsF64() bool // Cert només si la instància és de 64 bits en punt flotant.
	IsU64() bool // Cert només si la instància és de 64 bits en enter sense signe.
}
