// Interfície dels Valors de càlcul admessos dins els corresponents rangs ([-1.0,+1.0] i [-2·π, +2·π]).
// CreatedAt: 2025/01/04 ds. JIQ

package base

// Interfície dels Valors de càlcul admessos dins els corresponents rangs ([-1.0,+1.0] i [-2·π, +2·π]).
type GroupAIntf interface {
	Sign() bool          // Cert només si el valor és negatiu.
	IsSubnormal() bool   // Cert només si la instància pertany al subgrup A1.
	IsNormalNear0() bool // Cert només si la instància pertany al subgrup A2.
	IsNormalFar0() bool  // Cert només si la instància pertany al subgrup A3.
	IsInfinite() bool    // Cert només si la instància és un valor infinit.
	IsInfinitePos() bool // Cert només si la instància és un valor infinit positiu.
	IsInfiniteNeg() bool // Cert només si la instància és un valor infinit negatiu.

	Exponent() uint16 // Retorna l'exponent com a uint16 (11 bits)
	Mantissa() uint64 // Retorna la mantissa com a uint32 (52 bits)
	IsOne() bool      // Cert només si la instància és de rang [-1.0,+1.0]
	IsTwoPi() bool    // Cert només si la instància es de rang [-2·π, +2·π].
}
