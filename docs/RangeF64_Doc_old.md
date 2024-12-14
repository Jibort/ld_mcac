# Especificació del Model RangeF64

## Introducció
El tipus `RangeF64` és una representació compacta basada en l'estàndard IEEE 754 per a codificar informació que combina valors numèrics, símbols i altres metadades en un format de 64 bits. Aquesta estructura s'ha dissenyat per ser utilitzada en xarxes neuronals i sistemes de processament avançat.

En la mesura del possible l'estructura de F64 serà compatible amb l'estructura de F32, I64, U32, ...

---

## **Grups disponibles**

Donat que els valors que farem servir durant els càlculs de qualsevol de les xarxes té un rang de [-1.0, +1.0] existeixen moltes combinacions possibles que mai es faran servir com a valor.
En el cas del format F64 disposem d'una mantissa de 52 bits que oferiran un rang de precisió de 15 a 17 dígits decimals (per a assegurar la precisió quantitzarem a 15 bits), però en tots els casos el bit 62 sempre serà '0'.

Això significa que qualsevol RangeF64 amb bit 62 igual a '0' (o amb exponent 2047) correspon a un valor de càlcul (Group A), mentre que quan el bit 62 sigui '1' (i l'exponent no sigui 2047) correspondrà a altres grups de dades (Groups B-C-D).

El fet que al format F32 només càpiguen 4 grups (2 bits) m'ha fet decidir reduir també els tipus a 4 en F64.

### **Group A (Rang de valors [-1.0,+1.0]):** 0bS + 11e + 52m
- **Nota important:** Els valors amb exponent 2047 (tot l'exponent a 1) són gestionats directament pel sistema segons l'estàndard IEEE 754 i s'inclouen dins del Group A. Això inclou ±inf i NaN, que són valors natius compatibles.
- **Estructura**: 1b signe + 11b exponent + 52b mantissa.
- **Funció:** Representa valors numèrics dins el rang [-1.0, +1.0] amb 15 dígits de precisió exacta.
- **Identificació:** Bits 63-62 = 0bS0, on:
    - 'S' és el signe del valor (0: positiu, 1: negatiu)
    - el bit 62 sempre és '0'.
- **Casos especials:** Els exponents (11e) estàn a restringits a +1022 i +1023.
- **exemples**
    - +0.625: 0b01111111_11100100_00000000_00000000_00000000_00000000_00000000_00000000
    - -0.12345678901234: 0b1_01111111_01111110_10110111_00001010_10001111_01011100_00101001_00111111

### **Group B (Valors Especials i Personalitzats):**
- **Restricció:** Els valors amb exponent 2047 no formen part del Grup B, ja que estan reservats per als valors natius IEEE 754 (±inf i NaN). El Grup B només conté configuracions personalitzades, com ±sat i nul, amb mantisses específiques que no es poden confondre amb els valors d'IEEE 754.
- **Funció:** Codifica valors especials personalitzats, com ±sat i nul, sense conflicir amb valors IEEE 754 estàndard.
- **Estructura**: 2b 'S1' ('01' o '11') + '00'g grup B + 'ss' subgrup + 58d dades.
- **Identificació:** Bits 63-58 = 0bS100_ss
    - 'S' només és el signe quan el subgrup és '01' o '11'.
    - '00' estableix el subgrup.
- **Subgrups**:
        - Nul  (0b1100_00....)
        - +inf (0b0100_00....)
        - -inf (0b1100_00....)
    - **Subgrup '10':** Símbols (0b0100_1000 + 6bytes a 0x00 + 16bits de símbol)
        - Amb 2 bytes (16b) tenim rang suficient per a introduir els símbols nous que vagin apareixent.
    - **Subgrup '11':** ±Saturació (0bS100_1100....)
        - Representa valors de saturació personalitzats dins del grup B. Manté la identificació diferenciada dels valors estàndard IEEE 754.
        - En aquest subgrup la representació del signe i dels primers 52b (51-0) corresponen al signe i la mantissa en el Grup A.
        - D'aquesta forma es pot mantenir tant la marca de saturació per a tot el rang [-1.0,+1.0] com la mateixa representació dels valors en ambdós grups.


### **Group C (Tokens i Percentatges):** 0bp1 + '01'g (grup C) + subgrup + dades
- **Estructura General**: 2b 'p1' + '01'g grup C + subgrup + dades (segons subgrup).
- **Funció**: Representa percentatges, rangs de percentatges, o identificadors clau únics associats a tokens.
- **Subgrups**:
  - **Subgrup 00 (Percentatge):**
    - Representa un percentatge pur entre 0.0 i 1.0.
    - **Estructura**: 33 bits per al percentatge com a fracció decimal (rang 0.000000% - 100.000000%).
    - **Exemple**: 50% es codifica com `0bp1_00_100000_000000_...`.

  - **Subgrup 01 (Rang de Percentatges):**
    - Representa un rang de percentatges amb mínim i màxim.
    - **Estructura**:
      - 16 bits inicials: Percentatge mínim.
      - 16 bits següents: Percentatge màxim.
      - Exemple: Rang 10%-90% es codifica com `0bp1_01_00011001_01011110_...`.

  - **Subgrup 10 (Token Clàssic):**
    - Identifica un token únic associat.
    - **Estructura**: 33 bits per dades del token i 27 bits d'identificador.

  - **Subgrup 11 (Reservat):**
    - Espai reservat per a futures extensions.

- **Funcions**:
  - **Creació de Percentatge**: `NewPercentage(float64) -> RangeF64`.
  - **Creació de Rang de Percentatges**: `NewPercentageRange(float64, float64) -> RangeF64`.
  - **Escalatge**: `Scale(RangeF64, float64) -> RangeF64`.
  - **Validació**:
    - `IsPercentage(RangeF64) -> bool`.
    - `IsWithinRange(RangeF64, RangeF64) -> bool`.

Aquest subgrup permet gestionar valors percentuals i els amplia amb rangs, tot mantenint compatibilitat amb tokens clàssics.

### **Group D (Meta significats):** 0bm1 + '10'g (grup D) + 60m
- **Estructura**: 2b 'm1' + '10'g grup D + 61 bits de meta valors.
- **Funció**: Els metavalors encara no estan definits.
- **Identificació:** Bits 63-60 = 0bm110...
    - 'm' Tots els bits p's (1+60) Determinaran metavalors.

### **Group E (extensió F64):** 0b?1 + '11'g (grup E) + 60?
- **Estructura**: 2b '?1' + '11'g grup E + 60 bits '?' meta valors.
- **Funció**: 
    - Aquest grup només està disponible als tipus Range?64 perquè no té espai als tipus de 32 bits.
    - La seva funció està per determinar si finalment es fa servir.
- **Identificació:** Bits 63-60 = 0bm111...
    - 'm' Tots els bits m's (1+60) per a determinar si cal.

## Relació de símbols
- Espais i control
    * ' ': 0x01, '\t': 0x02, '\n': 0x03, '\r': 0x04,

- Lletres majúscules (inclou accents i diacrítics)
	- 'A': 0x10, 'Á': 0x11, 'À': 0x12, 'Ä': 0x13, 'Â': 0x14,
	- 'B': 0x20,
	- 'C': 0x30, 'Ç': 0x31,
	- 'D': 0x40,
	- 'E': 0x50, 'É': 0x51, 'È': 0x52, 'Ê': 0x53, 'Ë': 0x54,
	- 'F': 0x60,
	- 'G': 0x70,
	- 'H': 0x80,
	- 'I': 0x90, 'Í': 0x91, 'Ì': 0x92, 'Ï': 0x93,
	- 'J': 0xA0,
	- 'K': 0xB0,
	- 'L': 0xC0,
	- 'M': 0xD0,
	- 'N': 0xE0, 'Ñ': 0xE1,
	- 'O': 0xF0, 'Ó': 0xF1, 'Ò': 0xF2, 'Ö': 0xF3, 'Ô': 0xF4,
	- 'P': 0x100,
	- 'Q': 0x110,
	- 'R': 0x120,
	- 'S': 0x130,
	- 'T': 0x140,
	- 'U': 0x150, 'Ú': 0x151, 'Ù': 0x152, 'Ü': 0x153, 'Û': 0x154,
	- 'V': 0x160,
	- 'W': 0x170,
	- 'X': 0x180,
	- 'Y': 0x190,
	- 'Z': 0x1A0,

- Lletres minúscules (inclou accents i diacrítics)
	- 'a': 0x1B0, 'á': 0x1B1, 'à': 0x1B2, 'ä': 0x1B3, 'â': 0x1B4,
	- 'b': 0x1C0,
	- 'c': 0x1D0, 'ç': 0x1D1,
	- 'd': 0x1E0,
	- 'e': 0x1F0, 'é': 0x1F1, 'è': 0x1F2, 'ê': 0x1F3, 'ë': 0x1F4,
	- 'f': 0x200,
	- 'g': 0x210,
	- 'h': 0x220,
	- 'i': 0x230, 'í': 0x231, 'ì': 0x232, 'ï': 0x233,
	- 'j': 0x240,
	- 'k': 0x250,
	- 'l': 0x260,
	- 'm': 0x270,
	- 'n': 0x280, 'ñ': 0x281,
	- 'o': 0x290, 'ó': 0x291, 'ò': 0x292, 'ö': 0x293, 'ô': 0x294,
	- 'p': 0x2A0,
	- 'q': 0x2B0,
	- 'r': 0x2C0,
	- 's': 0x2D0, 'ß': 0x2D1,
	- 't': 0x2E0,
	- 'u': 0x2F0, 'ú': 0x2F1, 'ù': 0x2F2, 'ü': 0x2F3, 'û': 0x2F4,
	- 'v': 0x300,
	- 'w': 0x310,
	- 'x': 0x320,
	- 'y': 0x330,
	- 'z': 0x340,

- Números
	- '0': 0x400, '1': 0x410, '2': 0x420, '3': 0x430, '4': 0x440,
	- '5': 0x450, '6': 0x460, '7': 0x470, '8': 0x480, '9': 0x490,

- Puntuació
	- '.': 0x500, ',': 0x510, ';': 0x520, ':': 0x530,
	- '!': 0x540, '?': 0x550,
	- '(': 0x560, ')': 0x570,
	- '[': 0x580, ']': 0x590,
	- '{': 0x5A0, '}': 0x5B0,
	- '-': 0x5C0, '_': 0x5D0,
	- '\'': 0x5E0, '"': 0x5F0,

- Nous símbols especials
	- 'º': 0x600, 'ª': 0x610, '·': 0x620, '|': 0x630,
	- '#': 0x640, '$': 0x650, '%': 0x660, '&': 0x670,
	- '¬': 0x680, '~': 0x690, '€': 0x6A0,

- Elements Especials
	- rune(0xFFFC): 0xFFF3, // Nul
	- rune(0xFFFD): 0xFFF0, // Error
	- rune(0xFFFE): 0xFFF1, // Desconegut
	- rune(0xFFFF): 0xFFF2, // Qualsevol
