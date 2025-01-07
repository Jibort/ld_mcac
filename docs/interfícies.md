# Estructura d'Interfícies del Projecte 'ld_mcac'
Aquest document identifica, localitza i descriu les diferents interfícies Go del projecte.

<style>
    table {
        border-collapse: collapse; /* Fusiona les vores */
        border: 2px solid black;   /* Vora per a tota la taula */
        width: 100%;               /* Amplada completa */
        max-width: 1000px;          /* Amplada màxima */
        margin: auto;              /* Centra la taula */
        text-align: center;        /* Alinea text al centre */
        page-break-after: auto;    /* Permet salts de pàgina */
    }
    th, td {
        border: 1px solid black;   /* Defineix les vores correctes */
        padding: 8px;              /* Espai dins de cada cel·la */
    }
    th {
        font-size: 18px;           /* Tamany de la font per a capçaleres */
    }
    td {
        font-size: 14px;           /* Tamany de la font per a cel·les */
    }
    .page-break {
        page-break-after: always;  /* Força un salt de pàgina */
    }
</style>

<table>
    <tr>
        <th colspan="2">Definició d'Interfícies</th>
    </tr>
    <tr>
        <td style="text-align:left"><b>ClonableIntf</b></td>
        <td style="text-align:left">Interfície general per a la clonació d'instàncies.</td>
    </tr>
    <tr>
        <td style="text-align:left"><b>ConversionsIntf</b></td>
        <td style="text-align:left">Interfície general per a la conversió cap a tipus base (float64, uint64, ...).</td>
    </tr>
    <tr>
        <td style="text-align:left"><b>ComparableIntf</b></td>
        <td style="text-align:left">Interfície general per a la comparació d'instàncies numèriques.</td>
    </tr>
    <tr>
        <td style="text-align:left"><b>GroupableIntf</b></td>
        <td style="text-align:left">Interfície general per a la validació de grup de les instàncies.</td>
    </tr>
    <tr>
        <td style="text-align:left"><b>RangebleIntf</b></td>
        <td style="text-align:left">Interfície per a identificar la longitud de les instàncies en bits.</td>
    </tr>
    <tr>
        <td style="text-align:left"><b>MathOperationsIntf</b></td>
        <td style="text-align:left">Interfície per a realitzar operacions matemàtiques com suma, resta, multiplicació, i divisió.</td>
    </tr>
    <tr>
        <td style="text-align:left"><b>NormalizeIntf</b></td>
        <td style="text-align:left">Interfície per a normalitzar valors dins d'un rang específic.</td>
    </tr>
    <tr>
        <td style="text-align:left"><b>RangeIntf</b></td>
        <td style="text-align:left">Interfície general per a treballar amb rangs numèrics.</td>
    </tr>
    <tr>
        <td style="text-align:left"><b>F64RangeIntf</b></td>
        <td style="text-align:left">Interfície específica per a treballar amb rangs float64.</td>
    </tr>
    <tr>
        <td style="text-align:left"><b>U64RangeIntf</b></td>
        <td style="text-align:left">Interfície específica per a treballar amb rangs uint64.</td>
    </tr>
    <tr>
        <td style="text-align:left"><b>F32RangeIntf</b></td>
        <td style="text-align:left">Interfície específica per a treballar amb rangs float32.</td>
    </tr>
    <tr>
        <td style="text-align:left"><b>U32RangeIntf</b></td>
        <td style="text-align:left">Interfície específica per a treballar amb rangs uint32.</td>
    </tr>
</table>

<table>
    <tr>
        <th colspan="8">Jerarquia d'Interfícies</th>
    </tr>
    <tr>
        <td>
            <b>F64RangeOneIntf</b><br>
            <i>ƒ AsF64TwoPi() F64RangeTwoPiIntf</i><br>
        </td>
        <td>
            <b>F64RangTwoPiIntf</b><br>
            <i>ƒ AsF64One() F64RangeOneIntf</i><br>
        </td>
        <td>
            <b>U64RangeOneIntf</b><br>
            <i>ƒ AsU64TwoPi() U64RangeTwoPiIntf</i><br>
        </td>
        <td>
            <b>U64RangeTwoPiIntf</b><br>
            <i>ƒ AsU64One() U64RangeOneIntf</i><br>
        </td>        
        <td>
            <b>F32RangeOneIntf</b><br>
            <i>ƒ AsF32TwoPi() F32RangeTwoPiIntf</i><br>
        </td>
        <td>
            <b>F32RangTwoPiIntf</b><br>
            <i>ƒ AsF32One() F32RangeOneIntf</i><br>
        </td>
        <td>   
            <b>U32RangeOneIntf</b><br>
            <i>ƒ AsU32TwoPi() U32RangeTwoPiIntf</i><br>
        </td>
        <td>
            <b>U32RangeTwoPiIntf</b><br>
            <i>ƒ AsU32One() U32RangeOneIntf</i><br>
        </td>
    </tr>
    <tr>
        <td colspan="2">
            <b>F64RangeIntf</b><br>
            <i>ƒ AsF32() F32RangeIntf</i><br>
        </td>
        <td colspan="2">
            <b>U64RangeIntf</b><br>
            <i>ƒ AsU32() U32RangeIntf</i><br>
        </td>
        <td colspan="2">
            <b>F32RangeIntf</b><br>
            <i>ƒ AsF64() F64RangeIntf</i><br>
        </td>
        <td colspan="2">
            <b>U32RangeIntf</b><br>
            <i>ƒ AsU64() U64RangeIntf</i><br>
        </td>
    </tr>
    <tr>
        <td colspan="4">
            <b>X64RangeIntf</b><br>
            <i>ƒ As32() X32RangeIntf</i><br>
        </td>
        <td colspan="4">
            <b>X32RangeIntf</b><br>
            <i>ƒ As64() X64RangeIntf</i><br>
        </td>
    </tr>
    <tr>
        <td colspan="2">
            <b>RangeIntf</b><br>
        </td>
        <td colspan="2">
            <b>MathOperationsIntf</b><br>
        </td>
        <td colspan="2">
            <b>RangeIntf</b><br>
        </td>
        <td colspan="2">
            <b>MathOperationsIntf</b><br>
        </td>
    </tr>
    <tr>
        <td colspan="2" text-align="right">
        <b>ClonableIntf</b><br><i>
            ƒ Clone() RangeIntf</i><br><br>
        <b>ConversionsIntf</b><br><i>
            ƒ AsFloat64() float64<br>
            ƒ SetFloat64(float64)<br>
            ƒ AsUint64() uint64<br>
            ƒ SetUint64(uint64)<br>
            ƒ AsFloat32() float32<br>
            ƒ SetFloat32(float32)<br>
            ƒ AsUint32() uint32<br>
            ƒ SetUint32(uint32)</i><br><br>
        <b>ComparableIntf</b><br><i>
            ƒ Equals(pOther RangeIntf) bool<br>
            ƒ LessThan(pOther RangeIntf) bool<br> 
            ƒ LessOrEqualThan(pOther RangeIntf) bool<br>
            ƒ GreaterThan(pOther RangeIntf) bool<br>
            ƒ GreaterOrEqualThan(pOther RangeIntf) bool<br><br></i>
        <b>GroupableIntf</b><br><i>
            ƒ IsGroupA() bool<br>
            ƒ IsGroupB() bool<br>
            ƒ IsGroupC() bool<br>
            ƒ IsGroupD() bool<br>
            ƒ IsGroupE() bool<br><br></i>
        <b>RangebleIntf</b><br><i>
            ƒ Is64() bool<br>
            ƒ Is32() bool<br></i>
        </td>
        <td colspan="2" style="vertical-align:top"><i>
            ƒ Add(pOther RangeIntf) RangeIntf<br>
            ƒ Sub(pOther RangeIntf) RangeIntf<br>
            ƒ Mul(pOther RangeIntf) RangeIntf<br>
            ƒ Div(pOther RangeIntf) RangeIntf<br>
            ƒ IsF32() bool<br>
            ƒ IsU32() bool<br>
            ƒ IsF64() bool<br>
            ƒ IsU64() bool<br></i>
        </td>
        <td colspan="2" text-align="left">
            <b>ClonableIntf</b><br><i>
                ƒ Clone() RangeIntf</i><br><br>
            <b>ConversionsIntf</b><br><i>
                ƒ AsFloat64() float64<br>
                ƒ SetFloat64(float64)<br>
                ƒ AsUint64() uint64<br>
                ƒ SetUint64(uint64)<br>
                ƒ AsFloat32() float32<br>
                ƒ SetFloat32(float32)<br>
                ƒ AsUint32() uint32<br>
                ƒ SetUint32(uint32)</i><br><br>
            <b>ComparableIntf</b><br><i>
                ƒ Equals(pOther RangeIntf) bool<br>
                ƒ LessThan(pOther RangeIntf) bool<br> 
                ƒ LessOrEqualThan(pOther RangeIntf) bool<br>
                ƒ GreaterThan(pOther RangeIntf) bool<br>
                ƒ GreaterOrEqualThan(pOther RangeIntf) bool<br><br></i>
            <b>GroupableIntf</b><br><i>
                ƒ IsGroupA() bool<br>
                ƒ IsGroupB() bool<br>
                ƒ IsGroupC() bool<br>
                ƒ IsGroupD() bool<br>
                ƒ IsGroupE() bool<br><br></i>
            <b>RangebleIntf</b><br><i>
                ƒ Is64() bool<br>
                ƒ Is32() bool<br></i>
            </td>
            <td colspan="2" style="vertical-align:top"><i>
                ƒ Add(pOther RangeIntf) RangeIntf<br>
                ƒ Sub(pOther RangeIntf) RangeIntf<br>
                ƒ Mul(pOther RangeIntf) RangeIntf<br>
                ƒ Div(pOther RangeIntf) RangeIntf<br>
                ƒ IsF32() bool<br>
                ƒ IsU32() bool<br>
                ƒ IsF64() bool<br>
                ƒ IsU64() bool<br></i>
            </td>   
    </tr>
</table>