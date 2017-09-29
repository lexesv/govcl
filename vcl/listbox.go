
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TListBox struct {
    IControl
    instance uintptr
}

func NewListBox(owner IComponent) *TListBox {
    l := new(TListBox)
    l.instance = ListBox_Create(owner.Instance())
    return l
}

func ListBoxFromInst(inst uintptr) *TListBox {
    l := new(TListBox)
    l.instance = inst
    return l
}

func ListBoxFromObj(obj IObject) *TListBox {
    l := new(TListBox)
    l.instance = CheckPtr(obj)
    return l
}

func (l *TListBox) Free() {
    if l.instance != 0 {
        ListBox_Free(l.instance)
    }
}

func (l *TListBox) Instance() uintptr {
    return l.instance
}

func (l *TListBox) IsValid() bool {
    return l.instance != 0
}

func (l *TListBox) AddItem(Item string, AObject IObject) {
    defer exceptionProc()
    ListBox_AddItem(l.instance, Item , CheckPtr(AObject))
}

func (l *TListBox) Clear() {
    defer exceptionProc()
    ListBox_Clear(l.instance)
}

func (l *TListBox) ClearSelection() {
    defer exceptionProc()
    ListBox_ClearSelection(l.instance)
}

func (l *TListBox) DeleteSelected() {
    defer exceptionProc()
    ListBox_DeleteSelected(l.instance)
}

func (l *TListBox) SelectAll() {
    defer exceptionProc()
    ListBox_SelectAll(l.instance)
}

func (l *TListBox) CanFocus() bool {
    defer exceptionProc()
    return ListBox_CanFocus(l.instance)
}

func (l *TListBox) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    ListBox_FlipChildren(l.instance, AllLevels )
}

func (l *TListBox) Focused() bool {
    defer exceptionProc()
    return ListBox_Focused(l.instance)
}

func (l *TListBox) HandleAllocated() bool {
    defer exceptionProc()
    return ListBox_HandleAllocated(l.instance)
}

func (l *TListBox) Invalidate() {
    defer exceptionProc()
    ListBox_Invalidate(l.instance)
}

func (l *TListBox) Realign() {
    defer exceptionProc()
    ListBox_Realign(l.instance)
}

func (l *TListBox) Repaint() {
    defer exceptionProc()
    ListBox_Repaint(l.instance)
}

func (l *TListBox) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    ListBox_ScaleBy(l.instance, M , D )
}

func (l *TListBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    ListBox_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight )
}

func (l *TListBox) SetFocus() {
    defer exceptionProc()
    ListBox_SetFocus(l.instance)
}

func (l *TListBox) Update() {
    defer exceptionProc()
    ListBox_Update(l.instance)
}

func (l *TListBox) BringToFront() {
    defer exceptionProc()
    ListBox_BringToFront(l.instance)
}

func (l *TListBox) HasParent() bool {
    defer exceptionProc()
    return ListBox_HasParent(l.instance)
}

func (l *TListBox) Hide() {
    defer exceptionProc()
    ListBox_Hide(l.instance)
}

func (l *TListBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return ListBox_Perform(l.instance, Msg , WParam , LParam )
}

func (l *TListBox) Refresh() {
    defer exceptionProc()
    ListBox_Refresh(l.instance)
}

func (l *TListBox) SendToBack() {
    defer exceptionProc()
    ListBox_SendToBack(l.instance)
}

func (l *TListBox) Show() {
    defer exceptionProc()
    ListBox_Show(l.instance)
}

func (l *TListBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return ListBox_GetTextBuf(l.instance, Buffer , BufSize )
}

func (l *TListBox) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ListBox_FindComponent(l.instance, AName ))
}

func (l *TListBox) GetNamePath() string {
    defer exceptionProc()
    return ListBox_GetNamePath(l.instance)
}

func (l *TListBox) Assign(Source IObject) {
    defer exceptionProc()
    ListBox_Assign(l.instance, CheckPtr(Source))
}

func (l *TListBox) ClassName() string {
    defer exceptionProc()
    return ListBox_ClassName(l.instance)
}

func (l *TListBox) Equals(Obj IObject) bool {
    defer exceptionProc()
    return ListBox_Equals(l.instance, CheckPtr(Obj))
}

func (l *TListBox) GetHashCode() int32 {
    defer exceptionProc()
    return ListBox_GetHashCode(l.instance)
}

func (l *TListBox) ToString() string {
    defer exceptionProc()
    return ListBox_ToString(l.instance)
}

func (l *TListBox) Style() TListBoxStyle {
    defer exceptionProc()
    return ListBox_GetStyle(l.instance)
}

func (l *TListBox) SetStyle(value TListBoxStyle) {
    defer exceptionProc()
    ListBox_SetStyle(l.instance, value)
}

func (l *TListBox) AutoComplete() bool {
    defer exceptionProc()
    return ListBox_GetAutoComplete(l.instance)
}

func (l *TListBox) SetAutoComplete(value bool) {
    defer exceptionProc()
    ListBox_SetAutoComplete(l.instance, value)
}

func (l *TListBox) AutoCompleteDelay() uint32 {
    defer exceptionProc()
    return ListBox_GetAutoCompleteDelay(l.instance)
}

func (l *TListBox) SetAutoCompleteDelay(value uint32) {
    defer exceptionProc()
    ListBox_SetAutoCompleteDelay(l.instance, value)
}

func (l *TListBox) Align() TAlign {
    defer exceptionProc()
    return ListBox_GetAlign(l.instance)
}

func (l *TListBox) SetAlign(value TAlign) {
    defer exceptionProc()
    ListBox_SetAlign(l.instance, value)
}

func (l *TListBox) Anchors() TAnchors {
    defer exceptionProc()
    return ListBox_GetAnchors(l.instance)
}

func (l *TListBox) SetAnchors(value TAnchors) {
    defer exceptionProc()
    ListBox_SetAnchors(l.instance, value)
}

func (l *TListBox) BorderStyle() TBorderStyle {
    defer exceptionProc()
    return ListBox_GetBorderStyle(l.instance)
}

func (l *TListBox) SetBorderStyle(value TBorderStyle) {
    defer exceptionProc()
    ListBox_SetBorderStyle(l.instance, value)
}

func (l *TListBox) Color() TColor {
    defer exceptionProc()
    return ListBox_GetColor(l.instance)
}

func (l *TListBox) SetColor(value TColor) {
    defer exceptionProc()
    ListBox_SetColor(l.instance, value)
}

func (l *TListBox) Columns() int32 {
    defer exceptionProc()
    return ListBox_GetColumns(l.instance)
}

func (l *TListBox) SetColumns(value int32) {
    defer exceptionProc()
    ListBox_SetColumns(l.instance, value)
}

func (l *TListBox) DoubleBuffered() bool {
    defer exceptionProc()
    return ListBox_GetDoubleBuffered(l.instance)
}

func (l *TListBox) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    ListBox_SetDoubleBuffered(l.instance, value)
}

func (l *TListBox) Enabled() bool {
    defer exceptionProc()
    return ListBox_GetEnabled(l.instance)
}

func (l *TListBox) SetEnabled(value bool) {
    defer exceptionProc()
    ListBox_SetEnabled(l.instance, value)
}

func (l *TListBox) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(ListBox_GetFont(l.instance))
}

func (l *TListBox) SetFont(value *TFont) {
    defer exceptionProc()
    ListBox_SetFont(l.instance, CheckPtr(value))
}

func (l *TListBox) ItemHeight() int32 {
    defer exceptionProc()
    return ListBox_GetItemHeight(l.instance)
}

func (l *TListBox) SetItemHeight(value int32) {
    defer exceptionProc()
    ListBox_SetItemHeight(l.instance, value)
}

func (l *TListBox) Items() *TStrings {
    defer exceptionProc()
    return StringsFromInst(ListBox_GetItems(l.instance))
}

func (l *TListBox) SetItems(value IObject) {
    defer exceptionProc()
    ListBox_SetItems(l.instance, CheckPtr(value))
}

func (l *TListBox) MultiSelect() bool {
    defer exceptionProc()
    return ListBox_GetMultiSelect(l.instance)
}

func (l *TListBox) SetMultiSelect(value bool) {
    defer exceptionProc()
    ListBox_SetMultiSelect(l.instance, value)
}

func (l *TListBox) ParentColor() bool {
    defer exceptionProc()
    return ListBox_GetParentColor(l.instance)
}

func (l *TListBox) SetParentColor(value bool) {
    defer exceptionProc()
    ListBox_SetParentColor(l.instance, value)
}

func (l *TListBox) ParentCtl3D() bool {
    defer exceptionProc()
    return ListBox_GetParentCtl3D(l.instance)
}

func (l *TListBox) SetParentCtl3D(value bool) {
    defer exceptionProc()
    ListBox_SetParentCtl3D(l.instance, value)
}

func (l *TListBox) ParentFont() bool {
    defer exceptionProc()
    return ListBox_GetParentFont(l.instance)
}

func (l *TListBox) SetParentFont(value bool) {
    defer exceptionProc()
    ListBox_SetParentFont(l.instance, value)
}

func (l *TListBox) ParentShowHint() bool {
    defer exceptionProc()
    return ListBox_GetParentShowHint(l.instance)
}

func (l *TListBox) SetParentShowHint(value bool) {
    defer exceptionProc()
    ListBox_SetParentShowHint(l.instance, value)
}

func (l *TListBox) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(ListBox_GetPopupMenu(l.instance))
}

func (l *TListBox) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    ListBox_SetPopupMenu(l.instance, CheckPtr(value))
}

func (l *TListBox) ShowHint() bool {
    defer exceptionProc()
    return ListBox_GetShowHint(l.instance)
}

func (l *TListBox) SetShowHint(value bool) {
    defer exceptionProc()
    ListBox_SetShowHint(l.instance, value)
}

func (l *TListBox) Sorted() bool {
    defer exceptionProc()
    return ListBox_GetSorted(l.instance)
}

func (l *TListBox) SetSorted(value bool) {
    defer exceptionProc()
    ListBox_SetSorted(l.instance, value)
}

func (l *TListBox) TabOrder() int16 {
    defer exceptionProc()
    return ListBox_GetTabOrder(l.instance)
}

func (l *TListBox) SetTabOrder(value int16) {
    defer exceptionProc()
    ListBox_SetTabOrder(l.instance, value)
}

func (l *TListBox) TabStop() bool {
    defer exceptionProc()
    return ListBox_GetTabStop(l.instance)
}

func (l *TListBox) SetTabStop(value bool) {
    defer exceptionProc()
    ListBox_SetTabStop(l.instance, value)
}

func (l *TListBox) TabWidth() int32 {
    defer exceptionProc()
    return ListBox_GetTabWidth(l.instance)
}

func (l *TListBox) SetTabWidth(value int32) {
    defer exceptionProc()
    ListBox_SetTabWidth(l.instance, value)
}

func (l *TListBox) Visible() bool {
    defer exceptionProc()
    return ListBox_GetVisible(l.instance)
}

func (l *TListBox) SetVisible(value bool) {
    defer exceptionProc()
    ListBox_SetVisible(l.instance, value)
}

func (l *TListBox) SetOnClick(fn TNotifyEvent) {
    ListBox_SetOnClick(l.instance, fn)
}

func (l *TListBox) SetOnDblClick(fn TNotifyEvent) {
    ListBox_SetOnDblClick(l.instance, fn)
}

func (l *TListBox) SetOnEnter(fn TNotifyEvent) {
    ListBox_SetOnEnter(l.instance, fn)
}

func (l *TListBox) SetOnExit(fn TNotifyEvent) {
    ListBox_SetOnExit(l.instance, fn)
}

func (l *TListBox) SetOnKeyDown(fn TKeyEvent) {
    ListBox_SetOnKeyDown(l.instance, fn)
}

func (l *TListBox) SetOnKeyPress(fn TKeyPressEvent) {
    ListBox_SetOnKeyPress(l.instance, fn)
}

func (l *TListBox) SetOnKeyUp(fn TKeyEvent) {
    ListBox_SetOnKeyUp(l.instance, fn)
}

func (l *TListBox) SetOnMouseDown(fn TMouseEvent) {
    ListBox_SetOnMouseDown(l.instance, fn)
}

func (l *TListBox) SetOnMouseEnter(fn TNotifyEvent) {
    ListBox_SetOnMouseEnter(l.instance, fn)
}

func (l *TListBox) SetOnMouseLeave(fn TNotifyEvent) {
    ListBox_SetOnMouseLeave(l.instance, fn)
}

func (l *TListBox) SetOnMouseMove(fn TMouseMoveEvent) {
    ListBox_SetOnMouseMove(l.instance, fn)
}

func (l *TListBox) SetOnMouseUp(fn TMouseEvent) {
    ListBox_SetOnMouseUp(l.instance, fn)
}

func (l *TListBox) Canvas() *TCanvas {
    defer exceptionProc()
    return CanvasFromInst(ListBox_GetCanvas(l.instance))
}

func (l *TListBox) SelCount() int32 {
    defer exceptionProc()
    return ListBox_GetSelCount(l.instance)
}

func (l *TListBox) ItemIndex() int32 {
    defer exceptionProc()
    return ListBox_GetItemIndex(l.instance)
}

func (l *TListBox) SetItemIndex(value int32) {
    defer exceptionProc()
    ListBox_SetItemIndex(l.instance, value)
}

func (l *TListBox) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(ListBox_GetBrush(l.instance))
}

func (l *TListBox) ControlCount() int32 {
    defer exceptionProc()
    return ListBox_GetControlCount(l.instance)
}

func (l *TListBox) Handle() HWND {
    defer exceptionProc()
    return ListBox_GetHandle(l.instance)
}

func (l *TListBox) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(ListBox_GetAction(l.instance))
}

func (l *TListBox) SetAction(value IComponent) {
    defer exceptionProc()
    ListBox_SetAction(l.instance, CheckPtr(value))
}

func (l *TListBox) BoundsRect() TRect {
    defer exceptionProc()
    return ListBox_GetBoundsRect(l.instance)
}

func (l *TListBox) SetBoundsRect(value TRect) {
    defer exceptionProc()
    ListBox_SetBoundsRect(l.instance, value)
}

func (l *TListBox) ClientHeight() int32 {
    defer exceptionProc()
    return ListBox_GetClientHeight(l.instance)
}

func (l *TListBox) SetClientHeight(value int32) {
    defer exceptionProc()
    ListBox_SetClientHeight(l.instance, value)
}

func (l *TListBox) ClientRect() TRect {
    defer exceptionProc()
    return ListBox_GetClientRect(l.instance)
}

func (l *TListBox) ClientWidth() int32 {
    defer exceptionProc()
    return ListBox_GetClientWidth(l.instance)
}

func (l *TListBox) SetClientWidth(value int32) {
    defer exceptionProc()
    ListBox_SetClientWidth(l.instance, value)
}

func (l *TListBox) ExplicitLeft() int32 {
    defer exceptionProc()
    return ListBox_GetExplicitLeft(l.instance)
}

func (l *TListBox) ExplicitTop() int32 {
    defer exceptionProc()
    return ListBox_GetExplicitTop(l.instance)
}

func (l *TListBox) ExplicitWidth() int32 {
    defer exceptionProc()
    return ListBox_GetExplicitWidth(l.instance)
}

func (l *TListBox) ExplicitHeight() int32 {
    defer exceptionProc()
    return ListBox_GetExplicitHeight(l.instance)
}

func (l *TListBox) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(ListBox_GetParent(l.instance))
}

func (l *TListBox) SetParent(value IControl) {
    defer exceptionProc()
    ListBox_SetParent(l.instance, CheckPtr(value))
}

func (l *TListBox) AlignWithMargins() bool {
    defer exceptionProc()
    return ListBox_GetAlignWithMargins(l.instance)
}

func (l *TListBox) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    ListBox_SetAlignWithMargins(l.instance, value)
}

func (l *TListBox) Left() int32 {
    defer exceptionProc()
    return ListBox_GetLeft(l.instance)
}

func (l *TListBox) SetLeft(value int32) {
    defer exceptionProc()
    ListBox_SetLeft(l.instance, value)
}

func (l *TListBox) Top() int32 {
    defer exceptionProc()
    return ListBox_GetTop(l.instance)
}

func (l *TListBox) SetTop(value int32) {
    defer exceptionProc()
    ListBox_SetTop(l.instance, value)
}

func (l *TListBox) Width() int32 {
    defer exceptionProc()
    return ListBox_GetWidth(l.instance)
}

func (l *TListBox) SetWidth(value int32) {
    defer exceptionProc()
    ListBox_SetWidth(l.instance, value)
}

func (l *TListBox) Height() int32 {
    defer exceptionProc()
    return ListBox_GetHeight(l.instance)
}

func (l *TListBox) SetHeight(value int32) {
    defer exceptionProc()
    ListBox_SetHeight(l.instance, value)
}

func (l *TListBox) Cursor() TCursor {
    defer exceptionProc()
    return ListBox_GetCursor(l.instance)
}

func (l *TListBox) SetCursor(value TCursor) {
    defer exceptionProc()
    ListBox_SetCursor(l.instance, value)
}

func (l *TListBox) Hint() string {
    defer exceptionProc()
    return ListBox_GetHint(l.instance)
}

func (l *TListBox) SetHint(value string) {
    defer exceptionProc()
    ListBox_SetHint(l.instance, value)
}

func (l *TListBox) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(ListBox_GetMargins(l.instance))
}

func (l *TListBox) SetMargins(value *TMargins) {
    defer exceptionProc()
    ListBox_SetMargins(l.instance, CheckPtr(value))
}

func (l *TListBox) ComponentCount() int32 {
    defer exceptionProc()
    return ListBox_GetComponentCount(l.instance)
}

func (l *TListBox) ComponentIndex() int32 {
    defer exceptionProc()
    return ListBox_GetComponentIndex(l.instance)
}

func (l *TListBox) SetComponentIndex(value int32) {
    defer exceptionProc()
    ListBox_SetComponentIndex(l.instance, value)
}

func (l *TListBox) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ListBox_GetOwner(l.instance))
}

func (l *TListBox) Name() string {
    defer exceptionProc()
    return ListBox_GetName(l.instance)
}

func (l *TListBox) SetName(value string) {
    defer exceptionProc()
    ListBox_SetName(l.instance, value)
}

func (l *TListBox) Tag() int {
    defer exceptionProc()
    return ListBox_GetTag(l.instance)
}

func (l *TListBox) SetTag(value int) {
    defer exceptionProc()
    ListBox_SetTag(l.instance, value)
}

func (l *TListBox) Selected(Index int32) bool {
    defer exceptionProc()
    return ListBox_GetSelected(l.instance, Index)
}

func (l *TListBox) SetSelected(Index int32, value bool) {
    defer exceptionProc()
    ListBox_SetSelected(l.instance, Index, value)
}

func (l *TListBox) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(ListBox_GetControls(l.instance, Index))
}

func (l *TListBox) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ListBox_GetComponents(l.instance, AIndex))
}

