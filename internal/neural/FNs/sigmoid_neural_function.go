// Funció Neuronal 'Sigmòide' segons el model del projecte.
// CreatedAt: 2024/12/13 dv. GPT

package FNs

import (
	"math"

	"github.com/jibort/ld_mcac/internal/core"
)

type Sigmoid_nf struct {
	nfs []NeuralFunctionIntf
}

func NewSigmoid_nf() *Sigmoid_nf {
	return &Sigmoid_nf{
		nfs: make([]NeuralFunctionIntf, 0),
	}
}

func (sNF *Sigmoid_nf) Forward(pInput core.RangeIntf) core.RangeIntf {
	value := pInput.GetF64Value()
	original := 1.0 / (1.0 + math.Exp(-value))
	scaled := -1.0 + (original-0.268941)*(2.0/0.462118) // Escalat al rang [-1.0, +1.0]
	pInput.SetF64Value(scaled)

	for _, nf := range sNF.nfs {
		pInput = nf.Forward(pInput)
	}
	return pInput
}

func (sNF *Sigmoid_nf) Backward(pOutput core.RangeIntf) core.RangeIntf {
	for idx := len(sNF.nfs) - 1; idx >= 0; idx-- {
		pOutput = sNF.nfs[idx].Backward(pOutput)
	}

	value := pOutput.GetF64Value()
	original := value * (1.0 - value) // Derivada de sigmoid
	pOutput.SetF64Value(original)
	return pOutput
}

func (sNF *Sigmoid_nf) Append(pNew NeuralFunctionIntf) {
	if pNew == nil {
		panic("Cannot append a nil NeuralFunctionIntf")
	}
	sNF.nfs = append(sNF.nfs, pNew)
}
