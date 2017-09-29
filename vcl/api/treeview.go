
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func TreeView_Create(obj uintptr) uintptr {
    ret, _, _ := treeView_Create.Call(obj)
    return ret
}

func TreeView_Free(obj uintptr) {
    treeView_Free.Call(obj)
}

func TreeView_AlphaSort(obj uintptr, ARecurse bool) bool {
    ret, _, _ := treeView_AlphaSort.Call(obj, GoBoolToDBool(ARecurse) )
    return DBoolToGoBool(ret)
}

func TreeView_FullCollapse(obj uintptr)  {
    treeView_FullCollapse.Call(obj)
}

func TreeView_FullExpand(obj uintptr)  {
    treeView_FullExpand.Call(obj)
}

func TreeView_GetNodeAt(obj uintptr, X int32, Y int32) uintptr {
    ret, _, _ := treeView_GetNodeAt.Call(obj, uintptr(X) , uintptr(Y) )
    return ret
}

func TreeView_IsEditing(obj uintptr) bool {
    ret, _, _ := treeView_IsEditing.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_LoadFromFile(obj uintptr, FileName string)  {
    treeView_LoadFromFile.Call(obj, GoStrToDStr(FileName) )
}

func TreeView_LoadFromStream(obj uintptr, Stream uintptr)  {
    treeView_LoadFromStream.Call(obj, Stream )
}

func TreeView_SaveToFile(obj uintptr, FileName string)  {
    treeView_SaveToFile.Call(obj, GoStrToDStr(FileName) )
}

func TreeView_SaveToStream(obj uintptr, Stream uintptr)  {
    treeView_SaveToStream.Call(obj, Stream )
}

func TreeView_Deselect(obj uintptr, Node uintptr)  {
    treeView_Deselect.Call(obj, Node )
}

func TreeView_Subselect(obj uintptr, Node uintptr, Validate bool)  {
    treeView_Subselect.Call(obj, Node , GoBoolToDBool(Validate) )
}

func TreeView_ClearSelection(obj uintptr, KeepPrimary bool)  {
    treeView_ClearSelection.Call(obj, GoBoolToDBool(KeepPrimary) )
}

func TreeView_FindNextToSelect(obj uintptr) uintptr {
    ret, _, _ := treeView_FindNextToSelect.Call(obj)
    return ret
}

func TreeView_CanFocus(obj uintptr) bool {
    ret, _, _ := treeView_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_FlipChildren(obj uintptr, AllLevels bool)  {
    treeView_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func TreeView_Focused(obj uintptr) bool {
    ret, _, _ := treeView_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_HandleAllocated(obj uintptr) bool {
    ret, _, _ := treeView_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_Invalidate(obj uintptr)  {
    treeView_Invalidate.Call(obj)
}

func TreeView_Realign(obj uintptr)  {
    treeView_Realign.Call(obj)
}

func TreeView_Repaint(obj uintptr)  {
    treeView_Repaint.Call(obj)
}

func TreeView_ScaleBy(obj uintptr, M int32, D int32)  {
    treeView_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func TreeView_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    treeView_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func TreeView_SetFocus(obj uintptr)  {
    treeView_SetFocus.Call(obj)
}

func TreeView_Update(obj uintptr)  {
    treeView_Update.Call(obj)
}

func TreeView_BringToFront(obj uintptr)  {
    treeView_BringToFront.Call(obj)
}

func TreeView_HasParent(obj uintptr) bool {
    ret, _, _ := treeView_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_Hide(obj uintptr)  {
    treeView_Hide.Call(obj)
}

func TreeView_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := treeView_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func TreeView_Refresh(obj uintptr)  {
    treeView_Refresh.Call(obj)
}

func TreeView_SendToBack(obj uintptr)  {
    treeView_SendToBack.Call(obj)
}

func TreeView_Show(obj uintptr)  {
    treeView_Show.Call(obj)
}

func TreeView_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := treeView_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func TreeView_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := treeView_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func TreeView_GetNamePath(obj uintptr) string {
    ret, _, _ := treeView_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func TreeView_Assign(obj uintptr, Source uintptr)  {
    treeView_Assign.Call(obj, Source )
}

func TreeView_ClassName(obj uintptr) string {
    ret, _, _ := treeView_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func TreeView_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := treeView_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func TreeView_GetHashCode(obj uintptr) int32 {
    ret, _, _ := treeView_GetHashCode.Call(obj)
    return int32(ret)
}

func TreeView_ToString(obj uintptr) string {
    ret, _, _ := treeView_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func TreeView_GetAlign(obj uintptr) TAlign {
    ret, _, _ := treeView_GetAlign.Call(obj)
    return TAlign(ret)
}

func TreeView_SetAlign(obj uintptr, value TAlign) {
   treeView_SetAlign.Call(obj, uintptr(value))
}

func TreeView_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := treeView_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func TreeView_SetAnchors(obj uintptr, value TAnchors) {
   treeView_SetAnchors.Call(obj, uintptr(value))
}

func TreeView_GetAutoExpand(obj uintptr) bool {
    ret, _, _ := treeView_GetAutoExpand.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetAutoExpand(obj uintptr, value bool) {
   treeView_SetAutoExpand.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetBorderStyle(obj uintptr) TBorderStyle {
    ret, _, _ := treeView_GetBorderStyle.Call(obj)
    return TBorderStyle(ret)
}

func TreeView_SetBorderStyle(obj uintptr, value TBorderStyle) {
   treeView_SetBorderStyle.Call(obj, uintptr(value))
}

func TreeView_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := treeView_GetBorderWidth.Call(obj)
    return int32(ret)
}

func TreeView_SetBorderWidth(obj uintptr, value int32) {
   treeView_SetBorderWidth.Call(obj, uintptr(value))
}

func TreeView_GetChangeDelay(obj uintptr) int32 {
    ret, _, _ := treeView_GetChangeDelay.Call(obj)
    return int32(ret)
}

func TreeView_SetChangeDelay(obj uintptr, value int32) {
   treeView_SetChangeDelay.Call(obj, uintptr(value))
}

func TreeView_GetColor(obj uintptr) TColor {
    ret, _, _ := treeView_GetColor.Call(obj)
    return TColor(ret)
}

func TreeView_SetColor(obj uintptr, value TColor) {
   treeView_SetColor.Call(obj, uintptr(value))
}

func TreeView_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := treeView_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetDoubleBuffered(obj uintptr, value bool) {
   treeView_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetEnabled(obj uintptr) bool {
    ret, _, _ := treeView_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetEnabled(obj uintptr, value bool) {
   treeView_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetFont(obj uintptr) uintptr {
    ret, _, _ := treeView_GetFont.Call(obj)
    return ret
}

func TreeView_SetFont(obj uintptr, value uintptr) {
   treeView_SetFont.Call(obj, value)
}

func TreeView_GetHideSelection(obj uintptr) bool {
    ret, _, _ := treeView_GetHideSelection.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetHideSelection(obj uintptr, value bool) {
   treeView_SetHideSelection.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetHotTrack(obj uintptr) bool {
    ret, _, _ := treeView_GetHotTrack.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetHotTrack(obj uintptr, value bool) {
   treeView_SetHotTrack.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetImages(obj uintptr) uintptr {
    ret, _, _ := treeView_GetImages.Call(obj)
    return ret
}

func TreeView_SetImages(obj uintptr, value uintptr) {
   treeView_SetImages.Call(obj, value)
}

func TreeView_GetIndent(obj uintptr) int32 {
    ret, _, _ := treeView_GetIndent.Call(obj)
    return int32(ret)
}

func TreeView_SetIndent(obj uintptr, value int32) {
   treeView_SetIndent.Call(obj, uintptr(value))
}

func TreeView_GetMultiSelect(obj uintptr) bool {
    ret, _, _ := treeView_GetMultiSelect.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetMultiSelect(obj uintptr, value bool) {
   treeView_SetMultiSelect.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetMultiSelectStyle(obj uintptr) TMultiSelectStyle {
    ret, _, _ := treeView_GetMultiSelectStyle.Call(obj)
    return TMultiSelectStyle(ret)
}

func TreeView_SetMultiSelectStyle(obj uintptr, value TMultiSelectStyle) {
   treeView_SetMultiSelectStyle.Call(obj, uintptr(value))
}

func TreeView_GetParentColor(obj uintptr) bool {
    ret, _, _ := treeView_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetParentColor(obj uintptr, value bool) {
   treeView_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := treeView_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetParentCtl3D(obj uintptr, value bool) {
   treeView_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetParentFont(obj uintptr) bool {
    ret, _, _ := treeView_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetParentFont(obj uintptr, value bool) {
   treeView_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := treeView_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetParentShowHint(obj uintptr, value bool) {
   treeView_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := treeView_GetPopupMenu.Call(obj)
    return ret
}

func TreeView_SetPopupMenu(obj uintptr, value uintptr) {
   treeView_SetPopupMenu.Call(obj, value)
}

func TreeView_GetReadOnly(obj uintptr) bool {
    ret, _, _ := treeView_GetReadOnly.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetReadOnly(obj uintptr, value bool) {
   treeView_SetReadOnly.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetRightClickSelect(obj uintptr) bool {
    ret, _, _ := treeView_GetRightClickSelect.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetRightClickSelect(obj uintptr, value bool) {
   treeView_SetRightClickSelect.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetRowSelect(obj uintptr) bool {
    ret, _, _ := treeView_GetRowSelect.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetRowSelect(obj uintptr, value bool) {
   treeView_SetRowSelect.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetShowButtons(obj uintptr) bool {
    ret, _, _ := treeView_GetShowButtons.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetShowButtons(obj uintptr, value bool) {
   treeView_SetShowButtons.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetShowHint(obj uintptr) bool {
    ret, _, _ := treeView_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetShowHint(obj uintptr, value bool) {
   treeView_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetShowLines(obj uintptr) bool {
    ret, _, _ := treeView_GetShowLines.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetShowLines(obj uintptr, value bool) {
   treeView_SetShowLines.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetShowRoot(obj uintptr) bool {
    ret, _, _ := treeView_GetShowRoot.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetShowRoot(obj uintptr, value bool) {
   treeView_SetShowRoot.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetSortType(obj uintptr) TSortType {
    ret, _, _ := treeView_GetSortType.Call(obj)
    return TSortType(ret)
}

func TreeView_SetSortType(obj uintptr, value TSortType) {
   treeView_SetSortType.Call(obj, uintptr(value))
}

func TreeView_GetStateImages(obj uintptr) uintptr {
    ret, _, _ := treeView_GetStateImages.Call(obj)
    return ret
}

func TreeView_SetStateImages(obj uintptr, value uintptr) {
   treeView_SetStateImages.Call(obj, value)
}

func TreeView_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := treeView_GetTabOrder.Call(obj)
    return int16(ret)
}

func TreeView_SetTabOrder(obj uintptr, value int16) {
   treeView_SetTabOrder.Call(obj, uintptr(value))
}

func TreeView_GetTabStop(obj uintptr) bool {
    ret, _, _ := treeView_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetTabStop(obj uintptr, value bool) {
   treeView_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetToolTips(obj uintptr) bool {
    ret, _, _ := treeView_GetToolTips.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetToolTips(obj uintptr, value bool) {
   treeView_SetToolTips.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetVisible(obj uintptr) bool {
    ret, _, _ := treeView_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetVisible(obj uintptr, value bool) {
   treeView_SetVisible.Call(obj, GoBoolToDBool(value))
}

func TreeView_SetOnChange(obj uintptr, fn interface{}) {
    treeView_SetOnChange.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnClick(obj uintptr, fn interface{}) {
    treeView_SetOnClick.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnDblClick(obj uintptr, fn interface{}) {
    treeView_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnEnter(obj uintptr, fn interface{}) {
    treeView_SetOnEnter.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnExit(obj uintptr, fn interface{}) {
    treeView_SetOnExit.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnKeyDown(obj uintptr, fn interface{}) {
    treeView_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnKeyPress(obj uintptr, fn interface{}) {
    treeView_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnKeyUp(obj uintptr, fn interface{}) {
    treeView_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnMouseDown(obj uintptr, fn interface{}) {
    treeView_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnMouseEnter(obj uintptr, fn interface{}) {
    treeView_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnMouseLeave(obj uintptr, fn interface{}) {
    treeView_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnMouseMove(obj uintptr, fn interface{}) {
    treeView_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnMouseUp(obj uintptr, fn interface{}) {
    treeView_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func TreeView_GetItems(obj uintptr) uintptr {
    ret, _, _ := treeView_GetItems.Call(obj)
    return ret
}

func TreeView_SetItems(obj uintptr, value uintptr) {
   treeView_SetItems.Call(obj, value)
}

func TreeView_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := treeView_GetCanvas.Call(obj)
    return ret
}

func TreeView_GetDropTarget(obj uintptr) uintptr {
    ret, _, _ := treeView_GetDropTarget.Call(obj)
    return ret
}

func TreeView_SetDropTarget(obj uintptr, value uintptr) {
   treeView_SetDropTarget.Call(obj, value)
}

func TreeView_GetSelected(obj uintptr) uintptr {
    ret, _, _ := treeView_GetSelected.Call(obj)
    return ret
}

func TreeView_SetSelected(obj uintptr, value uintptr) {
   treeView_SetSelected.Call(obj, value)
}

func TreeView_GetTopItem(obj uintptr) uintptr {
    ret, _, _ := treeView_GetTopItem.Call(obj)
    return ret
}

func TreeView_SetTopItem(obj uintptr, value uintptr) {
   treeView_SetTopItem.Call(obj, value)
}

func TreeView_GetSelectionCount(obj uintptr) uint32 {
    ret, _, _ := treeView_GetSelectionCount.Call(obj)
    return uint32(ret)
}

func TreeView_GetBrush(obj uintptr) uintptr {
    ret, _, _ := treeView_GetBrush.Call(obj)
    return ret
}

func TreeView_GetControlCount(obj uintptr) int32 {
    ret, _, _ := treeView_GetControlCount.Call(obj)
    return int32(ret)
}

func TreeView_GetHandle(obj uintptr) HWND {
    ret, _, _ := treeView_GetHandle.Call(obj)
    return HWND(ret)
}

func TreeView_GetAction(obj uintptr) uintptr {
    ret, _, _ := treeView_GetAction.Call(obj)
    return ret
}

func TreeView_SetAction(obj uintptr, value uintptr) {
   treeView_SetAction.Call(obj, value)
}

func TreeView_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    treeView_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func TreeView_SetBoundsRect(obj uintptr, value TRect) {
   treeView_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func TreeView_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := treeView_GetClientHeight.Call(obj)
    return int32(ret)
}

func TreeView_SetClientHeight(obj uintptr, value int32) {
   treeView_SetClientHeight.Call(obj, uintptr(value))
}

func TreeView_GetClientRect(obj uintptr) TRect {
    var ret TRect
    treeView_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func TreeView_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := treeView_GetClientWidth.Call(obj)
    return int32(ret)
}

func TreeView_SetClientWidth(obj uintptr, value int32) {
   treeView_SetClientWidth.Call(obj, uintptr(value))
}

func TreeView_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := treeView_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func TreeView_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := treeView_GetExplicitTop.Call(obj)
    return int32(ret)
}

func TreeView_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := treeView_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func TreeView_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := treeView_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func TreeView_GetParent(obj uintptr) uintptr {
    ret, _, _ := treeView_GetParent.Call(obj)
    return ret
}

func TreeView_SetParent(obj uintptr, value uintptr) {
   treeView_SetParent.Call(obj, value)
}

func TreeView_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := treeView_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetAlignWithMargins(obj uintptr, value bool) {
   treeView_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetLeft(obj uintptr) int32 {
    ret, _, _ := treeView_GetLeft.Call(obj)
    return int32(ret)
}

func TreeView_SetLeft(obj uintptr, value int32) {
   treeView_SetLeft.Call(obj, uintptr(value))
}

func TreeView_GetTop(obj uintptr) int32 {
    ret, _, _ := treeView_GetTop.Call(obj)
    return int32(ret)
}

func TreeView_SetTop(obj uintptr, value int32) {
   treeView_SetTop.Call(obj, uintptr(value))
}

func TreeView_GetWidth(obj uintptr) int32 {
    ret, _, _ := treeView_GetWidth.Call(obj)
    return int32(ret)
}

func TreeView_SetWidth(obj uintptr, value int32) {
   treeView_SetWidth.Call(obj, uintptr(value))
}

func TreeView_GetHeight(obj uintptr) int32 {
    ret, _, _ := treeView_GetHeight.Call(obj)
    return int32(ret)
}

func TreeView_SetHeight(obj uintptr, value int32) {
   treeView_SetHeight.Call(obj, uintptr(value))
}

func TreeView_GetCursor(obj uintptr) TCursor {
    ret, _, _ := treeView_GetCursor.Call(obj)
    return TCursor(ret)
}

func TreeView_SetCursor(obj uintptr, value TCursor) {
   treeView_SetCursor.Call(obj, uintptr(value))
}

func TreeView_GetHint(obj uintptr) string {
    ret, _, _ := treeView_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func TreeView_SetHint(obj uintptr, value string) {
   treeView_SetHint.Call(obj, GoStrToDStr(value))
}

func TreeView_GetMargins(obj uintptr) uintptr {
    ret, _, _ := treeView_GetMargins.Call(obj)
    return ret
}

func TreeView_SetMargins(obj uintptr, value uintptr) {
   treeView_SetMargins.Call(obj, value)
}

func TreeView_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := treeView_GetComponentCount.Call(obj)
    return int32(ret)
}

func TreeView_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := treeView_GetComponentIndex.Call(obj)
    return int32(ret)
}

func TreeView_SetComponentIndex(obj uintptr, value int32) {
   treeView_SetComponentIndex.Call(obj, uintptr(value))
}

func TreeView_GetOwner(obj uintptr) uintptr {
    ret, _, _ := treeView_GetOwner.Call(obj)
    return ret
}

func TreeView_GetName(obj uintptr) string {
    ret, _, _ := treeView_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func TreeView_SetName(obj uintptr, value string) {
   treeView_SetName.Call(obj, GoStrToDStr(value))
}

func TreeView_GetTag(obj uintptr) int {
    ret, _, _ := treeView_GetTag.Call(obj)
    return int(ret)
}

func TreeView_SetTag(obj uintptr, value int) {
   treeView_SetTag.Call(obj, uintptr(value))
}

func TreeView_GetSelections(obj uintptr, Index int32) uintptr {
    ret, _, _ := treeView_GetSelections.Call(obj, uintptr(Index))
    return ret
}

func TreeView_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := treeView_GetControls.Call(obj, uintptr(Index))
    return ret
}

func TreeView_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := treeView_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

