// Interfícies de Range pel Grup C en float64.
// CreatedAt: 2025/01/03 dv. JIQ

package Intf

type RangeF64TokenIntf interface {
	RangeF64Intf
	Category() int           // Retorna la categoria
	Fiability() int          // Retorna la fiabilitat
	RelativeWeight() float64 // Retorna el pes relatiu
	Token() uint32           // Retorna el token
}
