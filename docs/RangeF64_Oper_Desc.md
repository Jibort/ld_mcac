# Descripció de les Operacions de RangeF64

Aquest document detalla les operacions suportades pel tipus `RangeF64` i les seves combinacions entre els diferents grups i subgrups.

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
- Si els grups no són compatibles, la funció retorna un error.
- Funcions auxiliars:
  - `IsGroupA(RangeF64) -> bool`
  - `IsPercentage(RangeF64) -> bool`
  - `ToFloat(RangeF64) -> float64`

---

## Funcions Previstes

Aquest document es mantindrà actualitzat amb noves operacions i semàntiques per a cada grup de `RangeF64`.

- **Pendents d'integració**:
  - Operacions de multiplicació (`Mul`) i divisió (`Div`) amb percentatges.
  - Suport per combinacions de valors especials del Grup B.

---

## Notes Generals
1. Les operacions són dissenyades per garantir consistència semàntica entre els grups.
2. Es recomana validar els grups abans d'executar qualsevol operació.

