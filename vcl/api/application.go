
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func Application_Create(obj uintptr) uintptr {
    ret, _, _ := application_Create.Call(obj)
    return ret
}

func Application_Free(obj uintptr) {
    application_Free.Call(obj)
}

func Application_ActivateHint(obj uintptr, CursorPos TPoint)  {
    application_ActivateHint.Call(obj, uintptr(unsafe.Pointer(&CursorPos)))
}

func Application_BringToFront(obj uintptr)  {
    application_BringToFront.Call(obj)
}

func Application_CancelHint(obj uintptr)  {
    application_CancelHint.Call(obj)
}

func Application_HandleMessage(obj uintptr)  {
    application_HandleMessage.Call(obj)
}

func Application_HideHint(obj uintptr)  {
    application_HideHint.Call(obj)
}

func Application_Initialize(obj uintptr)  {
    application_Initialize.Call(obj)
}

func Application_Minimize(obj uintptr)  {
    application_Minimize.Call(obj)
}

func Application_ModalStarted(obj uintptr)  {
    application_ModalStarted.Call(obj)
}

func Application_ModalFinished(obj uintptr)  {
    application_ModalFinished.Call(obj)
}

func Application_NormalizeAllTopMosts(obj uintptr)  {
    application_NormalizeAllTopMosts.Call(obj)
}

func Application_NormalizeTopMosts(obj uintptr)  {
    application_NormalizeTopMosts.Call(obj)
}

func Application_ProcessMessages(obj uintptr)  {
    application_ProcessMessages.Call(obj)
}

func Application_Restore(obj uintptr)  {
    application_Restore.Call(obj)
}

func Application_RestoreTopMosts(obj uintptr)  {
    application_RestoreTopMosts.Call(obj)
}

func Application_Run(obj uintptr)  {
    application_Run.Call(obj)
}

func Application_Terminate(obj uintptr)  {
    application_Terminate.Call(obj)
}

func Application_MessageBox(obj uintptr, Text string, Caption string, Flags int32) int32 {
    ret, _, _ := application_MessageBox.Call(obj, GoStrToDStr(Text) , GoStrToDStr(Caption) , uintptr(Flags) )
    return int32(ret)
}

func Application_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := application_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Application_GetNamePath(obj uintptr) string {
    ret, _, _ := application_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Application_HasParent(obj uintptr) bool {
    ret, _, _ := application_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Application_Assign(obj uintptr, Source uintptr)  {
    application_Assign.Call(obj, Source )
}

func Application_ClassName(obj uintptr) string {
    ret, _, _ := application_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Application_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := application_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Application_GetHashCode(obj uintptr) int32 {
    ret, _, _ := application_GetHashCode.Call(obj)
    return int32(ret)
}

func Application_ToString(obj uintptr) string {
    ret, _, _ := application_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Application_GetDialogHandle(obj uintptr) HWND {
    ret, _, _ := application_GetDialogHandle.Call(obj)
    return HWND(ret)
}

func Application_SetDialogHandle(obj uintptr, value HWND) {
   application_SetDialogHandle.Call(obj, uintptr(value))
}

func Application_GetHint(obj uintptr) string {
    ret, _, _ := application_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Application_SetHint(obj uintptr, value string) {
   application_SetHint.Call(obj, GoStrToDStr(value))
}

func Application_GetHintColor(obj uintptr) TColor {
    ret, _, _ := application_GetHintColor.Call(obj)
    return TColor(ret)
}

func Application_SetHintColor(obj uintptr, value TColor) {
   application_SetHintColor.Call(obj, uintptr(value))
}

func Application_GetHintHidePause(obj uintptr) int32 {
    ret, _, _ := application_GetHintHidePause.Call(obj)
    return int32(ret)
}

func Application_SetHintHidePause(obj uintptr, value int32) {
   application_SetHintHidePause.Call(obj, uintptr(value))
}

func Application_GetHintPause(obj uintptr) int32 {
    ret, _, _ := application_GetHintPause.Call(obj)
    return int32(ret)
}

func Application_SetHintPause(obj uintptr, value int32) {
   application_SetHintPause.Call(obj, uintptr(value))
}

func Application_GetHintShortCuts(obj uintptr) bool {
    ret, _, _ := application_GetHintShortCuts.Call(obj)
    return DBoolToGoBool(ret)
}

func Application_SetHintShortCuts(obj uintptr, value bool) {
   application_SetHintShortCuts.Call(obj, GoBoolToDBool(value))
}

func Application_GetHintShortPause(obj uintptr) int32 {
    ret, _, _ := application_GetHintShortPause.Call(obj)
    return int32(ret)
}

func Application_SetHintShortPause(obj uintptr, value int32) {
   application_SetHintShortPause.Call(obj, uintptr(value))
}

func Application_GetIcon(obj uintptr) uintptr {
    ret, _, _ := application_GetIcon.Call(obj)
    return ret
}

func Application_SetIcon(obj uintptr, value uintptr) {
   application_SetIcon.Call(obj, value)
}

func Application_GetIsMetropolisUI(obj uintptr) bool {
    ret, _, _ := application_GetIsMetropolisUI.Call(obj)
    return DBoolToGoBool(ret)
}

func Application_GetMainForm(obj uintptr) uintptr {
    ret, _, _ := application_GetMainForm.Call(obj)
    return ret
}

func Application_GetMainFormHandle(obj uintptr) HWND {
    ret, _, _ := application_GetMainFormHandle.Call(obj)
    return HWND(ret)
}

func Application_GetMainFormOnTaskBar(obj uintptr) bool {
    ret, _, _ := application_GetMainFormOnTaskBar.Call(obj)
    return DBoolToGoBool(ret)
}

func Application_SetMainFormOnTaskBar(obj uintptr, value bool) {
   application_SetMainFormOnTaskBar.Call(obj, GoBoolToDBool(value))
}

func Application_GetShowHint(obj uintptr) bool {
    ret, _, _ := application_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Application_SetShowHint(obj uintptr, value bool) {
   application_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Application_GetShowMainForm(obj uintptr) bool {
    ret, _, _ := application_GetShowMainForm.Call(obj)
    return DBoolToGoBool(ret)
}

func Application_SetShowMainForm(obj uintptr, value bool) {
   application_SetShowMainForm.Call(obj, GoBoolToDBool(value))
}

func Application_GetTitle(obj uintptr) string {
    ret, _, _ := application_GetTitle.Call(obj)
    return DStrToGoStr(ret)
}

func Application_SetTitle(obj uintptr, value string) {
   application_SetTitle.Call(obj, GoStrToDStr(value))
}

func Application_SetOnException(obj uintptr, fn interface{}) {
    application_SetOnException.Call(obj, addEventToMap(fn))
}

func Application_SetOnMinimize(obj uintptr, fn interface{}) {
    application_SetOnMinimize.Call(obj, addEventToMap(fn))
}

func Application_SetOnRestore(obj uintptr, fn interface{}) {
    application_SetOnRestore.Call(obj, addEventToMap(fn))
}

func Application_GetHandle(obj uintptr) HWND {
    ret, _, _ := application_GetHandle.Call(obj)
    return HWND(ret)
}

func Application_SetHandle(obj uintptr, value HWND) {
   application_SetHandle.Call(obj, uintptr(value))
}

func Application_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := application_GetComponentCount.Call(obj)
    return int32(ret)
}

func Application_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := application_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Application_SetComponentIndex(obj uintptr, value int32) {
   application_SetComponentIndex.Call(obj, uintptr(value))
}

func Application_GetOwner(obj uintptr) uintptr {
    ret, _, _ := application_GetOwner.Call(obj)
    return ret
}

func Application_GetName(obj uintptr) string {
    ret, _, _ := application_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Application_SetName(obj uintptr, value string) {
   application_SetName.Call(obj, GoStrToDStr(value))
}

func Application_GetTag(obj uintptr) int {
    ret, _, _ := application_GetTag.Call(obj)
    return int(ret)
}

func Application_SetTag(obj uintptr, value int) {
   application_SetTag.Call(obj, uintptr(value))
}

func Application_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := application_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

