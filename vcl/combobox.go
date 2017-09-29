
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TComboBox struct {
    IControl
    instance uintptr
}

func NewComboBox(owner IComponent) *TComboBox {
    c := new(TComboBox)
    c.instance = ComboBox_Create(owner.Instance())
    return c
}

func ComboBoxFromInst(inst uintptr) *TComboBox {
    c := new(TComboBox)
    c.instance = inst
    return c
}

func ComboBoxFromObj(obj IObject) *TComboBox {
    c := new(TComboBox)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TComboBox) Free() {
    if c.instance != 0 {
        ComboBox_Free(c.instance)
    }
}

func (c *TComboBox) Instance() uintptr {
    return c.instance
}

func (c *TComboBox) IsValid() bool {
    return c.instance != 0
}

func (c *TComboBox) AddItem(Item string, AObject IObject) {
    defer exceptionProc()
    ComboBox_AddItem(c.instance, Item , CheckPtr(AObject))
}

func (c *TComboBox) Clear() {
    defer exceptionProc()
    ComboBox_Clear(c.instance)
}

func (c *TComboBox) ClearSelection() {
    defer exceptionProc()
    ComboBox_ClearSelection(c.instance)
}

func (c *TComboBox) DeleteSelected() {
    defer exceptionProc()
    ComboBox_DeleteSelected(c.instance)
}

func (c *TComboBox) Focused() bool {
    defer exceptionProc()
    return ComboBox_Focused(c.instance)
}

func (c *TComboBox) SelectAll() {
    defer exceptionProc()
    ComboBox_SelectAll(c.instance)
}

func (c *TComboBox) CanFocus() bool {
    defer exceptionProc()
    return ComboBox_CanFocus(c.instance)
}

func (c *TComboBox) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    ComboBox_FlipChildren(c.instance, AllLevels )
}

func (c *TComboBox) HandleAllocated() bool {
    defer exceptionProc()
    return ComboBox_HandleAllocated(c.instance)
}

func (c *TComboBox) Invalidate() {
    defer exceptionProc()
    ComboBox_Invalidate(c.instance)
}

func (c *TComboBox) Realign() {
    defer exceptionProc()
    ComboBox_Realign(c.instance)
}

func (c *TComboBox) Repaint() {
    defer exceptionProc()
    ComboBox_Repaint(c.instance)
}

func (c *TComboBox) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    ComboBox_ScaleBy(c.instance, M , D )
}

func (c *TComboBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    ComboBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight )
}

func (c *TComboBox) SetFocus() {
    defer exceptionProc()
    ComboBox_SetFocus(c.instance)
}

func (c *TComboBox) Update() {
    defer exceptionProc()
    ComboBox_Update(c.instance)
}

func (c *TComboBox) BringToFront() {
    defer exceptionProc()
    ComboBox_BringToFront(c.instance)
}

func (c *TComboBox) HasParent() bool {
    defer exceptionProc()
    return ComboBox_HasParent(c.instance)
}

func (c *TComboBox) Hide() {
    defer exceptionProc()
    ComboBox_Hide(c.instance)
}

func (c *TComboBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return ComboBox_Perform(c.instance, Msg , WParam , LParam )
}

func (c *TComboBox) Refresh() {
    defer exceptionProc()
    ComboBox_Refresh(c.instance)
}

func (c *TComboBox) SendToBack() {
    defer exceptionProc()
    ComboBox_SendToBack(c.instance)
}

func (c *TComboBox) Show() {
    defer exceptionProc()
    ComboBox_Show(c.instance)
}

func (c *TComboBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return ComboBox_GetTextBuf(c.instance, Buffer , BufSize )
}

func (c *TComboBox) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ComboBox_FindComponent(c.instance, AName ))
}

func (c *TComboBox) GetNamePath() string {
    defer exceptionProc()
    return ComboBox_GetNamePath(c.instance)
}

func (c *TComboBox) Assign(Source IObject) {
    defer exceptionProc()
    ComboBox_Assign(c.instance, CheckPtr(Source))
}

func (c *TComboBox) ClassName() string {
    defer exceptionProc()
    return ComboBox_ClassName(c.instance)
}

func (c *TComboBox) Equals(Obj IObject) bool {
    defer exceptionProc()
    return ComboBox_Equals(c.instance, CheckPtr(Obj))
}

func (c *TComboBox) GetHashCode() int32 {
    defer exceptionProc()
    return ComboBox_GetHashCode(c.instance)
}

func (c *TComboBox) ToString() string {
    defer exceptionProc()
    return ComboBox_ToString(c.instance)
}

func (c *TComboBox) Align() TAlign {
    defer exceptionProc()
    return ComboBox_GetAlign(c.instance)
}

func (c *TComboBox) SetAlign(value TAlign) {
    defer exceptionProc()
    ComboBox_SetAlign(c.instance, value)
}

func (c *TComboBox) AutoComplete() bool {
    defer exceptionProc()
    return ComboBox_GetAutoComplete(c.instance)
}

func (c *TComboBox) SetAutoComplete(value bool) {
    defer exceptionProc()
    ComboBox_SetAutoComplete(c.instance, value)
}

func (c *TComboBox) AutoCompleteDelay() uint32 {
    defer exceptionProc()
    return ComboBox_GetAutoCompleteDelay(c.instance)
}

func (c *TComboBox) SetAutoCompleteDelay(value uint32) {
    defer exceptionProc()
    ComboBox_SetAutoCompleteDelay(c.instance, value)
}

func (c *TComboBox) AutoDropDown() bool {
    defer exceptionProc()
    return ComboBox_GetAutoDropDown(c.instance)
}

func (c *TComboBox) SetAutoDropDown(value bool) {
    defer exceptionProc()
    ComboBox_SetAutoDropDown(c.instance, value)
}

func (c *TComboBox) AutoCloseUp() bool {
    defer exceptionProc()
    return ComboBox_GetAutoCloseUp(c.instance)
}

func (c *TComboBox) SetAutoCloseUp(value bool) {
    defer exceptionProc()
    ComboBox_SetAutoCloseUp(c.instance, value)
}

func (c *TComboBox) Style() TComboBoxStyle {
    defer exceptionProc()
    return ComboBox_GetStyle(c.instance)
}

func (c *TComboBox) SetStyle(value TComboBoxStyle) {
    defer exceptionProc()
    ComboBox_SetStyle(c.instance, value)
}

func (c *TComboBox) Anchors() TAnchors {
    defer exceptionProc()
    return ComboBox_GetAnchors(c.instance)
}

func (c *TComboBox) SetAnchors(value TAnchors) {
    defer exceptionProc()
    ComboBox_SetAnchors(c.instance, value)
}

func (c *TComboBox) Color() TColor {
    defer exceptionProc()
    return ComboBox_GetColor(c.instance)
}

func (c *TComboBox) SetColor(value TColor) {
    defer exceptionProc()
    ComboBox_SetColor(c.instance, value)
}

func (c *TComboBox) DoubleBuffered() bool {
    defer exceptionProc()
    return ComboBox_GetDoubleBuffered(c.instance)
}

func (c *TComboBox) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    ComboBox_SetDoubleBuffered(c.instance, value)
}

func (c *TComboBox) DropDownCount() int32 {
    defer exceptionProc()
    return ComboBox_GetDropDownCount(c.instance)
}

func (c *TComboBox) SetDropDownCount(value int32) {
    defer exceptionProc()
    ComboBox_SetDropDownCount(c.instance, value)
}

func (c *TComboBox) Enabled() bool {
    defer exceptionProc()
    return ComboBox_GetEnabled(c.instance)
}

func (c *TComboBox) SetEnabled(value bool) {
    defer exceptionProc()
    ComboBox_SetEnabled(c.instance, value)
}

func (c *TComboBox) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(ComboBox_GetFont(c.instance))
}

func (c *TComboBox) SetFont(value *TFont) {
    defer exceptionProc()
    ComboBox_SetFont(c.instance, CheckPtr(value))
}

func (c *TComboBox) ItemHeight() int32 {
    defer exceptionProc()
    return ComboBox_GetItemHeight(c.instance)
}

func (c *TComboBox) SetItemHeight(value int32) {
    defer exceptionProc()
    ComboBox_SetItemHeight(c.instance, value)
}

func (c *TComboBox) ItemIndex() int32 {
    defer exceptionProc()
    return ComboBox_GetItemIndex(c.instance)
}

func (c *TComboBox) SetItemIndex(value int32) {
    defer exceptionProc()
    ComboBox_SetItemIndex(c.instance, value)
}

func (c *TComboBox) MaxLength() int32 {
    defer exceptionProc()
    return ComboBox_GetMaxLength(c.instance)
}

func (c *TComboBox) SetMaxLength(value int32) {
    defer exceptionProc()
    ComboBox_SetMaxLength(c.instance, value)
}

func (c *TComboBox) ParentColor() bool {
    defer exceptionProc()
    return ComboBox_GetParentColor(c.instance)
}

func (c *TComboBox) SetParentColor(value bool) {
    defer exceptionProc()
    ComboBox_SetParentColor(c.instance, value)
}

func (c *TComboBox) ParentCtl3D() bool {
    defer exceptionProc()
    return ComboBox_GetParentCtl3D(c.instance)
}

func (c *TComboBox) SetParentCtl3D(value bool) {
    defer exceptionProc()
    ComboBox_SetParentCtl3D(c.instance, value)
}

func (c *TComboBox) ParentFont() bool {
    defer exceptionProc()
    return ComboBox_GetParentFont(c.instance)
}

func (c *TComboBox) SetParentFont(value bool) {
    defer exceptionProc()
    ComboBox_SetParentFont(c.instance, value)
}

func (c *TComboBox) ParentShowHint() bool {
    defer exceptionProc()
    return ComboBox_GetParentShowHint(c.instance)
}

func (c *TComboBox) SetParentShowHint(value bool) {
    defer exceptionProc()
    ComboBox_SetParentShowHint(c.instance, value)
}

func (c *TComboBox) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(ComboBox_GetPopupMenu(c.instance))
}

func (c *TComboBox) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    ComboBox_SetPopupMenu(c.instance, CheckPtr(value))
}

func (c *TComboBox) ShowHint() bool {
    defer exceptionProc()
    return ComboBox_GetShowHint(c.instance)
}

func (c *TComboBox) SetShowHint(value bool) {
    defer exceptionProc()
    ComboBox_SetShowHint(c.instance, value)
}

func (c *TComboBox) Sorted() bool {
    defer exceptionProc()
    return ComboBox_GetSorted(c.instance)
}

func (c *TComboBox) SetSorted(value bool) {
    defer exceptionProc()
    ComboBox_SetSorted(c.instance, value)
}

func (c *TComboBox) TabOrder() int16 {
    defer exceptionProc()
    return ComboBox_GetTabOrder(c.instance)
}

func (c *TComboBox) SetTabOrder(value int16) {
    defer exceptionProc()
    ComboBox_SetTabOrder(c.instance, value)
}

func (c *TComboBox) TabStop() bool {
    defer exceptionProc()
    return ComboBox_GetTabStop(c.instance)
}

func (c *TComboBox) SetTabStop(value bool) {
    defer exceptionProc()
    ComboBox_SetTabStop(c.instance, value)
}

func (c *TComboBox) Text() string {
    defer exceptionProc()
    return ComboBox_GetText(c.instance)
}

func (c *TComboBox) SetText(value string) {
    defer exceptionProc()
    ComboBox_SetText(c.instance, value)
}

func (c *TComboBox) TextHint() string {
    defer exceptionProc()
    return ComboBox_GetTextHint(c.instance)
}

func (c *TComboBox) SetTextHint(value string) {
    defer exceptionProc()
    ComboBox_SetTextHint(c.instance, value)
}

func (c *TComboBox) Visible() bool {
    defer exceptionProc()
    return ComboBox_GetVisible(c.instance)
}

func (c *TComboBox) SetVisible(value bool) {
    defer exceptionProc()
    ComboBox_SetVisible(c.instance, value)
}

func (c *TComboBox) SetOnChange(fn TNotifyEvent) {
    ComboBox_SetOnChange(c.instance, fn)
}

func (c *TComboBox) SetOnClick(fn TNotifyEvent) {
    ComboBox_SetOnClick(c.instance, fn)
}

func (c *TComboBox) SetOnDblClick(fn TNotifyEvent) {
    ComboBox_SetOnDblClick(c.instance, fn)
}

func (c *TComboBox) SetOnEnter(fn TNotifyEvent) {
    ComboBox_SetOnEnter(c.instance, fn)
}

func (c *TComboBox) SetOnExit(fn TNotifyEvent) {
    ComboBox_SetOnExit(c.instance, fn)
}

func (c *TComboBox) SetOnKeyDown(fn TKeyEvent) {
    ComboBox_SetOnKeyDown(c.instance, fn)
}

func (c *TComboBox) SetOnKeyPress(fn TKeyPressEvent) {
    ComboBox_SetOnKeyPress(c.instance, fn)
}

func (c *TComboBox) SetOnKeyUp(fn TKeyEvent) {
    ComboBox_SetOnKeyUp(c.instance, fn)
}

func (c *TComboBox) SetOnMouseEnter(fn TNotifyEvent) {
    ComboBox_SetOnMouseEnter(c.instance, fn)
}

func (c *TComboBox) SetOnMouseLeave(fn TNotifyEvent) {
    ComboBox_SetOnMouseLeave(c.instance, fn)
}

func (c *TComboBox) Items() *TStrings {
    defer exceptionProc()
    return StringsFromInst(ComboBox_GetItems(c.instance))
}

func (c *TComboBox) SetItems(value IObject) {
    defer exceptionProc()
    ComboBox_SetItems(c.instance, CheckPtr(value))
}

func (c *TComboBox) SelText() string {
    defer exceptionProc()
    return ComboBox_GetSelText(c.instance)
}

func (c *TComboBox) SetSelText(value string) {
    defer exceptionProc()
    ComboBox_SetSelText(c.instance, value)
}

func (c *TComboBox) Canvas() *TCanvas {
    defer exceptionProc()
    return CanvasFromInst(ComboBox_GetCanvas(c.instance))
}

func (c *TComboBox) SelLength() int32 {
    defer exceptionProc()
    return ComboBox_GetSelLength(c.instance)
}

func (c *TComboBox) SetSelLength(value int32) {
    defer exceptionProc()
    ComboBox_SetSelLength(c.instance, value)
}

func (c *TComboBox) SelStart() int32 {
    defer exceptionProc()
    return ComboBox_GetSelStart(c.instance)
}

func (c *TComboBox) SetSelStart(value int32) {
    defer exceptionProc()
    ComboBox_SetSelStart(c.instance, value)
}

func (c *TComboBox) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(ComboBox_GetBrush(c.instance))
}

func (c *TComboBox) ControlCount() int32 {
    defer exceptionProc()
    return ComboBox_GetControlCount(c.instance)
}

func (c *TComboBox) Handle() HWND {
    defer exceptionProc()
    return ComboBox_GetHandle(c.instance)
}

func (c *TComboBox) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(ComboBox_GetAction(c.instance))
}

func (c *TComboBox) SetAction(value IComponent) {
    defer exceptionProc()
    ComboBox_SetAction(c.instance, CheckPtr(value))
}

func (c *TComboBox) BoundsRect() TRect {
    defer exceptionProc()
    return ComboBox_GetBoundsRect(c.instance)
}

func (c *TComboBox) SetBoundsRect(value TRect) {
    defer exceptionProc()
    ComboBox_SetBoundsRect(c.instance, value)
}

func (c *TComboBox) ClientHeight() int32 {
    defer exceptionProc()
    return ComboBox_GetClientHeight(c.instance)
}

func (c *TComboBox) SetClientHeight(value int32) {
    defer exceptionProc()
    ComboBox_SetClientHeight(c.instance, value)
}

func (c *TComboBox) ClientRect() TRect {
    defer exceptionProc()
    return ComboBox_GetClientRect(c.instance)
}

func (c *TComboBox) ClientWidth() int32 {
    defer exceptionProc()
    return ComboBox_GetClientWidth(c.instance)
}

func (c *TComboBox) SetClientWidth(value int32) {
    defer exceptionProc()
    ComboBox_SetClientWidth(c.instance, value)
}

func (c *TComboBox) ExplicitLeft() int32 {
    defer exceptionProc()
    return ComboBox_GetExplicitLeft(c.instance)
}

func (c *TComboBox) ExplicitTop() int32 {
    defer exceptionProc()
    return ComboBox_GetExplicitTop(c.instance)
}

func (c *TComboBox) ExplicitWidth() int32 {
    defer exceptionProc()
    return ComboBox_GetExplicitWidth(c.instance)
}

func (c *TComboBox) ExplicitHeight() int32 {
    defer exceptionProc()
    return ComboBox_GetExplicitHeight(c.instance)
}

func (c *TComboBox) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(ComboBox_GetParent(c.instance))
}

func (c *TComboBox) SetParent(value IControl) {
    defer exceptionProc()
    ComboBox_SetParent(c.instance, CheckPtr(value))
}

func (c *TComboBox) AlignWithMargins() bool {
    defer exceptionProc()
    return ComboBox_GetAlignWithMargins(c.instance)
}

func (c *TComboBox) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    ComboBox_SetAlignWithMargins(c.instance, value)
}

func (c *TComboBox) Left() int32 {
    defer exceptionProc()
    return ComboBox_GetLeft(c.instance)
}

func (c *TComboBox) SetLeft(value int32) {
    defer exceptionProc()
    ComboBox_SetLeft(c.instance, value)
}

func (c *TComboBox) Top() int32 {
    defer exceptionProc()
    return ComboBox_GetTop(c.instance)
}

func (c *TComboBox) SetTop(value int32) {
    defer exceptionProc()
    ComboBox_SetTop(c.instance, value)
}

func (c *TComboBox) Width() int32 {
    defer exceptionProc()
    return ComboBox_GetWidth(c.instance)
}

func (c *TComboBox) SetWidth(value int32) {
    defer exceptionProc()
    ComboBox_SetWidth(c.instance, value)
}

func (c *TComboBox) Height() int32 {
    defer exceptionProc()
    return ComboBox_GetHeight(c.instance)
}

func (c *TComboBox) SetHeight(value int32) {
    defer exceptionProc()
    ComboBox_SetHeight(c.instance, value)
}

func (c *TComboBox) Cursor() TCursor {
    defer exceptionProc()
    return ComboBox_GetCursor(c.instance)
}

func (c *TComboBox) SetCursor(value TCursor) {
    defer exceptionProc()
    ComboBox_SetCursor(c.instance, value)
}

func (c *TComboBox) Hint() string {
    defer exceptionProc()
    return ComboBox_GetHint(c.instance)
}

func (c *TComboBox) SetHint(value string) {
    defer exceptionProc()
    ComboBox_SetHint(c.instance, value)
}

func (c *TComboBox) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(ComboBox_GetMargins(c.instance))
}

func (c *TComboBox) SetMargins(value *TMargins) {
    defer exceptionProc()
    ComboBox_SetMargins(c.instance, CheckPtr(value))
}

func (c *TComboBox) ComponentCount() int32 {
    defer exceptionProc()
    return ComboBox_GetComponentCount(c.instance)
}

func (c *TComboBox) ComponentIndex() int32 {
    defer exceptionProc()
    return ComboBox_GetComponentIndex(c.instance)
}

func (c *TComboBox) SetComponentIndex(value int32) {
    defer exceptionProc()
    ComboBox_SetComponentIndex(c.instance, value)
}

func (c *TComboBox) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ComboBox_GetOwner(c.instance))
}

func (c *TComboBox) Name() string {
    defer exceptionProc()
    return ComboBox_GetName(c.instance)
}

func (c *TComboBox) SetName(value string) {
    defer exceptionProc()
    ComboBox_SetName(c.instance, value)
}

func (c *TComboBox) Tag() int {
    defer exceptionProc()
    return ComboBox_GetTag(c.instance)
}

func (c *TComboBox) SetTag(value int) {
    defer exceptionProc()
    ComboBox_SetTag(c.instance, value)
}

func (c *TComboBox) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(ComboBox_GetControls(c.instance, Index))
}

func (c *TComboBox) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ComboBox_GetComponents(c.instance, AIndex))
}

