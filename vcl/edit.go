
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TEdit struct {
    IControl
    instance uintptr
}

func NewEdit(owner IComponent) *TEdit {
    e := new(TEdit)
    e.instance = Edit_Create(owner.Instance())
    return e
}

func EditFromInst(inst uintptr) *TEdit {
    e := new(TEdit)
    e.instance = inst
    return e
}

func EditFromObj(obj IObject) *TEdit {
    e := new(TEdit)
    e.instance = CheckPtr(obj)
    return e
}

func (e *TEdit) Free() {
    if e.instance != 0 {
        Edit_Free(e.instance)
    }
}

func (e *TEdit) Instance() uintptr {
    return e.instance
}

func (e *TEdit) IsValid() bool {
    return e.instance != 0
}

func (e *TEdit) Clear() {
    defer exceptionProc()
    Edit_Clear(e.instance)
}

func (e *TEdit) ClearSelection() {
    defer exceptionProc()
    Edit_ClearSelection(e.instance)
}

func (e *TEdit) CopyToClipboard() {
    defer exceptionProc()
    Edit_CopyToClipboard(e.instance)
}

func (e *TEdit) CutToClipboard() {
    defer exceptionProc()
    Edit_CutToClipboard(e.instance)
}

func (e *TEdit) PasteFromClipboard() {
    defer exceptionProc()
    Edit_PasteFromClipboard(e.instance)
}

func (e *TEdit) SelectAll() {
    defer exceptionProc()
    Edit_SelectAll(e.instance)
}

func (e *TEdit) CanFocus() bool {
    defer exceptionProc()
    return Edit_CanFocus(e.instance)
}

func (e *TEdit) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    Edit_FlipChildren(e.instance, AllLevels )
}

func (e *TEdit) Focused() bool {
    defer exceptionProc()
    return Edit_Focused(e.instance)
}

func (e *TEdit) HandleAllocated() bool {
    defer exceptionProc()
    return Edit_HandleAllocated(e.instance)
}

func (e *TEdit) Invalidate() {
    defer exceptionProc()
    Edit_Invalidate(e.instance)
}

func (e *TEdit) Realign() {
    defer exceptionProc()
    Edit_Realign(e.instance)
}

func (e *TEdit) Repaint() {
    defer exceptionProc()
    Edit_Repaint(e.instance)
}

func (e *TEdit) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    Edit_ScaleBy(e.instance, M , D )
}

func (e *TEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    Edit_SetBounds(e.instance, ALeft , ATop , AWidth , AHeight )
}

func (e *TEdit) SetFocus() {
    defer exceptionProc()
    Edit_SetFocus(e.instance)
}

func (e *TEdit) Update() {
    defer exceptionProc()
    Edit_Update(e.instance)
}

func (e *TEdit) BringToFront() {
    defer exceptionProc()
    Edit_BringToFront(e.instance)
}

func (e *TEdit) HasParent() bool {
    defer exceptionProc()
    return Edit_HasParent(e.instance)
}

func (e *TEdit) Hide() {
    defer exceptionProc()
    Edit_Hide(e.instance)
}

func (e *TEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return Edit_Perform(e.instance, Msg , WParam , LParam )
}

func (e *TEdit) Refresh() {
    defer exceptionProc()
    Edit_Refresh(e.instance)
}

func (e *TEdit) SendToBack() {
    defer exceptionProc()
    Edit_SendToBack(e.instance)
}

func (e *TEdit) Show() {
    defer exceptionProc()
    Edit_Show(e.instance)
}

func (e *TEdit) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return Edit_GetTextBuf(e.instance, Buffer , BufSize )
}

func (e *TEdit) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Edit_FindComponent(e.instance, AName ))
}

func (e *TEdit) GetNamePath() string {
    defer exceptionProc()
    return Edit_GetNamePath(e.instance)
}

func (e *TEdit) Assign(Source IObject) {
    defer exceptionProc()
    Edit_Assign(e.instance, CheckPtr(Source))
}

func (e *TEdit) ClassName() string {
    defer exceptionProc()
    return Edit_ClassName(e.instance)
}

func (e *TEdit) Equals(Obj IObject) bool {
    defer exceptionProc()
    return Edit_Equals(e.instance, CheckPtr(Obj))
}

func (e *TEdit) GetHashCode() int32 {
    defer exceptionProc()
    return Edit_GetHashCode(e.instance)
}

func (e *TEdit) ToString() string {
    defer exceptionProc()
    return Edit_ToString(e.instance)
}

func (e *TEdit) Align() TAlign {
    defer exceptionProc()
    return Edit_GetAlign(e.instance)
}

func (e *TEdit) SetAlign(value TAlign) {
    defer exceptionProc()
    Edit_SetAlign(e.instance, value)
}

func (e *TEdit) Alignment() TAlignment {
    defer exceptionProc()
    return Edit_GetAlignment(e.instance)
}

func (e *TEdit) SetAlignment(value TAlignment) {
    defer exceptionProc()
    Edit_SetAlignment(e.instance, value)
}

func (e *TEdit) Anchors() TAnchors {
    defer exceptionProc()
    return Edit_GetAnchors(e.instance)
}

func (e *TEdit) SetAnchors(value TAnchors) {
    defer exceptionProc()
    Edit_SetAnchors(e.instance, value)
}

func (e *TEdit) AutoSelect() bool {
    defer exceptionProc()
    return Edit_GetAutoSelect(e.instance)
}

func (e *TEdit) SetAutoSelect(value bool) {
    defer exceptionProc()
    Edit_SetAutoSelect(e.instance, value)
}

func (e *TEdit) AutoSize() bool {
    defer exceptionProc()
    return Edit_GetAutoSize(e.instance)
}

func (e *TEdit) SetAutoSize(value bool) {
    defer exceptionProc()
    Edit_SetAutoSize(e.instance, value)
}

func (e *TEdit) BorderStyle() TBorderStyle {
    defer exceptionProc()
    return Edit_GetBorderStyle(e.instance)
}

func (e *TEdit) SetBorderStyle(value TBorderStyle) {
    defer exceptionProc()
    Edit_SetBorderStyle(e.instance, value)
}

func (e *TEdit) Color() TColor {
    defer exceptionProc()
    return Edit_GetColor(e.instance)
}

func (e *TEdit) SetColor(value TColor) {
    defer exceptionProc()
    Edit_SetColor(e.instance, value)
}

func (e *TEdit) DoubleBuffered() bool {
    defer exceptionProc()
    return Edit_GetDoubleBuffered(e.instance)
}

func (e *TEdit) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    Edit_SetDoubleBuffered(e.instance, value)
}

func (e *TEdit) Enabled() bool {
    defer exceptionProc()
    return Edit_GetEnabled(e.instance)
}

func (e *TEdit) SetEnabled(value bool) {
    defer exceptionProc()
    Edit_SetEnabled(e.instance, value)
}

func (e *TEdit) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(Edit_GetFont(e.instance))
}

func (e *TEdit) SetFont(value *TFont) {
    defer exceptionProc()
    Edit_SetFont(e.instance, CheckPtr(value))
}

func (e *TEdit) HideSelection() bool {
    defer exceptionProc()
    return Edit_GetHideSelection(e.instance)
}

func (e *TEdit) SetHideSelection(value bool) {
    defer exceptionProc()
    Edit_SetHideSelection(e.instance, value)
}

func (e *TEdit) MaxLength() int32 {
    defer exceptionProc()
    return Edit_GetMaxLength(e.instance)
}

func (e *TEdit) SetMaxLength(value int32) {
    defer exceptionProc()
    Edit_SetMaxLength(e.instance, value)
}

func (e *TEdit) NumbersOnly() bool {
    defer exceptionProc()
    return Edit_GetNumbersOnly(e.instance)
}

func (e *TEdit) SetNumbersOnly(value bool) {
    defer exceptionProc()
    Edit_SetNumbersOnly(e.instance, value)
}

func (e *TEdit) ParentColor() bool {
    defer exceptionProc()
    return Edit_GetParentColor(e.instance)
}

func (e *TEdit) SetParentColor(value bool) {
    defer exceptionProc()
    Edit_SetParentColor(e.instance, value)
}

func (e *TEdit) ParentCtl3D() bool {
    defer exceptionProc()
    return Edit_GetParentCtl3D(e.instance)
}

func (e *TEdit) SetParentCtl3D(value bool) {
    defer exceptionProc()
    Edit_SetParentCtl3D(e.instance, value)
}

func (e *TEdit) ParentFont() bool {
    defer exceptionProc()
    return Edit_GetParentFont(e.instance)
}

func (e *TEdit) SetParentFont(value bool) {
    defer exceptionProc()
    Edit_SetParentFont(e.instance, value)
}

func (e *TEdit) ParentShowHint() bool {
    defer exceptionProc()
    return Edit_GetParentShowHint(e.instance)
}

func (e *TEdit) SetParentShowHint(value bool) {
    defer exceptionProc()
    Edit_SetParentShowHint(e.instance, value)
}

func (e *TEdit) PasswordChar() uint16 {
    defer exceptionProc()
    return Edit_GetPasswordChar(e.instance)
}

func (e *TEdit) SetPasswordChar(value uint16) {
    defer exceptionProc()
    Edit_SetPasswordChar(e.instance, value)
}

func (e *TEdit) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(Edit_GetPopupMenu(e.instance))
}

func (e *TEdit) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    Edit_SetPopupMenu(e.instance, CheckPtr(value))
}

func (e *TEdit) ReadOnly() bool {
    defer exceptionProc()
    return Edit_GetReadOnly(e.instance)
}

func (e *TEdit) SetReadOnly(value bool) {
    defer exceptionProc()
    Edit_SetReadOnly(e.instance, value)
}

func (e *TEdit) ShowHint() bool {
    defer exceptionProc()
    return Edit_GetShowHint(e.instance)
}

func (e *TEdit) SetShowHint(value bool) {
    defer exceptionProc()
    Edit_SetShowHint(e.instance, value)
}

func (e *TEdit) TabOrder() int16 {
    defer exceptionProc()
    return Edit_GetTabOrder(e.instance)
}

func (e *TEdit) SetTabOrder(value int16) {
    defer exceptionProc()
    Edit_SetTabOrder(e.instance, value)
}

func (e *TEdit) TabStop() bool {
    defer exceptionProc()
    return Edit_GetTabStop(e.instance)
}

func (e *TEdit) SetTabStop(value bool) {
    defer exceptionProc()
    Edit_SetTabStop(e.instance, value)
}

func (e *TEdit) Text() string {
    defer exceptionProc()
    return Edit_GetText(e.instance)
}

func (e *TEdit) SetText(value string) {
    defer exceptionProc()
    Edit_SetText(e.instance, value)
}

func (e *TEdit) TextHint() string {
    defer exceptionProc()
    return Edit_GetTextHint(e.instance)
}

func (e *TEdit) SetTextHint(value string) {
    defer exceptionProc()
    Edit_SetTextHint(e.instance, value)
}

func (e *TEdit) Visible() bool {
    defer exceptionProc()
    return Edit_GetVisible(e.instance)
}

func (e *TEdit) SetVisible(value bool) {
    defer exceptionProc()
    Edit_SetVisible(e.instance, value)
}

func (e *TEdit) SetOnChange(fn TNotifyEvent) {
    Edit_SetOnChange(e.instance, fn)
}

func (e *TEdit) SetOnClick(fn TNotifyEvent) {
    Edit_SetOnClick(e.instance, fn)
}

func (e *TEdit) SetOnDblClick(fn TNotifyEvent) {
    Edit_SetOnDblClick(e.instance, fn)
}

func (e *TEdit) SetOnEnter(fn TNotifyEvent) {
    Edit_SetOnEnter(e.instance, fn)
}

func (e *TEdit) SetOnExit(fn TNotifyEvent) {
    Edit_SetOnExit(e.instance, fn)
}

func (e *TEdit) SetOnKeyDown(fn TKeyEvent) {
    Edit_SetOnKeyDown(e.instance, fn)
}

func (e *TEdit) SetOnKeyPress(fn TKeyPressEvent) {
    Edit_SetOnKeyPress(e.instance, fn)
}

func (e *TEdit) SetOnKeyUp(fn TKeyEvent) {
    Edit_SetOnKeyUp(e.instance, fn)
}

func (e *TEdit) SetOnMouseDown(fn TMouseEvent) {
    Edit_SetOnMouseDown(e.instance, fn)
}

func (e *TEdit) SetOnMouseEnter(fn TNotifyEvent) {
    Edit_SetOnMouseEnter(e.instance, fn)
}

func (e *TEdit) SetOnMouseLeave(fn TNotifyEvent) {
    Edit_SetOnMouseLeave(e.instance, fn)
}

func (e *TEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    Edit_SetOnMouseMove(e.instance, fn)
}

func (e *TEdit) SetOnMouseUp(fn TMouseEvent) {
    Edit_SetOnMouseUp(e.instance, fn)
}

func (e *TEdit) Modified() bool {
    defer exceptionProc()
    return Edit_GetModified(e.instance)
}

func (e *TEdit) SetModified(value bool) {
    defer exceptionProc()
    Edit_SetModified(e.instance, value)
}

func (e *TEdit) SelLength() int32 {
    defer exceptionProc()
    return Edit_GetSelLength(e.instance)
}

func (e *TEdit) SetSelLength(value int32) {
    defer exceptionProc()
    Edit_SetSelLength(e.instance, value)
}

func (e *TEdit) SelStart() int32 {
    defer exceptionProc()
    return Edit_GetSelStart(e.instance)
}

func (e *TEdit) SetSelStart(value int32) {
    defer exceptionProc()
    Edit_SetSelStart(e.instance, value)
}

func (e *TEdit) SelText() string {
    defer exceptionProc()
    return Edit_GetSelText(e.instance)
}

func (e *TEdit) SetSelText(value string) {
    defer exceptionProc()
    Edit_SetSelText(e.instance, value)
}

func (e *TEdit) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(Edit_GetBrush(e.instance))
}

func (e *TEdit) ControlCount() int32 {
    defer exceptionProc()
    return Edit_GetControlCount(e.instance)
}

func (e *TEdit) Handle() HWND {
    defer exceptionProc()
    return Edit_GetHandle(e.instance)
}

func (e *TEdit) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(Edit_GetAction(e.instance))
}

func (e *TEdit) SetAction(value IComponent) {
    defer exceptionProc()
    Edit_SetAction(e.instance, CheckPtr(value))
}

func (e *TEdit) BoundsRect() TRect {
    defer exceptionProc()
    return Edit_GetBoundsRect(e.instance)
}

func (e *TEdit) SetBoundsRect(value TRect) {
    defer exceptionProc()
    Edit_SetBoundsRect(e.instance, value)
}

func (e *TEdit) ClientHeight() int32 {
    defer exceptionProc()
    return Edit_GetClientHeight(e.instance)
}

func (e *TEdit) SetClientHeight(value int32) {
    defer exceptionProc()
    Edit_SetClientHeight(e.instance, value)
}

func (e *TEdit) ClientRect() TRect {
    defer exceptionProc()
    return Edit_GetClientRect(e.instance)
}

func (e *TEdit) ClientWidth() int32 {
    defer exceptionProc()
    return Edit_GetClientWidth(e.instance)
}

func (e *TEdit) SetClientWidth(value int32) {
    defer exceptionProc()
    Edit_SetClientWidth(e.instance, value)
}

func (e *TEdit) ExplicitLeft() int32 {
    defer exceptionProc()
    return Edit_GetExplicitLeft(e.instance)
}

func (e *TEdit) ExplicitTop() int32 {
    defer exceptionProc()
    return Edit_GetExplicitTop(e.instance)
}

func (e *TEdit) ExplicitWidth() int32 {
    defer exceptionProc()
    return Edit_GetExplicitWidth(e.instance)
}

func (e *TEdit) ExplicitHeight() int32 {
    defer exceptionProc()
    return Edit_GetExplicitHeight(e.instance)
}

func (e *TEdit) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(Edit_GetParent(e.instance))
}

func (e *TEdit) SetParent(value IControl) {
    defer exceptionProc()
    Edit_SetParent(e.instance, CheckPtr(value))
}

func (e *TEdit) AlignWithMargins() bool {
    defer exceptionProc()
    return Edit_GetAlignWithMargins(e.instance)
}

func (e *TEdit) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    Edit_SetAlignWithMargins(e.instance, value)
}

func (e *TEdit) Left() int32 {
    defer exceptionProc()
    return Edit_GetLeft(e.instance)
}

func (e *TEdit) SetLeft(value int32) {
    defer exceptionProc()
    Edit_SetLeft(e.instance, value)
}

func (e *TEdit) Top() int32 {
    defer exceptionProc()
    return Edit_GetTop(e.instance)
}

func (e *TEdit) SetTop(value int32) {
    defer exceptionProc()
    Edit_SetTop(e.instance, value)
}

func (e *TEdit) Width() int32 {
    defer exceptionProc()
    return Edit_GetWidth(e.instance)
}

func (e *TEdit) SetWidth(value int32) {
    defer exceptionProc()
    Edit_SetWidth(e.instance, value)
}

func (e *TEdit) Height() int32 {
    defer exceptionProc()
    return Edit_GetHeight(e.instance)
}

func (e *TEdit) SetHeight(value int32) {
    defer exceptionProc()
    Edit_SetHeight(e.instance, value)
}

func (e *TEdit) Cursor() TCursor {
    defer exceptionProc()
    return Edit_GetCursor(e.instance)
}

func (e *TEdit) SetCursor(value TCursor) {
    defer exceptionProc()
    Edit_SetCursor(e.instance, value)
}

func (e *TEdit) Hint() string {
    defer exceptionProc()
    return Edit_GetHint(e.instance)
}

func (e *TEdit) SetHint(value string) {
    defer exceptionProc()
    Edit_SetHint(e.instance, value)
}

func (e *TEdit) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(Edit_GetMargins(e.instance))
}

func (e *TEdit) SetMargins(value *TMargins) {
    defer exceptionProc()
    Edit_SetMargins(e.instance, CheckPtr(value))
}

func (e *TEdit) ComponentCount() int32 {
    defer exceptionProc()
    return Edit_GetComponentCount(e.instance)
}

func (e *TEdit) ComponentIndex() int32 {
    defer exceptionProc()
    return Edit_GetComponentIndex(e.instance)
}

func (e *TEdit) SetComponentIndex(value int32) {
    defer exceptionProc()
    Edit_SetComponentIndex(e.instance, value)
}

func (e *TEdit) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Edit_GetOwner(e.instance))
}

func (e *TEdit) Name() string {
    defer exceptionProc()
    return Edit_GetName(e.instance)
}

func (e *TEdit) SetName(value string) {
    defer exceptionProc()
    Edit_SetName(e.instance, value)
}

func (e *TEdit) Tag() int {
    defer exceptionProc()
    return Edit_GetTag(e.instance)
}

func (e *TEdit) SetTag(value int) {
    defer exceptionProc()
    Edit_SetTag(e.instance, value)
}

func (e *TEdit) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(Edit_GetControls(e.instance, Index))
}

func (e *TEdit) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Edit_GetComponents(e.instance, AIndex))
}

