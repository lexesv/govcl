
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func JPEGImage_Create() uintptr {
    ret, _, _ := jPEGImage_Create.Call()
    return ret
}

func JPEGImage_Free(obj uintptr) {
    jPEGImage_Free.Call(obj)
}

func JPEGImage_Assign(obj uintptr, Source uintptr)  {
    jPEGImage_Assign.Call(obj, Source )
}

func JPEGImage_LoadFromStream(obj uintptr, Stream uintptr)  {
    jPEGImage_LoadFromStream.Call(obj, Stream )
}

func JPEGImage_SaveToStream(obj uintptr, Stream uintptr)  {
    jPEGImage_SaveToStream.Call(obj, Stream )
}

func JPEGImage_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := jPEGImage_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func JPEGImage_LoadFromFile(obj uintptr, Filename string)  {
    jPEGImage_LoadFromFile.Call(obj, GoStrToDStr(Filename) )
}

func JPEGImage_SaveToFile(obj uintptr, Filename string)  {
    jPEGImage_SaveToFile.Call(obj, GoStrToDStr(Filename) )
}

func JPEGImage_SetSize(obj uintptr, AWidth int32, AHeight int32)  {
    jPEGImage_SetSize.Call(obj, uintptr(AWidth) , uintptr(AHeight) )
}

func JPEGImage_GetNamePath(obj uintptr) string {
    ret, _, _ := jPEGImage_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func JPEGImage_ClassName(obj uintptr) string {
    ret, _, _ := jPEGImage_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func JPEGImage_GetHashCode(obj uintptr) int32 {
    ret, _, _ := jPEGImage_GetHashCode.Call(obj)
    return int32(ret)
}

func JPEGImage_ToString(obj uintptr) string {
    ret, _, _ := jPEGImage_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func JPEGImage_GetPixelFormat(obj uintptr) TJPEGPixelFormat {
    ret, _, _ := jPEGImage_GetPixelFormat.Call(obj)
    return TJPEGPixelFormat(ret)
}

func JPEGImage_SetPixelFormat(obj uintptr, value TJPEGPixelFormat) {
   jPEGImage_SetPixelFormat.Call(obj, uintptr(value))
}

func JPEGImage_GetProgressiveDisplay(obj uintptr) bool {
    ret, _, _ := jPEGImage_GetProgressiveDisplay.Call(obj)
    return DBoolToGoBool(ret)
}

func JPEGImage_SetProgressiveDisplay(obj uintptr, value bool) {
   jPEGImage_SetProgressiveDisplay.Call(obj, GoBoolToDBool(value))
}

func JPEGImage_GetPerformance(obj uintptr) TJPEGPerformance {
    ret, _, _ := jPEGImage_GetPerformance.Call(obj)
    return TJPEGPerformance(ret)
}

func JPEGImage_SetPerformance(obj uintptr, value TJPEGPerformance) {
   jPEGImage_SetPerformance.Call(obj, uintptr(value))
}

func JPEGImage_GetScale(obj uintptr) TJPEGScale {
    ret, _, _ := jPEGImage_GetScale.Call(obj)
    return TJPEGScale(ret)
}

func JPEGImage_SetScale(obj uintptr, value TJPEGScale) {
   jPEGImage_SetScale.Call(obj, uintptr(value))
}

func JPEGImage_GetSmoothing(obj uintptr) bool {
    ret, _, _ := jPEGImage_GetSmoothing.Call(obj)
    return DBoolToGoBool(ret)
}

func JPEGImage_SetSmoothing(obj uintptr, value bool) {
   jPEGImage_SetSmoothing.Call(obj, GoBoolToDBool(value))
}

func JPEGImage_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := jPEGImage_GetCanvas.Call(obj)
    return ret
}

func JPEGImage_GetEmpty(obj uintptr) bool {
    ret, _, _ := jPEGImage_GetEmpty.Call(obj)
    return DBoolToGoBool(ret)
}

func JPEGImage_GetHeight(obj uintptr) int32 {
    ret, _, _ := jPEGImage_GetHeight.Call(obj)
    return int32(ret)
}

func JPEGImage_SetHeight(obj uintptr, value int32) {
   jPEGImage_SetHeight.Call(obj, uintptr(value))
}

func JPEGImage_GetModified(obj uintptr) bool {
    ret, _, _ := jPEGImage_GetModified.Call(obj)
    return DBoolToGoBool(ret)
}

func JPEGImage_SetModified(obj uintptr, value bool) {
   jPEGImage_SetModified.Call(obj, GoBoolToDBool(value))
}

func JPEGImage_GetPaletteModified(obj uintptr) bool {
    ret, _, _ := jPEGImage_GetPaletteModified.Call(obj)
    return DBoolToGoBool(ret)
}

func JPEGImage_SetPaletteModified(obj uintptr, value bool) {
   jPEGImage_SetPaletteModified.Call(obj, GoBoolToDBool(value))
}

func JPEGImage_GetTransparent(obj uintptr) bool {
    ret, _, _ := jPEGImage_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func JPEGImage_SetTransparent(obj uintptr, value bool) {
   jPEGImage_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func JPEGImage_GetWidth(obj uintptr) int32 {
    ret, _, _ := jPEGImage_GetWidth.Call(obj)
    return int32(ret)
}

func JPEGImage_SetWidth(obj uintptr, value int32) {
   jPEGImage_SetWidth.Call(obj, uintptr(value))
}

func JPEGImage_SetOnChange(obj uintptr, fn interface{}) {
    jPEGImage_SetOnChange.Call(obj, addEventToMap(fn))
}

