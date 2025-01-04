package symbols

import (
	f64 "github.com/jibort/ld_mcac/internal/core/intf/groups/range64"
)

type U64RangeSymbolIntf interface {
	f64.U64GroupBIntf

	AsRune() rune
}
