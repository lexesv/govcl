
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TMainMenu struct {
    IComponent
    instance uintptr
}

func NewMainMenu(owner IComponent) *TMainMenu {
    m := new(TMainMenu)
    m.instance = MainMenu_Create(owner.Instance())
    return m
}

func MainMenuFromInst(inst uintptr) *TMainMenu {
    m := new(TMainMenu)
    m.instance = inst
    return m
}

func MainMenuFromObj(obj IObject) *TMainMenu {
    m := new(TMainMenu)
    m.instance = CheckPtr(obj)
    return m
}

func (m *TMainMenu) Free() {
    if m.instance != 0 {
        MainMenu_Free(m.instance)
    }
}

func (m *TMainMenu) Instance() uintptr {
    return m.instance
}

func (m *TMainMenu) IsValid() bool {
    return m.instance != 0
}

func (m *TMainMenu) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(MainMenu_FindComponent(m.instance, AName ))
}

func (m *TMainMenu) GetNamePath() string {
    defer exceptionProc()
    return MainMenu_GetNamePath(m.instance)
}

func (m *TMainMenu) HasParent() bool {
    defer exceptionProc()
    return MainMenu_HasParent(m.instance)
}

func (m *TMainMenu) Assign(Source IObject) {
    defer exceptionProc()
    MainMenu_Assign(m.instance, CheckPtr(Source))
}

func (m *TMainMenu) ClassName() string {
    defer exceptionProc()
    return MainMenu_ClassName(m.instance)
}

func (m *TMainMenu) Equals(Obj IObject) bool {
    defer exceptionProc()
    return MainMenu_Equals(m.instance, CheckPtr(Obj))
}

func (m *TMainMenu) GetHashCode() int32 {
    defer exceptionProc()
    return MainMenu_GetHashCode(m.instance)
}

func (m *TMainMenu) ToString() string {
    defer exceptionProc()
    return MainMenu_ToString(m.instance)
}

func (m *TMainMenu) AutoHotkeys() TMenuAutoFlag {
    defer exceptionProc()
    return MainMenu_GetAutoHotkeys(m.instance)
}

func (m *TMainMenu) SetAutoHotkeys(value TMenuAutoFlag) {
    defer exceptionProc()
    MainMenu_SetAutoHotkeys(m.instance, value)
}

func (m *TMainMenu) Images() *TImageList {
    defer exceptionProc()
    return ImageListFromInst(MainMenu_GetImages(m.instance))
}

func (m *TMainMenu) SetImages(value IComponent) {
    defer exceptionProc()
    MainMenu_SetImages(m.instance, CheckPtr(value))
}

func (m *TMainMenu) SetOnChange(fn TMenuChangeEvent) {
    MainMenu_SetOnChange(m.instance, fn)
}

func (m *TMainMenu) Handle() HMENU {
    defer exceptionProc()
    return MainMenu_GetHandle(m.instance)
}

func (m *TMainMenu) WindowHandle() HWND {
    defer exceptionProc()
    return MainMenu_GetWindowHandle(m.instance)
}

func (m *TMainMenu) SetWindowHandle(value HWND) {
    defer exceptionProc()
    MainMenu_SetWindowHandle(m.instance, value)
}

func (m *TMainMenu) Items() *TMenuItem {
    defer exceptionProc()
    return MenuItemFromInst(MainMenu_GetItems(m.instance))
}

func (m *TMainMenu) ComponentCount() int32 {
    defer exceptionProc()
    return MainMenu_GetComponentCount(m.instance)
}

func (m *TMainMenu) ComponentIndex() int32 {
    defer exceptionProc()
    return MainMenu_GetComponentIndex(m.instance)
}

func (m *TMainMenu) SetComponentIndex(value int32) {
    defer exceptionProc()
    MainMenu_SetComponentIndex(m.instance, value)
}

func (m *TMainMenu) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(MainMenu_GetOwner(m.instance))
}

func (m *TMainMenu) Name() string {
    defer exceptionProc()
    return MainMenu_GetName(m.instance)
}

func (m *TMainMenu) SetName(value string) {
    defer exceptionProc()
    MainMenu_SetName(m.instance, value)
}

func (m *TMainMenu) Tag() int {
    defer exceptionProc()
    return MainMenu_GetTag(m.instance)
}

func (m *TMainMenu) SetTag(value int) {
    defer exceptionProc()
    MainMenu_SetTag(m.instance, value)
}

func (m *TMainMenu) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(MainMenu_GetComponents(m.instance, AIndex))
}

