
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

type TListView struct {
    IControl
    instance uintptr
}

func NewListView(owner IComponent) *TListView {
    l := new(TListView)
    l.instance = ListView_Create(owner.Instance())
    return l
}

func ListViewFromInst(inst uintptr) *TListView {
    l := new(TListView)
    l.instance = inst
    return l
}

func ListViewFromObj(obj IObject) *TListView {
    l := new(TListView)
    l.instance = CheckPtr(obj)
    return l
}

func (l *TListView) Free() {
    if l.instance != 0 {
        ListView_Free(l.instance)
    }
}

func (l *TListView) Instance() uintptr {
    return l.instance
}

func (l *TListView) IsValid() bool {
    return l.instance != 0
}

func (l *TListView) AddItem(Item string, AObject IObject) {
    defer exceptionProc()
    ListView_AddItem(l.instance, Item , CheckPtr(AObject))
}

func (l *TListView) AlphaSort() bool {
    defer exceptionProc()
    return ListView_AlphaSort(l.instance)
}

func (l *TListView) Clear() {
    defer exceptionProc()
    ListView_Clear(l.instance)
}

func (l *TListView) ClearSelection() {
    defer exceptionProc()
    ListView_ClearSelection(l.instance)
}

func (l *TListView) DeleteSelected() {
    defer exceptionProc()
    ListView_DeleteSelected(l.instance)
}

func (l *TListView) GetSearchString() string {
    defer exceptionProc()
    return ListView_GetSearchString(l.instance)
}

func (l *TListView) IsEditing() bool {
    defer exceptionProc()
    return ListView_IsEditing(l.instance)
}

func (l *TListView) SelectAll() {
    defer exceptionProc()
    ListView_SelectAll(l.instance)
}

func (l *TListView) Scroll(DX int32, DY int32) {
    defer exceptionProc()
    ListView_Scroll(l.instance, DX , DY )
}

func (l *TListView) CanFocus() bool {
    defer exceptionProc()
    return ListView_CanFocus(l.instance)
}

func (l *TListView) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    ListView_FlipChildren(l.instance, AllLevels )
}

func (l *TListView) Focused() bool {
    defer exceptionProc()
    return ListView_Focused(l.instance)
}

func (l *TListView) HandleAllocated() bool {
    defer exceptionProc()
    return ListView_HandleAllocated(l.instance)
}

func (l *TListView) Invalidate() {
    defer exceptionProc()
    ListView_Invalidate(l.instance)
}

func (l *TListView) Realign() {
    defer exceptionProc()
    ListView_Realign(l.instance)
}

func (l *TListView) Repaint() {
    defer exceptionProc()
    ListView_Repaint(l.instance)
}

func (l *TListView) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    ListView_ScaleBy(l.instance, M , D )
}

func (l *TListView) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    ListView_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight )
}

func (l *TListView) SetFocus() {
    defer exceptionProc()
    ListView_SetFocus(l.instance)
}

func (l *TListView) Update() {
    defer exceptionProc()
    ListView_Update(l.instance)
}

func (l *TListView) BringToFront() {
    defer exceptionProc()
    ListView_BringToFront(l.instance)
}

func (l *TListView) HasParent() bool {
    defer exceptionProc()
    return ListView_HasParent(l.instance)
}

func (l *TListView) Hide() {
    defer exceptionProc()
    ListView_Hide(l.instance)
}

func (l *TListView) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return ListView_Perform(l.instance, Msg , WParam , LParam )
}

func (l *TListView) Refresh() {
    defer exceptionProc()
    ListView_Refresh(l.instance)
}

func (l *TListView) SendToBack() {
    defer exceptionProc()
    ListView_SendToBack(l.instance)
}

func (l *TListView) Show() {
    defer exceptionProc()
    ListView_Show(l.instance)
}

func (l *TListView) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return ListView_GetTextBuf(l.instance, Buffer , BufSize )
}

func (l *TListView) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ListView_FindComponent(l.instance, AName ))
}

func (l *TListView) GetNamePath() string {
    defer exceptionProc()
    return ListView_GetNamePath(l.instance)
}

func (l *TListView) Assign(Source IObject) {
    defer exceptionProc()
    ListView_Assign(l.instance, CheckPtr(Source))
}

func (l *TListView) ClassName() string {
    defer exceptionProc()
    return ListView_ClassName(l.instance)
}

func (l *TListView) Equals(Obj IObject) bool {
    defer exceptionProc()
    return ListView_Equals(l.instance, CheckPtr(Obj))
}

func (l *TListView) GetHashCode() int32 {
    defer exceptionProc()
    return ListView_GetHashCode(l.instance)
}

func (l *TListView) ToString() string {
    defer exceptionProc()
    return ListView_ToString(l.instance)
}

func (l *TListView) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(ListView_GetAction(l.instance))
}

func (l *TListView) SetAction(value IComponent) {
    defer exceptionProc()
    ListView_SetAction(l.instance, CheckPtr(value))
}

func (l *TListView) Align() TAlign {
    defer exceptionProc()
    return ListView_GetAlign(l.instance)
}

func (l *TListView) SetAlign(value TAlign) {
    defer exceptionProc()
    ListView_SetAlign(l.instance, value)
}

func (l *TListView) AllocBy() int32 {
    defer exceptionProc()
    return ListView_GetAllocBy(l.instance)
}

func (l *TListView) SetAllocBy(value int32) {
    defer exceptionProc()
    ListView_SetAllocBy(l.instance, value)
}

func (l *TListView) Anchors() TAnchors {
    defer exceptionProc()
    return ListView_GetAnchors(l.instance)
}

func (l *TListView) SetAnchors(value TAnchors) {
    defer exceptionProc()
    ListView_SetAnchors(l.instance, value)
}

func (l *TListView) BorderStyle() TBorderStyle {
    defer exceptionProc()
    return ListView_GetBorderStyle(l.instance)
}

func (l *TListView) SetBorderStyle(value TBorderStyle) {
    defer exceptionProc()
    ListView_SetBorderStyle(l.instance, value)
}

func (l *TListView) BorderWidth() int32 {
    defer exceptionProc()
    return ListView_GetBorderWidth(l.instance)
}

func (l *TListView) SetBorderWidth(value int32) {
    defer exceptionProc()
    ListView_SetBorderWidth(l.instance, value)
}

func (l *TListView) Checkboxes() bool {
    defer exceptionProc()
    return ListView_GetCheckboxes(l.instance)
}

func (l *TListView) SetCheckboxes(value bool) {
    defer exceptionProc()
    ListView_SetCheckboxes(l.instance, value)
}

func (l *TListView) Color() TColor {
    defer exceptionProc()
    return ListView_GetColor(l.instance)
}

func (l *TListView) SetColor(value TColor) {
    defer exceptionProc()
    ListView_SetColor(l.instance, value)
}

func (l *TListView) Columns() *TListColumns {
    defer exceptionProc()
    return ListColumnsFromInst(ListView_GetColumns(l.instance))
}

func (l *TListView) SetColumns(value *TListColumns) {
    defer exceptionProc()
    ListView_SetColumns(l.instance, CheckPtr(value))
}

func (l *TListView) ColumnClick() bool {
    defer exceptionProc()
    return ListView_GetColumnClick(l.instance)
}

func (l *TListView) SetColumnClick(value bool) {
    defer exceptionProc()
    ListView_SetColumnClick(l.instance, value)
}

func (l *TListView) DoubleBuffered() bool {
    defer exceptionProc()
    return ListView_GetDoubleBuffered(l.instance)
}

func (l *TListView) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    ListView_SetDoubleBuffered(l.instance, value)
}

func (l *TListView) Enabled() bool {
    defer exceptionProc()
    return ListView_GetEnabled(l.instance)
}

func (l *TListView) SetEnabled(value bool) {
    defer exceptionProc()
    ListView_SetEnabled(l.instance, value)
}

func (l *TListView) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(ListView_GetFont(l.instance))
}

func (l *TListView) SetFont(value *TFont) {
    defer exceptionProc()
    ListView_SetFont(l.instance, CheckPtr(value))
}

func (l *TListView) FlatScrollBars() bool {
    defer exceptionProc()
    return ListView_GetFlatScrollBars(l.instance)
}

func (l *TListView) SetFlatScrollBars(value bool) {
    defer exceptionProc()
    ListView_SetFlatScrollBars(l.instance, value)
}

func (l *TListView) FullDrag() bool {
    defer exceptionProc()
    return ListView_GetFullDrag(l.instance)
}

func (l *TListView) SetFullDrag(value bool) {
    defer exceptionProc()
    ListView_SetFullDrag(l.instance, value)
}

func (l *TListView) GridLines() bool {
    defer exceptionProc()
    return ListView_GetGridLines(l.instance)
}

func (l *TListView) SetGridLines(value bool) {
    defer exceptionProc()
    ListView_SetGridLines(l.instance, value)
}

func (l *TListView) Groups() *TListGroups {
    defer exceptionProc()
    return ListGroupsFromInst(ListView_GetGroups(l.instance))
}

func (l *TListView) SetGroups(value *TListGroups) {
    defer exceptionProc()
    ListView_SetGroups(l.instance, CheckPtr(value))
}

func (l *TListView) HideSelection() bool {
    defer exceptionProc()
    return ListView_GetHideSelection(l.instance)
}

func (l *TListView) SetHideSelection(value bool) {
    defer exceptionProc()
    ListView_SetHideSelection(l.instance, value)
}

func (l *TListView) HotTrack() bool {
    defer exceptionProc()
    return ListView_GetHotTrack(l.instance)
}

func (l *TListView) SetHotTrack(value bool) {
    defer exceptionProc()
    ListView_SetHotTrack(l.instance, value)
}

func (l *TListView) HoverTime() int32 {
    defer exceptionProc()
    return ListView_GetHoverTime(l.instance)
}

func (l *TListView) SetHoverTime(value int32) {
    defer exceptionProc()
    ListView_SetHoverTime(l.instance, value)
}

func (l *TListView) Items() *TListItems {
    defer exceptionProc()
    return ListItemsFromInst(ListView_GetItems(l.instance))
}

func (l *TListView) SetItems(value *TListItems) {
    defer exceptionProc()
    ListView_SetItems(l.instance, CheckPtr(value))
}

func (l *TListView) LargeImages() *TImageList {
    defer exceptionProc()
    return ImageListFromInst(ListView_GetLargeImages(l.instance))
}

func (l *TListView) SetLargeImages(value IComponent) {
    defer exceptionProc()
    ListView_SetLargeImages(l.instance, CheckPtr(value))
}

func (l *TListView) MultiSelect() bool {
    defer exceptionProc()
    return ListView_GetMultiSelect(l.instance)
}

func (l *TListView) SetMultiSelect(value bool) {
    defer exceptionProc()
    ListView_SetMultiSelect(l.instance, value)
}

func (l *TListView) GroupHeaderImages() *TImageList {
    defer exceptionProc()
    return ImageListFromInst(ListView_GetGroupHeaderImages(l.instance))
}

func (l *TListView) SetGroupHeaderImages(value IComponent) {
    defer exceptionProc()
    ListView_SetGroupHeaderImages(l.instance, CheckPtr(value))
}

func (l *TListView) ReadOnly() bool {
    defer exceptionProc()
    return ListView_GetReadOnly(l.instance)
}

func (l *TListView) SetReadOnly(value bool) {
    defer exceptionProc()
    ListView_SetReadOnly(l.instance, value)
}

func (l *TListView) RowSelect() bool {
    defer exceptionProc()
    return ListView_GetRowSelect(l.instance)
}

func (l *TListView) SetRowSelect(value bool) {
    defer exceptionProc()
    ListView_SetRowSelect(l.instance, value)
}

func (l *TListView) ParentColor() bool {
    defer exceptionProc()
    return ListView_GetParentColor(l.instance)
}

func (l *TListView) SetParentColor(value bool) {
    defer exceptionProc()
    ListView_SetParentColor(l.instance, value)
}

func (l *TListView) ParentFont() bool {
    defer exceptionProc()
    return ListView_GetParentFont(l.instance)
}

func (l *TListView) SetParentFont(value bool) {
    defer exceptionProc()
    ListView_SetParentFont(l.instance, value)
}

func (l *TListView) ParentShowHint() bool {
    defer exceptionProc()
    return ListView_GetParentShowHint(l.instance)
}

func (l *TListView) SetParentShowHint(value bool) {
    defer exceptionProc()
    ListView_SetParentShowHint(l.instance, value)
}

func (l *TListView) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(ListView_GetPopupMenu(l.instance))
}

func (l *TListView) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    ListView_SetPopupMenu(l.instance, CheckPtr(value))
}

func (l *TListView) ShowColumnHeaders() bool {
    defer exceptionProc()
    return ListView_GetShowColumnHeaders(l.instance)
}

func (l *TListView) SetShowColumnHeaders(value bool) {
    defer exceptionProc()
    ListView_SetShowColumnHeaders(l.instance, value)
}

func (l *TListView) ShowWorkAreas() bool {
    defer exceptionProc()
    return ListView_GetShowWorkAreas(l.instance)
}

func (l *TListView) SetShowWorkAreas(value bool) {
    defer exceptionProc()
    ListView_SetShowWorkAreas(l.instance, value)
}

func (l *TListView) ShowHint() bool {
    defer exceptionProc()
    return ListView_GetShowHint(l.instance)
}

func (l *TListView) SetShowHint(value bool) {
    defer exceptionProc()
    ListView_SetShowHint(l.instance, value)
}

func (l *TListView) SmallImages() *TImageList {
    defer exceptionProc()
    return ImageListFromInst(ListView_GetSmallImages(l.instance))
}

func (l *TListView) SetSmallImages(value IComponent) {
    defer exceptionProc()
    ListView_SetSmallImages(l.instance, CheckPtr(value))
}

func (l *TListView) SortType() TSortType {
    defer exceptionProc()
    return ListView_GetSortType(l.instance)
}

func (l *TListView) SetSortType(value TSortType) {
    defer exceptionProc()
    ListView_SetSortType(l.instance, value)
}

func (l *TListView) StateImages() *TImageList {
    defer exceptionProc()
    return ImageListFromInst(ListView_GetStateImages(l.instance))
}

func (l *TListView) SetStateImages(value IComponent) {
    defer exceptionProc()
    ListView_SetStateImages(l.instance, CheckPtr(value))
}

func (l *TListView) TabOrder() int16 {
    defer exceptionProc()
    return ListView_GetTabOrder(l.instance)
}

func (l *TListView) SetTabOrder(value int16) {
    defer exceptionProc()
    ListView_SetTabOrder(l.instance, value)
}

func (l *TListView) TabStop() bool {
    defer exceptionProc()
    return ListView_GetTabStop(l.instance)
}

func (l *TListView) SetTabStop(value bool) {
    defer exceptionProc()
    ListView_SetTabStop(l.instance, value)
}

func (l *TListView) ViewStyle() TViewStyle {
    defer exceptionProc()
    return ListView_GetViewStyle(l.instance)
}

func (l *TListView) SetViewStyle(value TViewStyle) {
    defer exceptionProc()
    ListView_SetViewStyle(l.instance, value)
}

func (l *TListView) Visible() bool {
    defer exceptionProc()
    return ListView_GetVisible(l.instance)
}

func (l *TListView) SetVisible(value bool) {
    defer exceptionProc()
    ListView_SetVisible(l.instance, value)
}

func (l *TListView) SetOnChange(fn TLVChangeEvent) {
    ListView_SetOnChange(l.instance, fn)
}

func (l *TListView) SetOnClick(fn TNotifyEvent) {
    ListView_SetOnClick(l.instance, fn)
}

func (l *TListView) SetOnDblClick(fn TNotifyEvent) {
    ListView_SetOnDblClick(l.instance, fn)
}

func (l *TListView) SetOnEnter(fn TNotifyEvent) {
    ListView_SetOnEnter(l.instance, fn)
}

func (l *TListView) SetOnExit(fn TNotifyEvent) {
    ListView_SetOnExit(l.instance, fn)
}

func (l *TListView) SetOnKeyDown(fn TKeyEvent) {
    ListView_SetOnKeyDown(l.instance, fn)
}

func (l *TListView) SetOnKeyPress(fn TKeyPressEvent) {
    ListView_SetOnKeyPress(l.instance, fn)
}

func (l *TListView) SetOnKeyUp(fn TKeyEvent) {
    ListView_SetOnKeyUp(l.instance, fn)
}

func (l *TListView) SetOnMouseDown(fn TMouseEvent) {
    ListView_SetOnMouseDown(l.instance, fn)
}

func (l *TListView) SetOnMouseEnter(fn TNotifyEvent) {
    ListView_SetOnMouseEnter(l.instance, fn)
}

func (l *TListView) SetOnMouseLeave(fn TNotifyEvent) {
    ListView_SetOnMouseLeave(l.instance, fn)
}

func (l *TListView) SetOnMouseMove(fn TMouseMoveEvent) {
    ListView_SetOnMouseMove(l.instance, fn)
}

func (l *TListView) SetOnMouseUp(fn TMouseEvent) {
    ListView_SetOnMouseUp(l.instance, fn)
}

func (l *TListView) SetOnResize(fn TNotifyEvent) {
    ListView_SetOnResize(l.instance, fn)
}

func (l *TListView) Canvas() *TCanvas {
    defer exceptionProc()
    return CanvasFromInst(ListView_GetCanvas(l.instance))
}

func (l *TListView) DropTarget() *TListItem {
    defer exceptionProc()
    return ListItemFromInst(ListView_GetDropTarget(l.instance))
}

func (l *TListView) SetDropTarget(value *TListItem) {
    defer exceptionProc()
    ListView_SetDropTarget(l.instance, CheckPtr(value))
}

func (l *TListView) ItemFocused() *TListItem {
    defer exceptionProc()
    return ListItemFromInst(ListView_GetItemFocused(l.instance))
}

func (l *TListView) SetItemFocused(value *TListItem) {
    defer exceptionProc()
    ListView_SetItemFocused(l.instance, CheckPtr(value))
}

func (l *TListView) SelCount() int32 {
    defer exceptionProc()
    return ListView_GetSelCount(l.instance)
}

func (l *TListView) Selected() *TListItem {
    defer exceptionProc()
    return ListItemFromInst(ListView_GetSelected(l.instance))
}

func (l *TListView) SetSelected(value *TListItem) {
    defer exceptionProc()
    ListView_SetSelected(l.instance, CheckPtr(value))
}

func (l *TListView) TopItem() *TListItem {
    defer exceptionProc()
    return ListItemFromInst(ListView_GetTopItem(l.instance))
}

func (l *TListView) VisibleRowCount() int32 {
    defer exceptionProc()
    return ListView_GetVisibleRowCount(l.instance)
}

func (l *TListView) ItemIndex() int32 {
    defer exceptionProc()
    return ListView_GetItemIndex(l.instance)
}

func (l *TListView) SetItemIndex(value int32) {
    defer exceptionProc()
    ListView_SetItemIndex(l.instance, value)
}

func (l *TListView) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(ListView_GetBrush(l.instance))
}

func (l *TListView) ControlCount() int32 {
    defer exceptionProc()
    return ListView_GetControlCount(l.instance)
}

func (l *TListView) Handle() HWND {
    defer exceptionProc()
    return ListView_GetHandle(l.instance)
}

func (l *TListView) BoundsRect() TRect {
    defer exceptionProc()
    return ListView_GetBoundsRect(l.instance)
}

func (l *TListView) SetBoundsRect(value TRect) {
    defer exceptionProc()
    ListView_SetBoundsRect(l.instance, value)
}

func (l *TListView) ClientHeight() int32 {
    defer exceptionProc()
    return ListView_GetClientHeight(l.instance)
}

func (l *TListView) SetClientHeight(value int32) {
    defer exceptionProc()
    ListView_SetClientHeight(l.instance, value)
}

func (l *TListView) ClientRect() TRect {
    defer exceptionProc()
    return ListView_GetClientRect(l.instance)
}

func (l *TListView) ClientWidth() int32 {
    defer exceptionProc()
    return ListView_GetClientWidth(l.instance)
}

func (l *TListView) SetClientWidth(value int32) {
    defer exceptionProc()
    ListView_SetClientWidth(l.instance, value)
}

func (l *TListView) ExplicitLeft() int32 {
    defer exceptionProc()
    return ListView_GetExplicitLeft(l.instance)
}

func (l *TListView) ExplicitTop() int32 {
    defer exceptionProc()
    return ListView_GetExplicitTop(l.instance)
}

func (l *TListView) ExplicitWidth() int32 {
    defer exceptionProc()
    return ListView_GetExplicitWidth(l.instance)
}

func (l *TListView) ExplicitHeight() int32 {
    defer exceptionProc()
    return ListView_GetExplicitHeight(l.instance)
}

func (l *TListView) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(ListView_GetParent(l.instance))
}

func (l *TListView) SetParent(value IControl) {
    defer exceptionProc()
    ListView_SetParent(l.instance, CheckPtr(value))
}

func (l *TListView) AlignWithMargins() bool {
    defer exceptionProc()
    return ListView_GetAlignWithMargins(l.instance)
}

func (l *TListView) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    ListView_SetAlignWithMargins(l.instance, value)
}

func (l *TListView) Left() int32 {
    defer exceptionProc()
    return ListView_GetLeft(l.instance)
}

func (l *TListView) SetLeft(value int32) {
    defer exceptionProc()
    ListView_SetLeft(l.instance, value)
}

func (l *TListView) Top() int32 {
    defer exceptionProc()
    return ListView_GetTop(l.instance)
}

func (l *TListView) SetTop(value int32) {
    defer exceptionProc()
    ListView_SetTop(l.instance, value)
}

func (l *TListView) Width() int32 {
    defer exceptionProc()
    return ListView_GetWidth(l.instance)
}

func (l *TListView) SetWidth(value int32) {
    defer exceptionProc()
    ListView_SetWidth(l.instance, value)
}

func (l *TListView) Height() int32 {
    defer exceptionProc()
    return ListView_GetHeight(l.instance)
}

func (l *TListView) SetHeight(value int32) {
    defer exceptionProc()
    ListView_SetHeight(l.instance, value)
}

func (l *TListView) Cursor() TCursor {
    defer exceptionProc()
    return ListView_GetCursor(l.instance)
}

func (l *TListView) SetCursor(value TCursor) {
    defer exceptionProc()
    ListView_SetCursor(l.instance, value)
}

func (l *TListView) Hint() string {
    defer exceptionProc()
    return ListView_GetHint(l.instance)
}

func (l *TListView) SetHint(value string) {
    defer exceptionProc()
    ListView_SetHint(l.instance, value)
}

func (l *TListView) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(ListView_GetMargins(l.instance))
}

func (l *TListView) SetMargins(value *TMargins) {
    defer exceptionProc()
    ListView_SetMargins(l.instance, CheckPtr(value))
}

func (l *TListView) ComponentCount() int32 {
    defer exceptionProc()
    return ListView_GetComponentCount(l.instance)
}

func (l *TListView) ComponentIndex() int32 {
    defer exceptionProc()
    return ListView_GetComponentIndex(l.instance)
}

func (l *TListView) SetComponentIndex(value int32) {
    defer exceptionProc()
    ListView_SetComponentIndex(l.instance, value)
}

func (l *TListView) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ListView_GetOwner(l.instance))
}

func (l *TListView) Name() string {
    defer exceptionProc()
    return ListView_GetName(l.instance)
}

func (l *TListView) SetName(value string) {
    defer exceptionProc()
    ListView_SetName(l.instance, value)
}

func (l *TListView) Tag() int {
    defer exceptionProc()
    return ListView_GetTag(l.instance)
}

func (l *TListView) SetTag(value int) {
    defer exceptionProc()
    ListView_SetTag(l.instance, value)
}

func (l *TListView) Column(Index int32) *TListColumn {
    defer exceptionProc()
    return ListColumnFromInst(ListView_GetColumn(l.instance, Index))
}

func (l *TListView) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(ListView_GetControls(l.instance, Index))
}

func (l *TListView) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ListView_GetComponents(l.instance, AIndex))
}

