
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TRichEdit struct {
    IControl
    instance uintptr
}

func NewRichEdit(owner IComponent) *TRichEdit {
    r := new(TRichEdit)
    r.instance = RichEdit_Create(owner.Instance())
    return r
}

func RichEditFromInst(inst uintptr) *TRichEdit {
    r := new(TRichEdit)
    r.instance = inst
    return r
}

func RichEditFromObj(obj IObject) *TRichEdit {
    r := new(TRichEdit)
    r.instance = CheckPtr(obj)
    return r
}

func (r *TRichEdit) Free() {
    if r.instance != 0 {
        RichEdit_Free(r.instance)
    }
}

func (r *TRichEdit) Instance() uintptr {
    return r.instance
}

func (r *TRichEdit) IsValid() bool {
    return r.instance != 0
}

func (r *TRichEdit) Clear() {
    defer exceptionProc()
    RichEdit_Clear(r.instance)
}

func (r *TRichEdit) ClearSelection() {
    defer exceptionProc()
    RichEdit_ClearSelection(r.instance)
}

func (r *TRichEdit) CopyToClipboard() {
    defer exceptionProc()
    RichEdit_CopyToClipboard(r.instance)
}

func (r *TRichEdit) CutToClipboard() {
    defer exceptionProc()
    RichEdit_CutToClipboard(r.instance)
}

func (r *TRichEdit) PasteFromClipboard() {
    defer exceptionProc()
    RichEdit_PasteFromClipboard(r.instance)
}

func (r *TRichEdit) SelectAll() {
    defer exceptionProc()
    RichEdit_SelectAll(r.instance)
}

func (r *TRichEdit) CanFocus() bool {
    defer exceptionProc()
    return RichEdit_CanFocus(r.instance)
}

func (r *TRichEdit) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    RichEdit_FlipChildren(r.instance, AllLevels )
}

func (r *TRichEdit) Focused() bool {
    defer exceptionProc()
    return RichEdit_Focused(r.instance)
}

func (r *TRichEdit) HandleAllocated() bool {
    defer exceptionProc()
    return RichEdit_HandleAllocated(r.instance)
}

func (r *TRichEdit) Invalidate() {
    defer exceptionProc()
    RichEdit_Invalidate(r.instance)
}

func (r *TRichEdit) Realign() {
    defer exceptionProc()
    RichEdit_Realign(r.instance)
}

func (r *TRichEdit) Repaint() {
    defer exceptionProc()
    RichEdit_Repaint(r.instance)
}

func (r *TRichEdit) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    RichEdit_ScaleBy(r.instance, M , D )
}

func (r *TRichEdit) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    RichEdit_SetBounds(r.instance, ALeft , ATop , AWidth , AHeight )
}

func (r *TRichEdit) SetFocus() {
    defer exceptionProc()
    RichEdit_SetFocus(r.instance)
}

func (r *TRichEdit) Update() {
    defer exceptionProc()
    RichEdit_Update(r.instance)
}

func (r *TRichEdit) BringToFront() {
    defer exceptionProc()
    RichEdit_BringToFront(r.instance)
}

func (r *TRichEdit) HasParent() bool {
    defer exceptionProc()
    return RichEdit_HasParent(r.instance)
}

func (r *TRichEdit) Hide() {
    defer exceptionProc()
    RichEdit_Hide(r.instance)
}

func (r *TRichEdit) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return RichEdit_Perform(r.instance, Msg , WParam , LParam )
}

func (r *TRichEdit) Refresh() {
    defer exceptionProc()
    RichEdit_Refresh(r.instance)
}

func (r *TRichEdit) SendToBack() {
    defer exceptionProc()
    RichEdit_SendToBack(r.instance)
}

func (r *TRichEdit) Show() {
    defer exceptionProc()
    RichEdit_Show(r.instance)
}

func (r *TRichEdit) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return RichEdit_GetTextBuf(r.instance, Buffer , BufSize )
}

func (r *TRichEdit) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(RichEdit_FindComponent(r.instance, AName ))
}

func (r *TRichEdit) GetNamePath() string {
    defer exceptionProc()
    return RichEdit_GetNamePath(r.instance)
}

func (r *TRichEdit) Assign(Source IObject) {
    defer exceptionProc()
    RichEdit_Assign(r.instance, CheckPtr(Source))
}

func (r *TRichEdit) ClassName() string {
    defer exceptionProc()
    return RichEdit_ClassName(r.instance)
}

func (r *TRichEdit) Equals(Obj IObject) bool {
    defer exceptionProc()
    return RichEdit_Equals(r.instance, CheckPtr(Obj))
}

func (r *TRichEdit) GetHashCode() int32 {
    defer exceptionProc()
    return RichEdit_GetHashCode(r.instance)
}

func (r *TRichEdit) ToString() string {
    defer exceptionProc()
    return RichEdit_ToString(r.instance)
}

func (r *TRichEdit) Align() TAlign {
    defer exceptionProc()
    return RichEdit_GetAlign(r.instance)
}

func (r *TRichEdit) SetAlign(value TAlign) {
    defer exceptionProc()
    RichEdit_SetAlign(r.instance, value)
}

func (r *TRichEdit) Alignment() TAlignment {
    defer exceptionProc()
    return RichEdit_GetAlignment(r.instance)
}

func (r *TRichEdit) SetAlignment(value TAlignment) {
    defer exceptionProc()
    RichEdit_SetAlignment(r.instance, value)
}

func (r *TRichEdit) Anchors() TAnchors {
    defer exceptionProc()
    return RichEdit_GetAnchors(r.instance)
}

func (r *TRichEdit) SetAnchors(value TAnchors) {
    defer exceptionProc()
    RichEdit_SetAnchors(r.instance, value)
}

func (r *TRichEdit) BorderStyle() TBorderStyle {
    defer exceptionProc()
    return RichEdit_GetBorderStyle(r.instance)
}

func (r *TRichEdit) SetBorderStyle(value TBorderStyle) {
    defer exceptionProc()
    RichEdit_SetBorderStyle(r.instance, value)
}

func (r *TRichEdit) BorderWidth() int32 {
    defer exceptionProc()
    return RichEdit_GetBorderWidth(r.instance)
}

func (r *TRichEdit) SetBorderWidth(value int32) {
    defer exceptionProc()
    RichEdit_SetBorderWidth(r.instance, value)
}

func (r *TRichEdit) Color() TColor {
    defer exceptionProc()
    return RichEdit_GetColor(r.instance)
}

func (r *TRichEdit) SetColor(value TColor) {
    defer exceptionProc()
    RichEdit_SetColor(r.instance, value)
}

func (r *TRichEdit) Enabled() bool {
    defer exceptionProc()
    return RichEdit_GetEnabled(r.instance)
}

func (r *TRichEdit) SetEnabled(value bool) {
    defer exceptionProc()
    RichEdit_SetEnabled(r.instance, value)
}

func (r *TRichEdit) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(RichEdit_GetFont(r.instance))
}

func (r *TRichEdit) SetFont(value *TFont) {
    defer exceptionProc()
    RichEdit_SetFont(r.instance, CheckPtr(value))
}

func (r *TRichEdit) HideSelection() bool {
    defer exceptionProc()
    return RichEdit_GetHideSelection(r.instance)
}

func (r *TRichEdit) SetHideSelection(value bool) {
    defer exceptionProc()
    RichEdit_SetHideSelection(r.instance, value)
}

func (r *TRichEdit) Lines() *TStrings {
    defer exceptionProc()
    return StringsFromInst(RichEdit_GetLines(r.instance))
}

func (r *TRichEdit) SetLines(value IObject) {
    defer exceptionProc()
    RichEdit_SetLines(r.instance, CheckPtr(value))
}

func (r *TRichEdit) MaxLength() int32 {
    defer exceptionProc()
    return RichEdit_GetMaxLength(r.instance)
}

func (r *TRichEdit) SetMaxLength(value int32) {
    defer exceptionProc()
    RichEdit_SetMaxLength(r.instance, value)
}

func (r *TRichEdit) ParentColor() bool {
    defer exceptionProc()
    return RichEdit_GetParentColor(r.instance)
}

func (r *TRichEdit) SetParentColor(value bool) {
    defer exceptionProc()
    RichEdit_SetParentColor(r.instance, value)
}

func (r *TRichEdit) ParentCtl3D() bool {
    defer exceptionProc()
    return RichEdit_GetParentCtl3D(r.instance)
}

func (r *TRichEdit) SetParentCtl3D(value bool) {
    defer exceptionProc()
    RichEdit_SetParentCtl3D(r.instance, value)
}

func (r *TRichEdit) ParentFont() bool {
    defer exceptionProc()
    return RichEdit_GetParentFont(r.instance)
}

func (r *TRichEdit) SetParentFont(value bool) {
    defer exceptionProc()
    RichEdit_SetParentFont(r.instance, value)
}

func (r *TRichEdit) ParentShowHint() bool {
    defer exceptionProc()
    return RichEdit_GetParentShowHint(r.instance)
}

func (r *TRichEdit) SetParentShowHint(value bool) {
    defer exceptionProc()
    RichEdit_SetParentShowHint(r.instance, value)
}

func (r *TRichEdit) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(RichEdit_GetPopupMenu(r.instance))
}

func (r *TRichEdit) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    RichEdit_SetPopupMenu(r.instance, CheckPtr(value))
}

func (r *TRichEdit) ReadOnly() bool {
    defer exceptionProc()
    return RichEdit_GetReadOnly(r.instance)
}

func (r *TRichEdit) SetReadOnly(value bool) {
    defer exceptionProc()
    RichEdit_SetReadOnly(r.instance, value)
}

func (r *TRichEdit) ScrollBars() TScrollStyle {
    defer exceptionProc()
    return RichEdit_GetScrollBars(r.instance)
}

func (r *TRichEdit) SetScrollBars(value TScrollStyle) {
    defer exceptionProc()
    RichEdit_SetScrollBars(r.instance, value)
}

func (r *TRichEdit) ShowHint() bool {
    defer exceptionProc()
    return RichEdit_GetShowHint(r.instance)
}

func (r *TRichEdit) SetShowHint(value bool) {
    defer exceptionProc()
    RichEdit_SetShowHint(r.instance, value)
}

func (r *TRichEdit) TabOrder() int16 {
    defer exceptionProc()
    return RichEdit_GetTabOrder(r.instance)
}

func (r *TRichEdit) SetTabOrder(value int16) {
    defer exceptionProc()
    RichEdit_SetTabOrder(r.instance, value)
}

func (r *TRichEdit) TabStop() bool {
    defer exceptionProc()
    return RichEdit_GetTabStop(r.instance)
}

func (r *TRichEdit) SetTabStop(value bool) {
    defer exceptionProc()
    RichEdit_SetTabStop(r.instance, value)
}

func (r *TRichEdit) Visible() bool {
    defer exceptionProc()
    return RichEdit_GetVisible(r.instance)
}

func (r *TRichEdit) SetVisible(value bool) {
    defer exceptionProc()
    RichEdit_SetVisible(r.instance, value)
}

func (r *TRichEdit) WantTabs() bool {
    defer exceptionProc()
    return RichEdit_GetWantTabs(r.instance)
}

func (r *TRichEdit) SetWantTabs(value bool) {
    defer exceptionProc()
    RichEdit_SetWantTabs(r.instance, value)
}

func (r *TRichEdit) WantReturns() bool {
    defer exceptionProc()
    return RichEdit_GetWantReturns(r.instance)
}

func (r *TRichEdit) SetWantReturns(value bool) {
    defer exceptionProc()
    RichEdit_SetWantReturns(r.instance, value)
}

func (r *TRichEdit) WordWrap() bool {
    defer exceptionProc()
    return RichEdit_GetWordWrap(r.instance)
}

func (r *TRichEdit) SetWordWrap(value bool) {
    defer exceptionProc()
    RichEdit_SetWordWrap(r.instance, value)
}

func (r *TRichEdit) SetOnChange(fn TNotifyEvent) {
    RichEdit_SetOnChange(r.instance, fn)
}

func (r *TRichEdit) SetOnClick(fn TNotifyEvent) {
    RichEdit_SetOnClick(r.instance, fn)
}

func (r *TRichEdit) SetOnDblClick(fn TNotifyEvent) {
    RichEdit_SetOnDblClick(r.instance, fn)
}

func (r *TRichEdit) SetOnEnter(fn TNotifyEvent) {
    RichEdit_SetOnEnter(r.instance, fn)
}

func (r *TRichEdit) SetOnExit(fn TNotifyEvent) {
    RichEdit_SetOnExit(r.instance, fn)
}

func (r *TRichEdit) SetOnKeyDown(fn TKeyEvent) {
    RichEdit_SetOnKeyDown(r.instance, fn)
}

func (r *TRichEdit) SetOnKeyPress(fn TKeyPressEvent) {
    RichEdit_SetOnKeyPress(r.instance, fn)
}

func (r *TRichEdit) SetOnKeyUp(fn TKeyEvent) {
    RichEdit_SetOnKeyUp(r.instance, fn)
}

func (r *TRichEdit) SetOnMouseDown(fn TMouseEvent) {
    RichEdit_SetOnMouseDown(r.instance, fn)
}

func (r *TRichEdit) SetOnMouseEnter(fn TNotifyEvent) {
    RichEdit_SetOnMouseEnter(r.instance, fn)
}

func (r *TRichEdit) SetOnMouseLeave(fn TNotifyEvent) {
    RichEdit_SetOnMouseLeave(r.instance, fn)
}

func (r *TRichEdit) SetOnMouseMove(fn TMouseMoveEvent) {
    RichEdit_SetOnMouseMove(r.instance, fn)
}

func (r *TRichEdit) SetOnMouseUp(fn TMouseEvent) {
    RichEdit_SetOnMouseUp(r.instance, fn)
}

func (r *TRichEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
    RichEdit_SetOnMouseWheel(r.instance, fn)
}

func (r *TRichEdit) Modified() bool {
    defer exceptionProc()
    return RichEdit_GetModified(r.instance)
}

func (r *TRichEdit) SetModified(value bool) {
    defer exceptionProc()
    RichEdit_SetModified(r.instance, value)
}

func (r *TRichEdit) SelLength() int32 {
    defer exceptionProc()
    return RichEdit_GetSelLength(r.instance)
}

func (r *TRichEdit) SetSelLength(value int32) {
    defer exceptionProc()
    RichEdit_SetSelLength(r.instance, value)
}

func (r *TRichEdit) SelStart() int32 {
    defer exceptionProc()
    return RichEdit_GetSelStart(r.instance)
}

func (r *TRichEdit) SetSelStart(value int32) {
    defer exceptionProc()
    RichEdit_SetSelStart(r.instance, value)
}

func (r *TRichEdit) SelText() string {
    defer exceptionProc()
    return RichEdit_GetSelText(r.instance)
}

func (r *TRichEdit) SetSelText(value string) {
    defer exceptionProc()
    RichEdit_SetSelText(r.instance, value)
}

func (r *TRichEdit) Text() string {
    defer exceptionProc()
    return RichEdit_GetText(r.instance)
}

func (r *TRichEdit) SetText(value string) {
    defer exceptionProc()
    RichEdit_SetText(r.instance, value)
}

func (r *TRichEdit) TextHint() string {
    defer exceptionProc()
    return RichEdit_GetTextHint(r.instance)
}

func (r *TRichEdit) SetTextHint(value string) {
    defer exceptionProc()
    RichEdit_SetTextHint(r.instance, value)
}

func (r *TRichEdit) DoubleBuffered() bool {
    defer exceptionProc()
    return RichEdit_GetDoubleBuffered(r.instance)
}

func (r *TRichEdit) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    RichEdit_SetDoubleBuffered(r.instance, value)
}

func (r *TRichEdit) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(RichEdit_GetBrush(r.instance))
}

func (r *TRichEdit) ControlCount() int32 {
    defer exceptionProc()
    return RichEdit_GetControlCount(r.instance)
}

func (r *TRichEdit) Handle() HWND {
    defer exceptionProc()
    return RichEdit_GetHandle(r.instance)
}

func (r *TRichEdit) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(RichEdit_GetAction(r.instance))
}

func (r *TRichEdit) SetAction(value IComponent) {
    defer exceptionProc()
    RichEdit_SetAction(r.instance, CheckPtr(value))
}

func (r *TRichEdit) BoundsRect() TRect {
    defer exceptionProc()
    return RichEdit_GetBoundsRect(r.instance)
}

func (r *TRichEdit) SetBoundsRect(value TRect) {
    defer exceptionProc()
    RichEdit_SetBoundsRect(r.instance, value)
}

func (r *TRichEdit) ClientHeight() int32 {
    defer exceptionProc()
    return RichEdit_GetClientHeight(r.instance)
}

func (r *TRichEdit) SetClientHeight(value int32) {
    defer exceptionProc()
    RichEdit_SetClientHeight(r.instance, value)
}

func (r *TRichEdit) ClientRect() TRect {
    defer exceptionProc()
    return RichEdit_GetClientRect(r.instance)
}

func (r *TRichEdit) ClientWidth() int32 {
    defer exceptionProc()
    return RichEdit_GetClientWidth(r.instance)
}

func (r *TRichEdit) SetClientWidth(value int32) {
    defer exceptionProc()
    RichEdit_SetClientWidth(r.instance, value)
}

func (r *TRichEdit) ExplicitLeft() int32 {
    defer exceptionProc()
    return RichEdit_GetExplicitLeft(r.instance)
}

func (r *TRichEdit) ExplicitTop() int32 {
    defer exceptionProc()
    return RichEdit_GetExplicitTop(r.instance)
}

func (r *TRichEdit) ExplicitWidth() int32 {
    defer exceptionProc()
    return RichEdit_GetExplicitWidth(r.instance)
}

func (r *TRichEdit) ExplicitHeight() int32 {
    defer exceptionProc()
    return RichEdit_GetExplicitHeight(r.instance)
}

func (r *TRichEdit) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(RichEdit_GetParent(r.instance))
}

func (r *TRichEdit) SetParent(value IControl) {
    defer exceptionProc()
    RichEdit_SetParent(r.instance, CheckPtr(value))
}

func (r *TRichEdit) AlignWithMargins() bool {
    defer exceptionProc()
    return RichEdit_GetAlignWithMargins(r.instance)
}

func (r *TRichEdit) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    RichEdit_SetAlignWithMargins(r.instance, value)
}

func (r *TRichEdit) Left() int32 {
    defer exceptionProc()
    return RichEdit_GetLeft(r.instance)
}

func (r *TRichEdit) SetLeft(value int32) {
    defer exceptionProc()
    RichEdit_SetLeft(r.instance, value)
}

func (r *TRichEdit) Top() int32 {
    defer exceptionProc()
    return RichEdit_GetTop(r.instance)
}

func (r *TRichEdit) SetTop(value int32) {
    defer exceptionProc()
    RichEdit_SetTop(r.instance, value)
}

func (r *TRichEdit) Width() int32 {
    defer exceptionProc()
    return RichEdit_GetWidth(r.instance)
}

func (r *TRichEdit) SetWidth(value int32) {
    defer exceptionProc()
    RichEdit_SetWidth(r.instance, value)
}

func (r *TRichEdit) Height() int32 {
    defer exceptionProc()
    return RichEdit_GetHeight(r.instance)
}

func (r *TRichEdit) SetHeight(value int32) {
    defer exceptionProc()
    RichEdit_SetHeight(r.instance, value)
}

func (r *TRichEdit) Cursor() TCursor {
    defer exceptionProc()
    return RichEdit_GetCursor(r.instance)
}

func (r *TRichEdit) SetCursor(value TCursor) {
    defer exceptionProc()
    RichEdit_SetCursor(r.instance, value)
}

func (r *TRichEdit) Hint() string {
    defer exceptionProc()
    return RichEdit_GetHint(r.instance)
}

func (r *TRichEdit) SetHint(value string) {
    defer exceptionProc()
    RichEdit_SetHint(r.instance, value)
}

func (r *TRichEdit) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(RichEdit_GetMargins(r.instance))
}

func (r *TRichEdit) SetMargins(value *TMargins) {
    defer exceptionProc()
    RichEdit_SetMargins(r.instance, CheckPtr(value))
}

func (r *TRichEdit) ComponentCount() int32 {
    defer exceptionProc()
    return RichEdit_GetComponentCount(r.instance)
}

func (r *TRichEdit) ComponentIndex() int32 {
    defer exceptionProc()
    return RichEdit_GetComponentIndex(r.instance)
}

func (r *TRichEdit) SetComponentIndex(value int32) {
    defer exceptionProc()
    RichEdit_SetComponentIndex(r.instance, value)
}

func (r *TRichEdit) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(RichEdit_GetOwner(r.instance))
}

func (r *TRichEdit) Name() string {
    defer exceptionProc()
    return RichEdit_GetName(r.instance)
}

func (r *TRichEdit) SetName(value string) {
    defer exceptionProc()
    RichEdit_SetName(r.instance, value)
}

func (r *TRichEdit) Tag() int {
    defer exceptionProc()
    return RichEdit_GetTag(r.instance)
}

func (r *TRichEdit) SetTag(value int) {
    defer exceptionProc()
    RichEdit_SetTag(r.instance, value)
}

func (r *TRichEdit) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(RichEdit_GetControls(r.instance, Index))
}

func (r *TRichEdit) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(RichEdit_GetComponents(r.instance, AIndex))
}

