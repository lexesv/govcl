
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TListColumn struct {
    IObject
    instance uintptr
}

func NewListColumn() *TListColumn {
    l := new(TListColumn)
    l.instance = ListColumn_Create()
    return l
}

func ListColumnFromInst(inst uintptr) *TListColumn {
    l := new(TListColumn)
    l.instance = inst
    return l
}

func ListColumnFromObj(obj IObject) *TListColumn {
    l := new(TListColumn)
    l.instance = CheckPtr(obj)
    return l
}

func (l *TListColumn) Free() {
    if l.instance != 0 {
        ListColumn_Free(l.instance)
    }
}

func (l *TListColumn) Instance() uintptr {
    return l.instance
}

func (l *TListColumn) IsValid() bool {
    return l.instance != 0
}

func (l *TListColumn) Assign(Source IObject) {
    defer exceptionProc()
    ListColumn_Assign(l.instance, CheckPtr(Source))
}

func (l *TListColumn) GetNamePath() string {
    defer exceptionProc()
    return ListColumn_GetNamePath(l.instance)
}

func (l *TListColumn) ClassName() string {
    defer exceptionProc()
    return ListColumn_ClassName(l.instance)
}

func (l *TListColumn) Equals(Obj IObject) bool {
    defer exceptionProc()
    return ListColumn_Equals(l.instance, CheckPtr(Obj))
}

func (l *TListColumn) GetHashCode() int32 {
    defer exceptionProc()
    return ListColumn_GetHashCode(l.instance)
}

func (l *TListColumn) ToString() string {
    defer exceptionProc()
    return ListColumn_ToString(l.instance)
}

func (l *TListColumn) Alignment() TAlignment {
    defer exceptionProc()
    return ListColumn_GetAlignment(l.instance)
}

func (l *TListColumn) SetAlignment(value TAlignment) {
    defer exceptionProc()
    ListColumn_SetAlignment(l.instance, value)
}

func (l *TListColumn) AutoSize() bool {
    defer exceptionProc()
    return ListColumn_GetAutoSize(l.instance)
}

func (l *TListColumn) SetAutoSize(value bool) {
    defer exceptionProc()
    ListColumn_SetAutoSize(l.instance, value)
}

func (l *TListColumn) Caption() string {
    defer exceptionProc()
    return ListColumn_GetCaption(l.instance)
}

func (l *TListColumn) SetCaption(value string) {
    defer exceptionProc()
    ListColumn_SetCaption(l.instance, value)
}

func (l *TListColumn) ImageIndex() int32 {
    defer exceptionProc()
    return ListColumn_GetImageIndex(l.instance)
}

func (l *TListColumn) SetImageIndex(value int32) {
    defer exceptionProc()
    ListColumn_SetImageIndex(l.instance, value)
}

func (l *TListColumn) Tag() int32 {
    defer exceptionProc()
    return ListColumn_GetTag(l.instance)
}

func (l *TListColumn) SetTag(value int32) {
    defer exceptionProc()
    ListColumn_SetTag(l.instance, value)
}

func (l *TListColumn) Width() int32 {
    defer exceptionProc()
    return ListColumn_GetWidth(l.instance)
}

func (l *TListColumn) SetWidth(value int32) {
    defer exceptionProc()
    ListColumn_SetWidth(l.instance, value)
}

func (l *TListColumn) Index() int32 {
    defer exceptionProc()
    return ListColumn_GetIndex(l.instance)
}

func (l *TListColumn) SetIndex(value int32) {
    defer exceptionProc()
    ListColumn_SetIndex(l.instance, value)
}

