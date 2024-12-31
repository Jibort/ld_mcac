# Especificació del Model RangeF64

## Introducció
El tipus `RangeF64` és una representació compacta basada en l'estàndard IEEE 754 per a codificar informació que combina valors numèrics, símbols i altres metadades en un format de 64 bits. Aquesta estructura s'ha dissenyat per ser utilitzada en xarxes neuronals i sistemes de processament avançat.

En la mesura del possible l'estructura de F64 serà compatible amb l'estructura de F32, I64, U32, ...

---

## **Grups disponibles**
Donat que els valors que farem servir durant els càlculs de qualsevol de les xarxes té un rang de [-1.0, +1.0] existeixen moltes combinacions possibles que mai es faran servir com a valor.

En el cas del format F64 disposem d'una mantissa de 52 bits que oferiran un rang de precisió de 15 a 17 dígits decimals (per a assegurar la precisió quantitzarem a 15 bits), però en tots els casos el bit 62 sempre serà '0' (tant per a números normalitzat com desnormalitzats).

Això significa que qualsevol RangeF64 on el bit 62 sigui igual a '0' o on l'exponent sigui 2047, ens trobarem amb un valor numèric estàndard (en cas de '0' per a valors en el rang [-1.0, +1.0]) i deixarem que el sistema (CPU o CUDA) operi amb ells amb normalitat. A aquest conjunt de valors float64 els anomenarem _Group A_.

El fet que al format F32 només càpiguen 4 grups (2 bits) ens limita també en floa64 a aquest número de grups (tot i que en 64bits encara hi hauria espai per a un cinquè grup).

### **Group A (Rang de valors [-1.0,+1.0]):** S + 11e + 52m
- __S__: Bit de signe
- __11e__: 11 bits d'exponent
- __52m__: 52 bits de mantisa
- **Nota:** Tots els subgrups del group A ja els gestiona la implementació float64 IEEE 754 nativament, tant en CPU com en CUDA.
- **Identificació:** bit 62 igual a '0' o exponent igual a 2047.

#### **Subgroup A.1 (normalitzats):** S + '0'e + 10e + 52m
- Subgrup per a valors que no s'apropen en excés a 0.0.
- En aquest subgrup trobarem sempre un exponent que comença per 0 i es troba en el rang [−1,−1022].
- **Exemples:**
	- ```+0.1234567 = 00111111_11011111_01111101_00110111_10001010_10110011_11101100_10101110```
	- ```-0.1234567 = 10111111_11011111_01111101_00110111_10001010_10110011_11101100_10101110```
	- ```-1.0: 10111111_111100000_00000000_00000000_00000000_00000000_00000000_00000000```
	- ```+1.0: 00111111_111100000_00000000_00000000_00000000_00000000_00000000_00000000```
		
#### **Group A.2 (subnormalitzats):** S + '00000000000'e + 52m
- Subgrup per a valors que s'apropen extraordinàriament a 0.
- En aquest cas l'exponent són 11 bits a zero.
- **Exemples:** 
	- +2*10<sup>−1074</sup>:```0_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000001```
	- -2*10<sup>−1073</sup>:```1_00000000_00000000_00000000_00000000_00000000_00000000_00000000_00000010```

#### **Group A.3 (±inf i NaN):** S + '11111111111'e + 52m
- El subggrup amb exponent 2047 (tots els bits a 1) són gestionats directament pel sistema segons l'estàndard IEEE 754 encara que no siguin números reals.
- **Funció:** Representar els valors conceptuals estàndard -inf, +inf i NaN.
	- ```+inf: 01111111_11110000_00000000_00000000_00000000_00000000_00000000_00000000```
	- ```-inf: 11111111_11110000_00000000_00000000_00000000_00000000_00000000_00000000```
	- ``` NaN:   01111111_11111000_00000000_00000000_00000000_00000000_00000000_00000000```
	

### **Group B (Valors Especials, Símbols i Personalitzats):** S + '100' + 2sg + 58d
- __S__: Bit de signe (hi ha casos en que aquest bit no representa res i d'altres on es pot afegir als 58 bits de dades).
- __'1_00'__: Identificació del group B (bits 62-60).
- __2sg__: 2 bits de subgrup (B.1 = '00', B.2 = '01', B.3 = '10' i B.4 = '11')
- __58d__: 58 bits per a codificar la informació segons el subgrup.
- El Grup B només conté configuracions personalitzades: símbols, estat de tecles especials del teclat, paddings, saturació i valor nul.

#### **Subgroup B.1:** (símbols)
- Aquest subgrup conté la codificació de tots els caràcters que poden trobar-nos en els missatges que rep la xarxa (tant en català com en alemany, francès, anglès i castellà).
- **Identificació:** Tots els elements continguts en aquest subgrup comencen amb els bits '110000'.
- **Format:**```1   1   0   0   0   0   F1   F2  |F3  F4  F5  F6  F7  F8  F9  F10 |F11 F12 tab bm  she shd cte ctd|fnc win osx alt alg del sup ?|? ? ? ? ? ? ? ?|? ? ? ? ? ? ? ?|c c c c c c c c|c c c c c c c c```, on:
	- Fx: Bandera de tecla de funció ‘x’ 
	- tab: Bandera de tecla ‘tabulador’ 
	- bm: Bandera de bloq. de majúscules 
	- she: Bandera de tecla ‘shift equerra’ 
	- shd: Bandera de tecla ‘shift dreta’ 
	- cte: Bandera de tecla ‘ctrl esquerra’ 
	- ctd: Bandera de tecla ‘ctrl dreta’ 
	- fnc: Bandera de tecla de ‘funció’
	- win: Bandera de tecla de ‘windows’
	- osx: Bandera de tecla de ‘osx’
	- alt: Bandera de tecla ‘alt’
	- alg: Bandera de tecla ‘alt gr’
	- del: Bandera de tecla ‘delete’
	- sup: Bandera de tecla ‘suprimir’
	- ?'s: Bits per a futures extensions
	- c’s: identificador del símbol (taula al final del document)
	
#### **Subgroup B.2:** (paddings i extensions futures)
- **Identificació:** Tots els elements continguts en aquest subgrup comencen amb els bits 'F10001tt', on 'F' pot ser '0' (padding genèric o inici de padding) o '1'(final de padding) i 'tt' per ara només està definit per a '00'.
- ```Padding d’inici:  01000100_00000000_00000000_00000000_00000000_00000000_11111111_00001010```
- ```Padding de final:11000100_00000000_00000000_00000000_00000000_00000000_11111111_00001011```
- ```Padding genèric:  01000100_00000000_00000000_00000000_00000000_00000000_00000000_00000000```
- La resta de possibles valors està reservat per a futures extensions.

#### **Subgroup B.3:** (saturació i nul)
- **Format:** ```S1_00_10_x_000un_52m```, on:
	- __S1__: Bandera de signe i bit 62 a '1'
	- __00__: Grup B
	- __10__: Subgrup B.3
	- __x__: Només '1' en cas de valor nul.
	- __000__: Tres bits a '0'.
	- __u__: Només '1' quan es tracta de de +1.0* o -1.0* (+1 i -1 saturats)
	- __n__: Només a '1' per a números float64 subnormalitzats.
	- __52m__: 52 bits de mantisa que coincidèixen amb la mantisa del grup A, però en aquest cas per  valors saturats.
- **Nota:** En aquest subgrup trobem el mateix rang de valors normalitzats i subnormalitzats però en la seva forma saturada. Es pot reconstruir el format del grup A tenint en compte el signe i el flag 'n'.
- D'aquesta forma es pot mantenir tant la marca de saturació per a tot el rang [-1.0,+1.0] com la mateixa representació dels valors en ambdós grups.
- **Exemples:**
	- ```nul: 01001010_00100000_00000000_00000000_00000000_00000000_00000000_00000000```
	- ```-1.0*: 11001000_000100000_00000000_00000000_00000000_00000000_00000000_00000000```
	- ```+1.0*: 01001000_000100000_00000000_00000000_00000000_00000000_00000000_00000000```
	- ```+0.1234567*: 01001000_00001111_01111101_00110111_10001010_10110011_11101100_10101110```
	- ```-0.1234567*: 01001000_00001111_01111101_00110111_10001010_10110011_11101100_10101110```

#### **Subgroup B.4:** (extensions)
- Subgrup reservat per a futures expansions.

### **Group C (Tokens i Percentatges):** 'p1' + '01'g (grup C) + subgrup + dades
- **Estructura General**: 2b 'p1' + '01'g grup C + subgrup + dades (segons subgrup), on 'p' només és '1' per a identificadors de tokens.
- **Funció**: Representa percentatges, rangs de percentatges, o identificadors clau únics associats a tokens.
- **Subgrups**:
  - **Subgrup 00 (Percentatge):**
    - Representa un percentatge pur entre 0.0 i 1.0.
    - **Estructura**: 33 bits per al percentatge com a fracció decimal (rang 0.000000% - 100.000000%).
    - **Exemple**: 50% es codifica com ```01_00_100000_000000_...```.

  - **Subgrup 01 (Rang de Percentatges):**
    - Representa un rang de percentatges amb mínim i màxim.
    - **Estructura**:
	  - ```01_01_00011001_...```
      - 16 bits inicials: Percentatge mínim.
      - 16 bits següents: Percentatge màxim.
      - **Exemple:** Rang 10%-90% es codifica com ```01_01_00011001_01011110_...`.

  - **Subgrup 10 (Token Clàssic):**
    - Identifica un token únic associat.
    - **Estructura:** 33 bits per dades del token i 27 bits d'identificador.
	- **Exemple:** ```'11_01_10'``` i 58 bits dels quals només els primers 27 codificaran l'identificador del token.
	
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
#### Estructura General:
Els identificadors es codifiquen dins d’un esquema jeràrquic.

Es reserven espais suficients per a identificar components amb granularitat fina.

#### Estructura:####
```code
0bm1 + '10'g (grup D) + tipus (3 bits) + jerarquia (57 bits)
```
- 0bm1: Identificador del grup de meta-significats (Grup D).
- 10g: Codifica que es tracta del Grup D.
- Tipus (3 bits): Distingir entre els diferents elements:
   - 000: Xarxa (Net).
   - 001: Capa (Layer).
   - 010: Neurona (N).
   - 011: Sinapsi (Syn).
   - 1xx: Altres (Reservat per extensions).
- Jerarquia (57 bits): Codificació específica per a identificar instàncies concretes dins la xarxa.
#### **Exemples:**
<u>Identificador d'una Xarxa:</u>

Xarxa 1 (Net-1).
Codificació:
Copia el codi
0bm1_10_000_0000000000000000000000000000000000000000000000000000001
Identificador d'una Capa:

Xarxa 1, Capa 3 (Net-1 -> Layer-3).
Codificació:
Copia el codi
0bm1_10_001_0000000000000000000000000000000000000000000000000000011
Identificador d'una Neurona:

Xarxa 1, Capa 3, Neurona 42 (Net-1 -> Layer-3 -> N-42).
Codificació:
Copia el codi
0bm1_10_010_000000000000000000000000000000000000000000000000101010
Referència a una Sinapsi:

Sinapsi entre N-42 i N-84 (Net-1 -> Layer-3 -> Syn-1).
Codificació:
Copia el codi
0bm1_10_011_0000000000000000000000000000000000000000000000000000001
Espai Jeràrquic i Flexibilitat
Espai Jeràrquic (57 bits):

Amb 57 bits, es poden identificar fins a 2*10<sup>57</sup> − 1 elements en cada categoria.
Això cobreix xarxes molt grans amb múltiples capes, neurones i sinapsis.
Tipus i Extensions:

Els 3 bits dedicats al tipus d'element permeten ampliar fàcilment el sistema en el futur.
Subgrup B.4 Amb Grup D
Ara, el subgrup B.4 pot utilitzar referències explícites als identificadors del Grup D per especificar connexions o dependències amb relacions clares.
Exemple:

Referència des de N-42 a N-84:
B.4 apunta als identificadors D -> N-42 i D -> N-84.


### **Group E (extensió F64):** 0b?1 + '11'g (grup E) + 60?
- **Estructura**: 2b '?1' + '11'g grup E + 60 bits '?' meta valors.
- **Funció**: 
    - Aquest grup només està disponible als tipus RangeF64 i RangeU64 perquè no té espai als tipus de 32 bits.
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
