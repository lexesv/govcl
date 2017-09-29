
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
    . "gitee.com/ying32/govcl/vcl/types"
)

type TIcon struct {
    IObject
    instance uintptr
}

func NewIcon() *TIcon {
    i := new(TIcon)
    i.instance = Icon_Create()
    return i
}

func IconFromInst(inst uintptr) *TIcon {
    i := new(TIcon)
    i.instance = inst
    return i
}

func IconFromObj(obj IObject) *TIcon {
    i := new(TIcon)
    i.instance = CheckPtr(obj)
    return i
}

func (i *TIcon) Free() {
    if i.instance != 0 {
        Icon_Free(i.instance)
    }
}

func (i *TIcon) Instance() uintptr {
    return i.instance
}

func (i *TIcon) IsValid() bool {
    return i.instance != 0
}

func (i *TIcon) Assign(Source IObject) {
    defer exceptionProc()
    Icon_Assign(i.instance, CheckPtr(Source))
}

func (i *TIcon) HandleAllocated() bool {
    defer exceptionProc()
    return Icon_HandleAllocated(i.instance)
}

func (i *TIcon) LoadFromStream(Stream IObject) {
    defer exceptionProc()
    Icon_LoadFromStream(i.instance, CheckPtr(Stream))
}

func (i *TIcon) SaveToStream(Stream IObject) {
    defer exceptionProc()
    Icon_SaveToStream(i.instance, CheckPtr(Stream))
}

func (i *TIcon) SetSize(AWidth int32, AHeight int32) {
    defer exceptionProc()
    Icon_SetSize(i.instance, AWidth , AHeight )
}

func (i *TIcon) LoadFromResourceName(Instance uintptr, ResName string) {
    defer exceptionProc()
    Icon_LoadFromResourceName(i.instance, Instance , ResName )
}

func (i *TIcon) LoadFromResourceID(Instance uintptr, ResID int32) {
    defer exceptionProc()
    Icon_LoadFromResourceID(i.instance, Instance , ResID )
}

func (i *TIcon) Equals(Obj IObject) bool {
    defer exceptionProc()
    return Icon_Equals(i.instance, CheckPtr(Obj))
}

func (i *TIcon) LoadFromFile(Filename string) {
    defer exceptionProc()
    Icon_LoadFromFile(i.instance, Filename )
}

func (i *TIcon) SaveToFile(Filename string) {
    defer exceptionProc()
    Icon_SaveToFile(i.instance, Filename )
}

func (i *TIcon) GetNamePath() string {
    defer exceptionProc()
    return Icon_GetNamePath(i.instance)
}

func (i *TIcon) ClassName() string {
    defer exceptionProc()
    return Icon_ClassName(i.instance)
}

func (i *TIcon) GetHashCode() int32 {
    defer exceptionProc()
    return Icon_GetHashCode(i.instance)
}

func (i *TIcon) ToString() string {
    defer exceptionProc()
    return Icon_ToString(i.instance)
}

func (i *TIcon) Handle() HICON {
    defer exceptionProc()
    return Icon_GetHandle(i.instance)
}

func (i *TIcon) SetHandle(value HICON) {
    defer exceptionProc()
    Icon_SetHandle(i.instance, value)
}

func (i *TIcon) Empty() bool {
    defer exceptionProc()
    return Icon_GetEmpty(i.instance)
}

func (i *TIcon) Height() int32 {
    defer exceptionProc()
    return Icon_GetHeight(i.instance)
}

func (i *TIcon) SetHeight(value int32) {
    defer exceptionProc()
    Icon_SetHeight(i.instance, value)
}

func (i *TIcon) Modified() bool {
    defer exceptionProc()
    return Icon_GetModified(i.instance)
}

func (i *TIcon) SetModified(value bool) {
    defer exceptionProc()
    Icon_SetModified(i.instance, value)
}

func (i *TIcon) PaletteModified() bool {
    defer exceptionProc()
    return Icon_GetPaletteModified(i.instance)
}

func (i *TIcon) SetPaletteModified(value bool) {
    defer exceptionProc()
    Icon_SetPaletteModified(i.instance, value)
}

func (i *TIcon) Transparent() bool {
    defer exceptionProc()
    return Icon_GetTransparent(i.instance)
}

func (i *TIcon) SetTransparent(value bool) {
    defer exceptionProc()
    Icon_SetTransparent(i.instance, value)
}

func (i *TIcon) Width() int32 {
    defer exceptionProc()
    return Icon_GetWidth(i.instance)
}

func (i *TIcon) SetWidth(value int32) {
    defer exceptionProc()
    Icon_SetWidth(i.instance, value)
}

func (i *TIcon) SetOnChange(fn TNotifyEvent) {
    Icon_SetOnChange(i.instance, fn)
}

