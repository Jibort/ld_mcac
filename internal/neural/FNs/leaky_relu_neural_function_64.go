// Funció Neuronal ReLU segons el model del projecte.
// CreatedAt: 2024/12/13 dv. GPT

package FNs

import intf "github.com/jibort/ld_mcac/internal/core/intf/base"

const (
	alpha = 0.01
)

// LeakyReLU_nf implementa la interfície NeuralFunctionIntf
type LeakyReLU_nf struct {
	// Llista de funcions neuronals a executar després de la pròpia funció.
	nfs []NeuralFunctionIntf
}

// Crea una nova instància de la Funció Neuronal ReLU.
func NewLeakyReLU_nf() *LeakyReLU_nf {
	return &LeakyReLU_nf{
		nfs: make([]NeuralFunctionIntf, 0),
	}
}

// [NeuralFunctionIntf] Forward aplica la funció ReLU i retorna el valor processat.
// Forward aplica la funció Leaky ReLU al valor d'entrada
func (sNF *LeakyReLU_nf) Forward(pInput intf.RangeIntf) intf.RangeIntf {
	value := pInput.AsFloat64()
	if value < 0 {
		pInput.SetFloat64(value * alpha) // Aplicar alpha per a valors negatius
	}
	return pInput // Retorna l'input modificat
}

// [NeuralFunctionIntf] Backward aplica la derivada de la funció Leaky ReLU i retorna el valor processat.
func (sNF *LeakyReLU_nf) Backward(pOutput intf.RangeIntf) intf.RangeIntf {
	value := pOutput.AsFloat64()
	if value < 0 {
		pOutput.SetFloat64(alpha) // Derivada per a valors negatius és alpha
	} else {
		pOutput.SetFloat64(1.0) // Derivada per a valors positius és 1
	}
	return pOutput // Retorna l'output modificat
}

// [NeuralFunctionIntf] Append afegeix una funció neuronal a la cua, amb validació.
func (sNF *LeakyReLU_nf) Append(pNew NeuralFunctionIntf) {
	if pNew == nil {
		panic("Cannot append a nil NeuralFunctionIntf")
	}

	sNF.nfs = append(sNF.nfs, pNew)
}
