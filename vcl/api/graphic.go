
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func Graphic_Create() uintptr {
    ret, _, _ := graphic_Create.Call()
    return ret
}

func Graphic_Free(obj uintptr) {
    graphic_Free.Call(obj)
}

func Graphic_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := graphic_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Graphic_LoadFromFile(obj uintptr, Filename string)  {
    graphic_LoadFromFile.Call(obj, GoStrToDStr(Filename) )
}

func Graphic_SaveToFile(obj uintptr, Filename string)  {
    graphic_SaveToFile.Call(obj, GoStrToDStr(Filename) )
}

func Graphic_LoadFromStream(obj uintptr, Stream uintptr)  {
    graphic_LoadFromStream.Call(obj, Stream )
}

func Graphic_SaveToStream(obj uintptr, Stream uintptr)  {
    graphic_SaveToStream.Call(obj, Stream )
}

func Graphic_SetSize(obj uintptr, AWidth int32, AHeight int32)  {
    graphic_SetSize.Call(obj, uintptr(AWidth) , uintptr(AHeight) )
}

func Graphic_Assign(obj uintptr, Source uintptr)  {
    graphic_Assign.Call(obj, Source )
}

func Graphic_GetNamePath(obj uintptr) string {
    ret, _, _ := graphic_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Graphic_ClassName(obj uintptr) string {
    ret, _, _ := graphic_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Graphic_GetHashCode(obj uintptr) int32 {
    ret, _, _ := graphic_GetHashCode.Call(obj)
    return int32(ret)
}

func Graphic_ToString(obj uintptr) string {
    ret, _, _ := graphic_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Graphic_GetEmpty(obj uintptr) bool {
    ret, _, _ := graphic_GetEmpty.Call(obj)
    return DBoolToGoBool(ret)
}

func Graphic_GetHeight(obj uintptr) int32 {
    ret, _, _ := graphic_GetHeight.Call(obj)
    return int32(ret)
}

func Graphic_SetHeight(obj uintptr, value int32) {
   graphic_SetHeight.Call(obj, uintptr(value))
}

func Graphic_GetModified(obj uintptr) bool {
    ret, _, _ := graphic_GetModified.Call(obj)
    return DBoolToGoBool(ret)
}

func Graphic_SetModified(obj uintptr, value bool) {
   graphic_SetModified.Call(obj, GoBoolToDBool(value))
}

func Graphic_GetPaletteModified(obj uintptr) bool {
    ret, _, _ := graphic_GetPaletteModified.Call(obj)
    return DBoolToGoBool(ret)
}

func Graphic_SetPaletteModified(obj uintptr, value bool) {
   graphic_SetPaletteModified.Call(obj, GoBoolToDBool(value))
}

func Graphic_GetTransparent(obj uintptr) bool {
    ret, _, _ := graphic_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func Graphic_SetTransparent(obj uintptr, value bool) {
   graphic_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func Graphic_GetWidth(obj uintptr) int32 {
    ret, _, _ := graphic_GetWidth.Call(obj)
    return int32(ret)
}

func Graphic_SetWidth(obj uintptr, value int32) {
   graphic_SetWidth.Call(obj, uintptr(value))
}

func Graphic_SetOnChange(obj uintptr, fn interface{}) {
    graphic_SetOnChange.Call(obj, addEventToMap(fn))
}

