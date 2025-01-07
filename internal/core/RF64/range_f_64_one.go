// Implementació del tipus F64RangeOne utilitzant constants consolidades
// CreatedAt: 2024-12-27 dc. GPT(JIQ)

package rf64

import (
	"fmt"

	cs "github.com/jibort/ld_mcac/internal/core/consts"
	base "github.com/jibort/ld_mcac/internal/core/intf/base"
	iF64 "github.com/jibort/ld_mcac/internal/core/intf/ranges"
	tools "github.com/jibort/ld_mcac/internal/core/tools"
)

// F64RangeOne representa un rang IEEE754 [-1.0, +1.0].
// És una extensió de F64Range que valida i opera dins del rang definit.
type F64RangeOne struct {
	// if64.F64RangeOneIntf
	F64Range
}

// Verificació que implementa les interfícies requerides.
var _ base.RangeIntf = (*F64RangeOne)(nil)
var _ iF64.F64RangeOneIntf = (*F64RangeOne)(nil)

// -------------------------
// CONSTRUCTORS
// -------------------------

// NewF64RangeOne crea una nova instància validant que el valor es troba dins del rang [-1.0, +1.0].
func NewF64RangeOne(value float64) F64RangeOne {
	if value < cs.Range64Configs.OneF64.Min || value > cs.Range64Configs.OneF64.Max {
		panic(fmt.Sprintf("NewF64RangeOne: valor fora del rang [-1.0, +1.0]: %f", value))
	}
	return F64RangeOne{F64Range{Value: value}}
}

// INTERFÍCIE 'TypeConversionsIntf' ---
func (sR64One F64RangeOne) AsFloat64() float64 {
	return sR64One.Value
}

func (sR64One *F64RangeOne) SetFloat64(pF64 float64) {
	sR64One.Value = pF64
}

func (sR64One F64RangeOne) AsUint64() uint64 {
	panic("F32Range.AsUint64() : not implemented") // TODO: Implement
}

func (sR64One *F64RangeOne) SetUint64(uint64) {
	panic("F32Range.SetUint64() : not implemented") // TODO: Implement
}

func (sR64One F64RangeOne) AsFloat32() float32 {
	return tools.F64ToF32(sR64One.Value)
}

func (sR64One *F64RangeOne) SetFloat32(pF32 float32) {
	sR64One.Value = tools.F32ToF64(pF32)
}

func (sR64One F64RangeOne) AsUint32() uint32 {
	return tools.F32ToU32(tools.F64ToF32(sR64One.Value))
}

func (sR64One *F64RangeOne) SetUint32(pU32 uint32) {
	sR64One.Value = tools.U64ToF64(uint64(pU32))
}

// INTERFÍCIE 'TypeConversionsIntf' ---
func (sR64One F64RangeOne) Is32() bool     { return false }
func (sR64One F64RangeOne) Is64() bool     { return true }
func (sR64One F64RangeOne) IsGroupB() bool { return false }
func (sR64One F64RangeOne) IsGroupC() bool { return false }
func (sR64One F64RangeOne) IsGroupD() bool { return false }

func (s F64RangeOne) AsF64TwoPi() iF64.F64RangeTwoPiIntf {
	panic("F64RangeOne.AsF64TwoPi: Not implemented!")
}

// -------------------------
// IMPLEMENTACIÓ DE X64RangeIntf
// -------------------------
// Conversió a 32bits
func (s F64RangeOne) As32() iF64.X32RangeIntf {
	panic("F64RangeOne.As32: Not implemented!")
}

// -------------------------
// IMPLEMENTACIÓ DE F64RangeIntf
// -------------------------
// Conversió a 32bits
func (s F64RangeOne) AsF32() iF64.F32RangeIntf {
	panic("F64RangeOne.AsF32: Not implemented!")
}

// -------------------------
// IMPLEMENTACIÓ DE CLONABLEINTF
// -------------------------

// Clone retorna una còpia de la instància actual.
func (s F64RangeOne) Clone() base.RangeIntf {
	return &F64RangeOne{F64Range{Value: s.Value}}
}

// -------------------------
// IMPLEMENTACIÓ DE COMPARABLEINTF
// -------------------------

// Equals comprova si dos valors són iguals.
func (s F64RangeOne) Equals(other base.RangeIntf) bool {
	if r, ok := other.(*F64RangeOne); ok {
		return tools.Equals64(s.Value, r.Value, &cs.Epsilon64)
	}
	panic(fmt.Sprintf("F64RangeOne.Equals: tipus no vàlid: %T", other))
}

// LessThan comprova si el valor actual és menor que el valor d'una altra instància.
func (s F64RangeOne) LessThan(other base.RangeIntf) bool {
	if r, ok := other.(*F64RangeOne); ok {
		return s.Value < r.Value
	}
	panic(fmt.Sprintf("F64RangeOne.LessThan: tipus no vàlid: %T", other))
}

// -------------------------
// IMPLEMENTACIÓ DE MATHOPERATIONSINTF
// -------------------------

// Add suma el valor actual amb una altra instància de F64RangeOne.
func (s F64RangeOne) Add(other base.RangeIntf) base.RangeIntf {
	if r, ok := other.(*F64RangeOne); ok {
		result := s.Value + r.Value
		return &F64RangeOne{F64Range{Value: tools.Clamp64(result, cs.Range64Configs.OneF64.Min, cs.Range64Configs.OneF64.Max)}}
	}
	panic(fmt.Sprintf("F64RangeOne.Add: tipus no vàlid: %T", other))
}

// Sub resta dos F64RangeOne i retorna un nou valor (saturat o no) dins del rang.
func (sR64One F64RangeOne) Sub(pOther base.RangeIntf) base.RangeIntf {
	var r *F64RangeOne
	var ok bool
	if r, ok = pOther.(*F64RangeOne); !ok {
		tools.FPanic("F64RangeOne.Sub: Invalid parameter type %T", pOther)
	}

	result := sR64One.Value - r.AsFloat64()
	if result > cs.Range64Configs.OneF64.Max {
		result = cs.Range64Configs.OneF64.Max
	} else if result < cs.Range64Configs.OneF64.Min {
		result = cs.Range64Configs.OneF64.Min
	}
	f64o := NewF64RangeOne(result)
	return &f64o
}

func (sR64One F64RangeOne) Mul(pOther base.RangeIntf) base.RangeIntf {
	var r *F64RangeOne
	var ok bool
	if r, ok = pOther.(*F64RangeOne); !ok {
		tools.FPanic("F64RangeOne.Mul: Invalid parameter type %T", pOther)
	}

	result := sR64One.Value * r.AsFloat64()
	if result > cs.Range64Configs.OneF64.Max {
		result = cs.Range64Configs.OneF64.Max
	} else if result < cs.Range64Configs.OneF64.Min {
		result = cs.Range64Configs.OneF64.Min
	}
	f64o := NewF64RangeOne(result)
	return &f64o
}

func (sR64One F64RangeOne) Div(pOther base.RangeIntf) base.RangeIntf {
	var r *F64RangeOne
	var ok bool
	if r, ok = pOther.(*F64RangeOne); !ok {
		tools.FPanic("F64RangeOne.Div: Invalid parameter type %T", pOther)
	}

	res := NewF64RangeOne(tools.Quantize64One(sR64One.AsFloat64() / r.AsFloat64()))
	return &res
}

// -------------------------
// DOCUMENTACIÓ DEL CODI
// -------------------------

// La implementació assegura que:
// 1. Es valida que els valors estan dins del rang [-1.0, +1.0].
// 2. Es compleixen totes les interfícies requerides.
// 3. Les operacions matemàtiques com `Add` i `Sub` es limiten al rang especificat.

// import (
// 	"fmt"
// 	"math"

// 	cs "github.com/jibort/ld_mcac/internal/core/consts"
// 	base "github.com/jibort/ld_mcac/internal/core/intf/base"
// 	iF64 "github.com/jibort/ld_mcac/internal/core/intf/ranges"
// 	tools "github.com/jibort/ld_mcac/internal/core/tools"
// )

// // Tipus F64RangeOne representa el rang [-1.0, +1.0].
// type F64RangeOne struct {
// 	// f64.F64RangeOneIntf
// 	F64Range
// }

// // Assegura que F64RangeOne implementa RangeIntf
// var _ base.RangeIntf = (*F64RangeOne)(nil)

// // CONSTRUCTORS ------------------------
// func newRangeF64One(pValue float64) F64RangeOne {
// 	f64 := NewF64Range(pValue)
// 	return F64RangeOne{f64}
// }

// // NewF64RangeOne crea una nova instància de F64RangeOne amb validació del rang.
// func NewF64RangeOne(value float64) F64RangeOne {
// 	if math.IsNaN(value) || math.IsInf(value, 0) || (value >= cs.Range64Configs.OneF64.Min && value <= cs.Range64Configs.OneF64.Max) {
// 		return newRangeF64One(value)

// 	}
// 	panic(fmt.Sprintf("NewF64RangeOne: valor fora del rang [-1.0, +1.0]: %f", value))
// }

// // INTERFÍCIE 'ClonableIntf' ----------
// // Clone retorna una còpia de la instància.
// func (sR64One F64RangeOne) Clone() base.RangeIntf {
// 	res := NewF64RangeOne(sR64One.Value)
// 	return &res
// }

// // INTERFÍCIE 'ComparableIntf' --------
// // Equals comprova si dos valors són iguals.
// func (sR64One F64RangeOne) Equals(pOther base.RangeIntf) bool {
// 	if r, ok := pOther.(iF64.F64RangeOneIntf); ok {
// 		return tools.Equals64(sR64One.AsFloat64(), r.AsFloat64(), &cs.Epsilon64)
// 	}
// 	panic(fmt.Sprintf("F64RangeOne.Equals: tipus no vàlid: %T", pOther))
// }

// // Cert només si la instància és menor que pOther.
// func (sR64One F64RangeOne) LessThan(pOther base.RangeIntf) bool {
// 	if r, ok := pOther.(iF64.F64RangeOneIntf); ok {
// 		return sR64One.AsFloat64() < r.AsFloat64()
// 	}
// 	panic(fmt.Sprintf("F64RangeOne.LessThan: tipus no vàlid: %T", pOther))
// }

// // Cert només si la instància és menor o igual que pOther.
// func (sR64One F64RangeOne) LessOrEqualThan(pOther base.RangeIntf) bool {
// 	if r, ok := pOther.(iF64.F64RangeOneIntf); ok {
// 		return sR64One.LessThan(r) || sR64One.Equals(r)
// 	}
// 	panic(fmt.Sprintf("F64RangeOne.LessOrEqualThan: tipus no vàlid: %T", pOther))
// }

// // Cert només si la instància és major que pOther.
// func (sR64One F64RangeOne) GreaterThan(pOther base.RangeIntf) bool {
// 	if r, ok := pOther.(iF64.F64RangeOneIntf); ok {
// 		return sR64One.AsFloat64() > r.AsFloat64()
// 	}
// 	panic(fmt.Sprintf("F64RangeOne.GreaterThan: tipus no vàlid: %T", pOther))
// }

// // Cert només si la instància és major o igual que pOther.
// func (sR64One F64RangeOne) GreaterOrEqualThan(pOther base.RangeIntf) bool {
// 	if r, ok := pOther.(iF64.F64RangeOneIntf); ok {
// 		return sR64One.LessThan(r) || sR64One.Equals(r)
// 	}
// 	panic(fmt.Sprintf("F64RangeOne.GreaterOrEqualThan: tipus no vàlid: %T", pOther))
// }

// // INTERFÍCIE 'TypeConversionsIntf' ---
// func (sR64One F64RangeOne) Is32() bool     { return false }
// func (sR64One F64RangeOne) Is64() bool     { return true }
// func (sR64One F64RangeOne) IsGroupB() bool { return false }
// func (sR64One F64RangeOne) IsGroupC() bool { return false }
// func (sR64One F64RangeOne) IsGroupD() bool { return false }

// // func (sR64One *F64RangeOne) SetFloat64(pF64 float64) {
// // 	sR64One.Value = pF64
// // }

// // INTERFÍCIE 'RangeIntf' -------------
// // IsGroupA verifica si el valor pertany al Grup A.
// func (sR64One F64RangeOne) IsGroupA() bool {
// 	_, exponent, mantissa := tools.DecomposeF64(sR64One.Value)

// 	// Valors normals
// 	if exponent >= 1 && exponent <= 2046 {
// 		return sR64One.Value >= cs.Range64Configs.OneF64.Min && sR64One.Value <= cs.Range64Configs.OneF64.Max
// 	}

// 	// Valors subnormals
// 	if exponent == 0 && mantissa != 0 {
// 		return true
// 	}

// 	// ±Inf i NaN
// 	if exponent == 2047 {
// 		return true
// 	}

// 	return false
// }

// func (sR64One F64RangeOne) AsFloat32() float32 {
// 	return float32(sR64One.Value)
// }

// // func (sR64One F64RangeOne) AsFloat64() float64 {
// // 	return sR64One.Value
// // }

// func (sT64One F64RangeOne) AsUint32() uint32 {
// 	panic("F64RangeOne) AsUint32(): TODO: Encara no implementada!")
// }

// func (sT64One F64RangeOne) AsUint64() uint64 {
// 	return tools.F64ToU64(sT64One.Value)
// }

// func (sT64One *F64RangeOne) SetUint32(value uint32) {
// 	sT64One.SetFloat64(float64(value))
// }

// func (sT64One *F64RangeOne) SetUint64(value uint64) {
// 	sT64One.SetFloat64(tools.U64ToF64(value))
// }

// // INTERFÍCIE 'MathOperationsIntf' ----
// // Add suma dos F64RangeOne i retorna un nou valor saturat dins del rang.
// func (sR64One F64RangeOne) Add(pOther base.RangeIntf) base.RangeIntf {
// 	var r *F64RangeOne
// 	var ok bool
// 	if r, ok = pOther.(*F64RangeOne); !ok {
// 		tools.FPanic("F64RangeOne.Add: Invalid parameter type %T", pOther)
// 	}

// 	result := sR64One.Value + r.AsFloat64()
// 	if result > cs.Range64Configs.OneF64.Max {
// 		result = cs.Range64Configs.OneF64.Max
// 	} else if result < cs.Range64Configs.OneF64.Min {
// 		result = cs.Range64Configs.OneF64.Min
// 	}
// 	f64o := NewF64RangeOne(result)
// 	return &f64o
// }

// // Add suma dos F64RangeOne i retorna un nou valor saturat dins del rang.
// func (sR64One F64RangeOne) Sub(pOther base.RangeIntf) base.RangeIntf {
// 	var r *F64RangeOne
// 	var ok bool
// 	if r, ok = pOther.(*F64RangeOne); !ok {
// 		tools.FPanic("F64RangeOne.Sub: Invalid parameter type %T", pOther)
// 	}

// 	result := sR64One.Value - r.AsFloat64()
// 	if result > cs.Range64Configs.OneF64.Max {
// 		result = cs.Range64Configs.OneF64.Max
// 	} else if result < cs.Range64Configs.OneF64.Min {
// 		result = cs.Range64Configs.OneF64.Min
// 	}
// 	f64o := NewF64RangeOne(result)
// 	return &f64o
// }

// func (sR64One F64RangeOne) Mul(pOther base.RangeIntf) base.RangeIntf {
// 	var r *F64RangeOne
// 	var ok bool
// 	if r, ok = pOther.(*F64RangeOne); !ok {
// 		tools.FPanic("F64RangeOne.Mul: Invalid parameter type %T", pOther)
// 	}

// 	result := sR64One.Value * r.AsFloat64()
// 	if result > cs.Range64Configs.OneF64.Max {
// 		result = cs.Range64Configs.OneF64.Max
// 	} else if result < cs.Range64Configs.OneF64.Min {
// 		result = cs.Range64Configs.OneF64.Min
// 	}
// 	f64o := NewF64RangeOne(result)
// 	return &f64o
// }

// func (sR64One F64RangeOne) Div(pOther base.RangeIntf) base.RangeIntf {
// 	var r *F64RangeOne
// 	var ok bool
// 	if r, ok = pOther.(*F64RangeOne); !ok {
// 		tools.FPanic("F64RangeOne.Div: Invalid parameter type %T", pOther)
// 	}

// 	if tools.IsZero64(r.Value, &cs.Epsilon64) {
// 		tools.FPanic("F64RangeOne.Div: Division by zero")
// 	}

// 	result := sR64One.Value / r.AsFloat64()
// 	if result > cs.Range64Configs.OneF64.Max {
// 		result = cs.Range64Configs.OneF64.Max
// 	} else if result < cs.Range64Configs.OneF64.Min {
// 		result = cs.Range64Configs.OneF64.Min
// 	}
// 	f64o := NewF64RangeOne(result)
// 	return &f64o
// }

// // INTERFÍCIE 'F64RangeOneIntf' -------
// // Cert només si el valor és negatiu.
// func (sR64One F64RangeOne) Sign() bool {
// 	return sR64One.AsFloat64() < 0.0
// }

// // IsInfinitePos comprova si el valor és +Inf.
// func (sR64One F64RangeOne) IsInfinitePos() bool {
// 	return math.IsInf(sR64One.AsFloat64(), 1)
// }

// // IsInfiniteNeg comprova si el valor és -Inf.
// func (sR64One F64RangeOne) IsInfiniteNeg() bool {
// 	return math.IsInf(sR64One.Value, -1)
// }

// // IsInfinite comprova si el valor és infinit (positiu o negatiu).
// func (sR64One F64RangeOne) IsInfinite() bool {
// 	return math.IsInf(sR64One.AsFloat64(), 0)
// }

// // SetF64Value modifica el valor float64 i retorna l'instància actualitzada.
// // func (sR64One *F64RangeOne) SetF64Value(value float64) F64RangeOne {
// // 	sR64One.Value = value
// // 	return *sR64One
// // }
