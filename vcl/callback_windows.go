package vcl

import (
	"syscall"
)

var (
	callbackStdcall = syscall.NewCallbackCDecl(callbackProc)
)
