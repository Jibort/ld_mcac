// Eines de quantització pel projecte i conversions.
// CreatedAt: 2024/12/08 dg. JIQ

package core

import (
	"encoding/binary"
	"fmt"
	"math"
	"strings"
	"unsafe"
)

// Inicialització del paquet 'core'.
func init() {
	// fmt.Printf("SmallThreshold64 = %v\n", SmallThreshold64)
}

// QUANTITZACIÓ -----------------------
// Assegura que els decimals mantenen la precisió desitjada.
func Quantize64(pVal float64) float64 {
	// Zero exacte
	if pVal == 0.0 {
		return 0.0
	}

	// Valors petits a zero
	absVal := math.Abs(pVal)
	if (absVal - SmallThreshold64) < Epsilon64 {
		return 0.0
	}

	// Valors propers a ±1
	if absVal > 0.9999999 { // Valor més explícit per proximitat a 1
		return math.Copysign(1.0, pVal)
	}

	// Multipliquem, arrodonim i dividim
	scaled := pVal * Factor64
	rounded := math.Round(scaled)
	return rounded / Factor64
}

// Assegura que els decimals mantenen la precisió desitjada.
func Quantize32(pVal float32) float32 {
	val := float64(pVal)
	return float32(math.Round(val*Factor32) / Factor32)
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

func F64ToU64(pFt float64) uint64 {
	return math.Float64bits(pFt)
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
func U64ToF64(pBits uint64) float64 {
	return math.Float64frombits(pBits)
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

// Exponents.
func (sSrc RangeF64) ExtractExponent() uint16 {
	bits := F64ToU64(sSrc.value)
	exponent := (bits >> 52) & 0x7FF // Obtenim els 11 bits de l'exponent
	return uint16(exponent)
}
