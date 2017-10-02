package api

import (
	"unsafe"
)

// TStyleManager

// styleManager_IsValidStyle        = libvcl.NewProc("StyleManager_IsValidStyle")
// styleManager_LoadFromFile        = libvcl.NewProc("StyleManager_LoadFromFile")
// styleManager_CheckSysClassName   = libvcl.NewProc("StyleManager_CheckSysClassName")
// styleManager_TrySetStyle         = libvcl.NewProc("StyleManager_TrySetStyle")
// styleManager_SetStyle            = libvcl.NewProc("StyleManager_SetStyle")
// styleManager_TryLoadFromResource = libvcl.NewProc("StyleManager_TryLoadFromResource")

// StyleManager_Initialize
func StyleManager_Initialize() {
	styleManager_Initialize.Call()
}

// StyleManager_UnInitialize
func StyleManager_UnInitialize() {
	styleManager_UnInitialize.Call()
}

// StyleManager_IsValidStyle
func StyleManager_IsValidStyle(filename string) bool {
	r, _, _ := styleManager_IsValidStyle.Call(GoStrToDStr(filename))
	return r != 0
}

// StyleManager_LoadFromFile
func StyleManager_LoadFromFile(filename string) uintptr {
	r, _, _ := styleManager_LoadFromFile.Call(GoStrToDStr(filename))
	return r
}

// StyleManager_CheckSysClassName
func StyleManager_CheckSysClassName(className string) bool {
	r, _, _ := styleManager_CheckSysClassName.Call(GoStrToDStr(className))
	return r != 0
}

// StyleManager_TrySetStyle
func StyleManager_TrySetStyle(name string, showErrorDialog bool) bool {
	r, _, _ := styleManager_TrySetStyle.Call(GoStrToDStr(name), GoBoolToDBool(showErrorDialog))
	return r != 0
}

// StyleManager_SetStyle1
func StyleManager_SetStyle1(handle uintptr) {
	styleManager_SetStyle1.Call(handle)
}

// StyleManager_SetStyle2
func StyleManager_SetStyle2(name string) {
	styleManager_SetStyle2.Call(GoStrToDStr(name))
}

// StyleManager_TryLoadFromResource
func StyleManager_TryLoadFromResource(instance uintptr, resName, resType string, handle *uintptr) bool {
	r, _, _ := styleManager_TryLoadFromResource.Call(instance, GoStrToDStr(resName), GoStrToDStr(resType), uintptr(unsafe.Pointer(handle)))
	return r != 0
}
