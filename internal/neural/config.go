package neural

type NetworkConfig struct {
	Layers          int    `json:"layers"`            // Nombre de capes
	NeuronsPerLayer []int  `json:"neurons_per_layer"` // Nombre de neurones per cada capa
	Activation      string `json:"activation"`        // Funció d'activació predeterminada
}
