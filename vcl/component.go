
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TComponent struct {
    IComponent
    instance uintptr
}

func NewComponent(owner IComponent) *TComponent {
    c := new(TComponent)
    c.instance = Component_Create(owner.Instance())
    return c
}

func ComponentFromInst(inst uintptr) *TComponent {
    c := new(TComponent)
    c.instance = inst
    return c
}

func ComponentFromObj(obj IObject) *TComponent {
    c := new(TComponent)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TComponent) Free() {
    if c.instance != 0 {
        Component_Free(c.instance)
    }
}

func (c *TComponent) Instance() uintptr {
    return c.instance
}

func (c *TComponent) IsValid() bool {
    return c.instance != 0
}

func (c *TComponent) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Component_FindComponent(c.instance, AName ))
}

func (c *TComponent) GetNamePath() string {
    defer exceptionProc()
    return Component_GetNamePath(c.instance)
}

func (c *TComponent) HasParent() bool {
    defer exceptionProc()
    return Component_HasParent(c.instance)
}

func (c *TComponent) Assign(Source IObject) {
    defer exceptionProc()
    Component_Assign(c.instance, CheckPtr(Source))
}

func (c *TComponent) ClassName() string {
    defer exceptionProc()
    return Component_ClassName(c.instance)
}

func (c *TComponent) Equals(Obj IObject) bool {
    defer exceptionProc()
    return Component_Equals(c.instance, CheckPtr(Obj))
}

func (c *TComponent) GetHashCode() int32 {
    defer exceptionProc()
    return Component_GetHashCode(c.instance)
}

func (c *TComponent) ToString() string {
    defer exceptionProc()
    return Component_ToString(c.instance)
}

func (c *TComponent) ComponentCount() int32 {
    defer exceptionProc()
    return Component_GetComponentCount(c.instance)
}

func (c *TComponent) ComponentIndex() int32 {
    defer exceptionProc()
    return Component_GetComponentIndex(c.instance)
}

func (c *TComponent) SetComponentIndex(value int32) {
    defer exceptionProc()
    Component_SetComponentIndex(c.instance, value)
}

func (c *TComponent) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Component_GetOwner(c.instance))
}

func (c *TComponent) Name() string {
    defer exceptionProc()
    return Component_GetName(c.instance)
}

func (c *TComponent) SetName(value string) {
    defer exceptionProc()
    Component_SetName(c.instance, value)
}

func (c *TComponent) Tag() int {
    defer exceptionProc()
    return Component_GetTag(c.instance)
}

func (c *TComponent) SetTag(value int) {
    defer exceptionProc()
    Component_SetTag(c.instance, value)
}

func (c *TComponent) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Component_GetComponents(c.instance, AIndex))
}

