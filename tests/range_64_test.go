// Joc de proves per validar el tipus Range64.
// CreatedAt: 2024/12/08 dg CLD

package tests

import (
	"math"
	"testing"

	"github.com/jibort/ld_mcac/internal/core"
)

// UTILITATS DE TEST ----------------------
// equals compara dos float64 amb una precisió donada
func equals(pA, pB float64, pPrecision float64) (rEquals bool) {
	return math.Abs(pA-pB) < pPrecision
}

// TEST DE CONSTRUCTORS -------------------
func TestNewRange64Basics(t *testing.T) {
	sTests := []struct {
		sName      string
		sInput     float64
		sExpected  float64
		sPrecision float64
	}{
		{"Zero", 0.0, 0.0, 1e-15},
		{"One", 1.0, 1.0, 1e-15},
		{"MinusOne", -1.0, -1.0, 1e-15},
		{"Half", 0.5, 0.5, 1e-15},
		{"MinusHalf", -0.5, -0.5, 1e-15},
	}

	for _, sTest := range sTests {
		t.Run(sTest.sName, func(t *testing.T) {
			rResult := core.NewRange64(sTest.sInput).ToFloat64()
			if !equals(rResult, sTest.sExpected, core.Epsilon64) {
				t.Errorf("NewRange64(%v) = %v; want %v",
					sTest.sInput, rResult, sTest.sExpected)
			}
		})
	}
}

func TestNewRange64Limits(t *testing.T) {
	sTests := []struct {
		sName      string
		sInput     float64
		sExpected  float64
		sPrecision float64
	}{
		{"OverMax", 1.5, 1.0, 1e-15},
		{"UnderMin", -1.5, -1.0, 1e-15},
		{"VeryLarge", 1e6, 1.0, 1e-15},
		{"VerySmall", -1e6, -1.0, 1e-15},
		{"Infinity", math.Inf(1), 1.0, 1e-15},
		{"MinusInfinity", math.Inf(-1), -1.0, 1e-15},
		{"NaN", math.NaN(), 0.0, 1e-15},
	}

	for _, sTest := range sTests {
		t.Run(sTest.sName, func(t *testing.T) {
			rResult := core.NewRange64(sTest.sInput).ToFloat64()
			if !equals(rResult, sTest.sExpected, sTest.sPrecision) {
				t.Errorf("NewRange64(%v) = %v; want %v",
					sTest.sInput, rResult, sTest.sExpected)
			}
		})
	}
}

// TEST DE PRECISIÓ ----------------------

func TestRange64Precision(t *testing.T) {
	// Testegem valors específics on sabem que float64 té precisió exacta
	sTests := []struct {
		sName      string
		sInput     float64
		sPrecision float64
	}{
		{"OneHalf", 0.5, 1e-15},
		{"OneQuarter", 0.25, 1e-15},
		{"OneEighth", 0.125, 1e-15},
		{"OneSixteenth", 0.0625, 1e-15},
		{"OneThirtySecond", 0.03125, 1e-15},
		// Potències negatives de 2 fins a 2^-15
		{"MinusHalf", -0.5, 1e-15},
		{"MinusQuarter", -0.25, 1e-15},
		{"MinusEighth", -0.125, 1e-15},
		{"MinusSixteenth", -0.0625, 1e-15},
		{"MinusThirtySecond", -0.03125, 1e-15},
	}

	for _, sTest := range sTests {
		t.Run(sTest.sName, func(t *testing.T) {
			sRange := core.NewRange64(sTest.sInput)
			rResult := sRange.ToFloat64()
			if !equals(rResult, sTest.sInput, sTest.sPrecision) {
				t.Errorf("Pèrdua de precisió: input=%v, output=%v, diff=%v",
					sTest.sInput, rResult, math.Abs(sTest.sInput-rResult))
			}
		})
	}
}

// TEST DE METADADES ---------------------

func TestRange64Metadata(t *testing.T) {
	sTests := []struct {
		sName     string
		sValue    float64
		sMeta     uint64
		sExpected uint64
	}{
		{"MetaZero", 0.0, 0, 0},
		{"MetaAllOnes", 0.0, 0x7FFF, 0x7FFF},
		{"MetaAlternatingOnes", 0.0, 0x5555, 0x5555},
		{"MetaAlternatingZeros", 0.0, 0x2AAA, 0x2AAA},
	}

	for _, sTest := range sTests {
		t.Run(sTest.sName, func(t *testing.T) {
			sRange := core.NewRange64(sTest.sValue).SetMeta(sTest.sMeta)
			rResult := sRange.GetMeta()
			if rResult != sTest.sExpected {
				t.Errorf("Meta %v: got %014b, want %014b",
					sTest.sName, rResult, sTest.sExpected)
			}
		})
	}
}

// TEST D'OPERACIONS BÀSIQUES ------------

func TestRange64BasicOps(t *testing.T) {
	sTests := []struct {
		sName      string
		sA         float64
		sB         float64
		sSum       float64
		sSub       float64
		sMul       float64
		sDiv       float64
		sPrecision float64
	}{
		{
			"PositiveValues",
			0.5, 0.25,
			0.75, 0.25, 0.125, 2.0,
			1e-15,
		},
		{
			"NegativeValues",
			-0.5, -0.25,
			-0.75, -0.25, 0.125, 2.0,
			1e-15,
		},
		{
			"MixedValues",
			0.5, -0.25,
			0.25, 0.75, -0.125, -2.0,
			1e-15,
		},
	}

	for _, sTest := range sTests {
		t.Run(sTest.sName, func(t *testing.T) {
			sA := core.NewRange64(sTest.sA)
			sB := core.NewRange64(sTest.sB)

			rSum := sA.Add(sB).ToFloat64()
			if !equals(rSum, sTest.sSum, sTest.sPrecision) {
				t.Errorf("Add: %v + %v = %v; want %v",
					sTest.sA, sTest.sB, rSum, sTest.sSum)
			}

			rSub := sA.Sub(sB).ToFloat64()
			if !equals(rSub, sTest.sSub, sTest.sPrecision) {
				t.Errorf("Sub: %v - %v = %v; want %v",
					sTest.sA, sTest.sB, rSub, sTest.sSub)
			}

			rMul := sA.Mul(sB).ToFloat64()
			if !equals(rMul, sTest.sMul, sTest.sPrecision) {
				t.Errorf("Mul: %v * %v = %v; want %v",
					sTest.sA, sTest.sB, rMul, sTest.sMul)
			}

			rDiv := sA.Div(sB).ToFloat64()
			if !equals(rDiv, sTest.sDiv, sTest.sPrecision) {
				t.Errorf("Div: %v / %v = %v; want %v",
					sTest.sA, sTest.sB, rDiv, sTest.sDiv)
			}
		})
	}
}

// TEST DE DIVISIÓ PER ZERO -------------

func TestRange64DivByZero(t *testing.T) {
	sTests := []struct {
		sName      string
		sNum       float64
		sExpected  float64
		sPrecision float64
	}{
		{"PositiveByZero", 0.5, 1.0, 1e-15},
		{"NegativeByZero", -0.5, -1.0, 1e-15},
		{"ZeroByZero", 0.0, 1.0, 1e-15},
	}

	for _, sTest := range sTests {
		t.Run(sTest.sName, func(t *testing.T) {
			sA := core.NewRange64(sTest.sNum)
			sB := core.NewRange64(0.0)
			rResult := sA.Div(sB).ToFloat64()
			if !equals(rResult, sTest.sExpected, sTest.sPrecision) {
				t.Errorf("%v/0 = %v; want %v",
					sTest.sNum, rResult, sTest.sExpected)
			}
		})
	}
}
