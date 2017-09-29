
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func ComboBox_Create(obj uintptr) uintptr {
    ret, _, _ := comboBox_Create.Call(obj)
    return ret
}

func ComboBox_Free(obj uintptr) {
    comboBox_Free.Call(obj)
}

func ComboBox_AddItem(obj uintptr, Item string, AObject uintptr)  {
    comboBox_AddItem.Call(obj, GoStrToDStr(Item) , AObject )
}

func ComboBox_Clear(obj uintptr)  {
    comboBox_Clear.Call(obj)
}

func ComboBox_ClearSelection(obj uintptr)  {
    comboBox_ClearSelection.Call(obj)
}

func ComboBox_DeleteSelected(obj uintptr)  {
    comboBox_DeleteSelected.Call(obj)
}

func ComboBox_Focused(obj uintptr) bool {
    ret, _, _ := comboBox_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SelectAll(obj uintptr)  {
    comboBox_SelectAll.Call(obj)
}

func ComboBox_CanFocus(obj uintptr) bool {
    ret, _, _ := comboBox_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_FlipChildren(obj uintptr, AllLevels bool)  {
    comboBox_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func ComboBox_HandleAllocated(obj uintptr) bool {
    ret, _, _ := comboBox_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_Invalidate(obj uintptr)  {
    comboBox_Invalidate.Call(obj)
}

func ComboBox_Realign(obj uintptr)  {
    comboBox_Realign.Call(obj)
}

func ComboBox_Repaint(obj uintptr)  {
    comboBox_Repaint.Call(obj)
}

func ComboBox_ScaleBy(obj uintptr, M int32, D int32)  {
    comboBox_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func ComboBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    comboBox_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func ComboBox_SetFocus(obj uintptr)  {
    comboBox_SetFocus.Call(obj)
}

func ComboBox_Update(obj uintptr)  {
    comboBox_Update.Call(obj)
}

func ComboBox_BringToFront(obj uintptr)  {
    comboBox_BringToFront.Call(obj)
}

func ComboBox_HasParent(obj uintptr) bool {
    ret, _, _ := comboBox_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_Hide(obj uintptr)  {
    comboBox_Hide.Call(obj)
}

func ComboBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := comboBox_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func ComboBox_Refresh(obj uintptr)  {
    comboBox_Refresh.Call(obj)
}

func ComboBox_SendToBack(obj uintptr)  {
    comboBox_SendToBack.Call(obj)
}

func ComboBox_Show(obj uintptr)  {
    comboBox_Show.Call(obj)
}

func ComboBox_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := comboBox_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func ComboBox_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := comboBox_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ComboBox_GetNamePath(obj uintptr) string {
    ret, _, _ := comboBox_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ComboBox_Assign(obj uintptr, Source uintptr)  {
    comboBox_Assign.Call(obj, Source )
}

func ComboBox_ClassName(obj uintptr) string {
    ret, _, _ := comboBox_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ComboBox_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := comboBox_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ComboBox_GetHashCode(obj uintptr) int32 {
    ret, _, _ := comboBox_GetHashCode.Call(obj)
    return int32(ret)
}

func ComboBox_ToString(obj uintptr) string {
    ret, _, _ := comboBox_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ComboBox_GetAlign(obj uintptr) TAlign {
    ret, _, _ := comboBox_GetAlign.Call(obj)
    return TAlign(ret)
}

func ComboBox_SetAlign(obj uintptr, value TAlign) {
   comboBox_SetAlign.Call(obj, uintptr(value))
}

func ComboBox_GetAutoComplete(obj uintptr) bool {
    ret, _, _ := comboBox_GetAutoComplete.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetAutoComplete(obj uintptr, value bool) {
   comboBox_SetAutoComplete.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetAutoCompleteDelay(obj uintptr) uint32 {
    ret, _, _ := comboBox_GetAutoCompleteDelay.Call(obj)
    return uint32(ret)
}

func ComboBox_SetAutoCompleteDelay(obj uintptr, value uint32) {
   comboBox_SetAutoCompleteDelay.Call(obj, uintptr(value))
}

func ComboBox_GetAutoDropDown(obj uintptr) bool {
    ret, _, _ := comboBox_GetAutoDropDown.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetAutoDropDown(obj uintptr, value bool) {
   comboBox_SetAutoDropDown.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetAutoCloseUp(obj uintptr) bool {
    ret, _, _ := comboBox_GetAutoCloseUp.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetAutoCloseUp(obj uintptr, value bool) {
   comboBox_SetAutoCloseUp.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetStyle(obj uintptr) TComboBoxStyle {
    ret, _, _ := comboBox_GetStyle.Call(obj)
    return TComboBoxStyle(ret)
}

func ComboBox_SetStyle(obj uintptr, value TComboBoxStyle) {
   comboBox_SetStyle.Call(obj, uintptr(value))
}

func ComboBox_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := comboBox_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func ComboBox_SetAnchors(obj uintptr, value TAnchors) {
   comboBox_SetAnchors.Call(obj, uintptr(value))
}

func ComboBox_GetColor(obj uintptr) TColor {
    ret, _, _ := comboBox_GetColor.Call(obj)
    return TColor(ret)
}

func ComboBox_SetColor(obj uintptr, value TColor) {
   comboBox_SetColor.Call(obj, uintptr(value))
}

func ComboBox_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := comboBox_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetDoubleBuffered(obj uintptr, value bool) {
   comboBox_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetDropDownCount(obj uintptr) int32 {
    ret, _, _ := comboBox_GetDropDownCount.Call(obj)
    return int32(ret)
}

func ComboBox_SetDropDownCount(obj uintptr, value int32) {
   comboBox_SetDropDownCount.Call(obj, uintptr(value))
}

func ComboBox_GetEnabled(obj uintptr) bool {
    ret, _, _ := comboBox_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetEnabled(obj uintptr, value bool) {
   comboBox_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetFont(obj uintptr) uintptr {
    ret, _, _ := comboBox_GetFont.Call(obj)
    return ret
}

func ComboBox_SetFont(obj uintptr, value uintptr) {
   comboBox_SetFont.Call(obj, value)
}

func ComboBox_GetItemHeight(obj uintptr) int32 {
    ret, _, _ := comboBox_GetItemHeight.Call(obj)
    return int32(ret)
}

func ComboBox_SetItemHeight(obj uintptr, value int32) {
   comboBox_SetItemHeight.Call(obj, uintptr(value))
}

func ComboBox_GetItemIndex(obj uintptr) int32 {
    ret, _, _ := comboBox_GetItemIndex.Call(obj)
    return int32(ret)
}

func ComboBox_SetItemIndex(obj uintptr, value int32) {
   comboBox_SetItemIndex.Call(obj, uintptr(value))
}

func ComboBox_GetMaxLength(obj uintptr) int32 {
    ret, _, _ := comboBox_GetMaxLength.Call(obj)
    return int32(ret)
}

func ComboBox_SetMaxLength(obj uintptr, value int32) {
   comboBox_SetMaxLength.Call(obj, uintptr(value))
}

func ComboBox_GetParentColor(obj uintptr) bool {
    ret, _, _ := comboBox_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetParentColor(obj uintptr, value bool) {
   comboBox_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := comboBox_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetParentCtl3D(obj uintptr, value bool) {
   comboBox_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetParentFont(obj uintptr) bool {
    ret, _, _ := comboBox_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetParentFont(obj uintptr, value bool) {
   comboBox_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := comboBox_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetParentShowHint(obj uintptr, value bool) {
   comboBox_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := comboBox_GetPopupMenu.Call(obj)
    return ret
}

func ComboBox_SetPopupMenu(obj uintptr, value uintptr) {
   comboBox_SetPopupMenu.Call(obj, value)
}

func ComboBox_GetShowHint(obj uintptr) bool {
    ret, _, _ := comboBox_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetShowHint(obj uintptr, value bool) {
   comboBox_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetSorted(obj uintptr) bool {
    ret, _, _ := comboBox_GetSorted.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetSorted(obj uintptr, value bool) {
   comboBox_SetSorted.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := comboBox_GetTabOrder.Call(obj)
    return int16(ret)
}

func ComboBox_SetTabOrder(obj uintptr, value int16) {
   comboBox_SetTabOrder.Call(obj, uintptr(value))
}

func ComboBox_GetTabStop(obj uintptr) bool {
    ret, _, _ := comboBox_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetTabStop(obj uintptr, value bool) {
   comboBox_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetText(obj uintptr) string {
    ret, _, _ := comboBox_GetText.Call(obj)
    return DStrToGoStr(ret)
}

func ComboBox_SetText(obj uintptr, value string) {
   comboBox_SetText.Call(obj, GoStrToDStr(value))
}

func ComboBox_GetTextHint(obj uintptr) string {
    ret, _, _ := comboBox_GetTextHint.Call(obj)
    return DStrToGoStr(ret)
}

func ComboBox_SetTextHint(obj uintptr, value string) {
   comboBox_SetTextHint.Call(obj, GoStrToDStr(value))
}

func ComboBox_GetVisible(obj uintptr) bool {
    ret, _, _ := comboBox_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetVisible(obj uintptr, value bool) {
   comboBox_SetVisible.Call(obj, GoBoolToDBool(value))
}

func ComboBox_SetOnChange(obj uintptr, fn interface{}) {
    comboBox_SetOnChange.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnClick(obj uintptr, fn interface{}) {
    comboBox_SetOnClick.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnDblClick(obj uintptr, fn interface{}) {
    comboBox_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnEnter(obj uintptr, fn interface{}) {
    comboBox_SetOnEnter.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnExit(obj uintptr, fn interface{}) {
    comboBox_SetOnExit.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnKeyDown(obj uintptr, fn interface{}) {
    comboBox_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnKeyPress(obj uintptr, fn interface{}) {
    comboBox_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnKeyUp(obj uintptr, fn interface{}) {
    comboBox_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
    comboBox_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
    comboBox_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func ComboBox_GetItems(obj uintptr) uintptr {
    ret, _, _ := comboBox_GetItems.Call(obj)
    return ret
}

func ComboBox_SetItems(obj uintptr, value uintptr) {
   comboBox_SetItems.Call(obj, value)
}

func ComboBox_GetSelText(obj uintptr) string {
    ret, _, _ := comboBox_GetSelText.Call(obj)
    return DStrToGoStr(ret)
}

func ComboBox_SetSelText(obj uintptr, value string) {
   comboBox_SetSelText.Call(obj, GoStrToDStr(value))
}

func ComboBox_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := comboBox_GetCanvas.Call(obj)
    return ret
}

func ComboBox_GetSelLength(obj uintptr) int32 {
    ret, _, _ := comboBox_GetSelLength.Call(obj)
    return int32(ret)
}

func ComboBox_SetSelLength(obj uintptr, value int32) {
   comboBox_SetSelLength.Call(obj, uintptr(value))
}

func ComboBox_GetSelStart(obj uintptr) int32 {
    ret, _, _ := comboBox_GetSelStart.Call(obj)
    return int32(ret)
}

func ComboBox_SetSelStart(obj uintptr, value int32) {
   comboBox_SetSelStart.Call(obj, uintptr(value))
}

func ComboBox_GetBrush(obj uintptr) uintptr {
    ret, _, _ := comboBox_GetBrush.Call(obj)
    return ret
}

func ComboBox_GetControlCount(obj uintptr) int32 {
    ret, _, _ := comboBox_GetControlCount.Call(obj)
    return int32(ret)
}

func ComboBox_GetHandle(obj uintptr) HWND {
    ret, _, _ := comboBox_GetHandle.Call(obj)
    return HWND(ret)
}

func ComboBox_GetAction(obj uintptr) uintptr {
    ret, _, _ := comboBox_GetAction.Call(obj)
    return ret
}

func ComboBox_SetAction(obj uintptr, value uintptr) {
   comboBox_SetAction.Call(obj, value)
}

func ComboBox_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    comboBox_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ComboBox_SetBoundsRect(obj uintptr, value TRect) {
   comboBox_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ComboBox_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := comboBox_GetClientHeight.Call(obj)
    return int32(ret)
}

func ComboBox_SetClientHeight(obj uintptr, value int32) {
   comboBox_SetClientHeight.Call(obj, uintptr(value))
}

func ComboBox_GetClientRect(obj uintptr) TRect {
    var ret TRect
    comboBox_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ComboBox_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := comboBox_GetClientWidth.Call(obj)
    return int32(ret)
}

func ComboBox_SetClientWidth(obj uintptr, value int32) {
   comboBox_SetClientWidth.Call(obj, uintptr(value))
}

func ComboBox_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := comboBox_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func ComboBox_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := comboBox_GetExplicitTop.Call(obj)
    return int32(ret)
}

func ComboBox_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := comboBox_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func ComboBox_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := comboBox_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func ComboBox_GetParent(obj uintptr) uintptr {
    ret, _, _ := comboBox_GetParent.Call(obj)
    return ret
}

func ComboBox_SetParent(obj uintptr, value uintptr) {
   comboBox_SetParent.Call(obj, value)
}

func ComboBox_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := comboBox_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetAlignWithMargins(obj uintptr, value bool) {
   comboBox_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetLeft(obj uintptr) int32 {
    ret, _, _ := comboBox_GetLeft.Call(obj)
    return int32(ret)
}

func ComboBox_SetLeft(obj uintptr, value int32) {
   comboBox_SetLeft.Call(obj, uintptr(value))
}

func ComboBox_GetTop(obj uintptr) int32 {
    ret, _, _ := comboBox_GetTop.Call(obj)
    return int32(ret)
}

func ComboBox_SetTop(obj uintptr, value int32) {
   comboBox_SetTop.Call(obj, uintptr(value))
}

func ComboBox_GetWidth(obj uintptr) int32 {
    ret, _, _ := comboBox_GetWidth.Call(obj)
    return int32(ret)
}

func ComboBox_SetWidth(obj uintptr, value int32) {
   comboBox_SetWidth.Call(obj, uintptr(value))
}

func ComboBox_GetHeight(obj uintptr) int32 {
    ret, _, _ := comboBox_GetHeight.Call(obj)
    return int32(ret)
}

func ComboBox_SetHeight(obj uintptr, value int32) {
   comboBox_SetHeight.Call(obj, uintptr(value))
}

func ComboBox_GetCursor(obj uintptr) TCursor {
    ret, _, _ := comboBox_GetCursor.Call(obj)
    return TCursor(ret)
}

func ComboBox_SetCursor(obj uintptr, value TCursor) {
   comboBox_SetCursor.Call(obj, uintptr(value))
}

func ComboBox_GetHint(obj uintptr) string {
    ret, _, _ := comboBox_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func ComboBox_SetHint(obj uintptr, value string) {
   comboBox_SetHint.Call(obj, GoStrToDStr(value))
}

func ComboBox_GetMargins(obj uintptr) uintptr {
    ret, _, _ := comboBox_GetMargins.Call(obj)
    return ret
}

func ComboBox_SetMargins(obj uintptr, value uintptr) {
   comboBox_SetMargins.Call(obj, value)
}

func ComboBox_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := comboBox_GetComponentCount.Call(obj)
    return int32(ret)
}

func ComboBox_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := comboBox_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ComboBox_SetComponentIndex(obj uintptr, value int32) {
   comboBox_SetComponentIndex.Call(obj, uintptr(value))
}

func ComboBox_GetOwner(obj uintptr) uintptr {
    ret, _, _ := comboBox_GetOwner.Call(obj)
    return ret
}

func ComboBox_GetName(obj uintptr) string {
    ret, _, _ := comboBox_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ComboBox_SetName(obj uintptr, value string) {
   comboBox_SetName.Call(obj, GoStrToDStr(value))
}

func ComboBox_GetTag(obj uintptr) int {
    ret, _, _ := comboBox_GetTag.Call(obj)
    return int(ret)
}

func ComboBox_SetTag(obj uintptr, value int) {
   comboBox_SetTag.Call(obj, uintptr(value))
}

func ComboBox_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := comboBox_GetControls.Call(obj, uintptr(Index))
    return ret
}

func ComboBox_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := comboBox_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

