
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

type TProgressBar struct {
    IControl
    instance uintptr
}

func NewProgressBar(owner IComponent) *TProgressBar {
    p := new(TProgressBar)
    p.instance = ProgressBar_Create(owner.Instance())
    return p
}

func ProgressBarFromInst(inst uintptr) *TProgressBar {
    p := new(TProgressBar)
    p.instance = inst
    return p
}

func ProgressBarFromObj(obj IObject) *TProgressBar {
    p := new(TProgressBar)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TProgressBar) Free() {
    if p.instance != 0 {
        ProgressBar_Free(p.instance)
    }
}

func (p *TProgressBar) Instance() uintptr {
    return p.instance
}

func (p *TProgressBar) IsValid() bool {
    return p.instance != 0
}

func (p *TProgressBar) CanFocus() bool {
    defer exceptionProc()
    return ProgressBar_CanFocus(p.instance)
}

func (p *TProgressBar) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    ProgressBar_FlipChildren(p.instance, AllLevels )
}

func (p *TProgressBar) Focused() bool {
    defer exceptionProc()
    return ProgressBar_Focused(p.instance)
}

func (p *TProgressBar) HandleAllocated() bool {
    defer exceptionProc()
    return ProgressBar_HandleAllocated(p.instance)
}

func (p *TProgressBar) Invalidate() {
    defer exceptionProc()
    ProgressBar_Invalidate(p.instance)
}

func (p *TProgressBar) Realign() {
    defer exceptionProc()
    ProgressBar_Realign(p.instance)
}

func (p *TProgressBar) Repaint() {
    defer exceptionProc()
    ProgressBar_Repaint(p.instance)
}

func (p *TProgressBar) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    ProgressBar_ScaleBy(p.instance, M , D )
}

func (p *TProgressBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    ProgressBar_SetBounds(p.instance, ALeft , ATop , AWidth , AHeight )
}

func (p *TProgressBar) SetFocus() {
    defer exceptionProc()
    ProgressBar_SetFocus(p.instance)
}

func (p *TProgressBar) Update() {
    defer exceptionProc()
    ProgressBar_Update(p.instance)
}

func (p *TProgressBar) BringToFront() {
    defer exceptionProc()
    ProgressBar_BringToFront(p.instance)
}

func (p *TProgressBar) HasParent() bool {
    defer exceptionProc()
    return ProgressBar_HasParent(p.instance)
}

func (p *TProgressBar) Hide() {
    defer exceptionProc()
    ProgressBar_Hide(p.instance)
}

func (p *TProgressBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return ProgressBar_Perform(p.instance, Msg , WParam , LParam )
}

func (p *TProgressBar) Refresh() {
    defer exceptionProc()
    ProgressBar_Refresh(p.instance)
}

func (p *TProgressBar) SendToBack() {
    defer exceptionProc()
    ProgressBar_SendToBack(p.instance)
}

func (p *TProgressBar) Show() {
    defer exceptionProc()
    ProgressBar_Show(p.instance)
}

func (p *TProgressBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return ProgressBar_GetTextBuf(p.instance, Buffer , BufSize )
}

func (p *TProgressBar) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ProgressBar_FindComponent(p.instance, AName ))
}

func (p *TProgressBar) GetNamePath() string {
    defer exceptionProc()
    return ProgressBar_GetNamePath(p.instance)
}

func (p *TProgressBar) Assign(Source IObject) {
    defer exceptionProc()
    ProgressBar_Assign(p.instance, CheckPtr(Source))
}

func (p *TProgressBar) ClassName() string {
    defer exceptionProc()
    return ProgressBar_ClassName(p.instance)
}

func (p *TProgressBar) Equals(Obj IObject) bool {
    defer exceptionProc()
    return ProgressBar_Equals(p.instance, CheckPtr(Obj))
}

func (p *TProgressBar) GetHashCode() int32 {
    defer exceptionProc()
    return ProgressBar_GetHashCode(p.instance)
}

func (p *TProgressBar) ToString() string {
    defer exceptionProc()
    return ProgressBar_ToString(p.instance)
}

func (p *TProgressBar) Align() TAlign {
    defer exceptionProc()
    return ProgressBar_GetAlign(p.instance)
}

func (p *TProgressBar) SetAlign(value TAlign) {
    defer exceptionProc()
    ProgressBar_SetAlign(p.instance, value)
}

func (p *TProgressBar) Anchors() TAnchors {
    defer exceptionProc()
    return ProgressBar_GetAnchors(p.instance)
}

func (p *TProgressBar) SetAnchors(value TAnchors) {
    defer exceptionProc()
    ProgressBar_SetAnchors(p.instance, value)
}

func (p *TProgressBar) BorderWidth() int32 {
    defer exceptionProc()
    return ProgressBar_GetBorderWidth(p.instance)
}

func (p *TProgressBar) SetBorderWidth(value int32) {
    defer exceptionProc()
    ProgressBar_SetBorderWidth(p.instance, value)
}

func (p *TProgressBar) DoubleBuffered() bool {
    defer exceptionProc()
    return ProgressBar_GetDoubleBuffered(p.instance)
}

func (p *TProgressBar) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    ProgressBar_SetDoubleBuffered(p.instance, value)
}

func (p *TProgressBar) Enabled() bool {
    defer exceptionProc()
    return ProgressBar_GetEnabled(p.instance)
}

func (p *TProgressBar) SetEnabled(value bool) {
    defer exceptionProc()
    ProgressBar_SetEnabled(p.instance, value)
}

func (p *TProgressBar) Hint() string {
    defer exceptionProc()
    return ProgressBar_GetHint(p.instance)
}

func (p *TProgressBar) SetHint(value string) {
    defer exceptionProc()
    ProgressBar_SetHint(p.instance, value)
}

func (p *TProgressBar) Orientation() TProgressBarOrientation {
    defer exceptionProc()
    return ProgressBar_GetOrientation(p.instance)
}

func (p *TProgressBar) SetOrientation(value TProgressBarOrientation) {
    defer exceptionProc()
    ProgressBar_SetOrientation(p.instance, value)
}

func (p *TProgressBar) ParentShowHint() bool {
    defer exceptionProc()
    return ProgressBar_GetParentShowHint(p.instance)
}

func (p *TProgressBar) SetParentShowHint(value bool) {
    defer exceptionProc()
    ProgressBar_SetParentShowHint(p.instance, value)
}

func (p *TProgressBar) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(ProgressBar_GetPopupMenu(p.instance))
}

func (p *TProgressBar) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    ProgressBar_SetPopupMenu(p.instance, CheckPtr(value))
}

func (p *TProgressBar) Position() int32 {
    defer exceptionProc()
    return ProgressBar_GetPosition(p.instance)
}

func (p *TProgressBar) SetPosition(value int32) {
    defer exceptionProc()
    ProgressBar_SetPosition(p.instance, value)
}

func (p *TProgressBar) Style() TProgressBarStyle {
    defer exceptionProc()
    return ProgressBar_GetStyle(p.instance)
}

func (p *TProgressBar) SetStyle(value TProgressBarStyle) {
    defer exceptionProc()
    ProgressBar_SetStyle(p.instance, value)
}

func (p *TProgressBar) State() TProgressBarState {
    defer exceptionProc()
    return ProgressBar_GetState(p.instance)
}

func (p *TProgressBar) SetState(value TProgressBarState) {
    defer exceptionProc()
    ProgressBar_SetState(p.instance, value)
}

func (p *TProgressBar) ShowHint() bool {
    defer exceptionProc()
    return ProgressBar_GetShowHint(p.instance)
}

func (p *TProgressBar) SetShowHint(value bool) {
    defer exceptionProc()
    ProgressBar_SetShowHint(p.instance, value)
}

func (p *TProgressBar) TabOrder() int16 {
    defer exceptionProc()
    return ProgressBar_GetTabOrder(p.instance)
}

func (p *TProgressBar) SetTabOrder(value int16) {
    defer exceptionProc()
    ProgressBar_SetTabOrder(p.instance, value)
}

func (p *TProgressBar) TabStop() bool {
    defer exceptionProc()
    return ProgressBar_GetTabStop(p.instance)
}

func (p *TProgressBar) SetTabStop(value bool) {
    defer exceptionProc()
    ProgressBar_SetTabStop(p.instance, value)
}

func (p *TProgressBar) Visible() bool {
    defer exceptionProc()
    return ProgressBar_GetVisible(p.instance)
}

func (p *TProgressBar) SetVisible(value bool) {
    defer exceptionProc()
    ProgressBar_SetVisible(p.instance, value)
}

func (p *TProgressBar) SetOnEnter(fn TNotifyEvent) {
    ProgressBar_SetOnEnter(p.instance, fn)
}

func (p *TProgressBar) SetOnExit(fn TNotifyEvent) {
    ProgressBar_SetOnExit(p.instance, fn)
}

func (p *TProgressBar) SetOnMouseDown(fn TMouseEvent) {
    ProgressBar_SetOnMouseDown(p.instance, fn)
}

func (p *TProgressBar) SetOnMouseEnter(fn TNotifyEvent) {
    ProgressBar_SetOnMouseEnter(p.instance, fn)
}

func (p *TProgressBar) SetOnMouseLeave(fn TNotifyEvent) {
    ProgressBar_SetOnMouseLeave(p.instance, fn)
}

func (p *TProgressBar) SetOnMouseMove(fn TMouseMoveEvent) {
    ProgressBar_SetOnMouseMove(p.instance, fn)
}

func (p *TProgressBar) SetOnMouseUp(fn TMouseEvent) {
    ProgressBar_SetOnMouseUp(p.instance, fn)
}

func (p *TProgressBar) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(ProgressBar_GetBrush(p.instance))
}

func (p *TProgressBar) ControlCount() int32 {
    defer exceptionProc()
    return ProgressBar_GetControlCount(p.instance)
}

func (p *TProgressBar) Handle() HWND {
    defer exceptionProc()
    return ProgressBar_GetHandle(p.instance)
}

func (p *TProgressBar) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(ProgressBar_GetAction(p.instance))
}

func (p *TProgressBar) SetAction(value IComponent) {
    defer exceptionProc()
    ProgressBar_SetAction(p.instance, CheckPtr(value))
}

func (p *TProgressBar) BoundsRect() TRect {
    defer exceptionProc()
    return ProgressBar_GetBoundsRect(p.instance)
}

func (p *TProgressBar) SetBoundsRect(value TRect) {
    defer exceptionProc()
    ProgressBar_SetBoundsRect(p.instance, value)
}

func (p *TProgressBar) ClientHeight() int32 {
    defer exceptionProc()
    return ProgressBar_GetClientHeight(p.instance)
}

func (p *TProgressBar) SetClientHeight(value int32) {
    defer exceptionProc()
    ProgressBar_SetClientHeight(p.instance, value)
}

func (p *TProgressBar) ClientRect() TRect {
    defer exceptionProc()
    return ProgressBar_GetClientRect(p.instance)
}

func (p *TProgressBar) ClientWidth() int32 {
    defer exceptionProc()
    return ProgressBar_GetClientWidth(p.instance)
}

func (p *TProgressBar) SetClientWidth(value int32) {
    defer exceptionProc()
    ProgressBar_SetClientWidth(p.instance, value)
}

func (p *TProgressBar) ExplicitLeft() int32 {
    defer exceptionProc()
    return ProgressBar_GetExplicitLeft(p.instance)
}

func (p *TProgressBar) ExplicitTop() int32 {
    defer exceptionProc()
    return ProgressBar_GetExplicitTop(p.instance)
}

func (p *TProgressBar) ExplicitWidth() int32 {
    defer exceptionProc()
    return ProgressBar_GetExplicitWidth(p.instance)
}

func (p *TProgressBar) ExplicitHeight() int32 {
    defer exceptionProc()
    return ProgressBar_GetExplicitHeight(p.instance)
}

func (p *TProgressBar) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(ProgressBar_GetParent(p.instance))
}

func (p *TProgressBar) SetParent(value IControl) {
    defer exceptionProc()
    ProgressBar_SetParent(p.instance, CheckPtr(value))
}

func (p *TProgressBar) AlignWithMargins() bool {
    defer exceptionProc()
    return ProgressBar_GetAlignWithMargins(p.instance)
}

func (p *TProgressBar) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    ProgressBar_SetAlignWithMargins(p.instance, value)
}

func (p *TProgressBar) Left() int32 {
    defer exceptionProc()
    return ProgressBar_GetLeft(p.instance)
}

func (p *TProgressBar) SetLeft(value int32) {
    defer exceptionProc()
    ProgressBar_SetLeft(p.instance, value)
}

func (p *TProgressBar) Top() int32 {
    defer exceptionProc()
    return ProgressBar_GetTop(p.instance)
}

func (p *TProgressBar) SetTop(value int32) {
    defer exceptionProc()
    ProgressBar_SetTop(p.instance, value)
}

func (p *TProgressBar) Width() int32 {
    defer exceptionProc()
    return ProgressBar_GetWidth(p.instance)
}

func (p *TProgressBar) SetWidth(value int32) {
    defer exceptionProc()
    ProgressBar_SetWidth(p.instance, value)
}

func (p *TProgressBar) Height() int32 {
    defer exceptionProc()
    return ProgressBar_GetHeight(p.instance)
}

func (p *TProgressBar) SetHeight(value int32) {
    defer exceptionProc()
    ProgressBar_SetHeight(p.instance, value)
}

func (p *TProgressBar) Cursor() TCursor {
    defer exceptionProc()
    return ProgressBar_GetCursor(p.instance)
}

func (p *TProgressBar) SetCursor(value TCursor) {
    defer exceptionProc()
    ProgressBar_SetCursor(p.instance, value)
}

func (p *TProgressBar) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(ProgressBar_GetMargins(p.instance))
}

func (p *TProgressBar) SetMargins(value *TMargins) {
    defer exceptionProc()
    ProgressBar_SetMargins(p.instance, CheckPtr(value))
}

func (p *TProgressBar) ComponentCount() int32 {
    defer exceptionProc()
    return ProgressBar_GetComponentCount(p.instance)
}

func (p *TProgressBar) ComponentIndex() int32 {
    defer exceptionProc()
    return ProgressBar_GetComponentIndex(p.instance)
}

func (p *TProgressBar) SetComponentIndex(value int32) {
    defer exceptionProc()
    ProgressBar_SetComponentIndex(p.instance, value)
}

func (p *TProgressBar) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ProgressBar_GetOwner(p.instance))
}

func (p *TProgressBar) Name() string {
    defer exceptionProc()
    return ProgressBar_GetName(p.instance)
}

func (p *TProgressBar) SetName(value string) {
    defer exceptionProc()
    ProgressBar_SetName(p.instance, value)
}

func (p *TProgressBar) Tag() int {
    defer exceptionProc()
    return ProgressBar_GetTag(p.instance)
}

func (p *TProgressBar) SetTag(value int) {
    defer exceptionProc()
    ProgressBar_SetTag(p.instance, value)
}

func (p *TProgressBar) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(ProgressBar_GetControls(p.instance, Index))
}

func (p *TProgressBar) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ProgressBar_GetComponents(p.instance, AIndex))
}

