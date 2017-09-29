
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func Action_Create(obj uintptr) uintptr {
    ret, _, _ := action_Create.Call(obj)
    return ret
}

func Action_Free(obj uintptr) {
    action_Free.Call(obj)
}

func Action_Execute(obj uintptr) bool {
    ret, _, _ := action_Execute.Call(obj)
    return DBoolToGoBool(ret)
}

func Action_Update(obj uintptr) bool {
    ret, _, _ := action_Update.Call(obj)
    return DBoolToGoBool(ret)
}

func Action_HasParent(obj uintptr) bool {
    ret, _, _ := action_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Action_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := action_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Action_GetNamePath(obj uintptr) string {
    ret, _, _ := action_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Action_Assign(obj uintptr, Source uintptr)  {
    action_Assign.Call(obj, Source )
}

func Action_ClassName(obj uintptr) string {
    ret, _, _ := action_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Action_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := action_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Action_GetHashCode(obj uintptr) int32 {
    ret, _, _ := action_GetHashCode.Call(obj)
    return int32(ret)
}

func Action_ToString(obj uintptr) string {
    ret, _, _ := action_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Action_GetCaption(obj uintptr) string {
    ret, _, _ := action_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func Action_SetCaption(obj uintptr, value string) {
   action_SetCaption.Call(obj, GoStrToDStr(value))
}

func Action_GetChecked(obj uintptr) bool {
    ret, _, _ := action_GetChecked.Call(obj)
    return DBoolToGoBool(ret)
}

func Action_SetChecked(obj uintptr, value bool) {
   action_SetChecked.Call(obj, GoBoolToDBool(value))
}

func Action_GetEnabled(obj uintptr) bool {
    ret, _, _ := action_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Action_SetEnabled(obj uintptr, value bool) {
   action_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Action_GetGroupIndex(obj uintptr) int32 {
    ret, _, _ := action_GetGroupIndex.Call(obj)
    return int32(ret)
}

func Action_SetGroupIndex(obj uintptr, value int32) {
   action_SetGroupIndex.Call(obj, uintptr(value))
}

func Action_GetHint(obj uintptr) string {
    ret, _, _ := action_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Action_SetHint(obj uintptr, value string) {
   action_SetHint.Call(obj, GoStrToDStr(value))
}

func Action_GetImageIndex(obj uintptr) int32 {
    ret, _, _ := action_GetImageIndex.Call(obj)
    return int32(ret)
}

func Action_SetImageIndex(obj uintptr, value int32) {
   action_SetImageIndex.Call(obj, uintptr(value))
}

func Action_GetShortCut(obj uintptr) TShortCut {
    ret, _, _ := action_GetShortCut.Call(obj)
    return TShortCut(ret)
}

func Action_SetShortCut(obj uintptr, value TShortCut) {
   action_SetShortCut.Call(obj, uintptr(value))
}

func Action_GetVisible(obj uintptr) bool {
    ret, _, _ := action_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Action_SetVisible(obj uintptr, value bool) {
   action_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Action_SetOnExecute(obj uintptr, fn interface{}) {
    action_SetOnExecute.Call(obj, addEventToMap(fn))
}

func Action_SetOnUpdate(obj uintptr, fn interface{}) {
    action_SetOnUpdate.Call(obj, addEventToMap(fn))
}

func Action_GetImages(obj uintptr) uintptr {
    ret, _, _ := action_GetImages.Call(obj)
    return ret
}

func Action_GetIndex(obj uintptr) int32 {
    ret, _, _ := action_GetIndex.Call(obj)
    return int32(ret)
}

func Action_SetIndex(obj uintptr, value int32) {
   action_SetIndex.Call(obj, uintptr(value))
}

func Action_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := action_GetComponentCount.Call(obj)
    return int32(ret)
}

func Action_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := action_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Action_SetComponentIndex(obj uintptr, value int32) {
   action_SetComponentIndex.Call(obj, uintptr(value))
}

func Action_GetOwner(obj uintptr) uintptr {
    ret, _, _ := action_GetOwner.Call(obj)
    return ret
}

func Action_GetName(obj uintptr) string {
    ret, _, _ := action_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Action_SetName(obj uintptr, value string) {
   action_SetName.Call(obj, GoStrToDStr(value))
}

func Action_GetTag(obj uintptr) int {
    ret, _, _ := action_GetTag.Call(obj)
    return int(ret)
}

func Action_SetTag(obj uintptr, value int) {
   action_SetTag.Call(obj, uintptr(value))
}

func Action_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := action_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

