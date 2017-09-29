
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TForm struct {
    IControl
    instance uintptr
}

func NewForm(owner IComponent) *TForm {
    f := new(TForm)
    f.instance = Form_Create(owner.Instance())
    return f
}

func FormFromInst(inst uintptr) *TForm {
    f := new(TForm)
    f.instance = inst
    return f
}

func FormFromObj(obj IObject) *TForm {
    f := new(TForm)
    f.instance = CheckPtr(obj)
    return f
}

func (f *TForm) Free() {
    if f.instance != 0 {
        Form_Free(f.instance)
    }
}

func (f *TForm) Instance() uintptr {
    return f.instance
}

func (f *TForm) IsValid() bool {
    return f.instance != 0
}

func (f *TForm) Close() {
    defer exceptionProc()
    Form_Close(f.instance)
}

func (f *TForm) Hide() {
    defer exceptionProc()
    Form_Hide(f.instance)
}

func (f *TForm) SetFocus() {
    defer exceptionProc()
    Form_SetFocus(f.instance)
}

func (f *TForm) Show() {
    defer exceptionProc()
    Form_Show(f.instance)
}

func (f *TForm) ShowModal() int32 {
    defer exceptionProc()
    return Form_ShowModal(f.instance)
}

func (f *TForm) CanFocus() bool {
    defer exceptionProc()
    return Form_CanFocus(f.instance)
}

func (f *TForm) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    Form_FlipChildren(f.instance, AllLevels )
}

func (f *TForm) Focused() bool {
    defer exceptionProc()
    return Form_Focused(f.instance)
}

func (f *TForm) HandleAllocated() bool {
    defer exceptionProc()
    return Form_HandleAllocated(f.instance)
}

func (f *TForm) Invalidate() {
    defer exceptionProc()
    Form_Invalidate(f.instance)
}

func (f *TForm) Realign() {
    defer exceptionProc()
    Form_Realign(f.instance)
}

func (f *TForm) Repaint() {
    defer exceptionProc()
    Form_Repaint(f.instance)
}

func (f *TForm) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    Form_ScaleBy(f.instance, M , D )
}

func (f *TForm) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    Form_SetBounds(f.instance, ALeft , ATop , AWidth , AHeight )
}

func (f *TForm) Update() {
    defer exceptionProc()
    Form_Update(f.instance)
}

func (f *TForm) BringToFront() {
    defer exceptionProc()
    Form_BringToFront(f.instance)
}

func (f *TForm) HasParent() bool {
    defer exceptionProc()
    return Form_HasParent(f.instance)
}

func (f *TForm) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return Form_Perform(f.instance, Msg , WParam , LParam )
}

func (f *TForm) Refresh() {
    defer exceptionProc()
    Form_Refresh(f.instance)
}

func (f *TForm) SendToBack() {
    defer exceptionProc()
    Form_SendToBack(f.instance)
}

func (f *TForm) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return Form_GetTextBuf(f.instance, Buffer , BufSize )
}

func (f *TForm) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Form_FindComponent(f.instance, AName ))
}

func (f *TForm) GetNamePath() string {
    defer exceptionProc()
    return Form_GetNamePath(f.instance)
}

func (f *TForm) Assign(Source IObject) {
    defer exceptionProc()
    Form_Assign(f.instance, CheckPtr(Source))
}

func (f *TForm) ClassName() string {
    defer exceptionProc()
    return Form_ClassName(f.instance)
}

func (f *TForm) Equals(Obj IObject) bool {
    defer exceptionProc()
    return Form_Equals(f.instance, CheckPtr(Obj))
}

func (f *TForm) GetHashCode() int32 {
    defer exceptionProc()
    return Form_GetHashCode(f.instance)
}

func (f *TForm) ToString() string {
    defer exceptionProc()
    return Form_ToString(f.instance)
}

func (f *TForm) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(Form_GetAction(f.instance))
}

func (f *TForm) SetAction(value IComponent) {
    defer exceptionProc()
    Form_SetAction(f.instance, CheckPtr(value))
}

func (f *TForm) Align() TAlign {
    defer exceptionProc()
    return Form_GetAlign(f.instance)
}

func (f *TForm) SetAlign(value TAlign) {
    defer exceptionProc()
    Form_SetAlign(f.instance, value)
}

func (f *TForm) AlphaBlend() bool {
    defer exceptionProc()
    return Form_GetAlphaBlend(f.instance)
}

func (f *TForm) SetAlphaBlend(value bool) {
    defer exceptionProc()
    Form_SetAlphaBlend(f.instance, value)
}

func (f *TForm) AlphaBlendValue() uint8 {
    defer exceptionProc()
    return Form_GetAlphaBlendValue(f.instance)
}

func (f *TForm) SetAlphaBlendValue(value uint8) {
    defer exceptionProc()
    Form_SetAlphaBlendValue(f.instance, value)
}

func (f *TForm) Anchors() TAnchors {
    defer exceptionProc()
    return Form_GetAnchors(f.instance)
}

func (f *TForm) SetAnchors(value TAnchors) {
    defer exceptionProc()
    Form_SetAnchors(f.instance, value)
}

func (f *TForm) AutoSize() bool {
    defer exceptionProc()
    return Form_GetAutoSize(f.instance)
}

func (f *TForm) SetAutoSize(value bool) {
    defer exceptionProc()
    Form_SetAutoSize(f.instance, value)
}

func (f *TForm) BorderIcons() TBorderIcons {
    defer exceptionProc()
    return Form_GetBorderIcons(f.instance)
}

func (f *TForm) SetBorderIcons(value TBorderIcons) {
    defer exceptionProc()
    Form_SetBorderIcons(f.instance, value)
}

func (f *TForm) BorderStyle() TFormBorderStyle {
    defer exceptionProc()
    return Form_GetBorderStyle(f.instance)
}

func (f *TForm) SetBorderStyle(value TFormBorderStyle) {
    defer exceptionProc()
    Form_SetBorderStyle(f.instance, value)
}

func (f *TForm) BorderWidth() int32 {
    defer exceptionProc()
    return Form_GetBorderWidth(f.instance)
}

func (f *TForm) SetBorderWidth(value int32) {
    defer exceptionProc()
    Form_SetBorderWidth(f.instance, value)
}

func (f *TForm) Caption() string {
    defer exceptionProc()
    return Form_GetCaption(f.instance)
}

func (f *TForm) SetCaption(value string) {
    defer exceptionProc()
    Form_SetCaption(f.instance, value)
}

func (f *TForm) ClientHeight() int32 {
    defer exceptionProc()
    return Form_GetClientHeight(f.instance)
}

func (f *TForm) SetClientHeight(value int32) {
    defer exceptionProc()
    Form_SetClientHeight(f.instance, value)
}

func (f *TForm) ClientWidth() int32 {
    defer exceptionProc()
    return Form_GetClientWidth(f.instance)
}

func (f *TForm) SetClientWidth(value int32) {
    defer exceptionProc()
    Form_SetClientWidth(f.instance, value)
}

func (f *TForm) Color() TColor {
    defer exceptionProc()
    return Form_GetColor(f.instance)
}

func (f *TForm) SetColor(value TColor) {
    defer exceptionProc()
    Form_SetColor(f.instance, value)
}

func (f *TForm) TransparentColor() bool {
    defer exceptionProc()
    return Form_GetTransparentColor(f.instance)
}

func (f *TForm) SetTransparentColor(value bool) {
    defer exceptionProc()
    Form_SetTransparentColor(f.instance, value)
}

func (f *TForm) TransparentColorValue() TColor {
    defer exceptionProc()
    return Form_GetTransparentColorValue(f.instance)
}

func (f *TForm) SetTransparentColorValue(value TColor) {
    defer exceptionProc()
    Form_SetTransparentColorValue(f.instance, value)
}

func (f *TForm) DoubleBuffered() bool {
    defer exceptionProc()
    return Form_GetDoubleBuffered(f.instance)
}

func (f *TForm) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    Form_SetDoubleBuffered(f.instance, value)
}

func (f *TForm) Enabled() bool {
    defer exceptionProc()
    return Form_GetEnabled(f.instance)
}

func (f *TForm) SetEnabled(value bool) {
    defer exceptionProc()
    Form_SetEnabled(f.instance, value)
}

func (f *TForm) ParentFont() bool {
    defer exceptionProc()
    return Form_GetParentFont(f.instance)
}

func (f *TForm) SetParentFont(value bool) {
    defer exceptionProc()
    Form_SetParentFont(f.instance, value)
}

func (f *TForm) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(Form_GetFont(f.instance))
}

func (f *TForm) SetFont(value *TFont) {
    defer exceptionProc()
    Form_SetFont(f.instance, CheckPtr(value))
}

func (f *TForm) FormStyle() TFormStyle {
    defer exceptionProc()
    return Form_GetFormStyle(f.instance)
}

func (f *TForm) SetFormStyle(value TFormStyle) {
    defer exceptionProc()
    Form_SetFormStyle(f.instance, value)
}

func (f *TForm) Height() int32 {
    defer exceptionProc()
    return Form_GetHeight(f.instance)
}

func (f *TForm) SetHeight(value int32) {
    defer exceptionProc()
    Form_SetHeight(f.instance, value)
}

func (f *TForm) Icon() *TIcon {
    defer exceptionProc()
    return IconFromInst(Form_GetIcon(f.instance))
}

func (f *TForm) SetIcon(value *TIcon) {
    defer exceptionProc()
    Form_SetIcon(f.instance, CheckPtr(value))
}

func (f *TForm) KeyPreview() bool {
    defer exceptionProc()
    return Form_GetKeyPreview(f.instance)
}

func (f *TForm) SetKeyPreview(value bool) {
    defer exceptionProc()
    Form_SetKeyPreview(f.instance, value)
}

func (f *TForm) Menu() *TMainMenu {
    defer exceptionProc()
    return MainMenuFromInst(Form_GetMenu(f.instance))
}

func (f *TForm) SetMenu(value IComponent) {
    defer exceptionProc()
    Form_SetMenu(f.instance, CheckPtr(value))
}

func (f *TForm) PixelsPerInch() int32 {
    defer exceptionProc()
    return Form_GetPixelsPerInch(f.instance)
}

func (f *TForm) SetPixelsPerInch(value int32) {
    defer exceptionProc()
    Form_SetPixelsPerInch(f.instance, value)
}

func (f *TForm) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(Form_GetPopupMenu(f.instance))
}

func (f *TForm) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    Form_SetPopupMenu(f.instance, CheckPtr(value))
}

func (f *TForm) Position() TPosition {
    defer exceptionProc()
    return Form_GetPosition(f.instance)
}

func (f *TForm) SetPosition(value TPosition) {
    defer exceptionProc()
    Form_SetPosition(f.instance, value)
}

func (f *TForm) Scaled() bool {
    defer exceptionProc()
    return Form_GetScaled(f.instance)
}

func (f *TForm) SetScaled(value bool) {
    defer exceptionProc()
    Form_SetScaled(f.instance, value)
}

func (f *TForm) ShowHint() bool {
    defer exceptionProc()
    return Form_GetShowHint(f.instance)
}

func (f *TForm) SetShowHint(value bool) {
    defer exceptionProc()
    Form_SetShowHint(f.instance, value)
}

func (f *TForm) Visible() bool {
    defer exceptionProc()
    return Form_GetVisible(f.instance)
}

func (f *TForm) SetVisible(value bool) {
    defer exceptionProc()
    Form_SetVisible(f.instance, value)
}

func (f *TForm) Width() int32 {
    defer exceptionProc()
    return Form_GetWidth(f.instance)
}

func (f *TForm) SetWidth(value int32) {
    defer exceptionProc()
    Form_SetWidth(f.instance, value)
}

func (f *TForm) WindowState() TWindowState {
    defer exceptionProc()
    return Form_GetWindowState(f.instance)
}

func (f *TForm) SetWindowState(value TWindowState) {
    defer exceptionProc()
    Form_SetWindowState(f.instance, value)
}

func (f *TForm) SetOnClick(fn TNotifyEvent) {
    Form_SetOnClick(f.instance, fn)
}

func (f *TForm) SetOnClose(fn TCloseEvent) {
    Form_SetOnClose(f.instance, fn)
}

func (f *TForm) SetOnCloseQuery(fn TCloseQueryEvent) {
    Form_SetOnCloseQuery(f.instance, fn)
}

func (f *TForm) SetOnDblClick(fn TNotifyEvent) {
    Form_SetOnDblClick(f.instance, fn)
}

func (f *TForm) SetOnHide(fn TNotifyEvent) {
    Form_SetOnHide(f.instance, fn)
}

func (f *TForm) SetOnKeyDown(fn TKeyEvent) {
    Form_SetOnKeyDown(f.instance, fn)
}

func (f *TForm) SetOnKeyPress(fn TKeyPressEvent) {
    Form_SetOnKeyPress(f.instance, fn)
}

func (f *TForm) SetOnKeyUp(fn TKeyEvent) {
    Form_SetOnKeyUp(f.instance, fn)
}

func (f *TForm) SetOnMouseDown(fn TMouseEvent) {
    Form_SetOnMouseDown(f.instance, fn)
}

func (f *TForm) SetOnMouseEnter(fn TNotifyEvent) {
    Form_SetOnMouseEnter(f.instance, fn)
}

func (f *TForm) SetOnMouseLeave(fn TNotifyEvent) {
    Form_SetOnMouseLeave(f.instance, fn)
}

func (f *TForm) SetOnMouseMove(fn TMouseMoveEvent) {
    Form_SetOnMouseMove(f.instance, fn)
}

func (f *TForm) SetOnMouseUp(fn TMouseEvent) {
    Form_SetOnMouseUp(f.instance, fn)
}

func (f *TForm) SetOnMouseWheel(fn TMouseWheelEvent) {
    Form_SetOnMouseWheel(f.instance, fn)
}

func (f *TForm) SetOnPaint(fn TNotifyEvent) {
    Form_SetOnPaint(f.instance, fn)
}

func (f *TForm) SetOnResize(fn TNotifyEvent) {
    Form_SetOnResize(f.instance, fn)
}

func (f *TForm) SetOnShow(fn TNotifyEvent) {
    Form_SetOnShow(f.instance, fn)
}

func (f *TForm) Canvas() *TCanvas {
    defer exceptionProc()
    return CanvasFromInst(Form_GetCanvas(f.instance))
}

func (f *TForm) DropTarget() bool {
    defer exceptionProc()
    return Form_GetDropTarget(f.instance)
}

func (f *TForm) SetDropTarget(value bool) {
    defer exceptionProc()
    Form_SetDropTarget(f.instance, value)
}

func (f *TForm) ModalResult() TModalResult {
    defer exceptionProc()
    return Form_GetModalResult(f.instance)
}

func (f *TForm) SetModalResult(value TModalResult) {
    defer exceptionProc()
    Form_SetModalResult(f.instance, value)
}

func (f *TForm) Left() int32 {
    defer exceptionProc()
    return Form_GetLeft(f.instance)
}

func (f *TForm) SetLeft(value int32) {
    defer exceptionProc()
    Form_SetLeft(f.instance, value)
}

func (f *TForm) Top() int32 {
    defer exceptionProc()
    return Form_GetTop(f.instance)
}

func (f *TForm) SetTop(value int32) {
    defer exceptionProc()
    Form_SetTop(f.instance, value)
}

func (f *TForm) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(Form_GetBrush(f.instance))
}

func (f *TForm) ControlCount() int32 {
    defer exceptionProc()
    return Form_GetControlCount(f.instance)
}

func (f *TForm) Handle() HWND {
    defer exceptionProc()
    return Form_GetHandle(f.instance)
}

func (f *TForm) TabOrder() int16 {
    defer exceptionProc()
    return Form_GetTabOrder(f.instance)
}

func (f *TForm) SetTabOrder(value int16) {
    defer exceptionProc()
    Form_SetTabOrder(f.instance, value)
}

func (f *TForm) TabStop() bool {
    defer exceptionProc()
    return Form_GetTabStop(f.instance)
}

func (f *TForm) SetTabStop(value bool) {
    defer exceptionProc()
    Form_SetTabStop(f.instance, value)
}

func (f *TForm) BoundsRect() TRect {
    defer exceptionProc()
    return Form_GetBoundsRect(f.instance)
}

func (f *TForm) SetBoundsRect(value TRect) {
    defer exceptionProc()
    Form_SetBoundsRect(f.instance, value)
}

func (f *TForm) ClientRect() TRect {
    defer exceptionProc()
    return Form_GetClientRect(f.instance)
}

func (f *TForm) ExplicitLeft() int32 {
    defer exceptionProc()
    return Form_GetExplicitLeft(f.instance)
}

func (f *TForm) ExplicitTop() int32 {
    defer exceptionProc()
    return Form_GetExplicitTop(f.instance)
}

func (f *TForm) ExplicitWidth() int32 {
    defer exceptionProc()
    return Form_GetExplicitWidth(f.instance)
}

func (f *TForm) ExplicitHeight() int32 {
    defer exceptionProc()
    return Form_GetExplicitHeight(f.instance)
}

func (f *TForm) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(Form_GetParent(f.instance))
}

func (f *TForm) SetParent(value IControl) {
    defer exceptionProc()
    Form_SetParent(f.instance, CheckPtr(value))
}

func (f *TForm) AlignWithMargins() bool {
    defer exceptionProc()
    return Form_GetAlignWithMargins(f.instance)
}

func (f *TForm) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    Form_SetAlignWithMargins(f.instance, value)
}

func (f *TForm) Cursor() TCursor {
    defer exceptionProc()
    return Form_GetCursor(f.instance)
}

func (f *TForm) SetCursor(value TCursor) {
    defer exceptionProc()
    Form_SetCursor(f.instance, value)
}

func (f *TForm) Hint() string {
    defer exceptionProc()
    return Form_GetHint(f.instance)
}

func (f *TForm) SetHint(value string) {
    defer exceptionProc()
    Form_SetHint(f.instance, value)
}

func (f *TForm) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(Form_GetMargins(f.instance))
}

func (f *TForm) SetMargins(value *TMargins) {
    defer exceptionProc()
    Form_SetMargins(f.instance, CheckPtr(value))
}

func (f *TForm) ComponentCount() int32 {
    defer exceptionProc()
    return Form_GetComponentCount(f.instance)
}

func (f *TForm) ComponentIndex() int32 {
    defer exceptionProc()
    return Form_GetComponentIndex(f.instance)
}

func (f *TForm) SetComponentIndex(value int32) {
    defer exceptionProc()
    Form_SetComponentIndex(f.instance, value)
}

func (f *TForm) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Form_GetOwner(f.instance))
}

func (f *TForm) Name() string {
    defer exceptionProc()
    return Form_GetName(f.instance)
}

func (f *TForm) SetName(value string) {
    defer exceptionProc()
    Form_SetName(f.instance, value)
}

func (f *TForm) Tag() int {
    defer exceptionProc()
    return Form_GetTag(f.instance)
}

func (f *TForm) SetTag(value int) {
    defer exceptionProc()
    Form_SetTag(f.instance, value)
}

func (f *TForm) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(Form_GetControls(f.instance, Index))
}

func (f *TForm) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Form_GetComponents(f.instance, AIndex))
}

