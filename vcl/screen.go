
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TScreen struct {
    IComponent
    instance uintptr
}

func NewScreen(owner IComponent) *TScreen {
    s := new(TScreen)
    s.instance = Screen_Create(owner.Instance())
    return s
}

func ScreenFromInst(inst uintptr) *TScreen {
    s := new(TScreen)
    s.instance = inst
    return s
}

func ScreenFromObj(obj IObject) *TScreen {
    s := new(TScreen)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TScreen) Free() {
    if s.instance != 0 {
        Screen_Free(s.instance)
    }
}

func (s *TScreen) Instance() uintptr {
    return s.instance
}

func (s *TScreen) IsValid() bool {
    return s.instance != 0
}

func (s *TScreen) Realign() {
    defer exceptionProc()
    Screen_Realign(s.instance)
}

func (s *TScreen) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Screen_FindComponent(s.instance, AName ))
}

func (s *TScreen) GetNamePath() string {
    defer exceptionProc()
    return Screen_GetNamePath(s.instance)
}

func (s *TScreen) HasParent() bool {
    defer exceptionProc()
    return Screen_HasParent(s.instance)
}

func (s *TScreen) Assign(Source IObject) {
    defer exceptionProc()
    Screen_Assign(s.instance, CheckPtr(Source))
}

func (s *TScreen) ClassName() string {
    defer exceptionProc()
    return Screen_ClassName(s.instance)
}

func (s *TScreen) Equals(Obj IObject) bool {
    defer exceptionProc()
    return Screen_Equals(s.instance, CheckPtr(Obj))
}

func (s *TScreen) GetHashCode() int32 {
    defer exceptionProc()
    return Screen_GetHashCode(s.instance)
}

func (s *TScreen) ToString() string {
    defer exceptionProc()
    return Screen_ToString(s.instance)
}

func (s *TScreen) ActiveForm() *TForm {
    defer exceptionProc()
    return FormFromInst(Screen_GetActiveForm(s.instance))
}

func (s *TScreen) CustomFormCount() int32 {
    defer exceptionProc()
    return Screen_GetCustomFormCount(s.instance)
}

func (s *TScreen) CursorCount() int32 {
    defer exceptionProc()
    return Screen_GetCursorCount(s.instance)
}

func (s *TScreen) Cursor() TCursor {
    defer exceptionProc()
    return Screen_GetCursor(s.instance)
}

func (s *TScreen) SetCursor(value TCursor) {
    defer exceptionProc()
    Screen_SetCursor(s.instance, value)
}

func (s *TScreen) FocusedForm() *TForm {
    defer exceptionProc()
    return FormFromInst(Screen_GetFocusedForm(s.instance))
}

func (s *TScreen) SetFocusedForm(value IControl) {
    defer exceptionProc()
    Screen_SetFocusedForm(s.instance, CheckPtr(value))
}

func (s *TScreen) MonitorCount() int32 {
    defer exceptionProc()
    return Screen_GetMonitorCount(s.instance)
}

func (s *TScreen) DesktopRect() TRect {
    defer exceptionProc()
    return Screen_GetDesktopRect(s.instance)
}

func (s *TScreen) DesktopHeight() int32 {
    defer exceptionProc()
    return Screen_GetDesktopHeight(s.instance)
}

func (s *TScreen) DesktopLeft() int32 {
    defer exceptionProc()
    return Screen_GetDesktopLeft(s.instance)
}

func (s *TScreen) DesktopTop() int32 {
    defer exceptionProc()
    return Screen_GetDesktopTop(s.instance)
}

func (s *TScreen) DesktopWidth() int32 {
    defer exceptionProc()
    return Screen_GetDesktopWidth(s.instance)
}

func (s *TScreen) WorkAreaRect() TRect {
    defer exceptionProc()
    return Screen_GetWorkAreaRect(s.instance)
}

func (s *TScreen) WorkAreaHeight() int32 {
    defer exceptionProc()
    return Screen_GetWorkAreaHeight(s.instance)
}

func (s *TScreen) WorkAreaLeft() int32 {
    defer exceptionProc()
    return Screen_GetWorkAreaLeft(s.instance)
}

func (s *TScreen) WorkAreaTop() int32 {
    defer exceptionProc()
    return Screen_GetWorkAreaTop(s.instance)
}

func (s *TScreen) WorkAreaWidth() int32 {
    defer exceptionProc()
    return Screen_GetWorkAreaWidth(s.instance)
}

func (s *TScreen) Fonts() *TStrings {
    defer exceptionProc()
    return StringsFromInst(Screen_GetFonts(s.instance))
}

func (s *TScreen) Imes() *TStrings {
    defer exceptionProc()
    return StringsFromInst(Screen_GetImes(s.instance))
}

func (s *TScreen) DefaultIme() string {
    defer exceptionProc()
    return Screen_GetDefaultIme(s.instance)
}

func (s *TScreen) Height() int32 {
    defer exceptionProc()
    return Screen_GetHeight(s.instance)
}

func (s *TScreen) PixelsPerInch() int32 {
    defer exceptionProc()
    return Screen_GetPixelsPerInch(s.instance)
}

func (s *TScreen) PrimaryMonitor() *TMonitor {
    defer exceptionProc()
    return MonitorFromInst(Screen_GetPrimaryMonitor(s.instance))
}

func (s *TScreen) Width() int32 {
    defer exceptionProc()
    return Screen_GetWidth(s.instance)
}

func (s *TScreen) ComponentCount() int32 {
    defer exceptionProc()
    return Screen_GetComponentCount(s.instance)
}

func (s *TScreen) ComponentIndex() int32 {
    defer exceptionProc()
    return Screen_GetComponentIndex(s.instance)
}

func (s *TScreen) SetComponentIndex(value int32) {
    defer exceptionProc()
    Screen_SetComponentIndex(s.instance, value)
}

func (s *TScreen) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Screen_GetOwner(s.instance))
}

func (s *TScreen) Name() string {
    defer exceptionProc()
    return Screen_GetName(s.instance)
}

func (s *TScreen) SetName(value string) {
    defer exceptionProc()
    Screen_SetName(s.instance, value)
}

func (s *TScreen) Tag() int {
    defer exceptionProc()
    return Screen_GetTag(s.instance)
}

func (s *TScreen) SetTag(value int) {
    defer exceptionProc()
    Screen_SetTag(s.instance, value)
}

func (s *TScreen) Cursors(Index int32) HICON {
    defer exceptionProc()
    return Screen_GetCursors(s.instance, Index)
}

func (s *TScreen) SetCursors(Index int32, value HICON) {
    defer exceptionProc()
    Screen_SetCursors(s.instance, Index, value)
}

func (s *TScreen) Monitors(Index int32) *TMonitor {
    defer exceptionProc()
    return MonitorFromInst(Screen_GetMonitors(s.instance, Index))
}

func (s *TScreen) Forms(Index int32) *TForm {
    defer exceptionProc()
    return FormFromInst(Screen_GetForms(s.instance, Index))
}

func (s *TScreen) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Screen_GetComponents(s.instance, AIndex))
}

