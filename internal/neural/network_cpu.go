// Representació d'una xarxa neuronal optimitzada per a l'execució sobre CPU.
// CreatedAt: 2024/12/13 dv. JIQ[GPT]

package neural

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	base "github.com/jibort/ld_mcac/internal/core/intf/base"
	rF64 "github.com/jibort/ld_mcac/internal/core/rf64"
	"github.com/jibort/ld_mcac/internal/neural/FNs"
)

type NetworkCPU struct {
	layers   []Layer
	neurons  [][]Neuron
	synapses [][][]Synapse
}

// Crea una nova xarxa buida.
func NewNetworkCPU(ndlPath string) (*NetworkCPU, error) {
	// Obrim i deserialitzem el fitxer .ndl
	file, err := os.Open(ndlPath)
	if err != nil {
		return nil, fmt.Errorf("error opening .ndl file: %v", err)
	}
	defer file.Close()

	var ndlConfig struct {
		Name   string `json:"name"`
		Layers []int  `json:"layers"`
		Links  []struct {
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"links"`
	}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&ndlConfig); err != nil {
		return nil, fmt.Errorf("error decoding .ndl file: %v", err)
	}

	// Inicialitzem la xarxa
	network := &NetworkCPU{
		neurons:  make([][]Neuron, len(ndlConfig.Layers)),
		synapses: make([][][]Synapse, len(ndlConfig.Layers)-1),
	}

	// Creem les capes i neurones
	for _, neuronCount := range ndlConfig.Layers {
		layer := Layer{
			neurons:  make([]*Neuron, neuronCount),
			synapses: make([][]*Synapse, neuronCount),
		}

		// Inicialitzem les neurones i assignem a la capa
		for neuronIdx := 0; neuronIdx < neuronCount; neuronIdx++ {
			res := rF64.NewF64Range(0.0)
			neuron := Neuron{
				Inputs: []*Synapse{},
				Bias:   &res,             // Bias inicialitzat a 0.0
				FNL:    FNs.NewReLU_nf(), // Funció neuronal predeterminada
			}
			layer.neurons[neuronIdx] = &neuron
		}

		// Afegim la capa a la xarxa
		network.layers = append(network.layers, layer)
	}

	// Creem les sinapsis segons els links
	for _, link := range ndlConfig.Links {
		fromLayer, fromNeuron, err := parseLinkEndpoint(link.From)
		if err != nil {
			return nil, fmt.Errorf("error parsing 'from' endpoint: %v", err)
		}
		toLayer, toNeuron, err := parseLinkEndpoint(link.To)
		if err != nil {
			return nil, fmt.Errorf("error parsing 'to' endpoint: %v", err)
		}

		if fromNeuron == "n" || toNeuron == "n" {
			// Connexió de totes les neurones
			for fIdx, fromNeuron := range network.layers[fromLayer].neurons {
				for tIdx := range network.layers[toLayer].neurons {
					w := rF64.NewF64Range(0.0)
					synapse := Synapse{
						Weight: &w, // Pes inicialitzat a 0.0
						Input:  fromNeuron,
					}
					network.layers[toLayer].synapses[tIdx] = append(network.layers[toLayer].synapses[tIdx], &synapse)
					network.layers[fromLayer].neurons[fIdx].Inputs = append(network.layers[fromLayer].neurons[fIdx].Inputs, &synapse)
				}
			}
		} else {
			// Connexió específica
			fIdx, _ := strconv.Atoi(fromNeuron)
			tIdx, _ := strconv.Atoi(toNeuron)
			w := rF64.NewF64Range(0.0)
			synapse := Synapse{
				Weight: &w,
				Input:  network.layers[fromLayer].neurons[fIdx],
			}
			network.layers[toLayer].synapses[tIdx] = append(network.layers[toLayer].synapses[tIdx], &synapse)
			network.layers[fromLayer].neurons[fIdx].Inputs = append(network.layers[fromLayer].neurons[fIdx].Inputs, &synapse)
		}
	}

	return network, nil
}

// AddLayer afegeix una capa a la xarxa i actualitza les instàncies.
// func (sNet *NetworkCPU) AddLayer(pLayer Layer) {
// 	sNet.layers = append(sNet.layers, pLayer)
// 	sNet.neurons = append(sNet.neurons, pLayer.Neurons())
// 	sNet.synapses = append(sNet.synapses, pLayer.Synapses())
// }

// Forward executa la propagació endavant.
func (sNet *NetworkCPU) Forward(inputs []base.RangeIntf) []base.RangeIntf {
	currentInputs := inputs
	for _, layer := range sNet.layers {
		currentInputs = layer.Forward(currentInputs)
	}
	return currentInputs
}

// Backward executa la retropropagació.
func (n *NetworkCPU) Backward(errors []base.RangeIntf) {
	currentErrors := errors
	for i := len(n.layers) - 1; i >= 0; i-- {
		currentErrors = n.layers[i].Backward(currentErrors)
	}
}

func parseLinkEndpoint(endpoint string) (layerIdx int, neuronIdx string, err error) {
	// Dividim l'endpoint en parts separades per ","
	parts := strings.Split(endpoint, ",")
	if len(parts) != 2 {
		return 0, "", fmt.Errorf("invalid format: expected 'L:x, N:y', got '%s'", endpoint)
	}

	// Analitzem la capa (L:x)
	if !strings.HasPrefix(parts[0], "L:") {
		return 0, "", fmt.Errorf("invalid layer format: missing 'L:' in '%s'", parts[0])
	}
	layerStr := strings.TrimPrefix(parts[0], "L:")
	layerIdx, err = strconv.Atoi(layerStr)
	if err != nil {
		return 0, "", fmt.Errorf("invalid layer index: '%s'", layerStr)
	}

	// Analitzem la neurona (N:y)
	if !strings.HasPrefix(parts[1], "N:") {
		return 0, "", fmt.Errorf("invalid neuron format: missing 'N:' in '%s'", parts[1])
	}
	neuronIdx = strings.TrimPrefix(parts[1], "N:")

	return layerIdx, neuronIdx, nil
}
