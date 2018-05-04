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

type TOSVersionInfoExW struct {
	OSVersionInfoSize uint32
	MajorVersion      uint32
	MinorVersion      uint32
	BuildNumber       uint32
	PlatformId        uint32
	CSDVersion        [128]uint16 // Maintenance UnicodeString for PSS usage
	ServicePackMajor  uint16
	ServicePackMinor  uint16
	SuiteMask         uint16
	ProductType       uint8
	Reserved          uint8
}

type TVSFixedFileInfo struct {
	Signature        uint32 // e.g. $feef04bd
	StrucVersion     uint32 // e.g. $00000042 = "0.42"
	FileVersionMS    uint32 // e.g. $00030075 = "3.75"
	FileVersionLS    uint32 // e.g. $00000031 = "0.31"
	ProductVersionMS uint32 // e.g. $00030010 = "3.10"
	ProductVersionLS uint32 // e.g. $00000031 = "0.31"
	FileFlagsMask    uint32 // = $3F for version "0.42"
	FileFlags        uint32 // e.g. VFF_DEBUG | VFF_PRERELEASE
	FileOS           uint32 // e.g. VOS_DOS_WINDOWS16
	FileType         uint32 // e.g. VFT_DRIVER
	FileSubtype      uint32 // e.g. VFT2_DRV_KEYBOARD
	FileDateMS       uint32 // e.g. 0
	FileDateLS       uint32 // e.g. 0
}

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
