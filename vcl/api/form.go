
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func Form_Create(obj uintptr) uintptr {
    ret, _, _ := form_Create.Call(obj)
    return ret
}

func Form_Free(obj uintptr) {
    form_Free.Call(obj)
}

func Form_Close(obj uintptr)  {
    form_Close.Call(obj)
}

func Form_Hide(obj uintptr)  {
    form_Hide.Call(obj)
}

func Form_SetFocus(obj uintptr)  {
    form_SetFocus.Call(obj)
}

func Form_Show(obj uintptr)  {
    form_Show.Call(obj)
}

func Form_ShowModal(obj uintptr) int32 {
    ret, _, _ := form_ShowModal.Call(obj)
    return int32(ret)
}

func Form_CanFocus(obj uintptr) bool {
    ret, _, _ := form_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_FlipChildren(obj uintptr, AllLevels bool)  {
    form_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func Form_Focused(obj uintptr) bool {
    ret, _, _ := form_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_HandleAllocated(obj uintptr) bool {
    ret, _, _ := form_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_Invalidate(obj uintptr)  {
    form_Invalidate.Call(obj)
}

func Form_Realign(obj uintptr)  {
    form_Realign.Call(obj)
}

func Form_Repaint(obj uintptr)  {
    form_Repaint.Call(obj)
}

func Form_ScaleBy(obj uintptr, M int32, D int32)  {
    form_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func Form_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    form_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func Form_Update(obj uintptr)  {
    form_Update.Call(obj)
}

func Form_BringToFront(obj uintptr)  {
    form_BringToFront.Call(obj)
}

func Form_HasParent(obj uintptr) bool {
    ret, _, _ := form_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := form_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func Form_Refresh(obj uintptr)  {
    form_Refresh.Call(obj)
}

func Form_SendToBack(obj uintptr)  {
    form_SendToBack.Call(obj)
}

func Form_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := form_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Form_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := form_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Form_GetNamePath(obj uintptr) string {
    ret, _, _ := form_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Form_Assign(obj uintptr, Source uintptr)  {
    form_Assign.Call(obj, Source )
}

func Form_ClassName(obj uintptr) string {
    ret, _, _ := form_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Form_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := form_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Form_GetHashCode(obj uintptr) int32 {
    ret, _, _ := form_GetHashCode.Call(obj)
    return int32(ret)
}

func Form_ToString(obj uintptr) string {
    ret, _, _ := form_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Form_GetAction(obj uintptr) uintptr {
    ret, _, _ := form_GetAction.Call(obj)
    return ret
}

func Form_SetAction(obj uintptr, value uintptr) {
   form_SetAction.Call(obj, value)
}

func Form_GetAlign(obj uintptr) TAlign {
    ret, _, _ := form_GetAlign.Call(obj)
    return TAlign(ret)
}

func Form_SetAlign(obj uintptr, value TAlign) {
   form_SetAlign.Call(obj, uintptr(value))
}

func Form_GetAlphaBlend(obj uintptr) bool {
    ret, _, _ := form_GetAlphaBlend.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetAlphaBlend(obj uintptr, value bool) {
   form_SetAlphaBlend.Call(obj, GoBoolToDBool(value))
}

func Form_GetAlphaBlendValue(obj uintptr) uint8 {
    ret, _, _ := form_GetAlphaBlendValue.Call(obj)
    return uint8(ret)
}

func Form_SetAlphaBlendValue(obj uintptr, value uint8) {
   form_SetAlphaBlendValue.Call(obj, uintptr(value))
}

func Form_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := form_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func Form_SetAnchors(obj uintptr, value TAnchors) {
   form_SetAnchors.Call(obj, uintptr(value))
}

func Form_GetAutoSize(obj uintptr) bool {
    ret, _, _ := form_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetAutoSize(obj uintptr, value bool) {
   form_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func Form_GetBorderIcons(obj uintptr) TBorderIcons {
    ret, _, _ := form_GetBorderIcons.Call(obj)
    return TBorderIcons(ret)
}

func Form_SetBorderIcons(obj uintptr, value TBorderIcons) {
   form_SetBorderIcons.Call(obj, uintptr(value))
}

func Form_GetBorderStyle(obj uintptr) TFormBorderStyle {
    ret, _, _ := form_GetBorderStyle.Call(obj)
    return TFormBorderStyle(ret)
}

func Form_SetBorderStyle(obj uintptr, value TFormBorderStyle) {
   form_SetBorderStyle.Call(obj, uintptr(value))
}

func Form_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := form_GetBorderWidth.Call(obj)
    return int32(ret)
}

func Form_SetBorderWidth(obj uintptr, value int32) {
   form_SetBorderWidth.Call(obj, uintptr(value))
}

func Form_GetCaption(obj uintptr) string {
    ret, _, _ := form_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func Form_SetCaption(obj uintptr, value string) {
   form_SetCaption.Call(obj, GoStrToDStr(value))
}

func Form_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := form_GetClientHeight.Call(obj)
    return int32(ret)
}

func Form_SetClientHeight(obj uintptr, value int32) {
   form_SetClientHeight.Call(obj, uintptr(value))
}

func Form_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := form_GetClientWidth.Call(obj)
    return int32(ret)
}

func Form_SetClientWidth(obj uintptr, value int32) {
   form_SetClientWidth.Call(obj, uintptr(value))
}

func Form_GetColor(obj uintptr) TColor {
    ret, _, _ := form_GetColor.Call(obj)
    return TColor(ret)
}

func Form_SetColor(obj uintptr, value TColor) {
   form_SetColor.Call(obj, uintptr(value))
}

func Form_GetTransparentColor(obj uintptr) bool {
    ret, _, _ := form_GetTransparentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetTransparentColor(obj uintptr, value bool) {
   form_SetTransparentColor.Call(obj, GoBoolToDBool(value))
}

func Form_GetTransparentColorValue(obj uintptr) TColor {
    ret, _, _ := form_GetTransparentColorValue.Call(obj)
    return TColor(ret)
}

func Form_SetTransparentColorValue(obj uintptr, value TColor) {
   form_SetTransparentColorValue.Call(obj, uintptr(value))
}

func Form_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := form_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetDoubleBuffered(obj uintptr, value bool) {
   form_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func Form_GetEnabled(obj uintptr) bool {
    ret, _, _ := form_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetEnabled(obj uintptr, value bool) {
   form_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Form_GetParentFont(obj uintptr) bool {
    ret, _, _ := form_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetParentFont(obj uintptr, value bool) {
   form_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func Form_GetFont(obj uintptr) uintptr {
    ret, _, _ := form_GetFont.Call(obj)
    return ret
}

func Form_SetFont(obj uintptr, value uintptr) {
   form_SetFont.Call(obj, value)
}

func Form_GetFormStyle(obj uintptr) TFormStyle {
    ret, _, _ := form_GetFormStyle.Call(obj)
    return TFormStyle(ret)
}

func Form_SetFormStyle(obj uintptr, value TFormStyle) {
   form_SetFormStyle.Call(obj, uintptr(value))
}

func Form_GetHeight(obj uintptr) int32 {
    ret, _, _ := form_GetHeight.Call(obj)
    return int32(ret)
}

func Form_SetHeight(obj uintptr, value int32) {
   form_SetHeight.Call(obj, uintptr(value))
}

func Form_GetIcon(obj uintptr) uintptr {
    ret, _, _ := form_GetIcon.Call(obj)
    return ret
}

func Form_SetIcon(obj uintptr, value uintptr) {
   form_SetIcon.Call(obj, value)
}

func Form_GetKeyPreview(obj uintptr) bool {
    ret, _, _ := form_GetKeyPreview.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetKeyPreview(obj uintptr, value bool) {
   form_SetKeyPreview.Call(obj, GoBoolToDBool(value))
}

func Form_GetMenu(obj uintptr) uintptr {
    ret, _, _ := form_GetMenu.Call(obj)
    return ret
}

func Form_SetMenu(obj uintptr, value uintptr) {
   form_SetMenu.Call(obj, value)
}

func Form_GetPixelsPerInch(obj uintptr) int32 {
    ret, _, _ := form_GetPixelsPerInch.Call(obj)
    return int32(ret)
}

func Form_SetPixelsPerInch(obj uintptr, value int32) {
   form_SetPixelsPerInch.Call(obj, uintptr(value))
}

func Form_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := form_GetPopupMenu.Call(obj)
    return ret
}

func Form_SetPopupMenu(obj uintptr, value uintptr) {
   form_SetPopupMenu.Call(obj, value)
}

func Form_GetPosition(obj uintptr) TPosition {
    ret, _, _ := form_GetPosition.Call(obj)
    return TPosition(ret)
}

func Form_SetPosition(obj uintptr, value TPosition) {
   form_SetPosition.Call(obj, uintptr(value))
}

func Form_GetScaled(obj uintptr) bool {
    ret, _, _ := form_GetScaled.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetScaled(obj uintptr, value bool) {
   form_SetScaled.Call(obj, GoBoolToDBool(value))
}

func Form_GetShowHint(obj uintptr) bool {
    ret, _, _ := form_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetShowHint(obj uintptr, value bool) {
   form_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Form_GetVisible(obj uintptr) bool {
    ret, _, _ := form_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetVisible(obj uintptr, value bool) {
   form_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Form_GetWidth(obj uintptr) int32 {
    ret, _, _ := form_GetWidth.Call(obj)
    return int32(ret)
}

func Form_SetWidth(obj uintptr, value int32) {
   form_SetWidth.Call(obj, uintptr(value))
}

func Form_GetWindowState(obj uintptr) TWindowState {
    ret, _, _ := form_GetWindowState.Call(obj)
    return TWindowState(ret)
}

func Form_SetWindowState(obj uintptr, value TWindowState) {
   form_SetWindowState.Call(obj, uintptr(value))
}

func Form_SetOnClick(obj uintptr, fn interface{}) {
    form_SetOnClick.Call(obj, addEventToMap(fn))
}

func Form_SetOnClose(obj uintptr, fn interface{}) {
    form_SetOnClose.Call(obj, addEventToMap(fn))
}

func Form_SetOnCloseQuery(obj uintptr, fn interface{}) {
    form_SetOnCloseQuery.Call(obj, addEventToMap(fn))
}

func Form_SetOnDblClick(obj uintptr, fn interface{}) {
    form_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func Form_SetOnHide(obj uintptr, fn interface{}) {
    form_SetOnHide.Call(obj, addEventToMap(fn))
}

func Form_SetOnKeyDown(obj uintptr, fn interface{}) {
    form_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func Form_SetOnKeyPress(obj uintptr, fn interface{}) {
    form_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func Form_SetOnKeyUp(obj uintptr, fn interface{}) {
    form_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func Form_SetOnMouseDown(obj uintptr, fn interface{}) {
    form_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func Form_SetOnMouseEnter(obj uintptr, fn interface{}) {
    form_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func Form_SetOnMouseLeave(obj uintptr, fn interface{}) {
    form_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func Form_SetOnMouseMove(obj uintptr, fn interface{}) {
    form_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func Form_SetOnMouseUp(obj uintptr, fn interface{}) {
    form_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func Form_SetOnMouseWheel(obj uintptr, fn interface{}) {
    form_SetOnMouseWheel.Call(obj, addEventToMap(fn))
}

func Form_SetOnPaint(obj uintptr, fn interface{}) {
    form_SetOnPaint.Call(obj, addEventToMap(fn))
}

func Form_SetOnResize(obj uintptr, fn interface{}) {
    form_SetOnResize.Call(obj, addEventToMap(fn))
}

func Form_SetOnShow(obj uintptr, fn interface{}) {
    form_SetOnShow.Call(obj, addEventToMap(fn))
}

func Form_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := form_GetCanvas.Call(obj)
    return ret
}

func Form_GetDropTarget(obj uintptr) bool {
    ret, _, _ := form_GetDropTarget.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetDropTarget(obj uintptr, value bool) {
   form_SetDropTarget.Call(obj, GoBoolToDBool(value))
}

func Form_GetModalResult(obj uintptr) TModalResult {
    ret, _, _ := form_GetModalResult.Call(obj)
    return TModalResult(ret)
}

func Form_SetModalResult(obj uintptr, value TModalResult) {
   form_SetModalResult.Call(obj, uintptr(value))
}

func Form_GetLeft(obj uintptr) int32 {
    ret, _, _ := form_GetLeft.Call(obj)
    return int32(ret)
}

func Form_SetLeft(obj uintptr, value int32) {
   form_SetLeft.Call(obj, uintptr(value))
}

func Form_GetTop(obj uintptr) int32 {
    ret, _, _ := form_GetTop.Call(obj)
    return int32(ret)
}

func Form_SetTop(obj uintptr, value int32) {
   form_SetTop.Call(obj, uintptr(value))
}

func Form_GetBrush(obj uintptr) uintptr {
    ret, _, _ := form_GetBrush.Call(obj)
    return ret
}

func Form_GetControlCount(obj uintptr) int32 {
    ret, _, _ := form_GetControlCount.Call(obj)
    return int32(ret)
}

func Form_GetHandle(obj uintptr) HWND {
    ret, _, _ := form_GetHandle.Call(obj)
    return HWND(ret)
}

func Form_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := form_GetTabOrder.Call(obj)
    return int16(ret)
}

func Form_SetTabOrder(obj uintptr, value int16) {
   form_SetTabOrder.Call(obj, uintptr(value))
}

func Form_GetTabStop(obj uintptr) bool {
    ret, _, _ := form_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetTabStop(obj uintptr, value bool) {
   form_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func Form_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    form_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Form_SetBoundsRect(obj uintptr, value TRect) {
   form_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Form_GetClientRect(obj uintptr) TRect {
    var ret TRect
    form_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Form_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := form_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Form_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := form_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Form_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := form_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Form_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := form_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Form_GetParent(obj uintptr) uintptr {
    ret, _, _ := form_GetParent.Call(obj)
    return ret
}

func Form_SetParent(obj uintptr, value uintptr) {
   form_SetParent.Call(obj, value)
}

func Form_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := form_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetAlignWithMargins(obj uintptr, value bool) {
   form_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func Form_GetCursor(obj uintptr) TCursor {
    ret, _, _ := form_GetCursor.Call(obj)
    return TCursor(ret)
}

func Form_SetCursor(obj uintptr, value TCursor) {
   form_SetCursor.Call(obj, uintptr(value))
}

func Form_GetHint(obj uintptr) string {
    ret, _, _ := form_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Form_SetHint(obj uintptr, value string) {
   form_SetHint.Call(obj, GoStrToDStr(value))
}

func Form_GetMargins(obj uintptr) uintptr {
    ret, _, _ := form_GetMargins.Call(obj)
    return ret
}

func Form_SetMargins(obj uintptr, value uintptr) {
   form_SetMargins.Call(obj, value)
}

func Form_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := form_GetComponentCount.Call(obj)
    return int32(ret)
}

func Form_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := form_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Form_SetComponentIndex(obj uintptr, value int32) {
   form_SetComponentIndex.Call(obj, uintptr(value))
}

func Form_GetOwner(obj uintptr) uintptr {
    ret, _, _ := form_GetOwner.Call(obj)
    return ret
}

func Form_GetName(obj uintptr) string {
    ret, _, _ := form_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Form_SetName(obj uintptr, value string) {
   form_SetName.Call(obj, GoStrToDStr(value))
}

func Form_GetTag(obj uintptr) int {
    ret, _, _ := form_GetTag.Call(obj)
    return int(ret)
}

func Form_SetTag(obj uintptr, value int) {
   form_SetTag.Call(obj, uintptr(value))
}

func Form_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := form_GetControls.Call(obj, uintptr(Index))
    return ret
}

func Form_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := form_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

