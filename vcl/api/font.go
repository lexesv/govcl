
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func Font_Create() uintptr {
    ret, _, _ := font_Create.Call()
    return ret
}

func Font_Free(obj uintptr) {
    font_Free.Call(obj)
}

func Font_Assign(obj uintptr, Source uintptr)  {
    font_Assign.Call(obj, Source )
}

func Font_HandleAllocated(obj uintptr) bool {
    ret, _, _ := font_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Font_GetNamePath(obj uintptr) string {
    ret, _, _ := font_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Font_ClassName(obj uintptr) string {
    ret, _, _ := font_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Font_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := font_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Font_GetHashCode(obj uintptr) int32 {
    ret, _, _ := font_GetHashCode.Call(obj)
    return int32(ret)
}

func Font_ToString(obj uintptr) string {
    ret, _, _ := font_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Font_GetHandle(obj uintptr) HFONT {
    ret, _, _ := font_GetHandle.Call(obj)
    return HFONT(ret)
}

func Font_SetHandle(obj uintptr, value HFONT) {
   font_SetHandle.Call(obj, uintptr(value))
}

func Font_GetPixelsPerInch(obj uintptr) int32 {
    ret, _, _ := font_GetPixelsPerInch.Call(obj)
    return int32(ret)
}

func Font_SetPixelsPerInch(obj uintptr, value int32) {
   font_SetPixelsPerInch.Call(obj, uintptr(value))
}

func Font_GetCharset(obj uintptr) TFontCharset {
    ret, _, _ := font_GetCharset.Call(obj)
    return TFontCharset(ret)
}

func Font_SetCharset(obj uintptr, value TFontCharset) {
   font_SetCharset.Call(obj, uintptr(value))
}

func Font_GetColor(obj uintptr) TColor {
    ret, _, _ := font_GetColor.Call(obj)
    return TColor(ret)
}

func Font_SetColor(obj uintptr, value TColor) {
   font_SetColor.Call(obj, uintptr(value))
}

func Font_GetHeight(obj uintptr) int32 {
    ret, _, _ := font_GetHeight.Call(obj)
    return int32(ret)
}

func Font_SetHeight(obj uintptr, value int32) {
   font_SetHeight.Call(obj, uintptr(value))
}

func Font_GetName(obj uintptr) string {
    ret, _, _ := font_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Font_SetName(obj uintptr, value string) {
   font_SetName.Call(obj, GoStrToDStr(value))
}

func Font_GetOrientation(obj uintptr) int32 {
    ret, _, _ := font_GetOrientation.Call(obj)
    return int32(ret)
}

func Font_SetOrientation(obj uintptr, value int32) {
   font_SetOrientation.Call(obj, uintptr(value))
}

func Font_GetPitch(obj uintptr) TFontPitch {
    ret, _, _ := font_GetPitch.Call(obj)
    return TFontPitch(ret)
}

func Font_SetPitch(obj uintptr, value TFontPitch) {
   font_SetPitch.Call(obj, uintptr(value))
}

func Font_GetSize(obj uintptr) int32 {
    ret, _, _ := font_GetSize.Call(obj)
    return int32(ret)
}

func Font_SetSize(obj uintptr, value int32) {
   font_SetSize.Call(obj, uintptr(value))
}

func Font_GetStyle(obj uintptr) TFontStyles {
    ret, _, _ := font_GetStyle.Call(obj)
    return TFontStyles(ret)
}

func Font_SetStyle(obj uintptr, value TFontStyles) {
   font_SetStyle.Call(obj, uintptr(value))
}

func Font_GetQuality(obj uintptr) TFontQuality {
    ret, _, _ := font_GetQuality.Call(obj)
    return TFontQuality(ret)
}

func Font_SetQuality(obj uintptr, value TFontQuality) {
   font_SetQuality.Call(obj, uintptr(value))
}

func Font_SetOnChange(obj uintptr, fn interface{}) {
    font_SetOnChange.Call(obj, addEventToMap(fn))
}

