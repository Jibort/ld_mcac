// Representació d'una neurona dins una xarxa neuronal.
// CreatedAt: 2024/12/13 dv. GPT

package neural

import (
	rF64 "github.com/jibort/ld_mcac/internal/core/RF64"
	base "github.com/jibort/ld_mcac/internal/core/intf/base"
)

type Synapse struct {
	Weight base.RangeIntf
	Input  *Neuron
}

func NewSynapse(pWeight base.RangeIntf, pInput *Neuron) *Synapse {
	return &Synapse{
		Weight: pWeight,
		Input:  pInput,
	}
}

func (sSyn *Synapse) Compute(pInputs []base.RangeIntf) base.RangeIntf {
	res := rF64.NewF64Range(sSyn.Weight.AsFloat64() * pInputs[0].AsFloat64())
	return &res
}
