
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

type TRadioGroup struct {
    IControl
    instance uintptr
}

func NewRadioGroup(owner IComponent) *TRadioGroup {
    r := new(TRadioGroup)
    r.instance = RadioGroup_Create(owner.Instance())
    return r
}

func RadioGroupFromInst(inst uintptr) *TRadioGroup {
    r := new(TRadioGroup)
    r.instance = inst
    return r
}

func RadioGroupFromObj(obj IObject) *TRadioGroup {
    r := new(TRadioGroup)
    r.instance = CheckPtr(obj)
    return r
}

func (r *TRadioGroup) Free() {
    if r.instance != 0 {
        RadioGroup_Free(r.instance)
    }
}

func (r *TRadioGroup) Instance() uintptr {
    return r.instance
}

func (r *TRadioGroup) IsValid() bool {
    return r.instance != 0
}

func (r *TRadioGroup) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    RadioGroup_FlipChildren(r.instance, AllLevels )
}

func (r *TRadioGroup) CanFocus() bool {
    defer exceptionProc()
    return RadioGroup_CanFocus(r.instance)
}

func (r *TRadioGroup) Focused() bool {
    defer exceptionProc()
    return RadioGroup_Focused(r.instance)
}

func (r *TRadioGroup) HandleAllocated() bool {
    defer exceptionProc()
    return RadioGroup_HandleAllocated(r.instance)
}

func (r *TRadioGroup) Invalidate() {
    defer exceptionProc()
    RadioGroup_Invalidate(r.instance)
}

func (r *TRadioGroup) Realign() {
    defer exceptionProc()
    RadioGroup_Realign(r.instance)
}

func (r *TRadioGroup) Repaint() {
    defer exceptionProc()
    RadioGroup_Repaint(r.instance)
}

func (r *TRadioGroup) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    RadioGroup_ScaleBy(r.instance, M , D )
}

func (r *TRadioGroup) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    RadioGroup_SetBounds(r.instance, ALeft , ATop , AWidth , AHeight )
}

func (r *TRadioGroup) SetFocus() {
    defer exceptionProc()
    RadioGroup_SetFocus(r.instance)
}

func (r *TRadioGroup) Update() {
    defer exceptionProc()
    RadioGroup_Update(r.instance)
}

func (r *TRadioGroup) BringToFront() {
    defer exceptionProc()
    RadioGroup_BringToFront(r.instance)
}

func (r *TRadioGroup) HasParent() bool {
    defer exceptionProc()
    return RadioGroup_HasParent(r.instance)
}

func (r *TRadioGroup) Hide() {
    defer exceptionProc()
    RadioGroup_Hide(r.instance)
}

func (r *TRadioGroup) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return RadioGroup_Perform(r.instance, Msg , WParam , LParam )
}

func (r *TRadioGroup) Refresh() {
    defer exceptionProc()
    RadioGroup_Refresh(r.instance)
}

func (r *TRadioGroup) SendToBack() {
    defer exceptionProc()
    RadioGroup_SendToBack(r.instance)
}

func (r *TRadioGroup) Show() {
    defer exceptionProc()
    RadioGroup_Show(r.instance)
}

func (r *TRadioGroup) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return RadioGroup_GetTextBuf(r.instance, Buffer , BufSize )
}

func (r *TRadioGroup) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(RadioGroup_FindComponent(r.instance, AName ))
}

func (r *TRadioGroup) GetNamePath() string {
    defer exceptionProc()
    return RadioGroup_GetNamePath(r.instance)
}

func (r *TRadioGroup) Assign(Source IObject) {
    defer exceptionProc()
    RadioGroup_Assign(r.instance, CheckPtr(Source))
}

func (r *TRadioGroup) ClassName() string {
    defer exceptionProc()
    return RadioGroup_ClassName(r.instance)
}

func (r *TRadioGroup) Equals(Obj IObject) bool {
    defer exceptionProc()
    return RadioGroup_Equals(r.instance, CheckPtr(Obj))
}

func (r *TRadioGroup) GetHashCode() int32 {
    defer exceptionProc()
    return RadioGroup_GetHashCode(r.instance)
}

func (r *TRadioGroup) ToString() string {
    defer exceptionProc()
    return RadioGroup_ToString(r.instance)
}

func (r *TRadioGroup) Align() TAlign {
    defer exceptionProc()
    return RadioGroup_GetAlign(r.instance)
}

func (r *TRadioGroup) SetAlign(value TAlign) {
    defer exceptionProc()
    RadioGroup_SetAlign(r.instance, value)
}

func (r *TRadioGroup) Anchors() TAnchors {
    defer exceptionProc()
    return RadioGroup_GetAnchors(r.instance)
}

func (r *TRadioGroup) SetAnchors(value TAnchors) {
    defer exceptionProc()
    RadioGroup_SetAnchors(r.instance, value)
}

func (r *TRadioGroup) Caption() string {
    defer exceptionProc()
    return RadioGroup_GetCaption(r.instance)
}

func (r *TRadioGroup) SetCaption(value string) {
    defer exceptionProc()
    RadioGroup_SetCaption(r.instance, value)
}

func (r *TRadioGroup) Color() TColor {
    defer exceptionProc()
    return RadioGroup_GetColor(r.instance)
}

func (r *TRadioGroup) SetColor(value TColor) {
    defer exceptionProc()
    RadioGroup_SetColor(r.instance, value)
}

func (r *TRadioGroup) Columns() int32 {
    defer exceptionProc()
    return RadioGroup_GetColumns(r.instance)
}

func (r *TRadioGroup) SetColumns(value int32) {
    defer exceptionProc()
    RadioGroup_SetColumns(r.instance, value)
}

func (r *TRadioGroup) DoubleBuffered() bool {
    defer exceptionProc()
    return RadioGroup_GetDoubleBuffered(r.instance)
}

func (r *TRadioGroup) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    RadioGroup_SetDoubleBuffered(r.instance, value)
}

func (r *TRadioGroup) Enabled() bool {
    defer exceptionProc()
    return RadioGroup_GetEnabled(r.instance)
}

func (r *TRadioGroup) SetEnabled(value bool) {
    defer exceptionProc()
    RadioGroup_SetEnabled(r.instance, value)
}

func (r *TRadioGroup) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(RadioGroup_GetFont(r.instance))
}

func (r *TRadioGroup) SetFont(value *TFont) {
    defer exceptionProc()
    RadioGroup_SetFont(r.instance, CheckPtr(value))
}

func (r *TRadioGroup) ItemIndex() int32 {
    defer exceptionProc()
    return RadioGroup_GetItemIndex(r.instance)
}

func (r *TRadioGroup) SetItemIndex(value int32) {
    defer exceptionProc()
    RadioGroup_SetItemIndex(r.instance, value)
}

func (r *TRadioGroup) Items() *TStrings {
    defer exceptionProc()
    return StringsFromInst(RadioGroup_GetItems(r.instance))
}

func (r *TRadioGroup) SetItems(value IObject) {
    defer exceptionProc()
    RadioGroup_SetItems(r.instance, CheckPtr(value))
}

func (r *TRadioGroup) ParentColor() bool {
    defer exceptionProc()
    return RadioGroup_GetParentColor(r.instance)
}

func (r *TRadioGroup) SetParentColor(value bool) {
    defer exceptionProc()
    RadioGroup_SetParentColor(r.instance, value)
}

func (r *TRadioGroup) ParentCtl3D() bool {
    defer exceptionProc()
    return RadioGroup_GetParentCtl3D(r.instance)
}

func (r *TRadioGroup) SetParentCtl3D(value bool) {
    defer exceptionProc()
    RadioGroup_SetParentCtl3D(r.instance, value)
}

func (r *TRadioGroup) ParentFont() bool {
    defer exceptionProc()
    return RadioGroup_GetParentFont(r.instance)
}

func (r *TRadioGroup) SetParentFont(value bool) {
    defer exceptionProc()
    RadioGroup_SetParentFont(r.instance, value)
}

func (r *TRadioGroup) ParentShowHint() bool {
    defer exceptionProc()
    return RadioGroup_GetParentShowHint(r.instance)
}

func (r *TRadioGroup) SetParentShowHint(value bool) {
    defer exceptionProc()
    RadioGroup_SetParentShowHint(r.instance, value)
}

func (r *TRadioGroup) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(RadioGroup_GetPopupMenu(r.instance))
}

func (r *TRadioGroup) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    RadioGroup_SetPopupMenu(r.instance, CheckPtr(value))
}

func (r *TRadioGroup) ShowHint() bool {
    defer exceptionProc()
    return RadioGroup_GetShowHint(r.instance)
}

func (r *TRadioGroup) SetShowHint(value bool) {
    defer exceptionProc()
    RadioGroup_SetShowHint(r.instance, value)
}

func (r *TRadioGroup) TabOrder() int16 {
    defer exceptionProc()
    return RadioGroup_GetTabOrder(r.instance)
}

func (r *TRadioGroup) SetTabOrder(value int16) {
    defer exceptionProc()
    RadioGroup_SetTabOrder(r.instance, value)
}

func (r *TRadioGroup) TabStop() bool {
    defer exceptionProc()
    return RadioGroup_GetTabStop(r.instance)
}

func (r *TRadioGroup) SetTabStop(value bool) {
    defer exceptionProc()
    RadioGroup_SetTabStop(r.instance, value)
}

func (r *TRadioGroup) Visible() bool {
    defer exceptionProc()
    return RadioGroup_GetVisible(r.instance)
}

func (r *TRadioGroup) SetVisible(value bool) {
    defer exceptionProc()
    RadioGroup_SetVisible(r.instance, value)
}

func (r *TRadioGroup) WordWrap() bool {
    defer exceptionProc()
    return RadioGroup_GetWordWrap(r.instance)
}

func (r *TRadioGroup) SetWordWrap(value bool) {
    defer exceptionProc()
    RadioGroup_SetWordWrap(r.instance, value)
}

func (r *TRadioGroup) SetOnClick(fn TNotifyEvent) {
    RadioGroup_SetOnClick(r.instance, fn)
}

func (r *TRadioGroup) SetOnEnter(fn TNotifyEvent) {
    RadioGroup_SetOnEnter(r.instance, fn)
}

func (r *TRadioGroup) SetOnExit(fn TNotifyEvent) {
    RadioGroup_SetOnExit(r.instance, fn)
}

func (r *TRadioGroup) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(RadioGroup_GetBrush(r.instance))
}

func (r *TRadioGroup) ControlCount() int32 {
    defer exceptionProc()
    return RadioGroup_GetControlCount(r.instance)
}

func (r *TRadioGroup) Handle() HWND {
    defer exceptionProc()
    return RadioGroup_GetHandle(r.instance)
}

func (r *TRadioGroup) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(RadioGroup_GetAction(r.instance))
}

func (r *TRadioGroup) SetAction(value IComponent) {
    defer exceptionProc()
    RadioGroup_SetAction(r.instance, CheckPtr(value))
}

func (r *TRadioGroup) BoundsRect() TRect {
    defer exceptionProc()
    return RadioGroup_GetBoundsRect(r.instance)
}

func (r *TRadioGroup) SetBoundsRect(value TRect) {
    defer exceptionProc()
    RadioGroup_SetBoundsRect(r.instance, value)
}

func (r *TRadioGroup) ClientHeight() int32 {
    defer exceptionProc()
    return RadioGroup_GetClientHeight(r.instance)
}

func (r *TRadioGroup) SetClientHeight(value int32) {
    defer exceptionProc()
    RadioGroup_SetClientHeight(r.instance, value)
}

func (r *TRadioGroup) ClientRect() TRect {
    defer exceptionProc()
    return RadioGroup_GetClientRect(r.instance)
}

func (r *TRadioGroup) ClientWidth() int32 {
    defer exceptionProc()
    return RadioGroup_GetClientWidth(r.instance)
}

func (r *TRadioGroup) SetClientWidth(value int32) {
    defer exceptionProc()
    RadioGroup_SetClientWidth(r.instance, value)
}

func (r *TRadioGroup) ExplicitLeft() int32 {
    defer exceptionProc()
    return RadioGroup_GetExplicitLeft(r.instance)
}

func (r *TRadioGroup) ExplicitTop() int32 {
    defer exceptionProc()
    return RadioGroup_GetExplicitTop(r.instance)
}

func (r *TRadioGroup) ExplicitWidth() int32 {
    defer exceptionProc()
    return RadioGroup_GetExplicitWidth(r.instance)
}

func (r *TRadioGroup) ExplicitHeight() int32 {
    defer exceptionProc()
    return RadioGroup_GetExplicitHeight(r.instance)
}

func (r *TRadioGroup) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(RadioGroup_GetParent(r.instance))
}

func (r *TRadioGroup) SetParent(value IControl) {
    defer exceptionProc()
    RadioGroup_SetParent(r.instance, CheckPtr(value))
}

func (r *TRadioGroup) AlignWithMargins() bool {
    defer exceptionProc()
    return RadioGroup_GetAlignWithMargins(r.instance)
}

func (r *TRadioGroup) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    RadioGroup_SetAlignWithMargins(r.instance, value)
}

func (r *TRadioGroup) Left() int32 {
    defer exceptionProc()
    return RadioGroup_GetLeft(r.instance)
}

func (r *TRadioGroup) SetLeft(value int32) {
    defer exceptionProc()
    RadioGroup_SetLeft(r.instance, value)
}

func (r *TRadioGroup) Top() int32 {
    defer exceptionProc()
    return RadioGroup_GetTop(r.instance)
}

func (r *TRadioGroup) SetTop(value int32) {
    defer exceptionProc()
    RadioGroup_SetTop(r.instance, value)
}

func (r *TRadioGroup) Width() int32 {
    defer exceptionProc()
    return RadioGroup_GetWidth(r.instance)
}

func (r *TRadioGroup) SetWidth(value int32) {
    defer exceptionProc()
    RadioGroup_SetWidth(r.instance, value)
}

func (r *TRadioGroup) Height() int32 {
    defer exceptionProc()
    return RadioGroup_GetHeight(r.instance)
}

func (r *TRadioGroup) SetHeight(value int32) {
    defer exceptionProc()
    RadioGroup_SetHeight(r.instance, value)
}

func (r *TRadioGroup) Cursor() TCursor {
    defer exceptionProc()
    return RadioGroup_GetCursor(r.instance)
}

func (r *TRadioGroup) SetCursor(value TCursor) {
    defer exceptionProc()
    RadioGroup_SetCursor(r.instance, value)
}

func (r *TRadioGroup) Hint() string {
    defer exceptionProc()
    return RadioGroup_GetHint(r.instance)
}

func (r *TRadioGroup) SetHint(value string) {
    defer exceptionProc()
    RadioGroup_SetHint(r.instance, value)
}

func (r *TRadioGroup) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(RadioGroup_GetMargins(r.instance))
}

func (r *TRadioGroup) SetMargins(value *TMargins) {
    defer exceptionProc()
    RadioGroup_SetMargins(r.instance, CheckPtr(value))
}

func (r *TRadioGroup) ComponentCount() int32 {
    defer exceptionProc()
    return RadioGroup_GetComponentCount(r.instance)
}

func (r *TRadioGroup) ComponentIndex() int32 {
    defer exceptionProc()
    return RadioGroup_GetComponentIndex(r.instance)
}

func (r *TRadioGroup) SetComponentIndex(value int32) {
    defer exceptionProc()
    RadioGroup_SetComponentIndex(r.instance, value)
}

func (r *TRadioGroup) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(RadioGroup_GetOwner(r.instance))
}

func (r *TRadioGroup) Name() string {
    defer exceptionProc()
    return RadioGroup_GetName(r.instance)
}

func (r *TRadioGroup) SetName(value string) {
    defer exceptionProc()
    RadioGroup_SetName(r.instance, value)
}

func (r *TRadioGroup) Tag() int {
    defer exceptionProc()
    return RadioGroup_GetTag(r.instance)
}

func (r *TRadioGroup) SetTag(value int) {
    defer exceptionProc()
    RadioGroup_SetTag(r.instance, value)
}

func (r *TRadioGroup) Buttons(Index int32) *TRadioButton {
    defer exceptionProc()
    return RadioButtonFromInst(RadioGroup_GetButtons(r.instance, Index))
}

func (r *TRadioGroup) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(RadioGroup_GetControls(r.instance, Index))
}

func (r *TRadioGroup) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(RadioGroup_GetComponents(r.instance, AIndex))
}

