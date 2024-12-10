// Implementació dels valors Range en float32.
// CreatedAt: 2024/12/08 dg. JIQ

package core

import "math"

// Tipus pels Range float64.
type RangeF32 struct {
	Range32Intf
	value float32
}

// CONSTRUCTORS -----------------------
func NewRangeF32(pVal float32) RangeF32 {
	val := Quantize32(pVal)
	if val < -1.0 {
		val = -1.0

	} else if val > +1.0 {
		val = 1.0
	}

	return RangeF32{value: Quantize32(val)}
}

// INTERFÍCIE 'RangeIntf' -------------
// Retorna cert només si la diferència entre els valors és inferior a epsilon32.
func (sSrc RangeF32) Equals(pOther RangeIntf) bool {
	src := Quantize32(float32(math.Abs(sSrc.ValueF64())))
	tgt := Quantize32(float32(math.Abs(pOther.ValueF64())))
	return (src - tgt) < Epsilon32
}

// Retorna cert només en cas que el valor de la instància sigui inferior al de pOther.
func (sSrc RangeF32) LessThan(pOther RangeIntf) bool {
	src := Quantize32(float32(math.Abs(sSrc.ValueF64())))
	tgt := Quantize32(float32(math.Abs(pOther.ValueF64())))
	return src < tgt
}

// Retorna cert només si el valor de la instància és inferior o igual al de pOtheer.
func (sSrc RangeF32) LessOrEqualThan(pOther RangeIntf) bool {
	src := math.Abs(sSrc.ValueF64())
	tgt := math.Abs(pOther.ValueF64())
	return (src < tgt) || (src-tgt) < Epsilon64
}
