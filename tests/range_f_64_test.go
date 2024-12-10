// Proves del tipus RangeF64.
// CreatedAt: 2024/12/08 dg. CLD

package tests

import (
	"math"
	"testing"

	"github.com/jibort/ld_mcac/internal/core"
)

// func TestDebugMasks(t *testing.T) {
// 	inf := core.NewRangeF64Infinite(true)
// 	bits := core.F64ToU64(inf.GetValue())
// 	shiftedBits := (bits << 1) >> 1

// 	t.Logf("Original:  %064b", bits)
// 	t.Logf("Shifted:   %064b", shiftedBits)
// 	t.Logf("GroupMask: %064b", core.GroupMask)
// 	t.Logf("GroupB:    %064b", core.GroupBMask)
// 	t.Logf("Subgroup:  %064b", core.SubgroupMask)
// 	t.Logf("SubInf:    %064b", core.SubgroupInfMask)
// }

func TestDebugGroupBits(t *testing.T) {
	regular := core.NewRangeF64(0.5)
	bits := core.F64ToU64(regular.GetValue())
	t.Logf("Regular Value Bits: %064b", bits)
	t.Logf("GroupMask:         %064b", core.GroupMask)
	t.Logf("GroupA check: %v", (bits&core.GroupMask) == 0)
	t.Logf("GroupB check: %v", (bits&core.GroupMask) == core.GroupBMask)
}
func TestDebugInfinite(t *testing.T) {
	inf := core.NewRangeF64Infinite(true)
	bits := core.F64ToU64(inf.GetValue())
	t.Logf("Infinite Bits: %064b", bits)
	t.Logf("GroupB | SubInf: %064b", core.GroupBMask|core.SubgroupInfMask)
	t.Logf("Group check: %v", (bits&core.GroupMask) == core.GroupBMask)
	t.Logf("Full check: %v", (bits&(core.GroupMask|core.SubgroupMask)) == (core.GroupBMask|core.SubgroupInfMask))
}

func TestRangeF64Debug(t *testing.T) {
	t.Run("Debug Saturat Positiu", func(t *testing.T) {
		val := core.NewRangeF64Saturated(true)
		t.Logf("Saturat Positiu:\n%s", val.DebugBits())
	})

	t.Run("Debug Saturat Negatiu", func(t *testing.T) {
		val := core.NewRangeF64Saturated(false)
		t.Logf("Saturat Negatiu:\n%s", val.DebugBits())
	})

	t.Run("Debug Valor Normal", func(t *testing.T) {
		val := core.NewRangeF64(0.5)
		t.Logf("Valor Normal:\n%s", val.DebugBits())
	})

	t.Run("Debug Infinit Positiu", func(t *testing.T) {
		val := core.NewRangeF64Infinite(true)
		t.Logf("Infinit Positiu:\n%s", val.DebugBits())
	})
}

func TestQuantize64(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  float64
	}{
		{"Zero", 0.0, 0.0},
		{"One", 1.0, 1.0},
		{"MinusOne", -1.0, -1.0},
		{"Half", 0.5, 0.5},
		{"Small", 0.0000001, 0.0},
		{"NearOne", 0.999999999, 1.0},
		{"NearMinusOne", -0.999999999, -1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := core.Quantize64(tt.input)
			if math.Abs(got-tt.want) > core.Epsilon64 {
				t.Errorf("Quantize64(%v) = %v; want %v",
					tt.input, got, tt.want)
			}
		})
	}
}

func TestRangeF64BasicOperations(t *testing.T) {
	// Casos de prova per cada operació
	tests := []struct {
		name    string
		val1    float64
		val2    float64
		wantAdd float64
		wantSub float64
		wantMul float64
		wantDiv float64
		wantErr bool
	}{
		{
			name:    "Valors positius menors que 1",
			val1:    0.3,
			val2:    0.2,
			wantAdd: 0.5,
			wantSub: 0.1,
			wantMul: 0.06,
			wantDiv: 1.5, // Serà saturat a 1.0
			wantErr: false,
		},
		{
			name:    "Suma que satura",
			val1:    0.6,
			val2:    0.7,
			wantAdd: 1.0, // Saturat
			wantSub: -0.1,
			wantMul: 0.42,
			wantDiv: 0.857142857142857, // 0.857142857,
			wantErr: false,
		},
		{
			name:    "Valors negatius",
			val1:    -0.5,
			val2:    -0.3,
			wantAdd: -0.8,
			wantSub: -0.2,
			wantMul: 0.15,
			wantDiv: 1.666666667, // Serà saturat a 1.0
			wantErr: false,
		},
		{
			name:    "Zero com segon operand",
			val1:    0.5,
			val2:    0.0,
			wantAdd: 0.5,
			wantSub: 0.5,
			wantMul: 0.0,
			wantDiv: 0.0,
			wantErr: true, // Error en divisió
		},
		{
			name:    "Valors extrems",
			val1:    1.0,
			val2:    1.0,
			wantAdd: 1.0, // Saturat
			wantSub: 0.0,
			wantMul: 1.0,
			wantDiv: 1.0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Crear els operands
			r1 := core.NewRangeF64(tt.val1)
			r2 := core.NewRangeF64(tt.val2)

			// Provar Add
			resultAdd := r1.Add(r2)
			if math.Abs(resultAdd.ValueF64()-tt.wantAdd) > core.Epsilon64 {
				t.Errorf("Add() = %v, want %v", resultAdd.ValueF64(), tt.wantAdd)
			}

			// Provar Sub
			resultSub := r1.Sub(r2)
			if math.Abs(resultSub.ValueF64()-tt.wantSub) > core.Epsilon64 {
				t.Errorf("Sub() = %v, want %v", resultSub.ValueF64(), tt.wantSub)
			}

			// Provar Mul
			resultMul := r1.Mul(r2)
			if math.Abs(resultMul.ValueF64()-tt.wantMul) > core.Epsilon64 {
				t.Errorf("Mul() = %v, want %v", resultMul.ValueF64(), tt.wantMul)
			}

			// Provar Div
			resultDiv, err := r1.Div(r2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Div() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if math.Abs(resultDiv.ValueF64()-math.Min(1.0, tt.wantDiv)) > core.Epsilon64 {
					t.Errorf("Div() = %v, want %v", resultDiv.ValueF64(), math.Min(1.0, tt.wantDiv))
				}
			}
		})
	}
}

func TestRangeF64Comparisons(t *testing.T) {
	tests := []struct {
		name          string
		val1          float64
		val2          float64
		wantEquals    bool
		wantLess      bool
		wantLessEq    bool
		wantGreater   bool
		wantGreaterEq bool
	}{
		{
			name:          "Valors iguals",
			val1:          0.5,
			val2:          0.5,
			wantEquals:    true,
			wantLess:      false,
			wantLessEq:    true,
			wantGreater:   false,
			wantGreaterEq: true,
		},
		{
			name:          "Primer menor",
			val1:          0.3,
			val2:          0.5,
			wantEquals:    false,
			wantLess:      true,
			wantLessEq:    true,
			wantGreater:   false,
			wantGreaterEq: false,
		},
		{
			name:          "Primer major",
			val1:          0.7,
			val2:          0.5,
			wantEquals:    false,
			wantLess:      false,
			wantLessEq:    false,
			wantGreater:   true,
			wantGreaterEq: true,
		},
		{
			name:          "Valors molt propers",
			val1:          0.5 + core.Epsilon64/2, // TODO: 0.5000001,
			val2:          0.5,
			wantEquals:    true, // Dins Epsilon64
			wantLess:      false,
			wantLessEq:    true,
			wantGreater:   false,
			wantGreaterEq: true,
		},
		{
			name:          "Valors extrems iguals",
			val1:          1.0,
			val2:          1.0,
			wantEquals:    true,
			wantLess:      false,
			wantLessEq:    true,
			wantGreater:   false,
			wantGreaterEq: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r1 := core.NewRangeF64(tt.val1)
			r2 := core.NewRangeF64(tt.val2)

			if got := r1.Equals(r2); got != tt.wantEquals {
				t.Errorf("Equals() = %v, want %v", got, tt.wantEquals)
			}

			if got := r1.LessThan(r2); got != tt.wantLess {
				t.Errorf("LessThan() = %v, want %v", got, tt.wantLess)
			}

			if got := r1.LessOrEqualThan(r2); got != tt.wantLessEq {
				t.Errorf("LessOrEqualThan() = %v, want %v", got, tt.wantLessEq)
			}

			if got := r1.GreaterThan(r2); got != tt.wantGreater {
				t.Errorf("GreaterThan() = %v, want %v", got, tt.wantGreater)
			}

			if got := r1.GreaterOrEqualThan(r2); got != tt.wantGreaterEq {
				t.Errorf("GreaterOrEqualThan() = %v, want %v", got, tt.wantGreaterEq)
			}
		})
	}
}

func TestRangeF64Constructor(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  float64
	}{
		{"Zero", 0.0, 0.0},
		{"Un", 1.0, 1.0},
		{"Menys un", -1.0, -1.0},
		{"Valor positiu", 0.5, 0.5},
		{"Valor negatiu", -0.5, -0.5},
		{"Valor molt petit", 1e-16, 0.0},
		{"Proper a u", 0.999_999_999, 1.0},
		{"Proper a menys u", -0.999_999_999, -1.0},
		{"Major que u", 1.5, 1.0},
		{"Menor que menys u", -1.5, -1.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := core.NewRangeF64(tt.input)
			if math.Abs(got.ValueF64()-tt.want) > core.Epsilon64 {
				t.Errorf("NewRangeF64(%v) = %v, want %v, [%.9f]",
					tt.input, got.ValueF64(), tt.want, math.Abs(got.ValueF64()-tt.want))
			}
		})
	}
}

func TestRangeF64Conversions(t *testing.T) {
	tests := []struct {
		name  string
		input float64
	}{
		{"Zero", 0.0},
		{"Un", 1.0},
		{"Menys un", -1.0},
		{"Mig", 0.5},
		{"Negatiu", -0.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := core.NewRangeF64(tt.input)

			// Test ValueF64
			if got := r.ValueF64(); math.Abs(got-tt.input) > core.Epsilon64 {
				t.Errorf("ValueF64() = %v, want %v", got, tt.input)
			}

			// Test ValueF32
			wantF32 := float32(tt.input)
			if got := r.ValueF32(); math.Abs(float64(got-wantF32)) > float64(core.Epsilon32) {
				t.Errorf("ValueF32() = %v, want %v", got, wantF32)
			}

			// Test conversió i tornada
			if got := r.AsF64(); !got.Equals(r) {
				t.Errorf("AsF64() = %v, want %v", got.ValueF64(), r.ValueF64())
			}
		})
	}
}

func TestRangeF64SpecialValues(t *testing.T) {
	tests := []struct {
		name       string
		value      core.RangeF64
		wantNull   bool
		wantInfPos bool
		wantInfNeg bool
		wantGroupA bool
		wantGroupB bool
		wantGroupC bool
		wantGroupD bool
	}{
		{
			name:       "Null Value",
			value:      core.R64_NULL,
			wantNull:   true,
			wantGroupB: true,
		},
		{
			name:       "Positive Infinity",
			value:      core.R64_INF_POS,
			wantInfPos: true,
			wantGroupB: true,
		},
		{
			name:       "Negative Infinity",
			value:      core.R64_INF_NEG,
			wantInfNeg: true,
			wantGroupB: true,
		},
		{
			name:       "Regular Value",
			value:      core.NewRangeF64(0.5),
			wantGroupA: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.value.IsNullValue(); got != tt.wantNull {
				t.Errorf("IsNullValue() = %v, want %v", got, tt.wantNull)
			}
			if got := tt.value.IsInfinitePos(); got != tt.wantInfPos {
				t.Errorf("IsInfinitePos() = %v, want %v", got, tt.wantInfPos)
			}
			if got := tt.value.IsInfiniteNeg(); got != tt.wantInfNeg {
				t.Errorf("IsInfiniteNeg() = %v, want %v", got, tt.wantInfNeg)
			}
			if got := tt.value.IsGroupA(); got != tt.wantGroupA {
				t.Errorf("IsGroupA() = %v, want %v", got, tt.wantGroupA)
			}
			if got := tt.value.IsGroupB(); got != tt.wantGroupB {
				t.Errorf("IsGroupB() = %v, want %v", got, tt.wantGroupB)
			}
			if got := tt.value.IsGroupC(); got != tt.wantGroupC {
				t.Errorf("IsGroupC() = %v, want %v", got, tt.wantGroupC)
			}
			if got := tt.value.IsGroupD(); got != tt.wantGroupD {
				t.Errorf("IsGroupD() = %v, want %v", got, tt.wantGroupD)
			}
		})
	}
}

func TestRangeF64Saturation(t *testing.T) {
	tests := []struct {
		name       string
		value      core.RangeF64
		wantSatPos bool
		wantSatNeg bool
		wantSat    bool // Nou camp
		wantGroupD bool
	}{
		{
			name:       "Saturació Positiva",
			value:      core.R64_SAT_POS,
			wantSatPos: true,
			wantSatNeg: false,
			wantSat:    true, // Saturat
			wantGroupD: true,
		},
		{
			name:       "Saturació Negativa",
			value:      core.R64_SAT_NEG,
			wantSatPos: false,
			wantSatNeg: true,
			wantSat:    true, // Saturat
			wantGroupD: true,
		},
		{
			name:       "Valor Regular",
			value:      core.NewRangeF64(0.5),
			wantSatPos: false,
			wantSatNeg: false,
			wantSat:    false, // No saturat
			wantGroupD: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.value.IsSaturatedPos(); got != tt.wantSatPos {
				t.Errorf("IsSaturatedPos() = %v, want %v", got, tt.wantSatPos)
			}
			if got := tt.value.IsSaturatedNeg(); got != tt.wantSatNeg {
				t.Errorf("IsSaturatedNeg() = %v, want %v", got, tt.wantSatNeg)
			}
			if got := tt.value.IsSaturated(); got != tt.wantSat {
				t.Errorf("IsSaturated() = %v, want %v", got, tt.wantSat)
			}
			if got := tt.value.IsGroupD(); got != tt.wantGroupD {
				t.Errorf("IsGroupD() = %v, want %v", got, tt.wantGroupD)
			}
		})
	}
}
