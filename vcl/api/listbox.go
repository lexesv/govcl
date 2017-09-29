
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func ListBox_Create(obj uintptr) uintptr {
    ret, _, _ := listBox_Create.Call(obj)
    return ret
}

func ListBox_Free(obj uintptr) {
    listBox_Free.Call(obj)
}

func ListBox_AddItem(obj uintptr, Item string, AObject uintptr)  {
    listBox_AddItem.Call(obj, GoStrToDStr(Item) , AObject )
}

func ListBox_Clear(obj uintptr)  {
    listBox_Clear.Call(obj)
}

func ListBox_ClearSelection(obj uintptr)  {
    listBox_ClearSelection.Call(obj)
}

func ListBox_DeleteSelected(obj uintptr)  {
    listBox_DeleteSelected.Call(obj)
}

func ListBox_SelectAll(obj uintptr)  {
    listBox_SelectAll.Call(obj)
}

func ListBox_CanFocus(obj uintptr) bool {
    ret, _, _ := listBox_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_FlipChildren(obj uintptr, AllLevels bool)  {
    listBox_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func ListBox_Focused(obj uintptr) bool {
    ret, _, _ := listBox_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_HandleAllocated(obj uintptr) bool {
    ret, _, _ := listBox_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_Invalidate(obj uintptr)  {
    listBox_Invalidate.Call(obj)
}

func ListBox_Realign(obj uintptr)  {
    listBox_Realign.Call(obj)
}

func ListBox_Repaint(obj uintptr)  {
    listBox_Repaint.Call(obj)
}

func ListBox_ScaleBy(obj uintptr, M int32, D int32)  {
    listBox_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func ListBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    listBox_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func ListBox_SetFocus(obj uintptr)  {
    listBox_SetFocus.Call(obj)
}

func ListBox_Update(obj uintptr)  {
    listBox_Update.Call(obj)
}

func ListBox_BringToFront(obj uintptr)  {
    listBox_BringToFront.Call(obj)
}

func ListBox_HasParent(obj uintptr) bool {
    ret, _, _ := listBox_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_Hide(obj uintptr)  {
    listBox_Hide.Call(obj)
}

func ListBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := listBox_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func ListBox_Refresh(obj uintptr)  {
    listBox_Refresh.Call(obj)
}

func ListBox_SendToBack(obj uintptr)  {
    listBox_SendToBack.Call(obj)
}

func ListBox_Show(obj uintptr)  {
    listBox_Show.Call(obj)
}

func ListBox_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := listBox_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func ListBox_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := listBox_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ListBox_GetNamePath(obj uintptr) string {
    ret, _, _ := listBox_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ListBox_Assign(obj uintptr, Source uintptr)  {
    listBox_Assign.Call(obj, Source )
}

func ListBox_ClassName(obj uintptr) string {
    ret, _, _ := listBox_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ListBox_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := listBox_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ListBox_GetHashCode(obj uintptr) int32 {
    ret, _, _ := listBox_GetHashCode.Call(obj)
    return int32(ret)
}

func ListBox_ToString(obj uintptr) string {
    ret, _, _ := listBox_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ListBox_GetStyle(obj uintptr) TListBoxStyle {
    ret, _, _ := listBox_GetStyle.Call(obj)
    return TListBoxStyle(ret)
}

func ListBox_SetStyle(obj uintptr, value TListBoxStyle) {
   listBox_SetStyle.Call(obj, uintptr(value))
}

func ListBox_GetAutoComplete(obj uintptr) bool {
    ret, _, _ := listBox_GetAutoComplete.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetAutoComplete(obj uintptr, value bool) {
   listBox_SetAutoComplete.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetAutoCompleteDelay(obj uintptr) uint32 {
    ret, _, _ := listBox_GetAutoCompleteDelay.Call(obj)
    return uint32(ret)
}

func ListBox_SetAutoCompleteDelay(obj uintptr, value uint32) {
   listBox_SetAutoCompleteDelay.Call(obj, uintptr(value))
}

func ListBox_GetAlign(obj uintptr) TAlign {
    ret, _, _ := listBox_GetAlign.Call(obj)
    return TAlign(ret)
}

func ListBox_SetAlign(obj uintptr, value TAlign) {
   listBox_SetAlign.Call(obj, uintptr(value))
}

func ListBox_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := listBox_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func ListBox_SetAnchors(obj uintptr, value TAnchors) {
   listBox_SetAnchors.Call(obj, uintptr(value))
}

func ListBox_GetBorderStyle(obj uintptr) TBorderStyle {
    ret, _, _ := listBox_GetBorderStyle.Call(obj)
    return TBorderStyle(ret)
}

func ListBox_SetBorderStyle(obj uintptr, value TBorderStyle) {
   listBox_SetBorderStyle.Call(obj, uintptr(value))
}

func ListBox_GetColor(obj uintptr) TColor {
    ret, _, _ := listBox_GetColor.Call(obj)
    return TColor(ret)
}

func ListBox_SetColor(obj uintptr, value TColor) {
   listBox_SetColor.Call(obj, uintptr(value))
}

func ListBox_GetColumns(obj uintptr) int32 {
    ret, _, _ := listBox_GetColumns.Call(obj)
    return int32(ret)
}

func ListBox_SetColumns(obj uintptr, value int32) {
   listBox_SetColumns.Call(obj, uintptr(value))
}

func ListBox_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := listBox_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetDoubleBuffered(obj uintptr, value bool) {
   listBox_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetEnabled(obj uintptr) bool {
    ret, _, _ := listBox_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetEnabled(obj uintptr, value bool) {
   listBox_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetFont(obj uintptr) uintptr {
    ret, _, _ := listBox_GetFont.Call(obj)
    return ret
}

func ListBox_SetFont(obj uintptr, value uintptr) {
   listBox_SetFont.Call(obj, value)
}

func ListBox_GetItemHeight(obj uintptr) int32 {
    ret, _, _ := listBox_GetItemHeight.Call(obj)
    return int32(ret)
}

func ListBox_SetItemHeight(obj uintptr, value int32) {
   listBox_SetItemHeight.Call(obj, uintptr(value))
}

func ListBox_GetItems(obj uintptr) uintptr {
    ret, _, _ := listBox_GetItems.Call(obj)
    return ret
}

func ListBox_SetItems(obj uintptr, value uintptr) {
   listBox_SetItems.Call(obj, value)
}

func ListBox_GetMultiSelect(obj uintptr) bool {
    ret, _, _ := listBox_GetMultiSelect.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetMultiSelect(obj uintptr, value bool) {
   listBox_SetMultiSelect.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetParentColor(obj uintptr) bool {
    ret, _, _ := listBox_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetParentColor(obj uintptr, value bool) {
   listBox_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := listBox_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetParentCtl3D(obj uintptr, value bool) {
   listBox_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetParentFont(obj uintptr) bool {
    ret, _, _ := listBox_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetParentFont(obj uintptr, value bool) {
   listBox_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := listBox_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetParentShowHint(obj uintptr, value bool) {
   listBox_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := listBox_GetPopupMenu.Call(obj)
    return ret
}

func ListBox_SetPopupMenu(obj uintptr, value uintptr) {
   listBox_SetPopupMenu.Call(obj, value)
}

func ListBox_GetShowHint(obj uintptr) bool {
    ret, _, _ := listBox_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetShowHint(obj uintptr, value bool) {
   listBox_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetSorted(obj uintptr) bool {
    ret, _, _ := listBox_GetSorted.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetSorted(obj uintptr, value bool) {
   listBox_SetSorted.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := listBox_GetTabOrder.Call(obj)
    return int16(ret)
}

func ListBox_SetTabOrder(obj uintptr, value int16) {
   listBox_SetTabOrder.Call(obj, uintptr(value))
}

func ListBox_GetTabStop(obj uintptr) bool {
    ret, _, _ := listBox_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetTabStop(obj uintptr, value bool) {
   listBox_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetTabWidth(obj uintptr) int32 {
    ret, _, _ := listBox_GetTabWidth.Call(obj)
    return int32(ret)
}

func ListBox_SetTabWidth(obj uintptr, value int32) {
   listBox_SetTabWidth.Call(obj, uintptr(value))
}

func ListBox_GetVisible(obj uintptr) bool {
    ret, _, _ := listBox_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetVisible(obj uintptr, value bool) {
   listBox_SetVisible.Call(obj, GoBoolToDBool(value))
}

func ListBox_SetOnClick(obj uintptr, fn interface{}) {
    listBox_SetOnClick.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnDblClick(obj uintptr, fn interface{}) {
    listBox_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnEnter(obj uintptr, fn interface{}) {
    listBox_SetOnEnter.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnExit(obj uintptr, fn interface{}) {
    listBox_SetOnExit.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnKeyDown(obj uintptr, fn interface{}) {
    listBox_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnKeyPress(obj uintptr, fn interface{}) {
    listBox_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnKeyUp(obj uintptr, fn interface{}) {
    listBox_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnMouseDown(obj uintptr, fn interface{}) {
    listBox_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
    listBox_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
    listBox_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnMouseMove(obj uintptr, fn interface{}) {
    listBox_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnMouseUp(obj uintptr, fn interface{}) {
    listBox_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func ListBox_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := listBox_GetCanvas.Call(obj)
    return ret
}

func ListBox_GetSelCount(obj uintptr) int32 {
    ret, _, _ := listBox_GetSelCount.Call(obj)
    return int32(ret)
}

func ListBox_GetItemIndex(obj uintptr) int32 {
    ret, _, _ := listBox_GetItemIndex.Call(obj)
    return int32(ret)
}

func ListBox_SetItemIndex(obj uintptr, value int32) {
   listBox_SetItemIndex.Call(obj, uintptr(value))
}

func ListBox_GetBrush(obj uintptr) uintptr {
    ret, _, _ := listBox_GetBrush.Call(obj)
    return ret
}

func ListBox_GetControlCount(obj uintptr) int32 {
    ret, _, _ := listBox_GetControlCount.Call(obj)
    return int32(ret)
}

func ListBox_GetHandle(obj uintptr) HWND {
    ret, _, _ := listBox_GetHandle.Call(obj)
    return HWND(ret)
}

func ListBox_GetAction(obj uintptr) uintptr {
    ret, _, _ := listBox_GetAction.Call(obj)
    return ret
}

func ListBox_SetAction(obj uintptr, value uintptr) {
   listBox_SetAction.Call(obj, value)
}

func ListBox_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    listBox_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ListBox_SetBoundsRect(obj uintptr, value TRect) {
   listBox_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ListBox_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := listBox_GetClientHeight.Call(obj)
    return int32(ret)
}

func ListBox_SetClientHeight(obj uintptr, value int32) {
   listBox_SetClientHeight.Call(obj, uintptr(value))
}

func ListBox_GetClientRect(obj uintptr) TRect {
    var ret TRect
    listBox_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ListBox_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := listBox_GetClientWidth.Call(obj)
    return int32(ret)
}

func ListBox_SetClientWidth(obj uintptr, value int32) {
   listBox_SetClientWidth.Call(obj, uintptr(value))
}

func ListBox_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := listBox_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func ListBox_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := listBox_GetExplicitTop.Call(obj)
    return int32(ret)
}

func ListBox_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := listBox_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func ListBox_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := listBox_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func ListBox_GetParent(obj uintptr) uintptr {
    ret, _, _ := listBox_GetParent.Call(obj)
    return ret
}

func ListBox_SetParent(obj uintptr, value uintptr) {
   listBox_SetParent.Call(obj, value)
}

func ListBox_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := listBox_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetAlignWithMargins(obj uintptr, value bool) {
   listBox_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetLeft(obj uintptr) int32 {
    ret, _, _ := listBox_GetLeft.Call(obj)
    return int32(ret)
}

func ListBox_SetLeft(obj uintptr, value int32) {
   listBox_SetLeft.Call(obj, uintptr(value))
}

func ListBox_GetTop(obj uintptr) int32 {
    ret, _, _ := listBox_GetTop.Call(obj)
    return int32(ret)
}

func ListBox_SetTop(obj uintptr, value int32) {
   listBox_SetTop.Call(obj, uintptr(value))
}

func ListBox_GetWidth(obj uintptr) int32 {
    ret, _, _ := listBox_GetWidth.Call(obj)
    return int32(ret)
}

func ListBox_SetWidth(obj uintptr, value int32) {
   listBox_SetWidth.Call(obj, uintptr(value))
}

func ListBox_GetHeight(obj uintptr) int32 {
    ret, _, _ := listBox_GetHeight.Call(obj)
    return int32(ret)
}

func ListBox_SetHeight(obj uintptr, value int32) {
   listBox_SetHeight.Call(obj, uintptr(value))
}

func ListBox_GetCursor(obj uintptr) TCursor {
    ret, _, _ := listBox_GetCursor.Call(obj)
    return TCursor(ret)
}

func ListBox_SetCursor(obj uintptr, value TCursor) {
   listBox_SetCursor.Call(obj, uintptr(value))
}

func ListBox_GetHint(obj uintptr) string {
    ret, _, _ := listBox_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func ListBox_SetHint(obj uintptr, value string) {
   listBox_SetHint.Call(obj, GoStrToDStr(value))
}

func ListBox_GetMargins(obj uintptr) uintptr {
    ret, _, _ := listBox_GetMargins.Call(obj)
    return ret
}

func ListBox_SetMargins(obj uintptr, value uintptr) {
   listBox_SetMargins.Call(obj, value)
}

func ListBox_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := listBox_GetComponentCount.Call(obj)
    return int32(ret)
}

func ListBox_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := listBox_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ListBox_SetComponentIndex(obj uintptr, value int32) {
   listBox_SetComponentIndex.Call(obj, uintptr(value))
}

func ListBox_GetOwner(obj uintptr) uintptr {
    ret, _, _ := listBox_GetOwner.Call(obj)
    return ret
}

func ListBox_GetName(obj uintptr) string {
    ret, _, _ := listBox_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ListBox_SetName(obj uintptr, value string) {
   listBox_SetName.Call(obj, GoStrToDStr(value))
}

func ListBox_GetTag(obj uintptr) int {
    ret, _, _ := listBox_GetTag.Call(obj)
    return int(ret)
}

func ListBox_SetTag(obj uintptr, value int) {
   listBox_SetTag.Call(obj, uintptr(value))
}

func ListBox_GetSelected(obj uintptr, Index int32) bool {
    ret, _, _ := listBox_GetSelected.Call(obj, uintptr(Index))
    return DBoolToGoBool(ret)
}

func ListBox_SetSelected(obj uintptr, Index int32, value bool) {
   listBox_SetSelected.Call(obj, uintptr(Index), GoBoolToDBool(value))
}

func ListBox_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := listBox_GetControls.Call(obj, uintptr(Index))
    return ret
}

func ListBox_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := listBox_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

