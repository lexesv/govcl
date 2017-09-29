
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TList struct {
    IObject
    instance uintptr
}

func NewList() *TList {
    l := new(TList)
    l.instance = List_Create()
    return l
}

func ListFromInst(inst uintptr) *TList {
    l := new(TList)
    l.instance = inst
    return l
}

func ListFromObj(obj IObject) *TList {
    l := new(TList)
    l.instance = CheckPtr(obj)
    return l
}

func (l *TList) Free() {
    if l.instance != 0 {
        List_Free(l.instance)
    }
}

func (l *TList) Instance() uintptr {
    return l.instance
}

func (l *TList) IsValid() bool {
    return l.instance != 0
}

func (l *TList) Add(Item uintptr) int32 {
    defer exceptionProc()
    return List_Add(l.instance, Item )
}

func (l *TList) Clear() {
    defer exceptionProc()
    List_Clear(l.instance)
}

func (l *TList) Delete(Index int32) {
    defer exceptionProc()
    List_Delete(l.instance, Index )
}

func (l *TList) IndexOf(Item uintptr) int32 {
    defer exceptionProc()
    return List_IndexOf(l.instance, Item )
}

func (l *TList) Insert(Index int32, Item uintptr) {
    defer exceptionProc()
    List_Insert(l.instance, Index , Item )
}

func (l *TList) Move(CurIndex int32, NewIndex int32) {
    defer exceptionProc()
    List_Move(l.instance, CurIndex , NewIndex )
}

func (l *TList) ClassName() string {
    defer exceptionProc()
    return List_ClassName(l.instance)
}

func (l *TList) Equals(Obj IObject) bool {
    defer exceptionProc()
    return List_Equals(l.instance, CheckPtr(Obj))
}

func (l *TList) GetHashCode() int32 {
    defer exceptionProc()
    return List_GetHashCode(l.instance)
}

func (l *TList) ToString() string {
    defer exceptionProc()
    return List_ToString(l.instance)
}

func (l *TList) List() uintptr {
    defer exceptionProc()
    return List_GetList(l.instance)
}

func (l *TList) Items(Index int32) uintptr {
    defer exceptionProc()
    return List_GetItems(l.instance, Index)
}

func (l *TList) SetItems(Index int32, value uintptr) {
    defer exceptionProc()
    List_SetItems(l.instance, Index, value)
}

