# Distribució de bits en Range64

Aquest document descriu la distribució de bits en el tipus `Range64`, incloent tots els grups i subgrups d'extensió definits.

## Grup A (Valors normals, infinits i NaN)

``` monospace
  B7       B6        B5       B4       B3       B2       B1       B0
  Seeeeeee-eeeeemmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm
  ||            ⌂ 52b de dades segons subgrup
  |⌂ 3b Exponent IEEE 754
  ⌂ 1b de signe
```

### Subgrup A.1: Valors subnormals

``` monospace
B7       B6        B5       B4       B3       B2       B1       B0
S0000000-00000mmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm
```

- **Descripció:** Valors extraordinaris propers a zero.
- **Nota**: Els valors subnormals NO inclouen un 1 implícit a la mantissa.
- **Camps:**
  - `S`: 1 bit de signe.
  - `0000000-0000`: Exponent IEEE 754 pels valors extraordinàriament propers a 0 (rang [−2<sup>−1022</sup>, +2<sup>−1022</sup>]).
  - `m`: 52 bits per a la mantissa.

### Subgrup A.2: Valors normals molt propers a 0

``` monospace
B7       B6       B5       B4       B3       B2       B1       B0
S0000000-0001mmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm
```

- **Descripció:** Inclou els valors en el rang principal que són normals però molt propers a 0.
- **Camps:** 
  - `S`: 1 bit de signe.
  - `0000000-0001 (1d)`: Exponent IEEE 754 pels valors normals molt propers a 0 (rangs [−1.0,+1.0] i [-2π,+2π]).
  - `m`: 52 bits per a la mantissa.

<p style="page-break-before: always;"></p>

### Subgrup A.3: Valors normals NO molt propers a 0

``` monospace
B7       B6       B5       B4       B3       B2       B1       B0
Sxxxxxxx-xx10mmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm
```

- **Descripció:** Inclou els valors en el rang principal que són normals però no molt propers a 0.
- **Camps:**
  - `S`: 1 bit de signe.
  - `xxxxxxx-xx10 (0000000-0010 a 1111111-11110 (-1.021d a +1.023d))`: Exponents IEEE 754 pels valors normals NO molt propers a 0, pels rangs [−1.0,+1.0] i [-2π,+2π].
  - `m`: 52 bits per a la mantissa.

### Subgrup A.4: Infinits

``` monospace
B7       B6       B5       B4       B3       B2       B1       B0
S1111111-11110000-00000000-00000000-00000000-00000000-00000000-00000000
```

- **Descripció:** Representa ±Inf.
- **Camps:**
  - `S`: 1 bit de signe (0: +inf, 1: -inf)
  - `1111111-11111 (2.047)`: Exponent IEEE 754 per a ambdós valors.
  - `000..0`: 52b a zero

### Subgrup A.5: NaN

``` monospace
B7       B6       B5       B4       B3       B2       B1       B0
S1111111-11111mmm-mmmmmmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm
```

- **Descripció:** Representa NaN (Not a Number).
- **Camps:**
  - `S`: 1 bit de signe.
  - `1m...m`: qualsevol combinació de bits que almenys mantingui un bit a 1.

---

<p style="page-break-before: always;"></p>

## Grup B (Símbols i Informació addicional)

``` monospace
B7       B6       B5       B4       B3       B2       B1       B0
s100ggdd-dddddddd-dddddddd-dddddddd-dddddddd-dddddddd-dddddddd-dddddddd
|   | ⌂ 58b de dades segons el subgrup
||  ⌂ Identificador de subgrup
|⌂ Identificació del grup B (100)
⌂ Depenent del subgrup: signe, inici o final de padding o tecla escape
```

### Subgrup B.1: Símbols

``` monospace
B7       B6       B5       B4       B3       B2       B1       B0
s100ggdd-dddddddd-dddddddd-dddddddd-dddddddd-dddddddd-dddddddd-dddddddd
```

- **Descripció:** Representa símbols UTF-32.
- **Camps:**
  - `?`: bits reservats per a futures extensions.
  - `100`: 3 bits id. del Grup B.
  - `00`: 2 bits id. del Subgrup B.1
  - `c`: 32 bits per identificar el símbol (estàndard UTF-32).

### Subgrup B.2: Paddings, Nul, Coordenades, i Tecles

``` monospace
B7       B6       B5       B4       B3       B2       B1       B0
t10001tt-dddddddd-dddddddd-dddddddd-dddddddd-dddddddd-dddddddd-dddddddd
```

**Camps generals:**

- `d`: 59b definits segons el tipus 't-tt'.
- `100`: Identificador del Grup B.
- `01`: Identificador del Subgrup B.2.
- `t-tt`: Assigna els següents valors:
  - `000`: Padding inicial (espai abans d'un bloc de dades).
  - `001`: Padding general (espai dins o entre blocs de dades).
  - `010`: Padding final (espai després d'un bloc de dades).
  - `011`: Valor Nul.
  - `100`: Coordenades.
  - `101`: Rang de percentatges.
  - `110`: Percentatge.
  - `111`: Estat de les tecles auxilliars.

#### Subgrup B.2.P: Paddings ('000' a '010'):

``` monospace
B7       B6       B5       B4       B3       B2       B1       B0
t10001tt-????????-????????-????????-????????-????????-????????-????????
```

- **Descripció:** Identifica els diferents tipus de padding.
- **Camps:**
  - `?`: bits reservats per a futures extensions.
  - `100`: Identificador del Grup B.
  - `01`: Identificador del Subgrup B.2.
  - `t-tt`: Tipus
    - `000`: Padding inicial (espai abans d'un bloc de dades).
    - `001`: Padding general (espai dins o entre blocs de dades).
    - `010`: Padding final (espai després d'un bloc de dades).

#### Subgrup B.2.N: Valor Nul ('011')

``` monospace
B7       B6       B5       B4       B3       B2       B1       B0
t10001tt-????????-????????-????????-????????-????????-????????-????????
```

- **Descripció:** Identifica els diferents tipus de padding.
- **Camps:**
  - `?`: bits reservats per a futures extensions.
  - `100`: Identificador del Grup B.
  - `01`: Identificador del Subgrup B.2.
  - `t-tt`: Tipus
    - `011`: Valor Nul.

#### Subgrup B.2.C: Coordenades ('100')

``` monospace
B7       B6       B5       B4       B3       B2       B1       B0
t10001tt-xxxxxxxx-xxxxxxxx-xxxxxxxx-xxxxyyyy-yyyyyyyy-yyyyyyyy-yyyyyyyy
```

- **Descripció:** Identifica les coordenades {x,y} en un pla.
  - `100`: Identificador del Grup B.
  - `01`: Identificador del Subgrup B.2.
  - `t-tt`: Tipus
    - `100`: Coordenades.
  - `x`: 1b+27b per a la coordenada _x_ (1b signe (0: +, 1: -),  27b mantissa).
  - `y`: 1b+27b per a la coordenada _y_ (1b signe (0: +, 1: -),  27b mantissa).

#### Subgrup B.2.R: Rang de Percentatges

``` monospace
B7       B6       B5       B4       B3       B2       B1       B0
t10001tt-vvvvvvvv-vvvvvvvv-vvvvvvvv-vvvvwwww-wwwwwwww-wwwwwwww-wwwwwwww
```

- **Descripció:** Defineix un rang de percentatges.
- **Camps:**
  - `S`: 1b de signe (0: positiu, 1: negatiu).
  - `100`: Identificador del Grup B.
  - `01`: Identificador del Subgrup B.2.
  - `t-tt`: Tipus
    - `101`: Rang de percentatges.
  - `v`: 28b per a definir el percentatge mínim del rang (1b signe, 27b mantissa).
  - `w`: 28b per a definir el percentatge màxim del rang (1b signe, 27b mantissa).

#### Subgrup B.2.E: Percentatge

``` monospace
B7       B6       B5       B4       B3       B2       B1       B0       
t10001tt-???Svvvv-vvvvvvvv-vvvvvvvv-vvvvvvvv-vvvvvvvv-vvvvvvvv-vvvvvvvv
```

- **Descripció:** Valor percentual en un rang [-100.0, +100.0]
- **Camps:**
  - `100`: Identificador del Grup B.
  - `01`: Identificador del Subgrup B.2.
  - `t-tt`: Tipus
    - `110`: Percentatge.
  - `???`: Reservats per a futures extensions.
  - `S`: 1b de signe (0: positiu, 1: negatiu).
  - `v`: 52b de mantissa dins el conjunt de valors del Subgrup A.3.
  - **nota:** La mantissa és igual, bit a bit, al valor segons les definicions del Subgrup A.3.

<p style="page-break-before: always;"></p>

#### Subgrup B.2.K: Tecles auxiliars i mouse

``` monospace
B7       B6       B5       B4       B3       B2       B1       B0       
t10001tt-abcdefgh-ijklmnop-qrstuvwx-yz12ABCD-EFGHIJKL-MNOPQRST-UVWXYZ34
```

- **Descripció:** Identifica l'estat de les tecles auxiliars del teclat (0: NO pressa, 1: pressa) i dels botons del ratolí.
  - `100`: Identificador del Grup B.
  - `01`: Identificador del Subgrup B.2.
  - `t-tt`: Tipus
    - `111`: Estat de les tecles especials.
- **Camps:**

<div style="display: flex; gap: 5px;">
  <div style="flex: 1;">
    <b>Bloc 1</b><br>
    `a`: Tecla F1<br>
    `b`: Tecla F2<br>
    `c`: Tecla F3<br>
    `d`: Tecla F4<br>
    `e`: Tecla F5<br>
    `f`: Tecla F6<br>
    `g`: Tecla F7<br>
    `h`: Tecla F8<br>
    `i`: Tecla F9<br>
    `j`: Tecla F10<br>
    `k`: Tecla F11<br>
  </div>
  <div style="flex: 1;">
    <b>Bloc 2</b><br>
    `k`: Tecla F11<br>
    `l`: Tecla F12<br>
    `m`: Tecla Tabulador<br>
    `n`: Tecla Bloqueig de Majúscules<br>
    `o`: Tecla Shift Esquerra<br>
    `p`: Tecla Shift Dreta<br>
    `q`: Tecla Control Esquerra<br>
    `r`: Tecla Control Dreta<br>
    `s`: Tecla Funció<br>
    `t`: Tecla Windows<br>
    `u`: Tecla OSX<br>
  </div>
  <div style="flex: 1;">
    <b>Bloc 3</b><br>
    `v`: Tecla Alt<br>
    `w`: Tecla Alt Graph.<br>
    `x`: Tecla Esborrar<br>
    `y`: Tecla Suprimir<br>
    `z`: Tecla Escape<br>
    `1`: Tecla Botó Esquerre del Ratolí<br>
    `2`: Tecla Botó Central del Ratolí<br>
    `A`: Tecla Botó Dret del Ratolí<br>
    `B..Z34`: Bit per a futures expansions.<br>
  </div>
</div>


### Subgrup B.3: Saturació

``` monospace
B7       B6       B5       B4       B3       B2       B1       B0       
S10010tt-????mmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm-mmmmmmmm
```

- **Descripció:** Indica si un valor del grup A està saturat.
  - `?`: bits reservats per a futures extensions (nous tipus de saturació, ...).
- **Camps:**
  - `S`: 1 bit de signe.
  - `100`: Identificador del Grup B.
  - `10`: Identificador del Subgrup B.3.
  - `tt`: Assigna els següents valors:
    - `00`: Saturació que segueix el format del subgrup A.1 (valors subnormals). Exemple: `S1000100-00001...`
    - `01`: Saturació que segueix el format del subgrup A.2 (normals molt propers a 0). Exemple: `S1000101-00010...`
    - `10`: Saturació que segueix el format del subgrup A.3 (normals no molt propers a 0). Exemple: `S1000110-00100...`
    - `11`: Saturació que segueix el format del subgrup A.4 (infinits ±Inf). Exemple: `S1000111-01000...`
    - **nota:** No té sentit emular el valor NaN del subgrup A.5.
  - `m`: 52b de mantissa, equivalent al mateix valor no saturat definit en el Grup A.
    - **nota:** La mantissa és igual, bit a bit, al mateix valor no saturat segons les definicions del Grup A.

### Subgrup B.4: Errors

``` monospace
B7       B6       B5       B4       B3       B2       B1       B0       
c10011ee-eeeeeeee-????aaaa-aaaaaaaa-aaaaaaaa-aaaaaaaa-aaaaaaaa-aaaaaaaa
```

- **Descripció:** Representa codis d'error i arguments associats.
  - `?`: Bits reservat per a extensions com:
    - Versions
    - Mòduls
    - i altres
  - `c`: Criticitat:
    - `0`: Error no crític.
    - `1`: Error crític que requereix atura el funcionament del sistema.
  - `100`: Identificador del Grup B.
  - `11`: Identificador del Subgrup B.4.
  - `e`: 10b per al codi d'error (veure la secció _"Desglosament d'Errors"_ posterior).
  - `a`: 48b per a possibles arguments associats, com ara:
    - Identificadors de components afectats.
    - Informació específica de l'error (rang esperat, valor rebut, etc.).
    - Altres detalls contextuals.

---

<p style="page-break-before: always;"></p>

## Grup C (Tokens)

``` monospace
  B7       b6       B5       B4       B3       B2       B1       B0
  ?101cccc-ccffffff-ffffffpp-pppppppp-pppptttt-tttttttt-tttttttt-tttttttt
  ||  |      |            |               ⌂ 28b: identificador únic de token
  ||  |      |            ⌂ 14b: pes relatiu del token dins el context propi
  ||  |      ⌂ 12b: percentatge de fiabilitat (del 0% (000000) al 100% (111111))
  ||  ⌂ 6b: categoria del token
  |⌂ '101': identificació del grup B (101) 
  ⌂ 1b: reservat per a futures ampliacions
```

**Nota important:** Com és literalment impossible en l'espai disponible poder identificar un token amb el seu corresponent embedding fem servir l'identificador únic del registre de la base de dades on quedi enregistrat el token (2<sup>28</sup> = 268.435.456 possibles identificadors únics).

**Exemple de Token:**

- <u>Format complet</u>: `?1010011-00110101-01110000-00000000011001-000000010010011001011111`

* Grup C: (`101`)
* Categoria: 12 (`0011-00`)
* Percentatge de fiabilitat: 87% (`1010111`)
* Pes relatiu: 25 (`0000-0000011001`)
* Identificador únic: 1.234.567 (`000000010010011001011111`)

---

<p style="page-break-before: always;"></p>

## Grup D (Metareferències)

``` monospace
  B7       b6       B5       B4       B3       B2       B1       B0
  x110xxxx-xxxxxxyy-yyyyyyyy-yyyyynnn-nnnnnnnn-nnnnnsss-ssssssss-ssssssss
    - 11b (x's): Identificador de xarxa (2.048 xarxes diferents).
    - `110`: Identificador del Grup D.
    - 15b (y's): Identificador de capa dins la xarxa (32.768 capes diferents).
    - 17b (n's): Identificador de neurona dins la capa (131.072 neurones individuals per capa).
    - 18b (s's): Identificador de sinapsi de sortida de la neurona (262.144) possibles connexions a neurones destí.
```

**Nota important**: 

- La configuració identifica una xarxa si la resta d'identificadors es troben a '0'. 
- La configuració identifica una capa si la neurona i la sinapsi es troben a '0'.
- La configuració identifica una neurona si la sinapsi es troba a '0'.
- Reservar els valors '0' dels identificadors de capa, neurona i sinapsi permet estalviar bits que d'altra forma serien necessaris per tal de determinar si el cas és una identificació de xarxa, de xarxa i capa; de xarxa, capa i neurona; o de xarxa, capa, neurona i sinapsi.

**Rangs:** 

- L'identificador de xarxa té el rang [1, .., 2.047] perquè el 0 està reservat per 'xarxa actual'.
- L'identificador de la capa té el rang [1, .., 32.767] perquè el 0 està reservat per 'capa nul·la'.
- L'identificador de la neurona té el rang [1, .., 131.071] perquè el 0 està reservat per 'neurona nul·la'.
- L'identificador de la sinapsi té el rang [1, .., 262.143] perquè el 0 està reservat per 'sinapsi nul·la'.

**Exemple desglossat d'identificació de xarxa-capa-neurona-sinapsi:** 

- <u>Format complet</u>: `01100000-00010100-00000100-00000000-00100111-00010000-01111000-10000000`

* _Grup D: '_110'
* _Xarxa_: núm. 5 (`0GGG0000-000101`)
* _Capa_: núm. 256 (`00-00000100-00000`)
* _Neurona_: núm. 10.000 (`000-00100111-000100`)
* _Sinapsi_: núm. 123.456 (`00-01111000-10000000`)

**Exemple desglossat d'identificació de xarxa-capa:** 

- <u>Format complet</u>: `01100000-00111000-00001100-00000000-00000000-00000000-00000000-00000000`

* _Grup D: '_110'
* _Xarxa_: núm. 14 (`0ggg0000-001110`), on ggg = Grup D ('110')
* _Capa_: núm. 768 (`00-00001100-00000`)
* _Neurona_: nul·la (`000-00000000-000000`)
* _Sinapsi_: nul·la (`00-00000000-00000000`)

**Exemple desglossat d'identificació de la xarxa actual:**  

- <u>Format complet</u>: `01100000-00000000-00000000-00000000-00000000-00000000-00000000-00000000`  

* _Grup D_: `_110`  
* _Xarxa_: núm. xarxa actual (`0ggg0000-000000`), on ggg = Grup D ('110')
* _Capa_: núm. nul·la (`00-00000000-00000`)
* _Neurona_: nul·la (`000-00000000-000000`)
* _Sinapsi_: nul·la (`00-00000000-00000000`)

---

<p style="page-break-before: always;"></p>

## Grup E (Extensions futures en float64)

``` monospace
  B7       b6       B5       B4       B3       B2       B1       B0
  ?1110???-????????-????????-????????-????????-????????-????????-????????
    - '?': Bits reservats per a futures extensions exclusives de formats de 64b.
    - `111`: Identificador del Grup E.
    - `_1110___`: Identificador del Grup E, evitant coincidències amb exponents que corresponen al Grup A. 
```

### Descripció general

- Grup reservat exclusivament per a posibles usos futurs.
- El dígit '0' posterior a la identificació del grup és imprescindible per a assegurar que no solapem algun exponent necessari pel Grup A.

### Exemples d'extensions futures possibles:

- Formateig avançat de dades, com codificació de metadades complexes.
- Suport per a nous estàndards de representació numèrica o simbòlica.
- Integració amb funcionalitats quàntiques en futures arquitectures.

---

## Operacions binàries entre subgrups

No tots els parell d'instàncies de qualsevol grup o subgrup tenen sentit a l'hora d'obtenir resultats per a les diferents operacions binàries que desenvolupem.

Per exemple, sumar una instància de Subgrup A.3 amb una de Grup C, o una del Grup D i una del Subgrup B.2.E, no té un sentit lògic definit. Per això necessitem de documentar adequadament com respondre a les diferents combinacions de forma explícita per a totes les operacions binàries.

### Operació: `+`

|   +   |  A.1  |  A.2  |  A.3  |  B.1  | B.2.P | B.2.N | B.2.C | B.2.R | B.2.E | B.2.K |  B.3  | B.4  |   C   |   D   |   E   |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :--: | :---: | :---: | :---: |
|  A.1  |  A.1  |  A.1  |  A.1  | B.4.a |  A.1  |  A.1  | B.2.C | B.2.R | B.2.E | B.4.b |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  A.2  |  A.1  |  A.2  |  A.2  | B.4.a |  A.2  |  A.2  | B.2.C | B.2.R | B.2.E | B.4.b |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  A.3  |  A.1  |  A.2  |  A.3  | B.4.a |  A.3  |  A.3  | B.2.C | B.2.R | B.2.E | B.4.b |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  B.1  | B.4.a | B.4.a | B.4.a |  B.1  | B.4.a | B.4.a | B.4.a | B.4.a | B.4.a | B.4.a | B.4.a | B.4  | B.4.a | B.4.a | B.4.a |
| B.2.P |  A.1  |  A.2  |  A.3  | B.4.a | B.2.P | B.2.P | B.2.C | B.2.R | B.2.E | B.2.K |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.N |  A.1  |  A.2  |  A.3  | B.4.a | B.2.P | B.2.N | B.2.C | B.2.R | B.2.E | B.2.K |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.C | B.2.C | B.2.C | B.2.C | B.4.a | B.2.C | B.2.C | B.2.C | B.2.C | B.2.C | B.2.C |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.R |  A.1  |  A.2  |  A.3  | B.4.a | B.2.R | B.2.R | B.2.C | B.2.R | B.2.R | B.2.R |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.E |  A.1  |  A.2  |  A.3  | B.4.a | B.2.E | B.2.E | B.2.C | B.2.R | B.2.E | B.2.E |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.K | B.4.b | B.4.b | B.4.b | B.4.a | B.2.K | B.2.K | B.2.K | B.2.K | B.2.K | B.2.K |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  B.3  |  B.3  |  B.3  |  B.3  | B.4.a |  B.3  |  B.3  |  B.3  |  B.3  |  B.3  |  B.3  |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  | B.4  |  B.4  |  B.4  |  B.4  |
|   C   | B.4.c | B.4.c | B.4.c | B.4.a | B.4.c | B.4.c | B.4.c | B.4.c | B.4.c | B.4.c |  B.4  | B.4  |   C   | B.4.c | B.4.d |
|   D   | B.4.c | B.4.c | B.4.c | B.4.a | B.4.c | B.4.c | B.4.c | B.4.c | B.4.c | B.4.c |  B.4  | B.4  | B.4.c |   D   | B.4.d |
|   E   | B.4.d | B.4.d | B.4.d | B.4.a | B.4.d | B.4.d | B.4.d | B.4.d | B.4.d | B.4.d |  B.4  | B.4  | B.4.d | B.4.d |   E   |

---

### Operació: `-`

|   -   |  A.1  |  A.2  |  A.3  |  B.1  | B.2.P | B.2.N | B.2.C | B.2.R | B.2.E | B.2.K |  B.3  | B.4  |   C   |   D   |   E   |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :--: | :---: | :---: | :---: |
|  A.1  |  A.1  |  A.1  |  A.1  | B.4.a |  A.1  |  A.1  | B.2.C | B.2.R | B.2.E | B.4.b |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  A.2  |  A.1  |  A.2  |  A.2  | B.4.a |  A.2  |  A.2  | B.2.C | B.2.R | B.2.E | B.4.b |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  A.3  |  A.1  |  A.2  |  A.3  | B.4.a |  A.3  |  A.3  | B.2.C | B.2.R | B.2.E | B.4.b |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  B.1  | B.4.a | B.4.a | B.4.a |  B.1  | B.4.a | B.4.a | B.4.a | B.4.a | B.4.a | B.4.a | B.4.a | B.4  | B.4.a | B.4.a | B.4.a |
| B.2.P |  A.1  |  A.2  |  A.3  | B.4.a | B.2.P | B.2.P | B.2.C | B.2.R | B.2.E | B.2.K |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.N |  A.1  |  A.2  |  A.3  | B.4.a | B.2.P | B.2.N | B.2.C | B.2.R | B.2.E | B.2.K |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.C | B.2.C | B.2.C | B.2.C | B.4.a | B.2.C | B.2.C | B.2.C | B.2.C | B.2.C | B.2.C |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.R |  A.1  |  A.2  |  A.3  | B.4.a | B.2.R | B.2.R | B.2.C | B.2.R | B.2.R | B.2.R |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.E |  A.1  |  A.2  |  A.3  | B.4.a | B.2.E | B.2.E | B.2.C | B.2.R | B.2.E | B.2.E |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.K | B.4.b | B.4.b | B.4.b | B.4.a | B.2.K | B.2.K | B.2.K | B.2.K | B.2.K | B.2.K |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  B.3  |  B.3  |  B.3  |  B.3  | B.4.a |  B.3  |  B.3  |  B.3  |  B.3  |  B.3  |  B.3  |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  | B.4  |  B.4  |  B.4  |  B.4  |
|   C   | B.4.c | B.4.c | B.4.c | B.4.a | B.4.c | B.4.c | B.4.c | B.4.c | B.4.c | B.4.c |  B.4  | B.4  |   C   | B.4.c | B.4.d |
|   D   | B.4.c | B.4.c | B.4.c | B.4.a | B.4.c | B.4.c | B.4.c | B.4.c | B.4.c | B.4.c |  B.4  | B.4  | B.4.c |   D   | B.4.d |
|   E   | B.4.d | B.4.d | B.4.d | B.4.a | B.4.d | B.4.d | B.4.d | B.4.d | B.4.d | B.4.d |  B.4  | B.4  | B.4.d | B.4.d |   E   |

---

### Operació: `*`

|   *   |  A.1  |  A.2  |  A.3  |  B.1  | B.2.P | B.2.N | B.2.C | B.2.R | B.2.E | B.2.K |  B.3  | B.4  |   C   |   D   |   E   |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :--: | :---: | :---: | :---: |
|  A.1  |  A.1  |  A.1  |  A.1  | B.4.a |  A.1  |  A.1  | B.2.C | B.2.R | B.2.E | B.4.b |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  A.2  |  A.1  |  A.2  |  A.2  | B.4.a |  A.2  |  A.2  | B.2.C | B.2.R | B.2.E | B.4.b |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  A.3  |  A.1  |  A.2  |  A.3  | B.4.a |  A.3  |  A.3  | B.2.C | B.2.R | B.2.E | B.4.b |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  B.1  | B.4.a | B.4.a | B.4.a |  B.1  | B.4.a | B.4.a | B.4.a | B.4.a | B.4.a | B.4.a | B.4.a | B.4  | B.4.a | B.4.a | B.4.a |
| B.2.P |  A.1  |  A.2  |  A.3  | B.4.a | B.2.P | B.2.P | B.2.C | B.2.R | B.2.E | B.2.K |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.N |  A.1  |  A.2  |  A.3  | B.4.a | B.2.P | B.2.N | B.2.C | B.2.R | B.2.E | B.2.K |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.C | B.2.C | B.2.C | B.2.C | B.4.a | B.2.C | B.2.C | B.2.C | B.2.C | B.2.C | B.2.C |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.R |  A.1  |  A.2  |  A.3  | B.4.a | B.2.R | B.2.R | B.2.C | B.2.R | B.2.R | B.2.R |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.E |  A.1  |  A.2  |  A.3  | B.4.a | B.2.E | B.2.E | B.2.C | B.2.R | B.2.E | B.2.E |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.K | B.4.b | B.4.b | B.4.b | B.4.a | B.2.K | B.2.K | B.2.K | B.2.K | B.2.K | B.2.K |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  B.3  |  B.3  |  B.3  |  B.3  | B.4.a |  B.3  |  B.3  |  B.3  |  B.3  |  B.3  |  B.3  |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  | B.4  |  B.4  |  B.4  |  B.4  |
|   C   | B.4.c | B.4.c | B.4.c | B.4.a | B.4.c | B.4.c | B.4.c | B.4.c | B.4.c | B.4.c |  B.4  | B.4  |   C   | B.4.c | B.4.d |
|   D   | B.4.c | B.4.c | B.4.c | B.4.a | B.4.c | B.4.c | B.4.c | B.4.c | B.4.c | B.4.c |  B.4  | B.4  | B.4.c |   D   | B.4.d |
|   E   | B.4.d | B.4.d | B.4.d | B.4.a | B.4.d | B.4.d | B.4.d | B.4.d | B.4.d | B.4.d |  B.4  | B.4  | B.4.d | B.4.d |   E   |

---

<p style="page-break-before: always;"></p>

### Operació: `/`

|   /   |  A.1  |  A.2  |  A.3  |  B.1  | B.2.P | B.2.N | B.2.C | B.2.R | B.2.E | B.2.K |  B.3  | B.4  |   C   |   D   |   E   |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :--: | :---: | :---: | :---: |
|  A.1  |  A.1  |  A.1  |  A.1  | B.4.a |  A.1  |  A.1  | B.2.C | B.2.R | B.2.E | B.4.b |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  A.2  |  A.1  |  A.2  |  A.2  | B.4.a |  A.2  |  A.2  | B.2.C | B.2.R | B.2.E | B.4.b |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  A.3  |  A.1  |  A.2  |  A.3  | B.4.a |  A.3  |  A.3  | B.2.C | B.2.R | B.2.E | B.4.b |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  B.1  | B.4.a | B.4.a | B.4.a |  B.1  | B.4.a | B.4.a | B.4.a | B.4.a | B.4.a | B.4.a | B.4.a | B.4  | B.4.a | B.4.a | B.4.a |
| B.2.P |  A.1  |  A.2  |  A.3  | B.4.a | B.2.P | B.2.P | B.2.C | B.2.R | B.2.E | B.2.K |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.N |  A.1  |  A.2  |  A.3  | B.4.a | B.2.P | B.2.N | B.2.C | B.2.R | B.2.E | B.2.K |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.C | B.2.C | B.2.C | B.2.C | B.4.a | B.2.C | B.2.C | B.2.C | B.2.C | B.2.C | B.2.C |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.R |  A.1  |  A.2  |  A.3  | B.4.a | B.2.R | B.2.R | B.2.C | B.2.R | B.2.R | B.2.R |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.E |  A.1  |  A.2  |  A.3  | B.4.a | B.2.E | B.2.E | B.2.C | B.2.R | B.2.E | B.2.E |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.K | B.4.b | B.4.b | B.4.b | B.4.a | B.2.K | B.2.K | B.2.K | B.2.K | B.2.K | B.2.K |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  B.3  |  B.3  |  B.3  |  B.3  | B.4.a |  B.3  |  B.3  |  B.3  |  B.3  |  B.3  |  B.3  |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |  B.4  | B.4  |  B.4  |  B.4  |  B.4  |
|   C   | B.4.c | B.4.c | B.4.c | B.4.a | B.4.c | B.4.c | B.4.c | B.4.c | B.4.c | B.4.c |  B.4  | B.4  |   C   | B.4.c | B.4.d |
|   D   | B.4.c | B.4.c | B.4.c | B.4.a | B.4.c | B.4.c | B.4.c | B.4.c | B.4.c | B.4.c |  B.4  | B.4  | B.4.c |   D   | B.4.d |
|   E   | B.4.d | B.4.d | B.4.d | B.4.a | B.4.d | B.4.d | B.4.d | B.4.d | B.4.d | B.4.d |  B.4  | B.4  | B.4.d | B.4.d |   E   |

---

### **Anotacions sobre els errors**

#### **B.4.a: Incompatibilitat semàntica**

- Operacions amb símbols (B.1), ja que no tenen cap significat numèric.

#### **B.4.b: Incompatibilitat amb estats físics**

- Tecles auxiliars o botons de ratolí (B.2.K).

#### **B.4.c: Ambigüitat en metareferències o tokens**

- Operacions amb tokens (C) o metareferències (D).

#### **B.4.d: Extensions desconegudes**

- Extensions (E) que no tenen comportaments definits.

---

Aquest document reflecteix l'assignació final de bits per a cada grup i subgrup del tipus `Range64`. Si hi ha ajustaments o noves extensions, caldrà actualitzar aquest document amb les modificacions corresponents.