// Conté la interfície Grup B per a 64 bits
// CreatedAt: 2025/01/04 ds. JIQ

package range64

import (
	grp "github.com/jibort/ld_mcac/internal/core/intf/groups/base"
	x64 "github.com/jibort/ld_mcac/internal/core/intf/ranges"
)

// Interfície per a valors float64 del Grup B.
type F64GroupBIntf interface {
	x64.X64RangeIntf // Hereta les funcions dels range de 64bits.
	grp.GroupBIntf   // Hereta les funcions generals per a Grup B.
}

// Interfície per a valors uint64 del Grup B.
type U64GroupBIntf interface {
	x64.X64RangeIntf // Hereta les funcions dels range de 64bits.
	grp.GroupBIntf   // Hereta les funcions generals per a Grup B.
}
