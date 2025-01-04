// Conté la interfície Grup D per a 32 bits
// CreatedAt: 2025/01/04 ds. JIQ

package range32

import (
	grp "github.com/jibort/ld_mcac/internal/core/intf/groups/base"
	x32 "github.com/jibort/ld_mcac/internal/core/intf/ranges"
)

// Interfície per a valors float32 del Grup D.
type F32GroupDIntf interface {
	x32.X32RangeIntf // Hereta les funcions dels range de 32bits.
	grp.GroupDIntf   // Hereta les funcions generals per a Grup D.
}

// Interfície per a valors uint32 del Grup D.
type U32GroupDIntf interface {
	x32.X32RangeIntf // Hereta les funcions dels range de 32bits.
	grp.GroupDIntf   // Hereta les funcions generals per a Grup D.
}
