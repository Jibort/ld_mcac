package errors

import (
	f64 "github.com/jibort/ld_mcac/internal/core/intf/groups/range64"
)

type F64ErrorIntf interface {
	f64.F64GroupBIntf
	ErrorIntf
}
