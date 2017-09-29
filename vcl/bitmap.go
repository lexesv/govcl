
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TBitmap struct {
    IObject
    instance uintptr
}

func NewBitmap() *TBitmap {
    b := new(TBitmap)
    b.instance = Bitmap_Create()
    return b
}

func BitmapFromInst(inst uintptr) *TBitmap {
    b := new(TBitmap)
    b.instance = inst
    return b
}

func BitmapFromObj(obj IObject) *TBitmap {
    b := new(TBitmap)
    b.instance = CheckPtr(obj)
    return b
}

func (b *TBitmap) Free() {
    if b.instance != 0 {
        Bitmap_Free(b.instance)
    }
}

func (b *TBitmap) Instance() uintptr {
    return b.instance
}

func (b *TBitmap) IsValid() bool {
    return b.instance != 0
}

func (b *TBitmap) Assign(Source IObject) {
    defer exceptionProc()
    Bitmap_Assign(b.instance, CheckPtr(Source))
}

func (b *TBitmap) HandleAllocated() bool {
    defer exceptionProc()
    return Bitmap_HandleAllocated(b.instance)
}

func (b *TBitmap) LoadFromStream(Stream IObject) {
    defer exceptionProc()
    Bitmap_LoadFromStream(b.instance, CheckPtr(Stream))
}

func (b *TBitmap) SaveToStream(Stream IObject) {
    defer exceptionProc()
    Bitmap_SaveToStream(b.instance, CheckPtr(Stream))
}

func (b *TBitmap) SetSize(AWidth int32, AHeight int32) {
    defer exceptionProc()
    Bitmap_SetSize(b.instance, AWidth , AHeight )
}

func (b *TBitmap) LoadFromResourceName(Instance uintptr, ResName string) {
    defer exceptionProc()
    Bitmap_LoadFromResourceName(b.instance, Instance , ResName )
}

func (b *TBitmap) LoadFromResourceID(Instance uintptr, ResID int32) {
    defer exceptionProc()
    Bitmap_LoadFromResourceID(b.instance, Instance , ResID )
}

func (b *TBitmap) Equals(Obj IObject) bool {
    defer exceptionProc()
    return Bitmap_Equals(b.instance, CheckPtr(Obj))
}

func (b *TBitmap) LoadFromFile(Filename string) {
    defer exceptionProc()
    Bitmap_LoadFromFile(b.instance, Filename )
}

func (b *TBitmap) SaveToFile(Filename string) {
    defer exceptionProc()
    Bitmap_SaveToFile(b.instance, Filename )
}

func (b *TBitmap) GetNamePath() string {
    defer exceptionProc()
    return Bitmap_GetNamePath(b.instance)
}

func (b *TBitmap) ClassName() string {
    defer exceptionProc()
    return Bitmap_ClassName(b.instance)
}

func (b *TBitmap) GetHashCode() int32 {
    defer exceptionProc()
    return Bitmap_GetHashCode(b.instance)
}

func (b *TBitmap) ToString() string {
    defer exceptionProc()
    return Bitmap_ToString(b.instance)
}

func (b *TBitmap) Canvas() *TCanvas {
    defer exceptionProc()
    return CanvasFromInst(Bitmap_GetCanvas(b.instance))
}

func (b *TBitmap) Handle() HBITMAP {
    defer exceptionProc()
    return Bitmap_GetHandle(b.instance)
}

func (b *TBitmap) SetHandle(value HBITMAP) {
    defer exceptionProc()
    Bitmap_SetHandle(b.instance, value)
}

func (b *TBitmap) PixelFormat() TPixelFormat {
    defer exceptionProc()
    return Bitmap_GetPixelFormat(b.instance)
}

func (b *TBitmap) SetPixelFormat(value TPixelFormat) {
    defer exceptionProc()
    Bitmap_SetPixelFormat(b.instance, value)
}

func (b *TBitmap) TransparentColor() TColor {
    defer exceptionProc()
    return Bitmap_GetTransparentColor(b.instance)
}

func (b *TBitmap) SetTransparentColor(value TColor) {
    defer exceptionProc()
    Bitmap_SetTransparentColor(b.instance, value)
}

func (b *TBitmap) Empty() bool {
    defer exceptionProc()
    return Bitmap_GetEmpty(b.instance)
}

func (b *TBitmap) Height() int32 {
    defer exceptionProc()
    return Bitmap_GetHeight(b.instance)
}

func (b *TBitmap) SetHeight(value int32) {
    defer exceptionProc()
    Bitmap_SetHeight(b.instance, value)
}

func (b *TBitmap) Modified() bool {
    defer exceptionProc()
    return Bitmap_GetModified(b.instance)
}

func (b *TBitmap) SetModified(value bool) {
    defer exceptionProc()
    Bitmap_SetModified(b.instance, value)
}

func (b *TBitmap) PaletteModified() bool {
    defer exceptionProc()
    return Bitmap_GetPaletteModified(b.instance)
}

func (b *TBitmap) SetPaletteModified(value bool) {
    defer exceptionProc()
    Bitmap_SetPaletteModified(b.instance, value)
}

func (b *TBitmap) Transparent() bool {
    defer exceptionProc()
    return Bitmap_GetTransparent(b.instance)
}

func (b *TBitmap) SetTransparent(value bool) {
    defer exceptionProc()
    Bitmap_SetTransparent(b.instance, value)
}

func (b *TBitmap) Width() int32 {
    defer exceptionProc()
    return Bitmap_GetWidth(b.instance)
}

func (b *TBitmap) SetWidth(value int32) {
    defer exceptionProc()
    Bitmap_SetWidth(b.instance, value)
}

func (b *TBitmap) SetOnChange(fn TNotifyEvent) {
    Bitmap_SetOnChange(b.instance, fn)
}

func (b *TBitmap) ScanLine(Row int32) uintptr {
    defer exceptionProc()
    return Bitmap_GetScanLine(b.instance, Row)
}

