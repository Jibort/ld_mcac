// Proves del tipus RangeF64.
// CreatedAt:

package tests

import (
	"math"
	"testing"

	"github.com/jibort/ld_mcac/internal/core"
)

func TestNewRangeF64(t *testing.T) {
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
		r := core.NewRangeF64(tt.input)
		if tt.expectError {
			if !r.IsError() {
				t.Errorf("NewRangeF64(%v) expected an error but got none", tt.input)
			}
		} else {
			if r.IsError() {
				t.Errorf("NewRangeF64(%v) returned an unexpected error: %v", tt.input, r.ErrorCode())
			}
			if !r.IsGroupA() {
				t.Errorf("NewRangeF64(%v) expected to be in Group A but was not", tt.input)
			}
		}
	}
}
