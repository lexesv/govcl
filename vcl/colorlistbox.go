
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TColorListBox struct {
    IControl
    instance uintptr
}

func NewColorListBox(owner IComponent) *TColorListBox {
    c := new(TColorListBox)
    c.instance = ColorListBox_Create(owner.Instance())
    return c
}

func ColorListBoxFromInst(inst uintptr) *TColorListBox {
    c := new(TColorListBox)
    c.instance = inst
    return c
}

func ColorListBoxFromObj(obj IObject) *TColorListBox {
    c := new(TColorListBox)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TColorListBox) Free() {
    if c.instance != 0 {
        ColorListBox_Free(c.instance)
    }
}

func (c *TColorListBox) Instance() uintptr {
    return c.instance
}

func (c *TColorListBox) IsValid() bool {
    return c.instance != 0
}

func (c *TColorListBox) AddItem(Item string, AObject IObject) {
    defer exceptionProc()
    ColorListBox_AddItem(c.instance, Item , CheckPtr(AObject))
}

func (c *TColorListBox) Clear() {
    defer exceptionProc()
    ColorListBox_Clear(c.instance)
}

func (c *TColorListBox) ClearSelection() {
    defer exceptionProc()
    ColorListBox_ClearSelection(c.instance)
}

func (c *TColorListBox) DeleteSelected() {
    defer exceptionProc()
    ColorListBox_DeleteSelected(c.instance)
}

func (c *TColorListBox) SelectAll() {
    defer exceptionProc()
    ColorListBox_SelectAll(c.instance)
}

func (c *TColorListBox) CanFocus() bool {
    defer exceptionProc()
    return ColorListBox_CanFocus(c.instance)
}

func (c *TColorListBox) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    ColorListBox_FlipChildren(c.instance, AllLevels )
}

func (c *TColorListBox) Focused() bool {
    defer exceptionProc()
    return ColorListBox_Focused(c.instance)
}

func (c *TColorListBox) HandleAllocated() bool {
    defer exceptionProc()
    return ColorListBox_HandleAllocated(c.instance)
}

func (c *TColorListBox) Invalidate() {
    defer exceptionProc()
    ColorListBox_Invalidate(c.instance)
}

func (c *TColorListBox) Realign() {
    defer exceptionProc()
    ColorListBox_Realign(c.instance)
}

func (c *TColorListBox) Repaint() {
    defer exceptionProc()
    ColorListBox_Repaint(c.instance)
}

func (c *TColorListBox) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    ColorListBox_ScaleBy(c.instance, M , D )
}

func (c *TColorListBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    ColorListBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight )
}

func (c *TColorListBox) SetFocus() {
    defer exceptionProc()
    ColorListBox_SetFocus(c.instance)
}

func (c *TColorListBox) Update() {
    defer exceptionProc()
    ColorListBox_Update(c.instance)
}

func (c *TColorListBox) BringToFront() {
    defer exceptionProc()
    ColorListBox_BringToFront(c.instance)
}

func (c *TColorListBox) HasParent() bool {
    defer exceptionProc()
    return ColorListBox_HasParent(c.instance)
}

func (c *TColorListBox) Hide() {
    defer exceptionProc()
    ColorListBox_Hide(c.instance)
}

func (c *TColorListBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return ColorListBox_Perform(c.instance, Msg , WParam , LParam )
}

func (c *TColorListBox) Refresh() {
    defer exceptionProc()
    ColorListBox_Refresh(c.instance)
}

func (c *TColorListBox) SendToBack() {
    defer exceptionProc()
    ColorListBox_SendToBack(c.instance)
}

func (c *TColorListBox) Show() {
    defer exceptionProc()
    ColorListBox_Show(c.instance)
}

func (c *TColorListBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return ColorListBox_GetTextBuf(c.instance, Buffer , BufSize )
}

func (c *TColorListBox) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ColorListBox_FindComponent(c.instance, AName ))
}

func (c *TColorListBox) GetNamePath() string {
    defer exceptionProc()
    return ColorListBox_GetNamePath(c.instance)
}

func (c *TColorListBox) Assign(Source IObject) {
    defer exceptionProc()
    ColorListBox_Assign(c.instance, CheckPtr(Source))
}

func (c *TColorListBox) ClassName() string {
    defer exceptionProc()
    return ColorListBox_ClassName(c.instance)
}

func (c *TColorListBox) Equals(Obj IObject) bool {
    defer exceptionProc()
    return ColorListBox_Equals(c.instance, CheckPtr(Obj))
}

func (c *TColorListBox) GetHashCode() int32 {
    defer exceptionProc()
    return ColorListBox_GetHashCode(c.instance)
}

func (c *TColorListBox) ToString() string {
    defer exceptionProc()
    return ColorListBox_ToString(c.instance)
}

func (c *TColorListBox) Align() TAlign {
    defer exceptionProc()
    return ColorListBox_GetAlign(c.instance)
}

func (c *TColorListBox) SetAlign(value TAlign) {
    defer exceptionProc()
    ColorListBox_SetAlign(c.instance, value)
}

func (c *TColorListBox) AutoComplete() bool {
    defer exceptionProc()
    return ColorListBox_GetAutoComplete(c.instance)
}

func (c *TColorListBox) SetAutoComplete(value bool) {
    defer exceptionProc()
    ColorListBox_SetAutoComplete(c.instance, value)
}

func (c *TColorListBox) DefaultColorColor() TColor {
    defer exceptionProc()
    return ColorListBox_GetDefaultColorColor(c.instance)
}

func (c *TColorListBox) SetDefaultColorColor(value TColor) {
    defer exceptionProc()
    ColorListBox_SetDefaultColorColor(c.instance, value)
}

func (c *TColorListBox) NoneColorColor() TColor {
    defer exceptionProc()
    return ColorListBox_GetNoneColorColor(c.instance)
}

func (c *TColorListBox) SetNoneColorColor(value TColor) {
    defer exceptionProc()
    ColorListBox_SetNoneColorColor(c.instance, value)
}

func (c *TColorListBox) Selected() TColor {
    defer exceptionProc()
    return ColorListBox_GetSelected(c.instance)
}

func (c *TColorListBox) SetSelected(value TColor) {
    defer exceptionProc()
    ColorListBox_SetSelected(c.instance, value)
}

func (c *TColorListBox) Style() TColorBoxStyle {
    defer exceptionProc()
    return ColorListBox_GetStyle(c.instance)
}

func (c *TColorListBox) SetStyle(value TColorBoxStyle) {
    defer exceptionProc()
    ColorListBox_SetStyle(c.instance, value)
}

func (c *TColorListBox) Anchors() TAnchors {
    defer exceptionProc()
    return ColorListBox_GetAnchors(c.instance)
}

func (c *TColorListBox) SetAnchors(value TAnchors) {
    defer exceptionProc()
    ColorListBox_SetAnchors(c.instance, value)
}

func (c *TColorListBox) Color() TColor {
    defer exceptionProc()
    return ColorListBox_GetColor(c.instance)
}

func (c *TColorListBox) SetColor(value TColor) {
    defer exceptionProc()
    ColorListBox_SetColor(c.instance, value)
}

func (c *TColorListBox) DoubleBuffered() bool {
    defer exceptionProc()
    return ColorListBox_GetDoubleBuffered(c.instance)
}

func (c *TColorListBox) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    ColorListBox_SetDoubleBuffered(c.instance, value)
}

func (c *TColorListBox) Enabled() bool {
    defer exceptionProc()
    return ColorListBox_GetEnabled(c.instance)
}

func (c *TColorListBox) SetEnabled(value bool) {
    defer exceptionProc()
    ColorListBox_SetEnabled(c.instance, value)
}

func (c *TColorListBox) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(ColorListBox_GetFont(c.instance))
}

func (c *TColorListBox) SetFont(value *TFont) {
    defer exceptionProc()
    ColorListBox_SetFont(c.instance, CheckPtr(value))
}

func (c *TColorListBox) ItemHeight() int32 {
    defer exceptionProc()
    return ColorListBox_GetItemHeight(c.instance)
}

func (c *TColorListBox) SetItemHeight(value int32) {
    defer exceptionProc()
    ColorListBox_SetItemHeight(c.instance, value)
}

func (c *TColorListBox) ParentColor() bool {
    defer exceptionProc()
    return ColorListBox_GetParentColor(c.instance)
}

func (c *TColorListBox) SetParentColor(value bool) {
    defer exceptionProc()
    ColorListBox_SetParentColor(c.instance, value)
}

func (c *TColorListBox) ParentCtl3D() bool {
    defer exceptionProc()
    return ColorListBox_GetParentCtl3D(c.instance)
}

func (c *TColorListBox) SetParentCtl3D(value bool) {
    defer exceptionProc()
    ColorListBox_SetParentCtl3D(c.instance, value)
}

func (c *TColorListBox) ParentFont() bool {
    defer exceptionProc()
    return ColorListBox_GetParentFont(c.instance)
}

func (c *TColorListBox) SetParentFont(value bool) {
    defer exceptionProc()
    ColorListBox_SetParentFont(c.instance, value)
}

func (c *TColorListBox) ParentShowHint() bool {
    defer exceptionProc()
    return ColorListBox_GetParentShowHint(c.instance)
}

func (c *TColorListBox) SetParentShowHint(value bool) {
    defer exceptionProc()
    ColorListBox_SetParentShowHint(c.instance, value)
}

func (c *TColorListBox) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(ColorListBox_GetPopupMenu(c.instance))
}

func (c *TColorListBox) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    ColorListBox_SetPopupMenu(c.instance, CheckPtr(value))
}

func (c *TColorListBox) ShowHint() bool {
    defer exceptionProc()
    return ColorListBox_GetShowHint(c.instance)
}

func (c *TColorListBox) SetShowHint(value bool) {
    defer exceptionProc()
    ColorListBox_SetShowHint(c.instance, value)
}

func (c *TColorListBox) TabOrder() int16 {
    defer exceptionProc()
    return ColorListBox_GetTabOrder(c.instance)
}

func (c *TColorListBox) SetTabOrder(value int16) {
    defer exceptionProc()
    ColorListBox_SetTabOrder(c.instance, value)
}

func (c *TColorListBox) TabStop() bool {
    defer exceptionProc()
    return ColorListBox_GetTabStop(c.instance)
}

func (c *TColorListBox) SetTabStop(value bool) {
    defer exceptionProc()
    ColorListBox_SetTabStop(c.instance, value)
}

func (c *TColorListBox) Visible() bool {
    defer exceptionProc()
    return ColorListBox_GetVisible(c.instance)
}

func (c *TColorListBox) SetVisible(value bool) {
    defer exceptionProc()
    ColorListBox_SetVisible(c.instance, value)
}

func (c *TColorListBox) SetOnClick(fn TNotifyEvent) {
    ColorListBox_SetOnClick(c.instance, fn)
}

func (c *TColorListBox) SetOnDblClick(fn TNotifyEvent) {
    ColorListBox_SetOnDblClick(c.instance, fn)
}

func (c *TColorListBox) SetOnEnter(fn TNotifyEvent) {
    ColorListBox_SetOnEnter(c.instance, fn)
}

func (c *TColorListBox) SetOnExit(fn TNotifyEvent) {
    ColorListBox_SetOnExit(c.instance, fn)
}

func (c *TColorListBox) SetOnKeyDown(fn TKeyEvent) {
    ColorListBox_SetOnKeyDown(c.instance, fn)
}

func (c *TColorListBox) SetOnKeyPress(fn TKeyPressEvent) {
    ColorListBox_SetOnKeyPress(c.instance, fn)
}

func (c *TColorListBox) SetOnKeyUp(fn TKeyEvent) {
    ColorListBox_SetOnKeyUp(c.instance, fn)
}

func (c *TColorListBox) SetOnMouseDown(fn TMouseEvent) {
    ColorListBox_SetOnMouseDown(c.instance, fn)
}

func (c *TColorListBox) SetOnMouseEnter(fn TNotifyEvent) {
    ColorListBox_SetOnMouseEnter(c.instance, fn)
}

func (c *TColorListBox) SetOnMouseLeave(fn TNotifyEvent) {
    ColorListBox_SetOnMouseLeave(c.instance, fn)
}

func (c *TColorListBox) SetOnMouseMove(fn TMouseMoveEvent) {
    ColorListBox_SetOnMouseMove(c.instance, fn)
}

func (c *TColorListBox) SetOnMouseUp(fn TMouseEvent) {
    ColorListBox_SetOnMouseUp(c.instance, fn)
}

func (c *TColorListBox) AutoCompleteDelay() uint32 {
    defer exceptionProc()
    return ColorListBox_GetAutoCompleteDelay(c.instance)
}

func (c *TColorListBox) SetAutoCompleteDelay(value uint32) {
    defer exceptionProc()
    ColorListBox_SetAutoCompleteDelay(c.instance, value)
}

func (c *TColorListBox) Canvas() *TCanvas {
    defer exceptionProc()
    return CanvasFromInst(ColorListBox_GetCanvas(c.instance))
}

func (c *TColorListBox) Items() *TStrings {
    defer exceptionProc()
    return StringsFromInst(ColorListBox_GetItems(c.instance))
}

func (c *TColorListBox) SetItems(value IObject) {
    defer exceptionProc()
    ColorListBox_SetItems(c.instance, CheckPtr(value))
}

func (c *TColorListBox) MultiSelect() bool {
    defer exceptionProc()
    return ColorListBox_GetMultiSelect(c.instance)
}

func (c *TColorListBox) SetMultiSelect(value bool) {
    defer exceptionProc()
    ColorListBox_SetMultiSelect(c.instance, value)
}

func (c *TColorListBox) SelCount() int32 {
    defer exceptionProc()
    return ColorListBox_GetSelCount(c.instance)
}

func (c *TColorListBox) ItemIndex() int32 {
    defer exceptionProc()
    return ColorListBox_GetItemIndex(c.instance)
}

func (c *TColorListBox) SetItemIndex(value int32) {
    defer exceptionProc()
    ColorListBox_SetItemIndex(c.instance, value)
}

func (c *TColorListBox) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(ColorListBox_GetBrush(c.instance))
}

func (c *TColorListBox) ControlCount() int32 {
    defer exceptionProc()
    return ColorListBox_GetControlCount(c.instance)
}

func (c *TColorListBox) Handle() HWND {
    defer exceptionProc()
    return ColorListBox_GetHandle(c.instance)
}

func (c *TColorListBox) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(ColorListBox_GetAction(c.instance))
}

func (c *TColorListBox) SetAction(value IComponent) {
    defer exceptionProc()
    ColorListBox_SetAction(c.instance, CheckPtr(value))
}

func (c *TColorListBox) BoundsRect() TRect {
    defer exceptionProc()
    return ColorListBox_GetBoundsRect(c.instance)
}

func (c *TColorListBox) SetBoundsRect(value TRect) {
    defer exceptionProc()
    ColorListBox_SetBoundsRect(c.instance, value)
}

func (c *TColorListBox) ClientHeight() int32 {
    defer exceptionProc()
    return ColorListBox_GetClientHeight(c.instance)
}

func (c *TColorListBox) SetClientHeight(value int32) {
    defer exceptionProc()
    ColorListBox_SetClientHeight(c.instance, value)
}

func (c *TColorListBox) ClientRect() TRect {
    defer exceptionProc()
    return ColorListBox_GetClientRect(c.instance)
}

func (c *TColorListBox) ClientWidth() int32 {
    defer exceptionProc()
    return ColorListBox_GetClientWidth(c.instance)
}

func (c *TColorListBox) SetClientWidth(value int32) {
    defer exceptionProc()
    ColorListBox_SetClientWidth(c.instance, value)
}

func (c *TColorListBox) ExplicitLeft() int32 {
    defer exceptionProc()
    return ColorListBox_GetExplicitLeft(c.instance)
}

func (c *TColorListBox) ExplicitTop() int32 {
    defer exceptionProc()
    return ColorListBox_GetExplicitTop(c.instance)
}

func (c *TColorListBox) ExplicitWidth() int32 {
    defer exceptionProc()
    return ColorListBox_GetExplicitWidth(c.instance)
}

func (c *TColorListBox) ExplicitHeight() int32 {
    defer exceptionProc()
    return ColorListBox_GetExplicitHeight(c.instance)
}

func (c *TColorListBox) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(ColorListBox_GetParent(c.instance))
}

func (c *TColorListBox) SetParent(value IControl) {
    defer exceptionProc()
    ColorListBox_SetParent(c.instance, CheckPtr(value))
}

func (c *TColorListBox) AlignWithMargins() bool {
    defer exceptionProc()
    return ColorListBox_GetAlignWithMargins(c.instance)
}

func (c *TColorListBox) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    ColorListBox_SetAlignWithMargins(c.instance, value)
}

func (c *TColorListBox) Left() int32 {
    defer exceptionProc()
    return ColorListBox_GetLeft(c.instance)
}

func (c *TColorListBox) SetLeft(value int32) {
    defer exceptionProc()
    ColorListBox_SetLeft(c.instance, value)
}

func (c *TColorListBox) Top() int32 {
    defer exceptionProc()
    return ColorListBox_GetTop(c.instance)
}

func (c *TColorListBox) SetTop(value int32) {
    defer exceptionProc()
    ColorListBox_SetTop(c.instance, value)
}

func (c *TColorListBox) Width() int32 {
    defer exceptionProc()
    return ColorListBox_GetWidth(c.instance)
}

func (c *TColorListBox) SetWidth(value int32) {
    defer exceptionProc()
    ColorListBox_SetWidth(c.instance, value)
}

func (c *TColorListBox) Height() int32 {
    defer exceptionProc()
    return ColorListBox_GetHeight(c.instance)
}

func (c *TColorListBox) SetHeight(value int32) {
    defer exceptionProc()
    ColorListBox_SetHeight(c.instance, value)
}

func (c *TColorListBox) Cursor() TCursor {
    defer exceptionProc()
    return ColorListBox_GetCursor(c.instance)
}

func (c *TColorListBox) SetCursor(value TCursor) {
    defer exceptionProc()
    ColorListBox_SetCursor(c.instance, value)
}

func (c *TColorListBox) Hint() string {
    defer exceptionProc()
    return ColorListBox_GetHint(c.instance)
}

func (c *TColorListBox) SetHint(value string) {
    defer exceptionProc()
    ColorListBox_SetHint(c.instance, value)
}

func (c *TColorListBox) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(ColorListBox_GetMargins(c.instance))
}

func (c *TColorListBox) SetMargins(value *TMargins) {
    defer exceptionProc()
    ColorListBox_SetMargins(c.instance, CheckPtr(value))
}

func (c *TColorListBox) ComponentCount() int32 {
    defer exceptionProc()
    return ColorListBox_GetComponentCount(c.instance)
}

func (c *TColorListBox) ComponentIndex() int32 {
    defer exceptionProc()
    return ColorListBox_GetComponentIndex(c.instance)
}

func (c *TColorListBox) SetComponentIndex(value int32) {
    defer exceptionProc()
    ColorListBox_SetComponentIndex(c.instance, value)
}

func (c *TColorListBox) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ColorListBox_GetOwner(c.instance))
}

func (c *TColorListBox) Name() string {
    defer exceptionProc()
    return ColorListBox_GetName(c.instance)
}

func (c *TColorListBox) SetName(value string) {
    defer exceptionProc()
    ColorListBox_SetName(c.instance, value)
}

func (c *TColorListBox) Tag() int {
    defer exceptionProc()
    return ColorListBox_GetTag(c.instance)
}

func (c *TColorListBox) SetTag(value int) {
    defer exceptionProc()
    ColorListBox_SetTag(c.instance, value)
}

func (c *TColorListBox) Colors(Index int32) TColor {
    defer exceptionProc()
    return ColorListBox_GetColors(c.instance, Index)
}

func (c *TColorListBox) ColorNames(Index int32) string {
    defer exceptionProc()
    return ColorListBox_GetColorNames(c.instance, Index)
}

func (c *TColorListBox) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(ColorListBox_GetControls(c.instance, Index))
}

func (c *TColorListBox) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ColorListBox_GetComponents(c.instance, AIndex))
}

