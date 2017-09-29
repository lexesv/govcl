
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TMenuItem struct {
    IComponent
    instance uintptr
}

func NewMenuItem(owner IComponent) *TMenuItem {
    m := new(TMenuItem)
    m.instance = MenuItem_Create(owner.Instance())
    return m
}

func MenuItemFromInst(inst uintptr) *TMenuItem {
    m := new(TMenuItem)
    m.instance = inst
    return m
}

func MenuItemFromObj(obj IObject) *TMenuItem {
    m := new(TMenuItem)
    m.instance = CheckPtr(obj)
    return m
}

func (m *TMenuItem) Free() {
    if m.instance != 0 {
        MenuItem_Free(m.instance)
    }
}

func (m *TMenuItem) Instance() uintptr {
    return m.instance
}

func (m *TMenuItem) IsValid() bool {
    return m.instance != 0
}

func (m *TMenuItem) Insert(Index int32, Item IComponent) {
    defer exceptionProc()
    MenuItem_Insert(m.instance, Index , CheckPtr(Item))
}

func (m *TMenuItem) Delete(Index int32) {
    defer exceptionProc()
    MenuItem_Delete(m.instance, Index )
}

func (m *TMenuItem) Clear() {
    defer exceptionProc()
    MenuItem_Clear(m.instance)
}

func (m *TMenuItem) Click() {
    defer exceptionProc()
    MenuItem_Click(m.instance)
}

func (m *TMenuItem) IndexOf(Item IComponent) int32 {
    defer exceptionProc()
    return MenuItem_IndexOf(m.instance, CheckPtr(Item))
}

func (m *TMenuItem) HasParent() bool {
    defer exceptionProc()
    return MenuItem_HasParent(m.instance)
}

func (m *TMenuItem) Add(Item IComponent) {
    defer exceptionProc()
    MenuItem_Add(m.instance, CheckPtr(Item))
}

func (m *TMenuItem) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(MenuItem_FindComponent(m.instance, AName ))
}

func (m *TMenuItem) GetNamePath() string {
    defer exceptionProc()
    return MenuItem_GetNamePath(m.instance)
}

func (m *TMenuItem) Assign(Source IObject) {
    defer exceptionProc()
    MenuItem_Assign(m.instance, CheckPtr(Source))
}

func (m *TMenuItem) ClassName() string {
    defer exceptionProc()
    return MenuItem_ClassName(m.instance)
}

func (m *TMenuItem) Equals(Obj IObject) bool {
    defer exceptionProc()
    return MenuItem_Equals(m.instance, CheckPtr(Obj))
}

func (m *TMenuItem) GetHashCode() int32 {
    defer exceptionProc()
    return MenuItem_GetHashCode(m.instance)
}

func (m *TMenuItem) ToString() string {
    defer exceptionProc()
    return MenuItem_ToString(m.instance)
}

func (m *TMenuItem) Handle() HMENU {
    defer exceptionProc()
    return MenuItem_GetHandle(m.instance)
}

func (m *TMenuItem) Parent() *TMenuItem {
    defer exceptionProc()
    return MenuItemFromInst(MenuItem_GetParent(m.instance))
}

func (m *TMenuItem) Action() *TAction {
    defer exceptionProc()
    return ActionFromInst(MenuItem_GetAction(m.instance))
}

func (m *TMenuItem) SetAction(value IComponent) {
    defer exceptionProc()
    MenuItem_SetAction(m.instance, CheckPtr(value))
}

func (m *TMenuItem) AutoHotkeys() TMenuItemAutoFlag {
    defer exceptionProc()
    return MenuItem_GetAutoHotkeys(m.instance)
}

func (m *TMenuItem) SetAutoHotkeys(value TMenuItemAutoFlag) {
    defer exceptionProc()
    MenuItem_SetAutoHotkeys(m.instance, value)
}

func (m *TMenuItem) Bitmap() *TBitmap {
    defer exceptionProc()
    return BitmapFromInst(MenuItem_GetBitmap(m.instance))
}

func (m *TMenuItem) SetBitmap(value *TBitmap) {
    defer exceptionProc()
    MenuItem_SetBitmap(m.instance, CheckPtr(value))
}

func (m *TMenuItem) Caption() string {
    defer exceptionProc()
    return MenuItem_GetCaption(m.instance)
}

func (m *TMenuItem) SetCaption(value string) {
    defer exceptionProc()
    MenuItem_SetCaption(m.instance, value)
}

func (m *TMenuItem) Checked() bool {
    defer exceptionProc()
    return MenuItem_GetChecked(m.instance)
}

func (m *TMenuItem) SetChecked(value bool) {
    defer exceptionProc()
    MenuItem_SetChecked(m.instance, value)
}

func (m *TMenuItem) Default() bool {
    defer exceptionProc()
    return MenuItem_GetDefault(m.instance)
}

func (m *TMenuItem) SetDefault(value bool) {
    defer exceptionProc()
    MenuItem_SetDefault(m.instance, value)
}

func (m *TMenuItem) Enabled() bool {
    defer exceptionProc()
    return MenuItem_GetEnabled(m.instance)
}

func (m *TMenuItem) SetEnabled(value bool) {
    defer exceptionProc()
    MenuItem_SetEnabled(m.instance, value)
}

func (m *TMenuItem) GroupIndex() uint8 {
    defer exceptionProc()
    return MenuItem_GetGroupIndex(m.instance)
}

func (m *TMenuItem) SetGroupIndex(value uint8) {
    defer exceptionProc()
    MenuItem_SetGroupIndex(m.instance, value)
}

func (m *TMenuItem) Hint() string {
    defer exceptionProc()
    return MenuItem_GetHint(m.instance)
}

func (m *TMenuItem) SetHint(value string) {
    defer exceptionProc()
    MenuItem_SetHint(m.instance, value)
}

func (m *TMenuItem) ImageIndex() int32 {
    defer exceptionProc()
    return MenuItem_GetImageIndex(m.instance)
}

func (m *TMenuItem) SetImageIndex(value int32) {
    defer exceptionProc()
    MenuItem_SetImageIndex(m.instance, value)
}

func (m *TMenuItem) ShortCut() TShortCut {
    defer exceptionProc()
    return MenuItem_GetShortCut(m.instance)
}

func (m *TMenuItem) SetShortCut(value TShortCut) {
    defer exceptionProc()
    MenuItem_SetShortCut(m.instance, value)
}

func (m *TMenuItem) Visible() bool {
    defer exceptionProc()
    return MenuItem_GetVisible(m.instance)
}

func (m *TMenuItem) SetVisible(value bool) {
    defer exceptionProc()
    MenuItem_SetVisible(m.instance, value)
}

func (m *TMenuItem) SetOnClick(fn TNotifyEvent) {
    MenuItem_SetOnClick(m.instance, fn)
}

func (m *TMenuItem) ComponentCount() int32 {
    defer exceptionProc()
    return MenuItem_GetComponentCount(m.instance)
}

func (m *TMenuItem) ComponentIndex() int32 {
    defer exceptionProc()
    return MenuItem_GetComponentIndex(m.instance)
}

func (m *TMenuItem) SetComponentIndex(value int32) {
    defer exceptionProc()
    MenuItem_SetComponentIndex(m.instance, value)
}

func (m *TMenuItem) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(MenuItem_GetOwner(m.instance))
}

func (m *TMenuItem) Name() string {
    defer exceptionProc()
    return MenuItem_GetName(m.instance)
}

func (m *TMenuItem) SetName(value string) {
    defer exceptionProc()
    MenuItem_SetName(m.instance, value)
}

func (m *TMenuItem) Tag() int {
    defer exceptionProc()
    return MenuItem_GetTag(m.instance)
}

func (m *TMenuItem) SetTag(value int) {
    defer exceptionProc()
    MenuItem_SetTag(m.instance, value)
}

func (m *TMenuItem) Items(Index int32) *TMenuItem {
    defer exceptionProc()
    return MenuItemFromInst(MenuItem_GetItems(m.instance, Index))
}

func (m *TMenuItem) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(MenuItem_GetComponents(m.instance, AIndex))
}

