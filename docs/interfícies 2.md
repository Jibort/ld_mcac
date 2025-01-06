# Jerarquia de interfícies i tipus de 32 i 64 bits
Aquest document descriu les relacions entre les interfícies del projecte i com els tipus en fan us i es relacionen entre ells.

<style>
    table {
        border-collapse: collapse; /* Fusiona les vores */
        border: 2px solid black;
        width: 120%;               /* Amplada completa */
        max-width: 600px;          /* Amplada màxima */
        margin: auto;              /* Centra la taula */
        text-align: center;
    }

    th {
        borde: 1px solid black;
        text-align: center;
        font-size: 20px;           /* Tamany de la font */
        padding: 8px;
    }
    
    td {
        borde: 1px solid black;
        text-align: center;
        font-size: 12px;           /* Tamany de la font */
        padding: 8px;
    }
</style>

<table>
    <tr>
        <th colspan="8">Jerarquia d'Interfícies</th>
    </tr>
    <tr>
        <td><b>F64RangeOneIntf</b></td>
        <td><b>F64RangTwoPiIntf</b></td>
        <td><b>U64RangeOneIntf</b></td>
        <td><b>U64RangeTwoPiIntf</b></td>        
        <td><b>F32RangeOneIntf</b></td>
        <td><b>F32RangTwoPiIntf</b></td>
        <td><b>U32RangeOneIntf</b></td>
        <td><b>U32RangeTwoPiIntf</b></td>        
    </tr>
    <tr>
        <td colspan="2"><b>F64RangeIntf</b></td>
        <td colspan="2"><b>U64RangeIntf</b></td>
        <td colspan="2"><b>F32RangeIntf</b></td>
        <td colspan="2"><b>U32RangeIntf</b></td>
    </tr>
    <tr>
        <td colspan="4"><b>X64RangeIntf</b></td>
        <td colspan="4"><b>X32RangeIntf</b></td>
    </tr>
    <tr>
        <td colspan="2"><b>RangeIntf</b></td>
        <td colspan="2"><b>MathOperationsIntf</b></td>
        <td colspan="2"><b>RangeIntf</b></td>
        <td colspan="2"><b>MathOperationsIntf</b></td>
    </tr>
    <tr>
        <td colspan="2" text-align="right">
        <b>ClonableIntf</b><br><i>
            Clone() RangeIntf</i><br><br>
        <b>ConversionsIntf</b><br><i>
            AsFloat64() float64<br>
            SetFloat64(float64)<br>
            AsUint64() uint64<br>
            SetUint64(uint64)<br>
            AsFloat32() float32<br>
            SetFloat32(float32)<br>
            AsUint32() uint32<br>
            SetUint32(uint32)</i><br><br>
        <b>ComparableIntf</b><br><i>
            Equals(pOther RangeIntf) bool<br>
            LessThan(pOther RangeIntf) bool<br> 
            LessOrEqualThan(pOther RangeIntf) bool<br>
            GreaterThan(pOther RangeIntf) bool<br>
            GreaterOrEqualThan(pOther RangeIntf) bool<br><br></i>
        <b>GroupableIntf</b><br><i>
            IsGroupA() bool<br>
            IsGroupB() bool<br>
            IsGroupC() bool<br>
            IsGroupD() bool<br>
            IsGroupE() bool<br><br></i>
        <b>RangebleIntf</b><br><i>
            Is64() bool<br>
            Is32() bool<br></i>
        </td>
        <td colspan="2" style="vertical-align:top"><i>
            Add(pOther RangeIntf) RangeIntf<br>
            Sub(pOther RangeIntf) RangeIntf<br>
            Mul(pOther RangeIntf) RangeIntf<br>
            Div(pOther RangeIntf) RangeIntf<br>
            IsF32() bool<br>
            IsU32() bool<br>
            IsF64() bool<br>
            IsU64() bool<br></i>
        </td>
        <td colspan="2" text-align="left">
            <b>ClonableIntf</b><br><i>
            Clone() RangeIntf</i><br><br>
        <b>ConversionsIntf</b><br><i>
            AsFloat64() float64<br>
            SetFloat64(float64)<br>
            AsUint64() uint64<br>
            SetUint64(uint64)<br>
            AsFloat32() float32<br>
            SetFloat32(float32)<br>
            AsUint32() uint32<br>
            SetUint32(uint32)</i><br><br>
        <b>ComparableIntf</b><br><i>
            Equals(pOther RangeIntf) bool<br>
            LessThan(pOther RangeIntf) bool<br> 
            LessOrEqualThan(pOther RangeIntf) bool<br>
            GreaterThan(pOther RangeIntf) bool<br>
            GreaterOrEqualThan(pOther RangeIntf) bool<br><br></i>
        <b>GroupableIntf</b><br><i>
            IsGroupA() bool<br>
            IsGroupB() bool<br>
            IsGroupC() bool<br>
            IsGroupD() bool<br>
            IsGroupE() bool<br><br></i>
        <b>RangebleIntf</b><br><i>
            Is64() bool<br>
            Is32() bool<br></i>
        </td>
        <td colspan="2" style="vertical-align:top"><i>
            Add(pOther RangeIntf) RangeIntf<br>
            Sub(pOther RangeIntf) RangeIntf<br>
            Mul(pOther RangeIntf) RangeIntf<br>
            Div(pOther RangeIntf) RangeIntf<br>
            IsF32() bool<br>
            IsU32() bool<br>
            IsF64() bool<br>
            IsU64() bool<br></i>
        </td>   
    </tr>
</table>
