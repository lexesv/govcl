
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

type TBalloonHint struct {
    IComponent
    instance uintptr
}

func NewBalloonHint(owner IComponent) *TBalloonHint {
    b := new(TBalloonHint)
    b.instance = BalloonHint_Create(owner.Instance())
    return b
}

func BalloonHintFromInst(inst uintptr) *TBalloonHint {
    b := new(TBalloonHint)
    b.instance = inst
    return b
}

func BalloonHintFromObj(obj IObject) *TBalloonHint {
    b := new(TBalloonHint)
    b.instance = CheckPtr(obj)
    return b
}

func (b *TBalloonHint) Free() {
    if b.instance != 0 {
        BalloonHint_Free(b.instance)
    }
}

func (b *TBalloonHint) Instance() uintptr {
    return b.instance
}

func (b *TBalloonHint) IsValid() bool {
    return b.instance != 0
}

func (b *TBalloonHint) ShowHint() {
    defer exceptionProc()
    BalloonHint_ShowHint(b.instance)
}

func (b *TBalloonHint) HideHint() {
    defer exceptionProc()
    BalloonHint_HideHint(b.instance)
}

func (b *TBalloonHint) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(BalloonHint_FindComponent(b.instance, AName ))
}

func (b *TBalloonHint) GetNamePath() string {
    defer exceptionProc()
    return BalloonHint_GetNamePath(b.instance)
}

func (b *TBalloonHint) HasParent() bool {
    defer exceptionProc()
    return BalloonHint_HasParent(b.instance)
}

func (b *TBalloonHint) Assign(Source IObject) {
    defer exceptionProc()
    BalloonHint_Assign(b.instance, CheckPtr(Source))
}

func (b *TBalloonHint) ClassName() string {
    defer exceptionProc()
    return BalloonHint_ClassName(b.instance)
}

func (b *TBalloonHint) Equals(Obj IObject) bool {
    defer exceptionProc()
    return BalloonHint_Equals(b.instance, CheckPtr(Obj))
}

func (b *TBalloonHint) GetHashCode() int32 {
    defer exceptionProc()
    return BalloonHint_GetHashCode(b.instance)
}

func (b *TBalloonHint) ToString() string {
    defer exceptionProc()
    return BalloonHint_ToString(b.instance)
}

func (b *TBalloonHint) Title() string {
    defer exceptionProc()
    return BalloonHint_GetTitle(b.instance)
}

func (b *TBalloonHint) SetTitle(value string) {
    defer exceptionProc()
    BalloonHint_SetTitle(b.instance, value)
}

func (b *TBalloonHint) Description() string {
    defer exceptionProc()
    return BalloonHint_GetDescription(b.instance)
}

func (b *TBalloonHint) SetDescription(value string) {
    defer exceptionProc()
    BalloonHint_SetDescription(b.instance, value)
}

func (b *TBalloonHint) ImageIndex() int32 {
    defer exceptionProc()
    return BalloonHint_GetImageIndex(b.instance)
}

func (b *TBalloonHint) SetImageIndex(value int32) {
    defer exceptionProc()
    BalloonHint_SetImageIndex(b.instance, value)
}

func (b *TBalloonHint) Images() *TImageList {
    defer exceptionProc()
    return ImageListFromInst(BalloonHint_GetImages(b.instance))
}

func (b *TBalloonHint) SetImages(value IComponent) {
    defer exceptionProc()
    BalloonHint_SetImages(b.instance, CheckPtr(value))
}

func (b *TBalloonHint) Style() TBalloonHintStyle {
    defer exceptionProc()
    return BalloonHint_GetStyle(b.instance)
}

func (b *TBalloonHint) SetStyle(value TBalloonHintStyle) {
    defer exceptionProc()
    BalloonHint_SetStyle(b.instance, value)
}

func (b *TBalloonHint) Delay() uint32 {
    defer exceptionProc()
    return BalloonHint_GetDelay(b.instance)
}

func (b *TBalloonHint) SetDelay(value uint32) {
    defer exceptionProc()
    BalloonHint_SetDelay(b.instance, value)
}

func (b *TBalloonHint) HideAfter() int32 {
    defer exceptionProc()
    return BalloonHint_GetHideAfter(b.instance)
}

func (b *TBalloonHint) SetHideAfter(value int32) {
    defer exceptionProc()
    BalloonHint_SetHideAfter(b.instance, value)
}

func (b *TBalloonHint) ComponentCount() int32 {
    defer exceptionProc()
    return BalloonHint_GetComponentCount(b.instance)
}

func (b *TBalloonHint) ComponentIndex() int32 {
    defer exceptionProc()
    return BalloonHint_GetComponentIndex(b.instance)
}

func (b *TBalloonHint) SetComponentIndex(value int32) {
    defer exceptionProc()
    BalloonHint_SetComponentIndex(b.instance, value)
}

func (b *TBalloonHint) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(BalloonHint_GetOwner(b.instance))
}

func (b *TBalloonHint) Name() string {
    defer exceptionProc()
    return BalloonHint_GetName(b.instance)
}

func (b *TBalloonHint) SetName(value string) {
    defer exceptionProc()
    BalloonHint_SetName(b.instance, value)
}

func (b *TBalloonHint) Tag() int {
    defer exceptionProc()
    return BalloonHint_GetTag(b.instance)
}

func (b *TBalloonHint) SetTag(value int) {
    defer exceptionProc()
    BalloonHint_SetTag(b.instance, value)
}

func (b *TBalloonHint) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(BalloonHint_GetComponents(b.instance, AIndex))
}

