
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TSplitter struct {
    IControl
    instance uintptr
}

func NewSplitter(owner IComponent) *TSplitter {
    s := new(TSplitter)
    s.instance = Splitter_Create(owner.Instance())
    return s
}

func SplitterFromInst(inst uintptr) *TSplitter {
    s := new(TSplitter)
    s.instance = inst
    return s
}

func SplitterFromObj(obj IObject) *TSplitter {
    s := new(TSplitter)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TSplitter) Free() {
    if s.instance != 0 {
        Splitter_Free(s.instance)
    }
}

func (s *TSplitter) Instance() uintptr {
    return s.instance
}

func (s *TSplitter) IsValid() bool {
    return s.instance != 0
}

func (s *TSplitter) BringToFront() {
    defer exceptionProc()
    Splitter_BringToFront(s.instance)
}

func (s *TSplitter) HasParent() bool {
    defer exceptionProc()
    return Splitter_HasParent(s.instance)
}

func (s *TSplitter) Hide() {
    defer exceptionProc()
    Splitter_Hide(s.instance)
}

func (s *TSplitter) Invalidate() {
    defer exceptionProc()
    Splitter_Invalidate(s.instance)
}

func (s *TSplitter) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return Splitter_Perform(s.instance, Msg , WParam , LParam )
}

func (s *TSplitter) Refresh() {
    defer exceptionProc()
    Splitter_Refresh(s.instance)
}

func (s *TSplitter) Repaint() {
    defer exceptionProc()
    Splitter_Repaint(s.instance)
}

func (s *TSplitter) SendToBack() {
    defer exceptionProc()
    Splitter_SendToBack(s.instance)
}

func (s *TSplitter) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    Splitter_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight )
}

func (s *TSplitter) Show() {
    defer exceptionProc()
    Splitter_Show(s.instance)
}

func (s *TSplitter) Update() {
    defer exceptionProc()
    Splitter_Update(s.instance)
}

func (s *TSplitter) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return Splitter_GetTextBuf(s.instance, Buffer , BufSize )
}

func (s *TSplitter) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Splitter_FindComponent(s.instance, AName ))
}

func (s *TSplitter) GetNamePath() string {
    defer exceptionProc()
    return Splitter_GetNamePath(s.instance)
}

func (s *TSplitter) Assign(Source IObject) {
    defer exceptionProc()
    Splitter_Assign(s.instance, CheckPtr(Source))
}

func (s *TSplitter) ClassName() string {
    defer exceptionProc()
    return Splitter_ClassName(s.instance)
}

func (s *TSplitter) Equals(Obj IObject) bool {
    defer exceptionProc()
    return Splitter_Equals(s.instance, CheckPtr(Obj))
}

func (s *TSplitter) GetHashCode() int32 {
    defer exceptionProc()
    return Splitter_GetHashCode(s.instance)
}

func (s *TSplitter) ToString() string {
    defer exceptionProc()
    return Splitter_ToString(s.instance)
}

func (s *TSplitter) Canvas() *TCanvas {
    defer exceptionProc()
    return CanvasFromInst(Splitter_GetCanvas(s.instance))
}

func (s *TSplitter) Align() TAlign {
    defer exceptionProc()
    return Splitter_GetAlign(s.instance)
}

func (s *TSplitter) SetAlign(value TAlign) {
    defer exceptionProc()
    Splitter_SetAlign(s.instance, value)
}

func (s *TSplitter) Color() TColor {
    defer exceptionProc()
    return Splitter_GetColor(s.instance)
}

func (s *TSplitter) SetColor(value TColor) {
    defer exceptionProc()
    Splitter_SetColor(s.instance, value)
}

func (s *TSplitter) Cursor() TCursor {
    defer exceptionProc()
    return Splitter_GetCursor(s.instance)
}

func (s *TSplitter) SetCursor(value TCursor) {
    defer exceptionProc()
    Splitter_SetCursor(s.instance, value)
}

func (s *TSplitter) ParentColor() bool {
    defer exceptionProc()
    return Splitter_GetParentColor(s.instance)
}

func (s *TSplitter) SetParentColor(value bool) {
    defer exceptionProc()
    Splitter_SetParentColor(s.instance, value)
}

func (s *TSplitter) Visible() bool {
    defer exceptionProc()
    return Splitter_GetVisible(s.instance)
}

func (s *TSplitter) SetVisible(value bool) {
    defer exceptionProc()
    Splitter_SetVisible(s.instance, value)
}

func (s *TSplitter) Width() int32 {
    defer exceptionProc()
    return Splitter_GetWidth(s.instance)
}

func (s *TSplitter) SetWidth(value int32) {
    defer exceptionProc()
    Splitter_SetWidth(s.instance, value)
}

func (s *TSplitter) SetOnPaint(fn TNotifyEvent) {
    Splitter_SetOnPaint(s.instance, fn)
}

func (s *TSplitter) Enabled() bool {
    defer exceptionProc()
    return Splitter_GetEnabled(s.instance)
}

func (s *TSplitter) SetEnabled(value bool) {
    defer exceptionProc()
    Splitter_SetEnabled(s.instance, value)
}

func (s *TSplitter) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(Splitter_GetAction(s.instance))
}

func (s *TSplitter) SetAction(value IComponent) {
    defer exceptionProc()
    Splitter_SetAction(s.instance, CheckPtr(value))
}

func (s *TSplitter) Anchors() TAnchors {
    defer exceptionProc()
    return Splitter_GetAnchors(s.instance)
}

func (s *TSplitter) SetAnchors(value TAnchors) {
    defer exceptionProc()
    Splitter_SetAnchors(s.instance, value)
}

func (s *TSplitter) BoundsRect() TRect {
    defer exceptionProc()
    return Splitter_GetBoundsRect(s.instance)
}

func (s *TSplitter) SetBoundsRect(value TRect) {
    defer exceptionProc()
    Splitter_SetBoundsRect(s.instance, value)
}

func (s *TSplitter) ClientHeight() int32 {
    defer exceptionProc()
    return Splitter_GetClientHeight(s.instance)
}

func (s *TSplitter) SetClientHeight(value int32) {
    defer exceptionProc()
    Splitter_SetClientHeight(s.instance, value)
}

func (s *TSplitter) ClientRect() TRect {
    defer exceptionProc()
    return Splitter_GetClientRect(s.instance)
}

func (s *TSplitter) ClientWidth() int32 {
    defer exceptionProc()
    return Splitter_GetClientWidth(s.instance)
}

func (s *TSplitter) SetClientWidth(value int32) {
    defer exceptionProc()
    Splitter_SetClientWidth(s.instance, value)
}

func (s *TSplitter) ExplicitLeft() int32 {
    defer exceptionProc()
    return Splitter_GetExplicitLeft(s.instance)
}

func (s *TSplitter) ExplicitTop() int32 {
    defer exceptionProc()
    return Splitter_GetExplicitTop(s.instance)
}

func (s *TSplitter) ExplicitWidth() int32 {
    defer exceptionProc()
    return Splitter_GetExplicitWidth(s.instance)
}

func (s *TSplitter) ExplicitHeight() int32 {
    defer exceptionProc()
    return Splitter_GetExplicitHeight(s.instance)
}

func (s *TSplitter) ShowHint() bool {
    defer exceptionProc()
    return Splitter_GetShowHint(s.instance)
}

func (s *TSplitter) SetShowHint(value bool) {
    defer exceptionProc()
    Splitter_SetShowHint(s.instance, value)
}

func (s *TSplitter) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(Splitter_GetParent(s.instance))
}

func (s *TSplitter) SetParent(value IControl) {
    defer exceptionProc()
    Splitter_SetParent(s.instance, CheckPtr(value))
}

func (s *TSplitter) AlignWithMargins() bool {
    defer exceptionProc()
    return Splitter_GetAlignWithMargins(s.instance)
}

func (s *TSplitter) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    Splitter_SetAlignWithMargins(s.instance, value)
}

func (s *TSplitter) Left() int32 {
    defer exceptionProc()
    return Splitter_GetLeft(s.instance)
}

func (s *TSplitter) SetLeft(value int32) {
    defer exceptionProc()
    Splitter_SetLeft(s.instance, value)
}

func (s *TSplitter) Top() int32 {
    defer exceptionProc()
    return Splitter_GetTop(s.instance)
}

func (s *TSplitter) SetTop(value int32) {
    defer exceptionProc()
    Splitter_SetTop(s.instance, value)
}

func (s *TSplitter) Height() int32 {
    defer exceptionProc()
    return Splitter_GetHeight(s.instance)
}

func (s *TSplitter) SetHeight(value int32) {
    defer exceptionProc()
    Splitter_SetHeight(s.instance, value)
}

func (s *TSplitter) Hint() string {
    defer exceptionProc()
    return Splitter_GetHint(s.instance)
}

func (s *TSplitter) SetHint(value string) {
    defer exceptionProc()
    Splitter_SetHint(s.instance, value)
}

func (s *TSplitter) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(Splitter_GetMargins(s.instance))
}

func (s *TSplitter) SetMargins(value *TMargins) {
    defer exceptionProc()
    Splitter_SetMargins(s.instance, CheckPtr(value))
}

func (s *TSplitter) ComponentCount() int32 {
    defer exceptionProc()
    return Splitter_GetComponentCount(s.instance)
}

func (s *TSplitter) ComponentIndex() int32 {
    defer exceptionProc()
    return Splitter_GetComponentIndex(s.instance)
}

func (s *TSplitter) SetComponentIndex(value int32) {
    defer exceptionProc()
    Splitter_SetComponentIndex(s.instance, value)
}

func (s *TSplitter) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Splitter_GetOwner(s.instance))
}

func (s *TSplitter) Name() string {
    defer exceptionProc()
    return Splitter_GetName(s.instance)
}

func (s *TSplitter) SetName(value string) {
    defer exceptionProc()
    Splitter_SetName(s.instance, value)
}

func (s *TSplitter) Tag() int {
    defer exceptionProc()
    return Splitter_GetTag(s.instance)
}

func (s *TSplitter) SetTag(value int) {
    defer exceptionProc()
    Splitter_SetTag(s.instance, value)
}

func (s *TSplitter) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Splitter_GetComponents(s.instance, AIndex))
}

