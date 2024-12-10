
# Especificació del Model RangeF64

## Introducció

El tipus `RangeF64` és una representació compacta basada en l'estàndard IEEE 754 per a codificar informació que combina valors numèrics, símbols i altres metadades en un format de 64 bits. Aquesta estructura s'ha dissenyat per ser utilitzada en xarxes neuronals i sistemes de processament avançat.

---

## **Grups i Subgrups**

### **Grup A (R[-+]):**
- **Funció:** Representa valors numèrics dins el rang [-1.0, +1.0].
- **Identificació:** Bit 62 = 0.
- **Casos especials:** Valors amb exponents restringits a +1022 i +1023.

### **Grup B (Valors especials):**
- **Funció:** Codifica valors especials com `NaN`, `+Inf`, `-Inf`, i `Null`.
- **Identificació:** Bits 62-61 = `01`.
- **Subgrups:**
  - `B1`: `NaN`
  - `B2`: `+Inf`, `-Inf`
  - `B3`: `Null`
  - **Futur:** Identificadors de control i delimitadors.

### **Grup C (Tokens i símbols):**
- **Funció:** Representa símbols, tokens i meta informació.
- **Identificació:** Bits 62-61 = `10`.
- **Subgrups:**
  - `C1`: Tokens simples.
  - `C2`: Tokens compostos.
  - `C3`: Símbols (lletres, signes de puntuació, etc.).

#### **Símbols del Subgrup C3**
Els símbols inclosos actualment són:
- Lletres: `A-Z`, `a-z` (inclou caràcters específics d'idiomes: `ñ`, `ç`, `ä`, etc.).
- Dígits: `0-9`.
- Signes: `@`, `#`, `$`, `%`, `&`, `*`, `|`, etc.
- Parèntesis i altres delimitadors: `()`, `[]`, `{}`, `< >`.
- Espai i tabulador.
- Símbols especials: `€`, `£`, `¬`, etc.

#### **Bit 63: Meta Informació**
El bit 63 s'utilitza per indicar si un token és estàndard o porta informació associada:
- **Bit 63 = 0:** Tokens estàndard.
- **Bit 63 = 1:** Tokens de meta informació (debugging, context, traçabilitat).

**Exemple per `%`:**
- **Estàndard:** `0x400C000000000660`
- **Meta Informació:** `0xC00C000000000660`

#### **Càlcul del Valors RangeF64**
Els símbols es codifiquen combinant:
1. `GroupCMask`: `0x4000000000000000` (Bits 62-61).
2. `SubGroupC3Mask`: `0x00C0000000000000` (Bits 60-57).
3. `Bit 63`: Meta informació o estàndard.
4. `ID`: Identificador únic assignat (Bits 11-0).

---

### **Grup D (Valors saturats):**
- **Funció:** Marca valors numèrics com saturats.
- **Identificació:** Bits 62-61 = `11`.
- **Casos:** `+1.0*`, `-1.0*`.

### **Grup E (Expansions futures):**
- **Funció:** Reservat per a funcionalitats no implementades encara.

---

## **Codificació i Decodificació de Símbols**

### **Codificació**
La funció `EncodeSymbol(symbol rune) RangeF64` crea una instància de `RangeF64` per al símbol especificat. Exemple:

```go
symbol := '%'
encoded := EncodeSymbol(symbol)
fmt.Printf("RangeF64: %x
", encoded.GetValue())
```

Resultat:
```
RangeF64: 0x40000C660
```

### **Decodificació**
La funció `DecodeSymbol(r RangeF64) rune` recupera el símbol codificat en un `RangeF64`. Exemple:

```go
encoded := NewRangeU64(0x40000C660)
decoded := DecodeSymbol(encoded)
fmt.Printf("Symbol: %c
", decoded)
```

Resultat:
```
Symbol: %
```

---

## **Proves i Validació**

### **Tests Implementats**
- **`TestEncodeDecodeSymbol`:** Valida que els símbols es codifiquen i decodifiquen correctament.
- **`TestNewSpecialSymbols`:** Comprova que els valors esperats per símbols especials coincideixen amb els resultats calculats.
- **Validació del bit 63:** Assegura que els tokens de meta informació es diferencien correctament dels estàndards.

---

## **Consideracions Finals**
Aquest model és flexible i pot ser ampliat per incloure nous tipus d'entitats al Grup C (per exemple, rangs de dates) o funcionalitats addicionals al Grup E.

---

## **Pròximes Passes**
1. Definir nous subgrups dins del Grup C.
2. Desenvolupar casos d'ús per al Grup E.
3. Optimitzar el codi per a entorns amb recursos limitats.

---

## **Autoria**
- **Data de creació:** 2024/12/10 dt JIQ
- **Contribució:** GPT assistit per JIQ
