
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

type TGIFFrame struct {
    IObject
    instance uintptr
}

func NewGIFFrame() *TGIFFrame {
    g := new(TGIFFrame)
    g.instance = GIFFrame_Create()
    return g
}

func GIFFrameFromInst(inst uintptr) *TGIFFrame {
    g := new(TGIFFrame)
    g.instance = inst
    return g
}

func GIFFrameFromObj(obj IObject) *TGIFFrame {
    g := new(TGIFFrame)
    g.instance = CheckPtr(obj)
    return g
}

func (g *TGIFFrame) Free() {
    if g.instance != 0 {
        GIFFrame_Free(g.instance)
    }
}

func (g *TGIFFrame) Instance() uintptr {
    return g.instance
}

func (g *TGIFFrame) IsValid() bool {
    return g.instance != 0
}

func (g *TGIFFrame) Clear() {
    defer exceptionProc()
    GIFFrame_Clear(g.instance)
}

func (g *TGIFFrame) SaveToStream(Stream IObject) {
    defer exceptionProc()
    GIFFrame_SaveToStream(g.instance, CheckPtr(Stream))
}

func (g *TGIFFrame) LoadFromStream(Stream IObject) {
    defer exceptionProc()
    GIFFrame_LoadFromStream(g.instance, CheckPtr(Stream))
}

func (g *TGIFFrame) Assign(Source IObject) {
    defer exceptionProc()
    GIFFrame_Assign(g.instance, CheckPtr(Source))
}

func (g *TGIFFrame) SaveToFile(Filename string) {
    defer exceptionProc()
    GIFFrame_SaveToFile(g.instance, Filename )
}

func (g *TGIFFrame) LoadFromFile(Filename string) {
    defer exceptionProc()
    GIFFrame_LoadFromFile(g.instance, Filename )
}

func (g *TGIFFrame) GetNamePath() string {
    defer exceptionProc()
    return GIFFrame_GetNamePath(g.instance)
}

func (g *TGIFFrame) ClassName() string {
    defer exceptionProc()
    return GIFFrame_ClassName(g.instance)
}

func (g *TGIFFrame) Equals(Obj IObject) bool {
    defer exceptionProc()
    return GIFFrame_Equals(g.instance, CheckPtr(Obj))
}

func (g *TGIFFrame) GetHashCode() int32 {
    defer exceptionProc()
    return GIFFrame_GetHashCode(g.instance)
}

func (g *TGIFFrame) ToString() string {
    defer exceptionProc()
    return GIFFrame_ToString(g.instance)
}

func (g *TGIFFrame) HasBitmap() bool {
    defer exceptionProc()
    return GIFFrame_GetHasBitmap(g.instance)
}

func (g *TGIFFrame) SetHasBitmap(value bool) {
    defer exceptionProc()
    GIFFrame_SetHasBitmap(g.instance, value)
}

func (g *TGIFFrame) Left() uint16 {
    defer exceptionProc()
    return GIFFrame_GetLeft(g.instance)
}

func (g *TGIFFrame) SetLeft(value uint16) {
    defer exceptionProc()
    GIFFrame_SetLeft(g.instance, value)
}

func (g *TGIFFrame) Top() uint16 {
    defer exceptionProc()
    return GIFFrame_GetTop(g.instance)
}

func (g *TGIFFrame) SetTop(value uint16) {
    defer exceptionProc()
    GIFFrame_SetTop(g.instance, value)
}

func (g *TGIFFrame) Width() uint16 {
    defer exceptionProc()
    return GIFFrame_GetWidth(g.instance)
}

func (g *TGIFFrame) SetWidth(value uint16) {
    defer exceptionProc()
    GIFFrame_SetWidth(g.instance, value)
}

func (g *TGIFFrame) Height() uint16 {
    defer exceptionProc()
    return GIFFrame_GetHeight(g.instance)
}

func (g *TGIFFrame) SetHeight(value uint16) {
    defer exceptionProc()
    GIFFrame_SetHeight(g.instance, value)
}

func (g *TGIFFrame) BoundsRect() TRect {
    defer exceptionProc()
    return GIFFrame_GetBoundsRect(g.instance)
}

func (g *TGIFFrame) SetBoundsRect(value TRect) {
    defer exceptionProc()
    GIFFrame_SetBoundsRect(g.instance, value)
}

func (g *TGIFFrame) ClientRect() TRect {
    defer exceptionProc()
    return GIFFrame_GetClientRect(g.instance)
}

func (g *TGIFFrame) Data() uintptr {
    defer exceptionProc()
    return GIFFrame_GetData(g.instance)
}

func (g *TGIFFrame) DataSize() int32 {
    defer exceptionProc()
    return GIFFrame_GetDataSize(g.instance)
}

func (g *TGIFFrame) Version() TGIFVersion {
    defer exceptionProc()
    return GIFFrame_GetVersion(g.instance)
}

func (g *TGIFFrame) BitsPerPixel() int32 {
    defer exceptionProc()
    return GIFFrame_GetBitsPerPixel(g.instance)
}

func (g *TGIFFrame) Bitmap() *TBitmap {
    defer exceptionProc()
    return BitmapFromInst(GIFFrame_GetBitmap(g.instance))
}

func (g *TGIFFrame) SetBitmap(value *TBitmap) {
    defer exceptionProc()
    GIFFrame_SetBitmap(g.instance, CheckPtr(value))
}

func (g *TGIFFrame) Empty() bool {
    defer exceptionProc()
    return GIFFrame_GetEmpty(g.instance)
}

func (g *TGIFFrame) Transparent() bool {
    defer exceptionProc()
    return GIFFrame_GetTransparent(g.instance)
}

