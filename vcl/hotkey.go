
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type THotKey struct {
    IControl
    instance uintptr
}

func NewHotKey(owner IComponent) *THotKey {
    h := new(THotKey)
    h.instance = HotKey_Create(owner.Instance())
    return h
}

func HotKeyFromInst(inst uintptr) *THotKey {
    h := new(THotKey)
    h.instance = inst
    return h
}

func HotKeyFromObj(obj IObject) *THotKey {
    h := new(THotKey)
    h.instance = CheckPtr(obj)
    return h
}

func (h *THotKey) Free() {
    if h.instance != 0 {
        HotKey_Free(h.instance)
    }
}

func (h *THotKey) Instance() uintptr {
    return h.instance
}

func (h *THotKey) IsValid() bool {
    return h.instance != 0
}

func (h *THotKey) CanFocus() bool {
    defer exceptionProc()
    return HotKey_CanFocus(h.instance)
}

func (h *THotKey) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    HotKey_FlipChildren(h.instance, AllLevels )
}

func (h *THotKey) Focused() bool {
    defer exceptionProc()
    return HotKey_Focused(h.instance)
}

func (h *THotKey) HandleAllocated() bool {
    defer exceptionProc()
    return HotKey_HandleAllocated(h.instance)
}

func (h *THotKey) Invalidate() {
    defer exceptionProc()
    HotKey_Invalidate(h.instance)
}

func (h *THotKey) Realign() {
    defer exceptionProc()
    HotKey_Realign(h.instance)
}

func (h *THotKey) Repaint() {
    defer exceptionProc()
    HotKey_Repaint(h.instance)
}

func (h *THotKey) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    HotKey_ScaleBy(h.instance, M , D )
}

func (h *THotKey) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    HotKey_SetBounds(h.instance, ALeft , ATop , AWidth , AHeight )
}

func (h *THotKey) SetFocus() {
    defer exceptionProc()
    HotKey_SetFocus(h.instance)
}

func (h *THotKey) Update() {
    defer exceptionProc()
    HotKey_Update(h.instance)
}

func (h *THotKey) BringToFront() {
    defer exceptionProc()
    HotKey_BringToFront(h.instance)
}

func (h *THotKey) HasParent() bool {
    defer exceptionProc()
    return HotKey_HasParent(h.instance)
}

func (h *THotKey) Hide() {
    defer exceptionProc()
    HotKey_Hide(h.instance)
}

func (h *THotKey) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return HotKey_Perform(h.instance, Msg , WParam , LParam )
}

func (h *THotKey) Refresh() {
    defer exceptionProc()
    HotKey_Refresh(h.instance)
}

func (h *THotKey) SendToBack() {
    defer exceptionProc()
    HotKey_SendToBack(h.instance)
}

func (h *THotKey) Show() {
    defer exceptionProc()
    HotKey_Show(h.instance)
}

func (h *THotKey) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return HotKey_GetTextBuf(h.instance, Buffer , BufSize )
}

func (h *THotKey) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(HotKey_FindComponent(h.instance, AName ))
}

func (h *THotKey) GetNamePath() string {
    defer exceptionProc()
    return HotKey_GetNamePath(h.instance)
}

func (h *THotKey) Assign(Source IObject) {
    defer exceptionProc()
    HotKey_Assign(h.instance, CheckPtr(Source))
}

func (h *THotKey) ClassName() string {
    defer exceptionProc()
    return HotKey_ClassName(h.instance)
}

func (h *THotKey) Equals(Obj IObject) bool {
    defer exceptionProc()
    return HotKey_Equals(h.instance, CheckPtr(Obj))
}

func (h *THotKey) GetHashCode() int32 {
    defer exceptionProc()
    return HotKey_GetHashCode(h.instance)
}

func (h *THotKey) ToString() string {
    defer exceptionProc()
    return HotKey_ToString(h.instance)
}

func (h *THotKey) Anchors() TAnchors {
    defer exceptionProc()
    return HotKey_GetAnchors(h.instance)
}

func (h *THotKey) SetAnchors(value TAnchors) {
    defer exceptionProc()
    HotKey_SetAnchors(h.instance, value)
}

func (h *THotKey) AutoSize() bool {
    defer exceptionProc()
    return HotKey_GetAutoSize(h.instance)
}

func (h *THotKey) SetAutoSize(value bool) {
    defer exceptionProc()
    HotKey_SetAutoSize(h.instance, value)
}

func (h *THotKey) Enabled() bool {
    defer exceptionProc()
    return HotKey_GetEnabled(h.instance)
}

func (h *THotKey) SetEnabled(value bool) {
    defer exceptionProc()
    HotKey_SetEnabled(h.instance, value)
}

func (h *THotKey) Hint() string {
    defer exceptionProc()
    return HotKey_GetHint(h.instance)
}

func (h *THotKey) SetHint(value string) {
    defer exceptionProc()
    HotKey_SetHint(h.instance, value)
}

func (h *THotKey) ParentShowHint() bool {
    defer exceptionProc()
    return HotKey_GetParentShowHint(h.instance)
}

func (h *THotKey) SetParentShowHint(value bool) {
    defer exceptionProc()
    HotKey_SetParentShowHint(h.instance, value)
}

func (h *THotKey) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(HotKey_GetPopupMenu(h.instance))
}

func (h *THotKey) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    HotKey_SetPopupMenu(h.instance, CheckPtr(value))
}

func (h *THotKey) ShowHint() bool {
    defer exceptionProc()
    return HotKey_GetShowHint(h.instance)
}

func (h *THotKey) SetShowHint(value bool) {
    defer exceptionProc()
    HotKey_SetShowHint(h.instance, value)
}

func (h *THotKey) TabOrder() int16 {
    defer exceptionProc()
    return HotKey_GetTabOrder(h.instance)
}

func (h *THotKey) SetTabOrder(value int16) {
    defer exceptionProc()
    HotKey_SetTabOrder(h.instance, value)
}

func (h *THotKey) TabStop() bool {
    defer exceptionProc()
    return HotKey_GetTabStop(h.instance)
}

func (h *THotKey) SetTabStop(value bool) {
    defer exceptionProc()
    HotKey_SetTabStop(h.instance, value)
}

func (h *THotKey) Visible() bool {
    defer exceptionProc()
    return HotKey_GetVisible(h.instance)
}

func (h *THotKey) SetVisible(value bool) {
    defer exceptionProc()
    HotKey_SetVisible(h.instance, value)
}

func (h *THotKey) SetOnChange(fn TNotifyEvent) {
    HotKey_SetOnChange(h.instance, fn)
}

func (h *THotKey) SetOnEnter(fn TNotifyEvent) {
    HotKey_SetOnEnter(h.instance, fn)
}

func (h *THotKey) SetOnExit(fn TNotifyEvent) {
    HotKey_SetOnExit(h.instance, fn)
}

func (h *THotKey) SetOnMouseDown(fn TMouseEvent) {
    HotKey_SetOnMouseDown(h.instance, fn)
}

func (h *THotKey) SetOnMouseEnter(fn TNotifyEvent) {
    HotKey_SetOnMouseEnter(h.instance, fn)
}

func (h *THotKey) SetOnMouseLeave(fn TNotifyEvent) {
    HotKey_SetOnMouseLeave(h.instance, fn)
}

func (h *THotKey) SetOnMouseMove(fn TMouseMoveEvent) {
    HotKey_SetOnMouseMove(h.instance, fn)
}

func (h *THotKey) SetOnMouseUp(fn TMouseEvent) {
    HotKey_SetOnMouseUp(h.instance, fn)
}

func (h *THotKey) DoubleBuffered() bool {
    defer exceptionProc()
    return HotKey_GetDoubleBuffered(h.instance)
}

func (h *THotKey) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    HotKey_SetDoubleBuffered(h.instance, value)
}

func (h *THotKey) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(HotKey_GetBrush(h.instance))
}

func (h *THotKey) ControlCount() int32 {
    defer exceptionProc()
    return HotKey_GetControlCount(h.instance)
}

func (h *THotKey) Handle() HWND {
    defer exceptionProc()
    return HotKey_GetHandle(h.instance)
}

func (h *THotKey) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(HotKey_GetAction(h.instance))
}

func (h *THotKey) SetAction(value IComponent) {
    defer exceptionProc()
    HotKey_SetAction(h.instance, CheckPtr(value))
}

func (h *THotKey) Align() TAlign {
    defer exceptionProc()
    return HotKey_GetAlign(h.instance)
}

func (h *THotKey) SetAlign(value TAlign) {
    defer exceptionProc()
    HotKey_SetAlign(h.instance, value)
}

func (h *THotKey) BoundsRect() TRect {
    defer exceptionProc()
    return HotKey_GetBoundsRect(h.instance)
}

func (h *THotKey) SetBoundsRect(value TRect) {
    defer exceptionProc()
    HotKey_SetBoundsRect(h.instance, value)
}

func (h *THotKey) ClientHeight() int32 {
    defer exceptionProc()
    return HotKey_GetClientHeight(h.instance)
}

func (h *THotKey) SetClientHeight(value int32) {
    defer exceptionProc()
    HotKey_SetClientHeight(h.instance, value)
}

func (h *THotKey) ClientRect() TRect {
    defer exceptionProc()
    return HotKey_GetClientRect(h.instance)
}

func (h *THotKey) ClientWidth() int32 {
    defer exceptionProc()
    return HotKey_GetClientWidth(h.instance)
}

func (h *THotKey) SetClientWidth(value int32) {
    defer exceptionProc()
    HotKey_SetClientWidth(h.instance, value)
}

func (h *THotKey) ExplicitLeft() int32 {
    defer exceptionProc()
    return HotKey_GetExplicitLeft(h.instance)
}

func (h *THotKey) ExplicitTop() int32 {
    defer exceptionProc()
    return HotKey_GetExplicitTop(h.instance)
}

func (h *THotKey) ExplicitWidth() int32 {
    defer exceptionProc()
    return HotKey_GetExplicitWidth(h.instance)
}

func (h *THotKey) ExplicitHeight() int32 {
    defer exceptionProc()
    return HotKey_GetExplicitHeight(h.instance)
}

func (h *THotKey) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(HotKey_GetParent(h.instance))
}

func (h *THotKey) SetParent(value IControl) {
    defer exceptionProc()
    HotKey_SetParent(h.instance, CheckPtr(value))
}

func (h *THotKey) AlignWithMargins() bool {
    defer exceptionProc()
    return HotKey_GetAlignWithMargins(h.instance)
}

func (h *THotKey) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    HotKey_SetAlignWithMargins(h.instance, value)
}

func (h *THotKey) Left() int32 {
    defer exceptionProc()
    return HotKey_GetLeft(h.instance)
}

func (h *THotKey) SetLeft(value int32) {
    defer exceptionProc()
    HotKey_SetLeft(h.instance, value)
}

func (h *THotKey) Top() int32 {
    defer exceptionProc()
    return HotKey_GetTop(h.instance)
}

func (h *THotKey) SetTop(value int32) {
    defer exceptionProc()
    HotKey_SetTop(h.instance, value)
}

func (h *THotKey) Width() int32 {
    defer exceptionProc()
    return HotKey_GetWidth(h.instance)
}

func (h *THotKey) SetWidth(value int32) {
    defer exceptionProc()
    HotKey_SetWidth(h.instance, value)
}

func (h *THotKey) Height() int32 {
    defer exceptionProc()
    return HotKey_GetHeight(h.instance)
}

func (h *THotKey) SetHeight(value int32) {
    defer exceptionProc()
    HotKey_SetHeight(h.instance, value)
}

func (h *THotKey) Cursor() TCursor {
    defer exceptionProc()
    return HotKey_GetCursor(h.instance)
}

func (h *THotKey) SetCursor(value TCursor) {
    defer exceptionProc()
    HotKey_SetCursor(h.instance, value)
}

func (h *THotKey) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(HotKey_GetMargins(h.instance))
}

func (h *THotKey) SetMargins(value *TMargins) {
    defer exceptionProc()
    HotKey_SetMargins(h.instance, CheckPtr(value))
}

func (h *THotKey) ComponentCount() int32 {
    defer exceptionProc()
    return HotKey_GetComponentCount(h.instance)
}

func (h *THotKey) ComponentIndex() int32 {
    defer exceptionProc()
    return HotKey_GetComponentIndex(h.instance)
}

func (h *THotKey) SetComponentIndex(value int32) {
    defer exceptionProc()
    HotKey_SetComponentIndex(h.instance, value)
}

func (h *THotKey) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(HotKey_GetOwner(h.instance))
}

func (h *THotKey) Name() string {
    defer exceptionProc()
    return HotKey_GetName(h.instance)
}

func (h *THotKey) SetName(value string) {
    defer exceptionProc()
    HotKey_SetName(h.instance, value)
}

func (h *THotKey) Tag() int {
    defer exceptionProc()
    return HotKey_GetTag(h.instance)
}

func (h *THotKey) SetTag(value int) {
    defer exceptionProc()
    HotKey_SetTag(h.instance, value)
}

func (h *THotKey) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(HotKey_GetControls(h.instance, Index))
}

func (h *THotKey) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(HotKey_GetComponents(h.instance, AIndex))
}

