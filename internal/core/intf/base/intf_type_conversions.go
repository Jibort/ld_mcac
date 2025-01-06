// Interfície per a la conversió entre instàncies dels tipus Rang i tipus base de Golang.
// CreatedAt: 2025/01/04 ds. JIQ

package base

// Interfície 'ConversionsIntf'
type ConversionsIntf interface {
	AsFloat64() float64
	SetFloat64(float64)
	AsUint64() uint64
	SetUint64(uint64)

	AsFloat32() float32
	SetFloat32(float32)
	AsUint32() uint32
	SetUint32(uint32)
}
