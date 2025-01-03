// Implementació dels valors Range en float32.
// CreatedAt: 2024/12/08 dg. JIQ

package RF32

import (
	"math"

	cs "github.com/jibort/ld_mcac/internal/core/Consts"
	intf "github.com/jibort/ld_mcac/internal/core/intf"
	tools "github.com/jibort/ld_mcac/internal/core/tools"
)

// Tipus pels Range float64.
type RangeF32 struct {
	intf.Range32Intf
	value float32
}

// CONSTRUCTORS -----------------------
func NewRangeF32(pVal float32) RangeF32 {
	val := tools.Quantize32(pVal)
	if val < -1.0 {
		val = -1.0

	} else if val > +1.0 {
		val = 1.0
	}

	return RangeF32{value: tools.Quantize32(val)}
}

// INTERFÍCIE 'RangeIntf' -------------
// Retorna cert només si la diferència entre els valors és inferior a epsilon32.
func (sSrc RangeF32) Equals(pOther intf.RangeIntf) bool {
	var src, tgt float32

	switch other := pOther.(type) {
	case RangeF32:
		src = tools.Quantize32(float32(math.Abs(float64(sSrc.value))))
		tgt = tools.Quantize32(float32(math.Abs(other.AsFloat64())))
		break
	default:
		return false
	}

	return (src - tgt) < cs.Epsilon32
}

// Retorna cert només en cas que el valor de la instància sigui inferior al de pOther.
func (sSrc RangeF32) LessThan(pOther intf.RangeIntf) bool {
	src := tools.Quantize32(float32(math.Abs(float64(sSrc.value))))
	tgt := tools.Quantize32(float32(math.Abs(pOther.AsFloat64())))
	return src < tgt
}

// Retorna cert només si el valor de la instància és inferior o igual al de pOtheer.
func (sSrc RangeF32) LessOrEqualThan(pOther intf.RangeIntf) bool {
	src := float32(math.Abs(float64(sSrc.value)))
	tgt := float32(math.Abs(pOther.AsFloat64()))
	return (src < tgt) || (src-tgt) < cs.Epsilon32
}
