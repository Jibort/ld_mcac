// Conté la interfície Grup B per a 32 bits
// CreatedAt: 2025/01/04 ds. JIQ

package range32

import (
	grp "github.com/jibort/ld_mcac/internal/core/intf/groups/base"
	x32 "github.com/jibort/ld_mcac/internal/core/intf/ranges"
)

// Interfície per a valors float32 del Grup B.
type F32GroupBIntf interface {
	x32.X32RangeIntf // Hereta les funcions dels range de 32bits.
	grp.GroupBIntf   // Hereta les funcions generals per a Grup B.
}

// Interfície per a valors uint32 del Grup B.
type U32GroupBIntf interface {
	x32.X32RangeIntf // Hereta les funcions dels range de 32bits.
	grp.GroupBIntf   // Hereta les funcions generals per a Grup B.
}
