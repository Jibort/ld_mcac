# Esquema de les Interfícies des de les més Completes fins a les més Bàsiques
Aquest enfocament comença amb les interfícies més especialitzades (com les del Grup B i els valors IEEE 754) i baixa fins a les interfícies més bàsiques com RangeIntf.

## Interfícies més Completes
### RangeF64InfiniteIntf:
- Depèn de: RangeF64IE754Intf, RangeComparableIntf
- Funcions addicionals: IsInfinitePos(), IsInfiniteNeg()

### RangeF64IE754Intf:
- Depèn de: RangeMathOperations, RangeF64Intf
- Funcions addicionals: IsSubnormal(), IsNormalNear0(), IsNormalFar0(), Value(), SetValue()

### RangeF64GroupBIntf:
- Depèn de: RangeF64Intf
- Subinterfícies:
    - RangeF64SymbolIntf: Symbol()
    - RangeF64SaturatedIntf: IsSaturatedPos(), IsSaturatedNeg()
    - RangeF64PaddingIntf: IsStartPadding(), IsCommonPadding(),   - IsEndPadding(), PaddingRune()
    - RangeF64NullIntf: IsNull()
    - RangeF64CoordsIntf: Coordinates()
    - RangeF64PercRangeIntf: Margins()
    - RangeF64ErrorIntf: Code(), Arguments()

### RangeF64TokenIntf:
- Depèn de: RangeF64Intf
- Funcions addicionals: Category(), Fiability(), RelativeWeight(), Token()

### RangeF64MetaRefIntf:
- Depèn de: RangeF64Intf
- Funcions addicionals: Network(), Layer(), Neuron(), Synapse()

## Interfícies Intermèdies
### RangeF64GroupAIntf:
- Depèn de: RangeF64Intf
- Funcions addicionals: Sign(), Exponent(), Mantissa(), IsSubnormal(), IsNormalNear0(), IsNormalFar0(), IsInfinite(), IsInfinitePos(), IsInfiniteNeg()

### RangeF32GroupAIntf:
- Depèn de: RangeF32Intf
- Similar a RangeF64GroupAIntf però adaptada a float32.

## Interfícies Generals
### RangeF64Intf:
- Depèn de: Range64Intf
- Funcions addicionals: ToIntf(), ToF32()

### RangeF32Intf:
- Depèn de: Range32Intf
- Funcions addicionals: ToIntf(), ToF64()

### Range64Intf:
- Depèn de: RangeIntf
- Funcions addicionals: IsF64(), IsU64(), As32()

### Range32Intf:
- Depèn de: RangeIntf
- Funcions addicionals: IsF32(), IsU32(), As64()

## Interfícies Bàsiques
### RangeMathOperations:
- Depèn de: RangeComparableIntf
- Funcions: Add(), Sub(), Mul(), Div()

### RangeComparableIntf:
- Depèn de: RangeIntf
- Funcions: LessThan(), LessOrEqualThan(), GreaterThan(), GreaterOrEqualThan()

### RangeIntf:
- Interfície base amb funcions com:
    - Comparació: Equals()
    - Grups: IsGroupA(), IsGroupB(), IsGroupC(), IsGroupD()
    - Extracció: ExtractMantissa(), ExtractExponent()
    - Conversió: Is64(), Is32()

## Avantatges d'aquest Esquema
### Visió Jeràrquica Descendent:
- Comença amb les interfícies més completes (amb dependències explícites) i acaba amb les més bàsiques.
Claredat de les Dependències:

- És fàcil veure quines interfícies s’estenen per proporcionar més funcionalitat.
Esportivitat entre Tipus (F64 i F32):

 - Els paral·lelismes entre interfícies per a float64 i float32 són evidents, facilitant l'extensió i el manteniment.
