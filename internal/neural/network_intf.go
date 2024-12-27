// Representació funcional d'una xarxa neuronal.
// CreatedAt: 2024/12/13 dv. JIQ[GPT]

package neural

import (
	"github.com/jibort/ld_mcac/internal/core"
)

type NetworkIntf interface {
	Train(pConfig core.NetworkConfig) (rSuccess bool)
	Evaluate(pInput []core.RangeF64) (rOutput []core.RangeF64)
	Forward(inputs []core.RangeIntf) []core.RangeIntf // Propaga els inputs a través de la xarxa
	Backward(errors []core.RangeIntf)                 // Retropropaga els errors
	AddLayer(layer Layer)                             // Afegeix una capa a la xarxa
	Layers() []Layer                                  // Retorna totes les capes de la xarxa
}
