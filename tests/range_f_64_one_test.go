// Fitxer de tests per RangeF64One
// CreatedAt: 2024/12/28 ds. GPT(JIQ)

package tests

import (
	"math"
	"testing"

	rF64One "github.com/jibort/ld_mcac/internal/core/rf64"
)

func TestNewRangeF64One(t *testing.T) {
	tests := []struct {
		input       float64
		expectError bool
	}{
		{-1.0, false},       // Valor mínim vàlid
		{1.0, false},        // Valor màxim vàlid
		{0.5, false},        // Valor dins del rang
		{-1.5, true},        // Fora del rang mínim
		{1.5, true},         // Fora del rang màxim
		{math.Inf(1), true}, // Infinit positiu
		{math.NaN(), true},  // NaN
	}

	for _, tt := range tests {
		rOne := rF64One.NewF64RangeOne(tt.input)
		if tt.expectError {
			if !rOne.IsError() {
				t.Errorf("NewRangeF64One(%v) expected an error but got none", tt.input)
			}
		} else {
			if rOne.IsError() {
				t.Errorf("NewRangeF64One(%v) returned an unexpected error: %v", tt.input, rOne)
			}
			if !rOne.IsGroupA() {
				t.Errorf("NewRangeF64One(%v) expected to be in Group A but was not", tt.input)
			}
		}
	}
}
