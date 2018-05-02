// +build windows

package win

import (
	"syscall"
	"unsafe"
)

var (

	// kernel32.dll
	versiondll              = syscall.NewLazyDLL("version.dll")
	getFileVersionInfoSizeW = versiondll.NewProc("GetFileVersionInfoSizeW")
	getFileVersionInfoW     = versiondll.NewProc("GetFileVersionInfoW")
	verQueryValueW          = versiondll.NewProc("VerQueryValueW")
)

// GetFileVersionInfoSizeW
func GetFileVersionInfoSizeW(lptstrFileName string, lpdwhandle *uint32) uint32 {
	r, _, _ := getFileVersionInfoSizeW.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lptstrFileName))),
		uintptr(unsafe.Pointer(lpdwhandle)))
	return uint32(r)
}

// GetFileVersionInfoW
func GetFileVersionInfoW(lptstrFilename string, dwHandle, dwLen uint32, lpData uintptr) bool {
	r, _, _ := getFileVersionInfoW.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lptstrFilename))),
		uintptr(dwHandle), uintptr(dwLen), lpData)
	return r != 0
}

// VerQueryValueW
func VerQueryValueW(pBlock uintptr, lpSubBlock string, lplpBuffer *uintptr, puLen *uint32) bool {
	r, _, _ := verQueryValueW.Call(uintptr(pBlock),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpSubBlock))),
		uintptr(unsafe.Pointer(lplpBuffer)), uintptr(unsafe.Pointer(puLen)))
	return r != 0
}
