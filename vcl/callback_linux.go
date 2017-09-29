package vcl

// extern __stdcall *void doCallbackProc(f *void, args *void, argcount long);
import "C"

import (
	"unsafe"
)

//export doCallbackProc
func doCallbackProc(f unsafe.Pointer, args unsafe.Pointer, argcount C.long) unsafe.Pointer {
	callbackProc(uintptr(f), uintptr(args), int32(argcount))
	return C.NULL
}

var (
	callbackStdcall = uintptr(unsafe.Pointer(&doCallbackProc))
)
