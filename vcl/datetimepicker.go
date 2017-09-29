
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

type TDateTimePicker struct {
    IControl
    instance uintptr
}

func NewDateTimePicker(owner IComponent) *TDateTimePicker {
    d := new(TDateTimePicker)
    d.instance = DateTimePicker_Create(owner.Instance())
    return d
}

func DateTimePickerFromInst(inst uintptr) *TDateTimePicker {
    d := new(TDateTimePicker)
    d.instance = inst
    return d
}

func DateTimePickerFromObj(obj IObject) *TDateTimePicker {
    d := new(TDateTimePicker)
    d.instance = CheckPtr(obj)
    return d
}

func (d *TDateTimePicker) Free() {
    if d.instance != 0 {
        DateTimePicker_Free(d.instance)
    }
}

func (d *TDateTimePicker) Instance() uintptr {
    return d.instance
}

func (d *TDateTimePicker) IsValid() bool {
    return d.instance != 0
}

func (d *TDateTimePicker) CanFocus() bool {
    defer exceptionProc()
    return DateTimePicker_CanFocus(d.instance)
}

func (d *TDateTimePicker) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    DateTimePicker_FlipChildren(d.instance, AllLevels )
}

func (d *TDateTimePicker) Focused() bool {
    defer exceptionProc()
    return DateTimePicker_Focused(d.instance)
}

func (d *TDateTimePicker) HandleAllocated() bool {
    defer exceptionProc()
    return DateTimePicker_HandleAllocated(d.instance)
}

func (d *TDateTimePicker) Invalidate() {
    defer exceptionProc()
    DateTimePicker_Invalidate(d.instance)
}

func (d *TDateTimePicker) Realign() {
    defer exceptionProc()
    DateTimePicker_Realign(d.instance)
}

func (d *TDateTimePicker) Repaint() {
    defer exceptionProc()
    DateTimePicker_Repaint(d.instance)
}

func (d *TDateTimePicker) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    DateTimePicker_ScaleBy(d.instance, M , D )
}

func (d *TDateTimePicker) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    DateTimePicker_SetBounds(d.instance, ALeft , ATop , AWidth , AHeight )
}

func (d *TDateTimePicker) SetFocus() {
    defer exceptionProc()
    DateTimePicker_SetFocus(d.instance)
}

func (d *TDateTimePicker) Update() {
    defer exceptionProc()
    DateTimePicker_Update(d.instance)
}

func (d *TDateTimePicker) BringToFront() {
    defer exceptionProc()
    DateTimePicker_BringToFront(d.instance)
}

func (d *TDateTimePicker) HasParent() bool {
    defer exceptionProc()
    return DateTimePicker_HasParent(d.instance)
}

func (d *TDateTimePicker) Hide() {
    defer exceptionProc()
    DateTimePicker_Hide(d.instance)
}

func (d *TDateTimePicker) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return DateTimePicker_Perform(d.instance, Msg , WParam , LParam )
}

func (d *TDateTimePicker) Refresh() {
    defer exceptionProc()
    DateTimePicker_Refresh(d.instance)
}

func (d *TDateTimePicker) SendToBack() {
    defer exceptionProc()
    DateTimePicker_SendToBack(d.instance)
}

func (d *TDateTimePicker) Show() {
    defer exceptionProc()
    DateTimePicker_Show(d.instance)
}

func (d *TDateTimePicker) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return DateTimePicker_GetTextBuf(d.instance, Buffer , BufSize )
}

func (d *TDateTimePicker) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(DateTimePicker_FindComponent(d.instance, AName ))
}

func (d *TDateTimePicker) GetNamePath() string {
    defer exceptionProc()
    return DateTimePicker_GetNamePath(d.instance)
}

func (d *TDateTimePicker) Assign(Source IObject) {
    defer exceptionProc()
    DateTimePicker_Assign(d.instance, CheckPtr(Source))
}

func (d *TDateTimePicker) ClassName() string {
    defer exceptionProc()
    return DateTimePicker_ClassName(d.instance)
}

func (d *TDateTimePicker) Equals(Obj IObject) bool {
    defer exceptionProc()
    return DateTimePicker_Equals(d.instance, CheckPtr(Obj))
}

func (d *TDateTimePicker) GetHashCode() int32 {
    defer exceptionProc()
    return DateTimePicker_GetHashCode(d.instance)
}

func (d *TDateTimePicker) ToString() string {
    defer exceptionProc()
    return DateTimePicker_ToString(d.instance)
}

func (d *TDateTimePicker) Align() TAlign {
    defer exceptionProc()
    return DateTimePicker_GetAlign(d.instance)
}

func (d *TDateTimePicker) SetAlign(value TAlign) {
    defer exceptionProc()
    DateTimePicker_SetAlign(d.instance, value)
}

func (d *TDateTimePicker) Anchors() TAnchors {
    defer exceptionProc()
    return DateTimePicker_GetAnchors(d.instance)
}

func (d *TDateTimePicker) SetAnchors(value TAnchors) {
    defer exceptionProc()
    DateTimePicker_SetAnchors(d.instance, value)
}

func (d *TDateTimePicker) Checked() bool {
    defer exceptionProc()
    return DateTimePicker_GetChecked(d.instance)
}

func (d *TDateTimePicker) SetChecked(value bool) {
    defer exceptionProc()
    DateTimePicker_SetChecked(d.instance, value)
}

func (d *TDateTimePicker) Color() TColor {
    defer exceptionProc()
    return DateTimePicker_GetColor(d.instance)
}

func (d *TDateTimePicker) SetColor(value TColor) {
    defer exceptionProc()
    DateTimePicker_SetColor(d.instance, value)
}

func (d *TDateTimePicker) DoubleBuffered() bool {
    defer exceptionProc()
    return DateTimePicker_GetDoubleBuffered(d.instance)
}

func (d *TDateTimePicker) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    DateTimePicker_SetDoubleBuffered(d.instance, value)
}

func (d *TDateTimePicker) Enabled() bool {
    defer exceptionProc()
    return DateTimePicker_GetEnabled(d.instance)
}

func (d *TDateTimePicker) SetEnabled(value bool) {
    defer exceptionProc()
    DateTimePicker_SetEnabled(d.instance, value)
}

func (d *TDateTimePicker) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(DateTimePicker_GetFont(d.instance))
}

func (d *TDateTimePicker) SetFont(value *TFont) {
    defer exceptionProc()
    DateTimePicker_SetFont(d.instance, CheckPtr(value))
}

func (d *TDateTimePicker) ParentColor() bool {
    defer exceptionProc()
    return DateTimePicker_GetParentColor(d.instance)
}

func (d *TDateTimePicker) SetParentColor(value bool) {
    defer exceptionProc()
    DateTimePicker_SetParentColor(d.instance, value)
}

func (d *TDateTimePicker) ParentFont() bool {
    defer exceptionProc()
    return DateTimePicker_GetParentFont(d.instance)
}

func (d *TDateTimePicker) SetParentFont(value bool) {
    defer exceptionProc()
    DateTimePicker_SetParentFont(d.instance, value)
}

func (d *TDateTimePicker) ParentShowHint() bool {
    defer exceptionProc()
    return DateTimePicker_GetParentShowHint(d.instance)
}

func (d *TDateTimePicker) SetParentShowHint(value bool) {
    defer exceptionProc()
    DateTimePicker_SetParentShowHint(d.instance, value)
}

func (d *TDateTimePicker) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(DateTimePicker_GetPopupMenu(d.instance))
}

func (d *TDateTimePicker) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    DateTimePicker_SetPopupMenu(d.instance, CheckPtr(value))
}

func (d *TDateTimePicker) ShowHint() bool {
    defer exceptionProc()
    return DateTimePicker_GetShowHint(d.instance)
}

func (d *TDateTimePicker) SetShowHint(value bool) {
    defer exceptionProc()
    DateTimePicker_SetShowHint(d.instance, value)
}

func (d *TDateTimePicker) TabOrder() int16 {
    defer exceptionProc()
    return DateTimePicker_GetTabOrder(d.instance)
}

func (d *TDateTimePicker) SetTabOrder(value int16) {
    defer exceptionProc()
    DateTimePicker_SetTabOrder(d.instance, value)
}

func (d *TDateTimePicker) TabStop() bool {
    defer exceptionProc()
    return DateTimePicker_GetTabStop(d.instance)
}

func (d *TDateTimePicker) SetTabStop(value bool) {
    defer exceptionProc()
    DateTimePicker_SetTabStop(d.instance, value)
}

func (d *TDateTimePicker) Visible() bool {
    defer exceptionProc()
    return DateTimePicker_GetVisible(d.instance)
}

func (d *TDateTimePicker) SetVisible(value bool) {
    defer exceptionProc()
    DateTimePicker_SetVisible(d.instance, value)
}

func (d *TDateTimePicker) SetOnClick(fn TNotifyEvent) {
    DateTimePicker_SetOnClick(d.instance, fn)
}

func (d *TDateTimePicker) SetOnChange(fn TNotifyEvent) {
    DateTimePicker_SetOnChange(d.instance, fn)
}

func (d *TDateTimePicker) SetOnEnter(fn TNotifyEvent) {
    DateTimePicker_SetOnEnter(d.instance, fn)
}

func (d *TDateTimePicker) SetOnExit(fn TNotifyEvent) {
    DateTimePicker_SetOnExit(d.instance, fn)
}

func (d *TDateTimePicker) SetOnKeyDown(fn TKeyEvent) {
    DateTimePicker_SetOnKeyDown(d.instance, fn)
}

func (d *TDateTimePicker) SetOnKeyPress(fn TKeyPressEvent) {
    DateTimePicker_SetOnKeyPress(d.instance, fn)
}

func (d *TDateTimePicker) SetOnKeyUp(fn TKeyEvent) {
    DateTimePicker_SetOnKeyUp(d.instance, fn)
}

func (d *TDateTimePicker) SetOnMouseEnter(fn TNotifyEvent) {
    DateTimePicker_SetOnMouseEnter(d.instance, fn)
}

func (d *TDateTimePicker) SetOnMouseLeave(fn TNotifyEvent) {
    DateTimePicker_SetOnMouseLeave(d.instance, fn)
}

func (d *TDateTimePicker) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(DateTimePicker_GetBrush(d.instance))
}

func (d *TDateTimePicker) ControlCount() int32 {
    defer exceptionProc()
    return DateTimePicker_GetControlCount(d.instance)
}

func (d *TDateTimePicker) Handle() HWND {
    defer exceptionProc()
    return DateTimePicker_GetHandle(d.instance)
}

func (d *TDateTimePicker) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(DateTimePicker_GetAction(d.instance))
}

func (d *TDateTimePicker) SetAction(value IComponent) {
    defer exceptionProc()
    DateTimePicker_SetAction(d.instance, CheckPtr(value))
}

func (d *TDateTimePicker) BoundsRect() TRect {
    defer exceptionProc()
    return DateTimePicker_GetBoundsRect(d.instance)
}

func (d *TDateTimePicker) SetBoundsRect(value TRect) {
    defer exceptionProc()
    DateTimePicker_SetBoundsRect(d.instance, value)
}

func (d *TDateTimePicker) ClientHeight() int32 {
    defer exceptionProc()
    return DateTimePicker_GetClientHeight(d.instance)
}

func (d *TDateTimePicker) SetClientHeight(value int32) {
    defer exceptionProc()
    DateTimePicker_SetClientHeight(d.instance, value)
}

func (d *TDateTimePicker) ClientRect() TRect {
    defer exceptionProc()
    return DateTimePicker_GetClientRect(d.instance)
}

func (d *TDateTimePicker) ClientWidth() int32 {
    defer exceptionProc()
    return DateTimePicker_GetClientWidth(d.instance)
}

func (d *TDateTimePicker) SetClientWidth(value int32) {
    defer exceptionProc()
    DateTimePicker_SetClientWidth(d.instance, value)
}

func (d *TDateTimePicker) ExplicitLeft() int32 {
    defer exceptionProc()
    return DateTimePicker_GetExplicitLeft(d.instance)
}

func (d *TDateTimePicker) ExplicitTop() int32 {
    defer exceptionProc()
    return DateTimePicker_GetExplicitTop(d.instance)
}

func (d *TDateTimePicker) ExplicitWidth() int32 {
    defer exceptionProc()
    return DateTimePicker_GetExplicitWidth(d.instance)
}

func (d *TDateTimePicker) ExplicitHeight() int32 {
    defer exceptionProc()
    return DateTimePicker_GetExplicitHeight(d.instance)
}

func (d *TDateTimePicker) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(DateTimePicker_GetParent(d.instance))
}

func (d *TDateTimePicker) SetParent(value IControl) {
    defer exceptionProc()
    DateTimePicker_SetParent(d.instance, CheckPtr(value))
}

func (d *TDateTimePicker) AlignWithMargins() bool {
    defer exceptionProc()
    return DateTimePicker_GetAlignWithMargins(d.instance)
}

func (d *TDateTimePicker) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    DateTimePicker_SetAlignWithMargins(d.instance, value)
}

func (d *TDateTimePicker) Left() int32 {
    defer exceptionProc()
    return DateTimePicker_GetLeft(d.instance)
}

func (d *TDateTimePicker) SetLeft(value int32) {
    defer exceptionProc()
    DateTimePicker_SetLeft(d.instance, value)
}

func (d *TDateTimePicker) Top() int32 {
    defer exceptionProc()
    return DateTimePicker_GetTop(d.instance)
}

func (d *TDateTimePicker) SetTop(value int32) {
    defer exceptionProc()
    DateTimePicker_SetTop(d.instance, value)
}

func (d *TDateTimePicker) Width() int32 {
    defer exceptionProc()
    return DateTimePicker_GetWidth(d.instance)
}

func (d *TDateTimePicker) SetWidth(value int32) {
    defer exceptionProc()
    DateTimePicker_SetWidth(d.instance, value)
}

func (d *TDateTimePicker) Height() int32 {
    defer exceptionProc()
    return DateTimePicker_GetHeight(d.instance)
}

func (d *TDateTimePicker) SetHeight(value int32) {
    defer exceptionProc()
    DateTimePicker_SetHeight(d.instance, value)
}

func (d *TDateTimePicker) Cursor() TCursor {
    defer exceptionProc()
    return DateTimePicker_GetCursor(d.instance)
}

func (d *TDateTimePicker) SetCursor(value TCursor) {
    defer exceptionProc()
    DateTimePicker_SetCursor(d.instance, value)
}

func (d *TDateTimePicker) Hint() string {
    defer exceptionProc()
    return DateTimePicker_GetHint(d.instance)
}

func (d *TDateTimePicker) SetHint(value string) {
    defer exceptionProc()
    DateTimePicker_SetHint(d.instance, value)
}

func (d *TDateTimePicker) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(DateTimePicker_GetMargins(d.instance))
}

func (d *TDateTimePicker) SetMargins(value *TMargins) {
    defer exceptionProc()
    DateTimePicker_SetMargins(d.instance, CheckPtr(value))
}

func (d *TDateTimePicker) ComponentCount() int32 {
    defer exceptionProc()
    return DateTimePicker_GetComponentCount(d.instance)
}

func (d *TDateTimePicker) ComponentIndex() int32 {
    defer exceptionProc()
    return DateTimePicker_GetComponentIndex(d.instance)
}

func (d *TDateTimePicker) SetComponentIndex(value int32) {
    defer exceptionProc()
    DateTimePicker_SetComponentIndex(d.instance, value)
}

func (d *TDateTimePicker) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(DateTimePicker_GetOwner(d.instance))
}

func (d *TDateTimePicker) Name() string {
    defer exceptionProc()
    return DateTimePicker_GetName(d.instance)
}

func (d *TDateTimePicker) SetName(value string) {
    defer exceptionProc()
    DateTimePicker_SetName(d.instance, value)
}

func (d *TDateTimePicker) Tag() int {
    defer exceptionProc()
    return DateTimePicker_GetTag(d.instance)
}

func (d *TDateTimePicker) SetTag(value int) {
    defer exceptionProc()
    DateTimePicker_SetTag(d.instance, value)
}

func (d *TDateTimePicker) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(DateTimePicker_GetControls(d.instance, Index))
}

func (d *TDateTimePicker) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(DateTimePicker_GetComponents(d.instance, AIndex))
}

