// Funció Neuronal 'Linear' segons el model del projecte.
// CreatedAt: 2024/12/13 dv. GPT

package FNs

import intf "github.com/jibort/ld_mcac/internal/core/intf"

type Linear_nf struct {
	nfs []NeuralFunctionIntf
}

func NewLinear_nf() *Linear_nf {
	return &Linear_nf{
		nfs: make([]NeuralFunctionIntf, 0),
	}
}

func (lNF *Linear_nf) Forward(pInput intf.RangeIntf) intf.RangeIntf {
	for _, nf := range lNF.nfs {
		pInput = nf.Forward(pInput)
	}
	return pInput
}

func (lNF *Linear_nf) Backward(pOutput intf.RangeIntf) intf.RangeIntf {
	for idx := len(lNF.nfs) - 1; idx >= 0; idx-- {
		pOutput = lNF.nfs[idx].Backward(pOutput)
	}
	return pOutput
}

func (lNF *Linear_nf) Append(pNew NeuralFunctionIntf) {
	if pNew == nil {
		panic("Cannot append a nil NeuralFunctionIntf")
	}
	lNF.nfs = append(lNF.nfs, pNew)
}
