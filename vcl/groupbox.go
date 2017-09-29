
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TGroupBox struct {
    IControl
    instance uintptr
}

func NewGroupBox(owner IComponent) *TGroupBox {
    g := new(TGroupBox)
    g.instance = GroupBox_Create(owner.Instance())
    return g
}

func GroupBoxFromInst(inst uintptr) *TGroupBox {
    g := new(TGroupBox)
    g.instance = inst
    return g
}

func GroupBoxFromObj(obj IObject) *TGroupBox {
    g := new(TGroupBox)
    g.instance = CheckPtr(obj)
    return g
}

func (g *TGroupBox) Free() {
    if g.instance != 0 {
        GroupBox_Free(g.instance)
    }
}

func (g *TGroupBox) Instance() uintptr {
    return g.instance
}

func (g *TGroupBox) IsValid() bool {
    return g.instance != 0
}

func (g *TGroupBox) CanFocus() bool {
    defer exceptionProc()
    return GroupBox_CanFocus(g.instance)
}

func (g *TGroupBox) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    GroupBox_FlipChildren(g.instance, AllLevels )
}

func (g *TGroupBox) Focused() bool {
    defer exceptionProc()
    return GroupBox_Focused(g.instance)
}

func (g *TGroupBox) HandleAllocated() bool {
    defer exceptionProc()
    return GroupBox_HandleAllocated(g.instance)
}

func (g *TGroupBox) Invalidate() {
    defer exceptionProc()
    GroupBox_Invalidate(g.instance)
}

func (g *TGroupBox) Realign() {
    defer exceptionProc()
    GroupBox_Realign(g.instance)
}

func (g *TGroupBox) Repaint() {
    defer exceptionProc()
    GroupBox_Repaint(g.instance)
}

func (g *TGroupBox) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    GroupBox_ScaleBy(g.instance, M , D )
}

func (g *TGroupBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    GroupBox_SetBounds(g.instance, ALeft , ATop , AWidth , AHeight )
}

func (g *TGroupBox) SetFocus() {
    defer exceptionProc()
    GroupBox_SetFocus(g.instance)
}

func (g *TGroupBox) Update() {
    defer exceptionProc()
    GroupBox_Update(g.instance)
}

func (g *TGroupBox) BringToFront() {
    defer exceptionProc()
    GroupBox_BringToFront(g.instance)
}

func (g *TGroupBox) HasParent() bool {
    defer exceptionProc()
    return GroupBox_HasParent(g.instance)
}

func (g *TGroupBox) Hide() {
    defer exceptionProc()
    GroupBox_Hide(g.instance)
}

func (g *TGroupBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return GroupBox_Perform(g.instance, Msg , WParam , LParam )
}

func (g *TGroupBox) Refresh() {
    defer exceptionProc()
    GroupBox_Refresh(g.instance)
}

func (g *TGroupBox) SendToBack() {
    defer exceptionProc()
    GroupBox_SendToBack(g.instance)
}

func (g *TGroupBox) Show() {
    defer exceptionProc()
    GroupBox_Show(g.instance)
}

func (g *TGroupBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return GroupBox_GetTextBuf(g.instance, Buffer , BufSize )
}

func (g *TGroupBox) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(GroupBox_FindComponent(g.instance, AName ))
}

func (g *TGroupBox) GetNamePath() string {
    defer exceptionProc()
    return GroupBox_GetNamePath(g.instance)
}

func (g *TGroupBox) Assign(Source IObject) {
    defer exceptionProc()
    GroupBox_Assign(g.instance, CheckPtr(Source))
}

func (g *TGroupBox) ClassName() string {
    defer exceptionProc()
    return GroupBox_ClassName(g.instance)
}

func (g *TGroupBox) Equals(Obj IObject) bool {
    defer exceptionProc()
    return GroupBox_Equals(g.instance, CheckPtr(Obj))
}

func (g *TGroupBox) GetHashCode() int32 {
    defer exceptionProc()
    return GroupBox_GetHashCode(g.instance)
}

func (g *TGroupBox) ToString() string {
    defer exceptionProc()
    return GroupBox_ToString(g.instance)
}

func (g *TGroupBox) Align() TAlign {
    defer exceptionProc()
    return GroupBox_GetAlign(g.instance)
}

func (g *TGroupBox) SetAlign(value TAlign) {
    defer exceptionProc()
    GroupBox_SetAlign(g.instance, value)
}

func (g *TGroupBox) Anchors() TAnchors {
    defer exceptionProc()
    return GroupBox_GetAnchors(g.instance)
}

func (g *TGroupBox) SetAnchors(value TAnchors) {
    defer exceptionProc()
    GroupBox_SetAnchors(g.instance, value)
}

func (g *TGroupBox) Caption() string {
    defer exceptionProc()
    return GroupBox_GetCaption(g.instance)
}

func (g *TGroupBox) SetCaption(value string) {
    defer exceptionProc()
    GroupBox_SetCaption(g.instance, value)
}

func (g *TGroupBox) Color() TColor {
    defer exceptionProc()
    return GroupBox_GetColor(g.instance)
}

func (g *TGroupBox) SetColor(value TColor) {
    defer exceptionProc()
    GroupBox_SetColor(g.instance, value)
}

func (g *TGroupBox) DoubleBuffered() bool {
    defer exceptionProc()
    return GroupBox_GetDoubleBuffered(g.instance)
}

func (g *TGroupBox) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    GroupBox_SetDoubleBuffered(g.instance, value)
}

func (g *TGroupBox) Enabled() bool {
    defer exceptionProc()
    return GroupBox_GetEnabled(g.instance)
}

func (g *TGroupBox) SetEnabled(value bool) {
    defer exceptionProc()
    GroupBox_SetEnabled(g.instance, value)
}

func (g *TGroupBox) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(GroupBox_GetFont(g.instance))
}

func (g *TGroupBox) SetFont(value *TFont) {
    defer exceptionProc()
    GroupBox_SetFont(g.instance, CheckPtr(value))
}

func (g *TGroupBox) ParentColor() bool {
    defer exceptionProc()
    return GroupBox_GetParentColor(g.instance)
}

func (g *TGroupBox) SetParentColor(value bool) {
    defer exceptionProc()
    GroupBox_SetParentColor(g.instance, value)
}

func (g *TGroupBox) ParentCtl3D() bool {
    defer exceptionProc()
    return GroupBox_GetParentCtl3D(g.instance)
}

func (g *TGroupBox) SetParentCtl3D(value bool) {
    defer exceptionProc()
    GroupBox_SetParentCtl3D(g.instance, value)
}

func (g *TGroupBox) ParentFont() bool {
    defer exceptionProc()
    return GroupBox_GetParentFont(g.instance)
}

func (g *TGroupBox) SetParentFont(value bool) {
    defer exceptionProc()
    GroupBox_SetParentFont(g.instance, value)
}

func (g *TGroupBox) ParentShowHint() bool {
    defer exceptionProc()
    return GroupBox_GetParentShowHint(g.instance)
}

func (g *TGroupBox) SetParentShowHint(value bool) {
    defer exceptionProc()
    GroupBox_SetParentShowHint(g.instance, value)
}

func (g *TGroupBox) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(GroupBox_GetPopupMenu(g.instance))
}

func (g *TGroupBox) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    GroupBox_SetPopupMenu(g.instance, CheckPtr(value))
}

func (g *TGroupBox) ShowHint() bool {
    defer exceptionProc()
    return GroupBox_GetShowHint(g.instance)
}

func (g *TGroupBox) SetShowHint(value bool) {
    defer exceptionProc()
    GroupBox_SetShowHint(g.instance, value)
}

func (g *TGroupBox) TabOrder() int16 {
    defer exceptionProc()
    return GroupBox_GetTabOrder(g.instance)
}

func (g *TGroupBox) SetTabOrder(value int16) {
    defer exceptionProc()
    GroupBox_SetTabOrder(g.instance, value)
}

func (g *TGroupBox) TabStop() bool {
    defer exceptionProc()
    return GroupBox_GetTabStop(g.instance)
}

func (g *TGroupBox) SetTabStop(value bool) {
    defer exceptionProc()
    GroupBox_SetTabStop(g.instance, value)
}

func (g *TGroupBox) Visible() bool {
    defer exceptionProc()
    return GroupBox_GetVisible(g.instance)
}

func (g *TGroupBox) SetVisible(value bool) {
    defer exceptionProc()
    GroupBox_SetVisible(g.instance, value)
}

func (g *TGroupBox) SetOnClick(fn TNotifyEvent) {
    GroupBox_SetOnClick(g.instance, fn)
}

func (g *TGroupBox) SetOnDblClick(fn TNotifyEvent) {
    GroupBox_SetOnDblClick(g.instance, fn)
}

func (g *TGroupBox) SetOnEnter(fn TNotifyEvent) {
    GroupBox_SetOnEnter(g.instance, fn)
}

func (g *TGroupBox) SetOnExit(fn TNotifyEvent) {
    GroupBox_SetOnExit(g.instance, fn)
}

func (g *TGroupBox) SetOnMouseDown(fn TMouseEvent) {
    GroupBox_SetOnMouseDown(g.instance, fn)
}

func (g *TGroupBox) SetOnMouseEnter(fn TNotifyEvent) {
    GroupBox_SetOnMouseEnter(g.instance, fn)
}

func (g *TGroupBox) SetOnMouseLeave(fn TNotifyEvent) {
    GroupBox_SetOnMouseLeave(g.instance, fn)
}

func (g *TGroupBox) SetOnMouseMove(fn TMouseMoveEvent) {
    GroupBox_SetOnMouseMove(g.instance, fn)
}

func (g *TGroupBox) SetOnMouseUp(fn TMouseEvent) {
    GroupBox_SetOnMouseUp(g.instance, fn)
}

func (g *TGroupBox) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(GroupBox_GetBrush(g.instance))
}

func (g *TGroupBox) ControlCount() int32 {
    defer exceptionProc()
    return GroupBox_GetControlCount(g.instance)
}

func (g *TGroupBox) Handle() HWND {
    defer exceptionProc()
    return GroupBox_GetHandle(g.instance)
}

func (g *TGroupBox) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(GroupBox_GetAction(g.instance))
}

func (g *TGroupBox) SetAction(value IComponent) {
    defer exceptionProc()
    GroupBox_SetAction(g.instance, CheckPtr(value))
}

func (g *TGroupBox) BoundsRect() TRect {
    defer exceptionProc()
    return GroupBox_GetBoundsRect(g.instance)
}

func (g *TGroupBox) SetBoundsRect(value TRect) {
    defer exceptionProc()
    GroupBox_SetBoundsRect(g.instance, value)
}

func (g *TGroupBox) ClientHeight() int32 {
    defer exceptionProc()
    return GroupBox_GetClientHeight(g.instance)
}

func (g *TGroupBox) SetClientHeight(value int32) {
    defer exceptionProc()
    GroupBox_SetClientHeight(g.instance, value)
}

func (g *TGroupBox) ClientRect() TRect {
    defer exceptionProc()
    return GroupBox_GetClientRect(g.instance)
}

func (g *TGroupBox) ClientWidth() int32 {
    defer exceptionProc()
    return GroupBox_GetClientWidth(g.instance)
}

func (g *TGroupBox) SetClientWidth(value int32) {
    defer exceptionProc()
    GroupBox_SetClientWidth(g.instance, value)
}

func (g *TGroupBox) ExplicitLeft() int32 {
    defer exceptionProc()
    return GroupBox_GetExplicitLeft(g.instance)
}

func (g *TGroupBox) ExplicitTop() int32 {
    defer exceptionProc()
    return GroupBox_GetExplicitTop(g.instance)
}

func (g *TGroupBox) ExplicitWidth() int32 {
    defer exceptionProc()
    return GroupBox_GetExplicitWidth(g.instance)
}

func (g *TGroupBox) ExplicitHeight() int32 {
    defer exceptionProc()
    return GroupBox_GetExplicitHeight(g.instance)
}

func (g *TGroupBox) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(GroupBox_GetParent(g.instance))
}

func (g *TGroupBox) SetParent(value IControl) {
    defer exceptionProc()
    GroupBox_SetParent(g.instance, CheckPtr(value))
}

func (g *TGroupBox) AlignWithMargins() bool {
    defer exceptionProc()
    return GroupBox_GetAlignWithMargins(g.instance)
}

func (g *TGroupBox) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    GroupBox_SetAlignWithMargins(g.instance, value)
}

func (g *TGroupBox) Left() int32 {
    defer exceptionProc()
    return GroupBox_GetLeft(g.instance)
}

func (g *TGroupBox) SetLeft(value int32) {
    defer exceptionProc()
    GroupBox_SetLeft(g.instance, value)
}

func (g *TGroupBox) Top() int32 {
    defer exceptionProc()
    return GroupBox_GetTop(g.instance)
}

func (g *TGroupBox) SetTop(value int32) {
    defer exceptionProc()
    GroupBox_SetTop(g.instance, value)
}

func (g *TGroupBox) Width() int32 {
    defer exceptionProc()
    return GroupBox_GetWidth(g.instance)
}

func (g *TGroupBox) SetWidth(value int32) {
    defer exceptionProc()
    GroupBox_SetWidth(g.instance, value)
}

func (g *TGroupBox) Height() int32 {
    defer exceptionProc()
    return GroupBox_GetHeight(g.instance)
}

func (g *TGroupBox) SetHeight(value int32) {
    defer exceptionProc()
    GroupBox_SetHeight(g.instance, value)
}

func (g *TGroupBox) Cursor() TCursor {
    defer exceptionProc()
    return GroupBox_GetCursor(g.instance)
}

func (g *TGroupBox) SetCursor(value TCursor) {
    defer exceptionProc()
    GroupBox_SetCursor(g.instance, value)
}

func (g *TGroupBox) Hint() string {
    defer exceptionProc()
    return GroupBox_GetHint(g.instance)
}

func (g *TGroupBox) SetHint(value string) {
    defer exceptionProc()
    GroupBox_SetHint(g.instance, value)
}

func (g *TGroupBox) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(GroupBox_GetMargins(g.instance))
}

func (g *TGroupBox) SetMargins(value *TMargins) {
    defer exceptionProc()
    GroupBox_SetMargins(g.instance, CheckPtr(value))
}

func (g *TGroupBox) ComponentCount() int32 {
    defer exceptionProc()
    return GroupBox_GetComponentCount(g.instance)
}

func (g *TGroupBox) ComponentIndex() int32 {
    defer exceptionProc()
    return GroupBox_GetComponentIndex(g.instance)
}

func (g *TGroupBox) SetComponentIndex(value int32) {
    defer exceptionProc()
    GroupBox_SetComponentIndex(g.instance, value)
}

func (g *TGroupBox) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(GroupBox_GetOwner(g.instance))
}

func (g *TGroupBox) Name() string {
    defer exceptionProc()
    return GroupBox_GetName(g.instance)
}

func (g *TGroupBox) SetName(value string) {
    defer exceptionProc()
    GroupBox_SetName(g.instance, value)
}

func (g *TGroupBox) Tag() int {
    defer exceptionProc()
    return GroupBox_GetTag(g.instance)
}

func (g *TGroupBox) SetTag(value int) {
    defer exceptionProc()
    GroupBox_SetTag(g.instance, value)
}

func (g *TGroupBox) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(GroupBox_GetControls(g.instance, Index))
}

func (g *TGroupBox) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(GroupBox_GetComponents(g.instance, AIndex))
}

