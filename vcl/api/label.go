
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func Label_Create(obj uintptr) uintptr {
    ret, _, _ := label_Create.Call(obj)
    return ret
}

func Label_Free(obj uintptr) {
    label_Free.Call(obj)
}

func Label_BringToFront(obj uintptr)  {
    label_BringToFront.Call(obj)
}

func Label_HasParent(obj uintptr) bool {
    ret, _, _ := label_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_Hide(obj uintptr)  {
    label_Hide.Call(obj)
}

func Label_Invalidate(obj uintptr)  {
    label_Invalidate.Call(obj)
}

func Label_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := label_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func Label_Refresh(obj uintptr)  {
    label_Refresh.Call(obj)
}

func Label_Repaint(obj uintptr)  {
    label_Repaint.Call(obj)
}

func Label_SendToBack(obj uintptr)  {
    label_SendToBack.Call(obj)
}

func Label_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    label_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func Label_Show(obj uintptr)  {
    label_Show.Call(obj)
}

func Label_Update(obj uintptr)  {
    label_Update.Call(obj)
}

func Label_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := label_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Label_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := label_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Label_GetNamePath(obj uintptr) string {
    ret, _, _ := label_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Label_Assign(obj uintptr, Source uintptr)  {
    label_Assign.Call(obj, Source )
}

func Label_ClassName(obj uintptr) string {
    ret, _, _ := label_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Label_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := label_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Label_GetHashCode(obj uintptr) int32 {
    ret, _, _ := label_GetHashCode.Call(obj)
    return int32(ret)
}

func Label_ToString(obj uintptr) string {
    ret, _, _ := label_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Label_GetAlign(obj uintptr) TAlign {
    ret, _, _ := label_GetAlign.Call(obj)
    return TAlign(ret)
}

func Label_SetAlign(obj uintptr, value TAlign) {
   label_SetAlign.Call(obj, uintptr(value))
}

func Label_GetAlignment(obj uintptr) TAlignment {
    ret, _, _ := label_GetAlignment.Call(obj)
    return TAlignment(ret)
}

func Label_SetAlignment(obj uintptr, value TAlignment) {
   label_SetAlignment.Call(obj, uintptr(value))
}

func Label_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := label_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func Label_SetAnchors(obj uintptr, value TAnchors) {
   label_SetAnchors.Call(obj, uintptr(value))
}

func Label_GetAutoSize(obj uintptr) bool {
    ret, _, _ := label_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetAutoSize(obj uintptr, value bool) {
   label_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func Label_GetCaption(obj uintptr) string {
    ret, _, _ := label_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func Label_SetCaption(obj uintptr, value string) {
   label_SetCaption.Call(obj, GoStrToDStr(value))
}

func Label_GetColor(obj uintptr) TColor {
    ret, _, _ := label_GetColor.Call(obj)
    return TColor(ret)
}

func Label_SetColor(obj uintptr, value TColor) {
   label_SetColor.Call(obj, uintptr(value))
}

func Label_GetEllipsisPosition(obj uintptr) TEllipsisPosition {
    ret, _, _ := label_GetEllipsisPosition.Call(obj)
    return TEllipsisPosition(ret)
}

func Label_SetEllipsisPosition(obj uintptr, value TEllipsisPosition) {
   label_SetEllipsisPosition.Call(obj, uintptr(value))
}

func Label_GetEnabled(obj uintptr) bool {
    ret, _, _ := label_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetEnabled(obj uintptr, value bool) {
   label_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Label_GetFont(obj uintptr) uintptr {
    ret, _, _ := label_GetFont.Call(obj)
    return ret
}

func Label_SetFont(obj uintptr, value uintptr) {
   label_SetFont.Call(obj, value)
}

func Label_GetGlowSize(obj uintptr) int32 {
    ret, _, _ := label_GetGlowSize.Call(obj)
    return int32(ret)
}

func Label_SetGlowSize(obj uintptr, value int32) {
   label_SetGlowSize.Call(obj, uintptr(value))
}

func Label_GetParentColor(obj uintptr) bool {
    ret, _, _ := label_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetParentColor(obj uintptr, value bool) {
   label_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func Label_GetParentFont(obj uintptr) bool {
    ret, _, _ := label_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetParentFont(obj uintptr, value bool) {
   label_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func Label_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := label_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetParentShowHint(obj uintptr, value bool) {
   label_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func Label_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := label_GetPopupMenu.Call(obj)
    return ret
}

func Label_SetPopupMenu(obj uintptr, value uintptr) {
   label_SetPopupMenu.Call(obj, value)
}

func Label_GetShowAccelChar(obj uintptr) bool {
    ret, _, _ := label_GetShowAccelChar.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetShowAccelChar(obj uintptr, value bool) {
   label_SetShowAccelChar.Call(obj, GoBoolToDBool(value))
}

func Label_GetShowHint(obj uintptr) bool {
    ret, _, _ := label_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetShowHint(obj uintptr, value bool) {
   label_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Label_GetTransparent(obj uintptr) bool {
    ret, _, _ := label_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetTransparent(obj uintptr, value bool) {
   label_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func Label_GetLayout(obj uintptr) TTextLayout {
    ret, _, _ := label_GetLayout.Call(obj)
    return TTextLayout(ret)
}

func Label_SetLayout(obj uintptr, value TTextLayout) {
   label_SetLayout.Call(obj, uintptr(value))
}

func Label_GetVisible(obj uintptr) bool {
    ret, _, _ := label_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetVisible(obj uintptr, value bool) {
   label_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Label_GetWordWrap(obj uintptr) bool {
    ret, _, _ := label_GetWordWrap.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetWordWrap(obj uintptr, value bool) {
   label_SetWordWrap.Call(obj, GoBoolToDBool(value))
}

func Label_SetOnClick(obj uintptr, fn interface{}) {
    label_SetOnClick.Call(obj, addEventToMap(fn))
}

func Label_SetOnDblClick(obj uintptr, fn interface{}) {
    label_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func Label_SetOnMouseDown(obj uintptr, fn interface{}) {
    label_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func Label_SetOnMouseMove(obj uintptr, fn interface{}) {
    label_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func Label_SetOnMouseUp(obj uintptr, fn interface{}) {
    label_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func Label_SetOnMouseEnter(obj uintptr, fn interface{}) {
    label_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func Label_SetOnMouseLeave(obj uintptr, fn interface{}) {
    label_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func Label_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := label_GetCanvas.Call(obj)
    return ret
}

func Label_GetAction(obj uintptr) uintptr {
    ret, _, _ := label_GetAction.Call(obj)
    return ret
}

func Label_SetAction(obj uintptr, value uintptr) {
   label_SetAction.Call(obj, value)
}

func Label_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    label_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Label_SetBoundsRect(obj uintptr, value TRect) {
   label_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Label_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := label_GetClientHeight.Call(obj)
    return int32(ret)
}

func Label_SetClientHeight(obj uintptr, value int32) {
   label_SetClientHeight.Call(obj, uintptr(value))
}

func Label_GetClientRect(obj uintptr) TRect {
    var ret TRect
    label_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Label_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := label_GetClientWidth.Call(obj)
    return int32(ret)
}

func Label_SetClientWidth(obj uintptr, value int32) {
   label_SetClientWidth.Call(obj, uintptr(value))
}

func Label_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := label_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Label_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := label_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Label_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := label_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Label_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := label_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Label_GetParent(obj uintptr) uintptr {
    ret, _, _ := label_GetParent.Call(obj)
    return ret
}

func Label_SetParent(obj uintptr, value uintptr) {
   label_SetParent.Call(obj, value)
}

func Label_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := label_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetAlignWithMargins(obj uintptr, value bool) {
   label_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func Label_GetLeft(obj uintptr) int32 {
    ret, _, _ := label_GetLeft.Call(obj)
    return int32(ret)
}

func Label_SetLeft(obj uintptr, value int32) {
   label_SetLeft.Call(obj, uintptr(value))
}

func Label_GetTop(obj uintptr) int32 {
    ret, _, _ := label_GetTop.Call(obj)
    return int32(ret)
}

func Label_SetTop(obj uintptr, value int32) {
   label_SetTop.Call(obj, uintptr(value))
}

func Label_GetWidth(obj uintptr) int32 {
    ret, _, _ := label_GetWidth.Call(obj)
    return int32(ret)
}

func Label_SetWidth(obj uintptr, value int32) {
   label_SetWidth.Call(obj, uintptr(value))
}

func Label_GetHeight(obj uintptr) int32 {
    ret, _, _ := label_GetHeight.Call(obj)
    return int32(ret)
}

func Label_SetHeight(obj uintptr, value int32) {
   label_SetHeight.Call(obj, uintptr(value))
}

func Label_GetCursor(obj uintptr) TCursor {
    ret, _, _ := label_GetCursor.Call(obj)
    return TCursor(ret)
}

func Label_SetCursor(obj uintptr, value TCursor) {
   label_SetCursor.Call(obj, uintptr(value))
}

func Label_GetHint(obj uintptr) string {
    ret, _, _ := label_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Label_SetHint(obj uintptr, value string) {
   label_SetHint.Call(obj, GoStrToDStr(value))
}

func Label_GetMargins(obj uintptr) uintptr {
    ret, _, _ := label_GetMargins.Call(obj)
    return ret
}

func Label_SetMargins(obj uintptr, value uintptr) {
   label_SetMargins.Call(obj, value)
}

func Label_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := label_GetComponentCount.Call(obj)
    return int32(ret)
}

func Label_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := label_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Label_SetComponentIndex(obj uintptr, value int32) {
   label_SetComponentIndex.Call(obj, uintptr(value))
}

func Label_GetOwner(obj uintptr) uintptr {
    ret, _, _ := label_GetOwner.Call(obj)
    return ret
}

func Label_GetName(obj uintptr) string {
    ret, _, _ := label_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Label_SetName(obj uintptr, value string) {
   label_SetName.Call(obj, GoStrToDStr(value))
}

func Label_GetTag(obj uintptr) int {
    ret, _, _ := label_GetTag.Call(obj)
    return int(ret)
}

func Label_SetTag(obj uintptr, value int) {
   label_SetTag.Call(obj, uintptr(value))
}

func Label_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := label_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

