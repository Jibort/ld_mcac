// Representació funcional d'una xarxa neuronal.
// CreatedAt: 2024/12/13 dv. JIQ[GPT]

package neural

import (
	base "github.com/jibort/ld_mcac/internal/core/intf/base"
)

// type LayerIntf interface {
// 	Forward(inputs []core.RangeIntf) []core.RangeIntf  // Processa els inputs
// 	Backward(errors []core.RangeIntf) []core.RangeIntf // Retropropaga els errors
// }

type Layer struct {
	neurons  []*Neuron
	synapses [][]*Synapse
}

// Crea una nova capa.
func NewLayer(neurons []*Neuron, synapses [][]*Synapse) Layer {
	return Layer{
		neurons:  neurons,
		synapses: synapses,
	}
}

// Neurons retorna les neurones de la capa.
func (sLay *Layer) Neurons() []*Neuron {
	return sLay.neurons
}

// Synapses retorna les sinapsis de la capa.
func (sLay *Layer) Synapses() [][]*Synapse {
	return sLay.synapses
}

// Forward executa la propagació endavant per la capa.
func (sLay *Layer) Forward(inputs []base.RangeIntf) []base.RangeIntf {
	outputs := make([]base.RangeIntf, len(sLay.neurons))
	for i, neuron := range sLay.neurons {
		outputs[i] = neuron.Compute(inputs)
	}
	return outputs
}

// Backward executa la retropropagació per la capa.
func (sLay *Layer) Backward(errors []base.RangeIntf) []base.RangeIntf {
	// Retropropagació per les neurones
	// Retorna els errors per la capa anterior.
	return errors
}
