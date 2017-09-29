
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func ToolBar_Create(obj uintptr) uintptr {
    ret, _, _ := toolBar_Create.Call(obj)
    return ret
}

func ToolBar_Free(obj uintptr) {
    toolBar_Free.Call(obj)
}

func ToolBar_FlipChildren(obj uintptr, AllLevels bool)  {
    toolBar_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func ToolBar_CanFocus(obj uintptr) bool {
    ret, _, _ := toolBar_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_Focused(obj uintptr) bool {
    ret, _, _ := toolBar_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_HandleAllocated(obj uintptr) bool {
    ret, _, _ := toolBar_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_Invalidate(obj uintptr)  {
    toolBar_Invalidate.Call(obj)
}

func ToolBar_Realign(obj uintptr)  {
    toolBar_Realign.Call(obj)
}

func ToolBar_Repaint(obj uintptr)  {
    toolBar_Repaint.Call(obj)
}

func ToolBar_ScaleBy(obj uintptr, M int32, D int32)  {
    toolBar_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func ToolBar_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    toolBar_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func ToolBar_SetFocus(obj uintptr)  {
    toolBar_SetFocus.Call(obj)
}

func ToolBar_Update(obj uintptr)  {
    toolBar_Update.Call(obj)
}

func ToolBar_BringToFront(obj uintptr)  {
    toolBar_BringToFront.Call(obj)
}

func ToolBar_HasParent(obj uintptr) bool {
    ret, _, _ := toolBar_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_Hide(obj uintptr)  {
    toolBar_Hide.Call(obj)
}

func ToolBar_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := toolBar_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func ToolBar_Refresh(obj uintptr)  {
    toolBar_Refresh.Call(obj)
}

func ToolBar_SendToBack(obj uintptr)  {
    toolBar_SendToBack.Call(obj)
}

func ToolBar_Show(obj uintptr)  {
    toolBar_Show.Call(obj)
}

func ToolBar_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := toolBar_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func ToolBar_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := toolBar_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ToolBar_GetNamePath(obj uintptr) string {
    ret, _, _ := toolBar_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ToolBar_Assign(obj uintptr, Source uintptr)  {
    toolBar_Assign.Call(obj, Source )
}

func ToolBar_ClassName(obj uintptr) string {
    ret, _, _ := toolBar_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ToolBar_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := toolBar_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ToolBar_GetHashCode(obj uintptr) int32 {
    ret, _, _ := toolBar_GetHashCode.Call(obj)
    return int32(ret)
}

func ToolBar_ToString(obj uintptr) string {
    ret, _, _ := toolBar_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ToolBar_GetButtonCount(obj uintptr) int32 {
    ret, _, _ := toolBar_GetButtonCount.Call(obj)
    return int32(ret)
}

func ToolBar_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetCanvas.Call(obj)
    return ret
}

func ToolBar_GetCustomizeKeyName(obj uintptr) string {
    ret, _, _ := toolBar_GetCustomizeKeyName.Call(obj)
    return DStrToGoStr(ret)
}

func ToolBar_SetCustomizeKeyName(obj uintptr, value string) {
   toolBar_SetCustomizeKeyName.Call(obj, GoStrToDStr(value))
}

func ToolBar_GetCustomizeValueName(obj uintptr) string {
    ret, _, _ := toolBar_GetCustomizeValueName.Call(obj)
    return DStrToGoStr(ret)
}

func ToolBar_SetCustomizeValueName(obj uintptr, value string) {
   toolBar_SetCustomizeValueName.Call(obj, GoStrToDStr(value))
}

func ToolBar_GetRowCount(obj uintptr) int32 {
    ret, _, _ := toolBar_GetRowCount.Call(obj)
    return int32(ret)
}

func ToolBar_GetAlign(obj uintptr) TAlign {
    ret, _, _ := toolBar_GetAlign.Call(obj)
    return TAlign(ret)
}

func ToolBar_SetAlign(obj uintptr, value TAlign) {
   toolBar_SetAlign.Call(obj, uintptr(value))
}

func ToolBar_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := toolBar_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func ToolBar_SetAnchors(obj uintptr, value TAnchors) {
   toolBar_SetAnchors.Call(obj, uintptr(value))
}

func ToolBar_GetAutoSize(obj uintptr) bool {
    ret, _, _ := toolBar_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetAutoSize(obj uintptr, value bool) {
   toolBar_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := toolBar_GetBorderWidth.Call(obj)
    return int32(ret)
}

func ToolBar_SetBorderWidth(obj uintptr, value int32) {
   toolBar_SetBorderWidth.Call(obj, uintptr(value))
}

func ToolBar_GetButtonHeight(obj uintptr) int32 {
    ret, _, _ := toolBar_GetButtonHeight.Call(obj)
    return int32(ret)
}

func ToolBar_SetButtonHeight(obj uintptr, value int32) {
   toolBar_SetButtonHeight.Call(obj, uintptr(value))
}

func ToolBar_GetButtonWidth(obj uintptr) int32 {
    ret, _, _ := toolBar_GetButtonWidth.Call(obj)
    return int32(ret)
}

func ToolBar_SetButtonWidth(obj uintptr, value int32) {
   toolBar_SetButtonWidth.Call(obj, uintptr(value))
}

func ToolBar_GetCaption(obj uintptr) string {
    ret, _, _ := toolBar_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func ToolBar_SetCaption(obj uintptr, value string) {
   toolBar_SetCaption.Call(obj, GoStrToDStr(value))
}

func ToolBar_GetColor(obj uintptr) TColor {
    ret, _, _ := toolBar_GetColor.Call(obj)
    return TColor(ret)
}

func ToolBar_SetColor(obj uintptr, value TColor) {
   toolBar_SetColor.Call(obj, uintptr(value))
}

func ToolBar_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := toolBar_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetDoubleBuffered(obj uintptr, value bool) {
   toolBar_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetDrawingStyle(obj uintptr) TTBDrawingStyle {
    ret, _, _ := toolBar_GetDrawingStyle.Call(obj)
    return TTBDrawingStyle(ret)
}

func ToolBar_SetDrawingStyle(obj uintptr, value TTBDrawingStyle) {
   toolBar_SetDrawingStyle.Call(obj, uintptr(value))
}

func ToolBar_GetEnabled(obj uintptr) bool {
    ret, _, _ := toolBar_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetEnabled(obj uintptr, value bool) {
   toolBar_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetFlat(obj uintptr) bool {
    ret, _, _ := toolBar_GetFlat.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetFlat(obj uintptr, value bool) {
   toolBar_SetFlat.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetFont(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetFont.Call(obj)
    return ret
}

func ToolBar_SetFont(obj uintptr, value uintptr) {
   toolBar_SetFont.Call(obj, value)
}

func ToolBar_GetGradientEndColor(obj uintptr) TColor {
    ret, _, _ := toolBar_GetGradientEndColor.Call(obj)
    return TColor(ret)
}

func ToolBar_SetGradientEndColor(obj uintptr, value TColor) {
   toolBar_SetGradientEndColor.Call(obj, uintptr(value))
}

func ToolBar_GetGradientStartColor(obj uintptr) TColor {
    ret, _, _ := toolBar_GetGradientStartColor.Call(obj)
    return TColor(ret)
}

func ToolBar_SetGradientStartColor(obj uintptr, value TColor) {
   toolBar_SetGradientStartColor.Call(obj, uintptr(value))
}

func ToolBar_GetHeight(obj uintptr) int32 {
    ret, _, _ := toolBar_GetHeight.Call(obj)
    return int32(ret)
}

func ToolBar_SetHeight(obj uintptr, value int32) {
   toolBar_SetHeight.Call(obj, uintptr(value))
}

func ToolBar_GetHideClippedButtons(obj uintptr) bool {
    ret, _, _ := toolBar_GetHideClippedButtons.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetHideClippedButtons(obj uintptr, value bool) {
   toolBar_SetHideClippedButtons.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetHotImages(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetHotImages.Call(obj)
    return ret
}

func ToolBar_SetHotImages(obj uintptr, value uintptr) {
   toolBar_SetHotImages.Call(obj, value)
}

func ToolBar_GetImages(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetImages.Call(obj)
    return ret
}

func ToolBar_SetImages(obj uintptr, value uintptr) {
   toolBar_SetImages.Call(obj, value)
}

func ToolBar_GetIndent(obj uintptr) int32 {
    ret, _, _ := toolBar_GetIndent.Call(obj)
    return int32(ret)
}

func ToolBar_SetIndent(obj uintptr, value int32) {
   toolBar_SetIndent.Call(obj, uintptr(value))
}

func ToolBar_GetList(obj uintptr) bool {
    ret, _, _ := toolBar_GetList.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetList(obj uintptr, value bool) {
   toolBar_SetList.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetMenu(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetMenu.Call(obj)
    return ret
}

func ToolBar_SetMenu(obj uintptr, value uintptr) {
   toolBar_SetMenu.Call(obj, value)
}

func ToolBar_GetGradientDirection(obj uintptr) TGradientDirection {
    ret, _, _ := toolBar_GetGradientDirection.Call(obj)
    return TGradientDirection(ret)
}

func ToolBar_SetGradientDirection(obj uintptr, value TGradientDirection) {
   toolBar_SetGradientDirection.Call(obj, uintptr(value))
}

func ToolBar_GetGradientDrawingOptions(obj uintptr) TTBGradientDrawingOptions {
    ret, _, _ := toolBar_GetGradientDrawingOptions.Call(obj)
    return TTBGradientDrawingOptions(ret)
}

func ToolBar_SetGradientDrawingOptions(obj uintptr, value TTBGradientDrawingOptions) {
   toolBar_SetGradientDrawingOptions.Call(obj, uintptr(value))
}

func ToolBar_GetParentColor(obj uintptr) bool {
    ret, _, _ := toolBar_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetParentColor(obj uintptr, value bool) {
   toolBar_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetParentFont(obj uintptr) bool {
    ret, _, _ := toolBar_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetParentFont(obj uintptr, value bool) {
   toolBar_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := toolBar_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetParentShowHint(obj uintptr, value bool) {
   toolBar_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetPopupMenu.Call(obj)
    return ret
}

func ToolBar_SetPopupMenu(obj uintptr, value uintptr) {
   toolBar_SetPopupMenu.Call(obj, value)
}

func ToolBar_GetShowCaptions(obj uintptr) bool {
    ret, _, _ := toolBar_GetShowCaptions.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetShowCaptions(obj uintptr, value bool) {
   toolBar_SetShowCaptions.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetShowHint(obj uintptr) bool {
    ret, _, _ := toolBar_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetShowHint(obj uintptr, value bool) {
   toolBar_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := toolBar_GetTabOrder.Call(obj)
    return int16(ret)
}

func ToolBar_SetTabOrder(obj uintptr, value int16) {
   toolBar_SetTabOrder.Call(obj, uintptr(value))
}

func ToolBar_GetTabStop(obj uintptr) bool {
    ret, _, _ := toolBar_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetTabStop(obj uintptr, value bool) {
   toolBar_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetTransparent(obj uintptr) bool {
    ret, _, _ := toolBar_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetTransparent(obj uintptr, value bool) {
   toolBar_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetVisible(obj uintptr) bool {
    ret, _, _ := toolBar_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetVisible(obj uintptr, value bool) {
   toolBar_SetVisible.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetWrapable(obj uintptr) bool {
    ret, _, _ := toolBar_GetWrapable.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetWrapable(obj uintptr, value bool) {
   toolBar_SetWrapable.Call(obj, GoBoolToDBool(value))
}

func ToolBar_SetOnClick(obj uintptr, fn interface{}) {
    toolBar_SetOnClick.Call(obj, addEventToMap(fn))
}

func ToolBar_SetOnDblClick(obj uintptr, fn interface{}) {
    toolBar_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func ToolBar_SetOnEnter(obj uintptr, fn interface{}) {
    toolBar_SetOnEnter.Call(obj, addEventToMap(fn))
}

func ToolBar_SetOnExit(obj uintptr, fn interface{}) {
    toolBar_SetOnExit.Call(obj, addEventToMap(fn))
}

func ToolBar_SetOnMouseDown(obj uintptr, fn interface{}) {
    toolBar_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func ToolBar_SetOnMouseEnter(obj uintptr, fn interface{}) {
    toolBar_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func ToolBar_SetOnMouseLeave(obj uintptr, fn interface{}) {
    toolBar_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func ToolBar_SetOnMouseMove(obj uintptr, fn interface{}) {
    toolBar_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func ToolBar_SetOnMouseUp(obj uintptr, fn interface{}) {
    toolBar_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func ToolBar_SetOnResize(obj uintptr, fn interface{}) {
    toolBar_SetOnResize.Call(obj, addEventToMap(fn))
}

func ToolBar_GetBrush(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetBrush.Call(obj)
    return ret
}

func ToolBar_GetControlCount(obj uintptr) int32 {
    ret, _, _ := toolBar_GetControlCount.Call(obj)
    return int32(ret)
}

func ToolBar_GetHandle(obj uintptr) HWND {
    ret, _, _ := toolBar_GetHandle.Call(obj)
    return HWND(ret)
}

func ToolBar_GetAction(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetAction.Call(obj)
    return ret
}

func ToolBar_SetAction(obj uintptr, value uintptr) {
   toolBar_SetAction.Call(obj, value)
}

func ToolBar_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    toolBar_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ToolBar_SetBoundsRect(obj uintptr, value TRect) {
   toolBar_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ToolBar_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := toolBar_GetClientHeight.Call(obj)
    return int32(ret)
}

func ToolBar_SetClientHeight(obj uintptr, value int32) {
   toolBar_SetClientHeight.Call(obj, uintptr(value))
}

func ToolBar_GetClientRect(obj uintptr) TRect {
    var ret TRect
    toolBar_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ToolBar_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := toolBar_GetClientWidth.Call(obj)
    return int32(ret)
}

func ToolBar_SetClientWidth(obj uintptr, value int32) {
   toolBar_SetClientWidth.Call(obj, uintptr(value))
}

func ToolBar_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := toolBar_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func ToolBar_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := toolBar_GetExplicitTop.Call(obj)
    return int32(ret)
}

func ToolBar_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := toolBar_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func ToolBar_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := toolBar_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func ToolBar_GetParent(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetParent.Call(obj)
    return ret
}

func ToolBar_SetParent(obj uintptr, value uintptr) {
   toolBar_SetParent.Call(obj, value)
}

func ToolBar_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := toolBar_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetAlignWithMargins(obj uintptr, value bool) {
   toolBar_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetLeft(obj uintptr) int32 {
    ret, _, _ := toolBar_GetLeft.Call(obj)
    return int32(ret)
}

func ToolBar_SetLeft(obj uintptr, value int32) {
   toolBar_SetLeft.Call(obj, uintptr(value))
}

func ToolBar_GetTop(obj uintptr) int32 {
    ret, _, _ := toolBar_GetTop.Call(obj)
    return int32(ret)
}

func ToolBar_SetTop(obj uintptr, value int32) {
   toolBar_SetTop.Call(obj, uintptr(value))
}

func ToolBar_GetWidth(obj uintptr) int32 {
    ret, _, _ := toolBar_GetWidth.Call(obj)
    return int32(ret)
}

func ToolBar_SetWidth(obj uintptr, value int32) {
   toolBar_SetWidth.Call(obj, uintptr(value))
}

func ToolBar_GetCursor(obj uintptr) TCursor {
    ret, _, _ := toolBar_GetCursor.Call(obj)
    return TCursor(ret)
}

func ToolBar_SetCursor(obj uintptr, value TCursor) {
   toolBar_SetCursor.Call(obj, uintptr(value))
}

func ToolBar_GetHint(obj uintptr) string {
    ret, _, _ := toolBar_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func ToolBar_SetHint(obj uintptr, value string) {
   toolBar_SetHint.Call(obj, GoStrToDStr(value))
}

func ToolBar_GetMargins(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetMargins.Call(obj)
    return ret
}

func ToolBar_SetMargins(obj uintptr, value uintptr) {
   toolBar_SetMargins.Call(obj, value)
}

func ToolBar_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := toolBar_GetComponentCount.Call(obj)
    return int32(ret)
}

func ToolBar_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := toolBar_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ToolBar_SetComponentIndex(obj uintptr, value int32) {
   toolBar_SetComponentIndex.Call(obj, uintptr(value))
}

func ToolBar_GetOwner(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetOwner.Call(obj)
    return ret
}

func ToolBar_GetName(obj uintptr) string {
    ret, _, _ := toolBar_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ToolBar_SetName(obj uintptr, value string) {
   toolBar_SetName.Call(obj, GoStrToDStr(value))
}

func ToolBar_GetTag(obj uintptr) int {
    ret, _, _ := toolBar_GetTag.Call(obj)
    return int(ret)
}

func ToolBar_SetTag(obj uintptr, value int) {
   toolBar_SetTag.Call(obj, uintptr(value))
}

func ToolBar_GetButtons(obj uintptr, Index int32) uintptr {
    ret, _, _ := toolBar_GetButtons.Call(obj, uintptr(Index))
    return ret
}

func ToolBar_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := toolBar_GetControls.Call(obj, uintptr(Index))
    return ret
}

func ToolBar_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := toolBar_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

