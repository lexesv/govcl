
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TSaveTextFileDialog struct {
    IComponent
    instance uintptr
}

func NewSaveTextFileDialog(owner IComponent) *TSaveTextFileDialog {
    s := new(TSaveTextFileDialog)
    s.instance = SaveTextFileDialog_Create(owner.Instance())
    return s
}

func SaveTextFileDialogFromInst(inst uintptr) *TSaveTextFileDialog {
    s := new(TSaveTextFileDialog)
    s.instance = inst
    return s
}

func SaveTextFileDialogFromObj(obj IObject) *TSaveTextFileDialog {
    s := new(TSaveTextFileDialog)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TSaveTextFileDialog) Free() {
    if s.instance != 0 {
        SaveTextFileDialog_Free(s.instance)
    }
}

func (s *TSaveTextFileDialog) Instance() uintptr {
    return s.instance
}

func (s *TSaveTextFileDialog) IsValid() bool {
    return s.instance != 0
}

func (s *TSaveTextFileDialog) Execute() bool {
    defer exceptionProc()
    return SaveTextFileDialog_Execute(s.instance)
}

func (s *TSaveTextFileDialog) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(SaveTextFileDialog_FindComponent(s.instance, AName ))
}

func (s *TSaveTextFileDialog) GetNamePath() string {
    defer exceptionProc()
    return SaveTextFileDialog_GetNamePath(s.instance)
}

func (s *TSaveTextFileDialog) HasParent() bool {
    defer exceptionProc()
    return SaveTextFileDialog_HasParent(s.instance)
}

func (s *TSaveTextFileDialog) Assign(Source IObject) {
    defer exceptionProc()
    SaveTextFileDialog_Assign(s.instance, CheckPtr(Source))
}

func (s *TSaveTextFileDialog) ClassName() string {
    defer exceptionProc()
    return SaveTextFileDialog_ClassName(s.instance)
}

func (s *TSaveTextFileDialog) Equals(Obj IObject) bool {
    defer exceptionProc()
    return SaveTextFileDialog_Equals(s.instance, CheckPtr(Obj))
}

func (s *TSaveTextFileDialog) GetHashCode() int32 {
    defer exceptionProc()
    return SaveTextFileDialog_GetHashCode(s.instance)
}

func (s *TSaveTextFileDialog) ToString() string {
    defer exceptionProc()
    return SaveTextFileDialog_ToString(s.instance)
}

func (s *TSaveTextFileDialog) Files() *TStrings {
    defer exceptionProc()
    return StringsFromInst(SaveTextFileDialog_GetFiles(s.instance))
}

func (s *TSaveTextFileDialog) DefaultExt() string {
    defer exceptionProc()
    return SaveTextFileDialog_GetDefaultExt(s.instance)
}

func (s *TSaveTextFileDialog) SetDefaultExt(value string) {
    defer exceptionProc()
    SaveTextFileDialog_SetDefaultExt(s.instance, value)
}

func (s *TSaveTextFileDialog) FileName() string {
    defer exceptionProc()
    return SaveTextFileDialog_GetFileName(s.instance)
}

func (s *TSaveTextFileDialog) SetFileName(value string) {
    defer exceptionProc()
    SaveTextFileDialog_SetFileName(s.instance, value)
}

func (s *TSaveTextFileDialog) Filter() string {
    defer exceptionProc()
    return SaveTextFileDialog_GetFilter(s.instance)
}

func (s *TSaveTextFileDialog) SetFilter(value string) {
    defer exceptionProc()
    SaveTextFileDialog_SetFilter(s.instance, value)
}

func (s *TSaveTextFileDialog) FilterIndex() int32 {
    defer exceptionProc()
    return SaveTextFileDialog_GetFilterIndex(s.instance)
}

func (s *TSaveTextFileDialog) SetFilterIndex(value int32) {
    defer exceptionProc()
    SaveTextFileDialog_SetFilterIndex(s.instance, value)
}

func (s *TSaveTextFileDialog) InitialDir() string {
    defer exceptionProc()
    return SaveTextFileDialog_GetInitialDir(s.instance)
}

func (s *TSaveTextFileDialog) SetInitialDir(value string) {
    defer exceptionProc()
    SaveTextFileDialog_SetInitialDir(s.instance, value)
}

func (s *TSaveTextFileDialog) Options() TOpenOptions {
    defer exceptionProc()
    return SaveTextFileDialog_GetOptions(s.instance)
}

func (s *TSaveTextFileDialog) SetOptions(value TOpenOptions) {
    defer exceptionProc()
    SaveTextFileDialog_SetOptions(s.instance, value)
}

func (s *TSaveTextFileDialog) OptionsEx() TOpenOptionsEx {
    defer exceptionProc()
    return SaveTextFileDialog_GetOptionsEx(s.instance)
}

func (s *TSaveTextFileDialog) SetOptionsEx(value TOpenOptionsEx) {
    defer exceptionProc()
    SaveTextFileDialog_SetOptionsEx(s.instance, value)
}

func (s *TSaveTextFileDialog) Title() string {
    defer exceptionProc()
    return SaveTextFileDialog_GetTitle(s.instance)
}

func (s *TSaveTextFileDialog) SetTitle(value string) {
    defer exceptionProc()
    SaveTextFileDialog_SetTitle(s.instance, value)
}

func (s *TSaveTextFileDialog) Handle() HWND {
    defer exceptionProc()
    return SaveTextFileDialog_GetHandle(s.instance)
}

func (s *TSaveTextFileDialog) SetOnClose(fn TNotifyEvent) {
    SaveTextFileDialog_SetOnClose(s.instance, fn)
}

func (s *TSaveTextFileDialog) SetOnShow(fn TNotifyEvent) {
    SaveTextFileDialog_SetOnShow(s.instance, fn)
}

func (s *TSaveTextFileDialog) ComponentCount() int32 {
    defer exceptionProc()
    return SaveTextFileDialog_GetComponentCount(s.instance)
}

func (s *TSaveTextFileDialog) ComponentIndex() int32 {
    defer exceptionProc()
    return SaveTextFileDialog_GetComponentIndex(s.instance)
}

func (s *TSaveTextFileDialog) SetComponentIndex(value int32) {
    defer exceptionProc()
    SaveTextFileDialog_SetComponentIndex(s.instance, value)
}

func (s *TSaveTextFileDialog) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(SaveTextFileDialog_GetOwner(s.instance))
}

func (s *TSaveTextFileDialog) Name() string {
    defer exceptionProc()
    return SaveTextFileDialog_GetName(s.instance)
}

func (s *TSaveTextFileDialog) SetName(value string) {
    defer exceptionProc()
    SaveTextFileDialog_SetName(s.instance, value)
}

func (s *TSaveTextFileDialog) Tag() int {
    defer exceptionProc()
    return SaveTextFileDialog_GetTag(s.instance)
}

func (s *TSaveTextFileDialog) SetTag(value int) {
    defer exceptionProc()
    SaveTextFileDialog_SetTag(s.instance, value)
}

func (s *TSaveTextFileDialog) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(SaveTextFileDialog_GetComponents(s.instance, AIndex))
}

