// Representació d'una neurona dins una xarxa neuronal.
// CreatedAt: 2024/12/13 dv. GPT

package neural

import (
	intf "github.com/jibort/ld_mcac/internal/core/Intf"
	"github.com/jibort/ld_mcac/internal/neural/FNs"
)

type Neuron struct {
	Inputs []*Synapse
	Bias   intf.RangeIntf
	FNL    FNs.NeuralFunctionIntf
}

func (n *Neuron) Compute(inputs []intf.RangeIntf) intf.RangeIntf {
	sum := n.Bias
	for _, synapse := range n.Inputs {
		sum.SetFloat64(sum.AsFloat64() + synapse.Compute(inputs).AsFloat64())
	}
	return n.FNL.Forward(sum)
}
