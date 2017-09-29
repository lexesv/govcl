
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TFont struct {
    IObject
    instance uintptr
}

func NewFont() *TFont {
    f := new(TFont)
    f.instance = Font_Create()
    return f
}

func FontFromInst(inst uintptr) *TFont {
    f := new(TFont)
    f.instance = inst
    return f
}

func FontFromObj(obj IObject) *TFont {
    f := new(TFont)
    f.instance = CheckPtr(obj)
    return f
}

func (f *TFont) Free() {
    if f.instance != 0 {
        Font_Free(f.instance)
    }
}

func (f *TFont) Instance() uintptr {
    return f.instance
}

func (f *TFont) IsValid() bool {
    return f.instance != 0
}

func (f *TFont) Assign(Source IObject) {
    defer exceptionProc()
    Font_Assign(f.instance, CheckPtr(Source))
}

func (f *TFont) HandleAllocated() bool {
    defer exceptionProc()
    return Font_HandleAllocated(f.instance)
}

func (f *TFont) GetNamePath() string {
    defer exceptionProc()
    return Font_GetNamePath(f.instance)
}

func (f *TFont) ClassName() string {
    defer exceptionProc()
    return Font_ClassName(f.instance)
}

func (f *TFont) Equals(Obj IObject) bool {
    defer exceptionProc()
    return Font_Equals(f.instance, CheckPtr(Obj))
}

func (f *TFont) GetHashCode() int32 {
    defer exceptionProc()
    return Font_GetHashCode(f.instance)
}

func (f *TFont) ToString() string {
    defer exceptionProc()
    return Font_ToString(f.instance)
}

func (f *TFont) Handle() HFONT {
    defer exceptionProc()
    return Font_GetHandle(f.instance)
}

func (f *TFont) SetHandle(value HFONT) {
    defer exceptionProc()
    Font_SetHandle(f.instance, value)
}

func (f *TFont) PixelsPerInch() int32 {
    defer exceptionProc()
    return Font_GetPixelsPerInch(f.instance)
}

func (f *TFont) SetPixelsPerInch(value int32) {
    defer exceptionProc()
    Font_SetPixelsPerInch(f.instance, value)
}

func (f *TFont) Charset() TFontCharset {
    defer exceptionProc()
    return Font_GetCharset(f.instance)
}

func (f *TFont) SetCharset(value TFontCharset) {
    defer exceptionProc()
    Font_SetCharset(f.instance, value)
}

func (f *TFont) Color() TColor {
    defer exceptionProc()
    return Font_GetColor(f.instance)
}

func (f *TFont) SetColor(value TColor) {
    defer exceptionProc()
    Font_SetColor(f.instance, value)
}

func (f *TFont) Height() int32 {
    defer exceptionProc()
    return Font_GetHeight(f.instance)
}

func (f *TFont) SetHeight(value int32) {
    defer exceptionProc()
    Font_SetHeight(f.instance, value)
}

func (f *TFont) Name() string {
    defer exceptionProc()
    return Font_GetName(f.instance)
}

func (f *TFont) SetName(value string) {
    defer exceptionProc()
    Font_SetName(f.instance, value)
}

func (f *TFont) Orientation() int32 {
    defer exceptionProc()
    return Font_GetOrientation(f.instance)
}

func (f *TFont) SetOrientation(value int32) {
    defer exceptionProc()
    Font_SetOrientation(f.instance, value)
}

func (f *TFont) Pitch() TFontPitch {
    defer exceptionProc()
    return Font_GetPitch(f.instance)
}

func (f *TFont) SetPitch(value TFontPitch) {
    defer exceptionProc()
    Font_SetPitch(f.instance, value)
}

func (f *TFont) Size() int32 {
    defer exceptionProc()
    return Font_GetSize(f.instance)
}

func (f *TFont) SetSize(value int32) {
    defer exceptionProc()
    Font_SetSize(f.instance, value)
}

func (f *TFont) Style() TFontStyles {
    defer exceptionProc()
    return Font_GetStyle(f.instance)
}

func (f *TFont) SetStyle(value TFontStyles) {
    defer exceptionProc()
    Font_SetStyle(f.instance, value)
}

func (f *TFont) Quality() TFontQuality {
    defer exceptionProc()
    return Font_GetQuality(f.instance)
}

func (f *TFont) SetQuality(value TFontQuality) {
    defer exceptionProc()
    Font_SetQuality(f.instance, value)
}

func (f *TFont) SetOnChange(fn TNotifyEvent) {
    Font_SetOnChange(f.instance, fn)
}

