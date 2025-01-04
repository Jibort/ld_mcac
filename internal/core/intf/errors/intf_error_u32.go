package errors

import (
	f32 "github.com/jibort/ld_mcac/internal/core/intf/groups/range32"
)

type U32ErrorIntf interface {
	f32.U32GroupBIntf
	ErrorIntf
}
