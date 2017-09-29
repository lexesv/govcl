
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TPopupMenu struct {
    IComponent
    instance uintptr
}

func NewPopupMenu(owner IComponent) *TPopupMenu {
    p := new(TPopupMenu)
    p.instance = PopupMenu_Create(owner.Instance())
    return p
}

func PopupMenuFromInst(inst uintptr) *TPopupMenu {
    p := new(TPopupMenu)
    p.instance = inst
    return p
}

func PopupMenuFromObj(obj IObject) *TPopupMenu {
    p := new(TPopupMenu)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TPopupMenu) Free() {
    if p.instance != 0 {
        PopupMenu_Free(p.instance)
    }
}

func (p *TPopupMenu) Instance() uintptr {
    return p.instance
}

func (p *TPopupMenu) IsValid() bool {
    return p.instance != 0
}

func (p *TPopupMenu) CloseMenu() {
    defer exceptionProc()
    PopupMenu_CloseMenu(p.instance)
}

func (p *TPopupMenu) Popup(X int32, Y int32) {
    defer exceptionProc()
    PopupMenu_Popup(p.instance, X , Y )
}

func (p *TPopupMenu) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(PopupMenu_FindComponent(p.instance, AName ))
}

func (p *TPopupMenu) GetNamePath() string {
    defer exceptionProc()
    return PopupMenu_GetNamePath(p.instance)
}

func (p *TPopupMenu) HasParent() bool {
    defer exceptionProc()
    return PopupMenu_HasParent(p.instance)
}

func (p *TPopupMenu) Assign(Source IObject) {
    defer exceptionProc()
    PopupMenu_Assign(p.instance, CheckPtr(Source))
}

func (p *TPopupMenu) ClassName() string {
    defer exceptionProc()
    return PopupMenu_ClassName(p.instance)
}

func (p *TPopupMenu) Equals(Obj IObject) bool {
    defer exceptionProc()
    return PopupMenu_Equals(p.instance, CheckPtr(Obj))
}

func (p *TPopupMenu) GetHashCode() int32 {
    defer exceptionProc()
    return PopupMenu_GetHashCode(p.instance)
}

func (p *TPopupMenu) ToString() string {
    defer exceptionProc()
    return PopupMenu_ToString(p.instance)
}

func (p *TPopupMenu) Alignment() TPopupAlignment {
    defer exceptionProc()
    return PopupMenu_GetAlignment(p.instance)
}

func (p *TPopupMenu) SetAlignment(value TPopupAlignment) {
    defer exceptionProc()
    PopupMenu_SetAlignment(p.instance, value)
}

func (p *TPopupMenu) AutoHotkeys() TMenuAutoFlag {
    defer exceptionProc()
    return PopupMenu_GetAutoHotkeys(p.instance)
}

func (p *TPopupMenu) SetAutoHotkeys(value TMenuAutoFlag) {
    defer exceptionProc()
    PopupMenu_SetAutoHotkeys(p.instance, value)
}

func (p *TPopupMenu) Images() *TImageList {
    defer exceptionProc()
    return ImageListFromInst(PopupMenu_GetImages(p.instance))
}

func (p *TPopupMenu) SetImages(value IComponent) {
    defer exceptionProc()
    PopupMenu_SetImages(p.instance, CheckPtr(value))
}

func (p *TPopupMenu) SetOnChange(fn TMenuChangeEvent) {
    PopupMenu_SetOnChange(p.instance, fn)
}

func (p *TPopupMenu) SetOnPopup(fn TNotifyEvent) {
    PopupMenu_SetOnPopup(p.instance, fn)
}

func (p *TPopupMenu) Handle() HMENU {
    defer exceptionProc()
    return PopupMenu_GetHandle(p.instance)
}

func (p *TPopupMenu) WindowHandle() HWND {
    defer exceptionProc()
    return PopupMenu_GetWindowHandle(p.instance)
}

func (p *TPopupMenu) SetWindowHandle(value HWND) {
    defer exceptionProc()
    PopupMenu_SetWindowHandle(p.instance, value)
}

func (p *TPopupMenu) Items() *TMenuItem {
    defer exceptionProc()
    return MenuItemFromInst(PopupMenu_GetItems(p.instance))
}

func (p *TPopupMenu) ComponentCount() int32 {
    defer exceptionProc()
    return PopupMenu_GetComponentCount(p.instance)
}

func (p *TPopupMenu) ComponentIndex() int32 {
    defer exceptionProc()
    return PopupMenu_GetComponentIndex(p.instance)
}

func (p *TPopupMenu) SetComponentIndex(value int32) {
    defer exceptionProc()
    PopupMenu_SetComponentIndex(p.instance, value)
}

func (p *TPopupMenu) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(PopupMenu_GetOwner(p.instance))
}

func (p *TPopupMenu) Name() string {
    defer exceptionProc()
    return PopupMenu_GetName(p.instance)
}

func (p *TPopupMenu) SetName(value string) {
    defer exceptionProc()
    PopupMenu_SetName(p.instance, value)
}

func (p *TPopupMenu) Tag() int {
    defer exceptionProc()
    return PopupMenu_GetTag(p.instance)
}

func (p *TPopupMenu) SetTag(value int) {
    defer exceptionProc()
    PopupMenu_SetTag(p.instance, value)
}

func (p *TPopupMenu) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(PopupMenu_GetComponents(p.instance, AIndex))
}

