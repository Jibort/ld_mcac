// Eines de quantització pel projecte i conversions.
// CreatedAt: 2024/12/08 dg. JIQ

package tools

import (
	"encoding/binary"
	"fmt"
	"math"
	"strings"
	"unsafe"

	cs "github.com/jibort/ld_mcac/internal/core/Consts"
)

// QUANTITZACIÓ -----------------------
// Assegura que els decimals mantenen la precisió desitjada.
func Quantize64(pVal float64) float64 {
	// Zero exacte
	if pVal == 0.0 {
		return 0.0
	}

	// Valors petits a zero
	absVal := math.Abs(pVal)
	if (absVal - cs.SmallThreshold64) < cs.Epsilon64 {
		return 0.0
	}

	// Valors propers a ±1
	if absVal > 0.9999999 { // Valor més explícit per proximitat a 1
		return math.Copysign(1.0, pVal)
	}

	// Multipliquem, arrodonim i dividim
	scaled := pVal * cs.Factor64
	rounded := math.Round(scaled)
	return rounded / cs.Factor64
}

// Assegura que els decimals mantenen la precisió desitjada.
func Quantize32(pVal float32) float32 {
	val := float64(pVal)
	return float32(math.Round(val*cs.Factor32) / cs.Factor32)
}

// CONVERSIONS ------------------------
// F64 a ..............................
func F64ToF64(pFt float64) float64 {
	return pFt
}

func F64ToF32(pFt float64) float32 {
	return Quantize32(float32(pFt))
}

func F64ToI64(pFt float64) int64 {
	bits := math.Float64bits(pFt)
	return *(*int64)(unsafe.Pointer(&bits))
}

func F64ToI32(pFt float64) int32 {
	bits := math.Float32bits(Quantize32(float32(pFt)))
	return *(*int32)(unsafe.Pointer(&bits))
}

func F64ToU64(pF64 float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&pF64))
}

func F64ToU32(pFt float64) uint32 {
	return math.Float32bits(Quantize32(float32(pFt)))
}

func F64ToB64(pFt float64) string {
	u64 := F64ToU64(pFt)
	// Crear un slice de bytes de mida 8
	bytes := make([]byte, 8)
	// Convertir el valor uint64 a bytes en ordre BigEndian
	binary.BigEndian.PutUint64(bytes, u64)

	// Utilitzar strings.Builder per construir la cadena binària
	var builder strings.Builder
	for _, b := range bytes {
		// Afegir la representació binària de cada byte al builder
		fmt.Fprintf(&builder, "%08b ", b)
	}

	// Retornar la cadena resultant, eliminant l'espai final
	return strings.TrimSpace(builder.String())
}

// U64 a ..............................
// Conversió des de uint64 a float64
func U64ToF64(pU64 uint64) float64 {
	if (pU64>>63) == 1 && (pU64&0x7FFFFFFFFFFFFFFF) == 0x7FFFFFFFFFFFFFFF {
		panic(fmt.Sprintf("Invalid IEEE 754 representation: %064b", pU64))
	}
	return math.Float64frombits(pU64)
}

func U64ToB64(pVal uint64) string {
	// Crear un slice de bytes de mida 8
	bytes := make([]byte, 8)
	// Convertir el valor uint64 a bytes en ordre BigEndian
	binary.BigEndian.PutUint64(bytes, pVal)

	// Utilitzar strings.Builder per construir la cadena binària
	var builder strings.Builder
	for _, b := range bytes {
		// Afegir la representació binària de cada byte al builder
		fmt.Fprintf(&builder, "%08b ", b)
	}

	// Retornar la cadena resultant, eliminant l'espai final
	return strings.TrimSpace(builder.String())
}

func ValidateIEEE754(pU64 uint64) bool {
	// Comprova si els bits representen un valor vàlid en format IEEE 754
	exponent := (pU64 >> 52) & 0x7FF
	mantissa := pU64 & 0xFFFFFFFFFFFFF

	// El valor ha de tenir un exponent vàlid (diferent de 0x7FF que indicaria NaN o infinit)
	if exponent == 0x7FF && mantissa != 0 {
		return false // És un NaN
	}
	return true
}

func DecomposeF64(value float64) (sign int, exponent int, mantissa uint64) {
	// Converteix el float64 a la seva representació en uint64
	bits := F64ToU64(value)

	// Bit de signe: el bit 63
	sign = int((bits >> 63) & 1)

	// Exponent: els bits 62-52
	exponent = int((bits >> 52) & 0x7FF)

	// Mantissa: els bits 51-0
	mantissa = bits & ((1 << 52) - 1)

	return
}

// COMPARACIONS AMB 'float' -----------
// Retorna cert només si la diferència entre pA i pB és inferior a pEpsilon.
func Equals64(pA, pB float64, pEpsilon *float64) (rEquals bool) {
	if pEpsilon == nil {
		pEpsilon = &cs.Epsilon64
	}
	return math.Abs(pA-pB) < *pEpsilon
}

// MÀSCARES ---------------------------
func ApplyMask(pValue uint64, pMask uint64) (rU64 uint64) {
	return (pValue | pMask) & pMask
}