
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func Component_Create(obj uintptr) uintptr {
    ret, _, _ := component_Create.Call(obj)
    return ret
}

func Component_Free(obj uintptr) {
    component_Free.Call(obj)
}

func Component_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := component_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Component_GetNamePath(obj uintptr) string {
    ret, _, _ := component_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Component_HasParent(obj uintptr) bool {
    ret, _, _ := component_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Component_Assign(obj uintptr, Source uintptr)  {
    component_Assign.Call(obj, Source )
}

func Component_ClassName(obj uintptr) string {
    ret, _, _ := component_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Component_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := component_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Component_GetHashCode(obj uintptr) int32 {
    ret, _, _ := component_GetHashCode.Call(obj)
    return int32(ret)
}

func Component_ToString(obj uintptr) string {
    ret, _, _ := component_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Component_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := component_GetComponentCount.Call(obj)
    return int32(ret)
}

func Component_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := component_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Component_SetComponentIndex(obj uintptr, value int32) {
   component_SetComponentIndex.Call(obj, uintptr(value))
}

func Component_GetOwner(obj uintptr) uintptr {
    ret, _, _ := component_GetOwner.Call(obj)
    return ret
}

func Component_GetName(obj uintptr) string {
    ret, _, _ := component_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Component_SetName(obj uintptr, value string) {
   component_SetName.Call(obj, GoStrToDStr(value))
}

func Component_GetTag(obj uintptr) int {
    ret, _, _ := component_GetTag.Call(obj)
    return int(ret)
}

func Component_SetTag(obj uintptr, value int) {
   component_SetTag.Call(obj, uintptr(value))
}

func Component_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := component_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

