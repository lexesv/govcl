
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func GIFFrame_Create() uintptr {
    ret, _, _ := gIFFrame_Create.Call()
    return ret
}

func GIFFrame_Free(obj uintptr) {
    gIFFrame_Free.Call(obj)
}

func GIFFrame_Clear(obj uintptr)  {
    gIFFrame_Clear.Call(obj)
}

func GIFFrame_SaveToStream(obj uintptr, Stream uintptr)  {
    gIFFrame_SaveToStream.Call(obj, Stream )
}

func GIFFrame_LoadFromStream(obj uintptr, Stream uintptr)  {
    gIFFrame_LoadFromStream.Call(obj, Stream )
}

func GIFFrame_Assign(obj uintptr, Source uintptr)  {
    gIFFrame_Assign.Call(obj, Source )
}

func GIFFrame_SaveToFile(obj uintptr, Filename string)  {
    gIFFrame_SaveToFile.Call(obj, GoStrToDStr(Filename) )
}

func GIFFrame_LoadFromFile(obj uintptr, Filename string)  {
    gIFFrame_LoadFromFile.Call(obj, GoStrToDStr(Filename) )
}

func GIFFrame_GetNamePath(obj uintptr) string {
    ret, _, _ := gIFFrame_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func GIFFrame_ClassName(obj uintptr) string {
    ret, _, _ := gIFFrame_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func GIFFrame_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := gIFFrame_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func GIFFrame_GetHashCode(obj uintptr) int32 {
    ret, _, _ := gIFFrame_GetHashCode.Call(obj)
    return int32(ret)
}

func GIFFrame_ToString(obj uintptr) string {
    ret, _, _ := gIFFrame_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func GIFFrame_GetHasBitmap(obj uintptr) bool {
    ret, _, _ := gIFFrame_GetHasBitmap.Call(obj)
    return DBoolToGoBool(ret)
}

func GIFFrame_SetHasBitmap(obj uintptr, value bool) {
   gIFFrame_SetHasBitmap.Call(obj, GoBoolToDBool(value))
}

func GIFFrame_GetLeft(obj uintptr) uint16 {
    ret, _, _ := gIFFrame_GetLeft.Call(obj)
    return uint16(ret)
}

func GIFFrame_SetLeft(obj uintptr, value uint16) {
   gIFFrame_SetLeft.Call(obj, uintptr(value))
}

func GIFFrame_GetTop(obj uintptr) uint16 {
    ret, _, _ := gIFFrame_GetTop.Call(obj)
    return uint16(ret)
}

func GIFFrame_SetTop(obj uintptr, value uint16) {
   gIFFrame_SetTop.Call(obj, uintptr(value))
}

func GIFFrame_GetWidth(obj uintptr) uint16 {
    ret, _, _ := gIFFrame_GetWidth.Call(obj)
    return uint16(ret)
}

func GIFFrame_SetWidth(obj uintptr, value uint16) {
   gIFFrame_SetWidth.Call(obj, uintptr(value))
}

func GIFFrame_GetHeight(obj uintptr) uint16 {
    ret, _, _ := gIFFrame_GetHeight.Call(obj)
    return uint16(ret)
}

func GIFFrame_SetHeight(obj uintptr, value uint16) {
   gIFFrame_SetHeight.Call(obj, uintptr(value))
}

func GIFFrame_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    gIFFrame_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func GIFFrame_SetBoundsRect(obj uintptr, value TRect) {
   gIFFrame_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func GIFFrame_GetClientRect(obj uintptr) TRect {
    var ret TRect
    gIFFrame_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func GIFFrame_GetData(obj uintptr) uintptr {
    ret, _, _ := gIFFrame_GetData.Call(obj)
    return ret
}

func GIFFrame_GetDataSize(obj uintptr) int32 {
    ret, _, _ := gIFFrame_GetDataSize.Call(obj)
    return int32(ret)
}

func GIFFrame_GetVersion(obj uintptr) TGIFVersion {
    ret, _, _ := gIFFrame_GetVersion.Call(obj)
    return TGIFVersion(ret)
}

func GIFFrame_GetBitsPerPixel(obj uintptr) int32 {
    ret, _, _ := gIFFrame_GetBitsPerPixel.Call(obj)
    return int32(ret)
}

func GIFFrame_GetBitmap(obj uintptr) uintptr {
    ret, _, _ := gIFFrame_GetBitmap.Call(obj)
    return ret
}

func GIFFrame_SetBitmap(obj uintptr, value uintptr) {
   gIFFrame_SetBitmap.Call(obj, value)
}

func GIFFrame_GetEmpty(obj uintptr) bool {
    ret, _, _ := gIFFrame_GetEmpty.Call(obj)
    return DBoolToGoBool(ret)
}

func GIFFrame_GetTransparent(obj uintptr) bool {
    ret, _, _ := gIFFrame_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

