
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func Edit_Create(obj uintptr) uintptr {
    ret, _, _ := edit_Create.Call(obj)
    return ret
}

func Edit_Free(obj uintptr) {
    edit_Free.Call(obj)
}

func Edit_Clear(obj uintptr)  {
    edit_Clear.Call(obj)
}

func Edit_ClearSelection(obj uintptr)  {
    edit_ClearSelection.Call(obj)
}

func Edit_CopyToClipboard(obj uintptr)  {
    edit_CopyToClipboard.Call(obj)
}

func Edit_CutToClipboard(obj uintptr)  {
    edit_CutToClipboard.Call(obj)
}

func Edit_PasteFromClipboard(obj uintptr)  {
    edit_PasteFromClipboard.Call(obj)
}

func Edit_SelectAll(obj uintptr)  {
    edit_SelectAll.Call(obj)
}

func Edit_CanFocus(obj uintptr) bool {
    ret, _, _ := edit_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_FlipChildren(obj uintptr, AllLevels bool)  {
    edit_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func Edit_Focused(obj uintptr) bool {
    ret, _, _ := edit_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_HandleAllocated(obj uintptr) bool {
    ret, _, _ := edit_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_Invalidate(obj uintptr)  {
    edit_Invalidate.Call(obj)
}

func Edit_Realign(obj uintptr)  {
    edit_Realign.Call(obj)
}

func Edit_Repaint(obj uintptr)  {
    edit_Repaint.Call(obj)
}

func Edit_ScaleBy(obj uintptr, M int32, D int32)  {
    edit_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func Edit_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    edit_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func Edit_SetFocus(obj uintptr)  {
    edit_SetFocus.Call(obj)
}

func Edit_Update(obj uintptr)  {
    edit_Update.Call(obj)
}

func Edit_BringToFront(obj uintptr)  {
    edit_BringToFront.Call(obj)
}

func Edit_HasParent(obj uintptr) bool {
    ret, _, _ := edit_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_Hide(obj uintptr)  {
    edit_Hide.Call(obj)
}

func Edit_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := edit_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func Edit_Refresh(obj uintptr)  {
    edit_Refresh.Call(obj)
}

func Edit_SendToBack(obj uintptr)  {
    edit_SendToBack.Call(obj)
}

func Edit_Show(obj uintptr)  {
    edit_Show.Call(obj)
}

func Edit_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := edit_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Edit_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := edit_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Edit_GetNamePath(obj uintptr) string {
    ret, _, _ := edit_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Edit_Assign(obj uintptr, Source uintptr)  {
    edit_Assign.Call(obj, Source )
}

func Edit_ClassName(obj uintptr) string {
    ret, _, _ := edit_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Edit_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := edit_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Edit_GetHashCode(obj uintptr) int32 {
    ret, _, _ := edit_GetHashCode.Call(obj)
    return int32(ret)
}

func Edit_ToString(obj uintptr) string {
    ret, _, _ := edit_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Edit_GetAlign(obj uintptr) TAlign {
    ret, _, _ := edit_GetAlign.Call(obj)
    return TAlign(ret)
}

func Edit_SetAlign(obj uintptr, value TAlign) {
   edit_SetAlign.Call(obj, uintptr(value))
}

func Edit_GetAlignment(obj uintptr) TAlignment {
    ret, _, _ := edit_GetAlignment.Call(obj)
    return TAlignment(ret)
}

func Edit_SetAlignment(obj uintptr, value TAlignment) {
   edit_SetAlignment.Call(obj, uintptr(value))
}

func Edit_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := edit_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func Edit_SetAnchors(obj uintptr, value TAnchors) {
   edit_SetAnchors.Call(obj, uintptr(value))
}

func Edit_GetAutoSelect(obj uintptr) bool {
    ret, _, _ := edit_GetAutoSelect.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetAutoSelect(obj uintptr, value bool) {
   edit_SetAutoSelect.Call(obj, GoBoolToDBool(value))
}

func Edit_GetAutoSize(obj uintptr) bool {
    ret, _, _ := edit_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetAutoSize(obj uintptr, value bool) {
   edit_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func Edit_GetBorderStyle(obj uintptr) TBorderStyle {
    ret, _, _ := edit_GetBorderStyle.Call(obj)
    return TBorderStyle(ret)
}

func Edit_SetBorderStyle(obj uintptr, value TBorderStyle) {
   edit_SetBorderStyle.Call(obj, uintptr(value))
}

func Edit_GetColor(obj uintptr) TColor {
    ret, _, _ := edit_GetColor.Call(obj)
    return TColor(ret)
}

func Edit_SetColor(obj uintptr, value TColor) {
   edit_SetColor.Call(obj, uintptr(value))
}

func Edit_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := edit_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetDoubleBuffered(obj uintptr, value bool) {
   edit_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func Edit_GetEnabled(obj uintptr) bool {
    ret, _, _ := edit_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetEnabled(obj uintptr, value bool) {
   edit_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Edit_GetFont(obj uintptr) uintptr {
    ret, _, _ := edit_GetFont.Call(obj)
    return ret
}

func Edit_SetFont(obj uintptr, value uintptr) {
   edit_SetFont.Call(obj, value)
}

func Edit_GetHideSelection(obj uintptr) bool {
    ret, _, _ := edit_GetHideSelection.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetHideSelection(obj uintptr, value bool) {
   edit_SetHideSelection.Call(obj, GoBoolToDBool(value))
}

func Edit_GetMaxLength(obj uintptr) int32 {
    ret, _, _ := edit_GetMaxLength.Call(obj)
    return int32(ret)
}

func Edit_SetMaxLength(obj uintptr, value int32) {
   edit_SetMaxLength.Call(obj, uintptr(value))
}

func Edit_GetNumbersOnly(obj uintptr) bool {
    ret, _, _ := edit_GetNumbersOnly.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetNumbersOnly(obj uintptr, value bool) {
   edit_SetNumbersOnly.Call(obj, GoBoolToDBool(value))
}

func Edit_GetParentColor(obj uintptr) bool {
    ret, _, _ := edit_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetParentColor(obj uintptr, value bool) {
   edit_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func Edit_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := edit_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetParentCtl3D(obj uintptr, value bool) {
   edit_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func Edit_GetParentFont(obj uintptr) bool {
    ret, _, _ := edit_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetParentFont(obj uintptr, value bool) {
   edit_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func Edit_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := edit_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetParentShowHint(obj uintptr, value bool) {
   edit_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func Edit_GetPasswordChar(obj uintptr) uint16 {
    ret, _, _ := edit_GetPasswordChar.Call(obj)
    return uint16(ret)
}

func Edit_SetPasswordChar(obj uintptr, value uint16) {
   edit_SetPasswordChar.Call(obj, uintptr(value))
}

func Edit_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := edit_GetPopupMenu.Call(obj)
    return ret
}

func Edit_SetPopupMenu(obj uintptr, value uintptr) {
   edit_SetPopupMenu.Call(obj, value)
}

func Edit_GetReadOnly(obj uintptr) bool {
    ret, _, _ := edit_GetReadOnly.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetReadOnly(obj uintptr, value bool) {
   edit_SetReadOnly.Call(obj, GoBoolToDBool(value))
}

func Edit_GetShowHint(obj uintptr) bool {
    ret, _, _ := edit_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetShowHint(obj uintptr, value bool) {
   edit_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Edit_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := edit_GetTabOrder.Call(obj)
    return int16(ret)
}

func Edit_SetTabOrder(obj uintptr, value int16) {
   edit_SetTabOrder.Call(obj, uintptr(value))
}

func Edit_GetTabStop(obj uintptr) bool {
    ret, _, _ := edit_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetTabStop(obj uintptr, value bool) {
   edit_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func Edit_GetText(obj uintptr) string {
    ret, _, _ := edit_GetText.Call(obj)
    return DStrToGoStr(ret)
}

func Edit_SetText(obj uintptr, value string) {
   edit_SetText.Call(obj, GoStrToDStr(value))
}

func Edit_GetTextHint(obj uintptr) string {
    ret, _, _ := edit_GetTextHint.Call(obj)
    return DStrToGoStr(ret)
}

func Edit_SetTextHint(obj uintptr, value string) {
   edit_SetTextHint.Call(obj, GoStrToDStr(value))
}

func Edit_GetVisible(obj uintptr) bool {
    ret, _, _ := edit_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetVisible(obj uintptr, value bool) {
   edit_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Edit_SetOnChange(obj uintptr, fn interface{}) {
    edit_SetOnChange.Call(obj, addEventToMap(fn))
}

func Edit_SetOnClick(obj uintptr, fn interface{}) {
    edit_SetOnClick.Call(obj, addEventToMap(fn))
}

func Edit_SetOnDblClick(obj uintptr, fn interface{}) {
    edit_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func Edit_SetOnEnter(obj uintptr, fn interface{}) {
    edit_SetOnEnter.Call(obj, addEventToMap(fn))
}

func Edit_SetOnExit(obj uintptr, fn interface{}) {
    edit_SetOnExit.Call(obj, addEventToMap(fn))
}

func Edit_SetOnKeyDown(obj uintptr, fn interface{}) {
    edit_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func Edit_SetOnKeyPress(obj uintptr, fn interface{}) {
    edit_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func Edit_SetOnKeyUp(obj uintptr, fn interface{}) {
    edit_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func Edit_SetOnMouseDown(obj uintptr, fn interface{}) {
    edit_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func Edit_SetOnMouseEnter(obj uintptr, fn interface{}) {
    edit_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func Edit_SetOnMouseLeave(obj uintptr, fn interface{}) {
    edit_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func Edit_SetOnMouseMove(obj uintptr, fn interface{}) {
    edit_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func Edit_SetOnMouseUp(obj uintptr, fn interface{}) {
    edit_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func Edit_GetModified(obj uintptr) bool {
    ret, _, _ := edit_GetModified.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetModified(obj uintptr, value bool) {
   edit_SetModified.Call(obj, GoBoolToDBool(value))
}

func Edit_GetSelLength(obj uintptr) int32 {
    ret, _, _ := edit_GetSelLength.Call(obj)
    return int32(ret)
}

func Edit_SetSelLength(obj uintptr, value int32) {
   edit_SetSelLength.Call(obj, uintptr(value))
}

func Edit_GetSelStart(obj uintptr) int32 {
    ret, _, _ := edit_GetSelStart.Call(obj)
    return int32(ret)
}

func Edit_SetSelStart(obj uintptr, value int32) {
   edit_SetSelStart.Call(obj, uintptr(value))
}

func Edit_GetSelText(obj uintptr) string {
    ret, _, _ := edit_GetSelText.Call(obj)
    return DStrToGoStr(ret)
}

func Edit_SetSelText(obj uintptr, value string) {
   edit_SetSelText.Call(obj, GoStrToDStr(value))
}

func Edit_GetBrush(obj uintptr) uintptr {
    ret, _, _ := edit_GetBrush.Call(obj)
    return ret
}

func Edit_GetControlCount(obj uintptr) int32 {
    ret, _, _ := edit_GetControlCount.Call(obj)
    return int32(ret)
}

func Edit_GetHandle(obj uintptr) HWND {
    ret, _, _ := edit_GetHandle.Call(obj)
    return HWND(ret)
}

func Edit_GetAction(obj uintptr) uintptr {
    ret, _, _ := edit_GetAction.Call(obj)
    return ret
}

func Edit_SetAction(obj uintptr, value uintptr) {
   edit_SetAction.Call(obj, value)
}

func Edit_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    edit_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Edit_SetBoundsRect(obj uintptr, value TRect) {
   edit_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Edit_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := edit_GetClientHeight.Call(obj)
    return int32(ret)
}

func Edit_SetClientHeight(obj uintptr, value int32) {
   edit_SetClientHeight.Call(obj, uintptr(value))
}

func Edit_GetClientRect(obj uintptr) TRect {
    var ret TRect
    edit_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Edit_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := edit_GetClientWidth.Call(obj)
    return int32(ret)
}

func Edit_SetClientWidth(obj uintptr, value int32) {
   edit_SetClientWidth.Call(obj, uintptr(value))
}

func Edit_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := edit_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Edit_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := edit_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Edit_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := edit_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Edit_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := edit_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Edit_GetParent(obj uintptr) uintptr {
    ret, _, _ := edit_GetParent.Call(obj)
    return ret
}

func Edit_SetParent(obj uintptr, value uintptr) {
   edit_SetParent.Call(obj, value)
}

func Edit_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := edit_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetAlignWithMargins(obj uintptr, value bool) {
   edit_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func Edit_GetLeft(obj uintptr) int32 {
    ret, _, _ := edit_GetLeft.Call(obj)
    return int32(ret)
}

func Edit_SetLeft(obj uintptr, value int32) {
   edit_SetLeft.Call(obj, uintptr(value))
}

func Edit_GetTop(obj uintptr) int32 {
    ret, _, _ := edit_GetTop.Call(obj)
    return int32(ret)
}

func Edit_SetTop(obj uintptr, value int32) {
   edit_SetTop.Call(obj, uintptr(value))
}

func Edit_GetWidth(obj uintptr) int32 {
    ret, _, _ := edit_GetWidth.Call(obj)
    return int32(ret)
}

func Edit_SetWidth(obj uintptr, value int32) {
   edit_SetWidth.Call(obj, uintptr(value))
}

func Edit_GetHeight(obj uintptr) int32 {
    ret, _, _ := edit_GetHeight.Call(obj)
    return int32(ret)
}

func Edit_SetHeight(obj uintptr, value int32) {
   edit_SetHeight.Call(obj, uintptr(value))
}

func Edit_GetCursor(obj uintptr) TCursor {
    ret, _, _ := edit_GetCursor.Call(obj)
    return TCursor(ret)
}

func Edit_SetCursor(obj uintptr, value TCursor) {
   edit_SetCursor.Call(obj, uintptr(value))
}

func Edit_GetHint(obj uintptr) string {
    ret, _, _ := edit_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Edit_SetHint(obj uintptr, value string) {
   edit_SetHint.Call(obj, GoStrToDStr(value))
}

func Edit_GetMargins(obj uintptr) uintptr {
    ret, _, _ := edit_GetMargins.Call(obj)
    return ret
}

func Edit_SetMargins(obj uintptr, value uintptr) {
   edit_SetMargins.Call(obj, value)
}

func Edit_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := edit_GetComponentCount.Call(obj)
    return int32(ret)
}

func Edit_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := edit_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Edit_SetComponentIndex(obj uintptr, value int32) {
   edit_SetComponentIndex.Call(obj, uintptr(value))
}

func Edit_GetOwner(obj uintptr) uintptr {
    ret, _, _ := edit_GetOwner.Call(obj)
    return ret
}

func Edit_GetName(obj uintptr) string {
    ret, _, _ := edit_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Edit_SetName(obj uintptr, value string) {
   edit_SetName.Call(obj, GoStrToDStr(value))
}

func Edit_GetTag(obj uintptr) int {
    ret, _, _ := edit_GetTag.Call(obj)
    return int(ret)
}

func Edit_SetTag(obj uintptr, value int) {
   edit_SetTag.Call(obj, uintptr(value))
}

func Edit_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := edit_GetControls.Call(obj, uintptr(Index))
    return ret
}

func Edit_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := edit_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

