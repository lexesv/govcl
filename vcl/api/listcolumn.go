
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func ListColumn_Create() uintptr {
    ret, _, _ := listColumn_Create.Call()
    return ret
}

func ListColumn_Free(obj uintptr) {
    listColumn_Free.Call(obj)
}

func ListColumn_Assign(obj uintptr, Source uintptr)  {
    listColumn_Assign.Call(obj, Source )
}

func ListColumn_GetNamePath(obj uintptr) string {
    ret, _, _ := listColumn_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ListColumn_ClassName(obj uintptr) string {
    ret, _, _ := listColumn_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ListColumn_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := listColumn_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ListColumn_GetHashCode(obj uintptr) int32 {
    ret, _, _ := listColumn_GetHashCode.Call(obj)
    return int32(ret)
}

func ListColumn_ToString(obj uintptr) string {
    ret, _, _ := listColumn_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ListColumn_GetAlignment(obj uintptr) TAlignment {
    ret, _, _ := listColumn_GetAlignment.Call(obj)
    return TAlignment(ret)
}

func ListColumn_SetAlignment(obj uintptr, value TAlignment) {
   listColumn_SetAlignment.Call(obj, uintptr(value))
}

func ListColumn_GetAutoSize(obj uintptr) bool {
    ret, _, _ := listColumn_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func ListColumn_SetAutoSize(obj uintptr, value bool) {
   listColumn_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func ListColumn_GetCaption(obj uintptr) string {
    ret, _, _ := listColumn_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func ListColumn_SetCaption(obj uintptr, value string) {
   listColumn_SetCaption.Call(obj, GoStrToDStr(value))
}

func ListColumn_GetImageIndex(obj uintptr) int32 {
    ret, _, _ := listColumn_GetImageIndex.Call(obj)
    return int32(ret)
}

func ListColumn_SetImageIndex(obj uintptr, value int32) {
   listColumn_SetImageIndex.Call(obj, uintptr(value))
}

func ListColumn_GetTag(obj uintptr) int32 {
    ret, _, _ := listColumn_GetTag.Call(obj)
    return int32(ret)
}

func ListColumn_SetTag(obj uintptr, value int32) {
   listColumn_SetTag.Call(obj, uintptr(value))
}

func ListColumn_GetWidth(obj uintptr) int32 {
    ret, _, _ := listColumn_GetWidth.Call(obj)
    return int32(ret)
}

func ListColumn_SetWidth(obj uintptr, value int32) {
   listColumn_SetWidth.Call(obj, uintptr(value))
}

func ListColumn_GetIndex(obj uintptr) int32 {
    ret, _, _ := listColumn_GetIndex.Call(obj)
    return int32(ret)
}

func ListColumn_SetIndex(obj uintptr, value int32) {
   listColumn_SetIndex.Call(obj, uintptr(value))
}

