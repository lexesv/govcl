
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TJPEGImage struct {
    IObject
    instance uintptr
}

func NewJPEGImage() *TJPEGImage {
    j := new(TJPEGImage)
    j.instance = JPEGImage_Create()
    return j
}

func JPEGImageFromInst(inst uintptr) *TJPEGImage {
    j := new(TJPEGImage)
    j.instance = inst
    return j
}

func JPEGImageFromObj(obj IObject) *TJPEGImage {
    j := new(TJPEGImage)
    j.instance = CheckPtr(obj)
    return j
}

func (j *TJPEGImage) Free() {
    if j.instance != 0 {
        JPEGImage_Free(j.instance)
    }
}

func (j *TJPEGImage) Instance() uintptr {
    return j.instance
}

func (j *TJPEGImage) IsValid() bool {
    return j.instance != 0
}

func (j *TJPEGImage) Assign(Source IObject) {
    defer exceptionProc()
    JPEGImage_Assign(j.instance, CheckPtr(Source))
}

func (j *TJPEGImage) LoadFromStream(Stream IObject) {
    defer exceptionProc()
    JPEGImage_LoadFromStream(j.instance, CheckPtr(Stream))
}

func (j *TJPEGImage) SaveToStream(Stream IObject) {
    defer exceptionProc()
    JPEGImage_SaveToStream(j.instance, CheckPtr(Stream))
}

func (j *TJPEGImage) Equals(Obj IObject) bool {
    defer exceptionProc()
    return JPEGImage_Equals(j.instance, CheckPtr(Obj))
}

func (j *TJPEGImage) LoadFromFile(Filename string) {
    defer exceptionProc()
    JPEGImage_LoadFromFile(j.instance, Filename )
}

func (j *TJPEGImage) SaveToFile(Filename string) {
    defer exceptionProc()
    JPEGImage_SaveToFile(j.instance, Filename )
}

func (j *TJPEGImage) SetSize(AWidth int32, AHeight int32) {
    defer exceptionProc()
    JPEGImage_SetSize(j.instance, AWidth , AHeight )
}

func (j *TJPEGImage) GetNamePath() string {
    defer exceptionProc()
    return JPEGImage_GetNamePath(j.instance)
}

func (j *TJPEGImage) ClassName() string {
    defer exceptionProc()
    return JPEGImage_ClassName(j.instance)
}

func (j *TJPEGImage) GetHashCode() int32 {
    defer exceptionProc()
    return JPEGImage_GetHashCode(j.instance)
}

func (j *TJPEGImage) ToString() string {
    defer exceptionProc()
    return JPEGImage_ToString(j.instance)
}

func (j *TJPEGImage) PixelFormat() TJPEGPixelFormat {
    defer exceptionProc()
    return JPEGImage_GetPixelFormat(j.instance)
}

func (j *TJPEGImage) SetPixelFormat(value TJPEGPixelFormat) {
    defer exceptionProc()
    JPEGImage_SetPixelFormat(j.instance, value)
}

func (j *TJPEGImage) ProgressiveDisplay() bool {
    defer exceptionProc()
    return JPEGImage_GetProgressiveDisplay(j.instance)
}

func (j *TJPEGImage) SetProgressiveDisplay(value bool) {
    defer exceptionProc()
    JPEGImage_SetProgressiveDisplay(j.instance, value)
}

func (j *TJPEGImage) Performance() TJPEGPerformance {
    defer exceptionProc()
    return JPEGImage_GetPerformance(j.instance)
}

func (j *TJPEGImage) SetPerformance(value TJPEGPerformance) {
    defer exceptionProc()
    JPEGImage_SetPerformance(j.instance, value)
}

func (j *TJPEGImage) Scale() TJPEGScale {
    defer exceptionProc()
    return JPEGImage_GetScale(j.instance)
}

func (j *TJPEGImage) SetScale(value TJPEGScale) {
    defer exceptionProc()
    JPEGImage_SetScale(j.instance, value)
}

func (j *TJPEGImage) Smoothing() bool {
    defer exceptionProc()
    return JPEGImage_GetSmoothing(j.instance)
}

func (j *TJPEGImage) SetSmoothing(value bool) {
    defer exceptionProc()
    JPEGImage_SetSmoothing(j.instance, value)
}

func (j *TJPEGImage) Canvas() *TCanvas {
    defer exceptionProc()
    return CanvasFromInst(JPEGImage_GetCanvas(j.instance))
}

func (j *TJPEGImage) Empty() bool {
    defer exceptionProc()
    return JPEGImage_GetEmpty(j.instance)
}

func (j *TJPEGImage) Height() int32 {
    defer exceptionProc()
    return JPEGImage_GetHeight(j.instance)
}

func (j *TJPEGImage) SetHeight(value int32) {
    defer exceptionProc()
    JPEGImage_SetHeight(j.instance, value)
}

func (j *TJPEGImage) Modified() bool {
    defer exceptionProc()
    return JPEGImage_GetModified(j.instance)
}

func (j *TJPEGImage) SetModified(value bool) {
    defer exceptionProc()
    JPEGImage_SetModified(j.instance, value)
}

func (j *TJPEGImage) PaletteModified() bool {
    defer exceptionProc()
    return JPEGImage_GetPaletteModified(j.instance)
}

func (j *TJPEGImage) SetPaletteModified(value bool) {
    defer exceptionProc()
    JPEGImage_SetPaletteModified(j.instance, value)
}

func (j *TJPEGImage) Transparent() bool {
    defer exceptionProc()
    return JPEGImage_GetTransparent(j.instance)
}

func (j *TJPEGImage) SetTransparent(value bool) {
    defer exceptionProc()
    JPEGImage_SetTransparent(j.instance, value)
}

func (j *TJPEGImage) Width() int32 {
    defer exceptionProc()
    return JPEGImage_GetWidth(j.instance)
}

func (j *TJPEGImage) SetWidth(value int32) {
    defer exceptionProc()
    JPEGImage_SetWidth(j.instance, value)
}

func (j *TJPEGImage) SetOnChange(fn TNotifyEvent) {
    JPEGImage_SetOnChange(j.instance, fn)
}

