
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func Timer_Create(obj uintptr) uintptr {
    ret, _, _ := timer_Create.Call(obj)
    return ret
}

func Timer_Free(obj uintptr) {
    timer_Free.Call(obj)
}

func Timer_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := timer_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Timer_GetNamePath(obj uintptr) string {
    ret, _, _ := timer_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Timer_HasParent(obj uintptr) bool {
    ret, _, _ := timer_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Timer_Assign(obj uintptr, Source uintptr)  {
    timer_Assign.Call(obj, Source )
}

func Timer_ClassName(obj uintptr) string {
    ret, _, _ := timer_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Timer_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := timer_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Timer_GetHashCode(obj uintptr) int32 {
    ret, _, _ := timer_GetHashCode.Call(obj)
    return int32(ret)
}

func Timer_ToString(obj uintptr) string {
    ret, _, _ := timer_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Timer_GetEnabled(obj uintptr) bool {
    ret, _, _ := timer_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Timer_SetEnabled(obj uintptr, value bool) {
   timer_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Timer_GetInterval(obj uintptr) uint32 {
    ret, _, _ := timer_GetInterval.Call(obj)
    return uint32(ret)
}

func Timer_SetInterval(obj uintptr, value uint32) {
   timer_SetInterval.Call(obj, uintptr(value))
}

func Timer_SetOnTimer(obj uintptr, fn interface{}) {
    timer_SetOnTimer.Call(obj, addEventToMap(fn))
}

func Timer_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := timer_GetComponentCount.Call(obj)
    return int32(ret)
}

func Timer_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := timer_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Timer_SetComponentIndex(obj uintptr, value int32) {
   timer_SetComponentIndex.Call(obj, uintptr(value))
}

func Timer_GetOwner(obj uintptr) uintptr {
    ret, _, _ := timer_GetOwner.Call(obj)
    return ret
}

func Timer_GetName(obj uintptr) string {
    ret, _, _ := timer_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Timer_SetName(obj uintptr, value string) {
   timer_SetName.Call(obj, GoStrToDStr(value))
}

func Timer_GetTag(obj uintptr) int {
    ret, _, _ := timer_GetTag.Call(obj)
    return int(ret)
}

func Timer_SetTag(obj uintptr, value int) {
   timer_SetTag.Call(obj, uintptr(value))
}

func Timer_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := timer_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

