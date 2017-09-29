
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func CategoryPanelGroup_Create(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_Create.Call(obj)
    return ret
}

func CategoryPanelGroup_Free(obj uintptr) {
    categoryPanelGroup_Free.Call(obj)
}

func CategoryPanelGroup_CanFocus(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_FlipChildren(obj uintptr, AllLevels bool)  {
    categoryPanelGroup_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func CategoryPanelGroup_Focused(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_HandleAllocated(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_Invalidate(obj uintptr)  {
    categoryPanelGroup_Invalidate.Call(obj)
}

func CategoryPanelGroup_Realign(obj uintptr)  {
    categoryPanelGroup_Realign.Call(obj)
}

func CategoryPanelGroup_Repaint(obj uintptr)  {
    categoryPanelGroup_Repaint.Call(obj)
}

func CategoryPanelGroup_ScaleBy(obj uintptr, M int32, D int32)  {
    categoryPanelGroup_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func CategoryPanelGroup_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    categoryPanelGroup_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func CategoryPanelGroup_SetFocus(obj uintptr)  {
    categoryPanelGroup_SetFocus.Call(obj)
}

func CategoryPanelGroup_Update(obj uintptr)  {
    categoryPanelGroup_Update.Call(obj)
}

func CategoryPanelGroup_BringToFront(obj uintptr)  {
    categoryPanelGroup_BringToFront.Call(obj)
}

func CategoryPanelGroup_HasParent(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_Hide(obj uintptr)  {
    categoryPanelGroup_Hide.Call(obj)
}

func CategoryPanelGroup_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := categoryPanelGroup_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func CategoryPanelGroup_Refresh(obj uintptr)  {
    categoryPanelGroup_Refresh.Call(obj)
}

func CategoryPanelGroup_SendToBack(obj uintptr)  {
    categoryPanelGroup_SendToBack.Call(obj)
}

func CategoryPanelGroup_Show(obj uintptr)  {
    categoryPanelGroup_Show.Call(obj)
}

func CategoryPanelGroup_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := categoryPanelGroup_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func CategoryPanelGroup_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := categoryPanelGroup_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func CategoryPanelGroup_GetNamePath(obj uintptr) string {
    ret, _, _ := categoryPanelGroup_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func CategoryPanelGroup_Assign(obj uintptr, Source uintptr)  {
    categoryPanelGroup_Assign.Call(obj, Source )
}

func CategoryPanelGroup_ClassName(obj uintptr) string {
    ret, _, _ := categoryPanelGroup_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func CategoryPanelGroup_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_GetHashCode(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetHashCode.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_ToString(obj uintptr) string {
    ret, _, _ := categoryPanelGroup_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func CategoryPanelGroup_GetAlign(obj uintptr) TAlign {
    ret, _, _ := categoryPanelGroup_GetAlign.Call(obj)
    return TAlign(ret)
}

func CategoryPanelGroup_SetAlign(obj uintptr, value TAlign) {
   categoryPanelGroup_SetAlign.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := categoryPanelGroup_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func CategoryPanelGroup_SetAnchors(obj uintptr, value TAnchors) {
   categoryPanelGroup_SetAnchors.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetDoubleBuffered(obj uintptr, value bool) {
   categoryPanelGroup_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetEnabled(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetEnabled(obj uintptr, value bool) {
   categoryPanelGroup_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetColor(obj uintptr) TColor {
    ret, _, _ := categoryPanelGroup_GetColor.Call(obj)
    return TColor(ret)
}

func CategoryPanelGroup_SetColor(obj uintptr, value TColor) {
   categoryPanelGroup_SetColor.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetFont(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetFont.Call(obj)
    return ret
}

func CategoryPanelGroup_SetFont(obj uintptr, value uintptr) {
   categoryPanelGroup_SetFont.Call(obj, value)
}

func CategoryPanelGroup_GetGradientDirection(obj uintptr) TGradientDirection {
    ret, _, _ := categoryPanelGroup_GetGradientDirection.Call(obj)
    return TGradientDirection(ret)
}

func CategoryPanelGroup_SetGradientDirection(obj uintptr, value TGradientDirection) {
   categoryPanelGroup_SetGradientDirection.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetHeight(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetHeight.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_SetHeight(obj uintptr, value int32) {
   categoryPanelGroup_SetHeight.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetImages(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetImages.Call(obj)
    return ret
}

func CategoryPanelGroup_SetImages(obj uintptr, value uintptr) {
   categoryPanelGroup_SetImages.Call(obj, value)
}

func CategoryPanelGroup_GetParentColor(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetParentColor(obj uintptr, value bool) {
   categoryPanelGroup_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetParentCtl3D(obj uintptr, value bool) {
   categoryPanelGroup_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetParentFont(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetParentFont(obj uintptr, value bool) {
   categoryPanelGroup_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetParentShowHint(obj uintptr, value bool) {
   categoryPanelGroup_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetPopupMenu.Call(obj)
    return ret
}

func CategoryPanelGroup_SetPopupMenu(obj uintptr, value uintptr) {
   categoryPanelGroup_SetPopupMenu.Call(obj, value)
}

func CategoryPanelGroup_GetShowHint(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetShowHint(obj uintptr, value bool) {
   categoryPanelGroup_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := categoryPanelGroup_GetTabOrder.Call(obj)
    return int16(ret)
}

func CategoryPanelGroup_SetTabOrder(obj uintptr, value int16) {
   categoryPanelGroup_SetTabOrder.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetTabStop(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetTabStop(obj uintptr, value bool) {
   categoryPanelGroup_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetVisible(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetVisible(obj uintptr, value bool) {
   categoryPanelGroup_SetVisible.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetWidth(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetWidth.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_SetWidth(obj uintptr, value int32) {
   categoryPanelGroup_SetWidth.Call(obj, uintptr(value))
}

func CategoryPanelGroup_SetOnClick(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnClick.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnDblClick(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnEnter(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnEnter.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnExit(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnExit.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnMouseDown(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnMouseEnter(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnMouseLeave(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnMouseMove(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnMouseUp(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnMouseWheel(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnMouseWheel.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnResize(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnResize.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_GetPanels(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetPanels.Call(obj)
    return ret
}

func CategoryPanelGroup_GetBrush(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetBrush.Call(obj)
    return ret
}

func CategoryPanelGroup_GetControlCount(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetControlCount.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_GetHandle(obj uintptr) HWND {
    ret, _, _ := categoryPanelGroup_GetHandle.Call(obj)
    return HWND(ret)
}

func CategoryPanelGroup_GetAction(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetAction.Call(obj)
    return ret
}

func CategoryPanelGroup_SetAction(obj uintptr, value uintptr) {
   categoryPanelGroup_SetAction.Call(obj, value)
}

func CategoryPanelGroup_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    categoryPanelGroup_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func CategoryPanelGroup_SetBoundsRect(obj uintptr, value TRect) {
   categoryPanelGroup_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func CategoryPanelGroup_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetClientHeight.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_SetClientHeight(obj uintptr, value int32) {
   categoryPanelGroup_SetClientHeight.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetClientRect(obj uintptr) TRect {
    var ret TRect
    categoryPanelGroup_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func CategoryPanelGroup_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetClientWidth.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_SetClientWidth(obj uintptr, value int32) {
   categoryPanelGroup_SetClientWidth.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetExplicitTop.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_GetParent(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetParent.Call(obj)
    return ret
}

func CategoryPanelGroup_SetParent(obj uintptr, value uintptr) {
   categoryPanelGroup_SetParent.Call(obj, value)
}

func CategoryPanelGroup_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetAlignWithMargins(obj uintptr, value bool) {
   categoryPanelGroup_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetLeft(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetLeft.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_SetLeft(obj uintptr, value int32) {
   categoryPanelGroup_SetLeft.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetTop(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetTop.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_SetTop(obj uintptr, value int32) {
   categoryPanelGroup_SetTop.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetCursor(obj uintptr) TCursor {
    ret, _, _ := categoryPanelGroup_GetCursor.Call(obj)
    return TCursor(ret)
}

func CategoryPanelGroup_SetCursor(obj uintptr, value TCursor) {
   categoryPanelGroup_SetCursor.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetHint(obj uintptr) string {
    ret, _, _ := categoryPanelGroup_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func CategoryPanelGroup_SetHint(obj uintptr, value string) {
   categoryPanelGroup_SetHint.Call(obj, GoStrToDStr(value))
}

func CategoryPanelGroup_GetMargins(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetMargins.Call(obj)
    return ret
}

func CategoryPanelGroup_SetMargins(obj uintptr, value uintptr) {
   categoryPanelGroup_SetMargins.Call(obj, value)
}

func CategoryPanelGroup_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetComponentCount.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetComponentIndex.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_SetComponentIndex(obj uintptr, value int32) {
   categoryPanelGroup_SetComponentIndex.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetOwner(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetOwner.Call(obj)
    return ret
}

func CategoryPanelGroup_GetName(obj uintptr) string {
    ret, _, _ := categoryPanelGroup_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func CategoryPanelGroup_SetName(obj uintptr, value string) {
   categoryPanelGroup_SetName.Call(obj, GoStrToDStr(value))
}

func CategoryPanelGroup_GetTag(obj uintptr) int {
    ret, _, _ := categoryPanelGroup_GetTag.Call(obj)
    return int(ret)
}

func CategoryPanelGroup_SetTag(obj uintptr, value int) {
   categoryPanelGroup_SetTag.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := categoryPanelGroup_GetControls.Call(obj, uintptr(Index))
    return ret
}

func CategoryPanelGroup_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := categoryPanelGroup_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

