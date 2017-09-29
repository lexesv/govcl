
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

type TSavePictureDialog struct {
    IComponent
    instance uintptr
}

func NewSavePictureDialog(owner IComponent) *TSavePictureDialog {
    s := new(TSavePictureDialog)
    s.instance = SavePictureDialog_Create(owner.Instance())
    return s
}

func SavePictureDialogFromInst(inst uintptr) *TSavePictureDialog {
    s := new(TSavePictureDialog)
    s.instance = inst
    return s
}

func SavePictureDialogFromObj(obj IObject) *TSavePictureDialog {
    s := new(TSavePictureDialog)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TSavePictureDialog) Free() {
    if s.instance != 0 {
        SavePictureDialog_Free(s.instance)
    }
}

func (s *TSavePictureDialog) Instance() uintptr {
    return s.instance
}

func (s *TSavePictureDialog) IsValid() bool {
    return s.instance != 0
}

func (s *TSavePictureDialog) Execute() bool {
    defer exceptionProc()
    return SavePictureDialog_Execute(s.instance)
}

func (s *TSavePictureDialog) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(SavePictureDialog_FindComponent(s.instance, AName ))
}

func (s *TSavePictureDialog) GetNamePath() string {
    defer exceptionProc()
    return SavePictureDialog_GetNamePath(s.instance)
}

func (s *TSavePictureDialog) HasParent() bool {
    defer exceptionProc()
    return SavePictureDialog_HasParent(s.instance)
}

func (s *TSavePictureDialog) Assign(Source IObject) {
    defer exceptionProc()
    SavePictureDialog_Assign(s.instance, CheckPtr(Source))
}

func (s *TSavePictureDialog) ClassName() string {
    defer exceptionProc()
    return SavePictureDialog_ClassName(s.instance)
}

func (s *TSavePictureDialog) Equals(Obj IObject) bool {
    defer exceptionProc()
    return SavePictureDialog_Equals(s.instance, CheckPtr(Obj))
}

func (s *TSavePictureDialog) GetHashCode() int32 {
    defer exceptionProc()
    return SavePictureDialog_GetHashCode(s.instance)
}

func (s *TSavePictureDialog) ToString() string {
    defer exceptionProc()
    return SavePictureDialog_ToString(s.instance)
}

func (s *TSavePictureDialog) Filter() string {
    defer exceptionProc()
    return SavePictureDialog_GetFilter(s.instance)
}

func (s *TSavePictureDialog) SetFilter(value string) {
    defer exceptionProc()
    SavePictureDialog_SetFilter(s.instance, value)
}

func (s *TSavePictureDialog) Files() *TStrings {
    defer exceptionProc()
    return StringsFromInst(SavePictureDialog_GetFiles(s.instance))
}

func (s *TSavePictureDialog) DefaultExt() string {
    defer exceptionProc()
    return SavePictureDialog_GetDefaultExt(s.instance)
}

func (s *TSavePictureDialog) SetDefaultExt(value string) {
    defer exceptionProc()
    SavePictureDialog_SetDefaultExt(s.instance, value)
}

func (s *TSavePictureDialog) FileName() string {
    defer exceptionProc()
    return SavePictureDialog_GetFileName(s.instance)
}

func (s *TSavePictureDialog) SetFileName(value string) {
    defer exceptionProc()
    SavePictureDialog_SetFileName(s.instance, value)
}

func (s *TSavePictureDialog) FilterIndex() int32 {
    defer exceptionProc()
    return SavePictureDialog_GetFilterIndex(s.instance)
}

func (s *TSavePictureDialog) SetFilterIndex(value int32) {
    defer exceptionProc()
    SavePictureDialog_SetFilterIndex(s.instance, value)
}

func (s *TSavePictureDialog) InitialDir() string {
    defer exceptionProc()
    return SavePictureDialog_GetInitialDir(s.instance)
}

func (s *TSavePictureDialog) SetInitialDir(value string) {
    defer exceptionProc()
    SavePictureDialog_SetInitialDir(s.instance, value)
}

func (s *TSavePictureDialog) Options() TOpenOptions {
    defer exceptionProc()
    return SavePictureDialog_GetOptions(s.instance)
}

func (s *TSavePictureDialog) SetOptions(value TOpenOptions) {
    defer exceptionProc()
    SavePictureDialog_SetOptions(s.instance, value)
}

func (s *TSavePictureDialog) OptionsEx() TOpenOptionsEx {
    defer exceptionProc()
    return SavePictureDialog_GetOptionsEx(s.instance)
}

func (s *TSavePictureDialog) SetOptionsEx(value TOpenOptionsEx) {
    defer exceptionProc()
    SavePictureDialog_SetOptionsEx(s.instance, value)
}

func (s *TSavePictureDialog) Title() string {
    defer exceptionProc()
    return SavePictureDialog_GetTitle(s.instance)
}

func (s *TSavePictureDialog) SetTitle(value string) {
    defer exceptionProc()
    SavePictureDialog_SetTitle(s.instance, value)
}

func (s *TSavePictureDialog) Handle() HWND {
    defer exceptionProc()
    return SavePictureDialog_GetHandle(s.instance)
}

func (s *TSavePictureDialog) SetOnClose(fn TNotifyEvent) {
    SavePictureDialog_SetOnClose(s.instance, fn)
}

func (s *TSavePictureDialog) SetOnShow(fn TNotifyEvent) {
    SavePictureDialog_SetOnShow(s.instance, fn)
}

func (s *TSavePictureDialog) ComponentCount() int32 {
    defer exceptionProc()
    return SavePictureDialog_GetComponentCount(s.instance)
}

func (s *TSavePictureDialog) ComponentIndex() int32 {
    defer exceptionProc()
    return SavePictureDialog_GetComponentIndex(s.instance)
}

func (s *TSavePictureDialog) SetComponentIndex(value int32) {
    defer exceptionProc()
    SavePictureDialog_SetComponentIndex(s.instance, value)
}

func (s *TSavePictureDialog) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(SavePictureDialog_GetOwner(s.instance))
}

func (s *TSavePictureDialog) Name() string {
    defer exceptionProc()
    return SavePictureDialog_GetName(s.instance)
}

func (s *TSavePictureDialog) SetName(value string) {
    defer exceptionProc()
    SavePictureDialog_SetName(s.instance, value)
}

func (s *TSavePictureDialog) Tag() int {
    defer exceptionProc()
    return SavePictureDialog_GetTag(s.instance)
}

func (s *TSavePictureDialog) SetTag(value int) {
    defer exceptionProc()
    SavePictureDialog_SetTag(s.instance, value)
}

func (s *TSavePictureDialog) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(SavePictureDialog_GetComponents(s.instance, AIndex))
}

