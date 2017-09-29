
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TStatusBar struct {
    IControl
    instance uintptr
}

func NewStatusBar(owner IComponent) *TStatusBar {
    s := new(TStatusBar)
    s.instance = StatusBar_Create(owner.Instance())
    return s
}

func StatusBarFromInst(inst uintptr) *TStatusBar {
    s := new(TStatusBar)
    s.instance = inst
    return s
}

func StatusBarFromObj(obj IObject) *TStatusBar {
    s := new(TStatusBar)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TStatusBar) Free() {
    if s.instance != 0 {
        StatusBar_Free(s.instance)
    }
}

func (s *TStatusBar) Instance() uintptr {
    return s.instance
}

func (s *TStatusBar) IsValid() bool {
    return s.instance != 0
}

func (s *TStatusBar) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    StatusBar_FlipChildren(s.instance, AllLevels )
}

func (s *TStatusBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    StatusBar_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight )
}

func (s *TStatusBar) CanFocus() bool {
    defer exceptionProc()
    return StatusBar_CanFocus(s.instance)
}

func (s *TStatusBar) Focused() bool {
    defer exceptionProc()
    return StatusBar_Focused(s.instance)
}

func (s *TStatusBar) HandleAllocated() bool {
    defer exceptionProc()
    return StatusBar_HandleAllocated(s.instance)
}

func (s *TStatusBar) Invalidate() {
    defer exceptionProc()
    StatusBar_Invalidate(s.instance)
}

func (s *TStatusBar) Realign() {
    defer exceptionProc()
    StatusBar_Realign(s.instance)
}

func (s *TStatusBar) Repaint() {
    defer exceptionProc()
    StatusBar_Repaint(s.instance)
}

func (s *TStatusBar) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    StatusBar_ScaleBy(s.instance, M , D )
}

func (s *TStatusBar) SetFocus() {
    defer exceptionProc()
    StatusBar_SetFocus(s.instance)
}

func (s *TStatusBar) Update() {
    defer exceptionProc()
    StatusBar_Update(s.instance)
}

func (s *TStatusBar) BringToFront() {
    defer exceptionProc()
    StatusBar_BringToFront(s.instance)
}

func (s *TStatusBar) HasParent() bool {
    defer exceptionProc()
    return StatusBar_HasParent(s.instance)
}

func (s *TStatusBar) Hide() {
    defer exceptionProc()
    StatusBar_Hide(s.instance)
}

func (s *TStatusBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return StatusBar_Perform(s.instance, Msg , WParam , LParam )
}

func (s *TStatusBar) Refresh() {
    defer exceptionProc()
    StatusBar_Refresh(s.instance)
}

func (s *TStatusBar) SendToBack() {
    defer exceptionProc()
    StatusBar_SendToBack(s.instance)
}

func (s *TStatusBar) Show() {
    defer exceptionProc()
    StatusBar_Show(s.instance)
}

func (s *TStatusBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return StatusBar_GetTextBuf(s.instance, Buffer , BufSize )
}

func (s *TStatusBar) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(StatusBar_FindComponent(s.instance, AName ))
}

func (s *TStatusBar) GetNamePath() string {
    defer exceptionProc()
    return StatusBar_GetNamePath(s.instance)
}

func (s *TStatusBar) Assign(Source IObject) {
    defer exceptionProc()
    StatusBar_Assign(s.instance, CheckPtr(Source))
}

func (s *TStatusBar) ClassName() string {
    defer exceptionProc()
    return StatusBar_ClassName(s.instance)
}

func (s *TStatusBar) Equals(Obj IObject) bool {
    defer exceptionProc()
    return StatusBar_Equals(s.instance, CheckPtr(Obj))
}

func (s *TStatusBar) GetHashCode() int32 {
    defer exceptionProc()
    return StatusBar_GetHashCode(s.instance)
}

func (s *TStatusBar) ToString() string {
    defer exceptionProc()
    return StatusBar_ToString(s.instance)
}

func (s *TStatusBar) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(StatusBar_GetAction(s.instance))
}

func (s *TStatusBar) SetAction(value IComponent) {
    defer exceptionProc()
    StatusBar_SetAction(s.instance, CheckPtr(value))
}

func (s *TStatusBar) AutoHint() bool {
    defer exceptionProc()
    return StatusBar_GetAutoHint(s.instance)
}

func (s *TStatusBar) SetAutoHint(value bool) {
    defer exceptionProc()
    StatusBar_SetAutoHint(s.instance, value)
}

func (s *TStatusBar) Align() TAlign {
    defer exceptionProc()
    return StatusBar_GetAlign(s.instance)
}

func (s *TStatusBar) SetAlign(value TAlign) {
    defer exceptionProc()
    StatusBar_SetAlign(s.instance, value)
}

func (s *TStatusBar) Anchors() TAnchors {
    defer exceptionProc()
    return StatusBar_GetAnchors(s.instance)
}

func (s *TStatusBar) SetAnchors(value TAnchors) {
    defer exceptionProc()
    StatusBar_SetAnchors(s.instance, value)
}

func (s *TStatusBar) BorderWidth() int32 {
    defer exceptionProc()
    return StatusBar_GetBorderWidth(s.instance)
}

func (s *TStatusBar) SetBorderWidth(value int32) {
    defer exceptionProc()
    StatusBar_SetBorderWidth(s.instance, value)
}

func (s *TStatusBar) Color() TColor {
    defer exceptionProc()
    return StatusBar_GetColor(s.instance)
}

func (s *TStatusBar) SetColor(value TColor) {
    defer exceptionProc()
    StatusBar_SetColor(s.instance, value)
}

func (s *TStatusBar) DoubleBuffered() bool {
    defer exceptionProc()
    return StatusBar_GetDoubleBuffered(s.instance)
}

func (s *TStatusBar) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    StatusBar_SetDoubleBuffered(s.instance, value)
}

func (s *TStatusBar) Enabled() bool {
    defer exceptionProc()
    return StatusBar_GetEnabled(s.instance)
}

func (s *TStatusBar) SetEnabled(value bool) {
    defer exceptionProc()
    StatusBar_SetEnabled(s.instance, value)
}

func (s *TStatusBar) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(StatusBar_GetFont(s.instance))
}

func (s *TStatusBar) SetFont(value *TFont) {
    defer exceptionProc()
    StatusBar_SetFont(s.instance, CheckPtr(value))
}

func (s *TStatusBar) Panels() *TStatusPanels {
    defer exceptionProc()
    return StatusPanelsFromInst(StatusBar_GetPanels(s.instance))
}

func (s *TStatusBar) SetPanels(value *TStatusPanels) {
    defer exceptionProc()
    StatusBar_SetPanels(s.instance, CheckPtr(value))
}

func (s *TStatusBar) ParentColor() bool {
    defer exceptionProc()
    return StatusBar_GetParentColor(s.instance)
}

func (s *TStatusBar) SetParentColor(value bool) {
    defer exceptionProc()
    StatusBar_SetParentColor(s.instance, value)
}

func (s *TStatusBar) ParentFont() bool {
    defer exceptionProc()
    return StatusBar_GetParentFont(s.instance)
}

func (s *TStatusBar) SetParentFont(value bool) {
    defer exceptionProc()
    StatusBar_SetParentFont(s.instance, value)
}

func (s *TStatusBar) ParentShowHint() bool {
    defer exceptionProc()
    return StatusBar_GetParentShowHint(s.instance)
}

func (s *TStatusBar) SetParentShowHint(value bool) {
    defer exceptionProc()
    StatusBar_SetParentShowHint(s.instance, value)
}

func (s *TStatusBar) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(StatusBar_GetPopupMenu(s.instance))
}

func (s *TStatusBar) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    StatusBar_SetPopupMenu(s.instance, CheckPtr(value))
}

func (s *TStatusBar) ShowHint() bool {
    defer exceptionProc()
    return StatusBar_GetShowHint(s.instance)
}

func (s *TStatusBar) SetShowHint(value bool) {
    defer exceptionProc()
    StatusBar_SetShowHint(s.instance, value)
}

func (s *TStatusBar) SimplePanel() bool {
    defer exceptionProc()
    return StatusBar_GetSimplePanel(s.instance)
}

func (s *TStatusBar) SetSimplePanel(value bool) {
    defer exceptionProc()
    StatusBar_SetSimplePanel(s.instance, value)
}

func (s *TStatusBar) SimpleText() string {
    defer exceptionProc()
    return StatusBar_GetSimpleText(s.instance)
}

func (s *TStatusBar) SetSimpleText(value string) {
    defer exceptionProc()
    StatusBar_SetSimpleText(s.instance, value)
}

func (s *TStatusBar) SizeGrip() bool {
    defer exceptionProc()
    return StatusBar_GetSizeGrip(s.instance)
}

func (s *TStatusBar) SetSizeGrip(value bool) {
    defer exceptionProc()
    StatusBar_SetSizeGrip(s.instance, value)
}

func (s *TStatusBar) UseSystemFont() bool {
    defer exceptionProc()
    return StatusBar_GetUseSystemFont(s.instance)
}

func (s *TStatusBar) SetUseSystemFont(value bool) {
    defer exceptionProc()
    StatusBar_SetUseSystemFont(s.instance, value)
}

func (s *TStatusBar) Visible() bool {
    defer exceptionProc()
    return StatusBar_GetVisible(s.instance)
}

func (s *TStatusBar) SetVisible(value bool) {
    defer exceptionProc()
    StatusBar_SetVisible(s.instance, value)
}

func (s *TStatusBar) SetOnClick(fn TNotifyEvent) {
    StatusBar_SetOnClick(s.instance, fn)
}

func (s *TStatusBar) SetOnDblClick(fn TNotifyEvent) {
    StatusBar_SetOnDblClick(s.instance, fn)
}

func (s *TStatusBar) SetOnMouseDown(fn TMouseEvent) {
    StatusBar_SetOnMouseDown(s.instance, fn)
}

func (s *TStatusBar) SetOnMouseEnter(fn TNotifyEvent) {
    StatusBar_SetOnMouseEnter(s.instance, fn)
}

func (s *TStatusBar) SetOnMouseLeave(fn TNotifyEvent) {
    StatusBar_SetOnMouseLeave(s.instance, fn)
}

func (s *TStatusBar) SetOnMouseMove(fn TMouseMoveEvent) {
    StatusBar_SetOnMouseMove(s.instance, fn)
}

func (s *TStatusBar) SetOnMouseUp(fn TMouseEvent) {
    StatusBar_SetOnMouseUp(s.instance, fn)
}

func (s *TStatusBar) SetOnResize(fn TNotifyEvent) {
    StatusBar_SetOnResize(s.instance, fn)
}

func (s *TStatusBar) Canvas() *TCanvas {
    defer exceptionProc()
    return CanvasFromInst(StatusBar_GetCanvas(s.instance))
}

func (s *TStatusBar) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(StatusBar_GetBrush(s.instance))
}

func (s *TStatusBar) ControlCount() int32 {
    defer exceptionProc()
    return StatusBar_GetControlCount(s.instance)
}

func (s *TStatusBar) Handle() HWND {
    defer exceptionProc()
    return StatusBar_GetHandle(s.instance)
}

func (s *TStatusBar) TabOrder() int16 {
    defer exceptionProc()
    return StatusBar_GetTabOrder(s.instance)
}

func (s *TStatusBar) SetTabOrder(value int16) {
    defer exceptionProc()
    StatusBar_SetTabOrder(s.instance, value)
}

func (s *TStatusBar) TabStop() bool {
    defer exceptionProc()
    return StatusBar_GetTabStop(s.instance)
}

func (s *TStatusBar) SetTabStop(value bool) {
    defer exceptionProc()
    StatusBar_SetTabStop(s.instance, value)
}

func (s *TStatusBar) BoundsRect() TRect {
    defer exceptionProc()
    return StatusBar_GetBoundsRect(s.instance)
}

func (s *TStatusBar) SetBoundsRect(value TRect) {
    defer exceptionProc()
    StatusBar_SetBoundsRect(s.instance, value)
}

func (s *TStatusBar) ClientHeight() int32 {
    defer exceptionProc()
    return StatusBar_GetClientHeight(s.instance)
}

func (s *TStatusBar) SetClientHeight(value int32) {
    defer exceptionProc()
    StatusBar_SetClientHeight(s.instance, value)
}

func (s *TStatusBar) ClientRect() TRect {
    defer exceptionProc()
    return StatusBar_GetClientRect(s.instance)
}

func (s *TStatusBar) ClientWidth() int32 {
    defer exceptionProc()
    return StatusBar_GetClientWidth(s.instance)
}

func (s *TStatusBar) SetClientWidth(value int32) {
    defer exceptionProc()
    StatusBar_SetClientWidth(s.instance, value)
}

func (s *TStatusBar) ExplicitLeft() int32 {
    defer exceptionProc()
    return StatusBar_GetExplicitLeft(s.instance)
}

func (s *TStatusBar) ExplicitTop() int32 {
    defer exceptionProc()
    return StatusBar_GetExplicitTop(s.instance)
}

func (s *TStatusBar) ExplicitWidth() int32 {
    defer exceptionProc()
    return StatusBar_GetExplicitWidth(s.instance)
}

func (s *TStatusBar) ExplicitHeight() int32 {
    defer exceptionProc()
    return StatusBar_GetExplicitHeight(s.instance)
}

func (s *TStatusBar) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(StatusBar_GetParent(s.instance))
}

func (s *TStatusBar) SetParent(value IControl) {
    defer exceptionProc()
    StatusBar_SetParent(s.instance, CheckPtr(value))
}

func (s *TStatusBar) AlignWithMargins() bool {
    defer exceptionProc()
    return StatusBar_GetAlignWithMargins(s.instance)
}

func (s *TStatusBar) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    StatusBar_SetAlignWithMargins(s.instance, value)
}

func (s *TStatusBar) Left() int32 {
    defer exceptionProc()
    return StatusBar_GetLeft(s.instance)
}

func (s *TStatusBar) SetLeft(value int32) {
    defer exceptionProc()
    StatusBar_SetLeft(s.instance, value)
}

func (s *TStatusBar) Top() int32 {
    defer exceptionProc()
    return StatusBar_GetTop(s.instance)
}

func (s *TStatusBar) SetTop(value int32) {
    defer exceptionProc()
    StatusBar_SetTop(s.instance, value)
}

func (s *TStatusBar) Width() int32 {
    defer exceptionProc()
    return StatusBar_GetWidth(s.instance)
}

func (s *TStatusBar) SetWidth(value int32) {
    defer exceptionProc()
    StatusBar_SetWidth(s.instance, value)
}

func (s *TStatusBar) Height() int32 {
    defer exceptionProc()
    return StatusBar_GetHeight(s.instance)
}

func (s *TStatusBar) SetHeight(value int32) {
    defer exceptionProc()
    StatusBar_SetHeight(s.instance, value)
}

func (s *TStatusBar) Cursor() TCursor {
    defer exceptionProc()
    return StatusBar_GetCursor(s.instance)
}

func (s *TStatusBar) SetCursor(value TCursor) {
    defer exceptionProc()
    StatusBar_SetCursor(s.instance, value)
}

func (s *TStatusBar) Hint() string {
    defer exceptionProc()
    return StatusBar_GetHint(s.instance)
}

func (s *TStatusBar) SetHint(value string) {
    defer exceptionProc()
    StatusBar_SetHint(s.instance, value)
}

func (s *TStatusBar) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(StatusBar_GetMargins(s.instance))
}

func (s *TStatusBar) SetMargins(value *TMargins) {
    defer exceptionProc()
    StatusBar_SetMargins(s.instance, CheckPtr(value))
}

func (s *TStatusBar) ComponentCount() int32 {
    defer exceptionProc()
    return StatusBar_GetComponentCount(s.instance)
}

func (s *TStatusBar) ComponentIndex() int32 {
    defer exceptionProc()
    return StatusBar_GetComponentIndex(s.instance)
}

func (s *TStatusBar) SetComponentIndex(value int32) {
    defer exceptionProc()
    StatusBar_SetComponentIndex(s.instance, value)
}

func (s *TStatusBar) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(StatusBar_GetOwner(s.instance))
}

func (s *TStatusBar) Name() string {
    defer exceptionProc()
    return StatusBar_GetName(s.instance)
}

func (s *TStatusBar) SetName(value string) {
    defer exceptionProc()
    StatusBar_SetName(s.instance, value)
}

func (s *TStatusBar) Tag() int {
    defer exceptionProc()
    return StatusBar_GetTag(s.instance)
}

func (s *TStatusBar) SetTag(value int) {
    defer exceptionProc()
    StatusBar_SetTag(s.instance, value)
}

func (s *TStatusBar) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(StatusBar_GetControls(s.instance, Index))
}

func (s *TStatusBar) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(StatusBar_GetComponents(s.instance, AIndex))
}

