
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TTrackBar struct {
    IControl
    instance uintptr
}

func NewTrackBar(owner IComponent) *TTrackBar {
    t := new(TTrackBar)
    t.instance = TrackBar_Create(owner.Instance())
    return t
}

func TrackBarFromInst(inst uintptr) *TTrackBar {
    t := new(TTrackBar)
    t.instance = inst
    return t
}

func TrackBarFromObj(obj IObject) *TTrackBar {
    t := new(TTrackBar)
    t.instance = CheckPtr(obj)
    return t
}

func (t *TTrackBar) Free() {
    if t.instance != 0 {
        TrackBar_Free(t.instance)
    }
}

func (t *TTrackBar) Instance() uintptr {
    return t.instance
}

func (t *TTrackBar) IsValid() bool {
    return t.instance != 0
}

func (t *TTrackBar) CanFocus() bool {
    defer exceptionProc()
    return TrackBar_CanFocus(t.instance)
}

func (t *TTrackBar) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    TrackBar_FlipChildren(t.instance, AllLevels )
}

func (t *TTrackBar) Focused() bool {
    defer exceptionProc()
    return TrackBar_Focused(t.instance)
}

func (t *TTrackBar) HandleAllocated() bool {
    defer exceptionProc()
    return TrackBar_HandleAllocated(t.instance)
}

func (t *TTrackBar) Invalidate() {
    defer exceptionProc()
    TrackBar_Invalidate(t.instance)
}

func (t *TTrackBar) Realign() {
    defer exceptionProc()
    TrackBar_Realign(t.instance)
}

func (t *TTrackBar) Repaint() {
    defer exceptionProc()
    TrackBar_Repaint(t.instance)
}

func (t *TTrackBar) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    TrackBar_ScaleBy(t.instance, M , D )
}

func (t *TTrackBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    TrackBar_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight )
}

func (t *TTrackBar) SetFocus() {
    defer exceptionProc()
    TrackBar_SetFocus(t.instance)
}

func (t *TTrackBar) Update() {
    defer exceptionProc()
    TrackBar_Update(t.instance)
}

func (t *TTrackBar) BringToFront() {
    defer exceptionProc()
    TrackBar_BringToFront(t.instance)
}

func (t *TTrackBar) HasParent() bool {
    defer exceptionProc()
    return TrackBar_HasParent(t.instance)
}

func (t *TTrackBar) Hide() {
    defer exceptionProc()
    TrackBar_Hide(t.instance)
}

func (t *TTrackBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return TrackBar_Perform(t.instance, Msg , WParam , LParam )
}

func (t *TTrackBar) Refresh() {
    defer exceptionProc()
    TrackBar_Refresh(t.instance)
}

func (t *TTrackBar) SendToBack() {
    defer exceptionProc()
    TrackBar_SendToBack(t.instance)
}

func (t *TTrackBar) Show() {
    defer exceptionProc()
    TrackBar_Show(t.instance)
}

func (t *TTrackBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return TrackBar_GetTextBuf(t.instance, Buffer , BufSize )
}

func (t *TTrackBar) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(TrackBar_FindComponent(t.instance, AName ))
}

func (t *TTrackBar) GetNamePath() string {
    defer exceptionProc()
    return TrackBar_GetNamePath(t.instance)
}

func (t *TTrackBar) Assign(Source IObject) {
    defer exceptionProc()
    TrackBar_Assign(t.instance, CheckPtr(Source))
}

func (t *TTrackBar) ClassName() string {
    defer exceptionProc()
    return TrackBar_ClassName(t.instance)
}

func (t *TTrackBar) Equals(Obj IObject) bool {
    defer exceptionProc()
    return TrackBar_Equals(t.instance, CheckPtr(Obj))
}

func (t *TTrackBar) GetHashCode() int32 {
    defer exceptionProc()
    return TrackBar_GetHashCode(t.instance)
}

func (t *TTrackBar) ToString() string {
    defer exceptionProc()
    return TrackBar_ToString(t.instance)
}

func (t *TTrackBar) Align() TAlign {
    defer exceptionProc()
    return TrackBar_GetAlign(t.instance)
}

func (t *TTrackBar) SetAlign(value TAlign) {
    defer exceptionProc()
    TrackBar_SetAlign(t.instance, value)
}

func (t *TTrackBar) Anchors() TAnchors {
    defer exceptionProc()
    return TrackBar_GetAnchors(t.instance)
}

func (t *TTrackBar) SetAnchors(value TAnchors) {
    defer exceptionProc()
    TrackBar_SetAnchors(t.instance, value)
}

func (t *TTrackBar) BorderWidth() int32 {
    defer exceptionProc()
    return TrackBar_GetBorderWidth(t.instance)
}

func (t *TTrackBar) SetBorderWidth(value int32) {
    defer exceptionProc()
    TrackBar_SetBorderWidth(t.instance, value)
}

func (t *TTrackBar) DoubleBuffered() bool {
    defer exceptionProc()
    return TrackBar_GetDoubleBuffered(t.instance)
}

func (t *TTrackBar) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    TrackBar_SetDoubleBuffered(t.instance, value)
}

func (t *TTrackBar) Enabled() bool {
    defer exceptionProc()
    return TrackBar_GetEnabled(t.instance)
}

func (t *TTrackBar) SetEnabled(value bool) {
    defer exceptionProc()
    TrackBar_SetEnabled(t.instance, value)
}

func (t *TTrackBar) Orientation() TTrackBarOrientation {
    defer exceptionProc()
    return TrackBar_GetOrientation(t.instance)
}

func (t *TTrackBar) SetOrientation(value TTrackBarOrientation) {
    defer exceptionProc()
    TrackBar_SetOrientation(t.instance, value)
}

func (t *TTrackBar) ParentCtl3D() bool {
    defer exceptionProc()
    return TrackBar_GetParentCtl3D(t.instance)
}

func (t *TTrackBar) SetParentCtl3D(value bool) {
    defer exceptionProc()
    TrackBar_SetParentCtl3D(t.instance, value)
}

func (t *TTrackBar) ParentShowHint() bool {
    defer exceptionProc()
    return TrackBar_GetParentShowHint(t.instance)
}

func (t *TTrackBar) SetParentShowHint(value bool) {
    defer exceptionProc()
    TrackBar_SetParentShowHint(t.instance, value)
}

func (t *TTrackBar) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(TrackBar_GetPopupMenu(t.instance))
}

func (t *TTrackBar) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    TrackBar_SetPopupMenu(t.instance, CheckPtr(value))
}

func (t *TTrackBar) Position() int32 {
    defer exceptionProc()
    return TrackBar_GetPosition(t.instance)
}

func (t *TTrackBar) SetPosition(value int32) {
    defer exceptionProc()
    TrackBar_SetPosition(t.instance, value)
}

func (t *TTrackBar) SelStart() int32 {
    defer exceptionProc()
    return TrackBar_GetSelStart(t.instance)
}

func (t *TTrackBar) SetSelStart(value int32) {
    defer exceptionProc()
    TrackBar_SetSelStart(t.instance, value)
}

func (t *TTrackBar) ShowHint() bool {
    defer exceptionProc()
    return TrackBar_GetShowHint(t.instance)
}

func (t *TTrackBar) SetShowHint(value bool) {
    defer exceptionProc()
    TrackBar_SetShowHint(t.instance, value)
}

func (t *TTrackBar) TabOrder() int16 {
    defer exceptionProc()
    return TrackBar_GetTabOrder(t.instance)
}

func (t *TTrackBar) SetTabOrder(value int16) {
    defer exceptionProc()
    TrackBar_SetTabOrder(t.instance, value)
}

func (t *TTrackBar) TabStop() bool {
    defer exceptionProc()
    return TrackBar_GetTabStop(t.instance)
}

func (t *TTrackBar) SetTabStop(value bool) {
    defer exceptionProc()
    TrackBar_SetTabStop(t.instance, value)
}

func (t *TTrackBar) Visible() bool {
    defer exceptionProc()
    return TrackBar_GetVisible(t.instance)
}

func (t *TTrackBar) SetVisible(value bool) {
    defer exceptionProc()
    TrackBar_SetVisible(t.instance, value)
}

func (t *TTrackBar) SetOnChange(fn TNotifyEvent) {
    TrackBar_SetOnChange(t.instance, fn)
}

func (t *TTrackBar) SetOnEnter(fn TNotifyEvent) {
    TrackBar_SetOnEnter(t.instance, fn)
}

func (t *TTrackBar) SetOnExit(fn TNotifyEvent) {
    TrackBar_SetOnExit(t.instance, fn)
}

func (t *TTrackBar) SetOnKeyDown(fn TKeyEvent) {
    TrackBar_SetOnKeyDown(t.instance, fn)
}

func (t *TTrackBar) SetOnKeyPress(fn TKeyPressEvent) {
    TrackBar_SetOnKeyPress(t.instance, fn)
}

func (t *TTrackBar) SetOnKeyUp(fn TKeyEvent) {
    TrackBar_SetOnKeyUp(t.instance, fn)
}

func (t *TTrackBar) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(TrackBar_GetBrush(t.instance))
}

func (t *TTrackBar) ControlCount() int32 {
    defer exceptionProc()
    return TrackBar_GetControlCount(t.instance)
}

func (t *TTrackBar) Handle() HWND {
    defer exceptionProc()
    return TrackBar_GetHandle(t.instance)
}

func (t *TTrackBar) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(TrackBar_GetAction(t.instance))
}

func (t *TTrackBar) SetAction(value IComponent) {
    defer exceptionProc()
    TrackBar_SetAction(t.instance, CheckPtr(value))
}

func (t *TTrackBar) BoundsRect() TRect {
    defer exceptionProc()
    return TrackBar_GetBoundsRect(t.instance)
}

func (t *TTrackBar) SetBoundsRect(value TRect) {
    defer exceptionProc()
    TrackBar_SetBoundsRect(t.instance, value)
}

func (t *TTrackBar) ClientHeight() int32 {
    defer exceptionProc()
    return TrackBar_GetClientHeight(t.instance)
}

func (t *TTrackBar) SetClientHeight(value int32) {
    defer exceptionProc()
    TrackBar_SetClientHeight(t.instance, value)
}

func (t *TTrackBar) ClientRect() TRect {
    defer exceptionProc()
    return TrackBar_GetClientRect(t.instance)
}

func (t *TTrackBar) ClientWidth() int32 {
    defer exceptionProc()
    return TrackBar_GetClientWidth(t.instance)
}

func (t *TTrackBar) SetClientWidth(value int32) {
    defer exceptionProc()
    TrackBar_SetClientWidth(t.instance, value)
}

func (t *TTrackBar) ExplicitLeft() int32 {
    defer exceptionProc()
    return TrackBar_GetExplicitLeft(t.instance)
}

func (t *TTrackBar) ExplicitTop() int32 {
    defer exceptionProc()
    return TrackBar_GetExplicitTop(t.instance)
}

func (t *TTrackBar) ExplicitWidth() int32 {
    defer exceptionProc()
    return TrackBar_GetExplicitWidth(t.instance)
}

func (t *TTrackBar) ExplicitHeight() int32 {
    defer exceptionProc()
    return TrackBar_GetExplicitHeight(t.instance)
}

func (t *TTrackBar) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(TrackBar_GetParent(t.instance))
}

func (t *TTrackBar) SetParent(value IControl) {
    defer exceptionProc()
    TrackBar_SetParent(t.instance, CheckPtr(value))
}

func (t *TTrackBar) AlignWithMargins() bool {
    defer exceptionProc()
    return TrackBar_GetAlignWithMargins(t.instance)
}

func (t *TTrackBar) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    TrackBar_SetAlignWithMargins(t.instance, value)
}

func (t *TTrackBar) Left() int32 {
    defer exceptionProc()
    return TrackBar_GetLeft(t.instance)
}

func (t *TTrackBar) SetLeft(value int32) {
    defer exceptionProc()
    TrackBar_SetLeft(t.instance, value)
}

func (t *TTrackBar) Top() int32 {
    defer exceptionProc()
    return TrackBar_GetTop(t.instance)
}

func (t *TTrackBar) SetTop(value int32) {
    defer exceptionProc()
    TrackBar_SetTop(t.instance, value)
}

func (t *TTrackBar) Width() int32 {
    defer exceptionProc()
    return TrackBar_GetWidth(t.instance)
}

func (t *TTrackBar) SetWidth(value int32) {
    defer exceptionProc()
    TrackBar_SetWidth(t.instance, value)
}

func (t *TTrackBar) Height() int32 {
    defer exceptionProc()
    return TrackBar_GetHeight(t.instance)
}

func (t *TTrackBar) SetHeight(value int32) {
    defer exceptionProc()
    TrackBar_SetHeight(t.instance, value)
}

func (t *TTrackBar) Cursor() TCursor {
    defer exceptionProc()
    return TrackBar_GetCursor(t.instance)
}

func (t *TTrackBar) SetCursor(value TCursor) {
    defer exceptionProc()
    TrackBar_SetCursor(t.instance, value)
}

func (t *TTrackBar) Hint() string {
    defer exceptionProc()
    return TrackBar_GetHint(t.instance)
}

func (t *TTrackBar) SetHint(value string) {
    defer exceptionProc()
    TrackBar_SetHint(t.instance, value)
}

func (t *TTrackBar) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(TrackBar_GetMargins(t.instance))
}

func (t *TTrackBar) SetMargins(value *TMargins) {
    defer exceptionProc()
    TrackBar_SetMargins(t.instance, CheckPtr(value))
}

func (t *TTrackBar) ComponentCount() int32 {
    defer exceptionProc()
    return TrackBar_GetComponentCount(t.instance)
}

func (t *TTrackBar) ComponentIndex() int32 {
    defer exceptionProc()
    return TrackBar_GetComponentIndex(t.instance)
}

func (t *TTrackBar) SetComponentIndex(value int32) {
    defer exceptionProc()
    TrackBar_SetComponentIndex(t.instance, value)
}

func (t *TTrackBar) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(TrackBar_GetOwner(t.instance))
}

func (t *TTrackBar) Name() string {
    defer exceptionProc()
    return TrackBar_GetName(t.instance)
}

func (t *TTrackBar) SetName(value string) {
    defer exceptionProc()
    TrackBar_SetName(t.instance, value)
}

func (t *TTrackBar) Tag() int {
    defer exceptionProc()
    return TrackBar_GetTag(t.instance)
}

func (t *TTrackBar) SetTag(value int) {
    defer exceptionProc()
    TrackBar_SetTag(t.instance, value)
}

func (t *TTrackBar) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(TrackBar_GetControls(t.instance, Index))
}

func (t *TTrackBar) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(TrackBar_GetComponents(t.instance, AIndex))
}

