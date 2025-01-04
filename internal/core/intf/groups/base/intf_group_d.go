// Interfície general de tots els Subgrups D (metareferències).
// CreatedAt: 2025/01/04 ds. JIQ

package base

// Interfície general de tots els Subgrups D metareferències).
type GroupDIntf interface {
	IsNetworkId() bool // Cert només si la instància identifica una xarxa.
	IsLayerId() bool   // Cert només si la instància identifica una capa de xarxa.
	IsNeuron() bool    // Cert només si la instància identifica una neurona dins una capa de xarxa.
	IsSynapse() bool   // Cert només si la instància identifica una sinapsi d'entrada d'una neurona dins una capa de xarxa.
}
