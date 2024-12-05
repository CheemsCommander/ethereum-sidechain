//go:build !darwin
// +build !darwin

package drivechain

import "C"

func newUlong(in uint64) C.ulong {
	return C.ulong(in)
}