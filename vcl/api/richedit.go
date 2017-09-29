
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func RichEdit_Create(obj uintptr) uintptr {
    ret, _, _ := richEdit_Create.Call(obj)
    return ret
}

func RichEdit_Free(obj uintptr) {
    richEdit_Free.Call(obj)
}

func RichEdit_Clear(obj uintptr)  {
    richEdit_Clear.Call(obj)
}

func RichEdit_ClearSelection(obj uintptr)  {
    richEdit_ClearSelection.Call(obj)
}

func RichEdit_CopyToClipboard(obj uintptr)  {
    richEdit_CopyToClipboard.Call(obj)
}

func RichEdit_CutToClipboard(obj uintptr)  {
    richEdit_CutToClipboard.Call(obj)
}

func RichEdit_PasteFromClipboard(obj uintptr)  {
    richEdit_PasteFromClipboard.Call(obj)
}

func RichEdit_SelectAll(obj uintptr)  {
    richEdit_SelectAll.Call(obj)
}

func RichEdit_CanFocus(obj uintptr) bool {
    ret, _, _ := richEdit_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_FlipChildren(obj uintptr, AllLevels bool)  {
    richEdit_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func RichEdit_Focused(obj uintptr) bool {
    ret, _, _ := richEdit_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_HandleAllocated(obj uintptr) bool {
    ret, _, _ := richEdit_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_Invalidate(obj uintptr)  {
    richEdit_Invalidate.Call(obj)
}

func RichEdit_Realign(obj uintptr)  {
    richEdit_Realign.Call(obj)
}

func RichEdit_Repaint(obj uintptr)  {
    richEdit_Repaint.Call(obj)
}

func RichEdit_ScaleBy(obj uintptr, M int32, D int32)  {
    richEdit_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func RichEdit_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    richEdit_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func RichEdit_SetFocus(obj uintptr)  {
    richEdit_SetFocus.Call(obj)
}

func RichEdit_Update(obj uintptr)  {
    richEdit_Update.Call(obj)
}

func RichEdit_BringToFront(obj uintptr)  {
    richEdit_BringToFront.Call(obj)
}

func RichEdit_HasParent(obj uintptr) bool {
    ret, _, _ := richEdit_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_Hide(obj uintptr)  {
    richEdit_Hide.Call(obj)
}

func RichEdit_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := richEdit_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func RichEdit_Refresh(obj uintptr)  {
    richEdit_Refresh.Call(obj)
}

func RichEdit_SendToBack(obj uintptr)  {
    richEdit_SendToBack.Call(obj)
}

func RichEdit_Show(obj uintptr)  {
    richEdit_Show.Call(obj)
}

func RichEdit_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := richEdit_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func RichEdit_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := richEdit_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func RichEdit_GetNamePath(obj uintptr) string {
    ret, _, _ := richEdit_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func RichEdit_Assign(obj uintptr, Source uintptr)  {
    richEdit_Assign.Call(obj, Source )
}

func RichEdit_ClassName(obj uintptr) string {
    ret, _, _ := richEdit_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func RichEdit_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := richEdit_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func RichEdit_GetHashCode(obj uintptr) int32 {
    ret, _, _ := richEdit_GetHashCode.Call(obj)
    return int32(ret)
}

func RichEdit_ToString(obj uintptr) string {
    ret, _, _ := richEdit_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func RichEdit_GetAlign(obj uintptr) TAlign {
    ret, _, _ := richEdit_GetAlign.Call(obj)
    return TAlign(ret)
}

func RichEdit_SetAlign(obj uintptr, value TAlign) {
   richEdit_SetAlign.Call(obj, uintptr(value))
}

func RichEdit_GetAlignment(obj uintptr) TAlignment {
    ret, _, _ := richEdit_GetAlignment.Call(obj)
    return TAlignment(ret)
}

func RichEdit_SetAlignment(obj uintptr, value TAlignment) {
   richEdit_SetAlignment.Call(obj, uintptr(value))
}

func RichEdit_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := richEdit_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func RichEdit_SetAnchors(obj uintptr, value TAnchors) {
   richEdit_SetAnchors.Call(obj, uintptr(value))
}

func RichEdit_GetBorderStyle(obj uintptr) TBorderStyle {
    ret, _, _ := richEdit_GetBorderStyle.Call(obj)
    return TBorderStyle(ret)
}

func RichEdit_SetBorderStyle(obj uintptr, value TBorderStyle) {
   richEdit_SetBorderStyle.Call(obj, uintptr(value))
}

func RichEdit_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := richEdit_GetBorderWidth.Call(obj)
    return int32(ret)
}

func RichEdit_SetBorderWidth(obj uintptr, value int32) {
   richEdit_SetBorderWidth.Call(obj, uintptr(value))
}

func RichEdit_GetColor(obj uintptr) TColor {
    ret, _, _ := richEdit_GetColor.Call(obj)
    return TColor(ret)
}

func RichEdit_SetColor(obj uintptr, value TColor) {
   richEdit_SetColor.Call(obj, uintptr(value))
}

func RichEdit_GetEnabled(obj uintptr) bool {
    ret, _, _ := richEdit_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetEnabled(obj uintptr, value bool) {
   richEdit_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetFont(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetFont.Call(obj)
    return ret
}

func RichEdit_SetFont(obj uintptr, value uintptr) {
   richEdit_SetFont.Call(obj, value)
}

func RichEdit_GetHideSelection(obj uintptr) bool {
    ret, _, _ := richEdit_GetHideSelection.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetHideSelection(obj uintptr, value bool) {
   richEdit_SetHideSelection.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetLines(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetLines.Call(obj)
    return ret
}

func RichEdit_SetLines(obj uintptr, value uintptr) {
   richEdit_SetLines.Call(obj, value)
}

func RichEdit_GetMaxLength(obj uintptr) int32 {
    ret, _, _ := richEdit_GetMaxLength.Call(obj)
    return int32(ret)
}

func RichEdit_SetMaxLength(obj uintptr, value int32) {
   richEdit_SetMaxLength.Call(obj, uintptr(value))
}

func RichEdit_GetParentColor(obj uintptr) bool {
    ret, _, _ := richEdit_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetParentColor(obj uintptr, value bool) {
   richEdit_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := richEdit_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetParentCtl3D(obj uintptr, value bool) {
   richEdit_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetParentFont(obj uintptr) bool {
    ret, _, _ := richEdit_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetParentFont(obj uintptr, value bool) {
   richEdit_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := richEdit_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetParentShowHint(obj uintptr, value bool) {
   richEdit_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetPopupMenu.Call(obj)
    return ret
}

func RichEdit_SetPopupMenu(obj uintptr, value uintptr) {
   richEdit_SetPopupMenu.Call(obj, value)
}

func RichEdit_GetReadOnly(obj uintptr) bool {
    ret, _, _ := richEdit_GetReadOnly.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetReadOnly(obj uintptr, value bool) {
   richEdit_SetReadOnly.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetScrollBars(obj uintptr) TScrollStyle {
    ret, _, _ := richEdit_GetScrollBars.Call(obj)
    return TScrollStyle(ret)
}

func RichEdit_SetScrollBars(obj uintptr, value TScrollStyle) {
   richEdit_SetScrollBars.Call(obj, uintptr(value))
}

func RichEdit_GetShowHint(obj uintptr) bool {
    ret, _, _ := richEdit_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetShowHint(obj uintptr, value bool) {
   richEdit_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := richEdit_GetTabOrder.Call(obj)
    return int16(ret)
}

func RichEdit_SetTabOrder(obj uintptr, value int16) {
   richEdit_SetTabOrder.Call(obj, uintptr(value))
}

func RichEdit_GetTabStop(obj uintptr) bool {
    ret, _, _ := richEdit_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetTabStop(obj uintptr, value bool) {
   richEdit_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetVisible(obj uintptr) bool {
    ret, _, _ := richEdit_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetVisible(obj uintptr, value bool) {
   richEdit_SetVisible.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetWantTabs(obj uintptr) bool {
    ret, _, _ := richEdit_GetWantTabs.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetWantTabs(obj uintptr, value bool) {
   richEdit_SetWantTabs.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetWantReturns(obj uintptr) bool {
    ret, _, _ := richEdit_GetWantReturns.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetWantReturns(obj uintptr, value bool) {
   richEdit_SetWantReturns.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetWordWrap(obj uintptr) bool {
    ret, _, _ := richEdit_GetWordWrap.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetWordWrap(obj uintptr, value bool) {
   richEdit_SetWordWrap.Call(obj, GoBoolToDBool(value))
}

func RichEdit_SetOnChange(obj uintptr, fn interface{}) {
    richEdit_SetOnChange.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnClick(obj uintptr, fn interface{}) {
    richEdit_SetOnClick.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnDblClick(obj uintptr, fn interface{}) {
    richEdit_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnEnter(obj uintptr, fn interface{}) {
    richEdit_SetOnEnter.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnExit(obj uintptr, fn interface{}) {
    richEdit_SetOnExit.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnKeyDown(obj uintptr, fn interface{}) {
    richEdit_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnKeyPress(obj uintptr, fn interface{}) {
    richEdit_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnKeyUp(obj uintptr, fn interface{}) {
    richEdit_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnMouseDown(obj uintptr, fn interface{}) {
    richEdit_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnMouseEnter(obj uintptr, fn interface{}) {
    richEdit_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnMouseLeave(obj uintptr, fn interface{}) {
    richEdit_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnMouseMove(obj uintptr, fn interface{}) {
    richEdit_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnMouseUp(obj uintptr, fn interface{}) {
    richEdit_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnMouseWheel(obj uintptr, fn interface{}) {
    richEdit_SetOnMouseWheel.Call(obj, addEventToMap(fn))
}

func RichEdit_GetModified(obj uintptr) bool {
    ret, _, _ := richEdit_GetModified.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetModified(obj uintptr, value bool) {
   richEdit_SetModified.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetSelLength(obj uintptr) int32 {
    ret, _, _ := richEdit_GetSelLength.Call(obj)
    return int32(ret)
}

func RichEdit_SetSelLength(obj uintptr, value int32) {
   richEdit_SetSelLength.Call(obj, uintptr(value))
}

func RichEdit_GetSelStart(obj uintptr) int32 {
    ret, _, _ := richEdit_GetSelStart.Call(obj)
    return int32(ret)
}

func RichEdit_SetSelStart(obj uintptr, value int32) {
   richEdit_SetSelStart.Call(obj, uintptr(value))
}

func RichEdit_GetSelText(obj uintptr) string {
    ret, _, _ := richEdit_GetSelText.Call(obj)
    return DStrToGoStr(ret)
}

func RichEdit_SetSelText(obj uintptr, value string) {
   richEdit_SetSelText.Call(obj, GoStrToDStr(value))
}

func RichEdit_GetText(obj uintptr) string {
    ret, _, _ := richEdit_GetText.Call(obj)
    return DStrToGoStr(ret)
}

func RichEdit_SetText(obj uintptr, value string) {
   richEdit_SetText.Call(obj, GoStrToDStr(value))
}

func RichEdit_GetTextHint(obj uintptr) string {
    ret, _, _ := richEdit_GetTextHint.Call(obj)
    return DStrToGoStr(ret)
}

func RichEdit_SetTextHint(obj uintptr, value string) {
   richEdit_SetTextHint.Call(obj, GoStrToDStr(value))
}

func RichEdit_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := richEdit_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetDoubleBuffered(obj uintptr, value bool) {
   richEdit_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetBrush(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetBrush.Call(obj)
    return ret
}

func RichEdit_GetControlCount(obj uintptr) int32 {
    ret, _, _ := richEdit_GetControlCount.Call(obj)
    return int32(ret)
}

func RichEdit_GetHandle(obj uintptr) HWND {
    ret, _, _ := richEdit_GetHandle.Call(obj)
    return HWND(ret)
}

func RichEdit_GetAction(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetAction.Call(obj)
    return ret
}

func RichEdit_SetAction(obj uintptr, value uintptr) {
   richEdit_SetAction.Call(obj, value)
}

func RichEdit_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    richEdit_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func RichEdit_SetBoundsRect(obj uintptr, value TRect) {
   richEdit_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func RichEdit_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := richEdit_GetClientHeight.Call(obj)
    return int32(ret)
}

func RichEdit_SetClientHeight(obj uintptr, value int32) {
   richEdit_SetClientHeight.Call(obj, uintptr(value))
}

func RichEdit_GetClientRect(obj uintptr) TRect {
    var ret TRect
    richEdit_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func RichEdit_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := richEdit_GetClientWidth.Call(obj)
    return int32(ret)
}

func RichEdit_SetClientWidth(obj uintptr, value int32) {
   richEdit_SetClientWidth.Call(obj, uintptr(value))
}

func RichEdit_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := richEdit_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func RichEdit_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := richEdit_GetExplicitTop.Call(obj)
    return int32(ret)
}

func RichEdit_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := richEdit_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func RichEdit_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := richEdit_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func RichEdit_GetParent(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetParent.Call(obj)
    return ret
}

func RichEdit_SetParent(obj uintptr, value uintptr) {
   richEdit_SetParent.Call(obj, value)
}

func RichEdit_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := richEdit_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetAlignWithMargins(obj uintptr, value bool) {
   richEdit_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetLeft(obj uintptr) int32 {
    ret, _, _ := richEdit_GetLeft.Call(obj)
    return int32(ret)
}

func RichEdit_SetLeft(obj uintptr, value int32) {
   richEdit_SetLeft.Call(obj, uintptr(value))
}

func RichEdit_GetTop(obj uintptr) int32 {
    ret, _, _ := richEdit_GetTop.Call(obj)
    return int32(ret)
}

func RichEdit_SetTop(obj uintptr, value int32) {
   richEdit_SetTop.Call(obj, uintptr(value))
}

func RichEdit_GetWidth(obj uintptr) int32 {
    ret, _, _ := richEdit_GetWidth.Call(obj)
    return int32(ret)
}

func RichEdit_SetWidth(obj uintptr, value int32) {
   richEdit_SetWidth.Call(obj, uintptr(value))
}

func RichEdit_GetHeight(obj uintptr) int32 {
    ret, _, _ := richEdit_GetHeight.Call(obj)
    return int32(ret)
}

func RichEdit_SetHeight(obj uintptr, value int32) {
   richEdit_SetHeight.Call(obj, uintptr(value))
}

func RichEdit_GetCursor(obj uintptr) TCursor {
    ret, _, _ := richEdit_GetCursor.Call(obj)
    return TCursor(ret)
}

func RichEdit_SetCursor(obj uintptr, value TCursor) {
   richEdit_SetCursor.Call(obj, uintptr(value))
}

func RichEdit_GetHint(obj uintptr) string {
    ret, _, _ := richEdit_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func RichEdit_SetHint(obj uintptr, value string) {
   richEdit_SetHint.Call(obj, GoStrToDStr(value))
}

func RichEdit_GetMargins(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetMargins.Call(obj)
    return ret
}

func RichEdit_SetMargins(obj uintptr, value uintptr) {
   richEdit_SetMargins.Call(obj, value)
}

func RichEdit_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := richEdit_GetComponentCount.Call(obj)
    return int32(ret)
}

func RichEdit_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := richEdit_GetComponentIndex.Call(obj)
    return int32(ret)
}

func RichEdit_SetComponentIndex(obj uintptr, value int32) {
   richEdit_SetComponentIndex.Call(obj, uintptr(value))
}

func RichEdit_GetOwner(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetOwner.Call(obj)
    return ret
}

func RichEdit_GetName(obj uintptr) string {
    ret, _, _ := richEdit_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func RichEdit_SetName(obj uintptr, value string) {
   richEdit_SetName.Call(obj, GoStrToDStr(value))
}

func RichEdit_GetTag(obj uintptr) int {
    ret, _, _ := richEdit_GetTag.Call(obj)
    return int(ret)
}

func RichEdit_SetTag(obj uintptr, value int) {
   richEdit_SetTag.Call(obj, uintptr(value))
}

func RichEdit_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := richEdit_GetControls.Call(obj, uintptr(Index))
    return ret
}

func RichEdit_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := richEdit_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

