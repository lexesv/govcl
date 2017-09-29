
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TTabSheet struct {
    IControl
    instance uintptr
}

func NewTabSheet(owner IComponent) *TTabSheet {
    t := new(TTabSheet)
    t.instance = TabSheet_Create(owner.Instance())
    return t
}

func TabSheetFromInst(inst uintptr) *TTabSheet {
    t := new(TTabSheet)
    t.instance = inst
    return t
}

func TabSheetFromObj(obj IObject) *TTabSheet {
    t := new(TTabSheet)
    t.instance = CheckPtr(obj)
    return t
}

func (t *TTabSheet) Free() {
    if t.instance != 0 {
        TabSheet_Free(t.instance)
    }
}

func (t *TTabSheet) Instance() uintptr {
    return t.instance
}

func (t *TTabSheet) IsValid() bool {
    return t.instance != 0
}

func (t *TTabSheet) CanFocus() bool {
    defer exceptionProc()
    return TabSheet_CanFocus(t.instance)
}

func (t *TTabSheet) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    TabSheet_FlipChildren(t.instance, AllLevels )
}

func (t *TTabSheet) Focused() bool {
    defer exceptionProc()
    return TabSheet_Focused(t.instance)
}

func (t *TTabSheet) HandleAllocated() bool {
    defer exceptionProc()
    return TabSheet_HandleAllocated(t.instance)
}

func (t *TTabSheet) Invalidate() {
    defer exceptionProc()
    TabSheet_Invalidate(t.instance)
}

func (t *TTabSheet) Realign() {
    defer exceptionProc()
    TabSheet_Realign(t.instance)
}

func (t *TTabSheet) Repaint() {
    defer exceptionProc()
    TabSheet_Repaint(t.instance)
}

func (t *TTabSheet) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    TabSheet_ScaleBy(t.instance, M , D )
}

func (t *TTabSheet) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    TabSheet_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight )
}

func (t *TTabSheet) SetFocus() {
    defer exceptionProc()
    TabSheet_SetFocus(t.instance)
}

func (t *TTabSheet) Update() {
    defer exceptionProc()
    TabSheet_Update(t.instance)
}

func (t *TTabSheet) BringToFront() {
    defer exceptionProc()
    TabSheet_BringToFront(t.instance)
}

func (t *TTabSheet) HasParent() bool {
    defer exceptionProc()
    return TabSheet_HasParent(t.instance)
}

func (t *TTabSheet) Hide() {
    defer exceptionProc()
    TabSheet_Hide(t.instance)
}

func (t *TTabSheet) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return TabSheet_Perform(t.instance, Msg , WParam , LParam )
}

func (t *TTabSheet) Refresh() {
    defer exceptionProc()
    TabSheet_Refresh(t.instance)
}

func (t *TTabSheet) SendToBack() {
    defer exceptionProc()
    TabSheet_SendToBack(t.instance)
}

func (t *TTabSheet) Show() {
    defer exceptionProc()
    TabSheet_Show(t.instance)
}

func (t *TTabSheet) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return TabSheet_GetTextBuf(t.instance, Buffer , BufSize )
}

func (t *TTabSheet) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(TabSheet_FindComponent(t.instance, AName ))
}

func (t *TTabSheet) GetNamePath() string {
    defer exceptionProc()
    return TabSheet_GetNamePath(t.instance)
}

func (t *TTabSheet) Assign(Source IObject) {
    defer exceptionProc()
    TabSheet_Assign(t.instance, CheckPtr(Source))
}

func (t *TTabSheet) ClassName() string {
    defer exceptionProc()
    return TabSheet_ClassName(t.instance)
}

func (t *TTabSheet) Equals(Obj IObject) bool {
    defer exceptionProc()
    return TabSheet_Equals(t.instance, CheckPtr(Obj))
}

func (t *TTabSheet) GetHashCode() int32 {
    defer exceptionProc()
    return TabSheet_GetHashCode(t.instance)
}

func (t *TTabSheet) ToString() string {
    defer exceptionProc()
    return TabSheet_ToString(t.instance)
}

func (t *TTabSheet) PageControl() *TPageControl {
    defer exceptionProc()
    return PageControlFromInst(TabSheet_GetPageControl(t.instance))
}

func (t *TTabSheet) SetPageControl(value IControl) {
    defer exceptionProc()
    TabSheet_SetPageControl(t.instance, CheckPtr(value))
}

func (t *TTabSheet) TabIndex() int32 {
    defer exceptionProc()
    return TabSheet_GetTabIndex(t.instance)
}

func (t *TTabSheet) BorderWidth() int32 {
    defer exceptionProc()
    return TabSheet_GetBorderWidth(t.instance)
}

func (t *TTabSheet) SetBorderWidth(value int32) {
    defer exceptionProc()
    TabSheet_SetBorderWidth(t.instance, value)
}

func (t *TTabSheet) Caption() string {
    defer exceptionProc()
    return TabSheet_GetCaption(t.instance)
}

func (t *TTabSheet) SetCaption(value string) {
    defer exceptionProc()
    TabSheet_SetCaption(t.instance, value)
}

func (t *TTabSheet) DoubleBuffered() bool {
    defer exceptionProc()
    return TabSheet_GetDoubleBuffered(t.instance)
}

func (t *TTabSheet) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    TabSheet_SetDoubleBuffered(t.instance, value)
}

func (t *TTabSheet) Enabled() bool {
    defer exceptionProc()
    return TabSheet_GetEnabled(t.instance)
}

func (t *TTabSheet) SetEnabled(value bool) {
    defer exceptionProc()
    TabSheet_SetEnabled(t.instance, value)
}

func (t *TTabSheet) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(TabSheet_GetFont(t.instance))
}

func (t *TTabSheet) SetFont(value *TFont) {
    defer exceptionProc()
    TabSheet_SetFont(t.instance, CheckPtr(value))
}

func (t *TTabSheet) Height() int32 {
    defer exceptionProc()
    return TabSheet_GetHeight(t.instance)
}

func (t *TTabSheet) SetHeight(value int32) {
    defer exceptionProc()
    TabSheet_SetHeight(t.instance, value)
}

func (t *TTabSheet) Highlighted() bool {
    defer exceptionProc()
    return TabSheet_GetHighlighted(t.instance)
}

func (t *TTabSheet) SetHighlighted(value bool) {
    defer exceptionProc()
    TabSheet_SetHighlighted(t.instance, value)
}

func (t *TTabSheet) ImageIndex() int32 {
    defer exceptionProc()
    return TabSheet_GetImageIndex(t.instance)
}

func (t *TTabSheet) SetImageIndex(value int32) {
    defer exceptionProc()
    TabSheet_SetImageIndex(t.instance, value)
}

func (t *TTabSheet) Left() int32 {
    defer exceptionProc()
    return TabSheet_GetLeft(t.instance)
}

func (t *TTabSheet) SetLeft(value int32) {
    defer exceptionProc()
    TabSheet_SetLeft(t.instance, value)
}

func (t *TTabSheet) PageIndex() int32 {
    defer exceptionProc()
    return TabSheet_GetPageIndex(t.instance)
}

func (t *TTabSheet) SetPageIndex(value int32) {
    defer exceptionProc()
    TabSheet_SetPageIndex(t.instance, value)
}

func (t *TTabSheet) ParentFont() bool {
    defer exceptionProc()
    return TabSheet_GetParentFont(t.instance)
}

func (t *TTabSheet) SetParentFont(value bool) {
    defer exceptionProc()
    TabSheet_SetParentFont(t.instance, value)
}

func (t *TTabSheet) ParentShowHint() bool {
    defer exceptionProc()
    return TabSheet_GetParentShowHint(t.instance)
}

func (t *TTabSheet) SetParentShowHint(value bool) {
    defer exceptionProc()
    TabSheet_SetParentShowHint(t.instance, value)
}

func (t *TTabSheet) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(TabSheet_GetPopupMenu(t.instance))
}

func (t *TTabSheet) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    TabSheet_SetPopupMenu(t.instance, CheckPtr(value))
}

func (t *TTabSheet) ShowHint() bool {
    defer exceptionProc()
    return TabSheet_GetShowHint(t.instance)
}

func (t *TTabSheet) SetShowHint(value bool) {
    defer exceptionProc()
    TabSheet_SetShowHint(t.instance, value)
}

func (t *TTabSheet) TabVisible() bool {
    defer exceptionProc()
    return TabSheet_GetTabVisible(t.instance)
}

func (t *TTabSheet) SetTabVisible(value bool) {
    defer exceptionProc()
    TabSheet_SetTabVisible(t.instance, value)
}

func (t *TTabSheet) Top() int32 {
    defer exceptionProc()
    return TabSheet_GetTop(t.instance)
}

func (t *TTabSheet) SetTop(value int32) {
    defer exceptionProc()
    TabSheet_SetTop(t.instance, value)
}

func (t *TTabSheet) Visible() bool {
    defer exceptionProc()
    return TabSheet_GetVisible(t.instance)
}

func (t *TTabSheet) SetVisible(value bool) {
    defer exceptionProc()
    TabSheet_SetVisible(t.instance, value)
}

func (t *TTabSheet) Width() int32 {
    defer exceptionProc()
    return TabSheet_GetWidth(t.instance)
}

func (t *TTabSheet) SetWidth(value int32) {
    defer exceptionProc()
    TabSheet_SetWidth(t.instance, value)
}

func (t *TTabSheet) SetOnEnter(fn TNotifyEvent) {
    TabSheet_SetOnEnter(t.instance, fn)
}

func (t *TTabSheet) SetOnExit(fn TNotifyEvent) {
    TabSheet_SetOnExit(t.instance, fn)
}

func (t *TTabSheet) SetOnHide(fn TNotifyEvent) {
    TabSheet_SetOnHide(t.instance, fn)
}

func (t *TTabSheet) SetOnMouseDown(fn TMouseEvent) {
    TabSheet_SetOnMouseDown(t.instance, fn)
}

func (t *TTabSheet) SetOnMouseEnter(fn TNotifyEvent) {
    TabSheet_SetOnMouseEnter(t.instance, fn)
}

func (t *TTabSheet) SetOnMouseLeave(fn TNotifyEvent) {
    TabSheet_SetOnMouseLeave(t.instance, fn)
}

func (t *TTabSheet) SetOnMouseMove(fn TMouseMoveEvent) {
    TabSheet_SetOnMouseMove(t.instance, fn)
}

func (t *TTabSheet) SetOnMouseUp(fn TMouseEvent) {
    TabSheet_SetOnMouseUp(t.instance, fn)
}

func (t *TTabSheet) SetOnResize(fn TNotifyEvent) {
    TabSheet_SetOnResize(t.instance, fn)
}

func (t *TTabSheet) SetOnShow(fn TNotifyEvent) {
    TabSheet_SetOnShow(t.instance, fn)
}

func (t *TTabSheet) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(TabSheet_GetBrush(t.instance))
}

func (t *TTabSheet) ControlCount() int32 {
    defer exceptionProc()
    return TabSheet_GetControlCount(t.instance)
}

func (t *TTabSheet) Handle() HWND {
    defer exceptionProc()
    return TabSheet_GetHandle(t.instance)
}

func (t *TTabSheet) TabOrder() int16 {
    defer exceptionProc()
    return TabSheet_GetTabOrder(t.instance)
}

func (t *TTabSheet) SetTabOrder(value int16) {
    defer exceptionProc()
    TabSheet_SetTabOrder(t.instance, value)
}

func (t *TTabSheet) TabStop() bool {
    defer exceptionProc()
    return TabSheet_GetTabStop(t.instance)
}

func (t *TTabSheet) SetTabStop(value bool) {
    defer exceptionProc()
    TabSheet_SetTabStop(t.instance, value)
}

func (t *TTabSheet) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(TabSheet_GetAction(t.instance))
}

func (t *TTabSheet) SetAction(value IComponent) {
    defer exceptionProc()
    TabSheet_SetAction(t.instance, CheckPtr(value))
}

func (t *TTabSheet) Align() TAlign {
    defer exceptionProc()
    return TabSheet_GetAlign(t.instance)
}

func (t *TTabSheet) SetAlign(value TAlign) {
    defer exceptionProc()
    TabSheet_SetAlign(t.instance, value)
}

func (t *TTabSheet) Anchors() TAnchors {
    defer exceptionProc()
    return TabSheet_GetAnchors(t.instance)
}

func (t *TTabSheet) SetAnchors(value TAnchors) {
    defer exceptionProc()
    TabSheet_SetAnchors(t.instance, value)
}

func (t *TTabSheet) BoundsRect() TRect {
    defer exceptionProc()
    return TabSheet_GetBoundsRect(t.instance)
}

func (t *TTabSheet) SetBoundsRect(value TRect) {
    defer exceptionProc()
    TabSheet_SetBoundsRect(t.instance, value)
}

func (t *TTabSheet) ClientHeight() int32 {
    defer exceptionProc()
    return TabSheet_GetClientHeight(t.instance)
}

func (t *TTabSheet) SetClientHeight(value int32) {
    defer exceptionProc()
    TabSheet_SetClientHeight(t.instance, value)
}

func (t *TTabSheet) ClientRect() TRect {
    defer exceptionProc()
    return TabSheet_GetClientRect(t.instance)
}

func (t *TTabSheet) ClientWidth() int32 {
    defer exceptionProc()
    return TabSheet_GetClientWidth(t.instance)
}

func (t *TTabSheet) SetClientWidth(value int32) {
    defer exceptionProc()
    TabSheet_SetClientWidth(t.instance, value)
}

func (t *TTabSheet) ExplicitLeft() int32 {
    defer exceptionProc()
    return TabSheet_GetExplicitLeft(t.instance)
}

func (t *TTabSheet) ExplicitTop() int32 {
    defer exceptionProc()
    return TabSheet_GetExplicitTop(t.instance)
}

func (t *TTabSheet) ExplicitWidth() int32 {
    defer exceptionProc()
    return TabSheet_GetExplicitWidth(t.instance)
}

func (t *TTabSheet) ExplicitHeight() int32 {
    defer exceptionProc()
    return TabSheet_GetExplicitHeight(t.instance)
}

func (t *TTabSheet) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(TabSheet_GetParent(t.instance))
}

func (t *TTabSheet) SetParent(value IControl) {
    defer exceptionProc()
    TabSheet_SetParent(t.instance, CheckPtr(value))
}

func (t *TTabSheet) AlignWithMargins() bool {
    defer exceptionProc()
    return TabSheet_GetAlignWithMargins(t.instance)
}

func (t *TTabSheet) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    TabSheet_SetAlignWithMargins(t.instance, value)
}

func (t *TTabSheet) Cursor() TCursor {
    defer exceptionProc()
    return TabSheet_GetCursor(t.instance)
}

func (t *TTabSheet) SetCursor(value TCursor) {
    defer exceptionProc()
    TabSheet_SetCursor(t.instance, value)
}

func (t *TTabSheet) Hint() string {
    defer exceptionProc()
    return TabSheet_GetHint(t.instance)
}

func (t *TTabSheet) SetHint(value string) {
    defer exceptionProc()
    TabSheet_SetHint(t.instance, value)
}

func (t *TTabSheet) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(TabSheet_GetMargins(t.instance))
}

func (t *TTabSheet) SetMargins(value *TMargins) {
    defer exceptionProc()
    TabSheet_SetMargins(t.instance, CheckPtr(value))
}

func (t *TTabSheet) ComponentCount() int32 {
    defer exceptionProc()
    return TabSheet_GetComponentCount(t.instance)
}

func (t *TTabSheet) ComponentIndex() int32 {
    defer exceptionProc()
    return TabSheet_GetComponentIndex(t.instance)
}

func (t *TTabSheet) SetComponentIndex(value int32) {
    defer exceptionProc()
    TabSheet_SetComponentIndex(t.instance, value)
}

func (t *TTabSheet) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(TabSheet_GetOwner(t.instance))
}

func (t *TTabSheet) Name() string {
    defer exceptionProc()
    return TabSheet_GetName(t.instance)
}

func (t *TTabSheet) SetName(value string) {
    defer exceptionProc()
    TabSheet_SetName(t.instance, value)
}

func (t *TTabSheet) Tag() int {
    defer exceptionProc()
    return TabSheet_GetTag(t.instance)
}

func (t *TTabSheet) SetTag(value int) {
    defer exceptionProc()
    TabSheet_SetTag(t.instance, value)
}

func (t *TTabSheet) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(TabSheet_GetControls(t.instance, Index))
}

func (t *TTabSheet) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(TabSheet_GetComponents(t.instance, AIndex))
}

