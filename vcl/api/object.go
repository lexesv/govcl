
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func Object_Create() uintptr {
    ret, _, _ := object_Create.Call()
    return ret
}

func Object_Free(obj uintptr) {
    object_Free.Call(obj)
}

func Object_ClassName(obj uintptr) string {
    ret, _, _ := object_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Object_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := object_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Object_GetHashCode(obj uintptr) int32 {
    ret, _, _ := object_GetHashCode.Call(obj)
    return int32(ret)
}

func Object_ToString(obj uintptr) string {
    ret, _, _ := object_ToString.Call(obj)
    return DStrToGoStr(ret)
}

