
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func TabSheet_Create(obj uintptr) uintptr {
    ret, _, _ := tabSheet_Create.Call(obj)
    return ret
}

func TabSheet_Free(obj uintptr) {
    tabSheet_Free.Call(obj)
}

func TabSheet_CanFocus(obj uintptr) bool {
    ret, _, _ := tabSheet_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_FlipChildren(obj uintptr, AllLevels bool)  {
    tabSheet_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func TabSheet_Focused(obj uintptr) bool {
    ret, _, _ := tabSheet_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_HandleAllocated(obj uintptr) bool {
    ret, _, _ := tabSheet_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_Invalidate(obj uintptr)  {
    tabSheet_Invalidate.Call(obj)
}

func TabSheet_Realign(obj uintptr)  {
    tabSheet_Realign.Call(obj)
}

func TabSheet_Repaint(obj uintptr)  {
    tabSheet_Repaint.Call(obj)
}

func TabSheet_ScaleBy(obj uintptr, M int32, D int32)  {
    tabSheet_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func TabSheet_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    tabSheet_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func TabSheet_SetFocus(obj uintptr)  {
    tabSheet_SetFocus.Call(obj)
}

func TabSheet_Update(obj uintptr)  {
    tabSheet_Update.Call(obj)
}

func TabSheet_BringToFront(obj uintptr)  {
    tabSheet_BringToFront.Call(obj)
}

func TabSheet_HasParent(obj uintptr) bool {
    ret, _, _ := tabSheet_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_Hide(obj uintptr)  {
    tabSheet_Hide.Call(obj)
}

func TabSheet_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := tabSheet_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func TabSheet_Refresh(obj uintptr)  {
    tabSheet_Refresh.Call(obj)
}

func TabSheet_SendToBack(obj uintptr)  {
    tabSheet_SendToBack.Call(obj)
}

func TabSheet_Show(obj uintptr)  {
    tabSheet_Show.Call(obj)
}

func TabSheet_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := tabSheet_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func TabSheet_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := tabSheet_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func TabSheet_GetNamePath(obj uintptr) string {
    ret, _, _ := tabSheet_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func TabSheet_Assign(obj uintptr, Source uintptr)  {
    tabSheet_Assign.Call(obj, Source )
}

func TabSheet_ClassName(obj uintptr) string {
    ret, _, _ := tabSheet_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func TabSheet_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := tabSheet_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func TabSheet_GetHashCode(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetHashCode.Call(obj)
    return int32(ret)
}

func TabSheet_ToString(obj uintptr) string {
    ret, _, _ := tabSheet_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func TabSheet_GetPageControl(obj uintptr) uintptr {
    ret, _, _ := tabSheet_GetPageControl.Call(obj)
    return ret
}

func TabSheet_SetPageControl(obj uintptr, value uintptr) {
   tabSheet_SetPageControl.Call(obj, value)
}

func TabSheet_GetTabIndex(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetTabIndex.Call(obj)
    return int32(ret)
}

func TabSheet_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetBorderWidth.Call(obj)
    return int32(ret)
}

func TabSheet_SetBorderWidth(obj uintptr, value int32) {
   tabSheet_SetBorderWidth.Call(obj, uintptr(value))
}

func TabSheet_GetCaption(obj uintptr) string {
    ret, _, _ := tabSheet_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func TabSheet_SetCaption(obj uintptr, value string) {
   tabSheet_SetCaption.Call(obj, GoStrToDStr(value))
}

func TabSheet_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := tabSheet_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetDoubleBuffered(obj uintptr, value bool) {
   tabSheet_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetEnabled(obj uintptr) bool {
    ret, _, _ := tabSheet_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetEnabled(obj uintptr, value bool) {
   tabSheet_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetFont(obj uintptr) uintptr {
    ret, _, _ := tabSheet_GetFont.Call(obj)
    return ret
}

func TabSheet_SetFont(obj uintptr, value uintptr) {
   tabSheet_SetFont.Call(obj, value)
}

func TabSheet_GetHeight(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetHeight.Call(obj)
    return int32(ret)
}

func TabSheet_SetHeight(obj uintptr, value int32) {
   tabSheet_SetHeight.Call(obj, uintptr(value))
}

func TabSheet_GetHighlighted(obj uintptr) bool {
    ret, _, _ := tabSheet_GetHighlighted.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetHighlighted(obj uintptr, value bool) {
   tabSheet_SetHighlighted.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetImageIndex(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetImageIndex.Call(obj)
    return int32(ret)
}

func TabSheet_SetImageIndex(obj uintptr, value int32) {
   tabSheet_SetImageIndex.Call(obj, uintptr(value))
}

func TabSheet_GetLeft(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetLeft.Call(obj)
    return int32(ret)
}

func TabSheet_SetLeft(obj uintptr, value int32) {
   tabSheet_SetLeft.Call(obj, uintptr(value))
}

func TabSheet_GetPageIndex(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetPageIndex.Call(obj)
    return int32(ret)
}

func TabSheet_SetPageIndex(obj uintptr, value int32) {
   tabSheet_SetPageIndex.Call(obj, uintptr(value))
}

func TabSheet_GetParentFont(obj uintptr) bool {
    ret, _, _ := tabSheet_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetParentFont(obj uintptr, value bool) {
   tabSheet_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := tabSheet_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetParentShowHint(obj uintptr, value bool) {
   tabSheet_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := tabSheet_GetPopupMenu.Call(obj)
    return ret
}

func TabSheet_SetPopupMenu(obj uintptr, value uintptr) {
   tabSheet_SetPopupMenu.Call(obj, value)
}

func TabSheet_GetShowHint(obj uintptr) bool {
    ret, _, _ := tabSheet_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetShowHint(obj uintptr, value bool) {
   tabSheet_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetTabVisible(obj uintptr) bool {
    ret, _, _ := tabSheet_GetTabVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetTabVisible(obj uintptr, value bool) {
   tabSheet_SetTabVisible.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetTop(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetTop.Call(obj)
    return int32(ret)
}

func TabSheet_SetTop(obj uintptr, value int32) {
   tabSheet_SetTop.Call(obj, uintptr(value))
}

func TabSheet_GetVisible(obj uintptr) bool {
    ret, _, _ := tabSheet_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetVisible(obj uintptr, value bool) {
   tabSheet_SetVisible.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetWidth(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetWidth.Call(obj)
    return int32(ret)
}

func TabSheet_SetWidth(obj uintptr, value int32) {
   tabSheet_SetWidth.Call(obj, uintptr(value))
}

func TabSheet_SetOnEnter(obj uintptr, fn interface{}) {
    tabSheet_SetOnEnter.Call(obj, addEventToMap(fn))
}

func TabSheet_SetOnExit(obj uintptr, fn interface{}) {
    tabSheet_SetOnExit.Call(obj, addEventToMap(fn))
}

func TabSheet_SetOnHide(obj uintptr, fn interface{}) {
    tabSheet_SetOnHide.Call(obj, addEventToMap(fn))
}

func TabSheet_SetOnMouseDown(obj uintptr, fn interface{}) {
    tabSheet_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func TabSheet_SetOnMouseEnter(obj uintptr, fn interface{}) {
    tabSheet_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func TabSheet_SetOnMouseLeave(obj uintptr, fn interface{}) {
    tabSheet_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func TabSheet_SetOnMouseMove(obj uintptr, fn interface{}) {
    tabSheet_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func TabSheet_SetOnMouseUp(obj uintptr, fn interface{}) {
    tabSheet_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func TabSheet_SetOnResize(obj uintptr, fn interface{}) {
    tabSheet_SetOnResize.Call(obj, addEventToMap(fn))
}

func TabSheet_SetOnShow(obj uintptr, fn interface{}) {
    tabSheet_SetOnShow.Call(obj, addEventToMap(fn))
}

func TabSheet_GetBrush(obj uintptr) uintptr {
    ret, _, _ := tabSheet_GetBrush.Call(obj)
    return ret
}

func TabSheet_GetControlCount(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetControlCount.Call(obj)
    return int32(ret)
}

func TabSheet_GetHandle(obj uintptr) HWND {
    ret, _, _ := tabSheet_GetHandle.Call(obj)
    return HWND(ret)
}

func TabSheet_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := tabSheet_GetTabOrder.Call(obj)
    return int16(ret)
}

func TabSheet_SetTabOrder(obj uintptr, value int16) {
   tabSheet_SetTabOrder.Call(obj, uintptr(value))
}

func TabSheet_GetTabStop(obj uintptr) bool {
    ret, _, _ := tabSheet_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetTabStop(obj uintptr, value bool) {
   tabSheet_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetAction(obj uintptr) uintptr {
    ret, _, _ := tabSheet_GetAction.Call(obj)
    return ret
}

func TabSheet_SetAction(obj uintptr, value uintptr) {
   tabSheet_SetAction.Call(obj, value)
}

func TabSheet_GetAlign(obj uintptr) TAlign {
    ret, _, _ := tabSheet_GetAlign.Call(obj)
    return TAlign(ret)
}

func TabSheet_SetAlign(obj uintptr, value TAlign) {
   tabSheet_SetAlign.Call(obj, uintptr(value))
}

func TabSheet_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := tabSheet_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func TabSheet_SetAnchors(obj uintptr, value TAnchors) {
   tabSheet_SetAnchors.Call(obj, uintptr(value))
}

func TabSheet_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    tabSheet_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func TabSheet_SetBoundsRect(obj uintptr, value TRect) {
   tabSheet_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func TabSheet_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetClientHeight.Call(obj)
    return int32(ret)
}

func TabSheet_SetClientHeight(obj uintptr, value int32) {
   tabSheet_SetClientHeight.Call(obj, uintptr(value))
}

func TabSheet_GetClientRect(obj uintptr) TRect {
    var ret TRect
    tabSheet_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func TabSheet_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetClientWidth.Call(obj)
    return int32(ret)
}

func TabSheet_SetClientWidth(obj uintptr, value int32) {
   tabSheet_SetClientWidth.Call(obj, uintptr(value))
}

func TabSheet_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func TabSheet_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetExplicitTop.Call(obj)
    return int32(ret)
}

func TabSheet_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func TabSheet_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func TabSheet_GetParent(obj uintptr) uintptr {
    ret, _, _ := tabSheet_GetParent.Call(obj)
    return ret
}

func TabSheet_SetParent(obj uintptr, value uintptr) {
   tabSheet_SetParent.Call(obj, value)
}

func TabSheet_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := tabSheet_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetAlignWithMargins(obj uintptr, value bool) {
   tabSheet_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetCursor(obj uintptr) TCursor {
    ret, _, _ := tabSheet_GetCursor.Call(obj)
    return TCursor(ret)
}

func TabSheet_SetCursor(obj uintptr, value TCursor) {
   tabSheet_SetCursor.Call(obj, uintptr(value))
}

func TabSheet_GetHint(obj uintptr) string {
    ret, _, _ := tabSheet_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func TabSheet_SetHint(obj uintptr, value string) {
   tabSheet_SetHint.Call(obj, GoStrToDStr(value))
}

func TabSheet_GetMargins(obj uintptr) uintptr {
    ret, _, _ := tabSheet_GetMargins.Call(obj)
    return ret
}

func TabSheet_SetMargins(obj uintptr, value uintptr) {
   tabSheet_SetMargins.Call(obj, value)
}

func TabSheet_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetComponentCount.Call(obj)
    return int32(ret)
}

func TabSheet_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetComponentIndex.Call(obj)
    return int32(ret)
}

func TabSheet_SetComponentIndex(obj uintptr, value int32) {
   tabSheet_SetComponentIndex.Call(obj, uintptr(value))
}

func TabSheet_GetOwner(obj uintptr) uintptr {
    ret, _, _ := tabSheet_GetOwner.Call(obj)
    return ret
}

func TabSheet_GetName(obj uintptr) string {
    ret, _, _ := tabSheet_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func TabSheet_SetName(obj uintptr, value string) {
   tabSheet_SetName.Call(obj, GoStrToDStr(value))
}

func TabSheet_GetTag(obj uintptr) int {
    ret, _, _ := tabSheet_GetTag.Call(obj)
    return int(ret)
}

func TabSheet_SetTag(obj uintptr, value int) {
   tabSheet_SetTag.Call(obj, uintptr(value))
}

func TabSheet_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := tabSheet_GetControls.Call(obj, uintptr(Index))
    return ret
}

func TabSheet_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := tabSheet_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

