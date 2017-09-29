
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TRadioButton struct {
    IControl
    instance uintptr
}

func NewRadioButton(owner IComponent) *TRadioButton {
    r := new(TRadioButton)
    r.instance = RadioButton_Create(owner.Instance())
    return r
}

func RadioButtonFromInst(inst uintptr) *TRadioButton {
    r := new(TRadioButton)
    r.instance = inst
    return r
}

func RadioButtonFromObj(obj IObject) *TRadioButton {
    r := new(TRadioButton)
    r.instance = CheckPtr(obj)
    return r
}

func (r *TRadioButton) Free() {
    if r.instance != 0 {
        RadioButton_Free(r.instance)
    }
}

func (r *TRadioButton) Instance() uintptr {
    return r.instance
}

func (r *TRadioButton) IsValid() bool {
    return r.instance != 0
}

func (r *TRadioButton) CanFocus() bool {
    defer exceptionProc()
    return RadioButton_CanFocus(r.instance)
}

func (r *TRadioButton) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    RadioButton_FlipChildren(r.instance, AllLevels )
}

func (r *TRadioButton) Focused() bool {
    defer exceptionProc()
    return RadioButton_Focused(r.instance)
}

func (r *TRadioButton) HandleAllocated() bool {
    defer exceptionProc()
    return RadioButton_HandleAllocated(r.instance)
}

func (r *TRadioButton) Invalidate() {
    defer exceptionProc()
    RadioButton_Invalidate(r.instance)
}

func (r *TRadioButton) Realign() {
    defer exceptionProc()
    RadioButton_Realign(r.instance)
}

func (r *TRadioButton) Repaint() {
    defer exceptionProc()
    RadioButton_Repaint(r.instance)
}

func (r *TRadioButton) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    RadioButton_ScaleBy(r.instance, M , D )
}

func (r *TRadioButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    RadioButton_SetBounds(r.instance, ALeft , ATop , AWidth , AHeight )
}

func (r *TRadioButton) SetFocus() {
    defer exceptionProc()
    RadioButton_SetFocus(r.instance)
}

func (r *TRadioButton) Update() {
    defer exceptionProc()
    RadioButton_Update(r.instance)
}

func (r *TRadioButton) BringToFront() {
    defer exceptionProc()
    RadioButton_BringToFront(r.instance)
}

func (r *TRadioButton) HasParent() bool {
    defer exceptionProc()
    return RadioButton_HasParent(r.instance)
}

func (r *TRadioButton) Hide() {
    defer exceptionProc()
    RadioButton_Hide(r.instance)
}

func (r *TRadioButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return RadioButton_Perform(r.instance, Msg , WParam , LParam )
}

func (r *TRadioButton) Refresh() {
    defer exceptionProc()
    RadioButton_Refresh(r.instance)
}

func (r *TRadioButton) SendToBack() {
    defer exceptionProc()
    RadioButton_SendToBack(r.instance)
}

func (r *TRadioButton) Show() {
    defer exceptionProc()
    RadioButton_Show(r.instance)
}

func (r *TRadioButton) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return RadioButton_GetTextBuf(r.instance, Buffer , BufSize )
}

func (r *TRadioButton) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(RadioButton_FindComponent(r.instance, AName ))
}

func (r *TRadioButton) GetNamePath() string {
    defer exceptionProc()
    return RadioButton_GetNamePath(r.instance)
}

func (r *TRadioButton) Assign(Source IObject) {
    defer exceptionProc()
    RadioButton_Assign(r.instance, CheckPtr(Source))
}

func (r *TRadioButton) ClassName() string {
    defer exceptionProc()
    return RadioButton_ClassName(r.instance)
}

func (r *TRadioButton) Equals(Obj IObject) bool {
    defer exceptionProc()
    return RadioButton_Equals(r.instance, CheckPtr(Obj))
}

func (r *TRadioButton) GetHashCode() int32 {
    defer exceptionProc()
    return RadioButton_GetHashCode(r.instance)
}

func (r *TRadioButton) ToString() string {
    defer exceptionProc()
    return RadioButton_ToString(r.instance)
}

func (r *TRadioButton) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(RadioButton_GetAction(r.instance))
}

func (r *TRadioButton) SetAction(value IComponent) {
    defer exceptionProc()
    RadioButton_SetAction(r.instance, CheckPtr(value))
}

func (r *TRadioButton) Align() TAlign {
    defer exceptionProc()
    return RadioButton_GetAlign(r.instance)
}

func (r *TRadioButton) SetAlign(value TAlign) {
    defer exceptionProc()
    RadioButton_SetAlign(r.instance, value)
}

func (r *TRadioButton) Alignment() TLeftRight {
    defer exceptionProc()
    return RadioButton_GetAlignment(r.instance)
}

func (r *TRadioButton) SetAlignment(value TLeftRight) {
    defer exceptionProc()
    RadioButton_SetAlignment(r.instance, value)
}

func (r *TRadioButton) Anchors() TAnchors {
    defer exceptionProc()
    return RadioButton_GetAnchors(r.instance)
}

func (r *TRadioButton) SetAnchors(value TAnchors) {
    defer exceptionProc()
    RadioButton_SetAnchors(r.instance, value)
}

func (r *TRadioButton) Caption() string {
    defer exceptionProc()
    return RadioButton_GetCaption(r.instance)
}

func (r *TRadioButton) SetCaption(value string) {
    defer exceptionProc()
    RadioButton_SetCaption(r.instance, value)
}

func (r *TRadioButton) Checked() bool {
    defer exceptionProc()
    return RadioButton_GetChecked(r.instance)
}

func (r *TRadioButton) SetChecked(value bool) {
    defer exceptionProc()
    RadioButton_SetChecked(r.instance, value)
}

func (r *TRadioButton) Color() TColor {
    defer exceptionProc()
    return RadioButton_GetColor(r.instance)
}

func (r *TRadioButton) SetColor(value TColor) {
    defer exceptionProc()
    RadioButton_SetColor(r.instance, value)
}

func (r *TRadioButton) DoubleBuffered() bool {
    defer exceptionProc()
    return RadioButton_GetDoubleBuffered(r.instance)
}

func (r *TRadioButton) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    RadioButton_SetDoubleBuffered(r.instance, value)
}

func (r *TRadioButton) Enabled() bool {
    defer exceptionProc()
    return RadioButton_GetEnabled(r.instance)
}

func (r *TRadioButton) SetEnabled(value bool) {
    defer exceptionProc()
    RadioButton_SetEnabled(r.instance, value)
}

func (r *TRadioButton) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(RadioButton_GetFont(r.instance))
}

func (r *TRadioButton) SetFont(value *TFont) {
    defer exceptionProc()
    RadioButton_SetFont(r.instance, CheckPtr(value))
}

func (r *TRadioButton) ParentColor() bool {
    defer exceptionProc()
    return RadioButton_GetParentColor(r.instance)
}

func (r *TRadioButton) SetParentColor(value bool) {
    defer exceptionProc()
    RadioButton_SetParentColor(r.instance, value)
}

func (r *TRadioButton) ParentCtl3D() bool {
    defer exceptionProc()
    return RadioButton_GetParentCtl3D(r.instance)
}

func (r *TRadioButton) SetParentCtl3D(value bool) {
    defer exceptionProc()
    RadioButton_SetParentCtl3D(r.instance, value)
}

func (r *TRadioButton) ParentFont() bool {
    defer exceptionProc()
    return RadioButton_GetParentFont(r.instance)
}

func (r *TRadioButton) SetParentFont(value bool) {
    defer exceptionProc()
    RadioButton_SetParentFont(r.instance, value)
}

func (r *TRadioButton) ParentShowHint() bool {
    defer exceptionProc()
    return RadioButton_GetParentShowHint(r.instance)
}

func (r *TRadioButton) SetParentShowHint(value bool) {
    defer exceptionProc()
    RadioButton_SetParentShowHint(r.instance, value)
}

func (r *TRadioButton) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(RadioButton_GetPopupMenu(r.instance))
}

func (r *TRadioButton) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    RadioButton_SetPopupMenu(r.instance, CheckPtr(value))
}

func (r *TRadioButton) ShowHint() bool {
    defer exceptionProc()
    return RadioButton_GetShowHint(r.instance)
}

func (r *TRadioButton) SetShowHint(value bool) {
    defer exceptionProc()
    RadioButton_SetShowHint(r.instance, value)
}

func (r *TRadioButton) TabOrder() int16 {
    defer exceptionProc()
    return RadioButton_GetTabOrder(r.instance)
}

func (r *TRadioButton) SetTabOrder(value int16) {
    defer exceptionProc()
    RadioButton_SetTabOrder(r.instance, value)
}

func (r *TRadioButton) TabStop() bool {
    defer exceptionProc()
    return RadioButton_GetTabStop(r.instance)
}

func (r *TRadioButton) SetTabStop(value bool) {
    defer exceptionProc()
    RadioButton_SetTabStop(r.instance, value)
}

func (r *TRadioButton) Visible() bool {
    defer exceptionProc()
    return RadioButton_GetVisible(r.instance)
}

func (r *TRadioButton) SetVisible(value bool) {
    defer exceptionProc()
    RadioButton_SetVisible(r.instance, value)
}

func (r *TRadioButton) WordWrap() bool {
    defer exceptionProc()
    return RadioButton_GetWordWrap(r.instance)
}

func (r *TRadioButton) SetWordWrap(value bool) {
    defer exceptionProc()
    RadioButton_SetWordWrap(r.instance, value)
}

func (r *TRadioButton) SetOnClick(fn TNotifyEvent) {
    RadioButton_SetOnClick(r.instance, fn)
}

func (r *TRadioButton) SetOnDblClick(fn TNotifyEvent) {
    RadioButton_SetOnDblClick(r.instance, fn)
}

func (r *TRadioButton) SetOnEnter(fn TNotifyEvent) {
    RadioButton_SetOnEnter(r.instance, fn)
}

func (r *TRadioButton) SetOnExit(fn TNotifyEvent) {
    RadioButton_SetOnExit(r.instance, fn)
}

func (r *TRadioButton) SetOnKeyDown(fn TKeyEvent) {
    RadioButton_SetOnKeyDown(r.instance, fn)
}

func (r *TRadioButton) SetOnKeyPress(fn TKeyPressEvent) {
    RadioButton_SetOnKeyPress(r.instance, fn)
}

func (r *TRadioButton) SetOnKeyUp(fn TKeyEvent) {
    RadioButton_SetOnKeyUp(r.instance, fn)
}

func (r *TRadioButton) SetOnMouseDown(fn TMouseEvent) {
    RadioButton_SetOnMouseDown(r.instance, fn)
}

func (r *TRadioButton) SetOnMouseEnter(fn TNotifyEvent) {
    RadioButton_SetOnMouseEnter(r.instance, fn)
}

func (r *TRadioButton) SetOnMouseLeave(fn TNotifyEvent) {
    RadioButton_SetOnMouseLeave(r.instance, fn)
}

func (r *TRadioButton) SetOnMouseMove(fn TMouseMoveEvent) {
    RadioButton_SetOnMouseMove(r.instance, fn)
}

func (r *TRadioButton) SetOnMouseUp(fn TMouseEvent) {
    RadioButton_SetOnMouseUp(r.instance, fn)
}

func (r *TRadioButton) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(RadioButton_GetBrush(r.instance))
}

func (r *TRadioButton) ControlCount() int32 {
    defer exceptionProc()
    return RadioButton_GetControlCount(r.instance)
}

func (r *TRadioButton) Handle() HWND {
    defer exceptionProc()
    return RadioButton_GetHandle(r.instance)
}

func (r *TRadioButton) BoundsRect() TRect {
    defer exceptionProc()
    return RadioButton_GetBoundsRect(r.instance)
}

func (r *TRadioButton) SetBoundsRect(value TRect) {
    defer exceptionProc()
    RadioButton_SetBoundsRect(r.instance, value)
}

func (r *TRadioButton) ClientHeight() int32 {
    defer exceptionProc()
    return RadioButton_GetClientHeight(r.instance)
}

func (r *TRadioButton) SetClientHeight(value int32) {
    defer exceptionProc()
    RadioButton_SetClientHeight(r.instance, value)
}

func (r *TRadioButton) ClientRect() TRect {
    defer exceptionProc()
    return RadioButton_GetClientRect(r.instance)
}

func (r *TRadioButton) ClientWidth() int32 {
    defer exceptionProc()
    return RadioButton_GetClientWidth(r.instance)
}

func (r *TRadioButton) SetClientWidth(value int32) {
    defer exceptionProc()
    RadioButton_SetClientWidth(r.instance, value)
}

func (r *TRadioButton) ExplicitLeft() int32 {
    defer exceptionProc()
    return RadioButton_GetExplicitLeft(r.instance)
}

func (r *TRadioButton) ExplicitTop() int32 {
    defer exceptionProc()
    return RadioButton_GetExplicitTop(r.instance)
}

func (r *TRadioButton) ExplicitWidth() int32 {
    defer exceptionProc()
    return RadioButton_GetExplicitWidth(r.instance)
}

func (r *TRadioButton) ExplicitHeight() int32 {
    defer exceptionProc()
    return RadioButton_GetExplicitHeight(r.instance)
}

func (r *TRadioButton) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(RadioButton_GetParent(r.instance))
}

func (r *TRadioButton) SetParent(value IControl) {
    defer exceptionProc()
    RadioButton_SetParent(r.instance, CheckPtr(value))
}

func (r *TRadioButton) AlignWithMargins() bool {
    defer exceptionProc()
    return RadioButton_GetAlignWithMargins(r.instance)
}

func (r *TRadioButton) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    RadioButton_SetAlignWithMargins(r.instance, value)
}

func (r *TRadioButton) Left() int32 {
    defer exceptionProc()
    return RadioButton_GetLeft(r.instance)
}

func (r *TRadioButton) SetLeft(value int32) {
    defer exceptionProc()
    RadioButton_SetLeft(r.instance, value)
}

func (r *TRadioButton) Top() int32 {
    defer exceptionProc()
    return RadioButton_GetTop(r.instance)
}

func (r *TRadioButton) SetTop(value int32) {
    defer exceptionProc()
    RadioButton_SetTop(r.instance, value)
}

func (r *TRadioButton) Width() int32 {
    defer exceptionProc()
    return RadioButton_GetWidth(r.instance)
}

func (r *TRadioButton) SetWidth(value int32) {
    defer exceptionProc()
    RadioButton_SetWidth(r.instance, value)
}

func (r *TRadioButton) Height() int32 {
    defer exceptionProc()
    return RadioButton_GetHeight(r.instance)
}

func (r *TRadioButton) SetHeight(value int32) {
    defer exceptionProc()
    RadioButton_SetHeight(r.instance, value)
}

func (r *TRadioButton) Cursor() TCursor {
    defer exceptionProc()
    return RadioButton_GetCursor(r.instance)
}

func (r *TRadioButton) SetCursor(value TCursor) {
    defer exceptionProc()
    RadioButton_SetCursor(r.instance, value)
}

func (r *TRadioButton) Hint() string {
    defer exceptionProc()
    return RadioButton_GetHint(r.instance)
}

func (r *TRadioButton) SetHint(value string) {
    defer exceptionProc()
    RadioButton_SetHint(r.instance, value)
}

func (r *TRadioButton) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(RadioButton_GetMargins(r.instance))
}

func (r *TRadioButton) SetMargins(value *TMargins) {
    defer exceptionProc()
    RadioButton_SetMargins(r.instance, CheckPtr(value))
}

func (r *TRadioButton) ComponentCount() int32 {
    defer exceptionProc()
    return RadioButton_GetComponentCount(r.instance)
}

func (r *TRadioButton) ComponentIndex() int32 {
    defer exceptionProc()
    return RadioButton_GetComponentIndex(r.instance)
}

func (r *TRadioButton) SetComponentIndex(value int32) {
    defer exceptionProc()
    RadioButton_SetComponentIndex(r.instance, value)
}

func (r *TRadioButton) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(RadioButton_GetOwner(r.instance))
}

func (r *TRadioButton) Name() string {
    defer exceptionProc()
    return RadioButton_GetName(r.instance)
}

func (r *TRadioButton) SetName(value string) {
    defer exceptionProc()
    RadioButton_SetName(r.instance, value)
}

func (r *TRadioButton) Tag() int {
    defer exceptionProc()
    return RadioButton_GetTag(r.instance)
}

func (r *TRadioButton) SetTag(value int) {
    defer exceptionProc()
    RadioButton_SetTag(r.instance, value)
}

func (r *TRadioButton) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(RadioButton_GetControls(r.instance, Index))
}

func (r *TRadioButton) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(RadioButton_GetComponents(r.instance, AIndex))
}

