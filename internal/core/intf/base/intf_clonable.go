// Interfície per a la clonació d'instàncies dels tipus Rang.
// CreatedAt: 2025/01/04 ds. JIQ

package base

// Interfície 'Clonable'
type ClonableIntf interface {
	Clone() RangeIntf // Clona la instància de Rang que implementa 'ClonableIntf'.
}
