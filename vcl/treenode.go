
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
    . "gitee.com/ying32/govcl/vcl/types"
)

type TTreeNode struct {
    IObject
    instance uintptr
}

func NewTreeNode() *TTreeNode {
    t := new(TTreeNode)
    t.instance = TreeNode_Create()
    return t
}

func TreeNodeFromInst(inst uintptr) *TTreeNode {
    t := new(TTreeNode)
    t.instance = inst
    return t
}

func TreeNodeFromObj(obj IObject) *TTreeNode {
    t := new(TTreeNode)
    t.instance = CheckPtr(obj)
    return t
}

func (t *TTreeNode) Free() {
    if t.instance != 0 {
        TreeNode_Free(t.instance)
    }
}

func (t *TTreeNode) Instance() uintptr {
    return t.instance
}

func (t *TTreeNode) IsValid() bool {
    return t.instance != 0
}

func (t *TTreeNode) AlphaSort(ARecurse bool) bool {
    defer exceptionProc()
    return TreeNode_AlphaSort(t.instance, ARecurse )
}

func (t *TTreeNode) Assign(Source IObject) {
    defer exceptionProc()
    TreeNode_Assign(t.instance, CheckPtr(Source))
}

func (t *TTreeNode) Delete() {
    defer exceptionProc()
    TreeNode_Delete(t.instance)
}

func (t *TTreeNode) IndexOf(Value *TTreeNode) int32 {
    defer exceptionProc()
    return TreeNode_IndexOf(t.instance, CheckPtr(Value))
}

func (t *TTreeNode) MakeVisible() {
    defer exceptionProc()
    TreeNode_MakeVisible(t.instance)
}

func (t *TTreeNode) MoveTo(Destination *TTreeNode, Mode TNodeAttachMode) {
    defer exceptionProc()
    TreeNode_MoveTo(t.instance, CheckPtr(Destination), Mode )
}

func (t *TTreeNode) GetNamePath() string {
    defer exceptionProc()
    return TreeNode_GetNamePath(t.instance)
}

func (t *TTreeNode) ClassName() string {
    defer exceptionProc()
    return TreeNode_ClassName(t.instance)
}

func (t *TTreeNode) Equals(Obj IObject) bool {
    defer exceptionProc()
    return TreeNode_Equals(t.instance, CheckPtr(Obj))
}

func (t *TTreeNode) GetHashCode() int32 {
    defer exceptionProc()
    return TreeNode_GetHashCode(t.instance)
}

func (t *TTreeNode) ToString() string {
    defer exceptionProc()
    return TreeNode_ToString(t.instance)
}

func (t *TTreeNode) AbsoluteIndex() int32 {
    defer exceptionProc()
    return TreeNode_GetAbsoluteIndex(t.instance)
}

func (t *TTreeNode) Cut() bool {
    defer exceptionProc()
    return TreeNode_GetCut(t.instance)
}

func (t *TTreeNode) SetCut(value bool) {
    defer exceptionProc()
    TreeNode_SetCut(t.instance, value)
}

func (t *TTreeNode) Data() uintptr {
    defer exceptionProc()
    return TreeNode_GetData(t.instance)
}

func (t *TTreeNode) SetData(value uintptr) {
    defer exceptionProc()
    TreeNode_SetData(t.instance, value)
}

func (t *TTreeNode) Deleting() bool {
    defer exceptionProc()
    return TreeNode_GetDeleting(t.instance)
}

func (t *TTreeNode) Focused() bool {
    defer exceptionProc()
    return TreeNode_GetFocused(t.instance)
}

func (t *TTreeNode) SetFocused(value bool) {
    defer exceptionProc()
    TreeNode_SetFocused(t.instance, value)
}

func (t *TTreeNode) DropTarget() bool {
    defer exceptionProc()
    return TreeNode_GetDropTarget(t.instance)
}

func (t *TTreeNode) SetDropTarget(value bool) {
    defer exceptionProc()
    TreeNode_SetDropTarget(t.instance, value)
}

func (t *TTreeNode) Selected() bool {
    defer exceptionProc()
    return TreeNode_GetSelected(t.instance)
}

func (t *TTreeNode) SetSelected(value bool) {
    defer exceptionProc()
    TreeNode_SetSelected(t.instance, value)
}

func (t *TTreeNode) Expanded() bool {
    defer exceptionProc()
    return TreeNode_GetExpanded(t.instance)
}

func (t *TTreeNode) SetExpanded(value bool) {
    defer exceptionProc()
    TreeNode_SetExpanded(t.instance, value)
}

func (t *TTreeNode) ExpandedImageIndex() int32 {
    defer exceptionProc()
    return TreeNode_GetExpandedImageIndex(t.instance)
}

func (t *TTreeNode) SetExpandedImageIndex(value int32) {
    defer exceptionProc()
    TreeNode_SetExpandedImageIndex(t.instance, value)
}

func (t *TTreeNode) Handle() HWND {
    defer exceptionProc()
    return TreeNode_GetHandle(t.instance)
}

func (t *TTreeNode) HasChildren() bool {
    defer exceptionProc()
    return TreeNode_GetHasChildren(t.instance)
}

func (t *TTreeNode) SetHasChildren(value bool) {
    defer exceptionProc()
    TreeNode_SetHasChildren(t.instance, value)
}

func (t *TTreeNode) ImageIndex() int32 {
    defer exceptionProc()
    return TreeNode_GetImageIndex(t.instance)
}

func (t *TTreeNode) SetImageIndex(value int32) {
    defer exceptionProc()
    TreeNode_SetImageIndex(t.instance, value)
}

func (t *TTreeNode) Index() int32 {
    defer exceptionProc()
    return TreeNode_GetIndex(t.instance)
}

func (t *TTreeNode) IsVisible() bool {
    defer exceptionProc()
    return TreeNode_GetIsVisible(t.instance)
}

func (t *TTreeNode) ItemId() uintptr {
    defer exceptionProc()
    return TreeNode_GetItemId(t.instance)
}

func (t *TTreeNode) Level() int32 {
    defer exceptionProc()
    return TreeNode_GetLevel(t.instance)
}

func (t *TTreeNode) OverlayIndex() int32 {
    defer exceptionProc()
    return TreeNode_GetOverlayIndex(t.instance)
}

func (t *TTreeNode) SetOverlayIndex(value int32) {
    defer exceptionProc()
    TreeNode_SetOverlayIndex(t.instance, value)
}

func (t *TTreeNode) Owner() *TTreeNodes {
    defer exceptionProc()
    return TreeNodesFromInst(TreeNode_GetOwner(t.instance))
}

func (t *TTreeNode) Parent() *TTreeNode {
    defer exceptionProc()
    return TreeNodeFromInst(TreeNode_GetParent(t.instance))
}

func (t *TTreeNode) SelectedIndex() int32 {
    defer exceptionProc()
    return TreeNode_GetSelectedIndex(t.instance)
}

func (t *TTreeNode) SetSelectedIndex(value int32) {
    defer exceptionProc()
    TreeNode_SetSelectedIndex(t.instance, value)
}

func (t *TTreeNode) Enabled() bool {
    defer exceptionProc()
    return TreeNode_GetEnabled(t.instance)
}

func (t *TTreeNode) SetEnabled(value bool) {
    defer exceptionProc()
    TreeNode_SetEnabled(t.instance, value)
}

func (t *TTreeNode) StateIndex() int32 {
    defer exceptionProc()
    return TreeNode_GetStateIndex(t.instance)
}

func (t *TTreeNode) SetStateIndex(value int32) {
    defer exceptionProc()
    TreeNode_SetStateIndex(t.instance, value)
}

func (t *TTreeNode) Text() string {
    defer exceptionProc()
    return TreeNode_GetText(t.instance)
}

func (t *TTreeNode) SetText(value string) {
    defer exceptionProc()
    TreeNode_SetText(t.instance, value)
}

func (t *TTreeNode) Item(Index int32) *TTreeNode {
    defer exceptionProc()
    return TreeNodeFromInst(TreeNode_GetItem(t.instance, Index))
}

func (t *TTreeNode) SetItem(Index int32, value *TTreeNode) {
    defer exceptionProc()
    TreeNode_SetItem(t.instance, Index, CheckPtr(value))
}

