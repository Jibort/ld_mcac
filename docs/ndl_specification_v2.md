# Neural Description Language (NDL)

El **Neural Description Language (NDL)** és un llenguatge dissenyat per descriure, configurar i gestionar xarxes neuronals. Aquest llenguatge es compon de diversos tipus de fitxers que permeten separar l'estructura, el contingut, la configuració d'entrenament i els punts de control.

---

## Tipus de fitxers del Llenguatge NDL

### 1. `.ndl` - Neural Description Language (Estructura)
El format .ndl s'utilitza per definir l'arquitectura d'una xarxa neuronal, incloent-hi el nombre de capes, el nombre de neurones per capa i les connexions entre elles.

#### Estructura General
Un fitxer .ndl és un document JSON que conté els següents camps:

- _name_: (Obligatori) Nom de la xarxa.
- _description_: (Opcional) Descripció de la xarxa.
- _layers_: (Obligatori) Llista que especifica el nombre de neurones en cada capa.
- _links_: (Obligatori) Llista de connexions entre neurones.

__Exemple__:
```json
{
  "name": "my_network",
  "description": "Descripció opcional",
  "layers": [40, 10, 10, 5],
  "links": [
    {"from": "L:0, N:n", "to": "L:1, N:n"},
    {"from": "L:1, N:n", "to": "L:2, N:n"},
    {"from": "L:y, N:n", "to": "L:y+1, N:n"},
    {"from": "L:4, N:3", "to": "L:2, N:7"},
    {"from": "L:2, N:3", "to": "L:4, N:n"}
  ]
}
```

#### Descripció dels Camps
1. __name__ (obligatori)
_Tipus_: Cadena de text
_Descripció_: Nom identificatiu de la xarxa neuronal.
2. __description__ (opcional)
_Tipus_: Cadena de text
_Descripció_: Descripció opcional que proporciona informació addicional sobre la xarxa.
3. __layers__ (obligatori)
_Tipus_: Llista d'enters
_Descripció_: Cada element de la llista representa el nombre de neurones en la capa corresponent,   començant per la capa d'entrada.
<u>Nota</u>: La longitud de la llista determina el nombre total de capes de la xarxa. Per exemple, "layers": [40, 10, 10, 5] indica una xarxa amb quatre capes: una capa d'entrada amb 40 neurones, dues capes ocultes amb 10 neurones cadascuna i una capa de sortida amb 5 neurones.
4. __links__ (obligatori)
_Tipus_: Llista d'objectes
_Descripció_: Cada objecte defineix una connexió entre neurones, amb els camps:
    - from: Especifica la neurona d'origen.
    - to: Especifica la neurona de destinació.

##### Format per a from i to: "L:x, N:y"
- _L:x_: Indica la capa (x).
  - x pot ser un número específic o la lletra y per referir-se a totes les capes.
- _N:y_: Indica la neurona (y) dins de la capa.
  - y pot ser un número específic o la lletra n per referir-se a totes les neurones.

__Exemples__:
- "L:0, N:n": Totes les neurones de la capa 0.
- "L:y, N:n": Totes les neurones de totes les capes.
- "L:2, N:3": La quarta neurona de la tercera capa.

- "L:2, N:3", "L:4" és equivalent a "L:2, N:3", "L:4, N:n".
- "L:y[^1]", "L:y+1" és equivalent a "L:y, N:n", "L:y+1, N:n".
[^1]: Simplificació: Si no s'especifica N:y, s'assumeix N:n. Per exemple:

### 2. `.ncl` - Neural Content Language (Contingut)
Estructura General
Un fitxer .ncl és un document JSON que conté els següents camps:

- _specification_: (Obligatori) Ruta o identificador del fitxer .ndl corresponent.
- _NFs_: (Opcional) Llista de funcions d'activació assignades a neurones o capes. Per defecte la funció escollida serà "ReLU"
- _weights_: (Opcional) Llista de pesos per a les connexions sinàptiques. Per defecte els pesos de _from_ a _to_ es crearan amb valors al·leatoris Xavier/Gorot.
- _biases_: (Opcional) Llista de valors de biaix per a les neurones. Per defecte els biaixos de _from_ a _to_ es crearan amb valor '0.0'.

Exemple de my_network_values.ncl:
```json
{
  "specification": "github.com/jibort/.../a.ndl",
  "NFs": [
    {"L:1, N:n": ["ReLU", "Tanh"]},
    {"L:1, N:3": ["Sigmoid"]},
    {"L:y, N:1": ["Tanh", "ReLU"]},
    {"L:1": ["Sigmoid", "ReLU", "Linear"]}
  ],
  "weights": [
    {"from": "L:0", "to": "L:1", "value": "-0.332"},
    {"from": "L:2", "to": "L:4, N:3", "value": "0.52"},
    {"from": "L:y, N:0", "to": "L:y+1, N:0", "value": "0.1*"},
    {"from": "L:2", "value": "RXG"},
    {"from": "L:2", "value": ["RNG"]}
  ],
  "biases": [
    {"from": "L:y", "value": "-0.59999"},
    {"from": "L:2", "value": "0.59999"},
    {"from": "L:y", "to": "N:3", value: "-0.31415"}
  ]
}
```

#### Descripció dels Camps
- _specification_
  - Tipus: Cadena de text
  - Descripció: Identifica el fitxer .ndl que defineix l'estructura de la xarxa a la qual es refereixen els valors proporcionats.
- NFs (Neural Functions)
  - Tipus: Llista d'objectes
  - Descripció: Assigna funcions d'activació a neurones o capes específiques. Cada objecte conté una clau que especifica la neurona o capa i un valor que és una llista de funcions d'activació.
Format:
```json
"NFs": [
  {"L:x, N:y": ["Funció1", "Funció2"]},
  ...
]
```

On:
- L:x: Indica la capa x. Pot ser un número específic o la lletra y per referir-se a totes les capes.
- N:y: Indica la neurona y dins de la capa. Pot ser un número específic o la lletra n per referir-se a totes les neurones.

_Exemples_:
- {"L:1, N:n": ["ReLU", "Tanh"]}: Assigna les funcions d'activació ReLU i Tanh a totes les neurones de la capa 1.
- {"L:y, N:1": ["Tanh", "ReLU"]}: Assigna les funcions d'activació Tanh i ReLU a la primera neurona de totes les capes.

- weights (Pesos)
  - Tipus: Llista d'objectes
  - Descripció: Defineix els pesos de les connexions sinàptiques entre neurones. Cada objecte especifica la neurona d'origen, la neurona de destinació i el valor del pes.
  - Format:
```json
"weights": [
  {"from": "L:x, N:y", "to": "L:z, N:w", "value": "pes"},
  ...
]
```
On:
- from: Neurona d'origen, especificada com "L:x, N:y".
- to: Neurona de destinació, especificada de la mateixa manera.
- value: Valor del pes, que pot ser un número o una cadena que indica un mètode d'inicialització aleatòria (per exemple, "RXG" per a Xavier/Glorot).

_Exemples_:
- {"from": "L:0", "to": "L:1", "value": "-0.332"}: Assigna un pes de -0.332 a totes les connexions de la capa 0 a la capa 1.
- {"from": "L:2", "value": "RXG"}: Inicialitza aleatòriament els pesos de totes les connexions que surten de la capa 2 segons el mètode Xavier/Glorot.

- biases (Biaixos)
  - Tipus: Llista d'objectes
  - Descripció: Especifica els valors de biaix per a neurones individuals o grups de neurones. Cada objecte indica la capa i/o neurona i el valor del biaix.
  - Format:
```json
"biases": [
  {"L:x, N:y": "biaix"},
  ...
]
'''
On:
- _L:x_: Indica la capa x.
- _N:y_: Indica la neurona y dins de la capa.
- _biaix_: Valor del biaix assignat.

__Exemples__:

  - {"L:y": "-0.59999"}: Assigna un biaix de -0.59999 a totes les neurones de totes les capes.
  - {"L:2": "0.59999"}: Assigna un biaix de 0.59999 a totes les neurones de la capa 2.

### 3. `.ntr` - Neural Training Configuration (Configuració d'Entrenament)
Aquest fitxer defineix els paràmetres per configurar l'entrenament de la xarxa.

Exemple de my_training_config.ntr:
json
Copia el codi
{
  "learningRate": 0.001,
  "batchSize": 32,
  "epochs": 100,
  "optimizer": "Adam",
  "lossFunction": "MSE"
}
### 4. `.chk` - Neural Checkpoint (Punt de Control)
Aquest fitxer desa l'estat complet de la xarxa.

Exemple de my_network_checkpoint.chk:
{
  "epoch": 42,
  "weights": [
    [
      [0.52, 0.21], [0.13, 0.79], [0.34, 0.42]
    ]
  ],
  "biases": [
    [0.12, 0.22, 0.32], [0.45, 0.53]
  ]
}
Especificacions Addicionals
.ndl - Connexions
Connexions per capa:
fromLayer: Índex de la capa origen.
toLayer: Índex de la capa destí.
Connexions específiques:
details: Llista de connexions entre neurones, incloent pesos.
.ncl - Funcions Neuronals
Funcions neuronals específiques per neurona.
Seqüències de funcions suportades.
Flux de Treball
Definir l'Estructura: Crear un fitxer .ndl amb les capes i connexions.
Afegir els Valors: Crear un fitxer .ncl amb els pesos, biaixos i funcions neuronals.
Configurar l'Entrenament: Definir un fitxer .ntr per ajustar els paràmetres.
Desar i Carregar Estat: Utilitzar fitxers .chk per reprendre l'entrenament.

Exemple d'Ús

// Crear una xarxa des de l'estructura
network, err := NewNetworkCPU("my_network.ndl")
if err != nil {
    log.Fatal(err)
}

// Carregar els valors inicials
err = network.LoadContent("my_network_values.ncl")
if err != nil {
    log.Fatal(err)
}

// Entrenar la xarxa
trainer := NewTrainer("my_training_config.ntr")
trainer.Train(network)

// Desar el punt de control
err = network.SaveCheckpoint("my_network_checkpoint.chk")
if err != nil {
    log.Fatal(err)
}
