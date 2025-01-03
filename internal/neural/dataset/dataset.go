package dataset

import (
	rF64 "github.com/jibort/ld_mcac/internal/core/RF64"
)

// Dataset representa el conjunt de dades amb valors RangeF64.
type Dataset struct {
	Inputs  [][]rF64.RangeF64 // Entrades: Paraules com a seqüències de RangeF64.
	Targets [][]rF64.RangeF64 // Targets: Síl·labes com a seqüències de RangeF64.
}

// NewDatasetFromCorpus genera un Dataset a partir d'un map de paraules i síl·labes.
func NewDatasetFromCorpus(corpus map[string]string) *Dataset {
	var inputs [][]rF64.RangeF64
	var targets [][]rF64.RangeF64

	for word, syllables := range corpus {
		// Convertim la paraula a RangeF64 (símbols).
		wordSymbols := convertToRangeF64(word)

		// Convertim les síl·labes a RangeF64 (tokens separats per '_').
		syllableSymbols := convertToRangeF64(syllables)

		// Afegim les dades al Dataset.
		inputs = append(inputs, wordSymbols)
		targets = append(targets, syllableSymbols)
	}

	return &Dataset{
		Inputs:  inputs,
		Targets: targets,
	}
}

// convertToRangeF64 converteix una cadena en una seqüència de símbols RangeF64.
func convertToRangeF64(input string) []rF64.RangeF64 {
	var result []rF64.RangeF64
	for _, char := range input {
		// Convertim cada caràcter a un símbol del grup B.1 (definit en RangeF64).
		result = append(result, rF64.NewRangeF64FromSymbol(char))
	}
	return result
}
