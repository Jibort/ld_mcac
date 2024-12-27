// Interfícies generals dels elements individuals de càlcul pel projecte.
// CreatedAt: 2024/12/08 dg. JIQ

package core

import (
	"math"
)

const (
	Precision32 = 6
	Precision64 = 16
	// SmallThreshold64 = 1e-8 // Nou llindar per valors petits
)

var (
	Epsilon32 = float32(math.Pow(10, -Precision32))
	Factor32  = math.Pow(10, Precision32)
	// Epsilon64 = math.Pow(10, -Precision64)
	// Factor64  = math.Pow(10, Precision64)
)

const (
	Epsilon64        = 1e-13                  // Redueix el llindar per detectar errors petits
	SmallThreshold64 = 1e-8                   // Augmenta la sensibilitat per valors petits
	Factor64         = 10_000_000_000_000_000 // Augmenta la precisió en l'arrodoniment
)

// Interfície global
type RangeIntf interface {
	// Comparacions
	Equals(pOther RangeIntf) bool
	LessThan(pOther RangeIntf) bool
	LessOrEqualThan(pOther RangeIntf) bool
	GreaterThan(pOther RangeIntf) bool
	GreaterOrEqualThan(pOther RangeIntf) bool

	// Valors especials
	IsInfinitePos() bool
	IsInfiniteNeg() bool
	IsNullValue() bool
	IsPadding() bool
	IsPaddingStart() bool
	IsPaddingEnd() bool
	IsNull() bool

	// Grups
	IsGroupA() bool
	IsGroupB() bool
	IsGroupC() bool
	IsGroupD() bool

	// Operacions
	Add(pOther RangeIntf) RangeIntf
	Sub(pOther RangeIntf) RangeIntf
	Mul(pOther RangeIntf) RangeIntf
	Div(pOther RangeIntf) (RangeIntf, error)

	// Saturacions
	IsSaturated() bool
	IsSaturatedPos() bool
	IsSaturatedNeg() bool

	// Gestió d'errors
	IsError() bool           // Retorna si el valor representa un error
	ErrorCode() int          // Retorna el codi d'error si n'hi ha
	ErroneousValue() float64 // Retorna el valor associat a l'error (si escau)

	// Conversió entre tipus
	AsF64() RangeF64   // Conversió a RangeF64
	As32() Range32Intf // Conversió a Range32Intf

	// Getters/Setters
	GetF64Value() float64               // Conversió a float64
	SetF64Value(pVal float64) RangeIntf // Modificació del valor de la instància.
	GetU64Value() uint64                // Conversió a uint64
	SetU64Value(pVal uint64) RangeIntf  // Modificació del valor de la instància.
	GetPercentage() (float64, bool)
	SetPercentage(pVal float64) RangeIntf // Modificació del percentatge de la instància.

	// Helpers opcionals
	ValueF64() float64 // Retorna el valor com a float64
	ValueF32() float32 // Retorna el valor com a float32

	// Símbols
	IsSymbol() bool // Retorna cert només si la instància defineix un símbol
	ToSymbol() rune // Retorna el símbol RangeIntf si és un símbol (0 en cas contrari)

	IsIdentifier() bool
	IsSubnormal() bool
	ExtractMantissa() uint64
	ExtractExponent() int

	// Altres segons necessitats
}

// Interfície per a valors fixed o floating points de 64bits.
type Range64Intf interface {
	RangeIntf
	As32() Range32Intf // conversió a 32 bits
}

// Interfície per a valors fixed o floating points de 32bits.
type Range32Intf interface {
	RangeIntf
	As64() Range64Intf // conversió a 64 bits
}

type NetworkConfig struct {
	Layers          int    `json:"layers"`            // Nombre de capes
	NeuronsPerLayer []int  `json:"neurons_per_layer"` // Nombre de neurones per cada capa
	Activation      string `json:"activation"`        // Funció d'activació predeterminada
}
