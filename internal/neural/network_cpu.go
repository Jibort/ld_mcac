// Representació d'una xarxa neuronal optimitzada per a l'execució sobre CPU.
// CreatedAt: 2024/12/13 dv. JIQ[GPT]

package neural

import "github.com/jibort/ld_mcac/internal/core"

type NetworkCPU struct {
	layers   []Layer
	neurons  [][]Neuron
	synapses [][][]Synapse
}

// Crea una nova xarxa buida.
func NewNetworkCPU(pConfigPath string) (*NetworkCPU, error) {
	// Llegeix i parseja el fitxer descriptiu.
	config, err := parseConfig(pConfigPath)
	if err != nil {
		return nil, err
	}

	// Crea la xarxa a partir del fitxer.
	network := &NetworkCPU{
		layers:   make([]Layer, len(config.Layers)),
		neurons:  make([][]Neuron, len(config.Layers)),
		synapses: make([][][]Synapse, len(config.Layers)-1),
	}

	// Crea les capes i neurones.
	for i, layerConfig := range config.Layers {
		neurons := make([]Neuron, layerConfig.Neurons)
		for j := 0; j < layerConfig.Neurons; j++ {
			neurons[j] = Neuron{
				Bias:   core.NewRangeF64(0.0),
				FNL:    createActivationFunction(layerConfig.Activation),
				Inputs: make([]*Synapse, 0),
			}
		}
		network.neurons[i] = neurons
		network.layers[i] = Layer{neurons: neurons}
	}

	// Crea les sinapsis.
	for _, conn := range config.Connections {
		fromLayer := conn.FromLayer
		toLayer := conn.ToLayer

		synapses := make([][]Synapse, len(network.neurons[fromLayer]))
		for i := range network.neurons[fromLayer] {
			synapses[i] = make([]Synapse, len(network.neurons[toLayer]))
			for j := range network.neurons[toLayer] {
				synapses[i][j] = Synapse{
					Weight: core.NewRangeF64(conn.Weights[i][j]),
					Input:  &network.neurons[fromLayer][i],
				}
				network.neurons[toLayer][j].Inputs = append(network.neurons[toLayer][j].Inputs, &synapses[i][j])
			}
		}
		network.synapses[fromLayer] = synapses
	}

	return network, nil
}

// AddLayer afegeix una capa a la xarxa i actualitza les instàncies.
func (sNet *NetworkCPU) AddLayer(pLayer Layer) {
	sNet.layers = append(sNet.layers, pLayer)
	sNet.neurons = append(sNet.neurons, pLayer.Neurons())
	sNet.synapses = append(sNet.synapses, pLayer.Synapses())
}

// Forward executa la propagació endavant.
func (sNet *NetworkCPU) Forward(inputs []core.RangeIntf) []core.RangeIntf {
	currentInputs := inputs
	for _, layer := range sNet.layers {
		currentInputs = layer.Forward(currentInputs)
	}
	return currentInputs
}

// Backward executa la retropropagació.
func (n *NetworkCPU) Backward(errors []core.RangeIntf) {
	currentErrors := errors
	for i := len(n.layers) - 1; i >= 0; i-- {
		currentErrors = n.layers[i].Backward(currentErrors)
	}
}
