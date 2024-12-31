
# Operacions Binàries Matemàtiques en RangeF64

Aquest document detalla les operacions binàries matemàtiques (`+`, `-`, `*`, `/`) per al model **RangeF64**. S'hi inclouen les taules de combinació i les explicacions dels casos d'error identificats com **B.4.x**.

---

## Operació: `+`

|   +   | A.1  | A.2  | A.3  | B.1   | B.2.P  | B.2.N  | B.2.C  | B.2.R  | B.2.E  | B.2.K  | B.3   | B.4  | C   | D   | E   |
|:-----:|:----:|:----:|:----:|:-----:|:------:|:------:|:------:|:------:|:------:|:------:|:-----:|:----:|:---:|:---:|:---:|
|  A.1  | A.1  | A.1  | A.1  |  B.4.a |   A.1  |   A.1  |   B.2.C |   B.2.R |   B.2.E |   B.4.b |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  A.2  | A.1  | A.2  | A.2  |  B.4.a |   A.2  |   A.2  |   B.2.C |   B.2.R |   B.2.E |   B.4.b |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  A.3  | A.1  | A.2  | A.3  |  B.4.a |   A.3  |   A.3  |   B.2.C |   B.2.R |   B.2.E |   B.4.b |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  B.1  |  B.4.a |  B.4.a |  B.4.a |  B.1  |   B.4.a |   B.4.a |   B.4.a |   B.4.a |   B.4.a |   B.4.a |  B.4.a | B.4  | B.4.a | B.4.a | B.4.a |
| B.2.P |  A.1  | A.2  | A.3  |  B.4.a |   B.2.P |   B.2.P |   B.2.C |   B.2.R |   B.2.E |   B.2.K |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.N |  A.1  | A.2  | A.3  |  B.4.a |   B.2.P |   B.2.N |   B.2.C |   B.2.R |   B.2.E |   B.2.K |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.C |  B.2.C | B.2.C | B.2.C |  B.4.a |   B.2.C |   B.2.C |   B.2.C |   B.2.C |   B.2.C |   B.2.C |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.R |  A.1  | A.2  | A.3  |  B.4.a |   B.2.R |   B.2.R |   B.2.C |   B.2.R |   B.2.R |   B.2.R |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.E |  A.1  | A.2  | A.3  |  B.4.a |   B.2.E |   B.2.E |   B.2.C |   B.2.R |   B.2.E |   B.2.E |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
| B.2.K |  B.4.b | B.4.b | B.4.b |  B.4.a |   B.2.K |   B.2.K |   B.2.K |   B.2.K |   B.2.K |   B.2.K |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  B.3  |  B.3  |  B.3  |  B.3  |  B.4.a |   B.3  |   B.3  |   B.3  |   B.3  |   B.3  |   B.3  |  B.3  | B.4  | B.4.c | B.4.c | B.4.d |
|  B.4  |  B.4  |  B.4  |  B.4  |  B.4  |   B.4  |   B.4  |   B.4  |   B.4  |   B.4  |   B.4  |  B.4  |  B.4 | B.4  | B.4  | B.4 |
|   C   |  B.4.c |  B.4.c |  B.4.c |  B.4.a |   B.4.c |   B.4.c |   B.4.c |   B.4.c |   B.4.c |   B.4.c |  B.4  | B.4  |  C   | B.4.c | B.4.d |
|   D   |  B.4.c |  B.4.c |  B.4.c |  B.4.a |   B.4.c |   B.4.c |   B.4.c |   B.4.c |   B.4.c |   B.4.c |  B.4  | B.4  | B.4.c |  D   | B.4.d |
|   E   |  B.4.d |  B.4.d |  B.4.d |  B.4.a |   B.4.d |   B.4.d |   B.4.d |   B.4.d |   B.4.d |   B.4.d |  B.4  | B.4  | B.4.d | B.4.d |  E   |

---

## Operació: `-`

|   -   | A.1  | A.2  | A.3  | B.1   | B.2.P  | B.2.N  | B.2.C  | B.2.R  | B.2.E  | B.2.K  | B.3   | B.4  | C   | D   | E   |
|:-----:|:----:|:----:|:----:|:-----:|:------:|:------:|:------:|:------:|:------:|:------:|:-----:|:----:|:---:|:---:|:---:|
| **(Mateixa taula que +)** |

---

## Operació: `*`

|   *   | A.1  | A.2  | A.3  | B.1   | B.2.P  | B.2.N  | B.2.C  | B.2.R  | B.2.E  | B.2.K  | B.3   | B.4  | C   | D   | E   |
|:-----:|:----:|:----:|:----:|:-----:|:------:|:------:|:------:|:------:|:------:|:------:|:-----:|:----:|:---:|:---:|:---:|
| **(Mateixa taula que +)** |

---

## Operació: `/`

|   /   | A.1  | A.2  | A.3  | B.1   | B.2.P  | B.2.N  | B.2.C  | B.2.R  | B.2.E  | B.2.K  | B.3   | B.4  | C   | D   | E   |
|:-----:|:----:|:----:|:----:|:-----:|:------:|:------:|:------:|:------:|:------:|:------:|:-----:|:----:|:---:|:---:|:---:|
| **(Mateixa taula que +)** |

---

### **Anotacions sobre els errors**

#### **B.4.a: Incompatibilitat semàntica**
- Operacions amb símbols (\( B.1 \)), ja que no tenen cap significat numèric.

#### **B.4.b: Incompatibilitat amb estats físics**
- Tecles auxiliars o botons de ratolí (\( B.2.K \)).

#### **B.4.c: Ambigüitat en metareferències o tokens**
- Operacions amb tokens (\( C \)) o metareferències (\( D \)).

#### **B.4.d: Extensions desconegudes**
- Extensions (\( E \)) que no tenen comportaments definits.

---

Aquest document cobreix les regles per a totes les operacions binàries matemàtiques i identifica els casos específics d'error amb detalls sobre el seu significat i la seva causa.
