# Neural Description Language (NDL)

El **Neural Description Language (NDL)** és un llenguatge dissenyat per descriure, configurar i gestionar xarxes neuronals. Aquest llenguatge es compon de diversos tipus de fitxers que permeten separar l'estructura, el contingut, la configuració d'entrenament i els punts de control.

---

## Fitxers del Llenguatge NDL

### 1. `.ndl` - Neural Description Language (Estructura)
Aquest format de fitxer defineix l'estructura de la xarxa, incloent capes, neurones i connexions.

#### Exemple de `my_network.ndl`:
```code
{
  "name": "my_network",
  "description: "descripció opcional",
  "layers": [ 40, 10, 10, 5 ],
  "connections": [
    { "L:0, N:n", "L:1, N:n" },
    { "L:1, N:n", "L:2, N:n" },
    { "L:y, N:n", "L:y+1, N:n" },
    { "L:4, N:3", "L:2, N:7" },
    { "L:2, N:3", "L:4, N:n" },
  ],
}
```

L'estructura conté les següents seccions:
  - `name`: Nom de la xarxa.
  - `description`: Descripció opcional de la xarxa.
  - `layers`: Tamany de neurones per capa (implícitament també determina el número de capes per a la xarxa).
  - `connections`: És una llista de definició de les connexions entre neurones. L'estructura base és:
    1. `Cadena de sortida` ("_L:a, N:b_")
    - L:x quan x és un número, identifica la capa a partir d'on neix la connexió.
    - L:y quan y és de facto la lletra 'y', estableix totes les capes com a font de la definició de les connexions.
    - N:x, quan x és un número, identifica la neurona de la(les) capa(es) especificades per 'L' font de la(les) connexió(ons).
    2. `Cadena d'entrada` ("_L:a, N:b_")
    - Mateix format de cadenes que en la sortida.

Aquesta forma d'establir les connexions físiques entre neurones ofereix molta flexibilitat i compactibilitat, tant per a definir xarxes completament connectades entre capes ({"L:y, N:n", "L:y+1, N:n"}), fins a connexions específiques entre dues neurones qualsevol ({L:1, N:3", "L:3, N:1"}), fins a diferents tipus intermitjos ({L:1, N:3", "L:3, N:n"}).

Si en alguna definició de Layer's i Neurona's No apareix la part "N:a" és equivalent a escriure "N:n". Per exemple:

- "_L:2, N:3", "L:4_" és equivalent a "_L:2, N:3", "L:4, N:n_"
- "_L:y", "L:y+1_" és equivalent a "_L:y, N:n", "L:y+1, N:n_"

### 2. `.ncl` - Neural Content Language (Contingut)
Per a assegurar que un fitxer _ncl_ ha de correspondre necessàriament a una estructura específica descrita en un fitxer _ndl_ el primer paràmetre de l'estructura del fitxer _ncl_ ha de ser el nom del fitxer _ndl_.

Aquest format defineix els valors associats al contingut de la xarxa identificada amb _"specification", com ara pesos, biaixos, funcions neuronals, ...

Exemple de my_network_values.ncl:
```code
{
  "specification": "github.com/jibort/.../a.ndl",
  "NFs": [
    { "L:1, N:n", ["ReLU", "Tanh"] },
    { "L:1, N:3", ["Sigmoid"] },
    { "L:y, N:1", ["Tanh", "ReLU"] },
    { "L:1", ["Sigmoid", "ReLU", "Linear" ]}
  ],
  "weights": [
    { "L:0", "L:1", "-0.332" }, # El pes de totes les sinapsis que surten de neurones de la primera capa (_capa d'input_) i arriben a neurones de la capa "1" serà -0.332
    { "L:2", "L4:N3", "0.52" }, # El pes de totes les synapsis que siguin compatibles amb l'especificació inicial tindran valor 0.52 
    { "L:y, N:0", "L:y+1, N:0", , "0.1*" }, # El pes de la primera sinapsis de cada capa que tingui una connexió entre la primera neurona de capes adjacents tindrà pes "0.1*", és a dir 0.10000 saturat
    { "L:2", "RXG" }, # Els pesos de cada neurona de la capa 2 s'estableix al·latòriament segons el protocol Xavier/Gorot, es definiran més mecanismes d'al·leatorietat a mesura que aparegui la necessitat
    { "L:2", ["RNG"] }, # El pes de totes les sinapsis que neixen a neurones de la capa "2" s'estableix al·latòriament segons el protocol segur de creació de valors al·leatoris del paquet Go "rand"
  ],
  "biases": [
    { "L:y", "-0.59999" }, # els biases de totes les neurones de totes les capes tindrà el valor -0.59999
    { "L:2", "0.59999" }, # els biases de totes les neurones de la capa "2", tindrà el valor -0.59999
    { "L:y, N:3", "-0.31415" }, # els biases de les quartes neurones de totes les capes tindrà el valor -0.59999. El padding sempre serà 0.0
  ]
}
```

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
