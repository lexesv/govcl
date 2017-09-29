
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func ListView_Create(obj uintptr) uintptr {
    ret, _, _ := listView_Create.Call(obj)
    return ret
}

func ListView_Free(obj uintptr) {
    listView_Free.Call(obj)
}

func ListView_AddItem(obj uintptr, Item string, AObject uintptr)  {
    listView_AddItem.Call(obj, GoStrToDStr(Item) , AObject )
}

func ListView_AlphaSort(obj uintptr) bool {
    ret, _, _ := listView_AlphaSort.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_Clear(obj uintptr)  {
    listView_Clear.Call(obj)
}

func ListView_ClearSelection(obj uintptr)  {
    listView_ClearSelection.Call(obj)
}

func ListView_DeleteSelected(obj uintptr)  {
    listView_DeleteSelected.Call(obj)
}

func ListView_GetSearchString(obj uintptr) string {
    ret, _, _ := listView_GetSearchString.Call(obj)
    return DStrToGoStr(ret)
}

func ListView_IsEditing(obj uintptr) bool {
    ret, _, _ := listView_IsEditing.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SelectAll(obj uintptr)  {
    listView_SelectAll.Call(obj)
}

func ListView_Scroll(obj uintptr, DX int32, DY int32)  {
    listView_Scroll.Call(obj, uintptr(DX) , uintptr(DY) )
}

func ListView_CanFocus(obj uintptr) bool {
    ret, _, _ := listView_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_FlipChildren(obj uintptr, AllLevels bool)  {
    listView_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func ListView_Focused(obj uintptr) bool {
    ret, _, _ := listView_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_HandleAllocated(obj uintptr) bool {
    ret, _, _ := listView_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_Invalidate(obj uintptr)  {
    listView_Invalidate.Call(obj)
}

func ListView_Realign(obj uintptr)  {
    listView_Realign.Call(obj)
}

func ListView_Repaint(obj uintptr)  {
    listView_Repaint.Call(obj)
}

func ListView_ScaleBy(obj uintptr, M int32, D int32)  {
    listView_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func ListView_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    listView_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func ListView_SetFocus(obj uintptr)  {
    listView_SetFocus.Call(obj)
}

func ListView_Update(obj uintptr)  {
    listView_Update.Call(obj)
}

func ListView_BringToFront(obj uintptr)  {
    listView_BringToFront.Call(obj)
}

func ListView_HasParent(obj uintptr) bool {
    ret, _, _ := listView_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_Hide(obj uintptr)  {
    listView_Hide.Call(obj)
}

func ListView_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := listView_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func ListView_Refresh(obj uintptr)  {
    listView_Refresh.Call(obj)
}

func ListView_SendToBack(obj uintptr)  {
    listView_SendToBack.Call(obj)
}

func ListView_Show(obj uintptr)  {
    listView_Show.Call(obj)
}

func ListView_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := listView_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func ListView_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := listView_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ListView_GetNamePath(obj uintptr) string {
    ret, _, _ := listView_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ListView_Assign(obj uintptr, Source uintptr)  {
    listView_Assign.Call(obj, Source )
}

func ListView_ClassName(obj uintptr) string {
    ret, _, _ := listView_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ListView_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := listView_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ListView_GetHashCode(obj uintptr) int32 {
    ret, _, _ := listView_GetHashCode.Call(obj)
    return int32(ret)
}

func ListView_ToString(obj uintptr) string {
    ret, _, _ := listView_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ListView_GetAction(obj uintptr) uintptr {
    ret, _, _ := listView_GetAction.Call(obj)
    return ret
}

func ListView_SetAction(obj uintptr, value uintptr) {
   listView_SetAction.Call(obj, value)
}

func ListView_GetAlign(obj uintptr) TAlign {
    ret, _, _ := listView_GetAlign.Call(obj)
    return TAlign(ret)
}

func ListView_SetAlign(obj uintptr, value TAlign) {
   listView_SetAlign.Call(obj, uintptr(value))
}

func ListView_GetAllocBy(obj uintptr) int32 {
    ret, _, _ := listView_GetAllocBy.Call(obj)
    return int32(ret)
}

func ListView_SetAllocBy(obj uintptr, value int32) {
   listView_SetAllocBy.Call(obj, uintptr(value))
}

func ListView_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := listView_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func ListView_SetAnchors(obj uintptr, value TAnchors) {
   listView_SetAnchors.Call(obj, uintptr(value))
}

func ListView_GetBorderStyle(obj uintptr) TBorderStyle {
    ret, _, _ := listView_GetBorderStyle.Call(obj)
    return TBorderStyle(ret)
}

func ListView_SetBorderStyle(obj uintptr, value TBorderStyle) {
   listView_SetBorderStyle.Call(obj, uintptr(value))
}

func ListView_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := listView_GetBorderWidth.Call(obj)
    return int32(ret)
}

func ListView_SetBorderWidth(obj uintptr, value int32) {
   listView_SetBorderWidth.Call(obj, uintptr(value))
}

func ListView_GetCheckboxes(obj uintptr) bool {
    ret, _, _ := listView_GetCheckboxes.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetCheckboxes(obj uintptr, value bool) {
   listView_SetCheckboxes.Call(obj, GoBoolToDBool(value))
}

func ListView_GetColor(obj uintptr) TColor {
    ret, _, _ := listView_GetColor.Call(obj)
    return TColor(ret)
}

func ListView_SetColor(obj uintptr, value TColor) {
   listView_SetColor.Call(obj, uintptr(value))
}

func ListView_GetColumns(obj uintptr) uintptr {
    ret, _, _ := listView_GetColumns.Call(obj)
    return ret
}

func ListView_SetColumns(obj uintptr, value uintptr) {
   listView_SetColumns.Call(obj, value)
}

func ListView_GetColumnClick(obj uintptr) bool {
    ret, _, _ := listView_GetColumnClick.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetColumnClick(obj uintptr, value bool) {
   listView_SetColumnClick.Call(obj, GoBoolToDBool(value))
}

func ListView_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := listView_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetDoubleBuffered(obj uintptr, value bool) {
   listView_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ListView_GetEnabled(obj uintptr) bool {
    ret, _, _ := listView_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetEnabled(obj uintptr, value bool) {
   listView_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func ListView_GetFont(obj uintptr) uintptr {
    ret, _, _ := listView_GetFont.Call(obj)
    return ret
}

func ListView_SetFont(obj uintptr, value uintptr) {
   listView_SetFont.Call(obj, value)
}

func ListView_GetFlatScrollBars(obj uintptr) bool {
    ret, _, _ := listView_GetFlatScrollBars.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetFlatScrollBars(obj uintptr, value bool) {
   listView_SetFlatScrollBars.Call(obj, GoBoolToDBool(value))
}

func ListView_GetFullDrag(obj uintptr) bool {
    ret, _, _ := listView_GetFullDrag.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetFullDrag(obj uintptr, value bool) {
   listView_SetFullDrag.Call(obj, GoBoolToDBool(value))
}

func ListView_GetGridLines(obj uintptr) bool {
    ret, _, _ := listView_GetGridLines.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetGridLines(obj uintptr, value bool) {
   listView_SetGridLines.Call(obj, GoBoolToDBool(value))
}

func ListView_GetGroups(obj uintptr) uintptr {
    ret, _, _ := listView_GetGroups.Call(obj)
    return ret
}

func ListView_SetGroups(obj uintptr, value uintptr) {
   listView_SetGroups.Call(obj, value)
}

func ListView_GetHideSelection(obj uintptr) bool {
    ret, _, _ := listView_GetHideSelection.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetHideSelection(obj uintptr, value bool) {
   listView_SetHideSelection.Call(obj, GoBoolToDBool(value))
}

func ListView_GetHotTrack(obj uintptr) bool {
    ret, _, _ := listView_GetHotTrack.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetHotTrack(obj uintptr, value bool) {
   listView_SetHotTrack.Call(obj, GoBoolToDBool(value))
}

func ListView_GetHoverTime(obj uintptr) int32 {
    ret, _, _ := listView_GetHoverTime.Call(obj)
    return int32(ret)
}

func ListView_SetHoverTime(obj uintptr, value int32) {
   listView_SetHoverTime.Call(obj, uintptr(value))
}

func ListView_GetItems(obj uintptr) uintptr {
    ret, _, _ := listView_GetItems.Call(obj)
    return ret
}

func ListView_SetItems(obj uintptr, value uintptr) {
   listView_SetItems.Call(obj, value)
}

func ListView_GetLargeImages(obj uintptr) uintptr {
    ret, _, _ := listView_GetLargeImages.Call(obj)
    return ret
}

func ListView_SetLargeImages(obj uintptr, value uintptr) {
   listView_SetLargeImages.Call(obj, value)
}

func ListView_GetMultiSelect(obj uintptr) bool {
    ret, _, _ := listView_GetMultiSelect.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetMultiSelect(obj uintptr, value bool) {
   listView_SetMultiSelect.Call(obj, GoBoolToDBool(value))
}

func ListView_GetGroupHeaderImages(obj uintptr) uintptr {
    ret, _, _ := listView_GetGroupHeaderImages.Call(obj)
    return ret
}

func ListView_SetGroupHeaderImages(obj uintptr, value uintptr) {
   listView_SetGroupHeaderImages.Call(obj, value)
}

func ListView_GetReadOnly(obj uintptr) bool {
    ret, _, _ := listView_GetReadOnly.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetReadOnly(obj uintptr, value bool) {
   listView_SetReadOnly.Call(obj, GoBoolToDBool(value))
}

func ListView_GetRowSelect(obj uintptr) bool {
    ret, _, _ := listView_GetRowSelect.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetRowSelect(obj uintptr, value bool) {
   listView_SetRowSelect.Call(obj, GoBoolToDBool(value))
}

func ListView_GetParentColor(obj uintptr) bool {
    ret, _, _ := listView_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetParentColor(obj uintptr, value bool) {
   listView_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func ListView_GetParentFont(obj uintptr) bool {
    ret, _, _ := listView_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetParentFont(obj uintptr, value bool) {
   listView_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func ListView_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := listView_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetParentShowHint(obj uintptr, value bool) {
   listView_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func ListView_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := listView_GetPopupMenu.Call(obj)
    return ret
}

func ListView_SetPopupMenu(obj uintptr, value uintptr) {
   listView_SetPopupMenu.Call(obj, value)
}

func ListView_GetShowColumnHeaders(obj uintptr) bool {
    ret, _, _ := listView_GetShowColumnHeaders.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetShowColumnHeaders(obj uintptr, value bool) {
   listView_SetShowColumnHeaders.Call(obj, GoBoolToDBool(value))
}

func ListView_GetShowWorkAreas(obj uintptr) bool {
    ret, _, _ := listView_GetShowWorkAreas.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetShowWorkAreas(obj uintptr, value bool) {
   listView_SetShowWorkAreas.Call(obj, GoBoolToDBool(value))
}

func ListView_GetShowHint(obj uintptr) bool {
    ret, _, _ := listView_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetShowHint(obj uintptr, value bool) {
   listView_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func ListView_GetSmallImages(obj uintptr) uintptr {
    ret, _, _ := listView_GetSmallImages.Call(obj)
    return ret
}

func ListView_SetSmallImages(obj uintptr, value uintptr) {
   listView_SetSmallImages.Call(obj, value)
}

func ListView_GetSortType(obj uintptr) TSortType {
    ret, _, _ := listView_GetSortType.Call(obj)
    return TSortType(ret)
}

func ListView_SetSortType(obj uintptr, value TSortType) {
   listView_SetSortType.Call(obj, uintptr(value))
}

func ListView_GetStateImages(obj uintptr) uintptr {
    ret, _, _ := listView_GetStateImages.Call(obj)
    return ret
}

func ListView_SetStateImages(obj uintptr, value uintptr) {
   listView_SetStateImages.Call(obj, value)
}

func ListView_GetTabOrder(obj uintptr) int16 {
    ret, _, _ := listView_GetTabOrder.Call(obj)
    return int16(ret)
}

func ListView_SetTabOrder(obj uintptr, value int16) {
   listView_SetTabOrder.Call(obj, uintptr(value))
}

func ListView_GetTabStop(obj uintptr) bool {
    ret, _, _ := listView_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetTabStop(obj uintptr, value bool) {
   listView_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func ListView_GetViewStyle(obj uintptr) TViewStyle {
    ret, _, _ := listView_GetViewStyle.Call(obj)
    return TViewStyle(ret)
}

func ListView_SetViewStyle(obj uintptr, value TViewStyle) {
   listView_SetViewStyle.Call(obj, uintptr(value))
}

func ListView_GetVisible(obj uintptr) bool {
    ret, _, _ := listView_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetVisible(obj uintptr, value bool) {
   listView_SetVisible.Call(obj, GoBoolToDBool(value))
}

func ListView_SetOnChange(obj uintptr, fn interface{}) {
    listView_SetOnChange.Call(obj, addEventToMap(fn))
}

func ListView_SetOnClick(obj uintptr, fn interface{}) {
    listView_SetOnClick.Call(obj, addEventToMap(fn))
}

func ListView_SetOnDblClick(obj uintptr, fn interface{}) {
    listView_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func ListView_SetOnEnter(obj uintptr, fn interface{}) {
    listView_SetOnEnter.Call(obj, addEventToMap(fn))
}

func ListView_SetOnExit(obj uintptr, fn interface{}) {
    listView_SetOnExit.Call(obj, addEventToMap(fn))
}

func ListView_SetOnKeyDown(obj uintptr, fn interface{}) {
    listView_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func ListView_SetOnKeyPress(obj uintptr, fn interface{}) {
    listView_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func ListView_SetOnKeyUp(obj uintptr, fn interface{}) {
    listView_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func ListView_SetOnMouseDown(obj uintptr, fn interface{}) {
    listView_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func ListView_SetOnMouseEnter(obj uintptr, fn interface{}) {
    listView_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func ListView_SetOnMouseLeave(obj uintptr, fn interface{}) {
    listView_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func ListView_SetOnMouseMove(obj uintptr, fn interface{}) {
    listView_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func ListView_SetOnMouseUp(obj uintptr, fn interface{}) {
    listView_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func ListView_SetOnResize(obj uintptr, fn interface{}) {
    listView_SetOnResize.Call(obj, addEventToMap(fn))
}

func ListView_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := listView_GetCanvas.Call(obj)
    return ret
}

func ListView_GetDropTarget(obj uintptr) uintptr {
    ret, _, _ := listView_GetDropTarget.Call(obj)
    return ret
}

func ListView_SetDropTarget(obj uintptr, value uintptr) {
   listView_SetDropTarget.Call(obj, value)
}

func ListView_GetItemFocused(obj uintptr) uintptr {
    ret, _, _ := listView_GetItemFocused.Call(obj)
    return ret
}

func ListView_SetItemFocused(obj uintptr, value uintptr) {
   listView_SetItemFocused.Call(obj, value)
}

func ListView_GetSelCount(obj uintptr) int32 {
    ret, _, _ := listView_GetSelCount.Call(obj)
    return int32(ret)
}

func ListView_GetSelected(obj uintptr) uintptr {
    ret, _, _ := listView_GetSelected.Call(obj)
    return ret
}

func ListView_SetSelected(obj uintptr, value uintptr) {
   listView_SetSelected.Call(obj, value)
}

func ListView_GetTopItem(obj uintptr) uintptr {
    ret, _, _ := listView_GetTopItem.Call(obj)
    return ret
}

func ListView_GetVisibleRowCount(obj uintptr) int32 {
    ret, _, _ := listView_GetVisibleRowCount.Call(obj)
    return int32(ret)
}

func ListView_GetItemIndex(obj uintptr) int32 {
    ret, _, _ := listView_GetItemIndex.Call(obj)
    return int32(ret)
}

func ListView_SetItemIndex(obj uintptr, value int32) {
   listView_SetItemIndex.Call(obj, uintptr(value))
}

func ListView_GetBrush(obj uintptr) uintptr {
    ret, _, _ := listView_GetBrush.Call(obj)
    return ret
}

func ListView_GetControlCount(obj uintptr) int32 {
    ret, _, _ := listView_GetControlCount.Call(obj)
    return int32(ret)
}

func ListView_GetHandle(obj uintptr) HWND {
    ret, _, _ := listView_GetHandle.Call(obj)
    return HWND(ret)
}

func ListView_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    listView_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ListView_SetBoundsRect(obj uintptr, value TRect) {
   listView_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ListView_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := listView_GetClientHeight.Call(obj)
    return int32(ret)
}

func ListView_SetClientHeight(obj uintptr, value int32) {
   listView_SetClientHeight.Call(obj, uintptr(value))
}

func ListView_GetClientRect(obj uintptr) TRect {
    var ret TRect
    listView_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ListView_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := listView_GetClientWidth.Call(obj)
    return int32(ret)
}

func ListView_SetClientWidth(obj uintptr, value int32) {
   listView_SetClientWidth.Call(obj, uintptr(value))
}

func ListView_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := listView_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func ListView_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := listView_GetExplicitTop.Call(obj)
    return int32(ret)
}

func ListView_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := listView_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func ListView_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := listView_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func ListView_GetParent(obj uintptr) uintptr {
    ret, _, _ := listView_GetParent.Call(obj)
    return ret
}

func ListView_SetParent(obj uintptr, value uintptr) {
   listView_SetParent.Call(obj, value)
}

func ListView_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := listView_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetAlignWithMargins(obj uintptr, value bool) {
   listView_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func ListView_GetLeft(obj uintptr) int32 {
    ret, _, _ := listView_GetLeft.Call(obj)
    return int32(ret)
}

func ListView_SetLeft(obj uintptr, value int32) {
   listView_SetLeft.Call(obj, uintptr(value))
}

func ListView_GetTop(obj uintptr) int32 {
    ret, _, _ := listView_GetTop.Call(obj)
    return int32(ret)
}

func ListView_SetTop(obj uintptr, value int32) {
   listView_SetTop.Call(obj, uintptr(value))
}

func ListView_GetWidth(obj uintptr) int32 {
    ret, _, _ := listView_GetWidth.Call(obj)
    return int32(ret)
}

func ListView_SetWidth(obj uintptr, value int32) {
   listView_SetWidth.Call(obj, uintptr(value))
}

func ListView_GetHeight(obj uintptr) int32 {
    ret, _, _ := listView_GetHeight.Call(obj)
    return int32(ret)
}

func ListView_SetHeight(obj uintptr, value int32) {
   listView_SetHeight.Call(obj, uintptr(value))
}

func ListView_GetCursor(obj uintptr) TCursor {
    ret, _, _ := listView_GetCursor.Call(obj)
    return TCursor(ret)
}

func ListView_SetCursor(obj uintptr, value TCursor) {
   listView_SetCursor.Call(obj, uintptr(value))
}

func ListView_GetHint(obj uintptr) string {
    ret, _, _ := listView_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func ListView_SetHint(obj uintptr, value string) {
   listView_SetHint.Call(obj, GoStrToDStr(value))
}

func ListView_GetMargins(obj uintptr) uintptr {
    ret, _, _ := listView_GetMargins.Call(obj)
    return ret
}

func ListView_SetMargins(obj uintptr, value uintptr) {
   listView_SetMargins.Call(obj, value)
}

func ListView_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := listView_GetComponentCount.Call(obj)
    return int32(ret)
}

func ListView_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := listView_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ListView_SetComponentIndex(obj uintptr, value int32) {
   listView_SetComponentIndex.Call(obj, uintptr(value))
}

func ListView_GetOwner(obj uintptr) uintptr {
    ret, _, _ := listView_GetOwner.Call(obj)
    return ret
}

func ListView_GetName(obj uintptr) string {
    ret, _, _ := listView_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ListView_SetName(obj uintptr, value string) {
   listView_SetName.Call(obj, GoStrToDStr(value))
}

func ListView_GetTag(obj uintptr) int {
    ret, _, _ := listView_GetTag.Call(obj)
    return int(ret)
}

func ListView_SetTag(obj uintptr, value int) {
   listView_SetTag.Call(obj, uintptr(value))
}

func ListView_GetColumn(obj uintptr, Index int32) uintptr {
    ret, _, _ := listView_GetColumn.Call(obj, uintptr(Index))
    return ret
}

func ListView_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := listView_GetControls.Call(obj, uintptr(Index))
    return ret
}

func ListView_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := listView_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}

