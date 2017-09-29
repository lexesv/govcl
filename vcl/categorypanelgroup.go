
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

type TCategoryPanelGroup struct {
    IControl
    instance uintptr
}

func NewCategoryPanelGroup(owner IComponent) *TCategoryPanelGroup {
    c := new(TCategoryPanelGroup)
    c.instance = CategoryPanelGroup_Create(owner.Instance())
    return c
}

func CategoryPanelGroupFromInst(inst uintptr) *TCategoryPanelGroup {
    c := new(TCategoryPanelGroup)
    c.instance = inst
    return c
}

func CategoryPanelGroupFromObj(obj IObject) *TCategoryPanelGroup {
    c := new(TCategoryPanelGroup)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TCategoryPanelGroup) Free() {
    if c.instance != 0 {
        CategoryPanelGroup_Free(c.instance)
    }
}

func (c *TCategoryPanelGroup) Instance() uintptr {
    return c.instance
}

func (c *TCategoryPanelGroup) IsValid() bool {
    return c.instance != 0
}

func (c *TCategoryPanelGroup) CanFocus() bool {
    defer exceptionProc()
    return CategoryPanelGroup_CanFocus(c.instance)
}

func (c *TCategoryPanelGroup) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    CategoryPanelGroup_FlipChildren(c.instance, AllLevels )
}

func (c *TCategoryPanelGroup) Focused() bool {
    defer exceptionProc()
    return CategoryPanelGroup_Focused(c.instance)
}

func (c *TCategoryPanelGroup) HandleAllocated() bool {
    defer exceptionProc()
    return CategoryPanelGroup_HandleAllocated(c.instance)
}

func (c *TCategoryPanelGroup) Invalidate() {
    defer exceptionProc()
    CategoryPanelGroup_Invalidate(c.instance)
}

func (c *TCategoryPanelGroup) Realign() {
    defer exceptionProc()
    CategoryPanelGroup_Realign(c.instance)
}

func (c *TCategoryPanelGroup) Repaint() {
    defer exceptionProc()
    CategoryPanelGroup_Repaint(c.instance)
}

func (c *TCategoryPanelGroup) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    CategoryPanelGroup_ScaleBy(c.instance, M , D )
}

func (c *TCategoryPanelGroup) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    CategoryPanelGroup_SetBounds(c.instance, ALeft , ATop , AWidth , AHeight )
}

func (c *TCategoryPanelGroup) SetFocus() {
    defer exceptionProc()
    CategoryPanelGroup_SetFocus(c.instance)
}

func (c *TCategoryPanelGroup) Update() {
    defer exceptionProc()
    CategoryPanelGroup_Update(c.instance)
}

func (c *TCategoryPanelGroup) BringToFront() {
    defer exceptionProc()
    CategoryPanelGroup_BringToFront(c.instance)
}

func (c *TCategoryPanelGroup) HasParent() bool {
    defer exceptionProc()
    return CategoryPanelGroup_HasParent(c.instance)
}

func (c *TCategoryPanelGroup) Hide() {
    defer exceptionProc()
    CategoryPanelGroup_Hide(c.instance)
}

func (c *TCategoryPanelGroup) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return CategoryPanelGroup_Perform(c.instance, Msg , WParam , LParam )
}

func (c *TCategoryPanelGroup) Refresh() {
    defer exceptionProc()
    CategoryPanelGroup_Refresh(c.instance)
}

func (c *TCategoryPanelGroup) SendToBack() {
    defer exceptionProc()
    CategoryPanelGroup_SendToBack(c.instance)
}

func (c *TCategoryPanelGroup) Show() {
    defer exceptionProc()
    CategoryPanelGroup_Show(c.instance)
}

func (c *TCategoryPanelGroup) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return CategoryPanelGroup_GetTextBuf(c.instance, Buffer , BufSize )
}

func (c *TCategoryPanelGroup) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(CategoryPanelGroup_FindComponent(c.instance, AName ))
}

func (c *TCategoryPanelGroup) GetNamePath() string {
    defer exceptionProc()
    return CategoryPanelGroup_GetNamePath(c.instance)
}

func (c *TCategoryPanelGroup) Assign(Source IObject) {
    defer exceptionProc()
    CategoryPanelGroup_Assign(c.instance, CheckPtr(Source))
}

func (c *TCategoryPanelGroup) ClassName() string {
    defer exceptionProc()
    return CategoryPanelGroup_ClassName(c.instance)
}

func (c *TCategoryPanelGroup) Equals(Obj IObject) bool {
    defer exceptionProc()
    return CategoryPanelGroup_Equals(c.instance, CheckPtr(Obj))
}

func (c *TCategoryPanelGroup) GetHashCode() int32 {
    defer exceptionProc()
    return CategoryPanelGroup_GetHashCode(c.instance)
}

func (c *TCategoryPanelGroup) ToString() string {
    defer exceptionProc()
    return CategoryPanelGroup_ToString(c.instance)
}

func (c *TCategoryPanelGroup) Align() TAlign {
    defer exceptionProc()
    return CategoryPanelGroup_GetAlign(c.instance)
}

func (c *TCategoryPanelGroup) SetAlign(value TAlign) {
    defer exceptionProc()
    CategoryPanelGroup_SetAlign(c.instance, value)
}

func (c *TCategoryPanelGroup) Anchors() TAnchors {
    defer exceptionProc()
    return CategoryPanelGroup_GetAnchors(c.instance)
}

func (c *TCategoryPanelGroup) SetAnchors(value TAnchors) {
    defer exceptionProc()
    CategoryPanelGroup_SetAnchors(c.instance, value)
}

func (c *TCategoryPanelGroup) DoubleBuffered() bool {
    defer exceptionProc()
    return CategoryPanelGroup_GetDoubleBuffered(c.instance)
}

func (c *TCategoryPanelGroup) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    CategoryPanelGroup_SetDoubleBuffered(c.instance, value)
}

func (c *TCategoryPanelGroup) Enabled() bool {
    defer exceptionProc()
    return CategoryPanelGroup_GetEnabled(c.instance)
}

func (c *TCategoryPanelGroup) SetEnabled(value bool) {
    defer exceptionProc()
    CategoryPanelGroup_SetEnabled(c.instance, value)
}

func (c *TCategoryPanelGroup) Color() TColor {
    defer exceptionProc()
    return CategoryPanelGroup_GetColor(c.instance)
}

func (c *TCategoryPanelGroup) SetColor(value TColor) {
    defer exceptionProc()
    CategoryPanelGroup_SetColor(c.instance, value)
}

func (c *TCategoryPanelGroup) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(CategoryPanelGroup_GetFont(c.instance))
}

func (c *TCategoryPanelGroup) SetFont(value *TFont) {
    defer exceptionProc()
    CategoryPanelGroup_SetFont(c.instance, CheckPtr(value))
}

func (c *TCategoryPanelGroup) GradientDirection() TGradientDirection {
    defer exceptionProc()
    return CategoryPanelGroup_GetGradientDirection(c.instance)
}

func (c *TCategoryPanelGroup) SetGradientDirection(value TGradientDirection) {
    defer exceptionProc()
    CategoryPanelGroup_SetGradientDirection(c.instance, value)
}

func (c *TCategoryPanelGroup) Height() int32 {
    defer exceptionProc()
    return CategoryPanelGroup_GetHeight(c.instance)
}

func (c *TCategoryPanelGroup) SetHeight(value int32) {
    defer exceptionProc()
    CategoryPanelGroup_SetHeight(c.instance, value)
}

func (c *TCategoryPanelGroup) Images() *TImageList {
    defer exceptionProc()
    return ImageListFromInst(CategoryPanelGroup_GetImages(c.instance))
}

func (c *TCategoryPanelGroup) SetImages(value IComponent) {
    defer exceptionProc()
    CategoryPanelGroup_SetImages(c.instance, CheckPtr(value))
}

func (c *TCategoryPanelGroup) ParentColor() bool {
    defer exceptionProc()
    return CategoryPanelGroup_GetParentColor(c.instance)
}

func (c *TCategoryPanelGroup) SetParentColor(value bool) {
    defer exceptionProc()
    CategoryPanelGroup_SetParentColor(c.instance, value)
}

func (c *TCategoryPanelGroup) ParentCtl3D() bool {
    defer exceptionProc()
    return CategoryPanelGroup_GetParentCtl3D(c.instance)
}

func (c *TCategoryPanelGroup) SetParentCtl3D(value bool) {
    defer exceptionProc()
    CategoryPanelGroup_SetParentCtl3D(c.instance, value)
}

func (c *TCategoryPanelGroup) ParentFont() bool {
    defer exceptionProc()
    return CategoryPanelGroup_GetParentFont(c.instance)
}

func (c *TCategoryPanelGroup) SetParentFont(value bool) {
    defer exceptionProc()
    CategoryPanelGroup_SetParentFont(c.instance, value)
}

func (c *TCategoryPanelGroup) ParentShowHint() bool {
    defer exceptionProc()
    return CategoryPanelGroup_GetParentShowHint(c.instance)
}

func (c *TCategoryPanelGroup) SetParentShowHint(value bool) {
    defer exceptionProc()
    CategoryPanelGroup_SetParentShowHint(c.instance, value)
}

func (c *TCategoryPanelGroup) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(CategoryPanelGroup_GetPopupMenu(c.instance))
}

func (c *TCategoryPanelGroup) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    CategoryPanelGroup_SetPopupMenu(c.instance, CheckPtr(value))
}

func (c *TCategoryPanelGroup) ShowHint() bool {
    defer exceptionProc()
    return CategoryPanelGroup_GetShowHint(c.instance)
}

func (c *TCategoryPanelGroup) SetShowHint(value bool) {
    defer exceptionProc()
    CategoryPanelGroup_SetShowHint(c.instance, value)
}

func (c *TCategoryPanelGroup) TabOrder() int16 {
    defer exceptionProc()
    return CategoryPanelGroup_GetTabOrder(c.instance)
}

func (c *TCategoryPanelGroup) SetTabOrder(value int16) {
    defer exceptionProc()
    CategoryPanelGroup_SetTabOrder(c.instance, value)
}

func (c *TCategoryPanelGroup) TabStop() bool {
    defer exceptionProc()
    return CategoryPanelGroup_GetTabStop(c.instance)
}

func (c *TCategoryPanelGroup) SetTabStop(value bool) {
    defer exceptionProc()
    CategoryPanelGroup_SetTabStop(c.instance, value)
}

func (c *TCategoryPanelGroup) Visible() bool {
    defer exceptionProc()
    return CategoryPanelGroup_GetVisible(c.instance)
}

func (c *TCategoryPanelGroup) SetVisible(value bool) {
    defer exceptionProc()
    CategoryPanelGroup_SetVisible(c.instance, value)
}

func (c *TCategoryPanelGroup) Width() int32 {
    defer exceptionProc()
    return CategoryPanelGroup_GetWidth(c.instance)
}

func (c *TCategoryPanelGroup) SetWidth(value int32) {
    defer exceptionProc()
    CategoryPanelGroup_SetWidth(c.instance, value)
}

func (c *TCategoryPanelGroup) SetOnClick(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnClick(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnDblClick(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnDblClick(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnEnter(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnEnter(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnExit(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnExit(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnMouseDown(fn TMouseEvent) {
    CategoryPanelGroup_SetOnMouseDown(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnMouseEnter(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnMouseEnter(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnMouseLeave(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnMouseLeave(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnMouseMove(fn TMouseMoveEvent) {
    CategoryPanelGroup_SetOnMouseMove(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnMouseUp(fn TMouseEvent) {
    CategoryPanelGroup_SetOnMouseUp(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnMouseWheel(fn TMouseWheelEvent) {
    CategoryPanelGroup_SetOnMouseWheel(c.instance, fn)
}

func (c *TCategoryPanelGroup) SetOnResize(fn TNotifyEvent) {
    CategoryPanelGroup_SetOnResize(c.instance, fn)
}

func (c *TCategoryPanelGroup) Panels() *TList {
    defer exceptionProc()
    return ListFromInst(CategoryPanelGroup_GetPanels(c.instance))
}

func (c *TCategoryPanelGroup) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(CategoryPanelGroup_GetBrush(c.instance))
}

func (c *TCategoryPanelGroup) ControlCount() int32 {
    defer exceptionProc()
    return CategoryPanelGroup_GetControlCount(c.instance)
}

func (c *TCategoryPanelGroup) Handle() HWND {
    defer exceptionProc()
    return CategoryPanelGroup_GetHandle(c.instance)
}

func (c *TCategoryPanelGroup) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(CategoryPanelGroup_GetAction(c.instance))
}

func (c *TCategoryPanelGroup) SetAction(value IComponent) {
    defer exceptionProc()
    CategoryPanelGroup_SetAction(c.instance, CheckPtr(value))
}

func (c *TCategoryPanelGroup) BoundsRect() TRect {
    defer exceptionProc()
    return CategoryPanelGroup_GetBoundsRect(c.instance)
}

func (c *TCategoryPanelGroup) SetBoundsRect(value TRect) {
    defer exceptionProc()
    CategoryPanelGroup_SetBoundsRect(c.instance, value)
}

func (c *TCategoryPanelGroup) ClientHeight() int32 {
    defer exceptionProc()
    return CategoryPanelGroup_GetClientHeight(c.instance)
}

func (c *TCategoryPanelGroup) SetClientHeight(value int32) {
    defer exceptionProc()
    CategoryPanelGroup_SetClientHeight(c.instance, value)
}

func (c *TCategoryPanelGroup) ClientRect() TRect {
    defer exceptionProc()
    return CategoryPanelGroup_GetClientRect(c.instance)
}

func (c *TCategoryPanelGroup) ClientWidth() int32 {
    defer exceptionProc()
    return CategoryPanelGroup_GetClientWidth(c.instance)
}

func (c *TCategoryPanelGroup) SetClientWidth(value int32) {
    defer exceptionProc()
    CategoryPanelGroup_SetClientWidth(c.instance, value)
}

func (c *TCategoryPanelGroup) ExplicitLeft() int32 {
    defer exceptionProc()
    return CategoryPanelGroup_GetExplicitLeft(c.instance)
}

func (c *TCategoryPanelGroup) ExplicitTop() int32 {
    defer exceptionProc()
    return CategoryPanelGroup_GetExplicitTop(c.instance)
}

func (c *TCategoryPanelGroup) ExplicitWidth() int32 {
    defer exceptionProc()
    return CategoryPanelGroup_GetExplicitWidth(c.instance)
}

func (c *TCategoryPanelGroup) ExplicitHeight() int32 {
    defer exceptionProc()
    return CategoryPanelGroup_GetExplicitHeight(c.instance)
}

func (c *TCategoryPanelGroup) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(CategoryPanelGroup_GetParent(c.instance))
}

func (c *TCategoryPanelGroup) SetParent(value IControl) {
    defer exceptionProc()
    CategoryPanelGroup_SetParent(c.instance, CheckPtr(value))
}

func (c *TCategoryPanelGroup) AlignWithMargins() bool {
    defer exceptionProc()
    return CategoryPanelGroup_GetAlignWithMargins(c.instance)
}

func (c *TCategoryPanelGroup) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    CategoryPanelGroup_SetAlignWithMargins(c.instance, value)
}

func (c *TCategoryPanelGroup) Left() int32 {
    defer exceptionProc()
    return CategoryPanelGroup_GetLeft(c.instance)
}

func (c *TCategoryPanelGroup) SetLeft(value int32) {
    defer exceptionProc()
    CategoryPanelGroup_SetLeft(c.instance, value)
}

func (c *TCategoryPanelGroup) Top() int32 {
    defer exceptionProc()
    return CategoryPanelGroup_GetTop(c.instance)
}

func (c *TCategoryPanelGroup) SetTop(value int32) {
    defer exceptionProc()
    CategoryPanelGroup_SetTop(c.instance, value)
}

func (c *TCategoryPanelGroup) Cursor() TCursor {
    defer exceptionProc()
    return CategoryPanelGroup_GetCursor(c.instance)
}

func (c *TCategoryPanelGroup) SetCursor(value TCursor) {
    defer exceptionProc()
    CategoryPanelGroup_SetCursor(c.instance, value)
}

func (c *TCategoryPanelGroup) Hint() string {
    defer exceptionProc()
    return CategoryPanelGroup_GetHint(c.instance)
}

func (c *TCategoryPanelGroup) SetHint(value string) {
    defer exceptionProc()
    CategoryPanelGroup_SetHint(c.instance, value)
}

func (c *TCategoryPanelGroup) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(CategoryPanelGroup_GetMargins(c.instance))
}

func (c *TCategoryPanelGroup) SetMargins(value *TMargins) {
    defer exceptionProc()
    CategoryPanelGroup_SetMargins(c.instance, CheckPtr(value))
}

func (c *TCategoryPanelGroup) ComponentCount() int32 {
    defer exceptionProc()
    return CategoryPanelGroup_GetComponentCount(c.instance)
}

func (c *TCategoryPanelGroup) ComponentIndex() int32 {
    defer exceptionProc()
    return CategoryPanelGroup_GetComponentIndex(c.instance)
}

func (c *TCategoryPanelGroup) SetComponentIndex(value int32) {
    defer exceptionProc()
    CategoryPanelGroup_SetComponentIndex(c.instance, value)
}

func (c *TCategoryPanelGroup) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(CategoryPanelGroup_GetOwner(c.instance))
}

func (c *TCategoryPanelGroup) Name() string {
    defer exceptionProc()
    return CategoryPanelGroup_GetName(c.instance)
}

func (c *TCategoryPanelGroup) SetName(value string) {
    defer exceptionProc()
    CategoryPanelGroup_SetName(c.instance, value)
}

func (c *TCategoryPanelGroup) Tag() int {
    defer exceptionProc()
    return CategoryPanelGroup_GetTag(c.instance)
}

func (c *TCategoryPanelGroup) SetTag(value int) {
    defer exceptionProc()
    CategoryPanelGroup_SetTag(c.instance, value)
}

func (c *TCategoryPanelGroup) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(CategoryPanelGroup_GetControls(c.instance, Index))
}

func (c *TCategoryPanelGroup) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(CategoryPanelGroup_GetComponents(c.instance, AIndex))
}

