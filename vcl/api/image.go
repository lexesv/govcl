
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func Image_Create(obj uintptr) uintptr {
    ret, _, _ := image_Create.Call(obj)
    return ret
}

func Image_Free(obj uintptr) {
    image_Free.Call(obj)
}

func Image_BringToFront(obj uintptr)  {
    image_BringToFront.Call(obj)
}

func Image_HasParent(obj uintptr) bool {
    ret, _, _ := image_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_Hide(obj uintptr)  {
    image_Hide.Call(obj)
}

func Image_Invalidate(obj uintptr)  {
    image_Invalidate.Call(obj)
}

func Image_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := image_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func Image_Refresh(obj uintptr)  {
    image_Refresh.Call(obj)
}

func Image_Repaint(obj uintptr)  {
    image_Repaint.Call(obj)
}

func Image_SendToBack(obj uintptr)  {
    image_SendToBack.Call(obj)
}

func Image_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    image_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func Image_Show(obj uintptr)  {
    image_Show.Call(obj)
}

func Image_Update(obj uintptr)  {
    image_Update.Call(obj)
}

func Image_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := image_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Image_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := image_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Image_GetNamePath(obj uintptr) string {
    ret, _, _ := image_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Image_Assign(obj uintptr, Source uintptr)  {
    image_Assign.Call(obj, Source )
}

func Image_ClassName(obj uintptr) string {
    ret, _, _ := image_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Image_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := image_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Image_GetHashCode(obj uintptr) int32 {
    ret, _, _ := image_GetHashCode.Call(obj)
    return int32(ret)
}

func Image_ToString(obj uintptr) string {
    ret, _, _ := image_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Image_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := image_GetCanvas.Call(obj)
    return ret
}

func Image_GetAlign(obj uintptr) TAlign {
    ret, _, _ := image_GetAlign.Call(obj)
    return TAlign(ret)
}

func Image_SetAlign(obj uintptr, value TAlign) {
   image_SetAlign.Call(obj, uintptr(value))
}

func Image_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := image_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func Image_SetAnchors(obj uintptr, value TAnchors) {
   image_SetAnchors.Call(obj, uintptr(value))
}

func Image_GetAutoSize(obj uintptr) bool {
    ret, _, _ := image_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetAutoSize(obj uintptr, value bool) {
   image_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func Image_GetCenter(obj uintptr) bool {
    ret, _, _ := image_GetCenter.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetCenter(obj uintptr, value bool) {
   image_SetCenter.Call(obj, GoBoolToDBool(value))
}

func Image_GetEnabled(obj uintptr) bool {
    ret, _, _ := image_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetEnabled(obj uintptr, value bool) {
   image_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Image_GetIncrementalDisplay(obj uintptr) bool {
    ret, _, _ := image_GetIncrementalDisplay.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetIncrementalDisplay(obj uintptr, value bool) {
   image_SetIncrementalDisplay.Call(obj, GoBoolToDBool(value))
}

func Image_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := image_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetParentShowHint(obj uintptr, value bool) {
   image_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func Image_GetPicture(obj uintptr) uintptr {
    ret, _, _ := image_GetPicture.Call(obj)
    return ret
}

func Image_SetPicture(obj uintptr, value uintptr) {
   image_SetPicture.Call(obj, value)
}

func Image_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := image_GetPopupMenu.Call(obj)
    return ret
}

func Image_SetPopupMenu(obj uintptr, value uintptr) {
   image_SetPopupMenu.Call(obj, value)
}

func Image_GetProportional(obj uintptr) bool {
    ret, _, _ := image_GetProportional.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetProportional(obj uintptr, value bool) {
   image_SetProportional.Call(obj, GoBoolToDBool(value))
}

func Image_GetShowHint(obj uintptr) bool {
    ret, _, _ := image_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetShowHint(obj uintptr, value bool) {
   image_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Image_GetStretch(obj uintptr) bool {
    ret, _, _ := image_GetStretch.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetStretch(obj uintptr, value bool) {
   image_SetStretch.Call(obj, GoBoolToDBool(value))
}

func Image_GetTransparent(obj uintptr) bool {
    ret, _, _ := image_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetTransparent(obj uintptr, value bool) {
   image_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func Image_GetVisible(obj uintptr) bool {
    ret, _, _ := image_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetVisible(obj uintptr, value bool) {
   image_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Image_SetOnClick(obj uintptr, fn interface{}) {
    image_SetOnClick.Call(obj, addEventToMap(fn))
}

func Image_SetOnDblClick(obj uintptr, fn interface{}) {
    image_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func Image_SetOnMouseDown(obj uintptr, fn interface{}) {
    image_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func Image_SetOnMouseEnter(obj uintptr, fn interface{}) {
    image_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func Image_SetOnMouseLeave(obj uintptr, fn interface{}) {
    image_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func Image_SetOnMouseMove(obj uintptr, fn interface{}) {
    image_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func Image_SetOnMouseUp(obj uintptr, fn interface{}) {
    image_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func Image_GetAction(obj uintptr) uintptr {
    ret, _, _ := image_GetAction.Call(obj)
    return ret
}

func Image_SetAction(obj uintptr, value uintptr) {
   image_SetAction.Call(obj, value)
}

func Image_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    image_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Image_SetBoundsRect(obj uintptr, value TRect) {
   image_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Image_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := image_GetClientHeight.Call(obj)
    return int32(ret)
}

func Image_SetClientHeight(obj uintptr, value int32) {
   image_SetClientHeight.Call(obj, uintptr(value))
}

func Image_GetClientRect(obj uintptr) TRect {
    var ret TRect
    image_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Image_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := image_GetClientWidth.Call(obj)
    return int32(ret)
}

func Image_SetClientWidth(obj uintptr, value int32) {
   image_SetClientWidth.Call(obj, uintptr(value))
}

func Image_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := image_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Image_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := image_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Image_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := image_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Image_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := image_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Image_GetParent(obj uintptr) uintptr {
    ret, _, _ := image_GetParent.Call(obj)
    return ret
}

func Image_SetParent(obj uintptr, value uintptr) {
   image_SetParent.Call(obj, value)
}

func Image_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := image_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetAlignWithMargins(obj uintptr, value bool) {
   image_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func Image_GetLeft(obj uintptr) int32 {
    ret, _, _ := image_GetLeft.Call(obj)
    return int32(ret)
}

func Image_SetLeft(obj uintptr, value int32) {
   image_SetLeft.Call(obj, uintptr(value))
}

func Image_GetTop(obj uintptr) int32 {
    ret, _, _ := image_GetTop.Call(obj)
    return int32(ret)
}

func Image_SetTop(obj uintptr, value int32) {
   image_SetTop.Call(obj, uintptr(value))
}

func Image_GetWidth(obj uintptr) int32 {
    ret, _, _ := image_GetWidth.Call(obj)
    return int32(ret)
}

func Image_SetWidth(obj uintptr, value int32) {
   image_SetWidth.Call(obj, uintptr(value))
}

func Image_GetHeight(obj uintptr) int32 {
    ret, _, _ := image_GetHeight.Call(obj)
    return int32(ret)
}

func Image_SetHeight(obj uintptr, value int32) {
   image_SetHeight.Call(obj, uintptr(value))
}

func Image_GetCursor(obj uintptr) TCursor {
    ret, _, _ := image_GetCursor.Call(obj)
    return TCursor(ret)
}

func Image_SetCursor(obj uintptr, value TCursor) {
   image_SetCursor.Call(obj, uintptr(value))
}

func Image_GetHint(obj uintptr) string {
    ret, _, _ := image_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Image_SetHint(obj uintptr, value string) {
   image_SetHint.Call(obj, GoStrToDStr(value))
}

func Image_GetMargins(obj uintptr) uintptr {
    ret, _, _ := image_GetMargins.Call(obj)
    return ret
}

func Image_SetMargins(obj uintptr, value uintptr) {
   image_SetMargins.Call(obj, value)
}

func Image_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := image_GetComponentCount.Call(obj)
    return int32(ret)
}

func Image_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := image_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Image_SetComponentIndex(obj uintptr, value int32) {
   image_SetComponentIndex.Call(obj, uintptr(value))
}

func Image_GetOwner(obj uintptr) uintptr {
    ret, _, _ := image_GetOwner.Call(obj)
    return ret
}

func Image_GetName(obj uintptr) string {
    ret, _, _ := image_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Image_SetName(obj uintptr, value string) {
   image_SetName.Call(obj, GoStrToDStr(value))
}

func Image_GetTag(obj uintptr) int {
    ret, _, _ := image_GetTag.Call(obj)
    return int(ret)
}

func Image_SetTag(obj uintptr, value int) {
   image_SetTag.Call(obj, uintptr(value))
}

func Image_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := image_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

