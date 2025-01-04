// Conté la interfície Grup A per a 64 bits
// CreatedAt: 2025/01/04 ds. JIQ

package range64

import (
	grp "github.com/jibort/ld_mcac/internal/core/intf/groups/base"
	x64 "github.com/jibort/ld_mcac/internal/core/intf/ranges"
)

// Interfície per a valors float64 del Grup A.
type F64GroupAIntf interface {
	x64.X64RangeIntf // Hereta les funcions dels range de 64bits.
	grp.GroupAIntf   // Hereta les funcions generals per a Grup A.
}

// Interfície per a valors uint64 del Grup A.
type U64GroupAIntf interface {
	x64.X64RangeIntf // Hereta les funcions dels range de 64bits.
	grp.GroupAIntf   // Hereta les funcions generals per a Grup A.
}
