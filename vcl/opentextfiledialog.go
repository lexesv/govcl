
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TOpenTextFileDialog struct {
    IComponent
    instance uintptr
}

func NewOpenTextFileDialog(owner IComponent) *TOpenTextFileDialog {
    o := new(TOpenTextFileDialog)
    o.instance = OpenTextFileDialog_Create(owner.Instance())
    return o
}

func OpenTextFileDialogFromInst(inst uintptr) *TOpenTextFileDialog {
    o := new(TOpenTextFileDialog)
    o.instance = inst
    return o
}

func OpenTextFileDialogFromObj(obj IObject) *TOpenTextFileDialog {
    o := new(TOpenTextFileDialog)
    o.instance = CheckPtr(obj)
    return o
}

func (o *TOpenTextFileDialog) Free() {
    if o.instance != 0 {
        OpenTextFileDialog_Free(o.instance)
    }
}

func (o *TOpenTextFileDialog) Instance() uintptr {
    return o.instance
}

func (o *TOpenTextFileDialog) IsValid() bool {
    return o.instance != 0
}

func (o *TOpenTextFileDialog) Execute(ParentWnd HWND) bool {
    defer exceptionProc()
    return OpenTextFileDialog_Execute(o.instance, ParentWnd )
}

func (o *TOpenTextFileDialog) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(OpenTextFileDialog_FindComponent(o.instance, AName ))
}

func (o *TOpenTextFileDialog) GetNamePath() string {
    defer exceptionProc()
    return OpenTextFileDialog_GetNamePath(o.instance)
}

func (o *TOpenTextFileDialog) HasParent() bool {
    defer exceptionProc()
    return OpenTextFileDialog_HasParent(o.instance)
}

func (o *TOpenTextFileDialog) Assign(Source IObject) {
    defer exceptionProc()
    OpenTextFileDialog_Assign(o.instance, CheckPtr(Source))
}

func (o *TOpenTextFileDialog) ClassName() string {
    defer exceptionProc()
    return OpenTextFileDialog_ClassName(o.instance)
}

func (o *TOpenTextFileDialog) Equals(Obj IObject) bool {
    defer exceptionProc()
    return OpenTextFileDialog_Equals(o.instance, CheckPtr(Obj))
}

func (o *TOpenTextFileDialog) GetHashCode() int32 {
    defer exceptionProc()
    return OpenTextFileDialog_GetHashCode(o.instance)
}

func (o *TOpenTextFileDialog) ToString() string {
    defer exceptionProc()
    return OpenTextFileDialog_ToString(o.instance)
}

func (o *TOpenTextFileDialog) Files() *TStrings {
    defer exceptionProc()
    return StringsFromInst(OpenTextFileDialog_GetFiles(o.instance))
}

func (o *TOpenTextFileDialog) DefaultExt() string {
    defer exceptionProc()
    return OpenTextFileDialog_GetDefaultExt(o.instance)
}

func (o *TOpenTextFileDialog) SetDefaultExt(value string) {
    defer exceptionProc()
    OpenTextFileDialog_SetDefaultExt(o.instance, value)
}

func (o *TOpenTextFileDialog) FileName() string {
    defer exceptionProc()
    return OpenTextFileDialog_GetFileName(o.instance)
}

func (o *TOpenTextFileDialog) SetFileName(value string) {
    defer exceptionProc()
    OpenTextFileDialog_SetFileName(o.instance, value)
}

func (o *TOpenTextFileDialog) Filter() string {
    defer exceptionProc()
    return OpenTextFileDialog_GetFilter(o.instance)
}

func (o *TOpenTextFileDialog) SetFilter(value string) {
    defer exceptionProc()
    OpenTextFileDialog_SetFilter(o.instance, value)
}

func (o *TOpenTextFileDialog) FilterIndex() int32 {
    defer exceptionProc()
    return OpenTextFileDialog_GetFilterIndex(o.instance)
}

func (o *TOpenTextFileDialog) SetFilterIndex(value int32) {
    defer exceptionProc()
    OpenTextFileDialog_SetFilterIndex(o.instance, value)
}

func (o *TOpenTextFileDialog) InitialDir() string {
    defer exceptionProc()
    return OpenTextFileDialog_GetInitialDir(o.instance)
}

func (o *TOpenTextFileDialog) SetInitialDir(value string) {
    defer exceptionProc()
    OpenTextFileDialog_SetInitialDir(o.instance, value)
}

func (o *TOpenTextFileDialog) Options() TOpenOptions {
    defer exceptionProc()
    return OpenTextFileDialog_GetOptions(o.instance)
}

func (o *TOpenTextFileDialog) SetOptions(value TOpenOptions) {
    defer exceptionProc()
    OpenTextFileDialog_SetOptions(o.instance, value)
}

func (o *TOpenTextFileDialog) OptionsEx() TOpenOptionsEx {
    defer exceptionProc()
    return OpenTextFileDialog_GetOptionsEx(o.instance)
}

func (o *TOpenTextFileDialog) SetOptionsEx(value TOpenOptionsEx) {
    defer exceptionProc()
    OpenTextFileDialog_SetOptionsEx(o.instance, value)
}

func (o *TOpenTextFileDialog) Title() string {
    defer exceptionProc()
    return OpenTextFileDialog_GetTitle(o.instance)
}

func (o *TOpenTextFileDialog) SetTitle(value string) {
    defer exceptionProc()
    OpenTextFileDialog_SetTitle(o.instance, value)
}

func (o *TOpenTextFileDialog) Handle() HWND {
    defer exceptionProc()
    return OpenTextFileDialog_GetHandle(o.instance)
}

func (o *TOpenTextFileDialog) SetOnClose(fn TNotifyEvent) {
    OpenTextFileDialog_SetOnClose(o.instance, fn)
}

func (o *TOpenTextFileDialog) SetOnShow(fn TNotifyEvent) {
    OpenTextFileDialog_SetOnShow(o.instance, fn)
}

func (o *TOpenTextFileDialog) ComponentCount() int32 {
    defer exceptionProc()
    return OpenTextFileDialog_GetComponentCount(o.instance)
}

func (o *TOpenTextFileDialog) ComponentIndex() int32 {
    defer exceptionProc()
    return OpenTextFileDialog_GetComponentIndex(o.instance)
}

func (o *TOpenTextFileDialog) SetComponentIndex(value int32) {
    defer exceptionProc()
    OpenTextFileDialog_SetComponentIndex(o.instance, value)
}

func (o *TOpenTextFileDialog) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(OpenTextFileDialog_GetOwner(o.instance))
}

func (o *TOpenTextFileDialog) Name() string {
    defer exceptionProc()
    return OpenTextFileDialog_GetName(o.instance)
}

func (o *TOpenTextFileDialog) SetName(value string) {
    defer exceptionProc()
    OpenTextFileDialog_SetName(o.instance, value)
}

func (o *TOpenTextFileDialog) Tag() int {
    defer exceptionProc()
    return OpenTextFileDialog_GetTag(o.instance)
}

func (o *TOpenTextFileDialog) SetTag(value int) {
    defer exceptionProc()
    OpenTextFileDialog_SetTag(o.instance, value)
}

func (o *TOpenTextFileDialog) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(OpenTextFileDialog_GetComponents(o.instance, AIndex))
}

