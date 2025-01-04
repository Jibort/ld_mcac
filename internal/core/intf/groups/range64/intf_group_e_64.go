// Conté la interfície Grup E per a 64 bits (reservada per a noves extensions).
// CreatedAt: 2025/01/04 ds. JIQ

package range64

import (
	grp "github.com/jibort/ld_mcac/internal/core/intf/groups/base"
	x64 "github.com/jibort/ld_mcac/internal/core/intf/ranges"
)

// Interfície per a valors float64 del Grup E.
type F64GroupEIntf interface {
	x64.X64RangeIntf // Hereta les funcions dels range de 64bits.
	grp.GroupEIntf   // Hereta les funcions generals per a Grup E.
}

// Interfície per a valors uint64 del Grup E.
type U64GroupEIntf interface {
	x64.X64RangeIntf // Hereta les funcions dels range de 64bits.
	grp.GroupEIntf   // Hereta les funcions generals per a Grup E.
}
