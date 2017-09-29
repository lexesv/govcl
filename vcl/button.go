
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TButton struct {
    IControl
    instance uintptr
}

func NewButton(owner IComponent) *TButton {
    b := new(TButton)
    b.instance = Button_Create(owner.Instance())
    return b
}

func ButtonFromInst(inst uintptr) *TButton {
    b := new(TButton)
    b.instance = inst
    return b
}

func ButtonFromObj(obj IObject) *TButton {
    b := new(TButton)
    b.instance = CheckPtr(obj)
    return b
}

func (b *TButton) Free() {
    if b.instance != 0 {
        Button_Free(b.instance)
    }
}

func (b *TButton) Instance() uintptr {
    return b.instance
}

func (b *TButton) IsValid() bool {
    return b.instance != 0
}

func (b *TButton) Click() {
    defer exceptionProc()
    Button_Click(b.instance)
}

func (b *TButton) CanFocus() bool {
    defer exceptionProc()
    return Button_CanFocus(b.instance)
}

func (b *TButton) FlipChildren(AllLevels bool) {
    defer exceptionProc()
    Button_FlipChildren(b.instance, AllLevels )
}

func (b *TButton) Focused() bool {
    defer exceptionProc()
    return Button_Focused(b.instance)
}

func (b *TButton) HandleAllocated() bool {
    defer exceptionProc()
    return Button_HandleAllocated(b.instance)
}

func (b *TButton) Invalidate() {
    defer exceptionProc()
    Button_Invalidate(b.instance)
}

func (b *TButton) Realign() {
    defer exceptionProc()
    Button_Realign(b.instance)
}

func (b *TButton) Repaint() {
    defer exceptionProc()
    Button_Repaint(b.instance)
}

func (b *TButton) ScaleBy(M int32, D int32) {
    defer exceptionProc()
    Button_ScaleBy(b.instance, M , D )
}

func (b *TButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    Button_SetBounds(b.instance, ALeft , ATop , AWidth , AHeight )
}

func (b *TButton) SetFocus() {
    defer exceptionProc()
    Button_SetFocus(b.instance)
}

func (b *TButton) Update() {
    defer exceptionProc()
    Button_Update(b.instance)
}

func (b *TButton) BringToFront() {
    defer exceptionProc()
    Button_BringToFront(b.instance)
}

func (b *TButton) HasParent() bool {
    defer exceptionProc()
    return Button_HasParent(b.instance)
}

func (b *TButton) Hide() {
    defer exceptionProc()
    Button_Hide(b.instance)
}

func (b *TButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return Button_Perform(b.instance, Msg , WParam , LParam )
}

func (b *TButton) Refresh() {
    defer exceptionProc()
    Button_Refresh(b.instance)
}

func (b *TButton) SendToBack() {
    defer exceptionProc()
    Button_SendToBack(b.instance)
}

func (b *TButton) Show() {
    defer exceptionProc()
    Button_Show(b.instance)
}

func (b *TButton) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return Button_GetTextBuf(b.instance, Buffer , BufSize )
}

func (b *TButton) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Button_FindComponent(b.instance, AName ))
}

func (b *TButton) GetNamePath() string {
    defer exceptionProc()
    return Button_GetNamePath(b.instance)
}

func (b *TButton) Assign(Source IObject) {
    defer exceptionProc()
    Button_Assign(b.instance, CheckPtr(Source))
}

func (b *TButton) ClassName() string {
    defer exceptionProc()
    return Button_ClassName(b.instance)
}

func (b *TButton) Equals(Obj IObject) bool {
    defer exceptionProc()
    return Button_Equals(b.instance, CheckPtr(Obj))
}

func (b *TButton) GetHashCode() int32 {
    defer exceptionProc()
    return Button_GetHashCode(b.instance)
}

func (b *TButton) ToString() string {
    defer exceptionProc()
    return Button_ToString(b.instance)
}

func (b *TButton) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(Button_GetAction(b.instance))
}

func (b *TButton) SetAction(value IComponent) {
    defer exceptionProc()
    Button_SetAction(b.instance, CheckPtr(value))
}

func (b *TButton) Align() TAlign {
    defer exceptionProc()
    return Button_GetAlign(b.instance)
}

func (b *TButton) SetAlign(value TAlign) {
    defer exceptionProc()
    Button_SetAlign(b.instance, value)
}

func (b *TButton) Anchors() TAnchors {
    defer exceptionProc()
    return Button_GetAnchors(b.instance)
}

func (b *TButton) SetAnchors(value TAnchors) {
    defer exceptionProc()
    Button_SetAnchors(b.instance, value)
}

func (b *TButton) Cancel() bool {
    defer exceptionProc()
    return Button_GetCancel(b.instance)
}

func (b *TButton) SetCancel(value bool) {
    defer exceptionProc()
    Button_SetCancel(b.instance, value)
}

func (b *TButton) Caption() string {
    defer exceptionProc()
    return Button_GetCaption(b.instance)
}

func (b *TButton) SetCaption(value string) {
    defer exceptionProc()
    Button_SetCaption(b.instance, value)
}

func (b *TButton) Default() bool {
    defer exceptionProc()
    return Button_GetDefault(b.instance)
}

func (b *TButton) SetDefault(value bool) {
    defer exceptionProc()
    Button_SetDefault(b.instance, value)
}

func (b *TButton) DoubleBuffered() bool {
    defer exceptionProc()
    return Button_GetDoubleBuffered(b.instance)
}

func (b *TButton) SetDoubleBuffered(value bool) {
    defer exceptionProc()
    Button_SetDoubleBuffered(b.instance, value)
}

func (b *TButton) Enabled() bool {
    defer exceptionProc()
    return Button_GetEnabled(b.instance)
}

func (b *TButton) SetEnabled(value bool) {
    defer exceptionProc()
    Button_SetEnabled(b.instance, value)
}

func (b *TButton) Font() *TFont {
    defer exceptionProc()
    return FontFromInst(Button_GetFont(b.instance))
}

func (b *TButton) SetFont(value *TFont) {
    defer exceptionProc()
    Button_SetFont(b.instance, CheckPtr(value))
}

func (b *TButton) ImageIndex() int32 {
    defer exceptionProc()
    return Button_GetImageIndex(b.instance)
}

func (b *TButton) SetImageIndex(value int32) {
    defer exceptionProc()
    Button_SetImageIndex(b.instance, value)
}

func (b *TButton) Images() *TImageList {
    defer exceptionProc()
    return ImageListFromInst(Button_GetImages(b.instance))
}

func (b *TButton) SetImages(value IComponent) {
    defer exceptionProc()
    Button_SetImages(b.instance, CheckPtr(value))
}

func (b *TButton) ModalResult() TModalResult {
    defer exceptionProc()
    return Button_GetModalResult(b.instance)
}

func (b *TButton) SetModalResult(value TModalResult) {
    defer exceptionProc()
    Button_SetModalResult(b.instance, value)
}

func (b *TButton) ParentFont() bool {
    defer exceptionProc()
    return Button_GetParentFont(b.instance)
}

func (b *TButton) SetParentFont(value bool) {
    defer exceptionProc()
    Button_SetParentFont(b.instance, value)
}

func (b *TButton) ParentShowHint() bool {
    defer exceptionProc()
    return Button_GetParentShowHint(b.instance)
}

func (b *TButton) SetParentShowHint(value bool) {
    defer exceptionProc()
    Button_SetParentShowHint(b.instance, value)
}

func (b *TButton) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(Button_GetPopupMenu(b.instance))
}

func (b *TButton) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    Button_SetPopupMenu(b.instance, CheckPtr(value))
}

func (b *TButton) ShowHint() bool {
    defer exceptionProc()
    return Button_GetShowHint(b.instance)
}

func (b *TButton) SetShowHint(value bool) {
    defer exceptionProc()
    Button_SetShowHint(b.instance, value)
}

func (b *TButton) Style() TButtonStyle {
    defer exceptionProc()
    return Button_GetStyle(b.instance)
}

func (b *TButton) SetStyle(value TButtonStyle) {
    defer exceptionProc()
    Button_SetStyle(b.instance, value)
}

func (b *TButton) TabOrder() int16 {
    defer exceptionProc()
    return Button_GetTabOrder(b.instance)
}

func (b *TButton) SetTabOrder(value int16) {
    defer exceptionProc()
    Button_SetTabOrder(b.instance, value)
}

func (b *TButton) TabStop() bool {
    defer exceptionProc()
    return Button_GetTabStop(b.instance)
}

func (b *TButton) SetTabStop(value bool) {
    defer exceptionProc()
    Button_SetTabStop(b.instance, value)
}

func (b *TButton) Visible() bool {
    defer exceptionProc()
    return Button_GetVisible(b.instance)
}

func (b *TButton) SetVisible(value bool) {
    defer exceptionProc()
    Button_SetVisible(b.instance, value)
}

func (b *TButton) WordWrap() bool {
    defer exceptionProc()
    return Button_GetWordWrap(b.instance)
}

func (b *TButton) SetWordWrap(value bool) {
    defer exceptionProc()
    Button_SetWordWrap(b.instance, value)
}

func (b *TButton) SetOnClick(fn TNotifyEvent) {
    Button_SetOnClick(b.instance, fn)
}

func (b *TButton) SetOnEnter(fn TNotifyEvent) {
    Button_SetOnEnter(b.instance, fn)
}

func (b *TButton) SetOnExit(fn TNotifyEvent) {
    Button_SetOnExit(b.instance, fn)
}

func (b *TButton) SetOnKeyDown(fn TKeyEvent) {
    Button_SetOnKeyDown(b.instance, fn)
}

func (b *TButton) SetOnKeyPress(fn TKeyPressEvent) {
    Button_SetOnKeyPress(b.instance, fn)
}

func (b *TButton) SetOnKeyUp(fn TKeyEvent) {
    Button_SetOnKeyUp(b.instance, fn)
}

func (b *TButton) SetOnMouseDown(fn TMouseEvent) {
    Button_SetOnMouseDown(b.instance, fn)
}

func (b *TButton) SetOnMouseEnter(fn TNotifyEvent) {
    Button_SetOnMouseEnter(b.instance, fn)
}

func (b *TButton) SetOnMouseLeave(fn TNotifyEvent) {
    Button_SetOnMouseLeave(b.instance, fn)
}

func (b *TButton) SetOnMouseMove(fn TMouseMoveEvent) {
    Button_SetOnMouseMove(b.instance, fn)
}

func (b *TButton) SetOnMouseUp(fn TMouseEvent) {
    Button_SetOnMouseUp(b.instance, fn)
}

func (b *TButton) Brush() *TBrush {
    defer exceptionProc()
    return BrushFromInst(Button_GetBrush(b.instance))
}

func (b *TButton) ControlCount() int32 {
    defer exceptionProc()
    return Button_GetControlCount(b.instance)
}

func (b *TButton) Handle() HWND {
    defer exceptionProc()
    return Button_GetHandle(b.instance)
}

func (b *TButton) BoundsRect() TRect {
    defer exceptionProc()
    return Button_GetBoundsRect(b.instance)
}

func (b *TButton) SetBoundsRect(value TRect) {
    defer exceptionProc()
    Button_SetBoundsRect(b.instance, value)
}

func (b *TButton) ClientHeight() int32 {
    defer exceptionProc()
    return Button_GetClientHeight(b.instance)
}

func (b *TButton) SetClientHeight(value int32) {
    defer exceptionProc()
    Button_SetClientHeight(b.instance, value)
}

func (b *TButton) ClientRect() TRect {
    defer exceptionProc()
    return Button_GetClientRect(b.instance)
}

func (b *TButton) ClientWidth() int32 {
    defer exceptionProc()
    return Button_GetClientWidth(b.instance)
}

func (b *TButton) SetClientWidth(value int32) {
    defer exceptionProc()
    Button_SetClientWidth(b.instance, value)
}

func (b *TButton) ExplicitLeft() int32 {
    defer exceptionProc()
    return Button_GetExplicitLeft(b.instance)
}

func (b *TButton) ExplicitTop() int32 {
    defer exceptionProc()
    return Button_GetExplicitTop(b.instance)
}

func (b *TButton) ExplicitWidth() int32 {
    defer exceptionProc()
    return Button_GetExplicitWidth(b.instance)
}

func (b *TButton) ExplicitHeight() int32 {
    defer exceptionProc()
    return Button_GetExplicitHeight(b.instance)
}

func (b *TButton) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(Button_GetParent(b.instance))
}

func (b *TButton) SetParent(value IControl) {
    defer exceptionProc()
    Button_SetParent(b.instance, CheckPtr(value))
}

func (b *TButton) AlignWithMargins() bool {
    defer exceptionProc()
    return Button_GetAlignWithMargins(b.instance)
}

func (b *TButton) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    Button_SetAlignWithMargins(b.instance, value)
}

func (b *TButton) Left() int32 {
    defer exceptionProc()
    return Button_GetLeft(b.instance)
}

func (b *TButton) SetLeft(value int32) {
    defer exceptionProc()
    Button_SetLeft(b.instance, value)
}

func (b *TButton) Top() int32 {
    defer exceptionProc()
    return Button_GetTop(b.instance)
}

func (b *TButton) SetTop(value int32) {
    defer exceptionProc()
    Button_SetTop(b.instance, value)
}

func (b *TButton) Width() int32 {
    defer exceptionProc()
    return Button_GetWidth(b.instance)
}

func (b *TButton) SetWidth(value int32) {
    defer exceptionProc()
    Button_SetWidth(b.instance, value)
}

func (b *TButton) Height() int32 {
    defer exceptionProc()
    return Button_GetHeight(b.instance)
}

func (b *TButton) SetHeight(value int32) {
    defer exceptionProc()
    Button_SetHeight(b.instance, value)
}

func (b *TButton) Cursor() TCursor {
    defer exceptionProc()
    return Button_GetCursor(b.instance)
}

func (b *TButton) SetCursor(value TCursor) {
    defer exceptionProc()
    Button_SetCursor(b.instance, value)
}

func (b *TButton) Hint() string {
    defer exceptionProc()
    return Button_GetHint(b.instance)
}

func (b *TButton) SetHint(value string) {
    defer exceptionProc()
    Button_SetHint(b.instance, value)
}

func (b *TButton) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(Button_GetMargins(b.instance))
}

func (b *TButton) SetMargins(value *TMargins) {
    defer exceptionProc()
    Button_SetMargins(b.instance, CheckPtr(value))
}

func (b *TButton) ComponentCount() int32 {
    defer exceptionProc()
    return Button_GetComponentCount(b.instance)
}

func (b *TButton) ComponentIndex() int32 {
    defer exceptionProc()
    return Button_GetComponentIndex(b.instance)
}

func (b *TButton) SetComponentIndex(value int32) {
    defer exceptionProc()
    Button_SetComponentIndex(b.instance, value)
}

func (b *TButton) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Button_GetOwner(b.instance))
}

func (b *TButton) Name() string {
    defer exceptionProc()
    return Button_GetName(b.instance)
}

func (b *TButton) SetName(value string) {
    defer exceptionProc()
    Button_SetName(b.instance, value)
}

func (b *TButton) Tag() int {
    defer exceptionProc()
    return Button_GetTag(b.instance)
}

func (b *TButton) SetTag(value int) {
    defer exceptionProc()
    Button_SetTag(b.instance, value)
}

func (b *TButton) Controls(Index int32) *TControl {
    defer exceptionProc()
    return ControlFromInst(Button_GetControls(b.instance, Index))
}

func (b *TButton) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Button_GetComponents(b.instance, AIndex))
}

