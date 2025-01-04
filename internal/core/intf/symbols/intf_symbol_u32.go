package symbols

import (
	f32 "github.com/jibort/ld_mcac/internal/core/intf/groups/range32"
)

type U32SymbolIntf interface {
	f32.U32GroupBIntf

	AsRune() rune
}
