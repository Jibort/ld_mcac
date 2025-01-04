// Interfície per a la comparació d'instàncies dels tipus Rang.
// CreatedAt: 2024/12/08 dg. JIQ

package base

// Interície 'Comparable'
type ComparableIntf interface {
	Equals(pOther RangeIntf) bool             // Cert només si pOther correspon al mateix tipus i contingut.
	LessThan(pOther RangeIntf) bool           // Cert només si la instància és menor que pOther.
	LessOrEqualThan(pOther RangeIntf) bool    // Cert només si la instància és menor o igual que pOther.
	GreaterThan(pOther RangeIntf) bool        // Cert només si la instància és major que pOther.
	GreaterOrEqualThan(pOther RangeIntf) bool // Cert només si la instància és major o igual que pOther.
}
