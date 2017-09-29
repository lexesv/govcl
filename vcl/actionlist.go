
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TActionList struct {
    IComponent
    instance uintptr
}

func NewActionList(owner IComponent) *TActionList {
    a := new(TActionList)
    a.instance = ActionList_Create(owner.Instance())
    return a
}

func ActionListFromInst(inst uintptr) *TActionList {
    a := new(TActionList)
    a.instance = inst
    return a
}

func ActionListFromObj(obj IObject) *TActionList {
    a := new(TActionList)
    a.instance = CheckPtr(obj)
    return a
}

func (a *TActionList) Free() {
    if a.instance != 0 {
        ActionList_Free(a.instance)
    }
}

func (a *TActionList) Instance() uintptr {
    return a.instance
}

func (a *TActionList) IsValid() bool {
    return a.instance != 0
}

func (a *TActionList) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ActionList_FindComponent(a.instance, AName ))
}

func (a *TActionList) GetNamePath() string {
    defer exceptionProc()
    return ActionList_GetNamePath(a.instance)
}

func (a *TActionList) HasParent() bool {
    defer exceptionProc()
    return ActionList_HasParent(a.instance)
}

func (a *TActionList) Assign(Source IObject) {
    defer exceptionProc()
    ActionList_Assign(a.instance, CheckPtr(Source))
}

func (a *TActionList) ClassName() string {
    defer exceptionProc()
    return ActionList_ClassName(a.instance)
}

func (a *TActionList) Equals(Obj IObject) bool {
    defer exceptionProc()
    return ActionList_Equals(a.instance, CheckPtr(Obj))
}

func (a *TActionList) GetHashCode() int32 {
    defer exceptionProc()
    return ActionList_GetHashCode(a.instance)
}

func (a *TActionList) ToString() string {
    defer exceptionProc()
    return ActionList_ToString(a.instance)
}

func (a *TActionList) Images() *TImageList {
    defer exceptionProc()
    return ImageListFromInst(ActionList_GetImages(a.instance))
}

func (a *TActionList) SetImages(value IComponent) {
    defer exceptionProc()
    ActionList_SetImages(a.instance, CheckPtr(value))
}

func (a *TActionList) State() TActionListState {
    defer exceptionProc()
    return ActionList_GetState(a.instance)
}

func (a *TActionList) SetState(value TActionListState) {
    defer exceptionProc()
    ActionList_SetState(a.instance, value)
}

func (a *TActionList) SetOnChange(fn TNotifyEvent) {
    ActionList_SetOnChange(a.instance, fn)
}

func (a *TActionList) ComponentCount() int32 {
    defer exceptionProc()
    return ActionList_GetComponentCount(a.instance)
}

func (a *TActionList) ComponentIndex() int32 {
    defer exceptionProc()
    return ActionList_GetComponentIndex(a.instance)
}

func (a *TActionList) SetComponentIndex(value int32) {
    defer exceptionProc()
    ActionList_SetComponentIndex(a.instance, value)
}

func (a *TActionList) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ActionList_GetOwner(a.instance))
}

func (a *TActionList) Name() string {
    defer exceptionProc()
    return ActionList_GetName(a.instance)
}

func (a *TActionList) SetName(value string) {
    defer exceptionProc()
    ActionList_SetName(a.instance, value)
}

func (a *TActionList) Tag() int {
    defer exceptionProc()
    return ActionList_GetTag(a.instance)
}

func (a *TActionList) SetTag(value int) {
    defer exceptionProc()
    ActionList_SetTag(a.instance, value)
}

func (a *TActionList) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ActionList_GetComponents(a.instance, AIndex))
}

