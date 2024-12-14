// Representació d'una neurona dins una xarxa neuronal.
// CreatedAt: 2024/12/13 dv. GPT

package neural

import (
	"github.com/jibort/ld_mcac/internal/core"
	"github.com/jibort/ld_mcac/internal/neural/FNs"
)

type Neuron struct {
	Inputs []*Synapse
	Bias   *core.RangeIntf
	FNL    FNs.NeuralFunctionIntf
}

func (n *Neuron) Compute(inputs []core.RangeIntf) core.RangeIntf {
	sum := *n.Bias
	for _, synapse := range n.Inputs {
		sum.SetF64Value(sum.GetF64Value() + synapse.Compute(inputs).GetF64Value())
	}
	return n.FNL.Forward(sum)
}
