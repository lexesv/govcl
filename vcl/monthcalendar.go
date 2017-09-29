
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TMonthCalendar struct {
    IControl
    instance uintptr
}

func NewMonthCalendar(owner IComponent) *TMonthCalendar {
    m := new(TMonthCalendar)
    m.instance = MonthCalendar_Create(owner.Instance())
    return m
}

func MonthCalendarFromInst(inst uintptr) *TMonthCalendar {
    m := new(TMonthCalendar)
    m.instance = inst
    return m
}

func MonthCalendarFromObj(obj IObject) *TMonthCalendar {
    m := new(TMonthCalendar)
    m.instance = CheckPtr(obj)
    return m
}

func (m *TMonthCalendar) Free() {
    if m.instance != 0 {
        MonthCalendar_Free(m.instance)
    }
}

func (m *TMonthCalendar) Instance() uintptr {
    return m.instance
}

func (m *TMonthCalendar) IsValid() bool {
    return m.instance != 0
}

func (m *TMonthCalendar) CanFocus() bool {
    defer exceptionProc()
    return MonthCalendar_CanFocus(m.instance)
}

func (m *TMonthCalendar) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    MonthCalendar_FlipChildren(m.instance, AllLevels )
}

func (m *TMonthCalendar) Focused() bool {
    defer exceptionProc()
    return MonthCalendar_Focused(m.instance)
}

func (m *TMonthCalendar) HandleAllocated() bool {
    defer exceptionProc()
    return MonthCalendar_HandleAllocated(m.instance)
}

func (m *TMonthCalendar) Invalidate() {
    defer exceptionProc()
    MonthCalendar_Invalidate(m.instance)
}

func (m *TMonthCalendar) Realign() {
    defer exceptionProc()
    MonthCalendar_Realign(m.instance)
}

func (m *TMonthCalendar) Repaint() {
    defer exceptionProc()
    MonthCalendar_Repaint(m.instance)
}

func (m *TMonthCalendar) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    MonthCalendar_ScaleBy(m.instance, M , D )
}

func (m *TMonthCalendar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    MonthCalendar_SetBounds(m.instance, ALeft , ATop , AWidth , AHeight )
}

func (m *TMonthCalendar) SetFocus() {
    defer exceptionProc()
    MonthCalendar_SetFocus(m.instance)
}

func (m *TMonthCalendar) Update() {
    defer exceptionProc()
    MonthCalendar_Update(m.instance)
}

func (m *TMonthCalendar) BringToFront() {
    defer exceptionProc()
    MonthCalendar_BringToFront(m.instance)
}

func (m *TMonthCalendar) HasParent() bool {
    defer exceptionProc()
    return MonthCalendar_HasParent(m.instance)
}

func (m *TMonthCalendar) Hide() {
    defer exceptionProc()
    MonthCalendar_Hide(m.instance)
}

func (m *TMonthCalendar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return MonthCalendar_Perform(m.instance, Msg , WParam , LParam )
}

func (m *TMonthCalendar) Refresh() {
    defer exceptionProc()
    MonthCalendar_Refresh(m.instance)
}

func (m *TMonthCalendar) SendToBack() {
    defer exceptionProc()
    MonthCalendar_SendToBack(m.instance)
}

func (m *TMonthCalendar) Show() {
    defer exceptionProc()
    MonthCalendar_Show(m.instance)
}

func (m *TMonthCalendar) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return MonthCalendar_GetTextBuf(m.instance, Buffer , BufSize )
}

func (m *TMonthCalendar) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(MonthCalendar_FindComponent(m.instance, AName ))
}

func (m *TMonthCalendar) GetNamePath() string {
    defer exceptionProc()
    return MonthCalendar_GetNamePath(m.instance)
}

func (m *TMonthCalendar) Assign(Source IObject) {
    defer exceptionProc()
    MonthCalendar_Assign(m.instance, CheckPtr(Source))
}

func (m *TMonthCalendar) ClassName() string {
    defer exceptionProc()
    return MonthCalendar_ClassName(m.instance)
}

func (m *TMonthCalendar) Equals(Obj IObject) bool {
    defer exceptionProc()
    return MonthCalendar_Equals(m.instance, CheckPtr(Obj))
}

func (m *TMonthCalendar) GetHashCode() int32 {
    defer exceptionProc()
    return MonthCalendar_GetHashCode(m.instance)
}

func (m *TMonthCalendar) ToString() string {
    defer exceptionProc()
    return MonthCalendar_ToString(m.instance)
}

func (m *TMonthCalendar) Align() TAlign {
    defer exceptionProc()
    return MonthCalendar_GetAlign(m.instance)
}

func (m *TMonthCalendar) SetAlign(value TAlign) {
    defer exceptionProc()
    MonthCalendar_SetAlign(m.instance, value)
}

func (m *TMonthCalendar) Anchors() TAnchors {
    defer exceptionProc()
    return MonthCalendar_GetAnchors(m.instance)
}

func (m *TMonthCalendar) SetAnchors(value TAnchors) {
    defer exceptionProc()
    MonthCalendar_SetAnchors(m.instance, value)
}

func (m *TMonthCalendar) AutoSize() bool {
    defer exceptionProc()
    return MonthCalendar_GetAutoSize(m.instance)
}

func (m *TMonthCalendar) SetAutoSize(value bool) {
    defer exceptionProc()
    MonthCalendar_SetAutoSize(m.instance, value)
}

func (m *TMonthCalendar) BorderWidth() int32 {
    defer exceptionProc()
    return MonthCalendar_GetBorderWidth(m.instance)
}

func (m *TMonthCalendar) SetBorderWidth(value int32) {
    defer exceptionProc()
    MonthCalendar_SetBorderWidth(m.instance, value)
}

func (m *TMonthCalendar) MultiSelect() bool {
    defer exceptionProc()
    return MonthCalendar_GetMultiSelect(m.instance)
}

func (m *TMonthCalendar) SetMultiSelect(value bool) {
    defer exceptionProc()
    MonthCalendar_SetMultiSelect(m.instance, value)
}

func (m *TMonthCalendar) DoubleBuffered() bool {
    defer exceptionProc()
    return MonthCalendar_GetDoubleBuffered(m.instance)
}

func (m *TMonthCalendar) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    MonthCalendar_SetDoubleBuffered(m.instance, value)
}

func (m *TMonthCalendar) Enabled() bool {
    defer exceptionProc()
    return MonthCalendar_GetEnabled(m.instance)
}

func (m *TMonthCalendar) SetEnabled(value bool) {
    defer exceptionProc()
    MonthCalendar_SetEnabled(m.instance, value)
}

func (m *TMonthCalendar) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(MonthCalendar_GetFont(m.instance))
}

func (m *TMonthCalendar) SetFont(value *TFont) {
    defer exceptionProc()
    MonthCalendar_SetFont(m.instance, CheckPtr(value))
}

func (m *TMonthCalendar) ParentFont() bool {
    defer exceptionProc()
    return MonthCalendar_GetParentFont(m.instance)
}

func (m *TMonthCalendar) SetParentFont(value bool) {
    defer exceptionProc()
    MonthCalendar_SetParentFont(m.instance, value)
}

func (m *TMonthCalendar) ParentShowHint() bool {
    defer exceptionProc()
    return MonthCalendar_GetParentShowHint(m.instance)
}

func (m *TMonthCalendar) SetParentShowHint(value bool) {
    defer exceptionProc()
    MonthCalendar_SetParentShowHint(m.instance, value)
}

func (m *TMonthCalendar) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(MonthCalendar_GetPopupMenu(m.instance))
}

func (m *TMonthCalendar) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    MonthCalendar_SetPopupMenu(m.instance, CheckPtr(value))
}

func (m *TMonthCalendar) ShowHint() bool {
    defer exceptionProc()
    return MonthCalendar_GetShowHint(m.instance)
}

func (m *TMonthCalendar) SetShowHint(value bool) {
    defer exceptionProc()
    MonthCalendar_SetShowHint(m.instance, value)
}

func (m *TMonthCalendar) TabOrder() int16 {
    defer exceptionProc()
    return MonthCalendar_GetTabOrder(m.instance)
}

func (m *TMonthCalendar) SetTabOrder(value int16) {
    defer exceptionProc()
    MonthCalendar_SetTabOrder(m.instance, value)
}

func (m *TMonthCalendar) TabStop() bool {
    defer exceptionProc()
    return MonthCalendar_GetTabStop(m.instance)
}

func (m *TMonthCalendar) SetTabStop(value bool) {
    defer exceptionProc()
    MonthCalendar_SetTabStop(m.instance, value)
}

func (m *TMonthCalendar) Visible() bool {
    defer exceptionProc()
    return MonthCalendar_GetVisible(m.instance)
}

func (m *TMonthCalendar) SetVisible(value bool) {
    defer exceptionProc()
    MonthCalendar_SetVisible(m.instance, value)
}

func (m *TMonthCalendar) SetOnClick(fn TNotifyEvent) {
    MonthCalendar_SetOnClick(m.instance, fn)
}

func (m *TMonthCalendar) SetOnDblClick(fn TNotifyEvent) {
    MonthCalendar_SetOnDblClick(m.instance, fn)
}

func (m *TMonthCalendar) SetOnEnter(fn TNotifyEvent) {
    MonthCalendar_SetOnEnter(m.instance, fn)
}

func (m *TMonthCalendar) SetOnExit(fn TNotifyEvent) {
    MonthCalendar_SetOnExit(m.instance, fn)
}

func (m *TMonthCalendar) SetOnKeyDown(fn TKeyEvent) {
    MonthCalendar_SetOnKeyDown(m.instance, fn)
}

func (m *TMonthCalendar) SetOnKeyPress(fn TKeyPressEvent) {
    MonthCalendar_SetOnKeyPress(m.instance, fn)
}

func (m *TMonthCalendar) SetOnKeyUp(fn TKeyEvent) {
    MonthCalendar_SetOnKeyUp(m.instance, fn)
}

func (m *TMonthCalendar) SetOnMouseEnter(fn TNotifyEvent) {
    MonthCalendar_SetOnMouseEnter(m.instance, fn)
}

func (m *TMonthCalendar) SetOnMouseLeave(fn TNotifyEvent) {
    MonthCalendar_SetOnMouseLeave(m.instance, fn)
}

func (m *TMonthCalendar) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(MonthCalendar_GetBrush(m.instance))
}

func (m *TMonthCalendar) ControlCount() int32 {
    defer exceptionProc()
    return MonthCalendar_GetControlCount(m.instance)
}

func (m *TMonthCalendar) Handle() HWND {
    defer exceptionProc()
    return MonthCalendar_GetHandle(m.instance)
}

func (m *TMonthCalendar) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(MonthCalendar_GetAction(m.instance))
}

func (m *TMonthCalendar) SetAction(value IComponent) {
    defer exceptionProc()
    MonthCalendar_SetAction(m.instance, CheckPtr(value))
}

func (m *TMonthCalendar) BoundsRect() TRect {
    defer exceptionProc()
    return MonthCalendar_GetBoundsRect(m.instance)
}

func (m *TMonthCalendar) SetBoundsRect(value TRect) {
    defer exceptionProc()
    MonthCalendar_SetBoundsRect(m.instance, value)
}

func (m *TMonthCalendar) ClientHeight() int32 {
    defer exceptionProc()
    return MonthCalendar_GetClientHeight(m.instance)
}

func (m *TMonthCalendar) SetClientHeight(value int32) {
    defer exceptionProc()
    MonthCalendar_SetClientHeight(m.instance, value)
}

func (m *TMonthCalendar) ClientRect() TRect {
    defer exceptionProc()
    return MonthCalendar_GetClientRect(m.instance)
}

func (m *TMonthCalendar) ClientWidth() int32 {
    defer exceptionProc()
    return MonthCalendar_GetClientWidth(m.instance)
}

func (m *TMonthCalendar) SetClientWidth(value int32) {
    defer exceptionProc()
    MonthCalendar_SetClientWidth(m.instance, value)
}

func (m *TMonthCalendar) ExplicitLeft() int32 {
    defer exceptionProc()
    return MonthCalendar_GetExplicitLeft(m.instance)
}

func (m *TMonthCalendar) ExplicitTop() int32 {
    defer exceptionProc()
    return MonthCalendar_GetExplicitTop(m.instance)
}

func (m *TMonthCalendar) ExplicitWidth() int32 {
    defer exceptionProc()
    return MonthCalendar_GetExplicitWidth(m.instance)
}

func (m *TMonthCalendar) ExplicitHeight() int32 {
    defer exceptionProc()
    return MonthCalendar_GetExplicitHeight(m.instance)
}

func (m *TMonthCalendar) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(MonthCalendar_GetParent(m.instance))
}

func (m *TMonthCalendar) SetParent(value IControl) {
    defer exceptionProc()
    MonthCalendar_SetParent(m.instance, CheckPtr(value))
}

func (m *TMonthCalendar) AlignWithMargins() bool {
    defer exceptionProc()
    return MonthCalendar_GetAlignWithMargins(m.instance)
}

func (m *TMonthCalendar) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    MonthCalendar_SetAlignWithMargins(m.instance, value)
}

func (m *TMonthCalendar) Left() int32 {
    defer exceptionProc()
    return MonthCalendar_GetLeft(m.instance)
}

func (m *TMonthCalendar) SetLeft(value int32) {
    defer exceptionProc()
    MonthCalendar_SetLeft(m.instance, value)
}

func (m *TMonthCalendar) Top() int32 {
    defer exceptionProc()
    return MonthCalendar_GetTop(m.instance)
}

func (m *TMonthCalendar) SetTop(value int32) {
    defer exceptionProc()
    MonthCalendar_SetTop(m.instance, value)
}

func (m *TMonthCalendar) Width() int32 {
    defer exceptionProc()
    return MonthCalendar_GetWidth(m.instance)
}

func (m *TMonthCalendar) SetWidth(value int32) {
    defer exceptionProc()
    MonthCalendar_SetWidth(m.instance, value)
}

func (m *TMonthCalendar) Height() int32 {
    defer exceptionProc()
    return MonthCalendar_GetHeight(m.instance)
}

func (m *TMonthCalendar) SetHeight(value int32) {
    defer exceptionProc()
    MonthCalendar_SetHeight(m.instance, value)
}

func (m *TMonthCalendar) Cursor() TCursor {
    defer exceptionProc()
    return MonthCalendar_GetCursor(m.instance)
}

func (m *TMonthCalendar) SetCursor(value TCursor) {
    defer exceptionProc()
    MonthCalendar_SetCursor(m.instance, value)
}

func (m *TMonthCalendar) Hint() string {
    defer exceptionProc()
    return MonthCalendar_GetHint(m.instance)
}

func (m *TMonthCalendar) SetHint(value string) {
    defer exceptionProc()
    MonthCalendar_SetHint(m.instance, value)
}

func (m *TMonthCalendar) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(MonthCalendar_GetMargins(m.instance))
}

func (m *TMonthCalendar) SetMargins(value *TMargins) {
    defer exceptionProc()
    MonthCalendar_SetMargins(m.instance, CheckPtr(value))
}

func (m *TMonthCalendar) ComponentCount() int32 {
    defer exceptionProc()
    return MonthCalendar_GetComponentCount(m.instance)
}

func (m *TMonthCalendar) ComponentIndex() int32 {
    defer exceptionProc()
    return MonthCalendar_GetComponentIndex(m.instance)
}

func (m *TMonthCalendar) SetComponentIndex(value int32) {
    defer exceptionProc()
    MonthCalendar_SetComponentIndex(m.instance, value)
}

func (m *TMonthCalendar) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(MonthCalendar_GetOwner(m.instance))
}

func (m *TMonthCalendar) Name() string {
    defer exceptionProc()
    return MonthCalendar_GetName(m.instance)
}

func (m *TMonthCalendar) SetName(value string) {
    defer exceptionProc()
    MonthCalendar_SetName(m.instance, value)
}

func (m *TMonthCalendar) Tag() int {
    defer exceptionProc()
    return MonthCalendar_GetTag(m.instance)
}

func (m *TMonthCalendar) SetTag(value int) {
    defer exceptionProc()
    MonthCalendar_SetTag(m.instance, value)
}

func (m *TMonthCalendar) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(MonthCalendar_GetControls(m.instance, Index))
}

func (m *TMonthCalendar) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(MonthCalendar_GetComponents(m.instance, AIndex))
}

