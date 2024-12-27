// Implementació dels valors Range en float64.
// CreatedAt: 2024/12/08 dg. JIQ

package core

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

// Tipus pels Range float64.
type RangeF64 struct {
	Range64Intf
	value float64
}

// CONSTRUCTORS -----------------------
// Constructor general a partir d'un valor float64 (64bits).
func NewRangeF64(pF64 float64) RangeF64 {
	return RangeF64{value: pF64}
}

// Constructor a partir d'un uint64 (64bits).
func NewRangeF64FromU64(pU64 uint64) RangeF64 {
	// fmt.Printf("NewRangeU64: Bits rebuts: %064b\n", pU64)
	return RangeF64{value: U64ToF64(pU64)}
}

// Constructor de RangeF64 per a 0.0.
func NewRangeF64Zero() RangeF64 {
	return NewRangeF64FromU64(IEEE754ZeroBits)
}

func NewRangeF64Null() RangeF64 {
	bits := GroupBMask | SubgroupNullMask
	return NewRangeF64FromU64(bits)
}

// Crea un valor RangeF64 per a errors, codificant el codi d'error i el valor erroni.
func NewRangeF64Error(pCode int, pErroneousValue float64) RangeF64 {
	groupBits := GroupBMask                                    // Grup B
	subgroupBits := SubgroupNullMask                           // Subgrup genèric
	errorCodeBits := uint64(pCode) << 48                       // Desplaçar el codi d'error a la seva posició
	valueBits := math.Float64bits(pErroneousValue) & ValueMask // Valor associat a l'error

	// Combinar els bits
	finalBits := groupBits | subgroupBits | errorCodeBits | valueBits

	return NewRangeF64FromU64(finalBits)
}

// Crea valors saturats
func NewRangeF64Saturated(value float64, isNull bool) RangeF64 {
	var u64 uint64

	// Bit de signe i saturació (S1)
	if value < 0 {
		u64 |= SaturationMask | SignBitMask
	} else {
		u64 |= SaturationMask
	}

	// Valors especials (nul o +1.0* / -1.0*)
	if isNull {
		u64 |= NullFlagMask // Flag 'x' per a valors nuls
	} else if value == 1.0 || value == -1.0 {
		u64 |= UnitFlagMask // Flag 'u' per a +1.0* o -1.0*
	}

	// Subnormalitzat
	if IsFloat64Subnormal(value) {
		u64 |= SubnormalFlagMask // Flag 'n' per a subnormalitzats
	}

	// Mantisa (52 bits)
	mantissa := ExtractFloat64tMantissa(value)
	u64 |= mantissa

	return NewRangeF64FromU64(u64)
}

func NewRangeF64Saturated_(pVal float64) RangeF64 {
	// Converteix el valor proporcionat a bits IEEE 754
	bits := F64ToU64(pVal)

	// Força que el valor estigui dins del grup saturat (Grup D)
	bits = (bits &^ GroupMask) | GroupDMask

	// Limita el valor a ±1.0
	if pVal > 0 {
		bits |= SaturationMask // Saturació positiva
	} else {
		bits |= SignMask | SaturationMask // Saturació negativa
	}

	// Retorna un RangeF64 utilitzant el constructor base
	return NewRangeF64FromU64(bits)
}

// Crea ±infinit
func NewRangeF64Infinite(pIsPositive bool) RangeF64 {
	bits := GroupBMask | SubgroupInfMask
	if !pIsPositive {
		bits |= SignMask
	}
	return NewRangeF64FromU64(bits)
}

// Crea un padding
func NewRangeF64Padding(pEnd *bool) RangeF64 {
	if pEnd == nil {
		return NewRangeF64FromU64(PaddingMask)
	} else if *pEnd {
		return NewRangeF64FromU64(PaddingStartMask)
	} else {
		return NewRangeF64FromU64(PaddingEndMask)
	}
}

// Crea un símbol
func NewRangeF64FromSymbol(pSym rune) RangeF64 {
	sym := uint64(pSym)
	u64 := GroupCMask | GroupCMask | sym

	return NewRangeF64FromU64(u64)
}

func NewRangeF64Percentage(Pf64 float64) RangeF64 {
	// Valida que el percentatge està dins del rang [0.0, 1.0]
	if Pf64 < 0.0 || Pf64 > 1.0 {
		return NewRangeF64Error(ERR_INVALID_PERCENTAGE, Pf64)
	}

	// Converteix el percentatge en valor binari brut
	// valueBits := uint64(pct * (1 << 52)) // Escala a 2^-52
	valueBits := uint64(math.Round(Pf64 * (1 << 52)))
	// fmt.Printf("valueBits abans 'AND' > pct: %f, raw_value_bits: %d, valueBits: %d\n", pct, uint64(pct*(1<<52)), valueBits)
	valueBits &= ValueMask // Manté només els bits del valor
	//  fmt.Printf("valueBits després 'AND' > pct: %f, raw_value_bits: %d, valueBits: %d\n", pct, uint64(pct*(1<<52)), valueBits)

	// Configura el Grup C i Subgrup C3
	finalBits := GroupCMask | SubGroupC3Mask | valueBits
	// fmt.Printf("finalBits després de combinar màscares: %064b\n", finalBits)
	// fmt.Printf("valueBits abans d'enviar: %064b\n", valueBits)

	return NewRangeF64FromU64(finalBits)
}

func NewRangeF64Identifier(sequenceType uint8, elementType uint8, elementID uint64) RangeF64 {
	// Validació
	if sequenceType > 3 {
		panic("sequenceType ha de ser entre 0 i 3")
	}
	if elementType > 7 {
		panic("elementType ha de ser entre 0 i 7")
	}
	if elementID >= (1 << 52) {
		panic("elementID excedeix els 52 bits disponibles")
	}

	// Construcció del valor
	value := GroupDMask | uint64(sequenceType)<<SequenceTypeShift | uint64(elementType)<<ElementTypeShift | (elementID & ElementIDMask)
	return NewRangeF64FromU64(value)
}

// Crea un RangeF64 per a símbols (caràcters) del grup B.1.
func NewRangeF64Symbol(pChar rune) RangeF64 {
	const groupB1Mask uint64 = 0b011 << 61 // Grup B amb subgrup 1 (símbols)
	const symbolMask uint64 = 0x7F         // 7 bits per al caràcter ASCII

	// Validació: Assegurem que pChar està dins del rang ASCII
	if pChar > 127 {
		panic("NewRangeF64Symbol: pChar fora del rang ASCII (0-127)")
	}

	// Codifiquem el símbol en el format RangeF64
	value := groupB1Mask | (uint64(pChar) & symbolMask)

	return NewRangeF64FromU64(value) // El camp "value" ha de ser intern en RangeF64
}

// PARSERS ----------------------------
// parseConfig llegeix un fitxer JSON i retorna una estructura NetworkConfig
func ParseConfig(filePath string) (*NetworkConfig, error) {
	// Obrim el fitxer
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %v", err)
	}
	defer file.Close()

	// Decodifiquem el JSON
	var config NetworkConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("error decoding config file: %v", err)
	}

	// Validem la configuració mínima
	if config.Layers <= 0 || len(config.NeuronsPerLayer) != config.Layers {
		return nil, fmt.Errorf("invalid configuration: mismatch between layers and neurons per layer")
	}

	return &config, nil
}

// GETTERS/SETTERS --------------------
func (sF64 RangeF64) GetF64Value() float64 { return sF64.value }
func (sF64 RangeF64) SetF64Value(pVal float64) RangeIntf {
	sF64.value = pVal
	return sF64
}
func (sF64 RangeF64) GetU64Value() uint64 { return F64ToU64(sF64.value) }
func (sF64 RangeF64) SetU64Value(pVal uint64) RangeIntf {
	sF64.value = U64ToF64(pVal)
	return sF64
}
func (sF64 RangeF64) GetPercentage() (float64, bool) {
	bits := F64ToU64(sF64.value)

	// Comprova si és del Grup C i Subgrup C3
	if (bits&GroupMask) != GroupCMask || (bits&SubgroupMask) != SubGroupC3Mask {
		return 0.0, false
	}

	// Extreu els bits del valor
	valueBits := bits & ValueMask
	decodedValue := float64(valueBits) / (1 << 52)

	// Comprova si el valor extret està dins del rang vàlid [0.0, 1.0]
	if decodedValue < 0.0 || decodedValue > 1.0 {
		return 0.0, false
	}

	decodedValue = math.Round(float64(valueBits)/(1<<52)*1000000000) / 1000000000
	return decodedValue, true
}
func (sF64 RangeF64) SetPercentage(pVal float64) RangeIntf {
	// TODO: Cal implementar la funció.
	return sF64
}

func (sF64 RangeF64) GetSequenceType() uint8 {
	if !sF64.IsIdentifier() {
		panic("El valor no pertany al Grup D")
	}
	return uint8((sF64.GetU64Value() & SequenceTypeMask) >> SequenceTypeShift)
}

func (r RangeF64) GetElementType() uint8 {
	if !r.IsIdentifier() {
		panic("El valor no pertany al Grup D")
	}
	return uint8((r.GetU64Value() & ElementTypeMask) >> ElementTypeShift)
}

// GetElementID retorna l'identificador únic de l'element (dins del Grup D).
func (sF64 RangeF64) GetElementID() uint64 {
	if !sF64.IsIdentifier() {
		panic("El valor no pertany al Grup D")
	}
	return sF64.GetU64Value() & ElementIDMask
}

// GetF64Symbol retorna el símbol (com a rune) d'un RangeF64 del Grup B.1.
func (sF64 RangeF64) GetF64Symbol() (rSymbol rune) {
	const groupB1Mask uint64 = 0b011 << 61 // Grup B amb subgrup 1
	const symbolMask uint64 = 0x7F         // 7 bits per al caràcter ASCII

	// Validem que el RangeF64 és un símbol del Grup B.1
	if (sF64.GetU64Value() & groupB1Mask) != groupB1Mask {
		panic("GetF64Symbol: El valor no és un símbol del Grup B.1")
	}

	// Extreiem el símbol dels 7 bits inferiors
	rSymbol = rune(sF64.GetU64Value() & symbolMask)
	return
}

// INTERFÍCIE 'RangeIntf' -------------
// Retorna cert només si la diferència entre els valors és inferior a Precisio64.
func (sF64 RangeF64) Equals(pOther RangeIntf) bool {
	if sF64.IsSaturated() || pOther.IsSaturated() {
		return sF64.IsSaturatedPos() == pOther.IsSaturatedPos() &&
			sF64.IsSaturatedNeg() == pOther.IsSaturatedNeg()
	}
	return math.Abs(sF64.ValueF64()-pOther.ValueF64()) < Epsilon64
}

// Retorna cert només en cas que el valor de la instància sigui inferior al de pOther.
func (sF64 RangeF64) LessThan(pOther RangeIntf) bool {
	if sF64.IsSaturated() || pOther.IsSaturated() {
		return sF64.IsSaturatedNeg() && !pOther.IsSaturatedNeg()
	}
	return (pOther.ValueF64() - sF64.ValueF64()) > Epsilon64
}

func (sF64 RangeF64) LessOrEqualThan(pOther RangeIntf) bool {
	if sF64.IsSaturated() || pOther.IsSaturated() {
		return sF64.IsSaturatedNeg() || sF64.Equals(pOther)
	}
	return sF64.LessThan(pOther) || sF64.Equals(pOther)
}

// Retorna cert només si el valor de la instància es superior al de pOther.
func (sF64 RangeF64) GreaterThan(pOther RangeIntf) bool {
	if sF64.IsSaturated() || pOther.IsSaturated() {
		return sF64.IsSaturatedPos() && !pOther.IsSaturatedPos()
	}
	return (sF64.ValueF64() - pOther.ValueF64()) > Epsilon64
}

// Retorna cert només si el valor de la instància és superior o igual al de pOther.
func (sF64 RangeF64) GreaterOrEqualThan(pOther RangeIntf) bool {
	if sF64.IsSaturated() || pOther.IsSaturated() {
		return sF64.IsSaturatedPos() || sF64.Equals(pOther)
	}
	return sF64.GreaterThan(pOther) || sF64.Equals(pOther)
}

// Funcions Is....()
func (sF64 RangeF64) IsNullValue() bool {
	bits := F64ToU64(sF64.value)
	return (bits & (GroupMask | SubgroupMask)) == (GroupBMask | SubgroupNullMask)
}

func (sF64 RangeF64) IsInfinitePos() bool {
	bits := F64ToU64(sF64.value)
	expectedBits := GroupBMask | SubgroupInfMask
	return bits == expectedBits && (bits&SignMask) == 0
}

func (sF64 RangeF64) IsInfiniteNeg() bool {
	bits := F64ToU64(sF64.value)
	expectedBits := SignMask | GroupBMask | SubgroupInfMask
	return bits == expectedBits
}

func (sF64 RangeF64) IsPadding() bool      { return sF64.IsPaddingStart() || sF64.IsPaddingEnd() }
func (sF64 RangeF64) IsPaddingStart() bool { return sF64.GetU64Value() == PaddingStartMask }
func (sF64 RangeF64) IsPaddingEnd() bool   { return sF64.GetU64Value() == PaddingEndMask }

// Símbols
func (sF64 RangeF64) IsSymbol() bool {
	u64 := sF64.GetU64Value() & 0b11111100_00000000_00000000_00000000_00000000_00000000_00000000_00000000
	u64 = u64 >> 56
	return u64&0b11000000 == 0b11000000
}

func (sF64 RangeF64) ToSymbol() rune {
	var rn rune = 0

	if sF64.IsSymbol() {
		u64 := sF64.GetU64Value() & 0b00000000_00000000_00000000_00000000_00000000_00000000_11111111_11111111
		rn = rune(u64)
	}

	return rn
}

func (sF64 RangeF64) IsIdentifier() bool {
	return (sF64.GetU64Value() & GroupMask) == GroupDMask
}

func IsFloat64Subnormal(pValue float64) bool {
	exponent := ExtractFloat64Exponent(pValue)
	return exponent == 0 // Exponent zero indica valors subnormalitzats
}
func (sF64 RangeF64) IsSubnormal() bool {
	return IsFloat64Subnormal(sF64.value)
}

func ExtractFloat64tMantissa(value float64) uint64 {
	return math.Float64bits(value) & 0x000FFFFFFFFFFFFF // Últims 52 bits
}

func (sF64 RangeF64) ExtractMantissa() uint64 {
	return ExtractFloat64tMantissa(sF64.value)
}

func ExtractFloat64Exponent(value float64) int {
	return int((math.Float64bits(value) >> 52) & 0x7FF) // Bits 52-62
}
func (sF64 RangeF64) ExtractExponent() int {
	return ExtractFloat64Exponent(sF64.value)
}

// Grups.
// [GPT] IsGroupA comprova si el valor és del Grup A
func (sF64 RangeF64) IsGroupA() bool {
	bits := F64ToU64(sF64.value)

	bits = ((bits << 1) >> 63)
	return bits == 0

	// // Exclou valors amb el bit de signe
	// if (bits & SignMask) != 0 {
	// 	return false
	// }

	// // Verifica que pertany exclusivament al Grup A
	// return (bits & GroupMask) == GroupAMask
}

// [GPT] IsGroupB comprova si el valor és del Grup B
func (sF64 RangeF64) IsGroupB() bool {
	return (F64ToU64(sF64.value) & GroupMask) == GroupBMask
}

// [GPT] IsGroupC comprova si el valor és del Grup C
func (sF64 RangeF64) IsGroupC() bool {
	bits := F64ToU64(sF64.value)
	return (bits & GroupMask) == GroupCMask
}

// [GPT] IsGroupD comprova si el valor és del Grup D
func (sF64 RangeF64) IsGroupD() bool {
	return (F64ToU64(sF64.value) & GroupMask) == GroupDMask
}

// // [GPT] IsGroupE comprova si el valor és del Grup E
// func (sF64 RangeF64) IsGroupE() bool {
// 	return (F64ToU64(sF64.value) & GroupMask) == GroupEMask
// }

// Operacions
// Suma dos valors Range mantenint el resultat dins [-1.0, +1.0]
func (sF64 RangeF64) Add(pOther RangeIntf) RangeIntf {
	// Verifica si qualsevol dels valors és null
	if sF64.IsNullValue() || pOther.IsNullValue() {
		return NewRangeF64Error(ERR_NULL_VALUE, sF64.GetF64Value())
	}

	// Verifica que lhs és del Grup A
	if sF64.IsGroupA() {
		// Si rhs és del Grup C (Percentatge)
		if pOther.IsGroupC() {
			rhsValue, valid := pOther.GetPercentage()
			// fmt.Printf("Extracted Percentage: %f (Valid: %v)\n", rhsValue, valid)
			if !valid {
				return NewRangeF64(math.NaN()) // Fora de rang: retorna NaN
			}

			// Escalatge correcte: lhs + (lhs * percentatge)
			lhsValue := sF64.GetF64Value()
			scaledValue := lhsValue + (lhsValue * rhsValue)

			// Comprova límits de saturació
			if scaledValue > 1.0 {
				return NewRangeF64Saturated(1.0, false)
			}
			if scaledValue < -1.0 {
				return NewRangeF64Saturated(-1.0, false)
			}

			return NewRangeF64(scaledValue)
		}

		// Si rhs no és vàlid, retorna NaN
		return NewRangeF64(math.NaN())
	}

	// Retorna un error si els grups són incompatibles
	return NewRangeF64Error(ERR_INVALID_ADD_OPERATION, pOther.ValueF64())
}

// Resta dos valors Range mantenint el resultat dins [-1.0, +1.0]
func (sF64 RangeF64) Sub(pOther RangeIntf) RangeIntf {
	result := sF64.value - pOther.ValueF64()

	// Si el resultat supera els límits, retornem valor saturat
	if result > 1.0 {
		return NewRangeF64Saturated(1, false)
	}
	if result < -1.0 {
		return NewRangeF64Saturated(-1, false)
	}

	return NewRangeF64(result)
}

// Multiplica dos valors Range mantenint el resultat dins [-1.0, +1.0]
func (sF64 RangeF64) Mul(pOther RangeIntf) RangeIntf {
	result := sF64.value * pOther.ValueF64()

	// Si el resultat supera els límits, retornem valor saturat
	if result > 1.0 {
		return NewRangeF64Saturated(1, false)
	}
	if result < -1.0 {
		return NewRangeF64Saturated(-1, false)
	}

	return NewRangeF64(result)
}

// Divideix dos valors Range mantenint el resultat dins [-1.0, +1.0]
func (sF64 RangeF64) Div(pOther RangeIntf) (RangeIntf, error) {
	// Comprovem divisió per zero
	if math.Abs(pOther.ValueF64()) < SmallThreshold64 {
		return sF64, fmt.Errorf("divisió per zero")
	}

	result := sF64.value / pOther.ValueF64()

	// Si el resultat supera els límits, retornem valor saturat
	if result > 1.0 {
		return NewRangeF64Saturated(1, false), nil
	}
	if result < -1.0 {
		return NewRangeF64Saturated(-1, false), nil
	}

	return NewRangeF64(result), nil
}

// ... i altres operacions comunes necessàries.

// Funcions de desencapçulament.
func (sF64 RangeF64) ValueF64() float64 {
	return F64ToF64(sF64.value)
}

func (sF64 RangeF64) ValueF32() float32 {
	return F64ToF32(sF64.value)
}

func (sF64 RangeF64) ValueI64() int64 {
	return F64ToI64(sF64.value)
}

func (sF64 RangeF64) ValueI32() int32 {
	return F64ToI32(sF64.value)
}

func (sF64 RangeF64) ValueU64() uint64 {
	return F64ToU64(sF64.value)
}

func (sF64 RangeF64) ValueU32() uint32 {
	return F64ToU32(sF64.value)
}

// Aquestes funcions poden ser útils en un futur.
func (sF64 RangeF64) AsF64() RangeF64 { return sF64 }

// func (sF64 RangeF64) AsF32() RangeF32 { return sF64 }
// func (sF64 RangeF64) AsI64() RangeI64 { return sF64 }
// func (sF64 RangeF64) AsI32() RangeI32 { return sF64 }
// func (sF64 RangeF64) AsU64() RangeU64 { return sF64 }
// func (sF64 RangeF64) AsU32() RangeU32 { return sF64 }

// Funcions de saturació
func (sF64 RangeF64) IsSaturated() bool {
	return sF64.IsSaturatedPos() || sF64.IsSaturatedNeg()
}

func (sF64 RangeF64) IsSaturatedPos() bool {
	bits := F64ToU64(sF64.value)
	return (bits & (SignMask | GroupMask)) == GroupDMask // No signe + Group D
}

func (sF64 RangeF64) IsSaturatedNeg() bool {
	bits := F64ToU64(sF64.value)
	return (bits & (SignMask | GroupMask)) == (SignMask | GroupDMask) // Amb signe + Group D
}

// Gestió d'errors...
// Retorna cert només si el RangeF64 representa un error.
func (sF64 RangeF64) IsError() bool {
	bits := F64ToU64(sF64.value)
	return (bits & GroupMask) == GroupBMask
}

// Retorna el codi d'error a partir dels bits definits.
func (sF64 RangeF64) ErrorCode() int {
	bits := F64ToU64(sF64.value)
	if !sF64.IsError() {
		return 0 // No és un error
	}
	return int((bits >> 48) & 0xFFFF) // Bits 48-63
}

// Retorna la resta de bits per a poder incorporar un paràmetre a l'error.
func (sF64 RangeF64) ErroneousValue() float64 {
	bits := F64ToU64(sF64.value)
	if !sF64.IsError() {
		return 0 // No és un error
	}
	valueBits := bits & ValueMask
	return U64ToF64(valueBits)
}
