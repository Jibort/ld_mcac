// Representació funcional d'una xarxa neuronal.
// CreatedAt: 2024/12/13 dv. JIQ[GPT]

package neural

import (
	intf "github.com/jibort/ld_mcac/internal/core/Intf"
)

type NetworkIntf interface {
	Train(pConfig NetworkConfig) (rSuccess bool)
	Evaluate(pInput []intf.RangeIntf) (rOutput []intf.RangeIntf)
	Forward(inputs []intf.RangeIntf) []intf.RangeIntf // Propaga els inputs a través de la xarxa
	Backward(errors []intf.RangeIntf)                 // Retropropaga els errors
	AddLayer(layer Layer)                             // Afegeix una capa a la xarxa
	Layers() []Layer                                  // Retorna totes les capes de la xarxa
}
