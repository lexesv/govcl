
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TPageControl struct {
    IControl
    instance uintptr
}

func NewPageControl(owner IComponent) *TPageControl {
    p := new(TPageControl)
    p.instance = PageControl_Create(owner.Instance())
    return p
}

func PageControlFromInst(inst uintptr) *TPageControl {
    p := new(TPageControl)
    p.instance = inst
    return p
}

func PageControlFromObj(obj IObject) *TPageControl {
    p := new(TPageControl)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TPageControl) Free() {
    if p.instance != 0 {
        PageControl_Free(p.instance)
    }
}

func (p *TPageControl) Instance() uintptr {
    return p.instance
}

func (p *TPageControl) IsValid() bool {
    return p.instance != 0
}

func (p *TPageControl) SelectNextPage(GoForward bool, CheckTabVisible bool) {
    defer exceptionProc()
    PageControl_SelectNextPage(p.instance, GoForward , CheckTabVisible )
}

func (p *TPageControl) RowCount() int32 {
    defer exceptionProc()
    return PageControl_RowCount(p.instance)
}

func (p *TPageControl) CanFocus() bool {
    defer exceptionProc()
    return PageControl_CanFocus(p.instance)
}

func (p *TPageControl) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    PageControl_FlipChildren(p.instance, AllLevels )
}

func (p *TPageControl) Focused() bool {
    defer exceptionProc()
    return PageControl_Focused(p.instance)
}

func (p *TPageControl) HandleAllocated() bool {
    defer exceptionProc()
    return PageControl_HandleAllocated(p.instance)
}

func (p *TPageControl) Invalidate() {
    defer exceptionProc()
    PageControl_Invalidate(p.instance)
}

func (p *TPageControl) Realign() {
    defer exceptionProc()
    PageControl_Realign(p.instance)
}

func (p *TPageControl) Repaint() {
    defer exceptionProc()
    PageControl_Repaint(p.instance)
}

func (p *TPageControl) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    PageControl_ScaleBy(p.instance, M , D )
}

func (p *TPageControl) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    PageControl_SetBounds(p.instance, ALeft , ATop , AWidth , AHeight )
}

func (p *TPageControl) SetFocus() {
    defer exceptionProc()
    PageControl_SetFocus(p.instance)
}

func (p *TPageControl) Update() {
    defer exceptionProc()
    PageControl_Update(p.instance)
}

func (p *TPageControl) BringToFront() {
    defer exceptionProc()
    PageControl_BringToFront(p.instance)
}

func (p *TPageControl) HasParent() bool {
    defer exceptionProc()
    return PageControl_HasParent(p.instance)
}

func (p *TPageControl) Hide() {
    defer exceptionProc()
    PageControl_Hide(p.instance)
}

func (p *TPageControl) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return PageControl_Perform(p.instance, Msg , WParam , LParam )
}

func (p *TPageControl) Refresh() {
    defer exceptionProc()
    PageControl_Refresh(p.instance)
}

func (p *TPageControl) SendToBack() {
    defer exceptionProc()
    PageControl_SendToBack(p.instance)
}

func (p *TPageControl) Show() {
    defer exceptionProc()
    PageControl_Show(p.instance)
}

func (p *TPageControl) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return PageControl_GetTextBuf(p.instance, Buffer , BufSize )
}

func (p *TPageControl) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(PageControl_FindComponent(p.instance, AName ))
}

func (p *TPageControl) GetNamePath() string {
    defer exceptionProc()
    return PageControl_GetNamePath(p.instance)
}

func (p *TPageControl) Assign(Source IObject) {
    defer exceptionProc()
    PageControl_Assign(p.instance, CheckPtr(Source))
}

func (p *TPageControl) ClassName() string {
    defer exceptionProc()
    return PageControl_ClassName(p.instance)
}

func (p *TPageControl) Equals(Obj IObject) bool {
    defer exceptionProc()
    return PageControl_Equals(p.instance, CheckPtr(Obj))
}

func (p *TPageControl) GetHashCode() int32 {
    defer exceptionProc()
    return PageControl_GetHashCode(p.instance)
}

func (p *TPageControl) ToString() string {
    defer exceptionProc()
    return PageControl_ToString(p.instance)
}

func (p *TPageControl) ActivePageIndex() int32 {
    defer exceptionProc()
    return PageControl_GetActivePageIndex(p.instance)
}

func (p *TPageControl) SetActivePageIndex(value int32) {
    defer exceptionProc()
    PageControl_SetActivePageIndex(p.instance, value)
}

func (p *TPageControl) PageCount() int32 {
    defer exceptionProc()
    return PageControl_GetPageCount(p.instance)
}

func (p *TPageControl) Align() TAlign {
    defer exceptionProc()
    return PageControl_GetAlign(p.instance)
}

func (p *TPageControl) SetAlign(value TAlign) {
    defer exceptionProc()
    PageControl_SetAlign(p.instance, value)
}

func (p *TPageControl) Anchors() TAnchors {
    defer exceptionProc()
    return PageControl_GetAnchors(p.instance)
}

func (p *TPageControl) SetAnchors(value TAnchors) {
    defer exceptionProc()
    PageControl_SetAnchors(p.instance, value)
}

func (p *TPageControl) DoubleBuffered() bool {
    defer exceptionProc()
    return PageControl_GetDoubleBuffered(p.instance)
}

func (p *TPageControl) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    PageControl_SetDoubleBuffered(p.instance, value)
}

func (p *TPageControl) Enabled() bool {
    defer exceptionProc()
    return PageControl_GetEnabled(p.instance)
}

func (p *TPageControl) SetEnabled(value bool) {
    defer exceptionProc()
    PageControl_SetEnabled(p.instance, value)
}

func (p *TPageControl) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(PageControl_GetFont(p.instance))
}

func (p *TPageControl) SetFont(value *TFont) {
    defer exceptionProc()
    PageControl_SetFont(p.instance, CheckPtr(value))
}

func (p *TPageControl) HotTrack() bool {
    defer exceptionProc()
    return PageControl_GetHotTrack(p.instance)
}

func (p *TPageControl) SetHotTrack(value bool) {
    defer exceptionProc()
    PageControl_SetHotTrack(p.instance, value)
}

func (p *TPageControl) Images() *TImageList {
    defer exceptionProc()
    return ImageListFromInst(PageControl_GetImages(p.instance))
}

func (p *TPageControl) SetImages(value IComponent) {
    defer exceptionProc()
    PageControl_SetImages(p.instance, CheckPtr(value))
}

func (p *TPageControl) MultiLine() bool {
    defer exceptionProc()
    return PageControl_GetMultiLine(p.instance)
}

func (p *TPageControl) SetMultiLine(value bool) {
    defer exceptionProc()
    PageControl_SetMultiLine(p.instance, value)
}

func (p *TPageControl) ParentFont() bool {
    defer exceptionProc()
    return PageControl_GetParentFont(p.instance)
}

func (p *TPageControl) SetParentFont(value bool) {
    defer exceptionProc()
    PageControl_SetParentFont(p.instance, value)
}

func (p *TPageControl) ParentShowHint() bool {
    defer exceptionProc()
    return PageControl_GetParentShowHint(p.instance)
}

func (p *TPageControl) SetParentShowHint(value bool) {
    defer exceptionProc()
    PageControl_SetParentShowHint(p.instance, value)
}

func (p *TPageControl) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(PageControl_GetPopupMenu(p.instance))
}

func (p *TPageControl) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    PageControl_SetPopupMenu(p.instance, CheckPtr(value))
}

func (p *TPageControl) ShowHint() bool {
    defer exceptionProc()
    return PageControl_GetShowHint(p.instance)
}

func (p *TPageControl) SetShowHint(value bool) {
    defer exceptionProc()
    PageControl_SetShowHint(p.instance, value)
}

func (p *TPageControl) Style() TTabStyle {
    defer exceptionProc()
    return PageControl_GetStyle(p.instance)
}

func (p *TPageControl) SetStyle(value TTabStyle) {
    defer exceptionProc()
    PageControl_SetStyle(p.instance, value)
}

func (p *TPageControl) TabHeight() int16 {
    defer exceptionProc()
    return PageControl_GetTabHeight(p.instance)
}

func (p *TPageControl) SetTabHeight(value int16) {
    defer exceptionProc()
    PageControl_SetTabHeight(p.instance, value)
}

func (p *TPageControl) TabIndex() int32 {
    defer exceptionProc()
    return PageControl_GetTabIndex(p.instance)
}

func (p *TPageControl) SetTabIndex(value int32) {
    defer exceptionProc()
    PageControl_SetTabIndex(p.instance, value)
}

func (p *TPageControl) TabOrder() int16 {
    defer exceptionProc()
    return PageControl_GetTabOrder(p.instance)
}

func (p *TPageControl) SetTabOrder(value int16) {
    defer exceptionProc()
    PageControl_SetTabOrder(p.instance, value)
}

func (p *TPageControl) TabPosition() TTabPosition {
    defer exceptionProc()
    return PageControl_GetTabPosition(p.instance)
}

func (p *TPageControl) SetTabPosition(value TTabPosition) {
    defer exceptionProc()
    PageControl_SetTabPosition(p.instance, value)
}

func (p *TPageControl) TabStop() bool {
    defer exceptionProc()
    return PageControl_GetTabStop(p.instance)
}

func (p *TPageControl) SetTabStop(value bool) {
    defer exceptionProc()
    PageControl_SetTabStop(p.instance, value)
}

func (p *TPageControl) TabWidth() int16 {
    defer exceptionProc()
    return PageControl_GetTabWidth(p.instance)
}

func (p *TPageControl) SetTabWidth(value int16) {
    defer exceptionProc()
    PageControl_SetTabWidth(p.instance, value)
}

func (p *TPageControl) Visible() bool {
    defer exceptionProc()
    return PageControl_GetVisible(p.instance)
}

func (p *TPageControl) SetVisible(value bool) {
    defer exceptionProc()
    PageControl_SetVisible(p.instance, value)
}

func (p *TPageControl) SetOnChange(fn TNotifyEvent) {
    PageControl_SetOnChange(p.instance, fn)
}

func (p *TPageControl) SetOnEnter(fn TNotifyEvent) {
    PageControl_SetOnEnter(p.instance, fn)
}

func (p *TPageControl) SetOnExit(fn TNotifyEvent) {
    PageControl_SetOnExit(p.instance, fn)
}

func (p *TPageControl) SetOnMouseDown(fn TMouseEvent) {
    PageControl_SetOnMouseDown(p.instance, fn)
}

func (p *TPageControl) SetOnMouseEnter(fn TNotifyEvent) {
    PageControl_SetOnMouseEnter(p.instance, fn)
}

func (p *TPageControl) SetOnMouseLeave(fn TNotifyEvent) {
    PageControl_SetOnMouseLeave(p.instance, fn)
}

func (p *TPageControl) SetOnMouseMove(fn TMouseMoveEvent) {
    PageControl_SetOnMouseMove(p.instance, fn)
}

func (p *TPageControl) SetOnMouseUp(fn TMouseEvent) {
    PageControl_SetOnMouseUp(p.instance, fn)
}

func (p *TPageControl) SetOnResize(fn TNotifyEvent) {
    PageControl_SetOnResize(p.instance, fn)
}

func (p *TPageControl) Canvas() *TCanvas {
    defer exceptionProc()
    return CanvasFromInst(PageControl_GetCanvas(p.instance))
}

func (p *TPageControl) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(PageControl_GetBrush(p.instance))
}

func (p *TPageControl) ControlCount() int32 {
    defer exceptionProc()
    return PageControl_GetControlCount(p.instance)
}

func (p *TPageControl) Handle() HWND {
    defer exceptionProc()
    return PageControl_GetHandle(p.instance)
}

func (p *TPageControl) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(PageControl_GetAction(p.instance))
}

func (p *TPageControl) SetAction(value IComponent) {
    defer exceptionProc()
    PageControl_SetAction(p.instance, CheckPtr(value))
}

func (p *TPageControl) BoundsRect() TRect {
    defer exceptionProc()
    return PageControl_GetBoundsRect(p.instance)
}

func (p *TPageControl) SetBoundsRect(value TRect) {
    defer exceptionProc()
    PageControl_SetBoundsRect(p.instance, value)
}

func (p *TPageControl) ClientHeight() int32 {
    defer exceptionProc()
    return PageControl_GetClientHeight(p.instance)
}

func (p *TPageControl) SetClientHeight(value int32) {
    defer exceptionProc()
    PageControl_SetClientHeight(p.instance, value)
}

func (p *TPageControl) ClientRect() TRect {
    defer exceptionProc()
    return PageControl_GetClientRect(p.instance)
}

func (p *TPageControl) ClientWidth() int32 {
    defer exceptionProc()
    return PageControl_GetClientWidth(p.instance)
}

func (p *TPageControl) SetClientWidth(value int32) {
    defer exceptionProc()
    PageControl_SetClientWidth(p.instance, value)
}

func (p *TPageControl) ExplicitLeft() int32 {
    defer exceptionProc()
    return PageControl_GetExplicitLeft(p.instance)
}

func (p *TPageControl) ExplicitTop() int32 {
    defer exceptionProc()
    return PageControl_GetExplicitTop(p.instance)
}

func (p *TPageControl) ExplicitWidth() int32 {
    defer exceptionProc()
    return PageControl_GetExplicitWidth(p.instance)
}

func (p *TPageControl) ExplicitHeight() int32 {
    defer exceptionProc()
    return PageControl_GetExplicitHeight(p.instance)
}

func (p *TPageControl) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(PageControl_GetParent(p.instance))
}

func (p *TPageControl) SetParent(value IControl) {
    defer exceptionProc()
    PageControl_SetParent(p.instance, CheckPtr(value))
}

func (p *TPageControl) AlignWithMargins() bool {
    defer exceptionProc()
    return PageControl_GetAlignWithMargins(p.instance)
}

func (p *TPageControl) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    PageControl_SetAlignWithMargins(p.instance, value)
}

func (p *TPageControl) Left() int32 {
    defer exceptionProc()
    return PageControl_GetLeft(p.instance)
}

func (p *TPageControl) SetLeft(value int32) {
    defer exceptionProc()
    PageControl_SetLeft(p.instance, value)
}

func (p *TPageControl) Top() int32 {
    defer exceptionProc()
    return PageControl_GetTop(p.instance)
}

func (p *TPageControl) SetTop(value int32) {
    defer exceptionProc()
    PageControl_SetTop(p.instance, value)
}

func (p *TPageControl) Width() int32 {
    defer exceptionProc()
    return PageControl_GetWidth(p.instance)
}

func (p *TPageControl) SetWidth(value int32) {
    defer exceptionProc()
    PageControl_SetWidth(p.instance, value)
}

func (p *TPageControl) Height() int32 {
    defer exceptionProc()
    return PageControl_GetHeight(p.instance)
}

func (p *TPageControl) SetHeight(value int32) {
    defer exceptionProc()
    PageControl_SetHeight(p.instance, value)
}

func (p *TPageControl) Cursor() TCursor {
    defer exceptionProc()
    return PageControl_GetCursor(p.instance)
}

func (p *TPageControl) SetCursor(value TCursor) {
    defer exceptionProc()
    PageControl_SetCursor(p.instance, value)
}

func (p *TPageControl) Hint() string {
    defer exceptionProc()
    return PageControl_GetHint(p.instance)
}

func (p *TPageControl) SetHint(value string) {
    defer exceptionProc()
    PageControl_SetHint(p.instance, value)
}

func (p *TPageControl) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(PageControl_GetMargins(p.instance))
}

func (p *TPageControl) SetMargins(value *TMargins) {
    defer exceptionProc()
    PageControl_SetMargins(p.instance, CheckPtr(value))
}

func (p *TPageControl) ComponentCount() int32 {
    defer exceptionProc()
    return PageControl_GetComponentCount(p.instance)
}

func (p *TPageControl) ComponentIndex() int32 {
    defer exceptionProc()
    return PageControl_GetComponentIndex(p.instance)
}

func (p *TPageControl) SetComponentIndex(value int32) {
    defer exceptionProc()
    PageControl_SetComponentIndex(p.instance, value)
}

func (p *TPageControl) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(PageControl_GetOwner(p.instance))
}

func (p *TPageControl) Name() string {
    defer exceptionProc()
    return PageControl_GetName(p.instance)
}

func (p *TPageControl) SetName(value string) {
    defer exceptionProc()
    PageControl_SetName(p.instance, value)
}

func (p *TPageControl) Tag() int {
    defer exceptionProc()
    return PageControl_GetTag(p.instance)
}

func (p *TPageControl) SetTag(value int) {
    defer exceptionProc()
    PageControl_SetTag(p.instance, value)
}

func (p *TPageControl) Pages(Index int32) *TTabSheet {
    defer exceptionProc()
    return TabSheetFromInst(PageControl_GetPages(p.instance, Index))
}

func (p *TPageControl) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(PageControl_GetControls(p.instance, Index))
}

func (p *TPageControl) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(PageControl_GetComponents(p.instance, AIndex))
}

