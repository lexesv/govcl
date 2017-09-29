
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TLinkLabel struct {
    IControl
    instance uintptr
}

func NewLinkLabel(owner IComponent) *TLinkLabel {
    l := new(TLinkLabel)
    l.instance = LinkLabel_Create(owner.Instance())
    return l
}

func LinkLabelFromInst(inst uintptr) *TLinkLabel {
    l := new(TLinkLabel)
    l.instance = inst
    return l
}

func LinkLabelFromObj(obj IObject) *TLinkLabel {
    l := new(TLinkLabel)
    l.instance = CheckPtr(obj)
    return l
}

func (l *TLinkLabel) Free() {
    if l.instance != 0 {
        LinkLabel_Free(l.instance)
    }
}

func (l *TLinkLabel) Instance() uintptr {
    return l.instance
}

func (l *TLinkLabel) IsValid() bool {
    return l.instance != 0
}

func (l *TLinkLabel) CanFocus() bool {
    defer exceptionProc()
    return LinkLabel_CanFocus(l.instance)
}

func (l *TLinkLabel) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    LinkLabel_FlipChildren(l.instance, AllLevels )
}

func (l *TLinkLabel) Focused() bool {
    defer exceptionProc()
    return LinkLabel_Focused(l.instance)
}

func (l *TLinkLabel) HandleAllocated() bool {
    defer exceptionProc()
    return LinkLabel_HandleAllocated(l.instance)
}

func (l *TLinkLabel) Invalidate() {
    defer exceptionProc()
    LinkLabel_Invalidate(l.instance)
}

func (l *TLinkLabel) Realign() {
    defer exceptionProc()
    LinkLabel_Realign(l.instance)
}

func (l *TLinkLabel) Repaint() {
    defer exceptionProc()
    LinkLabel_Repaint(l.instance)
}

func (l *TLinkLabel) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    LinkLabel_ScaleBy(l.instance, M , D )
}

func (l *TLinkLabel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    LinkLabel_SetBounds(l.instance, ALeft , ATop , AWidth , AHeight )
}

func (l *TLinkLabel) SetFocus() {
    defer exceptionProc()
    LinkLabel_SetFocus(l.instance)
}

func (l *TLinkLabel) Update() {
    defer exceptionProc()
    LinkLabel_Update(l.instance)
}

func (l *TLinkLabel) BringToFront() {
    defer exceptionProc()
    LinkLabel_BringToFront(l.instance)
}

func (l *TLinkLabel) HasParent() bool {
    defer exceptionProc()
    return LinkLabel_HasParent(l.instance)
}

func (l *TLinkLabel) Hide() {
    defer exceptionProc()
    LinkLabel_Hide(l.instance)
}

func (l *TLinkLabel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return LinkLabel_Perform(l.instance, Msg , WParam , LParam )
}

func (l *TLinkLabel) Refresh() {
    defer exceptionProc()
    LinkLabel_Refresh(l.instance)
}

func (l *TLinkLabel) SendToBack() {
    defer exceptionProc()
    LinkLabel_SendToBack(l.instance)
}

func (l *TLinkLabel) Show() {
    defer exceptionProc()
    LinkLabel_Show(l.instance)
}

func (l *TLinkLabel) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return LinkLabel_GetTextBuf(l.instance, Buffer , BufSize )
}

func (l *TLinkLabel) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(LinkLabel_FindComponent(l.instance, AName ))
}

func (l *TLinkLabel) GetNamePath() string {
    defer exceptionProc()
    return LinkLabel_GetNamePath(l.instance)
}

func (l *TLinkLabel) Assign(Source IObject) {
    defer exceptionProc()
    LinkLabel_Assign(l.instance, CheckPtr(Source))
}

func (l *TLinkLabel) ClassName() string {
    defer exceptionProc()
    return LinkLabel_ClassName(l.instance)
}

func (l *TLinkLabel) Equals(Obj IObject) bool {
    defer exceptionProc()
    return LinkLabel_Equals(l.instance, CheckPtr(Obj))
}

func (l *TLinkLabel) GetHashCode() int32 {
    defer exceptionProc()
    return LinkLabel_GetHashCode(l.instance)
}

func (l *TLinkLabel) ToString() string {
    defer exceptionProc()
    return LinkLabel_ToString(l.instance)
}

func (l *TLinkLabel) Align() TAlign {
    defer exceptionProc()
    return LinkLabel_GetAlign(l.instance)
}

func (l *TLinkLabel) SetAlign(value TAlign) {
    defer exceptionProc()
    LinkLabel_SetAlign(l.instance, value)
}

func (l *TLinkLabel) Alignment() TLinkAlignment {
    defer exceptionProc()
    return LinkLabel_GetAlignment(l.instance)
}

func (l *TLinkLabel) SetAlignment(value TLinkAlignment) {
    defer exceptionProc()
    LinkLabel_SetAlignment(l.instance, value)
}

func (l *TLinkLabel) Anchors() TAnchors {
    defer exceptionProc()
    return LinkLabel_GetAnchors(l.instance)
}

func (l *TLinkLabel) SetAnchors(value TAnchors) {
    defer exceptionProc()
    LinkLabel_SetAnchors(l.instance, value)
}

func (l *TLinkLabel) AutoSize() bool {
    defer exceptionProc()
    return LinkLabel_GetAutoSize(l.instance)
}

func (l *TLinkLabel) SetAutoSize(value bool) {
    defer exceptionProc()
    LinkLabel_SetAutoSize(l.instance, value)
}

func (l *TLinkLabel) Caption() string {
    defer exceptionProc()
    return LinkLabel_GetCaption(l.instance)
}

func (l *TLinkLabel) SetCaption(value string) {
    defer exceptionProc()
    LinkLabel_SetCaption(l.instance, value)
}

func (l *TLinkLabel) Color() TColor {
    defer exceptionProc()
    return LinkLabel_GetColor(l.instance)
}

func (l *TLinkLabel) SetColor(value TColor) {
    defer exceptionProc()
    LinkLabel_SetColor(l.instance, value)
}

func (l *TLinkLabel) Enabled() bool {
    defer exceptionProc()
    return LinkLabel_GetEnabled(l.instance)
}

func (l *TLinkLabel) SetEnabled(value bool) {
    defer exceptionProc()
    LinkLabel_SetEnabled(l.instance, value)
}

func (l *TLinkLabel) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(LinkLabel_GetFont(l.instance))
}

func (l *TLinkLabel) SetFont(value *TFont) {
    defer exceptionProc()
    LinkLabel_SetFont(l.instance, CheckPtr(value))
}

func (l *TLinkLabel) ParentColor() bool {
    defer exceptionProc()
    return LinkLabel_GetParentColor(l.instance)
}

func (l *TLinkLabel) SetParentColor(value bool) {
    defer exceptionProc()
    LinkLabel_SetParentColor(l.instance, value)
}

func (l *TLinkLabel) ParentFont() bool {
    defer exceptionProc()
    return LinkLabel_GetParentFont(l.instance)
}

func (l *TLinkLabel) SetParentFont(value bool) {
    defer exceptionProc()
    LinkLabel_SetParentFont(l.instance, value)
}

func (l *TLinkLabel) ParentShowHint() bool {
    defer exceptionProc()
    return LinkLabel_GetParentShowHint(l.instance)
}

func (l *TLinkLabel) SetParentShowHint(value bool) {
    defer exceptionProc()
    LinkLabel_SetParentShowHint(l.instance, value)
}

func (l *TLinkLabel) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(LinkLabel_GetPopupMenu(l.instance))
}

func (l *TLinkLabel) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    LinkLabel_SetPopupMenu(l.instance, CheckPtr(value))
}

func (l *TLinkLabel) ShowHint() bool {
    defer exceptionProc()
    return LinkLabel_GetShowHint(l.instance)
}

func (l *TLinkLabel) SetShowHint(value bool) {
    defer exceptionProc()
    LinkLabel_SetShowHint(l.instance, value)
}

func (l *TLinkLabel) TabOrder() int16 {
    defer exceptionProc()
    return LinkLabel_GetTabOrder(l.instance)
}

func (l *TLinkLabel) SetTabOrder(value int16) {
    defer exceptionProc()
    LinkLabel_SetTabOrder(l.instance, value)
}

func (l *TLinkLabel) TabStop() bool {
    defer exceptionProc()
    return LinkLabel_GetTabStop(l.instance)
}

func (l *TLinkLabel) SetTabStop(value bool) {
    defer exceptionProc()
    LinkLabel_SetTabStop(l.instance, value)
}

func (l *TLinkLabel) UseVisualStyle() bool {
    defer exceptionProc()
    return LinkLabel_GetUseVisualStyle(l.instance)
}

func (l *TLinkLabel) SetUseVisualStyle(value bool) {
    defer exceptionProc()
    LinkLabel_SetUseVisualStyle(l.instance, value)
}

func (l *TLinkLabel) Visible() bool {
    defer exceptionProc()
    return LinkLabel_GetVisible(l.instance)
}

func (l *TLinkLabel) SetVisible(value bool) {
    defer exceptionProc()
    LinkLabel_SetVisible(l.instance, value)
}

func (l *TLinkLabel) SetOnClick(fn TNotifyEvent) {
    LinkLabel_SetOnClick(l.instance, fn)
}

func (l *TLinkLabel) SetOnDblClick(fn TNotifyEvent) {
    LinkLabel_SetOnDblClick(l.instance, fn)
}

func (l *TLinkLabel) SetOnMouseDown(fn TMouseEvent) {
    LinkLabel_SetOnMouseDown(l.instance, fn)
}

func (l *TLinkLabel) SetOnMouseEnter(fn TNotifyEvent) {
    LinkLabel_SetOnMouseEnter(l.instance, fn)
}

func (l *TLinkLabel) SetOnMouseLeave(fn TNotifyEvent) {
    LinkLabel_SetOnMouseLeave(l.instance, fn)
}

func (l *TLinkLabel) SetOnMouseMove(fn TMouseMoveEvent) {
    LinkLabel_SetOnMouseMove(l.instance, fn)
}

func (l *TLinkLabel) SetOnMouseUp(fn TMouseEvent) {
    LinkLabel_SetOnMouseUp(l.instance, fn)
}

func (l *TLinkLabel) SetOnLinkClick(fn TSysLinkEvent) {
    LinkLabel_SetOnLinkClick(l.instance, fn)
}

func (l *TLinkLabel) DoubleBuffered() bool {
    defer exceptionProc()
    return LinkLabel_GetDoubleBuffered(l.instance)
}

func (l *TLinkLabel) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    LinkLabel_SetDoubleBuffered(l.instance, value)
}

func (l *TLinkLabel) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(LinkLabel_GetBrush(l.instance))
}

func (l *TLinkLabel) ControlCount() int32 {
    defer exceptionProc()
    return LinkLabel_GetControlCount(l.instance)
}

func (l *TLinkLabel) Handle() HWND {
    defer exceptionProc()
    return LinkLabel_GetHandle(l.instance)
}

func (l *TLinkLabel) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(LinkLabel_GetAction(l.instance))
}

func (l *TLinkLabel) SetAction(value IComponent) {
    defer exceptionProc()
    LinkLabel_SetAction(l.instance, CheckPtr(value))
}

func (l *TLinkLabel) BoundsRect() TRect {
    defer exceptionProc()
    return LinkLabel_GetBoundsRect(l.instance)
}

func (l *TLinkLabel) SetBoundsRect(value TRect) {
    defer exceptionProc()
    LinkLabel_SetBoundsRect(l.instance, value)
}

func (l *TLinkLabel) ClientHeight() int32 {
    defer exceptionProc()
    return LinkLabel_GetClientHeight(l.instance)
}

func (l *TLinkLabel) SetClientHeight(value int32) {
    defer exceptionProc()
    LinkLabel_SetClientHeight(l.instance, value)
}

func (l *TLinkLabel) ClientRect() TRect {
    defer exceptionProc()
    return LinkLabel_GetClientRect(l.instance)
}

func (l *TLinkLabel) ClientWidth() int32 {
    defer exceptionProc()
    return LinkLabel_GetClientWidth(l.instance)
}

func (l *TLinkLabel) SetClientWidth(value int32) {
    defer exceptionProc()
    LinkLabel_SetClientWidth(l.instance, value)
}

func (l *TLinkLabel) ExplicitLeft() int32 {
    defer exceptionProc()
    return LinkLabel_GetExplicitLeft(l.instance)
}

func (l *TLinkLabel) ExplicitTop() int32 {
    defer exceptionProc()
    return LinkLabel_GetExplicitTop(l.instance)
}

func (l *TLinkLabel) ExplicitWidth() int32 {
    defer exceptionProc()
    return LinkLabel_GetExplicitWidth(l.instance)
}

func (l *TLinkLabel) ExplicitHeight() int32 {
    defer exceptionProc()
    return LinkLabel_GetExplicitHeight(l.instance)
}

func (l *TLinkLabel) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(LinkLabel_GetParent(l.instance))
}

func (l *TLinkLabel) SetParent(value IControl) {
    defer exceptionProc()
    LinkLabel_SetParent(l.instance, CheckPtr(value))
}

func (l *TLinkLabel) AlignWithMargins() bool {
    defer exceptionProc()
    return LinkLabel_GetAlignWithMargins(l.instance)
}

func (l *TLinkLabel) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    LinkLabel_SetAlignWithMargins(l.instance, value)
}

func (l *TLinkLabel) Left() int32 {
    defer exceptionProc()
    return LinkLabel_GetLeft(l.instance)
}

func (l *TLinkLabel) SetLeft(value int32) {
    defer exceptionProc()
    LinkLabel_SetLeft(l.instance, value)
}

func (l *TLinkLabel) Top() int32 {
    defer exceptionProc()
    return LinkLabel_GetTop(l.instance)
}

func (l *TLinkLabel) SetTop(value int32) {
    defer exceptionProc()
    LinkLabel_SetTop(l.instance, value)
}

func (l *TLinkLabel) Width() int32 {
    defer exceptionProc()
    return LinkLabel_GetWidth(l.instance)
}

func (l *TLinkLabel) SetWidth(value int32) {
    defer exceptionProc()
    LinkLabel_SetWidth(l.instance, value)
}

func (l *TLinkLabel) Height() int32 {
    defer exceptionProc()
    return LinkLabel_GetHeight(l.instance)
}

func (l *TLinkLabel) SetHeight(value int32) {
    defer exceptionProc()
    LinkLabel_SetHeight(l.instance, value)
}

func (l *TLinkLabel) Cursor() TCursor {
    defer exceptionProc()
    return LinkLabel_GetCursor(l.instance)
}

func (l *TLinkLabel) SetCursor(value TCursor) {
    defer exceptionProc()
    LinkLabel_SetCursor(l.instance, value)
}

func (l *TLinkLabel) Hint() string {
    defer exceptionProc()
    return LinkLabel_GetHint(l.instance)
}

func (l *TLinkLabel) SetHint(value string) {
    defer exceptionProc()
    LinkLabel_SetHint(l.instance, value)
}

func (l *TLinkLabel) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(LinkLabel_GetMargins(l.instance))
}

func (l *TLinkLabel) SetMargins(value *TMargins) {
    defer exceptionProc()
    LinkLabel_SetMargins(l.instance, CheckPtr(value))
}

func (l *TLinkLabel) ComponentCount() int32 {
    defer exceptionProc()
    return LinkLabel_GetComponentCount(l.instance)
}

func (l *TLinkLabel) ComponentIndex() int32 {
    defer exceptionProc()
    return LinkLabel_GetComponentIndex(l.instance)
}

func (l *TLinkLabel) SetComponentIndex(value int32) {
    defer exceptionProc()
    LinkLabel_SetComponentIndex(l.instance, value)
}

func (l *TLinkLabel) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(LinkLabel_GetOwner(l.instance))
}

func (l *TLinkLabel) Name() string {
    defer exceptionProc()
    return LinkLabel_GetName(l.instance)
}

func (l *TLinkLabel) SetName(value string) {
    defer exceptionProc()
    LinkLabel_SetName(l.instance, value)
}

func (l *TLinkLabel) Tag() int {
    defer exceptionProc()
    return LinkLabel_GetTag(l.instance)
}

func (l *TLinkLabel) SetTag(value int) {
    defer exceptionProc()
    LinkLabel_SetTag(l.instance, value)
}

func (l *TLinkLabel) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(LinkLabel_GetControls(l.instance, Index))
}

func (l *TLinkLabel) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(LinkLabel_GetComponents(l.instance, AIndex))
}

