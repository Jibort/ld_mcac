// Tests de comprovació de la gestió de símbols de l'abecedari del projecte.
// CreatedAt: 2024/12/10 dt. GPT

package tests

import (
	"testing"

	"github.com/jibort/ld_mcac/internal/core"
)

func TestNewSpecialSymbols(t *testing.T) {
	specialSymbols := map[rune]uint64{
		'%': core.GroupCMask | core.SubGroupC3Mask | 0x660,
		'º': core.GroupCMask | core.SubGroupC3Mask | 0x600,
		'ª': core.GroupCMask | core.SubGroupC3Mask | 0x610,
		'·': core.GroupCMask | core.SubGroupC3Mask | 0x620,
		'|': core.GroupCMask | core.SubGroupC3Mask | 0x630,
		'#': core.GroupCMask | core.SubGroupC3Mask | 0x640,
		'$': core.GroupCMask | core.SubGroupC3Mask | 0x650,
		'&': core.GroupCMask | core.SubGroupC3Mask | 0x670,
		'¬': core.GroupCMask | core.SubGroupC3Mask | 0x680,
		'~': core.GroupCMask | core.SubGroupC3Mask | 0x690,
		'€': core.GroupCMask | core.SubGroupC3Mask | 0x6A0,
	}

	for symbol, expected := range specialSymbols {
		// Codificació
		encoded := core.EncodeSymbol(symbol)
		got := core.F64ToU64(encoded.GetF64Value())
		if got != expected {
			t.Errorf("EncodeSymbol failed for symbol: %c, expected: %v, got: %v", symbol, expected, got)
		}

		// Decodificació
		decoded := core.DecodeSymbol(core.NewRangeF64FromU64(expected))
		if decoded != symbol {
			t.Errorf("DecodeSymbol failed for ID: %v, expected symbol: %c, got: %c", expected, symbol, decoded)
		}
	}
}

// 4000000000000000
// 00c0000000000000

func TestEncodeDecodeSymbol(t *testing.T) {
	specialSymbols := map[rune]uint64{
		'%': core.GroupCMask | core.SubGroupC3Mask | 0x660,
		'º': core.GroupCMask | core.SubGroupC3Mask | 0x600,
		'ª': core.GroupCMask | core.SubGroupC3Mask | 0x610,
		'·': core.GroupCMask | core.SubGroupC3Mask | 0x620,
		'|': core.GroupCMask | core.SubGroupC3Mask | 0x630,
		'#': core.GroupCMask | core.SubGroupC3Mask | 0x640,
		'$': core.GroupCMask | core.SubGroupC3Mask | 0x650,
		'&': core.GroupCMask | core.SubGroupC3Mask | 0x670,
		'¬': core.GroupCMask | core.SubGroupC3Mask | 0x680,
		'~': core.GroupCMask | core.SubGroupC3Mask | 0x690,
		'€': core.GroupCMask | core.SubGroupC3Mask | 0x6A0,
	}

	for symbol, expected := range specialSymbols {
		encoded := core.EncodeSymbol(symbol)
		got := core.F64ToU64(encoded.GetF64Value())
		t.Logf("Symbol: %c, GroupCMask: %x, SubGroupC3Mask: %x, ID: %x, Got: %x",
			symbol, core.GroupCMask, core.SubGroupC3Mask, core.SymbolToID[symbol], got)

		if got != expected {
			t.Errorf("EncodeSymbol failed for symbol: %c, expected: %v, got: %v", symbol, expected, got)
		}

		decoded := core.DecodeSymbol(core.NewRangeF64FromU64(expected))
		if decoded != symbol {
			t.Errorf("DecodeSymbol failed for ID: %v, expected symbol: %c, got: %c", expected, symbol, decoded)
		}
	}
}
