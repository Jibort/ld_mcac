// Representació d'una neurona dins una xarxa neuronal.
// CreatedAt: 2024/12/13 dv. GPT

package neural

import "github.com/jibort/ld_mcac/internal/core"

type Synapse struct {
	Weight core.RangeIntf
	Input  *Neuron
}

func NewSynapse(pWeight core.RangeIntf, pInput *Neuron) *Synapse {
	return &Synapse{
		Weight: pWeight,
		Input:  pInput,
	}
}

func (sSyn *Synapse) Compute(pInputs []core.RangeIntf) core.RangeIntf {
	return core.NewRangeF64(sSyn.Weight.GetF64Value() * pInputs[0].GetF64Value())
}
