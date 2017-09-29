
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TPaintBox struct {
    IControl
    instance uintptr
}

func NewPaintBox(owner IComponent) *TPaintBox {
    p := new(TPaintBox)
    p.instance = PaintBox_Create(owner.Instance())
    return p
}

func PaintBoxFromInst(inst uintptr) *TPaintBox {
    p := new(TPaintBox)
    p.instance = inst
    return p
}

func PaintBoxFromObj(obj IObject) *TPaintBox {
    p := new(TPaintBox)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TPaintBox) Free() {
    if p.instance != 0 {
        PaintBox_Free(p.instance)
    }
}

func (p *TPaintBox) Instance() uintptr {
    return p.instance
}

func (p *TPaintBox) IsValid() bool {
    return p.instance != 0
}

func (p *TPaintBox) BringToFront() {
    defer exceptionProc()
    PaintBox_BringToFront(p.instance)
}

func (p *TPaintBox) HasParent() bool {
    defer exceptionProc()
    return PaintBox_HasParent(p.instance)
}

func (p *TPaintBox) Hide() {
    defer exceptionProc()
    PaintBox_Hide(p.instance)
}

func (p *TPaintBox) Invalidate() {
    defer exceptionProc()
    PaintBox_Invalidate(p.instance)
}

func (p *TPaintBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return PaintBox_Perform(p.instance, Msg , WParam , LParam )
}

func (p *TPaintBox) Refresh() {
    defer exceptionProc()
    PaintBox_Refresh(p.instance)
}

func (p *TPaintBox) Repaint() {
    defer exceptionProc()
    PaintBox_Repaint(p.instance)
}

func (p *TPaintBox) SendToBack() {
    defer exceptionProc()
    PaintBox_SendToBack(p.instance)
}

func (p *TPaintBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    PaintBox_SetBounds(p.instance, ALeft , ATop , AWidth , AHeight )
}

func (p *TPaintBox) Show() {
    defer exceptionProc()
    PaintBox_Show(p.instance)
}

func (p *TPaintBox) Update() {
    defer exceptionProc()
    PaintBox_Update(p.instance)
}

func (p *TPaintBox) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return PaintBox_GetTextBuf(p.instance, Buffer , BufSize )
}

func (p *TPaintBox) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(PaintBox_FindComponent(p.instance, AName ))
}

func (p *TPaintBox) GetNamePath() string {
    defer exceptionProc()
    return PaintBox_GetNamePath(p.instance)
}

func (p *TPaintBox) Assign(Source IObject) {
    defer exceptionProc()
    PaintBox_Assign(p.instance, CheckPtr(Source))
}

func (p *TPaintBox) ClassName() string {
    defer exceptionProc()
    return PaintBox_ClassName(p.instance)
}

func (p *TPaintBox) Equals(Obj IObject) bool {
    defer exceptionProc()
    return PaintBox_Equals(p.instance, CheckPtr(Obj))
}

func (p *TPaintBox) GetHashCode() int32 {
    defer exceptionProc()
    return PaintBox_GetHashCode(p.instance)
}

func (p *TPaintBox) ToString() string {
    defer exceptionProc()
    return PaintBox_ToString(p.instance)
}

func (p *TPaintBox) Canvas() *TCanvas {
    defer exceptionProc()
    return CanvasFromInst(PaintBox_GetCanvas(p.instance))
}

func (p *TPaintBox) Align() TAlign {
    defer exceptionProc()
    return PaintBox_GetAlign(p.instance)
}

func (p *TPaintBox) SetAlign(value TAlign) {
    defer exceptionProc()
    PaintBox_SetAlign(p.instance, value)
}

func (p *TPaintBox) Anchors() TAnchors {
    defer exceptionProc()
    return PaintBox_GetAnchors(p.instance)
}

func (p *TPaintBox) SetAnchors(value TAnchors) {
    defer exceptionProc()
    PaintBox_SetAnchors(p.instance, value)
}

func (p *TPaintBox) Color() TColor {
    defer exceptionProc()
    return PaintBox_GetColor(p.instance)
}

func (p *TPaintBox) SetColor(value TColor) {
    defer exceptionProc()
    PaintBox_SetColor(p.instance, value)
}

func (p *TPaintBox) Enabled() bool {
    defer exceptionProc()
    return PaintBox_GetEnabled(p.instance)
}

func (p *TPaintBox) SetEnabled(value bool) {
    defer exceptionProc()
    PaintBox_SetEnabled(p.instance, value)
}

func (p *TPaintBox) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(PaintBox_GetFont(p.instance))
}

func (p *TPaintBox) SetFont(value *TFont) {
    defer exceptionProc()
    PaintBox_SetFont(p.instance, CheckPtr(value))
}

func (p *TPaintBox) ParentColor() bool {
    defer exceptionProc()
    return PaintBox_GetParentColor(p.instance)
}

func (p *TPaintBox) SetParentColor(value bool) {
    defer exceptionProc()
    PaintBox_SetParentColor(p.instance, value)
}

func (p *TPaintBox) ParentFont() bool {
    defer exceptionProc()
    return PaintBox_GetParentFont(p.instance)
}

func (p *TPaintBox) SetParentFont(value bool) {
    defer exceptionProc()
    PaintBox_SetParentFont(p.instance, value)
}

func (p *TPaintBox) ParentShowHint() bool {
    defer exceptionProc()
    return PaintBox_GetParentShowHint(p.instance)
}

func (p *TPaintBox) SetParentShowHint(value bool) {
    defer exceptionProc()
    PaintBox_SetParentShowHint(p.instance, value)
}

func (p *TPaintBox) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(PaintBox_GetPopupMenu(p.instance))
}

func (p *TPaintBox) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    PaintBox_SetPopupMenu(p.instance, CheckPtr(value))
}

func (p *TPaintBox) ShowHint() bool {
    defer exceptionProc()
    return PaintBox_GetShowHint(p.instance)
}

func (p *TPaintBox) SetShowHint(value bool) {
    defer exceptionProc()
    PaintBox_SetShowHint(p.instance, value)
}

func (p *TPaintBox) Visible() bool {
    defer exceptionProc()
    return PaintBox_GetVisible(p.instance)
}

func (p *TPaintBox) SetVisible(value bool) {
    defer exceptionProc()
    PaintBox_SetVisible(p.instance, value)
}

func (p *TPaintBox) SetOnClick(fn TNotifyEvent) {
    PaintBox_SetOnClick(p.instance, fn)
}

func (p *TPaintBox) SetOnDblClick(fn TNotifyEvent) {
    PaintBox_SetOnDblClick(p.instance, fn)
}

func (p *TPaintBox) SetOnMouseDown(fn TMouseEvent) {
    PaintBox_SetOnMouseDown(p.instance, fn)
}

func (p *TPaintBox) SetOnMouseEnter(fn TNotifyEvent) {
    PaintBox_SetOnMouseEnter(p.instance, fn)
}

func (p *TPaintBox) SetOnMouseLeave(fn TNotifyEvent) {
    PaintBox_SetOnMouseLeave(p.instance, fn)
}

func (p *TPaintBox) SetOnMouseMove(fn TMouseMoveEvent) {
    PaintBox_SetOnMouseMove(p.instance, fn)
}

func (p *TPaintBox) SetOnMouseUp(fn TMouseEvent) {
    PaintBox_SetOnMouseUp(p.instance, fn)
}

func (p *TPaintBox) SetOnPaint(fn TNotifyEvent) {
    PaintBox_SetOnPaint(p.instance, fn)
}

func (p *TPaintBox) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(PaintBox_GetAction(p.instance))
}

func (p *TPaintBox) SetAction(value IComponent) {
    defer exceptionProc()
    PaintBox_SetAction(p.instance, CheckPtr(value))
}

func (p *TPaintBox) BoundsRect() TRect {
    defer exceptionProc()
    return PaintBox_GetBoundsRect(p.instance)
}

func (p *TPaintBox) SetBoundsRect(value TRect) {
    defer exceptionProc()
    PaintBox_SetBoundsRect(p.instance, value)
}

func (p *TPaintBox) ClientHeight() int32 {
    defer exceptionProc()
    return PaintBox_GetClientHeight(p.instance)
}

func (p *TPaintBox) SetClientHeight(value int32) {
    defer exceptionProc()
    PaintBox_SetClientHeight(p.instance, value)
}

func (p *TPaintBox) ClientRect() TRect {
    defer exceptionProc()
    return PaintBox_GetClientRect(p.instance)
}

func (p *TPaintBox) ClientWidth() int32 {
    defer exceptionProc()
    return PaintBox_GetClientWidth(p.instance)
}

func (p *TPaintBox) SetClientWidth(value int32) {
    defer exceptionProc()
    PaintBox_SetClientWidth(p.instance, value)
}

func (p *TPaintBox) ExplicitLeft() int32 {
    defer exceptionProc()
    return PaintBox_GetExplicitLeft(p.instance)
}

func (p *TPaintBox) ExplicitTop() int32 {
    defer exceptionProc()
    return PaintBox_GetExplicitTop(p.instance)
}

func (p *TPaintBox) ExplicitWidth() int32 {
    defer exceptionProc()
    return PaintBox_GetExplicitWidth(p.instance)
}

func (p *TPaintBox) ExplicitHeight() int32 {
    defer exceptionProc()
    return PaintBox_GetExplicitHeight(p.instance)
}

func (p *TPaintBox) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(PaintBox_GetParent(p.instance))
}

func (p *TPaintBox) SetParent(value IControl) {
    defer exceptionProc()
    PaintBox_SetParent(p.instance, CheckPtr(value))
}

func (p *TPaintBox) AlignWithMargins() bool {
    defer exceptionProc()
    return PaintBox_GetAlignWithMargins(p.instance)
}

func (p *TPaintBox) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    PaintBox_SetAlignWithMargins(p.instance, value)
}

func (p *TPaintBox) Left() int32 {
    defer exceptionProc()
    return PaintBox_GetLeft(p.instance)
}

func (p *TPaintBox) SetLeft(value int32) {
    defer exceptionProc()
    PaintBox_SetLeft(p.instance, value)
}

func (p *TPaintBox) Top() int32 {
    defer exceptionProc()
    return PaintBox_GetTop(p.instance)
}

func (p *TPaintBox) SetTop(value int32) {
    defer exceptionProc()
    PaintBox_SetTop(p.instance, value)
}

func (p *TPaintBox) Width() int32 {
    defer exceptionProc()
    return PaintBox_GetWidth(p.instance)
}

func (p *TPaintBox) SetWidth(value int32) {
    defer exceptionProc()
    PaintBox_SetWidth(p.instance, value)
}

func (p *TPaintBox) Height() int32 {
    defer exceptionProc()
    return PaintBox_GetHeight(p.instance)
}

func (p *TPaintBox) SetHeight(value int32) {
    defer exceptionProc()
    PaintBox_SetHeight(p.instance, value)
}

func (p *TPaintBox) Cursor() TCursor {
    defer exceptionProc()
    return PaintBox_GetCursor(p.instance)
}

func (p *TPaintBox) SetCursor(value TCursor) {
    defer exceptionProc()
    PaintBox_SetCursor(p.instance, value)
}

func (p *TPaintBox) Hint() string {
    defer exceptionProc()
    return PaintBox_GetHint(p.instance)
}

func (p *TPaintBox) SetHint(value string) {
    defer exceptionProc()
    PaintBox_SetHint(p.instance, value)
}

func (p *TPaintBox) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(PaintBox_GetMargins(p.instance))
}

func (p *TPaintBox) SetMargins(value *TMargins) {
    defer exceptionProc()
    PaintBox_SetMargins(p.instance, CheckPtr(value))
}

func (p *TPaintBox) ComponentCount() int32 {
    defer exceptionProc()
    return PaintBox_GetComponentCount(p.instance)
}

func (p *TPaintBox) ComponentIndex() int32 {
    defer exceptionProc()
    return PaintBox_GetComponentIndex(p.instance)
}

func (p *TPaintBox) SetComponentIndex(value int32) {
    defer exceptionProc()
    PaintBox_SetComponentIndex(p.instance, value)
}

func (p *TPaintBox) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(PaintBox_GetOwner(p.instance))
}

func (p *TPaintBox) Name() string {
    defer exceptionProc()
    return PaintBox_GetName(p.instance)
}

func (p *TPaintBox) SetName(value string) {
    defer exceptionProc()
    PaintBox_SetName(p.instance, value)
}

func (p *TPaintBox) Tag() int {
    defer exceptionProc()
    return PaintBox_GetTag(p.instance)
}

func (p *TPaintBox) SetTag(value int) {
    defer exceptionProc()
    PaintBox_SetTag(p.instance, value)
}

func (p *TPaintBox) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(PaintBox_GetComponents(p.instance, AIndex))
}

