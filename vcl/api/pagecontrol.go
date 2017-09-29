
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func PageControl_Create(obj uintptr) uintptr {
    ret, _, _ := pageControl_Create.Call(obj)
    return ret
}

func PageControl_Free(obj uintptr) {
    pageControl_Free.Call(obj)
}

func PageControl_SelectNextPage(obj uintptr, GoForward bool, CheckTabVisible bool)  {
    pageControl_SelectNextPage.Call(obj, GoBoolToDBool(GoForward) , GoBoolToDBool(CheckTabVisible) )
}

func PageControl_RowCount(obj uintptr) int32 {
    ret, _, _ := pageControl_RowCount.Call(obj)
    return int32(ret)
}

func PageControl_CanFocus(obj uintptr) bool {
    ret, _, _ := pageControl_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_FlipChildren(obj uintptr, AllLevels bool)  {
    pageControl_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func PageControl_Focused(obj uintptr) bool {
    ret, _, _ := pageControl_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_HandleAllocated(obj uintptr) bool {
    ret, _, _ := pageControl_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_Invalidate(obj uintptr)  {
    pageControl_Invalidate.Call(obj)
}

func PageControl_Realign(obj uintptr)  {
    pageControl_Realign.Call(obj)
}

func PageControl_Repaint(obj uintptr)  {
    pageControl_Repaint.Call(obj)
}

func PageControl_ScaleBy(obj uintptr, M int32, D int32)  {
    pageControl_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func PageControl_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    pageControl_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func PageControl_SetFocus(obj uintptr)  {
    pageControl_SetFocus.Call(obj)
}

func PageControl_Update(obj uintptr)  {
    pageControl_Update.Call(obj)
}

func PageControl_BringToFront(obj uintptr)  {
    pageControl_BringToFront.Call(obj)
}

func PageControl_HasParent(obj uintptr) bool {
    ret, _, _ := pageControl_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_Hide(obj uintptr)  {
    pageControl_Hide.Call(obj)
}

func PageControl_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := pageControl_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func PageControl_Refresh(obj uintptr)  {
    pageControl_Refresh.Call(obj)
}

func PageControl_SendToBack(obj uintptr)  {
    pageControl_SendToBack.Call(obj)
}

func PageControl_Show(obj uintptr)  {
    pageControl_Show.Call(obj)
}

func PageControl_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := pageControl_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func PageControl_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := pageControl_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func PageControl_GetNamePath(obj uintptr) string {
    ret, _, _ := pageControl_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func PageControl_Assign(obj uintptr, Source uintptr)  {
    pageControl_Assign.Call(obj, Source )
}

func PageControl_ClassName(obj uintptr) string {
    ret, _, _ := pageControl_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func PageControl_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := pageControl_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func PageControl_GetHashCode(obj uintptr) int32 {
    ret, _, _ := pageControl_GetHashCode.Call(obj)
    return int32(ret)
}

func PageControl_ToString(obj uintptr) string {
    ret, _, _ := pageControl_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func PageControl_GetActivePageIndex(obj uintptr) int32 {
    ret, _, _ := pageControl_GetActivePageIndex.Call(obj)
    return int32(ret)
}

func PageControl_SetActivePageIndex(obj uintptr, value int32) {
   pageControl_SetActivePageIndex.Call(obj, uintptr(value))
}

func PageControl_GetPageCount(obj uintptr) int32 {
    ret, _, _ := pageControl_GetPageCount.Call(obj)
    return int32(ret)
}

func PageControl_GetAlign(obj uintptr) TAlign {
    ret, _, _ := pageControl_GetAlign.Call(obj)
    return TAlign(ret)
}

func PageControl_SetAlign(obj uintptr, value TAlign) {
   pageControl_SetAlign.Call(obj, uintptr(value))
}

func PageControl_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := pageControl_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func PageControl_SetAnchors(obj uintptr, value TAnchors) {
   pageControl_SetAnchors.Call(obj, uintptr(value))
}

func PageControl_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := pageControl_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetDoubleBuffered(obj uintptr, value bool) {
   pageControl_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetEnabled(obj uintptr) bool {
    ret, _, _ := pageControl_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetEnabled(obj uintptr, value bool) {
   pageControl_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetFont(obj uintptr) uintptr {
    ret, _, _ := pageControl_GetFont.Call(obj)
    return ret
}

func PageControl_SetFont(obj uintptr, value uintptr) {
   pageControl_SetFont.Call(obj, value)
}

func PageControl_GetHotTrack(obj uintptr) bool {
    ret, _, _ := pageControl_GetHotTrack.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetHotTrack(obj uintptr, value bool) {
   pageControl_SetHotTrack.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetImages(obj uintptr) uintptr {
    ret, _, _ := pageControl_GetImages.Call(obj)
    return ret
}

func PageControl_SetImages(obj uintptr, value uintptr) {
   pageControl_SetImages.Call(obj, value)
}

func PageControl_GetMultiLine(obj uintptr) bool {
    ret, _, _ := pageControl_GetMultiLine.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetMultiLine(obj uintptr, value bool) {
   pageControl_SetMultiLine.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetParentFont(obj uintptr) bool {
    ret, _, _ := pageControl_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetParentFont(obj uintptr, value bool) {
   pageControl_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := pageControl_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetParentShowHint(obj uintptr, value bool) {
   pageControl_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := pageControl_GetPopupMenu.Call(obj)
    return ret
}

func PageControl_SetPopupMenu(obj uintptr, value uintptr) {
   pageControl_SetPopupMenu.Call(obj, value)
}

func PageControl_GetShowHint(obj uintptr) bool {
    ret, _, _ := pageControl_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetShowHint(obj uintptr, value bool) {
   pageControl_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetStyle(obj uintptr) TTabStyle {
    ret, _, _ := pageControl_GetStyle.Call(obj)
    return TTabStyle(ret)
}

func PageControl_SetStyle(obj uintptr, value TTabStyle) {
   pageControl_SetStyle.Call(obj, uintptr(value))
}

func PageControl_GetTabHeight(obj uintptr) int16 {
    ret, _, _ := pageControl_GetTabHeight.Call(obj)
    return int16(ret)
}

func PageControl_SetTabHeight(obj uintptr, value int16) {
   pageControl_SetTabHeight.Call(obj, uintptr(value))
}

func PageControl_GetTabIndex(obj uintptr) int32 {
    ret, _, _ := pageControl_GetTabIndex.Call(obj)
    return int32(ret)
}

func PageControl_SetTabIndex(obj uintptr, value int32) {
   pageControl_SetTabIndex.Call(obj, uintptr(value))
}

func PageControl_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := pageControl_GetTabOrder.Call(obj)
    return int16(ret)
}

func PageControl_SetTabOrder(obj uintptr, value int16) {
   pageControl_SetTabOrder.Call(obj, uintptr(value))
}

func PageControl_GetTabPosition(obj uintptr) TTabPosition {
    ret, _, _ := pageControl_GetTabPosition.Call(obj)
    return TTabPosition(ret)
}

func PageControl_SetTabPosition(obj uintptr, value TTabPosition) {
   pageControl_SetTabPosition.Call(obj, uintptr(value))
}

func PageControl_GetTabStop(obj uintptr) bool {
    ret, _, _ := pageControl_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetTabStop(obj uintptr, value bool) {
   pageControl_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetTabWidth(obj uintptr) int16 {
    ret, _, _ := pageControl_GetTabWidth.Call(obj)
    return int16(ret)
}

func PageControl_SetTabWidth(obj uintptr, value int16) {
   pageControl_SetTabWidth.Call(obj, uintptr(value))
}

func PageControl_GetVisible(obj uintptr) bool {
    ret, _, _ := pageControl_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetVisible(obj uintptr, value bool) {
   pageControl_SetVisible.Call(obj, GoBoolToDBool(value))
}

func PageControl_SetOnChange(obj uintptr, fn interface{}) {
    pageControl_SetOnChange.Call(obj, addEventToMap(fn))
}

func PageControl_SetOnEnter(obj uintptr, fn interface{}) {
    pageControl_SetOnEnter.Call(obj, addEventToMap(fn))
}

func PageControl_SetOnExit(obj uintptr, fn interface{}) {
    pageControl_SetOnExit.Call(obj, addEventToMap(fn))
}

func PageControl_SetOnMouseDown(obj uintptr, fn interface{}) {
    pageControl_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func PageControl_SetOnMouseEnter(obj uintptr, fn interface{}) {
    pageControl_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func PageControl_SetOnMouseLeave(obj uintptr, fn interface{}) {
    pageControl_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func PageControl_SetOnMouseMove(obj uintptr, fn interface{}) {
    pageControl_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func PageControl_SetOnMouseUp(obj uintptr, fn interface{}) {
    pageControl_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func PageControl_SetOnResize(obj uintptr, fn interface{}) {
    pageControl_SetOnResize.Call(obj, addEventToMap(fn))
}

func PageControl_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := pageControl_GetCanvas.Call(obj)
    return ret
}

func PageControl_GetBrush(obj uintptr) uintptr {
    ret, _, _ := pageControl_GetBrush.Call(obj)
    return ret
}

func PageControl_GetControlCount(obj uintptr) int32 {
    ret, _, _ := pageControl_GetControlCount.Call(obj)
    return int32(ret)
}

func PageControl_GetHandle(obj uintptr) HWND {
    ret, _, _ := pageControl_GetHandle.Call(obj)
    return HWND(ret)
}

func PageControl_GetAction(obj uintptr) uintptr {
    ret, _, _ := pageControl_GetAction.Call(obj)
    return ret
}

func PageControl_SetAction(obj uintptr, value uintptr) {
   pageControl_SetAction.Call(obj, value)
}

func PageControl_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    pageControl_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func PageControl_SetBoundsRect(obj uintptr, value TRect) {
   pageControl_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func PageControl_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := pageControl_GetClientHeight.Call(obj)
    return int32(ret)
}

func PageControl_SetClientHeight(obj uintptr, value int32) {
   pageControl_SetClientHeight.Call(obj, uintptr(value))
}

func PageControl_GetClientRect(obj uintptr) TRect {
    var ret TRect
    pageControl_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func PageControl_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := pageControl_GetClientWidth.Call(obj)
    return int32(ret)
}

func PageControl_SetClientWidth(obj uintptr, value int32) {
   pageControl_SetClientWidth.Call(obj, uintptr(value))
}

func PageControl_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := pageControl_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func PageControl_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := pageControl_GetExplicitTop.Call(obj)
    return int32(ret)
}

func PageControl_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := pageControl_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func PageControl_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := pageControl_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func PageControl_GetParent(obj uintptr) uintptr {
    ret, _, _ := pageControl_GetParent.Call(obj)
    return ret
}

func PageControl_SetParent(obj uintptr, value uintptr) {
   pageControl_SetParent.Call(obj, value)
}

func PageControl_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := pageControl_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetAlignWithMargins(obj uintptr, value bool) {
   pageControl_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetLeft(obj uintptr) int32 {
    ret, _, _ := pageControl_GetLeft.Call(obj)
    return int32(ret)
}

func PageControl_SetLeft(obj uintptr, value int32) {
   pageControl_SetLeft.Call(obj, uintptr(value))
}

func PageControl_GetTop(obj uintptr) int32 {
    ret, _, _ := pageControl_GetTop.Call(obj)
    return int32(ret)
}

func PageControl_SetTop(obj uintptr, value int32) {
   pageControl_SetTop.Call(obj, uintptr(value))
}

func PageControl_GetWidth(obj uintptr) int32 {
    ret, _, _ := pageControl_GetWidth.Call(obj)
    return int32(ret)
}

func PageControl_SetWidth(obj uintptr, value int32) {
   pageControl_SetWidth.Call(obj, uintptr(value))
}

func PageControl_GetHeight(obj uintptr) int32 {
    ret, _, _ := pageControl_GetHeight.Call(obj)
    return int32(ret)
}

func PageControl_SetHeight(obj uintptr, value int32) {
   pageControl_SetHeight.Call(obj, uintptr(value))
}

func PageControl_GetCursor(obj uintptr) TCursor {
    ret, _, _ := pageControl_GetCursor.Call(obj)
    return TCursor(ret)
}

func PageControl_SetCursor(obj uintptr, value TCursor) {
   pageControl_SetCursor.Call(obj, uintptr(value))
}

func PageControl_GetHint(obj uintptr) string {
    ret, _, _ := pageControl_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func PageControl_SetHint(obj uintptr, value string) {
   pageControl_SetHint.Call(obj, GoStrToDStr(value))
}

func PageControl_GetMargins(obj uintptr) uintptr {
    ret, _, _ := pageControl_GetMargins.Call(obj)
    return ret
}

func PageControl_SetMargins(obj uintptr, value uintptr) {
   pageControl_SetMargins.Call(obj, value)
}

func PageControl_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := pageControl_GetComponentCount.Call(obj)
    return int32(ret)
}

func PageControl_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := pageControl_GetComponentIndex.Call(obj)
    return int32(ret)
}

func PageControl_SetComponentIndex(obj uintptr, value int32) {
   pageControl_SetComponentIndex.Call(obj, uintptr(value))
}

func PageControl_GetOwner(obj uintptr) uintptr {
    ret, _, _ := pageControl_GetOwner.Call(obj)
    return ret
}

func PageControl_GetName(obj uintptr) string {
    ret, _, _ := pageControl_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func PageControl_SetName(obj uintptr, value string) {
   pageControl_SetName.Call(obj, GoStrToDStr(value))
}

func PageControl_GetTag(obj uintptr) int {
    ret, _, _ := pageControl_GetTag.Call(obj)
    return int(ret)
}

func PageControl_SetTag(obj uintptr, value int) {
   pageControl_SetTag.Call(obj, uintptr(value))
}

func PageControl_GetPages(obj uintptr, Index int32) uintptr {
    ret, _, _ := pageControl_GetPages.Call(obj, uintptr(Index))
    return ret
}

func PageControl_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := pageControl_GetControls.Call(obj, uintptr(Index))
    return ret
}

func PageControl_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := pageControl_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

