
# Descripció de les Operacions de RangeF64

Aquest document detalla les operacions suportades pel tipus `RangeF64` i les seves combinacions entre els diferents grups i subgrups. També inclou la gestió d'errors associats.

---

## Funcions Generals

### `Add`

#### Descripció:
Permet sumar dos valors `RangeF64` amb semàntiques que depenen del grup dels operadors.

#### Especificacions:
- **Cas 1**: Si `sF64` pertany al **Grup A** i `pF64` és un **percentatge** (Grup C):
  - El percentatge escala proporcionalment el valor del **Grup A**.
  - **Fórmula**: 
    \[
    rF64 = sF64 	imes (1 + pF64)
    \]

  **Exemple:**
  - `sF64 = 0.5` (Grup A)
  - `pF64 = 10%` (Grup C)
  - Resultat: \( 0.5 	imes (1 + 0.1) = 0.55 \)

- **Cas 2**: Si `sF64` és un **percentatge** (Grup C) i `pF64` pertany al **Grup A** o és un altre **percentatge**:
  - Es realitza una suma literal dels valors numèrics dels percentatges.
  - **Fórmula**:
    \[
    rF64 = sF64 + pF64
    \]

  **Exemple:**
  - `sF64 = 20%` (Grup C)
  - `pF64 = 30%` (Grup C)
  - Resultat: \( 0.2 + 0.3 = 0.5 \) (50%)

#### Validacions:
- Si els grups no són compatibles, la funció retorna un `RangeF64` configurat com un error del Grup B (veure secció d'errors).
- Funcions auxiliars:
  - `IsGroupA(RangeF64) -> bool`
  - `IsPercentage(RangeF64) -> bool`
  - `ToFloat(RangeF64) -> float64`

---

## Errors en les Operacions

Els errors detectats durant les operacions amb `RangeF64` es representen com valors del Grup B (subgrup `00`) amb codis específics i informació addicional.

### Format d'Error
| Camp          | Descripció                                    |
|---------------|----------------------------------------------|
| **Grup (B)**  | Identifica que el valor és un error (`0bS1`). |
| **Codi d'Error** | Codi numèric únic per al tipus d'error.    |
| **Metadades** | Informació addicional per descriure l'error. |

### Codis d'Error Definits
| Codi               | Descripció                                  |
|--------------------|--------------------------------------------|
| **ERR_INVALID_ADD_OPERATION** | Operació `Add` no vàlida entre grups incompatibles. |

### Exemples
- **Error Add**:
  - Si `sF64` és del Grup A i `pF64` és un símbol (Grup B):
    - Codi: `ERR_INVALID_ADD_OPERATION`.
    - Metadades: Conté els valors d'entrada per depuració.

---

## Notes Generals
1. Les operacions són dissenyades per garantir consistència semàntica entre els grups.
2. Es recomana validar els grups abans d'executar qualsevol operació.
3. La xarxa pot processar els errors com a valors natius per facilitar el diagnòstic i depuració.
