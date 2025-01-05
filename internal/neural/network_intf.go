// Representació funcional d'una xarxa neuronal.
// CreatedAt: 2024/12/13 dv. JIQ[GPT]

package neural

import (
	base "github.com/jibort/ld_mcac/internal/core/intf/base"
)

type NetworkIntf interface {
	Train(pConfig NetworkConfig) (rSuccess bool)
	Evaluate(pInput []base.RangeIntf) (rOutput []base.RangeIntf)
	Forward(inputs []base.RangeIntf) []base.RangeIntf // Propaga els inputs a través de la xarxa
	Backward(errors []base.RangeIntf)                 // Retropropaga els errors
	AddLayer(layer Layer)                             // Afegeix una capa a la xarxa
	Layers() []Layer                                  // Retorna totes les capes de la xarxa
}
