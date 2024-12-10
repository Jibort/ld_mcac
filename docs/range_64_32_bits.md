# Range64 Bit Layout

## Format float64 IEEE 754 estandard
```
Byte    |           7           |           6           |           5           |           4           | 3..0
Bit     |63 62 61 60|59 58 57 56|55 54 53 52|51 50 49 48|47 46 45 44|43 42 41 40|39 38 37 36|35 34 33 32| ...
Value   |S  e  e  e |e  e  e  e |e  e  e  e |m  m  m  m |m  m  m  x |m  m  m  m |m  m  m  m |m  m  m  m | ...

on S: Flag de signe: 0 = positiu, 1 = negatiu
e's:  Bits d'exponent
m's:  Bits de mantissa
```

## Grup A: Valors regulars dins el rang [-1.0, +1.0] ('1|11')
```
Byte    |           7           |           6           |           5           |           4           | 3..0
Bit     |63 62 61 60|59 58 57 56|55 54 53 52|51 50 49 48|47 46 45 44|43 42 41 40|39 38 37 36|35 34 33 32| ...
Value   |S  1  1  1 |1  1  1  1 |1  1  1  E |m  m  m  m |m  m  m  x |m  m  m  m |m  m  m  m |m  m  m  m | ...

S: Flag de signe
E: +1_022 (en E = 0) o +1_023 -8 on E = 1): 111_1111_111E
m: Bits de mantissa
b[62]: Només és grup A si 62 = '1'
```

## Groups B-C-D-E (altres tipus d'informació) ('0|gg')
```
Byte    |           7           |           6           |           5           |           4           | 3..0
Bit     |63 62 61 60|59 58 57 56|55 54 53 52|51 50 49 48|47 46 45 44|43 42 41 40|39 38 37 36|35 34 33 32| ...
Value   |S? 0  g  g |?  ?  ?  ? |?  ?  ?  ? |?  ?  ?  ? |?  ?  ?  ? |?  ?  ?  ? |?  ?  ?  ? |?  ?  ?  ? | ...

S?: Flag de signe per casos especials com +inf, -inf, +sat i -sat
    En la resta de casos es comporta com el bit 60 '?' de dades i/o paràmetres
?:  59 o 60 bits de dades i/o paràmetres
b[62]: Només és '0' per als grups B-C-D-E
```

### Grup B ('0|00')
```
Byte    |           7           |           6           |           5           |           4           | 3..0
Bit     |63 62 61 60|59 58 57 56|55 54 53 52|51 50 49 48|47 46 45 44|43 42 41 40|39 38 37 36|35 34 33 32| ...
Value   |S? 0  0  0 |sg sg sg sg|?  ?  ?  ? |?  ?  ?  ? |?  ?  ?  ? |?  ?  ?  ? |?  ?  ?  ? |?  ?  ?  ? | ...

S?: Flag de signe per casos especials com +inf, -inf, +sat i -sat
    En la resta de casos es comporta com el bit 60 '?' de dades i/o paràmetres
sg: Identificador de subgrup (0000-0111: 7 subgrups)
?:  59 o 60 bits de dades i/o paràmetres
b[62]: Només és '0' per als grups B-C-D-E
```

#### Grup B.1 NULL & ANY ('0|01'+'0000')
```
Byte    |           7           |           6           |           5           |           4           | 3..0
Bit     |63 62 61 60|59 58 57 56|55 54 53 52|51 50 49 48|47 46 45 44|43 42 41 40|39 38 37 36|35 34 33 32| ...
Value   |na 0  0  0 |0  0  0  0 |?  ?  ?  ? |?  ?  ?  ? |?  ?  ?  ? |?  ?  ?  ? |?  ?  ?  ? |?  ?  ?  ? | ...

na: nul = 0, any = 1
?:  56 bits sense un significat assignat, per ara
```

#### Grup B.2 PADs ('0|01'+'0001')
```
Byte    |           7           |           6           |           5           |           4           | 3..0
Bit     |63 62 61 60|59 58 57 56|55 54 53 52|51 50 49 48|47 46 45 44|43 42 41 40|39 38 37 36|35 34 33 32| ...
Value   |na 0  0  0 |0  0  0  1 |?  ?  ?  ? |?  ?  ?  ? |?  ?  ?  ? |?  ?  ?  ? |?  ?  ?  ? |?  ?  ?  ? | ...

na: nul = 0, any = 1
?:  56 bits sense un significat assignat, per ara
```


S: Sign bit o incorporat al tipus 'd' en alguns subgrups.
01: Group B identifier
s: Subgroup identifier (0000-0111)
p: Parameter bits
d: Data bits

Subgroups:
0000: NULL & ANY
0001: PAD values
0010: Error values
0011: Sequence markers
0100: Separators
0101: Controls
0110: INF values
0111: Reserved
```

## Group C: Token Values ('10')
```
Byte    |           7           |           6           |           5           |           4           |
Bit     |63 62 61 60|59 58 57 56|55 54 53 52|51 50 49 48|47 46 45 44|43 42 41 40|39 38 37 36|35 34 33 32|
Value   |t  1  0  t |t  t  t  t |t  t  t  t |t  t  t  t |t  t  t  t |t  t  t  t |t  t  t  t |t  t  t  t |

        |           3           |           2           |           1           |           0           |
Bit     |31 30 29 28|27 26 25 24|23 22 21 20|19 18 17 16|15 14 13 12|11 10 09 08|07 06 05 04|03 02 01 00|
Value   |t  t  t  t |t  t  t  t |t  t  t  t |t  t  t  t |t  t  t  t |t  t  t  t |t  t  t  t |t  t  t  t |

S: Sign bit
10: Group C identifier
t: Token identifier bits (allows for > 2^50 tokens)
```

## Group D: Saturated Values ('11')
```
Byte    |           7           |           6           |           5           |           4           |
Bit     |63 62 61 60|59 58 57 56|55 54 53 52|51 50 49 48|47 46 45 44|43 42 41 40|39 38 37 36|35 34 33 32|
Value   |S  1  1  j |j  j  j  j |j  j  j  j |j  j  j  j |j  j  j  j |j  j  j  j |j  j  j  j |j  j  j  j |

        |           3           |           2           |           1           |           0           |
Bit     |31 30 29 28|27 26 25 24|23 22 21 20|19 18 17 16|15 14 13 12|11 10 09 08|07 06 05 04|03 02 01 00|
Value   |j  j  j  j |j  j  j  j |j  j  j  j |j  j  j  j |j  j  j  j |j  j  j  j |j  j  j  j |j  j  j  j |

S: Sign bit (0 = positive saturation, 1 = negative saturation)
11: Group D identifier
j: All bits set to 1 for saturation
```

### Examples:

```
R64_ZERO     = 0b0_00_00000_00000000_00000000_00000000_00000000_00000000_00000000_00000000
R64_NULL     = 0b0_01_0000_00000000_00000000_00000000_00000000_00000000_00000000_00000000
R64_PAD_SPACE= 0b0_01_0001_00000000_00100000_00000000_00000000_00000000_00000000_00000000
R64_TOKEN_MIN= 0b0_10_00000_00000000_00000000_00000000_00000000_00000000_00000000_00000000
R64_SAT_POS  = 0b0_11_11111_11111111_11111111_11111111_11111111_11111111_11111111_11111111
```