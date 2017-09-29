
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TListColumns struct {
    IObject
    instance uintptr
}

func NewListColumns() *TListColumns {
    l := new(TListColumns)
    l.instance = ListColumns_Create()
    return l
}

func ListColumnsFromInst(inst uintptr) *TListColumns {
    l := new(TListColumns)
    l.instance = inst
    return l
}

func ListColumnsFromObj(obj IObject) *TListColumns {
    l := new(TListColumns)
    l.instance = CheckPtr(obj)
    return l
}

func (l *TListColumns) Free() {
    if l.instance != 0 {
        ListColumns_Free(l.instance)
    }
}

func (l *TListColumns) Instance() uintptr {
    return l.instance
}

func (l *TListColumns) IsValid() bool {
    return l.instance != 0
}

func (l *TListColumns) Add() *TListColumn {
    defer exceptionProc()
    return ListColumnFromInst(ListColumns_Add(l.instance))
}

func (l *TListColumns) Owner() *TControl {
    defer exceptionProc()
    return ControlFromInst(ListColumns_Owner(l.instance))
}

func (l *TListColumns) Assign(Source IObject) {
    defer exceptionProc()
    ListColumns_Assign(l.instance, CheckPtr(Source))
}

func (l *TListColumns) BeginUpdate() {
    defer exceptionProc()
    ListColumns_BeginUpdate(l.instance)
}

func (l *TListColumns) Clear() {
    defer exceptionProc()
    ListColumns_Clear(l.instance)
}

func (l *TListColumns) Delete(Index int32) {
    defer exceptionProc()
    ListColumns_Delete(l.instance, Index )
}

func (l *TListColumns) EndUpdate() {
    defer exceptionProc()
    ListColumns_EndUpdate(l.instance)
}

func (l *TListColumns) GetNamePath() string {
    defer exceptionProc()
    return ListColumns_GetNamePath(l.instance)
}

func (l *TListColumns) Insert(Index int32) *TCollectionItem {
    defer exceptionProc()
    return CollectionItemFromInst(ListColumns_Insert(l.instance, Index ))
}

func (l *TListColumns) ClassName() string {
    defer exceptionProc()
    return ListColumns_ClassName(l.instance)
}

func (l *TListColumns) Equals(Obj IObject) bool {
    defer exceptionProc()
    return ListColumns_Equals(l.instance, CheckPtr(Obj))
}

func (l *TListColumns) GetHashCode() int32 {
    defer exceptionProc()
    return ListColumns_GetHashCode(l.instance)
}

func (l *TListColumns) ToString() string {
    defer exceptionProc()
    return ListColumns_ToString(l.instance)
}

func (l *TListColumns) Items(Index int32) *TListColumn {
    defer exceptionProc()
    return ListColumnFromInst(ListColumns_GetItems(l.instance, Index))
}

func (l *TListColumns) SetItems(Index int32, value *TListColumn) {
    defer exceptionProc()
    ListColumns_SetItems(l.instance, Index, CheckPtr(value))
}

