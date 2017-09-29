
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func Memo_Create(obj uintptr) uintptr {
    ret, _, _ := memo_Create.Call(obj)
    return ret
}

func Memo_Free(obj uintptr) {
    memo_Free.Call(obj)
}

func Memo_Clear(obj uintptr)  {
    memo_Clear.Call(obj)
}

func Memo_ClearSelection(obj uintptr)  {
    memo_ClearSelection.Call(obj)
}

func Memo_CopyToClipboard(obj uintptr)  {
    memo_CopyToClipboard.Call(obj)
}

func Memo_CutToClipboard(obj uintptr)  {
    memo_CutToClipboard.Call(obj)
}

func Memo_PasteFromClipboard(obj uintptr)  {
    memo_PasteFromClipboard.Call(obj)
}

func Memo_SelectAll(obj uintptr)  {
    memo_SelectAll.Call(obj)
}

func Memo_CanFocus(obj uintptr) bool {
    ret, _, _ := memo_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_FlipChildren(obj uintptr, AllLevels bool)  {
    memo_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func Memo_Focused(obj uintptr) bool {
    ret, _, _ := memo_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_HandleAllocated(obj uintptr) bool {
    ret, _, _ := memo_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_Invalidate(obj uintptr)  {
    memo_Invalidate.Call(obj)
}

func Memo_Realign(obj uintptr)  {
    memo_Realign.Call(obj)
}

func Memo_Repaint(obj uintptr)  {
    memo_Repaint.Call(obj)
}

func Memo_ScaleBy(obj uintptr, M int32, D int32)  {
    memo_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func Memo_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    memo_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func Memo_SetFocus(obj uintptr)  {
    memo_SetFocus.Call(obj)
}

func Memo_Update(obj uintptr)  {
    memo_Update.Call(obj)
}

func Memo_BringToFront(obj uintptr)  {
    memo_BringToFront.Call(obj)
}

func Memo_HasParent(obj uintptr) bool {
    ret, _, _ := memo_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_Hide(obj uintptr)  {
    memo_Hide.Call(obj)
}

func Memo_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := memo_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func Memo_Refresh(obj uintptr)  {
    memo_Refresh.Call(obj)
}

func Memo_SendToBack(obj uintptr)  {
    memo_SendToBack.Call(obj)
}

func Memo_Show(obj uintptr)  {
    memo_Show.Call(obj)
}

func Memo_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := memo_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Memo_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := memo_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Memo_GetNamePath(obj uintptr) string {
    ret, _, _ := memo_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Memo_Assign(obj uintptr, Source uintptr)  {
    memo_Assign.Call(obj, Source )
}

func Memo_ClassName(obj uintptr) string {
    ret, _, _ := memo_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Memo_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := memo_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Memo_GetHashCode(obj uintptr) int32 {
    ret, _, _ := memo_GetHashCode.Call(obj)
    return int32(ret)
}

func Memo_ToString(obj uintptr) string {
    ret, _, _ := memo_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Memo_GetAlign(obj uintptr) TAlign {
    ret, _, _ := memo_GetAlign.Call(obj)
    return TAlign(ret)
}

func Memo_SetAlign(obj uintptr, value TAlign) {
   memo_SetAlign.Call(obj, uintptr(value))
}

func Memo_GetAlignment(obj uintptr) TAlignment {
    ret, _, _ := memo_GetAlignment.Call(obj)
    return TAlignment(ret)
}

func Memo_SetAlignment(obj uintptr, value TAlignment) {
   memo_SetAlignment.Call(obj, uintptr(value))
}

func Memo_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := memo_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func Memo_SetAnchors(obj uintptr, value TAnchors) {
   memo_SetAnchors.Call(obj, uintptr(value))
}

func Memo_GetBorderStyle(obj uintptr) TBorderStyle {
    ret, _, _ := memo_GetBorderStyle.Call(obj)
    return TBorderStyle(ret)
}

func Memo_SetBorderStyle(obj uintptr, value TBorderStyle) {
   memo_SetBorderStyle.Call(obj, uintptr(value))
}

func Memo_GetColor(obj uintptr) TColor {
    ret, _, _ := memo_GetColor.Call(obj)
    return TColor(ret)
}

func Memo_SetColor(obj uintptr, value TColor) {
   memo_SetColor.Call(obj, uintptr(value))
}

func Memo_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := memo_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetDoubleBuffered(obj uintptr, value bool) {
   memo_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func Memo_GetEnabled(obj uintptr) bool {
    ret, _, _ := memo_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetEnabled(obj uintptr, value bool) {
   memo_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Memo_GetFont(obj uintptr) uintptr {
    ret, _, _ := memo_GetFont.Call(obj)
    return ret
}

func Memo_SetFont(obj uintptr, value uintptr) {
   memo_SetFont.Call(obj, value)
}

func Memo_GetHideSelection(obj uintptr) bool {
    ret, _, _ := memo_GetHideSelection.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetHideSelection(obj uintptr, value bool) {
   memo_SetHideSelection.Call(obj, GoBoolToDBool(value))
}

func Memo_GetLines(obj uintptr) uintptr {
    ret, _, _ := memo_GetLines.Call(obj)
    return ret
}

func Memo_SetLines(obj uintptr, value uintptr) {
   memo_SetLines.Call(obj, value)
}

func Memo_GetMaxLength(obj uintptr) int32 {
    ret, _, _ := memo_GetMaxLength.Call(obj)
    return int32(ret)
}

func Memo_SetMaxLength(obj uintptr, value int32) {
   memo_SetMaxLength.Call(obj, uintptr(value))
}

func Memo_GetParentColor(obj uintptr) bool {
    ret, _, _ := memo_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetParentColor(obj uintptr, value bool) {
   memo_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func Memo_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := memo_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetParentCtl3D(obj uintptr, value bool) {
   memo_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func Memo_GetParentFont(obj uintptr) bool {
    ret, _, _ := memo_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetParentFont(obj uintptr, value bool) {
   memo_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func Memo_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := memo_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetParentShowHint(obj uintptr, value bool) {
   memo_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func Memo_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := memo_GetPopupMenu.Call(obj)
    return ret
}

func Memo_SetPopupMenu(obj uintptr, value uintptr) {
   memo_SetPopupMenu.Call(obj, value)
}

func Memo_GetReadOnly(obj uintptr) bool {
    ret, _, _ := memo_GetReadOnly.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetReadOnly(obj uintptr, value bool) {
   memo_SetReadOnly.Call(obj, GoBoolToDBool(value))
}

func Memo_GetScrollBars(obj uintptr) TScrollStyle {
    ret, _, _ := memo_GetScrollBars.Call(obj)
    return TScrollStyle(ret)
}

func Memo_SetScrollBars(obj uintptr, value TScrollStyle) {
   memo_SetScrollBars.Call(obj, uintptr(value))
}

func Memo_GetShowHint(obj uintptr) bool {
    ret, _, _ := memo_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetShowHint(obj uintptr, value bool) {
   memo_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Memo_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := memo_GetTabOrder.Call(obj)
    return int16(ret)
}

func Memo_SetTabOrder(obj uintptr, value int16) {
   memo_SetTabOrder.Call(obj, uintptr(value))
}

func Memo_GetTabStop(obj uintptr) bool {
    ret, _, _ := memo_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetTabStop(obj uintptr, value bool) {
   memo_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func Memo_GetVisible(obj uintptr) bool {
    ret, _, _ := memo_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetVisible(obj uintptr, value bool) {
   memo_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Memo_GetWantReturns(obj uintptr) bool {
    ret, _, _ := memo_GetWantReturns.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetWantReturns(obj uintptr, value bool) {
   memo_SetWantReturns.Call(obj, GoBoolToDBool(value))
}

func Memo_GetWantTabs(obj uintptr) bool {
    ret, _, _ := memo_GetWantTabs.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetWantTabs(obj uintptr, value bool) {
   memo_SetWantTabs.Call(obj, GoBoolToDBool(value))
}

func Memo_GetWordWrap(obj uintptr) bool {
    ret, _, _ := memo_GetWordWrap.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetWordWrap(obj uintptr, value bool) {
   memo_SetWordWrap.Call(obj, GoBoolToDBool(value))
}

func Memo_SetOnChange(obj uintptr, fn interface{}) {
    memo_SetOnChange.Call(obj, addEventToMap(fn))
}

func Memo_SetOnClick(obj uintptr, fn interface{}) {
    memo_SetOnClick.Call(obj, addEventToMap(fn))
}

func Memo_SetOnDblClick(obj uintptr, fn interface{}) {
    memo_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func Memo_SetOnEnter(obj uintptr, fn interface{}) {
    memo_SetOnEnter.Call(obj, addEventToMap(fn))
}

func Memo_SetOnExit(obj uintptr, fn interface{}) {
    memo_SetOnExit.Call(obj, addEventToMap(fn))
}

func Memo_SetOnKeyDown(obj uintptr, fn interface{}) {
    memo_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func Memo_SetOnKeyPress(obj uintptr, fn interface{}) {
    memo_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func Memo_SetOnKeyUp(obj uintptr, fn interface{}) {
    memo_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func Memo_SetOnMouseDown(obj uintptr, fn interface{}) {
    memo_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func Memo_SetOnMouseEnter(obj uintptr, fn interface{}) {
    memo_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func Memo_SetOnMouseLeave(obj uintptr, fn interface{}) {
    memo_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func Memo_SetOnMouseMove(obj uintptr, fn interface{}) {
    memo_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func Memo_SetOnMouseUp(obj uintptr, fn interface{}) {
    memo_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func Memo_GetModified(obj uintptr) bool {
    ret, _, _ := memo_GetModified.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetModified(obj uintptr, value bool) {
   memo_SetModified.Call(obj, GoBoolToDBool(value))
}

func Memo_GetSelLength(obj uintptr) int32 {
    ret, _, _ := memo_GetSelLength.Call(obj)
    return int32(ret)
}

func Memo_SetSelLength(obj uintptr, value int32) {
   memo_SetSelLength.Call(obj, uintptr(value))
}

func Memo_GetSelStart(obj uintptr) int32 {
    ret, _, _ := memo_GetSelStart.Call(obj)
    return int32(ret)
}

func Memo_SetSelStart(obj uintptr, value int32) {
   memo_SetSelStart.Call(obj, uintptr(value))
}

func Memo_GetSelText(obj uintptr) string {
    ret, _, _ := memo_GetSelText.Call(obj)
    return DStrToGoStr(ret)
}

func Memo_SetSelText(obj uintptr, value string) {
   memo_SetSelText.Call(obj, GoStrToDStr(value))
}

func Memo_GetText(obj uintptr) string {
    ret, _, _ := memo_GetText.Call(obj)
    return DStrToGoStr(ret)
}

func Memo_SetText(obj uintptr, value string) {
   memo_SetText.Call(obj, GoStrToDStr(value))
}

func Memo_GetTextHint(obj uintptr) string {
    ret, _, _ := memo_GetTextHint.Call(obj)
    return DStrToGoStr(ret)
}

func Memo_SetTextHint(obj uintptr, value string) {
   memo_SetTextHint.Call(obj, GoStrToDStr(value))
}

func Memo_GetBrush(obj uintptr) uintptr {
    ret, _, _ := memo_GetBrush.Call(obj)
    return ret
}

func Memo_GetControlCount(obj uintptr) int32 {
    ret, _, _ := memo_GetControlCount.Call(obj)
    return int32(ret)
}

func Memo_GetHandle(obj uintptr) HWND {
    ret, _, _ := memo_GetHandle.Call(obj)
    return HWND(ret)
}

func Memo_GetAction(obj uintptr) uintptr {
    ret, _, _ := memo_GetAction.Call(obj)
    return ret
}

func Memo_SetAction(obj uintptr, value uintptr) {
   memo_SetAction.Call(obj, value)
}

func Memo_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    memo_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Memo_SetBoundsRect(obj uintptr, value TRect) {
   memo_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Memo_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := memo_GetClientHeight.Call(obj)
    return int32(ret)
}

func Memo_SetClientHeight(obj uintptr, value int32) {
   memo_SetClientHeight.Call(obj, uintptr(value))
}

func Memo_GetClientRect(obj uintptr) TRect {
    var ret TRect
    memo_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Memo_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := memo_GetClientWidth.Call(obj)
    return int32(ret)
}

func Memo_SetClientWidth(obj uintptr, value int32) {
   memo_SetClientWidth.Call(obj, uintptr(value))
}

func Memo_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := memo_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Memo_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := memo_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Memo_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := memo_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Memo_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := memo_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Memo_GetParent(obj uintptr) uintptr {
    ret, _, _ := memo_GetParent.Call(obj)
    return ret
}

func Memo_SetParent(obj uintptr, value uintptr) {
   memo_SetParent.Call(obj, value)
}

func Memo_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := memo_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetAlignWithMargins(obj uintptr, value bool) {
   memo_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func Memo_GetLeft(obj uintptr) int32 {
    ret, _, _ := memo_GetLeft.Call(obj)
    return int32(ret)
}

func Memo_SetLeft(obj uintptr, value int32) {
   memo_SetLeft.Call(obj, uintptr(value))
}

func Memo_GetTop(obj uintptr) int32 {
    ret, _, _ := memo_GetTop.Call(obj)
    return int32(ret)
}

func Memo_SetTop(obj uintptr, value int32) {
   memo_SetTop.Call(obj, uintptr(value))
}

func Memo_GetWidth(obj uintptr) int32 {
    ret, _, _ := memo_GetWidth.Call(obj)
    return int32(ret)
}

func Memo_SetWidth(obj uintptr, value int32) {
   memo_SetWidth.Call(obj, uintptr(value))
}

func Memo_GetHeight(obj uintptr) int32 {
    ret, _, _ := memo_GetHeight.Call(obj)
    return int32(ret)
}

func Memo_SetHeight(obj uintptr, value int32) {
   memo_SetHeight.Call(obj, uintptr(value))
}

func Memo_GetCursor(obj uintptr) TCursor {
    ret, _, _ := memo_GetCursor.Call(obj)
    return TCursor(ret)
}

func Memo_SetCursor(obj uintptr, value TCursor) {
   memo_SetCursor.Call(obj, uintptr(value))
}

func Memo_GetHint(obj uintptr) string {
    ret, _, _ := memo_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Memo_SetHint(obj uintptr, value string) {
   memo_SetHint.Call(obj, GoStrToDStr(value))
}

func Memo_GetMargins(obj uintptr) uintptr {
    ret, _, _ := memo_GetMargins.Call(obj)
    return ret
}

func Memo_SetMargins(obj uintptr, value uintptr) {
   memo_SetMargins.Call(obj, value)
}

func Memo_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := memo_GetComponentCount.Call(obj)
    return int32(ret)
}

func Memo_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := memo_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Memo_SetComponentIndex(obj uintptr, value int32) {
   memo_SetComponentIndex.Call(obj, uintptr(value))
}

func Memo_GetOwner(obj uintptr) uintptr {
    ret, _, _ := memo_GetOwner.Call(obj)
    return ret
}

func Memo_GetName(obj uintptr) string {
    ret, _, _ := memo_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Memo_SetName(obj uintptr, value string) {
   memo_SetName.Call(obj, GoStrToDStr(value))
}

func Memo_GetTag(obj uintptr) int {
    ret, _, _ := memo_GetTag.Call(obj)
    return int(ret)
}

func Memo_SetTag(obj uintptr, value int) {
   memo_SetTag.Call(obj, uintptr(value))
}

func Memo_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := memo_GetControls.Call(obj, uintptr(Index))
    return ret
}

func Memo_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := memo_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

