
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TColorBox struct {
    IControl
    instance uintptr
}

func NewColorBox(owner IComponent) *TColorBox {
    c := new(TColorBox)
    c.instance = ColorBox_Create(owner.Instance())
    return c
}

func ColorBoxFromInst(inst uintptr) *TColorBox {
    c := new(TColorBox)
    c.instance = inst
    return c
}

func ColorBoxFromObj(obj IObject) *TColorBox {
    c := new(TColorBox)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TColorBox) Free() {
    if c.instance != 0 {
        ColorBox_Free(c.instance)
    }
}

func (c *TColorBox) Instance() uintptr {
    return c.instance
}

func (c *TColorBox) IsValid() bool {
    return c.instance != 0
}

func (c *TColorBox) AddItem(Item string, AObject IObject) {
    defer exceptionProc()
    ColorBox_AddItem(c.instance, Item , CheckPtr(AObject))
}

func (c *TColorBox) Clear() {
    defer exceptionProc()
    ColorBox_Clear(c.instance)
}

func (c *TColorBox) ClearSelection() {
    defer exceptionProc()
    ColorBox_ClearSelection(c.instance)
}

func (c *TColorBox) DeleteSelected() {
    defer exceptionProc()
    ColorBox_DeleteSelected(c.instance)
}

func (c *TColorBox) Focused() bool {
    defer exceptionProc()
    return ColorBox_Focused(c.instance)
}

func (c *TColorBox) SelectAll() {
    defer exceptionProc()
    ColorBox_SelectAll(c.instance)
}

func (c *TColorBox) CanFocus() bool {
    defer exceptionProc()
    return ColorBox_CanFocus(c.instance)
}

func (c *TColorBox) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    ColorBox_FlipChildren(c.instance, AllLevels )
}

func (c *TColorBox) HandleAllocated() bool {
    defer exceptionProc()
    return ColorBox_HandleAllocated(c.instance)
}

func (c *TColorBox) Invalidate() {
    defer exceptionProc()
    ColorBox_Invalidate(c.instance)
}

func (c *TColorBox) Realign() {
    defer exceptionProc()
    ColorBox_Realign(c.instance)
}

func (c *TColorBox) Repaint() {
    defer exceptionProc()
    ColorBox_Repaint(c.instance)
}

func (c *TColorBox) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    ColorBox_ScaleBy(c.instance, M , D )
}

func (c *TColorBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    ColorBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight )
}

func (c *TColorBox) SetFocus() {
    defer exceptionProc()
    ColorBox_SetFocus(c.instance)
}

func (c *TColorBox) Update() {
    defer exceptionProc()
    ColorBox_Update(c.instance)
}

func (c *TColorBox) BringToFront() {
    defer exceptionProc()
    ColorBox_BringToFront(c.instance)
}

func (c *TColorBox) HasParent() bool {
    defer exceptionProc()
    return ColorBox_HasParent(c.instance)
}

func (c *TColorBox) Hide() {
    defer exceptionProc()
    ColorBox_Hide(c.instance)
}

func (c *TColorBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return ColorBox_Perform(c.instance, Msg , WParam , LParam )
}

func (c *TColorBox) Refresh() {
    defer exceptionProc()
    ColorBox_Refresh(c.instance)
}

func (c *TColorBox) SendToBack() {
    defer exceptionProc()
    ColorBox_SendToBack(c.instance)
}

func (c *TColorBox) Show() {
    defer exceptionProc()
    ColorBox_Show(c.instance)
}

func (c *TColorBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return ColorBox_GetTextBuf(c.instance, Buffer , BufSize )
}

func (c *TColorBox) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ColorBox_FindComponent(c.instance, AName ))
}

func (c *TColorBox) GetNamePath() string {
    defer exceptionProc()
    return ColorBox_GetNamePath(c.instance)
}

func (c *TColorBox) Assign(Source IObject) {
    defer exceptionProc()
    ColorBox_Assign(c.instance, CheckPtr(Source))
}

func (c *TColorBox) ClassName() string {
    defer exceptionProc()
    return ColorBox_ClassName(c.instance)
}

func (c *TColorBox) Equals(Obj IObject) bool {
    defer exceptionProc()
    return ColorBox_Equals(c.instance, CheckPtr(Obj))
}

func (c *TColorBox) GetHashCode() int32 {
    defer exceptionProc()
    return ColorBox_GetHashCode(c.instance)
}

func (c *TColorBox) ToString() string {
    defer exceptionProc()
    return ColorBox_ToString(c.instance)
}

func (c *TColorBox) Align() TAlign {
    defer exceptionProc()
    return ColorBox_GetAlign(c.instance)
}

func (c *TColorBox) SetAlign(value TAlign) {
    defer exceptionProc()
    ColorBox_SetAlign(c.instance, value)
}

func (c *TColorBox) AutoComplete() bool {
    defer exceptionProc()
    return ColorBox_GetAutoComplete(c.instance)
}

func (c *TColorBox) SetAutoComplete(value bool) {
    defer exceptionProc()
    ColorBox_SetAutoComplete(c.instance, value)
}

func (c *TColorBox) AutoDropDown() bool {
    defer exceptionProc()
    return ColorBox_GetAutoDropDown(c.instance)
}

func (c *TColorBox) SetAutoDropDown(value bool) {
    defer exceptionProc()
    ColorBox_SetAutoDropDown(c.instance, value)
}

func (c *TColorBox) DefaultColorColor() TColor {
    defer exceptionProc()
    return ColorBox_GetDefaultColorColor(c.instance)
}

func (c *TColorBox) SetDefaultColorColor(value TColor) {
    defer exceptionProc()
    ColorBox_SetDefaultColorColor(c.instance, value)
}

func (c *TColorBox) NoneColorColor() TColor {
    defer exceptionProc()
    return ColorBox_GetNoneColorColor(c.instance)
}

func (c *TColorBox) SetNoneColorColor(value TColor) {
    defer exceptionProc()
    ColorBox_SetNoneColorColor(c.instance, value)
}

func (c *TColorBox) Selected() TColor {
    defer exceptionProc()
    return ColorBox_GetSelected(c.instance)
}

func (c *TColorBox) SetSelected(value TColor) {
    defer exceptionProc()
    ColorBox_SetSelected(c.instance, value)
}

func (c *TColorBox) Style() TColorBoxStyle {
    defer exceptionProc()
    return ColorBox_GetStyle(c.instance)
}

func (c *TColorBox) SetStyle(value TColorBoxStyle) {
    defer exceptionProc()
    ColorBox_SetStyle(c.instance, value)
}

func (c *TColorBox) Anchors() TAnchors {
    defer exceptionProc()
    return ColorBox_GetAnchors(c.instance)
}

func (c *TColorBox) SetAnchors(value TAnchors) {
    defer exceptionProc()
    ColorBox_SetAnchors(c.instance, value)
}

func (c *TColorBox) Color() TColor {
    defer exceptionProc()
    return ColorBox_GetColor(c.instance)
}

func (c *TColorBox) SetColor(value TColor) {
    defer exceptionProc()
    ColorBox_SetColor(c.instance, value)
}

func (c *TColorBox) DoubleBuffered() bool {
    defer exceptionProc()
    return ColorBox_GetDoubleBuffered(c.instance)
}

func (c *TColorBox) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    ColorBox_SetDoubleBuffered(c.instance, value)
}

func (c *TColorBox) DropDownCount() int32 {
    defer exceptionProc()
    return ColorBox_GetDropDownCount(c.instance)
}

func (c *TColorBox) SetDropDownCount(value int32) {
    defer exceptionProc()
    ColorBox_SetDropDownCount(c.instance, value)
}

func (c *TColorBox) Enabled() bool {
    defer exceptionProc()
    return ColorBox_GetEnabled(c.instance)
}

func (c *TColorBox) SetEnabled(value bool) {
    defer exceptionProc()
    ColorBox_SetEnabled(c.instance, value)
}

func (c *TColorBox) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(ColorBox_GetFont(c.instance))
}

func (c *TColorBox) SetFont(value *TFont) {
    defer exceptionProc()
    ColorBox_SetFont(c.instance, CheckPtr(value))
}

func (c *TColorBox) ItemHeight() int32 {
    defer exceptionProc()
    return ColorBox_GetItemHeight(c.instance)
}

func (c *TColorBox) SetItemHeight(value int32) {
    defer exceptionProc()
    ColorBox_SetItemHeight(c.instance, value)
}

func (c *TColorBox) ParentColor() bool {
    defer exceptionProc()
    return ColorBox_GetParentColor(c.instance)
}

func (c *TColorBox) SetParentColor(value bool) {
    defer exceptionProc()
    ColorBox_SetParentColor(c.instance, value)
}

func (c *TColorBox) ParentCtl3D() bool {
    defer exceptionProc()
    return ColorBox_GetParentCtl3D(c.instance)
}

func (c *TColorBox) SetParentCtl3D(value bool) {
    defer exceptionProc()
    ColorBox_SetParentCtl3D(c.instance, value)
}

func (c *TColorBox) ParentFont() bool {
    defer exceptionProc()
    return ColorBox_GetParentFont(c.instance)
}

func (c *TColorBox) SetParentFont(value bool) {
    defer exceptionProc()
    ColorBox_SetParentFont(c.instance, value)
}

func (c *TColorBox) ParentShowHint() bool {
    defer exceptionProc()
    return ColorBox_GetParentShowHint(c.instance)
}

func (c *TColorBox) SetParentShowHint(value bool) {
    defer exceptionProc()
    ColorBox_SetParentShowHint(c.instance, value)
}

func (c *TColorBox) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(ColorBox_GetPopupMenu(c.instance))
}

func (c *TColorBox) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    ColorBox_SetPopupMenu(c.instance, CheckPtr(value))
}

func (c *TColorBox) ShowHint() bool {
    defer exceptionProc()
    return ColorBox_GetShowHint(c.instance)
}

func (c *TColorBox) SetShowHint(value bool) {
    defer exceptionProc()
    ColorBox_SetShowHint(c.instance, value)
}

func (c *TColorBox) TabOrder() int16 {
    defer exceptionProc()
    return ColorBox_GetTabOrder(c.instance)
}

func (c *TColorBox) SetTabOrder(value int16) {
    defer exceptionProc()
    ColorBox_SetTabOrder(c.instance, value)
}

func (c *TColorBox) TabStop() bool {
    defer exceptionProc()
    return ColorBox_GetTabStop(c.instance)
}

func (c *TColorBox) SetTabStop(value bool) {
    defer exceptionProc()
    ColorBox_SetTabStop(c.instance, value)
}

func (c *TColorBox) Visible() bool {
    defer exceptionProc()
    return ColorBox_GetVisible(c.instance)
}

func (c *TColorBox) SetVisible(value bool) {
    defer exceptionProc()
    ColorBox_SetVisible(c.instance, value)
}

func (c *TColorBox) SetOnChange(fn TNotifyEvent) {
    ColorBox_SetOnChange(c.instance, fn)
}

func (c *TColorBox) SetOnClick(fn TNotifyEvent) {
    ColorBox_SetOnClick(c.instance, fn)
}

func (c *TColorBox) SetOnEnter(fn TNotifyEvent) {
    ColorBox_SetOnEnter(c.instance, fn)
}

func (c *TColorBox) SetOnExit(fn TNotifyEvent) {
    ColorBox_SetOnExit(c.instance, fn)
}

func (c *TColorBox) SetOnKeyDown(fn TKeyEvent) {
    ColorBox_SetOnKeyDown(c.instance, fn)
}

func (c *TColorBox) SetOnKeyPress(fn TKeyPressEvent) {
    ColorBox_SetOnKeyPress(c.instance, fn)
}

func (c *TColorBox) SetOnKeyUp(fn TKeyEvent) {
    ColorBox_SetOnKeyUp(c.instance, fn)
}

func (c *TColorBox) SetOnMouseEnter(fn TNotifyEvent) {
    ColorBox_SetOnMouseEnter(c.instance, fn)
}

func (c *TColorBox) SetOnMouseLeave(fn TNotifyEvent) {
    ColorBox_SetOnMouseLeave(c.instance, fn)
}

func (c *TColorBox) AutoCompleteDelay() uint32 {
    defer exceptionProc()
    return ColorBox_GetAutoCompleteDelay(c.instance)
}

func (c *TColorBox) SetAutoCompleteDelay(value uint32) {
    defer exceptionProc()
    ColorBox_SetAutoCompleteDelay(c.instance, value)
}

func (c *TColorBox) AutoCloseUp() bool {
    defer exceptionProc()
    return ColorBox_GetAutoCloseUp(c.instance)
}

func (c *TColorBox) SetAutoCloseUp(value bool) {
    defer exceptionProc()
    ColorBox_SetAutoCloseUp(c.instance, value)
}

func (c *TColorBox) SelText() string {
    defer exceptionProc()
    return ColorBox_GetSelText(c.instance)
}

func (c *TColorBox) SetSelText(value string) {
    defer exceptionProc()
    ColorBox_SetSelText(c.instance, value)
}

func (c *TColorBox) TextHint() string {
    defer exceptionProc()
    return ColorBox_GetTextHint(c.instance)
}

func (c *TColorBox) SetTextHint(value string) {
    defer exceptionProc()
    ColorBox_SetTextHint(c.instance, value)
}

func (c *TColorBox) Canvas() *TCanvas {
    defer exceptionProc()
    return CanvasFromInst(ColorBox_GetCanvas(c.instance))
}

func (c *TColorBox) Items() *TStrings {
    defer exceptionProc()
    return StringsFromInst(ColorBox_GetItems(c.instance))
}

func (c *TColorBox) SetItems(value IObject) {
    defer exceptionProc()
    ColorBox_SetItems(c.instance, CheckPtr(value))
}

func (c *TColorBox) SelLength() int32 {
    defer exceptionProc()
    return ColorBox_GetSelLength(c.instance)
}

func (c *TColorBox) SetSelLength(value int32) {
    defer exceptionProc()
    ColorBox_SetSelLength(c.instance, value)
}

func (c *TColorBox) SelStart() int32 {
    defer exceptionProc()
    return ColorBox_GetSelStart(c.instance)
}

func (c *TColorBox) SetSelStart(value int32) {
    defer exceptionProc()
    ColorBox_SetSelStart(c.instance, value)
}

func (c *TColorBox) ItemIndex() int32 {
    defer exceptionProc()
    return ColorBox_GetItemIndex(c.instance)
}

func (c *TColorBox) SetItemIndex(value int32) {
    defer exceptionProc()
    ColorBox_SetItemIndex(c.instance, value)
}

func (c *TColorBox) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(ColorBox_GetBrush(c.instance))
}

func (c *TColorBox) ControlCount() int32 {
    defer exceptionProc()
    return ColorBox_GetControlCount(c.instance)
}

func (c *TColorBox) Handle() HWND {
    defer exceptionProc()
    return ColorBox_GetHandle(c.instance)
}

func (c *TColorBox) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(ColorBox_GetAction(c.instance))
}

func (c *TColorBox) SetAction(value IComponent) {
    defer exceptionProc()
    ColorBox_SetAction(c.instance, CheckPtr(value))
}

func (c *TColorBox) BoundsRect() TRect {
    defer exceptionProc()
    return ColorBox_GetBoundsRect(c.instance)
}

func (c *TColorBox) SetBoundsRect(value TRect) {
    defer exceptionProc()
    ColorBox_SetBoundsRect(c.instance, value)
}

func (c *TColorBox) ClientHeight() int32 {
    defer exceptionProc()
    return ColorBox_GetClientHeight(c.instance)
}

func (c *TColorBox) SetClientHeight(value int32) {
    defer exceptionProc()
    ColorBox_SetClientHeight(c.instance, value)
}

func (c *TColorBox) ClientRect() TRect {
    defer exceptionProc()
    return ColorBox_GetClientRect(c.instance)
}

func (c *TColorBox) ClientWidth() int32 {
    defer exceptionProc()
    return ColorBox_GetClientWidth(c.instance)
}

func (c *TColorBox) SetClientWidth(value int32) {
    defer exceptionProc()
    ColorBox_SetClientWidth(c.instance, value)
}

func (c *TColorBox) ExplicitLeft() int32 {
    defer exceptionProc()
    return ColorBox_GetExplicitLeft(c.instance)
}

func (c *TColorBox) ExplicitTop() int32 {
    defer exceptionProc()
    return ColorBox_GetExplicitTop(c.instance)
}

func (c *TColorBox) ExplicitWidth() int32 {
    defer exceptionProc()
    return ColorBox_GetExplicitWidth(c.instance)
}

func (c *TColorBox) ExplicitHeight() int32 {
    defer exceptionProc()
    return ColorBox_GetExplicitHeight(c.instance)
}

func (c *TColorBox) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(ColorBox_GetParent(c.instance))
}

func (c *TColorBox) SetParent(value IControl) {
    defer exceptionProc()
    ColorBox_SetParent(c.instance, CheckPtr(value))
}

func (c *TColorBox) AlignWithMargins() bool {
    defer exceptionProc()
    return ColorBox_GetAlignWithMargins(c.instance)
}

func (c *TColorBox) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    ColorBox_SetAlignWithMargins(c.instance, value)
}

func (c *TColorBox) Left() int32 {
    defer exceptionProc()
    return ColorBox_GetLeft(c.instance)
}

func (c *TColorBox) SetLeft(value int32) {
    defer exceptionProc()
    ColorBox_SetLeft(c.instance, value)
}

func (c *TColorBox) Top() int32 {
    defer exceptionProc()
    return ColorBox_GetTop(c.instance)
}

func (c *TColorBox) SetTop(value int32) {
    defer exceptionProc()
    ColorBox_SetTop(c.instance, value)
}

func (c *TColorBox) Width() int32 {
    defer exceptionProc()
    return ColorBox_GetWidth(c.instance)
}

func (c *TColorBox) SetWidth(value int32) {
    defer exceptionProc()
    ColorBox_SetWidth(c.instance, value)
}

func (c *TColorBox) Height() int32 {
    defer exceptionProc()
    return ColorBox_GetHeight(c.instance)
}

func (c *TColorBox) SetHeight(value int32) {
    defer exceptionProc()
    ColorBox_SetHeight(c.instance, value)
}

func (c *TColorBox) Cursor() TCursor {
    defer exceptionProc()
    return ColorBox_GetCursor(c.instance)
}

func (c *TColorBox) SetCursor(value TCursor) {
    defer exceptionProc()
    ColorBox_SetCursor(c.instance, value)
}

func (c *TColorBox) Hint() string {
    defer exceptionProc()
    return ColorBox_GetHint(c.instance)
}

func (c *TColorBox) SetHint(value string) {
    defer exceptionProc()
    ColorBox_SetHint(c.instance, value)
}

func (c *TColorBox) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(ColorBox_GetMargins(c.instance))
}

func (c *TColorBox) SetMargins(value *TMargins) {
    defer exceptionProc()
    ColorBox_SetMargins(c.instance, CheckPtr(value))
}

func (c *TColorBox) ComponentCount() int32 {
    defer exceptionProc()
    return ColorBox_GetComponentCount(c.instance)
}

func (c *TColorBox) ComponentIndex() int32 {
    defer exceptionProc()
    return ColorBox_GetComponentIndex(c.instance)
}

func (c *TColorBox) SetComponentIndex(value int32) {
    defer exceptionProc()
    ColorBox_SetComponentIndex(c.instance, value)
}

func (c *TColorBox) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ColorBox_GetOwner(c.instance))
}

func (c *TColorBox) Name() string {
    defer exceptionProc()
    return ColorBox_GetName(c.instance)
}

func (c *TColorBox) SetName(value string) {
    defer exceptionProc()
    ColorBox_SetName(c.instance, value)
}

func (c *TColorBox) Tag() int {
    defer exceptionProc()
    return ColorBox_GetTag(c.instance)
}

func (c *TColorBox) SetTag(value int) {
    defer exceptionProc()
    ColorBox_SetTag(c.instance, value)
}

func (c *TColorBox) Colors(Index int32) TColor {
    defer exceptionProc()
    return ColorBox_GetColors(c.instance, Index)
}

func (c *TColorBox) ColorNames(Index int32) string {
    defer exceptionProc()
    return ColorBox_GetColorNames(c.instance, Index)
}

func (c *TColorBox) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(ColorBox_GetControls(c.instance, Index))
}

func (c *TColorBox) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ColorBox_GetComponents(c.instance, AIndex))
}

