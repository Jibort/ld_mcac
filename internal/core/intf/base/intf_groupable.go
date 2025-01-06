// Interfície d'instàncies que pertanyen a algun grup.
// CreatedAt: 2025/01/04 ds. JIQ

package base

type GroupableIntf interface {
	IsGroupA() bool // Cert només si la instància pertany al grup A.
	IsGroupB() bool // Cert només si la instància pertany al grup B.
	IsGroupC() bool // Cert només si la instància pertany al grup C.
	IsGroupD() bool // Cert només si la instància pertany al grup D.
	IsGroupE() bool // Cert només si la instància pertany al grup E (només per a 64bits).
}
