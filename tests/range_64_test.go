// Fitxer de tests per Range64
// CreatedAt: 2024/12/28 ds. GPT(JIQ)

package tests

import (
	"math"
	"testing"

	rF64 "github.com/jibort/ld_mcac/internal/core/RF64"
)

func TestNewRange64(t *testing.T) {
	tests := []struct {
		input       float64
		expected    float64
		expectError bool
	}{
		{-1.5, -1.0, false},   // Fora del rang mínim
		{1.5, 1.0, false},     // Fora del rang màxim
		{0.5, 0.5, false},     // Dins del rang
		{-0.5, -0.5, false},   // Dins del rang
		{math.NaN(), 0, true}, // NaN
	}

	for _, tt := range tests {
		rf64 := rF64.NewRangeF64(tt.input)
		if rf64.AsFloat64() != tt.expected {
			t.Errorf("NewRange64(%v) = %v; want %v", tt.input, rf64.AsFloat64(), tt.expected)
		}
	}
}
