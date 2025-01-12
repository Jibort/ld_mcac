// Funció Neuronal ReLU segons el model del projecte.
// CreatedAt: 2024/12/13 dv. GPT

package FNs

import (
	intf "github.com/jibort/ld_mcac/internal/core/intf/base"
)

// ReLUNeuralFunc implementa la interfície NeuralFunctionIntf
type ReLU_nf struct {
	// Llista de funcions neuronals a executar després de la pròpia funció.
	nfs []NeuralFunctionIntf
}

// Crea una nova instància de la Funció Neuronal ReLU.
func NewReLU_nf() *ReLU_nf {
	return &ReLU_nf{
		nfs: make([]NeuralFunctionIntf, 0),
	}
}

// [NeuralFunctionIntf] Forward aplica la funció ReLU i retorna el valor processat.
func (sNF *ReLU_nf) Forward(pInput intf.RangeIntf) intf.RangeIntf {
	// Executem la funció
	value := pInput.AsFloat64()
	if value < 0.0 {
		value = 0.0
	}
	pInput.SetFloat64(value) // Modifiquem l'objecte existent

	// Fem forward de la llista ordenada de funcions neuronals.
	for _, nf := range sNF.nfs {
		pInput = nf.Forward(pInput)
	}

	return pInput
}

// [NeuralFunctionIntf] Backward aplica la derivada de la funció ReLU i retorna el valor processat.
func (sNF *ReLU_nf) Backward(pOutput intf.RangeIntf) intf.RangeIntf {
	// Fem backward de la llista inversa ordenada de funcions neuronals.
	for idx := len(sNF.nfs) - 1; idx >= 0; idx-- {
		pOutput = sNF.nfs[idx].Backward(pOutput)
	}

	// Executem el backward de la pròpia funció.
	value := pOutput.AsFloat64()
	if value > 0.0 {
		return pOutput
	}
	pOutput.SetFloat64(0.0) // Modifiquem l'objecte existent
	return pOutput
}

// [NeuralFunctionIntf] Append afegeix una funció neuronal a la cua, amb validació.
func (sNF *ReLU_nf) Append(pNew NeuralFunctionIntf) {
	if pNew == nil {
		panic("Cannot append a nil NeuralFunctionIntf")
	}

	sNF.nfs = append(sNF.nfs, pNew)
}
