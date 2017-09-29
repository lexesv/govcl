
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

type TTreeView struct {
    IControl
    instance uintptr
}

func NewTreeView(owner IComponent) *TTreeView {
    t := new(TTreeView)
    t.instance = TreeView_Create(owner.Instance())
    return t
}

func TreeViewFromInst(inst uintptr) *TTreeView {
    t := new(TTreeView)
    t.instance = inst
    return t
}

func TreeViewFromObj(obj IObject) *TTreeView {
    t := new(TTreeView)
    t.instance = CheckPtr(obj)
    return t
}

func (t *TTreeView) Free() {
    if t.instance != 0 {
        TreeView_Free(t.instance)
    }
}

func (t *TTreeView) Instance() uintptr {
    return t.instance
}

func (t *TTreeView) IsValid() bool {
    return t.instance != 0
}

func (t *TTreeView) AlphaSort(ARecurse bool) bool {
    defer exceptionProc()
    return TreeView_AlphaSort(t.instance, ARecurse )
}

func (t *TTreeView) FullCollapse() {
    defer exceptionProc()
    TreeView_FullCollapse(t.instance)
}

func (t *TTreeView) FullExpand() {
    defer exceptionProc()
    TreeView_FullExpand(t.instance)
}

func (t *TTreeView) GetNodeAt(X int32, Y int32) *TTreeNode {
    defer exceptionProc()
    return TreeNodeFromInst(TreeView_GetNodeAt(t.instance, X , Y ))
}

func (t *TTreeView) IsEditing() bool {
    defer exceptionProc()
    return TreeView_IsEditing(t.instance)
}

func (t *TTreeView) LoadFromFile(FileName string) {
    defer exceptionProc()
    TreeView_LoadFromFile(t.instance, FileName )
}

func (t *TTreeView) LoadFromStream(Stream IObject) {
    defer exceptionProc()
    TreeView_LoadFromStream(t.instance, CheckPtr(Stream))
}

func (t *TTreeView) SaveToFile(FileName string) {
    defer exceptionProc()
    TreeView_SaveToFile(t.instance, FileName )
}

func (t *TTreeView) SaveToStream(Stream IObject) {
    defer exceptionProc()
    TreeView_SaveToStream(t.instance, CheckPtr(Stream))
}

func (t *TTreeView) Deselect(Node *TTreeNode) {
    defer exceptionProc()
    TreeView_Deselect(t.instance, CheckPtr(Node))
}

func (t *TTreeView) Subselect(Node *TTreeNode, Validate bool) {
    defer exceptionProc()
    TreeView_Subselect(t.instance, CheckPtr(Node), Validate )
}

func (t *TTreeView) ClearSelection(KeepPrimary bool) {
    defer exceptionProc()
    TreeView_ClearSelection(t.instance, KeepPrimary )
}

func (t *TTreeView) FindNextToSelect() *TTreeNode {
    defer exceptionProc()
    return TreeNodeFromInst(TreeView_FindNextToSelect(t.instance))
}

func (t *TTreeView) CanFocus() bool {
    defer exceptionProc()
    return TreeView_CanFocus(t.instance)
}

func (t *TTreeView) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    TreeView_FlipChildren(t.instance, AllLevels )
}

func (t *TTreeView) Focused() bool {
    defer exceptionProc()
    return TreeView_Focused(t.instance)
}

func (t *TTreeView) HandleAllocated() bool {
    defer exceptionProc()
    return TreeView_HandleAllocated(t.instance)
}

func (t *TTreeView) Invalidate() {
    defer exceptionProc()
    TreeView_Invalidate(t.instance)
}

func (t *TTreeView) Realign() {
    defer exceptionProc()
    TreeView_Realign(t.instance)
}

func (t *TTreeView) Repaint() {
    defer exceptionProc()
    TreeView_Repaint(t.instance)
}

func (t *TTreeView) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    TreeView_ScaleBy(t.instance, M , D )
}

func (t *TTreeView) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    TreeView_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight )
}

func (t *TTreeView) SetFocus() {
    defer exceptionProc()
    TreeView_SetFocus(t.instance)
}

func (t *TTreeView) Update() {
    defer exceptionProc()
    TreeView_Update(t.instance)
}

func (t *TTreeView) BringToFront() {
    defer exceptionProc()
    TreeView_BringToFront(t.instance)
}

func (t *TTreeView) HasParent() bool {
    defer exceptionProc()
    return TreeView_HasParent(t.instance)
}

func (t *TTreeView) Hide() {
    defer exceptionProc()
    TreeView_Hide(t.instance)
}

func (t *TTreeView) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return TreeView_Perform(t.instance, Msg , WParam , LParam )
}

func (t *TTreeView) Refresh() {
    defer exceptionProc()
    TreeView_Refresh(t.instance)
}

func (t *TTreeView) SendToBack() {
    defer exceptionProc()
    TreeView_SendToBack(t.instance)
}

func (t *TTreeView) Show() {
    defer exceptionProc()
    TreeView_Show(t.instance)
}

func (t *TTreeView) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return TreeView_GetTextBuf(t.instance, Buffer , BufSize )
}

func (t *TTreeView) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(TreeView_FindComponent(t.instance, AName ))
}

func (t *TTreeView) GetNamePath() string {
    defer exceptionProc()
    return TreeView_GetNamePath(t.instance)
}

func (t *TTreeView) Assign(Source IObject) {
    defer exceptionProc()
    TreeView_Assign(t.instance, CheckPtr(Source))
}

func (t *TTreeView) ClassName() string {
    defer exceptionProc()
    return TreeView_ClassName(t.instance)
}

func (t *TTreeView) Equals(Obj IObject) bool {
    defer exceptionProc()
    return TreeView_Equals(t.instance, CheckPtr(Obj))
}

func (t *TTreeView) GetHashCode() int32 {
    defer exceptionProc()
    return TreeView_GetHashCode(t.instance)
}

func (t *TTreeView) ToString() string {
    defer exceptionProc()
    return TreeView_ToString(t.instance)
}

func (t *TTreeView) Align() TAlign {
    defer exceptionProc()
    return TreeView_GetAlign(t.instance)
}

func (t *TTreeView) SetAlign(value TAlign) {
    defer exceptionProc()
    TreeView_SetAlign(t.instance, value)
}

func (t *TTreeView) Anchors() TAnchors {
    defer exceptionProc()
    return TreeView_GetAnchors(t.instance)
}

func (t *TTreeView) SetAnchors(value TAnchors) {
    defer exceptionProc()
    TreeView_SetAnchors(t.instance, value)
}

func (t *TTreeView) AutoExpand() bool {
    defer exceptionProc()
    return TreeView_GetAutoExpand(t.instance)
}

func (t *TTreeView) SetAutoExpand(value bool) {
    defer exceptionProc()
    TreeView_SetAutoExpand(t.instance, value)
}

func (t *TTreeView) BorderStyle() TBorderStyle {
    defer exceptionProc()
    return TreeView_GetBorderStyle(t.instance)
}

func (t *TTreeView) SetBorderStyle(value TBorderStyle) {
    defer exceptionProc()
    TreeView_SetBorderStyle(t.instance, value)
}

func (t *TTreeView) BorderWidth() int32 {
    defer exceptionProc()
    return TreeView_GetBorderWidth(t.instance)
}

func (t *TTreeView) SetBorderWidth(value int32) {
    defer exceptionProc()
    TreeView_SetBorderWidth(t.instance, value)
}

func (t *TTreeView) ChangeDelay() int32 {
    defer exceptionProc()
    return TreeView_GetChangeDelay(t.instance)
}

func (t *TTreeView) SetChangeDelay(value int32) {
    defer exceptionProc()
    TreeView_SetChangeDelay(t.instance, value)
}

func (t *TTreeView) Color() TColor {
    defer exceptionProc()
    return TreeView_GetColor(t.instance)
}

func (t *TTreeView) SetColor(value TColor) {
    defer exceptionProc()
    TreeView_SetColor(t.instance, value)
}

func (t *TTreeView) DoubleBuffered() bool {
    defer exceptionProc()
    return TreeView_GetDoubleBuffered(t.instance)
}

func (t *TTreeView) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    TreeView_SetDoubleBuffered(t.instance, value)
}

func (t *TTreeView) Enabled() bool {
    defer exceptionProc()
    return TreeView_GetEnabled(t.instance)
}

func (t *TTreeView) SetEnabled(value bool) {
    defer exceptionProc()
    TreeView_SetEnabled(t.instance, value)
}

func (t *TTreeView) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(TreeView_GetFont(t.instance))
}

func (t *TTreeView) SetFont(value *TFont) {
    defer exceptionProc()
    TreeView_SetFont(t.instance, CheckPtr(value))
}

func (t *TTreeView) HideSelection() bool {
    defer exceptionProc()
    return TreeView_GetHideSelection(t.instance)
}

func (t *TTreeView) SetHideSelection(value bool) {
    defer exceptionProc()
    TreeView_SetHideSelection(t.instance, value)
}

func (t *TTreeView) HotTrack() bool {
    defer exceptionProc()
    return TreeView_GetHotTrack(t.instance)
}

func (t *TTreeView) SetHotTrack(value bool) {
    defer exceptionProc()
    TreeView_SetHotTrack(t.instance, value)
}

func (t *TTreeView) Images() *TImageList {
    defer exceptionProc()
    return ImageListFromInst(TreeView_GetImages(t.instance))
}

func (t *TTreeView) SetImages(value IComponent) {
    defer exceptionProc()
    TreeView_SetImages(t.instance, CheckPtr(value))
}

func (t *TTreeView) Indent() int32 {
    defer exceptionProc()
    return TreeView_GetIndent(t.instance)
}

func (t *TTreeView) SetIndent(value int32) {
    defer exceptionProc()
    TreeView_SetIndent(t.instance, value)
}

func (t *TTreeView) MultiSelect() bool {
    defer exceptionProc()
    return TreeView_GetMultiSelect(t.instance)
}

func (t *TTreeView) SetMultiSelect(value bool) {
    defer exceptionProc()
    TreeView_SetMultiSelect(t.instance, value)
}

func (t *TTreeView) MultiSelectStyle() TMultiSelectStyle {
    defer exceptionProc()
    return TreeView_GetMultiSelectStyle(t.instance)
}

func (t *TTreeView) SetMultiSelectStyle(value TMultiSelectStyle) {
    defer exceptionProc()
    TreeView_SetMultiSelectStyle(t.instance, value)
}

func (t *TTreeView) ParentColor() bool {
    defer exceptionProc()
    return TreeView_GetParentColor(t.instance)
}

func (t *TTreeView) SetParentColor(value bool) {
    defer exceptionProc()
    TreeView_SetParentColor(t.instance, value)
}

func (t *TTreeView) ParentCtl3D() bool {
    defer exceptionProc()
    return TreeView_GetParentCtl3D(t.instance)
}

func (t *TTreeView) SetParentCtl3D(value bool) {
    defer exceptionProc()
    TreeView_SetParentCtl3D(t.instance, value)
}

func (t *TTreeView) ParentFont() bool {
    defer exceptionProc()
    return TreeView_GetParentFont(t.instance)
}

func (t *TTreeView) SetParentFont(value bool) {
    defer exceptionProc()
    TreeView_SetParentFont(t.instance, value)
}

func (t *TTreeView) ParentShowHint() bool {
    defer exceptionProc()
    return TreeView_GetParentShowHint(t.instance)
}

func (t *TTreeView) SetParentShowHint(value bool) {
    defer exceptionProc()
    TreeView_SetParentShowHint(t.instance, value)
}

func (t *TTreeView) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(TreeView_GetPopupMenu(t.instance))
}

func (t *TTreeView) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    TreeView_SetPopupMenu(t.instance, CheckPtr(value))
}

func (t *TTreeView) ReadOnly() bool {
    defer exceptionProc()
    return TreeView_GetReadOnly(t.instance)
}

func (t *TTreeView) SetReadOnly(value bool) {
    defer exceptionProc()
    TreeView_SetReadOnly(t.instance, value)
}

func (t *TTreeView) RightClickSelect() bool {
    defer exceptionProc()
    return TreeView_GetRightClickSelect(t.instance)
}

func (t *TTreeView) SetRightClickSelect(value bool) {
    defer exceptionProc()
    TreeView_SetRightClickSelect(t.instance, value)
}

func (t *TTreeView) RowSelect() bool {
    defer exceptionProc()
    return TreeView_GetRowSelect(t.instance)
}

func (t *TTreeView) SetRowSelect(value bool) {
    defer exceptionProc()
    TreeView_SetRowSelect(t.instance, value)
}

func (t *TTreeView) ShowButtons() bool {
    defer exceptionProc()
    return TreeView_GetShowButtons(t.instance)
}

func (t *TTreeView) SetShowButtons(value bool) {
    defer exceptionProc()
    TreeView_SetShowButtons(t.instance, value)
}

func (t *TTreeView) ShowHint() bool {
    defer exceptionProc()
    return TreeView_GetShowHint(t.instance)
}

func (t *TTreeView) SetShowHint(value bool) {
    defer exceptionProc()
    TreeView_SetShowHint(t.instance, value)
}

func (t *TTreeView) ShowLines() bool {
    defer exceptionProc()
    return TreeView_GetShowLines(t.instance)
}

func (t *TTreeView) SetShowLines(value bool) {
    defer exceptionProc()
    TreeView_SetShowLines(t.instance, value)
}

func (t *TTreeView) ShowRoot() bool {
    defer exceptionProc()
    return TreeView_GetShowRoot(t.instance)
}

func (t *TTreeView) SetShowRoot(value bool) {
    defer exceptionProc()
    TreeView_SetShowRoot(t.instance, value)
}

func (t *TTreeView) SortType() TSortType {
    defer exceptionProc()
    return TreeView_GetSortType(t.instance)
}

func (t *TTreeView) SetSortType(value TSortType) {
    defer exceptionProc()
    TreeView_SetSortType(t.instance, value)
}

func (t *TTreeView) StateImages() *TImageList {
    defer exceptionProc()
    return ImageListFromInst(TreeView_GetStateImages(t.instance))
}

func (t *TTreeView) SetStateImages(value IComponent) {
    defer exceptionProc()
    TreeView_SetStateImages(t.instance, CheckPtr(value))
}

func (t *TTreeView) TabOrder() int16 {
    defer exceptionProc()
    return TreeView_GetTabOrder(t.instance)
}

func (t *TTreeView) SetTabOrder(value int16) {
    defer exceptionProc()
    TreeView_SetTabOrder(t.instance, value)
}

func (t *TTreeView) TabStop() bool {
    defer exceptionProc()
    return TreeView_GetTabStop(t.instance)
}

func (t *TTreeView) SetTabStop(value bool) {
    defer exceptionProc()
    TreeView_SetTabStop(t.instance, value)
}

func (t *TTreeView) ToolTips() bool {
    defer exceptionProc()
    return TreeView_GetToolTips(t.instance)
}

func (t *TTreeView) SetToolTips(value bool) {
    defer exceptionProc()
    TreeView_SetToolTips(t.instance, value)
}

func (t *TTreeView) Visible() bool {
    defer exceptionProc()
    return TreeView_GetVisible(t.instance)
}

func (t *TTreeView) SetVisible(value bool) {
    defer exceptionProc()
    TreeView_SetVisible(t.instance, value)
}

func (t *TTreeView) SetOnChange(fn TTVChangedEvent) {
    TreeView_SetOnChange(t.instance, fn)
}

func (t *TTreeView) SetOnClick(fn TNotifyEvent) {
    TreeView_SetOnClick(t.instance, fn)
}

func (t *TTreeView) SetOnDblClick(fn TNotifyEvent) {
    TreeView_SetOnDblClick(t.instance, fn)
}

func (t *TTreeView) SetOnEnter(fn TNotifyEvent) {
    TreeView_SetOnEnter(t.instance, fn)
}

func (t *TTreeView) SetOnExit(fn TNotifyEvent) {
    TreeView_SetOnExit(t.instance, fn)
}

func (t *TTreeView) SetOnKeyDown(fn TKeyEvent) {
    TreeView_SetOnKeyDown(t.instance, fn)
}

func (t *TTreeView) SetOnKeyPress(fn TKeyPressEvent) {
    TreeView_SetOnKeyPress(t.instance, fn)
}

func (t *TTreeView) SetOnKeyUp(fn TKeyEvent) {
    TreeView_SetOnKeyUp(t.instance, fn)
}

func (t *TTreeView) SetOnMouseDown(fn TMouseEvent) {
    TreeView_SetOnMouseDown(t.instance, fn)
}

func (t *TTreeView) SetOnMouseEnter(fn TNotifyEvent) {
    TreeView_SetOnMouseEnter(t.instance, fn)
}

func (t *TTreeView) SetOnMouseLeave(fn TNotifyEvent) {
    TreeView_SetOnMouseLeave(t.instance, fn)
}

func (t *TTreeView) SetOnMouseMove(fn TMouseMoveEvent) {
    TreeView_SetOnMouseMove(t.instance, fn)
}

func (t *TTreeView) SetOnMouseUp(fn TMouseEvent) {
    TreeView_SetOnMouseUp(t.instance, fn)
}

func (t *TTreeView) Items() *TTreeNodes {
    defer exceptionProc()
    return TreeNodesFromInst(TreeView_GetItems(t.instance))
}

func (t *TTreeView) SetItems(value *TTreeNodes) {
    defer exceptionProc()
    TreeView_SetItems(t.instance, CheckPtr(value))
}

func (t *TTreeView) Canvas() *TCanvas {
    defer exceptionProc()
    return CanvasFromInst(TreeView_GetCanvas(t.instance))
}

func (t *TTreeView) DropTarget() *TTreeNode {
    defer exceptionProc()
    return TreeNodeFromInst(TreeView_GetDropTarget(t.instance))
}

func (t *TTreeView) SetDropTarget(value *TTreeNode) {
    defer exceptionProc()
    TreeView_SetDropTarget(t.instance, CheckPtr(value))
}

func (t *TTreeView) Selected() *TTreeNode {
    defer exceptionProc()
    return TreeNodeFromInst(TreeView_GetSelected(t.instance))
}

func (t *TTreeView) SetSelected(value *TTreeNode) {
    defer exceptionProc()
    TreeView_SetSelected(t.instance, CheckPtr(value))
}

func (t *TTreeView) TopItem() *TTreeNode {
    defer exceptionProc()
    return TreeNodeFromInst(TreeView_GetTopItem(t.instance))
}

func (t *TTreeView) SetTopItem(value *TTreeNode) {
    defer exceptionProc()
    TreeView_SetTopItem(t.instance, CheckPtr(value))
}

func (t *TTreeView) SelectionCount() uint32 {
    defer exceptionProc()
    return TreeView_GetSelectionCount(t.instance)
}

func (t *TTreeView) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(TreeView_GetBrush(t.instance))
}

func (t *TTreeView) ControlCount() int32 {
    defer exceptionProc()
    return TreeView_GetControlCount(t.instance)
}

func (t *TTreeView) Handle() HWND {
    defer exceptionProc()
    return TreeView_GetHandle(t.instance)
}

func (t *TTreeView) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(TreeView_GetAction(t.instance))
}

func (t *TTreeView) SetAction(value IComponent) {
    defer exceptionProc()
    TreeView_SetAction(t.instance, CheckPtr(value))
}

func (t *TTreeView) BoundsRect() TRect {
    defer exceptionProc()
    return TreeView_GetBoundsRect(t.instance)
}

func (t *TTreeView) SetBoundsRect(value TRect) {
    defer exceptionProc()
    TreeView_SetBoundsRect(t.instance, value)
}

func (t *TTreeView) ClientHeight() int32 {
    defer exceptionProc()
    return TreeView_GetClientHeight(t.instance)
}

func (t *TTreeView) SetClientHeight(value int32) {
    defer exceptionProc()
    TreeView_SetClientHeight(t.instance, value)
}

func (t *TTreeView) ClientRect() TRect {
    defer exceptionProc()
    return TreeView_GetClientRect(t.instance)
}

func (t *TTreeView) ClientWidth() int32 {
    defer exceptionProc()
    return TreeView_GetClientWidth(t.instance)
}

func (t *TTreeView) SetClientWidth(value int32) {
    defer exceptionProc()
    TreeView_SetClientWidth(t.instance, value)
}

func (t *TTreeView) ExplicitLeft() int32 {
    defer exceptionProc()
    return TreeView_GetExplicitLeft(t.instance)
}

func (t *TTreeView) ExplicitTop() int32 {
    defer exceptionProc()
    return TreeView_GetExplicitTop(t.instance)
}

func (t *TTreeView) ExplicitWidth() int32 {
    defer exceptionProc()
    return TreeView_GetExplicitWidth(t.instance)
}

func (t *TTreeView) ExplicitHeight() int32 {
    defer exceptionProc()
    return TreeView_GetExplicitHeight(t.instance)
}

func (t *TTreeView) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(TreeView_GetParent(t.instance))
}

func (t *TTreeView) SetParent(value IControl) {
    defer exceptionProc()
    TreeView_SetParent(t.instance, CheckPtr(value))
}

func (t *TTreeView) AlignWithMargins() bool {
    defer exceptionProc()
    return TreeView_GetAlignWithMargins(t.instance)
}

func (t *TTreeView) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    TreeView_SetAlignWithMargins(t.instance, value)
}

func (t *TTreeView) Left() int32 {
    defer exceptionProc()
    return TreeView_GetLeft(t.instance)
}

func (t *TTreeView) SetLeft(value int32) {
    defer exceptionProc()
    TreeView_SetLeft(t.instance, value)
}

func (t *TTreeView) Top() int32 {
    defer exceptionProc()
    return TreeView_GetTop(t.instance)
}

func (t *TTreeView) SetTop(value int32) {
    defer exceptionProc()
    TreeView_SetTop(t.instance, value)
}

func (t *TTreeView) Width() int32 {
    defer exceptionProc()
    return TreeView_GetWidth(t.instance)
}

func (t *TTreeView) SetWidth(value int32) {
    defer exceptionProc()
    TreeView_SetWidth(t.instance, value)
}

func (t *TTreeView) Height() int32 {
    defer exceptionProc()
    return TreeView_GetHeight(t.instance)
}

func (t *TTreeView) SetHeight(value int32) {
    defer exceptionProc()
    TreeView_SetHeight(t.instance, value)
}

func (t *TTreeView) Cursor() TCursor {
    defer exceptionProc()
    return TreeView_GetCursor(t.instance)
}

func (t *TTreeView) SetCursor(value TCursor) {
    defer exceptionProc()
    TreeView_SetCursor(t.instance, value)
}

func (t *TTreeView) Hint() string {
    defer exceptionProc()
    return TreeView_GetHint(t.instance)
}

func (t *TTreeView) SetHint(value string) {
    defer exceptionProc()
    TreeView_SetHint(t.instance, value)
}

func (t *TTreeView) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(TreeView_GetMargins(t.instance))
}

func (t *TTreeView) SetMargins(value *TMargins) {
    defer exceptionProc()
    TreeView_SetMargins(t.instance, CheckPtr(value))
}

func (t *TTreeView) ComponentCount() int32 {
    defer exceptionProc()
    return TreeView_GetComponentCount(t.instance)
}

func (t *TTreeView) ComponentIndex() int32 {
    defer exceptionProc()
    return TreeView_GetComponentIndex(t.instance)
}

func (t *TTreeView) SetComponentIndex(value int32) {
    defer exceptionProc()
    TreeView_SetComponentIndex(t.instance, value)
}

func (t *TTreeView) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(TreeView_GetOwner(t.instance))
}

func (t *TTreeView) Name() string {
    defer exceptionProc()
    return TreeView_GetName(t.instance)
}

func (t *TTreeView) SetName(value string) {
    defer exceptionProc()
    TreeView_SetName(t.instance, value)
}

func (t *TTreeView) Tag() int {
    defer exceptionProc()
    return TreeView_GetTag(t.instance)
}

func (t *TTreeView) SetTag(value int) {
    defer exceptionProc()
    TreeView_SetTag(t.instance, value)
}

func (t *TTreeView) Selections(Index int32) *TTreeNode {
    defer exceptionProc()
    return TreeNodeFromInst(TreeView_GetSelections(t.instance, Index))
}

func (t *TTreeView) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(TreeView_GetControls(t.instance, Index))
}

func (t *TTreeView) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(TreeView_GetComponents(t.instance, AIndex))
}

