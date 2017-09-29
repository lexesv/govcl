
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func ColorBox_Create(obj uintptr) uintptr {
    ret, _, _ := colorBox_Create.Call(obj)
    return ret
}

func ColorBox_Free(obj uintptr) {
    colorBox_Free.Call(obj)
}

func ColorBox_AddItem(obj uintptr, Item string, AObject uintptr)  {
    colorBox_AddItem.Call(obj, GoStrToDStr(Item) , AObject )
}

func ColorBox_Clear(obj uintptr)  {
    colorBox_Clear.Call(obj)
}

func ColorBox_ClearSelection(obj uintptr)  {
    colorBox_ClearSelection.Call(obj)
}

func ColorBox_DeleteSelected(obj uintptr)  {
    colorBox_DeleteSelected.Call(obj)
}

func ColorBox_Focused(obj uintptr) bool {
    ret, _, _ := colorBox_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SelectAll(obj uintptr)  {
    colorBox_SelectAll.Call(obj)
}

func ColorBox_CanFocus(obj uintptr) bool {
    ret, _, _ := colorBox_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_FlipChildren(obj uintptr, AllLevels bool)  {
    colorBox_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func ColorBox_HandleAllocated(obj uintptr) bool {
    ret, _, _ := colorBox_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_Invalidate(obj uintptr)  {
    colorBox_Invalidate.Call(obj)
}

func ColorBox_Realign(obj uintptr)  {
    colorBox_Realign.Call(obj)
}

func ColorBox_Repaint(obj uintptr)  {
    colorBox_Repaint.Call(obj)
}

func ColorBox_ScaleBy(obj uintptr, M int32, D int32)  {
    colorBox_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func ColorBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    colorBox_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func ColorBox_SetFocus(obj uintptr)  {
    colorBox_SetFocus.Call(obj)
}

func ColorBox_Update(obj uintptr)  {
    colorBox_Update.Call(obj)
}

func ColorBox_BringToFront(obj uintptr)  {
    colorBox_BringToFront.Call(obj)
}

func ColorBox_HasParent(obj uintptr) bool {
    ret, _, _ := colorBox_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_Hide(obj uintptr)  {
    colorBox_Hide.Call(obj)
}

func ColorBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := colorBox_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func ColorBox_Refresh(obj uintptr)  {
    colorBox_Refresh.Call(obj)
}

func ColorBox_SendToBack(obj uintptr)  {
    colorBox_SendToBack.Call(obj)
}

func ColorBox_Show(obj uintptr)  {
    colorBox_Show.Call(obj)
}

func ColorBox_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := colorBox_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func ColorBox_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := colorBox_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ColorBox_GetNamePath(obj uintptr) string {
    ret, _, _ := colorBox_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ColorBox_Assign(obj uintptr, Source uintptr)  {
    colorBox_Assign.Call(obj, Source )
}

func ColorBox_ClassName(obj uintptr) string {
    ret, _, _ := colorBox_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ColorBox_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := colorBox_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ColorBox_GetHashCode(obj uintptr) int32 {
    ret, _, _ := colorBox_GetHashCode.Call(obj)
    return int32(ret)
}

func ColorBox_ToString(obj uintptr) string {
    ret, _, _ := colorBox_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ColorBox_GetAlign(obj uintptr) TAlign {
    ret, _, _ := colorBox_GetAlign.Call(obj)
    return TAlign(ret)
}

func ColorBox_SetAlign(obj uintptr, value TAlign) {
   colorBox_SetAlign.Call(obj, uintptr(value))
}

func ColorBox_GetAutoComplete(obj uintptr) bool {
    ret, _, _ := colorBox_GetAutoComplete.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetAutoComplete(obj uintptr, value bool) {
   colorBox_SetAutoComplete.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetAutoDropDown(obj uintptr) bool {
    ret, _, _ := colorBox_GetAutoDropDown.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetAutoDropDown(obj uintptr, value bool) {
   colorBox_SetAutoDropDown.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetDefaultColorColor(obj uintptr) TColor {
    ret, _, _ := colorBox_GetDefaultColorColor.Call(obj)
    return TColor(ret)
}

func ColorBox_SetDefaultColorColor(obj uintptr, value TColor) {
   colorBox_SetDefaultColorColor.Call(obj, uintptr(value))
}

func ColorBox_GetNoneColorColor(obj uintptr) TColor {
    ret, _, _ := colorBox_GetNoneColorColor.Call(obj)
    return TColor(ret)
}

func ColorBox_SetNoneColorColor(obj uintptr, value TColor) {
   colorBox_SetNoneColorColor.Call(obj, uintptr(value))
}

func ColorBox_GetSelected(obj uintptr) TColor {
    ret, _, _ := colorBox_GetSelected.Call(obj)
    return TColor(ret)
}

func ColorBox_SetSelected(obj uintptr, value TColor) {
   colorBox_SetSelected.Call(obj, uintptr(value))
}

func ColorBox_GetStyle(obj uintptr) TColorBoxStyle {
    ret, _, _ := colorBox_GetStyle.Call(obj)
    return TColorBoxStyle(ret)
}

func ColorBox_SetStyle(obj uintptr, value TColorBoxStyle) {
   colorBox_SetStyle.Call(obj, uintptr(value))
}

func ColorBox_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := colorBox_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func ColorBox_SetAnchors(obj uintptr, value TAnchors) {
   colorBox_SetAnchors.Call(obj, uintptr(value))
}

func ColorBox_GetColor(obj uintptr) TColor {
    ret, _, _ := colorBox_GetColor.Call(obj)
    return TColor(ret)
}

func ColorBox_SetColor(obj uintptr, value TColor) {
   colorBox_SetColor.Call(obj, uintptr(value))
}

func ColorBox_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := colorBox_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetDoubleBuffered(obj uintptr, value bool) {
   colorBox_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetDropDownCount(obj uintptr) int32 {
    ret, _, _ := colorBox_GetDropDownCount.Call(obj)
    return int32(ret)
}

func ColorBox_SetDropDownCount(obj uintptr, value int32) {
   colorBox_SetDropDownCount.Call(obj, uintptr(value))
}

func ColorBox_GetEnabled(obj uintptr) bool {
    ret, _, _ := colorBox_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetEnabled(obj uintptr, value bool) {
   colorBox_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetFont(obj uintptr) uintptr {
    ret, _, _ := colorBox_GetFont.Call(obj)
    return ret
}

func ColorBox_SetFont(obj uintptr, value uintptr) {
   colorBox_SetFont.Call(obj, value)
}

func ColorBox_GetItemHeight(obj uintptr) int32 {
    ret, _, _ := colorBox_GetItemHeight.Call(obj)
    return int32(ret)
}

func ColorBox_SetItemHeight(obj uintptr, value int32) {
   colorBox_SetItemHeight.Call(obj, uintptr(value))
}

func ColorBox_GetParentColor(obj uintptr) bool {
    ret, _, _ := colorBox_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetParentColor(obj uintptr, value bool) {
   colorBox_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := colorBox_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetParentCtl3D(obj uintptr, value bool) {
   colorBox_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetParentFont(obj uintptr) bool {
    ret, _, _ := colorBox_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetParentFont(obj uintptr, value bool) {
   colorBox_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := colorBox_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetParentShowHint(obj uintptr, value bool) {
   colorBox_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := colorBox_GetPopupMenu.Call(obj)
    return ret
}

func ColorBox_SetPopupMenu(obj uintptr, value uintptr) {
   colorBox_SetPopupMenu.Call(obj, value)
}

func ColorBox_GetShowHint(obj uintptr) bool {
    ret, _, _ := colorBox_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetShowHint(obj uintptr, value bool) {
   colorBox_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := colorBox_GetTabOrder.Call(obj)
    return int16(ret)
}

func ColorBox_SetTabOrder(obj uintptr, value int16) {
   colorBox_SetTabOrder.Call(obj, uintptr(value))
}

func ColorBox_GetTabStop(obj uintptr) bool {
    ret, _, _ := colorBox_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetTabStop(obj uintptr, value bool) {
   colorBox_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetVisible(obj uintptr) bool {
    ret, _, _ := colorBox_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetVisible(obj uintptr, value bool) {
   colorBox_SetVisible.Call(obj, GoBoolToDBool(value))
}

func ColorBox_SetOnChange(obj uintptr, fn interface{}) {
    colorBox_SetOnChange.Call(obj, addEventToMap(fn))
}

func ColorBox_SetOnClick(obj uintptr, fn interface{}) {
    colorBox_SetOnClick.Call(obj, addEventToMap(fn))
}

func ColorBox_SetOnEnter(obj uintptr, fn interface{}) {
    colorBox_SetOnEnter.Call(obj, addEventToMap(fn))
}

func ColorBox_SetOnExit(obj uintptr, fn interface{}) {
    colorBox_SetOnExit.Call(obj, addEventToMap(fn))
}

func ColorBox_SetOnKeyDown(obj uintptr, fn interface{}) {
    colorBox_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func ColorBox_SetOnKeyPress(obj uintptr, fn interface{}) {
    colorBox_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func ColorBox_SetOnKeyUp(obj uintptr, fn interface{}) {
    colorBox_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func ColorBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
    colorBox_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func ColorBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
    colorBox_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func ColorBox_GetAutoCompleteDelay(obj uintptr) uint32 {
    ret, _, _ := colorBox_GetAutoCompleteDelay.Call(obj)
    return uint32(ret)
}

func ColorBox_SetAutoCompleteDelay(obj uintptr, value uint32) {
   colorBox_SetAutoCompleteDelay.Call(obj, uintptr(value))
}

func ColorBox_GetAutoCloseUp(obj uintptr) bool {
    ret, _, _ := colorBox_GetAutoCloseUp.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetAutoCloseUp(obj uintptr, value bool) {
   colorBox_SetAutoCloseUp.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetSelText(obj uintptr) string {
    ret, _, _ := colorBox_GetSelText.Call(obj)
    return DStrToGoStr(ret)
}

func ColorBox_SetSelText(obj uintptr, value string) {
   colorBox_SetSelText.Call(obj, GoStrToDStr(value))
}

func ColorBox_GetTextHint(obj uintptr) string {
    ret, _, _ := colorBox_GetTextHint.Call(obj)
    return DStrToGoStr(ret)
}

func ColorBox_SetTextHint(obj uintptr, value string) {
   colorBox_SetTextHint.Call(obj, GoStrToDStr(value))
}

func ColorBox_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := colorBox_GetCanvas.Call(obj)
    return ret
}

func ColorBox_GetItems(obj uintptr) uintptr {
    ret, _, _ := colorBox_GetItems.Call(obj)
    return ret
}

func ColorBox_SetItems(obj uintptr, value uintptr) {
   colorBox_SetItems.Call(obj, value)
}

func ColorBox_GetSelLength(obj uintptr) int32 {
    ret, _, _ := colorBox_GetSelLength.Call(obj)
    return int32(ret)
}

func ColorBox_SetSelLength(obj uintptr, value int32) {
   colorBox_SetSelLength.Call(obj, uintptr(value))
}

func ColorBox_GetSelStart(obj uintptr) int32 {
    ret, _, _ := colorBox_GetSelStart.Call(obj)
    return int32(ret)
}

func ColorBox_SetSelStart(obj uintptr, value int32) {
   colorBox_SetSelStart.Call(obj, uintptr(value))
}

func ColorBox_GetItemIndex(obj uintptr) int32 {
    ret, _, _ := colorBox_GetItemIndex.Call(obj)
    return int32(ret)
}

func ColorBox_SetItemIndex(obj uintptr, value int32) {
   colorBox_SetItemIndex.Call(obj, uintptr(value))
}

func ColorBox_GetBrush(obj uintptr) uintptr {
    ret, _, _ := colorBox_GetBrush.Call(obj)
    return ret
}

func ColorBox_GetControlCount(obj uintptr) int32 {
    ret, _, _ := colorBox_GetControlCount.Call(obj)
    return int32(ret)
}

func ColorBox_GetHandle(obj uintptr) HWND {
    ret, _, _ := colorBox_GetHandle.Call(obj)
    return HWND(ret)
}

func ColorBox_GetAction(obj uintptr) uintptr {
    ret, _, _ := colorBox_GetAction.Call(obj)
    return ret
}

func ColorBox_SetAction(obj uintptr, value uintptr) {
   colorBox_SetAction.Call(obj, value)
}

func ColorBox_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    colorBox_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ColorBox_SetBoundsRect(obj uintptr, value TRect) {
   colorBox_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ColorBox_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := colorBox_GetClientHeight.Call(obj)
    return int32(ret)
}

func ColorBox_SetClientHeight(obj uintptr, value int32) {
   colorBox_SetClientHeight.Call(obj, uintptr(value))
}

func ColorBox_GetClientRect(obj uintptr) TRect {
    var ret TRect
    colorBox_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ColorBox_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := colorBox_GetClientWidth.Call(obj)
    return int32(ret)
}

func ColorBox_SetClientWidth(obj uintptr, value int32) {
   colorBox_SetClientWidth.Call(obj, uintptr(value))
}

func ColorBox_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := colorBox_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func ColorBox_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := colorBox_GetExplicitTop.Call(obj)
    return int32(ret)
}

func ColorBox_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := colorBox_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func ColorBox_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := colorBox_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func ColorBox_GetParent(obj uintptr) uintptr {
    ret, _, _ := colorBox_GetParent.Call(obj)
    return ret
}

func ColorBox_SetParent(obj uintptr, value uintptr) {
   colorBox_SetParent.Call(obj, value)
}

func ColorBox_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := colorBox_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetAlignWithMargins(obj uintptr, value bool) {
   colorBox_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetLeft(obj uintptr) int32 {
    ret, _, _ := colorBox_GetLeft.Call(obj)
    return int32(ret)
}

func ColorBox_SetLeft(obj uintptr, value int32) {
   colorBox_SetLeft.Call(obj, uintptr(value))
}

func ColorBox_GetTop(obj uintptr) int32 {
    ret, _, _ := colorBox_GetTop.Call(obj)
    return int32(ret)
}

func ColorBox_SetTop(obj uintptr, value int32) {
   colorBox_SetTop.Call(obj, uintptr(value))
}

func ColorBox_GetWidth(obj uintptr) int32 {
    ret, _, _ := colorBox_GetWidth.Call(obj)
    return int32(ret)
}

func ColorBox_SetWidth(obj uintptr, value int32) {
   colorBox_SetWidth.Call(obj, uintptr(value))
}

func ColorBox_GetHeight(obj uintptr) int32 {
    ret, _, _ := colorBox_GetHeight.Call(obj)
    return int32(ret)
}

func ColorBox_SetHeight(obj uintptr, value int32) {
   colorBox_SetHeight.Call(obj, uintptr(value))
}

func ColorBox_GetCursor(obj uintptr) TCursor {
    ret, _, _ := colorBox_GetCursor.Call(obj)
    return TCursor(ret)
}

func ColorBox_SetCursor(obj uintptr, value TCursor) {
   colorBox_SetCursor.Call(obj, uintptr(value))
}

func ColorBox_GetHint(obj uintptr) string {
    ret, _, _ := colorBox_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func ColorBox_SetHint(obj uintptr, value string) {
   colorBox_SetHint.Call(obj, GoStrToDStr(value))
}

func ColorBox_GetMargins(obj uintptr) uintptr {
    ret, _, _ := colorBox_GetMargins.Call(obj)
    return ret
}

func ColorBox_SetMargins(obj uintptr, value uintptr) {
   colorBox_SetMargins.Call(obj, value)
}

func ColorBox_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := colorBox_GetComponentCount.Call(obj)
    return int32(ret)
}

func ColorBox_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := colorBox_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ColorBox_SetComponentIndex(obj uintptr, value int32) {
   colorBox_SetComponentIndex.Call(obj, uintptr(value))
}

func ColorBox_GetOwner(obj uintptr) uintptr {
    ret, _, _ := colorBox_GetOwner.Call(obj)
    return ret
}

func ColorBox_GetName(obj uintptr) string {
    ret, _, _ := colorBox_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ColorBox_SetName(obj uintptr, value string) {
   colorBox_SetName.Call(obj, GoStrToDStr(value))
}

func ColorBox_GetTag(obj uintptr) int {
    ret, _, _ := colorBox_GetTag.Call(obj)
    return int(ret)
}

func ColorBox_SetTag(obj uintptr, value int) {
   colorBox_SetTag.Call(obj, uintptr(value))
}

func ColorBox_GetColors(obj uintptr, Index int32) TColor {
    ret, _, _ := colorBox_GetColors.Call(obj, uintptr(Index))
    return TColor(ret)
}

func ColorBox_GetColorNames(obj uintptr, Index int32) string {
    ret, _, _ := colorBox_GetColorNames.Call(obj, uintptr(Index))
    return DStrToGoStr(ret)
}

func ColorBox_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := colorBox_GetControls.Call(obj, uintptr(Index))
    return ret
}

func ColorBox_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := colorBox_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

