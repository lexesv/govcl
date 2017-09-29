
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TImage struct {
    IControl
    instance uintptr
}

func NewImage(owner IComponent) *TImage {
    i := new(TImage)
    i.instance = Image_Create(owner.Instance())
    return i
}

func ImageFromInst(inst uintptr) *TImage {
    i := new(TImage)
    i.instance = inst
    return i
}

func ImageFromObj(obj IObject) *TImage {
    i := new(TImage)
    i.instance = CheckPtr(obj)
    return i
}

func (i *TImage) Free() {
    if i.instance != 0 {
        Image_Free(i.instance)
    }
}

func (i *TImage) Instance() uintptr {
    return i.instance
}

func (i *TImage) IsValid() bool {
    return i.instance != 0
}

func (i *TImage) BringToFront() {
    defer exceptionProc()
    Image_BringToFront(i.instance)
}

func (i *TImage) HasParent() bool {
    defer exceptionProc()
    return Image_HasParent(i.instance)
}

func (i *TImage) Hide() {
    defer exceptionProc()
    Image_Hide(i.instance)
}

func (i *TImage) Invalidate() {
    defer exceptionProc()
    Image_Invalidate(i.instance)
}

func (i *TImage) Perform(Msg uint32, WParam uintptr, LParam int) int {
    defer exceptionProc()
    return Image_Perform(i.instance, Msg , WParam , LParam )
}

func (i *TImage) Refresh() {
    defer exceptionProc()
    Image_Refresh(i.instance)
}

func (i *TImage) Repaint() {
    defer exceptionProc()
    Image_Repaint(i.instance)
}

func (i *TImage) SendToBack() {
    defer exceptionProc()
    Image_SendToBack(i.instance)
}

func (i *TImage) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    defer exceptionProc()
    Image_SetBounds(i.instance, ALeft , ATop , AWidth , AHeight )
}

func (i *TImage) Show() {
    defer exceptionProc()
    Image_Show(i.instance)
}

func (i *TImage) Update() {
    defer exceptionProc()
    Image_Update(i.instance)
}

func (i *TImage) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return Image_GetTextBuf(i.instance, Buffer , BufSize )
}

func (i *TImage) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Image_FindComponent(i.instance, AName ))
}

func (i *TImage) GetNamePath() string {
    defer exceptionProc()
    return Image_GetNamePath(i.instance)
}

func (i *TImage) Assign(Source IObject) {
    defer exceptionProc()
    Image_Assign(i.instance, CheckPtr(Source))
}

func (i *TImage) ClassName() string {
    defer exceptionProc()
    return Image_ClassName(i.instance)
}

func (i *TImage) Equals(Obj IObject) bool {
    defer exceptionProc()
    return Image_Equals(i.instance, CheckPtr(Obj))
}

func (i *TImage) GetHashCode() int32 {
    defer exceptionProc()
    return Image_GetHashCode(i.instance)
}

func (i *TImage) ToString() string {
    defer exceptionProc()
    return Image_ToString(i.instance)
}

func (i *TImage) Canvas() *TCanvas {
    defer exceptionProc()
    return CanvasFromInst(Image_GetCanvas(i.instance))
}

func (i *TImage) Align() TAlign {
    defer exceptionProc()
    return Image_GetAlign(i.instance)
}

func (i *TImage) SetAlign(value TAlign) {
    defer exceptionProc()
    Image_SetAlign(i.instance, value)
}

func (i *TImage) Anchors() TAnchors {
    defer exceptionProc()
    return Image_GetAnchors(i.instance)
}

func (i *TImage) SetAnchors(value TAnchors) {
    defer exceptionProc()
    Image_SetAnchors(i.instance, value)
}

func (i *TImage) AutoSize() bool {
    defer exceptionProc()
    return Image_GetAutoSize(i.instance)
}

func (i *TImage) SetAutoSize(value bool) {
    defer exceptionProc()
    Image_SetAutoSize(i.instance, value)
}

func (i *TImage) Center() bool {
    defer exceptionProc()
    return Image_GetCenter(i.instance)
}

func (i *TImage) SetCenter(value bool) {
    defer exceptionProc()
    Image_SetCenter(i.instance, value)
}

func (i *TImage) Enabled() bool {
    defer exceptionProc()
    return Image_GetEnabled(i.instance)
}

func (i *TImage) SetEnabled(value bool) {
    defer exceptionProc()
    Image_SetEnabled(i.instance, value)
}

func (i *TImage) IncrementalDisplay() bool {
    defer exceptionProc()
    return Image_GetIncrementalDisplay(i.instance)
}

func (i *TImage) SetIncrementalDisplay(value bool) {
    defer exceptionProc()
    Image_SetIncrementalDisplay(i.instance, value)
}

func (i *TImage) ParentShowHint() bool {
    defer exceptionProc()
    return Image_GetParentShowHint(i.instance)
}

func (i *TImage) SetParentShowHint(value bool) {
    defer exceptionProc()
    Image_SetParentShowHint(i.instance, value)
}

func (i *TImage) Picture() *TPicture {
    defer exceptionProc()
    return PictureFromInst(Image_GetPicture(i.instance))
}

func (i *TImage) SetPicture(value *TPicture) {
    defer exceptionProc()
    Image_SetPicture(i.instance, CheckPtr(value))
}

func (i *TImage) PopupMenu() *TPopupMenu {
    defer exceptionProc()
    return PopupMenuFromInst(Image_GetPopupMenu(i.instance))
}

func (i *TImage) SetPopupMenu(value IComponent) {
    defer exceptionProc()
    Image_SetPopupMenu(i.instance, CheckPtr(value))
}

func (i *TImage) Proportional() bool {
    defer exceptionProc()
    return Image_GetProportional(i.instance)
}

func (i *TImage) SetProportional(value bool) {
    defer exceptionProc()
    Image_SetProportional(i.instance, value)
}

func (i *TImage) ShowHint() bool {
    defer exceptionProc()
    return Image_GetShowHint(i.instance)
}

func (i *TImage) SetShowHint(value bool) {
    defer exceptionProc()
    Image_SetShowHint(i.instance, value)
}

func (i *TImage) Stretch() bool {
    defer exceptionProc()
    return Image_GetStretch(i.instance)
}

func (i *TImage) SetStretch(value bool) {
    defer exceptionProc()
    Image_SetStretch(i.instance, value)
}

func (i *TImage) Transparent() bool {
    defer exceptionProc()
    return Image_GetTransparent(i.instance)
}

func (i *TImage) SetTransparent(value bool) {
    defer exceptionProc()
    Image_SetTransparent(i.instance, value)
}

func (i *TImage) Visible() bool {
    defer exceptionProc()
    return Image_GetVisible(i.instance)
}

func (i *TImage) SetVisible(value bool) {
    defer exceptionProc()
    Image_SetVisible(i.instance, value)
}

func (i *TImage) SetOnClick(fn TNotifyEvent) {
    Image_SetOnClick(i.instance, fn)
}

func (i *TImage) SetOnDblClick(fn TNotifyEvent) {
    Image_SetOnDblClick(i.instance, fn)
}

func (i *TImage) SetOnMouseDown(fn TMouseEvent) {
    Image_SetOnMouseDown(i.instance, fn)
}

func (i *TImage) SetOnMouseEnter(fn TNotifyEvent) {
    Image_SetOnMouseEnter(i.instance, fn)
}

func (i *TImage) SetOnMouseLeave(fn TNotifyEvent) {
    Image_SetOnMouseLeave(i.instance, fn)
}

func (i *TImage) SetOnMouseMove(fn TMouseMoveEvent) {
    Image_SetOnMouseMove(i.instance, fn)
}

func (i *TImage) SetOnMouseUp(fn TMouseEvent) {
    Image_SetOnMouseUp(i.instance, fn)
}

func (i *TImage) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(Image_GetAction(i.instance))
}

func (i *TImage) SetAction(value IComponent) {
    defer exceptionProc()
    Image_SetAction(i.instance, CheckPtr(value))
}

func (i *TImage) BoundsRect() TRect {
    defer exceptionProc()
    return Image_GetBoundsRect(i.instance)
}

func (i *TImage) SetBoundsRect(value TRect) {
    defer exceptionProc()
    Image_SetBoundsRect(i.instance, value)
}

func (i *TImage) ClientHeight() int32 {
    defer exceptionProc()
    return Image_GetClientHeight(i.instance)
}

func (i *TImage) SetClientHeight(value int32) {
    defer exceptionProc()
    Image_SetClientHeight(i.instance, value)
}

func (i *TImage) ClientRect() TRect {
    defer exceptionProc()
    return Image_GetClientRect(i.instance)
}

func (i *TImage) ClientWidth() int32 {
    defer exceptionProc()
    return Image_GetClientWidth(i.instance)
}

func (i *TImage) SetClientWidth(value int32) {
    defer exceptionProc()
    Image_SetClientWidth(i.instance, value)
}

func (i *TImage) ExplicitLeft() int32 {
    defer exceptionProc()
    return Image_GetExplicitLeft(i.instance)
}

func (i *TImage) ExplicitTop() int32 {
    defer exceptionProc()
    return Image_GetExplicitTop(i.instance)
}

func (i *TImage) ExplicitWidth() int32 {
    defer exceptionProc()
    return Image_GetExplicitWidth(i.instance)
}

func (i *TImage) ExplicitHeight() int32 {
    defer exceptionProc()
    return Image_GetExplicitHeight(i.instance)
}

func (i *TImage) Parent() *TControl {
    defer exceptionProc()
    return ControlFromInst(Image_GetParent(i.instance))
}

func (i *TImage) SetParent(value IControl) {
    defer exceptionProc()
    Image_SetParent(i.instance, CheckPtr(value))
}

func (i *TImage) AlignWithMargins() bool {
    defer exceptionProc()
    return Image_GetAlignWithMargins(i.instance)
}

func (i *TImage) SetAlignWithMargins(value bool) {
    defer exceptionProc()
    Image_SetAlignWithMargins(i.instance, value)
}

func (i *TImage) Left() int32 {
    defer exceptionProc()
    return Image_GetLeft(i.instance)
}

func (i *TImage) SetLeft(value int32) {
    defer exceptionProc()
    Image_SetLeft(i.instance, value)
}

func (i *TImage) Top() int32 {
    defer exceptionProc()
    return Image_GetTop(i.instance)
}

func (i *TImage) SetTop(value int32) {
    defer exceptionProc()
    Image_SetTop(i.instance, value)
}

func (i *TImage) Width() int32 {
    defer exceptionProc()
    return Image_GetWidth(i.instance)
}

func (i *TImage) SetWidth(value int32) {
    defer exceptionProc()
    Image_SetWidth(i.instance, value)
}

func (i *TImage) Height() int32 {
    defer exceptionProc()
    return Image_GetHeight(i.instance)
}

func (i *TImage) SetHeight(value int32) {
    defer exceptionProc()
    Image_SetHeight(i.instance, value)
}

func (i *TImage) Cursor() TCursor {
    defer exceptionProc()
    return Image_GetCursor(i.instance)
}

func (i *TImage) SetCursor(value TCursor) {
    defer exceptionProc()
    Image_SetCursor(i.instance, value)
}

func (i *TImage) Hint() string {
    defer exceptionProc()
    return Image_GetHint(i.instance)
}

func (i *TImage) SetHint(value string) {
    defer exceptionProc()
    Image_SetHint(i.instance, value)
}

func (i *TImage) Margins() *TMargins {
    defer exceptionProc()
    return MarginsFromInst(Image_GetMargins(i.instance))
}

func (i *TImage) SetMargins(value *TMargins) {
    defer exceptionProc()
    Image_SetMargins(i.instance, CheckPtr(value))
}

func (i *TImage) ComponentCount() int32 {
    defer exceptionProc()
    return Image_GetComponentCount(i.instance)
}

func (i *TImage) ComponentIndex() int32 {
    defer exceptionProc()
    return Image_GetComponentIndex(i.instance)
}

func (i *TImage) SetComponentIndex(value int32) {
    defer exceptionProc()
    Image_SetComponentIndex(i.instance, value)
}

func (i *TImage) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Image_GetOwner(i.instance))
}

func (i *TImage) Name() string {
    defer exceptionProc()
    return Image_GetName(i.instance)
}

func (i *TImage) SetName(value string) {
    defer exceptionProc()
    Image_SetName(i.instance, value)
}

func (i *TImage) Tag() int {
    defer exceptionProc()
    return Image_GetTag(i.instance)
}

func (i *TImage) SetTag(value int) {
    defer exceptionProc()
    Image_SetTag(i.instance, value)
}

func (i *TImage) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Image_GetComponents(i.instance, AIndex))
}

