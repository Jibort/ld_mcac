// Paquet de tests sobre el tipus RangeF64Error.
// CreatedAt: 2024/12/31 dt. GPT(JIQ)

package tests_range_f_64_error

import (
	"math"
	"testing"

	"github.com/jibort/ld_mcac/internal/core"
)

// Test per validar la creació d'un RangeF64Error
func TestNewRangeF64Error(t *testing.T) {
	errorInstance := core.NewRangeF64Error(false, 10, 42)

	if !errorInstance.IsError() {
		t.Errorf("Expected IsError() to return true, got false")
	}

	critic, code, args := errorInstance.Decode()
	if critic {
		t.Errorf("Expected critical flag to be false, got false")
	}
	if code != 10 {
		t.Errorf("Expected error code to be 10, got %d", code)
	}
	if len(args) != 1 || args[0].(uint64) != 42 {
		t.Errorf("Expected arguments to be [42], got %v", args)
	}
}

// Test per validar errors desconeguts
func TestNewRangeF64Error_InvalidCode(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid error code, got none")
		}
	}()

	_ = core.NewRangeF64Error(true, 999, 0) // Codi fora del rang vàlid
}

// Test per validar arguments invàlids
func TestNewRangeF64Error_InvalidArgs(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid arguments, got none")
		}
	}()

	_ = core.NewRangeF64Error(true, 10, math.MaxInt64) // Arguments fora del rang vàlid
}

// Test per validar el comportament de Decode
func TestDecode(t *testing.T) {
	errorInstance := core.NewRangeF64Error(false, 20, 128)

	critic, code, args := errorInstance.Decode()
	if critic {
		t.Errorf("Expected critical flag to be false, got true")
	}
	if code != 20 {
		t.Errorf("Expected error code to be 20, got %d", code)
	}
	if len(args) != 1 || args[0].(uint64) != 128 {
		t.Errorf("Expected arguments to be [128], got %v", args)
	}
}
