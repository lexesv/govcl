
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

type TSaveDialog struct {
    IComponent
    instance uintptr
}

func NewSaveDialog(owner IComponent) *TSaveDialog {
    s := new(TSaveDialog)
    s.instance = SaveDialog_Create(owner.Instance())
    return s
}

func SaveDialogFromInst(inst uintptr) *TSaveDialog {
    s := new(TSaveDialog)
    s.instance = inst
    return s
}

func SaveDialogFromObj(obj IObject) *TSaveDialog {
    s := new(TSaveDialog)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TSaveDialog) Free() {
    if s.instance != 0 {
        SaveDialog_Free(s.instance)
    }
}

func (s *TSaveDialog) Instance() uintptr {
    return s.instance
}

func (s *TSaveDialog) IsValid() bool {
    return s.instance != 0
}

func (s *TSaveDialog) Execute(ParentWnd HWND) bool {
    defer exceptionProc()
    return SaveDialog_Execute(s.instance, ParentWnd )
}

func (s *TSaveDialog) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(SaveDialog_FindComponent(s.instance, AName ))
}

func (s *TSaveDialog) GetNamePath() string {
    defer exceptionProc()
    return SaveDialog_GetNamePath(s.instance)
}

func (s *TSaveDialog) HasParent() bool {
    defer exceptionProc()
    return SaveDialog_HasParent(s.instance)
}

func (s *TSaveDialog) Assign(Source IObject) {
    defer exceptionProc()
    SaveDialog_Assign(s.instance, CheckPtr(Source))
}

func (s *TSaveDialog) ClassName() string {
    defer exceptionProc()
    return SaveDialog_ClassName(s.instance)
}

func (s *TSaveDialog) Equals(Obj IObject) bool {
    defer exceptionProc()
    return SaveDialog_Equals(s.instance, CheckPtr(Obj))
}

func (s *TSaveDialog) GetHashCode() int32 {
    defer exceptionProc()
    return SaveDialog_GetHashCode(s.instance)
}

func (s *TSaveDialog) ToString() string {
    defer exceptionProc()
    return SaveDialog_ToString(s.instance)
}

func (s *TSaveDialog) Files() *TStrings {
    defer exceptionProc()
    return StringsFromInst(SaveDialog_GetFiles(s.instance))
}

func (s *TSaveDialog) DefaultExt() string {
    defer exceptionProc()
    return SaveDialog_GetDefaultExt(s.instance)
}

func (s *TSaveDialog) SetDefaultExt(value string) {
    defer exceptionProc()
    SaveDialog_SetDefaultExt(s.instance, value)
}

func (s *TSaveDialog) FileName() string {
    defer exceptionProc()
    return SaveDialog_GetFileName(s.instance)
}

func (s *TSaveDialog) SetFileName(value string) {
    defer exceptionProc()
    SaveDialog_SetFileName(s.instance, value)
}

func (s *TSaveDialog) Filter() string {
    defer exceptionProc()
    return SaveDialog_GetFilter(s.instance)
}

func (s *TSaveDialog) SetFilter(value string) {
    defer exceptionProc()
    SaveDialog_SetFilter(s.instance, value)
}

func (s *TSaveDialog) FilterIndex() int32 {
    defer exceptionProc()
    return SaveDialog_GetFilterIndex(s.instance)
}

func (s *TSaveDialog) SetFilterIndex(value int32) {
    defer exceptionProc()
    SaveDialog_SetFilterIndex(s.instance, value)
}

func (s *TSaveDialog) InitialDir() string {
    defer exceptionProc()
    return SaveDialog_GetInitialDir(s.instance)
}

func (s *TSaveDialog) SetInitialDir(value string) {
    defer exceptionProc()
    SaveDialog_SetInitialDir(s.instance, value)
}

func (s *TSaveDialog) Options() TOpenOptions {
    defer exceptionProc()
    return SaveDialog_GetOptions(s.instance)
}

func (s *TSaveDialog) SetOptions(value TOpenOptions) {
    defer exceptionProc()
    SaveDialog_SetOptions(s.instance, value)
}

func (s *TSaveDialog) OptionsEx() TOpenOptionsEx {
    defer exceptionProc()
    return SaveDialog_GetOptionsEx(s.instance)
}

func (s *TSaveDialog) SetOptionsEx(value TOpenOptionsEx) {
    defer exceptionProc()
    SaveDialog_SetOptionsEx(s.instance, value)
}

func (s *TSaveDialog) Title() string {
    defer exceptionProc()
    return SaveDialog_GetTitle(s.instance)
}

func (s *TSaveDialog) SetTitle(value string) {
    defer exceptionProc()
    SaveDialog_SetTitle(s.instance, value)
}

func (s *TSaveDialog) Handle() HWND {
    defer exceptionProc()
    return SaveDialog_GetHandle(s.instance)
}

func (s *TSaveDialog) SetOnClose(fn TNotifyEvent) {
    SaveDialog_SetOnClose(s.instance, fn)
}

func (s *TSaveDialog) SetOnShow(fn TNotifyEvent) {
    SaveDialog_SetOnShow(s.instance, fn)
}

func (s *TSaveDialog) ComponentCount() int32 {
    defer exceptionProc()
    return SaveDialog_GetComponentCount(s.instance)
}

func (s *TSaveDialog) ComponentIndex() int32 {
    defer exceptionProc()
    return SaveDialog_GetComponentIndex(s.instance)
}

func (s *TSaveDialog) SetComponentIndex(value int32) {
    defer exceptionProc()
    SaveDialog_SetComponentIndex(s.instance, value)
}

func (s *TSaveDialog) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(SaveDialog_GetOwner(s.instance))
}

func (s *TSaveDialog) Name() string {
    defer exceptionProc()
    return SaveDialog_GetName(s.instance)
}

func (s *TSaveDialog) SetName(value string) {
    defer exceptionProc()
    SaveDialog_SetName(s.instance, value)
}

func (s *TSaveDialog) Tag() int {
    defer exceptionProc()
    return SaveDialog_GetTag(s.instance)
}

func (s *TSaveDialog) SetTag(value int) {
    defer exceptionProc()
    SaveDialog_SetTag(s.instance, value)
}

func (s *TSaveDialog) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(SaveDialog_GetComponents(s.instance, AIndex))
}

