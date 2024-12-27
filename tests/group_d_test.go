// Tests pel group D de l'estructura RangeF64.
// CreatedAt: 2024/12/16 dl. GPT

package tests

import (
	"testing"

	"github.com/jibort/ld_mcac/internal/core" // Importar RangeF64 des del paquet core
)

func TestNewRangeF64Identifier(t *testing.T) {
	// Crear un RangeF64 amb valors vàlids
	sF64 := core.NewRangeF64Identifier(1, 2, 42) // sequenceType=1, elementType=2, elementID=42

	// Comprovar que el valor creat és un identificador vàlid
	if !sF64.IsIdentifier() {
		t.Errorf("Expected IsIdentifier() to return true, got false")
	}

	// Comprovar que el valor generat té el format correcte
	if sF64.GetSequenceType() != 1 || sF64.GetElementType() != 2 || sF64.GetElementID() != 42 {
		t.Errorf("Identifier values do not match: got sequenceType=%d, elementType=%d, elementID=%d",
			sF64.GetSequenceType(), sF64.GetElementType(), sF64.GetElementID())
	}
}

func TestIsIdentifier(t *testing.T) {
	// Crear un identificador vàlid
	sF64 := core.NewRangeF64Identifier(0, 0, 0) // Identificador vàlid amb valors mínims
	if !sF64.IsIdentifier() {
		t.Errorf("Expected IsIdentifier() to return true for a valid identifier, got false")
	}

	// Crear un valor no identificador (per exemple, un valor normalitzat de Grup A)
	nonIdentifier := core.NewRangeF64(1.0) // Valor que no pertany al Grup D
	if nonIdentifier.IsIdentifier() {
		t.Errorf("Expected IsIdentifier() to return false for a non-identifier, got true")
	}
}

func TestGetElementFields(t *testing.T) {
	// Crear un identificador vàlid
	sF64 := core.NewRangeF64Identifier(2, 3, 12345) // sequenceType=2, elementType=3, elementID=12345

	// Comprovar els camps recuperats
	if sF64.GetSequenceType() != 2 {
		t.Errorf("Expected SequenceType=2, got %d", sF64.GetSequenceType())
	}

	if sF64.GetElementType() != 3 {
		t.Errorf("Expected ElementType=3, got %d", sF64.GetElementType())
	}

	if sF64.GetElementID() != 12345 {
		t.Errorf("Expected ElementID=12345, got %d", sF64.GetElementID())
	}
}
