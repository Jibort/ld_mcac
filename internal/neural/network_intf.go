// Representació funcional d'una xarxa neuronal.
// CreatedAt: 2024/12/13 dv. JIQ[GPT]

package neural

import "github.com/jibort/ld_mcac/internal/core"

type NetworkIntf interface {
	Forward(inputs []core.RangeIntf) []core.RangeIntf // Propaga els inputs a través de la xarxa
	Backward(errors []core.RangeIntf)                 // Retropropaga els errors
	AddLayer(layer LayerIntf)                         // Afegeix una capa a la xarxa
	Layers() []LayerIntf                              // Retorna totes les capes de la xarxa
}
