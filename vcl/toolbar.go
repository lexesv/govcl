
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TToolBar struct {
    IControl
    instance uintptr
}

func NewToolBar(owner IComponent) *TToolBar {
    t := new(TToolBar)
    t.instance = ToolBar_Create(owner.Instance())
    return t
}

func ToolBarFromInst(inst uintptr) *TToolBar {
    t := new(TToolBar)
    t.instance = inst
    return t
}

func ToolBarFromObj(obj IObject) *TToolBar {
    t := new(TToolBar)
    t.instance = CheckPtr(obj)
    return t
}

func (t *TToolBar) Free() {
    if t.instance != 0 {
        ToolBar_Free(t.instance)
    }
}

func (t *TToolBar) Instance() uintptr {
    return t.instance
}

func (t *TToolBar) IsValid() bool {
    return t.instance != 0
}

func (t *TToolBar) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    ToolBar_FlipChildren(t.instance, AllLevels )
}

func (t *TToolBar) CanFocus() bool {
    defer exceptionProc()
    return ToolBar_CanFocus(t.instance)
}

func (t *TToolBar) Focused() bool {
    defer exceptionProc()
    return ToolBar_Focused(t.instance)
}

func (t *TToolBar) HandleAllocated() bool {
    defer exceptionProc()
    return ToolBar_HandleAllocated(t.instance)
}

func (t *TToolBar) Invalidate() {
    defer exceptionProc()
    ToolBar_Invalidate(t.instance)
}

func (t *TToolBar) Realign() {
    defer exceptionProc()
    ToolBar_Realign(t.instance)
}

func (t *TToolBar) Repaint() {
    defer exceptionProc()
    ToolBar_Repaint(t.instance)
}

func (t *TToolBar) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    ToolBar_ScaleBy(t.instance, M , D )
}

func (t *TToolBar) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    ToolBar_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight )
}

func (t *TToolBar) SetFocus() {
    defer exceptionProc()
    ToolBar_SetFocus(t.instance)
}

func (t *TToolBar) Update() {
    defer exceptionProc()
    ToolBar_Update(t.instance)
}

func (t *TToolBar) BringToFront() {
    defer exceptionProc()
    ToolBar_BringToFront(t.instance)
}

func (t *TToolBar) HasParent() bool {
    defer exceptionProc()
    return ToolBar_HasParent(t.instance)
}

func (t *TToolBar) Hide() {
    defer exceptionProc()
    ToolBar_Hide(t.instance)
}

func (t *TToolBar) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return ToolBar_Perform(t.instance, Msg , WParam , LParam )
}

func (t *TToolBar) Refresh() {
    defer exceptionProc()
    ToolBar_Refresh(t.instance)
}

func (t *TToolBar) SendToBack() {
    defer exceptionProc()
    ToolBar_SendToBack(t.instance)
}

func (t *TToolBar) Show() {
    defer exceptionProc()
    ToolBar_Show(t.instance)
}

func (t *TToolBar) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return ToolBar_GetTextBuf(t.instance, Buffer , BufSize )
}

func (t *TToolBar) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ToolBar_FindComponent(t.instance, AName ))
}

func (t *TToolBar) GetNamePath() string {
    defer exceptionProc()
    return ToolBar_GetNamePath(t.instance)
}

func (t *TToolBar) Assign(Source IObject) {
    defer exceptionProc()
    ToolBar_Assign(t.instance, CheckPtr(Source))
}

func (t *TToolBar) ClassName() string {
    defer exceptionProc()
    return ToolBar_ClassName(t.instance)
}

func (t *TToolBar) Equals(Obj IObject) bool {
    defer exceptionProc()
    return ToolBar_Equals(t.instance, CheckPtr(Obj))
}

func (t *TToolBar) GetHashCode() int32 {
    defer exceptionProc()
    return ToolBar_GetHashCode(t.instance)
}

func (t *TToolBar) ToString() string {
    defer exceptionProc()
    return ToolBar_ToString(t.instance)
}

func (t *TToolBar) ButtonCount() int32 {
    defer exceptionProc()
    return ToolBar_GetButtonCount(t.instance)
}

func (t *TToolBar) Canvas() *TCanvas {
    defer exceptionProc()
    return CanvasFromInst(ToolBar_GetCanvas(t.instance))
}

func (t *TToolBar) CustomizeKeyName() string {
    defer exceptionProc()
    return ToolBar_GetCustomizeKeyName(t.instance)
}

func (t *TToolBar) SetCustomizeKeyName(value string) {
    defer exceptionProc()
    ToolBar_SetCustomizeKeyName(t.instance, value)
}

func (t *TToolBar) CustomizeValueName() string {
    defer exceptionProc()
    return ToolBar_GetCustomizeValueName(t.instance)
}

func (t *TToolBar) SetCustomizeValueName(value string) {
    defer exceptionProc()
    ToolBar_SetCustomizeValueName(t.instance, value)
}

func (t *TToolBar) RowCount() int32 {
    defer exceptionProc()
    return ToolBar_GetRowCount(t.instance)
}

func (t *TToolBar) Align() TAlign {
    defer exceptionProc()
    return ToolBar_GetAlign(t.instance)
}

func (t *TToolBar) SetAlign(value TAlign) {
    defer exceptionProc()
    ToolBar_SetAlign(t.instance, value)
}

func (t *TToolBar) Anchors() TAnchors {
    defer exceptionProc()
    return ToolBar_GetAnchors(t.instance)
}

func (t *TToolBar) SetAnchors(value TAnchors) {
    defer exceptionProc()
    ToolBar_SetAnchors(t.instance, value)
}

func (t *TToolBar) AutoSize() bool {
    defer exceptionProc()
    return ToolBar_GetAutoSize(t.instance)
}

func (t *TToolBar) SetAutoSize(value bool) {
    defer exceptionProc()
    ToolBar_SetAutoSize(t.instance, value)
}

func (t *TToolBar) BorderWidth() int32 {
    defer exceptionProc()
    return ToolBar_GetBorderWidth(t.instance)
}

func (t *TToolBar) SetBorderWidth(value int32) {
    defer exceptionProc()
    ToolBar_SetBorderWidth(t.instance, value)
}

func (t *TToolBar) ButtonHeight() int32 {
    defer exceptionProc()
    return ToolBar_GetButtonHeight(t.instance)
}

func (t *TToolBar) SetButtonHeight(value int32) {
    defer exceptionProc()
    ToolBar_SetButtonHeight(t.instance, value)
}

func (t *TToolBar) ButtonWidth() int32 {
    defer exceptionProc()
    return ToolBar_GetButtonWidth(t.instance)
}

func (t *TToolBar) SetButtonWidth(value int32) {
    defer exceptionProc()
    ToolBar_SetButtonWidth(t.instance, value)
}

func (t *TToolBar) Caption() string {
    defer exceptionProc()
    return ToolBar_GetCaption(t.instance)
}

func (t *TToolBar) SetCaption(value string) {
    defer exceptionProc()
    ToolBar_SetCaption(t.instance, value)
}

func (t *TToolBar) Color() TColor {
    defer exceptionProc()
    return ToolBar_GetColor(t.instance)
}

func (t *TToolBar) SetColor(value TColor) {
    defer exceptionProc()
    ToolBar_SetColor(t.instance, value)
}

func (t *TToolBar) DoubleBuffered() bool {
    defer exceptionProc()
    return ToolBar_GetDoubleBuffered(t.instance)
}

func (t *TToolBar) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    ToolBar_SetDoubleBuffered(t.instance, value)
}

func (t *TToolBar) DrawingStyle() TTBDrawingStyle {
    defer exceptionProc()
    return ToolBar_GetDrawingStyle(t.instance)
}

func (t *TToolBar) SetDrawingStyle(value TTBDrawingStyle) {
    defer exceptionProc()
    ToolBar_SetDrawingStyle(t.instance, value)
}

func (t *TToolBar) Enabled() bool {
    defer exceptionProc()
    return ToolBar_GetEnabled(t.instance)
}

func (t *TToolBar) SetEnabled(value bool) {
    defer exceptionProc()
    ToolBar_SetEnabled(t.instance, value)
}

func (t *TToolBar) Flat() bool {
    defer exceptionProc()
    return ToolBar_GetFlat(t.instance)
}

func (t *TToolBar) SetFlat(value bool) {
    defer exceptionProc()
    ToolBar_SetFlat(t.instance, value)
}

func (t *TToolBar) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(ToolBar_GetFont(t.instance))
}

func (t *TToolBar) SetFont(value *TFont) {
    defer exceptionProc()
    ToolBar_SetFont(t.instance, CheckPtr(value))
}

func (t *TToolBar) GradientEndColor() TColor {
    defer exceptionProc()
    return ToolBar_GetGradientEndColor(t.instance)
}

func (t *TToolBar) SetGradientEndColor(value TColor) {
    defer exceptionProc()
    ToolBar_SetGradientEndColor(t.instance, value)
}

func (t *TToolBar) GradientStartColor() TColor {
    defer exceptionProc()
    return ToolBar_GetGradientStartColor(t.instance)
}

func (t *TToolBar) SetGradientStartColor(value TColor) {
    defer exceptionProc()
    ToolBar_SetGradientStartColor(t.instance, value)
}

func (t *TToolBar) Height() int32 {
    defer exceptionProc()
    return ToolBar_GetHeight(t.instance)
}

func (t *TToolBar) SetHeight(value int32) {
    defer exceptionProc()
    ToolBar_SetHeight(t.instance, value)
}

func (t *TToolBar) HideClippedButtons() bool {
    defer exceptionProc()
    return ToolBar_GetHideClippedButtons(t.instance)
}

func (t *TToolBar) SetHideClippedButtons(value bool) {
    defer exceptionProc()
    ToolBar_SetHideClippedButtons(t.instance, value)
}

func (t *TToolBar) HotImages() *TImageList {
    defer exceptionProc()
    return ImageListFromInst(ToolBar_GetHotImages(t.instance))
}

func (t *TToolBar) SetHotImages(value IComponent) {
    defer exceptionProc()
    ToolBar_SetHotImages(t.instance, CheckPtr(value))
}

func (t *TToolBar) Images() *TImageList {
    defer exceptionProc()
    return ImageListFromInst(ToolBar_GetImages(t.instance))
}

func (t *TToolBar) SetImages(value IComponent) {
    defer exceptionProc()
    ToolBar_SetImages(t.instance, CheckPtr(value))
}

func (t *TToolBar) Indent() int32 {
    defer exceptionProc()
    return ToolBar_GetIndent(t.instance)
}

func (t *TToolBar) SetIndent(value int32) {
    defer exceptionProc()
    ToolBar_SetIndent(t.instance, value)
}

func (t *TToolBar) List() bool {
    defer exceptionProc()
    return ToolBar_GetList(t.instance)
}

func (t *TToolBar) SetList(value bool) {
    defer exceptionProc()
    ToolBar_SetList(t.instance, value)
}

func (t *TToolBar) Menu() *TMainMenu {
    defer exceptionProc()
    return MainMenuFromInst(ToolBar_GetMenu(t.instance))
}

func (t *TToolBar) SetMenu(value IComponent) {
    defer exceptionProc()
    ToolBar_SetMenu(t.instance, CheckPtr(value))
}

func (t *TToolBar) GradientDirection() TGradientDirection {
    defer exceptionProc()
    return ToolBar_GetGradientDirection(t.instance)
}

func (t *TToolBar) SetGradientDirection(value TGradientDirection) {
    defer exceptionProc()
    ToolBar_SetGradientDirection(t.instance, value)
}

func (t *TToolBar) GradientDrawingOptions() TTBGradientDrawingOptions {
    defer exceptionProc()
    return ToolBar_GetGradientDrawingOptions(t.instance)
}

func (t *TToolBar) SetGradientDrawingOptions(value TTBGradientDrawingOptions) {
    defer exceptionProc()
    ToolBar_SetGradientDrawingOptions(t.instance, value)
}

func (t *TToolBar) ParentColor() bool {
    defer exceptionProc()
    return ToolBar_GetParentColor(t.instance)
}

func (t *TToolBar) SetParentColor(value bool) {
    defer exceptionProc()
    ToolBar_SetParentColor(t.instance, value)
}

func (t *TToolBar) ParentFont() bool {
    defer exceptionProc()
    return ToolBar_GetParentFont(t.instance)
}

func (t *TToolBar) SetParentFont(value bool) {
    defer exceptionProc()
    ToolBar_SetParentFont(t.instance, value)
}

func (t *TToolBar) ParentShowHint() bool {
    defer exceptionProc()
    return ToolBar_GetParentShowHint(t.instance)
}

func (t *TToolBar) SetParentShowHint(value bool) {
    defer exceptionProc()
    ToolBar_SetParentShowHint(t.instance, value)
}

func (t *TToolBar) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(ToolBar_GetPopupMenu(t.instance))
}

func (t *TToolBar) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    ToolBar_SetPopupMenu(t.instance, CheckPtr(value))
}

func (t *TToolBar) ShowCaptions() bool {
    defer exceptionProc()
    return ToolBar_GetShowCaptions(t.instance)
}

func (t *TToolBar) SetShowCaptions(value bool) {
    defer exceptionProc()
    ToolBar_SetShowCaptions(t.instance, value)
}

func (t *TToolBar) ShowHint() bool {
    defer exceptionProc()
    return ToolBar_GetShowHint(t.instance)
}

func (t *TToolBar) SetShowHint(value bool) {
    defer exceptionProc()
    ToolBar_SetShowHint(t.instance, value)
}

func (t *TToolBar) TabOrder() int16 {
    defer exceptionProc()
    return ToolBar_GetTabOrder(t.instance)
}

func (t *TToolBar) SetTabOrder(value int16) {
    defer exceptionProc()
    ToolBar_SetTabOrder(t.instance, value)
}

func (t *TToolBar) TabStop() bool {
    defer exceptionProc()
    return ToolBar_GetTabStop(t.instance)
}

func (t *TToolBar) SetTabStop(value bool) {
    defer exceptionProc()
    ToolBar_SetTabStop(t.instance, value)
}

func (t *TToolBar) Transparent() bool {
    defer exceptionProc()
    return ToolBar_GetTransparent(t.instance)
}

func (t *TToolBar) SetTransparent(value bool) {
    defer exceptionProc()
    ToolBar_SetTransparent(t.instance, value)
}

func (t *TToolBar) Visible() bool {
    defer exceptionProc()
    return ToolBar_GetVisible(t.instance)
}

func (t *TToolBar) SetVisible(value bool) {
    defer exceptionProc()
    ToolBar_SetVisible(t.instance, value)
}

func (t *TToolBar) Wrapable() bool {
    defer exceptionProc()
    return ToolBar_GetWrapable(t.instance)
}

func (t *TToolBar) SetWrapable(value bool) {
    defer exceptionProc()
    ToolBar_SetWrapable(t.instance, value)
}

func (t *TToolBar) SetOnClick(fn TNotifyEvent) {
    ToolBar_SetOnClick(t.instance, fn)
}

func (t *TToolBar) SetOnDblClick(fn TNotifyEvent) {
    ToolBar_SetOnDblClick(t.instance, fn)
}

func (t *TToolBar) SetOnEnter(fn TNotifyEvent) {
    ToolBar_SetOnEnter(t.instance, fn)
}

func (t *TToolBar) SetOnExit(fn TNotifyEvent) {
    ToolBar_SetOnExit(t.instance, fn)
}

func (t *TToolBar) SetOnMouseDown(fn TMouseEvent) {
    ToolBar_SetOnMouseDown(t.instance, fn)
}

func (t *TToolBar) SetOnMouseEnter(fn TNotifyEvent) {
    ToolBar_SetOnMouseEnter(t.instance, fn)
}

func (t *TToolBar) SetOnMouseLeave(fn TNotifyEvent) {
    ToolBar_SetOnMouseLeave(t.instance, fn)
}

func (t *TToolBar) SetOnMouseMove(fn TMouseMoveEvent) {
    ToolBar_SetOnMouseMove(t.instance, fn)
}

func (t *TToolBar) SetOnMouseUp(fn TMouseEvent) {
    ToolBar_SetOnMouseUp(t.instance, fn)
}

func (t *TToolBar) SetOnResize(fn TNotifyEvent) {
    ToolBar_SetOnResize(t.instance, fn)
}

func (t *TToolBar) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(ToolBar_GetBrush(t.instance))
}

func (t *TToolBar) ControlCount() int32 {
    defer exceptionProc()
    return ToolBar_GetControlCount(t.instance)
}

func (t *TToolBar) Handle() HWND {
    defer exceptionProc()
    return ToolBar_GetHandle(t.instance)
}

func (t *TToolBar) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(ToolBar_GetAction(t.instance))
}

func (t *TToolBar) SetAction(value IComponent) {
    defer exceptionProc()
    ToolBar_SetAction(t.instance, CheckPtr(value))
}

func (t *TToolBar) BoundsRect() TRect {
    defer exceptionProc()
    return ToolBar_GetBoundsRect(t.instance)
}

func (t *TToolBar) SetBoundsRect(value TRect) {
    defer exceptionProc()
    ToolBar_SetBoundsRect(t.instance, value)
}

func (t *TToolBar) ClientHeight() int32 {
    defer exceptionProc()
    return ToolBar_GetClientHeight(t.instance)
}

func (t *TToolBar) SetClientHeight(value int32) {
    defer exceptionProc()
    ToolBar_SetClientHeight(t.instance, value)
}

func (t *TToolBar) ClientRect() TRect {
    defer exceptionProc()
    return ToolBar_GetClientRect(t.instance)
}

func (t *TToolBar) ClientWidth() int32 {
    defer exceptionProc()
    return ToolBar_GetClientWidth(t.instance)
}

func (t *TToolBar) SetClientWidth(value int32) {
    defer exceptionProc()
    ToolBar_SetClientWidth(t.instance, value)
}

func (t *TToolBar) ExplicitLeft() int32 {
    defer exceptionProc()
    return ToolBar_GetExplicitLeft(t.instance)
}

func (t *TToolBar) ExplicitTop() int32 {
    defer exceptionProc()
    return ToolBar_GetExplicitTop(t.instance)
}

func (t *TToolBar) ExplicitWidth() int32 {
    defer exceptionProc()
    return ToolBar_GetExplicitWidth(t.instance)
}

func (t *TToolBar) ExplicitHeight() int32 {
    defer exceptionProc()
    return ToolBar_GetExplicitHeight(t.instance)
}

func (t *TToolBar) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(ToolBar_GetParent(t.instance))
}

func (t *TToolBar) SetParent(value IControl) {
    defer exceptionProc()
    ToolBar_SetParent(t.instance, CheckPtr(value))
}

func (t *TToolBar) AlignWithMargins() bool {
    defer exceptionProc()
    return ToolBar_GetAlignWithMargins(t.instance)
}

func (t *TToolBar) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    ToolBar_SetAlignWithMargins(t.instance, value)
}

func (t *TToolBar) Left() int32 {
    defer exceptionProc()
    return ToolBar_GetLeft(t.instance)
}

func (t *TToolBar) SetLeft(value int32) {
    defer exceptionProc()
    ToolBar_SetLeft(t.instance, value)
}

func (t *TToolBar) Top() int32 {
    defer exceptionProc()
    return ToolBar_GetTop(t.instance)
}

func (t *TToolBar) SetTop(value int32) {
    defer exceptionProc()
    ToolBar_SetTop(t.instance, value)
}

func (t *TToolBar) Width() int32 {
    defer exceptionProc()
    return ToolBar_GetWidth(t.instance)
}

func (t *TToolBar) SetWidth(value int32) {
    defer exceptionProc()
    ToolBar_SetWidth(t.instance, value)
}

func (t *TToolBar) Cursor() TCursor {
    defer exceptionProc()
    return ToolBar_GetCursor(t.instance)
}

func (t *TToolBar) SetCursor(value TCursor) {
    defer exceptionProc()
    ToolBar_SetCursor(t.instance, value)
}

func (t *TToolBar) Hint() string {
    defer exceptionProc()
    return ToolBar_GetHint(t.instance)
}

func (t *TToolBar) SetHint(value string) {
    defer exceptionProc()
    ToolBar_SetHint(t.instance, value)
}

func (t *TToolBar) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(ToolBar_GetMargins(t.instance))
}

func (t *TToolBar) SetMargins(value *TMargins) {
    defer exceptionProc()
    ToolBar_SetMargins(t.instance, CheckPtr(value))
}

func (t *TToolBar) ComponentCount() int32 {
    defer exceptionProc()
    return ToolBar_GetComponentCount(t.instance)
}

func (t *TToolBar) ComponentIndex() int32 {
    defer exceptionProc()
    return ToolBar_GetComponentIndex(t.instance)
}

func (t *TToolBar) SetComponentIndex(value int32) {
    defer exceptionProc()
    ToolBar_SetComponentIndex(t.instance, value)
}

func (t *TToolBar) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ToolBar_GetOwner(t.instance))
}

func (t *TToolBar) Name() string {
    defer exceptionProc()
    return ToolBar_GetName(t.instance)
}

func (t *TToolBar) SetName(value string) {
    defer exceptionProc()
    ToolBar_SetName(t.instance, value)
}

func (t *TToolBar) Tag() int {
    defer exceptionProc()
    return ToolBar_GetTag(t.instance)
}

func (t *TToolBar) SetTag(value int) {
    defer exceptionProc()
    ToolBar_SetTag(t.instance, value)
}

func (t *TToolBar) Buttons(Index int32) *TToolButton {
    defer exceptionProc()
    return ToolButtonFromInst(ToolBar_GetButtons(t.instance, Index))
}

func (t *TToolBar) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(ToolBar_GetControls(t.instance, Index))
}

func (t *TToolBar) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ToolBar_GetComponents(t.instance, AIndex))
}

