// Implementació dels valors Range en uint64.
// CreatedAt: 2024/12/08 dg. JIQ

package ru764

import (
	r64 "github.com/jibort/ld_mcac/internal/core/r64"
)

// Tipus pels Range uint64.
type RangeU64 struct {
	// i64.Range64Intf
	r64.X64Range
	value uint64
}
