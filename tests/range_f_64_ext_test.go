// Tests per a la validació funcional a Range64.
// CreatedAt: 2024/12/10 dt GPT

package tests

import (
	"math"
	"testing"

	"github.com/jibort/ld_mcac/internal/core"
)

func TestRangeF64Constructors(t *testing.T) {
	epsilon := 1e-16 // Epsilon per comparar flotants

	t.Run("NewRangeF64Percentage", func(t *testing.T) {
		pct := 0.1
		r := core.NewRangeF64Percentage(pct)
		expectedFraction := pct
		expectedBits := uint64(0b0100110000000001100110011001100110011001100110011001100110011001)

		// Comprova els bits
		if r.GetU64Value() != expectedBits {
			t.Errorf("NewRangeF64Percentage failed: expected %064b, got %064b", expectedBits, r.GetU64Value())
		}

		// Comprova el valor flotant
		if diff := math.Abs(r.GetF64Value() - expectedFraction); diff > epsilon {
			t.Errorf("NewRangeF64Percentage value failed: expected ~%f, got %f (diff: %e)", expectedFraction, r.GetF64Value(), diff)
		}
	})

	t.Run("NewRangeU64", func(t *testing.T) {
		expectedBits := uint64(0b0100110000000001100110011001100110011001100110011001100110011001)
		r := core.NewRangeF64FromU64(expectedBits)
		expectedFraction := 0.1

		// Comprova els bits
		if r.GetU64Value() != expectedBits {
			t.Errorf("NewRangeU64 failed: expected %064b, got %064b", expectedBits, r.GetU64Value())
		}

		// Comprova el valor flotant
		if diff := math.Abs(r.GetF64Value() - expectedFraction); diff > epsilon {
			t.Errorf("NewRangeU64 value failed: expected ~%f, got %f (diff: %e)", expectedFraction, r.GetF64Value(), diff)
		}
	})
}

func TestRangeF64GroupIdentification(t *testing.T) {
	t.Run("Group Identification", func(t *testing.T) {
		r := core.NewRangeF64Percentage(0.1)
		groupBits := r.GetU64Value() & core.GroupMask
		expectedGroup := core.GroupCMask

		if groupBits != expectedGroup {
			t.Errorf("Group identification failed: expected %064b, got %064b", expectedGroup, groupBits)
		}

		subgroupBits := r.GetU64Value() & core.SubgroupMask
		expectedSubgroup := core.SubGroupC3Mask

		if subgroupBits != expectedSubgroup {
			t.Errorf("Subgroup identification failed: expected %064b, got %064b", expectedSubgroup, subgroupBits)
		}
	})
}

func TestAddRangeF64(t *testing.T) {
	epsilon := 1e-10

	t.Run("Add Percentages", func(t *testing.T) {
		r1 := core.NewRangeF64Percentage(0.3)
		r2 := core.NewRangeF64Percentage(0.2)
		result := AddRangeF64(r1, r2)

		expectedFraction := 0.5
		expectedBits := uint64(0b0100110000000100000000000000000000000000000000000000000000000000)

		// Comprova els bits
		if result.GetU64Value() != expectedBits {
			t.Errorf("AddRangeF64 failed: expected %064b, got %064b", expectedBits, result.GetU64Value())
		}

		// Comprova el valor flotant
		if diff := math.Abs(result.GetF64Value() - expectedFraction); diff > epsilon {
			t.Errorf("AddRangeF64 value failed: expected ~%f, got %f (diff: %e)", expectedFraction, result.GetF64Value(), diff)
		}
	})

	t.Run("Add Exceeds Maximum", func(t *testing.T) {
		r1 := core.NewRangeF64Percentage(0.8)
		r2 := core.NewRangeF64Percentage(0.3)
		result := AddRangeF64(r1, r2)

		expectedFraction := 1.0
		expectedBits := uint64(0b0100110000001000000000000000000000000000000000000000000000000000)

		// Comprova els bits
		if result.GetU64Value() != expectedBits {
			t.Errorf("AddRangeF64 failed (exceeds maximum): expected %064b, got %064b", expectedBits, result.GetU64Value())
		}

		// Comprova el valor flotant
		if diff := math.Abs(result.GetF64Value() - expectedFraction); diff > epsilon {
			t.Errorf("AddRangeF64 value failed (exceeds maximum): expected ~%f, got %f (diff: %e)", expectedFraction, result.GetF64Value(), diff)
		}
	})
}

func AddRangeF64(lhs, rhs core.RangeF64) core.RangeF64 {
	// Obtenim els bits dels valors d'entrada
	lhsBits := lhs.GetU64Value()
	rhsBits := rhs.GetU64Value()

	// Comprovem que els grups són compatibles
	if (lhsBits & core.GroupMask) != (rhsBits & core.GroupMask) {
		panic("Incompatible groups in AddRangeF64")
	}

	// Comprovem que els subgrups són compatibles
	if (lhsBits & core.SubgroupMask) != (rhsBits & core.SubgroupMask) {
		panic("Incompatible subgroups in AddRangeF64")
	}

	// Extreiem els bits de valor de cada operand
	lhsValueBits := lhsBits & core.ValueMask
	rhsValueBits := rhsBits & core.ValueMask

	// Sumem els valors
	sumValueBits := lhsValueBits + rhsValueBits

	// Comprovem si la suma excedeix el màxim permès per ValueMask
	if sumValueBits > core.ValueMask {
		sumValueBits = core.ValueMask // Truncar al màxim
	}

	// Construïm els bits finals per al resultat
	resultBits := (lhsBits & (core.GroupMask | core.SubgroupMask)) | sumValueBits

	// Retornem un nou RangeF64 amb els bits resultants
	return core.NewRangeF64FromU64(resultBits)
}

func TestDebugLargeNumber(t *testing.T) {
	input := uint64(0x40F86A0000000000) // Representació de 1000000.0
	expectedFloat := 1000000.0

	resultFloat := core.U64ToF64(input)
	if math.Abs(resultFloat-expectedFloat) > 1e-10 {
		t.Errorf("U64ToF64 failed: expected %f, got %f", expectedFloat, resultFloat)
	}

	resultBits := core.F64ToU64(resultFloat)
	if resultBits != input {
		t.Errorf("F64ToU64 failed: expected %064b, got %064b", input, resultBits)
	}
}

func TestU64ToF64AndF64ToU64(t *testing.T) {
	tests := []struct {
		name          string
		input         uint64
		expectedBits  uint64
		expectedFloat float64
	}{
		{
			name:          "Zero",
			input:         0x0,
			expectedBits:  0x0,
			expectedFloat: 0.0,
		},
		{
			name:          "Small Fraction",
			input:         0x3FB999999999999A, // float64(0.1)
			expectedBits:  0x3FB999999999999A,
			expectedFloat: 0.1,
		},
		{
			name:          "Half",
			input:         0x3FE0000000000000, // float64(0.5)
			expectedBits:  0x3FE0000000000000,
			expectedFloat: 0.5,
		},
		{
			name:          "One",
			input:         0x3FF0000000000000, // float64(1.0)
			expectedBits:  0x3FF0000000000000,
			expectedFloat: 1.0,
		},
		{
			name:          "Large Number",
			input:         0x40F86A0000000000, // float64(1e6)
			expectedBits:  0x40F86A0000000000,
			expectedFloat: 1e6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert uint64 to float64
			convertedFloat := core.U64ToF64(tt.input)
			if diff := math.Abs(convertedFloat - tt.expectedFloat); diff > 1e-10 {
				t.Errorf("U64ToF64 failed: expected %f, got %f (diff: %e)", tt.expectedFloat, convertedFloat, diff)
			}

			// Convert back to uint64
			convertedBits := core.F64ToU64(convertedFloat)
			if convertedBits != tt.expectedBits {
				t.Errorf("F64ToU64 failed: expected %064b, got %064b", tt.expectedBits, convertedBits)
			}
		})
	}
}
