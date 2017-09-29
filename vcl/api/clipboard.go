
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func Clipboard_Create() uintptr {
    ret, _, _ := clipboard_Create.Call()
    return ret
}

func Clipboard_Free(obj uintptr) {
    clipboard_Free.Call(obj)
}

func Clipboard_Assign(obj uintptr, Source uintptr)  {
    clipboard_Assign.Call(obj, Source )
}

func Clipboard_Clear(obj uintptr)  {
    clipboard_Clear.Call(obj)
}

func Clipboard_Close(obj uintptr)  {
    clipboard_Close.Call(obj)
}

func Clipboard_GetAsHandle(obj uintptr, Format uint16) uintptr {
    ret, _, _ := clipboard_GetAsHandle.Call(obj, uintptr(Format) )
    return ret
}

func Clipboard_HasFormat(obj uintptr, Format uint16) bool {
    ret, _, _ := clipboard_HasFormat.Call(obj, uintptr(Format) )
    return DBoolToGoBool(ret)
}

func Clipboard_Open(obj uintptr)  {
    clipboard_Open.Call(obj)
}

func Clipboard_SetAsHandle(obj uintptr, Format uint16, Value uintptr)  {
    clipboard_SetAsHandle.Call(obj, uintptr(Format) , Value )
}

func Clipboard_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := clipboard_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Clipboard_GetNamePath(obj uintptr) string {
    ret, _, _ := clipboard_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Clipboard_ClassName(obj uintptr) string {
    ret, _, _ := clipboard_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Clipboard_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := clipboard_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Clipboard_GetHashCode(obj uintptr) int32 {
    ret, _, _ := clipboard_GetHashCode.Call(obj)
    return int32(ret)
}

func Clipboard_ToString(obj uintptr) string {
    ret, _, _ := clipboard_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Clipboard_GetAsText(obj uintptr) string {
    ret, _, _ := clipboard_GetAsText.Call(obj)
    return DStrToGoStr(ret)
}

func Clipboard_SetAsText(obj uintptr, value string) {
   clipboard_SetAsText.Call(obj, GoStrToDStr(value))
}

func Clipboard_GetFormatCount(obj uintptr) int32 {
    ret, _, _ := clipboard_GetFormatCount.Call(obj)
    return int32(ret)
}

func Clipboard_GetFormats(obj uintptr, Index int32) uint16 {
    ret, _, _ := clipboard_GetFormats.Call(obj, uintptr(Index))
    return uint16(ret)
}

