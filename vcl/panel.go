
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TPanel struct {
    IControl
    instance uintptr
}

func NewPanel(owner IComponent) *TPanel {
    p := new(TPanel)
    p.instance = Panel_Create(owner.Instance())
    return p
}

func PanelFromInst(inst uintptr) *TPanel {
    p := new(TPanel)
    p.instance = inst
    return p
}

func PanelFromObj(obj IObject) *TPanel {
    p := new(TPanel)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TPanel) Free() {
    if p.instance != 0 {
        Panel_Free(p.instance)
    }
}

func (p *TPanel) Instance() uintptr {
    return p.instance
}

func (p *TPanel) IsValid() bool {
    return p.instance != 0
}

func (p *TPanel) CanFocus() bool {
    defer exceptionProc()
    return Panel_CanFocus(p.instance)
}

func (p *TPanel) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    Panel_FlipChildren(p.instance, AllLevels )
}

func (p *TPanel) Focused() bool {
    defer exceptionProc()
    return Panel_Focused(p.instance)
}

func (p *TPanel) HandleAllocated() bool {
    defer exceptionProc()
    return Panel_HandleAllocated(p.instance)
}

func (p *TPanel) Invalidate() {
    defer exceptionProc()
    Panel_Invalidate(p.instance)
}

func (p *TPanel) Realign() {
    defer exceptionProc()
    Panel_Realign(p.instance)
}

func (p *TPanel) Repaint() {
    defer exceptionProc()
    Panel_Repaint(p.instance)
}

func (p *TPanel) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    Panel_ScaleBy(p.instance, M , D )
}

func (p *TPanel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    Panel_SetBounds(p.instance, ALeft , ATop , AWidth , AHeight )
}

func (p *TPanel) SetFocus() {
    defer exceptionProc()
    Panel_SetFocus(p.instance)
}

func (p *TPanel) Update() {
    defer exceptionProc()
    Panel_Update(p.instance)
}

func (p *TPanel) BringToFront() {
    defer exceptionProc()
    Panel_BringToFront(p.instance)
}

func (p *TPanel) HasParent() bool {
    defer exceptionProc()
    return Panel_HasParent(p.instance)
}

func (p *TPanel) Hide() {
    defer exceptionProc()
    Panel_Hide(p.instance)
}

func (p *TPanel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return Panel_Perform(p.instance, Msg , WParam , LParam )
}

func (p *TPanel) Refresh() {
    defer exceptionProc()
    Panel_Refresh(p.instance)
}

func (p *TPanel) SendToBack() {
    defer exceptionProc()
    Panel_SendToBack(p.instance)
}

func (p *TPanel) Show() {
    defer exceptionProc()
    Panel_Show(p.instance)
}

func (p *TPanel) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return Panel_GetTextBuf(p.instance, Buffer , BufSize )
}

func (p *TPanel) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Panel_FindComponent(p.instance, AName ))
}

func (p *TPanel) GetNamePath() string {
    defer exceptionProc()
    return Panel_GetNamePath(p.instance)
}

func (p *TPanel) Assign(Source IObject) {
    defer exceptionProc()
    Panel_Assign(p.instance, CheckPtr(Source))
}

func (p *TPanel) ClassName() string {
    defer exceptionProc()
    return Panel_ClassName(p.instance)
}

func (p *TPanel) Equals(Obj IObject) bool {
    defer exceptionProc()
    return Panel_Equals(p.instance, CheckPtr(Obj))
}

func (p *TPanel) GetHashCode() int32 {
    defer exceptionProc()
    return Panel_GetHashCode(p.instance)
}

func (p *TPanel) ToString() string {
    defer exceptionProc()
    return Panel_ToString(p.instance)
}

func (p *TPanel) Align() TAlign {
    defer exceptionProc()
    return Panel_GetAlign(p.instance)
}

func (p *TPanel) SetAlign(value TAlign) {
    defer exceptionProc()
    Panel_SetAlign(p.instance, value)
}

func (p *TPanel) Alignment() TAlignment {
    defer exceptionProc()
    return Panel_GetAlignment(p.instance)
}

func (p *TPanel) SetAlignment(value TAlignment) {
    defer exceptionProc()
    Panel_SetAlignment(p.instance, value)
}

func (p *TPanel) Anchors() TAnchors {
    defer exceptionProc()
    return Panel_GetAnchors(p.instance)
}

func (p *TPanel) SetAnchors(value TAnchors) {
    defer exceptionProc()
    Panel_SetAnchors(p.instance, value)
}

func (p *TPanel) AutoSize() bool {
    defer exceptionProc()
    return Panel_GetAutoSize(p.instance)
}

func (p *TPanel) SetAutoSize(value bool) {
    defer exceptionProc()
    Panel_SetAutoSize(p.instance, value)
}

func (p *TPanel) BorderWidth() int32 {
    defer exceptionProc()
    return Panel_GetBorderWidth(p.instance)
}

func (p *TPanel) SetBorderWidth(value int32) {
    defer exceptionProc()
    Panel_SetBorderWidth(p.instance, value)
}

func (p *TPanel) BorderStyle() TBorderStyle {
    defer exceptionProc()
    return Panel_GetBorderStyle(p.instance)
}

func (p *TPanel) SetBorderStyle(value TBorderStyle) {
    defer exceptionProc()
    Panel_SetBorderStyle(p.instance, value)
}

func (p *TPanel) Caption() string {
    defer exceptionProc()
    return Panel_GetCaption(p.instance)
}

func (p *TPanel) SetCaption(value string) {
    defer exceptionProc()
    Panel_SetCaption(p.instance, value)
}

func (p *TPanel) Color() TColor {
    defer exceptionProc()
    return Panel_GetColor(p.instance)
}

func (p *TPanel) SetColor(value TColor) {
    defer exceptionProc()
    Panel_SetColor(p.instance, value)
}

func (p *TPanel) DoubleBuffered() bool {
    defer exceptionProc()
    return Panel_GetDoubleBuffered(p.instance)
}

func (p *TPanel) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    Panel_SetDoubleBuffered(p.instance, value)
}

func (p *TPanel) Enabled() bool {
    defer exceptionProc()
    return Panel_GetEnabled(p.instance)
}

func (p *TPanel) SetEnabled(value bool) {
    defer exceptionProc()
    Panel_SetEnabled(p.instance, value)
}

func (p *TPanel) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(Panel_GetFont(p.instance))
}

func (p *TPanel) SetFont(value *TFont) {
    defer exceptionProc()
    Panel_SetFont(p.instance, CheckPtr(value))
}

func (p *TPanel) ParentColor() bool {
    defer exceptionProc()
    return Panel_GetParentColor(p.instance)
}

func (p *TPanel) SetParentColor(value bool) {
    defer exceptionProc()
    Panel_SetParentColor(p.instance, value)
}

func (p *TPanel) ParentCtl3D() bool {
    defer exceptionProc()
    return Panel_GetParentCtl3D(p.instance)
}

func (p *TPanel) SetParentCtl3D(value bool) {
    defer exceptionProc()
    Panel_SetParentCtl3D(p.instance, value)
}

func (p *TPanel) ParentFont() bool {
    defer exceptionProc()
    return Panel_GetParentFont(p.instance)
}

func (p *TPanel) SetParentFont(value bool) {
    defer exceptionProc()
    Panel_SetParentFont(p.instance, value)
}

func (p *TPanel) ParentShowHint() bool {
    defer exceptionProc()
    return Panel_GetParentShowHint(p.instance)
}

func (p *TPanel) SetParentShowHint(value bool) {
    defer exceptionProc()
    Panel_SetParentShowHint(p.instance, value)
}

func (p *TPanel) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(Panel_GetPopupMenu(p.instance))
}

func (p *TPanel) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    Panel_SetPopupMenu(p.instance, CheckPtr(value))
}

func (p *TPanel) ShowCaption() bool {
    defer exceptionProc()
    return Panel_GetShowCaption(p.instance)
}

func (p *TPanel) SetShowCaption(value bool) {
    defer exceptionProc()
    Panel_SetShowCaption(p.instance, value)
}

func (p *TPanel) ShowHint() bool {
    defer exceptionProc()
    return Panel_GetShowHint(p.instance)
}

func (p *TPanel) SetShowHint(value bool) {
    defer exceptionProc()
    Panel_SetShowHint(p.instance, value)
}

func (p *TPanel) TabOrder() int16 {
    defer exceptionProc()
    return Panel_GetTabOrder(p.instance)
}

func (p *TPanel) SetTabOrder(value int16) {
    defer exceptionProc()
    Panel_SetTabOrder(p.instance, value)
}

func (p *TPanel) TabStop() bool {
    defer exceptionProc()
    return Panel_GetTabStop(p.instance)
}

func (p *TPanel) SetTabStop(value bool) {
    defer exceptionProc()
    Panel_SetTabStop(p.instance, value)
}

func (p *TPanel) Visible() bool {
    defer exceptionProc()
    return Panel_GetVisible(p.instance)
}

func (p *TPanel) SetVisible(value bool) {
    defer exceptionProc()
    Panel_SetVisible(p.instance, value)
}

func (p *TPanel) SetOnClick(fn TNotifyEvent) {
    Panel_SetOnClick(p.instance, fn)
}

func (p *TPanel) SetOnDblClick(fn TNotifyEvent) {
    Panel_SetOnDblClick(p.instance, fn)
}

func (p *TPanel) SetOnEnter(fn TNotifyEvent) {
    Panel_SetOnEnter(p.instance, fn)
}

func (p *TPanel) SetOnExit(fn TNotifyEvent) {
    Panel_SetOnExit(p.instance, fn)
}

func (p *TPanel) SetOnMouseDown(fn TMouseEvent) {
    Panel_SetOnMouseDown(p.instance, fn)
}

func (p *TPanel) SetOnMouseEnter(fn TNotifyEvent) {
    Panel_SetOnMouseEnter(p.instance, fn)
}

func (p *TPanel) SetOnMouseLeave(fn TNotifyEvent) {
    Panel_SetOnMouseLeave(p.instance, fn)
}

func (p *TPanel) SetOnMouseMove(fn TMouseMoveEvent) {
    Panel_SetOnMouseMove(p.instance, fn)
}

func (p *TPanel) SetOnMouseUp(fn TMouseEvent) {
    Panel_SetOnMouseUp(p.instance, fn)
}

func (p *TPanel) SetOnResize(fn TNotifyEvent) {
    Panel_SetOnResize(p.instance, fn)
}

func (p *TPanel) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(Panel_GetBrush(p.instance))
}

func (p *TPanel) ControlCount() int32 {
    defer exceptionProc()
    return Panel_GetControlCount(p.instance)
}

func (p *TPanel) Handle() HWND {
    defer exceptionProc()
    return Panel_GetHandle(p.instance)
}

func (p *TPanel) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(Panel_GetAction(p.instance))
}

func (p *TPanel) SetAction(value IComponent) {
    defer exceptionProc()
    Panel_SetAction(p.instance, CheckPtr(value))
}

func (p *TPanel) BoundsRect() TRect {
    defer exceptionProc()
    return Panel_GetBoundsRect(p.instance)
}

func (p *TPanel) SetBoundsRect(value TRect) {
    defer exceptionProc()
    Panel_SetBoundsRect(p.instance, value)
}

func (p *TPanel) ClientHeight() int32 {
    defer exceptionProc()
    return Panel_GetClientHeight(p.instance)
}

func (p *TPanel) SetClientHeight(value int32) {
    defer exceptionProc()
    Panel_SetClientHeight(p.instance, value)
}

func (p *TPanel) ClientRect() TRect {
    defer exceptionProc()
    return Panel_GetClientRect(p.instance)
}

func (p *TPanel) ClientWidth() int32 {
    defer exceptionProc()
    return Panel_GetClientWidth(p.instance)
}

func (p *TPanel) SetClientWidth(value int32) {
    defer exceptionProc()
    Panel_SetClientWidth(p.instance, value)
}

func (p *TPanel) ExplicitLeft() int32 {
    defer exceptionProc()
    return Panel_GetExplicitLeft(p.instance)
}

func (p *TPanel) ExplicitTop() int32 {
    defer exceptionProc()
    return Panel_GetExplicitTop(p.instance)
}

func (p *TPanel) ExplicitWidth() int32 {
    defer exceptionProc()
    return Panel_GetExplicitWidth(p.instance)
}

func (p *TPanel) ExplicitHeight() int32 {
    defer exceptionProc()
    return Panel_GetExplicitHeight(p.instance)
}

func (p *TPanel) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(Panel_GetParent(p.instance))
}

func (p *TPanel) SetParent(value IControl) {
    defer exceptionProc()
    Panel_SetParent(p.instance, CheckPtr(value))
}

func (p *TPanel) AlignWithMargins() bool {
    defer exceptionProc()
    return Panel_GetAlignWithMargins(p.instance)
}

func (p *TPanel) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    Panel_SetAlignWithMargins(p.instance, value)
}

func (p *TPanel) Left() int32 {
    defer exceptionProc()
    return Panel_GetLeft(p.instance)
}

func (p *TPanel) SetLeft(value int32) {
    defer exceptionProc()
    Panel_SetLeft(p.instance, value)
}

func (p *TPanel) Top() int32 {
    defer exceptionProc()
    return Panel_GetTop(p.instance)
}

func (p *TPanel) SetTop(value int32) {
    defer exceptionProc()
    Panel_SetTop(p.instance, value)
}

func (p *TPanel) Width() int32 {
    defer exceptionProc()
    return Panel_GetWidth(p.instance)
}

func (p *TPanel) SetWidth(value int32) {
    defer exceptionProc()
    Panel_SetWidth(p.instance, value)
}

func (p *TPanel) Height() int32 {
    defer exceptionProc()
    return Panel_GetHeight(p.instance)
}

func (p *TPanel) SetHeight(value int32) {
    defer exceptionProc()
    Panel_SetHeight(p.instance, value)
}

func (p *TPanel) Cursor() TCursor {
    defer exceptionProc()
    return Panel_GetCursor(p.instance)
}

func (p *TPanel) SetCursor(value TCursor) {
    defer exceptionProc()
    Panel_SetCursor(p.instance, value)
}

func (p *TPanel) Hint() string {
    defer exceptionProc()
    return Panel_GetHint(p.instance)
}

func (p *TPanel) SetHint(value string) {
    defer exceptionProc()
    Panel_SetHint(p.instance, value)
}

func (p *TPanel) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(Panel_GetMargins(p.instance))
}

func (p *TPanel) SetMargins(value *TMargins) {
    defer exceptionProc()
    Panel_SetMargins(p.instance, CheckPtr(value))
}

func (p *TPanel) ComponentCount() int32 {
    defer exceptionProc()
    return Panel_GetComponentCount(p.instance)
}

func (p *TPanel) ComponentIndex() int32 {
    defer exceptionProc()
    return Panel_GetComponentIndex(p.instance)
}

func (p *TPanel) SetComponentIndex(value int32) {
    defer exceptionProc()
    Panel_SetComponentIndex(p.instance, value)
}

func (p *TPanel) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Panel_GetOwner(p.instance))
}

func (p *TPanel) Name() string {
    defer exceptionProc()
    return Panel_GetName(p.instance)
}

func (p *TPanel) SetName(value string) {
    defer exceptionProc()
    Panel_SetName(p.instance, value)
}

func (p *TPanel) Tag() int {
    defer exceptionProc()
    return Panel_GetTag(p.instance)
}

func (p *TPanel) SetTag(value int) {
    defer exceptionProc()
    Panel_SetTag(p.instance, value)
}

func (p *TPanel) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(Panel_GetControls(p.instance, Index))
}

func (p *TPanel) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Panel_GetComponents(p.instance, AIndex))
}

