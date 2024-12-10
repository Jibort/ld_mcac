// Tests per a la identificació de grups a Range64.
// CreatedAt: 2024/12/10 dt GPT

package tests

import (
	"math"
	"testing"

	"github.com/jibort/ld_mcac/internal/core"
)

// [GPT] Test de -1.1
func TestIsGroupAWithNegativeValues(t *testing.T) {
	negativeValue := core.NewRangeF64(-1.1) // Exemple de valor negatiu fora de Grup A

	if negativeValue.IsGroupA() {
		t.Errorf("Valor negatiu (-1.1) incorrectament identificat com a Grup A")
	}
}

// [GPT] Tests de casos especials.
func TestSpecialCases(t *testing.T) {
	zero := core.NewRangeF64(0.0)
	inf := core.NewRangeF64(math.Inf(1))
	negInf := core.NewRangeF64(math.Inf(-1))
	nan := core.NewRangeF64(math.NaN())

	if !zero.IsGroupA() {
		t.Errorf("Zero (0.0) no identificat correctament com a Grup A")
	}
	if inf.IsGroupA() {
		t.Errorf("+Inf identificat incorrectament com a Grup A")
	}
	if negInf.IsGroupA() {
		t.Errorf("-Inf identificat incorrectament com a Grup A")
	}
	if nan.IsGroupA() {
		t.Errorf("NaN identificat incorrectament com a Grup A")
	}
}

// [GPT] TestIsGroupA valida la detecció del Grup A
func TestIsGroupA(t *testing.T) {
	testCases := []struct {
		name     string
		rangeF64 core.RangeF64
		expected bool
	}{
		{"F64 +zero", core.NewRangeF64(0.0), true},                   // Zero únic
		{"F64 GroupAMask", core.NewRangeU64(core.GroupAMask), true},  // Grup A vàlid
		{"F64 GroupBMask", core.NewRangeU64(core.GroupBMask), false}, // No Grup A
		{"F64 GroupCMask", core.NewRangeU64(core.GroupCMask), false}, // No Grup A
		{"F64 GroupDMask", core.NewRangeU64(core.GroupDMask), false}, // No Grup A
		{"F64 GroupEMask", core.NewRangeU64(core.GroupEMask), false}, // No Grup A
	}

	for _, tc := range testCases {
		result := tc.rangeF64.IsGroupA()
		if result != tc.expected {
			t.Errorf("Failed for %s > expected: %v, got: %v", tc.name, tc.expected, result)
		}
	}
}

// // TestIsGroupB valida la detecció del Grup B
// func TestIsGroupB(t *testing.T) {
// 	testCases := []struct {
// 		value    uint64
// 		expected bool
// 	}{
// 		{core.GroupAMask, false},
// 		{core.GroupBMask, true},
// 		{core.GroupCMask, false},
// 		{core.GroupDMask, false},
// 		{core.GroupEMask, false},
// 	}

// 	for _, tc := range testCases {
// 		result := core.IsGroupB(tc.value)
// 		if result != tc.expected {
// 			t.Errorf("IsGroupB failed for value: %064b, expected: %v, got: %v", tc.value, tc.expected, result)
// 		}
// 	}
// }

// // TestIsGroupC valida la detecció del Grup C
// func TestIsGroupC(t *testing.T) {
// 	testCases := []struct {
// 		value    uint64
// 		expected bool
// 	}{
// 		{core.GroupAMask, false},
// 		{core.GroupBMask, false},
// 		{core.GroupCMask, true},
// 		{core.GroupDMask, false},
// 		{core.GroupEMask, false},
// 	}

// 	for _, tc := range testCases {
// 		result := core.IsGroupC(tc.value)
// 		if result != tc.expected {
// 			t.Errorf("IsGroupC failed for value: %064b, expected: %v, got: %v", tc.value, tc.expected, result)
// 		}
// 	}
// }

// // TestIsGroupD valida la detecció del Grup D
// func TestIsGroupD(t *testing.T) {
// 	testCases := []struct {
// 		value    uint64
// 		expected bool
// 	}{
// 		{core.GroupAMask, false},
// 		{core.GroupBMask, false},
// 		{core.GroupCMask, false},
// 		{core.GroupDMask, true},
// 		{core.GroupEMask, false},
// 	}

// 	for _, tc := range testCases {
// 		result := core.IsGroupD(tc.value)
// 		if result != tc.expected {
// 			t.Errorf("IsGroupD failed for value: %064b, expected: %v, got: %v", tc.value, tc.expected, result)
// 		}
// 	}
// }

// // TestIsGroupE valida la detecció del Grup E
// func TestIsGroupE(t *testing.T) {
// 	testCases := []struct {
// 		value    uint64
// 		expected bool
// 	}{
// 		{core.GroupAMask, false},
// 		{core.GroupBMask, false},
// 		{core.GroupCMask, false},
// 		{core.GroupDMask, false},
// 		{core.GroupEMask, true},
// 	}

// 	for _, tc := range testCases {
// 		result := core.IsGroupE(tc.value)
// 		if result != tc.expected {
// 			t.Errorf("IsGroupE failed for value: %064b, expected: %v, got: %v", tc.value, tc.expected, result)
// 		}
// 	}
// }

// // TestRandomValues valida que els valors aleatoris no coincideixen amb cap grup incorrectament
// func TestRandomValues(t *testing.T) {
// 	randomValues := []uint64{
// 		0b00001111_11111111_00000000_11111111_00000000_11111111_00000000_11111111, // Aleatori
// 		0b11110000_00001111_11111111_11110000_00001111_11110000_11111111_00001111, // Aleatori
// 		0b00111111_11111111_11111111_11111111_11111111_11111111_11111111_11111111, // Aleatori
// 	}

// 	for _, val := range randomValues {
// 		if core.IsGroupA(val) || core.IsGroupB(val) || core.IsGroupC(val) || core.IsGroupD(val) || core.IsGroupE(val) {
// 			t.Errorf("Random value incorrectly identified as a group: %064b", val)
// 		}
// 	}
// }
