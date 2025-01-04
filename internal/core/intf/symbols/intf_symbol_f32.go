package symbols

import (
	f32 "github.com/jibort/ld_mcac/internal/core/intf/groups/range32"
)

type F32SymbolIntf interface {
	f32.F32GroupBIntf

	AsRune() rune
}
