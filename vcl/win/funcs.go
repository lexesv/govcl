// +build windows

package win

import (
	"unsafe"

	"fmt"
	"syscall"
)

const MAX_VERS = 20

// FixWindowsVersion
func FixWindowsVersion(MajorVersion, MinorVersion, ServicePackMajor *int) {

	var i int
	var osvi TOSVersionInfoEx
	var dwlConditionMask uint64
	var dwTypeMask uint32

	osvi.OSVersionInfoSize = uint32(unsafe.Sizeof(osvi))
	osvi.MinorVersion = 0
	osvi.ServicePackMajor = 0

	if *ServicePackMajor != -1 {
		dwTypeMask = VER_MAJORVERSION | VER_MINORVERSION | VER_SERVICEPACKMAJOR
	} else {
		dwTypeMask = VER_MAJORVERSION | VER_MINORVERSION
	}
	dwlConditionMask = VerSetConditionMask(0, VER_MAJORVERSION, VER_GREATER_EQUAL)
	dwlConditionMask = VerSetConditionMask(dwlConditionMask, VER_MINORVERSION, VER_GREATER_EQUAL)
	if *ServicePackMajor != -1 {
		dwlConditionMask = VerSetConditionMask(dwlConditionMask, VER_SERVICEPACKMAJOR, VER_GREATER_EQUAL)

	}

	//Try to find MajorVersion
	for i = 7; i <= MAX_VERS; i++ {
		osvi.MajorVersion = uint32(i)
		if VerifyVersionInfo(&osvi, dwTypeMask, dwlConditionMask) == false {
			*MajorVersion = i - 1
			break
		}
	}
	osvi.MajorVersion = uint32(*MajorVersion)

	//Try to find MinorVersion
	for i = 1; i <= MAX_VERS; i++ {
		osvi.MinorVersion = uint32(i)
		if VerifyVersionInfo(&osvi, dwTypeMask, dwlConditionMask) == false {
			*MinorVersion = i - 1
		}
	}
	osvi.MinorVersion = uint32(*MinorVersion)

	if *ServicePackMajor != -1 {
		//Try to find ServicePackMajor
		for i = 1; i <= MAX_VERS; i++ {
			osvi.ServicePackMajor = uint16(i)
			if VerifyVersionInfo(&osvi, dwTypeMask, dwlConditionMask) == false {
				*ServicePackMajor = i - 1
				break
			}
		}
	}
}

// CheckWin32Version
func CheckWin32Version(AMajor, AMinor int) bool {
	return (Win32MajorVersion > AMajor) ||
		((Win32MajorVersion == AMajor) &&
			(Win32MinorVersion >= AMinor))
}

// IsAdministrator 判断当前进程是否以Administrator权限运行中
func IsAdministrator() bool {
	var tokenHandle uintptr
	if CheckWin32Version(6, 0) {
		if OpenProcessToken(GetCurrentProcess(), TOKEN_QUERY, &tokenHandle) {
			defer CloseHandle(tokenHandle)
			var TokenIsElevated uint32
			var ReturnLength uint32
			if GetTokenInformation(tokenHandle, TokenElevation, uintptr(unsafe.Pointer(&TokenIsElevated)), uint32(unsafe.Sizeof(TokenIsElevated)), &ReturnLength) {
				if ReturnLength == uint32(unsafe.Sizeof(TokenIsElevated)) {
					return TokenIsElevated > 0
				}
			}
		}
		return false
	}
	return true
}

// OpenInExplorer
func OpenInExplorer(aFileName string) {
	ShellExecute(0, "OPEN", "Explorer.exe",
		fmt.Sprintf("/e, /select, \"%s\"", aFileName), "", SW_SHOW)
}

// 内部两个
// GoStrToCStr
func CStr(str string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
}

// CStrToGoStr
func GoStr(str []uint16) string {
	return syscall.UTF16ToString(str)
}

// GoBoolToCBool
func CBool(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

// CBoolToGoBool
func GoBool(b uintptr) bool {
	if b != 0 {
		return true
	}
	return false
}
