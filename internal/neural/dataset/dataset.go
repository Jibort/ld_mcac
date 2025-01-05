package dataset

import (
	syms "github.com/jibort/ld_mcac/internal/core/RF64Sym"
)

// Dataset representa el conjunt de dades amb valors F64Range.
type Dataset struct {
	Inputs  [][]syms.F64RangeSymbol // Entrades: Paraules com a seqüències de F64Range.
	Targets [][]syms.F64RangeSymbol // Targets: Síl·labes com a seqüències de F64Range.
}

// NewDatasetFromCorpus genera un Dataset a partir d'un map de paraules i síl·labes.
func NewDatasetFromCorpus(corpus map[string]string) *Dataset {
	var inputs [][]syms.F64RangeSymbol
	var targets [][]syms.F64RangeSymbol

	for word, syllables := range corpus {
		// Convertim la paraula a F64Range (símbols).
		wordSymbols := convertToF64Symbol(word)

		// Convertim les síl·labes a F64Range (tokens separats per '_').
		syllableSymbols := convertToF64Symbol(syllables)

		// Afegim les dades al Dataset.
		inputs = append(inputs, wordSymbols)
		targets = append(targets, syllableSymbols)
	}

	return &Dataset{
		Inputs:  inputs,
		Targets: targets,
	}
}

// convertToF64Symbol converteix una cadena en una seqüència de símbols F64Range.
func convertToF64Symbol(input string) []syms.F64RangeSymbol {
	var result []syms.F64RangeSymbol
	for _, char := range input {
		// Convertim cada caràcter a un símbol del grup B.1 (definit en F64Range).
		result = append(result, syms.NewF64RangeSymbol(char).(syms.F64RangeSymbol))
	}
	return result
}
