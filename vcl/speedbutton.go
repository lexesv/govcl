
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TSpeedButton struct {
    IControl
    instance uintptr
}

func NewSpeedButton(owner IComponent) *TSpeedButton {
    s := new(TSpeedButton)
    s.instance = SpeedButton_Create(owner.Instance())
    return s
}

func SpeedButtonFromInst(inst uintptr) *TSpeedButton {
    s := new(TSpeedButton)
    s.instance = inst
    return s
}

func SpeedButtonFromObj(obj IObject) *TSpeedButton {
    s := new(TSpeedButton)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TSpeedButton) Free() {
    if s.instance != 0 {
        SpeedButton_Free(s.instance)
    }
}

func (s *TSpeedButton) Instance() uintptr {
    return s.instance
}

func (s *TSpeedButton) IsValid() bool {
    return s.instance != 0
}

func (s *TSpeedButton) Click() {
    defer exceptionProc()
    SpeedButton_Click(s.instance)
}

func (s *TSpeedButton) BringToFront() {
    defer exceptionProc()
    SpeedButton_BringToFront(s.instance)
}

func (s *TSpeedButton) HasParent() bool {
    defer exceptionProc()
    return SpeedButton_HasParent(s.instance)
}

func (s *TSpeedButton) Hide() {
    defer exceptionProc()
    SpeedButton_Hide(s.instance)
}

func (s *TSpeedButton) Invalidate() {
    defer exceptionProc()
    SpeedButton_Invalidate(s.instance)
}

func (s *TSpeedButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return SpeedButton_Perform(s.instance, Msg , WParam , LParam )
}

func (s *TSpeedButton) Refresh() {
    defer exceptionProc()
    SpeedButton_Refresh(s.instance)
}

func (s *TSpeedButton) Repaint() {
    defer exceptionProc()
    SpeedButton_Repaint(s.instance)
}

func (s *TSpeedButton) SendToBack() {
    defer exceptionProc()
    SpeedButton_SendToBack(s.instance)
}

func (s *TSpeedButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    SpeedButton_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight )
}

func (s *TSpeedButton) Show() {
    defer exceptionProc()
    SpeedButton_Show(s.instance)
}

func (s *TSpeedButton) Update() {
    defer exceptionProc()
    SpeedButton_Update(s.instance)
}

func (s *TSpeedButton) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return SpeedButton_GetTextBuf(s.instance, Buffer , BufSize )
}

func (s *TSpeedButton) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(SpeedButton_FindComponent(s.instance, AName ))
}

func (s *TSpeedButton) GetNamePath() string {
    defer exceptionProc()
    return SpeedButton_GetNamePath(s.instance)
}

func (s *TSpeedButton) Assign(Source IObject) {
    defer exceptionProc()
    SpeedButton_Assign(s.instance, CheckPtr(Source))
}

func (s *TSpeedButton) ClassName() string {
    defer exceptionProc()
    return SpeedButton_ClassName(s.instance)
}

func (s *TSpeedButton) Equals(Obj IObject) bool {
    defer exceptionProc()
    return SpeedButton_Equals(s.instance, CheckPtr(Obj))
}

func (s *TSpeedButton) GetHashCode() int32 {
    defer exceptionProc()
    return SpeedButton_GetHashCode(s.instance)
}

func (s *TSpeedButton) ToString() string {
    defer exceptionProc()
    return SpeedButton_ToString(s.instance)
}

func (s *TSpeedButton) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(SpeedButton_GetAction(s.instance))
}

func (s *TSpeedButton) SetAction(value IComponent) {
    defer exceptionProc()
    SpeedButton_SetAction(s.instance, CheckPtr(value))
}

func (s *TSpeedButton) Align() TAlign {
    defer exceptionProc()
    return SpeedButton_GetAlign(s.instance)
}

func (s *TSpeedButton) SetAlign(value TAlign) {
    defer exceptionProc()
    SpeedButton_SetAlign(s.instance, value)
}

func (s *TSpeedButton) AllowAllUp() bool {
    defer exceptionProc()
    return SpeedButton_GetAllowAllUp(s.instance)
}

func (s *TSpeedButton) SetAllowAllUp(value bool) {
    defer exceptionProc()
    SpeedButton_SetAllowAllUp(s.instance, value)
}

func (s *TSpeedButton) Anchors() TAnchors {
    defer exceptionProc()
    return SpeedButton_GetAnchors(s.instance)
}

func (s *TSpeedButton) SetAnchors(value TAnchors) {
    defer exceptionProc()
    SpeedButton_SetAnchors(s.instance, value)
}

func (s *TSpeedButton) GroupIndex() int32 {
    defer exceptionProc()
    return SpeedButton_GetGroupIndex(s.instance)
}

func (s *TSpeedButton) SetGroupIndex(value int32) {
    defer exceptionProc()
    SpeedButton_SetGroupIndex(s.instance, value)
}

func (s *TSpeedButton) Down() bool {
    defer exceptionProc()
    return SpeedButton_GetDown(s.instance)
}

func (s *TSpeedButton) SetDown(value bool) {
    defer exceptionProc()
    SpeedButton_SetDown(s.instance, value)
}

func (s *TSpeedButton) Caption() string {
    defer exceptionProc()
    return SpeedButton_GetCaption(s.instance)
}

func (s *TSpeedButton) SetCaption(value string) {
    defer exceptionProc()
    SpeedButton_SetCaption(s.instance, value)
}

func (s *TSpeedButton) Enabled() bool {
    defer exceptionProc()
    return SpeedButton_GetEnabled(s.instance)
}

func (s *TSpeedButton) SetEnabled(value bool) {
    defer exceptionProc()
    SpeedButton_SetEnabled(s.instance, value)
}

func (s *TSpeedButton) Flat() bool {
    defer exceptionProc()
    return SpeedButton_GetFlat(s.instance)
}

func (s *TSpeedButton) SetFlat(value bool) {
    defer exceptionProc()
    SpeedButton_SetFlat(s.instance, value)
}

func (s *TSpeedButton) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(SpeedButton_GetFont(s.instance))
}

func (s *TSpeedButton) SetFont(value *TFont) {
    defer exceptionProc()
    SpeedButton_SetFont(s.instance, CheckPtr(value))
}

func (s *TSpeedButton) Layout() TButtonLayout {
    defer exceptionProc()
    return SpeedButton_GetLayout(s.instance)
}

func (s *TSpeedButton) SetLayout(value TButtonLayout) {
    defer exceptionProc()
    SpeedButton_SetLayout(s.instance, value)
}

func (s *TSpeedButton) ParentFont() bool {
    defer exceptionProc()
    return SpeedButton_GetParentFont(s.instance)
}

func (s *TSpeedButton) SetParentFont(value bool) {
    defer exceptionProc()
    SpeedButton_SetParentFont(s.instance, value)
}

func (s *TSpeedButton) ParentShowHint() bool {
    defer exceptionProc()
    return SpeedButton_GetParentShowHint(s.instance)
}

func (s *TSpeedButton) SetParentShowHint(value bool) {
    defer exceptionProc()
    SpeedButton_SetParentShowHint(s.instance, value)
}

func (s *TSpeedButton) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(SpeedButton_GetPopupMenu(s.instance))
}

func (s *TSpeedButton) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    SpeedButton_SetPopupMenu(s.instance, CheckPtr(value))
}

func (s *TSpeedButton) ShowHint() bool {
    defer exceptionProc()
    return SpeedButton_GetShowHint(s.instance)
}

func (s *TSpeedButton) SetShowHint(value bool) {
    defer exceptionProc()
    SpeedButton_SetShowHint(s.instance, value)
}

func (s *TSpeedButton) Transparent() bool {
    defer exceptionProc()
    return SpeedButton_GetTransparent(s.instance)
}

func (s *TSpeedButton) SetTransparent(value bool) {
    defer exceptionProc()
    SpeedButton_SetTransparent(s.instance, value)
}

func (s *TSpeedButton) Visible() bool {
    defer exceptionProc()
    return SpeedButton_GetVisible(s.instance)
}

func (s *TSpeedButton) SetVisible(value bool) {
    defer exceptionProc()
    SpeedButton_SetVisible(s.instance, value)
}

func (s *TSpeedButton) SetOnClick(fn TNotifyEvent) {
    SpeedButton_SetOnClick(s.instance, fn)
}

func (s *TSpeedButton) SetOnDblClick(fn TNotifyEvent) {
    SpeedButton_SetOnDblClick(s.instance, fn)
}

func (s *TSpeedButton) SetOnMouseDown(fn TMouseEvent) {
    SpeedButton_SetOnMouseDown(s.instance, fn)
}

func (s *TSpeedButton) SetOnMouseEnter(fn TNotifyEvent) {
    SpeedButton_SetOnMouseEnter(s.instance, fn)
}

func (s *TSpeedButton) SetOnMouseLeave(fn TNotifyEvent) {
    SpeedButton_SetOnMouseLeave(s.instance, fn)
}

func (s *TSpeedButton) SetOnMouseMove(fn TMouseMoveEvent) {
    SpeedButton_SetOnMouseMove(s.instance, fn)
}

func (s *TSpeedButton) SetOnMouseUp(fn TMouseEvent) {
    SpeedButton_SetOnMouseUp(s.instance, fn)
}

func (s *TSpeedButton) BoundsRect() TRect {
    defer exceptionProc()
    return SpeedButton_GetBoundsRect(s.instance)
}

func (s *TSpeedButton) SetBoundsRect(value TRect) {
    defer exceptionProc()
    SpeedButton_SetBoundsRect(s.instance, value)
}

func (s *TSpeedButton) ClientHeight() int32 {
    defer exceptionProc()
    return SpeedButton_GetClientHeight(s.instance)
}

func (s *TSpeedButton) SetClientHeight(value int32) {
    defer exceptionProc()
    SpeedButton_SetClientHeight(s.instance, value)
}

func (s *TSpeedButton) ClientRect() TRect {
    defer exceptionProc()
    return SpeedButton_GetClientRect(s.instance)
}

func (s *TSpeedButton) ClientWidth() int32 {
    defer exceptionProc()
    return SpeedButton_GetClientWidth(s.instance)
}

func (s *TSpeedButton) SetClientWidth(value int32) {
    defer exceptionProc()
    SpeedButton_SetClientWidth(s.instance, value)
}

func (s *TSpeedButton) ExplicitLeft() int32 {
    defer exceptionProc()
    return SpeedButton_GetExplicitLeft(s.instance)
}

func (s *TSpeedButton) ExplicitTop() int32 {
    defer exceptionProc()
    return SpeedButton_GetExplicitTop(s.instance)
}

func (s *TSpeedButton) ExplicitWidth() int32 {
    defer exceptionProc()
    return SpeedButton_GetExplicitWidth(s.instance)
}

func (s *TSpeedButton) ExplicitHeight() int32 {
    defer exceptionProc()
    return SpeedButton_GetExplicitHeight(s.instance)
}

func (s *TSpeedButton) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(SpeedButton_GetParent(s.instance))
}

func (s *TSpeedButton) SetParent(value IControl) {
    defer exceptionProc()
    SpeedButton_SetParent(s.instance, CheckPtr(value))
}

func (s *TSpeedButton) AlignWithMargins() bool {
    defer exceptionProc()
    return SpeedButton_GetAlignWithMargins(s.instance)
}

func (s *TSpeedButton) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    SpeedButton_SetAlignWithMargins(s.instance, value)
}

func (s *TSpeedButton) Left() int32 {
    defer exceptionProc()
    return SpeedButton_GetLeft(s.instance)
}

func (s *TSpeedButton) SetLeft(value int32) {
    defer exceptionProc()
    SpeedButton_SetLeft(s.instance, value)
}

func (s *TSpeedButton) Top() int32 {
    defer exceptionProc()
    return SpeedButton_GetTop(s.instance)
}

func (s *TSpeedButton) SetTop(value int32) {
    defer exceptionProc()
    SpeedButton_SetTop(s.instance, value)
}

func (s *TSpeedButton) Width() int32 {
    defer exceptionProc()
    return SpeedButton_GetWidth(s.instance)
}

func (s *TSpeedButton) SetWidth(value int32) {
    defer exceptionProc()
    SpeedButton_SetWidth(s.instance, value)
}

func (s *TSpeedButton) Height() int32 {
    defer exceptionProc()
    return SpeedButton_GetHeight(s.instance)
}

func (s *TSpeedButton) SetHeight(value int32) {
    defer exceptionProc()
    SpeedButton_SetHeight(s.instance, value)
}

func (s *TSpeedButton) Cursor() TCursor {
    defer exceptionProc()
    return SpeedButton_GetCursor(s.instance)
}

func (s *TSpeedButton) SetCursor(value TCursor) {
    defer exceptionProc()
    SpeedButton_SetCursor(s.instance, value)
}

func (s *TSpeedButton) Hint() string {
    defer exceptionProc()
    return SpeedButton_GetHint(s.instance)
}

func (s *TSpeedButton) SetHint(value string) {
    defer exceptionProc()
    SpeedButton_SetHint(s.instance, value)
}

func (s *TSpeedButton) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(SpeedButton_GetMargins(s.instance))
}

func (s *TSpeedButton) SetMargins(value *TMargins) {
    defer exceptionProc()
    SpeedButton_SetMargins(s.instance, CheckPtr(value))
}

func (s *TSpeedButton) ComponentCount() int32 {
    defer exceptionProc()
    return SpeedButton_GetComponentCount(s.instance)
}

func (s *TSpeedButton) ComponentIndex() int32 {
    defer exceptionProc()
    return SpeedButton_GetComponentIndex(s.instance)
}

func (s *TSpeedButton) SetComponentIndex(value int32) {
    defer exceptionProc()
    SpeedButton_SetComponentIndex(s.instance, value)
}

func (s *TSpeedButton) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(SpeedButton_GetOwner(s.instance))
}

func (s *TSpeedButton) Name() string {
    defer exceptionProc()
    return SpeedButton_GetName(s.instance)
}

func (s *TSpeedButton) SetName(value string) {
    defer exceptionProc()
    SpeedButton_SetName(s.instance, value)
}

func (s *TSpeedButton) Tag() int {
    defer exceptionProc()
    return SpeedButton_GetTag(s.instance)
}

func (s *TSpeedButton) SetTag(value int) {
    defer exceptionProc()
    SpeedButton_SetTag(s.instance, value)
}

func (s *TSpeedButton) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(SpeedButton_GetComponents(s.instance, AIndex))
}

