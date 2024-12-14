# Especificació del Model RangeF64

## Introducció
El tipus `RangeF64` és una representació compacta basada en l'estàndard IEEE 754 per a codificar informació que combina valors numèrics, símbols i altres metadades en un format de 64 bits. Aquesta estructura s'ha dissenyat per ser utilitzada en xarxes neuronals i sistemes de processament avançat.

En la mesura del possible, l'estructura de F64 serà compatible amb l'estructura de F32, I64, U32, ...

---

## **Grups disponibles**

Donat que els valors que farem servir durant els càlculs de qualsevol de les xarxes té un rang de [-1.0, +1.0], existeixen moltes combinacions possibles que mai es faran servir com a valor. En el cas del format F64, disposem d'una mantissa de 52 bits que ofereix un rang de precisió de 15 a 17 dígits decimals (per a assegurar la precisió, quantitzarem a 15 bits).

Això significa que qualsevol RangeF64 amb bit 62 igual a '0' (o amb exponent 2047) correspon a un valor de càlcul (**Group A**), mentre que quan el bit 62 sigui '1' (i l'exponent no sigui 2047) correspondrà a altres grups de dades (**Groups B-C-D**).

---

### **Group A (Rang de valors [-1.0,+1.0]):**
- **Nota important:** Els valors amb exponent 2047 (tot l'exponent a 1) són gestionats directament pel sistema segons l'estàndard IEEE 754 i s'inclouen dins del Group A. Això inclou ±inf i NaN, que són valors natius compatibles.
- **Estructura:** 1b signe + 11b exponent + 52b mantissa.
- **Funció:** Representa valors numèrics dins el rang [-1.0, +1.0] amb 15 dígits de precisió exacta.
- **Exemple de representació bit a bit:**
  | Bit   | 63 | 62-52 (Exponent) | 51-0 (Mantissa)            |
  |-------|----|------------------|---------------------------|
  | +inf  |  0 | 11111111111      | 0000000000000000000000000 |
  | -inf  |  1 | 11111111111      | 0000000000000000000000000 |
  | NaN   |  X | 11111111111      | !=0                       |

---

### **Group B (Valors Especials i Personalitzats):**
- **Funció:** Codifica valors especials personalitzats, com ±sat i nul, sense conflicir amb valors IEEE 754 estàndard.
- **Subgrups:**
  - **Subgrup B.1 (Símbols):**
    Assigna bits a símbols específics com funcions de teclat i banderes. Exemples:
    | Símbol | Bits           |
    |--------|----------------|
    | F3     | `F3`           |
    | Tab    | `tab`          |
    | Majúsc | `bm`           |
    - Els símbols futurs es marquen amb bits reservats `?` per evitar conflicència.

  - **Subgrup B.2 (Extensions futures):**
    - Espai reservat amb dades (`d`) per implementacions futures.
    - Els bits marcats com `?` asseguren que no hi hagi conflicència amb altres subgrups.

  - **Subgrup B.3 (±Saturació i Nul):**
    - **±Saturació:** Representació clara per valors màxims (+1.0) i mínims (-1.0).
    - **Nul:** Definició específica que utilitza un patró clar per evitar solapament.
    - Exemples:
      | Valor      | Bits                                    |
      |------------|----------------------------------------|
      | Saturació  | Patró binari clar amb mantissa i signe. |
      | Nul        | Patró fixat amb `0` en mantissa.       |

---

### **Group C (Tokens i Percentatges):**
- **Detalls:**
  - Defineix identificadors únics (`c`) amb espai suficient per múltiples aplicacions.
  - Percentatges representats amb alta precisió.
  - Exemple:
    | Percentatge | Bits          |
    |-------------|---------------|
    | 50%         | `0b_1100...`  |

---

### Notes addicionals:
- **Compatibilitat IEEE 754:** Els valors del Group A es gestionen completament segons l'estàndard, mentre que els Group B i C contenen extensions personalitzades.
- **Reservats:** Alguns bits en els grups B i C estan reservats per evitar conflicència amb futures ampliacions.
- **Precisions tècniques:** La documentació s'ha adaptat per reflectir amb exactitud la taula de representació bit a bit del fitxer ODS.

---

Aquest document està dissenyat per garantir que l'estructura del RangeF64 sigui compatible tant amb l'estàndard IEEE 754 com amb les necessitats específiques del projecte.
