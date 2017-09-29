
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func TreeNode_Create() uintptr {
    ret, _, _ := treeNode_Create.Call()
    return ret
}

func TreeNode_Free(obj uintptr) {
    treeNode_Free.Call(obj)
}

func TreeNode_AlphaSort(obj uintptr, ARecurse bool) bool {
    ret, _, _ := treeNode_AlphaSort.Call(obj, GoBoolToDBool(ARecurse) )
    return DBoolToGoBool(ret)
}

func TreeNode_Assign(obj uintptr, Source uintptr)  {
    treeNode_Assign.Call(obj, Source )
}

func TreeNode_Delete(obj uintptr)  {
    treeNode_Delete.Call(obj)
}

func TreeNode_IndexOf(obj uintptr, Value uintptr) int32 {
    ret, _, _ := treeNode_IndexOf.Call(obj, Value )
    return int32(ret)
}

func TreeNode_MakeVisible(obj uintptr)  {
    treeNode_MakeVisible.Call(obj)
}

func TreeNode_MoveTo(obj uintptr, Destination uintptr, Mode TNodeAttachMode)  {
    treeNode_MoveTo.Call(obj, Destination , uintptr(Mode) )
}

func TreeNode_GetNamePath(obj uintptr) string {
    ret, _, _ := treeNode_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func TreeNode_ClassName(obj uintptr) string {
    ret, _, _ := treeNode_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func TreeNode_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := treeNode_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func TreeNode_GetHashCode(obj uintptr) int32 {
    ret, _, _ := treeNode_GetHashCode.Call(obj)
    return int32(ret)
}

func TreeNode_ToString(obj uintptr) string {
    ret, _, _ := treeNode_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func TreeNode_GetAbsoluteIndex(obj uintptr) int32 {
    ret, _, _ := treeNode_GetAbsoluteIndex.Call(obj)
    return int32(ret)
}

func TreeNode_GetCut(obj uintptr) bool {
    ret, _, _ := treeNode_GetCut.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeNode_SetCut(obj uintptr, value bool) {
   treeNode_SetCut.Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetData(obj uintptr) uintptr {
    ret, _, _ := treeNode_GetData.Call(obj)
    return ret
}

func TreeNode_SetData(obj uintptr, value uintptr) {
   treeNode_SetData.Call(obj, value)
}

func TreeNode_GetDeleting(obj uintptr) bool {
    ret, _, _ := treeNode_GetDeleting.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeNode_GetFocused(obj uintptr) bool {
    ret, _, _ := treeNode_GetFocused.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeNode_SetFocused(obj uintptr, value bool) {
   treeNode_SetFocused.Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetDropTarget(obj uintptr) bool {
    ret, _, _ := treeNode_GetDropTarget.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeNode_SetDropTarget(obj uintptr, value bool) {
   treeNode_SetDropTarget.Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetSelected(obj uintptr) bool {
    ret, _, _ := treeNode_GetSelected.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeNode_SetSelected(obj uintptr, value bool) {
   treeNode_SetSelected.Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetExpanded(obj uintptr) bool {
    ret, _, _ := treeNode_GetExpanded.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeNode_SetExpanded(obj uintptr, value bool) {
   treeNode_SetExpanded.Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetExpandedImageIndex(obj uintptr) int32 {
    ret, _, _ := treeNode_GetExpandedImageIndex.Call(obj)
    return int32(ret)
}

func TreeNode_SetExpandedImageIndex(obj uintptr, value int32) {
   treeNode_SetExpandedImageIndex.Call(obj, uintptr(value))
}

func TreeNode_GetHandle(obj uintptr) HWND {
    ret, _, _ := treeNode_GetHandle.Call(obj)
    return HWND(ret)
}

func TreeNode_GetHasChildren(obj uintptr) bool {
    ret, _, _ := treeNode_GetHasChildren.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeNode_SetHasChildren(obj uintptr, value bool) {
   treeNode_SetHasChildren.Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetImageIndex(obj uintptr) int32 {
    ret, _, _ := treeNode_GetImageIndex.Call(obj)
    return int32(ret)
}

func TreeNode_SetImageIndex(obj uintptr, value int32) {
   treeNode_SetImageIndex.Call(obj, uintptr(value))
}

func TreeNode_GetIndex(obj uintptr) int32 {
    ret, _, _ := treeNode_GetIndex.Call(obj)
    return int32(ret)
}

func TreeNode_GetIsVisible(obj uintptr) bool {
    ret, _, _ := treeNode_GetIsVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeNode_GetItemId(obj uintptr) uintptr {
    ret, _, _ := treeNode_GetItemId.Call(obj)
    return ret
}

func TreeNode_GetLevel(obj uintptr) int32 {
    ret, _, _ := treeNode_GetLevel.Call(obj)
    return int32(ret)
}

func TreeNode_GetOverlayIndex(obj uintptr) int32 {
    ret, _, _ := treeNode_GetOverlayIndex.Call(obj)
    return int32(ret)
}

func TreeNode_SetOverlayIndex(obj uintptr, value int32) {
   treeNode_SetOverlayIndex.Call(obj, uintptr(value))
}

func TreeNode_GetOwner(obj uintptr) uintptr {
    ret, _, _ := treeNode_GetOwner.Call(obj)
    return ret
}

func TreeNode_GetParent(obj uintptr) uintptr {
    ret, _, _ := treeNode_GetParent.Call(obj)
    return ret
}

func TreeNode_GetSelectedIndex(obj uintptr) int32 {
    ret, _, _ := treeNode_GetSelectedIndex.Call(obj)
    return int32(ret)
}

func TreeNode_SetSelectedIndex(obj uintptr, value int32) {
   treeNode_SetSelectedIndex.Call(obj, uintptr(value))
}

func TreeNode_GetEnabled(obj uintptr) bool {
    ret, _, _ := treeNode_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeNode_SetEnabled(obj uintptr, value bool) {
   treeNode_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetStateIndex(obj uintptr) int32 {
    ret, _, _ := treeNode_GetStateIndex.Call(obj)
    return int32(ret)
}

func TreeNode_SetStateIndex(obj uintptr, value int32) {
   treeNode_SetStateIndex.Call(obj, uintptr(value))
}

func TreeNode_GetText(obj uintptr) string {
    ret, _, _ := treeNode_GetText.Call(obj)
    return DStrToGoStr(ret)
}

func TreeNode_SetText(obj uintptr, value string) {
   treeNode_SetText.Call(obj, GoStrToDStr(value))
}

func TreeNode_GetItem(obj uintptr, Index int32) uintptr {
    ret, _, _ := treeNode_GetItem.Call(obj, uintptr(Index))
    return ret
}

func TreeNode_SetItem(obj uintptr, Index int32, value uintptr) {
   treeNode_SetItem.Call(obj, uintptr(Index), value)
}

