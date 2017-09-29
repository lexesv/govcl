
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

type TMemo struct {
    IControl
    instance uintptr
}

func NewMemo(owner IComponent) *TMemo {
    m := new(TMemo)
    m.instance = Memo_Create(owner.Instance())
    return m
}

func MemoFromInst(inst uintptr) *TMemo {
    m := new(TMemo)
    m.instance = inst
    return m
}

func MemoFromObj(obj IObject) *TMemo {
    m := new(TMemo)
    m.instance = CheckPtr(obj)
    return m
}

func (m *TMemo) Free() {
    if m.instance != 0 {
        Memo_Free(m.instance)
    }
}

func (m *TMemo) Instance() uintptr {
    return m.instance
}

func (m *TMemo) IsValid() bool {
    return m.instance != 0
}

func (m *TMemo) Clear() {
    defer exceptionProc()
    Memo_Clear(m.instance)
}

func (m *TMemo) ClearSelection() {
    defer exceptionProc()
    Memo_ClearSelection(m.instance)
}

func (m *TMemo) CopyToClipboard() {
    defer exceptionProc()
    Memo_CopyToClipboard(m.instance)
}

func (m *TMemo) CutToClipboard() {
    defer exceptionProc()
    Memo_CutToClipboard(m.instance)
}

func (m *TMemo) PasteFromClipboard() {
    defer exceptionProc()
    Memo_PasteFromClipboard(m.instance)
}

func (m *TMemo) SelectAll() {
    defer exceptionProc()
    Memo_SelectAll(m.instance)
}

func (m *TMemo) CanFocus() bool {
    defer exceptionProc()
    return Memo_CanFocus(m.instance)
}

func (m *TMemo) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    Memo_FlipChildren(m.instance, AllLevels )
}

func (m *TMemo) Focused() bool {
    defer exceptionProc()
    return Memo_Focused(m.instance)
}

func (m *TMemo) HandleAllocated() bool {
    defer exceptionProc()
    return Memo_HandleAllocated(m.instance)
}

func (m *TMemo) Invalidate() {
    defer exceptionProc()
    Memo_Invalidate(m.instance)
}

func (m *TMemo) Realign() {
    defer exceptionProc()
    Memo_Realign(m.instance)
}

func (m *TMemo) Repaint() {
    defer exceptionProc()
    Memo_Repaint(m.instance)
}

func (m *TMemo) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    Memo_ScaleBy(m.instance, M , D )
}

func (m *TMemo) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    Memo_SetBounds(m.instance, ALeft , ATop , AWidth , AHeight )
}

func (m *TMemo) SetFocus() {
    defer exceptionProc()
    Memo_SetFocus(m.instance)
}

func (m *TMemo) Update() {
    defer exceptionProc()
    Memo_Update(m.instance)
}

func (m *TMemo) BringToFront() {
    defer exceptionProc()
    Memo_BringToFront(m.instance)
}

func (m *TMemo) HasParent() bool {
    defer exceptionProc()
    return Memo_HasParent(m.instance)
}

func (m *TMemo) Hide() {
    defer exceptionProc()
    Memo_Hide(m.instance)
}

func (m *TMemo) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return Memo_Perform(m.instance, Msg , WParam , LParam )
}

func (m *TMemo) Refresh() {
    defer exceptionProc()
    Memo_Refresh(m.instance)
}

func (m *TMemo) SendToBack() {
    defer exceptionProc()
    Memo_SendToBack(m.instance)
}

func (m *TMemo) Show() {
    defer exceptionProc()
    Memo_Show(m.instance)
}

func (m *TMemo) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return Memo_GetTextBuf(m.instance, Buffer , BufSize )
}

func (m *TMemo) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Memo_FindComponent(m.instance, AName ))
}

func (m *TMemo) GetNamePath() string {
    defer exceptionProc()
    return Memo_GetNamePath(m.instance)
}

func (m *TMemo) Assign(Source IObject) {
    defer exceptionProc()
    Memo_Assign(m.instance, CheckPtr(Source))
}

func (m *TMemo) ClassName() string {
    defer exceptionProc()
    return Memo_ClassName(m.instance)
}

func (m *TMemo) Equals(Obj IObject) bool {
    defer exceptionProc()
    return Memo_Equals(m.instance, CheckPtr(Obj))
}

func (m *TMemo) GetHashCode() int32 {
    defer exceptionProc()
    return Memo_GetHashCode(m.instance)
}

func (m *TMemo) ToString() string {
    defer exceptionProc()
    return Memo_ToString(m.instance)
}

func (m *TMemo) Align() TAlign {
    defer exceptionProc()
    return Memo_GetAlign(m.instance)
}

func (m *TMemo) SetAlign(value TAlign) {
    defer exceptionProc()
    Memo_SetAlign(m.instance, value)
}

func (m *TMemo) Alignment() TAlignment {
    defer exceptionProc()
    return Memo_GetAlignment(m.instance)
}

func (m *TMemo) SetAlignment(value TAlignment) {
    defer exceptionProc()
    Memo_SetAlignment(m.instance, value)
}

func (m *TMemo) Anchors() TAnchors {
    defer exceptionProc()
    return Memo_GetAnchors(m.instance)
}

func (m *TMemo) SetAnchors(value TAnchors) {
    defer exceptionProc()
    Memo_SetAnchors(m.instance, value)
}

func (m *TMemo) BorderStyle() TBorderStyle {
    defer exceptionProc()
    return Memo_GetBorderStyle(m.instance)
}

func (m *TMemo) SetBorderStyle(value TBorderStyle) {
    defer exceptionProc()
    Memo_SetBorderStyle(m.instance, value)
}

func (m *TMemo) Color() TColor {
    defer exceptionProc()
    return Memo_GetColor(m.instance)
}

func (m *TMemo) SetColor(value TColor) {
    defer exceptionProc()
    Memo_SetColor(m.instance, value)
}

func (m *TMemo) DoubleBuffered() bool {
    defer exceptionProc()
    return Memo_GetDoubleBuffered(m.instance)
}

func (m *TMemo) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    Memo_SetDoubleBuffered(m.instance, value)
}

func (m *TMemo) Enabled() bool {
    defer exceptionProc()
    return Memo_GetEnabled(m.instance)
}

func (m *TMemo) SetEnabled(value bool) {
    defer exceptionProc()
    Memo_SetEnabled(m.instance, value)
}

func (m *TMemo) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(Memo_GetFont(m.instance))
}

func (m *TMemo) SetFont(value *TFont) {
    defer exceptionProc()
    Memo_SetFont(m.instance, CheckPtr(value))
}

func (m *TMemo) HideSelection() bool {
    defer exceptionProc()
    return Memo_GetHideSelection(m.instance)
}

func (m *TMemo) SetHideSelection(value bool) {
    defer exceptionProc()
    Memo_SetHideSelection(m.instance, value)
}

func (m *TMemo) Lines() *TStrings {
    defer exceptionProc()
    return StringsFromInst(Memo_GetLines(m.instance))
}

func (m *TMemo) SetLines(value IObject) {
    defer exceptionProc()
    Memo_SetLines(m.instance, CheckPtr(value))
}

func (m *TMemo) MaxLength() int32 {
    defer exceptionProc()
    return Memo_GetMaxLength(m.instance)
}

func (m *TMemo) SetMaxLength(value int32) {
    defer exceptionProc()
    Memo_SetMaxLength(m.instance, value)
}

func (m *TMemo) ParentColor() bool {
    defer exceptionProc()
    return Memo_GetParentColor(m.instance)
}

func (m *TMemo) SetParentColor(value bool) {
    defer exceptionProc()
    Memo_SetParentColor(m.instance, value)
}

func (m *TMemo) ParentCtl3D() bool {
    defer exceptionProc()
    return Memo_GetParentCtl3D(m.instance)
}

func (m *TMemo) SetParentCtl3D(value bool) {
    defer exceptionProc()
    Memo_SetParentCtl3D(m.instance, value)
}

func (m *TMemo) ParentFont() bool {
    defer exceptionProc()
    return Memo_GetParentFont(m.instance)
}

func (m *TMemo) SetParentFont(value bool) {
    defer exceptionProc()
    Memo_SetParentFont(m.instance, value)
}

func (m *TMemo) ParentShowHint() bool {
    defer exceptionProc()
    return Memo_GetParentShowHint(m.instance)
}

func (m *TMemo) SetParentShowHint(value bool) {
    defer exceptionProc()
    Memo_SetParentShowHint(m.instance, value)
}

func (m *TMemo) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(Memo_GetPopupMenu(m.instance))
}

func (m *TMemo) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    Memo_SetPopupMenu(m.instance, CheckPtr(value))
}

func (m *TMemo) ReadOnly() bool {
    defer exceptionProc()
    return Memo_GetReadOnly(m.instance)
}

func (m *TMemo) SetReadOnly(value bool) {
    defer exceptionProc()
    Memo_SetReadOnly(m.instance, value)
}

func (m *TMemo) ScrollBars() TScrollStyle {
    defer exceptionProc()
    return Memo_GetScrollBars(m.instance)
}

func (m *TMemo) SetScrollBars(value TScrollStyle) {
    defer exceptionProc()
    Memo_SetScrollBars(m.instance, value)
}

func (m *TMemo) ShowHint() bool {
    defer exceptionProc()
    return Memo_GetShowHint(m.instance)
}

func (m *TMemo) SetShowHint(value bool) {
    defer exceptionProc()
    Memo_SetShowHint(m.instance, value)
}

func (m *TMemo) TabOrder() int16 {
    defer exceptionProc()
    return Memo_GetTabOrder(m.instance)
}

func (m *TMemo) SetTabOrder(value int16) {
    defer exceptionProc()
    Memo_SetTabOrder(m.instance, value)
}

func (m *TMemo) TabStop() bool {
    defer exceptionProc()
    return Memo_GetTabStop(m.instance)
}

func (m *TMemo) SetTabStop(value bool) {
    defer exceptionProc()
    Memo_SetTabStop(m.instance, value)
}

func (m *TMemo) Visible() bool {
    defer exceptionProc()
    return Memo_GetVisible(m.instance)
}

func (m *TMemo) SetVisible(value bool) {
    defer exceptionProc()
    Memo_SetVisible(m.instance, value)
}

func (m *TMemo) WantReturns() bool {
    defer exceptionProc()
    return Memo_GetWantReturns(m.instance)
}

func (m *TMemo) SetWantReturns(value bool) {
    defer exceptionProc()
    Memo_SetWantReturns(m.instance, value)
}

func (m *TMemo) WantTabs() bool {
    defer exceptionProc()
    return Memo_GetWantTabs(m.instance)
}

func (m *TMemo) SetWantTabs(value bool) {
    defer exceptionProc()
    Memo_SetWantTabs(m.instance, value)
}

func (m *TMemo) WordWrap() bool {
    defer exceptionProc()
    return Memo_GetWordWrap(m.instance)
}

func (m *TMemo) SetWordWrap(value bool) {
    defer exceptionProc()
    Memo_SetWordWrap(m.instance, value)
}

func (m *TMemo) SetOnChange(fn TNotifyEvent) {
    Memo_SetOnChange(m.instance, fn)
}

func (m *TMemo) SetOnClick(fn TNotifyEvent) {
    Memo_SetOnClick(m.instance, fn)
}

func (m *TMemo) SetOnDblClick(fn TNotifyEvent) {
    Memo_SetOnDblClick(m.instance, fn)
}

func (m *TMemo) SetOnEnter(fn TNotifyEvent) {
    Memo_SetOnEnter(m.instance, fn)
}

func (m *TMemo) SetOnExit(fn TNotifyEvent) {
    Memo_SetOnExit(m.instance, fn)
}

func (m *TMemo) SetOnKeyDown(fn TKeyEvent) {
    Memo_SetOnKeyDown(m.instance, fn)
}

func (m *TMemo) SetOnKeyPress(fn TKeyPressEvent) {
    Memo_SetOnKeyPress(m.instance, fn)
}

func (m *TMemo) SetOnKeyUp(fn TKeyEvent) {
    Memo_SetOnKeyUp(m.instance, fn)
}

func (m *TMemo) SetOnMouseDown(fn TMouseEvent) {
    Memo_SetOnMouseDown(m.instance, fn)
}

func (m *TMemo) SetOnMouseEnter(fn TNotifyEvent) {
    Memo_SetOnMouseEnter(m.instance, fn)
}

func (m *TMemo) SetOnMouseLeave(fn TNotifyEvent) {
    Memo_SetOnMouseLeave(m.instance, fn)
}

func (m *TMemo) SetOnMouseMove(fn TMouseMoveEvent) {
    Memo_SetOnMouseMove(m.instance, fn)
}

func (m *TMemo) SetOnMouseUp(fn TMouseEvent) {
    Memo_SetOnMouseUp(m.instance, fn)
}

func (m *TMemo) Modified() bool {
    defer exceptionProc()
    return Memo_GetModified(m.instance)
}

func (m *TMemo) SetModified(value bool) {
    defer exceptionProc()
    Memo_SetModified(m.instance, value)
}

func (m *TMemo) SelLength() int32 {
    defer exceptionProc()
    return Memo_GetSelLength(m.instance)
}

func (m *TMemo) SetSelLength(value int32) {
    defer exceptionProc()
    Memo_SetSelLength(m.instance, value)
}

func (m *TMemo) SelStart() int32 {
    defer exceptionProc()
    return Memo_GetSelStart(m.instance)
}

func (m *TMemo) SetSelStart(value int32) {
    defer exceptionProc()
    Memo_SetSelStart(m.instance, value)
}

func (m *TMemo) SelText() string {
    defer exceptionProc()
    return Memo_GetSelText(m.instance)
}

func (m *TMemo) SetSelText(value string) {
    defer exceptionProc()
    Memo_SetSelText(m.instance, value)
}

func (m *TMemo) Text() string {
    defer exceptionProc()
    return Memo_GetText(m.instance)
}

func (m *TMemo) SetText(value string) {
    defer exceptionProc()
    Memo_SetText(m.instance, value)
}

func (m *TMemo) TextHint() string {
    defer exceptionProc()
    return Memo_GetTextHint(m.instance)
}

func (m *TMemo) SetTextHint(value string) {
    defer exceptionProc()
    Memo_SetTextHint(m.instance, value)
}

func (m *TMemo) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(Memo_GetBrush(m.instance))
}

func (m *TMemo) ControlCount() int32 {
    defer exceptionProc()
    return Memo_GetControlCount(m.instance)
}

func (m *TMemo) Handle() HWND {
    defer exceptionProc()
    return Memo_GetHandle(m.instance)
}

func (m *TMemo) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(Memo_GetAction(m.instance))
}

func (m *TMemo) SetAction(value IComponent) {
    defer exceptionProc()
    Memo_SetAction(m.instance, CheckPtr(value))
}

func (m *TMemo) BoundsRect() TRect {
    defer exceptionProc()
    return Memo_GetBoundsRect(m.instance)
}

func (m *TMemo) SetBoundsRect(value TRect) {
    defer exceptionProc()
    Memo_SetBoundsRect(m.instance, value)
}

func (m *TMemo) ClientHeight() int32 {
    defer exceptionProc()
    return Memo_GetClientHeight(m.instance)
}

func (m *TMemo) SetClientHeight(value int32) {
    defer exceptionProc()
    Memo_SetClientHeight(m.instance, value)
}

func (m *TMemo) ClientRect() TRect {
    defer exceptionProc()
    return Memo_GetClientRect(m.instance)
}

func (m *TMemo) ClientWidth() int32 {
    defer exceptionProc()
    return Memo_GetClientWidth(m.instance)
}

func (m *TMemo) SetClientWidth(value int32) {
    defer exceptionProc()
    Memo_SetClientWidth(m.instance, value)
}

func (m *TMemo) ExplicitLeft() int32 {
    defer exceptionProc()
    return Memo_GetExplicitLeft(m.instance)
}

func (m *TMemo) ExplicitTop() int32 {
    defer exceptionProc()
    return Memo_GetExplicitTop(m.instance)
}

func (m *TMemo) ExplicitWidth() int32 {
    defer exceptionProc()
    return Memo_GetExplicitWidth(m.instance)
}

func (m *TMemo) ExplicitHeight() int32 {
    defer exceptionProc()
    return Memo_GetExplicitHeight(m.instance)
}

func (m *TMemo) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(Memo_GetParent(m.instance))
}

func (m *TMemo) SetParent(value IControl) {
    defer exceptionProc()
    Memo_SetParent(m.instance, CheckPtr(value))
}

func (m *TMemo) AlignWithMargins() bool {
    defer exceptionProc()
    return Memo_GetAlignWithMargins(m.instance)
}

func (m *TMemo) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    Memo_SetAlignWithMargins(m.instance, value)
}

func (m *TMemo) Left() int32 {
    defer exceptionProc()
    return Memo_GetLeft(m.instance)
}

func (m *TMemo) SetLeft(value int32) {
    defer exceptionProc()
    Memo_SetLeft(m.instance, value)
}

func (m *TMemo) Top() int32 {
    defer exceptionProc()
    return Memo_GetTop(m.instance)
}

func (m *TMemo) SetTop(value int32) {
    defer exceptionProc()
    Memo_SetTop(m.instance, value)
}

func (m *TMemo) Width() int32 {
    defer exceptionProc()
    return Memo_GetWidth(m.instance)
}

func (m *TMemo) SetWidth(value int32) {
    defer exceptionProc()
    Memo_SetWidth(m.instance, value)
}

func (m *TMemo) Height() int32 {
    defer exceptionProc()
    return Memo_GetHeight(m.instance)
}

func (m *TMemo) SetHeight(value int32) {
    defer exceptionProc()
    Memo_SetHeight(m.instance, value)
}

func (m *TMemo) Cursor() TCursor {
    defer exceptionProc()
    return Memo_GetCursor(m.instance)
}

func (m *TMemo) SetCursor(value TCursor) {
    defer exceptionProc()
    Memo_SetCursor(m.instance, value)
}

func (m *TMemo) Hint() string {
    defer exceptionProc()
    return Memo_GetHint(m.instance)
}

func (m *TMemo) SetHint(value string) {
    defer exceptionProc()
    Memo_SetHint(m.instance, value)
}

func (m *TMemo) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(Memo_GetMargins(m.instance))
}

func (m *TMemo) SetMargins(value *TMargins) {
    defer exceptionProc()
    Memo_SetMargins(m.instance, CheckPtr(value))
}

func (m *TMemo) ComponentCount() int32 {
    defer exceptionProc()
    return Memo_GetComponentCount(m.instance)
}

func (m *TMemo) ComponentIndex() int32 {
    defer exceptionProc()
    return Memo_GetComponentIndex(m.instance)
}

func (m *TMemo) SetComponentIndex(value int32) {
    defer exceptionProc()
    Memo_SetComponentIndex(m.instance, value)
}

func (m *TMemo) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Memo_GetOwner(m.instance))
}

func (m *TMemo) Name() string {
    defer exceptionProc()
    return Memo_GetName(m.instance)
}

func (m *TMemo) SetName(value string) {
    defer exceptionProc()
    Memo_SetName(m.instance, value)
}

func (m *TMemo) Tag() int {
    defer exceptionProc()
    return Memo_GetTag(m.instance)
}

func (m *TMemo) SetTag(value int) {
    defer exceptionProc()
    Memo_SetTag(m.instance, value)
}

func (m *TMemo) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(Memo_GetControls(m.instance, Index))
}

func (m *TMemo) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Memo_GetComponents(m.instance, AIndex))
}

