
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func ActionList_Create(obj uintptr) uintptr {
    ret, _, _ := actionList_Create.Call(obj)
    return ret
}

func ActionList_Free(obj uintptr) {
    actionList_Free.Call(obj)
}

func ActionList_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := actionList_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ActionList_GetNamePath(obj uintptr) string {
    ret, _, _ := actionList_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ActionList_HasParent(obj uintptr) bool {
    ret, _, _ := actionList_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ActionList_Assign(obj uintptr, Source uintptr)  {
    actionList_Assign.Call(obj, Source )
}

func ActionList_ClassName(obj uintptr) string {
    ret, _, _ := actionList_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ActionList_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := actionList_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ActionList_GetHashCode(obj uintptr) int32 {
    ret, _, _ := actionList_GetHashCode.Call(obj)
    return int32(ret)
}

func ActionList_ToString(obj uintptr) string {
    ret, _, _ := actionList_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ActionList_GetImages(obj uintptr) uintptr {
    ret, _, _ := actionList_GetImages.Call(obj)
    return ret
}

func ActionList_SetImages(obj uintptr, value uintptr) {
   actionList_SetImages.Call(obj, value)
}

func ActionList_GetState(obj uintptr) TActionListState {
    ret, _, _ := actionList_GetState.Call(obj)
    return TActionListState(ret)
}

func ActionList_SetState(obj uintptr, value TActionListState) {
   actionList_SetState.Call(obj, uintptr(value))
}

func ActionList_SetOnChange(obj uintptr, fn interface{}) {
    actionList_SetOnChange.Call(obj, addEventToMap(fn))
}

func ActionList_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := actionList_GetComponentCount.Call(obj)
    return int32(ret)
}

func ActionList_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := actionList_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ActionList_SetComponentIndex(obj uintptr, value int32) {
   actionList_SetComponentIndex.Call(obj, uintptr(value))
}

func ActionList_GetOwner(obj uintptr) uintptr {
    ret, _, _ := actionList_GetOwner.Call(obj)
    return ret
}

func ActionList_GetName(obj uintptr) string {
    ret, _, _ := actionList_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ActionList_SetName(obj uintptr, value string) {
   actionList_SetName.Call(obj, GoStrToDStr(value))
}

func ActionList_GetTag(obj uintptr) int {
    ret, _, _ := actionList_GetTag.Call(obj)
    return int(ret)
}

func ActionList_SetTag(obj uintptr, value int) {
   actionList_SetTag.Call(obj, uintptr(value))
}

func ActionList_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := actionList_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

