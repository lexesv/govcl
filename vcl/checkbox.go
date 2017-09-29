
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

type TCheckBox struct {
    IControl
    instance uintptr
}

func NewCheckBox(owner IComponent) *TCheckBox {
    c := new(TCheckBox)
    c.instance = CheckBox_Create(owner.Instance())
    return c
}

func CheckBoxFromInst(inst uintptr) *TCheckBox {
    c := new(TCheckBox)
    c.instance = inst
    return c
}

func CheckBoxFromObj(obj IObject) *TCheckBox {
    c := new(TCheckBox)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TCheckBox) Free() {
    if c.instance != 0 {
        CheckBox_Free(c.instance)
    }
}

func (c *TCheckBox) Instance() uintptr {
    return c.instance
}

func (c *TCheckBox) IsValid() bool {
    return c.instance != 0
}

func (c *TCheckBox) CanFocus() bool {
    defer exceptionProc()
    return CheckBox_CanFocus(c.instance)
}

func (c *TCheckBox) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    CheckBox_FlipChildren(c.instance, AllLevels )
}

func (c *TCheckBox) Focused() bool {
    defer exceptionProc()
    return CheckBox_Focused(c.instance)
}

func (c *TCheckBox) HandleAllocated() bool {
    defer exceptionProc()
    return CheckBox_HandleAllocated(c.instance)
}

func (c *TCheckBox) Invalidate() {
    defer exceptionProc()
    CheckBox_Invalidate(c.instance)
}

func (c *TCheckBox) Realign() {
    defer exceptionProc()
    CheckBox_Realign(c.instance)
}

func (c *TCheckBox) Repaint() {
    defer exceptionProc()
    CheckBox_Repaint(c.instance)
}

func (c *TCheckBox) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    CheckBox_ScaleBy(c.instance, M , D )
}

func (c *TCheckBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    CheckBox_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight )
}

func (c *TCheckBox) SetFocus() {
    defer exceptionProc()
    CheckBox_SetFocus(c.instance)
}

func (c *TCheckBox) Update() {
    defer exceptionProc()
    CheckBox_Update(c.instance)
}

func (c *TCheckBox) BringToFront() {
    defer exceptionProc()
    CheckBox_BringToFront(c.instance)
}

func (c *TCheckBox) HasParent() bool {
    defer exceptionProc()
    return CheckBox_HasParent(c.instance)
}

func (c *TCheckBox) Hide() {
    defer exceptionProc()
    CheckBox_Hide(c.instance)
}

func (c *TCheckBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return CheckBox_Perform(c.instance, Msg , WParam , LParam )
}

func (c *TCheckBox) Refresh() {
    defer exceptionProc()
    CheckBox_Refresh(c.instance)
}

func (c *TCheckBox) SendToBack() {
    defer exceptionProc()
    CheckBox_SendToBack(c.instance)
}

func (c *TCheckBox) Show() {
    defer exceptionProc()
    CheckBox_Show(c.instance)
}

func (c *TCheckBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return CheckBox_GetTextBuf(c.instance, Buffer , BufSize )
}

func (c *TCheckBox) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(CheckBox_FindComponent(c.instance, AName ))
}

func (c *TCheckBox) GetNamePath() string {
    defer exceptionProc()
    return CheckBox_GetNamePath(c.instance)
}

func (c *TCheckBox) Assign(Source IObject) {
    defer exceptionProc()
    CheckBox_Assign(c.instance, CheckPtr(Source))
}

func (c *TCheckBox) ClassName() string {
    defer exceptionProc()
    return CheckBox_ClassName(c.instance)
}

func (c *TCheckBox) Equals(Obj IObject) bool {
    defer exceptionProc()
    return CheckBox_Equals(c.instance, CheckPtr(Obj))
}

func (c *TCheckBox) GetHashCode() int32 {
    defer exceptionProc()
    return CheckBox_GetHashCode(c.instance)
}

func (c *TCheckBox) ToString() string {
    defer exceptionProc()
    return CheckBox_ToString(c.instance)
}

func (c *TCheckBox) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(CheckBox_GetAction(c.instance))
}

func (c *TCheckBox) SetAction(value IComponent) {
    defer exceptionProc()
    CheckBox_SetAction(c.instance, CheckPtr(value))
}

func (c *TCheckBox) Align() TAlign {
    defer exceptionProc()
    return CheckBox_GetAlign(c.instance)
}

func (c *TCheckBox) SetAlign(value TAlign) {
    defer exceptionProc()
    CheckBox_SetAlign(c.instance, value)
}

func (c *TCheckBox) Alignment() TLeftRight {
    defer exceptionProc()
    return CheckBox_GetAlignment(c.instance)
}

func (c *TCheckBox) SetAlignment(value TLeftRight) {
    defer exceptionProc()
    CheckBox_SetAlignment(c.instance, value)
}

func (c *TCheckBox) Anchors() TAnchors {
    defer exceptionProc()
    return CheckBox_GetAnchors(c.instance)
}

func (c *TCheckBox) SetAnchors(value TAnchors) {
    defer exceptionProc()
    CheckBox_SetAnchors(c.instance, value)
}

func (c *TCheckBox) Caption() string {
    defer exceptionProc()
    return CheckBox_GetCaption(c.instance)
}

func (c *TCheckBox) SetCaption(value string) {
    defer exceptionProc()
    CheckBox_SetCaption(c.instance, value)
}

func (c *TCheckBox) Checked() bool {
    defer exceptionProc()
    return CheckBox_GetChecked(c.instance)
}

func (c *TCheckBox) SetChecked(value bool) {
    defer exceptionProc()
    CheckBox_SetChecked(c.instance, value)
}

func (c *TCheckBox) Color() TColor {
    defer exceptionProc()
    return CheckBox_GetColor(c.instance)
}

func (c *TCheckBox) SetColor(value TColor) {
    defer exceptionProc()
    CheckBox_SetColor(c.instance, value)
}

func (c *TCheckBox) DoubleBuffered() bool {
    defer exceptionProc()
    return CheckBox_GetDoubleBuffered(c.instance)
}

func (c *TCheckBox) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    CheckBox_SetDoubleBuffered(c.instance, value)
}

func (c *TCheckBox) Enabled() bool {
    defer exceptionProc()
    return CheckBox_GetEnabled(c.instance)
}

func (c *TCheckBox) SetEnabled(value bool) {
    defer exceptionProc()
    CheckBox_SetEnabled(c.instance, value)
}

func (c *TCheckBox) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(CheckBox_GetFont(c.instance))
}

func (c *TCheckBox) SetFont(value *TFont) {
    defer exceptionProc()
    CheckBox_SetFont(c.instance, CheckPtr(value))
}

func (c *TCheckBox) ParentColor() bool {
    defer exceptionProc()
    return CheckBox_GetParentColor(c.instance)
}

func (c *TCheckBox) SetParentColor(value bool) {
    defer exceptionProc()
    CheckBox_SetParentColor(c.instance, value)
}

func (c *TCheckBox) ParentCtl3D() bool {
    defer exceptionProc()
    return CheckBox_GetParentCtl3D(c.instance)
}

func (c *TCheckBox) SetParentCtl3D(value bool) {
    defer exceptionProc()
    CheckBox_SetParentCtl3D(c.instance, value)
}

func (c *TCheckBox) ParentFont() bool {
    defer exceptionProc()
    return CheckBox_GetParentFont(c.instance)
}

func (c *TCheckBox) SetParentFont(value bool) {
    defer exceptionProc()
    CheckBox_SetParentFont(c.instance, value)
}

func (c *TCheckBox) ParentShowHint() bool {
    defer exceptionProc()
    return CheckBox_GetParentShowHint(c.instance)
}

func (c *TCheckBox) SetParentShowHint(value bool) {
    defer exceptionProc()
    CheckBox_SetParentShowHint(c.instance, value)
}

func (c *TCheckBox) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(CheckBox_GetPopupMenu(c.instance))
}

func (c *TCheckBox) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    CheckBox_SetPopupMenu(c.instance, CheckPtr(value))
}

func (c *TCheckBox) ShowHint() bool {
    defer exceptionProc()
    return CheckBox_GetShowHint(c.instance)
}

func (c *TCheckBox) SetShowHint(value bool) {
    defer exceptionProc()
    CheckBox_SetShowHint(c.instance, value)
}

func (c *TCheckBox) State() TCheckBoxState {
    defer exceptionProc()
    return CheckBox_GetState(c.instance)
}

func (c *TCheckBox) SetState(value TCheckBoxState) {
    defer exceptionProc()
    CheckBox_SetState(c.instance, value)
}

func (c *TCheckBox) TabOrder() int16 {
    defer exceptionProc()
    return CheckBox_GetTabOrder(c.instance)
}

func (c *TCheckBox) SetTabOrder(value int16) {
    defer exceptionProc()
    CheckBox_SetTabOrder(c.instance, value)
}

func (c *TCheckBox) TabStop() bool {
    defer exceptionProc()
    return CheckBox_GetTabStop(c.instance)
}

func (c *TCheckBox) SetTabStop(value bool) {
    defer exceptionProc()
    CheckBox_SetTabStop(c.instance, value)
}

func (c *TCheckBox) Visible() bool {
    defer exceptionProc()
    return CheckBox_GetVisible(c.instance)
}

func (c *TCheckBox) SetVisible(value bool) {
    defer exceptionProc()
    CheckBox_SetVisible(c.instance, value)
}

func (c *TCheckBox) WordWrap() bool {
    defer exceptionProc()
    return CheckBox_GetWordWrap(c.instance)
}

func (c *TCheckBox) SetWordWrap(value bool) {
    defer exceptionProc()
    CheckBox_SetWordWrap(c.instance, value)
}

func (c *TCheckBox) SetOnClick(fn TNotifyEvent) {
    CheckBox_SetOnClick(c.instance, fn)
}

func (c *TCheckBox) SetOnEnter(fn TNotifyEvent) {
    CheckBox_SetOnEnter(c.instance, fn)
}

func (c *TCheckBox) SetOnExit(fn TNotifyEvent) {
    CheckBox_SetOnExit(c.instance, fn)
}

func (c *TCheckBox) SetOnKeyDown(fn TKeyEvent) {
    CheckBox_SetOnKeyDown(c.instance, fn)
}

func (c *TCheckBox) SetOnKeyPress(fn TKeyPressEvent) {
    CheckBox_SetOnKeyPress(c.instance, fn)
}

func (c *TCheckBox) SetOnKeyUp(fn TKeyEvent) {
    CheckBox_SetOnKeyUp(c.instance, fn)
}

func (c *TCheckBox) SetOnMouseDown(fn TMouseEvent) {
    CheckBox_SetOnMouseDown(c.instance, fn)
}

func (c *TCheckBox) SetOnMouseEnter(fn TNotifyEvent) {
    CheckBox_SetOnMouseEnter(c.instance, fn)
}

func (c *TCheckBox) SetOnMouseLeave(fn TNotifyEvent) {
    CheckBox_SetOnMouseLeave(c.instance, fn)
}

func (c *TCheckBox) SetOnMouseMove(fn TMouseMoveEvent) {
    CheckBox_SetOnMouseMove(c.instance, fn)
}

func (c *TCheckBox) SetOnMouseUp(fn TMouseEvent) {
    CheckBox_SetOnMouseUp(c.instance, fn)
}

func (c *TCheckBox) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(CheckBox_GetBrush(c.instance))
}

func (c *TCheckBox) ControlCount() int32 {
    defer exceptionProc()
    return CheckBox_GetControlCount(c.instance)
}

func (c *TCheckBox) Handle() HWND {
    defer exceptionProc()
    return CheckBox_GetHandle(c.instance)
}

func (c *TCheckBox) BoundsRect() TRect {
    defer exceptionProc()
    return CheckBox_GetBoundsRect(c.instance)
}

func (c *TCheckBox) SetBoundsRect(value TRect) {
    defer exceptionProc()
    CheckBox_SetBoundsRect(c.instance, value)
}

func (c *TCheckBox) ClientHeight() int32 {
    defer exceptionProc()
    return CheckBox_GetClientHeight(c.instance)
}

func (c *TCheckBox) SetClientHeight(value int32) {
    defer exceptionProc()
    CheckBox_SetClientHeight(c.instance, value)
}

func (c *TCheckBox) ClientRect() TRect {
    defer exceptionProc()
    return CheckBox_GetClientRect(c.instance)
}

func (c *TCheckBox) ClientWidth() int32 {
    defer exceptionProc()
    return CheckBox_GetClientWidth(c.instance)
}

func (c *TCheckBox) SetClientWidth(value int32) {
    defer exceptionProc()
    CheckBox_SetClientWidth(c.instance, value)
}

func (c *TCheckBox) ExplicitLeft() int32 {
    defer exceptionProc()
    return CheckBox_GetExplicitLeft(c.instance)
}

func (c *TCheckBox) ExplicitTop() int32 {
    defer exceptionProc()
    return CheckBox_GetExplicitTop(c.instance)
}

func (c *TCheckBox) ExplicitWidth() int32 {
    defer exceptionProc()
    return CheckBox_GetExplicitWidth(c.instance)
}

func (c *TCheckBox) ExplicitHeight() int32 {
    defer exceptionProc()
    return CheckBox_GetExplicitHeight(c.instance)
}

func (c *TCheckBox) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(CheckBox_GetParent(c.instance))
}

func (c *TCheckBox) SetParent(value IControl) {
    defer exceptionProc()
    CheckBox_SetParent(c.instance, CheckPtr(value))
}

func (c *TCheckBox) AlignWithMargins() bool {
    defer exceptionProc()
    return CheckBox_GetAlignWithMargins(c.instance)
}

func (c *TCheckBox) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    CheckBox_SetAlignWithMargins(c.instance, value)
}

func (c *TCheckBox) Left() int32 {
    defer exceptionProc()
    return CheckBox_GetLeft(c.instance)
}

func (c *TCheckBox) SetLeft(value int32) {
    defer exceptionProc()
    CheckBox_SetLeft(c.instance, value)
}

func (c *TCheckBox) Top() int32 {
    defer exceptionProc()
    return CheckBox_GetTop(c.instance)
}

func (c *TCheckBox) SetTop(value int32) {
    defer exceptionProc()
    CheckBox_SetTop(c.instance, value)
}

func (c *TCheckBox) Width() int32 {
    defer exceptionProc()
    return CheckBox_GetWidth(c.instance)
}

func (c *TCheckBox) SetWidth(value int32) {
    defer exceptionProc()
    CheckBox_SetWidth(c.instance, value)
}

func (c *TCheckBox) Height() int32 {
    defer exceptionProc()
    return CheckBox_GetHeight(c.instance)
}

func (c *TCheckBox) SetHeight(value int32) {
    defer exceptionProc()
    CheckBox_SetHeight(c.instance, value)
}

func (c *TCheckBox) Cursor() TCursor {
    defer exceptionProc()
    return CheckBox_GetCursor(c.instance)
}

func (c *TCheckBox) SetCursor(value TCursor) {
    defer exceptionProc()
    CheckBox_SetCursor(c.instance, value)
}

func (c *TCheckBox) Hint() string {
    defer exceptionProc()
    return CheckBox_GetHint(c.instance)
}

func (c *TCheckBox) SetHint(value string) {
    defer exceptionProc()
    CheckBox_SetHint(c.instance, value)
}

func (c *TCheckBox) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(CheckBox_GetMargins(c.instance))
}

func (c *TCheckBox) SetMargins(value *TMargins) {
    defer exceptionProc()
    CheckBox_SetMargins(c.instance, CheckPtr(value))
}

func (c *TCheckBox) ComponentCount() int32 {
    defer exceptionProc()
    return CheckBox_GetComponentCount(c.instance)
}

func (c *TCheckBox) ComponentIndex() int32 {
    defer exceptionProc()
    return CheckBox_GetComponentIndex(c.instance)
}

func (c *TCheckBox) SetComponentIndex(value int32) {
    defer exceptionProc()
    CheckBox_SetComponentIndex(c.instance, value)
}

func (c *TCheckBox) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(CheckBox_GetOwner(c.instance))
}

func (c *TCheckBox) Name() string {
    defer exceptionProc()
    return CheckBox_GetName(c.instance)
}

func (c *TCheckBox) SetName(value string) {
    defer exceptionProc()
    CheckBox_SetName(c.instance, value)
}

func (c *TCheckBox) Tag() int {
    defer exceptionProc()
    return CheckBox_GetTag(c.instance)
}

func (c *TCheckBox) SetTag(value int) {
    defer exceptionProc()
    CheckBox_SetTag(c.instance, value)
}

func (c *TCheckBox) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(CheckBox_GetControls(c.instance, Index))
}

func (c *TCheckBox) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(CheckBox_GetComponents(c.instance, AIndex))
}

