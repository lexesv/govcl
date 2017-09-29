
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func Picture_Create() uintptr {
    ret, _, _ := picture_Create.Call()
    return ret
}

func Picture_Free(obj uintptr) {
    picture_Free.Call(obj)
}

func Picture_LoadFromFile(obj uintptr, Filename string)  {
    picture_LoadFromFile.Call(obj, GoStrToDStr(Filename) )
}

func Picture_SaveToFile(obj uintptr, Filename string)  {
    picture_SaveToFile.Call(obj, GoStrToDStr(Filename) )
}

func Picture_LoadFromStream(obj uintptr, Stream uintptr)  {
    picture_LoadFromStream.Call(obj, Stream )
}

func Picture_SaveToStream(obj uintptr, Stream uintptr)  {
    picture_SaveToStream.Call(obj, Stream )
}

func Picture_Assign(obj uintptr, Source uintptr)  {
    picture_Assign.Call(obj, Source )
}

func Picture_GetNamePath(obj uintptr) string {
    ret, _, _ := picture_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Picture_ClassName(obj uintptr) string {
    ret, _, _ := picture_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Picture_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := picture_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Picture_GetHashCode(obj uintptr) int32 {
    ret, _, _ := picture_GetHashCode.Call(obj)
    return int32(ret)
}

func Picture_ToString(obj uintptr) string {
    ret, _, _ := picture_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Picture_GetBitmap(obj uintptr) uintptr {
    ret, _, _ := picture_GetBitmap.Call(obj)
    return ret
}

func Picture_SetBitmap(obj uintptr, value uintptr) {
   picture_SetBitmap.Call(obj, value)
}

func Picture_GetGraphic(obj uintptr) uintptr {
    ret, _, _ := picture_GetGraphic.Call(obj)
    return ret
}

func Picture_SetGraphic(obj uintptr, value uintptr) {
   picture_SetGraphic.Call(obj, value)
}

func Picture_GetHeight(obj uintptr) int32 {
    ret, _, _ := picture_GetHeight.Call(obj)
    return int32(ret)
}

func Picture_GetIcon(obj uintptr) uintptr {
    ret, _, _ := picture_GetIcon.Call(obj)
    return ret
}

func Picture_SetIcon(obj uintptr, value uintptr) {
   picture_SetIcon.Call(obj, value)
}

func Picture_GetWidth(obj uintptr) int32 {
    ret, _, _ := picture_GetWidth.Call(obj)
    return int32(ret)
}

func Picture_SetOnChange(obj uintptr, fn interface{}) {
    picture_SetOnChange.Call(obj, addEventToMap(fn))
}

