package symbols

import (
	f64 "github.com/jibort/ld_mcac/internal/core/intf/groups/range64"
)

type F64SymbolIntf interface {
	f64.F64GroupBIntf

	AsRune() rune
}
