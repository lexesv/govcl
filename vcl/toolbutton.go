
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

type TToolButton struct {
    IControl
    instance uintptr
}

func NewToolButton(owner IComponent) *TToolButton {
    t := new(TToolButton)
    t.instance = ToolButton_Create(owner.Instance())
    return t
}

func ToolButtonFromInst(inst uintptr) *TToolButton {
    t := new(TToolButton)
    t.instance = inst
    return t
}

func ToolButtonFromObj(obj IObject) *TToolButton {
    t := new(TToolButton)
    t.instance = CheckPtr(obj)
    return t
}

func (t *TToolButton) Free() {
    if t.instance != 0 {
        ToolButton_Free(t.instance)
    }
}

func (t *TToolButton) Instance() uintptr {
    return t.instance
}

func (t *TToolButton) IsValid() bool {
    return t.instance != 0
}

func (t *TToolButton) CheckMenuDropdown() bool {
    defer exceptionProc()
    return ToolButton_CheckMenuDropdown(t.instance)
}

func (t *TToolButton) Click() {
    defer exceptionProc()
    ToolButton_Click(t.instance)
}

func (t *TToolButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    ToolButton_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight )
}

func (t *TToolButton) BringToFront() {
    defer exceptionProc()
    ToolButton_BringToFront(t.instance)
}

func (t *TToolButton) HasParent() bool {
    defer exceptionProc()
    return ToolButton_HasParent(t.instance)
}

func (t *TToolButton) Hide() {
    defer exceptionProc()
    ToolButton_Hide(t.instance)
}

func (t *TToolButton) Invalidate() {
    defer exceptionProc()
    ToolButton_Invalidate(t.instance)
}

func (t *TToolButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return ToolButton_Perform(t.instance, Msg , WParam , LParam )
}

func (t *TToolButton) Refresh() {
    defer exceptionProc()
    ToolButton_Refresh(t.instance)
}

func (t *TToolButton) Repaint() {
    defer exceptionProc()
    ToolButton_Repaint(t.instance)
}

func (t *TToolButton) SendToBack() {
    defer exceptionProc()
    ToolButton_SendToBack(t.instance)
}

func (t *TToolButton) Show() {
    defer exceptionProc()
    ToolButton_Show(t.instance)
}

func (t *TToolButton) Update() {
    defer exceptionProc()
    ToolButton_Update(t.instance)
}

func (t *TToolButton) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return ToolButton_GetTextBuf(t.instance, Buffer , BufSize )
}

func (t *TToolButton) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ToolButton_FindComponent(t.instance, AName ))
}

func (t *TToolButton) GetNamePath() string {
    defer exceptionProc()
    return ToolButton_GetNamePath(t.instance)
}

func (t *TToolButton) Assign(Source IObject) {
    defer exceptionProc()
    ToolButton_Assign(t.instance, CheckPtr(Source))
}

func (t *TToolButton) ClassName() string {
    defer exceptionProc()
    return ToolButton_ClassName(t.instance)
}

func (t *TToolButton) Equals(Obj IObject) bool {
    defer exceptionProc()
    return ToolButton_Equals(t.instance, CheckPtr(Obj))
}

func (t *TToolButton) GetHashCode() int32 {
    defer exceptionProc()
    return ToolButton_GetHashCode(t.instance)
}

func (t *TToolButton) ToString() string {
    defer exceptionProc()
    return ToolButton_ToString(t.instance)
}

func (t *TToolButton) Index() int32 {
    defer exceptionProc()
    return ToolButton_GetIndex(t.instance)
}

func (t *TToolButton) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(ToolButton_GetAction(t.instance))
}

func (t *TToolButton) SetAction(value IComponent) {
    defer exceptionProc()
    ToolButton_SetAction(t.instance, CheckPtr(value))
}

func (t *TToolButton) AllowAllUp() bool {
    defer exceptionProc()
    return ToolButton_GetAllowAllUp(t.instance)
}

func (t *TToolButton) SetAllowAllUp(value bool) {
    defer exceptionProc()
    ToolButton_SetAllowAllUp(t.instance, value)
}

func (t *TToolButton) AutoSize() bool {
    defer exceptionProc()
    return ToolButton_GetAutoSize(t.instance)
}

func (t *TToolButton) SetAutoSize(value bool) {
    defer exceptionProc()
    ToolButton_SetAutoSize(t.instance, value)
}

func (t *TToolButton) Caption() string {
    defer exceptionProc()
    return ToolButton_GetCaption(t.instance)
}

func (t *TToolButton) SetCaption(value string) {
    defer exceptionProc()
    ToolButton_SetCaption(t.instance, value)
}

func (t *TToolButton) Down() bool {
    defer exceptionProc()
    return ToolButton_GetDown(t.instance)
}

func (t *TToolButton) SetDown(value bool) {
    defer exceptionProc()
    ToolButton_SetDown(t.instance, value)
}

func (t *TToolButton) DropdownMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(ToolButton_GetDropdownMenu(t.instance))
}

func (t *TToolButton) SetDropdownMenu(value IComponent) {
    defer exceptionProc()
    ToolButton_SetDropdownMenu(t.instance, CheckPtr(value))
}

func (t *TToolButton) Enabled() bool {
    defer exceptionProc()
    return ToolButton_GetEnabled(t.instance)
}

func (t *TToolButton) SetEnabled(value bool) {
    defer exceptionProc()
    ToolButton_SetEnabled(t.instance, value)
}

func (t *TToolButton) EnableDropdown() bool {
    defer exceptionProc()
    return ToolButton_GetEnableDropdown(t.instance)
}

func (t *TToolButton) SetEnableDropdown(value bool) {
    defer exceptionProc()
    ToolButton_SetEnableDropdown(t.instance, value)
}

func (t *TToolButton) Grouped() bool {
    defer exceptionProc()
    return ToolButton_GetGrouped(t.instance)
}

func (t *TToolButton) SetGrouped(value bool) {
    defer exceptionProc()
    ToolButton_SetGrouped(t.instance, value)
}

func (t *TToolButton) Height() int32 {
    defer exceptionProc()
    return ToolButton_GetHeight(t.instance)
}

func (t *TToolButton) SetHeight(value int32) {
    defer exceptionProc()
    ToolButton_SetHeight(t.instance, value)
}

func (t *TToolButton) ImageIndex() int32 {
    defer exceptionProc()
    return ToolButton_GetImageIndex(t.instance)
}

func (t *TToolButton) SetImageIndex(value int32) {
    defer exceptionProc()
    ToolButton_SetImageIndex(t.instance, value)
}

func (t *TToolButton) Indeterminate() bool {
    defer exceptionProc()
    return ToolButton_GetIndeterminate(t.instance)
}

func (t *TToolButton) SetIndeterminate(value bool) {
    defer exceptionProc()
    ToolButton_SetIndeterminate(t.instance, value)
}

func (t *TToolButton) Marked() bool {
    defer exceptionProc()
    return ToolButton_GetMarked(t.instance)
}

func (t *TToolButton) SetMarked(value bool) {
    defer exceptionProc()
    ToolButton_SetMarked(t.instance, value)
}

func (t *TToolButton) ParentShowHint() bool {
    defer exceptionProc()
    return ToolButton_GetParentShowHint(t.instance)
}

func (t *TToolButton) SetParentShowHint(value bool) {
    defer exceptionProc()
    ToolButton_SetParentShowHint(t.instance, value)
}

func (t *TToolButton) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(ToolButton_GetPopupMenu(t.instance))
}

func (t *TToolButton) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    ToolButton_SetPopupMenu(t.instance, CheckPtr(value))
}

func (t *TToolButton) Wrap() bool {
    defer exceptionProc()
    return ToolButton_GetWrap(t.instance)
}

func (t *TToolButton) SetWrap(value bool) {
    defer exceptionProc()
    ToolButton_SetWrap(t.instance, value)
}

func (t *TToolButton) ShowHint() bool {
    defer exceptionProc()
    return ToolButton_GetShowHint(t.instance)
}

func (t *TToolButton) SetShowHint(value bool) {
    defer exceptionProc()
    ToolButton_SetShowHint(t.instance, value)
}

func (t *TToolButton) Style() TToolButtonStyle {
    defer exceptionProc()
    return ToolButton_GetStyle(t.instance)
}

func (t *TToolButton) SetStyle(value TToolButtonStyle) {
    defer exceptionProc()
    ToolButton_SetStyle(t.instance, value)
}

func (t *TToolButton) Visible() bool {
    defer exceptionProc()
    return ToolButton_GetVisible(t.instance)
}

func (t *TToolButton) SetVisible(value bool) {
    defer exceptionProc()
    ToolButton_SetVisible(t.instance, value)
}

func (t *TToolButton) Width() int32 {
    defer exceptionProc()
    return ToolButton_GetWidth(t.instance)
}

func (t *TToolButton) SetWidth(value int32) {
    defer exceptionProc()
    ToolButton_SetWidth(t.instance, value)
}

func (t *TToolButton) SetOnClick(fn TNotifyEvent) {
    ToolButton_SetOnClick(t.instance, fn)
}

func (t *TToolButton) SetOnMouseDown(fn TMouseEvent) {
    ToolButton_SetOnMouseDown(t.instance, fn)
}

func (t *TToolButton) SetOnMouseEnter(fn TNotifyEvent) {
    ToolButton_SetOnMouseEnter(t.instance, fn)
}

func (t *TToolButton) SetOnMouseLeave(fn TNotifyEvent) {
    ToolButton_SetOnMouseLeave(t.instance, fn)
}

func (t *TToolButton) SetOnMouseMove(fn TMouseMoveEvent) {
    ToolButton_SetOnMouseMove(t.instance, fn)
}

func (t *TToolButton) SetOnMouseUp(fn TMouseEvent) {
    ToolButton_SetOnMouseUp(t.instance, fn)
}

func (t *TToolButton) Align() TAlign {
    defer exceptionProc()
    return ToolButton_GetAlign(t.instance)
}

func (t *TToolButton) SetAlign(value TAlign) {
    defer exceptionProc()
    ToolButton_SetAlign(t.instance, value)
}

func (t *TToolButton) Anchors() TAnchors {
    defer exceptionProc()
    return ToolButton_GetAnchors(t.instance)
}

func (t *TToolButton) SetAnchors(value TAnchors) {
    defer exceptionProc()
    ToolButton_SetAnchors(t.instance, value)
}

func (t *TToolButton) BoundsRect() TRect {
    defer exceptionProc()
    return ToolButton_GetBoundsRect(t.instance)
}

func (t *TToolButton) SetBoundsRect(value TRect) {
    defer exceptionProc()
    ToolButton_SetBoundsRect(t.instance, value)
}

func (t *TToolButton) ClientHeight() int32 {
    defer exceptionProc()
    return ToolButton_GetClientHeight(t.instance)
}

func (t *TToolButton) SetClientHeight(value int32) {
    defer exceptionProc()
    ToolButton_SetClientHeight(t.instance, value)
}

func (t *TToolButton) ClientRect() TRect {
    defer exceptionProc()
    return ToolButton_GetClientRect(t.instance)
}

func (t *TToolButton) ClientWidth() int32 {
    defer exceptionProc()
    return ToolButton_GetClientWidth(t.instance)
}

func (t *TToolButton) SetClientWidth(value int32) {
    defer exceptionProc()
    ToolButton_SetClientWidth(t.instance, value)
}

func (t *TToolButton) ExplicitLeft() int32 {
    defer exceptionProc()
    return ToolButton_GetExplicitLeft(t.instance)
}

func (t *TToolButton) ExplicitTop() int32 {
    defer exceptionProc()
    return ToolButton_GetExplicitTop(t.instance)
}

func (t *TToolButton) ExplicitWidth() int32 {
    defer exceptionProc()
    return ToolButton_GetExplicitWidth(t.instance)
}

func (t *TToolButton) ExplicitHeight() int32 {
    defer exceptionProc()
    return ToolButton_GetExplicitHeight(t.instance)
}

func (t *TToolButton) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(ToolButton_GetParent(t.instance))
}

func (t *TToolButton) SetParent(value IControl) {
    defer exceptionProc()
    ToolButton_SetParent(t.instance, CheckPtr(value))
}

func (t *TToolButton) AlignWithMargins() bool {
    defer exceptionProc()
    return ToolButton_GetAlignWithMargins(t.instance)
}

func (t *TToolButton) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    ToolButton_SetAlignWithMargins(t.instance, value)
}

func (t *TToolButton) Left() int32 {
    defer exceptionProc()
    return ToolButton_GetLeft(t.instance)
}

func (t *TToolButton) SetLeft(value int32) {
    defer exceptionProc()
    ToolButton_SetLeft(t.instance, value)
}

func (t *TToolButton) Top() int32 {
    defer exceptionProc()
    return ToolButton_GetTop(t.instance)
}

func (t *TToolButton) SetTop(value int32) {
    defer exceptionProc()
    ToolButton_SetTop(t.instance, value)
}

func (t *TToolButton) Cursor() TCursor {
    defer exceptionProc()
    return ToolButton_GetCursor(t.instance)
}

func (t *TToolButton) SetCursor(value TCursor) {
    defer exceptionProc()
    ToolButton_SetCursor(t.instance, value)
}

func (t *TToolButton) Hint() string {
    defer exceptionProc()
    return ToolButton_GetHint(t.instance)
}

func (t *TToolButton) SetHint(value string) {
    defer exceptionProc()
    ToolButton_SetHint(t.instance, value)
}

func (t *TToolButton) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(ToolButton_GetMargins(t.instance))
}

func (t *TToolButton) SetMargins(value *TMargins) {
    defer exceptionProc()
    ToolButton_SetMargins(t.instance, CheckPtr(value))
}

func (t *TToolButton) ComponentCount() int32 {
    defer exceptionProc()
    return ToolButton_GetComponentCount(t.instance)
}

func (t *TToolButton) ComponentIndex() int32 {
    defer exceptionProc()
    return ToolButton_GetComponentIndex(t.instance)
}

func (t *TToolButton) SetComponentIndex(value int32) {
    defer exceptionProc()
    ToolButton_SetComponentIndex(t.instance, value)
}

func (t *TToolButton) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ToolButton_GetOwner(t.instance))
}

func (t *TToolButton) Name() string {
    defer exceptionProc()
    return ToolButton_GetName(t.instance)
}

func (t *TToolButton) SetName(value string) {
    defer exceptionProc()
    ToolButton_SetName(t.instance, value)
}

func (t *TToolButton) Tag() int {
    defer exceptionProc()
    return ToolButton_GetTag(t.instance)
}

func (t *TToolButton) SetTag(value int) {
    defer exceptionProc()
    ToolButton_SetTag(t.instance, value)
}

func (t *TToolButton) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ToolButton_GetComponents(t.instance, AIndex))
}

