
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

type TPngImage struct {
    IObject
    instance uintptr
}

func NewPngImage() *TPngImage {
    p := new(TPngImage)
    p.instance = PngImage_Create()
    return p
}

func PngImageFromInst(inst uintptr) *TPngImage {
    p := new(TPngImage)
    p.instance = inst
    return p
}

func PngImageFromObj(obj IObject) *TPngImage {
    p := new(TPngImage)
    p.instance = CheckPtr(obj)
    return p
}

func (p *TPngImage) Free() {
    if p.instance != 0 {
        PngImage_Free(p.instance)
    }
}

func (p *TPngImage) Instance() uintptr {
    return p.instance
}

func (p *TPngImage) IsValid() bool {
    return p.instance != 0
}

func (p *TPngImage) Assign(Source IObject) {
    defer exceptionProc()
    PngImage_Assign(p.instance, CheckPtr(Source))
}

func (p *TPngImage) LoadFromStream(Stream IObject) {
    defer exceptionProc()
    PngImage_LoadFromStream(p.instance, CheckPtr(Stream))
}

func (p *TPngImage) SaveToStream(Stream IObject) {
    defer exceptionProc()
    PngImage_SaveToStream(p.instance, CheckPtr(Stream))
}

func (p *TPngImage) LoadFromResourceName(Instance uintptr, Name string) {
    defer exceptionProc()
    PngImage_LoadFromResourceName(p.instance, Instance , Name )
}

func (p *TPngImage) LoadFromResourceID(Instance uintptr, ResID int32) {
    defer exceptionProc()
    PngImage_LoadFromResourceID(p.instance, Instance , ResID )
}

func (p *TPngImage) Equals(Obj IObject) bool {
    defer exceptionProc()
    return PngImage_Equals(p.instance, CheckPtr(Obj))
}

func (p *TPngImage) LoadFromFile(Filename string) {
    defer exceptionProc()
    PngImage_LoadFromFile(p.instance, Filename )
}

func (p *TPngImage) SaveToFile(Filename string) {
    defer exceptionProc()
    PngImage_SaveToFile(p.instance, Filename )
}

func (p *TPngImage) SetSize(AWidth int32, AHeight int32) {
    defer exceptionProc()
    PngImage_SetSize(p.instance, AWidth , AHeight )
}

func (p *TPngImage) GetNamePath() string {
    defer exceptionProc()
    return PngImage_GetNamePath(p.instance)
}

func (p *TPngImage) ClassName() string {
    defer exceptionProc()
    return PngImage_ClassName(p.instance)
}

func (p *TPngImage) GetHashCode() int32 {
    defer exceptionProc()
    return PngImage_GetHashCode(p.instance)
}

func (p *TPngImage) ToString() string {
    defer exceptionProc()
    return PngImage_ToString(p.instance)
}

func (p *TPngImage) TransparentColor() TColor {
    defer exceptionProc()
    return PngImage_GetTransparentColor(p.instance)
}

func (p *TPngImage) SetTransparentColor(value TColor) {
    defer exceptionProc()
    PngImage_SetTransparentColor(p.instance, value)
}

func (p *TPngImage) Canvas() *TCanvas {
    defer exceptionProc()
    return CanvasFromInst(PngImage_GetCanvas(p.instance))
}

func (p *TPngImage) Width() int32 {
    defer exceptionProc()
    return PngImage_GetWidth(p.instance)
}

func (p *TPngImage) Height() int32 {
    defer exceptionProc()
    return PngImage_GetHeight(p.instance)
}

func (p *TPngImage) MaxIdatSize() int32 {
    defer exceptionProc()
    return PngImage_GetMaxIdatSize(p.instance)
}

func (p *TPngImage) SetMaxIdatSize(value int32) {
    defer exceptionProc()
    PngImage_SetMaxIdatSize(p.instance, value)
}

func (p *TPngImage) Empty() bool {
    defer exceptionProc()
    return PngImage_GetEmpty(p.instance)
}

func (p *TPngImage) CompressionLevel() TCompressionLevel {
    defer exceptionProc()
    return PngImage_GetCompressionLevel(p.instance)
}

func (p *TPngImage) SetCompressionLevel(value TCompressionLevel) {
    defer exceptionProc()
    PngImage_SetCompressionLevel(p.instance, value)
}

func (p *TPngImage) Version() string {
    defer exceptionProc()
    return PngImage_GetVersion(p.instance)
}

func (p *TPngImage) Modified() bool {
    defer exceptionProc()
    return PngImage_GetModified(p.instance)
}

func (p *TPngImage) SetModified(value bool) {
    defer exceptionProc()
    PngImage_SetModified(p.instance, value)
}

func (p *TPngImage) PaletteModified() bool {
    defer exceptionProc()
    return PngImage_GetPaletteModified(p.instance)
}

func (p *TPngImage) SetPaletteModified(value bool) {
    defer exceptionProc()
    PngImage_SetPaletteModified(p.instance, value)
}

func (p *TPngImage) Transparent() bool {
    defer exceptionProc()
    return PngImage_GetTransparent(p.instance)
}

func (p *TPngImage) SetTransparent(value bool) {
    defer exceptionProc()
    PngImage_SetTransparent(p.instance, value)
}

func (p *TPngImage) SetOnChange(fn TNotifyEvent) {
    PngImage_SetOnChange(p.instance, fn)
}

