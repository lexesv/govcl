
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TControl struct {
    IControl
    instance uintptr
}

func NewControl(owner IComponent) *TControl {
    c := new(TControl)
    c.instance = Control_Create(owner.Instance())
    return c
}

func ControlFromInst(inst uintptr) *TControl {
    c := new(TControl)
    c.instance = inst
    return c
}

func ControlFromObj(obj IObject) *TControl {
    c := new(TControl)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TControl) Free() {
    if c.instance != 0 {
        Control_Free(c.instance)
    }
}

func (c *TControl) Instance() uintptr {
    return c.instance
}

func (c *TControl) IsValid() bool {
    return c.instance != 0
}

func (c *TControl) BringToFront() {
    defer exceptionProc()
    Control_BringToFront(c.instance)
}

func (c *TControl) HasParent() bool {
    defer exceptionProc()
    return Control_HasParent(c.instance)
}

func (c *TControl) Hide() {
    defer exceptionProc()
    Control_Hide(c.instance)
}

func (c *TControl) Invalidate() {
    defer exceptionProc()
    Control_Invalidate(c.instance)
}

func (c *TControl) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return Control_Perform(c.instance, Msg , WParam , LParam )
}

func (c *TControl) Refresh() {
    defer exceptionProc()
    Control_Refresh(c.instance)
}

func (c *TControl) Repaint() {
    defer exceptionProc()
    Control_Repaint(c.instance)
}

func (c *TControl) SendToBack() {
    defer exceptionProc()
    Control_SendToBack(c.instance)
}

func (c *TControl) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    Control_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight )
}

func (c *TControl) Show() {
    defer exceptionProc()
    Control_Show(c.instance)
}

func (c *TControl) Update() {
    defer exceptionProc()
    Control_Update(c.instance)
}

func (c *TControl) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return Control_GetTextBuf(c.instance, Buffer , BufSize )
}

func (c *TControl) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Control_FindComponent(c.instance, AName ))
}

func (c *TControl) GetNamePath() string {
    defer exceptionProc()
    return Control_GetNamePath(c.instance)
}

func (c *TControl) Assign(Source IObject) {
    defer exceptionProc()
    Control_Assign(c.instance, CheckPtr(Source))
}

func (c *TControl) ClassName() string {
    defer exceptionProc()
    return Control_ClassName(c.instance)
}

func (c *TControl) Equals(Obj IObject) bool {
    defer exceptionProc()
    return Control_Equals(c.instance, CheckPtr(Obj))
}

func (c *TControl) GetHashCode() int32 {
    defer exceptionProc()
    return Control_GetHashCode(c.instance)
}

func (c *TControl) ToString() string {
    defer exceptionProc()
    return Control_ToString(c.instance)
}

func (c *TControl) Enabled() bool {
    defer exceptionProc()
    return Control_GetEnabled(c.instance)
}

func (c *TControl) SetEnabled(value bool) {
    defer exceptionProc()
    Control_SetEnabled(c.instance, value)
}

func (c *TControl) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(Control_GetAction(c.instance))
}

func (c *TControl) SetAction(value IComponent) {
    defer exceptionProc()
    Control_SetAction(c.instance, CheckPtr(value))
}

func (c *TControl) Align() TAlign {
    defer exceptionProc()
    return Control_GetAlign(c.instance)
}

func (c *TControl) SetAlign(value TAlign) {
    defer exceptionProc()
    Control_SetAlign(c.instance, value)
}

func (c *TControl) Anchors() TAnchors {
    defer exceptionProc()
    return Control_GetAnchors(c.instance)
}

func (c *TControl) SetAnchors(value TAnchors) {
    defer exceptionProc()
    Control_SetAnchors(c.instance, value)
}

func (c *TControl) BoundsRect() TRect {
    defer exceptionProc()
    return Control_GetBoundsRect(c.instance)
}

func (c *TControl) SetBoundsRect(value TRect) {
    defer exceptionProc()
    Control_SetBoundsRect(c.instance, value)
}

func (c *TControl) ClientHeight() int32 {
    defer exceptionProc()
    return Control_GetClientHeight(c.instance)
}

func (c *TControl) SetClientHeight(value int32) {
    defer exceptionProc()
    Control_SetClientHeight(c.instance, value)
}

func (c *TControl) ClientRect() TRect {
    defer exceptionProc()
    return Control_GetClientRect(c.instance)
}

func (c *TControl) ClientWidth() int32 {
    defer exceptionProc()
    return Control_GetClientWidth(c.instance)
}

func (c *TControl) SetClientWidth(value int32) {
    defer exceptionProc()
    Control_SetClientWidth(c.instance, value)
}

func (c *TControl) ExplicitLeft() int32 {
    defer exceptionProc()
    return Control_GetExplicitLeft(c.instance)
}

func (c *TControl) ExplicitTop() int32 {
    defer exceptionProc()
    return Control_GetExplicitTop(c.instance)
}

func (c *TControl) ExplicitWidth() int32 {
    defer exceptionProc()
    return Control_GetExplicitWidth(c.instance)
}

func (c *TControl) ExplicitHeight() int32 {
    defer exceptionProc()
    return Control_GetExplicitHeight(c.instance)
}

func (c *TControl) ShowHint() bool {
    defer exceptionProc()
    return Control_GetShowHint(c.instance)
}

func (c *TControl) SetShowHint(value bool) {
    defer exceptionProc()
    Control_SetShowHint(c.instance, value)
}

func (c *TControl) Visible() bool {
    defer exceptionProc()
    return Control_GetVisible(c.instance)
}

func (c *TControl) SetVisible(value bool) {
    defer exceptionProc()
    Control_SetVisible(c.instance, value)
}

func (c *TControl) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(Control_GetParent(c.instance))
}

func (c *TControl) SetParent(value IControl) {
    defer exceptionProc()
    Control_SetParent(c.instance, CheckPtr(value))
}

func (c *TControl) AlignWithMargins() bool {
    defer exceptionProc()
    return Control_GetAlignWithMargins(c.instance)
}

func (c *TControl) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    Control_SetAlignWithMargins(c.instance, value)
}

func (c *TControl) Left() int32 {
    defer exceptionProc()
    return Control_GetLeft(c.instance)
}

func (c *TControl) SetLeft(value int32) {
    defer exceptionProc()
    Control_SetLeft(c.instance, value)
}

func (c *TControl) Top() int32 {
    defer exceptionProc()
    return Control_GetTop(c.instance)
}

func (c *TControl) SetTop(value int32) {
    defer exceptionProc()
    Control_SetTop(c.instance, value)
}

func (c *TControl) Width() int32 {
    defer exceptionProc()
    return Control_GetWidth(c.instance)
}

func (c *TControl) SetWidth(value int32) {
    defer exceptionProc()
    Control_SetWidth(c.instance, value)
}

func (c *TControl) Height() int32 {
    defer exceptionProc()
    return Control_GetHeight(c.instance)
}

func (c *TControl) SetHeight(value int32) {
    defer exceptionProc()
    Control_SetHeight(c.instance, value)
}

func (c *TControl) Cursor() TCursor {
    defer exceptionProc()
    return Control_GetCursor(c.instance)
}

func (c *TControl) SetCursor(value TCursor) {
    defer exceptionProc()
    Control_SetCursor(c.instance, value)
}

func (c *TControl) Hint() string {
    defer exceptionProc()
    return Control_GetHint(c.instance)
}

func (c *TControl) SetHint(value string) {
    defer exceptionProc()
    Control_SetHint(c.instance, value)
}

func (c *TControl) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(Control_GetMargins(c.instance))
}

func (c *TControl) SetMargins(value *TMargins) {
    defer exceptionProc()
    Control_SetMargins(c.instance, CheckPtr(value))
}

func (c *TControl) ComponentCount() int32 {
    defer exceptionProc()
    return Control_GetComponentCount(c.instance)
}

func (c *TControl) ComponentIndex() int32 {
    defer exceptionProc()
    return Control_GetComponentIndex(c.instance)
}

func (c *TControl) SetComponentIndex(value int32) {
    defer exceptionProc()
    Control_SetComponentIndex(c.instance, value)
}

func (c *TControl) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Control_GetOwner(c.instance))
}

func (c *TControl) Name() string {
    defer exceptionProc()
    return Control_GetName(c.instance)
}

func (c *TControl) SetName(value string) {
    defer exceptionProc()
    Control_SetName(c.instance, value)
}

func (c *TControl) Tag() int {
    defer exceptionProc()
    return Control_GetTag(c.instance)
}

func (c *TControl) SetTag(value int) {
    defer exceptionProc()
    Control_SetTag(c.instance, value)
}

func (c *TControl) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Control_GetComponents(c.instance, AIndex))
}

