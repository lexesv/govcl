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

// StyleManager_ActiveStyle
func StyleManager_ActiveStyle() uintptr {
	r, _, _ := styleManager_ActiveStyle.Call()
	return r
}

// StyleManager_SystemStyle
func StyleManager_SystemStyle() uintptr {
	r, _, _ := styleManager_SystemStyle.Call()
	return r
}

// StyleManager_Enabled
func StyleManager_Enabled() bool {
	r, _, _ := styleManager_Enabled.Call()
	return DBoolToGoBool(r)
}

// StyleManager_IsCustomStyleActive
func StyleManager_IsCustomStyleActive() bool {
	r, _, _ := styleManager_IsCustomStyleActive.Call()
	return DBoolToGoBool(r)
}

// StyleManager_UnRegisterStyle
func StyleManager_UnRegisterStyle(style uintptr) {
	styleManager_UnRegisterStyle.Call(style)
}

// StyleManager_RegisterStyle
func StyleManager_RegisterStyle(style uintptr) {
	styleManager_RegisterStyle.Call(style)
}

// StyleManager_Style
func StyleManager_Style(name string) uintptr {
	r, _, _ := styleManager_Style.Call(GoStrToDStr(name))
	return r
}

// // function StyleManager_StyleDescriptor(StyleName: PChar): TStyleManager.TStyleClassDescriptor; stdcall;
// func StyleManager_StyleDescriptor(styleName string) uintptr {
// styleManager_StyleDescriptor.Call(GoStrToDStr(styleName))
// }
