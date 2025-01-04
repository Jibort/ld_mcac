// Interfície general de tots els Grups C (tokens).
// CreatedAt: 2025/01/04 ds. JIQ

package base

// Interfície general de tots els Grups C (tokens).
type GroupCIntf interface {
	Category() int           // Retorna la categoria
	Fiability() int          // Retorna la fiabilitat
	RelativeWeight() float64 // Retorna el pes relatiu
	Token() uint32           // Retorna el token
}
