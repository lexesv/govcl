
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TOpenPictureDialog struct {
    IComponent
    instance uintptr
}

func NewOpenPictureDialog(owner IComponent) *TOpenPictureDialog {
    o := new(TOpenPictureDialog)
    o.instance = OpenPictureDialog_Create(owner.Instance())
    return o
}

func OpenPictureDialogFromInst(inst uintptr) *TOpenPictureDialog {
    o := new(TOpenPictureDialog)
    o.instance = inst
    return o
}

func OpenPictureDialogFromObj(obj IObject) *TOpenPictureDialog {
    o := new(TOpenPictureDialog)
    o.instance = CheckPtr(obj)
    return o
}

func (o *TOpenPictureDialog) Free() {
    if o.instance != 0 {
        OpenPictureDialog_Free(o.instance)
    }
}

func (o *TOpenPictureDialog) Instance() uintptr {
    return o.instance
}

func (o *TOpenPictureDialog) IsValid() bool {
    return o.instance != 0
}

func (o *TOpenPictureDialog) Execute(ParentWnd HWND) bool {
    defer exceptionProc()
    return OpenPictureDialog_Execute(o.instance, ParentWnd )
}

func (o *TOpenPictureDialog) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(OpenPictureDialog_FindComponent(o.instance, AName ))
}

func (o *TOpenPictureDialog) GetNamePath() string {
    defer exceptionProc()
    return OpenPictureDialog_GetNamePath(o.instance)
}

func (o *TOpenPictureDialog) HasParent() bool {
    defer exceptionProc()
    return OpenPictureDialog_HasParent(o.instance)
}

func (o *TOpenPictureDialog) Assign(Source IObject) {
    defer exceptionProc()
    OpenPictureDialog_Assign(o.instance, CheckPtr(Source))
}

func (o *TOpenPictureDialog) ClassName() string {
    defer exceptionProc()
    return OpenPictureDialog_ClassName(o.instance)
}

func (o *TOpenPictureDialog) Equals(Obj IObject) bool {
    defer exceptionProc()
    return OpenPictureDialog_Equals(o.instance, CheckPtr(Obj))
}

func (o *TOpenPictureDialog) GetHashCode() int32 {
    defer exceptionProc()
    return OpenPictureDialog_GetHashCode(o.instance)
}

func (o *TOpenPictureDialog) ToString() string {
    defer exceptionProc()
    return OpenPictureDialog_ToString(o.instance)
}

func (o *TOpenPictureDialog) Filter() string {
    defer exceptionProc()
    return OpenPictureDialog_GetFilter(o.instance)
}

func (o *TOpenPictureDialog) SetFilter(value string) {
    defer exceptionProc()
    OpenPictureDialog_SetFilter(o.instance, value)
}

func (o *TOpenPictureDialog) Files() *TStrings {
    defer exceptionProc()
    return StringsFromInst(OpenPictureDialog_GetFiles(o.instance))
}

func (o *TOpenPictureDialog) DefaultExt() string {
    defer exceptionProc()
    return OpenPictureDialog_GetDefaultExt(o.instance)
}

func (o *TOpenPictureDialog) SetDefaultExt(value string) {
    defer exceptionProc()
    OpenPictureDialog_SetDefaultExt(o.instance, value)
}

func (o *TOpenPictureDialog) FileName() string {
    defer exceptionProc()
    return OpenPictureDialog_GetFileName(o.instance)
}

func (o *TOpenPictureDialog) SetFileName(value string) {
    defer exceptionProc()
    OpenPictureDialog_SetFileName(o.instance, value)
}

func (o *TOpenPictureDialog) FilterIndex() int32 {
    defer exceptionProc()
    return OpenPictureDialog_GetFilterIndex(o.instance)
}

func (o *TOpenPictureDialog) SetFilterIndex(value int32) {
    defer exceptionProc()
    OpenPictureDialog_SetFilterIndex(o.instance, value)
}

func (o *TOpenPictureDialog) InitialDir() string {
    defer exceptionProc()
    return OpenPictureDialog_GetInitialDir(o.instance)
}

func (o *TOpenPictureDialog) SetInitialDir(value string) {
    defer exceptionProc()
    OpenPictureDialog_SetInitialDir(o.instance, value)
}

func (o *TOpenPictureDialog) Options() TOpenOptions {
    defer exceptionProc()
    return OpenPictureDialog_GetOptions(o.instance)
}

func (o *TOpenPictureDialog) SetOptions(value TOpenOptions) {
    defer exceptionProc()
    OpenPictureDialog_SetOptions(o.instance, value)
}

func (o *TOpenPictureDialog) OptionsEx() TOpenOptionsEx {
    defer exceptionProc()
    return OpenPictureDialog_GetOptionsEx(o.instance)
}

func (o *TOpenPictureDialog) SetOptionsEx(value TOpenOptionsEx) {
    defer exceptionProc()
    OpenPictureDialog_SetOptionsEx(o.instance, value)
}

func (o *TOpenPictureDialog) Title() string {
    defer exceptionProc()
    return OpenPictureDialog_GetTitle(o.instance)
}

func (o *TOpenPictureDialog) SetTitle(value string) {
    defer exceptionProc()
    OpenPictureDialog_SetTitle(o.instance, value)
}

func (o *TOpenPictureDialog) Handle() HWND {
    defer exceptionProc()
    return OpenPictureDialog_GetHandle(o.instance)
}

func (o *TOpenPictureDialog) SetOnClose(fn TNotifyEvent) {
    OpenPictureDialog_SetOnClose(o.instance, fn)
}

func (o *TOpenPictureDialog) SetOnShow(fn TNotifyEvent) {
    OpenPictureDialog_SetOnShow(o.instance, fn)
}

func (o *TOpenPictureDialog) ComponentCount() int32 {
    defer exceptionProc()
    return OpenPictureDialog_GetComponentCount(o.instance)
}

func (o *TOpenPictureDialog) ComponentIndex() int32 {
    defer exceptionProc()
    return OpenPictureDialog_GetComponentIndex(o.instance)
}

func (o *TOpenPictureDialog) SetComponentIndex(value int32) {
    defer exceptionProc()
    OpenPictureDialog_SetComponentIndex(o.instance, value)
}

func (o *TOpenPictureDialog) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(OpenPictureDialog_GetOwner(o.instance))
}

func (o *TOpenPictureDialog) Name() string {
    defer exceptionProc()
    return OpenPictureDialog_GetName(o.instance)
}

func (o *TOpenPictureDialog) SetName(value string) {
    defer exceptionProc()
    OpenPictureDialog_SetName(o.instance, value)
}

func (o *TOpenPictureDialog) Tag() int {
    defer exceptionProc()
    return OpenPictureDialog_GetTag(o.instance)
}

func (o *TOpenPictureDialog) SetTag(value int) {
    defer exceptionProc()
    OpenPictureDialog_SetTag(o.instance, value)
}

func (o *TOpenPictureDialog) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(OpenPictureDialog_GetComponents(o.instance, AIndex))
}

