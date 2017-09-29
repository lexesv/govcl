
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TListGroups struct {
    IObject
    instance uintptr
}

func NewListGroups() *TListGroups {
    l := new(TListGroups)
    l.instance = ListGroups_Create()
    return l
}

func ListGroupsFromInst(inst uintptr) *TListGroups {
    l := new(TListGroups)
    l.instance = inst
    return l
}

func ListGroupsFromObj(obj IObject) *TListGroups {
    l := new(TListGroups)
    l.instance = CheckPtr(obj)
    return l
}

func (l *TListGroups) Free() {
    if l.instance != 0 {
        ListGroups_Free(l.instance)
    }
}

func (l *TListGroups) Instance() uintptr {
    return l.instance
}

func (l *TListGroups) IsValid() bool {
    return l.instance != 0
}

func (l *TListGroups) Add() *TListGroup {
    defer exceptionProc()
    return ListGroupFromInst(ListGroups_Add(l.instance))
}

func (l *TListGroups) Owner() *TControl {
    defer exceptionProc()
    return ControlFromInst(ListGroups_Owner(l.instance))
}

func (l *TListGroups) Assign(Source IObject) {
    defer exceptionProc()
    ListGroups_Assign(l.instance, CheckPtr(Source))
}

func (l *TListGroups) BeginUpdate() {
    defer exceptionProc()
    ListGroups_BeginUpdate(l.instance)
}

func (l *TListGroups) Clear() {
    defer exceptionProc()
    ListGroups_Clear(l.instance)
}

func (l *TListGroups) Delete(Index int32) {
    defer exceptionProc()
    ListGroups_Delete(l.instance, Index )
}

func (l *TListGroups) EndUpdate() {
    defer exceptionProc()
    ListGroups_EndUpdate(l.instance)
}

func (l *TListGroups) GetNamePath() string {
    defer exceptionProc()
    return ListGroups_GetNamePath(l.instance)
}

func (l *TListGroups) Insert(Index int32) *TCollectionItem {
    defer exceptionProc()
    return CollectionItemFromInst(ListGroups_Insert(l.instance, Index ))
}

func (l *TListGroups) ClassName() string {
    defer exceptionProc()
    return ListGroups_ClassName(l.instance)
}

func (l *TListGroups) Equals(Obj IObject) bool {
    defer exceptionProc()
    return ListGroups_Equals(l.instance, CheckPtr(Obj))
}

func (l *TListGroups) GetHashCode() int32 {
    defer exceptionProc()
    return ListGroups_GetHashCode(l.instance)
}

func (l *TListGroups) ToString() string {
    defer exceptionProc()
    return ListGroups_ToString(l.instance)
}

func (l *TListGroups) Items(Index int32) *TListGroup {
    defer exceptionProc()
    return ListGroupFromInst(ListGroups_GetItems(l.instance, Index))
}

func (l *TListGroups) SetItems(Index int32, value *TListGroup) {
    defer exceptionProc()
    ListGroups_SetItems(l.instance, Index, CheckPtr(value))
}

