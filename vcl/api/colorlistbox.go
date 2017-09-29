
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func ColorListBox_Create(obj uintptr) uintptr {
    ret, _, _ := colorListBox_Create.Call(obj)
    return ret
}

func ColorListBox_Free(obj uintptr) {
    colorListBox_Free.Call(obj)
}

func ColorListBox_AddItem(obj uintptr, Item string, AObject uintptr)  {
    colorListBox_AddItem.Call(obj, GoStrToDStr(Item) , AObject )
}

func ColorListBox_Clear(obj uintptr)  {
    colorListBox_Clear.Call(obj)
}

func ColorListBox_ClearSelection(obj uintptr)  {
    colorListBox_ClearSelection.Call(obj)
}

func ColorListBox_DeleteSelected(obj uintptr)  {
    colorListBox_DeleteSelected.Call(obj)
}

func ColorListBox_SelectAll(obj uintptr)  {
    colorListBox_SelectAll.Call(obj)
}

func ColorListBox_CanFocus(obj uintptr) bool {
    ret, _, _ := colorListBox_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_FlipChildren(obj uintptr, AllLevels bool)  {
    colorListBox_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func ColorListBox_Focused(obj uintptr) bool {
    ret, _, _ := colorListBox_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_HandleAllocated(obj uintptr) bool {
    ret, _, _ := colorListBox_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_Invalidate(obj uintptr)  {
    colorListBox_Invalidate.Call(obj)
}

func ColorListBox_Realign(obj uintptr)  {
    colorListBox_Realign.Call(obj)
}

func ColorListBox_Repaint(obj uintptr)  {
    colorListBox_Repaint.Call(obj)
}

func ColorListBox_ScaleBy(obj uintptr, M int32, D int32)  {
    colorListBox_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func ColorListBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    colorListBox_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func ColorListBox_SetFocus(obj uintptr)  {
    colorListBox_SetFocus.Call(obj)
}

func ColorListBox_Update(obj uintptr)  {
    colorListBox_Update.Call(obj)
}

func ColorListBox_BringToFront(obj uintptr)  {
    colorListBox_BringToFront.Call(obj)
}

func ColorListBox_HasParent(obj uintptr) bool {
    ret, _, _ := colorListBox_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_Hide(obj uintptr)  {
    colorListBox_Hide.Call(obj)
}

func ColorListBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := colorListBox_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func ColorListBox_Refresh(obj uintptr)  {
    colorListBox_Refresh.Call(obj)
}

func ColorListBox_SendToBack(obj uintptr)  {
    colorListBox_SendToBack.Call(obj)
}

func ColorListBox_Show(obj uintptr)  {
    colorListBox_Show.Call(obj)
}

func ColorListBox_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := colorListBox_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func ColorListBox_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := colorListBox_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ColorListBox_GetNamePath(obj uintptr) string {
    ret, _, _ := colorListBox_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ColorListBox_Assign(obj uintptr, Source uintptr)  {
    colorListBox_Assign.Call(obj, Source )
}

func ColorListBox_ClassName(obj uintptr) string {
    ret, _, _ := colorListBox_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ColorListBox_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := colorListBox_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ColorListBox_GetHashCode(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetHashCode.Call(obj)
    return int32(ret)
}

func ColorListBox_ToString(obj uintptr) string {
    ret, _, _ := colorListBox_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ColorListBox_GetAlign(obj uintptr) TAlign {
    ret, _, _ := colorListBox_GetAlign.Call(obj)
    return TAlign(ret)
}

func ColorListBox_SetAlign(obj uintptr, value TAlign) {
   colorListBox_SetAlign.Call(obj, uintptr(value))
}

func ColorListBox_GetAutoComplete(obj uintptr) bool {
    ret, _, _ := colorListBox_GetAutoComplete.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetAutoComplete(obj uintptr, value bool) {
   colorListBox_SetAutoComplete.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetDefaultColorColor(obj uintptr) TColor {
    ret, _, _ := colorListBox_GetDefaultColorColor.Call(obj)
    return TColor(ret)
}

func ColorListBox_SetDefaultColorColor(obj uintptr, value TColor) {
   colorListBox_SetDefaultColorColor.Call(obj, uintptr(value))
}

func ColorListBox_GetNoneColorColor(obj uintptr) TColor {
    ret, _, _ := colorListBox_GetNoneColorColor.Call(obj)
    return TColor(ret)
}

func ColorListBox_SetNoneColorColor(obj uintptr, value TColor) {
   colorListBox_SetNoneColorColor.Call(obj, uintptr(value))
}

func ColorListBox_GetSelected(obj uintptr) TColor {
    ret, _, _ := colorListBox_GetSelected.Call(obj)
    return TColor(ret)
}

func ColorListBox_SetSelected(obj uintptr, value TColor) {
   colorListBox_SetSelected.Call(obj, uintptr(value))
}

func ColorListBox_GetStyle(obj uintptr) TColorBoxStyle {
    ret, _, _ := colorListBox_GetStyle.Call(obj)
    return TColorBoxStyle(ret)
}

func ColorListBox_SetStyle(obj uintptr, value TColorBoxStyle) {
   colorListBox_SetStyle.Call(obj, uintptr(value))
}

func ColorListBox_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := colorListBox_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func ColorListBox_SetAnchors(obj uintptr, value TAnchors) {
   colorListBox_SetAnchors.Call(obj, uintptr(value))
}

func ColorListBox_GetColor(obj uintptr) TColor {
    ret, _, _ := colorListBox_GetColor.Call(obj)
    return TColor(ret)
}

func ColorListBox_SetColor(obj uintptr, value TColor) {
   colorListBox_SetColor.Call(obj, uintptr(value))
}

func ColorListBox_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := colorListBox_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetDoubleBuffered(obj uintptr, value bool) {
   colorListBox_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetEnabled(obj uintptr) bool {
    ret, _, _ := colorListBox_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetEnabled(obj uintptr, value bool) {
   colorListBox_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetFont(obj uintptr) uintptr {
    ret, _, _ := colorListBox_GetFont.Call(obj)
    return ret
}

func ColorListBox_SetFont(obj uintptr, value uintptr) {
   colorListBox_SetFont.Call(obj, value)
}

func ColorListBox_GetItemHeight(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetItemHeight.Call(obj)
    return int32(ret)
}

func ColorListBox_SetItemHeight(obj uintptr, value int32) {
   colorListBox_SetItemHeight.Call(obj, uintptr(value))
}

func ColorListBox_GetParentColor(obj uintptr) bool {
    ret, _, _ := colorListBox_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetParentColor(obj uintptr, value bool) {
   colorListBox_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := colorListBox_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetParentCtl3D(obj uintptr, value bool) {
   colorListBox_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetParentFont(obj uintptr) bool {
    ret, _, _ := colorListBox_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetParentFont(obj uintptr, value bool) {
   colorListBox_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := colorListBox_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetParentShowHint(obj uintptr, value bool) {
   colorListBox_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := colorListBox_GetPopupMenu.Call(obj)
    return ret
}

func ColorListBox_SetPopupMenu(obj uintptr, value uintptr) {
   colorListBox_SetPopupMenu.Call(obj, value)
}

func ColorListBox_GetShowHint(obj uintptr) bool {
    ret, _, _ := colorListBox_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetShowHint(obj uintptr, value bool) {
   colorListBox_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := colorListBox_GetTabOrder.Call(obj)
    return int16(ret)
}

func ColorListBox_SetTabOrder(obj uintptr, value int16) {
   colorListBox_SetTabOrder.Call(obj, uintptr(value))
}

func ColorListBox_GetTabStop(obj uintptr) bool {
    ret, _, _ := colorListBox_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetTabStop(obj uintptr, value bool) {
   colorListBox_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetVisible(obj uintptr) bool {
    ret, _, _ := colorListBox_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetVisible(obj uintptr, value bool) {
   colorListBox_SetVisible.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_SetOnClick(obj uintptr, fn interface{}) {
    colorListBox_SetOnClick.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnDblClick(obj uintptr, fn interface{}) {
    colorListBox_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnEnter(obj uintptr, fn interface{}) {
    colorListBox_SetOnEnter.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnExit(obj uintptr, fn interface{}) {
    colorListBox_SetOnExit.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnKeyDown(obj uintptr, fn interface{}) {
    colorListBox_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnKeyPress(obj uintptr, fn interface{}) {
    colorListBox_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnKeyUp(obj uintptr, fn interface{}) {
    colorListBox_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnMouseDown(obj uintptr, fn interface{}) {
    colorListBox_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
    colorListBox_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
    colorListBox_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnMouseMove(obj uintptr, fn interface{}) {
    colorListBox_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnMouseUp(obj uintptr, fn interface{}) {
    colorListBox_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func ColorListBox_GetAutoCompleteDelay(obj uintptr) uint32 {
    ret, _, _ := colorListBox_GetAutoCompleteDelay.Call(obj)
    return uint32(ret)
}

func ColorListBox_SetAutoCompleteDelay(obj uintptr, value uint32) {
   colorListBox_SetAutoCompleteDelay.Call(obj, uintptr(value))
}

func ColorListBox_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := colorListBox_GetCanvas.Call(obj)
    return ret
}

func ColorListBox_GetItems(obj uintptr) uintptr {
    ret, _, _ := colorListBox_GetItems.Call(obj)
    return ret
}

func ColorListBox_SetItems(obj uintptr, value uintptr) {
   colorListBox_SetItems.Call(obj, value)
}

func ColorListBox_GetMultiSelect(obj uintptr) bool {
    ret, _, _ := colorListBox_GetMultiSelect.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetMultiSelect(obj uintptr, value bool) {
   colorListBox_SetMultiSelect.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetSelCount(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetSelCount.Call(obj)
    return int32(ret)
}

func ColorListBox_GetItemIndex(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetItemIndex.Call(obj)
    return int32(ret)
}

func ColorListBox_SetItemIndex(obj uintptr, value int32) {
   colorListBox_SetItemIndex.Call(obj, uintptr(value))
}

func ColorListBox_GetBrush(obj uintptr) uintptr {
    ret, _, _ := colorListBox_GetBrush.Call(obj)
    return ret
}

func ColorListBox_GetControlCount(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetControlCount.Call(obj)
    return int32(ret)
}

func ColorListBox_GetHandle(obj uintptr) HWND {
    ret, _, _ := colorListBox_GetHandle.Call(obj)
    return HWND(ret)
}

func ColorListBox_GetAction(obj uintptr) uintptr {
    ret, _, _ := colorListBox_GetAction.Call(obj)
    return ret
}

func ColorListBox_SetAction(obj uintptr, value uintptr) {
   colorListBox_SetAction.Call(obj, value)
}

func ColorListBox_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    colorListBox_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ColorListBox_SetBoundsRect(obj uintptr, value TRect) {
   colorListBox_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ColorListBox_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetClientHeight.Call(obj)
    return int32(ret)
}

func ColorListBox_SetClientHeight(obj uintptr, value int32) {
   colorListBox_SetClientHeight.Call(obj, uintptr(value))
}

func ColorListBox_GetClientRect(obj uintptr) TRect {
    var ret TRect
    colorListBox_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ColorListBox_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetClientWidth.Call(obj)
    return int32(ret)
}

func ColorListBox_SetClientWidth(obj uintptr, value int32) {
   colorListBox_SetClientWidth.Call(obj, uintptr(value))
}

func ColorListBox_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func ColorListBox_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetExplicitTop.Call(obj)
    return int32(ret)
}

func ColorListBox_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func ColorListBox_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func ColorListBox_GetParent(obj uintptr) uintptr {
    ret, _, _ := colorListBox_GetParent.Call(obj)
    return ret
}

func ColorListBox_SetParent(obj uintptr, value uintptr) {
   colorListBox_SetParent.Call(obj, value)
}

func ColorListBox_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := colorListBox_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetAlignWithMargins(obj uintptr, value bool) {
   colorListBox_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetLeft(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetLeft.Call(obj)
    return int32(ret)
}

func ColorListBox_SetLeft(obj uintptr, value int32) {
   colorListBox_SetLeft.Call(obj, uintptr(value))
}

func ColorListBox_GetTop(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetTop.Call(obj)
    return int32(ret)
}

func ColorListBox_SetTop(obj uintptr, value int32) {
   colorListBox_SetTop.Call(obj, uintptr(value))
}

func ColorListBox_GetWidth(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetWidth.Call(obj)
    return int32(ret)
}

func ColorListBox_SetWidth(obj uintptr, value int32) {
   colorListBox_SetWidth.Call(obj, uintptr(value))
}

func ColorListBox_GetHeight(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetHeight.Call(obj)
    return int32(ret)
}

func ColorListBox_SetHeight(obj uintptr, value int32) {
   colorListBox_SetHeight.Call(obj, uintptr(value))
}

func ColorListBox_GetCursor(obj uintptr) TCursor {
    ret, _, _ := colorListBox_GetCursor.Call(obj)
    return TCursor(ret)
}

func ColorListBox_SetCursor(obj uintptr, value TCursor) {
   colorListBox_SetCursor.Call(obj, uintptr(value))
}

func ColorListBox_GetHint(obj uintptr) string {
    ret, _, _ := colorListBox_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func ColorListBox_SetHint(obj uintptr, value string) {
   colorListBox_SetHint.Call(obj, GoStrToDStr(value))
}

func ColorListBox_GetMargins(obj uintptr) uintptr {
    ret, _, _ := colorListBox_GetMargins.Call(obj)
    return ret
}

func ColorListBox_SetMargins(obj uintptr, value uintptr) {
   colorListBox_SetMargins.Call(obj, value)
}

func ColorListBox_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetComponentCount.Call(obj)
    return int32(ret)
}

func ColorListBox_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ColorListBox_SetComponentIndex(obj uintptr, value int32) {
   colorListBox_SetComponentIndex.Call(obj, uintptr(value))
}

func ColorListBox_GetOwner(obj uintptr) uintptr {
    ret, _, _ := colorListBox_GetOwner.Call(obj)
    return ret
}

func ColorListBox_GetName(obj uintptr) string {
    ret, _, _ := colorListBox_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ColorListBox_SetName(obj uintptr, value string) {
   colorListBox_SetName.Call(obj, GoStrToDStr(value))
}

func ColorListBox_GetTag(obj uintptr) int {
    ret, _, _ := colorListBox_GetTag.Call(obj)
    return int(ret)
}

func ColorListBox_SetTag(obj uintptr, value int) {
   colorListBox_SetTag.Call(obj, uintptr(value))
}

func ColorListBox_GetColors(obj uintptr, Index int32) TColor {
    ret, _, _ := colorListBox_GetColors.Call(obj, uintptr(Index))
    return TColor(ret)
}

func ColorListBox_GetColorNames(obj uintptr, Index int32) string {
    ret, _, _ := colorListBox_GetColorNames.Call(obj, uintptr(Index))
    return DStrToGoStr(ret)
}

func ColorListBox_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := colorListBox_GetControls.Call(obj, uintptr(Index))
    return ret
}

func ColorListBox_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := colorListBox_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

