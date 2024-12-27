// Generació d'embeddings per a símbols textuals amb criteri fonètic i funcional
// CreatedAt: 24-12-2024 dg. GPT(JIQ)

package main

import (
	"fmt"

	"github.com/jibort/ld_mcac/internal/core"
)

// Tipus definit per RangF64 ja existent en un altre fitxer

type SymbolEmbedding struct {
	Symbol string
	Value  core.RangeF64
	Group  string // Categoria funcional o fonètica
}

func generateEmbeddings() []SymbolEmbedding {
	// Registre de valors usats per garantir que són únics
	usedValues := map[float64]string{}
	createRangeF64 := func(value float64, symbol string) core.RangeF64 {
		if existingSymbol, exists := usedValues[value]; exists {
			panic(fmt.Sprintf("El valor %.4f ja està assignat al símbol '%s' (intentat assignar a '%s')", value, existingSymbol, symbol))
		}
		usedValues[value] = symbol
		return core.NewRangeF64(value)
	}

	// Vocals distribuïdes segons lògica revisada
	alphabet := []SymbolEmbedding{
		{"a", createRangeF64(-0.90, "a"), "vowel"},
		{"à", createRangeF64(-0.9010, "à"), "vowel"},
		{"á", createRangeF64(-0.9020, "á"), "vowel"},
		{"e", createRangeF64(-0.60, "e"), "vowel"},
		{"è", createRangeF64(-0.6010, "è"), "vowel"},
		{"é", createRangeF64(-0.6020, "é"), "vowel"},
		{"i", createRangeF64(-0.30, "i"), "vowel"},
		{"í", createRangeF64(-0.3010, "í"), "vowel"},
		{"ï", createRangeF64(-0.3020, "ï"), "vowel"},
		{"o", createRangeF64(0.30, "o"), "vowel"},
		{"ò", createRangeF64(0.3010, "ò"), "vowel"},
		{"ó", createRangeF64(0.3020, "ó"), "vowel"},
		{"u", createRangeF64(0.60, "u"), "vowel"},
		{"ü", createRangeF64(0.6010, "ü"), "vowel"},
		{"ú", createRangeF64(0.6020, "ú"), "vowel"},
	}

	// Consonants distribuïdes segons grups revisats
	consonants := []SymbolEmbedding{
		{"b", createRangeF64(-0.85, "b"), "sonor_occlusive"},
		{"d", createRangeF64(-0.75, "d"), "sonor_occlusive"},
		{"g", createRangeF64(-0.65, "g"), "sonor_occlusive"},
		{"m", createRangeF64(-0.55, "m"), "nasal"},
		{"n", createRangeF64(-0.45, "n"), "nasal"},
		{"ñ", createRangeF64(-0.35, "ñ"), "nasal"},
		{"r", createRangeF64(-0.25, "r"), "liquid"},
		{"l", createRangeF64(-0.15, "l"), "liquid"},
		{"p", createRangeF64(0.34, "p"), "sorda_occlusive"},
		{"t", createRangeF64(0.44, "t"), "sorda_occlusive"},
		{"c", createRangeF64(0.50, "c"), "sorda_occlusive"},
		{"k", createRangeF64(0.54, "k"), "sorda_occlusive"},
		{"f", createRangeF64(0.64, "f"), "fricative"},
		{"s", createRangeF64(0.74, "s"), "fricative"},
		{"ç", createRangeF64(0.84, "ç"), "fricative"},
		{"ß", createRangeF64(0.8510, "ß"), "fricative"},
		{"x", createRangeF64(0.94, "x"), "fricative"},
		{"v", createRangeF64(-0.05, "v"), "especial"},
		{"z", createRangeF64(0.15, "z"), "especial"},
		{"h", createRangeF64(0.25, "h"), "especial"},
		{"w", createRangeF64(0.355, "w"), "especial"},
		{"j", createRangeF64(0.40, "j"), "especial"},
		{"y", createRangeF64(0.45, "y"), "especial"},
		{"q", createRangeF64(0.5050, "q"), "sorda_occlusive"},
	}

	// Números
	numbers := []SymbolEmbedding{
		{"0", createRangeF64(-0.95, "0"), "digit"},
		{"1", createRangeF64(-0.84, "1"), "digit"},
		{"2", createRangeF64(-0.75, "2"), "digit"},
		{"3", createRangeF64(-0.65, "3"), "digit"},
		{"4", createRangeF64(-0.55, "4"), "digit"},
		{"5", createRangeF64(-0.45, "5"), "digit"},
		{"6", createRangeF64(-0.35, "6"), "digit"},
		{"7", createRangeF64(-0.25, "7"), "digit"},
		{"8", createRangeF64(-0.15, "8"), "digit"},
		{"9", createRangeF64(-0.05, "9"), "digit"},
	}

	// Puntuació
	punctuation := []SymbolEmbedding{
		{".", createRangeF64(0.05, "."), "punctuation"},
		{",", createRangeF64(0.10, ","), "punctuation"},
		{";", createRangeF64(0.15, ";"), "punctuation"},
		{":", createRangeF64(0.20, ":"), "punctuation"},
		{"!", createRangeF64(0.25, "!"), "punctuation"},
		{"?", createRangeF64(0.30, "?"), "punctuation"},
		{"-", createRangeF64(0.35, "-"), "punctuation"},
		{"_", createRangeF64(0.40, "_"), "punctuation"},
		{"(", createRangeF64(0.45, "("), "punctuation"},
		{")", createRangeF64(0.50, ")"), "punctuation"},
		{"[", createRangeF64(0.55, "["), "punctuation"},
		{"]", createRangeF64(0.60, "]"), "punctuation"},
		{"{", createRangeF64(0.65, "{"), "punctuation"},
		{"}", createRangeF64(0.70, "}"), "punctuation"},
	}

	// Especials
	specials := []SymbolEmbedding{
		{"@", createRangeF64(0.75, "@"), "special"},
		{"#", createRangeF64(0.80, "#"), "special"},
		{"$", createRangeF64(0.85, "$"), "special"},
		{"%", createRangeF64(0.90, "%"), "special"},
		{"^", createRangeF64(0.95, "^"), "special"},
		{"&", createRangeF64(1.00, "&"), "special"},
		{"~", createRangeF64(-0.20, "~"), "special"},
		{"¬", createRangeF64(-0.10, "¬"), "special"},
	}

	// Concatenem tots els símbols
	return append(append(append(append(append(alphabet, consonants...), numbers...), punctuation...), specials...))
}

func main() {
	embeddings := generateEmbeddings()
	for _, e := range embeddings {
		fmt.Printf("Symbol: %s, Value: %.4f, Group: %s\n", e.Symbol, e.Value.GetF64Value(), e.Group)
	}
}
