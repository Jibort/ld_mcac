# Discussió sobre el Grup D com a contenidor de metareferències

## Introducció

### Comentari inicial

> Estic donant voltes a la idea del Grup D com a contenidor de metareferències tal i com ho havíem plantejat (network, layer, neuron, synapse, ...).
>
> Com no disposàvem d'espai suficient per a identificar amb prou detall la xarxa, la capa, la neurona i la sinapsi al mateix temps, vam acabar pensant en fer que cada RangeF64 del grup D identifiqués de forma individual cada element i que disposéssim d'un mecanisme per a indicar que hi havia una seqüència de RangeF64 que identificava en línia, un a un, els elements.
>
> Si volíem identificar una xarxa faríem servir un RangeF64 amb el bit de final de seqüència a 1. Si volíem identificar una neurona el primers dos RangeF64 identificarien la xarxa i la capa amb el flag a 0 i el següent identificaria la neurona amb flag igual a 1.
>
> Els problemes d'aquesta solució són dos al meu parer:
> 1. No té sentit fer servir 57b per a identificar aquests elements. Igual que en la primera aproximació ens vam adonar que no disposàvem de prou bits per a tot fer servir 57 tampoc no s'adapta a la realitat.
> 2. El concepte de seqüència de RangeF64 no encaixa realment amb el model de xarxes. La gràcia d'estendre el float64 amb extensions és, en gran mesura la capacitat de poder manipular tant valors de càlcul com altres tipus d'informació de forma atòmica, sense haver de canviar l'estructura i funcionament natural de les xarxes.
>
> Ara el model de bits del Grup D (que ara dediquem en exclusiva a les metareferències) està quedant així:
> "L110TTTr-rrrrrrrr-rrrrrrrr-rrrrrrrr-rrrrrrrr-rrrrrrrr-rrrrrrrr-rrrrrrrr", on 'L' és el flag d'últim element de la seqüència, 'TTT' 'és el tipus d'element (xarxa, capa, ...) i 'r' són els 57 bits d'identificació de l'element actual.
>
> Què et semblaria que ho féssim així?:
>
> ```
> B7           B6           B5            B4              B3             B2          B1         B0
> "x110xxxx-xxxxxxyy-yyyyyyyy-yyyyynnn-nnnnnnnn-nnnnnsss-ssssssss-ssssssss"
> ```
>
> on:
> - 11x: Identificador de xarxa (2.048 xarxes diferents).
> - 15y: Identificador de capa dins la xarxa (32.768 capes diferents).
> - 17n: Identificador de neurona dins la capa (131.072 neurones individuals per capa).
> - 18s: Identificador de sinapsi de sortida de la neurona (262.144 possibles connexions a neurones destí).
>
> El fet de dedicar només 3 bits a identificar el grup i prou ens dona 61 bits per a codificar amb precisió de sinapsi. Com ho veus tu? Arribem a un gra suficientment petit com per a ser útil?
>
> A més, sempre estarà informat el cap de xarxa, perquè és la mínima referència possible. La resta de bits dels 61 a 0. Això significaria tipus 'Xarxa'. Quan informem d'una capa de la xarxa disposaríem de 32.768 - 1 combinacions possibles, donat que 0 significaria nul. Igual per a neurones i sinapsis.

---

## Anàlisi inicial de la proposta

- **Resposta**:
    - La proposta de repartir els **61 bits** entre xarxa, capa, neurones i sinapsis és **clara i coherent**. Assignar menys bits a xarxes (àmbit més global) i més a elements més granulars com sinapsis (nivell micro) sembla **molt adequat**.
    
    - Els valors proposats ofereixen una jerarquia equilibrada:
      - Xarxes: \(2^{11} = 2,048\) possibles xarxes.
      - Capa: \(2^{15} = 32,768\) possibles capes per xarxa.
      - Neurones: \(2^{17} = 131,072\) possibles neurones per capa.
      - Sinapsis: \(2^{18} = 262,144\) connexions de sortida per neurona.

    - La distribució és suficient per a la majoria de xarxes reals actuals, i la proposta de flag per xarxa, capa, neurona i sinapsi ajuda a mantenir consistència en la representació.

---

## Millora de la proposta: Ampliació de sinapsis

- **Nova proposta**:
    > Reduir el nombre de bits per xarxes a **9 bits (512 xarxes)** i redistribuir-los per ampliar sinapsis:
    >
    > ```
    > xarxa    9       512 xarxes
    > capa    15     32,768 capes
    > neurona 18    262,144 neurones per capa
    > sinapsi 19    524,288 sinapsis per neurona
    > TOTAL    61
    > ```

- **Comentari**: Aquesta redistribució fa que el model sigui encara més **granular**, especialment en sinapsis, que solen ser l'àrea amb major necessitat de bits en xarxes profundes completament connectades.

---

## Consideracions sobre escenaris de futur

### **Impacte del Grup E per a expansibilitat**

- Es proposa que el **Grup E** (actualment no assignat) pugui utilitzar-se com a extensió per a escenaris on ni tan sols els **61 bits** del Grup D siguin suficients.
- Això permetria:
  - Augmentar el nombre de xarxes o altres nivells jeràrquics si fos necessari.
  - Adaptar-se a xarxes neuronals que requereixin una granularitat encara més fina en el futur.

### **Limitacions del model de 32 bits**

- El model de **32 bits** quedarà probablement restringit a:
  - Xarxes més petites.
  - Sistemes on la velocitat i el consum de memòria siguin crítics.

- **Proposta de distribució (32 bits):**
    - Xarxa: No necessari.
    - Capa: 10 bits (\(2^{10} = 1,024\) capes).
    - Neurona: 11 bits (\(2^{11} - 1 = 2,047\)).
    - Sinapsi: 11 bits (\(2^{11} - 1 = 2,047\)).

- Aquesta distribució permetria suportar xarxes totalment connectades de mida moderada.

---

## Conclusions

- El disseny del **Grup D** amb **61 bits** permet una codificació granular de metareferències de xarxes neuronals.
- La proposta d'assignar **19 bits a sinapsis** aporta la flexibilitat necessària per xarxes denses.
- El **model de 32 bits** serà una opció viable per a xarxes compactes i especialitzades, amb una distribució equilibrada entre capes, neurones i sinapsis.
- El **Grup E** s'hauria de reservar per a futures ampliacions, assegurant escalabilitat a llarg termini.

### Properes passes

1. Completar la implementació del model **RangeF64** per validar aquesta estructura.
2. Desenvolupar el model **RangeF32** i establir benchmarks per comparar-los.
3. Documentar les conclusions a mesura que avancem en altres exemples.

4. Establir mecanismes per a casos d’interacció entre xarxes en futurs dissenys.

### Establir mecanismes per a casos d’interacció entre xarxes en futurs dissenys

En aquesta fase, caldrà tenir en compte:

1. **Estructuració de metareferències comunes:**

   - Dissenyar un model que permeti compartir referències entre xarxes.
   - Utilitzar identificadors globals o reutilitzar els espais del **Grup E** per a aquest propòsit.

2. **Compatibilitat amb escalabilitat dinàmica:**

   - Assegurar que el sistema pugui adaptar-se a noves necessitats sense necessitat de redissenyar els nivells jeràrquics.
   - Permetre que xarxes autònomes puguin interactuar sense dependre de modificacions al disseny principal.

3. **Establir canals de comunicació entre elements de xarxes:**

   - Incloure referències que permetin identificar capes, neurones i sinapsis d’altres xarxes.
   - Garantir que la interacció es pugui representar de manera eficient sense comprometre la precisió.

4. **Documentació i proves exhaustives:**

   - Caldrà documentar clarament les regles d’interacció per evitar inconsistències.
   - Definir escenaris de test per validar comportaments inter-xarxes.

5. **Avaluació de les necessitats futures:**

   - Analitzar les possibles aplicacions d’aquests mecanismes per assegurar que les funcionalitats cobreixen tant necessitats immediates com escenaris potencials.

Aquestes passes seran crítiques per garantir que el sistema de metareferències es mantingui flexible i escalable en el temps.

---

### Reflexions sobre l’ús del Grup E i les limitacions del model de 32 bits

#### **1. Ús del Grup E per a expansibilitat**

El **Grup E** de `RangeF64`, no assignat actualment, pot utilitzar-se com a extensió en escenaris on el disseny original del Grup D (amb 61 bits) no sigui suficient. Això ofereix un gran avantatge per a futurs casos d'ús:

- **Escalabilitat dinàmica:**
  - Permet afegir més bits per identificar xarxes, capes, neurones o sinapsis, depenent de les necessitats.
  - Exemples: 
    - Xarxes amb més de 512 xarxes o capes.
    - Connexions molt complexes amb més sinapsis per neurona.

- **Compatibilitat amb el model actual:**
  - Mantenir el Grup E com una extensió permet que les aplicacions que no necessiten aquesta granularitat continuïn utilitzant el model estàndard del Grup D sense modificacions.

#### **2. Limitacions del model de 32 bits**

El model de **32 bits** té clares limitacions per a xarxes neuronals complexes:

- **Espai de representació reduït:**
  - Amb noms 32 bits, la capacitat de representar nivells jeràrquics amb precisó es redueix.
  - Exemples: 
    - Amb 9 bits per xarxes, gairebé una tercera part del rang ja queda ocupada.

- **Manca de flexibilitat:**
  - No permet configuracions dinàmiques com en el model de 64 bits.

- **Ús en subxarxes compactes:**
  - És una opció viable per xarxes petites, on la velocitat i eficiència siguin prioritaris.

#### **3. Avantatges del model de 64 bits**

- **Escalabilitat i precisó:**
  - Representació robusta per a jerarquies complexes.

- **Extensibilitat:**
  - Permet configuracions dinàmiques per a futures necessitats.

- **Integració amb `RangeF64`:**
  - Facilita la manipulació d'aquests valors dins les xarxes neuronals.

---

### Conclusions i properes passes

#### **Diferenciació per ús:**

- Utilitzar el model de **32 bits** per a subxarxes petites o especialitzades on la velocitat i la memòria siguin prioritaris.
- Reservar el model de **64 bits** com a estàndard per a xarxes més complexes i grans.

#### **Implementació progressiva:**

- Desenvolupar primer el model **RangeF64** per validar aquesta estructura.
- Posteriorment, adaptar una versió reduïda per al model **RangeF32**.

#### **Crear benchmarks:**

- Comparar **32 bits** i **64 bits** en termes de:
  - Precisó dels resultats.
  - Velocitat d’entrenament i execució.
  - Consum de memòria.

Aquest enfocament permetrà una anàlisi profunda de les necessitats i els avantatges de cada model.

