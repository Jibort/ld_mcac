// Conté la interfície Grup D per a 64 bits
// CreatedAt: 2025/01/04 ds. JIQ

package range64

import (
	grp "github.com/jibort/ld_mcac/internal/core/intf/groups/base"
	x64 "github.com/jibort/ld_mcac/internal/core/intf/ranges"
)

// Interfície per a valors float64 del Grup D.
type F64GroupDIntf interface {
	x64.X64RangeIntf // Hereta les funcions dels range de 64bits.
	grp.GroupDIntf   // Hereta les funcions generals per a Grup D.
}

// Interfície per a valors uint64 del Grup D.
type U64GroupDIntf interface {
	x64.X64RangeIntf // Hereta les funcions dels range de 64bits.
	grp.GroupDIntf   // Hereta les funcions generals per a Grup D.
}
