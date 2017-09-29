
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

type TStaticText struct {
    IControl
    instance uintptr
}

func NewStaticText(owner IComponent) *TStaticText {
    s := new(TStaticText)
    s.instance = StaticText_Create(owner.Instance())
    return s
}

func StaticTextFromInst(inst uintptr) *TStaticText {
    s := new(TStaticText)
    s.instance = inst
    return s
}

func StaticTextFromObj(obj IObject) *TStaticText {
    s := new(TStaticText)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TStaticText) Free() {
    if s.instance != 0 {
        StaticText_Free(s.instance)
    }
}

func (s *TStaticText) Instance() uintptr {
    return s.instance
}

func (s *TStaticText) IsValid() bool {
    return s.instance != 0
}

func (s *TStaticText) CanFocus() bool {
    defer exceptionProc()
    return StaticText_CanFocus(s.instance)
}

func (s *TStaticText) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    StaticText_FlipChildren(s.instance, AllLevels )
}

func (s *TStaticText) Focused() bool {
    defer exceptionProc()
    return StaticText_Focused(s.instance)
}

func (s *TStaticText) HandleAllocated() bool {
    defer exceptionProc()
    return StaticText_HandleAllocated(s.instance)
}

func (s *TStaticText) Invalidate() {
    defer exceptionProc()
    StaticText_Invalidate(s.instance)
}

func (s *TStaticText) Realign() {
    defer exceptionProc()
    StaticText_Realign(s.instance)
}

func (s *TStaticText) Repaint() {
    defer exceptionProc()
    StaticText_Repaint(s.instance)
}

func (s *TStaticText) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    StaticText_ScaleBy(s.instance, M , D )
}

func (s *TStaticText) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    StaticText_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight )
}

func (s *TStaticText) SetFocus() {
    defer exceptionProc()
    StaticText_SetFocus(s.instance)
}

func (s *TStaticText) Update() {
    defer exceptionProc()
    StaticText_Update(s.instance)
}

func (s *TStaticText) BringToFront() {
    defer exceptionProc()
    StaticText_BringToFront(s.instance)
}

func (s *TStaticText) HasParent() bool {
    defer exceptionProc()
    return StaticText_HasParent(s.instance)
}

func (s *TStaticText) Hide() {
    defer exceptionProc()
    StaticText_Hide(s.instance)
}

func (s *TStaticText) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return StaticText_Perform(s.instance, Msg , WParam , LParam )
}

func (s *TStaticText) Refresh() {
    defer exceptionProc()
    StaticText_Refresh(s.instance)
}

func (s *TStaticText) SendToBack() {
    defer exceptionProc()
    StaticText_SendToBack(s.instance)
}

func (s *TStaticText) Show() {
    defer exceptionProc()
    StaticText_Show(s.instance)
}

func (s *TStaticText) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return StaticText_GetTextBuf(s.instance, Buffer , BufSize )
}

func (s *TStaticText) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(StaticText_FindComponent(s.instance, AName ))
}

func (s *TStaticText) GetNamePath() string {
    defer exceptionProc()
    return StaticText_GetNamePath(s.instance)
}

func (s *TStaticText) Assign(Source IObject) {
    defer exceptionProc()
    StaticText_Assign(s.instance, CheckPtr(Source))
}

func (s *TStaticText) ClassName() string {
    defer exceptionProc()
    return StaticText_ClassName(s.instance)
}

func (s *TStaticText) Equals(Obj IObject) bool {
    defer exceptionProc()
    return StaticText_Equals(s.instance, CheckPtr(Obj))
}

func (s *TStaticText) GetHashCode() int32 {
    defer exceptionProc()
    return StaticText_GetHashCode(s.instance)
}

func (s *TStaticText) ToString() string {
    defer exceptionProc()
    return StaticText_ToString(s.instance)
}

func (s *TStaticText) Align() TAlign {
    defer exceptionProc()
    return StaticText_GetAlign(s.instance)
}

func (s *TStaticText) SetAlign(value TAlign) {
    defer exceptionProc()
    StaticText_SetAlign(s.instance, value)
}

func (s *TStaticText) Alignment() TAlignment {
    defer exceptionProc()
    return StaticText_GetAlignment(s.instance)
}

func (s *TStaticText) SetAlignment(value TAlignment) {
    defer exceptionProc()
    StaticText_SetAlignment(s.instance, value)
}

func (s *TStaticText) Anchors() TAnchors {
    defer exceptionProc()
    return StaticText_GetAnchors(s.instance)
}

func (s *TStaticText) SetAnchors(value TAnchors) {
    defer exceptionProc()
    StaticText_SetAnchors(s.instance, value)
}

func (s *TStaticText) AutoSize() bool {
    defer exceptionProc()
    return StaticText_GetAutoSize(s.instance)
}

func (s *TStaticText) SetAutoSize(value bool) {
    defer exceptionProc()
    StaticText_SetAutoSize(s.instance, value)
}

func (s *TStaticText) BorderStyle() TStaticBorderStyle {
    defer exceptionProc()
    return StaticText_GetBorderStyle(s.instance)
}

func (s *TStaticText) SetBorderStyle(value TStaticBorderStyle) {
    defer exceptionProc()
    StaticText_SetBorderStyle(s.instance, value)
}

func (s *TStaticText) Caption() string {
    defer exceptionProc()
    return StaticText_GetCaption(s.instance)
}

func (s *TStaticText) SetCaption(value string) {
    defer exceptionProc()
    StaticText_SetCaption(s.instance, value)
}

func (s *TStaticText) Color() TColor {
    defer exceptionProc()
    return StaticText_GetColor(s.instance)
}

func (s *TStaticText) SetColor(value TColor) {
    defer exceptionProc()
    StaticText_SetColor(s.instance, value)
}

func (s *TStaticText) DoubleBuffered() bool {
    defer exceptionProc()
    return StaticText_GetDoubleBuffered(s.instance)
}

func (s *TStaticText) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    StaticText_SetDoubleBuffered(s.instance, value)
}

func (s *TStaticText) Enabled() bool {
    defer exceptionProc()
    return StaticText_GetEnabled(s.instance)
}

func (s *TStaticText) SetEnabled(value bool) {
    defer exceptionProc()
    StaticText_SetEnabled(s.instance, value)
}

func (s *TStaticText) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(StaticText_GetFont(s.instance))
}

func (s *TStaticText) SetFont(value *TFont) {
    defer exceptionProc()
    StaticText_SetFont(s.instance, CheckPtr(value))
}

func (s *TStaticText) ParentColor() bool {
    defer exceptionProc()
    return StaticText_GetParentColor(s.instance)
}

func (s *TStaticText) SetParentColor(value bool) {
    defer exceptionProc()
    StaticText_SetParentColor(s.instance, value)
}

func (s *TStaticText) ParentFont() bool {
    defer exceptionProc()
    return StaticText_GetParentFont(s.instance)
}

func (s *TStaticText) SetParentFont(value bool) {
    defer exceptionProc()
    StaticText_SetParentFont(s.instance, value)
}

func (s *TStaticText) ParentShowHint() bool {
    defer exceptionProc()
    return StaticText_GetParentShowHint(s.instance)
}

func (s *TStaticText) SetParentShowHint(value bool) {
    defer exceptionProc()
    StaticText_SetParentShowHint(s.instance, value)
}

func (s *TStaticText) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(StaticText_GetPopupMenu(s.instance))
}

func (s *TStaticText) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    StaticText_SetPopupMenu(s.instance, CheckPtr(value))
}

func (s *TStaticText) ShowAccelChar() bool {
    defer exceptionProc()
    return StaticText_GetShowAccelChar(s.instance)
}

func (s *TStaticText) SetShowAccelChar(value bool) {
    defer exceptionProc()
    StaticText_SetShowAccelChar(s.instance, value)
}

func (s *TStaticText) ShowHint() bool {
    defer exceptionProc()
    return StaticText_GetShowHint(s.instance)
}

func (s *TStaticText) SetShowHint(value bool) {
    defer exceptionProc()
    StaticText_SetShowHint(s.instance, value)
}

func (s *TStaticText) TabOrder() int16 {
    defer exceptionProc()
    return StaticText_GetTabOrder(s.instance)
}

func (s *TStaticText) SetTabOrder(value int16) {
    defer exceptionProc()
    StaticText_SetTabOrder(s.instance, value)
}

func (s *TStaticText) TabStop() bool {
    defer exceptionProc()
    return StaticText_GetTabStop(s.instance)
}

func (s *TStaticText) SetTabStop(value bool) {
    defer exceptionProc()
    StaticText_SetTabStop(s.instance, value)
}

func (s *TStaticText) Transparent() bool {
    defer exceptionProc()
    return StaticText_GetTransparent(s.instance)
}

func (s *TStaticText) SetTransparent(value bool) {
    defer exceptionProc()
    StaticText_SetTransparent(s.instance, value)
}

func (s *TStaticText) Visible() bool {
    defer exceptionProc()
    return StaticText_GetVisible(s.instance)
}

func (s *TStaticText) SetVisible(value bool) {
    defer exceptionProc()
    StaticText_SetVisible(s.instance, value)
}

func (s *TStaticText) SetOnClick(fn TNotifyEvent) {
    StaticText_SetOnClick(s.instance, fn)
}

func (s *TStaticText) SetOnDblClick(fn TNotifyEvent) {
    StaticText_SetOnDblClick(s.instance, fn)
}

func (s *TStaticText) SetOnMouseDown(fn TMouseEvent) {
    StaticText_SetOnMouseDown(s.instance, fn)
}

func (s *TStaticText) SetOnMouseEnter(fn TNotifyEvent) {
    StaticText_SetOnMouseEnter(s.instance, fn)
}

func (s *TStaticText) SetOnMouseLeave(fn TNotifyEvent) {
    StaticText_SetOnMouseLeave(s.instance, fn)
}

func (s *TStaticText) SetOnMouseMove(fn TMouseMoveEvent) {
    StaticText_SetOnMouseMove(s.instance, fn)
}

func (s *TStaticText) SetOnMouseUp(fn TMouseEvent) {
    StaticText_SetOnMouseUp(s.instance, fn)
}

func (s *TStaticText) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(StaticText_GetBrush(s.instance))
}

func (s *TStaticText) ControlCount() int32 {
    defer exceptionProc()
    return StaticText_GetControlCount(s.instance)
}

func (s *TStaticText) Handle() HWND {
    defer exceptionProc()
    return StaticText_GetHandle(s.instance)
}

func (s *TStaticText) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(StaticText_GetAction(s.instance))
}

func (s *TStaticText) SetAction(value IComponent) {
    defer exceptionProc()
    StaticText_SetAction(s.instance, CheckPtr(value))
}

func (s *TStaticText) BoundsRect() TRect {
    defer exceptionProc()
    return StaticText_GetBoundsRect(s.instance)
}

func (s *TStaticText) SetBoundsRect(value TRect) {
    defer exceptionProc()
    StaticText_SetBoundsRect(s.instance, value)
}

func (s *TStaticText) ClientHeight() int32 {
    defer exceptionProc()
    return StaticText_GetClientHeight(s.instance)
}

func (s *TStaticText) SetClientHeight(value int32) {
    defer exceptionProc()
    StaticText_SetClientHeight(s.instance, value)
}

func (s *TStaticText) ClientRect() TRect {
    defer exceptionProc()
    return StaticText_GetClientRect(s.instance)
}

func (s *TStaticText) ClientWidth() int32 {
    defer exceptionProc()
    return StaticText_GetClientWidth(s.instance)
}

func (s *TStaticText) SetClientWidth(value int32) {
    defer exceptionProc()
    StaticText_SetClientWidth(s.instance, value)
}

func (s *TStaticText) ExplicitLeft() int32 {
    defer exceptionProc()
    return StaticText_GetExplicitLeft(s.instance)
}

func (s *TStaticText) ExplicitTop() int32 {
    defer exceptionProc()
    return StaticText_GetExplicitTop(s.instance)
}

func (s *TStaticText) ExplicitWidth() int32 {
    defer exceptionProc()
    return StaticText_GetExplicitWidth(s.instance)
}

func (s *TStaticText) ExplicitHeight() int32 {
    defer exceptionProc()
    return StaticText_GetExplicitHeight(s.instance)
}

func (s *TStaticText) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(StaticText_GetParent(s.instance))
}

func (s *TStaticText) SetParent(value IControl) {
    defer exceptionProc()
    StaticText_SetParent(s.instance, CheckPtr(value))
}

func (s *TStaticText) AlignWithMargins() bool {
    defer exceptionProc()
    return StaticText_GetAlignWithMargins(s.instance)
}

func (s *TStaticText) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    StaticText_SetAlignWithMargins(s.instance, value)
}

func (s *TStaticText) Left() int32 {
    defer exceptionProc()
    return StaticText_GetLeft(s.instance)
}

func (s *TStaticText) SetLeft(value int32) {
    defer exceptionProc()
    StaticText_SetLeft(s.instance, value)
}

func (s *TStaticText) Top() int32 {
    defer exceptionProc()
    return StaticText_GetTop(s.instance)
}

func (s *TStaticText) SetTop(value int32) {
    defer exceptionProc()
    StaticText_SetTop(s.instance, value)
}

func (s *TStaticText) Width() int32 {
    defer exceptionProc()
    return StaticText_GetWidth(s.instance)
}

func (s *TStaticText) SetWidth(value int32) {
    defer exceptionProc()
    StaticText_SetWidth(s.instance, value)
}

func (s *TStaticText) Height() int32 {
    defer exceptionProc()
    return StaticText_GetHeight(s.instance)
}

func (s *TStaticText) SetHeight(value int32) {
    defer exceptionProc()
    StaticText_SetHeight(s.instance, value)
}

func (s *TStaticText) Cursor() TCursor {
    defer exceptionProc()
    return StaticText_GetCursor(s.instance)
}

func (s *TStaticText) SetCursor(value TCursor) {
    defer exceptionProc()
    StaticText_SetCursor(s.instance, value)
}

func (s *TStaticText) Hint() string {
    defer exceptionProc()
    return StaticText_GetHint(s.instance)
}

func (s *TStaticText) SetHint(value string) {
    defer exceptionProc()
    StaticText_SetHint(s.instance, value)
}

func (s *TStaticText) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(StaticText_GetMargins(s.instance))
}

func (s *TStaticText) SetMargins(value *TMargins) {
    defer exceptionProc()
    StaticText_SetMargins(s.instance, CheckPtr(value))
}

func (s *TStaticText) ComponentCount() int32 {
    defer exceptionProc()
    return StaticText_GetComponentCount(s.instance)
}

func (s *TStaticText) ComponentIndex() int32 {
    defer exceptionProc()
    return StaticText_GetComponentIndex(s.instance)
}

func (s *TStaticText) SetComponentIndex(value int32) {
    defer exceptionProc()
    StaticText_SetComponentIndex(s.instance, value)
}

func (s *TStaticText) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(StaticText_GetOwner(s.instance))
}

func (s *TStaticText) Name() string {
    defer exceptionProc()
    return StaticText_GetName(s.instance)
}

func (s *TStaticText) SetName(value string) {
    defer exceptionProc()
    StaticText_SetName(s.instance, value)
}

func (s *TStaticText) Tag() int {
    defer exceptionProc()
    return StaticText_GetTag(s.instance)
}

func (s *TStaticText) SetTag(value int) {
    defer exceptionProc()
    StaticText_SetTag(s.instance, value)
}

func (s *TStaticText) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(StaticText_GetControls(s.instance, Index))
}

func (s *TStaticText) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(StaticText_GetComponents(s.instance, AIndex))
}

