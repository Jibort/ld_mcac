// Funció Neuronal 'TanH' segons el model del projecte.
// CreatedAt: 2024/12/13 dv. GPT

package FNs

import (
	"math"

	intf "github.com/jibort/ld_mcac/internal/core/intf"
)

type Tanh_nf struct {
	nfs []NeuralFunctionIntf
}

func NewTanh_nf() *Tanh_nf {
	return &Tanh_nf{
		nfs: make([]NeuralFunctionIntf, 0),
	}
}

func (tNF *Tanh_nf) Forward(pInput intf.RangeIntf) intf.RangeIntf {
	value := pInput.AsFloat64()
	original := math.Tanh(value)
	scaled := -1.0 + (original+0.761594)*(2.0/1.523188) // Escalat al rang [-1.0, +1.0]
	pInput.SetFloat64(scaled)

	for _, nf := range tNF.nfs {
		pInput = nf.Forward(pInput)
	}
	return pInput
}

func (tNF *Tanh_nf) Backward(pOutput intf.RangeIntf) intf.RangeIntf {
	for idx := len(tNF.nfs) - 1; idx >= 0; idx-- {
		pOutput = tNF.nfs[idx].Backward(pOutput)
	}

	value := pOutput.AsFloat64()
	original := 1.0 - math.Pow(math.Tanh(value), 2.0) // Derivada de tanh
	pOutput.SetFloat64(original)
	return pOutput
}

func (tNF *Tanh_nf) Append(pNew NeuralFunctionIntf) {
	if pNew == nil {
		panic("Cannot append a nil NeuralFunctionIntf")
	}
	tNF.nfs = append(tNF.nfs, pNew)
}
