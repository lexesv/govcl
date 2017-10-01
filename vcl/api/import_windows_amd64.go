package api

import (
	"syscall"
)

var libvcl = syscall.NewLazyDLL("libvclx64.dll")
