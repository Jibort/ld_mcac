// Interfícies de Range pel Grup D en float64.
// CreatedAt: 2025/01/03 dv. JIQ

package Intf

// Interfície per a valors de 64 bits (float64) del Grup D (metareferències).
type RangeF64MetaRefIntf interface {
	RangeF64Intf
	Network() *int // Retorna el número de xarxa (o 0 si és la xarxa actual)
	Layer() *int   // Retorna el número de capa (o 0 si la capa és nul)
	Neuron() *int  // Retorna el número de neurona (o 0 si la neurona és nul)
	Synapse() *int // Retorna el número de sinapsi (o 0 si la sinapsi és nul)
}
