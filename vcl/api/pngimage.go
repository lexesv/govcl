
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func PngImage_Create() uintptr {
    ret, _, _ := pngImage_Create.Call()
    return ret
}

func PngImage_Free(obj uintptr) {
    pngImage_Free.Call(obj)
}

func PngImage_Assign(obj uintptr, Source uintptr)  {
    pngImage_Assign.Call(obj, Source )
}

func PngImage_LoadFromStream(obj uintptr, Stream uintptr)  {
    pngImage_LoadFromStream.Call(obj, Stream )
}

func PngImage_SaveToStream(obj uintptr, Stream uintptr)  {
    pngImage_SaveToStream.Call(obj, Stream )
}

func PngImage_LoadFromResourceName(obj uintptr, Instance uintptr, Name string)  {
    pngImage_LoadFromResourceName.Call(obj, Instance , GoStrToDStr(Name) )
}

func PngImage_LoadFromResourceID(obj uintptr, Instance uintptr, ResID int32)  {
    pngImage_LoadFromResourceID.Call(obj, Instance , uintptr(ResID) )
}

func PngImage_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := pngImage_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func PngImage_LoadFromFile(obj uintptr, Filename string)  {
    pngImage_LoadFromFile.Call(obj, GoStrToDStr(Filename) )
}

func PngImage_SaveToFile(obj uintptr, Filename string)  {
    pngImage_SaveToFile.Call(obj, GoStrToDStr(Filename) )
}

func PngImage_SetSize(obj uintptr, AWidth int32, AHeight int32)  {
    pngImage_SetSize.Call(obj, uintptr(AWidth) , uintptr(AHeight) )
}

func PngImage_GetNamePath(obj uintptr) string {
    ret, _, _ := pngImage_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func PngImage_ClassName(obj uintptr) string {
    ret, _, _ := pngImage_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func PngImage_GetHashCode(obj uintptr) int32 {
    ret, _, _ := pngImage_GetHashCode.Call(obj)
    return int32(ret)
}

func PngImage_ToString(obj uintptr) string {
    ret, _, _ := pngImage_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func PngImage_GetTransparentColor(obj uintptr) TColor {
    ret, _, _ := pngImage_GetTransparentColor.Call(obj)
    return TColor(ret)
}

func PngImage_SetTransparentColor(obj uintptr, value TColor) {
   pngImage_SetTransparentColor.Call(obj, uintptr(value))
}

func PngImage_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := pngImage_GetCanvas.Call(obj)
    return ret
}

func PngImage_GetWidth(obj uintptr) int32 {
    ret, _, _ := pngImage_GetWidth.Call(obj)
    return int32(ret)
}

func PngImage_GetHeight(obj uintptr) int32 {
    ret, _, _ := pngImage_GetHeight.Call(obj)
    return int32(ret)
}

func PngImage_GetMaxIdatSize(obj uintptr) int32 {
    ret, _, _ := pngImage_GetMaxIdatSize.Call(obj)
    return int32(ret)
}

func PngImage_SetMaxIdatSize(obj uintptr, value int32) {
   pngImage_SetMaxIdatSize.Call(obj, uintptr(value))
}

func PngImage_GetEmpty(obj uintptr) bool {
    ret, _, _ := pngImage_GetEmpty.Call(obj)
    return DBoolToGoBool(ret)
}

func PngImage_GetCompressionLevel(obj uintptr) TCompressionLevel {
    ret, _, _ := pngImage_GetCompressionLevel.Call(obj)
    return TCompressionLevel(ret)
}

func PngImage_SetCompressionLevel(obj uintptr, value TCompressionLevel) {
   pngImage_SetCompressionLevel.Call(obj, uintptr(value))
}

func PngImage_GetVersion(obj uintptr) string {
    ret, _, _ := pngImage_GetVersion.Call(obj)
    return DStrToGoStr(ret)
}

func PngImage_GetModified(obj uintptr) bool {
    ret, _, _ := pngImage_GetModified.Call(obj)
    return DBoolToGoBool(ret)
}

func PngImage_SetModified(obj uintptr, value bool) {
   pngImage_SetModified.Call(obj, GoBoolToDBool(value))
}

func PngImage_GetPaletteModified(obj uintptr) bool {
    ret, _, _ := pngImage_GetPaletteModified.Call(obj)
    return DBoolToGoBool(ret)
}

func PngImage_SetPaletteModified(obj uintptr, value bool) {
   pngImage_SetPaletteModified.Call(obj, GoBoolToDBool(value))
}

func PngImage_GetTransparent(obj uintptr) bool {
    ret, _, _ := pngImage_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func PngImage_SetTransparent(obj uintptr, value bool) {
   pngImage_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func PngImage_SetOnChange(obj uintptr, fn interface{}) {
    pngImage_SetOnChange.Call(obj, addEventToMap(fn))
}

