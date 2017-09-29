
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

type TApplication struct {
    IComponent
    instance uintptr
}

func NewApplication(owner IComponent) *TApplication {
    a := new(TApplication)
    a.instance = Application_Create(owner.Instance())
    return a
}

func ApplicationFromInst(inst uintptr) *TApplication {
    a := new(TApplication)
    a.instance = inst
    return a
}

func ApplicationFromObj(obj IObject) *TApplication {
    a := new(TApplication)
    a.instance = CheckPtr(obj)
    return a
}

func (a *TApplication) Free() {
    if a.instance != 0 {
        Application_Free(a.instance)
    }
}

func (a *TApplication) Instance() uintptr {
    return a.instance
}

func (a *TApplication) IsValid() bool {
    return a.instance != 0
}

func (a *TApplication) ActivateHint(CursorPos TPoint) {
    defer exceptionProc()
    Application_ActivateHint(a.instance, CursorPos )
}

func (a *TApplication) BringToFront() {
    defer exceptionProc()
    Application_BringToFront(a.instance)
}

func (a *TApplication) CancelHint() {
    defer exceptionProc()
    Application_CancelHint(a.instance)
}

func (a *TApplication) HandleMessage() {
    defer exceptionProc()
    Application_HandleMessage(a.instance)
}

func (a *TApplication) HideHint() {
    defer exceptionProc()
    Application_HideHint(a.instance)
}

func (a *TApplication) Initialize() {
    defer exceptionProc()
    Application_Initialize(a.instance)
}

func (a *TApplication) Minimize() {
    defer exceptionProc()
    Application_Minimize(a.instance)
}

func (a *TApplication) ModalStarted() {
    defer exceptionProc()
    Application_ModalStarted(a.instance)
}

func (a *TApplication) ModalFinished() {
    defer exceptionProc()
    Application_ModalFinished(a.instance)
}

func (a *TApplication) NormalizeAllTopMosts() {
    defer exceptionProc()
    Application_NormalizeAllTopMosts(a.instance)
}

func (a *TApplication) NormalizeTopMosts() {
    defer exceptionProc()
    Application_NormalizeTopMosts(a.instance)
}

func (a *TApplication) ProcessMessages() {
    defer exceptionProc()
    Application_ProcessMessages(a.instance)
}

func (a *TApplication) Restore() {
    defer exceptionProc()
    Application_Restore(a.instance)
}

func (a *TApplication) RestoreTopMosts() {
    defer exceptionProc()
    Application_RestoreTopMosts(a.instance)
}

func (a *TApplication) Run() {
    defer exceptionProc()
    Application_Run(a.instance)
}

func (a *TApplication) Terminate() {
    defer exceptionProc()
    Application_Terminate(a.instance)
}

func (a *TApplication) MessageBox(Text string, Caption string, Flags int32) int32 {
    defer exceptionProc()
    return Application_MessageBox(a.instance, Text , Caption , Flags )
}

func (a *TApplication) FindComponent(AName string) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Application_FindComponent(a.instance, AName ))
}

func (a *TApplication) GetNamePath() string {
    defer exceptionProc()
    return Application_GetNamePath(a.instance)
}

func (a *TApplication) HasParent() bool {
    defer exceptionProc()
    return Application_HasParent(a.instance)
}

func (a *TApplication) Assign(Source IObject) {
    defer exceptionProc()
    Application_Assign(a.instance, CheckPtr(Source))
}

func (a *TApplication) ClassName() string {
    defer exceptionProc()
    return Application_ClassName(a.instance)
}

func (a *TApplication) Equals(Obj IObject) bool {
    defer exceptionProc()
    return Application_Equals(a.instance, CheckPtr(Obj))
}

func (a *TApplication) GetHashCode() int32 {
    defer exceptionProc()
    return Application_GetHashCode(a.instance)
}

func (a *TApplication) ToString() string {
    defer exceptionProc()
    return Application_ToString(a.instance)
}

func (a *TApplication) DialogHandle() HWND {
    defer exceptionProc()
    return Application_GetDialogHandle(a.instance)
}

func (a *TApplication) SetDialogHandle(value HWND) {
    defer exceptionProc()
    Application_SetDialogHandle(a.instance, value)
}

func (a *TApplication) Hint() string {
    defer exceptionProc()
    return Application_GetHint(a.instance)
}

func (a *TApplication) SetHint(value string) {
    defer exceptionProc()
    Application_SetHint(a.instance, value)
}

func (a *TApplication) HintColor() TColor {
    defer exceptionProc()
    return Application_GetHintColor(a.instance)
}

func (a *TApplication) SetHintColor(value TColor) {
    defer exceptionProc()
    Application_SetHintColor(a.instance, value)
}

func (a *TApplication) HintHidePause() int32 {
    defer exceptionProc()
    return Application_GetHintHidePause(a.instance)
}

func (a *TApplication) SetHintHidePause(value int32) {
    defer exceptionProc()
    Application_SetHintHidePause(a.instance, value)
}

func (a *TApplication) HintPause() int32 {
    defer exceptionProc()
    return Application_GetHintPause(a.instance)
}

func (a *TApplication) SetHintPause(value int32) {
    defer exceptionProc()
    Application_SetHintPause(a.instance, value)
}

func (a *TApplication) HintShortCuts() bool {
    defer exceptionProc()
    return Application_GetHintShortCuts(a.instance)
}

func (a *TApplication) SetHintShortCuts(value bool) {
    defer exceptionProc()
    Application_SetHintShortCuts(a.instance, value)
}

func (a *TApplication) HintShortPause() int32 {
    defer exceptionProc()
    return Application_GetHintShortPause(a.instance)
}

func (a *TApplication) SetHintShortPause(value int32) {
    defer exceptionProc()
    Application_SetHintShortPause(a.instance, value)
}

func (a *TApplication) Icon() *TIcon {
    defer exceptionProc()
    return IconFromInst(Application_GetIcon(a.instance))
}

func (a *TApplication) SetIcon(value *TIcon) {
    defer exceptionProc()
    Application_SetIcon(a.instance, CheckPtr(value))
}

func (a *TApplication) IsMetropolisUI() bool {
    defer exceptionProc()
    return Application_GetIsMetropolisUI(a.instance)
}

func (a *TApplication) MainForm() *TForm {
    defer exceptionProc()
    return FormFromInst(Application_GetMainForm(a.instance))
}

func (a *TApplication) MainFormHandle() HWND {
    defer exceptionProc()
    return Application_GetMainFormHandle(a.instance)
}

func (a *TApplication) MainFormOnTaskBar() bool {
    defer exceptionProc()
    return Application_GetMainFormOnTaskBar(a.instance)
}

func (a *TApplication) SetMainFormOnTaskBar(value bool) {
    defer exceptionProc()
    Application_SetMainFormOnTaskBar(a.instance, value)
}

func (a *TApplication) ShowHint() bool {
    defer exceptionProc()
    return Application_GetShowHint(a.instance)
}

func (a *TApplication) SetShowHint(value bool) {
    defer exceptionProc()
    Application_SetShowHint(a.instance, value)
}

func (a *TApplication) ShowMainForm() bool {
    defer exceptionProc()
    return Application_GetShowMainForm(a.instance)
}

func (a *TApplication) SetShowMainForm(value bool) {
    defer exceptionProc()
    Application_SetShowMainForm(a.instance, value)
}

func (a *TApplication) Title() string {
    defer exceptionProc()
    return Application_GetTitle(a.instance)
}

func (a *TApplication) SetTitle(value string) {
    defer exceptionProc()
    Application_SetTitle(a.instance, value)
}

func (a *TApplication) SetOnException(fn TExceptionEvent) {
    Application_SetOnException(a.instance, fn)
}

func (a *TApplication) SetOnMinimize(fn TNotifyEvent) {
    Application_SetOnMinimize(a.instance, fn)
}

func (a *TApplication) SetOnRestore(fn TNotifyEvent) {
    Application_SetOnRestore(a.instance, fn)
}

func (a *TApplication) Handle() HWND {
    defer exceptionProc()
    return Application_GetHandle(a.instance)
}

func (a *TApplication) SetHandle(value HWND) {
    defer exceptionProc()
    Application_SetHandle(a.instance, value)
}

func (a *TApplication) ComponentCount() int32 {
    defer exceptionProc()
    return Application_GetComponentCount(a.instance)
}

func (a *TApplication) ComponentIndex() int32 {
    defer exceptionProc()
    return Application_GetComponentIndex(a.instance)
}

func (a *TApplication) SetComponentIndex(value int32) {
    defer exceptionProc()
    Application_SetComponentIndex(a.instance, value)
}

func (a *TApplication) Owner() *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Application_GetOwner(a.instance))
}

func (a *TApplication) Name() string {
    defer exceptionProc()
    return Application_GetName(a.instance)
}

func (a *TApplication) SetName(value string) {
    defer exceptionProc()
    Application_SetName(a.instance, value)
}

func (a *TApplication) Tag() int {
    defer exceptionProc()
    return Application_GetTag(a.instance)
}

func (a *TApplication) SetTag(value int) {
    defer exceptionProc()
    Application_SetTag(a.instance, value)
}

func (a *TApplication) Components(AIndex int32) *TComponent {
    defer exceptionProc()
    return ComponentFromInst(Application_GetComponents(a.instance, AIndex))
}

