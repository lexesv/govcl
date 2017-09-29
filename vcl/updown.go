
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TUpDown struct {
    IControl
    instance uintptr
}

func NewUpDown(owner IComponent) *TUpDown {
    u := new(TUpDown)
    u.instance = UpDown_Create(owner.Instance())
    return u
}

func UpDownFromInst(inst uintptr) *TUpDown {
    u := new(TUpDown)
    u.instance = inst
    return u
}

func UpDownFromObj(obj IObject) *TUpDown {
    u := new(TUpDown)
    u.instance = CheckPtr(obj)
    return u
}

func (u *TUpDown) Free() {
    if u.instance != 0 {
        UpDown_Free(u.instance)
    }
}

func (u *TUpDown) Instance() uintptr {
    return u.instance
}

func (u *TUpDown) IsValid() bool {
    return u.instance != 0
}

func (u *TUpDown) CanFocus() bool {
    defer exceptionProc()
    return UpDown_CanFocus(u.instance)
}

func (u *TUpDown) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    UpDown_FlipChildren(u.instance, AllLevels )
}

func (u *TUpDown) Focused() bool {
    defer exceptionProc()
    return UpDown_Focused(u.instance)
}

func (u *TUpDown) HandleAllocated() bool {
    defer exceptionProc()
    return UpDown_HandleAllocated(u.instance)
}

func (u *TUpDown) Invalidate() {
    defer exceptionProc()
    UpDown_Invalidate(u.instance)
}

func (u *TUpDown) Realign() {
    defer exceptionProc()
    UpDown_Realign(u.instance)
}

func (u *TUpDown) Repaint() {
    defer exceptionProc()
    UpDown_Repaint(u.instance)
}

func (u *TUpDown) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    UpDown_ScaleBy(u.instance, M , D )
}

func (u *TUpDown) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    UpDown_SetBounds(u.instance, ALeft , ATop , AWidth , AHeight )
}

func (u *TUpDown) SetFocus() {
    defer exceptionProc()
    UpDown_SetFocus(u.instance)
}

func (u *TUpDown) Update() {
    defer exceptionProc()
    UpDown_Update(u.instance)
}

func (u *TUpDown) BringToFront() {
    defer exceptionProc()
    UpDown_BringToFront(u.instance)
}

func (u *TUpDown) HasParent() bool {
    defer exceptionProc()
    return UpDown_HasParent(u.instance)
}

func (u *TUpDown) Hide() {
    defer exceptionProc()
    UpDown_Hide(u.instance)
}

func (u *TUpDown) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return UpDown_Perform(u.instance, Msg , WParam , LParam )
}

func (u *TUpDown) Refresh() {
    defer exceptionProc()
    UpDown_Refresh(u.instance)
}

func (u *TUpDown) SendToBack() {
    defer exceptionProc()
    UpDown_SendToBack(u.instance)
}

func (u *TUpDown) Show() {
    defer exceptionProc()
    UpDown_Show(u.instance)
}

func (u *TUpDown) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return UpDown_GetTextBuf(u.instance, Buffer , BufSize )
}

func (u *TUpDown) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(UpDown_FindComponent(u.instance, AName ))
}

func (u *TUpDown) GetNamePath() string {
    defer exceptionProc()
    return UpDown_GetNamePath(u.instance)
}

func (u *TUpDown) Assign(Source IObject) {
    defer exceptionProc()
    UpDown_Assign(u.instance, CheckPtr(Source))
}

func (u *TUpDown) ClassName() string {
    defer exceptionProc()
    return UpDown_ClassName(u.instance)
}

func (u *TUpDown) Equals(Obj IObject) bool {
    defer exceptionProc()
    return UpDown_Equals(u.instance, CheckPtr(Obj))
}

func (u *TUpDown) GetHashCode() int32 {
    defer exceptionProc()
    return UpDown_GetHashCode(u.instance)
}

func (u *TUpDown) ToString() string {
    defer exceptionProc()
    return UpDown_ToString(u.instance)
}

func (u *TUpDown) Anchors() TAnchors {
    defer exceptionProc()
    return UpDown_GetAnchors(u.instance)
}

func (u *TUpDown) SetAnchors(value TAnchors) {
    defer exceptionProc()
    UpDown_SetAnchors(u.instance, value)
}

func (u *TUpDown) DoubleBuffered() bool {
    defer exceptionProc()
    return UpDown_GetDoubleBuffered(u.instance)
}

func (u *TUpDown) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    UpDown_SetDoubleBuffered(u.instance, value)
}

func (u *TUpDown) Enabled() bool {
    defer exceptionProc()
    return UpDown_GetEnabled(u.instance)
}

func (u *TUpDown) SetEnabled(value bool) {
    defer exceptionProc()
    UpDown_SetEnabled(u.instance, value)
}

func (u *TUpDown) Hint() string {
    defer exceptionProc()
    return UpDown_GetHint(u.instance)
}

func (u *TUpDown) SetHint(value string) {
    defer exceptionProc()
    UpDown_SetHint(u.instance, value)
}

func (u *TUpDown) Orientation() TUDOrientation {
    defer exceptionProc()
    return UpDown_GetOrientation(u.instance)
}

func (u *TUpDown) SetOrientation(value TUDOrientation) {
    defer exceptionProc()
    UpDown_SetOrientation(u.instance, value)
}

func (u *TUpDown) ParentShowHint() bool {
    defer exceptionProc()
    return UpDown_GetParentShowHint(u.instance)
}

func (u *TUpDown) SetParentShowHint(value bool) {
    defer exceptionProc()
    UpDown_SetParentShowHint(u.instance, value)
}

func (u *TUpDown) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(UpDown_GetPopupMenu(u.instance))
}

func (u *TUpDown) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    UpDown_SetPopupMenu(u.instance, CheckPtr(value))
}

func (u *TUpDown) Position() int32 {
    defer exceptionProc()
    return UpDown_GetPosition(u.instance)
}

func (u *TUpDown) SetPosition(value int32) {
    defer exceptionProc()
    UpDown_SetPosition(u.instance, value)
}

func (u *TUpDown) ShowHint() bool {
    defer exceptionProc()
    return UpDown_GetShowHint(u.instance)
}

func (u *TUpDown) SetShowHint(value bool) {
    defer exceptionProc()
    UpDown_SetShowHint(u.instance, value)
}

func (u *TUpDown) TabOrder() int16 {
    defer exceptionProc()
    return UpDown_GetTabOrder(u.instance)
}

func (u *TUpDown) SetTabOrder(value int16) {
    defer exceptionProc()
    UpDown_SetTabOrder(u.instance, value)
}

func (u *TUpDown) TabStop() bool {
    defer exceptionProc()
    return UpDown_GetTabStop(u.instance)
}

func (u *TUpDown) SetTabStop(value bool) {
    defer exceptionProc()
    UpDown_SetTabStop(u.instance, value)
}

func (u *TUpDown) Visible() bool {
    defer exceptionProc()
    return UpDown_GetVisible(u.instance)
}

func (u *TUpDown) SetVisible(value bool) {
    defer exceptionProc()
    UpDown_SetVisible(u.instance, value)
}

func (u *TUpDown) Wrap() bool {
    defer exceptionProc()
    return UpDown_GetWrap(u.instance)
}

func (u *TUpDown) SetWrap(value bool) {
    defer exceptionProc()
    UpDown_SetWrap(u.instance, value)
}

func (u *TUpDown) SetOnClick(fn TUDClickEvent) {
    UpDown_SetOnClick(u.instance, fn)
}

func (u *TUpDown) SetOnEnter(fn TNotifyEvent) {
    UpDown_SetOnEnter(u.instance, fn)
}

func (u *TUpDown) SetOnExit(fn TNotifyEvent) {
    UpDown_SetOnExit(u.instance, fn)
}

func (u *TUpDown) SetOnMouseDown(fn TMouseEvent) {
    UpDown_SetOnMouseDown(u.instance, fn)
}

func (u *TUpDown) SetOnMouseEnter(fn TNotifyEvent) {
    UpDown_SetOnMouseEnter(u.instance, fn)
}

func (u *TUpDown) SetOnMouseLeave(fn TNotifyEvent) {
    UpDown_SetOnMouseLeave(u.instance, fn)
}

func (u *TUpDown) SetOnMouseMove(fn TMouseMoveEvent) {
    UpDown_SetOnMouseMove(u.instance, fn)
}

func (u *TUpDown) SetOnMouseUp(fn TMouseEvent) {
    UpDown_SetOnMouseUp(u.instance, fn)
}

func (u *TUpDown) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(UpDown_GetBrush(u.instance))
}

func (u *TUpDown) ControlCount() int32 {
    defer exceptionProc()
    return UpDown_GetControlCount(u.instance)
}

func (u *TUpDown) Handle() HWND {
    defer exceptionProc()
    return UpDown_GetHandle(u.instance)
}

func (u *TUpDown) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(UpDown_GetAction(u.instance))
}

func (u *TUpDown) SetAction(value IComponent) {
    defer exceptionProc()
    UpDown_SetAction(u.instance, CheckPtr(value))
}

func (u *TUpDown) Align() TAlign {
    defer exceptionProc()
    return UpDown_GetAlign(u.instance)
}

func (u *TUpDown) SetAlign(value TAlign) {
    defer exceptionProc()
    UpDown_SetAlign(u.instance, value)
}

func (u *TUpDown) BoundsRect() TRect {
    defer exceptionProc()
    return UpDown_GetBoundsRect(u.instance)
}

func (u *TUpDown) SetBoundsRect(value TRect) {
    defer exceptionProc()
    UpDown_SetBoundsRect(u.instance, value)
}

func (u *TUpDown) ClientHeight() int32 {
    defer exceptionProc()
    return UpDown_GetClientHeight(u.instance)
}

func (u *TUpDown) SetClientHeight(value int32) {
    defer exceptionProc()
    UpDown_SetClientHeight(u.instance, value)
}

func (u *TUpDown) ClientRect() TRect {
    defer exceptionProc()
    return UpDown_GetClientRect(u.instance)
}

func (u *TUpDown) ClientWidth() int32 {
    defer exceptionProc()
    return UpDown_GetClientWidth(u.instance)
}

func (u *TUpDown) SetClientWidth(value int32) {
    defer exceptionProc()
    UpDown_SetClientWidth(u.instance, value)
}

func (u *TUpDown) ExplicitLeft() int32 {
    defer exceptionProc()
    return UpDown_GetExplicitLeft(u.instance)
}

func (u *TUpDown) ExplicitTop() int32 {
    defer exceptionProc()
    return UpDown_GetExplicitTop(u.instance)
}

func (u *TUpDown) ExplicitWidth() int32 {
    defer exceptionProc()
    return UpDown_GetExplicitWidth(u.instance)
}

func (u *TUpDown) ExplicitHeight() int32 {
    defer exceptionProc()
    return UpDown_GetExplicitHeight(u.instance)
}

func (u *TUpDown) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(UpDown_GetParent(u.instance))
}

func (u *TUpDown) SetParent(value IControl) {
    defer exceptionProc()
    UpDown_SetParent(u.instance, CheckPtr(value))
}

func (u *TUpDown) AlignWithMargins() bool {
    defer exceptionProc()
    return UpDown_GetAlignWithMargins(u.instance)
}

func (u *TUpDown) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    UpDown_SetAlignWithMargins(u.instance, value)
}

func (u *TUpDown) Left() int32 {
    defer exceptionProc()
    return UpDown_GetLeft(u.instance)
}

func (u *TUpDown) SetLeft(value int32) {
    defer exceptionProc()
    UpDown_SetLeft(u.instance, value)
}

func (u *TUpDown) Top() int32 {
    defer exceptionProc()
    return UpDown_GetTop(u.instance)
}

func (u *TUpDown) SetTop(value int32) {
    defer exceptionProc()
    UpDown_SetTop(u.instance, value)
}

func (u *TUpDown) Width() int32 {
    defer exceptionProc()
    return UpDown_GetWidth(u.instance)
}

func (u *TUpDown) SetWidth(value int32) {
    defer exceptionProc()
    UpDown_SetWidth(u.instance, value)
}

func (u *TUpDown) Height() int32 {
    defer exceptionProc()
    return UpDown_GetHeight(u.instance)
}

func (u *TUpDown) SetHeight(value int32) {
    defer exceptionProc()
    UpDown_SetHeight(u.instance, value)
}

func (u *TUpDown) Cursor() TCursor {
    defer exceptionProc()
    return UpDown_GetCursor(u.instance)
}

func (u *TUpDown) SetCursor(value TCursor) {
    defer exceptionProc()
    UpDown_SetCursor(u.instance, value)
}

func (u *TUpDown) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(UpDown_GetMargins(u.instance))
}

func (u *TUpDown) SetMargins(value *TMargins) {
    defer exceptionProc()
    UpDown_SetMargins(u.instance, CheckPtr(value))
}

func (u *TUpDown) ComponentCount() int32 {
    defer exceptionProc()
    return UpDown_GetComponentCount(u.instance)
}

func (u *TUpDown) ComponentIndex() int32 {
    defer exceptionProc()
    return UpDown_GetComponentIndex(u.instance)
}

func (u *TUpDown) SetComponentIndex(value int32) {
    defer exceptionProc()
    UpDown_SetComponentIndex(u.instance, value)
}

func (u *TUpDown) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(UpDown_GetOwner(u.instance))
}

func (u *TUpDown) Name() string {
    defer exceptionProc()
    return UpDown_GetName(u.instance)
}

func (u *TUpDown) SetName(value string) {
    defer exceptionProc()
    UpDown_SetName(u.instance, value)
}

func (u *TUpDown) Tag() int {
    defer exceptionProc()
    return UpDown_GetTag(u.instance)
}

func (u *TUpDown) SetTag(value int) {
    defer exceptionProc()
    UpDown_SetTag(u.instance, value)
}

func (u *TUpDown) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(UpDown_GetControls(u.instance, Index))
}

func (u *TUpDown) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(UpDown_GetComponents(u.instance, AIndex))
}

