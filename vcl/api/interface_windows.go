package api

import (
	"syscall"
)

var libvcl = syscall.NewLazyDLL("libvcl.dll")
