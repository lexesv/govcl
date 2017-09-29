
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TLabel struct {
    IControl
    instance uintptr
}

func NewLabel(owner IComponent) *TLabel {
    l := new(TLabel)
    l.instance = Label_Create(owner.Instance())
    return l
}

func LabelFromInst(inst uintptr) *TLabel {
    l := new(TLabel)
    l.instance = inst
    return l
}

func LabelFromObj(obj IObject) *TLabel {
    l := new(TLabel)
    l.instance = CheckPtr(obj)
    return l
}

func (l *TLabel) Free() {
    if l.instance != 0 {
        Label_Free(l.instance)
    }
}

func (l *TLabel) Instance() uintptr {
    return l.instance
}

func (l *TLabel) IsValid() bool {
    return l.instance != 0
}

func (l *TLabel) BringToFront() {
    defer exceptionProc()
    Label_BringToFront(l.instance)
}

func (l *TLabel) HasParent() bool {
    defer exceptionProc()
    return Label_HasParent(l.instance)
}

func (l *TLabel) Hide() {
    defer exceptionProc()
    Label_Hide(l.instance)
}

func (l *TLabel) Invalidate() {
    defer exceptionProc()
    Label_Invalidate(l.instance)
}

func (l *TLabel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return Label_Perform(l.instance, Msg , WParam , LParam )
}

func (l *TLabel) Refresh() {
    defer exceptionProc()
    Label_Refresh(l.instance)
}

func (l *TLabel) Repaint() {
    defer exceptionProc()
    Label_Repaint(l.instance)
}

func (l *TLabel) SendToBack() {
    defer exceptionProc()
    Label_SendToBack(l.instance)
}

func (l *TLabel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    Label_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight )
}

func (l *TLabel) Show() {
    defer exceptionProc()
    Label_Show(l.instance)
}

func (l *TLabel) Update() {
    defer exceptionProc()
    Label_Update(l.instance)
}

func (l *TLabel) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return Label_GetTextBuf(l.instance, Buffer , BufSize )
}

func (l *TLabel) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Label_FindComponent(l.instance, AName ))
}

func (l *TLabel) GetNamePath() string {
    defer exceptionProc()
    return Label_GetNamePath(l.instance)
}

func (l *TLabel) Assign(Source IObject) {
    defer exceptionProc()
    Label_Assign(l.instance, CheckPtr(Source))
}

func (l *TLabel) ClassName() string {
    defer exceptionProc()
    return Label_ClassName(l.instance)
}

func (l *TLabel) Equals(Obj IObject) bool {
    defer exceptionProc()
    return Label_Equals(l.instance, CheckPtr(Obj))
}

func (l *TLabel) GetHashCode() int32 {
    defer exceptionProc()
    return Label_GetHashCode(l.instance)
}

func (l *TLabel) ToString() string {
    defer exceptionProc()
    return Label_ToString(l.instance)
}

func (l *TLabel) Align() TAlign {
    defer exceptionProc()
    return Label_GetAlign(l.instance)
}

func (l *TLabel) SetAlign(value TAlign) {
    defer exceptionProc()
    Label_SetAlign(l.instance, value)
}

func (l *TLabel) Alignment() TAlignment {
    defer exceptionProc()
    return Label_GetAlignment(l.instance)
}

func (l *TLabel) SetAlignment(value TAlignment) {
    defer exceptionProc()
    Label_SetAlignment(l.instance, value)
}

func (l *TLabel) Anchors() TAnchors {
    defer exceptionProc()
    return Label_GetAnchors(l.instance)
}

func (l *TLabel) SetAnchors(value TAnchors) {
    defer exceptionProc()
    Label_SetAnchors(l.instance, value)
}

func (l *TLabel) AutoSize() bool {
    defer exceptionProc()
    return Label_GetAutoSize(l.instance)
}

func (l *TLabel) SetAutoSize(value bool) {
    defer exceptionProc()
    Label_SetAutoSize(l.instance, value)
}

func (l *TLabel) Caption() string {
    defer exceptionProc()
    return Label_GetCaption(l.instance)
}

func (l *TLabel) SetCaption(value string) {
    defer exceptionProc()
    Label_SetCaption(l.instance, value)
}

func (l *TLabel) Color() TColor {
    defer exceptionProc()
    return Label_GetColor(l.instance)
}

func (l *TLabel) SetColor(value TColor) {
    defer exceptionProc()
    Label_SetColor(l.instance, value)
}

func (l *TLabel) EllipsisPosition() TEllipsisPosition {
    defer exceptionProc()
    return Label_GetEllipsisPosition(l.instance)
}

func (l *TLabel) SetEllipsisPosition(value TEllipsisPosition) {
    defer exceptionProc()
    Label_SetEllipsisPosition(l.instance, value)
}

func (l *TLabel) Enabled() bool {
    defer exceptionProc()
    return Label_GetEnabled(l.instance)
}

func (l *TLabel) SetEnabled(value bool) {
    defer exceptionProc()
    Label_SetEnabled(l.instance, value)
}

func (l *TLabel) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(Label_GetFont(l.instance))
}

func (l *TLabel) SetFont(value *TFont) {
    defer exceptionProc()
    Label_SetFont(l.instance, CheckPtr(value))
}

func (l *TLabel) GlowSize() int32 {
    defer exceptionProc()
    return Label_GetGlowSize(l.instance)
}

func (l *TLabel) SetGlowSize(value int32) {
    defer exceptionProc()
    Label_SetGlowSize(l.instance, value)
}

func (l *TLabel) ParentColor() bool {
    defer exceptionProc()
    return Label_GetParentColor(l.instance)
}

func (l *TLabel) SetParentColor(value bool) {
    defer exceptionProc()
    Label_SetParentColor(l.instance, value)
}

func (l *TLabel) ParentFont() bool {
    defer exceptionProc()
    return Label_GetParentFont(l.instance)
}

func (l *TLabel) SetParentFont(value bool) {
    defer exceptionProc()
    Label_SetParentFont(l.instance, value)
}

func (l *TLabel) ParentShowHint() bool {
    defer exceptionProc()
    return Label_GetParentShowHint(l.instance)
}

func (l *TLabel) SetParentShowHint(value bool) {
    defer exceptionProc()
    Label_SetParentShowHint(l.instance, value)
}

func (l *TLabel) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(Label_GetPopupMenu(l.instance))
}

func (l *TLabel) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    Label_SetPopupMenu(l.instance, CheckPtr(value))
}

func (l *TLabel) ShowAccelChar() bool {
    defer exceptionProc()
    return Label_GetShowAccelChar(l.instance)
}

func (l *TLabel) SetShowAccelChar(value bool) {
    defer exceptionProc()
    Label_SetShowAccelChar(l.instance, value)
}

func (l *TLabel) ShowHint() bool {
    defer exceptionProc()
    return Label_GetShowHint(l.instance)
}

func (l *TLabel) SetShowHint(value bool) {
    defer exceptionProc()
    Label_SetShowHint(l.instance, value)
}

func (l *TLabel) Transparent() bool {
    defer exceptionProc()
    return Label_GetTransparent(l.instance)
}

func (l *TLabel) SetTransparent(value bool) {
    defer exceptionProc()
    Label_SetTransparent(l.instance, value)
}

func (l *TLabel) Layout() TTextLayout {
    defer exceptionProc()
    return Label_GetLayout(l.instance)
}

func (l *TLabel) SetLayout(value TTextLayout) {
    defer exceptionProc()
    Label_SetLayout(l.instance, value)
}

func (l *TLabel) Visible() bool {
    defer exceptionProc()
    return Label_GetVisible(l.instance)
}

func (l *TLabel) SetVisible(value bool) {
    defer exceptionProc()
    Label_SetVisible(l.instance, value)
}

func (l *TLabel) WordWrap() bool {
    defer exceptionProc()
    return Label_GetWordWrap(l.instance)
}

func (l *TLabel) SetWordWrap(value bool) {
    defer exceptionProc()
    Label_SetWordWrap(l.instance, value)
}

func (l *TLabel) SetOnClick(fn TNotifyEvent) {
    Label_SetOnClick(l.instance, fn)
}

func (l *TLabel) SetOnDblClick(fn TNotifyEvent) {
    Label_SetOnDblClick(l.instance, fn)
}

func (l *TLabel) SetOnMouseDown(fn TMouseEvent) {
    Label_SetOnMouseDown(l.instance, fn)
}

func (l *TLabel) SetOnMouseMove(fn TMouseMoveEvent) {
    Label_SetOnMouseMove(l.instance, fn)
}

func (l *TLabel) SetOnMouseUp(fn TMouseEvent) {
    Label_SetOnMouseUp(l.instance, fn)
}

func (l *TLabel) SetOnMouseEnter(fn TNotifyEvent) {
    Label_SetOnMouseEnter(l.instance, fn)
}

func (l *TLabel) SetOnMouseLeave(fn TNotifyEvent) {
    Label_SetOnMouseLeave(l.instance, fn)
}

func (l *TLabel) Canvas() *TCanvas {
    defer exceptionProc()
    return CanvasFromInst(Label_GetCanvas(l.instance))
}

func (l *TLabel) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(Label_GetAction(l.instance))
}

func (l *TLabel) SetAction(value IComponent) {
    defer exceptionProc()
    Label_SetAction(l.instance, CheckPtr(value))
}

func (l *TLabel) BoundsRect() TRect {
    defer exceptionProc()
    return Label_GetBoundsRect(l.instance)
}

func (l *TLabel) SetBoundsRect(value TRect) {
    defer exceptionProc()
    Label_SetBoundsRect(l.instance, value)
}

func (l *TLabel) ClientHeight() int32 {
    defer exceptionProc()
    return Label_GetClientHeight(l.instance)
}

func (l *TLabel) SetClientHeight(value int32) {
    defer exceptionProc()
    Label_SetClientHeight(l.instance, value)
}

func (l *TLabel) ClientRect() TRect {
    defer exceptionProc()
    return Label_GetClientRect(l.instance)
}

func (l *TLabel) ClientWidth() int32 {
    defer exceptionProc()
    return Label_GetClientWidth(l.instance)
}

func (l *TLabel) SetClientWidth(value int32) {
    defer exceptionProc()
    Label_SetClientWidth(l.instance, value)
}

func (l *TLabel) ExplicitLeft() int32 {
    defer exceptionProc()
    return Label_GetExplicitLeft(l.instance)
}

func (l *TLabel) ExplicitTop() int32 {
    defer exceptionProc()
    return Label_GetExplicitTop(l.instance)
}

func (l *TLabel) ExplicitWidth() int32 {
    defer exceptionProc()
    return Label_GetExplicitWidth(l.instance)
}

func (l *TLabel) ExplicitHeight() int32 {
    defer exceptionProc()
    return Label_GetExplicitHeight(l.instance)
}

func (l *TLabel) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(Label_GetParent(l.instance))
}

func (l *TLabel) SetParent(value IControl) {
    defer exceptionProc()
    Label_SetParent(l.instance, CheckPtr(value))
}

func (l *TLabel) AlignWithMargins() bool {
    defer exceptionProc()
    return Label_GetAlignWithMargins(l.instance)
}

func (l *TLabel) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    Label_SetAlignWithMargins(l.instance, value)
}

func (l *TLabel) Left() int32 {
    defer exceptionProc()
    return Label_GetLeft(l.instance)
}

func (l *TLabel) SetLeft(value int32) {
    defer exceptionProc()
    Label_SetLeft(l.instance, value)
}

func (l *TLabel) Top() int32 {
    defer exceptionProc()
    return Label_GetTop(l.instance)
}

func (l *TLabel) SetTop(value int32) {
    defer exceptionProc()
    Label_SetTop(l.instance, value)
}

func (l *TLabel) Width() int32 {
    defer exceptionProc()
    return Label_GetWidth(l.instance)
}

func (l *TLabel) SetWidth(value int32) {
    defer exceptionProc()
    Label_SetWidth(l.instance, value)
}

func (l *TLabel) Height() int32 {
    defer exceptionProc()
    return Label_GetHeight(l.instance)
}

func (l *TLabel) SetHeight(value int32) {
    defer exceptionProc()
    Label_SetHeight(l.instance, value)
}

func (l *TLabel) Cursor() TCursor {
    defer exceptionProc()
    return Label_GetCursor(l.instance)
}

func (l *TLabel) SetCursor(value TCursor) {
    defer exceptionProc()
    Label_SetCursor(l.instance, value)
}

func (l *TLabel) Hint() string {
    defer exceptionProc()
    return Label_GetHint(l.instance)
}

func (l *TLabel) SetHint(value string) {
    defer exceptionProc()
    Label_SetHint(l.instance, value)
}

func (l *TLabel) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(Label_GetMargins(l.instance))
}

func (l *TLabel) SetMargins(value *TMargins) {
    defer exceptionProc()
    Label_SetMargins(l.instance, CheckPtr(value))
}

func (l *TLabel) ComponentCount() int32 {
    defer exceptionProc()
    return Label_GetComponentCount(l.instance)
}

func (l *TLabel) ComponentIndex() int32 {
    defer exceptionProc()
    return Label_GetComponentIndex(l.instance)
}

func (l *TLabel) SetComponentIndex(value int32) {
    defer exceptionProc()
    Label_SetComponentIndex(l.instance, value)
}

func (l *TLabel) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Label_GetOwner(l.instance))
}

func (l *TLabel) Name() string {
    defer exceptionProc()
    return Label_GetName(l.instance)
}

func (l *TLabel) SetName(value string) {
    defer exceptionProc()
    Label_SetName(l.instance, value)
}

func (l *TLabel) Tag() int {
    defer exceptionProc()
    return Label_GetTag(l.instance)
}

func (l *TLabel) SetTag(value int) {
    defer exceptionProc()
    Label_SetTag(l.instance, value)
}

func (l *TLabel) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Label_GetComponents(l.instance, AIndex))
}

