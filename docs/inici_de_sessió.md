# Instruccions per a cada inici de sessió

## Objectiu del Projecte
Desenvolupar una implementació que permeti comprovar empíricament el Model Contextual d'Àmbits Cognitius (MCAC) utilitzant tecnologies de xarxes neuronals i transformers.

## Estratègia
1. Desenvolupar un LLM base des de zero per evitar dependències d'altres models existents
2. Començar amb un problema de sil·labificació en català com a primera aproximació:
   - Permet aprendre conceptes bàsics de xarxes neuronals
   - És verificable i té regles clares
   - Servirà com a base per la tokenització posterior

## Primer Pas: RangeIntf i implementació de RangeF64
Es defineix una interfície RangIntf que definirà les funcions que serviran per l'execució de càlculs dins les xarxes (pesos, bias, ...)

### Implementació de RangeF64
Abans d'abordar la implementació de totes les diferents formes possibles de RangeIntf desenvoluparem i provarem molt detalladament la inmplementació del tipus RangeF64 (->Rang64Intf -> RangeIntf).

Donat que el desenvolupament de les mecàniques de xarxes només dependràn de la interfície RangeIntf, un cop tinguem acabat RangeF64 podrem desenvolupar les xarxes i la resta de tipus en paral·lel.

L'objectiu principal de disposar de tants tipus diferents de representacions de dades és poder fer comparatives de rendiment, tamany, eficàcia i precissió executant xarxes de prova amb tots els tipus.

### Repositori GitHub del projecte
El repositori GitHub del projecte es troba en https://github.com/Jibort/ld_mcac.git

## Estructura del Projecte
```
ld_mcac/
├── cmd/
│   └── main.go
├── docs/
│   └── sessions/  # Sessions de conversa amb Claude abans de començar a treballar amb ChatGPT
├── internal/
│   ├── core/
│   │   └── *.go      # Interfícies, constants i diferents implementacions de tipus
│   ├── neural/       # Components de xarxa neuronal
│   └── syllables/    # Lògica de sil·labificació
├── pkg/
│   └── utils/            # Utilitats generals
├── tests/                 # Tests i benchmarks
├── go.mod
└── README.md
```

### Estructura de 64bits i 32bits
L'estructura de com farem servir tant els valors de 64 bits com els de 32bits es troba definida en el fitxer del repositori localitzat en: https://github.com/Jibort/ld_mcac.git\docs\range_64_32_bits.md

## Convencions de Codi
- Prefix 'p' per paràmetres d'entrada de funcions
- Prefix 'r' per valors de retorn
- Prefix 's' per el tipus en mètodes (source)
- Noms significatius començant amb majúscula
- Variables internes sense prefix

Exemple:
```go
func (sSyn Synapse) Mul(pVal Range64) (rRes Range64) { ... }
```