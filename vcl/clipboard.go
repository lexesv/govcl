
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TClipboard struct {
    IObject
    instance uintptr
}

func NewClipboard() *TClipboard {
    c := new(TClipboard)
    c.instance = Clipboard_Create()
    return c
}

func ClipboardFromInst(inst uintptr) *TClipboard {
    c := new(TClipboard)
    c.instance = inst
    return c
}

func ClipboardFromObj(obj IObject) *TClipboard {
    c := new(TClipboard)
    c.instance = CheckPtr(obj)
    return c
}

func (c *TClipboard) Free() {
    if c.instance != 0 {
        Clipboard_Free(c.instance)
    }
}

func (c *TClipboard) Instance() uintptr {
    return c.instance
}

func (c *TClipboard) IsValid() bool {
    return c.instance != 0
}

func (c *TClipboard) Assign(Source IObject) {
    defer exceptionProc()
    Clipboard_Assign(c.instance, CheckPtr(Source))
}

func (c *TClipboard) Clear() {
    defer exceptionProc()
    Clipboard_Clear(c.instance)
}

func (c *TClipboard) Close() {
    defer exceptionProc()
    Clipboard_Close(c.instance)
}

func (c *TClipboard) GetAsHandle(Format uint16) uintptr {
    defer exceptionProc()
    return Clipboard_GetAsHandle(c.instance, Format )
}

func (c *TClipboard) HasFormat(Format uint16) bool {
    defer exceptionProc()
    return Clipboard_HasFormat(c.instance, Format )
}

func (c *TClipboard) Open() {
    defer exceptionProc()
    Clipboard_Open(c.instance)
}

func (c *TClipboard) SetAsHandle(Format uint16, Value uintptr) {
    defer exceptionProc()
    Clipboard_SetAsHandle(c.instance, Format , Value )
}

func (c *TClipboard) GetTextBuf(Buffer string, BufSize int32) int32 {
    defer exceptionProc()
    return Clipboard_GetTextBuf(c.instance, Buffer , BufSize )
}

func (c *TClipboard) GetNamePath() string {
    defer exceptionProc()
    return Clipboard_GetNamePath(c.instance)
}

func (c *TClipboard) ClassName() string {
    defer exceptionProc()
    return Clipboard_ClassName(c.instance)
}

func (c *TClipboard) Equals(Obj IObject) bool {
    defer exceptionProc()
    return Clipboard_Equals(c.instance, CheckPtr(Obj))
}

func (c *TClipboard) GetHashCode() int32 {
    defer exceptionProc()
    return Clipboard_GetHashCode(c.instance)
}

func (c *TClipboard) ToString() string {
    defer exceptionProc()
    return Clipboard_ToString(c.instance)
}

func (c *TClipboard) AsText() string {
    defer exceptionProc()
    return Clipboard_GetAsText(c.instance)
}

func (c *TClipboard) SetAsText(value string) {
    defer exceptionProc()
    Clipboard_SetAsText(c.instance, value)
}

func (c *TClipboard) FormatCount() int32 {
    defer exceptionProc()
    return Clipboard_GetFormatCount(c.instance)
}

func (c *TClipboard) Formats(Index int32) uint16 {
    defer exceptionProc()
    return Clipboard_GetFormats(c.instance, Index)
}

