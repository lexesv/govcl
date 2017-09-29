
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func Icon_Create() uintptr {
    ret, _, _ := icon_Create.Call()
    return ret
}

func Icon_Free(obj uintptr) {
    icon_Free.Call(obj)
}

func Icon_Assign(obj uintptr, Source uintptr)  {
    icon_Assign.Call(obj, Source )
}

func Icon_HandleAllocated(obj uintptr) bool {
    ret, _, _ := icon_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Icon_LoadFromStream(obj uintptr, Stream uintptr)  {
    icon_LoadFromStream.Call(obj, Stream )
}

func Icon_SaveToStream(obj uintptr, Stream uintptr)  {
    icon_SaveToStream.Call(obj, Stream )
}

func Icon_SetSize(obj uintptr, AWidth int32, AHeight int32)  {
    icon_SetSize.Call(obj, uintptr(AWidth) , uintptr(AHeight) )
}

func Icon_LoadFromResourceName(obj uintptr, Instance uintptr, ResName string)  {
    icon_LoadFromResourceName.Call(obj, Instance , GoStrToDStr(ResName) )
}

func Icon_LoadFromResourceID(obj uintptr, Instance uintptr, ResID int32)  {
    icon_LoadFromResourceID.Call(obj, Instance , uintptr(ResID) )
}

func Icon_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := icon_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Icon_LoadFromFile(obj uintptr, Filename string)  {
    icon_LoadFromFile.Call(obj, GoStrToDStr(Filename) )
}

func Icon_SaveToFile(obj uintptr, Filename string)  {
    icon_SaveToFile.Call(obj, GoStrToDStr(Filename) )
}

func Icon_GetNamePath(obj uintptr) string {
    ret, _, _ := icon_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Icon_ClassName(obj uintptr) string {
    ret, _, _ := icon_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Icon_GetHashCode(obj uintptr) int32 {
    ret, _, _ := icon_GetHashCode.Call(obj)
    return int32(ret)
}

func Icon_ToString(obj uintptr) string {
    ret, _, _ := icon_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Icon_GetHandle(obj uintptr) HICON {
    ret, _, _ := icon_GetHandle.Call(obj)
    return HICON(ret)
}

func Icon_SetHandle(obj uintptr, value HICON) {
   icon_SetHandle.Call(obj, uintptr(value))
}

func Icon_GetEmpty(obj uintptr) bool {
    ret, _, _ := icon_GetEmpty.Call(obj)
    return DBoolToGoBool(ret)
}

func Icon_GetHeight(obj uintptr) int32 {
    ret, _, _ := icon_GetHeight.Call(obj)
    return int32(ret)
}

func Icon_SetHeight(obj uintptr, value int32) {
   icon_SetHeight.Call(obj, uintptr(value))
}

func Icon_GetModified(obj uintptr) bool {
    ret, _, _ := icon_GetModified.Call(obj)
    return DBoolToGoBool(ret)
}

func Icon_SetModified(obj uintptr, value bool) {
   icon_SetModified.Call(obj, GoBoolToDBool(value))
}

func Icon_GetPaletteModified(obj uintptr) bool {
    ret, _, _ := icon_GetPaletteModified.Call(obj)
    return DBoolToGoBool(ret)
}

func Icon_SetPaletteModified(obj uintptr, value bool) {
   icon_SetPaletteModified.Call(obj, GoBoolToDBool(value))
}

func Icon_GetTransparent(obj uintptr) bool {
    ret, _, _ := icon_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func Icon_SetTransparent(obj uintptr, value bool) {
   icon_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func Icon_GetWidth(obj uintptr) int32 {
    ret, _, _ := icon_GetWidth.Call(obj)
    return int32(ret)
}

func Icon_SetWidth(obj uintptr, value int32) {
   icon_SetWidth.Call(obj, uintptr(value))
}

func Icon_SetOnChange(obj uintptr, fn interface{}) {
    icon_SetOnChange.Call(obj, addEventToMap(fn))
}

