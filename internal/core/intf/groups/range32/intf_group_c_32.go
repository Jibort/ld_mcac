// Conté la interfície Grup C per a 32 bits
// CreatedAt: 2025/01/04 ds. JIQ

package range32

import (
	grp "github.com/jibort/ld_mcac/internal/core/intf/groups/base"
	x32 "github.com/jibort/ld_mcac/internal/core/intf/ranges"
)

// Interfície per a valors float32 del Grup C.
type F32GroupCIntf interface {
	x32.X32RangeIntf // Hereta les funcions dels range de 32bits.
	grp.GroupCIntf   // Hereta les funcions generals per a Grup C.
}

// Interfície per a valors uint32 del Grup C.
type U32GroupCIntf interface {
	x32.X32RangeIntf // Hereta les funcions dels range de 32bits.
	grp.GroupCIntf   // Hereta les funcions generals per a Grup C.
}
