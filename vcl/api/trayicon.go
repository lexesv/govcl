
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func TrayIcon_Create(obj uintptr) uintptr {
    ret, _, _ := trayIcon_Create.Call(obj)
    return ret
}

func TrayIcon_Free(obj uintptr) {
    trayIcon_Free.Call(obj)
}

func TrayIcon_Refresh(obj uintptr)  {
    trayIcon_Refresh.Call(obj)
}

func TrayIcon_SetDefaultIcon(obj uintptr)  {
    trayIcon_SetDefaultIcon.Call(obj)
}

func TrayIcon_ShowBalloonHint(obj uintptr)  {
    trayIcon_ShowBalloonHint.Call(obj)
}

func TrayIcon_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := trayIcon_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func TrayIcon_GetNamePath(obj uintptr) string {
    ret, _, _ := trayIcon_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func TrayIcon_HasParent(obj uintptr) bool {
    ret, _, _ := trayIcon_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func TrayIcon_Assign(obj uintptr, Source uintptr)  {
    trayIcon_Assign.Call(obj, Source )
}

func TrayIcon_ClassName(obj uintptr) string {
    ret, _, _ := trayIcon_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func TrayIcon_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := trayIcon_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func TrayIcon_GetHashCode(obj uintptr) int32 {
    ret, _, _ := trayIcon_GetHashCode.Call(obj)
    return int32(ret)
}

func TrayIcon_ToString(obj uintptr) string {
    ret, _, _ := trayIcon_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func TrayIcon_GetAnimate(obj uintptr) bool {
    ret, _, _ := trayIcon_GetAnimate.Call(obj)
    return DBoolToGoBool(ret)
}

func TrayIcon_SetAnimate(obj uintptr, value bool) {
   trayIcon_SetAnimate.Call(obj, GoBoolToDBool(value))
}

func TrayIcon_GetAnimateInterval(obj uintptr) uint32 {
    ret, _, _ := trayIcon_GetAnimateInterval.Call(obj)
    return uint32(ret)
}

func TrayIcon_SetAnimateInterval(obj uintptr, value uint32) {
   trayIcon_SetAnimateInterval.Call(obj, uintptr(value))
}

func TrayIcon_GetHint(obj uintptr) string {
    ret, _, _ := trayIcon_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func TrayIcon_SetHint(obj uintptr, value string) {
   trayIcon_SetHint.Call(obj, GoStrToDStr(value))
}

func TrayIcon_GetBalloonHint(obj uintptr) string {
    ret, _, _ := trayIcon_GetBalloonHint.Call(obj)
    return DStrToGoStr(ret)
}

func TrayIcon_SetBalloonHint(obj uintptr, value string) {
   trayIcon_SetBalloonHint.Call(obj, GoStrToDStr(value))
}

func TrayIcon_GetBalloonTitle(obj uintptr) string {
    ret, _, _ := trayIcon_GetBalloonTitle.Call(obj)
    return DStrToGoStr(ret)
}

func TrayIcon_SetBalloonTitle(obj uintptr, value string) {
   trayIcon_SetBalloonTitle.Call(obj, GoStrToDStr(value))
}

func TrayIcon_GetBalloonTimeout(obj uintptr) int32 {
    ret, _, _ := trayIcon_GetBalloonTimeout.Call(obj)
    return int32(ret)
}

func TrayIcon_SetBalloonTimeout(obj uintptr, value int32) {
   trayIcon_SetBalloonTimeout.Call(obj, uintptr(value))
}

func TrayIcon_GetBalloonFlags(obj uintptr) TBalloonFlags {
    ret, _, _ := trayIcon_GetBalloonFlags.Call(obj)
    return TBalloonFlags(ret)
}

func TrayIcon_SetBalloonFlags(obj uintptr, value TBalloonFlags) {
   trayIcon_SetBalloonFlags.Call(obj, uintptr(value))
}

func TrayIcon_GetIcon(obj uintptr) uintptr {
    ret, _, _ := trayIcon_GetIcon.Call(obj)
    return ret
}

func TrayIcon_SetIcon(obj uintptr, value uintptr) {
   trayIcon_SetIcon.Call(obj, value)
}

func TrayIcon_GetIconIndex(obj uintptr) int32 {
    ret, _, _ := trayIcon_GetIconIndex.Call(obj)
    return int32(ret)
}

func TrayIcon_SetIconIndex(obj uintptr, value int32) {
   trayIcon_SetIconIndex.Call(obj, uintptr(value))
}

func TrayIcon_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := trayIcon_GetPopupMenu.Call(obj)
    return ret
}

func TrayIcon_SetPopupMenu(obj uintptr, value uintptr) {
   trayIcon_SetPopupMenu.Call(obj, value)
}

func TrayIcon_GetVisible(obj uintptr) bool {
    ret, _, _ := trayIcon_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func TrayIcon_SetVisible(obj uintptr, value bool) {
   trayIcon_SetVisible.Call(obj, GoBoolToDBool(value))
}

func TrayIcon_SetOnBalloonClick(obj uintptr, fn interface{}) {
    trayIcon_SetOnBalloonClick.Call(obj, addEventToMap(fn))
}

func TrayIcon_SetOnClick(obj uintptr, fn interface{}) {
    trayIcon_SetOnClick.Call(obj, addEventToMap(fn))
}

func TrayIcon_SetOnDblClick(obj uintptr, fn interface{}) {
    trayIcon_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func TrayIcon_SetOnMouseMove(obj uintptr, fn interface{}) {
    trayIcon_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func TrayIcon_SetOnMouseUp(obj uintptr, fn interface{}) {
    trayIcon_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func TrayIcon_SetOnMouseDown(obj uintptr, fn interface{}) {
    trayIcon_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func TrayIcon_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := trayIcon_GetComponentCount.Call(obj)
    return int32(ret)
}

func TrayIcon_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := trayIcon_GetComponentIndex.Call(obj)
    return int32(ret)
}

func TrayIcon_SetComponentIndex(obj uintptr, value int32) {
   trayIcon_SetComponentIndex.Call(obj, uintptr(value))
}

func TrayIcon_GetOwner(obj uintptr) uintptr {
    ret, _, _ := trayIcon_GetOwner.Call(obj)
    return ret
}

func TrayIcon_GetName(obj uintptr) string {
    ret, _, _ := trayIcon_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func TrayIcon_SetName(obj uintptr, value string) {
   trayIcon_SetName.Call(obj, GoStrToDStr(value))
}

func TrayIcon_GetTag(obj uintptr) int {
    ret, _, _ := trayIcon_GetTag.Call(obj)
    return int(ret)
}

func TrayIcon_SetTag(obj uintptr, value int) {
   trayIcon_SetTag.Call(obj, uintptr(value))
}

func TrayIcon_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := trayIcon_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

