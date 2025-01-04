package errors

import (
	f64 "github.com/jibort/ld_mcac/internal/core/intf/groups/range64"
)

type U64ErrorIntf interface {
	f64.U64GroupBIntf
	ErrorIntf
}
