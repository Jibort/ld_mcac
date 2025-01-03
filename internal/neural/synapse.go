// Representació d'una neurona dins una xarxa neuronal.
// CreatedAt: 2024/12/13 dv. GPT

package neural

import (
	rF64 "github.com/jibort/ld_mcac/internal/core/RF64"
	intf "github.com/jibort/ld_mcac/internal/core/intf"
)

type Synapse struct {
	Weight intf.RangeIntf
	Input  *Neuron
}

func NewSynapse(pWeight intf.RangeIntf, pInput *Neuron) *Synapse {
	return &Synapse{
		Weight: pWeight,
		Input:  pInput,
	}
}

func (sSyn *Synapse) Compute(pInputs []intf.RangeIntf) intf.RangeIntf {
	return rF64.NewRangeF64(sSyn.Weight.AsFloat64() * pInputs[0].AsFloat64())
}
