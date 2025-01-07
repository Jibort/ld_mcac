package t_rf64_test

import (
	"testing"

	"github.com/jibort/ld_mcac/internal/core/consts"
	"github.com/jibort/ld_mcac/internal/core/rf64"
	"github.com/jibort/ld_mcac/internal/core/tools"
)

func TestNewF64RangeOne(t *testing.T) {
	t.Run("Valor dins del rang", func(t *testing.T) {
		value := 0.5
		r := rf64.NewF64RangeOne(value)
		if r.AsFloat64() != value {
			t.Errorf("Esperat: %f, Obtingut: %f", value, r.AsFloat64())
		}
	})

	t.Run("Valor fora del rang", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Esperada una recuperació per valor fora del rang")
			}
		}()
		_ = rf64.NewF64RangeOne(1.5) // Fora de rang
	})
}

func TestSetFloat64(t *testing.T) {
	r := rf64.NewF64RangeOne(0.0)

	t.Run("Valor vàlid", func(t *testing.T) {
		r.SetFloat64(0.8)
		if r.AsFloat64() != 0.8 {
			t.Errorf("Esperat: %f, Obtingut: %f", 0.8, r.AsFloat64())
		}
	})

	t.Run("Valor fora del rang", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Esperada una recuperació per valor fora del rang")
			}
		}()
		r.SetFloat64(-2.0) // Fora de rang
	})
}

func TestAdd(t *testing.T) {
	r1 := rf64.NewF64RangeOne(0.5)
	r2 := rf64.NewF64RangeOne(0.4)

	t.Run("Suma dins del rang", func(t *testing.T) {
		result := r1.Add(&r2).AsFloat64()
		expected := 0.9
		if !tools.Equals64(result, expected, &consts.Epsilon64) {
			t.Errorf("Esperat: %f, Obtingut: %f", expected, result)
		}
	})

	t.Run("Suma fora del rang", func(t *testing.T) {
		r2 := rf64.NewF64RangeOne(0.7)
		result := r1.Add(&r2).AsFloat64()
		expected := 1.0 // Clamp
		if result != expected {
			t.Errorf("Esperat: %f, Obtingut: %f", expected, result)
		}
	})
}

func TestClone(t *testing.T) {
	r := rf64.NewF64RangeOne(0.3)
	clone := r.Clone()

	t.Run("Còpia exacta", func(t *testing.T) {
		if r.AsFloat64() != clone.AsFloat64() {
			t.Errorf("La còpia no és exacta. Esperat: %f, Obtingut: %f", r.AsFloat64(), clone.AsFloat64())
		}
	})

	t.Run("Independència de la còpia", func(t *testing.T) {
		clone.SetFloat64(0.8)
		if r.AsFloat64() == clone.AsFloat64() {
			t.Errorf("Les instàncies haurien de ser independents")
		}
	})
}

func TestEquals(t *testing.T) {
	r1 := rf64.NewF64RangeOne(0.5)
	r2 := rf64.NewF64RangeOne(0.5)
	r3 := rf64.NewF64RangeOne(0.6)

	t.Run("Valors iguals", func(t *testing.T) {
		if !r1.Equals(&r2) {
			t.Errorf("Els valors haurien de ser iguals")
		}
	})

	t.Run("Valors diferents", func(t *testing.T) {
		if r1.Equals(&r3) {
			t.Errorf("Els valors haurien de ser diferents")
		}
	})
}

func TestLessThan(t *testing.T) {
	r1 := rf64.NewF64RangeOne(0.3)
	r2 := rf64.NewF64RangeOne(0.5)

	t.Run("És menor", func(t *testing.T) {
		if !r1.LessThan(&r2) {
			t.Errorf("%f hauria de ser menor que %f", r1.AsFloat64(), r2.AsFloat64())
		}
	})

	t.Run("No és menor", func(t *testing.T) {
		if r2.LessThan(&r1) {
			t.Errorf("%f no hauria de ser menor que %f", r2.AsFloat64(), r1.AsFloat64())
		}
	})
}
