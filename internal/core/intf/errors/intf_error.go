// Interfíce general d'error.
// CreatedAt: 2025/01/04 ds. JIQ

package errors

import (
	base "github.com/jibort/ld_mcac/internal/core/intf/base"
)

// Interfície per a valors de 64 bits (float64) del Subgrup B.4 (errors)
type ErrorIntf interface {
	base.RangeIntf

	IsCritical() bool    // Cert només si l'error és crític
	Code() uint16        // Retorna el codi d'error
	Arguments() []uint64 // Retorna els arguments de l'error
}
