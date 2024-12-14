// Definició de la interfície per a implementar funcions neuronals.
// CreatedAt: 2024/12/13 dv. JIQ

package FNs

import (
	"github.com/jibort/ld_mcac/internal/core"
)

type NeuralFunctionIntf interface {
	// Permet executar la funció neuronal i a continuació la resta ordenada de funcions.
	Forward(pInput core.RangeIntf) (rResult core.RangeIntf)

	// Permet executar en ordre invers la derivada de la funció neuronal i a continuació la resta ordenada de funcions.
	Backward(pOutput core.RangeIntf) (rResult core.RangeIntf)

	// Afegeix la següent Funció Neuronal a la cua de la llista.
	Append(NeuralFunctionIntf)

	// // Retorna la primera funció neuronal de la llista (l'única en que rPrev seria nul).
	// First() *NeuralFunctionIntf

	// // Retorna l'última funció neuronal de la llista (l'única en que rNext seria nul).
	// Last() *NeuralFunctionIntf
}
