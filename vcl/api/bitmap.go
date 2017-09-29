
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func Bitmap_Create() uintptr {
    ret, _, _ := bitmap_Create.Call()
    return ret
}

func Bitmap_Free(obj uintptr) {
    bitmap_Free.Call(obj)
}

func Bitmap_Assign(obj uintptr, Source uintptr)  {
    bitmap_Assign.Call(obj, Source )
}

func Bitmap_HandleAllocated(obj uintptr) bool {
    ret, _, _ := bitmap_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Bitmap_LoadFromStream(obj uintptr, Stream uintptr)  {
    bitmap_LoadFromStream.Call(obj, Stream )
}

func Bitmap_SaveToStream(obj uintptr, Stream uintptr)  {
    bitmap_SaveToStream.Call(obj, Stream )
}

func Bitmap_SetSize(obj uintptr, AWidth int32, AHeight int32)  {
    bitmap_SetSize.Call(obj, uintptr(AWidth) , uintptr(AHeight) )
}

func Bitmap_LoadFromResourceName(obj uintptr, Instance uintptr, ResName string)  {
    bitmap_LoadFromResourceName.Call(obj, Instance , GoStrToDStr(ResName) )
}

func Bitmap_LoadFromResourceID(obj uintptr, Instance uintptr, ResID int32)  {
    bitmap_LoadFromResourceID.Call(obj, Instance , uintptr(ResID) )
}

func Bitmap_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := bitmap_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Bitmap_LoadFromFile(obj uintptr, Filename string)  {
    bitmap_LoadFromFile.Call(obj, GoStrToDStr(Filename) )
}

func Bitmap_SaveToFile(obj uintptr, Filename string)  {
    bitmap_SaveToFile.Call(obj, GoStrToDStr(Filename) )
}

func Bitmap_GetNamePath(obj uintptr) string {
    ret, _, _ := bitmap_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Bitmap_ClassName(obj uintptr) string {
    ret, _, _ := bitmap_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Bitmap_GetHashCode(obj uintptr) int32 {
    ret, _, _ := bitmap_GetHashCode.Call(obj)
    return int32(ret)
}

func Bitmap_ToString(obj uintptr) string {
    ret, _, _ := bitmap_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Bitmap_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := bitmap_GetCanvas.Call(obj)
    return ret
}

func Bitmap_GetHandle(obj uintptr) HBITMAP {
    ret, _, _ := bitmap_GetHandle.Call(obj)
    return HBITMAP(ret)
}

func Bitmap_SetHandle(obj uintptr, value HBITMAP) {
   bitmap_SetHandle.Call(obj, uintptr(value))
}

func Bitmap_GetPixelFormat(obj uintptr) TPixelFormat {
    ret, _, _ := bitmap_GetPixelFormat.Call(obj)
    return TPixelFormat(ret)
}

func Bitmap_SetPixelFormat(obj uintptr, value TPixelFormat) {
   bitmap_SetPixelFormat.Call(obj, uintptr(value))
}

func Bitmap_GetTransparentColor(obj uintptr) TColor {
    ret, _, _ := bitmap_GetTransparentColor.Call(obj)
    return TColor(ret)
}

func Bitmap_SetTransparentColor(obj uintptr, value TColor) {
   bitmap_SetTransparentColor.Call(obj, uintptr(value))
}

func Bitmap_GetEmpty(obj uintptr) bool {
    ret, _, _ := bitmap_GetEmpty.Call(obj)
    return DBoolToGoBool(ret)
}

func Bitmap_GetHeight(obj uintptr) int32 {
    ret, _, _ := bitmap_GetHeight.Call(obj)
    return int32(ret)
}

func Bitmap_SetHeight(obj uintptr, value int32) {
   bitmap_SetHeight.Call(obj, uintptr(value))
}

func Bitmap_GetModified(obj uintptr) bool {
    ret, _, _ := bitmap_GetModified.Call(obj)
    return DBoolToGoBool(ret)
}

func Bitmap_SetModified(obj uintptr, value bool) {
   bitmap_SetModified.Call(obj, GoBoolToDBool(value))
}

func Bitmap_GetPaletteModified(obj uintptr) bool {
    ret, _, _ := bitmap_GetPaletteModified.Call(obj)
    return DBoolToGoBool(ret)
}

func Bitmap_SetPaletteModified(obj uintptr, value bool) {
   bitmap_SetPaletteModified.Call(obj, GoBoolToDBool(value))
}

func Bitmap_GetTransparent(obj uintptr) bool {
    ret, _, _ := bitmap_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func Bitmap_SetTransparent(obj uintptr, value bool) {
   bitmap_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func Bitmap_GetWidth(obj uintptr) int32 {
    ret, _, _ := bitmap_GetWidth.Call(obj)
    return int32(ret)
}

func Bitmap_SetWidth(obj uintptr, value int32) {
   bitmap_SetWidth.Call(obj, uintptr(value))
}

func Bitmap_SetOnChange(obj uintptr, fn interface{}) {
    bitmap_SetOnChange.Call(obj, addEventToMap(fn))
}

func Bitmap_GetScanLine(obj uintptr, Row int32) uintptr {
    ret, _, _ := bitmap_GetScanLine.Call(obj, uintptr(Row))
    return ret
}

