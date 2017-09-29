
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TImageList struct {
    IComponent
    instance uintptr
}

func NewImageList(owner IComponent) *TImageList {
    i := new(TImageList)
    i.instance = ImageList_Create(owner.Instance())
    return i
}

func ImageListFromInst(inst uintptr) *TImageList {
    i := new(TImageList)
    i.instance = inst
    return i
}

func ImageListFromObj(obj IObject) *TImageList {
    i := new(TImageList)
    i.instance = CheckPtr(obj)
    return i
}

func (i *TImageList) Free() {
    if i.instance != 0 {
        ImageList_Free(i.instance)
    }
}

func (i *TImageList) Instance() uintptr {
    return i.instance
}

func (i *TImageList) IsValid() bool {
    return i.instance != 0
}

func (i *TImageList) Assign(Source IObject) {
    defer exceptionProc()
    ImageList_Assign(i.instance, CheckPtr(Source))
}

func (i *TImageList) Add(Image *TBitmap, Mask *TBitmap) int32 {
    defer exceptionProc()
    return ImageList_Add(i.instance, CheckPtr(Image), CheckPtr(Mask))
}

func (i *TImageList) Clear() {
    defer exceptionProc()
    ImageList_Clear(i.instance)
}

func (i *TImageList) Delete(Index int32) {
    defer exceptionProc()
    ImageList_Delete(i.instance, Index )
}

func (i *TImageList) HandleAllocated() bool {
    defer exceptionProc()
    return ImageList_HandleAllocated(i.instance)
}

func (i *TImageList) Insert(Index int32, Image *TBitmap, Mask *TBitmap) {
    defer exceptionProc()
    ImageList_Insert(i.instance, Index , CheckPtr(Image), CheckPtr(Mask))
}

func (i *TImageList) Move(CurIndex int32, NewIndex int32) {
    defer exceptionProc()
    ImageList_Move(i.instance, CurIndex , NewIndex )
}

func (i *TImageList) SetSize(AWidth int32, AHeight int32) {
    defer exceptionProc()
    ImageList_SetSize(i.instance, AWidth , AHeight )
}

func (i *TImageList) BeginUpdate() {
    defer exceptionProc()
    ImageList_BeginUpdate(i.instance)
}

func (i *TImageList) EndUpdate() {
    defer exceptionProc()
    ImageList_EndUpdate(i.instance)
}

func (i *TImageList) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ImageList_FindComponent(i.instance, AName ))
}

func (i *TImageList) GetNamePath() string {
    defer exceptionProc()
    return ImageList_GetNamePath(i.instance)
}

func (i *TImageList) HasParent() bool {
    defer exceptionProc()
    return ImageList_HasParent(i.instance)
}

func (i *TImageList) ClassName() string {
    defer exceptionProc()
    return ImageList_ClassName(i.instance)
}

func (i *TImageList) Equals(Obj IObject) bool {
    defer exceptionProc()
    return ImageList_Equals(i.instance, CheckPtr(Obj))
}

func (i *TImageList) GetHashCode() int32 {
    defer exceptionProc()
    return ImageList_GetHashCode(i.instance)
}

func (i *TImageList) ToString() string {
    defer exceptionProc()
    return ImageList_ToString(i.instance)
}

func (i *TImageList) BlendColor() TColor {
    defer exceptionProc()
    return ImageList_GetBlendColor(i.instance)
}

func (i *TImageList) SetBlendColor(value TColor) {
    defer exceptionProc()
    ImageList_SetBlendColor(i.instance, value)
}

func (i *TImageList) BkColor() TColor {
    defer exceptionProc()
    return ImageList_GetBkColor(i.instance)
}

func (i *TImageList) SetBkColor(value TColor) {
    defer exceptionProc()
    ImageList_SetBkColor(i.instance, value)
}

func (i *TImageList) AllocBy() int32 {
    defer exceptionProc()
    return ImageList_GetAllocBy(i.instance)
}

func (i *TImageList) SetAllocBy(value int32) {
    defer exceptionProc()
    ImageList_SetAllocBy(i.instance, value)
}

func (i *TImageList) ColorDepth() TColorDepth {
    defer exceptionProc()
    return ImageList_GetColorDepth(i.instance)
}

func (i *TImageList) SetColorDepth(value TColorDepth) {
    defer exceptionProc()
    ImageList_SetColorDepth(i.instance, value)
}

func (i *TImageList) DrawingStyle() TDrawingStyle {
    defer exceptionProc()
    return ImageList_GetDrawingStyle(i.instance)
}

func (i *TImageList) SetDrawingStyle(value TDrawingStyle) {
    defer exceptionProc()
    ImageList_SetDrawingStyle(i.instance, value)
}

func (i *TImageList) GrayscaleFactor() uint8 {
    defer exceptionProc()
    return ImageList_GetGrayscaleFactor(i.instance)
}

func (i *TImageList) SetGrayscaleFactor(value uint8) {
    defer exceptionProc()
    ImageList_SetGrayscaleFactor(i.instance, value)
}

func (i *TImageList) Height() int32 {
    defer exceptionProc()
    return ImageList_GetHeight(i.instance)
}

func (i *TImageList) SetHeight(value int32) {
    defer exceptionProc()
    ImageList_SetHeight(i.instance, value)
}

func (i *TImageList) ImageType() TImageType {
    defer exceptionProc()
    return ImageList_GetImageType(i.instance)
}

func (i *TImageList) SetImageType(value TImageType) {
    defer exceptionProc()
    ImageList_SetImageType(i.instance, value)
}

func (i *TImageList) Masked() bool {
    defer exceptionProc()
    return ImageList_GetMasked(i.instance)
}

func (i *TImageList) SetMasked(value bool) {
    defer exceptionProc()
    ImageList_SetMasked(i.instance, value)
}

func (i *TImageList) SetOnChange(fn TNotifyEvent) {
    ImageList_SetOnChange(i.instance, fn)
}

func (i *TImageList) ShareImages() bool {
    defer exceptionProc()
    return ImageList_GetShareImages(i.instance)
}

func (i *TImageList) SetShareImages(value bool) {
    defer exceptionProc()
    ImageList_SetShareImages(i.instance, value)
}

func (i *TImageList) Width() int32 {
    defer exceptionProc()
    return ImageList_GetWidth(i.instance)
}

func (i *TImageList) SetWidth(value int32) {
    defer exceptionProc()
    ImageList_SetWidth(i.instance, value)
}

func (i *TImageList) Handle() uintptr {
    defer exceptionProc()
    return ImageList_GetHandle(i.instance)
}

func (i *TImageList) SetHandle(value uintptr) {
    defer exceptionProc()
    ImageList_SetHandle(i.instance, value)
}

func (i *TImageList) ComponentCount() int32 {
    defer exceptionProc()
    return ImageList_GetComponentCount(i.instance)
}

func (i *TImageList) ComponentIndex() int32 {
    defer exceptionProc()
    return ImageList_GetComponentIndex(i.instance)
}

func (i *TImageList) SetComponentIndex(value int32) {
    defer exceptionProc()
    ImageList_SetComponentIndex(i.instance, value)
}

func (i *TImageList) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ImageList_GetOwner(i.instance))
}

func (i *TImageList) Name() string {
    defer exceptionProc()
    return ImageList_GetName(i.instance)
}

func (i *TImageList) SetName(value string) {
    defer exceptionProc()
    ImageList_SetName(i.instance, value)
}

func (i *TImageList) Tag() int {
    defer exceptionProc()
    return ImageList_GetTag(i.instance)
}

func (i *TImageList) SetTag(value int) {
    defer exceptionProc()
    ImageList_SetTag(i.instance, value)
}

func (i *TImageList) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(ImageList_GetComponents(i.instance, AIndex))
}

