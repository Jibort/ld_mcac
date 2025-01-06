// Interfície general de Range.
// CreatedAt: 2025/01/04 ds. JIQ

package base

// INTERFÍCIE GLOBAL PER A QUALSEVOL INSTÀNCIA DE 'Range'
type RangeIntf interface {
	ClonableIntf
	ConversionsIntf
	ComparableIntf
	GroupableIntf
	RangebleIntf

	IsError() bool // Cert només si la instància correspon a un error.
}
