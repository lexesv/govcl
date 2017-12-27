
package api

import (
    "unsafe"
    "time"
    . "gitee.com/ying32/govcl/vcl/types"
)


//--------------------------- TApplication ---------------------------

func Application_Create(obj uintptr) uintptr {
    ret, _, _ := application_Create.Call(obj)
    return ret
}

func Application_Free(obj uintptr) {
    application_Free.Call(obj)
}

func Application_ActivateHint(obj uintptr, CursorPos TPoint)  {
    application_ActivateHint.Call(obj, uintptr(unsafe.Pointer(&CursorPos)))
}

func Application_BringToFront(obj uintptr)  {
    application_BringToFront.Call(obj)
}

func Application_CancelHint(obj uintptr)  {
    application_CancelHint.Call(obj)
}

func Application_HandleMessage(obj uintptr)  {
    application_HandleMessage.Call(obj)
}

func Application_HideHint(obj uintptr)  {
    application_HideHint.Call(obj)
}

func Application_Initialize(obj uintptr)  {
    application_Initialize.Call(obj)
}

func Application_Minimize(obj uintptr)  {
    application_Minimize.Call(obj)
}

func Application_ModalStarted(obj uintptr)  {
    application_ModalStarted.Call(obj)
}

func Application_ModalFinished(obj uintptr)  {
    application_ModalFinished.Call(obj)
}

func Application_NormalizeAllTopMosts(obj uintptr)  {
    application_NormalizeAllTopMosts.Call(obj)
}

func Application_NormalizeTopMosts(obj uintptr)  {
    application_NormalizeTopMosts.Call(obj)
}

func Application_ProcessMessages(obj uintptr)  {
    application_ProcessMessages.Call(obj)
}

func Application_Restore(obj uintptr)  {
    application_Restore.Call(obj)
}

func Application_RestoreTopMosts(obj uintptr)  {
    application_RestoreTopMosts.Call(obj)
}

func Application_Run(obj uintptr)  {
    application_Run.Call(obj)
}

func Application_Terminate(obj uintptr)  {
    application_Terminate.Call(obj)
}

func Application_MessageBox(obj uintptr, Text string, Caption string, Flags int32) int32 {
    ret, _, _ := application_MessageBox.Call(obj, GoStrToDStr(Text) , GoStrToDStr(Caption) , uintptr(Flags) )
    return int32(ret)
}

func Application_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := application_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Application_GetNamePath(obj uintptr) string {
    ret, _, _ := application_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Application_HasParent(obj uintptr) bool {
    ret, _, _ := application_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Application_Assign(obj uintptr, Source uintptr)  {
    application_Assign.Call(obj, Source )
}

func Application_ClassName(obj uintptr) string {
    ret, _, _ := application_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Application_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := application_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Application_GetHashCode(obj uintptr) int32 {
    ret, _, _ := application_GetHashCode.Call(obj)
    return int32(ret)
}

func Application_ToString(obj uintptr) string {
    ret, _, _ := application_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Application_GetDialogHandle(obj uintptr) HWND {
    ret, _, _ := application_GetDialogHandle.Call(obj)
    return HWND(ret)
}

func Application_SetDialogHandle(obj uintptr, value HWND) {
   application_SetDialogHandle.Call(obj, uintptr(value))
}

func Application_GetExeName(obj uintptr) string {
    ret, _, _ := application_GetExeName.Call(obj)
    return DStrToGoStr(ret)
}

func Application_GetHint(obj uintptr) string {
    ret, _, _ := application_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Application_SetHint(obj uintptr, value string) {
   application_SetHint.Call(obj, GoStrToDStr(value))
}

func Application_GetHintColor(obj uintptr) TColor {
    ret, _, _ := application_GetHintColor.Call(obj)
    return TColor(ret)
}

func Application_SetHintColor(obj uintptr, value TColor) {
   application_SetHintColor.Call(obj, uintptr(value))
}

func Application_GetHintHidePause(obj uintptr) int32 {
    ret, _, _ := application_GetHintHidePause.Call(obj)
    return int32(ret)
}

func Application_SetHintHidePause(obj uintptr, value int32) {
   application_SetHintHidePause.Call(obj, uintptr(value))
}

func Application_GetHintPause(obj uintptr) int32 {
    ret, _, _ := application_GetHintPause.Call(obj)
    return int32(ret)
}

func Application_SetHintPause(obj uintptr, value int32) {
   application_SetHintPause.Call(obj, uintptr(value))
}

func Application_GetHintShortCuts(obj uintptr) bool {
    ret, _, _ := application_GetHintShortCuts.Call(obj)
    return DBoolToGoBool(ret)
}

func Application_SetHintShortCuts(obj uintptr, value bool) {
   application_SetHintShortCuts.Call(obj, GoBoolToDBool(value))
}

func Application_GetHintShortPause(obj uintptr) int32 {
    ret, _, _ := application_GetHintShortPause.Call(obj)
    return int32(ret)
}

func Application_SetHintShortPause(obj uintptr, value int32) {
   application_SetHintShortPause.Call(obj, uintptr(value))
}

func Application_GetIcon(obj uintptr) uintptr {
    ret, _, _ := application_GetIcon.Call(obj)
    return ret
}

func Application_SetIcon(obj uintptr, value uintptr) {
   application_SetIcon.Call(obj, value)
}

func Application_GetIsMetropolisUI(obj uintptr) bool {
    ret, _, _ := application_GetIsMetropolisUI.Call(obj)
    return DBoolToGoBool(ret)
}

func Application_GetMainForm(obj uintptr) uintptr {
    ret, _, _ := application_GetMainForm.Call(obj)
    return ret
}

func Application_GetMainFormHandle(obj uintptr) HWND {
    ret, _, _ := application_GetMainFormHandle.Call(obj)
    return HWND(ret)
}

func Application_GetMainFormOnTaskBar(obj uintptr) bool {
    ret, _, _ := application_GetMainFormOnTaskBar.Call(obj)
    return DBoolToGoBool(ret)
}

func Application_SetMainFormOnTaskBar(obj uintptr, value bool) {
   application_SetMainFormOnTaskBar.Call(obj, GoBoolToDBool(value))
}

func Application_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := application_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func Application_SetBiDiMode(obj uintptr, value TBiDiMode) {
   application_SetBiDiMode.Call(obj, uintptr(value))
}

func Application_GetShowHint(obj uintptr) bool {
    ret, _, _ := application_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Application_SetShowHint(obj uintptr, value bool) {
   application_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Application_GetShowMainForm(obj uintptr) bool {
    ret, _, _ := application_GetShowMainForm.Call(obj)
    return DBoolToGoBool(ret)
}

func Application_SetShowMainForm(obj uintptr, value bool) {
   application_SetShowMainForm.Call(obj, GoBoolToDBool(value))
}

func Application_GetTitle(obj uintptr) string {
    ret, _, _ := application_GetTitle.Call(obj)
    return DStrToGoStr(ret)
}

func Application_SetTitle(obj uintptr, value string) {
   application_SetTitle.Call(obj, GoStrToDStr(value))
}

func Application_SetOnException(obj uintptr, fn interface{}) {
    application_SetOnException.Call(obj, addEventToMap(fn))
}

func Application_SetOnMinimize(obj uintptr, fn interface{}) {
    application_SetOnMinimize.Call(obj, addEventToMap(fn))
}

func Application_SetOnRestore(obj uintptr, fn interface{}) {
    application_SetOnRestore.Call(obj, addEventToMap(fn))
}

func Application_GetHandle(obj uintptr) HWND {
    ret, _, _ := application_GetHandle.Call(obj)
    return HWND(ret)
}

func Application_SetHandle(obj uintptr, value HWND) {
   application_SetHandle.Call(obj, uintptr(value))
}

func Application_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := application_GetComponentCount.Call(obj)
    return int32(ret)
}

func Application_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := application_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Application_SetComponentIndex(obj uintptr, value int32) {
   application_SetComponentIndex.Call(obj, uintptr(value))
}

func Application_GetOwner(obj uintptr) uintptr {
    ret, _, _ := application_GetOwner.Call(obj)
    return ret
}

func Application_GetName(obj uintptr) string {
    ret, _, _ := application_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Application_SetName(obj uintptr, value string) {
   application_SetName.Call(obj, GoStrToDStr(value))
}

func Application_GetTag(obj uintptr) int {
    ret, _, _ := application_GetTag.Call(obj)
    return int(ret)
}

func Application_SetTag(obj uintptr, value int) {
   application_SetTag.Call(obj, uintptr(value))
}

func Application_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := application_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TForm ---------------------------

func Form_Create(obj uintptr) uintptr {
    ret, _, _ := form_Create.Call(obj)
    return ret
}

func Form_Free(obj uintptr) {
    form_Free.Call(obj)
}

func Form_Close(obj uintptr)  {
    form_Close.Call(obj)
}

func Form_Hide(obj uintptr)  {
    form_Hide.Call(obj)
}

func Form_Print(obj uintptr)  {
    form_Print.Call(obj)
}

func Form_SetFocus(obj uintptr)  {
    form_SetFocus.Call(obj)
}

func Form_Show(obj uintptr)  {
    form_Show.Call(obj)
}

func Form_ShowModal(obj uintptr) int32 {
    ret, _, _ := form_ShowModal.Call(obj)
    return int32(ret)
}

func Form_CanFocus(obj uintptr) bool {
    ret, _, _ := form_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_FlipChildren(obj uintptr, AllLevels bool)  {
    form_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func Form_Focused(obj uintptr) bool {
    ret, _, _ := form_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_HandleAllocated(obj uintptr) bool {
    ret, _, _ := form_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_Invalidate(obj uintptr)  {
    form_Invalidate.Call(obj)
}

func Form_Realign(obj uintptr)  {
    form_Realign.Call(obj)
}

func Form_Repaint(obj uintptr)  {
    form_Repaint.Call(obj)
}

func Form_ScaleBy(obj uintptr, M int32, D int32)  {
    form_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func Form_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    form_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func Form_Update(obj uintptr)  {
    form_Update.Call(obj)
}

func Form_BringToFront(obj uintptr)  {
    form_BringToFront.Call(obj)
}

func Form_HasParent(obj uintptr) bool {
    ret, _, _ := form_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := form_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func Form_Refresh(obj uintptr)  {
    form_Refresh.Call(obj)
}

func Form_SendToBack(obj uintptr)  {
    form_SendToBack.Call(obj)
}

func Form_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := form_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Form_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := form_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Form_GetNamePath(obj uintptr) string {
    ret, _, _ := form_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Form_Assign(obj uintptr, Source uintptr)  {
    form_Assign.Call(obj, Source )
}

func Form_ClassName(obj uintptr) string {
    ret, _, _ := form_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Form_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := form_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Form_GetHashCode(obj uintptr) int32 {
    ret, _, _ := form_GetHashCode.Call(obj)
    return int32(ret)
}

func Form_ToString(obj uintptr) string {
    ret, _, _ := form_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Form_GetAction(obj uintptr) uintptr {
    ret, _, _ := form_GetAction.Call(obj)
    return ret
}

func Form_SetAction(obj uintptr, value uintptr) {
   form_SetAction.Call(obj, value)
}

func Form_GetAlign(obj uintptr) TAlign {
    ret, _, _ := form_GetAlign.Call(obj)
    return TAlign(ret)
}

func Form_SetAlign(obj uintptr, value TAlign) {
   form_SetAlign.Call(obj, uintptr(value))
}

func Form_GetAlphaBlend(obj uintptr) bool {
    ret, _, _ := form_GetAlphaBlend.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetAlphaBlend(obj uintptr, value bool) {
   form_SetAlphaBlend.Call(obj, GoBoolToDBool(value))
}

func Form_GetAlphaBlendValue(obj uintptr) uint8 {
    ret, _, _ := form_GetAlphaBlendValue.Call(obj)
    return uint8(ret)
}

func Form_SetAlphaBlendValue(obj uintptr, value uint8) {
   form_SetAlphaBlendValue.Call(obj, uintptr(value))
}

func Form_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := form_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func Form_SetAnchors(obj uintptr, value TAnchors) {
   form_SetAnchors.Call(obj, uintptr(value))
}

func Form_GetAutoSize(obj uintptr) bool {
    ret, _, _ := form_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetAutoSize(obj uintptr, value bool) {
   form_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func Form_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := form_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func Form_SetBiDiMode(obj uintptr, value TBiDiMode) {
   form_SetBiDiMode.Call(obj, uintptr(value))
}

func Form_GetBorderIcons(obj uintptr) TBorderIcons {
    ret, _, _ := form_GetBorderIcons.Call(obj)
    return TBorderIcons(ret)
}

func Form_SetBorderIcons(obj uintptr, value TBorderIcons) {
   form_SetBorderIcons.Call(obj, uintptr(value))
}

func Form_GetBorderStyle(obj uintptr) TFormBorderStyle {
    ret, _, _ := form_GetBorderStyle.Call(obj)
    return TFormBorderStyle(ret)
}

func Form_SetBorderStyle(obj uintptr, value TFormBorderStyle) {
   form_SetBorderStyle.Call(obj, uintptr(value))
}

func Form_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := form_GetBorderWidth.Call(obj)
    return int32(ret)
}

func Form_SetBorderWidth(obj uintptr, value int32) {
   form_SetBorderWidth.Call(obj, uintptr(value))
}

func Form_GetCaption(obj uintptr) string {
    ret, _, _ := form_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func Form_SetCaption(obj uintptr, value string) {
   form_SetCaption.Call(obj, GoStrToDStr(value))
}

func Form_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := form_GetClientHeight.Call(obj)
    return int32(ret)
}

func Form_SetClientHeight(obj uintptr, value int32) {
   form_SetClientHeight.Call(obj, uintptr(value))
}

func Form_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := form_GetClientWidth.Call(obj)
    return int32(ret)
}

func Form_SetClientWidth(obj uintptr, value int32) {
   form_SetClientWidth.Call(obj, uintptr(value))
}

func Form_GetColor(obj uintptr) TColor {
    ret, _, _ := form_GetColor.Call(obj)
    return TColor(ret)
}

func Form_SetColor(obj uintptr, value TColor) {
   form_SetColor.Call(obj, uintptr(value))
}

func Form_GetTransparentColor(obj uintptr) bool {
    ret, _, _ := form_GetTransparentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetTransparentColor(obj uintptr, value bool) {
   form_SetTransparentColor.Call(obj, GoBoolToDBool(value))
}

func Form_GetTransparentColorValue(obj uintptr) TColor {
    ret, _, _ := form_GetTransparentColorValue.Call(obj)
    return TColor(ret)
}

func Form_SetTransparentColorValue(obj uintptr, value TColor) {
   form_SetTransparentColorValue.Call(obj, uintptr(value))
}

func Form_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := form_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetDoubleBuffered(obj uintptr, value bool) {
   form_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func Form_GetEnabled(obj uintptr) bool {
    ret, _, _ := form_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetEnabled(obj uintptr, value bool) {
   form_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Form_GetParentFont(obj uintptr) bool {
    ret, _, _ := form_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetParentFont(obj uintptr, value bool) {
   form_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func Form_GetFont(obj uintptr) uintptr {
    ret, _, _ := form_GetFont.Call(obj)
    return ret
}

func Form_SetFont(obj uintptr, value uintptr) {
   form_SetFont.Call(obj, value)
}

func Form_GetFormStyle(obj uintptr) TFormStyle {
    ret, _, _ := form_GetFormStyle.Call(obj)
    return TFormStyle(ret)
}

func Form_SetFormStyle(obj uintptr, value TFormStyle) {
   form_SetFormStyle.Call(obj, uintptr(value))
}

func Form_GetHeight(obj uintptr) int32 {
    ret, _, _ := form_GetHeight.Call(obj)
    return int32(ret)
}

func Form_SetHeight(obj uintptr, value int32) {
   form_SetHeight.Call(obj, uintptr(value))
}

func Form_GetIcon(obj uintptr) uintptr {
    ret, _, _ := form_GetIcon.Call(obj)
    return ret
}

func Form_SetIcon(obj uintptr, value uintptr) {
   form_SetIcon.Call(obj, value)
}

func Form_GetKeyPreview(obj uintptr) bool {
    ret, _, _ := form_GetKeyPreview.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetKeyPreview(obj uintptr, value bool) {
   form_SetKeyPreview.Call(obj, GoBoolToDBool(value))
}

func Form_GetMenu(obj uintptr) uintptr {
    ret, _, _ := form_GetMenu.Call(obj)
    return ret
}

func Form_SetMenu(obj uintptr, value uintptr) {
   form_SetMenu.Call(obj, value)
}

func Form_GetPixelsPerInch(obj uintptr) int32 {
    ret, _, _ := form_GetPixelsPerInch.Call(obj)
    return int32(ret)
}

func Form_SetPixelsPerInch(obj uintptr, value int32) {
   form_SetPixelsPerInch.Call(obj, uintptr(value))
}

func Form_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := form_GetPopupMenu.Call(obj)
    return ret
}

func Form_SetPopupMenu(obj uintptr, value uintptr) {
   form_SetPopupMenu.Call(obj, value)
}

func Form_GetPosition(obj uintptr) TPosition {
    ret, _, _ := form_GetPosition.Call(obj)
    return TPosition(ret)
}

func Form_SetPosition(obj uintptr, value TPosition) {
   form_SetPosition.Call(obj, uintptr(value))
}

func Form_GetScaled(obj uintptr) bool {
    ret, _, _ := form_GetScaled.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetScaled(obj uintptr, value bool) {
   form_SetScaled.Call(obj, GoBoolToDBool(value))
}

func Form_GetShowHint(obj uintptr) bool {
    ret, _, _ := form_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetShowHint(obj uintptr, value bool) {
   form_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Form_GetVisible(obj uintptr) bool {
    ret, _, _ := form_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetVisible(obj uintptr, value bool) {
   form_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Form_GetWidth(obj uintptr) int32 {
    ret, _, _ := form_GetWidth.Call(obj)
    return int32(ret)
}

func Form_SetWidth(obj uintptr, value int32) {
   form_SetWidth.Call(obj, uintptr(value))
}

func Form_GetWindowState(obj uintptr) TWindowState {
    ret, _, _ := form_GetWindowState.Call(obj)
    return TWindowState(ret)
}

func Form_SetWindowState(obj uintptr, value TWindowState) {
   form_SetWindowState.Call(obj, uintptr(value))
}

func Form_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := form_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func Form_SetStyleElements(obj uintptr, value TStyleElements) {
   form_SetStyleElements.Call(obj, uintptr(value))
}

func Form_SetOnClick(obj uintptr, fn interface{}) {
    form_SetOnClick.Call(obj, addEventToMap(fn))
}

func Form_SetOnClose(obj uintptr, fn interface{}) {
    form_SetOnClose.Call(obj, addEventToMap(fn))
}

func Form_SetOnCloseQuery(obj uintptr, fn interface{}) {
    form_SetOnCloseQuery.Call(obj, addEventToMap(fn))
}

func Form_SetOnDblClick(obj uintptr, fn interface{}) {
    form_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func Form_SetOnHide(obj uintptr, fn interface{}) {
    form_SetOnHide.Call(obj, addEventToMap(fn))
}

func Form_SetOnKeyDown(obj uintptr, fn interface{}) {
    form_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func Form_SetOnKeyPress(obj uintptr, fn interface{}) {
    form_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func Form_SetOnKeyUp(obj uintptr, fn interface{}) {
    form_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func Form_SetOnMouseDown(obj uintptr, fn interface{}) {
    form_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func Form_SetOnMouseEnter(obj uintptr, fn interface{}) {
    form_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func Form_SetOnMouseLeave(obj uintptr, fn interface{}) {
    form_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func Form_SetOnMouseMove(obj uintptr, fn interface{}) {
    form_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func Form_SetOnMouseUp(obj uintptr, fn interface{}) {
    form_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func Form_SetOnMouseWheel(obj uintptr, fn interface{}) {
    form_SetOnMouseWheel.Call(obj, addEventToMap(fn))
}

func Form_SetOnPaint(obj uintptr, fn interface{}) {
    form_SetOnPaint.Call(obj, addEventToMap(fn))
}

func Form_SetOnResize(obj uintptr, fn interface{}) {
    form_SetOnResize.Call(obj, addEventToMap(fn))
}

func Form_SetOnShow(obj uintptr, fn interface{}) {
    form_SetOnShow.Call(obj, addEventToMap(fn))
}

func Form_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := form_GetCanvas.Call(obj)
    return ret
}

func Form_GetDropTarget(obj uintptr) bool {
    ret, _, _ := form_GetDropTarget.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetDropTarget(obj uintptr, value bool) {
   form_SetDropTarget.Call(obj, GoBoolToDBool(value))
}

func Form_GetModalResult(obj uintptr) TModalResult {
    ret, _, _ := form_GetModalResult.Call(obj)
    return TModalResult(ret)
}

func Form_SetModalResult(obj uintptr, value TModalResult) {
   form_SetModalResult.Call(obj, uintptr(value))
}

func Form_GetLeft(obj uintptr) int32 {
    ret, _, _ := form_GetLeft.Call(obj)
    return int32(ret)
}

func Form_SetLeft(obj uintptr, value int32) {
   form_SetLeft.Call(obj, uintptr(value))
}

func Form_GetTop(obj uintptr) int32 {
    ret, _, _ := form_GetTop.Call(obj)
    return int32(ret)
}

func Form_SetTop(obj uintptr, value int32) {
   form_SetTop.Call(obj, uintptr(value))
}

func Form_GetBrush(obj uintptr) uintptr {
    ret, _, _ := form_GetBrush.Call(obj)
    return ret
}

func Form_GetControlCount(obj uintptr) int32 {
    ret, _, _ := form_GetControlCount.Call(obj)
    return int32(ret)
}

func Form_GetHandle(obj uintptr) HWND {
    ret, _, _ := form_GetHandle.Call(obj)
    return HWND(ret)
}

func Form_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := form_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetParentDoubleBuffered(obj uintptr, value bool) {
   form_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func Form_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := form_GetParentWindow.Call(obj)
    return HWND(ret)
}

func Form_SetParentWindow(obj uintptr, value HWND) {
   form_SetParentWindow.Call(obj, uintptr(value))
}

func Form_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := form_GetTabOrder.Call(obj)
    return uint16(ret)
}

func Form_SetTabOrder(obj uintptr, value uint16) {
   form_SetTabOrder.Call(obj, uintptr(value))
}

func Form_GetTabStop(obj uintptr) bool {
    ret, _, _ := form_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetTabStop(obj uintptr, value bool) {
   form_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func Form_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    form_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Form_SetBoundsRect(obj uintptr, value TRect) {
   form_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Form_GetClientRect(obj uintptr) TRect {
    var ret TRect
    form_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Form_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := form_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Form_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := form_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Form_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := form_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Form_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := form_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Form_GetParent(obj uintptr) uintptr {
    ret, _, _ := form_GetParent.Call(obj)
    return ret
}

func Form_SetParent(obj uintptr, value uintptr) {
   form_SetParent.Call(obj, value)
}

func Form_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := form_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func Form_SetAlignWithMargins(obj uintptr, value bool) {
   form_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func Form_GetCursor(obj uintptr) TCursor {
    ret, _, _ := form_GetCursor.Call(obj)
    return TCursor(ret)
}

func Form_SetCursor(obj uintptr, value TCursor) {
   form_SetCursor.Call(obj, uintptr(value))
}

func Form_GetHint(obj uintptr) string {
    ret, _, _ := form_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Form_SetHint(obj uintptr, value string) {
   form_SetHint.Call(obj, GoStrToDStr(value))
}

func Form_GetMargins(obj uintptr) uintptr {
    ret, _, _ := form_GetMargins.Call(obj)
    return ret
}

func Form_SetMargins(obj uintptr, value uintptr) {
   form_SetMargins.Call(obj, value)
}

func Form_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := form_GetComponentCount.Call(obj)
    return int32(ret)
}

func Form_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := form_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Form_SetComponentIndex(obj uintptr, value int32) {
   form_SetComponentIndex.Call(obj, uintptr(value))
}

func Form_GetOwner(obj uintptr) uintptr {
    ret, _, _ := form_GetOwner.Call(obj)
    return ret
}

func Form_GetName(obj uintptr) string {
    ret, _, _ := form_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Form_SetName(obj uintptr, value string) {
   form_SetName.Call(obj, GoStrToDStr(value))
}

func Form_GetTag(obj uintptr) int {
    ret, _, _ := form_GetTag.Call(obj)
    return int(ret)
}

func Form_SetTag(obj uintptr, value int) {
   form_SetTag.Call(obj, uintptr(value))
}

func Form_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := form_GetControls.Call(obj, uintptr(Index))
    return ret
}

func Form_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := form_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TButton ---------------------------

func Button_Create(obj uintptr) uintptr {
    ret, _, _ := button_Create.Call(obj)
    return ret
}

func Button_Free(obj uintptr) {
    button_Free.Call(obj)
}

func Button_Click(obj uintptr)  {
    button_Click.Call(obj)
}

func Button_CanFocus(obj uintptr) bool {
    ret, _, _ := button_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_FlipChildren(obj uintptr, AllLevels bool)  {
    button_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func Button_Focused(obj uintptr) bool {
    ret, _, _ := button_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_HandleAllocated(obj uintptr) bool {
    ret, _, _ := button_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_Invalidate(obj uintptr)  {
    button_Invalidate.Call(obj)
}

func Button_Realign(obj uintptr)  {
    button_Realign.Call(obj)
}

func Button_Repaint(obj uintptr)  {
    button_Repaint.Call(obj)
}

func Button_ScaleBy(obj uintptr, M int32, D int32)  {
    button_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func Button_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    button_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func Button_SetFocus(obj uintptr)  {
    button_SetFocus.Call(obj)
}

func Button_Update(obj uintptr)  {
    button_Update.Call(obj)
}

func Button_BringToFront(obj uintptr)  {
    button_BringToFront.Call(obj)
}

func Button_HasParent(obj uintptr) bool {
    ret, _, _ := button_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_Hide(obj uintptr)  {
    button_Hide.Call(obj)
}

func Button_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := button_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func Button_Refresh(obj uintptr)  {
    button_Refresh.Call(obj)
}

func Button_SendToBack(obj uintptr)  {
    button_SendToBack.Call(obj)
}

func Button_Show(obj uintptr)  {
    button_Show.Call(obj)
}

func Button_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := button_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Button_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := button_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Button_GetNamePath(obj uintptr) string {
    ret, _, _ := button_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Button_Assign(obj uintptr, Source uintptr)  {
    button_Assign.Call(obj, Source )
}

func Button_ClassName(obj uintptr) string {
    ret, _, _ := button_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Button_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := button_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Button_GetHashCode(obj uintptr) int32 {
    ret, _, _ := button_GetHashCode.Call(obj)
    return int32(ret)
}

func Button_ToString(obj uintptr) string {
    ret, _, _ := button_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Button_GetAction(obj uintptr) uintptr {
    ret, _, _ := button_GetAction.Call(obj)
    return ret
}

func Button_SetAction(obj uintptr, value uintptr) {
   button_SetAction.Call(obj, value)
}

func Button_GetAlign(obj uintptr) TAlign {
    ret, _, _ := button_GetAlign.Call(obj)
    return TAlign(ret)
}

func Button_SetAlign(obj uintptr, value TAlign) {
   button_SetAlign.Call(obj, uintptr(value))
}

func Button_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := button_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func Button_SetAnchors(obj uintptr, value TAnchors) {
   button_SetAnchors.Call(obj, uintptr(value))
}

func Button_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := button_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func Button_SetBiDiMode(obj uintptr, value TBiDiMode) {
   button_SetBiDiMode.Call(obj, uintptr(value))
}

func Button_GetCancel(obj uintptr) bool {
    ret, _, _ := button_GetCancel.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetCancel(obj uintptr, value bool) {
   button_SetCancel.Call(obj, GoBoolToDBool(value))
}

func Button_GetCaption(obj uintptr) string {
    ret, _, _ := button_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func Button_SetCaption(obj uintptr, value string) {
   button_SetCaption.Call(obj, GoStrToDStr(value))
}

func Button_GetCommandLinkHint(obj uintptr) string {
    ret, _, _ := button_GetCommandLinkHint.Call(obj)
    return DStrToGoStr(ret)
}

func Button_SetCommandLinkHint(obj uintptr, value string) {
   button_SetCommandLinkHint.Call(obj, GoStrToDStr(value))
}

func Button_GetDefault(obj uintptr) bool {
    ret, _, _ := button_GetDefault.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetDefault(obj uintptr, value bool) {
   button_SetDefault.Call(obj, GoBoolToDBool(value))
}

func Button_GetDisabledImageIndex(obj uintptr) int32 {
    ret, _, _ := button_GetDisabledImageIndex.Call(obj)
    return int32(ret)
}

func Button_SetDisabledImageIndex(obj uintptr, value int32) {
   button_SetDisabledImageIndex.Call(obj, uintptr(value))
}

func Button_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := button_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetDoubleBuffered(obj uintptr, value bool) {
   button_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func Button_GetElevationRequired(obj uintptr) bool {
    ret, _, _ := button_GetElevationRequired.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetElevationRequired(obj uintptr, value bool) {
   button_SetElevationRequired.Call(obj, GoBoolToDBool(value))
}

func Button_GetEnabled(obj uintptr) bool {
    ret, _, _ := button_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetEnabled(obj uintptr, value bool) {
   button_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Button_GetFont(obj uintptr) uintptr {
    ret, _, _ := button_GetFont.Call(obj)
    return ret
}

func Button_SetFont(obj uintptr, value uintptr) {
   button_SetFont.Call(obj, value)
}

func Button_GetHotImageIndex(obj uintptr) int32 {
    ret, _, _ := button_GetHotImageIndex.Call(obj)
    return int32(ret)
}

func Button_SetHotImageIndex(obj uintptr, value int32) {
   button_SetHotImageIndex.Call(obj, uintptr(value))
}

func Button_GetImageAlignment(obj uintptr) TImageAlignment {
    ret, _, _ := button_GetImageAlignment.Call(obj)
    return TImageAlignment(ret)
}

func Button_SetImageAlignment(obj uintptr, value TImageAlignment) {
   button_SetImageAlignment.Call(obj, uintptr(value))
}

func Button_GetImageIndex(obj uintptr) int32 {
    ret, _, _ := button_GetImageIndex.Call(obj)
    return int32(ret)
}

func Button_SetImageIndex(obj uintptr, value int32) {
   button_SetImageIndex.Call(obj, uintptr(value))
}

func Button_GetImages(obj uintptr) uintptr {
    ret, _, _ := button_GetImages.Call(obj)
    return ret
}

func Button_SetImages(obj uintptr, value uintptr) {
   button_SetImages.Call(obj, value)
}

func Button_GetModalResult(obj uintptr) TModalResult {
    ret, _, _ := button_GetModalResult.Call(obj)
    return TModalResult(ret)
}

func Button_SetModalResult(obj uintptr, value TModalResult) {
   button_SetModalResult.Call(obj, uintptr(value))
}

func Button_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := button_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetParentDoubleBuffered(obj uintptr, value bool) {
   button_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func Button_GetParentFont(obj uintptr) bool {
    ret, _, _ := button_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetParentFont(obj uintptr, value bool) {
   button_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func Button_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := button_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetParentShowHint(obj uintptr, value bool) {
   button_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func Button_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := button_GetPopupMenu.Call(obj)
    return ret
}

func Button_SetPopupMenu(obj uintptr, value uintptr) {
   button_SetPopupMenu.Call(obj, value)
}

func Button_GetPressedImageIndex(obj uintptr) int32 {
    ret, _, _ := button_GetPressedImageIndex.Call(obj)
    return int32(ret)
}

func Button_SetPressedImageIndex(obj uintptr, value int32) {
   button_SetPressedImageIndex.Call(obj, uintptr(value))
}

func Button_GetSelectedImageIndex(obj uintptr) int32 {
    ret, _, _ := button_GetSelectedImageIndex.Call(obj)
    return int32(ret)
}

func Button_SetSelectedImageIndex(obj uintptr, value int32) {
   button_SetSelectedImageIndex.Call(obj, uintptr(value))
}

func Button_GetShowHint(obj uintptr) bool {
    ret, _, _ := button_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetShowHint(obj uintptr, value bool) {
   button_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Button_GetStyle(obj uintptr) TButtonStyle {
    ret, _, _ := button_GetStyle.Call(obj)
    return TButtonStyle(ret)
}

func Button_SetStyle(obj uintptr, value TButtonStyle) {
   button_SetStyle.Call(obj, uintptr(value))
}

func Button_GetStylusHotImageIndex(obj uintptr) int32 {
    ret, _, _ := button_GetStylusHotImageIndex.Call(obj)
    return int32(ret)
}

func Button_SetStylusHotImageIndex(obj uintptr, value int32) {
   button_SetStylusHotImageIndex.Call(obj, uintptr(value))
}

func Button_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := button_GetTabOrder.Call(obj)
    return uint16(ret)
}

func Button_SetTabOrder(obj uintptr, value uint16) {
   button_SetTabOrder.Call(obj, uintptr(value))
}

func Button_GetTabStop(obj uintptr) bool {
    ret, _, _ := button_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetTabStop(obj uintptr, value bool) {
   button_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func Button_GetVisible(obj uintptr) bool {
    ret, _, _ := button_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetVisible(obj uintptr, value bool) {
   button_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Button_GetWordWrap(obj uintptr) bool {
    ret, _, _ := button_GetWordWrap.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetWordWrap(obj uintptr, value bool) {
   button_SetWordWrap.Call(obj, GoBoolToDBool(value))
}

func Button_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := button_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func Button_SetStyleElements(obj uintptr, value TStyleElements) {
   button_SetStyleElements.Call(obj, uintptr(value))
}

func Button_SetOnClick(obj uintptr, fn interface{}) {
    button_SetOnClick.Call(obj, addEventToMap(fn))
}

func Button_SetOnEnter(obj uintptr, fn interface{}) {
    button_SetOnEnter.Call(obj, addEventToMap(fn))
}

func Button_SetOnExit(obj uintptr, fn interface{}) {
    button_SetOnExit.Call(obj, addEventToMap(fn))
}

func Button_SetOnKeyDown(obj uintptr, fn interface{}) {
    button_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func Button_SetOnKeyPress(obj uintptr, fn interface{}) {
    button_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func Button_SetOnKeyUp(obj uintptr, fn interface{}) {
    button_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func Button_SetOnMouseDown(obj uintptr, fn interface{}) {
    button_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func Button_SetOnMouseEnter(obj uintptr, fn interface{}) {
    button_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func Button_SetOnMouseLeave(obj uintptr, fn interface{}) {
    button_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func Button_SetOnMouseMove(obj uintptr, fn interface{}) {
    button_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func Button_SetOnMouseUp(obj uintptr, fn interface{}) {
    button_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func Button_GetBrush(obj uintptr) uintptr {
    ret, _, _ := button_GetBrush.Call(obj)
    return ret
}

func Button_GetControlCount(obj uintptr) int32 {
    ret, _, _ := button_GetControlCount.Call(obj)
    return int32(ret)
}

func Button_GetHandle(obj uintptr) HWND {
    ret, _, _ := button_GetHandle.Call(obj)
    return HWND(ret)
}

func Button_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := button_GetParentWindow.Call(obj)
    return HWND(ret)
}

func Button_SetParentWindow(obj uintptr, value HWND) {
   button_SetParentWindow.Call(obj, uintptr(value))
}

func Button_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    button_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Button_SetBoundsRect(obj uintptr, value TRect) {
   button_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Button_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := button_GetClientHeight.Call(obj)
    return int32(ret)
}

func Button_SetClientHeight(obj uintptr, value int32) {
   button_SetClientHeight.Call(obj, uintptr(value))
}

func Button_GetClientRect(obj uintptr) TRect {
    var ret TRect
    button_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Button_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := button_GetClientWidth.Call(obj)
    return int32(ret)
}

func Button_SetClientWidth(obj uintptr, value int32) {
   button_SetClientWidth.Call(obj, uintptr(value))
}

func Button_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := button_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Button_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := button_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Button_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := button_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Button_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := button_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Button_GetParent(obj uintptr) uintptr {
    ret, _, _ := button_GetParent.Call(obj)
    return ret
}

func Button_SetParent(obj uintptr, value uintptr) {
   button_SetParent.Call(obj, value)
}

func Button_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := button_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func Button_SetAlignWithMargins(obj uintptr, value bool) {
   button_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func Button_GetLeft(obj uintptr) int32 {
    ret, _, _ := button_GetLeft.Call(obj)
    return int32(ret)
}

func Button_SetLeft(obj uintptr, value int32) {
   button_SetLeft.Call(obj, uintptr(value))
}

func Button_GetTop(obj uintptr) int32 {
    ret, _, _ := button_GetTop.Call(obj)
    return int32(ret)
}

func Button_SetTop(obj uintptr, value int32) {
   button_SetTop.Call(obj, uintptr(value))
}

func Button_GetWidth(obj uintptr) int32 {
    ret, _, _ := button_GetWidth.Call(obj)
    return int32(ret)
}

func Button_SetWidth(obj uintptr, value int32) {
   button_SetWidth.Call(obj, uintptr(value))
}

func Button_GetHeight(obj uintptr) int32 {
    ret, _, _ := button_GetHeight.Call(obj)
    return int32(ret)
}

func Button_SetHeight(obj uintptr, value int32) {
   button_SetHeight.Call(obj, uintptr(value))
}

func Button_GetCursor(obj uintptr) TCursor {
    ret, _, _ := button_GetCursor.Call(obj)
    return TCursor(ret)
}

func Button_SetCursor(obj uintptr, value TCursor) {
   button_SetCursor.Call(obj, uintptr(value))
}

func Button_GetHint(obj uintptr) string {
    ret, _, _ := button_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Button_SetHint(obj uintptr, value string) {
   button_SetHint.Call(obj, GoStrToDStr(value))
}

func Button_GetMargins(obj uintptr) uintptr {
    ret, _, _ := button_GetMargins.Call(obj)
    return ret
}

func Button_SetMargins(obj uintptr, value uintptr) {
   button_SetMargins.Call(obj, value)
}

func Button_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := button_GetComponentCount.Call(obj)
    return int32(ret)
}

func Button_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := button_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Button_SetComponentIndex(obj uintptr, value int32) {
   button_SetComponentIndex.Call(obj, uintptr(value))
}

func Button_GetOwner(obj uintptr) uintptr {
    ret, _, _ := button_GetOwner.Call(obj)
    return ret
}

func Button_GetName(obj uintptr) string {
    ret, _, _ := button_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Button_SetName(obj uintptr, value string) {
   button_SetName.Call(obj, GoStrToDStr(value))
}

func Button_GetTag(obj uintptr) int {
    ret, _, _ := button_GetTag.Call(obj)
    return int(ret)
}

func Button_SetTag(obj uintptr, value int) {
   button_SetTag.Call(obj, uintptr(value))
}

func Button_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := button_GetControls.Call(obj, uintptr(Index))
    return ret
}

func Button_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := button_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TEdit ---------------------------

func Edit_Create(obj uintptr) uintptr {
    ret, _, _ := edit_Create.Call(obj)
    return ret
}

func Edit_Free(obj uintptr) {
    edit_Free.Call(obj)
}

func Edit_Clear(obj uintptr)  {
    edit_Clear.Call(obj)
}

func Edit_ClearSelection(obj uintptr)  {
    edit_ClearSelection.Call(obj)
}

func Edit_CopyToClipboard(obj uintptr)  {
    edit_CopyToClipboard.Call(obj)
}

func Edit_CutToClipboard(obj uintptr)  {
    edit_CutToClipboard.Call(obj)
}

func Edit_PasteFromClipboard(obj uintptr)  {
    edit_PasteFromClipboard.Call(obj)
}

func Edit_SelectAll(obj uintptr)  {
    edit_SelectAll.Call(obj)
}

func Edit_GetSelTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := edit_GetSelTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Edit_CanFocus(obj uintptr) bool {
    ret, _, _ := edit_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_FlipChildren(obj uintptr, AllLevels bool)  {
    edit_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func Edit_Focused(obj uintptr) bool {
    ret, _, _ := edit_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_HandleAllocated(obj uintptr) bool {
    ret, _, _ := edit_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_Invalidate(obj uintptr)  {
    edit_Invalidate.Call(obj)
}

func Edit_Realign(obj uintptr)  {
    edit_Realign.Call(obj)
}

func Edit_Repaint(obj uintptr)  {
    edit_Repaint.Call(obj)
}

func Edit_ScaleBy(obj uintptr, M int32, D int32)  {
    edit_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func Edit_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    edit_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func Edit_SetFocus(obj uintptr)  {
    edit_SetFocus.Call(obj)
}

func Edit_Update(obj uintptr)  {
    edit_Update.Call(obj)
}

func Edit_BringToFront(obj uintptr)  {
    edit_BringToFront.Call(obj)
}

func Edit_HasParent(obj uintptr) bool {
    ret, _, _ := edit_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_Hide(obj uintptr)  {
    edit_Hide.Call(obj)
}

func Edit_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := edit_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func Edit_Refresh(obj uintptr)  {
    edit_Refresh.Call(obj)
}

func Edit_SendToBack(obj uintptr)  {
    edit_SendToBack.Call(obj)
}

func Edit_Show(obj uintptr)  {
    edit_Show.Call(obj)
}

func Edit_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := edit_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Edit_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := edit_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Edit_GetNamePath(obj uintptr) string {
    ret, _, _ := edit_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Edit_Assign(obj uintptr, Source uintptr)  {
    edit_Assign.Call(obj, Source )
}

func Edit_ClassName(obj uintptr) string {
    ret, _, _ := edit_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Edit_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := edit_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Edit_GetHashCode(obj uintptr) int32 {
    ret, _, _ := edit_GetHashCode.Call(obj)
    return int32(ret)
}

func Edit_ToString(obj uintptr) string {
    ret, _, _ := edit_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Edit_GetAlign(obj uintptr) TAlign {
    ret, _, _ := edit_GetAlign.Call(obj)
    return TAlign(ret)
}

func Edit_SetAlign(obj uintptr, value TAlign) {
   edit_SetAlign.Call(obj, uintptr(value))
}

func Edit_GetAlignment(obj uintptr) TAlignment {
    ret, _, _ := edit_GetAlignment.Call(obj)
    return TAlignment(ret)
}

func Edit_SetAlignment(obj uintptr, value TAlignment) {
   edit_SetAlignment.Call(obj, uintptr(value))
}

func Edit_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := edit_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func Edit_SetAnchors(obj uintptr, value TAnchors) {
   edit_SetAnchors.Call(obj, uintptr(value))
}

func Edit_GetAutoSelect(obj uintptr) bool {
    ret, _, _ := edit_GetAutoSelect.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetAutoSelect(obj uintptr, value bool) {
   edit_SetAutoSelect.Call(obj, GoBoolToDBool(value))
}

func Edit_GetAutoSize(obj uintptr) bool {
    ret, _, _ := edit_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetAutoSize(obj uintptr, value bool) {
   edit_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func Edit_GetBevelEdges(obj uintptr) TBevelEdges {
    ret, _, _ := edit_GetBevelEdges.Call(obj)
    return TBevelEdges(ret)
}

func Edit_SetBevelEdges(obj uintptr, value TBevelEdges) {
   edit_SetBevelEdges.Call(obj, uintptr(value))
}

func Edit_GetBevelInner(obj uintptr) TBevelCut {
    ret, _, _ := edit_GetBevelInner.Call(obj)
    return TBevelCut(ret)
}

func Edit_SetBevelInner(obj uintptr, value TBevelCut) {
   edit_SetBevelInner.Call(obj, uintptr(value))
}

func Edit_GetBevelKind(obj uintptr) TBevelKind {
    ret, _, _ := edit_GetBevelKind.Call(obj)
    return TBevelKind(ret)
}

func Edit_SetBevelKind(obj uintptr, value TBevelKind) {
   edit_SetBevelKind.Call(obj, uintptr(value))
}

func Edit_GetBevelOuter(obj uintptr) TBevelCut {
    ret, _, _ := edit_GetBevelOuter.Call(obj)
    return TBevelCut(ret)
}

func Edit_SetBevelOuter(obj uintptr, value TBevelCut) {
   edit_SetBevelOuter.Call(obj, uintptr(value))
}

func Edit_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := edit_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func Edit_SetBiDiMode(obj uintptr, value TBiDiMode) {
   edit_SetBiDiMode.Call(obj, uintptr(value))
}

func Edit_GetBorderStyle(obj uintptr) TBorderStyle {
    ret, _, _ := edit_GetBorderStyle.Call(obj)
    return TBorderStyle(ret)
}

func Edit_SetBorderStyle(obj uintptr, value TBorderStyle) {
   edit_SetBorderStyle.Call(obj, uintptr(value))
}

func Edit_GetColor(obj uintptr) TColor {
    ret, _, _ := edit_GetColor.Call(obj)
    return TColor(ret)
}

func Edit_SetColor(obj uintptr, value TColor) {
   edit_SetColor.Call(obj, uintptr(value))
}

func Edit_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := edit_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetDoubleBuffered(obj uintptr, value bool) {
   edit_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func Edit_GetEnabled(obj uintptr) bool {
    ret, _, _ := edit_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetEnabled(obj uintptr, value bool) {
   edit_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Edit_GetFont(obj uintptr) uintptr {
    ret, _, _ := edit_GetFont.Call(obj)
    return ret
}

func Edit_SetFont(obj uintptr, value uintptr) {
   edit_SetFont.Call(obj, value)
}

func Edit_GetHideSelection(obj uintptr) bool {
    ret, _, _ := edit_GetHideSelection.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetHideSelection(obj uintptr, value bool) {
   edit_SetHideSelection.Call(obj, GoBoolToDBool(value))
}

func Edit_GetMaxLength(obj uintptr) int32 {
    ret, _, _ := edit_GetMaxLength.Call(obj)
    return int32(ret)
}

func Edit_SetMaxLength(obj uintptr, value int32) {
   edit_SetMaxLength.Call(obj, uintptr(value))
}

func Edit_GetNumbersOnly(obj uintptr) bool {
    ret, _, _ := edit_GetNumbersOnly.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetNumbersOnly(obj uintptr, value bool) {
   edit_SetNumbersOnly.Call(obj, GoBoolToDBool(value))
}

func Edit_GetParentColor(obj uintptr) bool {
    ret, _, _ := edit_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetParentColor(obj uintptr, value bool) {
   edit_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func Edit_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := edit_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetParentCtl3D(obj uintptr, value bool) {
   edit_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func Edit_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := edit_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetParentDoubleBuffered(obj uintptr, value bool) {
   edit_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func Edit_GetParentFont(obj uintptr) bool {
    ret, _, _ := edit_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetParentFont(obj uintptr, value bool) {
   edit_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func Edit_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := edit_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetParentShowHint(obj uintptr, value bool) {
   edit_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func Edit_GetPasswordChar(obj uintptr) uint16 {
    ret, _, _ := edit_GetPasswordChar.Call(obj)
    return uint16(ret)
}

func Edit_SetPasswordChar(obj uintptr, value uint16) {
   edit_SetPasswordChar.Call(obj, uintptr(value))
}

func Edit_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := edit_GetPopupMenu.Call(obj)
    return ret
}

func Edit_SetPopupMenu(obj uintptr, value uintptr) {
   edit_SetPopupMenu.Call(obj, value)
}

func Edit_GetReadOnly(obj uintptr) bool {
    ret, _, _ := edit_GetReadOnly.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetReadOnly(obj uintptr, value bool) {
   edit_SetReadOnly.Call(obj, GoBoolToDBool(value))
}

func Edit_GetShowHint(obj uintptr) bool {
    ret, _, _ := edit_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetShowHint(obj uintptr, value bool) {
   edit_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Edit_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := edit_GetTabOrder.Call(obj)
    return uint16(ret)
}

func Edit_SetTabOrder(obj uintptr, value uint16) {
   edit_SetTabOrder.Call(obj, uintptr(value))
}

func Edit_GetTabStop(obj uintptr) bool {
    ret, _, _ := edit_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetTabStop(obj uintptr, value bool) {
   edit_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func Edit_GetText(obj uintptr) string {
    ret, _, _ := edit_GetText.Call(obj)
    return DStrToGoStr(ret)
}

func Edit_SetText(obj uintptr, value string) {
   edit_SetText.Call(obj, GoStrToDStr(value))
}

func Edit_GetTextHint(obj uintptr) string {
    ret, _, _ := edit_GetTextHint.Call(obj)
    return DStrToGoStr(ret)
}

func Edit_SetTextHint(obj uintptr, value string) {
   edit_SetTextHint.Call(obj, GoStrToDStr(value))
}

func Edit_GetVisible(obj uintptr) bool {
    ret, _, _ := edit_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetVisible(obj uintptr, value bool) {
   edit_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Edit_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := edit_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func Edit_SetStyleElements(obj uintptr, value TStyleElements) {
   edit_SetStyleElements.Call(obj, uintptr(value))
}

func Edit_SetOnChange(obj uintptr, fn interface{}) {
    edit_SetOnChange.Call(obj, addEventToMap(fn))
}

func Edit_SetOnClick(obj uintptr, fn interface{}) {
    edit_SetOnClick.Call(obj, addEventToMap(fn))
}

func Edit_SetOnDblClick(obj uintptr, fn interface{}) {
    edit_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func Edit_SetOnEnter(obj uintptr, fn interface{}) {
    edit_SetOnEnter.Call(obj, addEventToMap(fn))
}

func Edit_SetOnExit(obj uintptr, fn interface{}) {
    edit_SetOnExit.Call(obj, addEventToMap(fn))
}

func Edit_SetOnKeyDown(obj uintptr, fn interface{}) {
    edit_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func Edit_SetOnKeyPress(obj uintptr, fn interface{}) {
    edit_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func Edit_SetOnKeyUp(obj uintptr, fn interface{}) {
    edit_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func Edit_SetOnMouseDown(obj uintptr, fn interface{}) {
    edit_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func Edit_SetOnMouseEnter(obj uintptr, fn interface{}) {
    edit_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func Edit_SetOnMouseLeave(obj uintptr, fn interface{}) {
    edit_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func Edit_SetOnMouseMove(obj uintptr, fn interface{}) {
    edit_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func Edit_SetOnMouseUp(obj uintptr, fn interface{}) {
    edit_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func Edit_GetModified(obj uintptr) bool {
    ret, _, _ := edit_GetModified.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetModified(obj uintptr, value bool) {
   edit_SetModified.Call(obj, GoBoolToDBool(value))
}

func Edit_GetSelLength(obj uintptr) int32 {
    ret, _, _ := edit_GetSelLength.Call(obj)
    return int32(ret)
}

func Edit_SetSelLength(obj uintptr, value int32) {
   edit_SetSelLength.Call(obj, uintptr(value))
}

func Edit_GetSelStart(obj uintptr) int32 {
    ret, _, _ := edit_GetSelStart.Call(obj)
    return int32(ret)
}

func Edit_SetSelStart(obj uintptr, value int32) {
   edit_SetSelStart.Call(obj, uintptr(value))
}

func Edit_GetSelText(obj uintptr) string {
    ret, _, _ := edit_GetSelText.Call(obj)
    return DStrToGoStr(ret)
}

func Edit_SetSelText(obj uintptr, value string) {
   edit_SetSelText.Call(obj, GoStrToDStr(value))
}

func Edit_GetBrush(obj uintptr) uintptr {
    ret, _, _ := edit_GetBrush.Call(obj)
    return ret
}

func Edit_GetControlCount(obj uintptr) int32 {
    ret, _, _ := edit_GetControlCount.Call(obj)
    return int32(ret)
}

func Edit_GetHandle(obj uintptr) HWND {
    ret, _, _ := edit_GetHandle.Call(obj)
    return HWND(ret)
}

func Edit_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := edit_GetParentWindow.Call(obj)
    return HWND(ret)
}

func Edit_SetParentWindow(obj uintptr, value HWND) {
   edit_SetParentWindow.Call(obj, uintptr(value))
}

func Edit_GetAction(obj uintptr) uintptr {
    ret, _, _ := edit_GetAction.Call(obj)
    return ret
}

func Edit_SetAction(obj uintptr, value uintptr) {
   edit_SetAction.Call(obj, value)
}

func Edit_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    edit_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Edit_SetBoundsRect(obj uintptr, value TRect) {
   edit_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Edit_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := edit_GetClientHeight.Call(obj)
    return int32(ret)
}

func Edit_SetClientHeight(obj uintptr, value int32) {
   edit_SetClientHeight.Call(obj, uintptr(value))
}

func Edit_GetClientRect(obj uintptr) TRect {
    var ret TRect
    edit_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Edit_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := edit_GetClientWidth.Call(obj)
    return int32(ret)
}

func Edit_SetClientWidth(obj uintptr, value int32) {
   edit_SetClientWidth.Call(obj, uintptr(value))
}

func Edit_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := edit_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Edit_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := edit_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Edit_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := edit_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Edit_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := edit_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Edit_GetParent(obj uintptr) uintptr {
    ret, _, _ := edit_GetParent.Call(obj)
    return ret
}

func Edit_SetParent(obj uintptr, value uintptr) {
   edit_SetParent.Call(obj, value)
}

func Edit_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := edit_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func Edit_SetAlignWithMargins(obj uintptr, value bool) {
   edit_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func Edit_GetLeft(obj uintptr) int32 {
    ret, _, _ := edit_GetLeft.Call(obj)
    return int32(ret)
}

func Edit_SetLeft(obj uintptr, value int32) {
   edit_SetLeft.Call(obj, uintptr(value))
}

func Edit_GetTop(obj uintptr) int32 {
    ret, _, _ := edit_GetTop.Call(obj)
    return int32(ret)
}

func Edit_SetTop(obj uintptr, value int32) {
   edit_SetTop.Call(obj, uintptr(value))
}

func Edit_GetWidth(obj uintptr) int32 {
    ret, _, _ := edit_GetWidth.Call(obj)
    return int32(ret)
}

func Edit_SetWidth(obj uintptr, value int32) {
   edit_SetWidth.Call(obj, uintptr(value))
}

func Edit_GetHeight(obj uintptr) int32 {
    ret, _, _ := edit_GetHeight.Call(obj)
    return int32(ret)
}

func Edit_SetHeight(obj uintptr, value int32) {
   edit_SetHeight.Call(obj, uintptr(value))
}

func Edit_GetCursor(obj uintptr) TCursor {
    ret, _, _ := edit_GetCursor.Call(obj)
    return TCursor(ret)
}

func Edit_SetCursor(obj uintptr, value TCursor) {
   edit_SetCursor.Call(obj, uintptr(value))
}

func Edit_GetHint(obj uintptr) string {
    ret, _, _ := edit_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Edit_SetHint(obj uintptr, value string) {
   edit_SetHint.Call(obj, GoStrToDStr(value))
}

func Edit_GetMargins(obj uintptr) uintptr {
    ret, _, _ := edit_GetMargins.Call(obj)
    return ret
}

func Edit_SetMargins(obj uintptr, value uintptr) {
   edit_SetMargins.Call(obj, value)
}

func Edit_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := edit_GetComponentCount.Call(obj)
    return int32(ret)
}

func Edit_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := edit_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Edit_SetComponentIndex(obj uintptr, value int32) {
   edit_SetComponentIndex.Call(obj, uintptr(value))
}

func Edit_GetOwner(obj uintptr) uintptr {
    ret, _, _ := edit_GetOwner.Call(obj)
    return ret
}

func Edit_GetName(obj uintptr) string {
    ret, _, _ := edit_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Edit_SetName(obj uintptr, value string) {
   edit_SetName.Call(obj, GoStrToDStr(value))
}

func Edit_GetTag(obj uintptr) int {
    ret, _, _ := edit_GetTag.Call(obj)
    return int(ret)
}

func Edit_SetTag(obj uintptr, value int) {
   edit_SetTag.Call(obj, uintptr(value))
}

func Edit_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := edit_GetControls.Call(obj, uintptr(Index))
    return ret
}

func Edit_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := edit_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TMainMenu ---------------------------

func MainMenu_Create(obj uintptr) uintptr {
    ret, _, _ := mainMenu_Create.Call(obj)
    return ret
}

func MainMenu_Free(obj uintptr) {
    mainMenu_Free.Call(obj)
}

func MainMenu_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := mainMenu_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func MainMenu_GetNamePath(obj uintptr) string {
    ret, _, _ := mainMenu_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func MainMenu_HasParent(obj uintptr) bool {
    ret, _, _ := mainMenu_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func MainMenu_Assign(obj uintptr, Source uintptr)  {
    mainMenu_Assign.Call(obj, Source )
}

func MainMenu_ClassName(obj uintptr) string {
    ret, _, _ := mainMenu_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func MainMenu_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := mainMenu_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func MainMenu_GetHashCode(obj uintptr) int32 {
    ret, _, _ := mainMenu_GetHashCode.Call(obj)
    return int32(ret)
}

func MainMenu_ToString(obj uintptr) string {
    ret, _, _ := mainMenu_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func MainMenu_GetAutoHotkeys(obj uintptr) TMenuAutoFlag {
    ret, _, _ := mainMenu_GetAutoHotkeys.Call(obj)
    return TMenuAutoFlag(ret)
}

func MainMenu_SetAutoHotkeys(obj uintptr, value TMenuAutoFlag) {
   mainMenu_SetAutoHotkeys.Call(obj, uintptr(value))
}

func MainMenu_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := mainMenu_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func MainMenu_SetBiDiMode(obj uintptr, value TBiDiMode) {
   mainMenu_SetBiDiMode.Call(obj, uintptr(value))
}

func MainMenu_GetImages(obj uintptr) uintptr {
    ret, _, _ := mainMenu_GetImages.Call(obj)
    return ret
}

func MainMenu_SetImages(obj uintptr, value uintptr) {
   mainMenu_SetImages.Call(obj, value)
}

func MainMenu_SetOnChange(obj uintptr, fn interface{}) {
    mainMenu_SetOnChange.Call(obj, addEventToMap(fn))
}

func MainMenu_GetHandle(obj uintptr) HMENU {
    ret, _, _ := mainMenu_GetHandle.Call(obj)
    return HMENU(ret)
}

func MainMenu_GetWindowHandle(obj uintptr) HWND {
    ret, _, _ := mainMenu_GetWindowHandle.Call(obj)
    return HWND(ret)
}

func MainMenu_SetWindowHandle(obj uintptr, value HWND) {
   mainMenu_SetWindowHandle.Call(obj, uintptr(value))
}

func MainMenu_GetItems(obj uintptr) uintptr {
    ret, _, _ := mainMenu_GetItems.Call(obj)
    return ret
}

func MainMenu_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := mainMenu_GetComponentCount.Call(obj)
    return int32(ret)
}

func MainMenu_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := mainMenu_GetComponentIndex.Call(obj)
    return int32(ret)
}

func MainMenu_SetComponentIndex(obj uintptr, value int32) {
   mainMenu_SetComponentIndex.Call(obj, uintptr(value))
}

func MainMenu_GetOwner(obj uintptr) uintptr {
    ret, _, _ := mainMenu_GetOwner.Call(obj)
    return ret
}

func MainMenu_GetName(obj uintptr) string {
    ret, _, _ := mainMenu_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func MainMenu_SetName(obj uintptr, value string) {
   mainMenu_SetName.Call(obj, GoStrToDStr(value))
}

func MainMenu_GetTag(obj uintptr) int {
    ret, _, _ := mainMenu_GetTag.Call(obj)
    return int(ret)
}

func MainMenu_SetTag(obj uintptr, value int) {
   mainMenu_SetTag.Call(obj, uintptr(value))
}

func MainMenu_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := mainMenu_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TPopupMenu ---------------------------

func PopupMenu_Create(obj uintptr) uintptr {
    ret, _, _ := popupMenu_Create.Call(obj)
    return ret
}

func PopupMenu_Free(obj uintptr) {
    popupMenu_Free.Call(obj)
}

func PopupMenu_CloseMenu(obj uintptr)  {
    popupMenu_CloseMenu.Call(obj)
}

func PopupMenu_Popup(obj uintptr, X int32, Y int32)  {
    popupMenu_Popup.Call(obj, uintptr(X) , uintptr(Y) )
}

func PopupMenu_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := popupMenu_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func PopupMenu_GetNamePath(obj uintptr) string {
    ret, _, _ := popupMenu_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func PopupMenu_HasParent(obj uintptr) bool {
    ret, _, _ := popupMenu_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func PopupMenu_Assign(obj uintptr, Source uintptr)  {
    popupMenu_Assign.Call(obj, Source )
}

func PopupMenu_ClassName(obj uintptr) string {
    ret, _, _ := popupMenu_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func PopupMenu_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := popupMenu_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func PopupMenu_GetHashCode(obj uintptr) int32 {
    ret, _, _ := popupMenu_GetHashCode.Call(obj)
    return int32(ret)
}

func PopupMenu_ToString(obj uintptr) string {
    ret, _, _ := popupMenu_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func PopupMenu_GetAlignment(obj uintptr) TPopupAlignment {
    ret, _, _ := popupMenu_GetAlignment.Call(obj)
    return TPopupAlignment(ret)
}

func PopupMenu_SetAlignment(obj uintptr, value TPopupAlignment) {
   popupMenu_SetAlignment.Call(obj, uintptr(value))
}

func PopupMenu_GetAutoHotkeys(obj uintptr) TMenuAutoFlag {
    ret, _, _ := popupMenu_GetAutoHotkeys.Call(obj)
    return TMenuAutoFlag(ret)
}

func PopupMenu_SetAutoHotkeys(obj uintptr, value TMenuAutoFlag) {
   popupMenu_SetAutoHotkeys.Call(obj, uintptr(value))
}

func PopupMenu_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := popupMenu_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func PopupMenu_SetBiDiMode(obj uintptr, value TBiDiMode) {
   popupMenu_SetBiDiMode.Call(obj, uintptr(value))
}

func PopupMenu_GetImages(obj uintptr) uintptr {
    ret, _, _ := popupMenu_GetImages.Call(obj)
    return ret
}

func PopupMenu_SetImages(obj uintptr, value uintptr) {
   popupMenu_SetImages.Call(obj, value)
}

func PopupMenu_SetOnChange(obj uintptr, fn interface{}) {
    popupMenu_SetOnChange.Call(obj, addEventToMap(fn))
}

func PopupMenu_SetOnPopup(obj uintptr, fn interface{}) {
    popupMenu_SetOnPopup.Call(obj, addEventToMap(fn))
}

func PopupMenu_GetHandle(obj uintptr) HMENU {
    ret, _, _ := popupMenu_GetHandle.Call(obj)
    return HMENU(ret)
}

func PopupMenu_GetWindowHandle(obj uintptr) HWND {
    ret, _, _ := popupMenu_GetWindowHandle.Call(obj)
    return HWND(ret)
}

func PopupMenu_SetWindowHandle(obj uintptr, value HWND) {
   popupMenu_SetWindowHandle.Call(obj, uintptr(value))
}

func PopupMenu_GetItems(obj uintptr) uintptr {
    ret, _, _ := popupMenu_GetItems.Call(obj)
    return ret
}

func PopupMenu_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := popupMenu_GetComponentCount.Call(obj)
    return int32(ret)
}

func PopupMenu_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := popupMenu_GetComponentIndex.Call(obj)
    return int32(ret)
}

func PopupMenu_SetComponentIndex(obj uintptr, value int32) {
   popupMenu_SetComponentIndex.Call(obj, uintptr(value))
}

func PopupMenu_GetOwner(obj uintptr) uintptr {
    ret, _, _ := popupMenu_GetOwner.Call(obj)
    return ret
}

func PopupMenu_GetName(obj uintptr) string {
    ret, _, _ := popupMenu_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func PopupMenu_SetName(obj uintptr, value string) {
   popupMenu_SetName.Call(obj, GoStrToDStr(value))
}

func PopupMenu_GetTag(obj uintptr) int {
    ret, _, _ := popupMenu_GetTag.Call(obj)
    return int(ret)
}

func PopupMenu_SetTag(obj uintptr, value int) {
   popupMenu_SetTag.Call(obj, uintptr(value))
}

func PopupMenu_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := popupMenu_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TMemo ---------------------------

func Memo_Create(obj uintptr) uintptr {
    ret, _, _ := memo_Create.Call(obj)
    return ret
}

func Memo_Free(obj uintptr) {
    memo_Free.Call(obj)
}

func Memo_Clear(obj uintptr)  {
    memo_Clear.Call(obj)
}

func Memo_ClearSelection(obj uintptr)  {
    memo_ClearSelection.Call(obj)
}

func Memo_CopyToClipboard(obj uintptr)  {
    memo_CopyToClipboard.Call(obj)
}

func Memo_CutToClipboard(obj uintptr)  {
    memo_CutToClipboard.Call(obj)
}

func Memo_PasteFromClipboard(obj uintptr)  {
    memo_PasteFromClipboard.Call(obj)
}

func Memo_SelectAll(obj uintptr)  {
    memo_SelectAll.Call(obj)
}

func Memo_GetSelTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := memo_GetSelTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Memo_CanFocus(obj uintptr) bool {
    ret, _, _ := memo_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_FlipChildren(obj uintptr, AllLevels bool)  {
    memo_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func Memo_Focused(obj uintptr) bool {
    ret, _, _ := memo_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_HandleAllocated(obj uintptr) bool {
    ret, _, _ := memo_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_Invalidate(obj uintptr)  {
    memo_Invalidate.Call(obj)
}

func Memo_Realign(obj uintptr)  {
    memo_Realign.Call(obj)
}

func Memo_Repaint(obj uintptr)  {
    memo_Repaint.Call(obj)
}

func Memo_ScaleBy(obj uintptr, M int32, D int32)  {
    memo_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func Memo_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    memo_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func Memo_SetFocus(obj uintptr)  {
    memo_SetFocus.Call(obj)
}

func Memo_Update(obj uintptr)  {
    memo_Update.Call(obj)
}

func Memo_BringToFront(obj uintptr)  {
    memo_BringToFront.Call(obj)
}

func Memo_HasParent(obj uintptr) bool {
    ret, _, _ := memo_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_Hide(obj uintptr)  {
    memo_Hide.Call(obj)
}

func Memo_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := memo_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func Memo_Refresh(obj uintptr)  {
    memo_Refresh.Call(obj)
}

func Memo_SendToBack(obj uintptr)  {
    memo_SendToBack.Call(obj)
}

func Memo_Show(obj uintptr)  {
    memo_Show.Call(obj)
}

func Memo_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := memo_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Memo_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := memo_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Memo_GetNamePath(obj uintptr) string {
    ret, _, _ := memo_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Memo_Assign(obj uintptr, Source uintptr)  {
    memo_Assign.Call(obj, Source )
}

func Memo_ClassName(obj uintptr) string {
    ret, _, _ := memo_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Memo_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := memo_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Memo_GetHashCode(obj uintptr) int32 {
    ret, _, _ := memo_GetHashCode.Call(obj)
    return int32(ret)
}

func Memo_ToString(obj uintptr) string {
    ret, _, _ := memo_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Memo_GetAlign(obj uintptr) TAlign {
    ret, _, _ := memo_GetAlign.Call(obj)
    return TAlign(ret)
}

func Memo_SetAlign(obj uintptr, value TAlign) {
   memo_SetAlign.Call(obj, uintptr(value))
}

func Memo_GetAlignment(obj uintptr) TAlignment {
    ret, _, _ := memo_GetAlignment.Call(obj)
    return TAlignment(ret)
}

func Memo_SetAlignment(obj uintptr, value TAlignment) {
   memo_SetAlignment.Call(obj, uintptr(value))
}

func Memo_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := memo_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func Memo_SetAnchors(obj uintptr, value TAnchors) {
   memo_SetAnchors.Call(obj, uintptr(value))
}

func Memo_GetBevelEdges(obj uintptr) TBevelEdges {
    ret, _, _ := memo_GetBevelEdges.Call(obj)
    return TBevelEdges(ret)
}

func Memo_SetBevelEdges(obj uintptr, value TBevelEdges) {
   memo_SetBevelEdges.Call(obj, uintptr(value))
}

func Memo_GetBevelInner(obj uintptr) TBevelCut {
    ret, _, _ := memo_GetBevelInner.Call(obj)
    return TBevelCut(ret)
}

func Memo_SetBevelInner(obj uintptr, value TBevelCut) {
   memo_SetBevelInner.Call(obj, uintptr(value))
}

func Memo_GetBevelKind(obj uintptr) TBevelKind {
    ret, _, _ := memo_GetBevelKind.Call(obj)
    return TBevelKind(ret)
}

func Memo_SetBevelKind(obj uintptr, value TBevelKind) {
   memo_SetBevelKind.Call(obj, uintptr(value))
}

func Memo_GetBevelOuter(obj uintptr) TBevelCut {
    ret, _, _ := memo_GetBevelOuter.Call(obj)
    return TBevelCut(ret)
}

func Memo_SetBevelOuter(obj uintptr, value TBevelCut) {
   memo_SetBevelOuter.Call(obj, uintptr(value))
}

func Memo_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := memo_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func Memo_SetBiDiMode(obj uintptr, value TBiDiMode) {
   memo_SetBiDiMode.Call(obj, uintptr(value))
}

func Memo_GetBorderStyle(obj uintptr) TBorderStyle {
    ret, _, _ := memo_GetBorderStyle.Call(obj)
    return TBorderStyle(ret)
}

func Memo_SetBorderStyle(obj uintptr, value TBorderStyle) {
   memo_SetBorderStyle.Call(obj, uintptr(value))
}

func Memo_GetColor(obj uintptr) TColor {
    ret, _, _ := memo_GetColor.Call(obj)
    return TColor(ret)
}

func Memo_SetColor(obj uintptr, value TColor) {
   memo_SetColor.Call(obj, uintptr(value))
}

func Memo_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := memo_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetDoubleBuffered(obj uintptr, value bool) {
   memo_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func Memo_GetEnabled(obj uintptr) bool {
    ret, _, _ := memo_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetEnabled(obj uintptr, value bool) {
   memo_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Memo_GetFont(obj uintptr) uintptr {
    ret, _, _ := memo_GetFont.Call(obj)
    return ret
}

func Memo_SetFont(obj uintptr, value uintptr) {
   memo_SetFont.Call(obj, value)
}

func Memo_GetHideSelection(obj uintptr) bool {
    ret, _, _ := memo_GetHideSelection.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetHideSelection(obj uintptr, value bool) {
   memo_SetHideSelection.Call(obj, GoBoolToDBool(value))
}

func Memo_GetLines(obj uintptr) uintptr {
    ret, _, _ := memo_GetLines.Call(obj)
    return ret
}

func Memo_SetLines(obj uintptr, value uintptr) {
   memo_SetLines.Call(obj, value)
}

func Memo_GetMaxLength(obj uintptr) int32 {
    ret, _, _ := memo_GetMaxLength.Call(obj)
    return int32(ret)
}

func Memo_SetMaxLength(obj uintptr, value int32) {
   memo_SetMaxLength.Call(obj, uintptr(value))
}

func Memo_GetParentColor(obj uintptr) bool {
    ret, _, _ := memo_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetParentColor(obj uintptr, value bool) {
   memo_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func Memo_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := memo_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetParentCtl3D(obj uintptr, value bool) {
   memo_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func Memo_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := memo_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetParentDoubleBuffered(obj uintptr, value bool) {
   memo_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func Memo_GetParentFont(obj uintptr) bool {
    ret, _, _ := memo_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetParentFont(obj uintptr, value bool) {
   memo_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func Memo_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := memo_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetParentShowHint(obj uintptr, value bool) {
   memo_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func Memo_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := memo_GetPopupMenu.Call(obj)
    return ret
}

func Memo_SetPopupMenu(obj uintptr, value uintptr) {
   memo_SetPopupMenu.Call(obj, value)
}

func Memo_GetReadOnly(obj uintptr) bool {
    ret, _, _ := memo_GetReadOnly.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetReadOnly(obj uintptr, value bool) {
   memo_SetReadOnly.Call(obj, GoBoolToDBool(value))
}

func Memo_GetScrollBars(obj uintptr) TScrollStyle {
    ret, _, _ := memo_GetScrollBars.Call(obj)
    return TScrollStyle(ret)
}

func Memo_SetScrollBars(obj uintptr, value TScrollStyle) {
   memo_SetScrollBars.Call(obj, uintptr(value))
}

func Memo_GetShowHint(obj uintptr) bool {
    ret, _, _ := memo_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetShowHint(obj uintptr, value bool) {
   memo_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Memo_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := memo_GetTabOrder.Call(obj)
    return uint16(ret)
}

func Memo_SetTabOrder(obj uintptr, value uint16) {
   memo_SetTabOrder.Call(obj, uintptr(value))
}

func Memo_GetTabStop(obj uintptr) bool {
    ret, _, _ := memo_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetTabStop(obj uintptr, value bool) {
   memo_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func Memo_GetVisible(obj uintptr) bool {
    ret, _, _ := memo_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetVisible(obj uintptr, value bool) {
   memo_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Memo_GetWantReturns(obj uintptr) bool {
    ret, _, _ := memo_GetWantReturns.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetWantReturns(obj uintptr, value bool) {
   memo_SetWantReturns.Call(obj, GoBoolToDBool(value))
}

func Memo_GetWantTabs(obj uintptr) bool {
    ret, _, _ := memo_GetWantTabs.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetWantTabs(obj uintptr, value bool) {
   memo_SetWantTabs.Call(obj, GoBoolToDBool(value))
}

func Memo_GetWordWrap(obj uintptr) bool {
    ret, _, _ := memo_GetWordWrap.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetWordWrap(obj uintptr, value bool) {
   memo_SetWordWrap.Call(obj, GoBoolToDBool(value))
}

func Memo_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := memo_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func Memo_SetStyleElements(obj uintptr, value TStyleElements) {
   memo_SetStyleElements.Call(obj, uintptr(value))
}

func Memo_SetOnChange(obj uintptr, fn interface{}) {
    memo_SetOnChange.Call(obj, addEventToMap(fn))
}

func Memo_SetOnClick(obj uintptr, fn interface{}) {
    memo_SetOnClick.Call(obj, addEventToMap(fn))
}

func Memo_SetOnDblClick(obj uintptr, fn interface{}) {
    memo_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func Memo_SetOnEnter(obj uintptr, fn interface{}) {
    memo_SetOnEnter.Call(obj, addEventToMap(fn))
}

func Memo_SetOnExit(obj uintptr, fn interface{}) {
    memo_SetOnExit.Call(obj, addEventToMap(fn))
}

func Memo_SetOnKeyDown(obj uintptr, fn interface{}) {
    memo_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func Memo_SetOnKeyPress(obj uintptr, fn interface{}) {
    memo_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func Memo_SetOnKeyUp(obj uintptr, fn interface{}) {
    memo_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func Memo_SetOnMouseDown(obj uintptr, fn interface{}) {
    memo_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func Memo_SetOnMouseEnter(obj uintptr, fn interface{}) {
    memo_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func Memo_SetOnMouseLeave(obj uintptr, fn interface{}) {
    memo_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func Memo_SetOnMouseMove(obj uintptr, fn interface{}) {
    memo_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func Memo_SetOnMouseUp(obj uintptr, fn interface{}) {
    memo_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func Memo_GetCaretPos(obj uintptr) TPoint {
    var ret TPoint
    memo_GetCaretPos.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Memo_SetCaretPos(obj uintptr, value TPoint) {
   memo_SetCaretPos.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Memo_GetModified(obj uintptr) bool {
    ret, _, _ := memo_GetModified.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetModified(obj uintptr, value bool) {
   memo_SetModified.Call(obj, GoBoolToDBool(value))
}

func Memo_GetSelLength(obj uintptr) int32 {
    ret, _, _ := memo_GetSelLength.Call(obj)
    return int32(ret)
}

func Memo_SetSelLength(obj uintptr, value int32) {
   memo_SetSelLength.Call(obj, uintptr(value))
}

func Memo_GetSelStart(obj uintptr) int32 {
    ret, _, _ := memo_GetSelStart.Call(obj)
    return int32(ret)
}

func Memo_SetSelStart(obj uintptr, value int32) {
   memo_SetSelStart.Call(obj, uintptr(value))
}

func Memo_GetSelText(obj uintptr) string {
    ret, _, _ := memo_GetSelText.Call(obj)
    return DStrToGoStr(ret)
}

func Memo_SetSelText(obj uintptr, value string) {
   memo_SetSelText.Call(obj, GoStrToDStr(value))
}

func Memo_GetText(obj uintptr) string {
    ret, _, _ := memo_GetText.Call(obj)
    return DStrToGoStr(ret)
}

func Memo_SetText(obj uintptr, value string) {
   memo_SetText.Call(obj, GoStrToDStr(value))
}

func Memo_GetTextHint(obj uintptr) string {
    ret, _, _ := memo_GetTextHint.Call(obj)
    return DStrToGoStr(ret)
}

func Memo_SetTextHint(obj uintptr, value string) {
   memo_SetTextHint.Call(obj, GoStrToDStr(value))
}

func Memo_GetBrush(obj uintptr) uintptr {
    ret, _, _ := memo_GetBrush.Call(obj)
    return ret
}

func Memo_GetControlCount(obj uintptr) int32 {
    ret, _, _ := memo_GetControlCount.Call(obj)
    return int32(ret)
}

func Memo_GetHandle(obj uintptr) HWND {
    ret, _, _ := memo_GetHandle.Call(obj)
    return HWND(ret)
}

func Memo_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := memo_GetParentWindow.Call(obj)
    return HWND(ret)
}

func Memo_SetParentWindow(obj uintptr, value HWND) {
   memo_SetParentWindow.Call(obj, uintptr(value))
}

func Memo_GetAction(obj uintptr) uintptr {
    ret, _, _ := memo_GetAction.Call(obj)
    return ret
}

func Memo_SetAction(obj uintptr, value uintptr) {
   memo_SetAction.Call(obj, value)
}

func Memo_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    memo_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Memo_SetBoundsRect(obj uintptr, value TRect) {
   memo_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Memo_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := memo_GetClientHeight.Call(obj)
    return int32(ret)
}

func Memo_SetClientHeight(obj uintptr, value int32) {
   memo_SetClientHeight.Call(obj, uintptr(value))
}

func Memo_GetClientRect(obj uintptr) TRect {
    var ret TRect
    memo_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Memo_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := memo_GetClientWidth.Call(obj)
    return int32(ret)
}

func Memo_SetClientWidth(obj uintptr, value int32) {
   memo_SetClientWidth.Call(obj, uintptr(value))
}

func Memo_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := memo_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Memo_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := memo_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Memo_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := memo_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Memo_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := memo_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Memo_GetParent(obj uintptr) uintptr {
    ret, _, _ := memo_GetParent.Call(obj)
    return ret
}

func Memo_SetParent(obj uintptr, value uintptr) {
   memo_SetParent.Call(obj, value)
}

func Memo_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := memo_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func Memo_SetAlignWithMargins(obj uintptr, value bool) {
   memo_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func Memo_GetLeft(obj uintptr) int32 {
    ret, _, _ := memo_GetLeft.Call(obj)
    return int32(ret)
}

func Memo_SetLeft(obj uintptr, value int32) {
   memo_SetLeft.Call(obj, uintptr(value))
}

func Memo_GetTop(obj uintptr) int32 {
    ret, _, _ := memo_GetTop.Call(obj)
    return int32(ret)
}

func Memo_SetTop(obj uintptr, value int32) {
   memo_SetTop.Call(obj, uintptr(value))
}

func Memo_GetWidth(obj uintptr) int32 {
    ret, _, _ := memo_GetWidth.Call(obj)
    return int32(ret)
}

func Memo_SetWidth(obj uintptr, value int32) {
   memo_SetWidth.Call(obj, uintptr(value))
}

func Memo_GetHeight(obj uintptr) int32 {
    ret, _, _ := memo_GetHeight.Call(obj)
    return int32(ret)
}

func Memo_SetHeight(obj uintptr, value int32) {
   memo_SetHeight.Call(obj, uintptr(value))
}

func Memo_GetCursor(obj uintptr) TCursor {
    ret, _, _ := memo_GetCursor.Call(obj)
    return TCursor(ret)
}

func Memo_SetCursor(obj uintptr, value TCursor) {
   memo_SetCursor.Call(obj, uintptr(value))
}

func Memo_GetHint(obj uintptr) string {
    ret, _, _ := memo_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Memo_SetHint(obj uintptr, value string) {
   memo_SetHint.Call(obj, GoStrToDStr(value))
}

func Memo_GetMargins(obj uintptr) uintptr {
    ret, _, _ := memo_GetMargins.Call(obj)
    return ret
}

func Memo_SetMargins(obj uintptr, value uintptr) {
   memo_SetMargins.Call(obj, value)
}

func Memo_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := memo_GetComponentCount.Call(obj)
    return int32(ret)
}

func Memo_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := memo_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Memo_SetComponentIndex(obj uintptr, value int32) {
   memo_SetComponentIndex.Call(obj, uintptr(value))
}

func Memo_GetOwner(obj uintptr) uintptr {
    ret, _, _ := memo_GetOwner.Call(obj)
    return ret
}

func Memo_GetName(obj uintptr) string {
    ret, _, _ := memo_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Memo_SetName(obj uintptr, value string) {
   memo_SetName.Call(obj, GoStrToDStr(value))
}

func Memo_GetTag(obj uintptr) int {
    ret, _, _ := memo_GetTag.Call(obj)
    return int(ret)
}

func Memo_SetTag(obj uintptr, value int) {
   memo_SetTag.Call(obj, uintptr(value))
}

func Memo_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := memo_GetControls.Call(obj, uintptr(Index))
    return ret
}

func Memo_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := memo_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TCheckBox ---------------------------

func CheckBox_Create(obj uintptr) uintptr {
    ret, _, _ := checkBox_Create.Call(obj)
    return ret
}

func CheckBox_Free(obj uintptr) {
    checkBox_Free.Call(obj)
}

func CheckBox_CanFocus(obj uintptr) bool {
    ret, _, _ := checkBox_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_FlipChildren(obj uintptr, AllLevels bool)  {
    checkBox_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func CheckBox_Focused(obj uintptr) bool {
    ret, _, _ := checkBox_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_HandleAllocated(obj uintptr) bool {
    ret, _, _ := checkBox_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_Invalidate(obj uintptr)  {
    checkBox_Invalidate.Call(obj)
}

func CheckBox_Realign(obj uintptr)  {
    checkBox_Realign.Call(obj)
}

func CheckBox_Repaint(obj uintptr)  {
    checkBox_Repaint.Call(obj)
}

func CheckBox_ScaleBy(obj uintptr, M int32, D int32)  {
    checkBox_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func CheckBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    checkBox_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func CheckBox_SetFocus(obj uintptr)  {
    checkBox_SetFocus.Call(obj)
}

func CheckBox_Update(obj uintptr)  {
    checkBox_Update.Call(obj)
}

func CheckBox_BringToFront(obj uintptr)  {
    checkBox_BringToFront.Call(obj)
}

func CheckBox_HasParent(obj uintptr) bool {
    ret, _, _ := checkBox_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_Hide(obj uintptr)  {
    checkBox_Hide.Call(obj)
}

func CheckBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := checkBox_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func CheckBox_Refresh(obj uintptr)  {
    checkBox_Refresh.Call(obj)
}

func CheckBox_SendToBack(obj uintptr)  {
    checkBox_SendToBack.Call(obj)
}

func CheckBox_Show(obj uintptr)  {
    checkBox_Show.Call(obj)
}

func CheckBox_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := checkBox_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func CheckBox_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := checkBox_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func CheckBox_GetNamePath(obj uintptr) string {
    ret, _, _ := checkBox_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func CheckBox_Assign(obj uintptr, Source uintptr)  {
    checkBox_Assign.Call(obj, Source )
}

func CheckBox_ClassName(obj uintptr) string {
    ret, _, _ := checkBox_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func CheckBox_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := checkBox_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func CheckBox_GetHashCode(obj uintptr) int32 {
    ret, _, _ := checkBox_GetHashCode.Call(obj)
    return int32(ret)
}

func CheckBox_ToString(obj uintptr) string {
    ret, _, _ := checkBox_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func CheckBox_GetAction(obj uintptr) uintptr {
    ret, _, _ := checkBox_GetAction.Call(obj)
    return ret
}

func CheckBox_SetAction(obj uintptr, value uintptr) {
   checkBox_SetAction.Call(obj, value)
}

func CheckBox_GetAlign(obj uintptr) TAlign {
    ret, _, _ := checkBox_GetAlign.Call(obj)
    return TAlign(ret)
}

func CheckBox_SetAlign(obj uintptr, value TAlign) {
   checkBox_SetAlign.Call(obj, uintptr(value))
}

func CheckBox_GetAlignment(obj uintptr) TLeftRight {
    ret, _, _ := checkBox_GetAlignment.Call(obj)
    return TLeftRight(ret)
}

func CheckBox_SetAlignment(obj uintptr, value TLeftRight) {
   checkBox_SetAlignment.Call(obj, uintptr(value))
}

func CheckBox_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := checkBox_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func CheckBox_SetAnchors(obj uintptr, value TAnchors) {
   checkBox_SetAnchors.Call(obj, uintptr(value))
}

func CheckBox_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := checkBox_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func CheckBox_SetBiDiMode(obj uintptr, value TBiDiMode) {
   checkBox_SetBiDiMode.Call(obj, uintptr(value))
}

func CheckBox_GetCaption(obj uintptr) string {
    ret, _, _ := checkBox_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func CheckBox_SetCaption(obj uintptr, value string) {
   checkBox_SetCaption.Call(obj, GoStrToDStr(value))
}

func CheckBox_GetChecked(obj uintptr) bool {
    ret, _, _ := checkBox_GetChecked.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetChecked(obj uintptr, value bool) {
   checkBox_SetChecked.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetColor(obj uintptr) TColor {
    ret, _, _ := checkBox_GetColor.Call(obj)
    return TColor(ret)
}

func CheckBox_SetColor(obj uintptr, value TColor) {
   checkBox_SetColor.Call(obj, uintptr(value))
}

func CheckBox_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := checkBox_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetDoubleBuffered(obj uintptr, value bool) {
   checkBox_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetEnabled(obj uintptr) bool {
    ret, _, _ := checkBox_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetEnabled(obj uintptr, value bool) {
   checkBox_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetFont(obj uintptr) uintptr {
    ret, _, _ := checkBox_GetFont.Call(obj)
    return ret
}

func CheckBox_SetFont(obj uintptr, value uintptr) {
   checkBox_SetFont.Call(obj, value)
}

func CheckBox_GetParentColor(obj uintptr) bool {
    ret, _, _ := checkBox_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetParentColor(obj uintptr, value bool) {
   checkBox_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := checkBox_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetParentCtl3D(obj uintptr, value bool) {
   checkBox_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := checkBox_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetParentDoubleBuffered(obj uintptr, value bool) {
   checkBox_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetParentFont(obj uintptr) bool {
    ret, _, _ := checkBox_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetParentFont(obj uintptr, value bool) {
   checkBox_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := checkBox_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetParentShowHint(obj uintptr, value bool) {
   checkBox_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := checkBox_GetPopupMenu.Call(obj)
    return ret
}

func CheckBox_SetPopupMenu(obj uintptr, value uintptr) {
   checkBox_SetPopupMenu.Call(obj, value)
}

func CheckBox_GetShowHint(obj uintptr) bool {
    ret, _, _ := checkBox_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetShowHint(obj uintptr, value bool) {
   checkBox_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetState(obj uintptr) TCheckBoxState {
    ret, _, _ := checkBox_GetState.Call(obj)
    return TCheckBoxState(ret)
}

func CheckBox_SetState(obj uintptr, value TCheckBoxState) {
   checkBox_SetState.Call(obj, uintptr(value))
}

func CheckBox_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := checkBox_GetTabOrder.Call(obj)
    return uint16(ret)
}

func CheckBox_SetTabOrder(obj uintptr, value uint16) {
   checkBox_SetTabOrder.Call(obj, uintptr(value))
}

func CheckBox_GetTabStop(obj uintptr) bool {
    ret, _, _ := checkBox_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetTabStop(obj uintptr, value bool) {
   checkBox_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetVisible(obj uintptr) bool {
    ret, _, _ := checkBox_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetVisible(obj uintptr, value bool) {
   checkBox_SetVisible.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetWordWrap(obj uintptr) bool {
    ret, _, _ := checkBox_GetWordWrap.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetWordWrap(obj uintptr, value bool) {
   checkBox_SetWordWrap.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := checkBox_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func CheckBox_SetStyleElements(obj uintptr, value TStyleElements) {
   checkBox_SetStyleElements.Call(obj, uintptr(value))
}

func CheckBox_SetOnClick(obj uintptr, fn interface{}) {
    checkBox_SetOnClick.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnEnter(obj uintptr, fn interface{}) {
    checkBox_SetOnEnter.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnExit(obj uintptr, fn interface{}) {
    checkBox_SetOnExit.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnKeyDown(obj uintptr, fn interface{}) {
    checkBox_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnKeyPress(obj uintptr, fn interface{}) {
    checkBox_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnKeyUp(obj uintptr, fn interface{}) {
    checkBox_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnMouseDown(obj uintptr, fn interface{}) {
    checkBox_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
    checkBox_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
    checkBox_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnMouseMove(obj uintptr, fn interface{}) {
    checkBox_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func CheckBox_SetOnMouseUp(obj uintptr, fn interface{}) {
    checkBox_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func CheckBox_GetBrush(obj uintptr) uintptr {
    ret, _, _ := checkBox_GetBrush.Call(obj)
    return ret
}

func CheckBox_GetControlCount(obj uintptr) int32 {
    ret, _, _ := checkBox_GetControlCount.Call(obj)
    return int32(ret)
}

func CheckBox_GetHandle(obj uintptr) HWND {
    ret, _, _ := checkBox_GetHandle.Call(obj)
    return HWND(ret)
}

func CheckBox_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := checkBox_GetParentWindow.Call(obj)
    return HWND(ret)
}

func CheckBox_SetParentWindow(obj uintptr, value HWND) {
   checkBox_SetParentWindow.Call(obj, uintptr(value))
}

func CheckBox_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    checkBox_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func CheckBox_SetBoundsRect(obj uintptr, value TRect) {
   checkBox_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func CheckBox_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := checkBox_GetClientHeight.Call(obj)
    return int32(ret)
}

func CheckBox_SetClientHeight(obj uintptr, value int32) {
   checkBox_SetClientHeight.Call(obj, uintptr(value))
}

func CheckBox_GetClientRect(obj uintptr) TRect {
    var ret TRect
    checkBox_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func CheckBox_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := checkBox_GetClientWidth.Call(obj)
    return int32(ret)
}

func CheckBox_SetClientWidth(obj uintptr, value int32) {
   checkBox_SetClientWidth.Call(obj, uintptr(value))
}

func CheckBox_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := checkBox_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func CheckBox_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := checkBox_GetExplicitTop.Call(obj)
    return int32(ret)
}

func CheckBox_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := checkBox_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func CheckBox_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := checkBox_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func CheckBox_GetParent(obj uintptr) uintptr {
    ret, _, _ := checkBox_GetParent.Call(obj)
    return ret
}

func CheckBox_SetParent(obj uintptr, value uintptr) {
   checkBox_SetParent.Call(obj, value)
}

func CheckBox_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := checkBox_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func CheckBox_SetAlignWithMargins(obj uintptr, value bool) {
   checkBox_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func CheckBox_GetLeft(obj uintptr) int32 {
    ret, _, _ := checkBox_GetLeft.Call(obj)
    return int32(ret)
}

func CheckBox_SetLeft(obj uintptr, value int32) {
   checkBox_SetLeft.Call(obj, uintptr(value))
}

func CheckBox_GetTop(obj uintptr) int32 {
    ret, _, _ := checkBox_GetTop.Call(obj)
    return int32(ret)
}

func CheckBox_SetTop(obj uintptr, value int32) {
   checkBox_SetTop.Call(obj, uintptr(value))
}

func CheckBox_GetWidth(obj uintptr) int32 {
    ret, _, _ := checkBox_GetWidth.Call(obj)
    return int32(ret)
}

func CheckBox_SetWidth(obj uintptr, value int32) {
   checkBox_SetWidth.Call(obj, uintptr(value))
}

func CheckBox_GetHeight(obj uintptr) int32 {
    ret, _, _ := checkBox_GetHeight.Call(obj)
    return int32(ret)
}

func CheckBox_SetHeight(obj uintptr, value int32) {
   checkBox_SetHeight.Call(obj, uintptr(value))
}

func CheckBox_GetCursor(obj uintptr) TCursor {
    ret, _, _ := checkBox_GetCursor.Call(obj)
    return TCursor(ret)
}

func CheckBox_SetCursor(obj uintptr, value TCursor) {
   checkBox_SetCursor.Call(obj, uintptr(value))
}

func CheckBox_GetHint(obj uintptr) string {
    ret, _, _ := checkBox_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func CheckBox_SetHint(obj uintptr, value string) {
   checkBox_SetHint.Call(obj, GoStrToDStr(value))
}

func CheckBox_GetMargins(obj uintptr) uintptr {
    ret, _, _ := checkBox_GetMargins.Call(obj)
    return ret
}

func CheckBox_SetMargins(obj uintptr, value uintptr) {
   checkBox_SetMargins.Call(obj, value)
}

func CheckBox_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := checkBox_GetComponentCount.Call(obj)
    return int32(ret)
}

func CheckBox_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := checkBox_GetComponentIndex.Call(obj)
    return int32(ret)
}

func CheckBox_SetComponentIndex(obj uintptr, value int32) {
   checkBox_SetComponentIndex.Call(obj, uintptr(value))
}

func CheckBox_GetOwner(obj uintptr) uintptr {
    ret, _, _ := checkBox_GetOwner.Call(obj)
    return ret
}

func CheckBox_GetName(obj uintptr) string {
    ret, _, _ := checkBox_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func CheckBox_SetName(obj uintptr, value string) {
   checkBox_SetName.Call(obj, GoStrToDStr(value))
}

func CheckBox_GetTag(obj uintptr) int {
    ret, _, _ := checkBox_GetTag.Call(obj)
    return int(ret)
}

func CheckBox_SetTag(obj uintptr, value int) {
   checkBox_SetTag.Call(obj, uintptr(value))
}

func CheckBox_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := checkBox_GetControls.Call(obj, uintptr(Index))
    return ret
}

func CheckBox_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := checkBox_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TRadioButton ---------------------------

func RadioButton_Create(obj uintptr) uintptr {
    ret, _, _ := radioButton_Create.Call(obj)
    return ret
}

func RadioButton_Free(obj uintptr) {
    radioButton_Free.Call(obj)
}

func RadioButton_CanFocus(obj uintptr) bool {
    ret, _, _ := radioButton_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_FlipChildren(obj uintptr, AllLevels bool)  {
    radioButton_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func RadioButton_Focused(obj uintptr) bool {
    ret, _, _ := radioButton_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_HandleAllocated(obj uintptr) bool {
    ret, _, _ := radioButton_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_Invalidate(obj uintptr)  {
    radioButton_Invalidate.Call(obj)
}

func RadioButton_Realign(obj uintptr)  {
    radioButton_Realign.Call(obj)
}

func RadioButton_Repaint(obj uintptr)  {
    radioButton_Repaint.Call(obj)
}

func RadioButton_ScaleBy(obj uintptr, M int32, D int32)  {
    radioButton_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func RadioButton_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    radioButton_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func RadioButton_SetFocus(obj uintptr)  {
    radioButton_SetFocus.Call(obj)
}

func RadioButton_Update(obj uintptr)  {
    radioButton_Update.Call(obj)
}

func RadioButton_BringToFront(obj uintptr)  {
    radioButton_BringToFront.Call(obj)
}

func RadioButton_HasParent(obj uintptr) bool {
    ret, _, _ := radioButton_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_Hide(obj uintptr)  {
    radioButton_Hide.Call(obj)
}

func RadioButton_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := radioButton_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func RadioButton_Refresh(obj uintptr)  {
    radioButton_Refresh.Call(obj)
}

func RadioButton_SendToBack(obj uintptr)  {
    radioButton_SendToBack.Call(obj)
}

func RadioButton_Show(obj uintptr)  {
    radioButton_Show.Call(obj)
}

func RadioButton_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := radioButton_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func RadioButton_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := radioButton_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func RadioButton_GetNamePath(obj uintptr) string {
    ret, _, _ := radioButton_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func RadioButton_Assign(obj uintptr, Source uintptr)  {
    radioButton_Assign.Call(obj, Source )
}

func RadioButton_ClassName(obj uintptr) string {
    ret, _, _ := radioButton_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func RadioButton_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := radioButton_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func RadioButton_GetHashCode(obj uintptr) int32 {
    ret, _, _ := radioButton_GetHashCode.Call(obj)
    return int32(ret)
}

func RadioButton_ToString(obj uintptr) string {
    ret, _, _ := radioButton_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func RadioButton_GetAction(obj uintptr) uintptr {
    ret, _, _ := radioButton_GetAction.Call(obj)
    return ret
}

func RadioButton_SetAction(obj uintptr, value uintptr) {
   radioButton_SetAction.Call(obj, value)
}

func RadioButton_GetAlign(obj uintptr) TAlign {
    ret, _, _ := radioButton_GetAlign.Call(obj)
    return TAlign(ret)
}

func RadioButton_SetAlign(obj uintptr, value TAlign) {
   radioButton_SetAlign.Call(obj, uintptr(value))
}

func RadioButton_GetAlignment(obj uintptr) TLeftRight {
    ret, _, _ := radioButton_GetAlignment.Call(obj)
    return TLeftRight(ret)
}

func RadioButton_SetAlignment(obj uintptr, value TLeftRight) {
   radioButton_SetAlignment.Call(obj, uintptr(value))
}

func RadioButton_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := radioButton_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func RadioButton_SetAnchors(obj uintptr, value TAnchors) {
   radioButton_SetAnchors.Call(obj, uintptr(value))
}

func RadioButton_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := radioButton_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func RadioButton_SetBiDiMode(obj uintptr, value TBiDiMode) {
   radioButton_SetBiDiMode.Call(obj, uintptr(value))
}

func RadioButton_GetCaption(obj uintptr) string {
    ret, _, _ := radioButton_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func RadioButton_SetCaption(obj uintptr, value string) {
   radioButton_SetCaption.Call(obj, GoStrToDStr(value))
}

func RadioButton_GetChecked(obj uintptr) bool {
    ret, _, _ := radioButton_GetChecked.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetChecked(obj uintptr, value bool) {
   radioButton_SetChecked.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetColor(obj uintptr) TColor {
    ret, _, _ := radioButton_GetColor.Call(obj)
    return TColor(ret)
}

func RadioButton_SetColor(obj uintptr, value TColor) {
   radioButton_SetColor.Call(obj, uintptr(value))
}

func RadioButton_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := radioButton_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetDoubleBuffered(obj uintptr, value bool) {
   radioButton_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetEnabled(obj uintptr) bool {
    ret, _, _ := radioButton_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetEnabled(obj uintptr, value bool) {
   radioButton_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetFont(obj uintptr) uintptr {
    ret, _, _ := radioButton_GetFont.Call(obj)
    return ret
}

func RadioButton_SetFont(obj uintptr, value uintptr) {
   radioButton_SetFont.Call(obj, value)
}

func RadioButton_GetParentColor(obj uintptr) bool {
    ret, _, _ := radioButton_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetParentColor(obj uintptr, value bool) {
   radioButton_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := radioButton_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetParentCtl3D(obj uintptr, value bool) {
   radioButton_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := radioButton_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetParentDoubleBuffered(obj uintptr, value bool) {
   radioButton_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetParentFont(obj uintptr) bool {
    ret, _, _ := radioButton_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetParentFont(obj uintptr, value bool) {
   radioButton_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := radioButton_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetParentShowHint(obj uintptr, value bool) {
   radioButton_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := radioButton_GetPopupMenu.Call(obj)
    return ret
}

func RadioButton_SetPopupMenu(obj uintptr, value uintptr) {
   radioButton_SetPopupMenu.Call(obj, value)
}

func RadioButton_GetShowHint(obj uintptr) bool {
    ret, _, _ := radioButton_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetShowHint(obj uintptr, value bool) {
   radioButton_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := radioButton_GetTabOrder.Call(obj)
    return uint16(ret)
}

func RadioButton_SetTabOrder(obj uintptr, value uint16) {
   radioButton_SetTabOrder.Call(obj, uintptr(value))
}

func RadioButton_GetTabStop(obj uintptr) bool {
    ret, _, _ := radioButton_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetTabStop(obj uintptr, value bool) {
   radioButton_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetVisible(obj uintptr) bool {
    ret, _, _ := radioButton_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetVisible(obj uintptr, value bool) {
   radioButton_SetVisible.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetWordWrap(obj uintptr) bool {
    ret, _, _ := radioButton_GetWordWrap.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetWordWrap(obj uintptr, value bool) {
   radioButton_SetWordWrap.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := radioButton_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func RadioButton_SetStyleElements(obj uintptr, value TStyleElements) {
   radioButton_SetStyleElements.Call(obj, uintptr(value))
}

func RadioButton_SetOnClick(obj uintptr, fn interface{}) {
    radioButton_SetOnClick.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnDblClick(obj uintptr, fn interface{}) {
    radioButton_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnEnter(obj uintptr, fn interface{}) {
    radioButton_SetOnEnter.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnExit(obj uintptr, fn interface{}) {
    radioButton_SetOnExit.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnKeyDown(obj uintptr, fn interface{}) {
    radioButton_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnKeyPress(obj uintptr, fn interface{}) {
    radioButton_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnKeyUp(obj uintptr, fn interface{}) {
    radioButton_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnMouseDown(obj uintptr, fn interface{}) {
    radioButton_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnMouseEnter(obj uintptr, fn interface{}) {
    radioButton_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnMouseLeave(obj uintptr, fn interface{}) {
    radioButton_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnMouseMove(obj uintptr, fn interface{}) {
    radioButton_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func RadioButton_SetOnMouseUp(obj uintptr, fn interface{}) {
    radioButton_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func RadioButton_GetBrush(obj uintptr) uintptr {
    ret, _, _ := radioButton_GetBrush.Call(obj)
    return ret
}

func RadioButton_GetControlCount(obj uintptr) int32 {
    ret, _, _ := radioButton_GetControlCount.Call(obj)
    return int32(ret)
}

func RadioButton_GetHandle(obj uintptr) HWND {
    ret, _, _ := radioButton_GetHandle.Call(obj)
    return HWND(ret)
}

func RadioButton_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := radioButton_GetParentWindow.Call(obj)
    return HWND(ret)
}

func RadioButton_SetParentWindow(obj uintptr, value HWND) {
   radioButton_SetParentWindow.Call(obj, uintptr(value))
}

func RadioButton_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    radioButton_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func RadioButton_SetBoundsRect(obj uintptr, value TRect) {
   radioButton_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func RadioButton_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := radioButton_GetClientHeight.Call(obj)
    return int32(ret)
}

func RadioButton_SetClientHeight(obj uintptr, value int32) {
   radioButton_SetClientHeight.Call(obj, uintptr(value))
}

func RadioButton_GetClientRect(obj uintptr) TRect {
    var ret TRect
    radioButton_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func RadioButton_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := radioButton_GetClientWidth.Call(obj)
    return int32(ret)
}

func RadioButton_SetClientWidth(obj uintptr, value int32) {
   radioButton_SetClientWidth.Call(obj, uintptr(value))
}

func RadioButton_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := radioButton_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func RadioButton_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := radioButton_GetExplicitTop.Call(obj)
    return int32(ret)
}

func RadioButton_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := radioButton_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func RadioButton_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := radioButton_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func RadioButton_GetParent(obj uintptr) uintptr {
    ret, _, _ := radioButton_GetParent.Call(obj)
    return ret
}

func RadioButton_SetParent(obj uintptr, value uintptr) {
   radioButton_SetParent.Call(obj, value)
}

func RadioButton_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := radioButton_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioButton_SetAlignWithMargins(obj uintptr, value bool) {
   radioButton_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func RadioButton_GetLeft(obj uintptr) int32 {
    ret, _, _ := radioButton_GetLeft.Call(obj)
    return int32(ret)
}

func RadioButton_SetLeft(obj uintptr, value int32) {
   radioButton_SetLeft.Call(obj, uintptr(value))
}

func RadioButton_GetTop(obj uintptr) int32 {
    ret, _, _ := radioButton_GetTop.Call(obj)
    return int32(ret)
}

func RadioButton_SetTop(obj uintptr, value int32) {
   radioButton_SetTop.Call(obj, uintptr(value))
}

func RadioButton_GetWidth(obj uintptr) int32 {
    ret, _, _ := radioButton_GetWidth.Call(obj)
    return int32(ret)
}

func RadioButton_SetWidth(obj uintptr, value int32) {
   radioButton_SetWidth.Call(obj, uintptr(value))
}

func RadioButton_GetHeight(obj uintptr) int32 {
    ret, _, _ := radioButton_GetHeight.Call(obj)
    return int32(ret)
}

func RadioButton_SetHeight(obj uintptr, value int32) {
   radioButton_SetHeight.Call(obj, uintptr(value))
}

func RadioButton_GetCursor(obj uintptr) TCursor {
    ret, _, _ := radioButton_GetCursor.Call(obj)
    return TCursor(ret)
}

func RadioButton_SetCursor(obj uintptr, value TCursor) {
   radioButton_SetCursor.Call(obj, uintptr(value))
}

func RadioButton_GetHint(obj uintptr) string {
    ret, _, _ := radioButton_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func RadioButton_SetHint(obj uintptr, value string) {
   radioButton_SetHint.Call(obj, GoStrToDStr(value))
}

func RadioButton_GetMargins(obj uintptr) uintptr {
    ret, _, _ := radioButton_GetMargins.Call(obj)
    return ret
}

func RadioButton_SetMargins(obj uintptr, value uintptr) {
   radioButton_SetMargins.Call(obj, value)
}

func RadioButton_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := radioButton_GetComponentCount.Call(obj)
    return int32(ret)
}

func RadioButton_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := radioButton_GetComponentIndex.Call(obj)
    return int32(ret)
}

func RadioButton_SetComponentIndex(obj uintptr, value int32) {
   radioButton_SetComponentIndex.Call(obj, uintptr(value))
}

func RadioButton_GetOwner(obj uintptr) uintptr {
    ret, _, _ := radioButton_GetOwner.Call(obj)
    return ret
}

func RadioButton_GetName(obj uintptr) string {
    ret, _, _ := radioButton_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func RadioButton_SetName(obj uintptr, value string) {
   radioButton_SetName.Call(obj, GoStrToDStr(value))
}

func RadioButton_GetTag(obj uintptr) int {
    ret, _, _ := radioButton_GetTag.Call(obj)
    return int(ret)
}

func RadioButton_SetTag(obj uintptr, value int) {
   radioButton_SetTag.Call(obj, uintptr(value))
}

func RadioButton_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := radioButton_GetControls.Call(obj, uintptr(Index))
    return ret
}

func RadioButton_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := radioButton_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TGroupBox ---------------------------

func GroupBox_Create(obj uintptr) uintptr {
    ret, _, _ := groupBox_Create.Call(obj)
    return ret
}

func GroupBox_Free(obj uintptr) {
    groupBox_Free.Call(obj)
}

func GroupBox_CanFocus(obj uintptr) bool {
    ret, _, _ := groupBox_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_FlipChildren(obj uintptr, AllLevels bool)  {
    groupBox_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func GroupBox_Focused(obj uintptr) bool {
    ret, _, _ := groupBox_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_HandleAllocated(obj uintptr) bool {
    ret, _, _ := groupBox_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_Invalidate(obj uintptr)  {
    groupBox_Invalidate.Call(obj)
}

func GroupBox_Realign(obj uintptr)  {
    groupBox_Realign.Call(obj)
}

func GroupBox_Repaint(obj uintptr)  {
    groupBox_Repaint.Call(obj)
}

func GroupBox_ScaleBy(obj uintptr, M int32, D int32)  {
    groupBox_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func GroupBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    groupBox_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func GroupBox_SetFocus(obj uintptr)  {
    groupBox_SetFocus.Call(obj)
}

func GroupBox_Update(obj uintptr)  {
    groupBox_Update.Call(obj)
}

func GroupBox_BringToFront(obj uintptr)  {
    groupBox_BringToFront.Call(obj)
}

func GroupBox_HasParent(obj uintptr) bool {
    ret, _, _ := groupBox_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_Hide(obj uintptr)  {
    groupBox_Hide.Call(obj)
}

func GroupBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := groupBox_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func GroupBox_Refresh(obj uintptr)  {
    groupBox_Refresh.Call(obj)
}

func GroupBox_SendToBack(obj uintptr)  {
    groupBox_SendToBack.Call(obj)
}

func GroupBox_Show(obj uintptr)  {
    groupBox_Show.Call(obj)
}

func GroupBox_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := groupBox_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func GroupBox_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := groupBox_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func GroupBox_GetNamePath(obj uintptr) string {
    ret, _, _ := groupBox_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func GroupBox_Assign(obj uintptr, Source uintptr)  {
    groupBox_Assign.Call(obj, Source )
}

func GroupBox_ClassName(obj uintptr) string {
    ret, _, _ := groupBox_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func GroupBox_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := groupBox_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func GroupBox_GetHashCode(obj uintptr) int32 {
    ret, _, _ := groupBox_GetHashCode.Call(obj)
    return int32(ret)
}

func GroupBox_ToString(obj uintptr) string {
    ret, _, _ := groupBox_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func GroupBox_GetAlign(obj uintptr) TAlign {
    ret, _, _ := groupBox_GetAlign.Call(obj)
    return TAlign(ret)
}

func GroupBox_SetAlign(obj uintptr, value TAlign) {
   groupBox_SetAlign.Call(obj, uintptr(value))
}

func GroupBox_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := groupBox_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func GroupBox_SetAnchors(obj uintptr, value TAnchors) {
   groupBox_SetAnchors.Call(obj, uintptr(value))
}

func GroupBox_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := groupBox_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func GroupBox_SetBiDiMode(obj uintptr, value TBiDiMode) {
   groupBox_SetBiDiMode.Call(obj, uintptr(value))
}

func GroupBox_GetCaption(obj uintptr) string {
    ret, _, _ := groupBox_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func GroupBox_SetCaption(obj uintptr, value string) {
   groupBox_SetCaption.Call(obj, GoStrToDStr(value))
}

func GroupBox_GetColor(obj uintptr) TColor {
    ret, _, _ := groupBox_GetColor.Call(obj)
    return TColor(ret)
}

func GroupBox_SetColor(obj uintptr, value TColor) {
   groupBox_SetColor.Call(obj, uintptr(value))
}

func GroupBox_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := groupBox_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetDoubleBuffered(obj uintptr, value bool) {
   groupBox_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetEnabled(obj uintptr) bool {
    ret, _, _ := groupBox_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetEnabled(obj uintptr, value bool) {
   groupBox_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetFont(obj uintptr) uintptr {
    ret, _, _ := groupBox_GetFont.Call(obj)
    return ret
}

func GroupBox_SetFont(obj uintptr, value uintptr) {
   groupBox_SetFont.Call(obj, value)
}

func GroupBox_GetParentBackground(obj uintptr) bool {
    ret, _, _ := groupBox_GetParentBackground.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetParentBackground(obj uintptr, value bool) {
   groupBox_SetParentBackground.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetParentColor(obj uintptr) bool {
    ret, _, _ := groupBox_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetParentColor(obj uintptr, value bool) {
   groupBox_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := groupBox_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetParentCtl3D(obj uintptr, value bool) {
   groupBox_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := groupBox_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetParentDoubleBuffered(obj uintptr, value bool) {
   groupBox_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetParentFont(obj uintptr) bool {
    ret, _, _ := groupBox_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetParentFont(obj uintptr, value bool) {
   groupBox_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := groupBox_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetParentShowHint(obj uintptr, value bool) {
   groupBox_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := groupBox_GetPopupMenu.Call(obj)
    return ret
}

func GroupBox_SetPopupMenu(obj uintptr, value uintptr) {
   groupBox_SetPopupMenu.Call(obj, value)
}

func GroupBox_GetShowHint(obj uintptr) bool {
    ret, _, _ := groupBox_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetShowHint(obj uintptr, value bool) {
   groupBox_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := groupBox_GetTabOrder.Call(obj)
    return uint16(ret)
}

func GroupBox_SetTabOrder(obj uintptr, value uint16) {
   groupBox_SetTabOrder.Call(obj, uintptr(value))
}

func GroupBox_GetTabStop(obj uintptr) bool {
    ret, _, _ := groupBox_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetTabStop(obj uintptr, value bool) {
   groupBox_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetVisible(obj uintptr) bool {
    ret, _, _ := groupBox_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetVisible(obj uintptr, value bool) {
   groupBox_SetVisible.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := groupBox_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func GroupBox_SetStyleElements(obj uintptr, value TStyleElements) {
   groupBox_SetStyleElements.Call(obj, uintptr(value))
}

func GroupBox_SetOnClick(obj uintptr, fn interface{}) {
    groupBox_SetOnClick.Call(obj, addEventToMap(fn))
}

func GroupBox_SetOnDblClick(obj uintptr, fn interface{}) {
    groupBox_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func GroupBox_SetOnEnter(obj uintptr, fn interface{}) {
    groupBox_SetOnEnter.Call(obj, addEventToMap(fn))
}

func GroupBox_SetOnExit(obj uintptr, fn interface{}) {
    groupBox_SetOnExit.Call(obj, addEventToMap(fn))
}

func GroupBox_SetOnMouseDown(obj uintptr, fn interface{}) {
    groupBox_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func GroupBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
    groupBox_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func GroupBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
    groupBox_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func GroupBox_SetOnMouseMove(obj uintptr, fn interface{}) {
    groupBox_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func GroupBox_SetOnMouseUp(obj uintptr, fn interface{}) {
    groupBox_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func GroupBox_GetBrush(obj uintptr) uintptr {
    ret, _, _ := groupBox_GetBrush.Call(obj)
    return ret
}

func GroupBox_GetControlCount(obj uintptr) int32 {
    ret, _, _ := groupBox_GetControlCount.Call(obj)
    return int32(ret)
}

func GroupBox_GetHandle(obj uintptr) HWND {
    ret, _, _ := groupBox_GetHandle.Call(obj)
    return HWND(ret)
}

func GroupBox_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := groupBox_GetParentWindow.Call(obj)
    return HWND(ret)
}

func GroupBox_SetParentWindow(obj uintptr, value HWND) {
   groupBox_SetParentWindow.Call(obj, uintptr(value))
}

func GroupBox_GetAction(obj uintptr) uintptr {
    ret, _, _ := groupBox_GetAction.Call(obj)
    return ret
}

func GroupBox_SetAction(obj uintptr, value uintptr) {
   groupBox_SetAction.Call(obj, value)
}

func GroupBox_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    groupBox_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func GroupBox_SetBoundsRect(obj uintptr, value TRect) {
   groupBox_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func GroupBox_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := groupBox_GetClientHeight.Call(obj)
    return int32(ret)
}

func GroupBox_SetClientHeight(obj uintptr, value int32) {
   groupBox_SetClientHeight.Call(obj, uintptr(value))
}

func GroupBox_GetClientRect(obj uintptr) TRect {
    var ret TRect
    groupBox_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func GroupBox_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := groupBox_GetClientWidth.Call(obj)
    return int32(ret)
}

func GroupBox_SetClientWidth(obj uintptr, value int32) {
   groupBox_SetClientWidth.Call(obj, uintptr(value))
}

func GroupBox_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := groupBox_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func GroupBox_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := groupBox_GetExplicitTop.Call(obj)
    return int32(ret)
}

func GroupBox_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := groupBox_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func GroupBox_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := groupBox_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func GroupBox_GetParent(obj uintptr) uintptr {
    ret, _, _ := groupBox_GetParent.Call(obj)
    return ret
}

func GroupBox_SetParent(obj uintptr, value uintptr) {
   groupBox_SetParent.Call(obj, value)
}

func GroupBox_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := groupBox_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func GroupBox_SetAlignWithMargins(obj uintptr, value bool) {
   groupBox_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func GroupBox_GetLeft(obj uintptr) int32 {
    ret, _, _ := groupBox_GetLeft.Call(obj)
    return int32(ret)
}

func GroupBox_SetLeft(obj uintptr, value int32) {
   groupBox_SetLeft.Call(obj, uintptr(value))
}

func GroupBox_GetTop(obj uintptr) int32 {
    ret, _, _ := groupBox_GetTop.Call(obj)
    return int32(ret)
}

func GroupBox_SetTop(obj uintptr, value int32) {
   groupBox_SetTop.Call(obj, uintptr(value))
}

func GroupBox_GetWidth(obj uintptr) int32 {
    ret, _, _ := groupBox_GetWidth.Call(obj)
    return int32(ret)
}

func GroupBox_SetWidth(obj uintptr, value int32) {
   groupBox_SetWidth.Call(obj, uintptr(value))
}

func GroupBox_GetHeight(obj uintptr) int32 {
    ret, _, _ := groupBox_GetHeight.Call(obj)
    return int32(ret)
}

func GroupBox_SetHeight(obj uintptr, value int32) {
   groupBox_SetHeight.Call(obj, uintptr(value))
}

func GroupBox_GetCursor(obj uintptr) TCursor {
    ret, _, _ := groupBox_GetCursor.Call(obj)
    return TCursor(ret)
}

func GroupBox_SetCursor(obj uintptr, value TCursor) {
   groupBox_SetCursor.Call(obj, uintptr(value))
}

func GroupBox_GetHint(obj uintptr) string {
    ret, _, _ := groupBox_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func GroupBox_SetHint(obj uintptr, value string) {
   groupBox_SetHint.Call(obj, GoStrToDStr(value))
}

func GroupBox_GetMargins(obj uintptr) uintptr {
    ret, _, _ := groupBox_GetMargins.Call(obj)
    return ret
}

func GroupBox_SetMargins(obj uintptr, value uintptr) {
   groupBox_SetMargins.Call(obj, value)
}

func GroupBox_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := groupBox_GetComponentCount.Call(obj)
    return int32(ret)
}

func GroupBox_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := groupBox_GetComponentIndex.Call(obj)
    return int32(ret)
}

func GroupBox_SetComponentIndex(obj uintptr, value int32) {
   groupBox_SetComponentIndex.Call(obj, uintptr(value))
}

func GroupBox_GetOwner(obj uintptr) uintptr {
    ret, _, _ := groupBox_GetOwner.Call(obj)
    return ret
}

func GroupBox_GetName(obj uintptr) string {
    ret, _, _ := groupBox_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func GroupBox_SetName(obj uintptr, value string) {
   groupBox_SetName.Call(obj, GoStrToDStr(value))
}

func GroupBox_GetTag(obj uintptr) int {
    ret, _, _ := groupBox_GetTag.Call(obj)
    return int(ret)
}

func GroupBox_SetTag(obj uintptr, value int) {
   groupBox_SetTag.Call(obj, uintptr(value))
}

func GroupBox_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := groupBox_GetControls.Call(obj, uintptr(Index))
    return ret
}

func GroupBox_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := groupBox_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TLabel ---------------------------

func Label_Create(obj uintptr) uintptr {
    ret, _, _ := label_Create.Call(obj)
    return ret
}

func Label_Free(obj uintptr) {
    label_Free.Call(obj)
}

func Label_BringToFront(obj uintptr)  {
    label_BringToFront.Call(obj)
}

func Label_HasParent(obj uintptr) bool {
    ret, _, _ := label_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_Hide(obj uintptr)  {
    label_Hide.Call(obj)
}

func Label_Invalidate(obj uintptr)  {
    label_Invalidate.Call(obj)
}

func Label_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := label_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func Label_Refresh(obj uintptr)  {
    label_Refresh.Call(obj)
}

func Label_Repaint(obj uintptr)  {
    label_Repaint.Call(obj)
}

func Label_SendToBack(obj uintptr)  {
    label_SendToBack.Call(obj)
}

func Label_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    label_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func Label_Show(obj uintptr)  {
    label_Show.Call(obj)
}

func Label_Update(obj uintptr)  {
    label_Update.Call(obj)
}

func Label_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := label_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Label_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := label_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Label_GetNamePath(obj uintptr) string {
    ret, _, _ := label_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Label_Assign(obj uintptr, Source uintptr)  {
    label_Assign.Call(obj, Source )
}

func Label_ClassName(obj uintptr) string {
    ret, _, _ := label_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Label_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := label_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Label_GetHashCode(obj uintptr) int32 {
    ret, _, _ := label_GetHashCode.Call(obj)
    return int32(ret)
}

func Label_ToString(obj uintptr) string {
    ret, _, _ := label_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Label_GetAlign(obj uintptr) TAlign {
    ret, _, _ := label_GetAlign.Call(obj)
    return TAlign(ret)
}

func Label_SetAlign(obj uintptr, value TAlign) {
   label_SetAlign.Call(obj, uintptr(value))
}

func Label_GetAlignment(obj uintptr) TAlignment {
    ret, _, _ := label_GetAlignment.Call(obj)
    return TAlignment(ret)
}

func Label_SetAlignment(obj uintptr, value TAlignment) {
   label_SetAlignment.Call(obj, uintptr(value))
}

func Label_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := label_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func Label_SetAnchors(obj uintptr, value TAnchors) {
   label_SetAnchors.Call(obj, uintptr(value))
}

func Label_GetAutoSize(obj uintptr) bool {
    ret, _, _ := label_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetAutoSize(obj uintptr, value bool) {
   label_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func Label_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := label_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func Label_SetBiDiMode(obj uintptr, value TBiDiMode) {
   label_SetBiDiMode.Call(obj, uintptr(value))
}

func Label_GetCaption(obj uintptr) string {
    ret, _, _ := label_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func Label_SetCaption(obj uintptr, value string) {
   label_SetCaption.Call(obj, GoStrToDStr(value))
}

func Label_GetColor(obj uintptr) TColor {
    ret, _, _ := label_GetColor.Call(obj)
    return TColor(ret)
}

func Label_SetColor(obj uintptr, value TColor) {
   label_SetColor.Call(obj, uintptr(value))
}

func Label_GetEllipsisPosition(obj uintptr) TEllipsisPosition {
    ret, _, _ := label_GetEllipsisPosition.Call(obj)
    return TEllipsisPosition(ret)
}

func Label_SetEllipsisPosition(obj uintptr, value TEllipsisPosition) {
   label_SetEllipsisPosition.Call(obj, uintptr(value))
}

func Label_GetEnabled(obj uintptr) bool {
    ret, _, _ := label_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetEnabled(obj uintptr, value bool) {
   label_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Label_GetFont(obj uintptr) uintptr {
    ret, _, _ := label_GetFont.Call(obj)
    return ret
}

func Label_SetFont(obj uintptr, value uintptr) {
   label_SetFont.Call(obj, value)
}

func Label_GetGlowSize(obj uintptr) int32 {
    ret, _, _ := label_GetGlowSize.Call(obj)
    return int32(ret)
}

func Label_SetGlowSize(obj uintptr, value int32) {
   label_SetGlowSize.Call(obj, uintptr(value))
}

func Label_GetParentColor(obj uintptr) bool {
    ret, _, _ := label_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetParentColor(obj uintptr, value bool) {
   label_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func Label_GetParentFont(obj uintptr) bool {
    ret, _, _ := label_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetParentFont(obj uintptr, value bool) {
   label_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func Label_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := label_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetParentShowHint(obj uintptr, value bool) {
   label_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func Label_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := label_GetPopupMenu.Call(obj)
    return ret
}

func Label_SetPopupMenu(obj uintptr, value uintptr) {
   label_SetPopupMenu.Call(obj, value)
}

func Label_GetShowAccelChar(obj uintptr) bool {
    ret, _, _ := label_GetShowAccelChar.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetShowAccelChar(obj uintptr, value bool) {
   label_SetShowAccelChar.Call(obj, GoBoolToDBool(value))
}

func Label_GetShowHint(obj uintptr) bool {
    ret, _, _ := label_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetShowHint(obj uintptr, value bool) {
   label_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Label_GetTransparent(obj uintptr) bool {
    ret, _, _ := label_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetTransparent(obj uintptr, value bool) {
   label_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func Label_GetLayout(obj uintptr) TTextLayout {
    ret, _, _ := label_GetLayout.Call(obj)
    return TTextLayout(ret)
}

func Label_SetLayout(obj uintptr, value TTextLayout) {
   label_SetLayout.Call(obj, uintptr(value))
}

func Label_GetVisible(obj uintptr) bool {
    ret, _, _ := label_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetVisible(obj uintptr, value bool) {
   label_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Label_GetWordWrap(obj uintptr) bool {
    ret, _, _ := label_GetWordWrap.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetWordWrap(obj uintptr, value bool) {
   label_SetWordWrap.Call(obj, GoBoolToDBool(value))
}

func Label_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := label_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func Label_SetStyleElements(obj uintptr, value TStyleElements) {
   label_SetStyleElements.Call(obj, uintptr(value))
}

func Label_SetOnClick(obj uintptr, fn interface{}) {
    label_SetOnClick.Call(obj, addEventToMap(fn))
}

func Label_SetOnDblClick(obj uintptr, fn interface{}) {
    label_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func Label_SetOnMouseDown(obj uintptr, fn interface{}) {
    label_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func Label_SetOnMouseMove(obj uintptr, fn interface{}) {
    label_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func Label_SetOnMouseUp(obj uintptr, fn interface{}) {
    label_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func Label_SetOnMouseEnter(obj uintptr, fn interface{}) {
    label_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func Label_SetOnMouseLeave(obj uintptr, fn interface{}) {
    label_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func Label_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := label_GetCanvas.Call(obj)
    return ret
}

func Label_GetAction(obj uintptr) uintptr {
    ret, _, _ := label_GetAction.Call(obj)
    return ret
}

func Label_SetAction(obj uintptr, value uintptr) {
   label_SetAction.Call(obj, value)
}

func Label_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    label_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Label_SetBoundsRect(obj uintptr, value TRect) {
   label_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Label_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := label_GetClientHeight.Call(obj)
    return int32(ret)
}

func Label_SetClientHeight(obj uintptr, value int32) {
   label_SetClientHeight.Call(obj, uintptr(value))
}

func Label_GetClientRect(obj uintptr) TRect {
    var ret TRect
    label_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Label_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := label_GetClientWidth.Call(obj)
    return int32(ret)
}

func Label_SetClientWidth(obj uintptr, value int32) {
   label_SetClientWidth.Call(obj, uintptr(value))
}

func Label_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := label_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Label_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := label_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Label_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := label_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Label_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := label_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Label_GetParent(obj uintptr) uintptr {
    ret, _, _ := label_GetParent.Call(obj)
    return ret
}

func Label_SetParent(obj uintptr, value uintptr) {
   label_SetParent.Call(obj, value)
}

func Label_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := label_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func Label_SetAlignWithMargins(obj uintptr, value bool) {
   label_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func Label_GetLeft(obj uintptr) int32 {
    ret, _, _ := label_GetLeft.Call(obj)
    return int32(ret)
}

func Label_SetLeft(obj uintptr, value int32) {
   label_SetLeft.Call(obj, uintptr(value))
}

func Label_GetTop(obj uintptr) int32 {
    ret, _, _ := label_GetTop.Call(obj)
    return int32(ret)
}

func Label_SetTop(obj uintptr, value int32) {
   label_SetTop.Call(obj, uintptr(value))
}

func Label_GetWidth(obj uintptr) int32 {
    ret, _, _ := label_GetWidth.Call(obj)
    return int32(ret)
}

func Label_SetWidth(obj uintptr, value int32) {
   label_SetWidth.Call(obj, uintptr(value))
}

func Label_GetHeight(obj uintptr) int32 {
    ret, _, _ := label_GetHeight.Call(obj)
    return int32(ret)
}

func Label_SetHeight(obj uintptr, value int32) {
   label_SetHeight.Call(obj, uintptr(value))
}

func Label_GetCursor(obj uintptr) TCursor {
    ret, _, _ := label_GetCursor.Call(obj)
    return TCursor(ret)
}

func Label_SetCursor(obj uintptr, value TCursor) {
   label_SetCursor.Call(obj, uintptr(value))
}

func Label_GetHint(obj uintptr) string {
    ret, _, _ := label_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Label_SetHint(obj uintptr, value string) {
   label_SetHint.Call(obj, GoStrToDStr(value))
}

func Label_GetMargins(obj uintptr) uintptr {
    ret, _, _ := label_GetMargins.Call(obj)
    return ret
}

func Label_SetMargins(obj uintptr, value uintptr) {
   label_SetMargins.Call(obj, value)
}

func Label_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := label_GetComponentCount.Call(obj)
    return int32(ret)
}

func Label_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := label_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Label_SetComponentIndex(obj uintptr, value int32) {
   label_SetComponentIndex.Call(obj, uintptr(value))
}

func Label_GetOwner(obj uintptr) uintptr {
    ret, _, _ := label_GetOwner.Call(obj)
    return ret
}

func Label_GetName(obj uintptr) string {
    ret, _, _ := label_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Label_SetName(obj uintptr, value string) {
   label_SetName.Call(obj, GoStrToDStr(value))
}

func Label_GetTag(obj uintptr) int {
    ret, _, _ := label_GetTag.Call(obj)
    return int(ret)
}

func Label_SetTag(obj uintptr, value int) {
   label_SetTag.Call(obj, uintptr(value))
}

func Label_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := label_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TListBox ---------------------------

func ListBox_Create(obj uintptr) uintptr {
    ret, _, _ := listBox_Create.Call(obj)
    return ret
}

func ListBox_Free(obj uintptr) {
    listBox_Free.Call(obj)
}

func ListBox_AddItem(obj uintptr, Item string, AObject uintptr)  {
    listBox_AddItem.Call(obj, GoStrToDStr(Item) , AObject )
}

func ListBox_Clear(obj uintptr)  {
    listBox_Clear.Call(obj)
}

func ListBox_ClearSelection(obj uintptr)  {
    listBox_ClearSelection.Call(obj)
}

func ListBox_DeleteSelected(obj uintptr)  {
    listBox_DeleteSelected.Call(obj)
}

func ListBox_SelectAll(obj uintptr)  {
    listBox_SelectAll.Call(obj)
}

func ListBox_CanFocus(obj uintptr) bool {
    ret, _, _ := listBox_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_FlipChildren(obj uintptr, AllLevels bool)  {
    listBox_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func ListBox_Focused(obj uintptr) bool {
    ret, _, _ := listBox_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_HandleAllocated(obj uintptr) bool {
    ret, _, _ := listBox_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_Invalidate(obj uintptr)  {
    listBox_Invalidate.Call(obj)
}

func ListBox_Realign(obj uintptr)  {
    listBox_Realign.Call(obj)
}

func ListBox_Repaint(obj uintptr)  {
    listBox_Repaint.Call(obj)
}

func ListBox_ScaleBy(obj uintptr, M int32, D int32)  {
    listBox_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func ListBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    listBox_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func ListBox_SetFocus(obj uintptr)  {
    listBox_SetFocus.Call(obj)
}

func ListBox_Update(obj uintptr)  {
    listBox_Update.Call(obj)
}

func ListBox_BringToFront(obj uintptr)  {
    listBox_BringToFront.Call(obj)
}

func ListBox_HasParent(obj uintptr) bool {
    ret, _, _ := listBox_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_Hide(obj uintptr)  {
    listBox_Hide.Call(obj)
}

func ListBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := listBox_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func ListBox_Refresh(obj uintptr)  {
    listBox_Refresh.Call(obj)
}

func ListBox_SendToBack(obj uintptr)  {
    listBox_SendToBack.Call(obj)
}

func ListBox_Show(obj uintptr)  {
    listBox_Show.Call(obj)
}

func ListBox_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := listBox_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func ListBox_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := listBox_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ListBox_GetNamePath(obj uintptr) string {
    ret, _, _ := listBox_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ListBox_Assign(obj uintptr, Source uintptr)  {
    listBox_Assign.Call(obj, Source )
}

func ListBox_ClassName(obj uintptr) string {
    ret, _, _ := listBox_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ListBox_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := listBox_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ListBox_GetHashCode(obj uintptr) int32 {
    ret, _, _ := listBox_GetHashCode.Call(obj)
    return int32(ret)
}

func ListBox_ToString(obj uintptr) string {
    ret, _, _ := listBox_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ListBox_GetStyle(obj uintptr) TListBoxStyle {
    ret, _, _ := listBox_GetStyle.Call(obj)
    return TListBoxStyle(ret)
}

func ListBox_SetStyle(obj uintptr, value TListBoxStyle) {
   listBox_SetStyle.Call(obj, uintptr(value))
}

func ListBox_GetAutoComplete(obj uintptr) bool {
    ret, _, _ := listBox_GetAutoComplete.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetAutoComplete(obj uintptr, value bool) {
   listBox_SetAutoComplete.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetAutoCompleteDelay(obj uintptr) uint32 {
    ret, _, _ := listBox_GetAutoCompleteDelay.Call(obj)
    return uint32(ret)
}

func ListBox_SetAutoCompleteDelay(obj uintptr, value uint32) {
   listBox_SetAutoCompleteDelay.Call(obj, uintptr(value))
}

func ListBox_GetAlign(obj uintptr) TAlign {
    ret, _, _ := listBox_GetAlign.Call(obj)
    return TAlign(ret)
}

func ListBox_SetAlign(obj uintptr, value TAlign) {
   listBox_SetAlign.Call(obj, uintptr(value))
}

func ListBox_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := listBox_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func ListBox_SetAnchors(obj uintptr, value TAnchors) {
   listBox_SetAnchors.Call(obj, uintptr(value))
}

func ListBox_GetBevelEdges(obj uintptr) TBevelEdges {
    ret, _, _ := listBox_GetBevelEdges.Call(obj)
    return TBevelEdges(ret)
}

func ListBox_SetBevelEdges(obj uintptr, value TBevelEdges) {
   listBox_SetBevelEdges.Call(obj, uintptr(value))
}

func ListBox_GetBevelInner(obj uintptr) TBevelCut {
    ret, _, _ := listBox_GetBevelInner.Call(obj)
    return TBevelCut(ret)
}

func ListBox_SetBevelInner(obj uintptr, value TBevelCut) {
   listBox_SetBevelInner.Call(obj, uintptr(value))
}

func ListBox_GetBevelKind(obj uintptr) TBevelKind {
    ret, _, _ := listBox_GetBevelKind.Call(obj)
    return TBevelKind(ret)
}

func ListBox_SetBevelKind(obj uintptr, value TBevelKind) {
   listBox_SetBevelKind.Call(obj, uintptr(value))
}

func ListBox_GetBevelOuter(obj uintptr) TBevelCut {
    ret, _, _ := listBox_GetBevelOuter.Call(obj)
    return TBevelCut(ret)
}

func ListBox_SetBevelOuter(obj uintptr, value TBevelCut) {
   listBox_SetBevelOuter.Call(obj, uintptr(value))
}

func ListBox_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := listBox_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func ListBox_SetBiDiMode(obj uintptr, value TBiDiMode) {
   listBox_SetBiDiMode.Call(obj, uintptr(value))
}

func ListBox_GetBorderStyle(obj uintptr) TBorderStyle {
    ret, _, _ := listBox_GetBorderStyle.Call(obj)
    return TBorderStyle(ret)
}

func ListBox_SetBorderStyle(obj uintptr, value TBorderStyle) {
   listBox_SetBorderStyle.Call(obj, uintptr(value))
}

func ListBox_GetColor(obj uintptr) TColor {
    ret, _, _ := listBox_GetColor.Call(obj)
    return TColor(ret)
}

func ListBox_SetColor(obj uintptr, value TColor) {
   listBox_SetColor.Call(obj, uintptr(value))
}

func ListBox_GetColumns(obj uintptr) int32 {
    ret, _, _ := listBox_GetColumns.Call(obj)
    return int32(ret)
}

func ListBox_SetColumns(obj uintptr, value int32) {
   listBox_SetColumns.Call(obj, uintptr(value))
}

func ListBox_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := listBox_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetDoubleBuffered(obj uintptr, value bool) {
   listBox_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetEnabled(obj uintptr) bool {
    ret, _, _ := listBox_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetEnabled(obj uintptr, value bool) {
   listBox_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetFont(obj uintptr) uintptr {
    ret, _, _ := listBox_GetFont.Call(obj)
    return ret
}

func ListBox_SetFont(obj uintptr, value uintptr) {
   listBox_SetFont.Call(obj, value)
}

func ListBox_GetItemHeight(obj uintptr) int32 {
    ret, _, _ := listBox_GetItemHeight.Call(obj)
    return int32(ret)
}

func ListBox_SetItemHeight(obj uintptr, value int32) {
   listBox_SetItemHeight.Call(obj, uintptr(value))
}

func ListBox_GetItems(obj uintptr) uintptr {
    ret, _, _ := listBox_GetItems.Call(obj)
    return ret
}

func ListBox_SetItems(obj uintptr, value uintptr) {
   listBox_SetItems.Call(obj, value)
}

func ListBox_GetMultiSelect(obj uintptr) bool {
    ret, _, _ := listBox_GetMultiSelect.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetMultiSelect(obj uintptr, value bool) {
   listBox_SetMultiSelect.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetParentColor(obj uintptr) bool {
    ret, _, _ := listBox_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetParentColor(obj uintptr, value bool) {
   listBox_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := listBox_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetParentCtl3D(obj uintptr, value bool) {
   listBox_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := listBox_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetParentDoubleBuffered(obj uintptr, value bool) {
   listBox_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetParentFont(obj uintptr) bool {
    ret, _, _ := listBox_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetParentFont(obj uintptr, value bool) {
   listBox_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := listBox_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetParentShowHint(obj uintptr, value bool) {
   listBox_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := listBox_GetPopupMenu.Call(obj)
    return ret
}

func ListBox_SetPopupMenu(obj uintptr, value uintptr) {
   listBox_SetPopupMenu.Call(obj, value)
}

func ListBox_GetShowHint(obj uintptr) bool {
    ret, _, _ := listBox_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetShowHint(obj uintptr, value bool) {
   listBox_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetSorted(obj uintptr) bool {
    ret, _, _ := listBox_GetSorted.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetSorted(obj uintptr, value bool) {
   listBox_SetSorted.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := listBox_GetTabOrder.Call(obj)
    return uint16(ret)
}

func ListBox_SetTabOrder(obj uintptr, value uint16) {
   listBox_SetTabOrder.Call(obj, uintptr(value))
}

func ListBox_GetTabStop(obj uintptr) bool {
    ret, _, _ := listBox_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetTabStop(obj uintptr, value bool) {
   listBox_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetTabWidth(obj uintptr) int32 {
    ret, _, _ := listBox_GetTabWidth.Call(obj)
    return int32(ret)
}

func ListBox_SetTabWidth(obj uintptr, value int32) {
   listBox_SetTabWidth.Call(obj, uintptr(value))
}

func ListBox_GetVisible(obj uintptr) bool {
    ret, _, _ := listBox_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetVisible(obj uintptr, value bool) {
   listBox_SetVisible.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := listBox_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func ListBox_SetStyleElements(obj uintptr, value TStyleElements) {
   listBox_SetStyleElements.Call(obj, uintptr(value))
}

func ListBox_SetOnClick(obj uintptr, fn interface{}) {
    listBox_SetOnClick.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnDblClick(obj uintptr, fn interface{}) {
    listBox_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnDrawItem(obj uintptr, fn interface{}) {
    listBox_SetOnDrawItem.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnEnter(obj uintptr, fn interface{}) {
    listBox_SetOnEnter.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnExit(obj uintptr, fn interface{}) {
    listBox_SetOnExit.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnKeyDown(obj uintptr, fn interface{}) {
    listBox_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnKeyPress(obj uintptr, fn interface{}) {
    listBox_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnKeyUp(obj uintptr, fn interface{}) {
    listBox_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnMouseDown(obj uintptr, fn interface{}) {
    listBox_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
    listBox_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
    listBox_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnMouseMove(obj uintptr, fn interface{}) {
    listBox_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func ListBox_SetOnMouseUp(obj uintptr, fn interface{}) {
    listBox_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func ListBox_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := listBox_GetCanvas.Call(obj)
    return ret
}

func ListBox_GetSelCount(obj uintptr) int32 {
    ret, _, _ := listBox_GetSelCount.Call(obj)
    return int32(ret)
}

func ListBox_GetItemIndex(obj uintptr) int32 {
    ret, _, _ := listBox_GetItemIndex.Call(obj)
    return int32(ret)
}

func ListBox_SetItemIndex(obj uintptr, value int32) {
   listBox_SetItemIndex.Call(obj, uintptr(value))
}

func ListBox_GetBrush(obj uintptr) uintptr {
    ret, _, _ := listBox_GetBrush.Call(obj)
    return ret
}

func ListBox_GetControlCount(obj uintptr) int32 {
    ret, _, _ := listBox_GetControlCount.Call(obj)
    return int32(ret)
}

func ListBox_GetHandle(obj uintptr) HWND {
    ret, _, _ := listBox_GetHandle.Call(obj)
    return HWND(ret)
}

func ListBox_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := listBox_GetParentWindow.Call(obj)
    return HWND(ret)
}

func ListBox_SetParentWindow(obj uintptr, value HWND) {
   listBox_SetParentWindow.Call(obj, uintptr(value))
}

func ListBox_GetAction(obj uintptr) uintptr {
    ret, _, _ := listBox_GetAction.Call(obj)
    return ret
}

func ListBox_SetAction(obj uintptr, value uintptr) {
   listBox_SetAction.Call(obj, value)
}

func ListBox_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    listBox_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ListBox_SetBoundsRect(obj uintptr, value TRect) {
   listBox_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ListBox_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := listBox_GetClientHeight.Call(obj)
    return int32(ret)
}

func ListBox_SetClientHeight(obj uintptr, value int32) {
   listBox_SetClientHeight.Call(obj, uintptr(value))
}

func ListBox_GetClientRect(obj uintptr) TRect {
    var ret TRect
    listBox_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ListBox_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := listBox_GetClientWidth.Call(obj)
    return int32(ret)
}

func ListBox_SetClientWidth(obj uintptr, value int32) {
   listBox_SetClientWidth.Call(obj, uintptr(value))
}

func ListBox_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := listBox_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func ListBox_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := listBox_GetExplicitTop.Call(obj)
    return int32(ret)
}

func ListBox_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := listBox_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func ListBox_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := listBox_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func ListBox_GetParent(obj uintptr) uintptr {
    ret, _, _ := listBox_GetParent.Call(obj)
    return ret
}

func ListBox_SetParent(obj uintptr, value uintptr) {
   listBox_SetParent.Call(obj, value)
}

func ListBox_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := listBox_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func ListBox_SetAlignWithMargins(obj uintptr, value bool) {
   listBox_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func ListBox_GetLeft(obj uintptr) int32 {
    ret, _, _ := listBox_GetLeft.Call(obj)
    return int32(ret)
}

func ListBox_SetLeft(obj uintptr, value int32) {
   listBox_SetLeft.Call(obj, uintptr(value))
}

func ListBox_GetTop(obj uintptr) int32 {
    ret, _, _ := listBox_GetTop.Call(obj)
    return int32(ret)
}

func ListBox_SetTop(obj uintptr, value int32) {
   listBox_SetTop.Call(obj, uintptr(value))
}

func ListBox_GetWidth(obj uintptr) int32 {
    ret, _, _ := listBox_GetWidth.Call(obj)
    return int32(ret)
}

func ListBox_SetWidth(obj uintptr, value int32) {
   listBox_SetWidth.Call(obj, uintptr(value))
}

func ListBox_GetHeight(obj uintptr) int32 {
    ret, _, _ := listBox_GetHeight.Call(obj)
    return int32(ret)
}

func ListBox_SetHeight(obj uintptr, value int32) {
   listBox_SetHeight.Call(obj, uintptr(value))
}

func ListBox_GetCursor(obj uintptr) TCursor {
    ret, _, _ := listBox_GetCursor.Call(obj)
    return TCursor(ret)
}

func ListBox_SetCursor(obj uintptr, value TCursor) {
   listBox_SetCursor.Call(obj, uintptr(value))
}

func ListBox_GetHint(obj uintptr) string {
    ret, _, _ := listBox_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func ListBox_SetHint(obj uintptr, value string) {
   listBox_SetHint.Call(obj, GoStrToDStr(value))
}

func ListBox_GetMargins(obj uintptr) uintptr {
    ret, _, _ := listBox_GetMargins.Call(obj)
    return ret
}

func ListBox_SetMargins(obj uintptr, value uintptr) {
   listBox_SetMargins.Call(obj, value)
}

func ListBox_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := listBox_GetComponentCount.Call(obj)
    return int32(ret)
}

func ListBox_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := listBox_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ListBox_SetComponentIndex(obj uintptr, value int32) {
   listBox_SetComponentIndex.Call(obj, uintptr(value))
}

func ListBox_GetOwner(obj uintptr) uintptr {
    ret, _, _ := listBox_GetOwner.Call(obj)
    return ret
}

func ListBox_GetName(obj uintptr) string {
    ret, _, _ := listBox_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ListBox_SetName(obj uintptr, value string) {
   listBox_SetName.Call(obj, GoStrToDStr(value))
}

func ListBox_GetTag(obj uintptr) int {
    ret, _, _ := listBox_GetTag.Call(obj)
    return int(ret)
}

func ListBox_SetTag(obj uintptr, value int) {
   listBox_SetTag.Call(obj, uintptr(value))
}

func ListBox_GetSelected(obj uintptr, Index int32) bool {
    ret, _, _ := listBox_GetSelected.Call(obj, uintptr(Index))
    return DBoolToGoBool(ret)
}

func ListBox_SetSelected(obj uintptr, Index int32, value bool) {
   listBox_SetSelected.Call(obj, uintptr(Index), GoBoolToDBool(value))
}

func ListBox_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := listBox_GetControls.Call(obj, uintptr(Index))
    return ret
}

func ListBox_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := listBox_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TComboBox ---------------------------

func ComboBox_Create(obj uintptr) uintptr {
    ret, _, _ := comboBox_Create.Call(obj)
    return ret
}

func ComboBox_Free(obj uintptr) {
    comboBox_Free.Call(obj)
}

func ComboBox_AddItem(obj uintptr, Item string, AObject uintptr)  {
    comboBox_AddItem.Call(obj, GoStrToDStr(Item) , AObject )
}

func ComboBox_Clear(obj uintptr)  {
    comboBox_Clear.Call(obj)
}

func ComboBox_ClearSelection(obj uintptr)  {
    comboBox_ClearSelection.Call(obj)
}

func ComboBox_DeleteSelected(obj uintptr)  {
    comboBox_DeleteSelected.Call(obj)
}

func ComboBox_Focused(obj uintptr) bool {
    ret, _, _ := comboBox_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SelectAll(obj uintptr)  {
    comboBox_SelectAll.Call(obj)
}

func ComboBox_CanFocus(obj uintptr) bool {
    ret, _, _ := comboBox_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_FlipChildren(obj uintptr, AllLevels bool)  {
    comboBox_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func ComboBox_HandleAllocated(obj uintptr) bool {
    ret, _, _ := comboBox_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_Invalidate(obj uintptr)  {
    comboBox_Invalidate.Call(obj)
}

func ComboBox_Realign(obj uintptr)  {
    comboBox_Realign.Call(obj)
}

func ComboBox_Repaint(obj uintptr)  {
    comboBox_Repaint.Call(obj)
}

func ComboBox_ScaleBy(obj uintptr, M int32, D int32)  {
    comboBox_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func ComboBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    comboBox_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func ComboBox_SetFocus(obj uintptr)  {
    comboBox_SetFocus.Call(obj)
}

func ComboBox_Update(obj uintptr)  {
    comboBox_Update.Call(obj)
}

func ComboBox_BringToFront(obj uintptr)  {
    comboBox_BringToFront.Call(obj)
}

func ComboBox_HasParent(obj uintptr) bool {
    ret, _, _ := comboBox_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_Hide(obj uintptr)  {
    comboBox_Hide.Call(obj)
}

func ComboBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := comboBox_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func ComboBox_Refresh(obj uintptr)  {
    comboBox_Refresh.Call(obj)
}

func ComboBox_SendToBack(obj uintptr)  {
    comboBox_SendToBack.Call(obj)
}

func ComboBox_Show(obj uintptr)  {
    comboBox_Show.Call(obj)
}

func ComboBox_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := comboBox_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func ComboBox_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := comboBox_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ComboBox_GetNamePath(obj uintptr) string {
    ret, _, _ := comboBox_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ComboBox_Assign(obj uintptr, Source uintptr)  {
    comboBox_Assign.Call(obj, Source )
}

func ComboBox_ClassName(obj uintptr) string {
    ret, _, _ := comboBox_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ComboBox_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := comboBox_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ComboBox_GetHashCode(obj uintptr) int32 {
    ret, _, _ := comboBox_GetHashCode.Call(obj)
    return int32(ret)
}

func ComboBox_ToString(obj uintptr) string {
    ret, _, _ := comboBox_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ComboBox_GetAlign(obj uintptr) TAlign {
    ret, _, _ := comboBox_GetAlign.Call(obj)
    return TAlign(ret)
}

func ComboBox_SetAlign(obj uintptr, value TAlign) {
   comboBox_SetAlign.Call(obj, uintptr(value))
}

func ComboBox_GetAutoComplete(obj uintptr) bool {
    ret, _, _ := comboBox_GetAutoComplete.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetAutoComplete(obj uintptr, value bool) {
   comboBox_SetAutoComplete.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetAutoCompleteDelay(obj uintptr) uint32 {
    ret, _, _ := comboBox_GetAutoCompleteDelay.Call(obj)
    return uint32(ret)
}

func ComboBox_SetAutoCompleteDelay(obj uintptr, value uint32) {
   comboBox_SetAutoCompleteDelay.Call(obj, uintptr(value))
}

func ComboBox_GetAutoDropDown(obj uintptr) bool {
    ret, _, _ := comboBox_GetAutoDropDown.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetAutoDropDown(obj uintptr, value bool) {
   comboBox_SetAutoDropDown.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetAutoCloseUp(obj uintptr) bool {
    ret, _, _ := comboBox_GetAutoCloseUp.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetAutoCloseUp(obj uintptr, value bool) {
   comboBox_SetAutoCloseUp.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetBevelEdges(obj uintptr) TBevelEdges {
    ret, _, _ := comboBox_GetBevelEdges.Call(obj)
    return TBevelEdges(ret)
}

func ComboBox_SetBevelEdges(obj uintptr, value TBevelEdges) {
   comboBox_SetBevelEdges.Call(obj, uintptr(value))
}

func ComboBox_GetBevelInner(obj uintptr) TBevelCut {
    ret, _, _ := comboBox_GetBevelInner.Call(obj)
    return TBevelCut(ret)
}

func ComboBox_SetBevelInner(obj uintptr, value TBevelCut) {
   comboBox_SetBevelInner.Call(obj, uintptr(value))
}

func ComboBox_GetBevelKind(obj uintptr) TBevelKind {
    ret, _, _ := comboBox_GetBevelKind.Call(obj)
    return TBevelKind(ret)
}

func ComboBox_SetBevelKind(obj uintptr, value TBevelKind) {
   comboBox_SetBevelKind.Call(obj, uintptr(value))
}

func ComboBox_GetBevelOuter(obj uintptr) TBevelCut {
    ret, _, _ := comboBox_GetBevelOuter.Call(obj)
    return TBevelCut(ret)
}

func ComboBox_SetBevelOuter(obj uintptr, value TBevelCut) {
   comboBox_SetBevelOuter.Call(obj, uintptr(value))
}

func ComboBox_GetStyle(obj uintptr) TComboBoxStyle {
    ret, _, _ := comboBox_GetStyle.Call(obj)
    return TComboBoxStyle(ret)
}

func ComboBox_SetStyle(obj uintptr, value TComboBoxStyle) {
   comboBox_SetStyle.Call(obj, uintptr(value))
}

func ComboBox_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := comboBox_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func ComboBox_SetAnchors(obj uintptr, value TAnchors) {
   comboBox_SetAnchors.Call(obj, uintptr(value))
}

func ComboBox_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := comboBox_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func ComboBox_SetBiDiMode(obj uintptr, value TBiDiMode) {
   comboBox_SetBiDiMode.Call(obj, uintptr(value))
}

func ComboBox_GetColor(obj uintptr) TColor {
    ret, _, _ := comboBox_GetColor.Call(obj)
    return TColor(ret)
}

func ComboBox_SetColor(obj uintptr, value TColor) {
   comboBox_SetColor.Call(obj, uintptr(value))
}

func ComboBox_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := comboBox_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetDoubleBuffered(obj uintptr, value bool) {
   comboBox_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetDropDownCount(obj uintptr) int32 {
    ret, _, _ := comboBox_GetDropDownCount.Call(obj)
    return int32(ret)
}

func ComboBox_SetDropDownCount(obj uintptr, value int32) {
   comboBox_SetDropDownCount.Call(obj, uintptr(value))
}

func ComboBox_GetEnabled(obj uintptr) bool {
    ret, _, _ := comboBox_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetEnabled(obj uintptr, value bool) {
   comboBox_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetFont(obj uintptr) uintptr {
    ret, _, _ := comboBox_GetFont.Call(obj)
    return ret
}

func ComboBox_SetFont(obj uintptr, value uintptr) {
   comboBox_SetFont.Call(obj, value)
}

func ComboBox_GetItemHeight(obj uintptr) int32 {
    ret, _, _ := comboBox_GetItemHeight.Call(obj)
    return int32(ret)
}

func ComboBox_SetItemHeight(obj uintptr, value int32) {
   comboBox_SetItemHeight.Call(obj, uintptr(value))
}

func ComboBox_GetItemIndex(obj uintptr) int32 {
    ret, _, _ := comboBox_GetItemIndex.Call(obj)
    return int32(ret)
}

func ComboBox_SetItemIndex(obj uintptr, value int32) {
   comboBox_SetItemIndex.Call(obj, uintptr(value))
}

func ComboBox_GetMaxLength(obj uintptr) int32 {
    ret, _, _ := comboBox_GetMaxLength.Call(obj)
    return int32(ret)
}

func ComboBox_SetMaxLength(obj uintptr, value int32) {
   comboBox_SetMaxLength.Call(obj, uintptr(value))
}

func ComboBox_GetParentColor(obj uintptr) bool {
    ret, _, _ := comboBox_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetParentColor(obj uintptr, value bool) {
   comboBox_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := comboBox_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetParentCtl3D(obj uintptr, value bool) {
   comboBox_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := comboBox_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetParentDoubleBuffered(obj uintptr, value bool) {
   comboBox_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetParentFont(obj uintptr) bool {
    ret, _, _ := comboBox_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetParentFont(obj uintptr, value bool) {
   comboBox_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := comboBox_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetParentShowHint(obj uintptr, value bool) {
   comboBox_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := comboBox_GetPopupMenu.Call(obj)
    return ret
}

func ComboBox_SetPopupMenu(obj uintptr, value uintptr) {
   comboBox_SetPopupMenu.Call(obj, value)
}

func ComboBox_GetShowHint(obj uintptr) bool {
    ret, _, _ := comboBox_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetShowHint(obj uintptr, value bool) {
   comboBox_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetSorted(obj uintptr) bool {
    ret, _, _ := comboBox_GetSorted.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetSorted(obj uintptr, value bool) {
   comboBox_SetSorted.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := comboBox_GetTabOrder.Call(obj)
    return uint16(ret)
}

func ComboBox_SetTabOrder(obj uintptr, value uint16) {
   comboBox_SetTabOrder.Call(obj, uintptr(value))
}

func ComboBox_GetTabStop(obj uintptr) bool {
    ret, _, _ := comboBox_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetTabStop(obj uintptr, value bool) {
   comboBox_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetText(obj uintptr) string {
    ret, _, _ := comboBox_GetText.Call(obj)
    return DStrToGoStr(ret)
}

func ComboBox_SetText(obj uintptr, value string) {
   comboBox_SetText.Call(obj, GoStrToDStr(value))
}

func ComboBox_GetTextHint(obj uintptr) string {
    ret, _, _ := comboBox_GetTextHint.Call(obj)
    return DStrToGoStr(ret)
}

func ComboBox_SetTextHint(obj uintptr, value string) {
   comboBox_SetTextHint.Call(obj, GoStrToDStr(value))
}

func ComboBox_GetVisible(obj uintptr) bool {
    ret, _, _ := comboBox_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetVisible(obj uintptr, value bool) {
   comboBox_SetVisible.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := comboBox_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func ComboBox_SetStyleElements(obj uintptr, value TStyleElements) {
   comboBox_SetStyleElements.Call(obj, uintptr(value))
}

func ComboBox_SetOnChange(obj uintptr, fn interface{}) {
    comboBox_SetOnChange.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnClick(obj uintptr, fn interface{}) {
    comboBox_SetOnClick.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnDblClick(obj uintptr, fn interface{}) {
    comboBox_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnDrawItem(obj uintptr, fn interface{}) {
    comboBox_SetOnDrawItem.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnEnter(obj uintptr, fn interface{}) {
    comboBox_SetOnEnter.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnExit(obj uintptr, fn interface{}) {
    comboBox_SetOnExit.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnKeyDown(obj uintptr, fn interface{}) {
    comboBox_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnKeyPress(obj uintptr, fn interface{}) {
    comboBox_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnKeyUp(obj uintptr, fn interface{}) {
    comboBox_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
    comboBox_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func ComboBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
    comboBox_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func ComboBox_GetItems(obj uintptr) uintptr {
    ret, _, _ := comboBox_GetItems.Call(obj)
    return ret
}

func ComboBox_SetItems(obj uintptr, value uintptr) {
   comboBox_SetItems.Call(obj, value)
}

func ComboBox_GetSelText(obj uintptr) string {
    ret, _, _ := comboBox_GetSelText.Call(obj)
    return DStrToGoStr(ret)
}

func ComboBox_SetSelText(obj uintptr, value string) {
   comboBox_SetSelText.Call(obj, GoStrToDStr(value))
}

func ComboBox_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := comboBox_GetCanvas.Call(obj)
    return ret
}

func ComboBox_GetDroppedDown(obj uintptr) bool {
    ret, _, _ := comboBox_GetDroppedDown.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetDroppedDown(obj uintptr, value bool) {
   comboBox_SetDroppedDown.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetSelLength(obj uintptr) int32 {
    ret, _, _ := comboBox_GetSelLength.Call(obj)
    return int32(ret)
}

func ComboBox_SetSelLength(obj uintptr, value int32) {
   comboBox_SetSelLength.Call(obj, uintptr(value))
}

func ComboBox_GetSelStart(obj uintptr) int32 {
    ret, _, _ := comboBox_GetSelStart.Call(obj)
    return int32(ret)
}

func ComboBox_SetSelStart(obj uintptr, value int32) {
   comboBox_SetSelStart.Call(obj, uintptr(value))
}

func ComboBox_GetBrush(obj uintptr) uintptr {
    ret, _, _ := comboBox_GetBrush.Call(obj)
    return ret
}

func ComboBox_GetControlCount(obj uintptr) int32 {
    ret, _, _ := comboBox_GetControlCount.Call(obj)
    return int32(ret)
}

func ComboBox_GetHandle(obj uintptr) HWND {
    ret, _, _ := comboBox_GetHandle.Call(obj)
    return HWND(ret)
}

func ComboBox_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := comboBox_GetParentWindow.Call(obj)
    return HWND(ret)
}

func ComboBox_SetParentWindow(obj uintptr, value HWND) {
   comboBox_SetParentWindow.Call(obj, uintptr(value))
}

func ComboBox_GetAction(obj uintptr) uintptr {
    ret, _, _ := comboBox_GetAction.Call(obj)
    return ret
}

func ComboBox_SetAction(obj uintptr, value uintptr) {
   comboBox_SetAction.Call(obj, value)
}

func ComboBox_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    comboBox_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ComboBox_SetBoundsRect(obj uintptr, value TRect) {
   comboBox_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ComboBox_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := comboBox_GetClientHeight.Call(obj)
    return int32(ret)
}

func ComboBox_SetClientHeight(obj uintptr, value int32) {
   comboBox_SetClientHeight.Call(obj, uintptr(value))
}

func ComboBox_GetClientRect(obj uintptr) TRect {
    var ret TRect
    comboBox_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ComboBox_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := comboBox_GetClientWidth.Call(obj)
    return int32(ret)
}

func ComboBox_SetClientWidth(obj uintptr, value int32) {
   comboBox_SetClientWidth.Call(obj, uintptr(value))
}

func ComboBox_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := comboBox_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func ComboBox_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := comboBox_GetExplicitTop.Call(obj)
    return int32(ret)
}

func ComboBox_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := comboBox_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func ComboBox_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := comboBox_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func ComboBox_GetParent(obj uintptr) uintptr {
    ret, _, _ := comboBox_GetParent.Call(obj)
    return ret
}

func ComboBox_SetParent(obj uintptr, value uintptr) {
   comboBox_SetParent.Call(obj, value)
}

func ComboBox_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := comboBox_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func ComboBox_SetAlignWithMargins(obj uintptr, value bool) {
   comboBox_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func ComboBox_GetLeft(obj uintptr) int32 {
    ret, _, _ := comboBox_GetLeft.Call(obj)
    return int32(ret)
}

func ComboBox_SetLeft(obj uintptr, value int32) {
   comboBox_SetLeft.Call(obj, uintptr(value))
}

func ComboBox_GetTop(obj uintptr) int32 {
    ret, _, _ := comboBox_GetTop.Call(obj)
    return int32(ret)
}

func ComboBox_SetTop(obj uintptr, value int32) {
   comboBox_SetTop.Call(obj, uintptr(value))
}

func ComboBox_GetWidth(obj uintptr) int32 {
    ret, _, _ := comboBox_GetWidth.Call(obj)
    return int32(ret)
}

func ComboBox_SetWidth(obj uintptr, value int32) {
   comboBox_SetWidth.Call(obj, uintptr(value))
}

func ComboBox_GetHeight(obj uintptr) int32 {
    ret, _, _ := comboBox_GetHeight.Call(obj)
    return int32(ret)
}

func ComboBox_SetHeight(obj uintptr, value int32) {
   comboBox_SetHeight.Call(obj, uintptr(value))
}

func ComboBox_GetCursor(obj uintptr) TCursor {
    ret, _, _ := comboBox_GetCursor.Call(obj)
    return TCursor(ret)
}

func ComboBox_SetCursor(obj uintptr, value TCursor) {
   comboBox_SetCursor.Call(obj, uintptr(value))
}

func ComboBox_GetHint(obj uintptr) string {
    ret, _, _ := comboBox_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func ComboBox_SetHint(obj uintptr, value string) {
   comboBox_SetHint.Call(obj, GoStrToDStr(value))
}

func ComboBox_GetMargins(obj uintptr) uintptr {
    ret, _, _ := comboBox_GetMargins.Call(obj)
    return ret
}

func ComboBox_SetMargins(obj uintptr, value uintptr) {
   comboBox_SetMargins.Call(obj, value)
}

func ComboBox_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := comboBox_GetComponentCount.Call(obj)
    return int32(ret)
}

func ComboBox_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := comboBox_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ComboBox_SetComponentIndex(obj uintptr, value int32) {
   comboBox_SetComponentIndex.Call(obj, uintptr(value))
}

func ComboBox_GetOwner(obj uintptr) uintptr {
    ret, _, _ := comboBox_GetOwner.Call(obj)
    return ret
}

func ComboBox_GetName(obj uintptr) string {
    ret, _, _ := comboBox_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ComboBox_SetName(obj uintptr, value string) {
   comboBox_SetName.Call(obj, GoStrToDStr(value))
}

func ComboBox_GetTag(obj uintptr) int {
    ret, _, _ := comboBox_GetTag.Call(obj)
    return int(ret)
}

func ComboBox_SetTag(obj uintptr, value int) {
   comboBox_SetTag.Call(obj, uintptr(value))
}

func ComboBox_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := comboBox_GetControls.Call(obj, uintptr(Index))
    return ret
}

func ComboBox_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := comboBox_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TPanel ---------------------------

func Panel_Create(obj uintptr) uintptr {
    ret, _, _ := panel_Create.Call(obj)
    return ret
}

func Panel_Free(obj uintptr) {
    panel_Free.Call(obj)
}

func Panel_CanFocus(obj uintptr) bool {
    ret, _, _ := panel_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_FlipChildren(obj uintptr, AllLevels bool)  {
    panel_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func Panel_Focused(obj uintptr) bool {
    ret, _, _ := panel_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_HandleAllocated(obj uintptr) bool {
    ret, _, _ := panel_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_Invalidate(obj uintptr)  {
    panel_Invalidate.Call(obj)
}

func Panel_Realign(obj uintptr)  {
    panel_Realign.Call(obj)
}

func Panel_Repaint(obj uintptr)  {
    panel_Repaint.Call(obj)
}

func Panel_ScaleBy(obj uintptr, M int32, D int32)  {
    panel_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func Panel_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    panel_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func Panel_SetFocus(obj uintptr)  {
    panel_SetFocus.Call(obj)
}

func Panel_Update(obj uintptr)  {
    panel_Update.Call(obj)
}

func Panel_BringToFront(obj uintptr)  {
    panel_BringToFront.Call(obj)
}

func Panel_HasParent(obj uintptr) bool {
    ret, _, _ := panel_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_Hide(obj uintptr)  {
    panel_Hide.Call(obj)
}

func Panel_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := panel_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func Panel_Refresh(obj uintptr)  {
    panel_Refresh.Call(obj)
}

func Panel_SendToBack(obj uintptr)  {
    panel_SendToBack.Call(obj)
}

func Panel_Show(obj uintptr)  {
    panel_Show.Call(obj)
}

func Panel_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := panel_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Panel_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := panel_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Panel_GetNamePath(obj uintptr) string {
    ret, _, _ := panel_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Panel_Assign(obj uintptr, Source uintptr)  {
    panel_Assign.Call(obj, Source )
}

func Panel_ClassName(obj uintptr) string {
    ret, _, _ := panel_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Panel_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := panel_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Panel_GetHashCode(obj uintptr) int32 {
    ret, _, _ := panel_GetHashCode.Call(obj)
    return int32(ret)
}

func Panel_ToString(obj uintptr) string {
    ret, _, _ := panel_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Panel_GetAlign(obj uintptr) TAlign {
    ret, _, _ := panel_GetAlign.Call(obj)
    return TAlign(ret)
}

func Panel_SetAlign(obj uintptr, value TAlign) {
   panel_SetAlign.Call(obj, uintptr(value))
}

func Panel_GetAlignment(obj uintptr) TAlignment {
    ret, _, _ := panel_GetAlignment.Call(obj)
    return TAlignment(ret)
}

func Panel_SetAlignment(obj uintptr, value TAlignment) {
   panel_SetAlignment.Call(obj, uintptr(value))
}

func Panel_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := panel_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func Panel_SetAnchors(obj uintptr, value TAnchors) {
   panel_SetAnchors.Call(obj, uintptr(value))
}

func Panel_GetAutoSize(obj uintptr) bool {
    ret, _, _ := panel_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetAutoSize(obj uintptr, value bool) {
   panel_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func Panel_GetBevelEdges(obj uintptr) TBevelEdges {
    ret, _, _ := panel_GetBevelEdges.Call(obj)
    return TBevelEdges(ret)
}

func Panel_SetBevelEdges(obj uintptr, value TBevelEdges) {
   panel_SetBevelEdges.Call(obj, uintptr(value))
}

func Panel_GetBevelInner(obj uintptr) TBevelCut {
    ret, _, _ := panel_GetBevelInner.Call(obj)
    return TBevelCut(ret)
}

func Panel_SetBevelInner(obj uintptr, value TBevelCut) {
   panel_SetBevelInner.Call(obj, uintptr(value))
}

func Panel_GetBevelKind(obj uintptr) TBevelKind {
    ret, _, _ := panel_GetBevelKind.Call(obj)
    return TBevelKind(ret)
}

func Panel_SetBevelKind(obj uintptr, value TBevelKind) {
   panel_SetBevelKind.Call(obj, uintptr(value))
}

func Panel_GetBevelOuter(obj uintptr) TBevelCut {
    ret, _, _ := panel_GetBevelOuter.Call(obj)
    return TBevelCut(ret)
}

func Panel_SetBevelOuter(obj uintptr, value TBevelCut) {
   panel_SetBevelOuter.Call(obj, uintptr(value))
}

func Panel_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := panel_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func Panel_SetBiDiMode(obj uintptr, value TBiDiMode) {
   panel_SetBiDiMode.Call(obj, uintptr(value))
}

func Panel_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := panel_GetBorderWidth.Call(obj)
    return int32(ret)
}

func Panel_SetBorderWidth(obj uintptr, value int32) {
   panel_SetBorderWidth.Call(obj, uintptr(value))
}

func Panel_GetBorderStyle(obj uintptr) TBorderStyle {
    ret, _, _ := panel_GetBorderStyle.Call(obj)
    return TBorderStyle(ret)
}

func Panel_SetBorderStyle(obj uintptr, value TBorderStyle) {
   panel_SetBorderStyle.Call(obj, uintptr(value))
}

func Panel_GetCaption(obj uintptr) string {
    ret, _, _ := panel_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func Panel_SetCaption(obj uintptr, value string) {
   panel_SetCaption.Call(obj, GoStrToDStr(value))
}

func Panel_GetColor(obj uintptr) TColor {
    ret, _, _ := panel_GetColor.Call(obj)
    return TColor(ret)
}

func Panel_SetColor(obj uintptr, value TColor) {
   panel_SetColor.Call(obj, uintptr(value))
}

func Panel_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := panel_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetDoubleBuffered(obj uintptr, value bool) {
   panel_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func Panel_GetEnabled(obj uintptr) bool {
    ret, _, _ := panel_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetEnabled(obj uintptr, value bool) {
   panel_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Panel_GetFullRepaint(obj uintptr) bool {
    ret, _, _ := panel_GetFullRepaint.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetFullRepaint(obj uintptr, value bool) {
   panel_SetFullRepaint.Call(obj, GoBoolToDBool(value))
}

func Panel_GetFont(obj uintptr) uintptr {
    ret, _, _ := panel_GetFont.Call(obj)
    return ret
}

func Panel_SetFont(obj uintptr, value uintptr) {
   panel_SetFont.Call(obj, value)
}

func Panel_GetLocked(obj uintptr) bool {
    ret, _, _ := panel_GetLocked.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetLocked(obj uintptr, value bool) {
   panel_SetLocked.Call(obj, GoBoolToDBool(value))
}

func Panel_GetParentBackground(obj uintptr) bool {
    ret, _, _ := panel_GetParentBackground.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetParentBackground(obj uintptr, value bool) {
   panel_SetParentBackground.Call(obj, GoBoolToDBool(value))
}

func Panel_GetParentColor(obj uintptr) bool {
    ret, _, _ := panel_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetParentColor(obj uintptr, value bool) {
   panel_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func Panel_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := panel_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetParentCtl3D(obj uintptr, value bool) {
   panel_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func Panel_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := panel_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetParentDoubleBuffered(obj uintptr, value bool) {
   panel_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func Panel_GetParentFont(obj uintptr) bool {
    ret, _, _ := panel_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetParentFont(obj uintptr, value bool) {
   panel_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func Panel_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := panel_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetParentShowHint(obj uintptr, value bool) {
   panel_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func Panel_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := panel_GetPopupMenu.Call(obj)
    return ret
}

func Panel_SetPopupMenu(obj uintptr, value uintptr) {
   panel_SetPopupMenu.Call(obj, value)
}

func Panel_GetShowCaption(obj uintptr) bool {
    ret, _, _ := panel_GetShowCaption.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetShowCaption(obj uintptr, value bool) {
   panel_SetShowCaption.Call(obj, GoBoolToDBool(value))
}

func Panel_GetShowHint(obj uintptr) bool {
    ret, _, _ := panel_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetShowHint(obj uintptr, value bool) {
   panel_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Panel_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := panel_GetTabOrder.Call(obj)
    return uint16(ret)
}

func Panel_SetTabOrder(obj uintptr, value uint16) {
   panel_SetTabOrder.Call(obj, uintptr(value))
}

func Panel_GetTabStop(obj uintptr) bool {
    ret, _, _ := panel_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetTabStop(obj uintptr, value bool) {
   panel_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func Panel_GetVisible(obj uintptr) bool {
    ret, _, _ := panel_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetVisible(obj uintptr, value bool) {
   panel_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Panel_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := panel_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func Panel_SetStyleElements(obj uintptr, value TStyleElements) {
   panel_SetStyleElements.Call(obj, uintptr(value))
}

func Panel_SetOnClick(obj uintptr, fn interface{}) {
    panel_SetOnClick.Call(obj, addEventToMap(fn))
}

func Panel_SetOnDblClick(obj uintptr, fn interface{}) {
    panel_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func Panel_SetOnEnter(obj uintptr, fn interface{}) {
    panel_SetOnEnter.Call(obj, addEventToMap(fn))
}

func Panel_SetOnExit(obj uintptr, fn interface{}) {
    panel_SetOnExit.Call(obj, addEventToMap(fn))
}

func Panel_SetOnMouseDown(obj uintptr, fn interface{}) {
    panel_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func Panel_SetOnMouseEnter(obj uintptr, fn interface{}) {
    panel_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func Panel_SetOnMouseLeave(obj uintptr, fn interface{}) {
    panel_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func Panel_SetOnMouseMove(obj uintptr, fn interface{}) {
    panel_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func Panel_SetOnMouseUp(obj uintptr, fn interface{}) {
    panel_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func Panel_SetOnResize(obj uintptr, fn interface{}) {
    panel_SetOnResize.Call(obj, addEventToMap(fn))
}

func Panel_GetBrush(obj uintptr) uintptr {
    ret, _, _ := panel_GetBrush.Call(obj)
    return ret
}

func Panel_GetControlCount(obj uintptr) int32 {
    ret, _, _ := panel_GetControlCount.Call(obj)
    return int32(ret)
}

func Panel_GetHandle(obj uintptr) HWND {
    ret, _, _ := panel_GetHandle.Call(obj)
    return HWND(ret)
}

func Panel_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := panel_GetParentWindow.Call(obj)
    return HWND(ret)
}

func Panel_SetParentWindow(obj uintptr, value HWND) {
   panel_SetParentWindow.Call(obj, uintptr(value))
}

func Panel_GetAction(obj uintptr) uintptr {
    ret, _, _ := panel_GetAction.Call(obj)
    return ret
}

func Panel_SetAction(obj uintptr, value uintptr) {
   panel_SetAction.Call(obj, value)
}

func Panel_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    panel_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Panel_SetBoundsRect(obj uintptr, value TRect) {
   panel_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Panel_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := panel_GetClientHeight.Call(obj)
    return int32(ret)
}

func Panel_SetClientHeight(obj uintptr, value int32) {
   panel_SetClientHeight.Call(obj, uintptr(value))
}

func Panel_GetClientRect(obj uintptr) TRect {
    var ret TRect
    panel_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Panel_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := panel_GetClientWidth.Call(obj)
    return int32(ret)
}

func Panel_SetClientWidth(obj uintptr, value int32) {
   panel_SetClientWidth.Call(obj, uintptr(value))
}

func Panel_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := panel_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Panel_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := panel_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Panel_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := panel_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Panel_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := panel_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Panel_GetParent(obj uintptr) uintptr {
    ret, _, _ := panel_GetParent.Call(obj)
    return ret
}

func Panel_SetParent(obj uintptr, value uintptr) {
   panel_SetParent.Call(obj, value)
}

func Panel_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := panel_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func Panel_SetAlignWithMargins(obj uintptr, value bool) {
   panel_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func Panel_GetLeft(obj uintptr) int32 {
    ret, _, _ := panel_GetLeft.Call(obj)
    return int32(ret)
}

func Panel_SetLeft(obj uintptr, value int32) {
   panel_SetLeft.Call(obj, uintptr(value))
}

func Panel_GetTop(obj uintptr) int32 {
    ret, _, _ := panel_GetTop.Call(obj)
    return int32(ret)
}

func Panel_SetTop(obj uintptr, value int32) {
   panel_SetTop.Call(obj, uintptr(value))
}

func Panel_GetWidth(obj uintptr) int32 {
    ret, _, _ := panel_GetWidth.Call(obj)
    return int32(ret)
}

func Panel_SetWidth(obj uintptr, value int32) {
   panel_SetWidth.Call(obj, uintptr(value))
}

func Panel_GetHeight(obj uintptr) int32 {
    ret, _, _ := panel_GetHeight.Call(obj)
    return int32(ret)
}

func Panel_SetHeight(obj uintptr, value int32) {
   panel_SetHeight.Call(obj, uintptr(value))
}

func Panel_GetCursor(obj uintptr) TCursor {
    ret, _, _ := panel_GetCursor.Call(obj)
    return TCursor(ret)
}

func Panel_SetCursor(obj uintptr, value TCursor) {
   panel_SetCursor.Call(obj, uintptr(value))
}

func Panel_GetHint(obj uintptr) string {
    ret, _, _ := panel_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Panel_SetHint(obj uintptr, value string) {
   panel_SetHint.Call(obj, GoStrToDStr(value))
}

func Panel_GetMargins(obj uintptr) uintptr {
    ret, _, _ := panel_GetMargins.Call(obj)
    return ret
}

func Panel_SetMargins(obj uintptr, value uintptr) {
   panel_SetMargins.Call(obj, value)
}

func Panel_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := panel_GetComponentCount.Call(obj)
    return int32(ret)
}

func Panel_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := panel_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Panel_SetComponentIndex(obj uintptr, value int32) {
   panel_SetComponentIndex.Call(obj, uintptr(value))
}

func Panel_GetOwner(obj uintptr) uintptr {
    ret, _, _ := panel_GetOwner.Call(obj)
    return ret
}

func Panel_GetName(obj uintptr) string {
    ret, _, _ := panel_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Panel_SetName(obj uintptr, value string) {
   panel_SetName.Call(obj, GoStrToDStr(value))
}

func Panel_GetTag(obj uintptr) int {
    ret, _, _ := panel_GetTag.Call(obj)
    return int(ret)
}

func Panel_SetTag(obj uintptr, value int) {
   panel_SetTag.Call(obj, uintptr(value))
}

func Panel_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := panel_GetControls.Call(obj, uintptr(Index))
    return ret
}

func Panel_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := panel_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TImage ---------------------------

func Image_Create(obj uintptr) uintptr {
    ret, _, _ := image_Create.Call(obj)
    return ret
}

func Image_Free(obj uintptr) {
    image_Free.Call(obj)
}

func Image_BringToFront(obj uintptr)  {
    image_BringToFront.Call(obj)
}

func Image_HasParent(obj uintptr) bool {
    ret, _, _ := image_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_Hide(obj uintptr)  {
    image_Hide.Call(obj)
}

func Image_Invalidate(obj uintptr)  {
    image_Invalidate.Call(obj)
}

func Image_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := image_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func Image_Refresh(obj uintptr)  {
    image_Refresh.Call(obj)
}

func Image_Repaint(obj uintptr)  {
    image_Repaint.Call(obj)
}

func Image_SendToBack(obj uintptr)  {
    image_SendToBack.Call(obj)
}

func Image_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    image_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func Image_Show(obj uintptr)  {
    image_Show.Call(obj)
}

func Image_Update(obj uintptr)  {
    image_Update.Call(obj)
}

func Image_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := image_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Image_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := image_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Image_GetNamePath(obj uintptr) string {
    ret, _, _ := image_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Image_Assign(obj uintptr, Source uintptr)  {
    image_Assign.Call(obj, Source )
}

func Image_ClassName(obj uintptr) string {
    ret, _, _ := image_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Image_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := image_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Image_GetHashCode(obj uintptr) int32 {
    ret, _, _ := image_GetHashCode.Call(obj)
    return int32(ret)
}

func Image_ToString(obj uintptr) string {
    ret, _, _ := image_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Image_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := image_GetCanvas.Call(obj)
    return ret
}

func Image_GetAlign(obj uintptr) TAlign {
    ret, _, _ := image_GetAlign.Call(obj)
    return TAlign(ret)
}

func Image_SetAlign(obj uintptr, value TAlign) {
   image_SetAlign.Call(obj, uintptr(value))
}

func Image_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := image_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func Image_SetAnchors(obj uintptr, value TAnchors) {
   image_SetAnchors.Call(obj, uintptr(value))
}

func Image_GetAutoSize(obj uintptr) bool {
    ret, _, _ := image_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetAutoSize(obj uintptr, value bool) {
   image_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func Image_GetCenter(obj uintptr) bool {
    ret, _, _ := image_GetCenter.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetCenter(obj uintptr, value bool) {
   image_SetCenter.Call(obj, GoBoolToDBool(value))
}

func Image_GetEnabled(obj uintptr) bool {
    ret, _, _ := image_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetEnabled(obj uintptr, value bool) {
   image_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Image_GetIncrementalDisplay(obj uintptr) bool {
    ret, _, _ := image_GetIncrementalDisplay.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetIncrementalDisplay(obj uintptr, value bool) {
   image_SetIncrementalDisplay.Call(obj, GoBoolToDBool(value))
}

func Image_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := image_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetParentShowHint(obj uintptr, value bool) {
   image_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func Image_GetPicture(obj uintptr) uintptr {
    ret, _, _ := image_GetPicture.Call(obj)
    return ret
}

func Image_SetPicture(obj uintptr, value uintptr) {
   image_SetPicture.Call(obj, value)
}

func Image_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := image_GetPopupMenu.Call(obj)
    return ret
}

func Image_SetPopupMenu(obj uintptr, value uintptr) {
   image_SetPopupMenu.Call(obj, value)
}

func Image_GetProportional(obj uintptr) bool {
    ret, _, _ := image_GetProportional.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetProportional(obj uintptr, value bool) {
   image_SetProportional.Call(obj, GoBoolToDBool(value))
}

func Image_GetShowHint(obj uintptr) bool {
    ret, _, _ := image_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetShowHint(obj uintptr, value bool) {
   image_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Image_GetStretch(obj uintptr) bool {
    ret, _, _ := image_GetStretch.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetStretch(obj uintptr, value bool) {
   image_SetStretch.Call(obj, GoBoolToDBool(value))
}

func Image_GetTransparent(obj uintptr) bool {
    ret, _, _ := image_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetTransparent(obj uintptr, value bool) {
   image_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func Image_GetVisible(obj uintptr) bool {
    ret, _, _ := image_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetVisible(obj uintptr, value bool) {
   image_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Image_SetOnClick(obj uintptr, fn interface{}) {
    image_SetOnClick.Call(obj, addEventToMap(fn))
}

func Image_SetOnDblClick(obj uintptr, fn interface{}) {
    image_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func Image_SetOnMouseDown(obj uintptr, fn interface{}) {
    image_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func Image_SetOnMouseEnter(obj uintptr, fn interface{}) {
    image_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func Image_SetOnMouseLeave(obj uintptr, fn interface{}) {
    image_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func Image_SetOnMouseMove(obj uintptr, fn interface{}) {
    image_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func Image_SetOnMouseUp(obj uintptr, fn interface{}) {
    image_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func Image_GetAction(obj uintptr) uintptr {
    ret, _, _ := image_GetAction.Call(obj)
    return ret
}

func Image_SetAction(obj uintptr, value uintptr) {
   image_SetAction.Call(obj, value)
}

func Image_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := image_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func Image_SetBiDiMode(obj uintptr, value TBiDiMode) {
   image_SetBiDiMode.Call(obj, uintptr(value))
}

func Image_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    image_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Image_SetBoundsRect(obj uintptr, value TRect) {
   image_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Image_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := image_GetClientHeight.Call(obj)
    return int32(ret)
}

func Image_SetClientHeight(obj uintptr, value int32) {
   image_SetClientHeight.Call(obj, uintptr(value))
}

func Image_GetClientRect(obj uintptr) TRect {
    var ret TRect
    image_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Image_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := image_GetClientWidth.Call(obj)
    return int32(ret)
}

func Image_SetClientWidth(obj uintptr, value int32) {
   image_SetClientWidth.Call(obj, uintptr(value))
}

func Image_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := image_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Image_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := image_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Image_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := image_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Image_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := image_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Image_GetParent(obj uintptr) uintptr {
    ret, _, _ := image_GetParent.Call(obj)
    return ret
}

func Image_SetParent(obj uintptr, value uintptr) {
   image_SetParent.Call(obj, value)
}

func Image_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := image_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func Image_SetStyleElements(obj uintptr, value TStyleElements) {
   image_SetStyleElements.Call(obj, uintptr(value))
}

func Image_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := image_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func Image_SetAlignWithMargins(obj uintptr, value bool) {
   image_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func Image_GetLeft(obj uintptr) int32 {
    ret, _, _ := image_GetLeft.Call(obj)
    return int32(ret)
}

func Image_SetLeft(obj uintptr, value int32) {
   image_SetLeft.Call(obj, uintptr(value))
}

func Image_GetTop(obj uintptr) int32 {
    ret, _, _ := image_GetTop.Call(obj)
    return int32(ret)
}

func Image_SetTop(obj uintptr, value int32) {
   image_SetTop.Call(obj, uintptr(value))
}

func Image_GetWidth(obj uintptr) int32 {
    ret, _, _ := image_GetWidth.Call(obj)
    return int32(ret)
}

func Image_SetWidth(obj uintptr, value int32) {
   image_SetWidth.Call(obj, uintptr(value))
}

func Image_GetHeight(obj uintptr) int32 {
    ret, _, _ := image_GetHeight.Call(obj)
    return int32(ret)
}

func Image_SetHeight(obj uintptr, value int32) {
   image_SetHeight.Call(obj, uintptr(value))
}

func Image_GetCursor(obj uintptr) TCursor {
    ret, _, _ := image_GetCursor.Call(obj)
    return TCursor(ret)
}

func Image_SetCursor(obj uintptr, value TCursor) {
   image_SetCursor.Call(obj, uintptr(value))
}

func Image_GetHint(obj uintptr) string {
    ret, _, _ := image_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Image_SetHint(obj uintptr, value string) {
   image_SetHint.Call(obj, GoStrToDStr(value))
}

func Image_GetMargins(obj uintptr) uintptr {
    ret, _, _ := image_GetMargins.Call(obj)
    return ret
}

func Image_SetMargins(obj uintptr, value uintptr) {
   image_SetMargins.Call(obj, value)
}

func Image_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := image_GetComponentCount.Call(obj)
    return int32(ret)
}

func Image_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := image_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Image_SetComponentIndex(obj uintptr, value int32) {
   image_SetComponentIndex.Call(obj, uintptr(value))
}

func Image_GetOwner(obj uintptr) uintptr {
    ret, _, _ := image_GetOwner.Call(obj)
    return ret
}

func Image_GetName(obj uintptr) string {
    ret, _, _ := image_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Image_SetName(obj uintptr, value string) {
   image_SetName.Call(obj, GoStrToDStr(value))
}

func Image_GetTag(obj uintptr) int {
    ret, _, _ := image_GetTag.Call(obj)
    return int(ret)
}

func Image_SetTag(obj uintptr, value int) {
   image_SetTag.Call(obj, uintptr(value))
}

func Image_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := image_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TLinkLabel ---------------------------

func LinkLabel_Create(obj uintptr) uintptr {
    ret, _, _ := linkLabel_Create.Call(obj)
    return ret
}

func LinkLabel_Free(obj uintptr) {
    linkLabel_Free.Call(obj)
}

func LinkLabel_CanFocus(obj uintptr) bool {
    ret, _, _ := linkLabel_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_FlipChildren(obj uintptr, AllLevels bool)  {
    linkLabel_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func LinkLabel_Focused(obj uintptr) bool {
    ret, _, _ := linkLabel_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_HandleAllocated(obj uintptr) bool {
    ret, _, _ := linkLabel_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_Invalidate(obj uintptr)  {
    linkLabel_Invalidate.Call(obj)
}

func LinkLabel_Realign(obj uintptr)  {
    linkLabel_Realign.Call(obj)
}

func LinkLabel_Repaint(obj uintptr)  {
    linkLabel_Repaint.Call(obj)
}

func LinkLabel_ScaleBy(obj uintptr, M int32, D int32)  {
    linkLabel_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func LinkLabel_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    linkLabel_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func LinkLabel_SetFocus(obj uintptr)  {
    linkLabel_SetFocus.Call(obj)
}

func LinkLabel_Update(obj uintptr)  {
    linkLabel_Update.Call(obj)
}

func LinkLabel_BringToFront(obj uintptr)  {
    linkLabel_BringToFront.Call(obj)
}

func LinkLabel_HasParent(obj uintptr) bool {
    ret, _, _ := linkLabel_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_Hide(obj uintptr)  {
    linkLabel_Hide.Call(obj)
}

func LinkLabel_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := linkLabel_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func LinkLabel_Refresh(obj uintptr)  {
    linkLabel_Refresh.Call(obj)
}

func LinkLabel_SendToBack(obj uintptr)  {
    linkLabel_SendToBack.Call(obj)
}

func LinkLabel_Show(obj uintptr)  {
    linkLabel_Show.Call(obj)
}

func LinkLabel_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := linkLabel_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func LinkLabel_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := linkLabel_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func LinkLabel_GetNamePath(obj uintptr) string {
    ret, _, _ := linkLabel_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func LinkLabel_Assign(obj uintptr, Source uintptr)  {
    linkLabel_Assign.Call(obj, Source )
}

func LinkLabel_ClassName(obj uintptr) string {
    ret, _, _ := linkLabel_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func LinkLabel_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := linkLabel_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func LinkLabel_GetHashCode(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetHashCode.Call(obj)
    return int32(ret)
}

func LinkLabel_ToString(obj uintptr) string {
    ret, _, _ := linkLabel_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func LinkLabel_GetAlign(obj uintptr) TAlign {
    ret, _, _ := linkLabel_GetAlign.Call(obj)
    return TAlign(ret)
}

func LinkLabel_SetAlign(obj uintptr, value TAlign) {
   linkLabel_SetAlign.Call(obj, uintptr(value))
}

func LinkLabel_GetAlignment(obj uintptr) TLinkAlignment {
    ret, _, _ := linkLabel_GetAlignment.Call(obj)
    return TLinkAlignment(ret)
}

func LinkLabel_SetAlignment(obj uintptr, value TLinkAlignment) {
   linkLabel_SetAlignment.Call(obj, uintptr(value))
}

func LinkLabel_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := linkLabel_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func LinkLabel_SetAnchors(obj uintptr, value TAnchors) {
   linkLabel_SetAnchors.Call(obj, uintptr(value))
}

func LinkLabel_GetAutoSize(obj uintptr) bool {
    ret, _, _ := linkLabel_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetAutoSize(obj uintptr, value bool) {
   linkLabel_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetBevelEdges(obj uintptr) TBevelEdges {
    ret, _, _ := linkLabel_GetBevelEdges.Call(obj)
    return TBevelEdges(ret)
}

func LinkLabel_SetBevelEdges(obj uintptr, value TBevelEdges) {
   linkLabel_SetBevelEdges.Call(obj, uintptr(value))
}

func LinkLabel_GetBevelInner(obj uintptr) TBevelCut {
    ret, _, _ := linkLabel_GetBevelInner.Call(obj)
    return TBevelCut(ret)
}

func LinkLabel_SetBevelInner(obj uintptr, value TBevelCut) {
   linkLabel_SetBevelInner.Call(obj, uintptr(value))
}

func LinkLabel_GetBevelKind(obj uintptr) TBevelKind {
    ret, _, _ := linkLabel_GetBevelKind.Call(obj)
    return TBevelKind(ret)
}

func LinkLabel_SetBevelKind(obj uintptr, value TBevelKind) {
   linkLabel_SetBevelKind.Call(obj, uintptr(value))
}

func LinkLabel_GetBevelOuter(obj uintptr) TBevelCut {
    ret, _, _ := linkLabel_GetBevelOuter.Call(obj)
    return TBevelCut(ret)
}

func LinkLabel_SetBevelOuter(obj uintptr, value TBevelCut) {
   linkLabel_SetBevelOuter.Call(obj, uintptr(value))
}

func LinkLabel_GetCaption(obj uintptr) string {
    ret, _, _ := linkLabel_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func LinkLabel_SetCaption(obj uintptr, value string) {
   linkLabel_SetCaption.Call(obj, GoStrToDStr(value))
}

func LinkLabel_GetColor(obj uintptr) TColor {
    ret, _, _ := linkLabel_GetColor.Call(obj)
    return TColor(ret)
}

func LinkLabel_SetColor(obj uintptr, value TColor) {
   linkLabel_SetColor.Call(obj, uintptr(value))
}

func LinkLabel_GetEnabled(obj uintptr) bool {
    ret, _, _ := linkLabel_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetEnabled(obj uintptr, value bool) {
   linkLabel_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetFont(obj uintptr) uintptr {
    ret, _, _ := linkLabel_GetFont.Call(obj)
    return ret
}

func LinkLabel_SetFont(obj uintptr, value uintptr) {
   linkLabel_SetFont.Call(obj, value)
}

func LinkLabel_GetParentColor(obj uintptr) bool {
    ret, _, _ := linkLabel_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetParentColor(obj uintptr, value bool) {
   linkLabel_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetParentFont(obj uintptr) bool {
    ret, _, _ := linkLabel_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetParentFont(obj uintptr, value bool) {
   linkLabel_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := linkLabel_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetParentShowHint(obj uintptr, value bool) {
   linkLabel_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := linkLabel_GetPopupMenu.Call(obj)
    return ret
}

func LinkLabel_SetPopupMenu(obj uintptr, value uintptr) {
   linkLabel_SetPopupMenu.Call(obj, value)
}

func LinkLabel_GetShowHint(obj uintptr) bool {
    ret, _, _ := linkLabel_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetShowHint(obj uintptr, value bool) {
   linkLabel_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := linkLabel_GetTabOrder.Call(obj)
    return uint16(ret)
}

func LinkLabel_SetTabOrder(obj uintptr, value uint16) {
   linkLabel_SetTabOrder.Call(obj, uintptr(value))
}

func LinkLabel_GetTabStop(obj uintptr) bool {
    ret, _, _ := linkLabel_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetTabStop(obj uintptr, value bool) {
   linkLabel_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetUseVisualStyle(obj uintptr) bool {
    ret, _, _ := linkLabel_GetUseVisualStyle.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetUseVisualStyle(obj uintptr, value bool) {
   linkLabel_SetUseVisualStyle.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetVisible(obj uintptr) bool {
    ret, _, _ := linkLabel_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetVisible(obj uintptr, value bool) {
   linkLabel_SetVisible.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_SetOnClick(obj uintptr, fn interface{}) {
    linkLabel_SetOnClick.Call(obj, addEventToMap(fn))
}

func LinkLabel_SetOnDblClick(obj uintptr, fn interface{}) {
    linkLabel_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func LinkLabel_SetOnMouseDown(obj uintptr, fn interface{}) {
    linkLabel_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func LinkLabel_SetOnMouseEnter(obj uintptr, fn interface{}) {
    linkLabel_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func LinkLabel_SetOnMouseLeave(obj uintptr, fn interface{}) {
    linkLabel_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func LinkLabel_SetOnMouseMove(obj uintptr, fn interface{}) {
    linkLabel_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func LinkLabel_SetOnMouseUp(obj uintptr, fn interface{}) {
    linkLabel_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func LinkLabel_SetOnLinkClick(obj uintptr, fn interface{}) {
    linkLabel_SetOnLinkClick.Call(obj, addEventToMap(fn))
}

func LinkLabel_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := linkLabel_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetDoubleBuffered(obj uintptr, value bool) {
   linkLabel_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetBrush(obj uintptr) uintptr {
    ret, _, _ := linkLabel_GetBrush.Call(obj)
    return ret
}

func LinkLabel_GetControlCount(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetControlCount.Call(obj)
    return int32(ret)
}

func LinkLabel_GetHandle(obj uintptr) HWND {
    ret, _, _ := linkLabel_GetHandle.Call(obj)
    return HWND(ret)
}

func LinkLabel_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := linkLabel_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetParentDoubleBuffered(obj uintptr, value bool) {
   linkLabel_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := linkLabel_GetParentWindow.Call(obj)
    return HWND(ret)
}

func LinkLabel_SetParentWindow(obj uintptr, value HWND) {
   linkLabel_SetParentWindow.Call(obj, uintptr(value))
}

func LinkLabel_GetAction(obj uintptr) uintptr {
    ret, _, _ := linkLabel_GetAction.Call(obj)
    return ret
}

func LinkLabel_SetAction(obj uintptr, value uintptr) {
   linkLabel_SetAction.Call(obj, value)
}

func LinkLabel_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := linkLabel_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func LinkLabel_SetBiDiMode(obj uintptr, value TBiDiMode) {
   linkLabel_SetBiDiMode.Call(obj, uintptr(value))
}

func LinkLabel_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    linkLabel_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func LinkLabel_SetBoundsRect(obj uintptr, value TRect) {
   linkLabel_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func LinkLabel_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetClientHeight.Call(obj)
    return int32(ret)
}

func LinkLabel_SetClientHeight(obj uintptr, value int32) {
   linkLabel_SetClientHeight.Call(obj, uintptr(value))
}

func LinkLabel_GetClientRect(obj uintptr) TRect {
    var ret TRect
    linkLabel_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func LinkLabel_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetClientWidth.Call(obj)
    return int32(ret)
}

func LinkLabel_SetClientWidth(obj uintptr, value int32) {
   linkLabel_SetClientWidth.Call(obj, uintptr(value))
}

func LinkLabel_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func LinkLabel_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetExplicitTop.Call(obj)
    return int32(ret)
}

func LinkLabel_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func LinkLabel_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func LinkLabel_GetParent(obj uintptr) uintptr {
    ret, _, _ := linkLabel_GetParent.Call(obj)
    return ret
}

func LinkLabel_SetParent(obj uintptr, value uintptr) {
   linkLabel_SetParent.Call(obj, value)
}

func LinkLabel_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := linkLabel_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func LinkLabel_SetStyleElements(obj uintptr, value TStyleElements) {
   linkLabel_SetStyleElements.Call(obj, uintptr(value))
}

func LinkLabel_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := linkLabel_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func LinkLabel_SetAlignWithMargins(obj uintptr, value bool) {
   linkLabel_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func LinkLabel_GetLeft(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetLeft.Call(obj)
    return int32(ret)
}

func LinkLabel_SetLeft(obj uintptr, value int32) {
   linkLabel_SetLeft.Call(obj, uintptr(value))
}

func LinkLabel_GetTop(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetTop.Call(obj)
    return int32(ret)
}

func LinkLabel_SetTop(obj uintptr, value int32) {
   linkLabel_SetTop.Call(obj, uintptr(value))
}

func LinkLabel_GetWidth(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetWidth.Call(obj)
    return int32(ret)
}

func LinkLabel_SetWidth(obj uintptr, value int32) {
   linkLabel_SetWidth.Call(obj, uintptr(value))
}

func LinkLabel_GetHeight(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetHeight.Call(obj)
    return int32(ret)
}

func LinkLabel_SetHeight(obj uintptr, value int32) {
   linkLabel_SetHeight.Call(obj, uintptr(value))
}

func LinkLabel_GetCursor(obj uintptr) TCursor {
    ret, _, _ := linkLabel_GetCursor.Call(obj)
    return TCursor(ret)
}

func LinkLabel_SetCursor(obj uintptr, value TCursor) {
   linkLabel_SetCursor.Call(obj, uintptr(value))
}

func LinkLabel_GetHint(obj uintptr) string {
    ret, _, _ := linkLabel_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func LinkLabel_SetHint(obj uintptr, value string) {
   linkLabel_SetHint.Call(obj, GoStrToDStr(value))
}

func LinkLabel_GetMargins(obj uintptr) uintptr {
    ret, _, _ := linkLabel_GetMargins.Call(obj)
    return ret
}

func LinkLabel_SetMargins(obj uintptr, value uintptr) {
   linkLabel_SetMargins.Call(obj, value)
}

func LinkLabel_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetComponentCount.Call(obj)
    return int32(ret)
}

func LinkLabel_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := linkLabel_GetComponentIndex.Call(obj)
    return int32(ret)
}

func LinkLabel_SetComponentIndex(obj uintptr, value int32) {
   linkLabel_SetComponentIndex.Call(obj, uintptr(value))
}

func LinkLabel_GetOwner(obj uintptr) uintptr {
    ret, _, _ := linkLabel_GetOwner.Call(obj)
    return ret
}

func LinkLabel_GetName(obj uintptr) string {
    ret, _, _ := linkLabel_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func LinkLabel_SetName(obj uintptr, value string) {
   linkLabel_SetName.Call(obj, GoStrToDStr(value))
}

func LinkLabel_GetTag(obj uintptr) int {
    ret, _, _ := linkLabel_GetTag.Call(obj)
    return int(ret)
}

func LinkLabel_SetTag(obj uintptr, value int) {
   linkLabel_SetTag.Call(obj, uintptr(value))
}

func LinkLabel_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := linkLabel_GetControls.Call(obj, uintptr(Index))
    return ret
}

func LinkLabel_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := linkLabel_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TSpeedButton ---------------------------

func SpeedButton_Create(obj uintptr) uintptr {
    ret, _, _ := speedButton_Create.Call(obj)
    return ret
}

func SpeedButton_Free(obj uintptr) {
    speedButton_Free.Call(obj)
}

func SpeedButton_Click(obj uintptr)  {
    speedButton_Click.Call(obj)
}

func SpeedButton_BringToFront(obj uintptr)  {
    speedButton_BringToFront.Call(obj)
}

func SpeedButton_HasParent(obj uintptr) bool {
    ret, _, _ := speedButton_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_Hide(obj uintptr)  {
    speedButton_Hide.Call(obj)
}

func SpeedButton_Invalidate(obj uintptr)  {
    speedButton_Invalidate.Call(obj)
}

func SpeedButton_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := speedButton_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func SpeedButton_Refresh(obj uintptr)  {
    speedButton_Refresh.Call(obj)
}

func SpeedButton_Repaint(obj uintptr)  {
    speedButton_Repaint.Call(obj)
}

func SpeedButton_SendToBack(obj uintptr)  {
    speedButton_SendToBack.Call(obj)
}

func SpeedButton_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    speedButton_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func SpeedButton_Show(obj uintptr)  {
    speedButton_Show.Call(obj)
}

func SpeedButton_Update(obj uintptr)  {
    speedButton_Update.Call(obj)
}

func SpeedButton_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := speedButton_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func SpeedButton_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := speedButton_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func SpeedButton_GetNamePath(obj uintptr) string {
    ret, _, _ := speedButton_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func SpeedButton_Assign(obj uintptr, Source uintptr)  {
    speedButton_Assign.Call(obj, Source )
}

func SpeedButton_ClassName(obj uintptr) string {
    ret, _, _ := speedButton_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func SpeedButton_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := speedButton_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func SpeedButton_GetHashCode(obj uintptr) int32 {
    ret, _, _ := speedButton_GetHashCode.Call(obj)
    return int32(ret)
}

func SpeedButton_ToString(obj uintptr) string {
    ret, _, _ := speedButton_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func SpeedButton_GetAction(obj uintptr) uintptr {
    ret, _, _ := speedButton_GetAction.Call(obj)
    return ret
}

func SpeedButton_SetAction(obj uintptr, value uintptr) {
   speedButton_SetAction.Call(obj, value)
}

func SpeedButton_GetAlign(obj uintptr) TAlign {
    ret, _, _ := speedButton_GetAlign.Call(obj)
    return TAlign(ret)
}

func SpeedButton_SetAlign(obj uintptr, value TAlign) {
   speedButton_SetAlign.Call(obj, uintptr(value))
}

func SpeedButton_GetAllowAllUp(obj uintptr) bool {
    ret, _, _ := speedButton_GetAllowAllUp.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetAllowAllUp(obj uintptr, value bool) {
   speedButton_SetAllowAllUp.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := speedButton_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func SpeedButton_SetAnchors(obj uintptr, value TAnchors) {
   speedButton_SetAnchors.Call(obj, uintptr(value))
}

func SpeedButton_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := speedButton_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func SpeedButton_SetBiDiMode(obj uintptr, value TBiDiMode) {
   speedButton_SetBiDiMode.Call(obj, uintptr(value))
}

func SpeedButton_GetGroupIndex(obj uintptr) int32 {
    ret, _, _ := speedButton_GetGroupIndex.Call(obj)
    return int32(ret)
}

func SpeedButton_SetGroupIndex(obj uintptr, value int32) {
   speedButton_SetGroupIndex.Call(obj, uintptr(value))
}

func SpeedButton_GetDown(obj uintptr) bool {
    ret, _, _ := speedButton_GetDown.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetDown(obj uintptr, value bool) {
   speedButton_SetDown.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetCaption(obj uintptr) string {
    ret, _, _ := speedButton_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func SpeedButton_SetCaption(obj uintptr, value string) {
   speedButton_SetCaption.Call(obj, GoStrToDStr(value))
}

func SpeedButton_GetEnabled(obj uintptr) bool {
    ret, _, _ := speedButton_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetEnabled(obj uintptr, value bool) {
   speedButton_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetFlat(obj uintptr) bool {
    ret, _, _ := speedButton_GetFlat.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetFlat(obj uintptr, value bool) {
   speedButton_SetFlat.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetFont(obj uintptr) uintptr {
    ret, _, _ := speedButton_GetFont.Call(obj)
    return ret
}

func SpeedButton_SetFont(obj uintptr, value uintptr) {
   speedButton_SetFont.Call(obj, value)
}

func SpeedButton_GetLayout(obj uintptr) TButtonLayout {
    ret, _, _ := speedButton_GetLayout.Call(obj)
    return TButtonLayout(ret)
}

func SpeedButton_SetLayout(obj uintptr, value TButtonLayout) {
   speedButton_SetLayout.Call(obj, uintptr(value))
}

func SpeedButton_GetParentFont(obj uintptr) bool {
    ret, _, _ := speedButton_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetParentFont(obj uintptr, value bool) {
   speedButton_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := speedButton_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetParentShowHint(obj uintptr, value bool) {
   speedButton_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := speedButton_GetPopupMenu.Call(obj)
    return ret
}

func SpeedButton_SetPopupMenu(obj uintptr, value uintptr) {
   speedButton_SetPopupMenu.Call(obj, value)
}

func SpeedButton_GetShowHint(obj uintptr) bool {
    ret, _, _ := speedButton_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetShowHint(obj uintptr, value bool) {
   speedButton_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetTransparent(obj uintptr) bool {
    ret, _, _ := speedButton_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetTransparent(obj uintptr, value bool) {
   speedButton_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetVisible(obj uintptr) bool {
    ret, _, _ := speedButton_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetVisible(obj uintptr, value bool) {
   speedButton_SetVisible.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := speedButton_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func SpeedButton_SetStyleElements(obj uintptr, value TStyleElements) {
   speedButton_SetStyleElements.Call(obj, uintptr(value))
}

func SpeedButton_SetOnClick(obj uintptr, fn interface{}) {
    speedButton_SetOnClick.Call(obj, addEventToMap(fn))
}

func SpeedButton_SetOnDblClick(obj uintptr, fn interface{}) {
    speedButton_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func SpeedButton_SetOnMouseDown(obj uintptr, fn interface{}) {
    speedButton_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func SpeedButton_SetOnMouseEnter(obj uintptr, fn interface{}) {
    speedButton_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func SpeedButton_SetOnMouseLeave(obj uintptr, fn interface{}) {
    speedButton_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func SpeedButton_SetOnMouseMove(obj uintptr, fn interface{}) {
    speedButton_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func SpeedButton_SetOnMouseUp(obj uintptr, fn interface{}) {
    speedButton_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func SpeedButton_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    speedButton_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func SpeedButton_SetBoundsRect(obj uintptr, value TRect) {
   speedButton_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func SpeedButton_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := speedButton_GetClientHeight.Call(obj)
    return int32(ret)
}

func SpeedButton_SetClientHeight(obj uintptr, value int32) {
   speedButton_SetClientHeight.Call(obj, uintptr(value))
}

func SpeedButton_GetClientRect(obj uintptr) TRect {
    var ret TRect
    speedButton_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func SpeedButton_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := speedButton_GetClientWidth.Call(obj)
    return int32(ret)
}

func SpeedButton_SetClientWidth(obj uintptr, value int32) {
   speedButton_SetClientWidth.Call(obj, uintptr(value))
}

func SpeedButton_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := speedButton_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func SpeedButton_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := speedButton_GetExplicitTop.Call(obj)
    return int32(ret)
}

func SpeedButton_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := speedButton_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func SpeedButton_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := speedButton_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func SpeedButton_GetParent(obj uintptr) uintptr {
    ret, _, _ := speedButton_GetParent.Call(obj)
    return ret
}

func SpeedButton_SetParent(obj uintptr, value uintptr) {
   speedButton_SetParent.Call(obj, value)
}

func SpeedButton_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := speedButton_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func SpeedButton_SetAlignWithMargins(obj uintptr, value bool) {
   speedButton_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func SpeedButton_GetLeft(obj uintptr) int32 {
    ret, _, _ := speedButton_GetLeft.Call(obj)
    return int32(ret)
}

func SpeedButton_SetLeft(obj uintptr, value int32) {
   speedButton_SetLeft.Call(obj, uintptr(value))
}

func SpeedButton_GetTop(obj uintptr) int32 {
    ret, _, _ := speedButton_GetTop.Call(obj)
    return int32(ret)
}

func SpeedButton_SetTop(obj uintptr, value int32) {
   speedButton_SetTop.Call(obj, uintptr(value))
}

func SpeedButton_GetWidth(obj uintptr) int32 {
    ret, _, _ := speedButton_GetWidth.Call(obj)
    return int32(ret)
}

func SpeedButton_SetWidth(obj uintptr, value int32) {
   speedButton_SetWidth.Call(obj, uintptr(value))
}

func SpeedButton_GetHeight(obj uintptr) int32 {
    ret, _, _ := speedButton_GetHeight.Call(obj)
    return int32(ret)
}

func SpeedButton_SetHeight(obj uintptr, value int32) {
   speedButton_SetHeight.Call(obj, uintptr(value))
}

func SpeedButton_GetCursor(obj uintptr) TCursor {
    ret, _, _ := speedButton_GetCursor.Call(obj)
    return TCursor(ret)
}

func SpeedButton_SetCursor(obj uintptr, value TCursor) {
   speedButton_SetCursor.Call(obj, uintptr(value))
}

func SpeedButton_GetHint(obj uintptr) string {
    ret, _, _ := speedButton_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func SpeedButton_SetHint(obj uintptr, value string) {
   speedButton_SetHint.Call(obj, GoStrToDStr(value))
}

func SpeedButton_GetMargins(obj uintptr) uintptr {
    ret, _, _ := speedButton_GetMargins.Call(obj)
    return ret
}

func SpeedButton_SetMargins(obj uintptr, value uintptr) {
   speedButton_SetMargins.Call(obj, value)
}

func SpeedButton_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := speedButton_GetComponentCount.Call(obj)
    return int32(ret)
}

func SpeedButton_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := speedButton_GetComponentIndex.Call(obj)
    return int32(ret)
}

func SpeedButton_SetComponentIndex(obj uintptr, value int32) {
   speedButton_SetComponentIndex.Call(obj, uintptr(value))
}

func SpeedButton_GetOwner(obj uintptr) uintptr {
    ret, _, _ := speedButton_GetOwner.Call(obj)
    return ret
}

func SpeedButton_GetName(obj uintptr) string {
    ret, _, _ := speedButton_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func SpeedButton_SetName(obj uintptr, value string) {
   speedButton_SetName.Call(obj, GoStrToDStr(value))
}

func SpeedButton_GetTag(obj uintptr) int {
    ret, _, _ := speedButton_GetTag.Call(obj)
    return int(ret)
}

func SpeedButton_SetTag(obj uintptr, value int) {
   speedButton_SetTag.Call(obj, uintptr(value))
}

func SpeedButton_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := speedButton_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TSplitter ---------------------------

func Splitter_Create(obj uintptr) uintptr {
    ret, _, _ := splitter_Create.Call(obj)
    return ret
}

func Splitter_Free(obj uintptr) {
    splitter_Free.Call(obj)
}

func Splitter_BringToFront(obj uintptr)  {
    splitter_BringToFront.Call(obj)
}

func Splitter_HasParent(obj uintptr) bool {
    ret, _, _ := splitter_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Splitter_Hide(obj uintptr)  {
    splitter_Hide.Call(obj)
}

func Splitter_Invalidate(obj uintptr)  {
    splitter_Invalidate.Call(obj)
}

func Splitter_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := splitter_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func Splitter_Refresh(obj uintptr)  {
    splitter_Refresh.Call(obj)
}

func Splitter_Repaint(obj uintptr)  {
    splitter_Repaint.Call(obj)
}

func Splitter_SendToBack(obj uintptr)  {
    splitter_SendToBack.Call(obj)
}

func Splitter_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    splitter_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func Splitter_Show(obj uintptr)  {
    splitter_Show.Call(obj)
}

func Splitter_Update(obj uintptr)  {
    splitter_Update.Call(obj)
}

func Splitter_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := splitter_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Splitter_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := splitter_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Splitter_GetNamePath(obj uintptr) string {
    ret, _, _ := splitter_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Splitter_Assign(obj uintptr, Source uintptr)  {
    splitter_Assign.Call(obj, Source )
}

func Splitter_ClassName(obj uintptr) string {
    ret, _, _ := splitter_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Splitter_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := splitter_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Splitter_GetHashCode(obj uintptr) int32 {
    ret, _, _ := splitter_GetHashCode.Call(obj)
    return int32(ret)
}

func Splitter_ToString(obj uintptr) string {
    ret, _, _ := splitter_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Splitter_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := splitter_GetCanvas.Call(obj)
    return ret
}

func Splitter_GetAlign(obj uintptr) TAlign {
    ret, _, _ := splitter_GetAlign.Call(obj)
    return TAlign(ret)
}

func Splitter_SetAlign(obj uintptr, value TAlign) {
   splitter_SetAlign.Call(obj, uintptr(value))
}

func Splitter_GetColor(obj uintptr) TColor {
    ret, _, _ := splitter_GetColor.Call(obj)
    return TColor(ret)
}

func Splitter_SetColor(obj uintptr, value TColor) {
   splitter_SetColor.Call(obj, uintptr(value))
}

func Splitter_GetCursor(obj uintptr) TCursor {
    ret, _, _ := splitter_GetCursor.Call(obj)
    return TCursor(ret)
}

func Splitter_SetCursor(obj uintptr, value TCursor) {
   splitter_SetCursor.Call(obj, uintptr(value))
}

func Splitter_GetParentColor(obj uintptr) bool {
    ret, _, _ := splitter_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func Splitter_SetParentColor(obj uintptr, value bool) {
   splitter_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func Splitter_GetVisible(obj uintptr) bool {
    ret, _, _ := splitter_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Splitter_SetVisible(obj uintptr, value bool) {
   splitter_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Splitter_GetWidth(obj uintptr) int32 {
    ret, _, _ := splitter_GetWidth.Call(obj)
    return int32(ret)
}

func Splitter_SetWidth(obj uintptr, value int32) {
   splitter_SetWidth.Call(obj, uintptr(value))
}

func Splitter_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := splitter_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func Splitter_SetStyleElements(obj uintptr, value TStyleElements) {
   splitter_SetStyleElements.Call(obj, uintptr(value))
}

func Splitter_SetOnPaint(obj uintptr, fn interface{}) {
    splitter_SetOnPaint.Call(obj, addEventToMap(fn))
}

func Splitter_GetEnabled(obj uintptr) bool {
    ret, _, _ := splitter_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Splitter_SetEnabled(obj uintptr, value bool) {
   splitter_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Splitter_GetAction(obj uintptr) uintptr {
    ret, _, _ := splitter_GetAction.Call(obj)
    return ret
}

func Splitter_SetAction(obj uintptr, value uintptr) {
   splitter_SetAction.Call(obj, value)
}

func Splitter_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := splitter_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func Splitter_SetAnchors(obj uintptr, value TAnchors) {
   splitter_SetAnchors.Call(obj, uintptr(value))
}

func Splitter_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := splitter_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func Splitter_SetBiDiMode(obj uintptr, value TBiDiMode) {
   splitter_SetBiDiMode.Call(obj, uintptr(value))
}

func Splitter_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    splitter_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Splitter_SetBoundsRect(obj uintptr, value TRect) {
   splitter_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Splitter_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := splitter_GetClientHeight.Call(obj)
    return int32(ret)
}

func Splitter_SetClientHeight(obj uintptr, value int32) {
   splitter_SetClientHeight.Call(obj, uintptr(value))
}

func Splitter_GetClientRect(obj uintptr) TRect {
    var ret TRect
    splitter_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Splitter_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := splitter_GetClientWidth.Call(obj)
    return int32(ret)
}

func Splitter_SetClientWidth(obj uintptr, value int32) {
   splitter_SetClientWidth.Call(obj, uintptr(value))
}

func Splitter_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := splitter_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Splitter_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := splitter_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Splitter_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := splitter_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Splitter_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := splitter_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Splitter_GetShowHint(obj uintptr) bool {
    ret, _, _ := splitter_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Splitter_SetShowHint(obj uintptr, value bool) {
   splitter_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Splitter_GetParent(obj uintptr) uintptr {
    ret, _, _ := splitter_GetParent.Call(obj)
    return ret
}

func Splitter_SetParent(obj uintptr, value uintptr) {
   splitter_SetParent.Call(obj, value)
}

func Splitter_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := splitter_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func Splitter_SetAlignWithMargins(obj uintptr, value bool) {
   splitter_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func Splitter_GetLeft(obj uintptr) int32 {
    ret, _, _ := splitter_GetLeft.Call(obj)
    return int32(ret)
}

func Splitter_SetLeft(obj uintptr, value int32) {
   splitter_SetLeft.Call(obj, uintptr(value))
}

func Splitter_GetTop(obj uintptr) int32 {
    ret, _, _ := splitter_GetTop.Call(obj)
    return int32(ret)
}

func Splitter_SetTop(obj uintptr, value int32) {
   splitter_SetTop.Call(obj, uintptr(value))
}

func Splitter_GetHeight(obj uintptr) int32 {
    ret, _, _ := splitter_GetHeight.Call(obj)
    return int32(ret)
}

func Splitter_SetHeight(obj uintptr, value int32) {
   splitter_SetHeight.Call(obj, uintptr(value))
}

func Splitter_GetHint(obj uintptr) string {
    ret, _, _ := splitter_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Splitter_SetHint(obj uintptr, value string) {
   splitter_SetHint.Call(obj, GoStrToDStr(value))
}

func Splitter_GetMargins(obj uintptr) uintptr {
    ret, _, _ := splitter_GetMargins.Call(obj)
    return ret
}

func Splitter_SetMargins(obj uintptr, value uintptr) {
   splitter_SetMargins.Call(obj, value)
}

func Splitter_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := splitter_GetComponentCount.Call(obj)
    return int32(ret)
}

func Splitter_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := splitter_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Splitter_SetComponentIndex(obj uintptr, value int32) {
   splitter_SetComponentIndex.Call(obj, uintptr(value))
}

func Splitter_GetOwner(obj uintptr) uintptr {
    ret, _, _ := splitter_GetOwner.Call(obj)
    return ret
}

func Splitter_GetName(obj uintptr) string {
    ret, _, _ := splitter_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Splitter_SetName(obj uintptr, value string) {
   splitter_SetName.Call(obj, GoStrToDStr(value))
}

func Splitter_GetTag(obj uintptr) int {
    ret, _, _ := splitter_GetTag.Call(obj)
    return int(ret)
}

func Splitter_SetTag(obj uintptr, value int) {
   splitter_SetTag.Call(obj, uintptr(value))
}

func Splitter_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := splitter_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TRadioGroup ---------------------------

func RadioGroup_Create(obj uintptr) uintptr {
    ret, _, _ := radioGroup_Create.Call(obj)
    return ret
}

func RadioGroup_Free(obj uintptr) {
    radioGroup_Free.Call(obj)
}

func RadioGroup_FlipChildren(obj uintptr, AllLevels bool)  {
    radioGroup_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func RadioGroup_CanFocus(obj uintptr) bool {
    ret, _, _ := radioGroup_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_Focused(obj uintptr) bool {
    ret, _, _ := radioGroup_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_HandleAllocated(obj uintptr) bool {
    ret, _, _ := radioGroup_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_Invalidate(obj uintptr)  {
    radioGroup_Invalidate.Call(obj)
}

func RadioGroup_Realign(obj uintptr)  {
    radioGroup_Realign.Call(obj)
}

func RadioGroup_Repaint(obj uintptr)  {
    radioGroup_Repaint.Call(obj)
}

func RadioGroup_ScaleBy(obj uintptr, M int32, D int32)  {
    radioGroup_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func RadioGroup_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    radioGroup_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func RadioGroup_SetFocus(obj uintptr)  {
    radioGroup_SetFocus.Call(obj)
}

func RadioGroup_Update(obj uintptr)  {
    radioGroup_Update.Call(obj)
}

func RadioGroup_BringToFront(obj uintptr)  {
    radioGroup_BringToFront.Call(obj)
}

func RadioGroup_HasParent(obj uintptr) bool {
    ret, _, _ := radioGroup_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_Hide(obj uintptr)  {
    radioGroup_Hide.Call(obj)
}

func RadioGroup_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := radioGroup_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func RadioGroup_Refresh(obj uintptr)  {
    radioGroup_Refresh.Call(obj)
}

func RadioGroup_SendToBack(obj uintptr)  {
    radioGroup_SendToBack.Call(obj)
}

func RadioGroup_Show(obj uintptr)  {
    radioGroup_Show.Call(obj)
}

func RadioGroup_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := radioGroup_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func RadioGroup_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := radioGroup_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func RadioGroup_GetNamePath(obj uintptr) string {
    ret, _, _ := radioGroup_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func RadioGroup_Assign(obj uintptr, Source uintptr)  {
    radioGroup_Assign.Call(obj, Source )
}

func RadioGroup_ClassName(obj uintptr) string {
    ret, _, _ := radioGroup_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func RadioGroup_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := radioGroup_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func RadioGroup_GetHashCode(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetHashCode.Call(obj)
    return int32(ret)
}

func RadioGroup_ToString(obj uintptr) string {
    ret, _, _ := radioGroup_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func RadioGroup_GetAlign(obj uintptr) TAlign {
    ret, _, _ := radioGroup_GetAlign.Call(obj)
    return TAlign(ret)
}

func RadioGroup_SetAlign(obj uintptr, value TAlign) {
   radioGroup_SetAlign.Call(obj, uintptr(value))
}

func RadioGroup_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := radioGroup_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func RadioGroup_SetAnchors(obj uintptr, value TAnchors) {
   radioGroup_SetAnchors.Call(obj, uintptr(value))
}

func RadioGroup_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := radioGroup_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func RadioGroup_SetBiDiMode(obj uintptr, value TBiDiMode) {
   radioGroup_SetBiDiMode.Call(obj, uintptr(value))
}

func RadioGroup_GetCaption(obj uintptr) string {
    ret, _, _ := radioGroup_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func RadioGroup_SetCaption(obj uintptr, value string) {
   radioGroup_SetCaption.Call(obj, GoStrToDStr(value))
}

func RadioGroup_GetColor(obj uintptr) TColor {
    ret, _, _ := radioGroup_GetColor.Call(obj)
    return TColor(ret)
}

func RadioGroup_SetColor(obj uintptr, value TColor) {
   radioGroup_SetColor.Call(obj, uintptr(value))
}

func RadioGroup_GetColumns(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetColumns.Call(obj)
    return int32(ret)
}

func RadioGroup_SetColumns(obj uintptr, value int32) {
   radioGroup_SetColumns.Call(obj, uintptr(value))
}

func RadioGroup_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := radioGroup_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetDoubleBuffered(obj uintptr, value bool) {
   radioGroup_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetEnabled(obj uintptr) bool {
    ret, _, _ := radioGroup_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetEnabled(obj uintptr, value bool) {
   radioGroup_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetFont(obj uintptr) uintptr {
    ret, _, _ := radioGroup_GetFont.Call(obj)
    return ret
}

func RadioGroup_SetFont(obj uintptr, value uintptr) {
   radioGroup_SetFont.Call(obj, value)
}

func RadioGroup_GetItemIndex(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetItemIndex.Call(obj)
    return int32(ret)
}

func RadioGroup_SetItemIndex(obj uintptr, value int32) {
   radioGroup_SetItemIndex.Call(obj, uintptr(value))
}

func RadioGroup_GetItems(obj uintptr) uintptr {
    ret, _, _ := radioGroup_GetItems.Call(obj)
    return ret
}

func RadioGroup_SetItems(obj uintptr, value uintptr) {
   radioGroup_SetItems.Call(obj, value)
}

func RadioGroup_GetParentBackground(obj uintptr) bool {
    ret, _, _ := radioGroup_GetParentBackground.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetParentBackground(obj uintptr, value bool) {
   radioGroup_SetParentBackground.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetParentColor(obj uintptr) bool {
    ret, _, _ := radioGroup_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetParentColor(obj uintptr, value bool) {
   radioGroup_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := radioGroup_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetParentCtl3D(obj uintptr, value bool) {
   radioGroup_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := radioGroup_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetParentDoubleBuffered(obj uintptr, value bool) {
   radioGroup_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetParentFont(obj uintptr) bool {
    ret, _, _ := radioGroup_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetParentFont(obj uintptr, value bool) {
   radioGroup_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := radioGroup_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetParentShowHint(obj uintptr, value bool) {
   radioGroup_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := radioGroup_GetPopupMenu.Call(obj)
    return ret
}

func RadioGroup_SetPopupMenu(obj uintptr, value uintptr) {
   radioGroup_SetPopupMenu.Call(obj, value)
}

func RadioGroup_GetShowHint(obj uintptr) bool {
    ret, _, _ := radioGroup_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetShowHint(obj uintptr, value bool) {
   radioGroup_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := radioGroup_GetTabOrder.Call(obj)
    return uint16(ret)
}

func RadioGroup_SetTabOrder(obj uintptr, value uint16) {
   radioGroup_SetTabOrder.Call(obj, uintptr(value))
}

func RadioGroup_GetTabStop(obj uintptr) bool {
    ret, _, _ := radioGroup_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetTabStop(obj uintptr, value bool) {
   radioGroup_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetVisible(obj uintptr) bool {
    ret, _, _ := radioGroup_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetVisible(obj uintptr, value bool) {
   radioGroup_SetVisible.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := radioGroup_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func RadioGroup_SetStyleElements(obj uintptr, value TStyleElements) {
   radioGroup_SetStyleElements.Call(obj, uintptr(value))
}

func RadioGroup_GetWordWrap(obj uintptr) bool {
    ret, _, _ := radioGroup_GetWordWrap.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetWordWrap(obj uintptr, value bool) {
   radioGroup_SetWordWrap.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_SetOnClick(obj uintptr, fn interface{}) {
    radioGroup_SetOnClick.Call(obj, addEventToMap(fn))
}

func RadioGroup_SetOnEnter(obj uintptr, fn interface{}) {
    radioGroup_SetOnEnter.Call(obj, addEventToMap(fn))
}

func RadioGroup_SetOnExit(obj uintptr, fn interface{}) {
    radioGroup_SetOnExit.Call(obj, addEventToMap(fn))
}

func RadioGroup_GetBrush(obj uintptr) uintptr {
    ret, _, _ := radioGroup_GetBrush.Call(obj)
    return ret
}

func RadioGroup_GetControlCount(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetControlCount.Call(obj)
    return int32(ret)
}

func RadioGroup_GetHandle(obj uintptr) HWND {
    ret, _, _ := radioGroup_GetHandle.Call(obj)
    return HWND(ret)
}

func RadioGroup_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := radioGroup_GetParentWindow.Call(obj)
    return HWND(ret)
}

func RadioGroup_SetParentWindow(obj uintptr, value HWND) {
   radioGroup_SetParentWindow.Call(obj, uintptr(value))
}

func RadioGroup_GetAction(obj uintptr) uintptr {
    ret, _, _ := radioGroup_GetAction.Call(obj)
    return ret
}

func RadioGroup_SetAction(obj uintptr, value uintptr) {
   radioGroup_SetAction.Call(obj, value)
}

func RadioGroup_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    radioGroup_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func RadioGroup_SetBoundsRect(obj uintptr, value TRect) {
   radioGroup_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func RadioGroup_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetClientHeight.Call(obj)
    return int32(ret)
}

func RadioGroup_SetClientHeight(obj uintptr, value int32) {
   radioGroup_SetClientHeight.Call(obj, uintptr(value))
}

func RadioGroup_GetClientRect(obj uintptr) TRect {
    var ret TRect
    radioGroup_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func RadioGroup_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetClientWidth.Call(obj)
    return int32(ret)
}

func RadioGroup_SetClientWidth(obj uintptr, value int32) {
   radioGroup_SetClientWidth.Call(obj, uintptr(value))
}

func RadioGroup_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func RadioGroup_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetExplicitTop.Call(obj)
    return int32(ret)
}

func RadioGroup_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func RadioGroup_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func RadioGroup_GetParent(obj uintptr) uintptr {
    ret, _, _ := radioGroup_GetParent.Call(obj)
    return ret
}

func RadioGroup_SetParent(obj uintptr, value uintptr) {
   radioGroup_SetParent.Call(obj, value)
}

func RadioGroup_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := radioGroup_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func RadioGroup_SetAlignWithMargins(obj uintptr, value bool) {
   radioGroup_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func RadioGroup_GetLeft(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetLeft.Call(obj)
    return int32(ret)
}

func RadioGroup_SetLeft(obj uintptr, value int32) {
   radioGroup_SetLeft.Call(obj, uintptr(value))
}

func RadioGroup_GetTop(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetTop.Call(obj)
    return int32(ret)
}

func RadioGroup_SetTop(obj uintptr, value int32) {
   radioGroup_SetTop.Call(obj, uintptr(value))
}

func RadioGroup_GetWidth(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetWidth.Call(obj)
    return int32(ret)
}

func RadioGroup_SetWidth(obj uintptr, value int32) {
   radioGroup_SetWidth.Call(obj, uintptr(value))
}

func RadioGroup_GetHeight(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetHeight.Call(obj)
    return int32(ret)
}

func RadioGroup_SetHeight(obj uintptr, value int32) {
   radioGroup_SetHeight.Call(obj, uintptr(value))
}

func RadioGroup_GetCursor(obj uintptr) TCursor {
    ret, _, _ := radioGroup_GetCursor.Call(obj)
    return TCursor(ret)
}

func RadioGroup_SetCursor(obj uintptr, value TCursor) {
   radioGroup_SetCursor.Call(obj, uintptr(value))
}

func RadioGroup_GetHint(obj uintptr) string {
    ret, _, _ := radioGroup_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func RadioGroup_SetHint(obj uintptr, value string) {
   radioGroup_SetHint.Call(obj, GoStrToDStr(value))
}

func RadioGroup_GetMargins(obj uintptr) uintptr {
    ret, _, _ := radioGroup_GetMargins.Call(obj)
    return ret
}

func RadioGroup_SetMargins(obj uintptr, value uintptr) {
   radioGroup_SetMargins.Call(obj, value)
}

func RadioGroup_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetComponentCount.Call(obj)
    return int32(ret)
}

func RadioGroup_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := radioGroup_GetComponentIndex.Call(obj)
    return int32(ret)
}

func RadioGroup_SetComponentIndex(obj uintptr, value int32) {
   radioGroup_SetComponentIndex.Call(obj, uintptr(value))
}

func RadioGroup_GetOwner(obj uintptr) uintptr {
    ret, _, _ := radioGroup_GetOwner.Call(obj)
    return ret
}

func RadioGroup_GetName(obj uintptr) string {
    ret, _, _ := radioGroup_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func RadioGroup_SetName(obj uintptr, value string) {
   radioGroup_SetName.Call(obj, GoStrToDStr(value))
}

func RadioGroup_GetTag(obj uintptr) int {
    ret, _, _ := radioGroup_GetTag.Call(obj)
    return int(ret)
}

func RadioGroup_SetTag(obj uintptr, value int) {
   radioGroup_SetTag.Call(obj, uintptr(value))
}

func RadioGroup_GetButtons(obj uintptr, Index int32) uintptr {
    ret, _, _ := radioGroup_GetButtons.Call(obj, uintptr(Index))
    return ret
}

func RadioGroup_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := radioGroup_GetControls.Call(obj, uintptr(Index))
    return ret
}

func RadioGroup_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := radioGroup_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TStaticText ---------------------------

func StaticText_Create(obj uintptr) uintptr {
    ret, _, _ := staticText_Create.Call(obj)
    return ret
}

func StaticText_Free(obj uintptr) {
    staticText_Free.Call(obj)
}

func StaticText_CanFocus(obj uintptr) bool {
    ret, _, _ := staticText_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_FlipChildren(obj uintptr, AllLevels bool)  {
    staticText_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func StaticText_Focused(obj uintptr) bool {
    ret, _, _ := staticText_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_HandleAllocated(obj uintptr) bool {
    ret, _, _ := staticText_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_Invalidate(obj uintptr)  {
    staticText_Invalidate.Call(obj)
}

func StaticText_Realign(obj uintptr)  {
    staticText_Realign.Call(obj)
}

func StaticText_Repaint(obj uintptr)  {
    staticText_Repaint.Call(obj)
}

func StaticText_ScaleBy(obj uintptr, M int32, D int32)  {
    staticText_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func StaticText_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    staticText_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func StaticText_SetFocus(obj uintptr)  {
    staticText_SetFocus.Call(obj)
}

func StaticText_Update(obj uintptr)  {
    staticText_Update.Call(obj)
}

func StaticText_BringToFront(obj uintptr)  {
    staticText_BringToFront.Call(obj)
}

func StaticText_HasParent(obj uintptr) bool {
    ret, _, _ := staticText_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_Hide(obj uintptr)  {
    staticText_Hide.Call(obj)
}

func StaticText_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := staticText_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func StaticText_Refresh(obj uintptr)  {
    staticText_Refresh.Call(obj)
}

func StaticText_SendToBack(obj uintptr)  {
    staticText_SendToBack.Call(obj)
}

func StaticText_Show(obj uintptr)  {
    staticText_Show.Call(obj)
}

func StaticText_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := staticText_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func StaticText_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := staticText_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func StaticText_GetNamePath(obj uintptr) string {
    ret, _, _ := staticText_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func StaticText_Assign(obj uintptr, Source uintptr)  {
    staticText_Assign.Call(obj, Source )
}

func StaticText_ClassName(obj uintptr) string {
    ret, _, _ := staticText_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func StaticText_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := staticText_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func StaticText_GetHashCode(obj uintptr) int32 {
    ret, _, _ := staticText_GetHashCode.Call(obj)
    return int32(ret)
}

func StaticText_ToString(obj uintptr) string {
    ret, _, _ := staticText_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func StaticText_GetAlign(obj uintptr) TAlign {
    ret, _, _ := staticText_GetAlign.Call(obj)
    return TAlign(ret)
}

func StaticText_SetAlign(obj uintptr, value TAlign) {
   staticText_SetAlign.Call(obj, uintptr(value))
}

func StaticText_GetAlignment(obj uintptr) TAlignment {
    ret, _, _ := staticText_GetAlignment.Call(obj)
    return TAlignment(ret)
}

func StaticText_SetAlignment(obj uintptr, value TAlignment) {
   staticText_SetAlignment.Call(obj, uintptr(value))
}

func StaticText_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := staticText_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func StaticText_SetAnchors(obj uintptr, value TAnchors) {
   staticText_SetAnchors.Call(obj, uintptr(value))
}

func StaticText_GetAutoSize(obj uintptr) bool {
    ret, _, _ := staticText_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetAutoSize(obj uintptr, value bool) {
   staticText_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetBevelEdges(obj uintptr) TBevelEdges {
    ret, _, _ := staticText_GetBevelEdges.Call(obj)
    return TBevelEdges(ret)
}

func StaticText_SetBevelEdges(obj uintptr, value TBevelEdges) {
   staticText_SetBevelEdges.Call(obj, uintptr(value))
}

func StaticText_GetBevelInner(obj uintptr) TBevelCut {
    ret, _, _ := staticText_GetBevelInner.Call(obj)
    return TBevelCut(ret)
}

func StaticText_SetBevelInner(obj uintptr, value TBevelCut) {
   staticText_SetBevelInner.Call(obj, uintptr(value))
}

func StaticText_GetBevelKind(obj uintptr) TBevelKind {
    ret, _, _ := staticText_GetBevelKind.Call(obj)
    return TBevelKind(ret)
}

func StaticText_SetBevelKind(obj uintptr, value TBevelKind) {
   staticText_SetBevelKind.Call(obj, uintptr(value))
}

func StaticText_GetBevelOuter(obj uintptr) TBevelCut {
    ret, _, _ := staticText_GetBevelOuter.Call(obj)
    return TBevelCut(ret)
}

func StaticText_SetBevelOuter(obj uintptr, value TBevelCut) {
   staticText_SetBevelOuter.Call(obj, uintptr(value))
}

func StaticText_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := staticText_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func StaticText_SetBiDiMode(obj uintptr, value TBiDiMode) {
   staticText_SetBiDiMode.Call(obj, uintptr(value))
}

func StaticText_GetBorderStyle(obj uintptr) TStaticBorderStyle {
    ret, _, _ := staticText_GetBorderStyle.Call(obj)
    return TStaticBorderStyle(ret)
}

func StaticText_SetBorderStyle(obj uintptr, value TStaticBorderStyle) {
   staticText_SetBorderStyle.Call(obj, uintptr(value))
}

func StaticText_GetCaption(obj uintptr) string {
    ret, _, _ := staticText_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func StaticText_SetCaption(obj uintptr, value string) {
   staticText_SetCaption.Call(obj, GoStrToDStr(value))
}

func StaticText_GetColor(obj uintptr) TColor {
    ret, _, _ := staticText_GetColor.Call(obj)
    return TColor(ret)
}

func StaticText_SetColor(obj uintptr, value TColor) {
   staticText_SetColor.Call(obj, uintptr(value))
}

func StaticText_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := staticText_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetDoubleBuffered(obj uintptr, value bool) {
   staticText_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetEnabled(obj uintptr) bool {
    ret, _, _ := staticText_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetEnabled(obj uintptr, value bool) {
   staticText_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetFont(obj uintptr) uintptr {
    ret, _, _ := staticText_GetFont.Call(obj)
    return ret
}

func StaticText_SetFont(obj uintptr, value uintptr) {
   staticText_SetFont.Call(obj, value)
}

func StaticText_GetParentColor(obj uintptr) bool {
    ret, _, _ := staticText_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetParentColor(obj uintptr, value bool) {
   staticText_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := staticText_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetParentDoubleBuffered(obj uintptr, value bool) {
   staticText_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetParentFont(obj uintptr) bool {
    ret, _, _ := staticText_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetParentFont(obj uintptr, value bool) {
   staticText_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := staticText_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetParentShowHint(obj uintptr, value bool) {
   staticText_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := staticText_GetPopupMenu.Call(obj)
    return ret
}

func StaticText_SetPopupMenu(obj uintptr, value uintptr) {
   staticText_SetPopupMenu.Call(obj, value)
}

func StaticText_GetShowAccelChar(obj uintptr) bool {
    ret, _, _ := staticText_GetShowAccelChar.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetShowAccelChar(obj uintptr, value bool) {
   staticText_SetShowAccelChar.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetShowHint(obj uintptr) bool {
    ret, _, _ := staticText_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetShowHint(obj uintptr, value bool) {
   staticText_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := staticText_GetTabOrder.Call(obj)
    return uint16(ret)
}

func StaticText_SetTabOrder(obj uintptr, value uint16) {
   staticText_SetTabOrder.Call(obj, uintptr(value))
}

func StaticText_GetTabStop(obj uintptr) bool {
    ret, _, _ := staticText_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetTabStop(obj uintptr, value bool) {
   staticText_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetTransparent(obj uintptr) bool {
    ret, _, _ := staticText_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetTransparent(obj uintptr, value bool) {
   staticText_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetVisible(obj uintptr) bool {
    ret, _, _ := staticText_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetVisible(obj uintptr, value bool) {
   staticText_SetVisible.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := staticText_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func StaticText_SetStyleElements(obj uintptr, value TStyleElements) {
   staticText_SetStyleElements.Call(obj, uintptr(value))
}

func StaticText_SetOnClick(obj uintptr, fn interface{}) {
    staticText_SetOnClick.Call(obj, addEventToMap(fn))
}

func StaticText_SetOnDblClick(obj uintptr, fn interface{}) {
    staticText_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func StaticText_SetOnMouseDown(obj uintptr, fn interface{}) {
    staticText_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func StaticText_SetOnMouseEnter(obj uintptr, fn interface{}) {
    staticText_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func StaticText_SetOnMouseLeave(obj uintptr, fn interface{}) {
    staticText_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func StaticText_SetOnMouseMove(obj uintptr, fn interface{}) {
    staticText_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func StaticText_SetOnMouseUp(obj uintptr, fn interface{}) {
    staticText_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func StaticText_GetBrush(obj uintptr) uintptr {
    ret, _, _ := staticText_GetBrush.Call(obj)
    return ret
}

func StaticText_GetControlCount(obj uintptr) int32 {
    ret, _, _ := staticText_GetControlCount.Call(obj)
    return int32(ret)
}

func StaticText_GetHandle(obj uintptr) HWND {
    ret, _, _ := staticText_GetHandle.Call(obj)
    return HWND(ret)
}

func StaticText_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := staticText_GetParentWindow.Call(obj)
    return HWND(ret)
}

func StaticText_SetParentWindow(obj uintptr, value HWND) {
   staticText_SetParentWindow.Call(obj, uintptr(value))
}

func StaticText_GetAction(obj uintptr) uintptr {
    ret, _, _ := staticText_GetAction.Call(obj)
    return ret
}

func StaticText_SetAction(obj uintptr, value uintptr) {
   staticText_SetAction.Call(obj, value)
}

func StaticText_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    staticText_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func StaticText_SetBoundsRect(obj uintptr, value TRect) {
   staticText_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func StaticText_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := staticText_GetClientHeight.Call(obj)
    return int32(ret)
}

func StaticText_SetClientHeight(obj uintptr, value int32) {
   staticText_SetClientHeight.Call(obj, uintptr(value))
}

func StaticText_GetClientRect(obj uintptr) TRect {
    var ret TRect
    staticText_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func StaticText_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := staticText_GetClientWidth.Call(obj)
    return int32(ret)
}

func StaticText_SetClientWidth(obj uintptr, value int32) {
   staticText_SetClientWidth.Call(obj, uintptr(value))
}

func StaticText_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := staticText_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func StaticText_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := staticText_GetExplicitTop.Call(obj)
    return int32(ret)
}

func StaticText_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := staticText_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func StaticText_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := staticText_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func StaticText_GetParent(obj uintptr) uintptr {
    ret, _, _ := staticText_GetParent.Call(obj)
    return ret
}

func StaticText_SetParent(obj uintptr, value uintptr) {
   staticText_SetParent.Call(obj, value)
}

func StaticText_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := staticText_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func StaticText_SetAlignWithMargins(obj uintptr, value bool) {
   staticText_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func StaticText_GetLeft(obj uintptr) int32 {
    ret, _, _ := staticText_GetLeft.Call(obj)
    return int32(ret)
}

func StaticText_SetLeft(obj uintptr, value int32) {
   staticText_SetLeft.Call(obj, uintptr(value))
}

func StaticText_GetTop(obj uintptr) int32 {
    ret, _, _ := staticText_GetTop.Call(obj)
    return int32(ret)
}

func StaticText_SetTop(obj uintptr, value int32) {
   staticText_SetTop.Call(obj, uintptr(value))
}

func StaticText_GetWidth(obj uintptr) int32 {
    ret, _, _ := staticText_GetWidth.Call(obj)
    return int32(ret)
}

func StaticText_SetWidth(obj uintptr, value int32) {
   staticText_SetWidth.Call(obj, uintptr(value))
}

func StaticText_GetHeight(obj uintptr) int32 {
    ret, _, _ := staticText_GetHeight.Call(obj)
    return int32(ret)
}

func StaticText_SetHeight(obj uintptr, value int32) {
   staticText_SetHeight.Call(obj, uintptr(value))
}

func StaticText_GetCursor(obj uintptr) TCursor {
    ret, _, _ := staticText_GetCursor.Call(obj)
    return TCursor(ret)
}

func StaticText_SetCursor(obj uintptr, value TCursor) {
   staticText_SetCursor.Call(obj, uintptr(value))
}

func StaticText_GetHint(obj uintptr) string {
    ret, _, _ := staticText_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func StaticText_SetHint(obj uintptr, value string) {
   staticText_SetHint.Call(obj, GoStrToDStr(value))
}

func StaticText_GetMargins(obj uintptr) uintptr {
    ret, _, _ := staticText_GetMargins.Call(obj)
    return ret
}

func StaticText_SetMargins(obj uintptr, value uintptr) {
   staticText_SetMargins.Call(obj, value)
}

func StaticText_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := staticText_GetComponentCount.Call(obj)
    return int32(ret)
}

func StaticText_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := staticText_GetComponentIndex.Call(obj)
    return int32(ret)
}

func StaticText_SetComponentIndex(obj uintptr, value int32) {
   staticText_SetComponentIndex.Call(obj, uintptr(value))
}

func StaticText_GetOwner(obj uintptr) uintptr {
    ret, _, _ := staticText_GetOwner.Call(obj)
    return ret
}

func StaticText_GetName(obj uintptr) string {
    ret, _, _ := staticText_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func StaticText_SetName(obj uintptr, value string) {
   staticText_SetName.Call(obj, GoStrToDStr(value))
}

func StaticText_GetTag(obj uintptr) int {
    ret, _, _ := staticText_GetTag.Call(obj)
    return int(ret)
}

func StaticText_SetTag(obj uintptr, value int) {
   staticText_SetTag.Call(obj, uintptr(value))
}

func StaticText_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := staticText_GetControls.Call(obj, uintptr(Index))
    return ret
}

func StaticText_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := staticText_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TColorBox ---------------------------

func ColorBox_Create(obj uintptr) uintptr {
    ret, _, _ := colorBox_Create.Call(obj)
    return ret
}

func ColorBox_Free(obj uintptr) {
    colorBox_Free.Call(obj)
}

func ColorBox_AddItem(obj uintptr, Item string, AObject uintptr)  {
    colorBox_AddItem.Call(obj, GoStrToDStr(Item) , AObject )
}

func ColorBox_Clear(obj uintptr)  {
    colorBox_Clear.Call(obj)
}

func ColorBox_ClearSelection(obj uintptr)  {
    colorBox_ClearSelection.Call(obj)
}

func ColorBox_DeleteSelected(obj uintptr)  {
    colorBox_DeleteSelected.Call(obj)
}

func ColorBox_Focused(obj uintptr) bool {
    ret, _, _ := colorBox_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SelectAll(obj uintptr)  {
    colorBox_SelectAll.Call(obj)
}

func ColorBox_CanFocus(obj uintptr) bool {
    ret, _, _ := colorBox_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_FlipChildren(obj uintptr, AllLevels bool)  {
    colorBox_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func ColorBox_HandleAllocated(obj uintptr) bool {
    ret, _, _ := colorBox_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_Invalidate(obj uintptr)  {
    colorBox_Invalidate.Call(obj)
}

func ColorBox_Realign(obj uintptr)  {
    colorBox_Realign.Call(obj)
}

func ColorBox_Repaint(obj uintptr)  {
    colorBox_Repaint.Call(obj)
}

func ColorBox_ScaleBy(obj uintptr, M int32, D int32)  {
    colorBox_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func ColorBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    colorBox_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func ColorBox_SetFocus(obj uintptr)  {
    colorBox_SetFocus.Call(obj)
}

func ColorBox_Update(obj uintptr)  {
    colorBox_Update.Call(obj)
}

func ColorBox_BringToFront(obj uintptr)  {
    colorBox_BringToFront.Call(obj)
}

func ColorBox_HasParent(obj uintptr) bool {
    ret, _, _ := colorBox_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_Hide(obj uintptr)  {
    colorBox_Hide.Call(obj)
}

func ColorBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := colorBox_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func ColorBox_Refresh(obj uintptr)  {
    colorBox_Refresh.Call(obj)
}

func ColorBox_SendToBack(obj uintptr)  {
    colorBox_SendToBack.Call(obj)
}

func ColorBox_Show(obj uintptr)  {
    colorBox_Show.Call(obj)
}

func ColorBox_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := colorBox_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func ColorBox_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := colorBox_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ColorBox_GetNamePath(obj uintptr) string {
    ret, _, _ := colorBox_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ColorBox_Assign(obj uintptr, Source uintptr)  {
    colorBox_Assign.Call(obj, Source )
}

func ColorBox_ClassName(obj uintptr) string {
    ret, _, _ := colorBox_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ColorBox_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := colorBox_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ColorBox_GetHashCode(obj uintptr) int32 {
    ret, _, _ := colorBox_GetHashCode.Call(obj)
    return int32(ret)
}

func ColorBox_ToString(obj uintptr) string {
    ret, _, _ := colorBox_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ColorBox_GetAlign(obj uintptr) TAlign {
    ret, _, _ := colorBox_GetAlign.Call(obj)
    return TAlign(ret)
}

func ColorBox_SetAlign(obj uintptr, value TAlign) {
   colorBox_SetAlign.Call(obj, uintptr(value))
}

func ColorBox_GetAutoComplete(obj uintptr) bool {
    ret, _, _ := colorBox_GetAutoComplete.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetAutoComplete(obj uintptr, value bool) {
   colorBox_SetAutoComplete.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetAutoDropDown(obj uintptr) bool {
    ret, _, _ := colorBox_GetAutoDropDown.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetAutoDropDown(obj uintptr, value bool) {
   colorBox_SetAutoDropDown.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetDefaultColorColor(obj uintptr) TColor {
    ret, _, _ := colorBox_GetDefaultColorColor.Call(obj)
    return TColor(ret)
}

func ColorBox_SetDefaultColorColor(obj uintptr, value TColor) {
   colorBox_SetDefaultColorColor.Call(obj, uintptr(value))
}

func ColorBox_GetNoneColorColor(obj uintptr) TColor {
    ret, _, _ := colorBox_GetNoneColorColor.Call(obj)
    return TColor(ret)
}

func ColorBox_SetNoneColorColor(obj uintptr, value TColor) {
   colorBox_SetNoneColorColor.Call(obj, uintptr(value))
}

func ColorBox_GetSelected(obj uintptr) TColor {
    ret, _, _ := colorBox_GetSelected.Call(obj)
    return TColor(ret)
}

func ColorBox_SetSelected(obj uintptr, value TColor) {
   colorBox_SetSelected.Call(obj, uintptr(value))
}

func ColorBox_GetStyle(obj uintptr) TColorBoxStyle {
    ret, _, _ := colorBox_GetStyle.Call(obj)
    return TColorBoxStyle(ret)
}

func ColorBox_SetStyle(obj uintptr, value TColorBoxStyle) {
   colorBox_SetStyle.Call(obj, uintptr(value))
}

func ColorBox_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := colorBox_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func ColorBox_SetAnchors(obj uintptr, value TAnchors) {
   colorBox_SetAnchors.Call(obj, uintptr(value))
}

func ColorBox_GetBevelEdges(obj uintptr) TBevelEdges {
    ret, _, _ := colorBox_GetBevelEdges.Call(obj)
    return TBevelEdges(ret)
}

func ColorBox_SetBevelEdges(obj uintptr, value TBevelEdges) {
   colorBox_SetBevelEdges.Call(obj, uintptr(value))
}

func ColorBox_GetBevelInner(obj uintptr) TBevelCut {
    ret, _, _ := colorBox_GetBevelInner.Call(obj)
    return TBevelCut(ret)
}

func ColorBox_SetBevelInner(obj uintptr, value TBevelCut) {
   colorBox_SetBevelInner.Call(obj, uintptr(value))
}

func ColorBox_GetBevelKind(obj uintptr) TBevelKind {
    ret, _, _ := colorBox_GetBevelKind.Call(obj)
    return TBevelKind(ret)
}

func ColorBox_SetBevelKind(obj uintptr, value TBevelKind) {
   colorBox_SetBevelKind.Call(obj, uintptr(value))
}

func ColorBox_GetBevelOuter(obj uintptr) TBevelCut {
    ret, _, _ := colorBox_GetBevelOuter.Call(obj)
    return TBevelCut(ret)
}

func ColorBox_SetBevelOuter(obj uintptr, value TBevelCut) {
   colorBox_SetBevelOuter.Call(obj, uintptr(value))
}

func ColorBox_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := colorBox_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func ColorBox_SetBiDiMode(obj uintptr, value TBiDiMode) {
   colorBox_SetBiDiMode.Call(obj, uintptr(value))
}

func ColorBox_GetColor(obj uintptr) TColor {
    ret, _, _ := colorBox_GetColor.Call(obj)
    return TColor(ret)
}

func ColorBox_SetColor(obj uintptr, value TColor) {
   colorBox_SetColor.Call(obj, uintptr(value))
}

func ColorBox_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := colorBox_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetDoubleBuffered(obj uintptr, value bool) {
   colorBox_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetDropDownCount(obj uintptr) int32 {
    ret, _, _ := colorBox_GetDropDownCount.Call(obj)
    return int32(ret)
}

func ColorBox_SetDropDownCount(obj uintptr, value int32) {
   colorBox_SetDropDownCount.Call(obj, uintptr(value))
}

func ColorBox_GetEnabled(obj uintptr) bool {
    ret, _, _ := colorBox_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetEnabled(obj uintptr, value bool) {
   colorBox_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetFont(obj uintptr) uintptr {
    ret, _, _ := colorBox_GetFont.Call(obj)
    return ret
}

func ColorBox_SetFont(obj uintptr, value uintptr) {
   colorBox_SetFont.Call(obj, value)
}

func ColorBox_GetItemHeight(obj uintptr) int32 {
    ret, _, _ := colorBox_GetItemHeight.Call(obj)
    return int32(ret)
}

func ColorBox_SetItemHeight(obj uintptr, value int32) {
   colorBox_SetItemHeight.Call(obj, uintptr(value))
}

func ColorBox_GetParentColor(obj uintptr) bool {
    ret, _, _ := colorBox_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetParentColor(obj uintptr, value bool) {
   colorBox_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := colorBox_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetParentCtl3D(obj uintptr, value bool) {
   colorBox_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := colorBox_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetParentDoubleBuffered(obj uintptr, value bool) {
   colorBox_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetParentFont(obj uintptr) bool {
    ret, _, _ := colorBox_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetParentFont(obj uintptr, value bool) {
   colorBox_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := colorBox_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetParentShowHint(obj uintptr, value bool) {
   colorBox_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := colorBox_GetPopupMenu.Call(obj)
    return ret
}

func ColorBox_SetPopupMenu(obj uintptr, value uintptr) {
   colorBox_SetPopupMenu.Call(obj, value)
}

func ColorBox_GetShowHint(obj uintptr) bool {
    ret, _, _ := colorBox_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetShowHint(obj uintptr, value bool) {
   colorBox_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := colorBox_GetTabOrder.Call(obj)
    return uint16(ret)
}

func ColorBox_SetTabOrder(obj uintptr, value uint16) {
   colorBox_SetTabOrder.Call(obj, uintptr(value))
}

func ColorBox_GetTabStop(obj uintptr) bool {
    ret, _, _ := colorBox_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetTabStop(obj uintptr, value bool) {
   colorBox_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetVisible(obj uintptr) bool {
    ret, _, _ := colorBox_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetVisible(obj uintptr, value bool) {
   colorBox_SetVisible.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := colorBox_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func ColorBox_SetStyleElements(obj uintptr, value TStyleElements) {
   colorBox_SetStyleElements.Call(obj, uintptr(value))
}

func ColorBox_SetOnChange(obj uintptr, fn interface{}) {
    colorBox_SetOnChange.Call(obj, addEventToMap(fn))
}

func ColorBox_SetOnClick(obj uintptr, fn interface{}) {
    colorBox_SetOnClick.Call(obj, addEventToMap(fn))
}

func ColorBox_SetOnEnter(obj uintptr, fn interface{}) {
    colorBox_SetOnEnter.Call(obj, addEventToMap(fn))
}

func ColorBox_SetOnExit(obj uintptr, fn interface{}) {
    colorBox_SetOnExit.Call(obj, addEventToMap(fn))
}

func ColorBox_SetOnKeyDown(obj uintptr, fn interface{}) {
    colorBox_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func ColorBox_SetOnKeyPress(obj uintptr, fn interface{}) {
    colorBox_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func ColorBox_SetOnKeyUp(obj uintptr, fn interface{}) {
    colorBox_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func ColorBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
    colorBox_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func ColorBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
    colorBox_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func ColorBox_GetAutoCompleteDelay(obj uintptr) uint32 {
    ret, _, _ := colorBox_GetAutoCompleteDelay.Call(obj)
    return uint32(ret)
}

func ColorBox_SetAutoCompleteDelay(obj uintptr, value uint32) {
   colorBox_SetAutoCompleteDelay.Call(obj, uintptr(value))
}

func ColorBox_GetAutoCloseUp(obj uintptr) bool {
    ret, _, _ := colorBox_GetAutoCloseUp.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetAutoCloseUp(obj uintptr, value bool) {
   colorBox_SetAutoCloseUp.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetSelText(obj uintptr) string {
    ret, _, _ := colorBox_GetSelText.Call(obj)
    return DStrToGoStr(ret)
}

func ColorBox_SetSelText(obj uintptr, value string) {
   colorBox_SetSelText.Call(obj, GoStrToDStr(value))
}

func ColorBox_GetTextHint(obj uintptr) string {
    ret, _, _ := colorBox_GetTextHint.Call(obj)
    return DStrToGoStr(ret)
}

func ColorBox_SetTextHint(obj uintptr, value string) {
   colorBox_SetTextHint.Call(obj, GoStrToDStr(value))
}

func ColorBox_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := colorBox_GetCanvas.Call(obj)
    return ret
}

func ColorBox_GetDroppedDown(obj uintptr) bool {
    ret, _, _ := colorBox_GetDroppedDown.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetDroppedDown(obj uintptr, value bool) {
   colorBox_SetDroppedDown.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetItems(obj uintptr) uintptr {
    ret, _, _ := colorBox_GetItems.Call(obj)
    return ret
}

func ColorBox_SetItems(obj uintptr, value uintptr) {
   colorBox_SetItems.Call(obj, value)
}

func ColorBox_GetSelLength(obj uintptr) int32 {
    ret, _, _ := colorBox_GetSelLength.Call(obj)
    return int32(ret)
}

func ColorBox_SetSelLength(obj uintptr, value int32) {
   colorBox_SetSelLength.Call(obj, uintptr(value))
}

func ColorBox_GetSelStart(obj uintptr) int32 {
    ret, _, _ := colorBox_GetSelStart.Call(obj)
    return int32(ret)
}

func ColorBox_SetSelStart(obj uintptr, value int32) {
   colorBox_SetSelStart.Call(obj, uintptr(value))
}

func ColorBox_GetItemIndex(obj uintptr) int32 {
    ret, _, _ := colorBox_GetItemIndex.Call(obj)
    return int32(ret)
}

func ColorBox_SetItemIndex(obj uintptr, value int32) {
   colorBox_SetItemIndex.Call(obj, uintptr(value))
}

func ColorBox_GetBrush(obj uintptr) uintptr {
    ret, _, _ := colorBox_GetBrush.Call(obj)
    return ret
}

func ColorBox_GetControlCount(obj uintptr) int32 {
    ret, _, _ := colorBox_GetControlCount.Call(obj)
    return int32(ret)
}

func ColorBox_GetHandle(obj uintptr) HWND {
    ret, _, _ := colorBox_GetHandle.Call(obj)
    return HWND(ret)
}

func ColorBox_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := colorBox_GetParentWindow.Call(obj)
    return HWND(ret)
}

func ColorBox_SetParentWindow(obj uintptr, value HWND) {
   colorBox_SetParentWindow.Call(obj, uintptr(value))
}

func ColorBox_GetAction(obj uintptr) uintptr {
    ret, _, _ := colorBox_GetAction.Call(obj)
    return ret
}

func ColorBox_SetAction(obj uintptr, value uintptr) {
   colorBox_SetAction.Call(obj, value)
}

func ColorBox_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    colorBox_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ColorBox_SetBoundsRect(obj uintptr, value TRect) {
   colorBox_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ColorBox_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := colorBox_GetClientHeight.Call(obj)
    return int32(ret)
}

func ColorBox_SetClientHeight(obj uintptr, value int32) {
   colorBox_SetClientHeight.Call(obj, uintptr(value))
}

func ColorBox_GetClientRect(obj uintptr) TRect {
    var ret TRect
    colorBox_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ColorBox_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := colorBox_GetClientWidth.Call(obj)
    return int32(ret)
}

func ColorBox_SetClientWidth(obj uintptr, value int32) {
   colorBox_SetClientWidth.Call(obj, uintptr(value))
}

func ColorBox_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := colorBox_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func ColorBox_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := colorBox_GetExplicitTop.Call(obj)
    return int32(ret)
}

func ColorBox_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := colorBox_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func ColorBox_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := colorBox_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func ColorBox_GetParent(obj uintptr) uintptr {
    ret, _, _ := colorBox_GetParent.Call(obj)
    return ret
}

func ColorBox_SetParent(obj uintptr, value uintptr) {
   colorBox_SetParent.Call(obj, value)
}

func ColorBox_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := colorBox_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorBox_SetAlignWithMargins(obj uintptr, value bool) {
   colorBox_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func ColorBox_GetLeft(obj uintptr) int32 {
    ret, _, _ := colorBox_GetLeft.Call(obj)
    return int32(ret)
}

func ColorBox_SetLeft(obj uintptr, value int32) {
   colorBox_SetLeft.Call(obj, uintptr(value))
}

func ColorBox_GetTop(obj uintptr) int32 {
    ret, _, _ := colorBox_GetTop.Call(obj)
    return int32(ret)
}

func ColorBox_SetTop(obj uintptr, value int32) {
   colorBox_SetTop.Call(obj, uintptr(value))
}

func ColorBox_GetWidth(obj uintptr) int32 {
    ret, _, _ := colorBox_GetWidth.Call(obj)
    return int32(ret)
}

func ColorBox_SetWidth(obj uintptr, value int32) {
   colorBox_SetWidth.Call(obj, uintptr(value))
}

func ColorBox_GetHeight(obj uintptr) int32 {
    ret, _, _ := colorBox_GetHeight.Call(obj)
    return int32(ret)
}

func ColorBox_SetHeight(obj uintptr, value int32) {
   colorBox_SetHeight.Call(obj, uintptr(value))
}

func ColorBox_GetCursor(obj uintptr) TCursor {
    ret, _, _ := colorBox_GetCursor.Call(obj)
    return TCursor(ret)
}

func ColorBox_SetCursor(obj uintptr, value TCursor) {
   colorBox_SetCursor.Call(obj, uintptr(value))
}

func ColorBox_GetHint(obj uintptr) string {
    ret, _, _ := colorBox_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func ColorBox_SetHint(obj uintptr, value string) {
   colorBox_SetHint.Call(obj, GoStrToDStr(value))
}

func ColorBox_GetMargins(obj uintptr) uintptr {
    ret, _, _ := colorBox_GetMargins.Call(obj)
    return ret
}

func ColorBox_SetMargins(obj uintptr, value uintptr) {
   colorBox_SetMargins.Call(obj, value)
}

func ColorBox_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := colorBox_GetComponentCount.Call(obj)
    return int32(ret)
}

func ColorBox_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := colorBox_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ColorBox_SetComponentIndex(obj uintptr, value int32) {
   colorBox_SetComponentIndex.Call(obj, uintptr(value))
}

func ColorBox_GetOwner(obj uintptr) uintptr {
    ret, _, _ := colorBox_GetOwner.Call(obj)
    return ret
}

func ColorBox_GetName(obj uintptr) string {
    ret, _, _ := colorBox_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ColorBox_SetName(obj uintptr, value string) {
   colorBox_SetName.Call(obj, GoStrToDStr(value))
}

func ColorBox_GetTag(obj uintptr) int {
    ret, _, _ := colorBox_GetTag.Call(obj)
    return int(ret)
}

func ColorBox_SetTag(obj uintptr, value int) {
   colorBox_SetTag.Call(obj, uintptr(value))
}

func ColorBox_GetColors(obj uintptr, Index int32) TColor {
    ret, _, _ := colorBox_GetColors.Call(obj, uintptr(Index))
    return TColor(ret)
}

func ColorBox_GetColorNames(obj uintptr, Index int32) string {
    ret, _, _ := colorBox_GetColorNames.Call(obj, uintptr(Index))
    return DStrToGoStr(ret)
}

func ColorBox_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := colorBox_GetControls.Call(obj, uintptr(Index))
    return ret
}

func ColorBox_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := colorBox_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TColorListBox ---------------------------

func ColorListBox_Create(obj uintptr) uintptr {
    ret, _, _ := colorListBox_Create.Call(obj)
    return ret
}

func ColorListBox_Free(obj uintptr) {
    colorListBox_Free.Call(obj)
}

func ColorListBox_AddItem(obj uintptr, Item string, AObject uintptr)  {
    colorListBox_AddItem.Call(obj, GoStrToDStr(Item) , AObject )
}

func ColorListBox_Clear(obj uintptr)  {
    colorListBox_Clear.Call(obj)
}

func ColorListBox_ClearSelection(obj uintptr)  {
    colorListBox_ClearSelection.Call(obj)
}

func ColorListBox_DeleteSelected(obj uintptr)  {
    colorListBox_DeleteSelected.Call(obj)
}

func ColorListBox_SelectAll(obj uintptr)  {
    colorListBox_SelectAll.Call(obj)
}

func ColorListBox_CanFocus(obj uintptr) bool {
    ret, _, _ := colorListBox_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_FlipChildren(obj uintptr, AllLevels bool)  {
    colorListBox_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func ColorListBox_Focused(obj uintptr) bool {
    ret, _, _ := colorListBox_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_HandleAllocated(obj uintptr) bool {
    ret, _, _ := colorListBox_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_Invalidate(obj uintptr)  {
    colorListBox_Invalidate.Call(obj)
}

func ColorListBox_Realign(obj uintptr)  {
    colorListBox_Realign.Call(obj)
}

func ColorListBox_Repaint(obj uintptr)  {
    colorListBox_Repaint.Call(obj)
}

func ColorListBox_ScaleBy(obj uintptr, M int32, D int32)  {
    colorListBox_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func ColorListBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    colorListBox_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func ColorListBox_SetFocus(obj uintptr)  {
    colorListBox_SetFocus.Call(obj)
}

func ColorListBox_Update(obj uintptr)  {
    colorListBox_Update.Call(obj)
}

func ColorListBox_BringToFront(obj uintptr)  {
    colorListBox_BringToFront.Call(obj)
}

func ColorListBox_HasParent(obj uintptr) bool {
    ret, _, _ := colorListBox_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_Hide(obj uintptr)  {
    colorListBox_Hide.Call(obj)
}

func ColorListBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := colorListBox_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func ColorListBox_Refresh(obj uintptr)  {
    colorListBox_Refresh.Call(obj)
}

func ColorListBox_SendToBack(obj uintptr)  {
    colorListBox_SendToBack.Call(obj)
}

func ColorListBox_Show(obj uintptr)  {
    colorListBox_Show.Call(obj)
}

func ColorListBox_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := colorListBox_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func ColorListBox_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := colorListBox_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ColorListBox_GetNamePath(obj uintptr) string {
    ret, _, _ := colorListBox_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ColorListBox_Assign(obj uintptr, Source uintptr)  {
    colorListBox_Assign.Call(obj, Source )
}

func ColorListBox_ClassName(obj uintptr) string {
    ret, _, _ := colorListBox_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ColorListBox_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := colorListBox_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ColorListBox_GetHashCode(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetHashCode.Call(obj)
    return int32(ret)
}

func ColorListBox_ToString(obj uintptr) string {
    ret, _, _ := colorListBox_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ColorListBox_GetAlign(obj uintptr) TAlign {
    ret, _, _ := colorListBox_GetAlign.Call(obj)
    return TAlign(ret)
}

func ColorListBox_SetAlign(obj uintptr, value TAlign) {
   colorListBox_SetAlign.Call(obj, uintptr(value))
}

func ColorListBox_GetAutoComplete(obj uintptr) bool {
    ret, _, _ := colorListBox_GetAutoComplete.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetAutoComplete(obj uintptr, value bool) {
   colorListBox_SetAutoComplete.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetDefaultColorColor(obj uintptr) TColor {
    ret, _, _ := colorListBox_GetDefaultColorColor.Call(obj)
    return TColor(ret)
}

func ColorListBox_SetDefaultColorColor(obj uintptr, value TColor) {
   colorListBox_SetDefaultColorColor.Call(obj, uintptr(value))
}

func ColorListBox_GetNoneColorColor(obj uintptr) TColor {
    ret, _, _ := colorListBox_GetNoneColorColor.Call(obj)
    return TColor(ret)
}

func ColorListBox_SetNoneColorColor(obj uintptr, value TColor) {
   colorListBox_SetNoneColorColor.Call(obj, uintptr(value))
}

func ColorListBox_GetSelected(obj uintptr) TColor {
    ret, _, _ := colorListBox_GetSelected.Call(obj)
    return TColor(ret)
}

func ColorListBox_SetSelected(obj uintptr, value TColor) {
   colorListBox_SetSelected.Call(obj, uintptr(value))
}

func ColorListBox_GetStyle(obj uintptr) TColorBoxStyle {
    ret, _, _ := colorListBox_GetStyle.Call(obj)
    return TColorBoxStyle(ret)
}

func ColorListBox_SetStyle(obj uintptr, value TColorBoxStyle) {
   colorListBox_SetStyle.Call(obj, uintptr(value))
}

func ColorListBox_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := colorListBox_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func ColorListBox_SetAnchors(obj uintptr, value TAnchors) {
   colorListBox_SetAnchors.Call(obj, uintptr(value))
}

func ColorListBox_GetBevelEdges(obj uintptr) TBevelEdges {
    ret, _, _ := colorListBox_GetBevelEdges.Call(obj)
    return TBevelEdges(ret)
}

func ColorListBox_SetBevelEdges(obj uintptr, value TBevelEdges) {
   colorListBox_SetBevelEdges.Call(obj, uintptr(value))
}

func ColorListBox_GetBevelInner(obj uintptr) TBevelCut {
    ret, _, _ := colorListBox_GetBevelInner.Call(obj)
    return TBevelCut(ret)
}

func ColorListBox_SetBevelInner(obj uintptr, value TBevelCut) {
   colorListBox_SetBevelInner.Call(obj, uintptr(value))
}

func ColorListBox_GetBevelKind(obj uintptr) TBevelKind {
    ret, _, _ := colorListBox_GetBevelKind.Call(obj)
    return TBevelKind(ret)
}

func ColorListBox_SetBevelKind(obj uintptr, value TBevelKind) {
   colorListBox_SetBevelKind.Call(obj, uintptr(value))
}

func ColorListBox_GetBevelOuter(obj uintptr) TBevelCut {
    ret, _, _ := colorListBox_GetBevelOuter.Call(obj)
    return TBevelCut(ret)
}

func ColorListBox_SetBevelOuter(obj uintptr, value TBevelCut) {
   colorListBox_SetBevelOuter.Call(obj, uintptr(value))
}

func ColorListBox_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := colorListBox_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func ColorListBox_SetBiDiMode(obj uintptr, value TBiDiMode) {
   colorListBox_SetBiDiMode.Call(obj, uintptr(value))
}

func ColorListBox_GetColor(obj uintptr) TColor {
    ret, _, _ := colorListBox_GetColor.Call(obj)
    return TColor(ret)
}

func ColorListBox_SetColor(obj uintptr, value TColor) {
   colorListBox_SetColor.Call(obj, uintptr(value))
}

func ColorListBox_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := colorListBox_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetDoubleBuffered(obj uintptr, value bool) {
   colorListBox_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetEnabled(obj uintptr) bool {
    ret, _, _ := colorListBox_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetEnabled(obj uintptr, value bool) {
   colorListBox_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetFont(obj uintptr) uintptr {
    ret, _, _ := colorListBox_GetFont.Call(obj)
    return ret
}

func ColorListBox_SetFont(obj uintptr, value uintptr) {
   colorListBox_SetFont.Call(obj, value)
}

func ColorListBox_GetItemHeight(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetItemHeight.Call(obj)
    return int32(ret)
}

func ColorListBox_SetItemHeight(obj uintptr, value int32) {
   colorListBox_SetItemHeight.Call(obj, uintptr(value))
}

func ColorListBox_GetParentColor(obj uintptr) bool {
    ret, _, _ := colorListBox_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetParentColor(obj uintptr, value bool) {
   colorListBox_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := colorListBox_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetParentCtl3D(obj uintptr, value bool) {
   colorListBox_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := colorListBox_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetParentDoubleBuffered(obj uintptr, value bool) {
   colorListBox_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetParentFont(obj uintptr) bool {
    ret, _, _ := colorListBox_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetParentFont(obj uintptr, value bool) {
   colorListBox_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := colorListBox_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetParentShowHint(obj uintptr, value bool) {
   colorListBox_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := colorListBox_GetPopupMenu.Call(obj)
    return ret
}

func ColorListBox_SetPopupMenu(obj uintptr, value uintptr) {
   colorListBox_SetPopupMenu.Call(obj, value)
}

func ColorListBox_GetShowHint(obj uintptr) bool {
    ret, _, _ := colorListBox_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetShowHint(obj uintptr, value bool) {
   colorListBox_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := colorListBox_GetTabOrder.Call(obj)
    return uint16(ret)
}

func ColorListBox_SetTabOrder(obj uintptr, value uint16) {
   colorListBox_SetTabOrder.Call(obj, uintptr(value))
}

func ColorListBox_GetTabStop(obj uintptr) bool {
    ret, _, _ := colorListBox_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetTabStop(obj uintptr, value bool) {
   colorListBox_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetVisible(obj uintptr) bool {
    ret, _, _ := colorListBox_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetVisible(obj uintptr, value bool) {
   colorListBox_SetVisible.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := colorListBox_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func ColorListBox_SetStyleElements(obj uintptr, value TStyleElements) {
   colorListBox_SetStyleElements.Call(obj, uintptr(value))
}

func ColorListBox_SetOnClick(obj uintptr, fn interface{}) {
    colorListBox_SetOnClick.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnDblClick(obj uintptr, fn interface{}) {
    colorListBox_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnEnter(obj uintptr, fn interface{}) {
    colorListBox_SetOnEnter.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnExit(obj uintptr, fn interface{}) {
    colorListBox_SetOnExit.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnKeyDown(obj uintptr, fn interface{}) {
    colorListBox_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnKeyPress(obj uintptr, fn interface{}) {
    colorListBox_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnKeyUp(obj uintptr, fn interface{}) {
    colorListBox_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnMouseDown(obj uintptr, fn interface{}) {
    colorListBox_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
    colorListBox_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
    colorListBox_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnMouseMove(obj uintptr, fn interface{}) {
    colorListBox_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func ColorListBox_SetOnMouseUp(obj uintptr, fn interface{}) {
    colorListBox_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func ColorListBox_GetAutoCompleteDelay(obj uintptr) uint32 {
    ret, _, _ := colorListBox_GetAutoCompleteDelay.Call(obj)
    return uint32(ret)
}

func ColorListBox_SetAutoCompleteDelay(obj uintptr, value uint32) {
   colorListBox_SetAutoCompleteDelay.Call(obj, uintptr(value))
}

func ColorListBox_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := colorListBox_GetCanvas.Call(obj)
    return ret
}

func ColorListBox_GetItems(obj uintptr) uintptr {
    ret, _, _ := colorListBox_GetItems.Call(obj)
    return ret
}

func ColorListBox_SetItems(obj uintptr, value uintptr) {
   colorListBox_SetItems.Call(obj, value)
}

func ColorListBox_GetMultiSelect(obj uintptr) bool {
    ret, _, _ := colorListBox_GetMultiSelect.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetMultiSelect(obj uintptr, value bool) {
   colorListBox_SetMultiSelect.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetSelCount(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetSelCount.Call(obj)
    return int32(ret)
}

func ColorListBox_GetItemIndex(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetItemIndex.Call(obj)
    return int32(ret)
}

func ColorListBox_SetItemIndex(obj uintptr, value int32) {
   colorListBox_SetItemIndex.Call(obj, uintptr(value))
}

func ColorListBox_GetBrush(obj uintptr) uintptr {
    ret, _, _ := colorListBox_GetBrush.Call(obj)
    return ret
}

func ColorListBox_GetControlCount(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetControlCount.Call(obj)
    return int32(ret)
}

func ColorListBox_GetHandle(obj uintptr) HWND {
    ret, _, _ := colorListBox_GetHandle.Call(obj)
    return HWND(ret)
}

func ColorListBox_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := colorListBox_GetParentWindow.Call(obj)
    return HWND(ret)
}

func ColorListBox_SetParentWindow(obj uintptr, value HWND) {
   colorListBox_SetParentWindow.Call(obj, uintptr(value))
}

func ColorListBox_GetAction(obj uintptr) uintptr {
    ret, _, _ := colorListBox_GetAction.Call(obj)
    return ret
}

func ColorListBox_SetAction(obj uintptr, value uintptr) {
   colorListBox_SetAction.Call(obj, value)
}

func ColorListBox_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    colorListBox_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ColorListBox_SetBoundsRect(obj uintptr, value TRect) {
   colorListBox_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ColorListBox_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetClientHeight.Call(obj)
    return int32(ret)
}

func ColorListBox_SetClientHeight(obj uintptr, value int32) {
   colorListBox_SetClientHeight.Call(obj, uintptr(value))
}

func ColorListBox_GetClientRect(obj uintptr) TRect {
    var ret TRect
    colorListBox_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ColorListBox_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetClientWidth.Call(obj)
    return int32(ret)
}

func ColorListBox_SetClientWidth(obj uintptr, value int32) {
   colorListBox_SetClientWidth.Call(obj, uintptr(value))
}

func ColorListBox_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func ColorListBox_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetExplicitTop.Call(obj)
    return int32(ret)
}

func ColorListBox_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func ColorListBox_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func ColorListBox_GetParent(obj uintptr) uintptr {
    ret, _, _ := colorListBox_GetParent.Call(obj)
    return ret
}

func ColorListBox_SetParent(obj uintptr, value uintptr) {
   colorListBox_SetParent.Call(obj, value)
}

func ColorListBox_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := colorListBox_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorListBox_SetAlignWithMargins(obj uintptr, value bool) {
   colorListBox_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func ColorListBox_GetLeft(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetLeft.Call(obj)
    return int32(ret)
}

func ColorListBox_SetLeft(obj uintptr, value int32) {
   colorListBox_SetLeft.Call(obj, uintptr(value))
}

func ColorListBox_GetTop(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetTop.Call(obj)
    return int32(ret)
}

func ColorListBox_SetTop(obj uintptr, value int32) {
   colorListBox_SetTop.Call(obj, uintptr(value))
}

func ColorListBox_GetWidth(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetWidth.Call(obj)
    return int32(ret)
}

func ColorListBox_SetWidth(obj uintptr, value int32) {
   colorListBox_SetWidth.Call(obj, uintptr(value))
}

func ColorListBox_GetHeight(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetHeight.Call(obj)
    return int32(ret)
}

func ColorListBox_SetHeight(obj uintptr, value int32) {
   colorListBox_SetHeight.Call(obj, uintptr(value))
}

func ColorListBox_GetCursor(obj uintptr) TCursor {
    ret, _, _ := colorListBox_GetCursor.Call(obj)
    return TCursor(ret)
}

func ColorListBox_SetCursor(obj uintptr, value TCursor) {
   colorListBox_SetCursor.Call(obj, uintptr(value))
}

func ColorListBox_GetHint(obj uintptr) string {
    ret, _, _ := colorListBox_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func ColorListBox_SetHint(obj uintptr, value string) {
   colorListBox_SetHint.Call(obj, GoStrToDStr(value))
}

func ColorListBox_GetMargins(obj uintptr) uintptr {
    ret, _, _ := colorListBox_GetMargins.Call(obj)
    return ret
}

func ColorListBox_SetMargins(obj uintptr, value uintptr) {
   colorListBox_SetMargins.Call(obj, value)
}

func ColorListBox_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetComponentCount.Call(obj)
    return int32(ret)
}

func ColorListBox_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := colorListBox_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ColorListBox_SetComponentIndex(obj uintptr, value int32) {
   colorListBox_SetComponentIndex.Call(obj, uintptr(value))
}

func ColorListBox_GetOwner(obj uintptr) uintptr {
    ret, _, _ := colorListBox_GetOwner.Call(obj)
    return ret
}

func ColorListBox_GetName(obj uintptr) string {
    ret, _, _ := colorListBox_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ColorListBox_SetName(obj uintptr, value string) {
   colorListBox_SetName.Call(obj, GoStrToDStr(value))
}

func ColorListBox_GetTag(obj uintptr) int {
    ret, _, _ := colorListBox_GetTag.Call(obj)
    return int(ret)
}

func ColorListBox_SetTag(obj uintptr, value int) {
   colorListBox_SetTag.Call(obj, uintptr(value))
}

func ColorListBox_GetColors(obj uintptr, Index int32) TColor {
    ret, _, _ := colorListBox_GetColors.Call(obj, uintptr(Index))
    return TColor(ret)
}

func ColorListBox_GetColorNames(obj uintptr, Index int32) string {
    ret, _, _ := colorListBox_GetColorNames.Call(obj, uintptr(Index))
    return DStrToGoStr(ret)
}

func ColorListBox_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := colorListBox_GetControls.Call(obj, uintptr(Index))
    return ret
}

func ColorListBox_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := colorListBox_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TTrayIcon ---------------------------

func TrayIcon_Create(obj uintptr) uintptr {
    ret, _, _ := trayIcon_Create.Call(obj)
    return ret
}

func TrayIcon_Free(obj uintptr) {
    trayIcon_Free.Call(obj)
}

func TrayIcon_Refresh(obj uintptr)  {
    trayIcon_Refresh.Call(obj)
}

func TrayIcon_SetDefaultIcon(obj uintptr)  {
    trayIcon_SetDefaultIcon.Call(obj)
}

func TrayIcon_ShowBalloonHint(obj uintptr)  {
    trayIcon_ShowBalloonHint.Call(obj)
}

func TrayIcon_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := trayIcon_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func TrayIcon_GetNamePath(obj uintptr) string {
    ret, _, _ := trayIcon_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func TrayIcon_HasParent(obj uintptr) bool {
    ret, _, _ := trayIcon_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func TrayIcon_Assign(obj uintptr, Source uintptr)  {
    trayIcon_Assign.Call(obj, Source )
}

func TrayIcon_ClassName(obj uintptr) string {
    ret, _, _ := trayIcon_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func TrayIcon_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := trayIcon_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func TrayIcon_GetHashCode(obj uintptr) int32 {
    ret, _, _ := trayIcon_GetHashCode.Call(obj)
    return int32(ret)
}

func TrayIcon_ToString(obj uintptr) string {
    ret, _, _ := trayIcon_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func TrayIcon_GetAnimate(obj uintptr) bool {
    ret, _, _ := trayIcon_GetAnimate.Call(obj)
    return DBoolToGoBool(ret)
}

func TrayIcon_SetAnimate(obj uintptr, value bool) {
   trayIcon_SetAnimate.Call(obj, GoBoolToDBool(value))
}

func TrayIcon_GetAnimateInterval(obj uintptr) uint32 {
    ret, _, _ := trayIcon_GetAnimateInterval.Call(obj)
    return uint32(ret)
}

func TrayIcon_SetAnimateInterval(obj uintptr, value uint32) {
   trayIcon_SetAnimateInterval.Call(obj, uintptr(value))
}

func TrayIcon_GetHint(obj uintptr) string {
    ret, _, _ := trayIcon_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func TrayIcon_SetHint(obj uintptr, value string) {
   trayIcon_SetHint.Call(obj, GoStrToDStr(value))
}

func TrayIcon_GetBalloonHint(obj uintptr) string {
    ret, _, _ := trayIcon_GetBalloonHint.Call(obj)
    return DStrToGoStr(ret)
}

func TrayIcon_SetBalloonHint(obj uintptr, value string) {
   trayIcon_SetBalloonHint.Call(obj, GoStrToDStr(value))
}

func TrayIcon_GetBalloonTitle(obj uintptr) string {
    ret, _, _ := trayIcon_GetBalloonTitle.Call(obj)
    return DStrToGoStr(ret)
}

func TrayIcon_SetBalloonTitle(obj uintptr, value string) {
   trayIcon_SetBalloonTitle.Call(obj, GoStrToDStr(value))
}

func TrayIcon_GetBalloonTimeout(obj uintptr) int32 {
    ret, _, _ := trayIcon_GetBalloonTimeout.Call(obj)
    return int32(ret)
}

func TrayIcon_SetBalloonTimeout(obj uintptr, value int32) {
   trayIcon_SetBalloonTimeout.Call(obj, uintptr(value))
}

func TrayIcon_GetBalloonFlags(obj uintptr) TBalloonFlags {
    ret, _, _ := trayIcon_GetBalloonFlags.Call(obj)
    return TBalloonFlags(ret)
}

func TrayIcon_SetBalloonFlags(obj uintptr, value TBalloonFlags) {
   trayIcon_SetBalloonFlags.Call(obj, uintptr(value))
}

func TrayIcon_GetIcon(obj uintptr) uintptr {
    ret, _, _ := trayIcon_GetIcon.Call(obj)
    return ret
}

func TrayIcon_SetIcon(obj uintptr, value uintptr) {
   trayIcon_SetIcon.Call(obj, value)
}

func TrayIcon_GetIconIndex(obj uintptr) int32 {
    ret, _, _ := trayIcon_GetIconIndex.Call(obj)
    return int32(ret)
}

func TrayIcon_SetIconIndex(obj uintptr, value int32) {
   trayIcon_SetIconIndex.Call(obj, uintptr(value))
}

func TrayIcon_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := trayIcon_GetPopupMenu.Call(obj)
    return ret
}

func TrayIcon_SetPopupMenu(obj uintptr, value uintptr) {
   trayIcon_SetPopupMenu.Call(obj, value)
}

func TrayIcon_GetVisible(obj uintptr) bool {
    ret, _, _ := trayIcon_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func TrayIcon_SetVisible(obj uintptr, value bool) {
   trayIcon_SetVisible.Call(obj, GoBoolToDBool(value))
}

func TrayIcon_SetOnBalloonClick(obj uintptr, fn interface{}) {
    trayIcon_SetOnBalloonClick.Call(obj, addEventToMap(fn))
}

func TrayIcon_SetOnClick(obj uintptr, fn interface{}) {
    trayIcon_SetOnClick.Call(obj, addEventToMap(fn))
}

func TrayIcon_SetOnDblClick(obj uintptr, fn interface{}) {
    trayIcon_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func TrayIcon_SetOnMouseMove(obj uintptr, fn interface{}) {
    trayIcon_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func TrayIcon_SetOnMouseUp(obj uintptr, fn interface{}) {
    trayIcon_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func TrayIcon_SetOnMouseDown(obj uintptr, fn interface{}) {
    trayIcon_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func TrayIcon_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := trayIcon_GetComponentCount.Call(obj)
    return int32(ret)
}

func TrayIcon_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := trayIcon_GetComponentIndex.Call(obj)
    return int32(ret)
}

func TrayIcon_SetComponentIndex(obj uintptr, value int32) {
   trayIcon_SetComponentIndex.Call(obj, uintptr(value))
}

func TrayIcon_GetOwner(obj uintptr) uintptr {
    ret, _, _ := trayIcon_GetOwner.Call(obj)
    return ret
}

func TrayIcon_GetName(obj uintptr) string {
    ret, _, _ := trayIcon_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func TrayIcon_SetName(obj uintptr, value string) {
   trayIcon_SetName.Call(obj, GoStrToDStr(value))
}

func TrayIcon_GetTag(obj uintptr) int {
    ret, _, _ := trayIcon_GetTag.Call(obj)
    return int(ret)
}

func TrayIcon_SetTag(obj uintptr, value int) {
   trayIcon_SetTag.Call(obj, uintptr(value))
}

func TrayIcon_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := trayIcon_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TBalloonHint ---------------------------

func BalloonHint_Create(obj uintptr) uintptr {
    ret, _, _ := balloonHint_Create.Call(obj)
    return ret
}

func BalloonHint_Free(obj uintptr) {
    balloonHint_Free.Call(obj)
}

func BalloonHint_ShowHint(obj uintptr)  {
    balloonHint_ShowHint.Call(obj)
}

func BalloonHint_HideHint(obj uintptr)  {
    balloonHint_HideHint.Call(obj)
}

func BalloonHint_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := balloonHint_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func BalloonHint_GetNamePath(obj uintptr) string {
    ret, _, _ := balloonHint_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func BalloonHint_HasParent(obj uintptr) bool {
    ret, _, _ := balloonHint_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func BalloonHint_Assign(obj uintptr, Source uintptr)  {
    balloonHint_Assign.Call(obj, Source )
}

func BalloonHint_ClassName(obj uintptr) string {
    ret, _, _ := balloonHint_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func BalloonHint_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := balloonHint_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func BalloonHint_GetHashCode(obj uintptr) int32 {
    ret, _, _ := balloonHint_GetHashCode.Call(obj)
    return int32(ret)
}

func BalloonHint_ToString(obj uintptr) string {
    ret, _, _ := balloonHint_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func BalloonHint_GetTitle(obj uintptr) string {
    ret, _, _ := balloonHint_GetTitle.Call(obj)
    return DStrToGoStr(ret)
}

func BalloonHint_SetTitle(obj uintptr, value string) {
   balloonHint_SetTitle.Call(obj, GoStrToDStr(value))
}

func BalloonHint_GetDescription(obj uintptr) string {
    ret, _, _ := balloonHint_GetDescription.Call(obj)
    return DStrToGoStr(ret)
}

func BalloonHint_SetDescription(obj uintptr, value string) {
   balloonHint_SetDescription.Call(obj, GoStrToDStr(value))
}

func BalloonHint_GetImageIndex(obj uintptr) int32 {
    ret, _, _ := balloonHint_GetImageIndex.Call(obj)
    return int32(ret)
}

func BalloonHint_SetImageIndex(obj uintptr, value int32) {
   balloonHint_SetImageIndex.Call(obj, uintptr(value))
}

func BalloonHint_GetImages(obj uintptr) uintptr {
    ret, _, _ := balloonHint_GetImages.Call(obj)
    return ret
}

func BalloonHint_SetImages(obj uintptr, value uintptr) {
   balloonHint_SetImages.Call(obj, value)
}

func BalloonHint_GetStyle(obj uintptr) TBalloonHintStyle {
    ret, _, _ := balloonHint_GetStyle.Call(obj)
    return TBalloonHintStyle(ret)
}

func BalloonHint_SetStyle(obj uintptr, value TBalloonHintStyle) {
   balloonHint_SetStyle.Call(obj, uintptr(value))
}

func BalloonHint_GetDelay(obj uintptr) uint32 {
    ret, _, _ := balloonHint_GetDelay.Call(obj)
    return uint32(ret)
}

func BalloonHint_SetDelay(obj uintptr, value uint32) {
   balloonHint_SetDelay.Call(obj, uintptr(value))
}

func BalloonHint_GetHideAfter(obj uintptr) int32 {
    ret, _, _ := balloonHint_GetHideAfter.Call(obj)
    return int32(ret)
}

func BalloonHint_SetHideAfter(obj uintptr, value int32) {
   balloonHint_SetHideAfter.Call(obj, uintptr(value))
}

func BalloonHint_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := balloonHint_GetComponentCount.Call(obj)
    return int32(ret)
}

func BalloonHint_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := balloonHint_GetComponentIndex.Call(obj)
    return int32(ret)
}

func BalloonHint_SetComponentIndex(obj uintptr, value int32) {
   balloonHint_SetComponentIndex.Call(obj, uintptr(value))
}

func BalloonHint_GetOwner(obj uintptr) uintptr {
    ret, _, _ := balloonHint_GetOwner.Call(obj)
    return ret
}

func BalloonHint_GetName(obj uintptr) string {
    ret, _, _ := balloonHint_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func BalloonHint_SetName(obj uintptr, value string) {
   balloonHint_SetName.Call(obj, GoStrToDStr(value))
}

func BalloonHint_GetTag(obj uintptr) int {
    ret, _, _ := balloonHint_GetTag.Call(obj)
    return int(ret)
}

func BalloonHint_SetTag(obj uintptr, value int) {
   balloonHint_SetTag.Call(obj, uintptr(value))
}

func BalloonHint_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := balloonHint_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TCategoryPanelGroup ---------------------------

func CategoryPanelGroup_Create(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_Create.Call(obj)
    return ret
}

func CategoryPanelGroup_Free(obj uintptr) {
    categoryPanelGroup_Free.Call(obj)
}

func CategoryPanelGroup_CollapseAll(obj uintptr)  {
    categoryPanelGroup_CollapseAll.Call(obj)
}

func CategoryPanelGroup_ExpandAll(obj uintptr)  {
    categoryPanelGroup_ExpandAll.Call(obj)
}

func CategoryPanelGroup_CanFocus(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_FlipChildren(obj uintptr, AllLevels bool)  {
    categoryPanelGroup_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func CategoryPanelGroup_Focused(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_HandleAllocated(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_Invalidate(obj uintptr)  {
    categoryPanelGroup_Invalidate.Call(obj)
}

func CategoryPanelGroup_Realign(obj uintptr)  {
    categoryPanelGroup_Realign.Call(obj)
}

func CategoryPanelGroup_Repaint(obj uintptr)  {
    categoryPanelGroup_Repaint.Call(obj)
}

func CategoryPanelGroup_ScaleBy(obj uintptr, M int32, D int32)  {
    categoryPanelGroup_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func CategoryPanelGroup_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    categoryPanelGroup_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func CategoryPanelGroup_SetFocus(obj uintptr)  {
    categoryPanelGroup_SetFocus.Call(obj)
}

func CategoryPanelGroup_Update(obj uintptr)  {
    categoryPanelGroup_Update.Call(obj)
}

func CategoryPanelGroup_BringToFront(obj uintptr)  {
    categoryPanelGroup_BringToFront.Call(obj)
}

func CategoryPanelGroup_HasParent(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_Hide(obj uintptr)  {
    categoryPanelGroup_Hide.Call(obj)
}

func CategoryPanelGroup_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := categoryPanelGroup_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func CategoryPanelGroup_Refresh(obj uintptr)  {
    categoryPanelGroup_Refresh.Call(obj)
}

func CategoryPanelGroup_SendToBack(obj uintptr)  {
    categoryPanelGroup_SendToBack.Call(obj)
}

func CategoryPanelGroup_Show(obj uintptr)  {
    categoryPanelGroup_Show.Call(obj)
}

func CategoryPanelGroup_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := categoryPanelGroup_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func CategoryPanelGroup_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := categoryPanelGroup_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func CategoryPanelGroup_GetNamePath(obj uintptr) string {
    ret, _, _ := categoryPanelGroup_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func CategoryPanelGroup_Assign(obj uintptr, Source uintptr)  {
    categoryPanelGroup_Assign.Call(obj, Source )
}

func CategoryPanelGroup_ClassName(obj uintptr) string {
    ret, _, _ := categoryPanelGroup_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func CategoryPanelGroup_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_GetHashCode(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetHashCode.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_ToString(obj uintptr) string {
    ret, _, _ := categoryPanelGroup_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func CategoryPanelGroup_GetAlign(obj uintptr) TAlign {
    ret, _, _ := categoryPanelGroup_GetAlign.Call(obj)
    return TAlign(ret)
}

func CategoryPanelGroup_SetAlign(obj uintptr, value TAlign) {
   categoryPanelGroup_SetAlign.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := categoryPanelGroup_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func CategoryPanelGroup_SetAnchors(obj uintptr, value TAnchors) {
   categoryPanelGroup_SetAnchors.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetBevelEdges(obj uintptr) TBevelEdges {
    ret, _, _ := categoryPanelGroup_GetBevelEdges.Call(obj)
    return TBevelEdges(ret)
}

func CategoryPanelGroup_SetBevelEdges(obj uintptr, value TBevelEdges) {
   categoryPanelGroup_SetBevelEdges.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetBevelInner(obj uintptr) TBevelCut {
    ret, _, _ := categoryPanelGroup_GetBevelInner.Call(obj)
    return TBevelCut(ret)
}

func CategoryPanelGroup_SetBevelInner(obj uintptr, value TBevelCut) {
   categoryPanelGroup_SetBevelInner.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetBevelOuter(obj uintptr) TBevelCut {
    ret, _, _ := categoryPanelGroup_GetBevelOuter.Call(obj)
    return TBevelCut(ret)
}

func CategoryPanelGroup_SetBevelOuter(obj uintptr, value TBevelCut) {
   categoryPanelGroup_SetBevelOuter.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetBevelKind(obj uintptr) TBevelKind {
    ret, _, _ := categoryPanelGroup_GetBevelKind.Call(obj)
    return TBevelKind(ret)
}

func CategoryPanelGroup_SetBevelKind(obj uintptr, value TBevelKind) {
   categoryPanelGroup_SetBevelKind.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := categoryPanelGroup_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func CategoryPanelGroup_SetBiDiMode(obj uintptr, value TBiDiMode) {
   categoryPanelGroup_SetBiDiMode.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetChevronAlignment(obj uintptr) TAlignment {
    ret, _, _ := categoryPanelGroup_GetChevronAlignment.Call(obj)
    return TAlignment(ret)
}

func CategoryPanelGroup_SetChevronAlignment(obj uintptr, value TAlignment) {
   categoryPanelGroup_SetChevronAlignment.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetChevronColor(obj uintptr) TColor {
    ret, _, _ := categoryPanelGroup_GetChevronColor.Call(obj)
    return TColor(ret)
}

func CategoryPanelGroup_SetChevronColor(obj uintptr, value TColor) {
   categoryPanelGroup_SetChevronColor.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetChevronHotColor(obj uintptr) TColor {
    ret, _, _ := categoryPanelGroup_GetChevronHotColor.Call(obj)
    return TColor(ret)
}

func CategoryPanelGroup_SetChevronHotColor(obj uintptr, value TColor) {
   categoryPanelGroup_SetChevronHotColor.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetDoubleBuffered(obj uintptr, value bool) {
   categoryPanelGroup_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetEnabled(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetEnabled(obj uintptr, value bool) {
   categoryPanelGroup_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetColor(obj uintptr) TColor {
    ret, _, _ := categoryPanelGroup_GetColor.Call(obj)
    return TColor(ret)
}

func CategoryPanelGroup_SetColor(obj uintptr, value TColor) {
   categoryPanelGroup_SetColor.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetFont(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetFont.Call(obj)
    return ret
}

func CategoryPanelGroup_SetFont(obj uintptr, value uintptr) {
   categoryPanelGroup_SetFont.Call(obj, value)
}

func CategoryPanelGroup_GetGradientBaseColor(obj uintptr) TColor {
    ret, _, _ := categoryPanelGroup_GetGradientBaseColor.Call(obj)
    return TColor(ret)
}

func CategoryPanelGroup_SetGradientBaseColor(obj uintptr, value TColor) {
   categoryPanelGroup_SetGradientBaseColor.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetGradientColor(obj uintptr) TColor {
    ret, _, _ := categoryPanelGroup_GetGradientColor.Call(obj)
    return TColor(ret)
}

func CategoryPanelGroup_SetGradientColor(obj uintptr, value TColor) {
   categoryPanelGroup_SetGradientColor.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetGradientDirection(obj uintptr) TGradientDirection {
    ret, _, _ := categoryPanelGroup_GetGradientDirection.Call(obj)
    return TGradientDirection(ret)
}

func CategoryPanelGroup_SetGradientDirection(obj uintptr, value TGradientDirection) {
   categoryPanelGroup_SetGradientDirection.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetHeaderAlignment(obj uintptr) TAlignment {
    ret, _, _ := categoryPanelGroup_GetHeaderAlignment.Call(obj)
    return TAlignment(ret)
}

func CategoryPanelGroup_SetHeaderAlignment(obj uintptr, value TAlignment) {
   categoryPanelGroup_SetHeaderAlignment.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetHeaderFont(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetHeaderFont.Call(obj)
    return ret
}

func CategoryPanelGroup_SetHeaderFont(obj uintptr, value uintptr) {
   categoryPanelGroup_SetHeaderFont.Call(obj, value)
}

func CategoryPanelGroup_GetHeaderHeight(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetHeaderHeight.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_SetHeaderHeight(obj uintptr, value int32) {
   categoryPanelGroup_SetHeaderHeight.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetHeaderImage(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetHeaderImage.Call(obj)
    return ret
}

func CategoryPanelGroup_SetHeaderImage(obj uintptr, value uintptr) {
   categoryPanelGroup_SetHeaderImage.Call(obj, value)
}

func CategoryPanelGroup_GetHeaderStyle(obj uintptr) THeaderStyle {
    ret, _, _ := categoryPanelGroup_GetHeaderStyle.Call(obj)
    return THeaderStyle(ret)
}

func CategoryPanelGroup_SetHeaderStyle(obj uintptr, value THeaderStyle) {
   categoryPanelGroup_SetHeaderStyle.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetHeight(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetHeight.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_SetHeight(obj uintptr, value int32) {
   categoryPanelGroup_SetHeight.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetImages(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetImages.Call(obj)
    return ret
}

func CategoryPanelGroup_SetImages(obj uintptr, value uintptr) {
   categoryPanelGroup_SetImages.Call(obj, value)
}

func CategoryPanelGroup_GetParentBackground(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetParentBackground.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetParentBackground(obj uintptr, value bool) {
   categoryPanelGroup_SetParentBackground.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetParentColor(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetParentColor(obj uintptr, value bool) {
   categoryPanelGroup_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetParentCtl3D(obj uintptr, value bool) {
   categoryPanelGroup_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetParentDoubleBuffered(obj uintptr, value bool) {
   categoryPanelGroup_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetParentFont(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetParentFont(obj uintptr, value bool) {
   categoryPanelGroup_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetParentShowHint(obj uintptr, value bool) {
   categoryPanelGroup_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetPopupMenu.Call(obj)
    return ret
}

func CategoryPanelGroup_SetPopupMenu(obj uintptr, value uintptr) {
   categoryPanelGroup_SetPopupMenu.Call(obj, value)
}

func CategoryPanelGroup_GetShowHint(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetShowHint(obj uintptr, value bool) {
   categoryPanelGroup_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := categoryPanelGroup_GetTabOrder.Call(obj)
    return uint16(ret)
}

func CategoryPanelGroup_SetTabOrder(obj uintptr, value uint16) {
   categoryPanelGroup_SetTabOrder.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetTabStop(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetTabStop(obj uintptr, value bool) {
   categoryPanelGroup_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetVisible(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetVisible(obj uintptr, value bool) {
   categoryPanelGroup_SetVisible.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := categoryPanelGroup_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func CategoryPanelGroup_SetStyleElements(obj uintptr, value TStyleElements) {
   categoryPanelGroup_SetStyleElements.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetWidth(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetWidth.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_SetWidth(obj uintptr, value int32) {
   categoryPanelGroup_SetWidth.Call(obj, uintptr(value))
}

func CategoryPanelGroup_SetOnClick(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnClick.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnDblClick(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnEnter(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnEnter.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnExit(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnExit.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnMouseDown(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnMouseEnter(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnMouseLeave(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnMouseMove(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnMouseUp(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnMouseWheel(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnMouseWheel.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_SetOnResize(obj uintptr, fn interface{}) {
    categoryPanelGroup_SetOnResize.Call(obj, addEventToMap(fn))
}

func CategoryPanelGroup_GetPanels(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetPanels.Call(obj)
    return ret
}

func CategoryPanelGroup_GetBrush(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetBrush.Call(obj)
    return ret
}

func CategoryPanelGroup_GetControlCount(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetControlCount.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_GetHandle(obj uintptr) HWND {
    ret, _, _ := categoryPanelGroup_GetHandle.Call(obj)
    return HWND(ret)
}

func CategoryPanelGroup_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := categoryPanelGroup_GetParentWindow.Call(obj)
    return HWND(ret)
}

func CategoryPanelGroup_SetParentWindow(obj uintptr, value HWND) {
   categoryPanelGroup_SetParentWindow.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetAction(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetAction.Call(obj)
    return ret
}

func CategoryPanelGroup_SetAction(obj uintptr, value uintptr) {
   categoryPanelGroup_SetAction.Call(obj, value)
}

func CategoryPanelGroup_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    categoryPanelGroup_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func CategoryPanelGroup_SetBoundsRect(obj uintptr, value TRect) {
   categoryPanelGroup_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func CategoryPanelGroup_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetClientHeight.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_SetClientHeight(obj uintptr, value int32) {
   categoryPanelGroup_SetClientHeight.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetClientRect(obj uintptr) TRect {
    var ret TRect
    categoryPanelGroup_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func CategoryPanelGroup_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetClientWidth.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_SetClientWidth(obj uintptr, value int32) {
   categoryPanelGroup_SetClientWidth.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetExplicitTop.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_GetParent(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetParent.Call(obj)
    return ret
}

func CategoryPanelGroup_SetParent(obj uintptr, value uintptr) {
   categoryPanelGroup_SetParent.Call(obj, value)
}

func CategoryPanelGroup_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := categoryPanelGroup_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanelGroup_SetAlignWithMargins(obj uintptr, value bool) {
   categoryPanelGroup_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func CategoryPanelGroup_GetLeft(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetLeft.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_SetLeft(obj uintptr, value int32) {
   categoryPanelGroup_SetLeft.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetTop(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetTop.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_SetTop(obj uintptr, value int32) {
   categoryPanelGroup_SetTop.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetCursor(obj uintptr) TCursor {
    ret, _, _ := categoryPanelGroup_GetCursor.Call(obj)
    return TCursor(ret)
}

func CategoryPanelGroup_SetCursor(obj uintptr, value TCursor) {
   categoryPanelGroup_SetCursor.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetHint(obj uintptr) string {
    ret, _, _ := categoryPanelGroup_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func CategoryPanelGroup_SetHint(obj uintptr, value string) {
   categoryPanelGroup_SetHint.Call(obj, GoStrToDStr(value))
}

func CategoryPanelGroup_GetMargins(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetMargins.Call(obj)
    return ret
}

func CategoryPanelGroup_SetMargins(obj uintptr, value uintptr) {
   categoryPanelGroup_SetMargins.Call(obj, value)
}

func CategoryPanelGroup_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetComponentCount.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := categoryPanelGroup_GetComponentIndex.Call(obj)
    return int32(ret)
}

func CategoryPanelGroup_SetComponentIndex(obj uintptr, value int32) {
   categoryPanelGroup_SetComponentIndex.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetOwner(obj uintptr) uintptr {
    ret, _, _ := categoryPanelGroup_GetOwner.Call(obj)
    return ret
}

func CategoryPanelGroup_GetName(obj uintptr) string {
    ret, _, _ := categoryPanelGroup_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func CategoryPanelGroup_SetName(obj uintptr, value string) {
   categoryPanelGroup_SetName.Call(obj, GoStrToDStr(value))
}

func CategoryPanelGroup_GetTag(obj uintptr) int {
    ret, _, _ := categoryPanelGroup_GetTag.Call(obj)
    return int(ret)
}

func CategoryPanelGroup_SetTag(obj uintptr, value int) {
   categoryPanelGroup_SetTag.Call(obj, uintptr(value))
}

func CategoryPanelGroup_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := categoryPanelGroup_GetControls.Call(obj, uintptr(Index))
    return ret
}

func CategoryPanelGroup_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := categoryPanelGroup_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TCategoryPanel ---------------------------

func CategoryPanel_Create(obj uintptr) uintptr {
    ret, _, _ := categoryPanel_Create.Call(obj)
    return ret
}

func CategoryPanel_Free(obj uintptr) {
    categoryPanel_Free.Call(obj)
}

func CategoryPanel_Collapse(obj uintptr)  {
    categoryPanel_Collapse.Call(obj)
}

func CategoryPanel_Expand(obj uintptr)  {
    categoryPanel_Expand.Call(obj)
}

func CategoryPanel_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    categoryPanel_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func CategoryPanel_CanFocus(obj uintptr) bool {
    ret, _, _ := categoryPanel_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_FlipChildren(obj uintptr, AllLevels bool)  {
    categoryPanel_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func CategoryPanel_Focused(obj uintptr) bool {
    ret, _, _ := categoryPanel_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_HandleAllocated(obj uintptr) bool {
    ret, _, _ := categoryPanel_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_Invalidate(obj uintptr)  {
    categoryPanel_Invalidate.Call(obj)
}

func CategoryPanel_Realign(obj uintptr)  {
    categoryPanel_Realign.Call(obj)
}

func CategoryPanel_Repaint(obj uintptr)  {
    categoryPanel_Repaint.Call(obj)
}

func CategoryPanel_ScaleBy(obj uintptr, M int32, D int32)  {
    categoryPanel_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func CategoryPanel_SetFocus(obj uintptr)  {
    categoryPanel_SetFocus.Call(obj)
}

func CategoryPanel_Update(obj uintptr)  {
    categoryPanel_Update.Call(obj)
}

func CategoryPanel_BringToFront(obj uintptr)  {
    categoryPanel_BringToFront.Call(obj)
}

func CategoryPanel_HasParent(obj uintptr) bool {
    ret, _, _ := categoryPanel_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_Hide(obj uintptr)  {
    categoryPanel_Hide.Call(obj)
}

func CategoryPanel_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := categoryPanel_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func CategoryPanel_Refresh(obj uintptr)  {
    categoryPanel_Refresh.Call(obj)
}

func CategoryPanel_SendToBack(obj uintptr)  {
    categoryPanel_SendToBack.Call(obj)
}

func CategoryPanel_Show(obj uintptr)  {
    categoryPanel_Show.Call(obj)
}

func CategoryPanel_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := categoryPanel_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func CategoryPanel_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := categoryPanel_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func CategoryPanel_GetNamePath(obj uintptr) string {
    ret, _, _ := categoryPanel_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func CategoryPanel_Assign(obj uintptr, Source uintptr)  {
    categoryPanel_Assign.Call(obj, Source )
}

func CategoryPanel_ClassName(obj uintptr) string {
    ret, _, _ := categoryPanel_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func CategoryPanel_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := categoryPanel_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func CategoryPanel_GetHashCode(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetHashCode.Call(obj)
    return int32(ret)
}

func CategoryPanel_ToString(obj uintptr) string {
    ret, _, _ := categoryPanel_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func CategoryPanel_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := categoryPanel_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func CategoryPanel_SetBiDiMode(obj uintptr, value TBiDiMode) {
   categoryPanel_SetBiDiMode.Call(obj, uintptr(value))
}

func CategoryPanel_GetCaption(obj uintptr) string {
    ret, _, _ := categoryPanel_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func CategoryPanel_SetCaption(obj uintptr, value string) {
   categoryPanel_SetCaption.Call(obj, GoStrToDStr(value))
}

func CategoryPanel_GetColor(obj uintptr) TColor {
    ret, _, _ := categoryPanel_GetColor.Call(obj)
    return TColor(ret)
}

func CategoryPanel_SetColor(obj uintptr, value TColor) {
   categoryPanel_SetColor.Call(obj, uintptr(value))
}

func CategoryPanel_GetCollapsed(obj uintptr) bool {
    ret, _, _ := categoryPanel_GetCollapsed.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_SetCollapsed(obj uintptr, value bool) {
   categoryPanel_SetCollapsed.Call(obj, GoBoolToDBool(value))
}

func CategoryPanel_GetCollapsedHotImageIndex(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetCollapsedHotImageIndex.Call(obj)
    return int32(ret)
}

func CategoryPanel_SetCollapsedHotImageIndex(obj uintptr, value int32) {
   categoryPanel_SetCollapsedHotImageIndex.Call(obj, uintptr(value))
}

func CategoryPanel_GetCollapsedImageIndex(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetCollapsedImageIndex.Call(obj)
    return int32(ret)
}

func CategoryPanel_SetCollapsedImageIndex(obj uintptr, value int32) {
   categoryPanel_SetCollapsedImageIndex.Call(obj, uintptr(value))
}

func CategoryPanel_GetCollapsedPressedImageIndex(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetCollapsedPressedImageIndex.Call(obj)
    return int32(ret)
}

func CategoryPanel_SetCollapsedPressedImageIndex(obj uintptr, value int32) {
   categoryPanel_SetCollapsedPressedImageIndex.Call(obj, uintptr(value))
}

func CategoryPanel_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := categoryPanel_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_SetDoubleBuffered(obj uintptr, value bool) {
   categoryPanel_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func CategoryPanel_GetEnabled(obj uintptr) bool {
    ret, _, _ := categoryPanel_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_SetEnabled(obj uintptr, value bool) {
   categoryPanel_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func CategoryPanel_GetExpandedHotImageIndex(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetExpandedHotImageIndex.Call(obj)
    return int32(ret)
}

func CategoryPanel_SetExpandedHotImageIndex(obj uintptr, value int32) {
   categoryPanel_SetExpandedHotImageIndex.Call(obj, uintptr(value))
}

func CategoryPanel_GetExpandedImageIndex(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetExpandedImageIndex.Call(obj)
    return int32(ret)
}

func CategoryPanel_SetExpandedImageIndex(obj uintptr, value int32) {
   categoryPanel_SetExpandedImageIndex.Call(obj, uintptr(value))
}

func CategoryPanel_GetExpandedPressedImageIndex(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetExpandedPressedImageIndex.Call(obj)
    return int32(ret)
}

func CategoryPanel_SetExpandedPressedImageIndex(obj uintptr, value int32) {
   categoryPanel_SetExpandedPressedImageIndex.Call(obj, uintptr(value))
}

func CategoryPanel_GetFullRepaint(obj uintptr) bool {
    ret, _, _ := categoryPanel_GetFullRepaint.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_SetFullRepaint(obj uintptr, value bool) {
   categoryPanel_SetFullRepaint.Call(obj, GoBoolToDBool(value))
}

func CategoryPanel_GetFont(obj uintptr) uintptr {
    ret, _, _ := categoryPanel_GetFont.Call(obj)
    return ret
}

func CategoryPanel_SetFont(obj uintptr, value uintptr) {
   categoryPanel_SetFont.Call(obj, value)
}

func CategoryPanel_GetHeight(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetHeight.Call(obj)
    return int32(ret)
}

func CategoryPanel_SetHeight(obj uintptr, value int32) {
   categoryPanel_SetHeight.Call(obj, uintptr(value))
}

func CategoryPanel_GetLeft(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetLeft.Call(obj)
    return int32(ret)
}

func CategoryPanel_SetLeft(obj uintptr, value int32) {
   categoryPanel_SetLeft.Call(obj, uintptr(value))
}

func CategoryPanel_GetLocked(obj uintptr) bool {
    ret, _, _ := categoryPanel_GetLocked.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_SetLocked(obj uintptr, value bool) {
   categoryPanel_SetLocked.Call(obj, GoBoolToDBool(value))
}

func CategoryPanel_GetParentBackground(obj uintptr) bool {
    ret, _, _ := categoryPanel_GetParentBackground.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_SetParentBackground(obj uintptr, value bool) {
   categoryPanel_SetParentBackground.Call(obj, GoBoolToDBool(value))
}

func CategoryPanel_GetParentColor(obj uintptr) bool {
    ret, _, _ := categoryPanel_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_SetParentColor(obj uintptr, value bool) {
   categoryPanel_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func CategoryPanel_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := categoryPanel_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_SetParentCtl3D(obj uintptr, value bool) {
   categoryPanel_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func CategoryPanel_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := categoryPanel_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_SetParentDoubleBuffered(obj uintptr, value bool) {
   categoryPanel_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func CategoryPanel_GetParentFont(obj uintptr) bool {
    ret, _, _ := categoryPanel_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_SetParentFont(obj uintptr, value bool) {
   categoryPanel_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func CategoryPanel_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := categoryPanel_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_SetParentShowHint(obj uintptr, value bool) {
   categoryPanel_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func CategoryPanel_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := categoryPanel_GetPopupMenu.Call(obj)
    return ret
}

func CategoryPanel_SetPopupMenu(obj uintptr, value uintptr) {
   categoryPanel_SetPopupMenu.Call(obj, value)
}

func CategoryPanel_GetShowHint(obj uintptr) bool {
    ret, _, _ := categoryPanel_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_SetShowHint(obj uintptr, value bool) {
   categoryPanel_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func CategoryPanel_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := categoryPanel_GetTabOrder.Call(obj)
    return uint16(ret)
}

func CategoryPanel_SetTabOrder(obj uintptr, value uint16) {
   categoryPanel_SetTabOrder.Call(obj, uintptr(value))
}

func CategoryPanel_GetTabStop(obj uintptr) bool {
    ret, _, _ := categoryPanel_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_SetTabStop(obj uintptr, value bool) {
   categoryPanel_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func CategoryPanel_GetVisible(obj uintptr) bool {
    ret, _, _ := categoryPanel_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_SetVisible(obj uintptr, value bool) {
   categoryPanel_SetVisible.Call(obj, GoBoolToDBool(value))
}

func CategoryPanel_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := categoryPanel_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func CategoryPanel_SetStyleElements(obj uintptr, value TStyleElements) {
   categoryPanel_SetStyleElements.Call(obj, uintptr(value))
}

func CategoryPanel_GetWidth(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetWidth.Call(obj)
    return int32(ret)
}

func CategoryPanel_SetWidth(obj uintptr, value int32) {
   categoryPanel_SetWidth.Call(obj, uintptr(value))
}

func CategoryPanel_SetOnClick(obj uintptr, fn interface{}) {
    categoryPanel_SetOnClick.Call(obj, addEventToMap(fn))
}

func CategoryPanel_SetOnDblClick(obj uintptr, fn interface{}) {
    categoryPanel_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func CategoryPanel_SetOnEnter(obj uintptr, fn interface{}) {
    categoryPanel_SetOnEnter.Call(obj, addEventToMap(fn))
}

func CategoryPanel_SetOnExit(obj uintptr, fn interface{}) {
    categoryPanel_SetOnExit.Call(obj, addEventToMap(fn))
}

func CategoryPanel_SetOnMouseDown(obj uintptr, fn interface{}) {
    categoryPanel_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func CategoryPanel_SetOnMouseEnter(obj uintptr, fn interface{}) {
    categoryPanel_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func CategoryPanel_SetOnMouseLeave(obj uintptr, fn interface{}) {
    categoryPanel_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func CategoryPanel_SetOnMouseMove(obj uintptr, fn interface{}) {
    categoryPanel_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func CategoryPanel_SetOnMouseUp(obj uintptr, fn interface{}) {
    categoryPanel_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func CategoryPanel_GetPanelGroup(obj uintptr) uintptr {
    ret, _, _ := categoryPanel_GetPanelGroup.Call(obj)
    return ret
}

func CategoryPanel_SetPanelGroup(obj uintptr, value uintptr) {
   categoryPanel_SetPanelGroup.Call(obj, value)
}

func CategoryPanel_GetBrush(obj uintptr) uintptr {
    ret, _, _ := categoryPanel_GetBrush.Call(obj)
    return ret
}

func CategoryPanel_GetControlCount(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetControlCount.Call(obj)
    return int32(ret)
}

func CategoryPanel_GetHandle(obj uintptr) HWND {
    ret, _, _ := categoryPanel_GetHandle.Call(obj)
    return HWND(ret)
}

func CategoryPanel_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := categoryPanel_GetParentWindow.Call(obj)
    return HWND(ret)
}

func CategoryPanel_SetParentWindow(obj uintptr, value HWND) {
   categoryPanel_SetParentWindow.Call(obj, uintptr(value))
}

func CategoryPanel_GetAction(obj uintptr) uintptr {
    ret, _, _ := categoryPanel_GetAction.Call(obj)
    return ret
}

func CategoryPanel_SetAction(obj uintptr, value uintptr) {
   categoryPanel_SetAction.Call(obj, value)
}

func CategoryPanel_GetAlign(obj uintptr) TAlign {
    ret, _, _ := categoryPanel_GetAlign.Call(obj)
    return TAlign(ret)
}

func CategoryPanel_SetAlign(obj uintptr, value TAlign) {
   categoryPanel_SetAlign.Call(obj, uintptr(value))
}

func CategoryPanel_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := categoryPanel_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func CategoryPanel_SetAnchors(obj uintptr, value TAnchors) {
   categoryPanel_SetAnchors.Call(obj, uintptr(value))
}

func CategoryPanel_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    categoryPanel_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func CategoryPanel_SetBoundsRect(obj uintptr, value TRect) {
   categoryPanel_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func CategoryPanel_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetClientHeight.Call(obj)
    return int32(ret)
}

func CategoryPanel_SetClientHeight(obj uintptr, value int32) {
   categoryPanel_SetClientHeight.Call(obj, uintptr(value))
}

func CategoryPanel_GetClientRect(obj uintptr) TRect {
    var ret TRect
    categoryPanel_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func CategoryPanel_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetClientWidth.Call(obj)
    return int32(ret)
}

func CategoryPanel_SetClientWidth(obj uintptr, value int32) {
   categoryPanel_SetClientWidth.Call(obj, uintptr(value))
}

func CategoryPanel_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func CategoryPanel_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetExplicitTop.Call(obj)
    return int32(ret)
}

func CategoryPanel_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func CategoryPanel_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func CategoryPanel_GetParent(obj uintptr) uintptr {
    ret, _, _ := categoryPanel_GetParent.Call(obj)
    return ret
}

func CategoryPanel_SetParent(obj uintptr, value uintptr) {
   categoryPanel_SetParent.Call(obj, value)
}

func CategoryPanel_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := categoryPanel_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func CategoryPanel_SetAlignWithMargins(obj uintptr, value bool) {
   categoryPanel_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func CategoryPanel_GetTop(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetTop.Call(obj)
    return int32(ret)
}

func CategoryPanel_SetTop(obj uintptr, value int32) {
   categoryPanel_SetTop.Call(obj, uintptr(value))
}

func CategoryPanel_GetCursor(obj uintptr) TCursor {
    ret, _, _ := categoryPanel_GetCursor.Call(obj)
    return TCursor(ret)
}

func CategoryPanel_SetCursor(obj uintptr, value TCursor) {
   categoryPanel_SetCursor.Call(obj, uintptr(value))
}

func CategoryPanel_GetHint(obj uintptr) string {
    ret, _, _ := categoryPanel_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func CategoryPanel_SetHint(obj uintptr, value string) {
   categoryPanel_SetHint.Call(obj, GoStrToDStr(value))
}

func CategoryPanel_GetMargins(obj uintptr) uintptr {
    ret, _, _ := categoryPanel_GetMargins.Call(obj)
    return ret
}

func CategoryPanel_SetMargins(obj uintptr, value uintptr) {
   categoryPanel_SetMargins.Call(obj, value)
}

func CategoryPanel_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetComponentCount.Call(obj)
    return int32(ret)
}

func CategoryPanel_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := categoryPanel_GetComponentIndex.Call(obj)
    return int32(ret)
}

func CategoryPanel_SetComponentIndex(obj uintptr, value int32) {
   categoryPanel_SetComponentIndex.Call(obj, uintptr(value))
}

func CategoryPanel_GetOwner(obj uintptr) uintptr {
    ret, _, _ := categoryPanel_GetOwner.Call(obj)
    return ret
}

func CategoryPanel_GetName(obj uintptr) string {
    ret, _, _ := categoryPanel_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func CategoryPanel_SetName(obj uintptr, value string) {
   categoryPanel_SetName.Call(obj, GoStrToDStr(value))
}

func CategoryPanel_GetTag(obj uintptr) int {
    ret, _, _ := categoryPanel_GetTag.Call(obj)
    return int(ret)
}

func CategoryPanel_SetTag(obj uintptr, value int) {
   categoryPanel_SetTag.Call(obj, uintptr(value))
}

func CategoryPanel_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := categoryPanel_GetControls.Call(obj, uintptr(Index))
    return ret
}

func CategoryPanel_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := categoryPanel_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TOpenDialog ---------------------------

func OpenDialog_Create(obj uintptr) uintptr {
    ret, _, _ := openDialog_Create.Call(obj)
    return ret
}

func OpenDialog_Free(obj uintptr) {
    openDialog_Free.Call(obj)
}

func OpenDialog_Execute(obj uintptr, ParentWnd HWND) bool {
    ret, _, _ := openDialog_Execute.Call(obj, uintptr(ParentWnd) )
    return DBoolToGoBool(ret)
}

func OpenDialog_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := openDialog_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func OpenDialog_GetNamePath(obj uintptr) string {
    ret, _, _ := openDialog_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func OpenDialog_HasParent(obj uintptr) bool {
    ret, _, _ := openDialog_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func OpenDialog_Assign(obj uintptr, Source uintptr)  {
    openDialog_Assign.Call(obj, Source )
}

func OpenDialog_ClassName(obj uintptr) string {
    ret, _, _ := openDialog_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func OpenDialog_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := openDialog_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func OpenDialog_GetHashCode(obj uintptr) int32 {
    ret, _, _ := openDialog_GetHashCode.Call(obj)
    return int32(ret)
}

func OpenDialog_ToString(obj uintptr) string {
    ret, _, _ := openDialog_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func OpenDialog_GetFiles(obj uintptr) uintptr {
    ret, _, _ := openDialog_GetFiles.Call(obj)
    return ret
}

func OpenDialog_GetDefaultExt(obj uintptr) string {
    ret, _, _ := openDialog_GetDefaultExt.Call(obj)
    return DStrToGoStr(ret)
}

func OpenDialog_SetDefaultExt(obj uintptr, value string) {
   openDialog_SetDefaultExt.Call(obj, GoStrToDStr(value))
}

func OpenDialog_GetFileName(obj uintptr) string {
    ret, _, _ := openDialog_GetFileName.Call(obj)
    return DStrToGoStr(ret)
}

func OpenDialog_SetFileName(obj uintptr, value string) {
   openDialog_SetFileName.Call(obj, GoStrToDStr(value))
}

func OpenDialog_GetFilter(obj uintptr) string {
    ret, _, _ := openDialog_GetFilter.Call(obj)
    return DStrToGoStr(ret)
}

func OpenDialog_SetFilter(obj uintptr, value string) {
   openDialog_SetFilter.Call(obj, GoStrToDStr(value))
}

func OpenDialog_GetFilterIndex(obj uintptr) int32 {
    ret, _, _ := openDialog_GetFilterIndex.Call(obj)
    return int32(ret)
}

func OpenDialog_SetFilterIndex(obj uintptr, value int32) {
   openDialog_SetFilterIndex.Call(obj, uintptr(value))
}

func OpenDialog_GetInitialDir(obj uintptr) string {
    ret, _, _ := openDialog_GetInitialDir.Call(obj)
    return DStrToGoStr(ret)
}

func OpenDialog_SetInitialDir(obj uintptr, value string) {
   openDialog_SetInitialDir.Call(obj, GoStrToDStr(value))
}

func OpenDialog_GetOptions(obj uintptr) TOpenOptions {
    ret, _, _ := openDialog_GetOptions.Call(obj)
    return TOpenOptions(ret)
}

func OpenDialog_SetOptions(obj uintptr, value TOpenOptions) {
   openDialog_SetOptions.Call(obj, uintptr(value))
}

func OpenDialog_GetOptionsEx(obj uintptr) TOpenOptionsEx {
    ret, _, _ := openDialog_GetOptionsEx.Call(obj)
    return TOpenOptionsEx(ret)
}

func OpenDialog_SetOptionsEx(obj uintptr, value TOpenOptionsEx) {
   openDialog_SetOptionsEx.Call(obj, uintptr(value))
}

func OpenDialog_GetTitle(obj uintptr) string {
    ret, _, _ := openDialog_GetTitle.Call(obj)
    return DStrToGoStr(ret)
}

func OpenDialog_SetTitle(obj uintptr, value string) {
   openDialog_SetTitle.Call(obj, GoStrToDStr(value))
}

func OpenDialog_GetHandle(obj uintptr) HWND {
    ret, _, _ := openDialog_GetHandle.Call(obj)
    return HWND(ret)
}

func OpenDialog_SetOnClose(obj uintptr, fn interface{}) {
    openDialog_SetOnClose.Call(obj, addEventToMap(fn))
}

func OpenDialog_SetOnShow(obj uintptr, fn interface{}) {
    openDialog_SetOnShow.Call(obj, addEventToMap(fn))
}

func OpenDialog_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := openDialog_GetComponentCount.Call(obj)
    return int32(ret)
}

func OpenDialog_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := openDialog_GetComponentIndex.Call(obj)
    return int32(ret)
}

func OpenDialog_SetComponentIndex(obj uintptr, value int32) {
   openDialog_SetComponentIndex.Call(obj, uintptr(value))
}

func OpenDialog_GetOwner(obj uintptr) uintptr {
    ret, _, _ := openDialog_GetOwner.Call(obj)
    return ret
}

func OpenDialog_GetName(obj uintptr) string {
    ret, _, _ := openDialog_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func OpenDialog_SetName(obj uintptr, value string) {
   openDialog_SetName.Call(obj, GoStrToDStr(value))
}

func OpenDialog_GetTag(obj uintptr) int {
    ret, _, _ := openDialog_GetTag.Call(obj)
    return int(ret)
}

func OpenDialog_SetTag(obj uintptr, value int) {
   openDialog_SetTag.Call(obj, uintptr(value))
}

func OpenDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := openDialog_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TSaveDialog ---------------------------

func SaveDialog_Create(obj uintptr) uintptr {
    ret, _, _ := saveDialog_Create.Call(obj)
    return ret
}

func SaveDialog_Free(obj uintptr) {
    saveDialog_Free.Call(obj)
}

func SaveDialog_Execute(obj uintptr, ParentWnd HWND) bool {
    ret, _, _ := saveDialog_Execute.Call(obj, uintptr(ParentWnd) )
    return DBoolToGoBool(ret)
}

func SaveDialog_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := saveDialog_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func SaveDialog_GetNamePath(obj uintptr) string {
    ret, _, _ := saveDialog_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func SaveDialog_HasParent(obj uintptr) bool {
    ret, _, _ := saveDialog_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func SaveDialog_Assign(obj uintptr, Source uintptr)  {
    saveDialog_Assign.Call(obj, Source )
}

func SaveDialog_ClassName(obj uintptr) string {
    ret, _, _ := saveDialog_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func SaveDialog_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := saveDialog_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func SaveDialog_GetHashCode(obj uintptr) int32 {
    ret, _, _ := saveDialog_GetHashCode.Call(obj)
    return int32(ret)
}

func SaveDialog_ToString(obj uintptr) string {
    ret, _, _ := saveDialog_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func SaveDialog_GetFiles(obj uintptr) uintptr {
    ret, _, _ := saveDialog_GetFiles.Call(obj)
    return ret
}

func SaveDialog_GetDefaultExt(obj uintptr) string {
    ret, _, _ := saveDialog_GetDefaultExt.Call(obj)
    return DStrToGoStr(ret)
}

func SaveDialog_SetDefaultExt(obj uintptr, value string) {
   saveDialog_SetDefaultExt.Call(obj, GoStrToDStr(value))
}

func SaveDialog_GetFileName(obj uintptr) string {
    ret, _, _ := saveDialog_GetFileName.Call(obj)
    return DStrToGoStr(ret)
}

func SaveDialog_SetFileName(obj uintptr, value string) {
   saveDialog_SetFileName.Call(obj, GoStrToDStr(value))
}

func SaveDialog_GetFilter(obj uintptr) string {
    ret, _, _ := saveDialog_GetFilter.Call(obj)
    return DStrToGoStr(ret)
}

func SaveDialog_SetFilter(obj uintptr, value string) {
   saveDialog_SetFilter.Call(obj, GoStrToDStr(value))
}

func SaveDialog_GetFilterIndex(obj uintptr) int32 {
    ret, _, _ := saveDialog_GetFilterIndex.Call(obj)
    return int32(ret)
}

func SaveDialog_SetFilterIndex(obj uintptr, value int32) {
   saveDialog_SetFilterIndex.Call(obj, uintptr(value))
}

func SaveDialog_GetInitialDir(obj uintptr) string {
    ret, _, _ := saveDialog_GetInitialDir.Call(obj)
    return DStrToGoStr(ret)
}

func SaveDialog_SetInitialDir(obj uintptr, value string) {
   saveDialog_SetInitialDir.Call(obj, GoStrToDStr(value))
}

func SaveDialog_GetOptions(obj uintptr) TOpenOptions {
    ret, _, _ := saveDialog_GetOptions.Call(obj)
    return TOpenOptions(ret)
}

func SaveDialog_SetOptions(obj uintptr, value TOpenOptions) {
   saveDialog_SetOptions.Call(obj, uintptr(value))
}

func SaveDialog_GetOptionsEx(obj uintptr) TOpenOptionsEx {
    ret, _, _ := saveDialog_GetOptionsEx.Call(obj)
    return TOpenOptionsEx(ret)
}

func SaveDialog_SetOptionsEx(obj uintptr, value TOpenOptionsEx) {
   saveDialog_SetOptionsEx.Call(obj, uintptr(value))
}

func SaveDialog_GetTitle(obj uintptr) string {
    ret, _, _ := saveDialog_GetTitle.Call(obj)
    return DStrToGoStr(ret)
}

func SaveDialog_SetTitle(obj uintptr, value string) {
   saveDialog_SetTitle.Call(obj, GoStrToDStr(value))
}

func SaveDialog_GetHandle(obj uintptr) HWND {
    ret, _, _ := saveDialog_GetHandle.Call(obj)
    return HWND(ret)
}

func SaveDialog_SetOnClose(obj uintptr, fn interface{}) {
    saveDialog_SetOnClose.Call(obj, addEventToMap(fn))
}

func SaveDialog_SetOnShow(obj uintptr, fn interface{}) {
    saveDialog_SetOnShow.Call(obj, addEventToMap(fn))
}

func SaveDialog_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := saveDialog_GetComponentCount.Call(obj)
    return int32(ret)
}

func SaveDialog_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := saveDialog_GetComponentIndex.Call(obj)
    return int32(ret)
}

func SaveDialog_SetComponentIndex(obj uintptr, value int32) {
   saveDialog_SetComponentIndex.Call(obj, uintptr(value))
}

func SaveDialog_GetOwner(obj uintptr) uintptr {
    ret, _, _ := saveDialog_GetOwner.Call(obj)
    return ret
}

func SaveDialog_GetName(obj uintptr) string {
    ret, _, _ := saveDialog_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func SaveDialog_SetName(obj uintptr, value string) {
   saveDialog_SetName.Call(obj, GoStrToDStr(value))
}

func SaveDialog_GetTag(obj uintptr) int {
    ret, _, _ := saveDialog_GetTag.Call(obj)
    return int(ret)
}

func SaveDialog_SetTag(obj uintptr, value int) {
   saveDialog_SetTag.Call(obj, uintptr(value))
}

func SaveDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := saveDialog_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TColorDialog ---------------------------

func ColorDialog_Create(obj uintptr) uintptr {
    ret, _, _ := colorDialog_Create.Call(obj)
    return ret
}

func ColorDialog_Free(obj uintptr) {
    colorDialog_Free.Call(obj)
}

func ColorDialog_Execute(obj uintptr, ParentWnd HWND) bool {
    ret, _, _ := colorDialog_Execute.Call(obj, uintptr(ParentWnd) )
    return DBoolToGoBool(ret)
}

func ColorDialog_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := colorDialog_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ColorDialog_GetNamePath(obj uintptr) string {
    ret, _, _ := colorDialog_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ColorDialog_HasParent(obj uintptr) bool {
    ret, _, _ := colorDialog_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ColorDialog_Assign(obj uintptr, Source uintptr)  {
    colorDialog_Assign.Call(obj, Source )
}

func ColorDialog_ClassName(obj uintptr) string {
    ret, _, _ := colorDialog_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ColorDialog_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := colorDialog_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ColorDialog_GetHashCode(obj uintptr) int32 {
    ret, _, _ := colorDialog_GetHashCode.Call(obj)
    return int32(ret)
}

func ColorDialog_ToString(obj uintptr) string {
    ret, _, _ := colorDialog_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ColorDialog_GetColor(obj uintptr) TColor {
    ret, _, _ := colorDialog_GetColor.Call(obj)
    return TColor(ret)
}

func ColorDialog_SetColor(obj uintptr, value TColor) {
   colorDialog_SetColor.Call(obj, uintptr(value))
}

func ColorDialog_GetOptions(obj uintptr) TColorDialogOptions {
    ret, _, _ := colorDialog_GetOptions.Call(obj)
    return TColorDialogOptions(ret)
}

func ColorDialog_SetOptions(obj uintptr, value TColorDialogOptions) {
   colorDialog_SetOptions.Call(obj, uintptr(value))
}

func ColorDialog_GetHandle(obj uintptr) HWND {
    ret, _, _ := colorDialog_GetHandle.Call(obj)
    return HWND(ret)
}

func ColorDialog_SetOnClose(obj uintptr, fn interface{}) {
    colorDialog_SetOnClose.Call(obj, addEventToMap(fn))
}

func ColorDialog_SetOnShow(obj uintptr, fn interface{}) {
    colorDialog_SetOnShow.Call(obj, addEventToMap(fn))
}

func ColorDialog_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := colorDialog_GetComponentCount.Call(obj)
    return int32(ret)
}

func ColorDialog_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := colorDialog_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ColorDialog_SetComponentIndex(obj uintptr, value int32) {
   colorDialog_SetComponentIndex.Call(obj, uintptr(value))
}

func ColorDialog_GetOwner(obj uintptr) uintptr {
    ret, _, _ := colorDialog_GetOwner.Call(obj)
    return ret
}

func ColorDialog_GetName(obj uintptr) string {
    ret, _, _ := colorDialog_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ColorDialog_SetName(obj uintptr, value string) {
   colorDialog_SetName.Call(obj, GoStrToDStr(value))
}

func ColorDialog_GetTag(obj uintptr) int {
    ret, _, _ := colorDialog_GetTag.Call(obj)
    return int(ret)
}

func ColorDialog_SetTag(obj uintptr, value int) {
   colorDialog_SetTag.Call(obj, uintptr(value))
}

func ColorDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := colorDialog_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TFontDialog ---------------------------

func FontDialog_Create(obj uintptr) uintptr {
    ret, _, _ := fontDialog_Create.Call(obj)
    return ret
}

func FontDialog_Free(obj uintptr) {
    fontDialog_Free.Call(obj)
}

func FontDialog_Execute(obj uintptr, ParentWnd HWND) bool {
    ret, _, _ := fontDialog_Execute.Call(obj, uintptr(ParentWnd) )
    return DBoolToGoBool(ret)
}

func FontDialog_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := fontDialog_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func FontDialog_GetNamePath(obj uintptr) string {
    ret, _, _ := fontDialog_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func FontDialog_HasParent(obj uintptr) bool {
    ret, _, _ := fontDialog_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func FontDialog_Assign(obj uintptr, Source uintptr)  {
    fontDialog_Assign.Call(obj, Source )
}

func FontDialog_ClassName(obj uintptr) string {
    ret, _, _ := fontDialog_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func FontDialog_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := fontDialog_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func FontDialog_GetHashCode(obj uintptr) int32 {
    ret, _, _ := fontDialog_GetHashCode.Call(obj)
    return int32(ret)
}

func FontDialog_ToString(obj uintptr) string {
    ret, _, _ := fontDialog_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func FontDialog_GetFont(obj uintptr) uintptr {
    ret, _, _ := fontDialog_GetFont.Call(obj)
    return ret
}

func FontDialog_SetFont(obj uintptr, value uintptr) {
   fontDialog_SetFont.Call(obj, value)
}

func FontDialog_GetOptions(obj uintptr) TFontDialogOptions {
    ret, _, _ := fontDialog_GetOptions.Call(obj)
    return TFontDialogOptions(ret)
}

func FontDialog_SetOptions(obj uintptr, value TFontDialogOptions) {
   fontDialog_SetOptions.Call(obj, uintptr(value))
}

func FontDialog_GetHandle(obj uintptr) HWND {
    ret, _, _ := fontDialog_GetHandle.Call(obj)
    return HWND(ret)
}

func FontDialog_SetOnClose(obj uintptr, fn interface{}) {
    fontDialog_SetOnClose.Call(obj, addEventToMap(fn))
}

func FontDialog_SetOnShow(obj uintptr, fn interface{}) {
    fontDialog_SetOnShow.Call(obj, addEventToMap(fn))
}

func FontDialog_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := fontDialog_GetComponentCount.Call(obj)
    return int32(ret)
}

func FontDialog_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := fontDialog_GetComponentIndex.Call(obj)
    return int32(ret)
}

func FontDialog_SetComponentIndex(obj uintptr, value int32) {
   fontDialog_SetComponentIndex.Call(obj, uintptr(value))
}

func FontDialog_GetOwner(obj uintptr) uintptr {
    ret, _, _ := fontDialog_GetOwner.Call(obj)
    return ret
}

func FontDialog_GetName(obj uintptr) string {
    ret, _, _ := fontDialog_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func FontDialog_SetName(obj uintptr, value string) {
   fontDialog_SetName.Call(obj, GoStrToDStr(value))
}

func FontDialog_GetTag(obj uintptr) int {
    ret, _, _ := fontDialog_GetTag.Call(obj)
    return int(ret)
}

func FontDialog_SetTag(obj uintptr, value int) {
   fontDialog_SetTag.Call(obj, uintptr(value))
}

func FontDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := fontDialog_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TPrintDialog ---------------------------

func PrintDialog_Create(obj uintptr) uintptr {
    ret, _, _ := printDialog_Create.Call(obj)
    return ret
}

func PrintDialog_Free(obj uintptr) {
    printDialog_Free.Call(obj)
}

func PrintDialog_Execute(obj uintptr, ParentWnd HWND) bool {
    ret, _, _ := printDialog_Execute.Call(obj, uintptr(ParentWnd) )
    return DBoolToGoBool(ret)
}

func PrintDialog_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := printDialog_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func PrintDialog_GetNamePath(obj uintptr) string {
    ret, _, _ := printDialog_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func PrintDialog_HasParent(obj uintptr) bool {
    ret, _, _ := printDialog_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func PrintDialog_Assign(obj uintptr, Source uintptr)  {
    printDialog_Assign.Call(obj, Source )
}

func PrintDialog_ClassName(obj uintptr) string {
    ret, _, _ := printDialog_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func PrintDialog_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := printDialog_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func PrintDialog_GetHashCode(obj uintptr) int32 {
    ret, _, _ := printDialog_GetHashCode.Call(obj)
    return int32(ret)
}

func PrintDialog_ToString(obj uintptr) string {
    ret, _, _ := printDialog_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func PrintDialog_GetCollate(obj uintptr) bool {
    ret, _, _ := printDialog_GetCollate.Call(obj)
    return DBoolToGoBool(ret)
}

func PrintDialog_SetCollate(obj uintptr, value bool) {
   printDialog_SetCollate.Call(obj, GoBoolToDBool(value))
}

func PrintDialog_GetCopies(obj uintptr) int32 {
    ret, _, _ := printDialog_GetCopies.Call(obj)
    return int32(ret)
}

func PrintDialog_SetCopies(obj uintptr, value int32) {
   printDialog_SetCopies.Call(obj, uintptr(value))
}

func PrintDialog_GetFromPage(obj uintptr) int32 {
    ret, _, _ := printDialog_GetFromPage.Call(obj)
    return int32(ret)
}

func PrintDialog_SetFromPage(obj uintptr, value int32) {
   printDialog_SetFromPage.Call(obj, uintptr(value))
}

func PrintDialog_GetMinPage(obj uintptr) int32 {
    ret, _, _ := printDialog_GetMinPage.Call(obj)
    return int32(ret)
}

func PrintDialog_SetMinPage(obj uintptr, value int32) {
   printDialog_SetMinPage.Call(obj, uintptr(value))
}

func PrintDialog_GetMaxPage(obj uintptr) int32 {
    ret, _, _ := printDialog_GetMaxPage.Call(obj)
    return int32(ret)
}

func PrintDialog_SetMaxPage(obj uintptr, value int32) {
   printDialog_SetMaxPage.Call(obj, uintptr(value))
}

func PrintDialog_GetOptions(obj uintptr) TPrintDialogOptions {
    ret, _, _ := printDialog_GetOptions.Call(obj)
    return TPrintDialogOptions(ret)
}

func PrintDialog_SetOptions(obj uintptr, value TPrintDialogOptions) {
   printDialog_SetOptions.Call(obj, uintptr(value))
}

func PrintDialog_GetPrintToFile(obj uintptr) bool {
    ret, _, _ := printDialog_GetPrintToFile.Call(obj)
    return DBoolToGoBool(ret)
}

func PrintDialog_SetPrintToFile(obj uintptr, value bool) {
   printDialog_SetPrintToFile.Call(obj, GoBoolToDBool(value))
}

func PrintDialog_GetPrintRange(obj uintptr) TPrintRange {
    ret, _, _ := printDialog_GetPrintRange.Call(obj)
    return TPrintRange(ret)
}

func PrintDialog_SetPrintRange(obj uintptr, value TPrintRange) {
   printDialog_SetPrintRange.Call(obj, uintptr(value))
}

func PrintDialog_GetToPage(obj uintptr) int32 {
    ret, _, _ := printDialog_GetToPage.Call(obj)
    return int32(ret)
}

func PrintDialog_SetToPage(obj uintptr, value int32) {
   printDialog_SetToPage.Call(obj, uintptr(value))
}

func PrintDialog_GetHandle(obj uintptr) HWND {
    ret, _, _ := printDialog_GetHandle.Call(obj)
    return HWND(ret)
}

func PrintDialog_SetOnClose(obj uintptr, fn interface{}) {
    printDialog_SetOnClose.Call(obj, addEventToMap(fn))
}

func PrintDialog_SetOnShow(obj uintptr, fn interface{}) {
    printDialog_SetOnShow.Call(obj, addEventToMap(fn))
}

func PrintDialog_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := printDialog_GetComponentCount.Call(obj)
    return int32(ret)
}

func PrintDialog_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := printDialog_GetComponentIndex.Call(obj)
    return int32(ret)
}

func PrintDialog_SetComponentIndex(obj uintptr, value int32) {
   printDialog_SetComponentIndex.Call(obj, uintptr(value))
}

func PrintDialog_GetOwner(obj uintptr) uintptr {
    ret, _, _ := printDialog_GetOwner.Call(obj)
    return ret
}

func PrintDialog_GetName(obj uintptr) string {
    ret, _, _ := printDialog_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func PrintDialog_SetName(obj uintptr, value string) {
   printDialog_SetName.Call(obj, GoStrToDStr(value))
}

func PrintDialog_GetTag(obj uintptr) int {
    ret, _, _ := printDialog_GetTag.Call(obj)
    return int(ret)
}

func PrintDialog_SetTag(obj uintptr, value int) {
   printDialog_SetTag.Call(obj, uintptr(value))
}

func PrintDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := printDialog_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TOpenPictureDialog ---------------------------

func OpenPictureDialog_Create(obj uintptr) uintptr {
    ret, _, _ := openPictureDialog_Create.Call(obj)
    return ret
}

func OpenPictureDialog_Free(obj uintptr) {
    openPictureDialog_Free.Call(obj)
}

func OpenPictureDialog_Execute(obj uintptr, ParentWnd HWND) bool {
    ret, _, _ := openPictureDialog_Execute.Call(obj, uintptr(ParentWnd) )
    return DBoolToGoBool(ret)
}

func OpenPictureDialog_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := openPictureDialog_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func OpenPictureDialog_GetNamePath(obj uintptr) string {
    ret, _, _ := openPictureDialog_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func OpenPictureDialog_HasParent(obj uintptr) bool {
    ret, _, _ := openPictureDialog_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func OpenPictureDialog_Assign(obj uintptr, Source uintptr)  {
    openPictureDialog_Assign.Call(obj, Source )
}

func OpenPictureDialog_ClassName(obj uintptr) string {
    ret, _, _ := openPictureDialog_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func OpenPictureDialog_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := openPictureDialog_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func OpenPictureDialog_GetHashCode(obj uintptr) int32 {
    ret, _, _ := openPictureDialog_GetHashCode.Call(obj)
    return int32(ret)
}

func OpenPictureDialog_ToString(obj uintptr) string {
    ret, _, _ := openPictureDialog_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func OpenPictureDialog_GetFilter(obj uintptr) string {
    ret, _, _ := openPictureDialog_GetFilter.Call(obj)
    return DStrToGoStr(ret)
}

func OpenPictureDialog_SetFilter(obj uintptr, value string) {
   openPictureDialog_SetFilter.Call(obj, GoStrToDStr(value))
}

func OpenPictureDialog_GetFiles(obj uintptr) uintptr {
    ret, _, _ := openPictureDialog_GetFiles.Call(obj)
    return ret
}

func OpenPictureDialog_GetDefaultExt(obj uintptr) string {
    ret, _, _ := openPictureDialog_GetDefaultExt.Call(obj)
    return DStrToGoStr(ret)
}

func OpenPictureDialog_SetDefaultExt(obj uintptr, value string) {
   openPictureDialog_SetDefaultExt.Call(obj, GoStrToDStr(value))
}

func OpenPictureDialog_GetFileName(obj uintptr) string {
    ret, _, _ := openPictureDialog_GetFileName.Call(obj)
    return DStrToGoStr(ret)
}

func OpenPictureDialog_SetFileName(obj uintptr, value string) {
   openPictureDialog_SetFileName.Call(obj, GoStrToDStr(value))
}

func OpenPictureDialog_GetFilterIndex(obj uintptr) int32 {
    ret, _, _ := openPictureDialog_GetFilterIndex.Call(obj)
    return int32(ret)
}

func OpenPictureDialog_SetFilterIndex(obj uintptr, value int32) {
   openPictureDialog_SetFilterIndex.Call(obj, uintptr(value))
}

func OpenPictureDialog_GetInitialDir(obj uintptr) string {
    ret, _, _ := openPictureDialog_GetInitialDir.Call(obj)
    return DStrToGoStr(ret)
}

func OpenPictureDialog_SetInitialDir(obj uintptr, value string) {
   openPictureDialog_SetInitialDir.Call(obj, GoStrToDStr(value))
}

func OpenPictureDialog_GetOptions(obj uintptr) TOpenOptions {
    ret, _, _ := openPictureDialog_GetOptions.Call(obj)
    return TOpenOptions(ret)
}

func OpenPictureDialog_SetOptions(obj uintptr, value TOpenOptions) {
   openPictureDialog_SetOptions.Call(obj, uintptr(value))
}

func OpenPictureDialog_GetOptionsEx(obj uintptr) TOpenOptionsEx {
    ret, _, _ := openPictureDialog_GetOptionsEx.Call(obj)
    return TOpenOptionsEx(ret)
}

func OpenPictureDialog_SetOptionsEx(obj uintptr, value TOpenOptionsEx) {
   openPictureDialog_SetOptionsEx.Call(obj, uintptr(value))
}

func OpenPictureDialog_GetTitle(obj uintptr) string {
    ret, _, _ := openPictureDialog_GetTitle.Call(obj)
    return DStrToGoStr(ret)
}

func OpenPictureDialog_SetTitle(obj uintptr, value string) {
   openPictureDialog_SetTitle.Call(obj, GoStrToDStr(value))
}

func OpenPictureDialog_GetHandle(obj uintptr) HWND {
    ret, _, _ := openPictureDialog_GetHandle.Call(obj)
    return HWND(ret)
}

func OpenPictureDialog_SetOnClose(obj uintptr, fn interface{}) {
    openPictureDialog_SetOnClose.Call(obj, addEventToMap(fn))
}

func OpenPictureDialog_SetOnShow(obj uintptr, fn interface{}) {
    openPictureDialog_SetOnShow.Call(obj, addEventToMap(fn))
}

func OpenPictureDialog_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := openPictureDialog_GetComponentCount.Call(obj)
    return int32(ret)
}

func OpenPictureDialog_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := openPictureDialog_GetComponentIndex.Call(obj)
    return int32(ret)
}

func OpenPictureDialog_SetComponentIndex(obj uintptr, value int32) {
   openPictureDialog_SetComponentIndex.Call(obj, uintptr(value))
}

func OpenPictureDialog_GetOwner(obj uintptr) uintptr {
    ret, _, _ := openPictureDialog_GetOwner.Call(obj)
    return ret
}

func OpenPictureDialog_GetName(obj uintptr) string {
    ret, _, _ := openPictureDialog_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func OpenPictureDialog_SetName(obj uintptr, value string) {
   openPictureDialog_SetName.Call(obj, GoStrToDStr(value))
}

func OpenPictureDialog_GetTag(obj uintptr) int {
    ret, _, _ := openPictureDialog_GetTag.Call(obj)
    return int(ret)
}

func OpenPictureDialog_SetTag(obj uintptr, value int) {
   openPictureDialog_SetTag.Call(obj, uintptr(value))
}

func OpenPictureDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := openPictureDialog_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TSavePictureDialog ---------------------------

func SavePictureDialog_Create(obj uintptr) uintptr {
    ret, _, _ := savePictureDialog_Create.Call(obj)
    return ret
}

func SavePictureDialog_Free(obj uintptr) {
    savePictureDialog_Free.Call(obj)
}

func SavePictureDialog_Execute(obj uintptr) bool {
    ret, _, _ := savePictureDialog_Execute.Call(obj)
    return DBoolToGoBool(ret)
}

func SavePictureDialog_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := savePictureDialog_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func SavePictureDialog_GetNamePath(obj uintptr) string {
    ret, _, _ := savePictureDialog_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func SavePictureDialog_HasParent(obj uintptr) bool {
    ret, _, _ := savePictureDialog_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func SavePictureDialog_Assign(obj uintptr, Source uintptr)  {
    savePictureDialog_Assign.Call(obj, Source )
}

func SavePictureDialog_ClassName(obj uintptr) string {
    ret, _, _ := savePictureDialog_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func SavePictureDialog_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := savePictureDialog_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func SavePictureDialog_GetHashCode(obj uintptr) int32 {
    ret, _, _ := savePictureDialog_GetHashCode.Call(obj)
    return int32(ret)
}

func SavePictureDialog_ToString(obj uintptr) string {
    ret, _, _ := savePictureDialog_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func SavePictureDialog_GetFilter(obj uintptr) string {
    ret, _, _ := savePictureDialog_GetFilter.Call(obj)
    return DStrToGoStr(ret)
}

func SavePictureDialog_SetFilter(obj uintptr, value string) {
   savePictureDialog_SetFilter.Call(obj, GoStrToDStr(value))
}

func SavePictureDialog_GetFiles(obj uintptr) uintptr {
    ret, _, _ := savePictureDialog_GetFiles.Call(obj)
    return ret
}

func SavePictureDialog_GetDefaultExt(obj uintptr) string {
    ret, _, _ := savePictureDialog_GetDefaultExt.Call(obj)
    return DStrToGoStr(ret)
}

func SavePictureDialog_SetDefaultExt(obj uintptr, value string) {
   savePictureDialog_SetDefaultExt.Call(obj, GoStrToDStr(value))
}

func SavePictureDialog_GetFileName(obj uintptr) string {
    ret, _, _ := savePictureDialog_GetFileName.Call(obj)
    return DStrToGoStr(ret)
}

func SavePictureDialog_SetFileName(obj uintptr, value string) {
   savePictureDialog_SetFileName.Call(obj, GoStrToDStr(value))
}

func SavePictureDialog_GetFilterIndex(obj uintptr) int32 {
    ret, _, _ := savePictureDialog_GetFilterIndex.Call(obj)
    return int32(ret)
}

func SavePictureDialog_SetFilterIndex(obj uintptr, value int32) {
   savePictureDialog_SetFilterIndex.Call(obj, uintptr(value))
}

func SavePictureDialog_GetInitialDir(obj uintptr) string {
    ret, _, _ := savePictureDialog_GetInitialDir.Call(obj)
    return DStrToGoStr(ret)
}

func SavePictureDialog_SetInitialDir(obj uintptr, value string) {
   savePictureDialog_SetInitialDir.Call(obj, GoStrToDStr(value))
}

func SavePictureDialog_GetOptions(obj uintptr) TOpenOptions {
    ret, _, _ := savePictureDialog_GetOptions.Call(obj)
    return TOpenOptions(ret)
}

func SavePictureDialog_SetOptions(obj uintptr, value TOpenOptions) {
   savePictureDialog_SetOptions.Call(obj, uintptr(value))
}

func SavePictureDialog_GetOptionsEx(obj uintptr) TOpenOptionsEx {
    ret, _, _ := savePictureDialog_GetOptionsEx.Call(obj)
    return TOpenOptionsEx(ret)
}

func SavePictureDialog_SetOptionsEx(obj uintptr, value TOpenOptionsEx) {
   savePictureDialog_SetOptionsEx.Call(obj, uintptr(value))
}

func SavePictureDialog_GetTitle(obj uintptr) string {
    ret, _, _ := savePictureDialog_GetTitle.Call(obj)
    return DStrToGoStr(ret)
}

func SavePictureDialog_SetTitle(obj uintptr, value string) {
   savePictureDialog_SetTitle.Call(obj, GoStrToDStr(value))
}

func SavePictureDialog_GetHandle(obj uintptr) HWND {
    ret, _, _ := savePictureDialog_GetHandle.Call(obj)
    return HWND(ret)
}

func SavePictureDialog_SetOnClose(obj uintptr, fn interface{}) {
    savePictureDialog_SetOnClose.Call(obj, addEventToMap(fn))
}

func SavePictureDialog_SetOnShow(obj uintptr, fn interface{}) {
    savePictureDialog_SetOnShow.Call(obj, addEventToMap(fn))
}

func SavePictureDialog_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := savePictureDialog_GetComponentCount.Call(obj)
    return int32(ret)
}

func SavePictureDialog_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := savePictureDialog_GetComponentIndex.Call(obj)
    return int32(ret)
}

func SavePictureDialog_SetComponentIndex(obj uintptr, value int32) {
   savePictureDialog_SetComponentIndex.Call(obj, uintptr(value))
}

func SavePictureDialog_GetOwner(obj uintptr) uintptr {
    ret, _, _ := savePictureDialog_GetOwner.Call(obj)
    return ret
}

func SavePictureDialog_GetName(obj uintptr) string {
    ret, _, _ := savePictureDialog_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func SavePictureDialog_SetName(obj uintptr, value string) {
   savePictureDialog_SetName.Call(obj, GoStrToDStr(value))
}

func SavePictureDialog_GetTag(obj uintptr) int {
    ret, _, _ := savePictureDialog_GetTag.Call(obj)
    return int(ret)
}

func SavePictureDialog_SetTag(obj uintptr, value int) {
   savePictureDialog_SetTag.Call(obj, uintptr(value))
}

func SavePictureDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := savePictureDialog_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TSaveTextFileDialog ---------------------------

func SaveTextFileDialog_Create(obj uintptr) uintptr {
    ret, _, _ := saveTextFileDialog_Create.Call(obj)
    return ret
}

func SaveTextFileDialog_Free(obj uintptr) {
    saveTextFileDialog_Free.Call(obj)
}

func SaveTextFileDialog_Execute(obj uintptr) bool {
    ret, _, _ := saveTextFileDialog_Execute.Call(obj)
    return DBoolToGoBool(ret)
}

func SaveTextFileDialog_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := saveTextFileDialog_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func SaveTextFileDialog_GetNamePath(obj uintptr) string {
    ret, _, _ := saveTextFileDialog_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func SaveTextFileDialog_HasParent(obj uintptr) bool {
    ret, _, _ := saveTextFileDialog_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func SaveTextFileDialog_Assign(obj uintptr, Source uintptr)  {
    saveTextFileDialog_Assign.Call(obj, Source )
}

func SaveTextFileDialog_ClassName(obj uintptr) string {
    ret, _, _ := saveTextFileDialog_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func SaveTextFileDialog_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := saveTextFileDialog_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func SaveTextFileDialog_GetHashCode(obj uintptr) int32 {
    ret, _, _ := saveTextFileDialog_GetHashCode.Call(obj)
    return int32(ret)
}

func SaveTextFileDialog_ToString(obj uintptr) string {
    ret, _, _ := saveTextFileDialog_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func SaveTextFileDialog_GetFiles(obj uintptr) uintptr {
    ret, _, _ := saveTextFileDialog_GetFiles.Call(obj)
    return ret
}

func SaveTextFileDialog_GetDefaultExt(obj uintptr) string {
    ret, _, _ := saveTextFileDialog_GetDefaultExt.Call(obj)
    return DStrToGoStr(ret)
}

func SaveTextFileDialog_SetDefaultExt(obj uintptr, value string) {
   saveTextFileDialog_SetDefaultExt.Call(obj, GoStrToDStr(value))
}

func SaveTextFileDialog_GetFileName(obj uintptr) string {
    ret, _, _ := saveTextFileDialog_GetFileName.Call(obj)
    return DStrToGoStr(ret)
}

func SaveTextFileDialog_SetFileName(obj uintptr, value string) {
   saveTextFileDialog_SetFileName.Call(obj, GoStrToDStr(value))
}

func SaveTextFileDialog_GetFilter(obj uintptr) string {
    ret, _, _ := saveTextFileDialog_GetFilter.Call(obj)
    return DStrToGoStr(ret)
}

func SaveTextFileDialog_SetFilter(obj uintptr, value string) {
   saveTextFileDialog_SetFilter.Call(obj, GoStrToDStr(value))
}

func SaveTextFileDialog_GetFilterIndex(obj uintptr) int32 {
    ret, _, _ := saveTextFileDialog_GetFilterIndex.Call(obj)
    return int32(ret)
}

func SaveTextFileDialog_SetFilterIndex(obj uintptr, value int32) {
   saveTextFileDialog_SetFilterIndex.Call(obj, uintptr(value))
}

func SaveTextFileDialog_GetInitialDir(obj uintptr) string {
    ret, _, _ := saveTextFileDialog_GetInitialDir.Call(obj)
    return DStrToGoStr(ret)
}

func SaveTextFileDialog_SetInitialDir(obj uintptr, value string) {
   saveTextFileDialog_SetInitialDir.Call(obj, GoStrToDStr(value))
}

func SaveTextFileDialog_GetOptions(obj uintptr) TOpenOptions {
    ret, _, _ := saveTextFileDialog_GetOptions.Call(obj)
    return TOpenOptions(ret)
}

func SaveTextFileDialog_SetOptions(obj uintptr, value TOpenOptions) {
   saveTextFileDialog_SetOptions.Call(obj, uintptr(value))
}

func SaveTextFileDialog_GetOptionsEx(obj uintptr) TOpenOptionsEx {
    ret, _, _ := saveTextFileDialog_GetOptionsEx.Call(obj)
    return TOpenOptionsEx(ret)
}

func SaveTextFileDialog_SetOptionsEx(obj uintptr, value TOpenOptionsEx) {
   saveTextFileDialog_SetOptionsEx.Call(obj, uintptr(value))
}

func SaveTextFileDialog_GetTitle(obj uintptr) string {
    ret, _, _ := saveTextFileDialog_GetTitle.Call(obj)
    return DStrToGoStr(ret)
}

func SaveTextFileDialog_SetTitle(obj uintptr, value string) {
   saveTextFileDialog_SetTitle.Call(obj, GoStrToDStr(value))
}

func SaveTextFileDialog_GetHandle(obj uintptr) HWND {
    ret, _, _ := saveTextFileDialog_GetHandle.Call(obj)
    return HWND(ret)
}

func SaveTextFileDialog_SetOnClose(obj uintptr, fn interface{}) {
    saveTextFileDialog_SetOnClose.Call(obj, addEventToMap(fn))
}

func SaveTextFileDialog_SetOnShow(obj uintptr, fn interface{}) {
    saveTextFileDialog_SetOnShow.Call(obj, addEventToMap(fn))
}

func SaveTextFileDialog_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := saveTextFileDialog_GetComponentCount.Call(obj)
    return int32(ret)
}

func SaveTextFileDialog_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := saveTextFileDialog_GetComponentIndex.Call(obj)
    return int32(ret)
}

func SaveTextFileDialog_SetComponentIndex(obj uintptr, value int32) {
   saveTextFileDialog_SetComponentIndex.Call(obj, uintptr(value))
}

func SaveTextFileDialog_GetOwner(obj uintptr) uintptr {
    ret, _, _ := saveTextFileDialog_GetOwner.Call(obj)
    return ret
}

func SaveTextFileDialog_GetName(obj uintptr) string {
    ret, _, _ := saveTextFileDialog_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func SaveTextFileDialog_SetName(obj uintptr, value string) {
   saveTextFileDialog_SetName.Call(obj, GoStrToDStr(value))
}

func SaveTextFileDialog_GetTag(obj uintptr) int {
    ret, _, _ := saveTextFileDialog_GetTag.Call(obj)
    return int(ret)
}

func SaveTextFileDialog_SetTag(obj uintptr, value int) {
   saveTextFileDialog_SetTag.Call(obj, uintptr(value))
}

func SaveTextFileDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := saveTextFileDialog_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TOpenTextFileDialog ---------------------------

func OpenTextFileDialog_Create(obj uintptr) uintptr {
    ret, _, _ := openTextFileDialog_Create.Call(obj)
    return ret
}

func OpenTextFileDialog_Free(obj uintptr) {
    openTextFileDialog_Free.Call(obj)
}

func OpenTextFileDialog_Execute(obj uintptr, ParentWnd HWND) bool {
    ret, _, _ := openTextFileDialog_Execute.Call(obj, uintptr(ParentWnd) )
    return DBoolToGoBool(ret)
}

func OpenTextFileDialog_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := openTextFileDialog_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func OpenTextFileDialog_GetNamePath(obj uintptr) string {
    ret, _, _ := openTextFileDialog_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func OpenTextFileDialog_HasParent(obj uintptr) bool {
    ret, _, _ := openTextFileDialog_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func OpenTextFileDialog_Assign(obj uintptr, Source uintptr)  {
    openTextFileDialog_Assign.Call(obj, Source )
}

func OpenTextFileDialog_ClassName(obj uintptr) string {
    ret, _, _ := openTextFileDialog_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func OpenTextFileDialog_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := openTextFileDialog_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func OpenTextFileDialog_GetHashCode(obj uintptr) int32 {
    ret, _, _ := openTextFileDialog_GetHashCode.Call(obj)
    return int32(ret)
}

func OpenTextFileDialog_ToString(obj uintptr) string {
    ret, _, _ := openTextFileDialog_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func OpenTextFileDialog_GetFiles(obj uintptr) uintptr {
    ret, _, _ := openTextFileDialog_GetFiles.Call(obj)
    return ret
}

func OpenTextFileDialog_GetDefaultExt(obj uintptr) string {
    ret, _, _ := openTextFileDialog_GetDefaultExt.Call(obj)
    return DStrToGoStr(ret)
}

func OpenTextFileDialog_SetDefaultExt(obj uintptr, value string) {
   openTextFileDialog_SetDefaultExt.Call(obj, GoStrToDStr(value))
}

func OpenTextFileDialog_GetFileName(obj uintptr) string {
    ret, _, _ := openTextFileDialog_GetFileName.Call(obj)
    return DStrToGoStr(ret)
}

func OpenTextFileDialog_SetFileName(obj uintptr, value string) {
   openTextFileDialog_SetFileName.Call(obj, GoStrToDStr(value))
}

func OpenTextFileDialog_GetFilter(obj uintptr) string {
    ret, _, _ := openTextFileDialog_GetFilter.Call(obj)
    return DStrToGoStr(ret)
}

func OpenTextFileDialog_SetFilter(obj uintptr, value string) {
   openTextFileDialog_SetFilter.Call(obj, GoStrToDStr(value))
}

func OpenTextFileDialog_GetFilterIndex(obj uintptr) int32 {
    ret, _, _ := openTextFileDialog_GetFilterIndex.Call(obj)
    return int32(ret)
}

func OpenTextFileDialog_SetFilterIndex(obj uintptr, value int32) {
   openTextFileDialog_SetFilterIndex.Call(obj, uintptr(value))
}

func OpenTextFileDialog_GetInitialDir(obj uintptr) string {
    ret, _, _ := openTextFileDialog_GetInitialDir.Call(obj)
    return DStrToGoStr(ret)
}

func OpenTextFileDialog_SetInitialDir(obj uintptr, value string) {
   openTextFileDialog_SetInitialDir.Call(obj, GoStrToDStr(value))
}

func OpenTextFileDialog_GetOptions(obj uintptr) TOpenOptions {
    ret, _, _ := openTextFileDialog_GetOptions.Call(obj)
    return TOpenOptions(ret)
}

func OpenTextFileDialog_SetOptions(obj uintptr, value TOpenOptions) {
   openTextFileDialog_SetOptions.Call(obj, uintptr(value))
}

func OpenTextFileDialog_GetOptionsEx(obj uintptr) TOpenOptionsEx {
    ret, _, _ := openTextFileDialog_GetOptionsEx.Call(obj)
    return TOpenOptionsEx(ret)
}

func OpenTextFileDialog_SetOptionsEx(obj uintptr, value TOpenOptionsEx) {
   openTextFileDialog_SetOptionsEx.Call(obj, uintptr(value))
}

func OpenTextFileDialog_GetTitle(obj uintptr) string {
    ret, _, _ := openTextFileDialog_GetTitle.Call(obj)
    return DStrToGoStr(ret)
}

func OpenTextFileDialog_SetTitle(obj uintptr, value string) {
   openTextFileDialog_SetTitle.Call(obj, GoStrToDStr(value))
}

func OpenTextFileDialog_GetHandle(obj uintptr) HWND {
    ret, _, _ := openTextFileDialog_GetHandle.Call(obj)
    return HWND(ret)
}

func OpenTextFileDialog_SetOnClose(obj uintptr, fn interface{}) {
    openTextFileDialog_SetOnClose.Call(obj, addEventToMap(fn))
}

func OpenTextFileDialog_SetOnShow(obj uintptr, fn interface{}) {
    openTextFileDialog_SetOnShow.Call(obj, addEventToMap(fn))
}

func OpenTextFileDialog_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := openTextFileDialog_GetComponentCount.Call(obj)
    return int32(ret)
}

func OpenTextFileDialog_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := openTextFileDialog_GetComponentIndex.Call(obj)
    return int32(ret)
}

func OpenTextFileDialog_SetComponentIndex(obj uintptr, value int32) {
   openTextFileDialog_SetComponentIndex.Call(obj, uintptr(value))
}

func OpenTextFileDialog_GetOwner(obj uintptr) uintptr {
    ret, _, _ := openTextFileDialog_GetOwner.Call(obj)
    return ret
}

func OpenTextFileDialog_GetName(obj uintptr) string {
    ret, _, _ := openTextFileDialog_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func OpenTextFileDialog_SetName(obj uintptr, value string) {
   openTextFileDialog_SetName.Call(obj, GoStrToDStr(value))
}

func OpenTextFileDialog_GetTag(obj uintptr) int {
    ret, _, _ := openTextFileDialog_GetTag.Call(obj)
    return int(ret)
}

func OpenTextFileDialog_SetTag(obj uintptr, value int) {
   openTextFileDialog_SetTag.Call(obj, uintptr(value))
}

func OpenTextFileDialog_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := openTextFileDialog_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TRichEdit ---------------------------

func RichEdit_Create(obj uintptr) uintptr {
    ret, _, _ := richEdit_Create.Call(obj)
    return ret
}

func RichEdit_Free(obj uintptr) {
    richEdit_Free.Call(obj)
}

func RichEdit_Clear(obj uintptr)  {
    richEdit_Clear.Call(obj)
}

func RichEdit_FindText(obj uintptr, SearchStr string, StartPos int32, Length int32, Options TSearchTypes) int32 {
    ret, _, _ := richEdit_FindText.Call(obj, GoStrToDStr(SearchStr) , uintptr(StartPos) , uintptr(Length) , uintptr(Options) )
    return int32(ret)
}

func RichEdit_Print(obj uintptr, Caption string)  {
    richEdit_Print.Call(obj, GoStrToDStr(Caption) )
}

func RichEdit_GetSelTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := richEdit_GetSelTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func RichEdit_ClearSelection(obj uintptr)  {
    richEdit_ClearSelection.Call(obj)
}

func RichEdit_CopyToClipboard(obj uintptr)  {
    richEdit_CopyToClipboard.Call(obj)
}

func RichEdit_CutToClipboard(obj uintptr)  {
    richEdit_CutToClipboard.Call(obj)
}

func RichEdit_PasteFromClipboard(obj uintptr)  {
    richEdit_PasteFromClipboard.Call(obj)
}

func RichEdit_SelectAll(obj uintptr)  {
    richEdit_SelectAll.Call(obj)
}

func RichEdit_CanFocus(obj uintptr) bool {
    ret, _, _ := richEdit_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_FlipChildren(obj uintptr, AllLevels bool)  {
    richEdit_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func RichEdit_Focused(obj uintptr) bool {
    ret, _, _ := richEdit_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_HandleAllocated(obj uintptr) bool {
    ret, _, _ := richEdit_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_Invalidate(obj uintptr)  {
    richEdit_Invalidate.Call(obj)
}

func RichEdit_Realign(obj uintptr)  {
    richEdit_Realign.Call(obj)
}

func RichEdit_Repaint(obj uintptr)  {
    richEdit_Repaint.Call(obj)
}

func RichEdit_ScaleBy(obj uintptr, M int32, D int32)  {
    richEdit_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func RichEdit_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    richEdit_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func RichEdit_SetFocus(obj uintptr)  {
    richEdit_SetFocus.Call(obj)
}

func RichEdit_Update(obj uintptr)  {
    richEdit_Update.Call(obj)
}

func RichEdit_BringToFront(obj uintptr)  {
    richEdit_BringToFront.Call(obj)
}

func RichEdit_HasParent(obj uintptr) bool {
    ret, _, _ := richEdit_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_Hide(obj uintptr)  {
    richEdit_Hide.Call(obj)
}

func RichEdit_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := richEdit_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func RichEdit_Refresh(obj uintptr)  {
    richEdit_Refresh.Call(obj)
}

func RichEdit_SendToBack(obj uintptr)  {
    richEdit_SendToBack.Call(obj)
}

func RichEdit_Show(obj uintptr)  {
    richEdit_Show.Call(obj)
}

func RichEdit_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := richEdit_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func RichEdit_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := richEdit_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func RichEdit_GetNamePath(obj uintptr) string {
    ret, _, _ := richEdit_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func RichEdit_Assign(obj uintptr, Source uintptr)  {
    richEdit_Assign.Call(obj, Source )
}

func RichEdit_ClassName(obj uintptr) string {
    ret, _, _ := richEdit_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func RichEdit_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := richEdit_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func RichEdit_GetHashCode(obj uintptr) int32 {
    ret, _, _ := richEdit_GetHashCode.Call(obj)
    return int32(ret)
}

func RichEdit_ToString(obj uintptr) string {
    ret, _, _ := richEdit_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func RichEdit_GetAlign(obj uintptr) TAlign {
    ret, _, _ := richEdit_GetAlign.Call(obj)
    return TAlign(ret)
}

func RichEdit_SetAlign(obj uintptr, value TAlign) {
   richEdit_SetAlign.Call(obj, uintptr(value))
}

func RichEdit_GetAlignment(obj uintptr) TAlignment {
    ret, _, _ := richEdit_GetAlignment.Call(obj)
    return TAlignment(ret)
}

func RichEdit_SetAlignment(obj uintptr, value TAlignment) {
   richEdit_SetAlignment.Call(obj, uintptr(value))
}

func RichEdit_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := richEdit_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func RichEdit_SetAnchors(obj uintptr, value TAnchors) {
   richEdit_SetAnchors.Call(obj, uintptr(value))
}

func RichEdit_GetBevelEdges(obj uintptr) TBevelEdges {
    ret, _, _ := richEdit_GetBevelEdges.Call(obj)
    return TBevelEdges(ret)
}

func RichEdit_SetBevelEdges(obj uintptr, value TBevelEdges) {
   richEdit_SetBevelEdges.Call(obj, uintptr(value))
}

func RichEdit_GetBevelInner(obj uintptr) TBevelCut {
    ret, _, _ := richEdit_GetBevelInner.Call(obj)
    return TBevelCut(ret)
}

func RichEdit_SetBevelInner(obj uintptr, value TBevelCut) {
   richEdit_SetBevelInner.Call(obj, uintptr(value))
}

func RichEdit_GetBevelOuter(obj uintptr) TBevelCut {
    ret, _, _ := richEdit_GetBevelOuter.Call(obj)
    return TBevelCut(ret)
}

func RichEdit_SetBevelOuter(obj uintptr, value TBevelCut) {
   richEdit_SetBevelOuter.Call(obj, uintptr(value))
}

func RichEdit_GetBevelKind(obj uintptr) TBevelKind {
    ret, _, _ := richEdit_GetBevelKind.Call(obj)
    return TBevelKind(ret)
}

func RichEdit_SetBevelKind(obj uintptr, value TBevelKind) {
   richEdit_SetBevelKind.Call(obj, uintptr(value))
}

func RichEdit_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := richEdit_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func RichEdit_SetBiDiMode(obj uintptr, value TBiDiMode) {
   richEdit_SetBiDiMode.Call(obj, uintptr(value))
}

func RichEdit_GetBorderStyle(obj uintptr) TBorderStyle {
    ret, _, _ := richEdit_GetBorderStyle.Call(obj)
    return TBorderStyle(ret)
}

func RichEdit_SetBorderStyle(obj uintptr, value TBorderStyle) {
   richEdit_SetBorderStyle.Call(obj, uintptr(value))
}

func RichEdit_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := richEdit_GetBorderWidth.Call(obj)
    return int32(ret)
}

func RichEdit_SetBorderWidth(obj uintptr, value int32) {
   richEdit_SetBorderWidth.Call(obj, uintptr(value))
}

func RichEdit_GetColor(obj uintptr) TColor {
    ret, _, _ := richEdit_GetColor.Call(obj)
    return TColor(ret)
}

func RichEdit_SetColor(obj uintptr, value TColor) {
   richEdit_SetColor.Call(obj, uintptr(value))
}

func RichEdit_GetEnabled(obj uintptr) bool {
    ret, _, _ := richEdit_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetEnabled(obj uintptr, value bool) {
   richEdit_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetFont(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetFont.Call(obj)
    return ret
}

func RichEdit_SetFont(obj uintptr, value uintptr) {
   richEdit_SetFont.Call(obj, value)
}

func RichEdit_GetHideSelection(obj uintptr) bool {
    ret, _, _ := richEdit_GetHideSelection.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetHideSelection(obj uintptr, value bool) {
   richEdit_SetHideSelection.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetHideScrollBars(obj uintptr) bool {
    ret, _, _ := richEdit_GetHideScrollBars.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetHideScrollBars(obj uintptr, value bool) {
   richEdit_SetHideScrollBars.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetLines(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetLines.Call(obj)
    return ret
}

func RichEdit_SetLines(obj uintptr, value uintptr) {
   richEdit_SetLines.Call(obj, value)
}

func RichEdit_GetMaxLength(obj uintptr) int32 {
    ret, _, _ := richEdit_GetMaxLength.Call(obj)
    return int32(ret)
}

func RichEdit_SetMaxLength(obj uintptr, value int32) {
   richEdit_SetMaxLength.Call(obj, uintptr(value))
}

func RichEdit_GetParentColor(obj uintptr) bool {
    ret, _, _ := richEdit_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetParentColor(obj uintptr, value bool) {
   richEdit_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := richEdit_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetParentCtl3D(obj uintptr, value bool) {
   richEdit_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetParentFont(obj uintptr) bool {
    ret, _, _ := richEdit_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetParentFont(obj uintptr, value bool) {
   richEdit_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := richEdit_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetParentShowHint(obj uintptr, value bool) {
   richEdit_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetPlainText(obj uintptr) bool {
    ret, _, _ := richEdit_GetPlainText.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetPlainText(obj uintptr, value bool) {
   richEdit_SetPlainText.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetPopupMenu.Call(obj)
    return ret
}

func RichEdit_SetPopupMenu(obj uintptr, value uintptr) {
   richEdit_SetPopupMenu.Call(obj, value)
}

func RichEdit_GetReadOnly(obj uintptr) bool {
    ret, _, _ := richEdit_GetReadOnly.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetReadOnly(obj uintptr, value bool) {
   richEdit_SetReadOnly.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetScrollBars(obj uintptr) TScrollStyle {
    ret, _, _ := richEdit_GetScrollBars.Call(obj)
    return TScrollStyle(ret)
}

func RichEdit_SetScrollBars(obj uintptr, value TScrollStyle) {
   richEdit_SetScrollBars.Call(obj, uintptr(value))
}

func RichEdit_GetShowHint(obj uintptr) bool {
    ret, _, _ := richEdit_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetShowHint(obj uintptr, value bool) {
   richEdit_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := richEdit_GetTabOrder.Call(obj)
    return uint16(ret)
}

func RichEdit_SetTabOrder(obj uintptr, value uint16) {
   richEdit_SetTabOrder.Call(obj, uintptr(value))
}

func RichEdit_GetTabStop(obj uintptr) bool {
    ret, _, _ := richEdit_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetTabStop(obj uintptr, value bool) {
   richEdit_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetVisible(obj uintptr) bool {
    ret, _, _ := richEdit_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetVisible(obj uintptr, value bool) {
   richEdit_SetVisible.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetWantTabs(obj uintptr) bool {
    ret, _, _ := richEdit_GetWantTabs.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetWantTabs(obj uintptr, value bool) {
   richEdit_SetWantTabs.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetWantReturns(obj uintptr) bool {
    ret, _, _ := richEdit_GetWantReturns.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetWantReturns(obj uintptr, value bool) {
   richEdit_SetWantReturns.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetWordWrap(obj uintptr) bool {
    ret, _, _ := richEdit_GetWordWrap.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetWordWrap(obj uintptr, value bool) {
   richEdit_SetWordWrap.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := richEdit_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func RichEdit_SetStyleElements(obj uintptr, value TStyleElements) {
   richEdit_SetStyleElements.Call(obj, uintptr(value))
}

func RichEdit_GetZoom(obj uintptr) int32 {
    ret, _, _ := richEdit_GetZoom.Call(obj)
    return int32(ret)
}

func RichEdit_SetZoom(obj uintptr, value int32) {
   richEdit_SetZoom.Call(obj, uintptr(value))
}

func RichEdit_SetOnChange(obj uintptr, fn interface{}) {
    richEdit_SetOnChange.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnClick(obj uintptr, fn interface{}) {
    richEdit_SetOnClick.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnDblClick(obj uintptr, fn interface{}) {
    richEdit_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnEnter(obj uintptr, fn interface{}) {
    richEdit_SetOnEnter.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnExit(obj uintptr, fn interface{}) {
    richEdit_SetOnExit.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnKeyDown(obj uintptr, fn interface{}) {
    richEdit_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnKeyPress(obj uintptr, fn interface{}) {
    richEdit_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnKeyUp(obj uintptr, fn interface{}) {
    richEdit_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnMouseDown(obj uintptr, fn interface{}) {
    richEdit_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnMouseEnter(obj uintptr, fn interface{}) {
    richEdit_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnMouseLeave(obj uintptr, fn interface{}) {
    richEdit_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnMouseMove(obj uintptr, fn interface{}) {
    richEdit_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnMouseUp(obj uintptr, fn interface{}) {
    richEdit_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func RichEdit_SetOnMouseWheel(obj uintptr, fn interface{}) {
    richEdit_SetOnMouseWheel.Call(obj, addEventToMap(fn))
}

func RichEdit_GetActiveLineNo(obj uintptr) uint32 {
    ret, _, _ := richEdit_GetActiveLineNo.Call(obj)
    return uint32(ret)
}

func RichEdit_GetDefAttributes(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetDefAttributes.Call(obj)
    return ret
}

func RichEdit_SetDefAttributes(obj uintptr, value uintptr) {
   richEdit_SetDefAttributes.Call(obj, value)
}

func RichEdit_GetSelAttributes(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetSelAttributes.Call(obj)
    return ret
}

func RichEdit_SetSelAttributes(obj uintptr, value uintptr) {
   richEdit_SetSelAttributes.Call(obj, value)
}

func RichEdit_GetPageRect(obj uintptr) TRect {
    var ret TRect
    richEdit_GetPageRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func RichEdit_SetPageRect(obj uintptr, value TRect) {
   richEdit_SetPageRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func RichEdit_GetParagraph(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetParagraph.Call(obj)
    return ret
}

func RichEdit_GetCaretPos(obj uintptr) TPoint {
    var ret TPoint
    richEdit_GetCaretPos.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func RichEdit_SetCaretPos(obj uintptr, value TPoint) {
   richEdit_SetCaretPos.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func RichEdit_GetModified(obj uintptr) bool {
    ret, _, _ := richEdit_GetModified.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetModified(obj uintptr, value bool) {
   richEdit_SetModified.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetSelLength(obj uintptr) int32 {
    ret, _, _ := richEdit_GetSelLength.Call(obj)
    return int32(ret)
}

func RichEdit_SetSelLength(obj uintptr, value int32) {
   richEdit_SetSelLength.Call(obj, uintptr(value))
}

func RichEdit_GetSelStart(obj uintptr) int32 {
    ret, _, _ := richEdit_GetSelStart.Call(obj)
    return int32(ret)
}

func RichEdit_SetSelStart(obj uintptr, value int32) {
   richEdit_SetSelStart.Call(obj, uintptr(value))
}

func RichEdit_GetSelText(obj uintptr) string {
    ret, _, _ := richEdit_GetSelText.Call(obj)
    return DStrToGoStr(ret)
}

func RichEdit_SetSelText(obj uintptr, value string) {
   richEdit_SetSelText.Call(obj, GoStrToDStr(value))
}

func RichEdit_GetText(obj uintptr) string {
    ret, _, _ := richEdit_GetText.Call(obj)
    return DStrToGoStr(ret)
}

func RichEdit_SetText(obj uintptr, value string) {
   richEdit_SetText.Call(obj, GoStrToDStr(value))
}

func RichEdit_GetTextHint(obj uintptr) string {
    ret, _, _ := richEdit_GetTextHint.Call(obj)
    return DStrToGoStr(ret)
}

func RichEdit_SetTextHint(obj uintptr, value string) {
   richEdit_SetTextHint.Call(obj, GoStrToDStr(value))
}

func RichEdit_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := richEdit_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetDoubleBuffered(obj uintptr, value bool) {
   richEdit_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetBrush(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetBrush.Call(obj)
    return ret
}

func RichEdit_GetControlCount(obj uintptr) int32 {
    ret, _, _ := richEdit_GetControlCount.Call(obj)
    return int32(ret)
}

func RichEdit_GetHandle(obj uintptr) HWND {
    ret, _, _ := richEdit_GetHandle.Call(obj)
    return HWND(ret)
}

func RichEdit_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := richEdit_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetParentDoubleBuffered(obj uintptr, value bool) {
   richEdit_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := richEdit_GetParentWindow.Call(obj)
    return HWND(ret)
}

func RichEdit_SetParentWindow(obj uintptr, value HWND) {
   richEdit_SetParentWindow.Call(obj, uintptr(value))
}

func RichEdit_GetAction(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetAction.Call(obj)
    return ret
}

func RichEdit_SetAction(obj uintptr, value uintptr) {
   richEdit_SetAction.Call(obj, value)
}

func RichEdit_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    richEdit_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func RichEdit_SetBoundsRect(obj uintptr, value TRect) {
   richEdit_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func RichEdit_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := richEdit_GetClientHeight.Call(obj)
    return int32(ret)
}

func RichEdit_SetClientHeight(obj uintptr, value int32) {
   richEdit_SetClientHeight.Call(obj, uintptr(value))
}

func RichEdit_GetClientRect(obj uintptr) TRect {
    var ret TRect
    richEdit_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func RichEdit_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := richEdit_GetClientWidth.Call(obj)
    return int32(ret)
}

func RichEdit_SetClientWidth(obj uintptr, value int32) {
   richEdit_SetClientWidth.Call(obj, uintptr(value))
}

func RichEdit_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := richEdit_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func RichEdit_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := richEdit_GetExplicitTop.Call(obj)
    return int32(ret)
}

func RichEdit_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := richEdit_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func RichEdit_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := richEdit_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func RichEdit_GetParent(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetParent.Call(obj)
    return ret
}

func RichEdit_SetParent(obj uintptr, value uintptr) {
   richEdit_SetParent.Call(obj, value)
}

func RichEdit_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := richEdit_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func RichEdit_SetAlignWithMargins(obj uintptr, value bool) {
   richEdit_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func RichEdit_GetLeft(obj uintptr) int32 {
    ret, _, _ := richEdit_GetLeft.Call(obj)
    return int32(ret)
}

func RichEdit_SetLeft(obj uintptr, value int32) {
   richEdit_SetLeft.Call(obj, uintptr(value))
}

func RichEdit_GetTop(obj uintptr) int32 {
    ret, _, _ := richEdit_GetTop.Call(obj)
    return int32(ret)
}

func RichEdit_SetTop(obj uintptr, value int32) {
   richEdit_SetTop.Call(obj, uintptr(value))
}

func RichEdit_GetWidth(obj uintptr) int32 {
    ret, _, _ := richEdit_GetWidth.Call(obj)
    return int32(ret)
}

func RichEdit_SetWidth(obj uintptr, value int32) {
   richEdit_SetWidth.Call(obj, uintptr(value))
}

func RichEdit_GetHeight(obj uintptr) int32 {
    ret, _, _ := richEdit_GetHeight.Call(obj)
    return int32(ret)
}

func RichEdit_SetHeight(obj uintptr, value int32) {
   richEdit_SetHeight.Call(obj, uintptr(value))
}

func RichEdit_GetCursor(obj uintptr) TCursor {
    ret, _, _ := richEdit_GetCursor.Call(obj)
    return TCursor(ret)
}

func RichEdit_SetCursor(obj uintptr, value TCursor) {
   richEdit_SetCursor.Call(obj, uintptr(value))
}

func RichEdit_GetHint(obj uintptr) string {
    ret, _, _ := richEdit_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func RichEdit_SetHint(obj uintptr, value string) {
   richEdit_SetHint.Call(obj, GoStrToDStr(value))
}

func RichEdit_GetMargins(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetMargins.Call(obj)
    return ret
}

func RichEdit_SetMargins(obj uintptr, value uintptr) {
   richEdit_SetMargins.Call(obj, value)
}

func RichEdit_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := richEdit_GetComponentCount.Call(obj)
    return int32(ret)
}

func RichEdit_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := richEdit_GetComponentIndex.Call(obj)
    return int32(ret)
}

func RichEdit_SetComponentIndex(obj uintptr, value int32) {
   richEdit_SetComponentIndex.Call(obj, uintptr(value))
}

func RichEdit_GetOwner(obj uintptr) uintptr {
    ret, _, _ := richEdit_GetOwner.Call(obj)
    return ret
}

func RichEdit_GetName(obj uintptr) string {
    ret, _, _ := richEdit_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func RichEdit_SetName(obj uintptr, value string) {
   richEdit_SetName.Call(obj, GoStrToDStr(value))
}

func RichEdit_GetTag(obj uintptr) int {
    ret, _, _ := richEdit_GetTag.Call(obj)
    return int(ret)
}

func RichEdit_SetTag(obj uintptr, value int) {
   richEdit_SetTag.Call(obj, uintptr(value))
}

func RichEdit_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := richEdit_GetControls.Call(obj, uintptr(Index))
    return ret
}

func RichEdit_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := richEdit_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TTrackBar ---------------------------

func TrackBar_Create(obj uintptr) uintptr {
    ret, _, _ := trackBar_Create.Call(obj)
    return ret
}

func TrackBar_Free(obj uintptr) {
    trackBar_Free.Call(obj)
}

func TrackBar_SetTick(obj uintptr, Value int32)  {
    trackBar_SetTick.Call(obj, uintptr(Value) )
}

func TrackBar_CanFocus(obj uintptr) bool {
    ret, _, _ := trackBar_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_FlipChildren(obj uintptr, AllLevels bool)  {
    trackBar_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func TrackBar_Focused(obj uintptr) bool {
    ret, _, _ := trackBar_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_HandleAllocated(obj uintptr) bool {
    ret, _, _ := trackBar_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_Invalidate(obj uintptr)  {
    trackBar_Invalidate.Call(obj)
}

func TrackBar_Realign(obj uintptr)  {
    trackBar_Realign.Call(obj)
}

func TrackBar_Repaint(obj uintptr)  {
    trackBar_Repaint.Call(obj)
}

func TrackBar_ScaleBy(obj uintptr, M int32, D int32)  {
    trackBar_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func TrackBar_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    trackBar_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func TrackBar_SetFocus(obj uintptr)  {
    trackBar_SetFocus.Call(obj)
}

func TrackBar_Update(obj uintptr)  {
    trackBar_Update.Call(obj)
}

func TrackBar_BringToFront(obj uintptr)  {
    trackBar_BringToFront.Call(obj)
}

func TrackBar_HasParent(obj uintptr) bool {
    ret, _, _ := trackBar_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_Hide(obj uintptr)  {
    trackBar_Hide.Call(obj)
}

func TrackBar_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := trackBar_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func TrackBar_Refresh(obj uintptr)  {
    trackBar_Refresh.Call(obj)
}

func TrackBar_SendToBack(obj uintptr)  {
    trackBar_SendToBack.Call(obj)
}

func TrackBar_Show(obj uintptr)  {
    trackBar_Show.Call(obj)
}

func TrackBar_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := trackBar_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func TrackBar_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := trackBar_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func TrackBar_GetNamePath(obj uintptr) string {
    ret, _, _ := trackBar_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func TrackBar_Assign(obj uintptr, Source uintptr)  {
    trackBar_Assign.Call(obj, Source )
}

func TrackBar_ClassName(obj uintptr) string {
    ret, _, _ := trackBar_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func TrackBar_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := trackBar_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func TrackBar_GetHashCode(obj uintptr) int32 {
    ret, _, _ := trackBar_GetHashCode.Call(obj)
    return int32(ret)
}

func TrackBar_ToString(obj uintptr) string {
    ret, _, _ := trackBar_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func TrackBar_GetAlign(obj uintptr) TAlign {
    ret, _, _ := trackBar_GetAlign.Call(obj)
    return TAlign(ret)
}

func TrackBar_SetAlign(obj uintptr, value TAlign) {
   trackBar_SetAlign.Call(obj, uintptr(value))
}

func TrackBar_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := trackBar_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func TrackBar_SetAnchors(obj uintptr, value TAnchors) {
   trackBar_SetAnchors.Call(obj, uintptr(value))
}

func TrackBar_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := trackBar_GetBorderWidth.Call(obj)
    return int32(ret)
}

func TrackBar_SetBorderWidth(obj uintptr, value int32) {
   trackBar_SetBorderWidth.Call(obj, uintptr(value))
}

func TrackBar_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := trackBar_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetDoubleBuffered(obj uintptr, value bool) {
   trackBar_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetEnabled(obj uintptr) bool {
    ret, _, _ := trackBar_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetEnabled(obj uintptr, value bool) {
   trackBar_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetLineSize(obj uintptr) int32 {
    ret, _, _ := trackBar_GetLineSize.Call(obj)
    return int32(ret)
}

func TrackBar_SetLineSize(obj uintptr, value int32) {
   trackBar_SetLineSize.Call(obj, uintptr(value))
}

func TrackBar_GetMax(obj uintptr) int32 {
    ret, _, _ := trackBar_GetMax.Call(obj)
    return int32(ret)
}

func TrackBar_SetMax(obj uintptr, value int32) {
   trackBar_SetMax.Call(obj, uintptr(value))
}

func TrackBar_GetMin(obj uintptr) int32 {
    ret, _, _ := trackBar_GetMin.Call(obj)
    return int32(ret)
}

func TrackBar_SetMin(obj uintptr, value int32) {
   trackBar_SetMin.Call(obj, uintptr(value))
}

func TrackBar_GetOrientation(obj uintptr) TTrackBarOrientation {
    ret, _, _ := trackBar_GetOrientation.Call(obj)
    return TTrackBarOrientation(ret)
}

func TrackBar_SetOrientation(obj uintptr, value TTrackBarOrientation) {
   trackBar_SetOrientation.Call(obj, uintptr(value))
}

func TrackBar_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := trackBar_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetParentCtl3D(obj uintptr, value bool) {
   trackBar_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := trackBar_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetParentDoubleBuffered(obj uintptr, value bool) {
   trackBar_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := trackBar_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetParentShowHint(obj uintptr, value bool) {
   trackBar_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetPageSize(obj uintptr) int32 {
    ret, _, _ := trackBar_GetPageSize.Call(obj)
    return int32(ret)
}

func TrackBar_SetPageSize(obj uintptr, value int32) {
   trackBar_SetPageSize.Call(obj, uintptr(value))
}

func TrackBar_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := trackBar_GetPopupMenu.Call(obj)
    return ret
}

func TrackBar_SetPopupMenu(obj uintptr, value uintptr) {
   trackBar_SetPopupMenu.Call(obj, value)
}

func TrackBar_GetFrequency(obj uintptr) int32 {
    ret, _, _ := trackBar_GetFrequency.Call(obj)
    return int32(ret)
}

func TrackBar_SetFrequency(obj uintptr, value int32) {
   trackBar_SetFrequency.Call(obj, uintptr(value))
}

func TrackBar_GetPosition(obj uintptr) int32 {
    ret, _, _ := trackBar_GetPosition.Call(obj)
    return int32(ret)
}

func TrackBar_SetPosition(obj uintptr, value int32) {
   trackBar_SetPosition.Call(obj, uintptr(value))
}

func TrackBar_GetPositionToolTip(obj uintptr) TPositionToolTip {
    ret, _, _ := trackBar_GetPositionToolTip.Call(obj)
    return TPositionToolTip(ret)
}

func TrackBar_SetPositionToolTip(obj uintptr, value TPositionToolTip) {
   trackBar_SetPositionToolTip.Call(obj, uintptr(value))
}

func TrackBar_GetSliderVisible(obj uintptr) bool {
    ret, _, _ := trackBar_GetSliderVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetSliderVisible(obj uintptr, value bool) {
   trackBar_SetSliderVisible.Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetSelEnd(obj uintptr) int32 {
    ret, _, _ := trackBar_GetSelEnd.Call(obj)
    return int32(ret)
}

func TrackBar_SetSelEnd(obj uintptr, value int32) {
   trackBar_SetSelEnd.Call(obj, uintptr(value))
}

func TrackBar_GetSelStart(obj uintptr) int32 {
    ret, _, _ := trackBar_GetSelStart.Call(obj)
    return int32(ret)
}

func TrackBar_SetSelStart(obj uintptr, value int32) {
   trackBar_SetSelStart.Call(obj, uintptr(value))
}

func TrackBar_GetShowHint(obj uintptr) bool {
    ret, _, _ := trackBar_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetShowHint(obj uintptr, value bool) {
   trackBar_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetShowSelRange(obj uintptr) bool {
    ret, _, _ := trackBar_GetShowSelRange.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetShowSelRange(obj uintptr, value bool) {
   trackBar_SetShowSelRange.Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := trackBar_GetTabOrder.Call(obj)
    return uint16(ret)
}

func TrackBar_SetTabOrder(obj uintptr, value uint16) {
   trackBar_SetTabOrder.Call(obj, uintptr(value))
}

func TrackBar_GetTabStop(obj uintptr) bool {
    ret, _, _ := trackBar_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetTabStop(obj uintptr, value bool) {
   trackBar_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetThumbLength(obj uintptr) int32 {
    ret, _, _ := trackBar_GetThumbLength.Call(obj)
    return int32(ret)
}

func TrackBar_SetThumbLength(obj uintptr, value int32) {
   trackBar_SetThumbLength.Call(obj, uintptr(value))
}

func TrackBar_GetTickMarks(obj uintptr) TTickMark {
    ret, _, _ := trackBar_GetTickMarks.Call(obj)
    return TTickMark(ret)
}

func TrackBar_SetTickMarks(obj uintptr, value TTickMark) {
   trackBar_SetTickMarks.Call(obj, uintptr(value))
}

func TrackBar_GetTickStyle(obj uintptr) TTickStyle {
    ret, _, _ := trackBar_GetTickStyle.Call(obj)
    return TTickStyle(ret)
}

func TrackBar_SetTickStyle(obj uintptr, value TTickStyle) {
   trackBar_SetTickStyle.Call(obj, uintptr(value))
}

func TrackBar_GetVisible(obj uintptr) bool {
    ret, _, _ := trackBar_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetVisible(obj uintptr, value bool) {
   trackBar_SetVisible.Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := trackBar_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func TrackBar_SetStyleElements(obj uintptr, value TStyleElements) {
   trackBar_SetStyleElements.Call(obj, uintptr(value))
}

func TrackBar_SetOnChange(obj uintptr, fn interface{}) {
    trackBar_SetOnChange.Call(obj, addEventToMap(fn))
}

func TrackBar_SetOnEnter(obj uintptr, fn interface{}) {
    trackBar_SetOnEnter.Call(obj, addEventToMap(fn))
}

func TrackBar_SetOnExit(obj uintptr, fn interface{}) {
    trackBar_SetOnExit.Call(obj, addEventToMap(fn))
}

func TrackBar_SetOnKeyDown(obj uintptr, fn interface{}) {
    trackBar_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func TrackBar_SetOnKeyPress(obj uintptr, fn interface{}) {
    trackBar_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func TrackBar_SetOnKeyUp(obj uintptr, fn interface{}) {
    trackBar_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func TrackBar_GetBrush(obj uintptr) uintptr {
    ret, _, _ := trackBar_GetBrush.Call(obj)
    return ret
}

func TrackBar_GetControlCount(obj uintptr) int32 {
    ret, _, _ := trackBar_GetControlCount.Call(obj)
    return int32(ret)
}

func TrackBar_GetHandle(obj uintptr) HWND {
    ret, _, _ := trackBar_GetHandle.Call(obj)
    return HWND(ret)
}

func TrackBar_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := trackBar_GetParentWindow.Call(obj)
    return HWND(ret)
}

func TrackBar_SetParentWindow(obj uintptr, value HWND) {
   trackBar_SetParentWindow.Call(obj, uintptr(value))
}

func TrackBar_GetAction(obj uintptr) uintptr {
    ret, _, _ := trackBar_GetAction.Call(obj)
    return ret
}

func TrackBar_SetAction(obj uintptr, value uintptr) {
   trackBar_SetAction.Call(obj, value)
}

func TrackBar_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := trackBar_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func TrackBar_SetBiDiMode(obj uintptr, value TBiDiMode) {
   trackBar_SetBiDiMode.Call(obj, uintptr(value))
}

func TrackBar_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    trackBar_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func TrackBar_SetBoundsRect(obj uintptr, value TRect) {
   trackBar_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func TrackBar_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := trackBar_GetClientHeight.Call(obj)
    return int32(ret)
}

func TrackBar_SetClientHeight(obj uintptr, value int32) {
   trackBar_SetClientHeight.Call(obj, uintptr(value))
}

func TrackBar_GetClientRect(obj uintptr) TRect {
    var ret TRect
    trackBar_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func TrackBar_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := trackBar_GetClientWidth.Call(obj)
    return int32(ret)
}

func TrackBar_SetClientWidth(obj uintptr, value int32) {
   trackBar_SetClientWidth.Call(obj, uintptr(value))
}

func TrackBar_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := trackBar_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func TrackBar_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := trackBar_GetExplicitTop.Call(obj)
    return int32(ret)
}

func TrackBar_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := trackBar_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func TrackBar_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := trackBar_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func TrackBar_GetParent(obj uintptr) uintptr {
    ret, _, _ := trackBar_GetParent.Call(obj)
    return ret
}

func TrackBar_SetParent(obj uintptr, value uintptr) {
   trackBar_SetParent.Call(obj, value)
}

func TrackBar_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := trackBar_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func TrackBar_SetAlignWithMargins(obj uintptr, value bool) {
   trackBar_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func TrackBar_GetLeft(obj uintptr) int32 {
    ret, _, _ := trackBar_GetLeft.Call(obj)
    return int32(ret)
}

func TrackBar_SetLeft(obj uintptr, value int32) {
   trackBar_SetLeft.Call(obj, uintptr(value))
}

func TrackBar_GetTop(obj uintptr) int32 {
    ret, _, _ := trackBar_GetTop.Call(obj)
    return int32(ret)
}

func TrackBar_SetTop(obj uintptr, value int32) {
   trackBar_SetTop.Call(obj, uintptr(value))
}

func TrackBar_GetWidth(obj uintptr) int32 {
    ret, _, _ := trackBar_GetWidth.Call(obj)
    return int32(ret)
}

func TrackBar_SetWidth(obj uintptr, value int32) {
   trackBar_SetWidth.Call(obj, uintptr(value))
}

func TrackBar_GetHeight(obj uintptr) int32 {
    ret, _, _ := trackBar_GetHeight.Call(obj)
    return int32(ret)
}

func TrackBar_SetHeight(obj uintptr, value int32) {
   trackBar_SetHeight.Call(obj, uintptr(value))
}

func TrackBar_GetCursor(obj uintptr) TCursor {
    ret, _, _ := trackBar_GetCursor.Call(obj)
    return TCursor(ret)
}

func TrackBar_SetCursor(obj uintptr, value TCursor) {
   trackBar_SetCursor.Call(obj, uintptr(value))
}

func TrackBar_GetHint(obj uintptr) string {
    ret, _, _ := trackBar_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func TrackBar_SetHint(obj uintptr, value string) {
   trackBar_SetHint.Call(obj, GoStrToDStr(value))
}

func TrackBar_GetMargins(obj uintptr) uintptr {
    ret, _, _ := trackBar_GetMargins.Call(obj)
    return ret
}

func TrackBar_SetMargins(obj uintptr, value uintptr) {
   trackBar_SetMargins.Call(obj, value)
}

func TrackBar_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := trackBar_GetComponentCount.Call(obj)
    return int32(ret)
}

func TrackBar_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := trackBar_GetComponentIndex.Call(obj)
    return int32(ret)
}

func TrackBar_SetComponentIndex(obj uintptr, value int32) {
   trackBar_SetComponentIndex.Call(obj, uintptr(value))
}

func TrackBar_GetOwner(obj uintptr) uintptr {
    ret, _, _ := trackBar_GetOwner.Call(obj)
    return ret
}

func TrackBar_GetName(obj uintptr) string {
    ret, _, _ := trackBar_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func TrackBar_SetName(obj uintptr, value string) {
   trackBar_SetName.Call(obj, GoStrToDStr(value))
}

func TrackBar_GetTag(obj uintptr) int {
    ret, _, _ := trackBar_GetTag.Call(obj)
    return int(ret)
}

func TrackBar_SetTag(obj uintptr, value int) {
   trackBar_SetTag.Call(obj, uintptr(value))
}

func TrackBar_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := trackBar_GetControls.Call(obj, uintptr(Index))
    return ret
}

func TrackBar_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := trackBar_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TImageList ---------------------------

func ImageList_Create(obj uintptr) uintptr {
    ret, _, _ := imageList_Create.Call(obj)
    return ret
}

func ImageList_Free(obj uintptr) {
    imageList_Free.Call(obj)
}

func ImageList_GetHotSpot(obj uintptr) TPoint {
    var ret TPoint
    imageList_GetHotSpot.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ImageList_Assign(obj uintptr, Source uintptr)  {
    imageList_Assign.Call(obj, Source )
}

func ImageList_Add(obj uintptr, Image uintptr, Mask uintptr) int32 {
    ret, _, _ := imageList_Add.Call(obj, Image , Mask )
    return int32(ret)
}

func ImageList_AddIcon(obj uintptr, Image uintptr) int32 {
    ret, _, _ := imageList_AddIcon.Call(obj, Image )
    return int32(ret)
}

func ImageList_AddImage(obj uintptr, Value uintptr, Index int32) int32 {
    ret, _, _ := imageList_AddImage.Call(obj, Value , uintptr(Index) )
    return int32(ret)
}

func ImageList_AddImages(obj uintptr, Value uintptr)  {
    imageList_AddImages.Call(obj, Value )
}

func ImageList_AddMasked(obj uintptr, Image uintptr, MaskColor TColor) int32 {
    ret, _, _ := imageList_AddMasked.Call(obj, Image , uintptr(MaskColor) )
    return int32(ret)
}

func ImageList_Clear(obj uintptr)  {
    imageList_Clear.Call(obj)
}

func ImageList_Delete(obj uintptr, Index int32)  {
    imageList_Delete.Call(obj, uintptr(Index) )
}

func ImageList_FileLoad(obj uintptr, ResType TResType, Name string, MaskColor TColor) bool {
    ret, _, _ := imageList_FileLoad.Call(obj, uintptr(ResType) , GoStrToDStr(Name) , uintptr(MaskColor) )
    return DBoolToGoBool(ret)
}

func ImageList_GetBitmap(obj uintptr, Index int32, Image uintptr) bool {
    ret, _, _ := imageList_GetBitmap.Call(obj, uintptr(Index) , Image )
    return DBoolToGoBool(ret)
}

func ImageList_GetImageBitmap(obj uintptr) HBITMAP {
    ret, _, _ := imageList_GetImageBitmap.Call(obj)
    return HBITMAP(ret)
}

func ImageList_GetMaskBitmap(obj uintptr) HBITMAP {
    ret, _, _ := imageList_GetMaskBitmap.Call(obj)
    return HBITMAP(ret)
}

func ImageList_GetResource(obj uintptr, ResType TResType, Name string, Width int32, LoadFlags TLoadResources, MaskColor TColor) bool {
    ret, _, _ := imageList_GetResource.Call(obj, uintptr(ResType) , GoStrToDStr(Name) , uintptr(Width) , uintptr(LoadFlags) , uintptr(MaskColor) )
    return DBoolToGoBool(ret)
}

func ImageList_HandleAllocated(obj uintptr) bool {
    ret, _, _ := imageList_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func ImageList_Insert(obj uintptr, Index int32, Image uintptr, Mask uintptr)  {
    imageList_Insert.Call(obj, uintptr(Index) , Image , Mask )
}

func ImageList_InsertIcon(obj uintptr, Index int32, Image uintptr)  {
    imageList_InsertIcon.Call(obj, uintptr(Index) , Image )
}

func ImageList_InsertMasked(obj uintptr, Index int32, Image uintptr, MaskColor TColor)  {
    imageList_InsertMasked.Call(obj, uintptr(Index) , Image , uintptr(MaskColor) )
}

func ImageList_Move(obj uintptr, CurIndex int32, NewIndex int32)  {
    imageList_Move.Call(obj, uintptr(CurIndex) , uintptr(NewIndex) )
}

func ImageList_Overlay(obj uintptr, ImageIndex int32, Overlay uint8) bool {
    ret, _, _ := imageList_Overlay.Call(obj, uintptr(ImageIndex) , uintptr(Overlay) )
    return DBoolToGoBool(ret)
}

func ImageList_ResourceLoad(obj uintptr, ResType TResType, Name string, MaskColor TColor) bool {
    ret, _, _ := imageList_ResourceLoad.Call(obj, uintptr(ResType) , GoStrToDStr(Name) , uintptr(MaskColor) )
    return DBoolToGoBool(ret)
}

func ImageList_ResInstLoad(obj uintptr, Instance uintptr, ResType TResType, Name string, MaskColor TColor) bool {
    ret, _, _ := imageList_ResInstLoad.Call(obj, Instance , uintptr(ResType) , GoStrToDStr(Name) , uintptr(MaskColor) )
    return DBoolToGoBool(ret)
}

func ImageList_Replace(obj uintptr, Index int32, Image uintptr, Mask uintptr)  {
    imageList_Replace.Call(obj, uintptr(Index) , Image , Mask )
}

func ImageList_ReplaceIcon(obj uintptr, Index int32, Image uintptr)  {
    imageList_ReplaceIcon.Call(obj, uintptr(Index) , Image )
}

func ImageList_ReplaceMasked(obj uintptr, Index int32, NewImage uintptr, MaskColor TColor)  {
    imageList_ReplaceMasked.Call(obj, uintptr(Index) , NewImage , uintptr(MaskColor) )
}

func ImageList_SetSize(obj uintptr, AWidth int32, AHeight int32)  {
    imageList_SetSize.Call(obj, uintptr(AWidth) , uintptr(AHeight) )
}

func ImageList_BeginUpdate(obj uintptr)  {
    imageList_BeginUpdate.Call(obj)
}

func ImageList_EndUpdate(obj uintptr)  {
    imageList_EndUpdate.Call(obj)
}

func ImageList_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := imageList_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ImageList_GetNamePath(obj uintptr) string {
    ret, _, _ := imageList_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ImageList_HasParent(obj uintptr) bool {
    ret, _, _ := imageList_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ImageList_ClassName(obj uintptr) string {
    ret, _, _ := imageList_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ImageList_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := imageList_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ImageList_GetHashCode(obj uintptr) int32 {
    ret, _, _ := imageList_GetHashCode.Call(obj)
    return int32(ret)
}

func ImageList_ToString(obj uintptr) string {
    ret, _, _ := imageList_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ImageList_GetBlendColor(obj uintptr) TColor {
    ret, _, _ := imageList_GetBlendColor.Call(obj)
    return TColor(ret)
}

func ImageList_SetBlendColor(obj uintptr, value TColor) {
   imageList_SetBlendColor.Call(obj, uintptr(value))
}

func ImageList_GetBkColor(obj uintptr) TColor {
    ret, _, _ := imageList_GetBkColor.Call(obj)
    return TColor(ret)
}

func ImageList_SetBkColor(obj uintptr, value TColor) {
   imageList_SetBkColor.Call(obj, uintptr(value))
}

func ImageList_GetAllocBy(obj uintptr) int32 {
    ret, _, _ := imageList_GetAllocBy.Call(obj)
    return int32(ret)
}

func ImageList_SetAllocBy(obj uintptr, value int32) {
   imageList_SetAllocBy.Call(obj, uintptr(value))
}

func ImageList_GetColorDepth(obj uintptr) TColorDepth {
    ret, _, _ := imageList_GetColorDepth.Call(obj)
    return TColorDepth(ret)
}

func ImageList_SetColorDepth(obj uintptr, value TColorDepth) {
   imageList_SetColorDepth.Call(obj, uintptr(value))
}

func ImageList_GetDrawingStyle(obj uintptr) TDrawingStyle {
    ret, _, _ := imageList_GetDrawingStyle.Call(obj)
    return TDrawingStyle(ret)
}

func ImageList_SetDrawingStyle(obj uintptr, value TDrawingStyle) {
   imageList_SetDrawingStyle.Call(obj, uintptr(value))
}

func ImageList_GetGrayscaleFactor(obj uintptr) uint8 {
    ret, _, _ := imageList_GetGrayscaleFactor.Call(obj)
    return uint8(ret)
}

func ImageList_SetGrayscaleFactor(obj uintptr, value uint8) {
   imageList_SetGrayscaleFactor.Call(obj, uintptr(value))
}

func ImageList_GetHeight(obj uintptr) int32 {
    ret, _, _ := imageList_GetHeight.Call(obj)
    return int32(ret)
}

func ImageList_SetHeight(obj uintptr, value int32) {
   imageList_SetHeight.Call(obj, uintptr(value))
}

func ImageList_GetImageType(obj uintptr) TImageType {
    ret, _, _ := imageList_GetImageType.Call(obj)
    return TImageType(ret)
}

func ImageList_SetImageType(obj uintptr, value TImageType) {
   imageList_SetImageType.Call(obj, uintptr(value))
}

func ImageList_GetMasked(obj uintptr) bool {
    ret, _, _ := imageList_GetMasked.Call(obj)
    return DBoolToGoBool(ret)
}

func ImageList_SetMasked(obj uintptr, value bool) {
   imageList_SetMasked.Call(obj, GoBoolToDBool(value))
}

func ImageList_SetOnChange(obj uintptr, fn interface{}) {
    imageList_SetOnChange.Call(obj, addEventToMap(fn))
}

func ImageList_GetShareImages(obj uintptr) bool {
    ret, _, _ := imageList_GetShareImages.Call(obj)
    return DBoolToGoBool(ret)
}

func ImageList_SetShareImages(obj uintptr, value bool) {
   imageList_SetShareImages.Call(obj, GoBoolToDBool(value))
}

func ImageList_GetWidth(obj uintptr) int32 {
    ret, _, _ := imageList_GetWidth.Call(obj)
    return int32(ret)
}

func ImageList_SetWidth(obj uintptr, value int32) {
   imageList_SetWidth.Call(obj, uintptr(value))
}

func ImageList_GetHandle(obj uintptr) uintptr {
    ret, _, _ := imageList_GetHandle.Call(obj)
    return ret
}

func ImageList_SetHandle(obj uintptr, value uintptr) {
   imageList_SetHandle.Call(obj, value)
}

func ImageList_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := imageList_GetComponentCount.Call(obj)
    return int32(ret)
}

func ImageList_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := imageList_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ImageList_SetComponentIndex(obj uintptr, value int32) {
   imageList_SetComponentIndex.Call(obj, uintptr(value))
}

func ImageList_GetOwner(obj uintptr) uintptr {
    ret, _, _ := imageList_GetOwner.Call(obj)
    return ret
}

func ImageList_GetName(obj uintptr) string {
    ret, _, _ := imageList_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ImageList_SetName(obj uintptr, value string) {
   imageList_SetName.Call(obj, GoStrToDStr(value))
}

func ImageList_GetTag(obj uintptr) int {
    ret, _, _ := imageList_GetTag.Call(obj)
    return int(ret)
}

func ImageList_SetTag(obj uintptr, value int) {
   imageList_SetTag.Call(obj, uintptr(value))
}

func ImageList_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := imageList_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TUpDown ---------------------------

func UpDown_Create(obj uintptr) uintptr {
    ret, _, _ := upDown_Create.Call(obj)
    return ret
}

func UpDown_Free(obj uintptr) {
    upDown_Free.Call(obj)
}

func UpDown_CanFocus(obj uintptr) bool {
    ret, _, _ := upDown_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_FlipChildren(obj uintptr, AllLevels bool)  {
    upDown_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func UpDown_Focused(obj uintptr) bool {
    ret, _, _ := upDown_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_HandleAllocated(obj uintptr) bool {
    ret, _, _ := upDown_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_Invalidate(obj uintptr)  {
    upDown_Invalidate.Call(obj)
}

func UpDown_Realign(obj uintptr)  {
    upDown_Realign.Call(obj)
}

func UpDown_Repaint(obj uintptr)  {
    upDown_Repaint.Call(obj)
}

func UpDown_ScaleBy(obj uintptr, M int32, D int32)  {
    upDown_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func UpDown_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    upDown_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func UpDown_SetFocus(obj uintptr)  {
    upDown_SetFocus.Call(obj)
}

func UpDown_Update(obj uintptr)  {
    upDown_Update.Call(obj)
}

func UpDown_BringToFront(obj uintptr)  {
    upDown_BringToFront.Call(obj)
}

func UpDown_HasParent(obj uintptr) bool {
    ret, _, _ := upDown_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_Hide(obj uintptr)  {
    upDown_Hide.Call(obj)
}

func UpDown_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := upDown_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func UpDown_Refresh(obj uintptr)  {
    upDown_Refresh.Call(obj)
}

func UpDown_SendToBack(obj uintptr)  {
    upDown_SendToBack.Call(obj)
}

func UpDown_Show(obj uintptr)  {
    upDown_Show.Call(obj)
}

func UpDown_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := upDown_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func UpDown_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := upDown_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func UpDown_GetNamePath(obj uintptr) string {
    ret, _, _ := upDown_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func UpDown_Assign(obj uintptr, Source uintptr)  {
    upDown_Assign.Call(obj, Source )
}

func UpDown_ClassName(obj uintptr) string {
    ret, _, _ := upDown_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func UpDown_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := upDown_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func UpDown_GetHashCode(obj uintptr) int32 {
    ret, _, _ := upDown_GetHashCode.Call(obj)
    return int32(ret)
}

func UpDown_ToString(obj uintptr) string {
    ret, _, _ := upDown_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func UpDown_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := upDown_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func UpDown_SetAnchors(obj uintptr, value TAnchors) {
   upDown_SetAnchors.Call(obj, uintptr(value))
}

func UpDown_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := upDown_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_SetDoubleBuffered(obj uintptr, value bool) {
   upDown_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func UpDown_GetEnabled(obj uintptr) bool {
    ret, _, _ := upDown_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_SetEnabled(obj uintptr, value bool) {
   upDown_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func UpDown_GetHint(obj uintptr) string {
    ret, _, _ := upDown_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func UpDown_SetHint(obj uintptr, value string) {
   upDown_SetHint.Call(obj, GoStrToDStr(value))
}

func UpDown_GetMin(obj uintptr) int32 {
    ret, _, _ := upDown_GetMin.Call(obj)
    return int32(ret)
}

func UpDown_SetMin(obj uintptr, value int32) {
   upDown_SetMin.Call(obj, uintptr(value))
}

func UpDown_GetMax(obj uintptr) int32 {
    ret, _, _ := upDown_GetMax.Call(obj)
    return int32(ret)
}

func UpDown_SetMax(obj uintptr, value int32) {
   upDown_SetMax.Call(obj, uintptr(value))
}

func UpDown_GetOrientation(obj uintptr) TUDOrientation {
    ret, _, _ := upDown_GetOrientation.Call(obj)
    return TUDOrientation(ret)
}

func UpDown_SetOrientation(obj uintptr, value TUDOrientation) {
   upDown_SetOrientation.Call(obj, uintptr(value))
}

func UpDown_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := upDown_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_SetParentDoubleBuffered(obj uintptr, value bool) {
   upDown_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func UpDown_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := upDown_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_SetParentShowHint(obj uintptr, value bool) {
   upDown_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func UpDown_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := upDown_GetPopupMenu.Call(obj)
    return ret
}

func UpDown_SetPopupMenu(obj uintptr, value uintptr) {
   upDown_SetPopupMenu.Call(obj, value)
}

func UpDown_GetPosition(obj uintptr) int32 {
    ret, _, _ := upDown_GetPosition.Call(obj)
    return int32(ret)
}

func UpDown_SetPosition(obj uintptr, value int32) {
   upDown_SetPosition.Call(obj, uintptr(value))
}

func UpDown_GetShowHint(obj uintptr) bool {
    ret, _, _ := upDown_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_SetShowHint(obj uintptr, value bool) {
   upDown_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func UpDown_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := upDown_GetTabOrder.Call(obj)
    return uint16(ret)
}

func UpDown_SetTabOrder(obj uintptr, value uint16) {
   upDown_SetTabOrder.Call(obj, uintptr(value))
}

func UpDown_GetTabStop(obj uintptr) bool {
    ret, _, _ := upDown_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_SetTabStop(obj uintptr, value bool) {
   upDown_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func UpDown_GetVisible(obj uintptr) bool {
    ret, _, _ := upDown_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_SetVisible(obj uintptr, value bool) {
   upDown_SetVisible.Call(obj, GoBoolToDBool(value))
}

func UpDown_GetWrap(obj uintptr) bool {
    ret, _, _ := upDown_GetWrap.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_SetWrap(obj uintptr, value bool) {
   upDown_SetWrap.Call(obj, GoBoolToDBool(value))
}

func UpDown_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := upDown_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func UpDown_SetStyleElements(obj uintptr, value TStyleElements) {
   upDown_SetStyleElements.Call(obj, uintptr(value))
}

func UpDown_SetOnClick(obj uintptr, fn interface{}) {
    upDown_SetOnClick.Call(obj, addEventToMap(fn))
}

func UpDown_SetOnEnter(obj uintptr, fn interface{}) {
    upDown_SetOnEnter.Call(obj, addEventToMap(fn))
}

func UpDown_SetOnExit(obj uintptr, fn interface{}) {
    upDown_SetOnExit.Call(obj, addEventToMap(fn))
}

func UpDown_SetOnMouseDown(obj uintptr, fn interface{}) {
    upDown_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func UpDown_SetOnMouseEnter(obj uintptr, fn interface{}) {
    upDown_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func UpDown_SetOnMouseLeave(obj uintptr, fn interface{}) {
    upDown_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func UpDown_SetOnMouseMove(obj uintptr, fn interface{}) {
    upDown_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func UpDown_SetOnMouseUp(obj uintptr, fn interface{}) {
    upDown_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func UpDown_GetBrush(obj uintptr) uintptr {
    ret, _, _ := upDown_GetBrush.Call(obj)
    return ret
}

func UpDown_GetControlCount(obj uintptr) int32 {
    ret, _, _ := upDown_GetControlCount.Call(obj)
    return int32(ret)
}

func UpDown_GetHandle(obj uintptr) HWND {
    ret, _, _ := upDown_GetHandle.Call(obj)
    return HWND(ret)
}

func UpDown_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := upDown_GetParentWindow.Call(obj)
    return HWND(ret)
}

func UpDown_SetParentWindow(obj uintptr, value HWND) {
   upDown_SetParentWindow.Call(obj, uintptr(value))
}

func UpDown_GetAction(obj uintptr) uintptr {
    ret, _, _ := upDown_GetAction.Call(obj)
    return ret
}

func UpDown_SetAction(obj uintptr, value uintptr) {
   upDown_SetAction.Call(obj, value)
}

func UpDown_GetAlign(obj uintptr) TAlign {
    ret, _, _ := upDown_GetAlign.Call(obj)
    return TAlign(ret)
}

func UpDown_SetAlign(obj uintptr, value TAlign) {
   upDown_SetAlign.Call(obj, uintptr(value))
}

func UpDown_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := upDown_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func UpDown_SetBiDiMode(obj uintptr, value TBiDiMode) {
   upDown_SetBiDiMode.Call(obj, uintptr(value))
}

func UpDown_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    upDown_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func UpDown_SetBoundsRect(obj uintptr, value TRect) {
   upDown_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func UpDown_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := upDown_GetClientHeight.Call(obj)
    return int32(ret)
}

func UpDown_SetClientHeight(obj uintptr, value int32) {
   upDown_SetClientHeight.Call(obj, uintptr(value))
}

func UpDown_GetClientRect(obj uintptr) TRect {
    var ret TRect
    upDown_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func UpDown_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := upDown_GetClientWidth.Call(obj)
    return int32(ret)
}

func UpDown_SetClientWidth(obj uintptr, value int32) {
   upDown_SetClientWidth.Call(obj, uintptr(value))
}

func UpDown_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := upDown_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func UpDown_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := upDown_GetExplicitTop.Call(obj)
    return int32(ret)
}

func UpDown_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := upDown_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func UpDown_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := upDown_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func UpDown_GetParent(obj uintptr) uintptr {
    ret, _, _ := upDown_GetParent.Call(obj)
    return ret
}

func UpDown_SetParent(obj uintptr, value uintptr) {
   upDown_SetParent.Call(obj, value)
}

func UpDown_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := upDown_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func UpDown_SetAlignWithMargins(obj uintptr, value bool) {
   upDown_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func UpDown_GetLeft(obj uintptr) int32 {
    ret, _, _ := upDown_GetLeft.Call(obj)
    return int32(ret)
}

func UpDown_SetLeft(obj uintptr, value int32) {
   upDown_SetLeft.Call(obj, uintptr(value))
}

func UpDown_GetTop(obj uintptr) int32 {
    ret, _, _ := upDown_GetTop.Call(obj)
    return int32(ret)
}

func UpDown_SetTop(obj uintptr, value int32) {
   upDown_SetTop.Call(obj, uintptr(value))
}

func UpDown_GetWidth(obj uintptr) int32 {
    ret, _, _ := upDown_GetWidth.Call(obj)
    return int32(ret)
}

func UpDown_SetWidth(obj uintptr, value int32) {
   upDown_SetWidth.Call(obj, uintptr(value))
}

func UpDown_GetHeight(obj uintptr) int32 {
    ret, _, _ := upDown_GetHeight.Call(obj)
    return int32(ret)
}

func UpDown_SetHeight(obj uintptr, value int32) {
   upDown_SetHeight.Call(obj, uintptr(value))
}

func UpDown_GetCursor(obj uintptr) TCursor {
    ret, _, _ := upDown_GetCursor.Call(obj)
    return TCursor(ret)
}

func UpDown_SetCursor(obj uintptr, value TCursor) {
   upDown_SetCursor.Call(obj, uintptr(value))
}

func UpDown_GetMargins(obj uintptr) uintptr {
    ret, _, _ := upDown_GetMargins.Call(obj)
    return ret
}

func UpDown_SetMargins(obj uintptr, value uintptr) {
   upDown_SetMargins.Call(obj, value)
}

func UpDown_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := upDown_GetComponentCount.Call(obj)
    return int32(ret)
}

func UpDown_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := upDown_GetComponentIndex.Call(obj)
    return int32(ret)
}

func UpDown_SetComponentIndex(obj uintptr, value int32) {
   upDown_SetComponentIndex.Call(obj, uintptr(value))
}

func UpDown_GetOwner(obj uintptr) uintptr {
    ret, _, _ := upDown_GetOwner.Call(obj)
    return ret
}

func UpDown_GetName(obj uintptr) string {
    ret, _, _ := upDown_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func UpDown_SetName(obj uintptr, value string) {
   upDown_SetName.Call(obj, GoStrToDStr(value))
}

func UpDown_GetTag(obj uintptr) int {
    ret, _, _ := upDown_GetTag.Call(obj)
    return int(ret)
}

func UpDown_SetTag(obj uintptr, value int) {
   upDown_SetTag.Call(obj, uintptr(value))
}

func UpDown_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := upDown_GetControls.Call(obj, uintptr(Index))
    return ret
}

func UpDown_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := upDown_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TProgressBar ---------------------------

func ProgressBar_Create(obj uintptr) uintptr {
    ret, _, _ := progressBar_Create.Call(obj)
    return ret
}

func ProgressBar_Free(obj uintptr) {
    progressBar_Free.Call(obj)
}

func ProgressBar_StepIt(obj uintptr)  {
    progressBar_StepIt.Call(obj)
}

func ProgressBar_StepBy(obj uintptr, Delta int32)  {
    progressBar_StepBy.Call(obj, uintptr(Delta) )
}

func ProgressBar_CanFocus(obj uintptr) bool {
    ret, _, _ := progressBar_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_FlipChildren(obj uintptr, AllLevels bool)  {
    progressBar_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func ProgressBar_Focused(obj uintptr) bool {
    ret, _, _ := progressBar_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_HandleAllocated(obj uintptr) bool {
    ret, _, _ := progressBar_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_Invalidate(obj uintptr)  {
    progressBar_Invalidate.Call(obj)
}

func ProgressBar_Realign(obj uintptr)  {
    progressBar_Realign.Call(obj)
}

func ProgressBar_Repaint(obj uintptr)  {
    progressBar_Repaint.Call(obj)
}

func ProgressBar_ScaleBy(obj uintptr, M int32, D int32)  {
    progressBar_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func ProgressBar_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    progressBar_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func ProgressBar_SetFocus(obj uintptr)  {
    progressBar_SetFocus.Call(obj)
}

func ProgressBar_Update(obj uintptr)  {
    progressBar_Update.Call(obj)
}

func ProgressBar_BringToFront(obj uintptr)  {
    progressBar_BringToFront.Call(obj)
}

func ProgressBar_HasParent(obj uintptr) bool {
    ret, _, _ := progressBar_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_Hide(obj uintptr)  {
    progressBar_Hide.Call(obj)
}

func ProgressBar_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := progressBar_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func ProgressBar_Refresh(obj uintptr)  {
    progressBar_Refresh.Call(obj)
}

func ProgressBar_SendToBack(obj uintptr)  {
    progressBar_SendToBack.Call(obj)
}

func ProgressBar_Show(obj uintptr)  {
    progressBar_Show.Call(obj)
}

func ProgressBar_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := progressBar_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func ProgressBar_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := progressBar_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ProgressBar_GetNamePath(obj uintptr) string {
    ret, _, _ := progressBar_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ProgressBar_Assign(obj uintptr, Source uintptr)  {
    progressBar_Assign.Call(obj, Source )
}

func ProgressBar_ClassName(obj uintptr) string {
    ret, _, _ := progressBar_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ProgressBar_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := progressBar_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ProgressBar_GetHashCode(obj uintptr) int32 {
    ret, _, _ := progressBar_GetHashCode.Call(obj)
    return int32(ret)
}

func ProgressBar_ToString(obj uintptr) string {
    ret, _, _ := progressBar_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ProgressBar_GetAlign(obj uintptr) TAlign {
    ret, _, _ := progressBar_GetAlign.Call(obj)
    return TAlign(ret)
}

func ProgressBar_SetAlign(obj uintptr, value TAlign) {
   progressBar_SetAlign.Call(obj, uintptr(value))
}

func ProgressBar_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := progressBar_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func ProgressBar_SetAnchors(obj uintptr, value TAnchors) {
   progressBar_SetAnchors.Call(obj, uintptr(value))
}

func ProgressBar_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := progressBar_GetBorderWidth.Call(obj)
    return int32(ret)
}

func ProgressBar_SetBorderWidth(obj uintptr, value int32) {
   progressBar_SetBorderWidth.Call(obj, uintptr(value))
}

func ProgressBar_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := progressBar_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_SetDoubleBuffered(obj uintptr, value bool) {
   progressBar_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetEnabled(obj uintptr) bool {
    ret, _, _ := progressBar_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_SetEnabled(obj uintptr, value bool) {
   progressBar_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetHint(obj uintptr) string {
    ret, _, _ := progressBar_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func ProgressBar_SetHint(obj uintptr, value string) {
   progressBar_SetHint.Call(obj, GoStrToDStr(value))
}

func ProgressBar_GetMin(obj uintptr) int32 {
    ret, _, _ := progressBar_GetMin.Call(obj)
    return int32(ret)
}

func ProgressBar_SetMin(obj uintptr, value int32) {
   progressBar_SetMin.Call(obj, uintptr(value))
}

func ProgressBar_GetMax(obj uintptr) int32 {
    ret, _, _ := progressBar_GetMax.Call(obj)
    return int32(ret)
}

func ProgressBar_SetMax(obj uintptr, value int32) {
   progressBar_SetMax.Call(obj, uintptr(value))
}

func ProgressBar_GetOrientation(obj uintptr) TProgressBarOrientation {
    ret, _, _ := progressBar_GetOrientation.Call(obj)
    return TProgressBarOrientation(ret)
}

func ProgressBar_SetOrientation(obj uintptr, value TProgressBarOrientation) {
   progressBar_SetOrientation.Call(obj, uintptr(value))
}

func ProgressBar_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := progressBar_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_SetParentDoubleBuffered(obj uintptr, value bool) {
   progressBar_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := progressBar_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_SetParentShowHint(obj uintptr, value bool) {
   progressBar_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := progressBar_GetPopupMenu.Call(obj)
    return ret
}

func ProgressBar_SetPopupMenu(obj uintptr, value uintptr) {
   progressBar_SetPopupMenu.Call(obj, value)
}

func ProgressBar_GetPosition(obj uintptr) int32 {
    ret, _, _ := progressBar_GetPosition.Call(obj)
    return int32(ret)
}

func ProgressBar_SetPosition(obj uintptr, value int32) {
   progressBar_SetPosition.Call(obj, uintptr(value))
}

func ProgressBar_GetSmooth(obj uintptr) bool {
    ret, _, _ := progressBar_GetSmooth.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_SetSmooth(obj uintptr, value bool) {
   progressBar_SetSmooth.Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetStyle(obj uintptr) TProgressBarStyle {
    ret, _, _ := progressBar_GetStyle.Call(obj)
    return TProgressBarStyle(ret)
}

func ProgressBar_SetStyle(obj uintptr, value TProgressBarStyle) {
   progressBar_SetStyle.Call(obj, uintptr(value))
}

func ProgressBar_GetMarqueeInterval(obj uintptr) int32 {
    ret, _, _ := progressBar_GetMarqueeInterval.Call(obj)
    return int32(ret)
}

func ProgressBar_SetMarqueeInterval(obj uintptr, value int32) {
   progressBar_SetMarqueeInterval.Call(obj, uintptr(value))
}

func ProgressBar_GetBarColor(obj uintptr) TColor {
    ret, _, _ := progressBar_GetBarColor.Call(obj)
    return TColor(ret)
}

func ProgressBar_SetBarColor(obj uintptr, value TColor) {
   progressBar_SetBarColor.Call(obj, uintptr(value))
}

func ProgressBar_GetBackgroundColor(obj uintptr) TColor {
    ret, _, _ := progressBar_GetBackgroundColor.Call(obj)
    return TColor(ret)
}

func ProgressBar_SetBackgroundColor(obj uintptr, value TColor) {
   progressBar_SetBackgroundColor.Call(obj, uintptr(value))
}

func ProgressBar_GetSmoothReverse(obj uintptr) bool {
    ret, _, _ := progressBar_GetSmoothReverse.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_SetSmoothReverse(obj uintptr, value bool) {
   progressBar_SetSmoothReverse.Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetStep(obj uintptr) int32 {
    ret, _, _ := progressBar_GetStep.Call(obj)
    return int32(ret)
}

func ProgressBar_SetStep(obj uintptr, value int32) {
   progressBar_SetStep.Call(obj, uintptr(value))
}

func ProgressBar_GetState(obj uintptr) TProgressBarState {
    ret, _, _ := progressBar_GetState.Call(obj)
    return TProgressBarState(ret)
}

func ProgressBar_SetState(obj uintptr, value TProgressBarState) {
   progressBar_SetState.Call(obj, uintptr(value))
}

func ProgressBar_GetShowHint(obj uintptr) bool {
    ret, _, _ := progressBar_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_SetShowHint(obj uintptr, value bool) {
   progressBar_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := progressBar_GetTabOrder.Call(obj)
    return uint16(ret)
}

func ProgressBar_SetTabOrder(obj uintptr, value uint16) {
   progressBar_SetTabOrder.Call(obj, uintptr(value))
}

func ProgressBar_GetTabStop(obj uintptr) bool {
    ret, _, _ := progressBar_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_SetTabStop(obj uintptr, value bool) {
   progressBar_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetVisible(obj uintptr) bool {
    ret, _, _ := progressBar_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_SetVisible(obj uintptr, value bool) {
   progressBar_SetVisible.Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := progressBar_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func ProgressBar_SetStyleElements(obj uintptr, value TStyleElements) {
   progressBar_SetStyleElements.Call(obj, uintptr(value))
}

func ProgressBar_SetOnEnter(obj uintptr, fn interface{}) {
    progressBar_SetOnEnter.Call(obj, addEventToMap(fn))
}

func ProgressBar_SetOnExit(obj uintptr, fn interface{}) {
    progressBar_SetOnExit.Call(obj, addEventToMap(fn))
}

func ProgressBar_SetOnMouseDown(obj uintptr, fn interface{}) {
    progressBar_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func ProgressBar_SetOnMouseEnter(obj uintptr, fn interface{}) {
    progressBar_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func ProgressBar_SetOnMouseLeave(obj uintptr, fn interface{}) {
    progressBar_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func ProgressBar_SetOnMouseMove(obj uintptr, fn interface{}) {
    progressBar_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func ProgressBar_SetOnMouseUp(obj uintptr, fn interface{}) {
    progressBar_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func ProgressBar_GetBrush(obj uintptr) uintptr {
    ret, _, _ := progressBar_GetBrush.Call(obj)
    return ret
}

func ProgressBar_GetControlCount(obj uintptr) int32 {
    ret, _, _ := progressBar_GetControlCount.Call(obj)
    return int32(ret)
}

func ProgressBar_GetHandle(obj uintptr) HWND {
    ret, _, _ := progressBar_GetHandle.Call(obj)
    return HWND(ret)
}

func ProgressBar_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := progressBar_GetParentWindow.Call(obj)
    return HWND(ret)
}

func ProgressBar_SetParentWindow(obj uintptr, value HWND) {
   progressBar_SetParentWindow.Call(obj, uintptr(value))
}

func ProgressBar_GetAction(obj uintptr) uintptr {
    ret, _, _ := progressBar_GetAction.Call(obj)
    return ret
}

func ProgressBar_SetAction(obj uintptr, value uintptr) {
   progressBar_SetAction.Call(obj, value)
}

func ProgressBar_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := progressBar_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func ProgressBar_SetBiDiMode(obj uintptr, value TBiDiMode) {
   progressBar_SetBiDiMode.Call(obj, uintptr(value))
}

func ProgressBar_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    progressBar_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ProgressBar_SetBoundsRect(obj uintptr, value TRect) {
   progressBar_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ProgressBar_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := progressBar_GetClientHeight.Call(obj)
    return int32(ret)
}

func ProgressBar_SetClientHeight(obj uintptr, value int32) {
   progressBar_SetClientHeight.Call(obj, uintptr(value))
}

func ProgressBar_GetClientRect(obj uintptr) TRect {
    var ret TRect
    progressBar_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ProgressBar_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := progressBar_GetClientWidth.Call(obj)
    return int32(ret)
}

func ProgressBar_SetClientWidth(obj uintptr, value int32) {
   progressBar_SetClientWidth.Call(obj, uintptr(value))
}

func ProgressBar_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := progressBar_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func ProgressBar_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := progressBar_GetExplicitTop.Call(obj)
    return int32(ret)
}

func ProgressBar_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := progressBar_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func ProgressBar_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := progressBar_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func ProgressBar_GetParent(obj uintptr) uintptr {
    ret, _, _ := progressBar_GetParent.Call(obj)
    return ret
}

func ProgressBar_SetParent(obj uintptr, value uintptr) {
   progressBar_SetParent.Call(obj, value)
}

func ProgressBar_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := progressBar_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func ProgressBar_SetAlignWithMargins(obj uintptr, value bool) {
   progressBar_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func ProgressBar_GetLeft(obj uintptr) int32 {
    ret, _, _ := progressBar_GetLeft.Call(obj)
    return int32(ret)
}

func ProgressBar_SetLeft(obj uintptr, value int32) {
   progressBar_SetLeft.Call(obj, uintptr(value))
}

func ProgressBar_GetTop(obj uintptr) int32 {
    ret, _, _ := progressBar_GetTop.Call(obj)
    return int32(ret)
}

func ProgressBar_SetTop(obj uintptr, value int32) {
   progressBar_SetTop.Call(obj, uintptr(value))
}

func ProgressBar_GetWidth(obj uintptr) int32 {
    ret, _, _ := progressBar_GetWidth.Call(obj)
    return int32(ret)
}

func ProgressBar_SetWidth(obj uintptr, value int32) {
   progressBar_SetWidth.Call(obj, uintptr(value))
}

func ProgressBar_GetHeight(obj uintptr) int32 {
    ret, _, _ := progressBar_GetHeight.Call(obj)
    return int32(ret)
}

func ProgressBar_SetHeight(obj uintptr, value int32) {
   progressBar_SetHeight.Call(obj, uintptr(value))
}

func ProgressBar_GetCursor(obj uintptr) TCursor {
    ret, _, _ := progressBar_GetCursor.Call(obj)
    return TCursor(ret)
}

func ProgressBar_SetCursor(obj uintptr, value TCursor) {
   progressBar_SetCursor.Call(obj, uintptr(value))
}

func ProgressBar_GetMargins(obj uintptr) uintptr {
    ret, _, _ := progressBar_GetMargins.Call(obj)
    return ret
}

func ProgressBar_SetMargins(obj uintptr, value uintptr) {
   progressBar_SetMargins.Call(obj, value)
}

func ProgressBar_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := progressBar_GetComponentCount.Call(obj)
    return int32(ret)
}

func ProgressBar_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := progressBar_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ProgressBar_SetComponentIndex(obj uintptr, value int32) {
   progressBar_SetComponentIndex.Call(obj, uintptr(value))
}

func ProgressBar_GetOwner(obj uintptr) uintptr {
    ret, _, _ := progressBar_GetOwner.Call(obj)
    return ret
}

func ProgressBar_GetName(obj uintptr) string {
    ret, _, _ := progressBar_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ProgressBar_SetName(obj uintptr, value string) {
   progressBar_SetName.Call(obj, GoStrToDStr(value))
}

func ProgressBar_GetTag(obj uintptr) int {
    ret, _, _ := progressBar_GetTag.Call(obj)
    return int(ret)
}

func ProgressBar_SetTag(obj uintptr, value int) {
   progressBar_SetTag.Call(obj, uintptr(value))
}

func ProgressBar_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := progressBar_GetControls.Call(obj, uintptr(Index))
    return ret
}

func ProgressBar_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := progressBar_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- THotKey ---------------------------

func HotKey_Create(obj uintptr) uintptr {
    ret, _, _ := hotKey_Create.Call(obj)
    return ret
}

func HotKey_Free(obj uintptr) {
    hotKey_Free.Call(obj)
}

func HotKey_CanFocus(obj uintptr) bool {
    ret, _, _ := hotKey_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_FlipChildren(obj uintptr, AllLevels bool)  {
    hotKey_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func HotKey_Focused(obj uintptr) bool {
    ret, _, _ := hotKey_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_HandleAllocated(obj uintptr) bool {
    ret, _, _ := hotKey_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_Invalidate(obj uintptr)  {
    hotKey_Invalidate.Call(obj)
}

func HotKey_Realign(obj uintptr)  {
    hotKey_Realign.Call(obj)
}

func HotKey_Repaint(obj uintptr)  {
    hotKey_Repaint.Call(obj)
}

func HotKey_ScaleBy(obj uintptr, M int32, D int32)  {
    hotKey_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func HotKey_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    hotKey_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func HotKey_SetFocus(obj uintptr)  {
    hotKey_SetFocus.Call(obj)
}

func HotKey_Update(obj uintptr)  {
    hotKey_Update.Call(obj)
}

func HotKey_BringToFront(obj uintptr)  {
    hotKey_BringToFront.Call(obj)
}

func HotKey_HasParent(obj uintptr) bool {
    ret, _, _ := hotKey_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_Hide(obj uintptr)  {
    hotKey_Hide.Call(obj)
}

func HotKey_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := hotKey_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func HotKey_Refresh(obj uintptr)  {
    hotKey_Refresh.Call(obj)
}

func HotKey_SendToBack(obj uintptr)  {
    hotKey_SendToBack.Call(obj)
}

func HotKey_Show(obj uintptr)  {
    hotKey_Show.Call(obj)
}

func HotKey_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := hotKey_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func HotKey_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := hotKey_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func HotKey_GetNamePath(obj uintptr) string {
    ret, _, _ := hotKey_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func HotKey_Assign(obj uintptr, Source uintptr)  {
    hotKey_Assign.Call(obj, Source )
}

func HotKey_ClassName(obj uintptr) string {
    ret, _, _ := hotKey_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func HotKey_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := hotKey_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func HotKey_GetHashCode(obj uintptr) int32 {
    ret, _, _ := hotKey_GetHashCode.Call(obj)
    return int32(ret)
}

func HotKey_ToString(obj uintptr) string {
    ret, _, _ := hotKey_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func HotKey_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := hotKey_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func HotKey_SetAnchors(obj uintptr, value TAnchors) {
   hotKey_SetAnchors.Call(obj, uintptr(value))
}

func HotKey_GetAutoSize(obj uintptr) bool {
    ret, _, _ := hotKey_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_SetAutoSize(obj uintptr, value bool) {
   hotKey_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func HotKey_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := hotKey_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func HotKey_SetBiDiMode(obj uintptr, value TBiDiMode) {
   hotKey_SetBiDiMode.Call(obj, uintptr(value))
}

func HotKey_GetEnabled(obj uintptr) bool {
    ret, _, _ := hotKey_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_SetEnabled(obj uintptr, value bool) {
   hotKey_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func HotKey_GetHint(obj uintptr) string {
    ret, _, _ := hotKey_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func HotKey_SetHint(obj uintptr, value string) {
   hotKey_SetHint.Call(obj, GoStrToDStr(value))
}

func HotKey_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := hotKey_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_SetParentShowHint(obj uintptr, value bool) {
   hotKey_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func HotKey_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := hotKey_GetPopupMenu.Call(obj)
    return ret
}

func HotKey_SetPopupMenu(obj uintptr, value uintptr) {
   hotKey_SetPopupMenu.Call(obj, value)
}

func HotKey_GetShowHint(obj uintptr) bool {
    ret, _, _ := hotKey_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_SetShowHint(obj uintptr, value bool) {
   hotKey_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func HotKey_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := hotKey_GetTabOrder.Call(obj)
    return uint16(ret)
}

func HotKey_SetTabOrder(obj uintptr, value uint16) {
   hotKey_SetTabOrder.Call(obj, uintptr(value))
}

func HotKey_GetTabStop(obj uintptr) bool {
    ret, _, _ := hotKey_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_SetTabStop(obj uintptr, value bool) {
   hotKey_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func HotKey_GetVisible(obj uintptr) bool {
    ret, _, _ := hotKey_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_SetVisible(obj uintptr, value bool) {
   hotKey_SetVisible.Call(obj, GoBoolToDBool(value))
}

func HotKey_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := hotKey_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func HotKey_SetStyleElements(obj uintptr, value TStyleElements) {
   hotKey_SetStyleElements.Call(obj, uintptr(value))
}

func HotKey_SetOnChange(obj uintptr, fn interface{}) {
    hotKey_SetOnChange.Call(obj, addEventToMap(fn))
}

func HotKey_SetOnEnter(obj uintptr, fn interface{}) {
    hotKey_SetOnEnter.Call(obj, addEventToMap(fn))
}

func HotKey_SetOnExit(obj uintptr, fn interface{}) {
    hotKey_SetOnExit.Call(obj, addEventToMap(fn))
}

func HotKey_SetOnMouseDown(obj uintptr, fn interface{}) {
    hotKey_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func HotKey_SetOnMouseEnter(obj uintptr, fn interface{}) {
    hotKey_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func HotKey_SetOnMouseLeave(obj uintptr, fn interface{}) {
    hotKey_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func HotKey_SetOnMouseMove(obj uintptr, fn interface{}) {
    hotKey_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func HotKey_SetOnMouseUp(obj uintptr, fn interface{}) {
    hotKey_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func HotKey_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := hotKey_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_SetDoubleBuffered(obj uintptr, value bool) {
   hotKey_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func HotKey_GetBrush(obj uintptr) uintptr {
    ret, _, _ := hotKey_GetBrush.Call(obj)
    return ret
}

func HotKey_GetControlCount(obj uintptr) int32 {
    ret, _, _ := hotKey_GetControlCount.Call(obj)
    return int32(ret)
}

func HotKey_GetHandle(obj uintptr) HWND {
    ret, _, _ := hotKey_GetHandle.Call(obj)
    return HWND(ret)
}

func HotKey_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := hotKey_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_SetParentDoubleBuffered(obj uintptr, value bool) {
   hotKey_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func HotKey_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := hotKey_GetParentWindow.Call(obj)
    return HWND(ret)
}

func HotKey_SetParentWindow(obj uintptr, value HWND) {
   hotKey_SetParentWindow.Call(obj, uintptr(value))
}

func HotKey_GetAction(obj uintptr) uintptr {
    ret, _, _ := hotKey_GetAction.Call(obj)
    return ret
}

func HotKey_SetAction(obj uintptr, value uintptr) {
   hotKey_SetAction.Call(obj, value)
}

func HotKey_GetAlign(obj uintptr) TAlign {
    ret, _, _ := hotKey_GetAlign.Call(obj)
    return TAlign(ret)
}

func HotKey_SetAlign(obj uintptr, value TAlign) {
   hotKey_SetAlign.Call(obj, uintptr(value))
}

func HotKey_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    hotKey_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func HotKey_SetBoundsRect(obj uintptr, value TRect) {
   hotKey_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func HotKey_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := hotKey_GetClientHeight.Call(obj)
    return int32(ret)
}

func HotKey_SetClientHeight(obj uintptr, value int32) {
   hotKey_SetClientHeight.Call(obj, uintptr(value))
}

func HotKey_GetClientRect(obj uintptr) TRect {
    var ret TRect
    hotKey_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func HotKey_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := hotKey_GetClientWidth.Call(obj)
    return int32(ret)
}

func HotKey_SetClientWidth(obj uintptr, value int32) {
   hotKey_SetClientWidth.Call(obj, uintptr(value))
}

func HotKey_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := hotKey_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func HotKey_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := hotKey_GetExplicitTop.Call(obj)
    return int32(ret)
}

func HotKey_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := hotKey_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func HotKey_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := hotKey_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func HotKey_GetParent(obj uintptr) uintptr {
    ret, _, _ := hotKey_GetParent.Call(obj)
    return ret
}

func HotKey_SetParent(obj uintptr, value uintptr) {
   hotKey_SetParent.Call(obj, value)
}

func HotKey_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := hotKey_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func HotKey_SetAlignWithMargins(obj uintptr, value bool) {
   hotKey_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func HotKey_GetLeft(obj uintptr) int32 {
    ret, _, _ := hotKey_GetLeft.Call(obj)
    return int32(ret)
}

func HotKey_SetLeft(obj uintptr, value int32) {
   hotKey_SetLeft.Call(obj, uintptr(value))
}

func HotKey_GetTop(obj uintptr) int32 {
    ret, _, _ := hotKey_GetTop.Call(obj)
    return int32(ret)
}

func HotKey_SetTop(obj uintptr, value int32) {
   hotKey_SetTop.Call(obj, uintptr(value))
}

func HotKey_GetWidth(obj uintptr) int32 {
    ret, _, _ := hotKey_GetWidth.Call(obj)
    return int32(ret)
}

func HotKey_SetWidth(obj uintptr, value int32) {
   hotKey_SetWidth.Call(obj, uintptr(value))
}

func HotKey_GetHeight(obj uintptr) int32 {
    ret, _, _ := hotKey_GetHeight.Call(obj)
    return int32(ret)
}

func HotKey_SetHeight(obj uintptr, value int32) {
   hotKey_SetHeight.Call(obj, uintptr(value))
}

func HotKey_GetCursor(obj uintptr) TCursor {
    ret, _, _ := hotKey_GetCursor.Call(obj)
    return TCursor(ret)
}

func HotKey_SetCursor(obj uintptr, value TCursor) {
   hotKey_SetCursor.Call(obj, uintptr(value))
}

func HotKey_GetMargins(obj uintptr) uintptr {
    ret, _, _ := hotKey_GetMargins.Call(obj)
    return ret
}

func HotKey_SetMargins(obj uintptr, value uintptr) {
   hotKey_SetMargins.Call(obj, value)
}

func HotKey_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := hotKey_GetComponentCount.Call(obj)
    return int32(ret)
}

func HotKey_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := hotKey_GetComponentIndex.Call(obj)
    return int32(ret)
}

func HotKey_SetComponentIndex(obj uintptr, value int32) {
   hotKey_SetComponentIndex.Call(obj, uintptr(value))
}

func HotKey_GetOwner(obj uintptr) uintptr {
    ret, _, _ := hotKey_GetOwner.Call(obj)
    return ret
}

func HotKey_GetName(obj uintptr) string {
    ret, _, _ := hotKey_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func HotKey_SetName(obj uintptr, value string) {
   hotKey_SetName.Call(obj, GoStrToDStr(value))
}

func HotKey_GetTag(obj uintptr) int {
    ret, _, _ := hotKey_GetTag.Call(obj)
    return int(ret)
}

func HotKey_SetTag(obj uintptr, value int) {
   hotKey_SetTag.Call(obj, uintptr(value))
}

func HotKey_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := hotKey_GetControls.Call(obj, uintptr(Index))
    return ret
}

func HotKey_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := hotKey_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TDateTimePicker ---------------------------

func DateTimePicker_Create(obj uintptr) uintptr {
    ret, _, _ := dateTimePicker_Create.Call(obj)
    return ret
}

func DateTimePicker_Free(obj uintptr) {
    dateTimePicker_Free.Call(obj)
}

func DateTimePicker_CanFocus(obj uintptr) bool {
    ret, _, _ := dateTimePicker_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_FlipChildren(obj uintptr, AllLevels bool)  {
    dateTimePicker_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func DateTimePicker_Focused(obj uintptr) bool {
    ret, _, _ := dateTimePicker_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_HandleAllocated(obj uintptr) bool {
    ret, _, _ := dateTimePicker_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_Invalidate(obj uintptr)  {
    dateTimePicker_Invalidate.Call(obj)
}

func DateTimePicker_Realign(obj uintptr)  {
    dateTimePicker_Realign.Call(obj)
}

func DateTimePicker_Repaint(obj uintptr)  {
    dateTimePicker_Repaint.Call(obj)
}

func DateTimePicker_ScaleBy(obj uintptr, M int32, D int32)  {
    dateTimePicker_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func DateTimePicker_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    dateTimePicker_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func DateTimePicker_SetFocus(obj uintptr)  {
    dateTimePicker_SetFocus.Call(obj)
}

func DateTimePicker_Update(obj uintptr)  {
    dateTimePicker_Update.Call(obj)
}

func DateTimePicker_BringToFront(obj uintptr)  {
    dateTimePicker_BringToFront.Call(obj)
}

func DateTimePicker_HasParent(obj uintptr) bool {
    ret, _, _ := dateTimePicker_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_Hide(obj uintptr)  {
    dateTimePicker_Hide.Call(obj)
}

func DateTimePicker_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := dateTimePicker_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func DateTimePicker_Refresh(obj uintptr)  {
    dateTimePicker_Refresh.Call(obj)
}

func DateTimePicker_SendToBack(obj uintptr)  {
    dateTimePicker_SendToBack.Call(obj)
}

func DateTimePicker_Show(obj uintptr)  {
    dateTimePicker_Show.Call(obj)
}

func DateTimePicker_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := dateTimePicker_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func DateTimePicker_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := dateTimePicker_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func DateTimePicker_GetNamePath(obj uintptr) string {
    ret, _, _ := dateTimePicker_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func DateTimePicker_Assign(obj uintptr, Source uintptr)  {
    dateTimePicker_Assign.Call(obj, Source )
}

func DateTimePicker_ClassName(obj uintptr) string {
    ret, _, _ := dateTimePicker_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func DateTimePicker_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := dateTimePicker_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func DateTimePicker_GetHashCode(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetHashCode.Call(obj)
    return int32(ret)
}

func DateTimePicker_ToString(obj uintptr) string {
    ret, _, _ := dateTimePicker_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func DateTimePicker_GetDateTime(obj uintptr) time.Time {
    ret, _, _ := dateTimePicker_GetDateTime.Call(obj)
    return time.Unix(int64(ret), 0)
}

func DateTimePicker_SetDateTime(obj uintptr, value time.Time) {
   dateTimePicker_SetDateTime.Call(obj, uintptr(value.Unix()))
}

func DateTimePicker_GetDroppedDown(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetDroppedDown.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_GetAlign(obj uintptr) TAlign {
    ret, _, _ := dateTimePicker_GetAlign.Call(obj)
    return TAlign(ret)
}

func DateTimePicker_SetAlign(obj uintptr, value TAlign) {
   dateTimePicker_SetAlign.Call(obj, uintptr(value))
}

func DateTimePicker_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := dateTimePicker_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func DateTimePicker_SetAnchors(obj uintptr, value TAnchors) {
   dateTimePicker_SetAnchors.Call(obj, uintptr(value))
}

func DateTimePicker_GetBevelEdges(obj uintptr) TBevelEdges {
    ret, _, _ := dateTimePicker_GetBevelEdges.Call(obj)
    return TBevelEdges(ret)
}

func DateTimePicker_SetBevelEdges(obj uintptr, value TBevelEdges) {
   dateTimePicker_SetBevelEdges.Call(obj, uintptr(value))
}

func DateTimePicker_GetBevelInner(obj uintptr) TBevelCut {
    ret, _, _ := dateTimePicker_GetBevelInner.Call(obj)
    return TBevelCut(ret)
}

func DateTimePicker_SetBevelInner(obj uintptr, value TBevelCut) {
   dateTimePicker_SetBevelInner.Call(obj, uintptr(value))
}

func DateTimePicker_GetBevelOuter(obj uintptr) TBevelCut {
    ret, _, _ := dateTimePicker_GetBevelOuter.Call(obj)
    return TBevelCut(ret)
}

func DateTimePicker_SetBevelOuter(obj uintptr, value TBevelCut) {
   dateTimePicker_SetBevelOuter.Call(obj, uintptr(value))
}

func DateTimePicker_GetBevelKind(obj uintptr) TBevelKind {
    ret, _, _ := dateTimePicker_GetBevelKind.Call(obj)
    return TBevelKind(ret)
}

func DateTimePicker_SetBevelKind(obj uintptr, value TBevelKind) {
   dateTimePicker_SetBevelKind.Call(obj, uintptr(value))
}

func DateTimePicker_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := dateTimePicker_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func DateTimePicker_SetBiDiMode(obj uintptr, value TBiDiMode) {
   dateTimePicker_SetBiDiMode.Call(obj, uintptr(value))
}

func DateTimePicker_GetCalAlignment(obj uintptr) TDTCalAlignment {
    ret, _, _ := dateTimePicker_GetCalAlignment.Call(obj)
    return TDTCalAlignment(ret)
}

func DateTimePicker_SetCalAlignment(obj uintptr, value TDTCalAlignment) {
   dateTimePicker_SetCalAlignment.Call(obj, uintptr(value))
}

func DateTimePicker_GetCalColors(obj uintptr) uintptr {
    ret, _, _ := dateTimePicker_GetCalColors.Call(obj)
    return ret
}

func DateTimePicker_SetCalColors(obj uintptr, value uintptr) {
   dateTimePicker_SetCalColors.Call(obj, value)
}

func DateTimePicker_GetDate(obj uintptr) time.Time {
    ret, _, _ := dateTimePicker_GetDate.Call(obj)
    return time.Unix(int64(ret), 0)
}

func DateTimePicker_SetDate(obj uintptr, value time.Time) {
   dateTimePicker_SetDate.Call(obj, uintptr(value.Unix()))
}

func DateTimePicker_GetFormat(obj uintptr) string {
    ret, _, _ := dateTimePicker_GetFormat.Call(obj)
    return DStrToGoStr(ret)
}

func DateTimePicker_SetFormat(obj uintptr, value string) {
   dateTimePicker_SetFormat.Call(obj, GoStrToDStr(value))
}

func DateTimePicker_GetTime(obj uintptr) time.Time {
    ret, _, _ := dateTimePicker_GetTime.Call(obj)
    return time.Unix(int64(ret), 0)
}

func DateTimePicker_SetTime(obj uintptr, value time.Time) {
   dateTimePicker_SetTime.Call(obj, uintptr(value.Unix()))
}

func DateTimePicker_GetChecked(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetChecked.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetChecked(obj uintptr, value bool) {
   dateTimePicker_SetChecked.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetColor(obj uintptr) TColor {
    ret, _, _ := dateTimePicker_GetColor.Call(obj)
    return TColor(ret)
}

func DateTimePicker_SetColor(obj uintptr, value TColor) {
   dateTimePicker_SetColor.Call(obj, uintptr(value))
}

func DateTimePicker_GetDateFormat(obj uintptr) TDTDateFormat {
    ret, _, _ := dateTimePicker_GetDateFormat.Call(obj)
    return TDTDateFormat(ret)
}

func DateTimePicker_SetDateFormat(obj uintptr, value TDTDateFormat) {
   dateTimePicker_SetDateFormat.Call(obj, uintptr(value))
}

func DateTimePicker_GetDateMode(obj uintptr) TDTDateMode {
    ret, _, _ := dateTimePicker_GetDateMode.Call(obj)
    return TDTDateMode(ret)
}

func DateTimePicker_SetDateMode(obj uintptr, value TDTDateMode) {
   dateTimePicker_SetDateMode.Call(obj, uintptr(value))
}

func DateTimePicker_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetDoubleBuffered(obj uintptr, value bool) {
   dateTimePicker_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetEnabled(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetEnabled(obj uintptr, value bool) {
   dateTimePicker_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetFont(obj uintptr) uintptr {
    ret, _, _ := dateTimePicker_GetFont.Call(obj)
    return ret
}

func DateTimePicker_SetFont(obj uintptr, value uintptr) {
   dateTimePicker_SetFont.Call(obj, value)
}

func DateTimePicker_GetKind(obj uintptr) TDateTimeKind {
    ret, _, _ := dateTimePicker_GetKind.Call(obj)
    return TDateTimeKind(ret)
}

func DateTimePicker_SetKind(obj uintptr, value TDateTimeKind) {
   dateTimePicker_SetKind.Call(obj, uintptr(value))
}

func DateTimePicker_GetMaxDate(obj uintptr) time.Time {
    ret, _, _ := dateTimePicker_GetMaxDate.Call(obj)
    return time.Unix(int64(ret), 0)
}

func DateTimePicker_SetMaxDate(obj uintptr, value time.Time) {
   dateTimePicker_SetMaxDate.Call(obj, uintptr(value.Unix()))
}

func DateTimePicker_GetMinDate(obj uintptr) time.Time {
    ret, _, _ := dateTimePicker_GetMinDate.Call(obj)
    return time.Unix(int64(ret), 0)
}

func DateTimePicker_SetMinDate(obj uintptr, value time.Time) {
   dateTimePicker_SetMinDate.Call(obj, uintptr(value.Unix()))
}

func DateTimePicker_GetParseInput(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetParseInput.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetParseInput(obj uintptr, value bool) {
   dateTimePicker_SetParseInput.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetParentColor(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetParentColor(obj uintptr, value bool) {
   dateTimePicker_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetParentDoubleBuffered(obj uintptr, value bool) {
   dateTimePicker_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetParentFont(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetParentFont(obj uintptr, value bool) {
   dateTimePicker_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetParentShowHint(obj uintptr, value bool) {
   dateTimePicker_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := dateTimePicker_GetPopupMenu.Call(obj)
    return ret
}

func DateTimePicker_SetPopupMenu(obj uintptr, value uintptr) {
   dateTimePicker_SetPopupMenu.Call(obj, value)
}

func DateTimePicker_GetShowHint(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetShowHint(obj uintptr, value bool) {
   dateTimePicker_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := dateTimePicker_GetTabOrder.Call(obj)
    return uint16(ret)
}

func DateTimePicker_SetTabOrder(obj uintptr, value uint16) {
   dateTimePicker_SetTabOrder.Call(obj, uintptr(value))
}

func DateTimePicker_GetTabStop(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetTabStop(obj uintptr, value bool) {
   dateTimePicker_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetVisible(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetVisible(obj uintptr, value bool) {
   dateTimePicker_SetVisible.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := dateTimePicker_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func DateTimePicker_SetStyleElements(obj uintptr, value TStyleElements) {
   dateTimePicker_SetStyleElements.Call(obj, uintptr(value))
}

func DateTimePicker_SetOnClick(obj uintptr, fn interface{}) {
    dateTimePicker_SetOnClick.Call(obj, addEventToMap(fn))
}

func DateTimePicker_SetOnChange(obj uintptr, fn interface{}) {
    dateTimePicker_SetOnChange.Call(obj, addEventToMap(fn))
}

func DateTimePicker_SetOnEnter(obj uintptr, fn interface{}) {
    dateTimePicker_SetOnEnter.Call(obj, addEventToMap(fn))
}

func DateTimePicker_SetOnExit(obj uintptr, fn interface{}) {
    dateTimePicker_SetOnExit.Call(obj, addEventToMap(fn))
}

func DateTimePicker_SetOnKeyDown(obj uintptr, fn interface{}) {
    dateTimePicker_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func DateTimePicker_SetOnKeyPress(obj uintptr, fn interface{}) {
    dateTimePicker_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func DateTimePicker_SetOnKeyUp(obj uintptr, fn interface{}) {
    dateTimePicker_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func DateTimePicker_SetOnMouseEnter(obj uintptr, fn interface{}) {
    dateTimePicker_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func DateTimePicker_SetOnMouseLeave(obj uintptr, fn interface{}) {
    dateTimePicker_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func DateTimePicker_GetBrush(obj uintptr) uintptr {
    ret, _, _ := dateTimePicker_GetBrush.Call(obj)
    return ret
}

func DateTimePicker_GetControlCount(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetControlCount.Call(obj)
    return int32(ret)
}

func DateTimePicker_GetHandle(obj uintptr) HWND {
    ret, _, _ := dateTimePicker_GetHandle.Call(obj)
    return HWND(ret)
}

func DateTimePicker_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := dateTimePicker_GetParentWindow.Call(obj)
    return HWND(ret)
}

func DateTimePicker_SetParentWindow(obj uintptr, value HWND) {
   dateTimePicker_SetParentWindow.Call(obj, uintptr(value))
}

func DateTimePicker_GetAction(obj uintptr) uintptr {
    ret, _, _ := dateTimePicker_GetAction.Call(obj)
    return ret
}

func DateTimePicker_SetAction(obj uintptr, value uintptr) {
   dateTimePicker_SetAction.Call(obj, value)
}

func DateTimePicker_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    dateTimePicker_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func DateTimePicker_SetBoundsRect(obj uintptr, value TRect) {
   dateTimePicker_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func DateTimePicker_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetClientHeight.Call(obj)
    return int32(ret)
}

func DateTimePicker_SetClientHeight(obj uintptr, value int32) {
   dateTimePicker_SetClientHeight.Call(obj, uintptr(value))
}

func DateTimePicker_GetClientRect(obj uintptr) TRect {
    var ret TRect
    dateTimePicker_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func DateTimePicker_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetClientWidth.Call(obj)
    return int32(ret)
}

func DateTimePicker_SetClientWidth(obj uintptr, value int32) {
   dateTimePicker_SetClientWidth.Call(obj, uintptr(value))
}

func DateTimePicker_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func DateTimePicker_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetExplicitTop.Call(obj)
    return int32(ret)
}

func DateTimePicker_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func DateTimePicker_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func DateTimePicker_GetParent(obj uintptr) uintptr {
    ret, _, _ := dateTimePicker_GetParent.Call(obj)
    return ret
}

func DateTimePicker_SetParent(obj uintptr, value uintptr) {
   dateTimePicker_SetParent.Call(obj, value)
}

func DateTimePicker_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := dateTimePicker_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func DateTimePicker_SetAlignWithMargins(obj uintptr, value bool) {
   dateTimePicker_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func DateTimePicker_GetLeft(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetLeft.Call(obj)
    return int32(ret)
}

func DateTimePicker_SetLeft(obj uintptr, value int32) {
   dateTimePicker_SetLeft.Call(obj, uintptr(value))
}

func DateTimePicker_GetTop(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetTop.Call(obj)
    return int32(ret)
}

func DateTimePicker_SetTop(obj uintptr, value int32) {
   dateTimePicker_SetTop.Call(obj, uintptr(value))
}

func DateTimePicker_GetWidth(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetWidth.Call(obj)
    return int32(ret)
}

func DateTimePicker_SetWidth(obj uintptr, value int32) {
   dateTimePicker_SetWidth.Call(obj, uintptr(value))
}

func DateTimePicker_GetHeight(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetHeight.Call(obj)
    return int32(ret)
}

func DateTimePicker_SetHeight(obj uintptr, value int32) {
   dateTimePicker_SetHeight.Call(obj, uintptr(value))
}

func DateTimePicker_GetCursor(obj uintptr) TCursor {
    ret, _, _ := dateTimePicker_GetCursor.Call(obj)
    return TCursor(ret)
}

func DateTimePicker_SetCursor(obj uintptr, value TCursor) {
   dateTimePicker_SetCursor.Call(obj, uintptr(value))
}

func DateTimePicker_GetHint(obj uintptr) string {
    ret, _, _ := dateTimePicker_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func DateTimePicker_SetHint(obj uintptr, value string) {
   dateTimePicker_SetHint.Call(obj, GoStrToDStr(value))
}

func DateTimePicker_GetMargins(obj uintptr) uintptr {
    ret, _, _ := dateTimePicker_GetMargins.Call(obj)
    return ret
}

func DateTimePicker_SetMargins(obj uintptr, value uintptr) {
   dateTimePicker_SetMargins.Call(obj, value)
}

func DateTimePicker_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetComponentCount.Call(obj)
    return int32(ret)
}

func DateTimePicker_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := dateTimePicker_GetComponentIndex.Call(obj)
    return int32(ret)
}

func DateTimePicker_SetComponentIndex(obj uintptr, value int32) {
   dateTimePicker_SetComponentIndex.Call(obj, uintptr(value))
}

func DateTimePicker_GetOwner(obj uintptr) uintptr {
    ret, _, _ := dateTimePicker_GetOwner.Call(obj)
    return ret
}

func DateTimePicker_GetName(obj uintptr) string {
    ret, _, _ := dateTimePicker_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func DateTimePicker_SetName(obj uintptr, value string) {
   dateTimePicker_SetName.Call(obj, GoStrToDStr(value))
}

func DateTimePicker_GetTag(obj uintptr) int {
    ret, _, _ := dateTimePicker_GetTag.Call(obj)
    return int(ret)
}

func DateTimePicker_SetTag(obj uintptr, value int) {
   dateTimePicker_SetTag.Call(obj, uintptr(value))
}

func DateTimePicker_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := dateTimePicker_GetControls.Call(obj, uintptr(Index))
    return ret
}

func DateTimePicker_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := dateTimePicker_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TMonthCalendar ---------------------------

func MonthCalendar_Create(obj uintptr) uintptr {
    ret, _, _ := monthCalendar_Create.Call(obj)
    return ret
}

func MonthCalendar_Free(obj uintptr) {
    monthCalendar_Free.Call(obj)
}

func MonthCalendar_CanFocus(obj uintptr) bool {
    ret, _, _ := monthCalendar_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_FlipChildren(obj uintptr, AllLevels bool)  {
    monthCalendar_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func MonthCalendar_Focused(obj uintptr) bool {
    ret, _, _ := monthCalendar_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_HandleAllocated(obj uintptr) bool {
    ret, _, _ := monthCalendar_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_Invalidate(obj uintptr)  {
    monthCalendar_Invalidate.Call(obj)
}

func MonthCalendar_Realign(obj uintptr)  {
    monthCalendar_Realign.Call(obj)
}

func MonthCalendar_Repaint(obj uintptr)  {
    monthCalendar_Repaint.Call(obj)
}

func MonthCalendar_ScaleBy(obj uintptr, M int32, D int32)  {
    monthCalendar_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func MonthCalendar_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    monthCalendar_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func MonthCalendar_SetFocus(obj uintptr)  {
    monthCalendar_SetFocus.Call(obj)
}

func MonthCalendar_Update(obj uintptr)  {
    monthCalendar_Update.Call(obj)
}

func MonthCalendar_BringToFront(obj uintptr)  {
    monthCalendar_BringToFront.Call(obj)
}

func MonthCalendar_HasParent(obj uintptr) bool {
    ret, _, _ := monthCalendar_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_Hide(obj uintptr)  {
    monthCalendar_Hide.Call(obj)
}

func MonthCalendar_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := monthCalendar_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func MonthCalendar_Refresh(obj uintptr)  {
    monthCalendar_Refresh.Call(obj)
}

func MonthCalendar_SendToBack(obj uintptr)  {
    monthCalendar_SendToBack.Call(obj)
}

func MonthCalendar_Show(obj uintptr)  {
    monthCalendar_Show.Call(obj)
}

func MonthCalendar_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := monthCalendar_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func MonthCalendar_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := monthCalendar_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func MonthCalendar_GetNamePath(obj uintptr) string {
    ret, _, _ := monthCalendar_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func MonthCalendar_Assign(obj uintptr, Source uintptr)  {
    monthCalendar_Assign.Call(obj, Source )
}

func MonthCalendar_ClassName(obj uintptr) string {
    ret, _, _ := monthCalendar_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func MonthCalendar_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := monthCalendar_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func MonthCalendar_GetHashCode(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetHashCode.Call(obj)
    return int32(ret)
}

func MonthCalendar_ToString(obj uintptr) string {
    ret, _, _ := monthCalendar_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func MonthCalendar_GetAlign(obj uintptr) TAlign {
    ret, _, _ := monthCalendar_GetAlign.Call(obj)
    return TAlign(ret)
}

func MonthCalendar_SetAlign(obj uintptr, value TAlign) {
   monthCalendar_SetAlign.Call(obj, uintptr(value))
}

func MonthCalendar_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := monthCalendar_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func MonthCalendar_SetAnchors(obj uintptr, value TAnchors) {
   monthCalendar_SetAnchors.Call(obj, uintptr(value))
}

func MonthCalendar_GetAutoSize(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetAutoSize(obj uintptr, value bool) {
   monthCalendar_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetBorderWidth.Call(obj)
    return int32(ret)
}

func MonthCalendar_SetBorderWidth(obj uintptr, value int32) {
   monthCalendar_SetBorderWidth.Call(obj, uintptr(value))
}

func MonthCalendar_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := monthCalendar_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func MonthCalendar_SetBiDiMode(obj uintptr, value TBiDiMode) {
   monthCalendar_SetBiDiMode.Call(obj, uintptr(value))
}

func MonthCalendar_GetCalColors(obj uintptr) uintptr {
    ret, _, _ := monthCalendar_GetCalColors.Call(obj)
    return ret
}

func MonthCalendar_SetCalColors(obj uintptr, value uintptr) {
   monthCalendar_SetCalColors.Call(obj, value)
}

func MonthCalendar_GetMultiSelect(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetMultiSelect.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetMultiSelect(obj uintptr, value bool) {
   monthCalendar_SetMultiSelect.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetDate(obj uintptr) time.Time {
    ret, _, _ := monthCalendar_GetDate.Call(obj)
    return time.Unix(int64(ret), 0)
}

func MonthCalendar_SetDate(obj uintptr, value time.Time) {
   monthCalendar_SetDate.Call(obj, uintptr(value.Unix()))
}

func MonthCalendar_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetDoubleBuffered(obj uintptr, value bool) {
   monthCalendar_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetEnabled(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetEnabled(obj uintptr, value bool) {
   monthCalendar_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetFirstDayOfWeek(obj uintptr) TCalDayOfWeek {
    ret, _, _ := monthCalendar_GetFirstDayOfWeek.Call(obj)
    return TCalDayOfWeek(ret)
}

func MonthCalendar_SetFirstDayOfWeek(obj uintptr, value TCalDayOfWeek) {
   monthCalendar_SetFirstDayOfWeek.Call(obj, uintptr(value))
}

func MonthCalendar_GetFont(obj uintptr) uintptr {
    ret, _, _ := monthCalendar_GetFont.Call(obj)
    return ret
}

func MonthCalendar_SetFont(obj uintptr, value uintptr) {
   monthCalendar_SetFont.Call(obj, value)
}

func MonthCalendar_GetMaxDate(obj uintptr) time.Time {
    ret, _, _ := monthCalendar_GetMaxDate.Call(obj)
    return time.Unix(int64(ret), 0)
}

func MonthCalendar_SetMaxDate(obj uintptr, value time.Time) {
   monthCalendar_SetMaxDate.Call(obj, uintptr(value.Unix()))
}

func MonthCalendar_GetMaxSelectRange(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetMaxSelectRange.Call(obj)
    return int32(ret)
}

func MonthCalendar_SetMaxSelectRange(obj uintptr, value int32) {
   monthCalendar_SetMaxSelectRange.Call(obj, uintptr(value))
}

func MonthCalendar_GetMinDate(obj uintptr) time.Time {
    ret, _, _ := monthCalendar_GetMinDate.Call(obj)
    return time.Unix(int64(ret), 0)
}

func MonthCalendar_SetMinDate(obj uintptr, value time.Time) {
   monthCalendar_SetMinDate.Call(obj, uintptr(value.Unix()))
}

func MonthCalendar_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetParentDoubleBuffered(obj uintptr, value bool) {
   monthCalendar_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetParentFont(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetParentFont(obj uintptr, value bool) {
   monthCalendar_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetParentShowHint(obj uintptr, value bool) {
   monthCalendar_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := monthCalendar_GetPopupMenu.Call(obj)
    return ret
}

func MonthCalendar_SetPopupMenu(obj uintptr, value uintptr) {
   monthCalendar_SetPopupMenu.Call(obj, value)
}

func MonthCalendar_GetShowHint(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetShowHint(obj uintptr, value bool) {
   monthCalendar_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetShowToday(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetShowToday.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetShowToday(obj uintptr, value bool) {
   monthCalendar_SetShowToday.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetShowTodayCircle(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetShowTodayCircle.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetShowTodayCircle(obj uintptr, value bool) {
   monthCalendar_SetShowTodayCircle.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := monthCalendar_GetTabOrder.Call(obj)
    return uint16(ret)
}

func MonthCalendar_SetTabOrder(obj uintptr, value uint16) {
   monthCalendar_SetTabOrder.Call(obj, uintptr(value))
}

func MonthCalendar_GetTabStop(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetTabStop(obj uintptr, value bool) {
   monthCalendar_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetVisible(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetVisible(obj uintptr, value bool) {
   monthCalendar_SetVisible.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetWeekNumbers(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetWeekNumbers.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetWeekNumbers(obj uintptr, value bool) {
   monthCalendar_SetWeekNumbers.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_SetOnClick(obj uintptr, fn interface{}) {
    monthCalendar_SetOnClick.Call(obj, addEventToMap(fn))
}

func MonthCalendar_SetOnDblClick(obj uintptr, fn interface{}) {
    monthCalendar_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func MonthCalendar_SetOnEnter(obj uintptr, fn interface{}) {
    monthCalendar_SetOnEnter.Call(obj, addEventToMap(fn))
}

func MonthCalendar_SetOnExit(obj uintptr, fn interface{}) {
    monthCalendar_SetOnExit.Call(obj, addEventToMap(fn))
}

func MonthCalendar_SetOnKeyDown(obj uintptr, fn interface{}) {
    monthCalendar_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func MonthCalendar_SetOnKeyPress(obj uintptr, fn interface{}) {
    monthCalendar_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func MonthCalendar_SetOnKeyUp(obj uintptr, fn interface{}) {
    monthCalendar_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func MonthCalendar_SetOnMouseEnter(obj uintptr, fn interface{}) {
    monthCalendar_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func MonthCalendar_SetOnMouseLeave(obj uintptr, fn interface{}) {
    monthCalendar_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func MonthCalendar_GetBrush(obj uintptr) uintptr {
    ret, _, _ := monthCalendar_GetBrush.Call(obj)
    return ret
}

func MonthCalendar_GetControlCount(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetControlCount.Call(obj)
    return int32(ret)
}

func MonthCalendar_GetHandle(obj uintptr) HWND {
    ret, _, _ := monthCalendar_GetHandle.Call(obj)
    return HWND(ret)
}

func MonthCalendar_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := monthCalendar_GetParentWindow.Call(obj)
    return HWND(ret)
}

func MonthCalendar_SetParentWindow(obj uintptr, value HWND) {
   monthCalendar_SetParentWindow.Call(obj, uintptr(value))
}

func MonthCalendar_GetAction(obj uintptr) uintptr {
    ret, _, _ := monthCalendar_GetAction.Call(obj)
    return ret
}

func MonthCalendar_SetAction(obj uintptr, value uintptr) {
   monthCalendar_SetAction.Call(obj, value)
}

func MonthCalendar_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    monthCalendar_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func MonthCalendar_SetBoundsRect(obj uintptr, value TRect) {
   monthCalendar_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func MonthCalendar_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetClientHeight.Call(obj)
    return int32(ret)
}

func MonthCalendar_SetClientHeight(obj uintptr, value int32) {
   monthCalendar_SetClientHeight.Call(obj, uintptr(value))
}

func MonthCalendar_GetClientRect(obj uintptr) TRect {
    var ret TRect
    monthCalendar_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func MonthCalendar_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetClientWidth.Call(obj)
    return int32(ret)
}

func MonthCalendar_SetClientWidth(obj uintptr, value int32) {
   monthCalendar_SetClientWidth.Call(obj, uintptr(value))
}

func MonthCalendar_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func MonthCalendar_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetExplicitTop.Call(obj)
    return int32(ret)
}

func MonthCalendar_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func MonthCalendar_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func MonthCalendar_GetParent(obj uintptr) uintptr {
    ret, _, _ := monthCalendar_GetParent.Call(obj)
    return ret
}

func MonthCalendar_SetParent(obj uintptr, value uintptr) {
   monthCalendar_SetParent.Call(obj, value)
}

func MonthCalendar_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := monthCalendar_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func MonthCalendar_SetStyleElements(obj uintptr, value TStyleElements) {
   monthCalendar_SetStyleElements.Call(obj, uintptr(value))
}

func MonthCalendar_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := monthCalendar_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func MonthCalendar_SetAlignWithMargins(obj uintptr, value bool) {
   monthCalendar_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func MonthCalendar_GetLeft(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetLeft.Call(obj)
    return int32(ret)
}

func MonthCalendar_SetLeft(obj uintptr, value int32) {
   monthCalendar_SetLeft.Call(obj, uintptr(value))
}

func MonthCalendar_GetTop(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetTop.Call(obj)
    return int32(ret)
}

func MonthCalendar_SetTop(obj uintptr, value int32) {
   monthCalendar_SetTop.Call(obj, uintptr(value))
}

func MonthCalendar_GetWidth(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetWidth.Call(obj)
    return int32(ret)
}

func MonthCalendar_SetWidth(obj uintptr, value int32) {
   monthCalendar_SetWidth.Call(obj, uintptr(value))
}

func MonthCalendar_GetHeight(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetHeight.Call(obj)
    return int32(ret)
}

func MonthCalendar_SetHeight(obj uintptr, value int32) {
   monthCalendar_SetHeight.Call(obj, uintptr(value))
}

func MonthCalendar_GetCursor(obj uintptr) TCursor {
    ret, _, _ := monthCalendar_GetCursor.Call(obj)
    return TCursor(ret)
}

func MonthCalendar_SetCursor(obj uintptr, value TCursor) {
   monthCalendar_SetCursor.Call(obj, uintptr(value))
}

func MonthCalendar_GetHint(obj uintptr) string {
    ret, _, _ := monthCalendar_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func MonthCalendar_SetHint(obj uintptr, value string) {
   monthCalendar_SetHint.Call(obj, GoStrToDStr(value))
}

func MonthCalendar_GetMargins(obj uintptr) uintptr {
    ret, _, _ := monthCalendar_GetMargins.Call(obj)
    return ret
}

func MonthCalendar_SetMargins(obj uintptr, value uintptr) {
   monthCalendar_SetMargins.Call(obj, value)
}

func MonthCalendar_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetComponentCount.Call(obj)
    return int32(ret)
}

func MonthCalendar_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := monthCalendar_GetComponentIndex.Call(obj)
    return int32(ret)
}

func MonthCalendar_SetComponentIndex(obj uintptr, value int32) {
   monthCalendar_SetComponentIndex.Call(obj, uintptr(value))
}

func MonthCalendar_GetOwner(obj uintptr) uintptr {
    ret, _, _ := monthCalendar_GetOwner.Call(obj)
    return ret
}

func MonthCalendar_GetName(obj uintptr) string {
    ret, _, _ := monthCalendar_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func MonthCalendar_SetName(obj uintptr, value string) {
   monthCalendar_SetName.Call(obj, GoStrToDStr(value))
}

func MonthCalendar_GetTag(obj uintptr) int {
    ret, _, _ := monthCalendar_GetTag.Call(obj)
    return int(ret)
}

func MonthCalendar_SetTag(obj uintptr, value int) {
   monthCalendar_SetTag.Call(obj, uintptr(value))
}

func MonthCalendar_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := monthCalendar_GetControls.Call(obj, uintptr(Index))
    return ret
}

func MonthCalendar_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := monthCalendar_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TListView ---------------------------

func ListView_Create(obj uintptr) uintptr {
    ret, _, _ := listView_Create.Call(obj)
    return ret
}

func ListView_Free(obj uintptr) {
    listView_Free.Call(obj)
}

func ListView_AddItem(obj uintptr, Item string, AObject uintptr)  {
    listView_AddItem.Call(obj, GoStrToDStr(Item) , AObject )
}

func ListView_AlphaSort(obj uintptr) bool {
    ret, _, _ := listView_AlphaSort.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_Clear(obj uintptr)  {
    listView_Clear.Call(obj)
}

func ListView_ClearSelection(obj uintptr)  {
    listView_ClearSelection.Call(obj)
}

func ListView_DeleteSelected(obj uintptr)  {
    listView_DeleteSelected.Call(obj)
}

func ListView_GetSearchString(obj uintptr) string {
    ret, _, _ := listView_GetSearchString.Call(obj)
    return DStrToGoStr(ret)
}

func ListView_IsEditing(obj uintptr) bool {
    ret, _, _ := listView_IsEditing.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SelectAll(obj uintptr)  {
    listView_SelectAll.Call(obj)
}

func ListView_Scroll(obj uintptr, DX int32, DY int32)  {
    listView_Scroll.Call(obj, uintptr(DX) , uintptr(DY) )
}

func ListView_CustomSort(obj uintptr, SortProc PFNLVCOMPARE, lParam int) bool {
    ret, _, _ := listView_CustomSort.Call(obj, uintptr(SortProc) , uintptr(lParam) )
    return DBoolToGoBool(ret)
}

func ListView_CanFocus(obj uintptr) bool {
    ret, _, _ := listView_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_FlipChildren(obj uintptr, AllLevels bool)  {
    listView_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func ListView_Focused(obj uintptr) bool {
    ret, _, _ := listView_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_HandleAllocated(obj uintptr) bool {
    ret, _, _ := listView_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_Invalidate(obj uintptr)  {
    listView_Invalidate.Call(obj)
}

func ListView_Realign(obj uintptr)  {
    listView_Realign.Call(obj)
}

func ListView_Repaint(obj uintptr)  {
    listView_Repaint.Call(obj)
}

func ListView_ScaleBy(obj uintptr, M int32, D int32)  {
    listView_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func ListView_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    listView_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func ListView_SetFocus(obj uintptr)  {
    listView_SetFocus.Call(obj)
}

func ListView_Update(obj uintptr)  {
    listView_Update.Call(obj)
}

func ListView_BringToFront(obj uintptr)  {
    listView_BringToFront.Call(obj)
}

func ListView_HasParent(obj uintptr) bool {
    ret, _, _ := listView_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_Hide(obj uintptr)  {
    listView_Hide.Call(obj)
}

func ListView_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := listView_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func ListView_Refresh(obj uintptr)  {
    listView_Refresh.Call(obj)
}

func ListView_SendToBack(obj uintptr)  {
    listView_SendToBack.Call(obj)
}

func ListView_Show(obj uintptr)  {
    listView_Show.Call(obj)
}

func ListView_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := listView_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func ListView_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := listView_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ListView_GetNamePath(obj uintptr) string {
    ret, _, _ := listView_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ListView_Assign(obj uintptr, Source uintptr)  {
    listView_Assign.Call(obj, Source )
}

func ListView_ClassName(obj uintptr) string {
    ret, _, _ := listView_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ListView_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := listView_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ListView_GetHashCode(obj uintptr) int32 {
    ret, _, _ := listView_GetHashCode.Call(obj)
    return int32(ret)
}

func ListView_ToString(obj uintptr) string {
    ret, _, _ := listView_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ListView_GetAction(obj uintptr) uintptr {
    ret, _, _ := listView_GetAction.Call(obj)
    return ret
}

func ListView_SetAction(obj uintptr, value uintptr) {
   listView_SetAction.Call(obj, value)
}

func ListView_GetAlign(obj uintptr) TAlign {
    ret, _, _ := listView_GetAlign.Call(obj)
    return TAlign(ret)
}

func ListView_SetAlign(obj uintptr, value TAlign) {
   listView_SetAlign.Call(obj, uintptr(value))
}

func ListView_GetAllocBy(obj uintptr) int32 {
    ret, _, _ := listView_GetAllocBy.Call(obj)
    return int32(ret)
}

func ListView_SetAllocBy(obj uintptr, value int32) {
   listView_SetAllocBy.Call(obj, uintptr(value))
}

func ListView_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := listView_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func ListView_SetAnchors(obj uintptr, value TAnchors) {
   listView_SetAnchors.Call(obj, uintptr(value))
}

func ListView_GetBevelEdges(obj uintptr) TBevelEdges {
    ret, _, _ := listView_GetBevelEdges.Call(obj)
    return TBevelEdges(ret)
}

func ListView_SetBevelEdges(obj uintptr, value TBevelEdges) {
   listView_SetBevelEdges.Call(obj, uintptr(value))
}

func ListView_GetBevelInner(obj uintptr) TBevelCut {
    ret, _, _ := listView_GetBevelInner.Call(obj)
    return TBevelCut(ret)
}

func ListView_SetBevelInner(obj uintptr, value TBevelCut) {
   listView_SetBevelInner.Call(obj, uintptr(value))
}

func ListView_GetBevelOuter(obj uintptr) TBevelCut {
    ret, _, _ := listView_GetBevelOuter.Call(obj)
    return TBevelCut(ret)
}

func ListView_SetBevelOuter(obj uintptr, value TBevelCut) {
   listView_SetBevelOuter.Call(obj, uintptr(value))
}

func ListView_GetBevelKind(obj uintptr) TBevelKind {
    ret, _, _ := listView_GetBevelKind.Call(obj)
    return TBevelKind(ret)
}

func ListView_SetBevelKind(obj uintptr, value TBevelKind) {
   listView_SetBevelKind.Call(obj, uintptr(value))
}

func ListView_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := listView_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func ListView_SetBiDiMode(obj uintptr, value TBiDiMode) {
   listView_SetBiDiMode.Call(obj, uintptr(value))
}

func ListView_GetBorderStyle(obj uintptr) TBorderStyle {
    ret, _, _ := listView_GetBorderStyle.Call(obj)
    return TBorderStyle(ret)
}

func ListView_SetBorderStyle(obj uintptr, value TBorderStyle) {
   listView_SetBorderStyle.Call(obj, uintptr(value))
}

func ListView_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := listView_GetBorderWidth.Call(obj)
    return int32(ret)
}

func ListView_SetBorderWidth(obj uintptr, value int32) {
   listView_SetBorderWidth.Call(obj, uintptr(value))
}

func ListView_GetCheckboxes(obj uintptr) bool {
    ret, _, _ := listView_GetCheckboxes.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetCheckboxes(obj uintptr, value bool) {
   listView_SetCheckboxes.Call(obj, GoBoolToDBool(value))
}

func ListView_GetColor(obj uintptr) TColor {
    ret, _, _ := listView_GetColor.Call(obj)
    return TColor(ret)
}

func ListView_SetColor(obj uintptr, value TColor) {
   listView_SetColor.Call(obj, uintptr(value))
}

func ListView_GetColumns(obj uintptr) uintptr {
    ret, _, _ := listView_GetColumns.Call(obj)
    return ret
}

func ListView_SetColumns(obj uintptr, value uintptr) {
   listView_SetColumns.Call(obj, value)
}

func ListView_GetColumnClick(obj uintptr) bool {
    ret, _, _ := listView_GetColumnClick.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetColumnClick(obj uintptr, value bool) {
   listView_SetColumnClick.Call(obj, GoBoolToDBool(value))
}

func ListView_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := listView_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetDoubleBuffered(obj uintptr, value bool) {
   listView_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ListView_GetEnabled(obj uintptr) bool {
    ret, _, _ := listView_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetEnabled(obj uintptr, value bool) {
   listView_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func ListView_GetFont(obj uintptr) uintptr {
    ret, _, _ := listView_GetFont.Call(obj)
    return ret
}

func ListView_SetFont(obj uintptr, value uintptr) {
   listView_SetFont.Call(obj, value)
}

func ListView_GetFlatScrollBars(obj uintptr) bool {
    ret, _, _ := listView_GetFlatScrollBars.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetFlatScrollBars(obj uintptr, value bool) {
   listView_SetFlatScrollBars.Call(obj, GoBoolToDBool(value))
}

func ListView_GetFullDrag(obj uintptr) bool {
    ret, _, _ := listView_GetFullDrag.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetFullDrag(obj uintptr, value bool) {
   listView_SetFullDrag.Call(obj, GoBoolToDBool(value))
}

func ListView_GetGridLines(obj uintptr) bool {
    ret, _, _ := listView_GetGridLines.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetGridLines(obj uintptr, value bool) {
   listView_SetGridLines.Call(obj, GoBoolToDBool(value))
}

func ListView_GetGroups(obj uintptr) uintptr {
    ret, _, _ := listView_GetGroups.Call(obj)
    return ret
}

func ListView_SetGroups(obj uintptr, value uintptr) {
   listView_SetGroups.Call(obj, value)
}

func ListView_GetHideSelection(obj uintptr) bool {
    ret, _, _ := listView_GetHideSelection.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetHideSelection(obj uintptr, value bool) {
   listView_SetHideSelection.Call(obj, GoBoolToDBool(value))
}

func ListView_GetHotTrack(obj uintptr) bool {
    ret, _, _ := listView_GetHotTrack.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetHotTrack(obj uintptr, value bool) {
   listView_SetHotTrack.Call(obj, GoBoolToDBool(value))
}

func ListView_GetHoverTime(obj uintptr) int32 {
    ret, _, _ := listView_GetHoverTime.Call(obj)
    return int32(ret)
}

func ListView_SetHoverTime(obj uintptr, value int32) {
   listView_SetHoverTime.Call(obj, uintptr(value))
}

func ListView_GetIconOptions(obj uintptr) uintptr {
    ret, _, _ := listView_GetIconOptions.Call(obj)
    return ret
}

func ListView_SetIconOptions(obj uintptr, value uintptr) {
   listView_SetIconOptions.Call(obj, value)
}

func ListView_GetItems(obj uintptr) uintptr {
    ret, _, _ := listView_GetItems.Call(obj)
    return ret
}

func ListView_SetItems(obj uintptr, value uintptr) {
   listView_SetItems.Call(obj, value)
}

func ListView_GetLargeImages(obj uintptr) uintptr {
    ret, _, _ := listView_GetLargeImages.Call(obj)
    return ret
}

func ListView_SetLargeImages(obj uintptr, value uintptr) {
   listView_SetLargeImages.Call(obj, value)
}

func ListView_GetMultiSelect(obj uintptr) bool {
    ret, _, _ := listView_GetMultiSelect.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetMultiSelect(obj uintptr, value bool) {
   listView_SetMultiSelect.Call(obj, GoBoolToDBool(value))
}

func ListView_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := listView_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func ListView_SetStyleElements(obj uintptr, value TStyleElements) {
   listView_SetStyleElements.Call(obj, uintptr(value))
}

func ListView_GetGroupHeaderImages(obj uintptr) uintptr {
    ret, _, _ := listView_GetGroupHeaderImages.Call(obj)
    return ret
}

func ListView_SetGroupHeaderImages(obj uintptr, value uintptr) {
   listView_SetGroupHeaderImages.Call(obj, value)
}

func ListView_GetGroupView(obj uintptr) bool {
    ret, _, _ := listView_GetGroupView.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetGroupView(obj uintptr, value bool) {
   listView_SetGroupView.Call(obj, GoBoolToDBool(value))
}

func ListView_GetReadOnly(obj uintptr) bool {
    ret, _, _ := listView_GetReadOnly.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetReadOnly(obj uintptr, value bool) {
   listView_SetReadOnly.Call(obj, GoBoolToDBool(value))
}

func ListView_GetRowSelect(obj uintptr) bool {
    ret, _, _ := listView_GetRowSelect.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetRowSelect(obj uintptr, value bool) {
   listView_SetRowSelect.Call(obj, GoBoolToDBool(value))
}

func ListView_GetParentColor(obj uintptr) bool {
    ret, _, _ := listView_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetParentColor(obj uintptr, value bool) {
   listView_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func ListView_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := listView_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetParentDoubleBuffered(obj uintptr, value bool) {
   listView_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ListView_GetParentFont(obj uintptr) bool {
    ret, _, _ := listView_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetParentFont(obj uintptr, value bool) {
   listView_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func ListView_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := listView_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetParentShowHint(obj uintptr, value bool) {
   listView_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func ListView_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := listView_GetPopupMenu.Call(obj)
    return ret
}

func ListView_SetPopupMenu(obj uintptr, value uintptr) {
   listView_SetPopupMenu.Call(obj, value)
}

func ListView_GetShowColumnHeaders(obj uintptr) bool {
    ret, _, _ := listView_GetShowColumnHeaders.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetShowColumnHeaders(obj uintptr, value bool) {
   listView_SetShowColumnHeaders.Call(obj, GoBoolToDBool(value))
}

func ListView_GetShowWorkAreas(obj uintptr) bool {
    ret, _, _ := listView_GetShowWorkAreas.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetShowWorkAreas(obj uintptr, value bool) {
   listView_SetShowWorkAreas.Call(obj, GoBoolToDBool(value))
}

func ListView_GetShowHint(obj uintptr) bool {
    ret, _, _ := listView_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetShowHint(obj uintptr, value bool) {
   listView_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func ListView_GetSmallImages(obj uintptr) uintptr {
    ret, _, _ := listView_GetSmallImages.Call(obj)
    return ret
}

func ListView_SetSmallImages(obj uintptr, value uintptr) {
   listView_SetSmallImages.Call(obj, value)
}

func ListView_GetSortType(obj uintptr) TSortType {
    ret, _, _ := listView_GetSortType.Call(obj)
    return TSortType(ret)
}

func ListView_SetSortType(obj uintptr, value TSortType) {
   listView_SetSortType.Call(obj, uintptr(value))
}

func ListView_GetStateImages(obj uintptr) uintptr {
    ret, _, _ := listView_GetStateImages.Call(obj)
    return ret
}

func ListView_SetStateImages(obj uintptr, value uintptr) {
   listView_SetStateImages.Call(obj, value)
}

func ListView_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := listView_GetTabOrder.Call(obj)
    return uint16(ret)
}

func ListView_SetTabOrder(obj uintptr, value uint16) {
   listView_SetTabOrder.Call(obj, uintptr(value))
}

func ListView_GetTabStop(obj uintptr) bool {
    ret, _, _ := listView_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetTabStop(obj uintptr, value bool) {
   listView_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func ListView_GetViewStyle(obj uintptr) TViewStyle {
    ret, _, _ := listView_GetViewStyle.Call(obj)
    return TViewStyle(ret)
}

func ListView_SetViewStyle(obj uintptr, value TViewStyle) {
   listView_SetViewStyle.Call(obj, uintptr(value))
}

func ListView_GetVisible(obj uintptr) bool {
    ret, _, _ := listView_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetVisible(obj uintptr, value bool) {
   listView_SetVisible.Call(obj, GoBoolToDBool(value))
}

func ListView_SetOnChange(obj uintptr, fn interface{}) {
    listView_SetOnChange.Call(obj, addEventToMap(fn))
}

func ListView_SetOnClick(obj uintptr, fn interface{}) {
    listView_SetOnClick.Call(obj, addEventToMap(fn))
}

func ListView_SetOnColumnClick(obj uintptr, fn interface{}) {
    listView_SetOnColumnClick.Call(obj, addEventToMap(fn))
}

func ListView_SetOnColumnRightClick(obj uintptr, fn interface{}) {
    listView_SetOnColumnRightClick.Call(obj, addEventToMap(fn))
}

func ListView_SetOnCompare(obj uintptr, fn interface{}) {
    listView_SetOnCompare.Call(obj, addEventToMap(fn))
}

func ListView_SetOnDblClick(obj uintptr, fn interface{}) {
    listView_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func ListView_SetOnEnter(obj uintptr, fn interface{}) {
    listView_SetOnEnter.Call(obj, addEventToMap(fn))
}

func ListView_SetOnExit(obj uintptr, fn interface{}) {
    listView_SetOnExit.Call(obj, addEventToMap(fn))
}

func ListView_SetOnGetImageIndex(obj uintptr, fn interface{}) {
    listView_SetOnGetImageIndex.Call(obj, addEventToMap(fn))
}

func ListView_SetOnKeyDown(obj uintptr, fn interface{}) {
    listView_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func ListView_SetOnKeyPress(obj uintptr, fn interface{}) {
    listView_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func ListView_SetOnKeyUp(obj uintptr, fn interface{}) {
    listView_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func ListView_SetOnMouseDown(obj uintptr, fn interface{}) {
    listView_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func ListView_SetOnMouseEnter(obj uintptr, fn interface{}) {
    listView_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func ListView_SetOnMouseLeave(obj uintptr, fn interface{}) {
    listView_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func ListView_SetOnMouseMove(obj uintptr, fn interface{}) {
    listView_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func ListView_SetOnMouseUp(obj uintptr, fn interface{}) {
    listView_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func ListView_SetOnResize(obj uintptr, fn interface{}) {
    listView_SetOnResize.Call(obj, addEventToMap(fn))
}

func ListView_SetOnSelectItem(obj uintptr, fn interface{}) {
    listView_SetOnSelectItem.Call(obj, addEventToMap(fn))
}

func ListView_SetOnItemChecked(obj uintptr, fn interface{}) {
    listView_SetOnItemChecked.Call(obj, addEventToMap(fn))
}

func ListView_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := listView_GetCanvas.Call(obj)
    return ret
}

func ListView_GetDropTarget(obj uintptr) uintptr {
    ret, _, _ := listView_GetDropTarget.Call(obj)
    return ret
}

func ListView_SetDropTarget(obj uintptr, value uintptr) {
   listView_SetDropTarget.Call(obj, value)
}

func ListView_GetItemFocused(obj uintptr) uintptr {
    ret, _, _ := listView_GetItemFocused.Call(obj)
    return ret
}

func ListView_SetItemFocused(obj uintptr, value uintptr) {
   listView_SetItemFocused.Call(obj, value)
}

func ListView_GetSelCount(obj uintptr) int32 {
    ret, _, _ := listView_GetSelCount.Call(obj)
    return int32(ret)
}

func ListView_GetSelected(obj uintptr) uintptr {
    ret, _, _ := listView_GetSelected.Call(obj)
    return ret
}

func ListView_SetSelected(obj uintptr, value uintptr) {
   listView_SetSelected.Call(obj, value)
}

func ListView_GetTopItem(obj uintptr) uintptr {
    ret, _, _ := listView_GetTopItem.Call(obj)
    return ret
}

func ListView_GetVisibleRowCount(obj uintptr) int32 {
    ret, _, _ := listView_GetVisibleRowCount.Call(obj)
    return int32(ret)
}

func ListView_GetItemIndex(obj uintptr) int32 {
    ret, _, _ := listView_GetItemIndex.Call(obj)
    return int32(ret)
}

func ListView_SetItemIndex(obj uintptr, value int32) {
   listView_SetItemIndex.Call(obj, uintptr(value))
}

func ListView_GetBrush(obj uintptr) uintptr {
    ret, _, _ := listView_GetBrush.Call(obj)
    return ret
}

func ListView_GetControlCount(obj uintptr) int32 {
    ret, _, _ := listView_GetControlCount.Call(obj)
    return int32(ret)
}

func ListView_GetHandle(obj uintptr) HWND {
    ret, _, _ := listView_GetHandle.Call(obj)
    return HWND(ret)
}

func ListView_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := listView_GetParentWindow.Call(obj)
    return HWND(ret)
}

func ListView_SetParentWindow(obj uintptr, value HWND) {
   listView_SetParentWindow.Call(obj, uintptr(value))
}

func ListView_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    listView_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ListView_SetBoundsRect(obj uintptr, value TRect) {
   listView_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ListView_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := listView_GetClientHeight.Call(obj)
    return int32(ret)
}

func ListView_SetClientHeight(obj uintptr, value int32) {
   listView_SetClientHeight.Call(obj, uintptr(value))
}

func ListView_GetClientRect(obj uintptr) TRect {
    var ret TRect
    listView_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ListView_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := listView_GetClientWidth.Call(obj)
    return int32(ret)
}

func ListView_SetClientWidth(obj uintptr, value int32) {
   listView_SetClientWidth.Call(obj, uintptr(value))
}

func ListView_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := listView_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func ListView_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := listView_GetExplicitTop.Call(obj)
    return int32(ret)
}

func ListView_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := listView_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func ListView_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := listView_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func ListView_GetParent(obj uintptr) uintptr {
    ret, _, _ := listView_GetParent.Call(obj)
    return ret
}

func ListView_SetParent(obj uintptr, value uintptr) {
   listView_SetParent.Call(obj, value)
}

func ListView_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := listView_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func ListView_SetAlignWithMargins(obj uintptr, value bool) {
   listView_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func ListView_GetLeft(obj uintptr) int32 {
    ret, _, _ := listView_GetLeft.Call(obj)
    return int32(ret)
}

func ListView_SetLeft(obj uintptr, value int32) {
   listView_SetLeft.Call(obj, uintptr(value))
}

func ListView_GetTop(obj uintptr) int32 {
    ret, _, _ := listView_GetTop.Call(obj)
    return int32(ret)
}

func ListView_SetTop(obj uintptr, value int32) {
   listView_SetTop.Call(obj, uintptr(value))
}

func ListView_GetWidth(obj uintptr) int32 {
    ret, _, _ := listView_GetWidth.Call(obj)
    return int32(ret)
}

func ListView_SetWidth(obj uintptr, value int32) {
   listView_SetWidth.Call(obj, uintptr(value))
}

func ListView_GetHeight(obj uintptr) int32 {
    ret, _, _ := listView_GetHeight.Call(obj)
    return int32(ret)
}

func ListView_SetHeight(obj uintptr, value int32) {
   listView_SetHeight.Call(obj, uintptr(value))
}

func ListView_GetCursor(obj uintptr) TCursor {
    ret, _, _ := listView_GetCursor.Call(obj)
    return TCursor(ret)
}

func ListView_SetCursor(obj uintptr, value TCursor) {
   listView_SetCursor.Call(obj, uintptr(value))
}

func ListView_GetHint(obj uintptr) string {
    ret, _, _ := listView_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func ListView_SetHint(obj uintptr, value string) {
   listView_SetHint.Call(obj, GoStrToDStr(value))
}

func ListView_GetMargins(obj uintptr) uintptr {
    ret, _, _ := listView_GetMargins.Call(obj)
    return ret
}

func ListView_SetMargins(obj uintptr, value uintptr) {
   listView_SetMargins.Call(obj, value)
}

func ListView_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := listView_GetComponentCount.Call(obj)
    return int32(ret)
}

func ListView_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := listView_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ListView_SetComponentIndex(obj uintptr, value int32) {
   listView_SetComponentIndex.Call(obj, uintptr(value))
}

func ListView_GetOwner(obj uintptr) uintptr {
    ret, _, _ := listView_GetOwner.Call(obj)
    return ret
}

func ListView_GetName(obj uintptr) string {
    ret, _, _ := listView_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ListView_SetName(obj uintptr, value string) {
   listView_SetName.Call(obj, GoStrToDStr(value))
}

func ListView_GetTag(obj uintptr) int {
    ret, _, _ := listView_GetTag.Call(obj)
    return int(ret)
}

func ListView_SetTag(obj uintptr, value int) {
   listView_SetTag.Call(obj, uintptr(value))
}

func ListView_GetColumn(obj uintptr, Index int32) uintptr {
    ret, _, _ := listView_GetColumn.Call(obj, uintptr(Index))
    return ret
}

func ListView_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := listView_GetControls.Call(obj, uintptr(Index))
    return ret
}

func ListView_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := listView_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TTreeView ---------------------------

func TreeView_Create(obj uintptr) uintptr {
    ret, _, _ := treeView_Create.Call(obj)
    return ret
}

func TreeView_Free(obj uintptr) {
    treeView_Free.Call(obj)
}

func TreeView_AlphaSort(obj uintptr, ARecurse bool) bool {
    ret, _, _ := treeView_AlphaSort.Call(obj, GoBoolToDBool(ARecurse) )
    return DBoolToGoBool(ret)
}

func TreeView_FullCollapse(obj uintptr)  {
    treeView_FullCollapse.Call(obj)
}

func TreeView_FullExpand(obj uintptr)  {
    treeView_FullExpand.Call(obj)
}

func TreeView_GetNodeAt(obj uintptr, X int32, Y int32) uintptr {
    ret, _, _ := treeView_GetNodeAt.Call(obj, uintptr(X) , uintptr(Y) )
    return ret
}

func TreeView_IsEditing(obj uintptr) bool {
    ret, _, _ := treeView_IsEditing.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_LoadFromFile(obj uintptr, FileName string)  {
    treeView_LoadFromFile.Call(obj, GoStrToDStr(FileName) )
}

func TreeView_LoadFromStream(obj uintptr, Stream uintptr)  {
    treeView_LoadFromStream.Call(obj, Stream )
}

func TreeView_SaveToFile(obj uintptr, FileName string)  {
    treeView_SaveToFile.Call(obj, GoStrToDStr(FileName) )
}

func TreeView_SaveToStream(obj uintptr, Stream uintptr)  {
    treeView_SaveToStream.Call(obj, Stream )
}

func TreeView_Deselect(obj uintptr, Node uintptr)  {
    treeView_Deselect.Call(obj, Node )
}

func TreeView_Subselect(obj uintptr, Node uintptr, Validate bool)  {
    treeView_Subselect.Call(obj, Node , GoBoolToDBool(Validate) )
}

func TreeView_ClearSelection(obj uintptr, KeepPrimary bool)  {
    treeView_ClearSelection.Call(obj, GoBoolToDBool(KeepPrimary) )
}

func TreeView_FindNextToSelect(obj uintptr) uintptr {
    ret, _, _ := treeView_FindNextToSelect.Call(obj)
    return ret
}

func TreeView_CustomSort(obj uintptr, SortProc PFNTVCOMPARE, Data int, ARecurse bool) bool {
    ret, _, _ := treeView_CustomSort.Call(obj, uintptr(SortProc) , uintptr(Data) , GoBoolToDBool(ARecurse) )
    return DBoolToGoBool(ret)
}

func TreeView_CanFocus(obj uintptr) bool {
    ret, _, _ := treeView_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_FlipChildren(obj uintptr, AllLevels bool)  {
    treeView_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func TreeView_Focused(obj uintptr) bool {
    ret, _, _ := treeView_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_HandleAllocated(obj uintptr) bool {
    ret, _, _ := treeView_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_Invalidate(obj uintptr)  {
    treeView_Invalidate.Call(obj)
}

func TreeView_Realign(obj uintptr)  {
    treeView_Realign.Call(obj)
}

func TreeView_Repaint(obj uintptr)  {
    treeView_Repaint.Call(obj)
}

func TreeView_ScaleBy(obj uintptr, M int32, D int32)  {
    treeView_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func TreeView_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    treeView_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func TreeView_SetFocus(obj uintptr)  {
    treeView_SetFocus.Call(obj)
}

func TreeView_Update(obj uintptr)  {
    treeView_Update.Call(obj)
}

func TreeView_BringToFront(obj uintptr)  {
    treeView_BringToFront.Call(obj)
}

func TreeView_HasParent(obj uintptr) bool {
    ret, _, _ := treeView_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_Hide(obj uintptr)  {
    treeView_Hide.Call(obj)
}

func TreeView_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := treeView_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func TreeView_Refresh(obj uintptr)  {
    treeView_Refresh.Call(obj)
}

func TreeView_SendToBack(obj uintptr)  {
    treeView_SendToBack.Call(obj)
}

func TreeView_Show(obj uintptr)  {
    treeView_Show.Call(obj)
}

func TreeView_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := treeView_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func TreeView_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := treeView_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func TreeView_GetNamePath(obj uintptr) string {
    ret, _, _ := treeView_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func TreeView_Assign(obj uintptr, Source uintptr)  {
    treeView_Assign.Call(obj, Source )
}

func TreeView_ClassName(obj uintptr) string {
    ret, _, _ := treeView_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func TreeView_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := treeView_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func TreeView_GetHashCode(obj uintptr) int32 {
    ret, _, _ := treeView_GetHashCode.Call(obj)
    return int32(ret)
}

func TreeView_ToString(obj uintptr) string {
    ret, _, _ := treeView_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func TreeView_GetAlign(obj uintptr) TAlign {
    ret, _, _ := treeView_GetAlign.Call(obj)
    return TAlign(ret)
}

func TreeView_SetAlign(obj uintptr, value TAlign) {
   treeView_SetAlign.Call(obj, uintptr(value))
}

func TreeView_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := treeView_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func TreeView_SetAnchors(obj uintptr, value TAnchors) {
   treeView_SetAnchors.Call(obj, uintptr(value))
}

func TreeView_GetAutoExpand(obj uintptr) bool {
    ret, _, _ := treeView_GetAutoExpand.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetAutoExpand(obj uintptr, value bool) {
   treeView_SetAutoExpand.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetBevelEdges(obj uintptr) TBevelEdges {
    ret, _, _ := treeView_GetBevelEdges.Call(obj)
    return TBevelEdges(ret)
}

func TreeView_SetBevelEdges(obj uintptr, value TBevelEdges) {
   treeView_SetBevelEdges.Call(obj, uintptr(value))
}

func TreeView_GetBevelInner(obj uintptr) TBevelCut {
    ret, _, _ := treeView_GetBevelInner.Call(obj)
    return TBevelCut(ret)
}

func TreeView_SetBevelInner(obj uintptr, value TBevelCut) {
   treeView_SetBevelInner.Call(obj, uintptr(value))
}

func TreeView_GetBevelOuter(obj uintptr) TBevelCut {
    ret, _, _ := treeView_GetBevelOuter.Call(obj)
    return TBevelCut(ret)
}

func TreeView_SetBevelOuter(obj uintptr, value TBevelCut) {
   treeView_SetBevelOuter.Call(obj, uintptr(value))
}

func TreeView_GetBevelKind(obj uintptr) TBevelKind {
    ret, _, _ := treeView_GetBevelKind.Call(obj)
    return TBevelKind(ret)
}

func TreeView_SetBevelKind(obj uintptr, value TBevelKind) {
   treeView_SetBevelKind.Call(obj, uintptr(value))
}

func TreeView_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := treeView_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func TreeView_SetBiDiMode(obj uintptr, value TBiDiMode) {
   treeView_SetBiDiMode.Call(obj, uintptr(value))
}

func TreeView_GetBorderStyle(obj uintptr) TBorderStyle {
    ret, _, _ := treeView_GetBorderStyle.Call(obj)
    return TBorderStyle(ret)
}

func TreeView_SetBorderStyle(obj uintptr, value TBorderStyle) {
   treeView_SetBorderStyle.Call(obj, uintptr(value))
}

func TreeView_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := treeView_GetBorderWidth.Call(obj)
    return int32(ret)
}

func TreeView_SetBorderWidth(obj uintptr, value int32) {
   treeView_SetBorderWidth.Call(obj, uintptr(value))
}

func TreeView_GetChangeDelay(obj uintptr) int32 {
    ret, _, _ := treeView_GetChangeDelay.Call(obj)
    return int32(ret)
}

func TreeView_SetChangeDelay(obj uintptr, value int32) {
   treeView_SetChangeDelay.Call(obj, uintptr(value))
}

func TreeView_GetColor(obj uintptr) TColor {
    ret, _, _ := treeView_GetColor.Call(obj)
    return TColor(ret)
}

func TreeView_SetColor(obj uintptr, value TColor) {
   treeView_SetColor.Call(obj, uintptr(value))
}

func TreeView_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := treeView_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetDoubleBuffered(obj uintptr, value bool) {
   treeView_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetEnabled(obj uintptr) bool {
    ret, _, _ := treeView_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetEnabled(obj uintptr, value bool) {
   treeView_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetFont(obj uintptr) uintptr {
    ret, _, _ := treeView_GetFont.Call(obj)
    return ret
}

func TreeView_SetFont(obj uintptr, value uintptr) {
   treeView_SetFont.Call(obj, value)
}

func TreeView_GetHideSelection(obj uintptr) bool {
    ret, _, _ := treeView_GetHideSelection.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetHideSelection(obj uintptr, value bool) {
   treeView_SetHideSelection.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetHotTrack(obj uintptr) bool {
    ret, _, _ := treeView_GetHotTrack.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetHotTrack(obj uintptr, value bool) {
   treeView_SetHotTrack.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetImages(obj uintptr) uintptr {
    ret, _, _ := treeView_GetImages.Call(obj)
    return ret
}

func TreeView_SetImages(obj uintptr, value uintptr) {
   treeView_SetImages.Call(obj, value)
}

func TreeView_GetIndent(obj uintptr) int32 {
    ret, _, _ := treeView_GetIndent.Call(obj)
    return int32(ret)
}

func TreeView_SetIndent(obj uintptr, value int32) {
   treeView_SetIndent.Call(obj, uintptr(value))
}

func TreeView_GetMultiSelect(obj uintptr) bool {
    ret, _, _ := treeView_GetMultiSelect.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetMultiSelect(obj uintptr, value bool) {
   treeView_SetMultiSelect.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetMultiSelectStyle(obj uintptr) TMultiSelectStyle {
    ret, _, _ := treeView_GetMultiSelectStyle.Call(obj)
    return TMultiSelectStyle(ret)
}

func TreeView_SetMultiSelectStyle(obj uintptr, value TMultiSelectStyle) {
   treeView_SetMultiSelectStyle.Call(obj, uintptr(value))
}

func TreeView_GetParentColor(obj uintptr) bool {
    ret, _, _ := treeView_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetParentColor(obj uintptr, value bool) {
   treeView_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetParentCtl3D(obj uintptr) bool {
    ret, _, _ := treeView_GetParentCtl3D.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetParentCtl3D(obj uintptr, value bool) {
   treeView_SetParentCtl3D.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := treeView_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetParentDoubleBuffered(obj uintptr, value bool) {
   treeView_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetParentFont(obj uintptr) bool {
    ret, _, _ := treeView_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetParentFont(obj uintptr, value bool) {
   treeView_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := treeView_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetParentShowHint(obj uintptr, value bool) {
   treeView_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := treeView_GetPopupMenu.Call(obj)
    return ret
}

func TreeView_SetPopupMenu(obj uintptr, value uintptr) {
   treeView_SetPopupMenu.Call(obj, value)
}

func TreeView_GetReadOnly(obj uintptr) bool {
    ret, _, _ := treeView_GetReadOnly.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetReadOnly(obj uintptr, value bool) {
   treeView_SetReadOnly.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetRightClickSelect(obj uintptr) bool {
    ret, _, _ := treeView_GetRightClickSelect.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetRightClickSelect(obj uintptr, value bool) {
   treeView_SetRightClickSelect.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetRowSelect(obj uintptr) bool {
    ret, _, _ := treeView_GetRowSelect.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetRowSelect(obj uintptr, value bool) {
   treeView_SetRowSelect.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetShowButtons(obj uintptr) bool {
    ret, _, _ := treeView_GetShowButtons.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetShowButtons(obj uintptr, value bool) {
   treeView_SetShowButtons.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetShowHint(obj uintptr) bool {
    ret, _, _ := treeView_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetShowHint(obj uintptr, value bool) {
   treeView_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetShowLines(obj uintptr) bool {
    ret, _, _ := treeView_GetShowLines.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetShowLines(obj uintptr, value bool) {
   treeView_SetShowLines.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetShowRoot(obj uintptr) bool {
    ret, _, _ := treeView_GetShowRoot.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetShowRoot(obj uintptr, value bool) {
   treeView_SetShowRoot.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetSortType(obj uintptr) TSortType {
    ret, _, _ := treeView_GetSortType.Call(obj)
    return TSortType(ret)
}

func TreeView_SetSortType(obj uintptr, value TSortType) {
   treeView_SetSortType.Call(obj, uintptr(value))
}

func TreeView_GetStateImages(obj uintptr) uintptr {
    ret, _, _ := treeView_GetStateImages.Call(obj)
    return ret
}

func TreeView_SetStateImages(obj uintptr, value uintptr) {
   treeView_SetStateImages.Call(obj, value)
}

func TreeView_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := treeView_GetTabOrder.Call(obj)
    return uint16(ret)
}

func TreeView_SetTabOrder(obj uintptr, value uint16) {
   treeView_SetTabOrder.Call(obj, uintptr(value))
}

func TreeView_GetTabStop(obj uintptr) bool {
    ret, _, _ := treeView_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetTabStop(obj uintptr, value bool) {
   treeView_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetToolTips(obj uintptr) bool {
    ret, _, _ := treeView_GetToolTips.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetToolTips(obj uintptr, value bool) {
   treeView_SetToolTips.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetVisible(obj uintptr) bool {
    ret, _, _ := treeView_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetVisible(obj uintptr, value bool) {
   treeView_SetVisible.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := treeView_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func TreeView_SetStyleElements(obj uintptr, value TStyleElements) {
   treeView_SetStyleElements.Call(obj, uintptr(value))
}

func TreeView_SetOnChange(obj uintptr, fn interface{}) {
    treeView_SetOnChange.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnClick(obj uintptr, fn interface{}) {
    treeView_SetOnClick.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnCompare(obj uintptr, fn interface{}) {
    treeView_SetOnCompare.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnDblClick(obj uintptr, fn interface{}) {
    treeView_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnEnter(obj uintptr, fn interface{}) {
    treeView_SetOnEnter.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnExit(obj uintptr, fn interface{}) {
    treeView_SetOnExit.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnGetImageIndex(obj uintptr, fn interface{}) {
    treeView_SetOnGetImageIndex.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnGetSelectedIndex(obj uintptr, fn interface{}) {
    treeView_SetOnGetSelectedIndex.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnKeyDown(obj uintptr, fn interface{}) {
    treeView_SetOnKeyDown.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnKeyPress(obj uintptr, fn interface{}) {
    treeView_SetOnKeyPress.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnKeyUp(obj uintptr, fn interface{}) {
    treeView_SetOnKeyUp.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnMouseDown(obj uintptr, fn interface{}) {
    treeView_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnMouseEnter(obj uintptr, fn interface{}) {
    treeView_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnMouseLeave(obj uintptr, fn interface{}) {
    treeView_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnMouseMove(obj uintptr, fn interface{}) {
    treeView_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func TreeView_SetOnMouseUp(obj uintptr, fn interface{}) {
    treeView_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func TreeView_GetItems(obj uintptr) uintptr {
    ret, _, _ := treeView_GetItems.Call(obj)
    return ret
}

func TreeView_SetItems(obj uintptr, value uintptr) {
   treeView_SetItems.Call(obj, value)
}

func TreeView_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := treeView_GetCanvas.Call(obj)
    return ret
}

func TreeView_GetDropTarget(obj uintptr) uintptr {
    ret, _, _ := treeView_GetDropTarget.Call(obj)
    return ret
}

func TreeView_SetDropTarget(obj uintptr, value uintptr) {
   treeView_SetDropTarget.Call(obj, value)
}

func TreeView_GetSelected(obj uintptr) uintptr {
    ret, _, _ := treeView_GetSelected.Call(obj)
    return ret
}

func TreeView_SetSelected(obj uintptr, value uintptr) {
   treeView_SetSelected.Call(obj, value)
}

func TreeView_GetTopItem(obj uintptr) uintptr {
    ret, _, _ := treeView_GetTopItem.Call(obj)
    return ret
}

func TreeView_SetTopItem(obj uintptr, value uintptr) {
   treeView_SetTopItem.Call(obj, value)
}

func TreeView_GetSelectionCount(obj uintptr) uint32 {
    ret, _, _ := treeView_GetSelectionCount.Call(obj)
    return uint32(ret)
}

func TreeView_GetBrush(obj uintptr) uintptr {
    ret, _, _ := treeView_GetBrush.Call(obj)
    return ret
}

func TreeView_GetControlCount(obj uintptr) int32 {
    ret, _, _ := treeView_GetControlCount.Call(obj)
    return int32(ret)
}

func TreeView_GetHandle(obj uintptr) HWND {
    ret, _, _ := treeView_GetHandle.Call(obj)
    return HWND(ret)
}

func TreeView_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := treeView_GetParentWindow.Call(obj)
    return HWND(ret)
}

func TreeView_SetParentWindow(obj uintptr, value HWND) {
   treeView_SetParentWindow.Call(obj, uintptr(value))
}

func TreeView_GetAction(obj uintptr) uintptr {
    ret, _, _ := treeView_GetAction.Call(obj)
    return ret
}

func TreeView_SetAction(obj uintptr, value uintptr) {
   treeView_SetAction.Call(obj, value)
}

func TreeView_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    treeView_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func TreeView_SetBoundsRect(obj uintptr, value TRect) {
   treeView_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func TreeView_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := treeView_GetClientHeight.Call(obj)
    return int32(ret)
}

func TreeView_SetClientHeight(obj uintptr, value int32) {
   treeView_SetClientHeight.Call(obj, uintptr(value))
}

func TreeView_GetClientRect(obj uintptr) TRect {
    var ret TRect
    treeView_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func TreeView_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := treeView_GetClientWidth.Call(obj)
    return int32(ret)
}

func TreeView_SetClientWidth(obj uintptr, value int32) {
   treeView_SetClientWidth.Call(obj, uintptr(value))
}

func TreeView_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := treeView_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func TreeView_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := treeView_GetExplicitTop.Call(obj)
    return int32(ret)
}

func TreeView_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := treeView_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func TreeView_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := treeView_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func TreeView_GetParent(obj uintptr) uintptr {
    ret, _, _ := treeView_GetParent.Call(obj)
    return ret
}

func TreeView_SetParent(obj uintptr, value uintptr) {
   treeView_SetParent.Call(obj, value)
}

func TreeView_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := treeView_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeView_SetAlignWithMargins(obj uintptr, value bool) {
   treeView_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func TreeView_GetLeft(obj uintptr) int32 {
    ret, _, _ := treeView_GetLeft.Call(obj)
    return int32(ret)
}

func TreeView_SetLeft(obj uintptr, value int32) {
   treeView_SetLeft.Call(obj, uintptr(value))
}

func TreeView_GetTop(obj uintptr) int32 {
    ret, _, _ := treeView_GetTop.Call(obj)
    return int32(ret)
}

func TreeView_SetTop(obj uintptr, value int32) {
   treeView_SetTop.Call(obj, uintptr(value))
}

func TreeView_GetWidth(obj uintptr) int32 {
    ret, _, _ := treeView_GetWidth.Call(obj)
    return int32(ret)
}

func TreeView_SetWidth(obj uintptr, value int32) {
   treeView_SetWidth.Call(obj, uintptr(value))
}

func TreeView_GetHeight(obj uintptr) int32 {
    ret, _, _ := treeView_GetHeight.Call(obj)
    return int32(ret)
}

func TreeView_SetHeight(obj uintptr, value int32) {
   treeView_SetHeight.Call(obj, uintptr(value))
}

func TreeView_GetCursor(obj uintptr) TCursor {
    ret, _, _ := treeView_GetCursor.Call(obj)
    return TCursor(ret)
}

func TreeView_SetCursor(obj uintptr, value TCursor) {
   treeView_SetCursor.Call(obj, uintptr(value))
}

func TreeView_GetHint(obj uintptr) string {
    ret, _, _ := treeView_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func TreeView_SetHint(obj uintptr, value string) {
   treeView_SetHint.Call(obj, GoStrToDStr(value))
}

func TreeView_GetMargins(obj uintptr) uintptr {
    ret, _, _ := treeView_GetMargins.Call(obj)
    return ret
}

func TreeView_SetMargins(obj uintptr, value uintptr) {
   treeView_SetMargins.Call(obj, value)
}

func TreeView_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := treeView_GetComponentCount.Call(obj)
    return int32(ret)
}

func TreeView_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := treeView_GetComponentIndex.Call(obj)
    return int32(ret)
}

func TreeView_SetComponentIndex(obj uintptr, value int32) {
   treeView_SetComponentIndex.Call(obj, uintptr(value))
}

func TreeView_GetOwner(obj uintptr) uintptr {
    ret, _, _ := treeView_GetOwner.Call(obj)
    return ret
}

func TreeView_GetName(obj uintptr) string {
    ret, _, _ := treeView_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func TreeView_SetName(obj uintptr, value string) {
   treeView_SetName.Call(obj, GoStrToDStr(value))
}

func TreeView_GetTag(obj uintptr) int {
    ret, _, _ := treeView_GetTag.Call(obj)
    return int(ret)
}

func TreeView_SetTag(obj uintptr, value int) {
   treeView_SetTag.Call(obj, uintptr(value))
}

func TreeView_GetSelections(obj uintptr, Index int32) uintptr {
    ret, _, _ := treeView_GetSelections.Call(obj, uintptr(Index))
    return ret
}

func TreeView_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := treeView_GetControls.Call(obj, uintptr(Index))
    return ret
}

func TreeView_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := treeView_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TStatusBar ---------------------------

func StatusBar_Create(obj uintptr) uintptr {
    ret, _, _ := statusBar_Create.Call(obj)
    return ret
}

func StatusBar_Free(obj uintptr) {
    statusBar_Free.Call(obj)
}

func StatusBar_FlipChildren(obj uintptr, AllLevels bool)  {
    statusBar_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func StatusBar_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    statusBar_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func StatusBar_CanFocus(obj uintptr) bool {
    ret, _, _ := statusBar_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_Focused(obj uintptr) bool {
    ret, _, _ := statusBar_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_HandleAllocated(obj uintptr) bool {
    ret, _, _ := statusBar_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_Invalidate(obj uintptr)  {
    statusBar_Invalidate.Call(obj)
}

func StatusBar_Realign(obj uintptr)  {
    statusBar_Realign.Call(obj)
}

func StatusBar_Repaint(obj uintptr)  {
    statusBar_Repaint.Call(obj)
}

func StatusBar_ScaleBy(obj uintptr, M int32, D int32)  {
    statusBar_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func StatusBar_SetFocus(obj uintptr)  {
    statusBar_SetFocus.Call(obj)
}

func StatusBar_Update(obj uintptr)  {
    statusBar_Update.Call(obj)
}

func StatusBar_BringToFront(obj uintptr)  {
    statusBar_BringToFront.Call(obj)
}

func StatusBar_HasParent(obj uintptr) bool {
    ret, _, _ := statusBar_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_Hide(obj uintptr)  {
    statusBar_Hide.Call(obj)
}

func StatusBar_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := statusBar_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func StatusBar_Refresh(obj uintptr)  {
    statusBar_Refresh.Call(obj)
}

func StatusBar_SendToBack(obj uintptr)  {
    statusBar_SendToBack.Call(obj)
}

func StatusBar_Show(obj uintptr)  {
    statusBar_Show.Call(obj)
}

func StatusBar_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := statusBar_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func StatusBar_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := statusBar_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func StatusBar_GetNamePath(obj uintptr) string {
    ret, _, _ := statusBar_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func StatusBar_Assign(obj uintptr, Source uintptr)  {
    statusBar_Assign.Call(obj, Source )
}

func StatusBar_ClassName(obj uintptr) string {
    ret, _, _ := statusBar_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func StatusBar_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := statusBar_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func StatusBar_GetHashCode(obj uintptr) int32 {
    ret, _, _ := statusBar_GetHashCode.Call(obj)
    return int32(ret)
}

func StatusBar_ToString(obj uintptr) string {
    ret, _, _ := statusBar_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func StatusBar_GetAction(obj uintptr) uintptr {
    ret, _, _ := statusBar_GetAction.Call(obj)
    return ret
}

func StatusBar_SetAction(obj uintptr, value uintptr) {
   statusBar_SetAction.Call(obj, value)
}

func StatusBar_GetAutoHint(obj uintptr) bool {
    ret, _, _ := statusBar_GetAutoHint.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetAutoHint(obj uintptr, value bool) {
   statusBar_SetAutoHint.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetAlign(obj uintptr) TAlign {
    ret, _, _ := statusBar_GetAlign.Call(obj)
    return TAlign(ret)
}

func StatusBar_SetAlign(obj uintptr, value TAlign) {
   statusBar_SetAlign.Call(obj, uintptr(value))
}

func StatusBar_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := statusBar_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func StatusBar_SetAnchors(obj uintptr, value TAnchors) {
   statusBar_SetAnchors.Call(obj, uintptr(value))
}

func StatusBar_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := statusBar_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func StatusBar_SetBiDiMode(obj uintptr, value TBiDiMode) {
   statusBar_SetBiDiMode.Call(obj, uintptr(value))
}

func StatusBar_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := statusBar_GetBorderWidth.Call(obj)
    return int32(ret)
}

func StatusBar_SetBorderWidth(obj uintptr, value int32) {
   statusBar_SetBorderWidth.Call(obj, uintptr(value))
}

func StatusBar_GetColor(obj uintptr) TColor {
    ret, _, _ := statusBar_GetColor.Call(obj)
    return TColor(ret)
}

func StatusBar_SetColor(obj uintptr, value TColor) {
   statusBar_SetColor.Call(obj, uintptr(value))
}

func StatusBar_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := statusBar_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetDoubleBuffered(obj uintptr, value bool) {
   statusBar_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetEnabled(obj uintptr) bool {
    ret, _, _ := statusBar_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetEnabled(obj uintptr, value bool) {
   statusBar_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetFont(obj uintptr) uintptr {
    ret, _, _ := statusBar_GetFont.Call(obj)
    return ret
}

func StatusBar_SetFont(obj uintptr, value uintptr) {
   statusBar_SetFont.Call(obj, value)
}

func StatusBar_GetPanels(obj uintptr) uintptr {
    ret, _, _ := statusBar_GetPanels.Call(obj)
    return ret
}

func StatusBar_SetPanels(obj uintptr, value uintptr) {
   statusBar_SetPanels.Call(obj, value)
}

func StatusBar_GetParentColor(obj uintptr) bool {
    ret, _, _ := statusBar_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetParentColor(obj uintptr, value bool) {
   statusBar_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := statusBar_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetParentDoubleBuffered(obj uintptr, value bool) {
   statusBar_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetParentFont(obj uintptr) bool {
    ret, _, _ := statusBar_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetParentFont(obj uintptr, value bool) {
   statusBar_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := statusBar_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetParentShowHint(obj uintptr, value bool) {
   statusBar_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := statusBar_GetPopupMenu.Call(obj)
    return ret
}

func StatusBar_SetPopupMenu(obj uintptr, value uintptr) {
   statusBar_SetPopupMenu.Call(obj, value)
}

func StatusBar_GetShowHint(obj uintptr) bool {
    ret, _, _ := statusBar_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetShowHint(obj uintptr, value bool) {
   statusBar_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetSimplePanel(obj uintptr) bool {
    ret, _, _ := statusBar_GetSimplePanel.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetSimplePanel(obj uintptr, value bool) {
   statusBar_SetSimplePanel.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetSimpleText(obj uintptr) string {
    ret, _, _ := statusBar_GetSimpleText.Call(obj)
    return DStrToGoStr(ret)
}

func StatusBar_SetSimpleText(obj uintptr, value string) {
   statusBar_SetSimpleText.Call(obj, GoStrToDStr(value))
}

func StatusBar_GetSizeGrip(obj uintptr) bool {
    ret, _, _ := statusBar_GetSizeGrip.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetSizeGrip(obj uintptr, value bool) {
   statusBar_SetSizeGrip.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetUseSystemFont(obj uintptr) bool {
    ret, _, _ := statusBar_GetUseSystemFont.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetUseSystemFont(obj uintptr, value bool) {
   statusBar_SetUseSystemFont.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetVisible(obj uintptr) bool {
    ret, _, _ := statusBar_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetVisible(obj uintptr, value bool) {
   statusBar_SetVisible.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := statusBar_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func StatusBar_SetStyleElements(obj uintptr, value TStyleElements) {
   statusBar_SetStyleElements.Call(obj, uintptr(value))
}

func StatusBar_SetOnClick(obj uintptr, fn interface{}) {
    statusBar_SetOnClick.Call(obj, addEventToMap(fn))
}

func StatusBar_SetOnDblClick(obj uintptr, fn interface{}) {
    statusBar_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func StatusBar_SetOnMouseDown(obj uintptr, fn interface{}) {
    statusBar_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func StatusBar_SetOnMouseEnter(obj uintptr, fn interface{}) {
    statusBar_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func StatusBar_SetOnMouseLeave(obj uintptr, fn interface{}) {
    statusBar_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func StatusBar_SetOnMouseMove(obj uintptr, fn interface{}) {
    statusBar_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func StatusBar_SetOnMouseUp(obj uintptr, fn interface{}) {
    statusBar_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func StatusBar_SetOnResize(obj uintptr, fn interface{}) {
    statusBar_SetOnResize.Call(obj, addEventToMap(fn))
}

func StatusBar_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := statusBar_GetCanvas.Call(obj)
    return ret
}

func StatusBar_GetBrush(obj uintptr) uintptr {
    ret, _, _ := statusBar_GetBrush.Call(obj)
    return ret
}

func StatusBar_GetControlCount(obj uintptr) int32 {
    ret, _, _ := statusBar_GetControlCount.Call(obj)
    return int32(ret)
}

func StatusBar_GetHandle(obj uintptr) HWND {
    ret, _, _ := statusBar_GetHandle.Call(obj)
    return HWND(ret)
}

func StatusBar_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := statusBar_GetParentWindow.Call(obj)
    return HWND(ret)
}

func StatusBar_SetParentWindow(obj uintptr, value HWND) {
   statusBar_SetParentWindow.Call(obj, uintptr(value))
}

func StatusBar_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := statusBar_GetTabOrder.Call(obj)
    return uint16(ret)
}

func StatusBar_SetTabOrder(obj uintptr, value uint16) {
   statusBar_SetTabOrder.Call(obj, uintptr(value))
}

func StatusBar_GetTabStop(obj uintptr) bool {
    ret, _, _ := statusBar_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetTabStop(obj uintptr, value bool) {
   statusBar_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    statusBar_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func StatusBar_SetBoundsRect(obj uintptr, value TRect) {
   statusBar_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func StatusBar_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := statusBar_GetClientHeight.Call(obj)
    return int32(ret)
}

func StatusBar_SetClientHeight(obj uintptr, value int32) {
   statusBar_SetClientHeight.Call(obj, uintptr(value))
}

func StatusBar_GetClientRect(obj uintptr) TRect {
    var ret TRect
    statusBar_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func StatusBar_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := statusBar_GetClientWidth.Call(obj)
    return int32(ret)
}

func StatusBar_SetClientWidth(obj uintptr, value int32) {
   statusBar_SetClientWidth.Call(obj, uintptr(value))
}

func StatusBar_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := statusBar_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func StatusBar_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := statusBar_GetExplicitTop.Call(obj)
    return int32(ret)
}

func StatusBar_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := statusBar_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func StatusBar_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := statusBar_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func StatusBar_GetParent(obj uintptr) uintptr {
    ret, _, _ := statusBar_GetParent.Call(obj)
    return ret
}

func StatusBar_SetParent(obj uintptr, value uintptr) {
   statusBar_SetParent.Call(obj, value)
}

func StatusBar_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := statusBar_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func StatusBar_SetAlignWithMargins(obj uintptr, value bool) {
   statusBar_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func StatusBar_GetLeft(obj uintptr) int32 {
    ret, _, _ := statusBar_GetLeft.Call(obj)
    return int32(ret)
}

func StatusBar_SetLeft(obj uintptr, value int32) {
   statusBar_SetLeft.Call(obj, uintptr(value))
}

func StatusBar_GetTop(obj uintptr) int32 {
    ret, _, _ := statusBar_GetTop.Call(obj)
    return int32(ret)
}

func StatusBar_SetTop(obj uintptr, value int32) {
   statusBar_SetTop.Call(obj, uintptr(value))
}

func StatusBar_GetWidth(obj uintptr) int32 {
    ret, _, _ := statusBar_GetWidth.Call(obj)
    return int32(ret)
}

func StatusBar_SetWidth(obj uintptr, value int32) {
   statusBar_SetWidth.Call(obj, uintptr(value))
}

func StatusBar_GetHeight(obj uintptr) int32 {
    ret, _, _ := statusBar_GetHeight.Call(obj)
    return int32(ret)
}

func StatusBar_SetHeight(obj uintptr, value int32) {
   statusBar_SetHeight.Call(obj, uintptr(value))
}

func StatusBar_GetCursor(obj uintptr) TCursor {
    ret, _, _ := statusBar_GetCursor.Call(obj)
    return TCursor(ret)
}

func StatusBar_SetCursor(obj uintptr, value TCursor) {
   statusBar_SetCursor.Call(obj, uintptr(value))
}

func StatusBar_GetHint(obj uintptr) string {
    ret, _, _ := statusBar_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func StatusBar_SetHint(obj uintptr, value string) {
   statusBar_SetHint.Call(obj, GoStrToDStr(value))
}

func StatusBar_GetMargins(obj uintptr) uintptr {
    ret, _, _ := statusBar_GetMargins.Call(obj)
    return ret
}

func StatusBar_SetMargins(obj uintptr, value uintptr) {
   statusBar_SetMargins.Call(obj, value)
}

func StatusBar_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := statusBar_GetComponentCount.Call(obj)
    return int32(ret)
}

func StatusBar_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := statusBar_GetComponentIndex.Call(obj)
    return int32(ret)
}

func StatusBar_SetComponentIndex(obj uintptr, value int32) {
   statusBar_SetComponentIndex.Call(obj, uintptr(value))
}

func StatusBar_GetOwner(obj uintptr) uintptr {
    ret, _, _ := statusBar_GetOwner.Call(obj)
    return ret
}

func StatusBar_GetName(obj uintptr) string {
    ret, _, _ := statusBar_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func StatusBar_SetName(obj uintptr, value string) {
   statusBar_SetName.Call(obj, GoStrToDStr(value))
}

func StatusBar_GetTag(obj uintptr) int {
    ret, _, _ := statusBar_GetTag.Call(obj)
    return int(ret)
}

func StatusBar_SetTag(obj uintptr, value int) {
   statusBar_SetTag.Call(obj, uintptr(value))
}

func StatusBar_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := statusBar_GetControls.Call(obj, uintptr(Index))
    return ret
}

func StatusBar_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := statusBar_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TToolBar ---------------------------

func ToolBar_Create(obj uintptr) uintptr {
    ret, _, _ := toolBar_Create.Call(obj)
    return ret
}

func ToolBar_Free(obj uintptr) {
    toolBar_Free.Call(obj)
}

func ToolBar_FlipChildren(obj uintptr, AllLevels bool)  {
    toolBar_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func ToolBar_CanFocus(obj uintptr) bool {
    ret, _, _ := toolBar_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_Focused(obj uintptr) bool {
    ret, _, _ := toolBar_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_HandleAllocated(obj uintptr) bool {
    ret, _, _ := toolBar_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_Invalidate(obj uintptr)  {
    toolBar_Invalidate.Call(obj)
}

func ToolBar_Realign(obj uintptr)  {
    toolBar_Realign.Call(obj)
}

func ToolBar_Repaint(obj uintptr)  {
    toolBar_Repaint.Call(obj)
}

func ToolBar_ScaleBy(obj uintptr, M int32, D int32)  {
    toolBar_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func ToolBar_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    toolBar_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func ToolBar_SetFocus(obj uintptr)  {
    toolBar_SetFocus.Call(obj)
}

func ToolBar_Update(obj uintptr)  {
    toolBar_Update.Call(obj)
}

func ToolBar_BringToFront(obj uintptr)  {
    toolBar_BringToFront.Call(obj)
}

func ToolBar_HasParent(obj uintptr) bool {
    ret, _, _ := toolBar_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_Hide(obj uintptr)  {
    toolBar_Hide.Call(obj)
}

func ToolBar_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := toolBar_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func ToolBar_Refresh(obj uintptr)  {
    toolBar_Refresh.Call(obj)
}

func ToolBar_SendToBack(obj uintptr)  {
    toolBar_SendToBack.Call(obj)
}

func ToolBar_Show(obj uintptr)  {
    toolBar_Show.Call(obj)
}

func ToolBar_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := toolBar_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func ToolBar_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := toolBar_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ToolBar_GetNamePath(obj uintptr) string {
    ret, _, _ := toolBar_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ToolBar_Assign(obj uintptr, Source uintptr)  {
    toolBar_Assign.Call(obj, Source )
}

func ToolBar_ClassName(obj uintptr) string {
    ret, _, _ := toolBar_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ToolBar_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := toolBar_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ToolBar_GetHashCode(obj uintptr) int32 {
    ret, _, _ := toolBar_GetHashCode.Call(obj)
    return int32(ret)
}

func ToolBar_ToString(obj uintptr) string {
    ret, _, _ := toolBar_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ToolBar_GetButtonCount(obj uintptr) int32 {
    ret, _, _ := toolBar_GetButtonCount.Call(obj)
    return int32(ret)
}

func ToolBar_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetCanvas.Call(obj)
    return ret
}

func ToolBar_GetCustomizeKeyName(obj uintptr) string {
    ret, _, _ := toolBar_GetCustomizeKeyName.Call(obj)
    return DStrToGoStr(ret)
}

func ToolBar_SetCustomizeKeyName(obj uintptr, value string) {
   toolBar_SetCustomizeKeyName.Call(obj, GoStrToDStr(value))
}

func ToolBar_GetCustomizeValueName(obj uintptr) string {
    ret, _, _ := toolBar_GetCustomizeValueName.Call(obj)
    return DStrToGoStr(ret)
}

func ToolBar_SetCustomizeValueName(obj uintptr, value string) {
   toolBar_SetCustomizeValueName.Call(obj, GoStrToDStr(value))
}

func ToolBar_GetRowCount(obj uintptr) int32 {
    ret, _, _ := toolBar_GetRowCount.Call(obj)
    return int32(ret)
}

func ToolBar_GetAlign(obj uintptr) TAlign {
    ret, _, _ := toolBar_GetAlign.Call(obj)
    return TAlign(ret)
}

func ToolBar_SetAlign(obj uintptr, value TAlign) {
   toolBar_SetAlign.Call(obj, uintptr(value))
}

func ToolBar_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := toolBar_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func ToolBar_SetAnchors(obj uintptr, value TAnchors) {
   toolBar_SetAnchors.Call(obj, uintptr(value))
}

func ToolBar_GetAutoSize(obj uintptr) bool {
    ret, _, _ := toolBar_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetAutoSize(obj uintptr, value bool) {
   toolBar_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := toolBar_GetBorderWidth.Call(obj)
    return int32(ret)
}

func ToolBar_SetBorderWidth(obj uintptr, value int32) {
   toolBar_SetBorderWidth.Call(obj, uintptr(value))
}

func ToolBar_GetButtonHeight(obj uintptr) int32 {
    ret, _, _ := toolBar_GetButtonHeight.Call(obj)
    return int32(ret)
}

func ToolBar_SetButtonHeight(obj uintptr, value int32) {
   toolBar_SetButtonHeight.Call(obj, uintptr(value))
}

func ToolBar_GetButtonWidth(obj uintptr) int32 {
    ret, _, _ := toolBar_GetButtonWidth.Call(obj)
    return int32(ret)
}

func ToolBar_SetButtonWidth(obj uintptr, value int32) {
   toolBar_SetButtonWidth.Call(obj, uintptr(value))
}

func ToolBar_GetCaption(obj uintptr) string {
    ret, _, _ := toolBar_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func ToolBar_SetCaption(obj uintptr, value string) {
   toolBar_SetCaption.Call(obj, GoStrToDStr(value))
}

func ToolBar_GetColor(obj uintptr) TColor {
    ret, _, _ := toolBar_GetColor.Call(obj)
    return TColor(ret)
}

func ToolBar_SetColor(obj uintptr, value TColor) {
   toolBar_SetColor.Call(obj, uintptr(value))
}

func ToolBar_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := toolBar_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetDoubleBuffered(obj uintptr, value bool) {
   toolBar_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetDrawingStyle(obj uintptr) TTBDrawingStyle {
    ret, _, _ := toolBar_GetDrawingStyle.Call(obj)
    return TTBDrawingStyle(ret)
}

func ToolBar_SetDrawingStyle(obj uintptr, value TTBDrawingStyle) {
   toolBar_SetDrawingStyle.Call(obj, uintptr(value))
}

func ToolBar_GetEnabled(obj uintptr) bool {
    ret, _, _ := toolBar_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetEnabled(obj uintptr, value bool) {
   toolBar_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetFlat(obj uintptr) bool {
    ret, _, _ := toolBar_GetFlat.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetFlat(obj uintptr, value bool) {
   toolBar_SetFlat.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetFont(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetFont.Call(obj)
    return ret
}

func ToolBar_SetFont(obj uintptr, value uintptr) {
   toolBar_SetFont.Call(obj, value)
}

func ToolBar_GetGradientEndColor(obj uintptr) TColor {
    ret, _, _ := toolBar_GetGradientEndColor.Call(obj)
    return TColor(ret)
}

func ToolBar_SetGradientEndColor(obj uintptr, value TColor) {
   toolBar_SetGradientEndColor.Call(obj, uintptr(value))
}

func ToolBar_GetGradientStartColor(obj uintptr) TColor {
    ret, _, _ := toolBar_GetGradientStartColor.Call(obj)
    return TColor(ret)
}

func ToolBar_SetGradientStartColor(obj uintptr, value TColor) {
   toolBar_SetGradientStartColor.Call(obj, uintptr(value))
}

func ToolBar_GetHeight(obj uintptr) int32 {
    ret, _, _ := toolBar_GetHeight.Call(obj)
    return int32(ret)
}

func ToolBar_SetHeight(obj uintptr, value int32) {
   toolBar_SetHeight.Call(obj, uintptr(value))
}

func ToolBar_GetHideClippedButtons(obj uintptr) bool {
    ret, _, _ := toolBar_GetHideClippedButtons.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetHideClippedButtons(obj uintptr, value bool) {
   toolBar_SetHideClippedButtons.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetHotImages(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetHotImages.Call(obj)
    return ret
}

func ToolBar_SetHotImages(obj uintptr, value uintptr) {
   toolBar_SetHotImages.Call(obj, value)
}

func ToolBar_GetImages(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetImages.Call(obj)
    return ret
}

func ToolBar_SetImages(obj uintptr, value uintptr) {
   toolBar_SetImages.Call(obj, value)
}

func ToolBar_GetIndent(obj uintptr) int32 {
    ret, _, _ := toolBar_GetIndent.Call(obj)
    return int32(ret)
}

func ToolBar_SetIndent(obj uintptr, value int32) {
   toolBar_SetIndent.Call(obj, uintptr(value))
}

func ToolBar_GetList(obj uintptr) bool {
    ret, _, _ := toolBar_GetList.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetList(obj uintptr, value bool) {
   toolBar_SetList.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetMenu(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetMenu.Call(obj)
    return ret
}

func ToolBar_SetMenu(obj uintptr, value uintptr) {
   toolBar_SetMenu.Call(obj, value)
}

func ToolBar_GetGradientDirection(obj uintptr) TGradientDirection {
    ret, _, _ := toolBar_GetGradientDirection.Call(obj)
    return TGradientDirection(ret)
}

func ToolBar_SetGradientDirection(obj uintptr, value TGradientDirection) {
   toolBar_SetGradientDirection.Call(obj, uintptr(value))
}

func ToolBar_GetGradientDrawingOptions(obj uintptr) TTBGradientDrawingOptions {
    ret, _, _ := toolBar_GetGradientDrawingOptions.Call(obj)
    return TTBGradientDrawingOptions(ret)
}

func ToolBar_SetGradientDrawingOptions(obj uintptr, value TTBGradientDrawingOptions) {
   toolBar_SetGradientDrawingOptions.Call(obj, uintptr(value))
}

func ToolBar_GetParentColor(obj uintptr) bool {
    ret, _, _ := toolBar_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetParentColor(obj uintptr, value bool) {
   toolBar_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := toolBar_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetParentDoubleBuffered(obj uintptr, value bool) {
   toolBar_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetParentFont(obj uintptr) bool {
    ret, _, _ := toolBar_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetParentFont(obj uintptr, value bool) {
   toolBar_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := toolBar_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetParentShowHint(obj uintptr, value bool) {
   toolBar_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetPopupMenu.Call(obj)
    return ret
}

func ToolBar_SetPopupMenu(obj uintptr, value uintptr) {
   toolBar_SetPopupMenu.Call(obj, value)
}

func ToolBar_GetShowCaptions(obj uintptr) bool {
    ret, _, _ := toolBar_GetShowCaptions.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetShowCaptions(obj uintptr, value bool) {
   toolBar_SetShowCaptions.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetShowHint(obj uintptr) bool {
    ret, _, _ := toolBar_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetShowHint(obj uintptr, value bool) {
   toolBar_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := toolBar_GetTabOrder.Call(obj)
    return uint16(ret)
}

func ToolBar_SetTabOrder(obj uintptr, value uint16) {
   toolBar_SetTabOrder.Call(obj, uintptr(value))
}

func ToolBar_GetTabStop(obj uintptr) bool {
    ret, _, _ := toolBar_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetTabStop(obj uintptr, value bool) {
   toolBar_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetTransparent(obj uintptr) bool {
    ret, _, _ := toolBar_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetTransparent(obj uintptr, value bool) {
   toolBar_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetVisible(obj uintptr) bool {
    ret, _, _ := toolBar_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetVisible(obj uintptr, value bool) {
   toolBar_SetVisible.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := toolBar_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func ToolBar_SetStyleElements(obj uintptr, value TStyleElements) {
   toolBar_SetStyleElements.Call(obj, uintptr(value))
}

func ToolBar_GetWrapable(obj uintptr) bool {
    ret, _, _ := toolBar_GetWrapable.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetWrapable(obj uintptr, value bool) {
   toolBar_SetWrapable.Call(obj, GoBoolToDBool(value))
}

func ToolBar_SetOnClick(obj uintptr, fn interface{}) {
    toolBar_SetOnClick.Call(obj, addEventToMap(fn))
}

func ToolBar_SetOnDblClick(obj uintptr, fn interface{}) {
    toolBar_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func ToolBar_SetOnEnter(obj uintptr, fn interface{}) {
    toolBar_SetOnEnter.Call(obj, addEventToMap(fn))
}

func ToolBar_SetOnExit(obj uintptr, fn interface{}) {
    toolBar_SetOnExit.Call(obj, addEventToMap(fn))
}

func ToolBar_SetOnMouseDown(obj uintptr, fn interface{}) {
    toolBar_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func ToolBar_SetOnMouseEnter(obj uintptr, fn interface{}) {
    toolBar_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func ToolBar_SetOnMouseLeave(obj uintptr, fn interface{}) {
    toolBar_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func ToolBar_SetOnMouseMove(obj uintptr, fn interface{}) {
    toolBar_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func ToolBar_SetOnMouseUp(obj uintptr, fn interface{}) {
    toolBar_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func ToolBar_SetOnResize(obj uintptr, fn interface{}) {
    toolBar_SetOnResize.Call(obj, addEventToMap(fn))
}

func ToolBar_GetBrush(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetBrush.Call(obj)
    return ret
}

func ToolBar_GetControlCount(obj uintptr) int32 {
    ret, _, _ := toolBar_GetControlCount.Call(obj)
    return int32(ret)
}

func ToolBar_GetHandle(obj uintptr) HWND {
    ret, _, _ := toolBar_GetHandle.Call(obj)
    return HWND(ret)
}

func ToolBar_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := toolBar_GetParentWindow.Call(obj)
    return HWND(ret)
}

func ToolBar_SetParentWindow(obj uintptr, value HWND) {
   toolBar_SetParentWindow.Call(obj, uintptr(value))
}

func ToolBar_GetAction(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetAction.Call(obj)
    return ret
}

func ToolBar_SetAction(obj uintptr, value uintptr) {
   toolBar_SetAction.Call(obj, value)
}

func ToolBar_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := toolBar_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func ToolBar_SetBiDiMode(obj uintptr, value TBiDiMode) {
   toolBar_SetBiDiMode.Call(obj, uintptr(value))
}

func ToolBar_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    toolBar_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ToolBar_SetBoundsRect(obj uintptr, value TRect) {
   toolBar_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ToolBar_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := toolBar_GetClientHeight.Call(obj)
    return int32(ret)
}

func ToolBar_SetClientHeight(obj uintptr, value int32) {
   toolBar_SetClientHeight.Call(obj, uintptr(value))
}

func ToolBar_GetClientRect(obj uintptr) TRect {
    var ret TRect
    toolBar_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ToolBar_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := toolBar_GetClientWidth.Call(obj)
    return int32(ret)
}

func ToolBar_SetClientWidth(obj uintptr, value int32) {
   toolBar_SetClientWidth.Call(obj, uintptr(value))
}

func ToolBar_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := toolBar_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func ToolBar_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := toolBar_GetExplicitTop.Call(obj)
    return int32(ret)
}

func ToolBar_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := toolBar_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func ToolBar_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := toolBar_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func ToolBar_GetParent(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetParent.Call(obj)
    return ret
}

func ToolBar_SetParent(obj uintptr, value uintptr) {
   toolBar_SetParent.Call(obj, value)
}

func ToolBar_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := toolBar_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolBar_SetAlignWithMargins(obj uintptr, value bool) {
   toolBar_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func ToolBar_GetLeft(obj uintptr) int32 {
    ret, _, _ := toolBar_GetLeft.Call(obj)
    return int32(ret)
}

func ToolBar_SetLeft(obj uintptr, value int32) {
   toolBar_SetLeft.Call(obj, uintptr(value))
}

func ToolBar_GetTop(obj uintptr) int32 {
    ret, _, _ := toolBar_GetTop.Call(obj)
    return int32(ret)
}

func ToolBar_SetTop(obj uintptr, value int32) {
   toolBar_SetTop.Call(obj, uintptr(value))
}

func ToolBar_GetWidth(obj uintptr) int32 {
    ret, _, _ := toolBar_GetWidth.Call(obj)
    return int32(ret)
}

func ToolBar_SetWidth(obj uintptr, value int32) {
   toolBar_SetWidth.Call(obj, uintptr(value))
}

func ToolBar_GetCursor(obj uintptr) TCursor {
    ret, _, _ := toolBar_GetCursor.Call(obj)
    return TCursor(ret)
}

func ToolBar_SetCursor(obj uintptr, value TCursor) {
   toolBar_SetCursor.Call(obj, uintptr(value))
}

func ToolBar_GetHint(obj uintptr) string {
    ret, _, _ := toolBar_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func ToolBar_SetHint(obj uintptr, value string) {
   toolBar_SetHint.Call(obj, GoStrToDStr(value))
}

func ToolBar_GetMargins(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetMargins.Call(obj)
    return ret
}

func ToolBar_SetMargins(obj uintptr, value uintptr) {
   toolBar_SetMargins.Call(obj, value)
}

func ToolBar_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := toolBar_GetComponentCount.Call(obj)
    return int32(ret)
}

func ToolBar_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := toolBar_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ToolBar_SetComponentIndex(obj uintptr, value int32) {
   toolBar_SetComponentIndex.Call(obj, uintptr(value))
}

func ToolBar_GetOwner(obj uintptr) uintptr {
    ret, _, _ := toolBar_GetOwner.Call(obj)
    return ret
}

func ToolBar_GetName(obj uintptr) string {
    ret, _, _ := toolBar_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ToolBar_SetName(obj uintptr, value string) {
   toolBar_SetName.Call(obj, GoStrToDStr(value))
}

func ToolBar_GetTag(obj uintptr) int {
    ret, _, _ := toolBar_GetTag.Call(obj)
    return int(ret)
}

func ToolBar_SetTag(obj uintptr, value int) {
   toolBar_SetTag.Call(obj, uintptr(value))
}

func ToolBar_GetButtons(obj uintptr, Index int32) uintptr {
    ret, _, _ := toolBar_GetButtons.Call(obj, uintptr(Index))
    return ret
}

func ToolBar_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := toolBar_GetControls.Call(obj, uintptr(Index))
    return ret
}

func ToolBar_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := toolBar_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TIcon ---------------------------

func Icon_Create() uintptr {
    ret, _, _ := icon_Create.Call()
    return ret
}

func Icon_Free(obj uintptr) {
    icon_Free.Call(obj)
}

func Icon_Assign(obj uintptr, Source uintptr)  {
    icon_Assign.Call(obj, Source )
}

func Icon_HandleAllocated(obj uintptr) bool {
    ret, _, _ := icon_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Icon_LoadFromStream(obj uintptr, Stream uintptr)  {
    icon_LoadFromStream.Call(obj, Stream )
}

func Icon_SaveToStream(obj uintptr, Stream uintptr)  {
    icon_SaveToStream.Call(obj, Stream )
}

func Icon_SetSize(obj uintptr, AWidth int32, AHeight int32)  {
    icon_SetSize.Call(obj, uintptr(AWidth) , uintptr(AHeight) )
}

func Icon_LoadFromResourceName(obj uintptr, Instance uintptr, ResName string)  {
    icon_LoadFromResourceName.Call(obj, Instance , GoStrToDStr(ResName) )
}

func Icon_LoadFromResourceID(obj uintptr, Instance uintptr, ResID int32)  {
    icon_LoadFromResourceID.Call(obj, Instance , uintptr(ResID) )
}

func Icon_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := icon_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Icon_LoadFromFile(obj uintptr, Filename string)  {
    icon_LoadFromFile.Call(obj, GoStrToDStr(Filename) )
}

func Icon_SaveToFile(obj uintptr, Filename string)  {
    icon_SaveToFile.Call(obj, GoStrToDStr(Filename) )
}

func Icon_GetNamePath(obj uintptr) string {
    ret, _, _ := icon_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Icon_ClassName(obj uintptr) string {
    ret, _, _ := icon_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Icon_GetHashCode(obj uintptr) int32 {
    ret, _, _ := icon_GetHashCode.Call(obj)
    return int32(ret)
}

func Icon_ToString(obj uintptr) string {
    ret, _, _ := icon_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Icon_GetHandle(obj uintptr) HICON {
    ret, _, _ := icon_GetHandle.Call(obj)
    return HICON(ret)
}

func Icon_SetHandle(obj uintptr, value HICON) {
   icon_SetHandle.Call(obj, uintptr(value))
}

func Icon_GetEmpty(obj uintptr) bool {
    ret, _, _ := icon_GetEmpty.Call(obj)
    return DBoolToGoBool(ret)
}

func Icon_GetHeight(obj uintptr) int32 {
    ret, _, _ := icon_GetHeight.Call(obj)
    return int32(ret)
}

func Icon_SetHeight(obj uintptr, value int32) {
   icon_SetHeight.Call(obj, uintptr(value))
}

func Icon_GetModified(obj uintptr) bool {
    ret, _, _ := icon_GetModified.Call(obj)
    return DBoolToGoBool(ret)
}

func Icon_SetModified(obj uintptr, value bool) {
   icon_SetModified.Call(obj, GoBoolToDBool(value))
}

func Icon_GetPaletteModified(obj uintptr) bool {
    ret, _, _ := icon_GetPaletteModified.Call(obj)
    return DBoolToGoBool(ret)
}

func Icon_SetPaletteModified(obj uintptr, value bool) {
   icon_SetPaletteModified.Call(obj, GoBoolToDBool(value))
}

func Icon_GetTransparent(obj uintptr) bool {
    ret, _, _ := icon_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func Icon_SetTransparent(obj uintptr, value bool) {
   icon_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func Icon_GetWidth(obj uintptr) int32 {
    ret, _, _ := icon_GetWidth.Call(obj)
    return int32(ret)
}

func Icon_SetWidth(obj uintptr, value int32) {
   icon_SetWidth.Call(obj, uintptr(value))
}

func Icon_SetOnChange(obj uintptr, fn interface{}) {
    icon_SetOnChange.Call(obj, addEventToMap(fn))
}


//--------------------------- TBitmap ---------------------------

func Bitmap_Create() uintptr {
    ret, _, _ := bitmap_Create.Call()
    return ret
}

func Bitmap_Free(obj uintptr) {
    bitmap_Free.Call(obj)
}

func Bitmap_Assign(obj uintptr, Source uintptr)  {
    bitmap_Assign.Call(obj, Source )
}

func Bitmap_HandleAllocated(obj uintptr) bool {
    ret, _, _ := bitmap_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Bitmap_LoadFromStream(obj uintptr, Stream uintptr)  {
    bitmap_LoadFromStream.Call(obj, Stream )
}

func Bitmap_SaveToStream(obj uintptr, Stream uintptr)  {
    bitmap_SaveToStream.Call(obj, Stream )
}

func Bitmap_SetSize(obj uintptr, AWidth int32, AHeight int32)  {
    bitmap_SetSize.Call(obj, uintptr(AWidth) , uintptr(AHeight) )
}

func Bitmap_LoadFromResourceName(obj uintptr, Instance uintptr, ResName string)  {
    bitmap_LoadFromResourceName.Call(obj, Instance , GoStrToDStr(ResName) )
}

func Bitmap_LoadFromResourceID(obj uintptr, Instance uintptr, ResID int32)  {
    bitmap_LoadFromResourceID.Call(obj, Instance , uintptr(ResID) )
}

func Bitmap_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := bitmap_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Bitmap_LoadFromFile(obj uintptr, Filename string)  {
    bitmap_LoadFromFile.Call(obj, GoStrToDStr(Filename) )
}

func Bitmap_SaveToFile(obj uintptr, Filename string)  {
    bitmap_SaveToFile.Call(obj, GoStrToDStr(Filename) )
}

func Bitmap_GetNamePath(obj uintptr) string {
    ret, _, _ := bitmap_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Bitmap_ClassName(obj uintptr) string {
    ret, _, _ := bitmap_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Bitmap_GetHashCode(obj uintptr) int32 {
    ret, _, _ := bitmap_GetHashCode.Call(obj)
    return int32(ret)
}

func Bitmap_ToString(obj uintptr) string {
    ret, _, _ := bitmap_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Bitmap_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := bitmap_GetCanvas.Call(obj)
    return ret
}

func Bitmap_GetHandle(obj uintptr) HBITMAP {
    ret, _, _ := bitmap_GetHandle.Call(obj)
    return HBITMAP(ret)
}

func Bitmap_SetHandle(obj uintptr, value HBITMAP) {
   bitmap_SetHandle.Call(obj, uintptr(value))
}

func Bitmap_GetPixelFormat(obj uintptr) TPixelFormat {
    ret, _, _ := bitmap_GetPixelFormat.Call(obj)
    return TPixelFormat(ret)
}

func Bitmap_SetPixelFormat(obj uintptr, value TPixelFormat) {
   bitmap_SetPixelFormat.Call(obj, uintptr(value))
}

func Bitmap_GetTransparentColor(obj uintptr) TColor {
    ret, _, _ := bitmap_GetTransparentColor.Call(obj)
    return TColor(ret)
}

func Bitmap_SetTransparentColor(obj uintptr, value TColor) {
   bitmap_SetTransparentColor.Call(obj, uintptr(value))
}

func Bitmap_GetEmpty(obj uintptr) bool {
    ret, _, _ := bitmap_GetEmpty.Call(obj)
    return DBoolToGoBool(ret)
}

func Bitmap_GetHeight(obj uintptr) int32 {
    ret, _, _ := bitmap_GetHeight.Call(obj)
    return int32(ret)
}

func Bitmap_SetHeight(obj uintptr, value int32) {
   bitmap_SetHeight.Call(obj, uintptr(value))
}

func Bitmap_GetModified(obj uintptr) bool {
    ret, _, _ := bitmap_GetModified.Call(obj)
    return DBoolToGoBool(ret)
}

func Bitmap_SetModified(obj uintptr, value bool) {
   bitmap_SetModified.Call(obj, GoBoolToDBool(value))
}

func Bitmap_GetPaletteModified(obj uintptr) bool {
    ret, _, _ := bitmap_GetPaletteModified.Call(obj)
    return DBoolToGoBool(ret)
}

func Bitmap_SetPaletteModified(obj uintptr, value bool) {
   bitmap_SetPaletteModified.Call(obj, GoBoolToDBool(value))
}

func Bitmap_GetTransparent(obj uintptr) bool {
    ret, _, _ := bitmap_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func Bitmap_SetTransparent(obj uintptr, value bool) {
   bitmap_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func Bitmap_GetWidth(obj uintptr) int32 {
    ret, _, _ := bitmap_GetWidth.Call(obj)
    return int32(ret)
}

func Bitmap_SetWidth(obj uintptr, value int32) {
   bitmap_SetWidth.Call(obj, uintptr(value))
}

func Bitmap_SetOnChange(obj uintptr, fn interface{}) {
    bitmap_SetOnChange.Call(obj, addEventToMap(fn))
}

func Bitmap_GetScanLine(obj uintptr, Row int32) uintptr {
    ret, _, _ := bitmap_GetScanLine.Call(obj, uintptr(Row))
    return ret
}


//--------------------------- TMemoryStream ---------------------------

func MemoryStream_Create() uintptr {
    ret, _, _ := memoryStream_Create.Call()
    return ret
}

func MemoryStream_Free(obj uintptr) {
    memoryStream_Free.Call(obj)
}

func MemoryStream_Clear(obj uintptr)  {
    memoryStream_Clear.Call(obj)
}

func MemoryStream_LoadFromStream(obj uintptr, Stream uintptr)  {
    memoryStream_LoadFromStream.Call(obj, Stream )
}

func MemoryStream_LoadFromFile(obj uintptr, FileName string)  {
    memoryStream_LoadFromFile.Call(obj, GoStrToDStr(FileName) )
}

func MemoryStream_Seek(obj uintptr, Offset int64, Origin TSeekOrigin) int64 {
    var ret int64
    memoryStream_Seek.Call(obj, uintptr(unsafe.Pointer(&Offset)), uintptr(Origin) , uintptr(unsafe.Pointer(&ret)))
    return ret
}

func MemoryStream_SaveToStream(obj uintptr, Stream uintptr)  {
    memoryStream_SaveToStream.Call(obj, Stream )
}

func MemoryStream_SaveToFile(obj uintptr, FileName string)  {
    memoryStream_SaveToFile.Call(obj, GoStrToDStr(FileName) )
}

func MemoryStream_CopyFrom(obj uintptr, Source uintptr, Count int64) int64 {
    var ret int64
    memoryStream_CopyFrom.Call(obj, Source , uintptr(unsafe.Pointer(&Count)), uintptr(unsafe.Pointer(&ret)))
    return ret
}

func MemoryStream_ClassName(obj uintptr) string {
    ret, _, _ := memoryStream_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func MemoryStream_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := memoryStream_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func MemoryStream_GetHashCode(obj uintptr) int32 {
    ret, _, _ := memoryStream_GetHashCode.Call(obj)
    return int32(ret)
}

func MemoryStream_ToString(obj uintptr) string {
    ret, _, _ := memoryStream_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func MemoryStream_GetMemory(obj uintptr) uintptr {
    ret, _, _ := memoryStream_GetMemory.Call(obj)
    return ret
}

func MemoryStream_GetPosition(obj uintptr) int64 {
    var ret int64
    memoryStream_GetPosition.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func MemoryStream_SetPosition(obj uintptr, value int64) {
   memoryStream_SetPosition.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func MemoryStream_GetSize(obj uintptr) int64 {
    var ret int64
    memoryStream_GetSize.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func MemoryStream_SetSize(obj uintptr, value int64) {
   memoryStream_SetSize.Call(obj, uintptr(unsafe.Pointer(&value)))
}


//--------------------------- TFont ---------------------------

func Font_Create() uintptr {
    ret, _, _ := font_Create.Call()
    return ret
}

func Font_Free(obj uintptr) {
    font_Free.Call(obj)
}

func Font_Assign(obj uintptr, Source uintptr)  {
    font_Assign.Call(obj, Source )
}

func Font_HandleAllocated(obj uintptr) bool {
    ret, _, _ := font_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Font_GetNamePath(obj uintptr) string {
    ret, _, _ := font_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Font_ClassName(obj uintptr) string {
    ret, _, _ := font_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Font_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := font_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Font_GetHashCode(obj uintptr) int32 {
    ret, _, _ := font_GetHashCode.Call(obj)
    return int32(ret)
}

func Font_ToString(obj uintptr) string {
    ret, _, _ := font_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Font_GetHandle(obj uintptr) HFONT {
    ret, _, _ := font_GetHandle.Call(obj)
    return HFONT(ret)
}

func Font_SetHandle(obj uintptr, value HFONT) {
   font_SetHandle.Call(obj, uintptr(value))
}

func Font_GetPixelsPerInch(obj uintptr) int32 {
    ret, _, _ := font_GetPixelsPerInch.Call(obj)
    return int32(ret)
}

func Font_SetPixelsPerInch(obj uintptr, value int32) {
   font_SetPixelsPerInch.Call(obj, uintptr(value))
}

func Font_GetCharset(obj uintptr) TFontCharset {
    ret, _, _ := font_GetCharset.Call(obj)
    return TFontCharset(ret)
}

func Font_SetCharset(obj uintptr, value TFontCharset) {
   font_SetCharset.Call(obj, uintptr(value))
}

func Font_GetColor(obj uintptr) TColor {
    ret, _, _ := font_GetColor.Call(obj)
    return TColor(ret)
}

func Font_SetColor(obj uintptr, value TColor) {
   font_SetColor.Call(obj, uintptr(value))
}

func Font_GetHeight(obj uintptr) int32 {
    ret, _, _ := font_GetHeight.Call(obj)
    return int32(ret)
}

func Font_SetHeight(obj uintptr, value int32) {
   font_SetHeight.Call(obj, uintptr(value))
}

func Font_GetName(obj uintptr) string {
    ret, _, _ := font_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Font_SetName(obj uintptr, value string) {
   font_SetName.Call(obj, GoStrToDStr(value))
}

func Font_GetOrientation(obj uintptr) int32 {
    ret, _, _ := font_GetOrientation.Call(obj)
    return int32(ret)
}

func Font_SetOrientation(obj uintptr, value int32) {
   font_SetOrientation.Call(obj, uintptr(value))
}

func Font_GetPitch(obj uintptr) TFontPitch {
    ret, _, _ := font_GetPitch.Call(obj)
    return TFontPitch(ret)
}

func Font_SetPitch(obj uintptr, value TFontPitch) {
   font_SetPitch.Call(obj, uintptr(value))
}

func Font_GetSize(obj uintptr) int32 {
    ret, _, _ := font_GetSize.Call(obj)
    return int32(ret)
}

func Font_SetSize(obj uintptr, value int32) {
   font_SetSize.Call(obj, uintptr(value))
}

func Font_GetStyle(obj uintptr) TFontStyles {
    ret, _, _ := font_GetStyle.Call(obj)
    return TFontStyles(ret)
}

func Font_SetStyle(obj uintptr, value TFontStyles) {
   font_SetStyle.Call(obj, uintptr(value))
}

func Font_GetQuality(obj uintptr) TFontQuality {
    ret, _, _ := font_GetQuality.Call(obj)
    return TFontQuality(ret)
}

func Font_SetQuality(obj uintptr, value TFontQuality) {
   font_SetQuality.Call(obj, uintptr(value))
}

func Font_SetOnChange(obj uintptr, fn interface{}) {
    font_SetOnChange.Call(obj, addEventToMap(fn))
}


//--------------------------- TStrings ---------------------------

func Strings_Create() uintptr {
    ret, _, _ := strings_Create.Call()
    return ret
}

func Strings_Free(obj uintptr) {
    strings_Free.Call(obj)
}

func Strings_Add(obj uintptr, S string) int32 {
    ret, _, _ := strings_Add.Call(obj, GoStrToDStr(S) )
    return int32(ret)
}

func Strings_AddObject(obj uintptr, S string, AObject uintptr) int32 {
    ret, _, _ := strings_AddObject.Call(obj, GoStrToDStr(S) , AObject )
    return int32(ret)
}

func Strings_Append(obj uintptr, S string)  {
    strings_Append.Call(obj, GoStrToDStr(S) )
}

func Strings_Assign(obj uintptr, Source uintptr)  {
    strings_Assign.Call(obj, Source )
}

func Strings_BeginUpdate(obj uintptr)  {
    strings_BeginUpdate.Call(obj)
}

func Strings_Clear(obj uintptr)  {
    strings_Clear.Call(obj)
}

func Strings_Delete(obj uintptr, Index int32)  {
    strings_Delete.Call(obj, uintptr(Index) )
}

func Strings_EndUpdate(obj uintptr)  {
    strings_EndUpdate.Call(obj)
}

func Strings_Equals(obj uintptr, Strings uintptr) bool {
    ret, _, _ := strings_Equals.Call(obj, Strings )
    return DBoolToGoBool(ret)
}

func Strings_IndexOf(obj uintptr, S string) int32 {
    ret, _, _ := strings_IndexOf.Call(obj, GoStrToDStr(S) )
    return int32(ret)
}

func Strings_IndexOfName(obj uintptr, Name string) int32 {
    ret, _, _ := strings_IndexOfName.Call(obj, GoStrToDStr(Name) )
    return int32(ret)
}

func Strings_IndexOfObject(obj uintptr, AObject uintptr) int32 {
    ret, _, _ := strings_IndexOfObject.Call(obj, AObject )
    return int32(ret)
}

func Strings_Insert(obj uintptr, Index int32, S string)  {
    strings_Insert.Call(obj, uintptr(Index) , GoStrToDStr(S) )
}

func Strings_InsertObject(obj uintptr, Index int32, S string, AObject uintptr)  {
    strings_InsertObject.Call(obj, uintptr(Index) , GoStrToDStr(S) , AObject )
}

func Strings_LoadFromFile(obj uintptr, FileName string)  {
    strings_LoadFromFile.Call(obj, GoStrToDStr(FileName) )
}

func Strings_LoadFromStream(obj uintptr, Stream uintptr)  {
    strings_LoadFromStream.Call(obj, Stream )
}

func Strings_Move(obj uintptr, CurIndex int32, NewIndex int32)  {
    strings_Move.Call(obj, uintptr(CurIndex) , uintptr(NewIndex) )
}

func Strings_SaveToFile(obj uintptr, FileName string)  {
    strings_SaveToFile.Call(obj, GoStrToDStr(FileName) )
}

func Strings_SaveToStream(obj uintptr, Stream uintptr)  {
    strings_SaveToStream.Call(obj, Stream )
}

func Strings_GetNamePath(obj uintptr) string {
    ret, _, _ := strings_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Strings_ClassName(obj uintptr) string {
    ret, _, _ := strings_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Strings_GetHashCode(obj uintptr) int32 {
    ret, _, _ := strings_GetHashCode.Call(obj)
    return int32(ret)
}

func Strings_ToString(obj uintptr) string {
    ret, _, _ := strings_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Strings_GetCommaText(obj uintptr) string {
    ret, _, _ := strings_GetCommaText.Call(obj)
    return DStrToGoStr(ret)
}

func Strings_SetCommaText(obj uintptr, value string) {
   strings_SetCommaText.Call(obj, GoStrToDStr(value))
}

func Strings_GetDelimiter(obj uintptr) uint16 {
    ret, _, _ := strings_GetDelimiter.Call(obj)
    return uint16(ret)
}

func Strings_SetDelimiter(obj uintptr, value uint16) {
   strings_SetDelimiter.Call(obj, uintptr(value))
}

func Strings_GetText(obj uintptr) string {
    ret, _, _ := strings_GetText.Call(obj)
    return DStrToGoStr(ret)
}

func Strings_SetText(obj uintptr, value string) {
   strings_SetText.Call(obj, GoStrToDStr(value))
}

func Strings_GetWriteBOM(obj uintptr) bool {
    ret, _, _ := strings_GetWriteBOM.Call(obj)
    return DBoolToGoBool(ret)
}

func Strings_SetWriteBOM(obj uintptr, value bool) {
   strings_SetWriteBOM.Call(obj, GoBoolToDBool(value))
}

func Strings_GetOptions(obj uintptr) TStringsOptions {
    ret, _, _ := strings_GetOptions.Call(obj)
    return TStringsOptions(ret)
}

func Strings_SetOptions(obj uintptr, value TStringsOptions) {
   strings_SetOptions.Call(obj, uintptr(value))
}

func Strings_GetValues(obj uintptr, Name string) string {
    ret, _, _ := strings_GetValues.Call(obj, GoStrToDStr(Name))
    return DStrToGoStr(ret)
}

func Strings_SetValues(obj uintptr, Name string, value string) {
   strings_SetValues.Call(obj, GoStrToDStr(Name), GoStrToDStr(value))
}

func Strings_GetValueFromIndex(obj uintptr, Index int32) string {
    ret, _, _ := strings_GetValueFromIndex.Call(obj, uintptr(Index))
    return DStrToGoStr(ret)
}

func Strings_SetValueFromIndex(obj uintptr, Index int32, value string) {
   strings_SetValueFromIndex.Call(obj, uintptr(Index), GoStrToDStr(value))
}

func Strings_GetStrings(obj uintptr, Index int32) string {
    ret, _, _ := strings_GetStrings.Call(obj, uintptr(Index))
    return DStrToGoStr(ret)
}

func Strings_SetStrings(obj uintptr, Index int32, value string) {
   strings_SetStrings.Call(obj, uintptr(Index), GoStrToDStr(value))
}


//--------------------------- TStringList ---------------------------

func StringList_Create() uintptr {
    ret, _, _ := stringList_Create.Call()
    return ret
}

func StringList_Free(obj uintptr) {
    stringList_Free.Call(obj)
}

func StringList_Add(obj uintptr, S string) int32 {
    ret, _, _ := stringList_Add.Call(obj, GoStrToDStr(S) )
    return int32(ret)
}

func StringList_AddObject(obj uintptr, S string, AObject uintptr) int32 {
    ret, _, _ := stringList_AddObject.Call(obj, GoStrToDStr(S) , AObject )
    return int32(ret)
}

func StringList_Assign(obj uintptr, Source uintptr)  {
    stringList_Assign.Call(obj, Source )
}

func StringList_Clear(obj uintptr)  {
    stringList_Clear.Call(obj)
}

func StringList_Delete(obj uintptr, Index int32)  {
    stringList_Delete.Call(obj, uintptr(Index) )
}

func StringList_IndexOf(obj uintptr, S string) int32 {
    ret, _, _ := stringList_IndexOf.Call(obj, GoStrToDStr(S) )
    return int32(ret)
}

func StringList_Insert(obj uintptr, Index int32, S string)  {
    stringList_Insert.Call(obj, uintptr(Index) , GoStrToDStr(S) )
}

func StringList_InsertObject(obj uintptr, Index int32, S string, AObject uintptr)  {
    stringList_InsertObject.Call(obj, uintptr(Index) , GoStrToDStr(S) , AObject )
}

func StringList_Append(obj uintptr, S string)  {
    stringList_Append.Call(obj, GoStrToDStr(S) )
}

func StringList_BeginUpdate(obj uintptr)  {
    stringList_BeginUpdate.Call(obj)
}

func StringList_EndUpdate(obj uintptr)  {
    stringList_EndUpdate.Call(obj)
}

func StringList_Equals(obj uintptr, Strings uintptr) bool {
    ret, _, _ := stringList_Equals.Call(obj, Strings )
    return DBoolToGoBool(ret)
}

func StringList_IndexOfName(obj uintptr, Name string) int32 {
    ret, _, _ := stringList_IndexOfName.Call(obj, GoStrToDStr(Name) )
    return int32(ret)
}

func StringList_IndexOfObject(obj uintptr, AObject uintptr) int32 {
    ret, _, _ := stringList_IndexOfObject.Call(obj, AObject )
    return int32(ret)
}

func StringList_LoadFromFile(obj uintptr, FileName string)  {
    stringList_LoadFromFile.Call(obj, GoStrToDStr(FileName) )
}

func StringList_LoadFromStream(obj uintptr, Stream uintptr)  {
    stringList_LoadFromStream.Call(obj, Stream )
}

func StringList_Move(obj uintptr, CurIndex int32, NewIndex int32)  {
    stringList_Move.Call(obj, uintptr(CurIndex) , uintptr(NewIndex) )
}

func StringList_SaveToFile(obj uintptr, FileName string)  {
    stringList_SaveToFile.Call(obj, GoStrToDStr(FileName) )
}

func StringList_SaveToStream(obj uintptr, Stream uintptr)  {
    stringList_SaveToStream.Call(obj, Stream )
}

func StringList_GetNamePath(obj uintptr) string {
    ret, _, _ := stringList_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func StringList_ClassName(obj uintptr) string {
    ret, _, _ := stringList_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func StringList_GetHashCode(obj uintptr) int32 {
    ret, _, _ := stringList_GetHashCode.Call(obj)
    return int32(ret)
}

func StringList_ToString(obj uintptr) string {
    ret, _, _ := stringList_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func StringList_GetSorted(obj uintptr) bool {
    ret, _, _ := stringList_GetSorted.Call(obj)
    return DBoolToGoBool(ret)
}

func StringList_SetSorted(obj uintptr, value bool) {
   stringList_SetSorted.Call(obj, GoBoolToDBool(value))
}

func StringList_SetOnChange(obj uintptr, fn interface{}) {
    stringList_SetOnChange.Call(obj, addEventToMap(fn))
}

func StringList_GetCommaText(obj uintptr) string {
    ret, _, _ := stringList_GetCommaText.Call(obj)
    return DStrToGoStr(ret)
}

func StringList_SetCommaText(obj uintptr, value string) {
   stringList_SetCommaText.Call(obj, GoStrToDStr(value))
}

func StringList_GetDelimiter(obj uintptr) uint16 {
    ret, _, _ := stringList_GetDelimiter.Call(obj)
    return uint16(ret)
}

func StringList_SetDelimiter(obj uintptr, value uint16) {
   stringList_SetDelimiter.Call(obj, uintptr(value))
}

func StringList_GetText(obj uintptr) string {
    ret, _, _ := stringList_GetText.Call(obj)
    return DStrToGoStr(ret)
}

func StringList_SetText(obj uintptr, value string) {
   stringList_SetText.Call(obj, GoStrToDStr(value))
}

func StringList_GetWriteBOM(obj uintptr) bool {
    ret, _, _ := stringList_GetWriteBOM.Call(obj)
    return DBoolToGoBool(ret)
}

func StringList_SetWriteBOM(obj uintptr, value bool) {
   stringList_SetWriteBOM.Call(obj, GoBoolToDBool(value))
}

func StringList_GetOptions(obj uintptr) TStringsOptions {
    ret, _, _ := stringList_GetOptions.Call(obj)
    return TStringsOptions(ret)
}

func StringList_SetOptions(obj uintptr, value TStringsOptions) {
   stringList_SetOptions.Call(obj, uintptr(value))
}

func StringList_GetValues(obj uintptr, Name string) string {
    ret, _, _ := stringList_GetValues.Call(obj, GoStrToDStr(Name))
    return DStrToGoStr(ret)
}

func StringList_SetValues(obj uintptr, Name string, value string) {
   stringList_SetValues.Call(obj, GoStrToDStr(Name), GoStrToDStr(value))
}

func StringList_GetValueFromIndex(obj uintptr, Index int32) string {
    ret, _, _ := stringList_GetValueFromIndex.Call(obj, uintptr(Index))
    return DStrToGoStr(ret)
}

func StringList_SetValueFromIndex(obj uintptr, Index int32, value string) {
   stringList_SetValueFromIndex.Call(obj, uintptr(Index), GoStrToDStr(value))
}

func StringList_GetStrings(obj uintptr, Index int32) string {
    ret, _, _ := stringList_GetStrings.Call(obj, uintptr(Index))
    return DStrToGoStr(ret)
}

func StringList_SetStrings(obj uintptr, Index int32, value string) {
   stringList_SetStrings.Call(obj, uintptr(Index), GoStrToDStr(value))
}


//--------------------------- TBrush ---------------------------

func Brush_Create() uintptr {
    ret, _, _ := brush_Create.Call()
    return ret
}

func Brush_Free(obj uintptr) {
    brush_Free.Call(obj)
}

func Brush_Assign(obj uintptr, Source uintptr)  {
    brush_Assign.Call(obj, Source )
}

func Brush_HandleAllocated(obj uintptr) bool {
    ret, _, _ := brush_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Brush_GetNamePath(obj uintptr) string {
    ret, _, _ := brush_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Brush_ClassName(obj uintptr) string {
    ret, _, _ := brush_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Brush_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := brush_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Brush_GetHashCode(obj uintptr) int32 {
    ret, _, _ := brush_GetHashCode.Call(obj)
    return int32(ret)
}

func Brush_ToString(obj uintptr) string {
    ret, _, _ := brush_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Brush_GetBitmap(obj uintptr) uintptr {
    ret, _, _ := brush_GetBitmap.Call(obj)
    return ret
}

func Brush_SetBitmap(obj uintptr, value uintptr) {
   brush_SetBitmap.Call(obj, value)
}

func Brush_GetHandle(obj uintptr) HBRUSH {
    ret, _, _ := brush_GetHandle.Call(obj)
    return HBRUSH(ret)
}

func Brush_SetHandle(obj uintptr, value HBRUSH) {
   brush_SetHandle.Call(obj, uintptr(value))
}

func Brush_GetColor(obj uintptr) TColor {
    ret, _, _ := brush_GetColor.Call(obj)
    return TColor(ret)
}

func Brush_SetColor(obj uintptr, value TColor) {
   brush_SetColor.Call(obj, uintptr(value))
}

func Brush_GetStyle(obj uintptr) TBrushStyle {
    ret, _, _ := brush_GetStyle.Call(obj)
    return TBrushStyle(ret)
}

func Brush_SetStyle(obj uintptr, value TBrushStyle) {
   brush_SetStyle.Call(obj, uintptr(value))
}

func Brush_SetOnChange(obj uintptr, fn interface{}) {
    brush_SetOnChange.Call(obj, addEventToMap(fn))
}


//--------------------------- TPen ---------------------------

func Pen_Create() uintptr {
    ret, _, _ := pen_Create.Call()
    return ret
}

func Pen_Free(obj uintptr) {
    pen_Free.Call(obj)
}

func Pen_Assign(obj uintptr, Source uintptr)  {
    pen_Assign.Call(obj, Source )
}

func Pen_HandleAllocated(obj uintptr) bool {
    ret, _, _ := pen_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Pen_GetNamePath(obj uintptr) string {
    ret, _, _ := pen_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Pen_ClassName(obj uintptr) string {
    ret, _, _ := pen_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Pen_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := pen_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Pen_GetHashCode(obj uintptr) int32 {
    ret, _, _ := pen_GetHashCode.Call(obj)
    return int32(ret)
}

func Pen_ToString(obj uintptr) string {
    ret, _, _ := pen_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Pen_GetHandle(obj uintptr) HPEN {
    ret, _, _ := pen_GetHandle.Call(obj)
    return HPEN(ret)
}

func Pen_SetHandle(obj uintptr, value HPEN) {
   pen_SetHandle.Call(obj, uintptr(value))
}

func Pen_GetColor(obj uintptr) TColor {
    ret, _, _ := pen_GetColor.Call(obj)
    return TColor(ret)
}

func Pen_SetColor(obj uintptr, value TColor) {
   pen_SetColor.Call(obj, uintptr(value))
}

func Pen_GetMode(obj uintptr) TPenMode {
    ret, _, _ := pen_GetMode.Call(obj)
    return TPenMode(ret)
}

func Pen_SetMode(obj uintptr, value TPenMode) {
   pen_SetMode.Call(obj, uintptr(value))
}

func Pen_GetStyle(obj uintptr) TPenStyle {
    ret, _, _ := pen_GetStyle.Call(obj)
    return TPenStyle(ret)
}

func Pen_SetStyle(obj uintptr, value TPenStyle) {
   pen_SetStyle.Call(obj, uintptr(value))
}

func Pen_GetWidth(obj uintptr) int32 {
    ret, _, _ := pen_GetWidth.Call(obj)
    return int32(ret)
}

func Pen_SetWidth(obj uintptr, value int32) {
   pen_SetWidth.Call(obj, uintptr(value))
}

func Pen_SetOnChange(obj uintptr, fn interface{}) {
    pen_SetOnChange.Call(obj, addEventToMap(fn))
}


//--------------------------- TMenuItem ---------------------------

func MenuItem_Create(obj uintptr) uintptr {
    ret, _, _ := menuItem_Create.Call(obj)
    return ret
}

func MenuItem_Free(obj uintptr) {
    menuItem_Free.Call(obj)
}

func MenuItem_Insert(obj uintptr, Index int32, Item uintptr)  {
    menuItem_Insert.Call(obj, uintptr(Index) , Item )
}

func MenuItem_Delete(obj uintptr, Index int32)  {
    menuItem_Delete.Call(obj, uintptr(Index) )
}

func MenuItem_Clear(obj uintptr)  {
    menuItem_Clear.Call(obj)
}

func MenuItem_Click(obj uintptr)  {
    menuItem_Click.Call(obj)
}

func MenuItem_IndexOf(obj uintptr, Item uintptr) int32 {
    ret, _, _ := menuItem_IndexOf.Call(obj, Item )
    return int32(ret)
}

func MenuItem_HasParent(obj uintptr) bool {
    ret, _, _ := menuItem_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func MenuItem_Add(obj uintptr, Item uintptr)  {
    menuItem_Add.Call(obj, Item )
}

func MenuItem_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := menuItem_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func MenuItem_GetNamePath(obj uintptr) string {
    ret, _, _ := menuItem_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func MenuItem_Assign(obj uintptr, Source uintptr)  {
    menuItem_Assign.Call(obj, Source )
}

func MenuItem_ClassName(obj uintptr) string {
    ret, _, _ := menuItem_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func MenuItem_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := menuItem_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func MenuItem_GetHashCode(obj uintptr) int32 {
    ret, _, _ := menuItem_GetHashCode.Call(obj)
    return int32(ret)
}

func MenuItem_ToString(obj uintptr) string {
    ret, _, _ := menuItem_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func MenuItem_GetHandle(obj uintptr) HMENU {
    ret, _, _ := menuItem_GetHandle.Call(obj)
    return HMENU(ret)
}

func MenuItem_GetParent(obj uintptr) uintptr {
    ret, _, _ := menuItem_GetParent.Call(obj)
    return ret
}

func MenuItem_GetAction(obj uintptr) uintptr {
    ret, _, _ := menuItem_GetAction.Call(obj)
    return ret
}

func MenuItem_SetAction(obj uintptr, value uintptr) {
   menuItem_SetAction.Call(obj, value)
}

func MenuItem_GetAutoHotkeys(obj uintptr) TMenuItemAutoFlag {
    ret, _, _ := menuItem_GetAutoHotkeys.Call(obj)
    return TMenuItemAutoFlag(ret)
}

func MenuItem_SetAutoHotkeys(obj uintptr, value TMenuItemAutoFlag) {
   menuItem_SetAutoHotkeys.Call(obj, uintptr(value))
}

func MenuItem_GetBitmap(obj uintptr) uintptr {
    ret, _, _ := menuItem_GetBitmap.Call(obj)
    return ret
}

func MenuItem_SetBitmap(obj uintptr, value uintptr) {
   menuItem_SetBitmap.Call(obj, value)
}

func MenuItem_GetCaption(obj uintptr) string {
    ret, _, _ := menuItem_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func MenuItem_SetCaption(obj uintptr, value string) {
   menuItem_SetCaption.Call(obj, GoStrToDStr(value))
}

func MenuItem_GetChecked(obj uintptr) bool {
    ret, _, _ := menuItem_GetChecked.Call(obj)
    return DBoolToGoBool(ret)
}

func MenuItem_SetChecked(obj uintptr, value bool) {
   menuItem_SetChecked.Call(obj, GoBoolToDBool(value))
}

func MenuItem_GetDefault(obj uintptr) bool {
    ret, _, _ := menuItem_GetDefault.Call(obj)
    return DBoolToGoBool(ret)
}

func MenuItem_SetDefault(obj uintptr, value bool) {
   menuItem_SetDefault.Call(obj, GoBoolToDBool(value))
}

func MenuItem_GetEnabled(obj uintptr) bool {
    ret, _, _ := menuItem_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func MenuItem_SetEnabled(obj uintptr, value bool) {
   menuItem_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func MenuItem_GetGroupIndex(obj uintptr) uint8 {
    ret, _, _ := menuItem_GetGroupIndex.Call(obj)
    return uint8(ret)
}

func MenuItem_SetGroupIndex(obj uintptr, value uint8) {
   menuItem_SetGroupIndex.Call(obj, uintptr(value))
}

func MenuItem_GetHint(obj uintptr) string {
    ret, _, _ := menuItem_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func MenuItem_SetHint(obj uintptr, value string) {
   menuItem_SetHint.Call(obj, GoStrToDStr(value))
}

func MenuItem_GetImageIndex(obj uintptr) int32 {
    ret, _, _ := menuItem_GetImageIndex.Call(obj)
    return int32(ret)
}

func MenuItem_SetImageIndex(obj uintptr, value int32) {
   menuItem_SetImageIndex.Call(obj, uintptr(value))
}

func MenuItem_GetShortCut(obj uintptr) TShortCut {
    ret, _, _ := menuItem_GetShortCut.Call(obj)
    return TShortCut(ret)
}

func MenuItem_SetShortCut(obj uintptr, value TShortCut) {
   menuItem_SetShortCut.Call(obj, uintptr(value))
}

func MenuItem_GetVisible(obj uintptr) bool {
    ret, _, _ := menuItem_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func MenuItem_SetVisible(obj uintptr, value bool) {
   menuItem_SetVisible.Call(obj, GoBoolToDBool(value))
}

func MenuItem_SetOnClick(obj uintptr, fn interface{}) {
    menuItem_SetOnClick.Call(obj, addEventToMap(fn))
}

func MenuItem_SetOnDrawItem(obj uintptr, fn interface{}) {
    menuItem_SetOnDrawItem.Call(obj, addEventToMap(fn))
}

func MenuItem_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := menuItem_GetComponentCount.Call(obj)
    return int32(ret)
}

func MenuItem_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := menuItem_GetComponentIndex.Call(obj)
    return int32(ret)
}

func MenuItem_SetComponentIndex(obj uintptr, value int32) {
   menuItem_SetComponentIndex.Call(obj, uintptr(value))
}

func MenuItem_GetOwner(obj uintptr) uintptr {
    ret, _, _ := menuItem_GetOwner.Call(obj)
    return ret
}

func MenuItem_GetName(obj uintptr) string {
    ret, _, _ := menuItem_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func MenuItem_SetName(obj uintptr, value string) {
   menuItem_SetName.Call(obj, GoStrToDStr(value))
}

func MenuItem_GetTag(obj uintptr) int {
    ret, _, _ := menuItem_GetTag.Call(obj)
    return int(ret)
}

func MenuItem_SetTag(obj uintptr, value int) {
   menuItem_SetTag.Call(obj, uintptr(value))
}

func MenuItem_GetItems(obj uintptr, Index int32) uintptr {
    ret, _, _ := menuItem_GetItems.Call(obj, uintptr(Index))
    return ret
}

func MenuItem_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := menuItem_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TListGroups ---------------------------

func ListGroups_Create() uintptr {
    ret, _, _ := listGroups_Create.Call()
    return ret
}

func ListGroups_Free(obj uintptr) {
    listGroups_Free.Call(obj)
}

func ListGroups_Add(obj uintptr) uintptr {
    ret, _, _ := listGroups_Add.Call(obj)
    return ret
}

func ListGroups_Owner(obj uintptr) uintptr {
    ret, _, _ := listGroups_Owner.Call(obj)
    return ret
}

func ListGroups_Assign(obj uintptr, Source uintptr)  {
    listGroups_Assign.Call(obj, Source )
}

func ListGroups_BeginUpdate(obj uintptr)  {
    listGroups_BeginUpdate.Call(obj)
}

func ListGroups_Clear(obj uintptr)  {
    listGroups_Clear.Call(obj)
}

func ListGroups_Delete(obj uintptr, Index int32)  {
    listGroups_Delete.Call(obj, uintptr(Index) )
}

func ListGroups_EndUpdate(obj uintptr)  {
    listGroups_EndUpdate.Call(obj)
}

func ListGroups_GetNamePath(obj uintptr) string {
    ret, _, _ := listGroups_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ListGroups_Insert(obj uintptr, Index int32) uintptr {
    ret, _, _ := listGroups_Insert.Call(obj, uintptr(Index) )
    return ret
}

func ListGroups_ClassName(obj uintptr) string {
    ret, _, _ := listGroups_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ListGroups_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := listGroups_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ListGroups_GetHashCode(obj uintptr) int32 {
    ret, _, _ := listGroups_GetHashCode.Call(obj)
    return int32(ret)
}

func ListGroups_ToString(obj uintptr) string {
    ret, _, _ := listGroups_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ListGroups_GetItems(obj uintptr, Index int32) uintptr {
    ret, _, _ := listGroups_GetItems.Call(obj, uintptr(Index))
    return ret
}

func ListGroups_SetItems(obj uintptr, Index int32, value uintptr) {
   listGroups_SetItems.Call(obj, uintptr(Index), value)
}


//--------------------------- TPicture ---------------------------

func Picture_Create() uintptr {
    ret, _, _ := picture_Create.Call()
    return ret
}

func Picture_Free(obj uintptr) {
    picture_Free.Call(obj)
}

func Picture_LoadFromFile(obj uintptr, Filename string)  {
    picture_LoadFromFile.Call(obj, GoStrToDStr(Filename) )
}

func Picture_SaveToFile(obj uintptr, Filename string)  {
    picture_SaveToFile.Call(obj, GoStrToDStr(Filename) )
}

func Picture_LoadFromStream(obj uintptr, Stream uintptr)  {
    picture_LoadFromStream.Call(obj, Stream )
}

func Picture_SaveToStream(obj uintptr, Stream uintptr)  {
    picture_SaveToStream.Call(obj, Stream )
}

func Picture_Assign(obj uintptr, Source uintptr)  {
    picture_Assign.Call(obj, Source )
}

func Picture_GetNamePath(obj uintptr) string {
    ret, _, _ := picture_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Picture_ClassName(obj uintptr) string {
    ret, _, _ := picture_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Picture_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := picture_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Picture_GetHashCode(obj uintptr) int32 {
    ret, _, _ := picture_GetHashCode.Call(obj)
    return int32(ret)
}

func Picture_ToString(obj uintptr) string {
    ret, _, _ := picture_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Picture_GetBitmap(obj uintptr) uintptr {
    ret, _, _ := picture_GetBitmap.Call(obj)
    return ret
}

func Picture_SetBitmap(obj uintptr, value uintptr) {
   picture_SetBitmap.Call(obj, value)
}

func Picture_GetGraphic(obj uintptr) uintptr {
    ret, _, _ := picture_GetGraphic.Call(obj)
    return ret
}

func Picture_SetGraphic(obj uintptr, value uintptr) {
   picture_SetGraphic.Call(obj, value)
}

func Picture_GetHeight(obj uintptr) int32 {
    ret, _, _ := picture_GetHeight.Call(obj)
    return int32(ret)
}

func Picture_GetIcon(obj uintptr) uintptr {
    ret, _, _ := picture_GetIcon.Call(obj)
    return ret
}

func Picture_SetIcon(obj uintptr, value uintptr) {
   picture_SetIcon.Call(obj, value)
}

func Picture_GetWidth(obj uintptr) int32 {
    ret, _, _ := picture_GetWidth.Call(obj)
    return int32(ret)
}

func Picture_SetOnChange(obj uintptr, fn interface{}) {
    picture_SetOnChange.Call(obj, addEventToMap(fn))
}


//--------------------------- TListColumns ---------------------------

func ListColumns_Create() uintptr {
    ret, _, _ := listColumns_Create.Call()
    return ret
}

func ListColumns_Free(obj uintptr) {
    listColumns_Free.Call(obj)
}

func ListColumns_Add(obj uintptr) uintptr {
    ret, _, _ := listColumns_Add.Call(obj)
    return ret
}

func ListColumns_Owner(obj uintptr) uintptr {
    ret, _, _ := listColumns_Owner.Call(obj)
    return ret
}

func ListColumns_Assign(obj uintptr, Source uintptr)  {
    listColumns_Assign.Call(obj, Source )
}

func ListColumns_BeginUpdate(obj uintptr)  {
    listColumns_BeginUpdate.Call(obj)
}

func ListColumns_Clear(obj uintptr)  {
    listColumns_Clear.Call(obj)
}

func ListColumns_Delete(obj uintptr, Index int32)  {
    listColumns_Delete.Call(obj, uintptr(Index) )
}

func ListColumns_EndUpdate(obj uintptr)  {
    listColumns_EndUpdate.Call(obj)
}

func ListColumns_GetNamePath(obj uintptr) string {
    ret, _, _ := listColumns_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ListColumns_Insert(obj uintptr, Index int32) uintptr {
    ret, _, _ := listColumns_Insert.Call(obj, uintptr(Index) )
    return ret
}

func ListColumns_ClassName(obj uintptr) string {
    ret, _, _ := listColumns_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ListColumns_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := listColumns_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ListColumns_GetHashCode(obj uintptr) int32 {
    ret, _, _ := listColumns_GetHashCode.Call(obj)
    return int32(ret)
}

func ListColumns_ToString(obj uintptr) string {
    ret, _, _ := listColumns_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ListColumns_GetItems(obj uintptr, Index int32) uintptr {
    ret, _, _ := listColumns_GetItems.Call(obj, uintptr(Index))
    return ret
}

func ListColumns_SetItems(obj uintptr, Index int32, value uintptr) {
   listColumns_SetItems.Call(obj, uintptr(Index), value)
}


//--------------------------- TListItems ---------------------------

func ListItems_Create() uintptr {
    ret, _, _ := listItems_Create.Call()
    return ret
}

func ListItems_Free(obj uintptr) {
    listItems_Free.Call(obj)
}

func ListItems_Add(obj uintptr) uintptr {
    ret, _, _ := listItems_Add.Call(obj)
    return ret
}

func ListItems_AddItem(obj uintptr, Item uintptr, Index int32) uintptr {
    ret, _, _ := listItems_AddItem.Call(obj, Item , uintptr(Index) )
    return ret
}

func ListItems_Assign(obj uintptr, Source uintptr)  {
    listItems_Assign.Call(obj, Source )
}

func ListItems_BeginUpdate(obj uintptr)  {
    listItems_BeginUpdate.Call(obj)
}

func ListItems_Clear(obj uintptr)  {
    listItems_Clear.Call(obj)
}

func ListItems_Delete(obj uintptr, Index int32)  {
    listItems_Delete.Call(obj, uintptr(Index) )
}

func ListItems_EndUpdate(obj uintptr)  {
    listItems_EndUpdate.Call(obj)
}

func ListItems_IndexOf(obj uintptr, Value uintptr) int32 {
    ret, _, _ := listItems_IndexOf.Call(obj, Value )
    return int32(ret)
}

func ListItems_Insert(obj uintptr, Index int32) uintptr {
    ret, _, _ := listItems_Insert.Call(obj, uintptr(Index) )
    return ret
}

func ListItems_GetNamePath(obj uintptr) string {
    ret, _, _ := listItems_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ListItems_ClassName(obj uintptr) string {
    ret, _, _ := listItems_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ListItems_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := listItems_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ListItems_GetHashCode(obj uintptr) int32 {
    ret, _, _ := listItems_GetHashCode.Call(obj)
    return int32(ret)
}

func ListItems_ToString(obj uintptr) string {
    ret, _, _ := listItems_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ListItems_GetHandle(obj uintptr) HWND {
    ret, _, _ := listItems_GetHandle.Call(obj)
    return HWND(ret)
}

func ListItems_GetOwner(obj uintptr) uintptr {
    ret, _, _ := listItems_GetOwner.Call(obj)
    return ret
}

func ListItems_GetItem(obj uintptr, Index int32) uintptr {
    ret, _, _ := listItems_GetItem.Call(obj, uintptr(Index))
    return ret
}

func ListItems_SetItem(obj uintptr, Index int32, value uintptr) {
   listItems_SetItem.Call(obj, uintptr(Index), value)
}


//--------------------------- TTreeNodes ---------------------------

func TreeNodes_Create() uintptr {
    ret, _, _ := treeNodes_Create.Call()
    return ret
}

func TreeNodes_Free(obj uintptr) {
    treeNodes_Free.Call(obj)
}

func TreeNodes_AddChildFirst(obj uintptr, Parent uintptr, S string) uintptr {
    ret, _, _ := treeNodes_AddChildFirst.Call(obj, Parent , GoStrToDStr(S) )
    return ret
}

func TreeNodes_AddChild(obj uintptr, Parent uintptr, S string) uintptr {
    ret, _, _ := treeNodes_AddChild.Call(obj, Parent , GoStrToDStr(S) )
    return ret
}

func TreeNodes_AddChildObjectFirst(obj uintptr, Parent uintptr, S string, Ptr uintptr) uintptr {
    ret, _, _ := treeNodes_AddChildObjectFirst.Call(obj, Parent , GoStrToDStr(S) , Ptr )
    return ret
}

func TreeNodes_AddChildObject(obj uintptr, Parent uintptr, S string, Ptr uintptr) uintptr {
    ret, _, _ := treeNodes_AddChildObject.Call(obj, Parent , GoStrToDStr(S) , Ptr )
    return ret
}

func TreeNodes_AddObjectFirst(obj uintptr, Sibling uintptr, S string, Ptr uintptr) uintptr {
    ret, _, _ := treeNodes_AddObjectFirst.Call(obj, Sibling , GoStrToDStr(S) , Ptr )
    return ret
}

func TreeNodes_AddObject(obj uintptr, Sibling uintptr, S string, Ptr uintptr) uintptr {
    ret, _, _ := treeNodes_AddObject.Call(obj, Sibling , GoStrToDStr(S) , Ptr )
    return ret
}

func TreeNodes_AddNode(obj uintptr, Node uintptr, Relative uintptr, S string, Ptr uintptr, Method TNodeAttachMode) uintptr {
    ret, _, _ := treeNodes_AddNode.Call(obj, Node , Relative , GoStrToDStr(S) , Ptr , uintptr(Method) )
    return ret
}

func TreeNodes_AddFirst(obj uintptr, Sibling uintptr, S string) uintptr {
    ret, _, _ := treeNodes_AddFirst.Call(obj, Sibling , GoStrToDStr(S) )
    return ret
}

func TreeNodes_Add(obj uintptr, Sibling uintptr, S string) uintptr {
    ret, _, _ := treeNodes_Add.Call(obj, Sibling , GoStrToDStr(S) )
    return ret
}

func TreeNodes_AlphaSort(obj uintptr, ARecurse bool) bool {
    ret, _, _ := treeNodes_AlphaSort.Call(obj, GoBoolToDBool(ARecurse) )
    return DBoolToGoBool(ret)
}

func TreeNodes_Assign(obj uintptr, Source uintptr)  {
    treeNodes_Assign.Call(obj, Source )
}

func TreeNodes_BeginUpdate(obj uintptr)  {
    treeNodes_BeginUpdate.Call(obj)
}

func TreeNodes_Clear(obj uintptr)  {
    treeNodes_Clear.Call(obj)
}

func TreeNodes_Delete(obj uintptr, Node uintptr)  {
    treeNodes_Delete.Call(obj, Node )
}

func TreeNodes_EndUpdate(obj uintptr)  {
    treeNodes_EndUpdate.Call(obj)
}

func TreeNodes_GetFirstNode(obj uintptr) uintptr {
    ret, _, _ := treeNodes_GetFirstNode.Call(obj)
    return ret
}

func TreeNodes_GetNode(obj uintptr, ItemId uintptr) uintptr {
    ret, _, _ := treeNodes_GetNode.Call(obj, ItemId )
    return ret
}

func TreeNodes_Insert(obj uintptr, Sibling uintptr, S string) uintptr {
    ret, _, _ := treeNodes_Insert.Call(obj, Sibling , GoStrToDStr(S) )
    return ret
}

func TreeNodes_InsertObject(obj uintptr, Sibling uintptr, S string, Ptr uintptr) uintptr {
    ret, _, _ := treeNodes_InsertObject.Call(obj, Sibling , GoStrToDStr(S) , Ptr )
    return ret
}

func TreeNodes_GetNamePath(obj uintptr) string {
    ret, _, _ := treeNodes_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func TreeNodes_ClassName(obj uintptr) string {
    ret, _, _ := treeNodes_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func TreeNodes_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := treeNodes_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func TreeNodes_GetHashCode(obj uintptr) int32 {
    ret, _, _ := treeNodes_GetHashCode.Call(obj)
    return int32(ret)
}

func TreeNodes_ToString(obj uintptr) string {
    ret, _, _ := treeNodes_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func TreeNodes_GetHandle(obj uintptr) HWND {
    ret, _, _ := treeNodes_GetHandle.Call(obj)
    return HWND(ret)
}

func TreeNodes_GetOwner(obj uintptr) uintptr {
    ret, _, _ := treeNodes_GetOwner.Call(obj)
    return ret
}

func TreeNodes_GetItem(obj uintptr, Index int32) uintptr {
    ret, _, _ := treeNodes_GetItem.Call(obj, uintptr(Index))
    return ret
}


//--------------------------- TListItem ---------------------------

func ListItem_Create() uintptr {
    ret, _, _ := listItem_Create.Call()
    return ret
}

func ListItem_Free(obj uintptr) {
    listItem_Free.Call(obj)
}

func ListItem_Assign(obj uintptr, Source uintptr)  {
    listItem_Assign.Call(obj, Source )
}

func ListItem_CancelEdit(obj uintptr)  {
    listItem_CancelEdit.Call(obj)
}

func ListItem_Delete(obj uintptr)  {
    listItem_Delete.Call(obj)
}

func ListItem_EditCaption(obj uintptr) bool {
    ret, _, _ := listItem_EditCaption.Call(obj)
    return DBoolToGoBool(ret)
}

func ListItem_MakeVisible(obj uintptr, PartialOK bool)  {
    listItem_MakeVisible.Call(obj, GoBoolToDBool(PartialOK) )
}

func ListItem_Update(obj uintptr)  {
    listItem_Update.Call(obj)
}

func ListItem_WorkArea(obj uintptr) int32 {
    ret, _, _ := listItem_WorkArea.Call(obj)
    return int32(ret)
}

func ListItem_GetNamePath(obj uintptr) string {
    ret, _, _ := listItem_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ListItem_ClassName(obj uintptr) string {
    ret, _, _ := listItem_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ListItem_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := listItem_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ListItem_GetHashCode(obj uintptr) int32 {
    ret, _, _ := listItem_GetHashCode.Call(obj)
    return int32(ret)
}

func ListItem_ToString(obj uintptr) string {
    ret, _, _ := listItem_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ListItem_GetCaption(obj uintptr) string {
    ret, _, _ := listItem_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func ListItem_SetCaption(obj uintptr, value string) {
   listItem_SetCaption.Call(obj, GoStrToDStr(value))
}

func ListItem_GetChecked(obj uintptr) bool {
    ret, _, _ := listItem_GetChecked.Call(obj)
    return DBoolToGoBool(ret)
}

func ListItem_SetChecked(obj uintptr, value bool) {
   listItem_SetChecked.Call(obj, GoBoolToDBool(value))
}

func ListItem_GetCut(obj uintptr) bool {
    ret, _, _ := listItem_GetCut.Call(obj)
    return DBoolToGoBool(ret)
}

func ListItem_SetCut(obj uintptr, value bool) {
   listItem_SetCut.Call(obj, GoBoolToDBool(value))
}

func ListItem_GetData(obj uintptr) uintptr {
    ret, _, _ := listItem_GetData.Call(obj)
    return ret
}

func ListItem_SetData(obj uintptr, value uintptr) {
   listItem_SetData.Call(obj, value)
}

func ListItem_GetDeleting(obj uintptr) bool {
    ret, _, _ := listItem_GetDeleting.Call(obj)
    return DBoolToGoBool(ret)
}

func ListItem_GetDropTarget(obj uintptr) bool {
    ret, _, _ := listItem_GetDropTarget.Call(obj)
    return DBoolToGoBool(ret)
}

func ListItem_SetDropTarget(obj uintptr, value bool) {
   listItem_SetDropTarget.Call(obj, GoBoolToDBool(value))
}

func ListItem_GetFocused(obj uintptr) bool {
    ret, _, _ := listItem_GetFocused.Call(obj)
    return DBoolToGoBool(ret)
}

func ListItem_SetFocused(obj uintptr, value bool) {
   listItem_SetFocused.Call(obj, GoBoolToDBool(value))
}

func ListItem_GetGroupID(obj uintptr) int32 {
    ret, _, _ := listItem_GetGroupID.Call(obj)
    return int32(ret)
}

func ListItem_SetGroupID(obj uintptr, value int32) {
   listItem_SetGroupID.Call(obj, uintptr(value))
}

func ListItem_GetHandle(obj uintptr) HWND {
    ret, _, _ := listItem_GetHandle.Call(obj)
    return HWND(ret)
}

func ListItem_GetImageIndex(obj uintptr) int32 {
    ret, _, _ := listItem_GetImageIndex.Call(obj)
    return int32(ret)
}

func ListItem_SetImageIndex(obj uintptr, value int32) {
   listItem_SetImageIndex.Call(obj, uintptr(value))
}

func ListItem_GetIndent(obj uintptr) int32 {
    ret, _, _ := listItem_GetIndent.Call(obj)
    return int32(ret)
}

func ListItem_SetIndent(obj uintptr, value int32) {
   listItem_SetIndent.Call(obj, uintptr(value))
}

func ListItem_GetIndex(obj uintptr) int32 {
    ret, _, _ := listItem_GetIndex.Call(obj)
    return int32(ret)
}

func ListItem_GetLeft(obj uintptr) int32 {
    ret, _, _ := listItem_GetLeft.Call(obj)
    return int32(ret)
}

func ListItem_SetLeft(obj uintptr, value int32) {
   listItem_SetLeft.Call(obj, uintptr(value))
}

func ListItem_GetOwner(obj uintptr) uintptr {
    ret, _, _ := listItem_GetOwner.Call(obj)
    return ret
}

func ListItem_GetOverlayIndex(obj uintptr) int32 {
    ret, _, _ := listItem_GetOverlayIndex.Call(obj)
    return int32(ret)
}

func ListItem_SetOverlayIndex(obj uintptr, value int32) {
   listItem_SetOverlayIndex.Call(obj, uintptr(value))
}

func ListItem_GetPosition(obj uintptr) TPoint {
    var ret TPoint
    listItem_GetPosition.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ListItem_SetPosition(obj uintptr, value TPoint) {
   listItem_SetPosition.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ListItem_GetSelected(obj uintptr) bool {
    ret, _, _ := listItem_GetSelected.Call(obj)
    return DBoolToGoBool(ret)
}

func ListItem_SetSelected(obj uintptr, value bool) {
   listItem_SetSelected.Call(obj, GoBoolToDBool(value))
}

func ListItem_GetStateIndex(obj uintptr) int32 {
    ret, _, _ := listItem_GetStateIndex.Call(obj)
    return int32(ret)
}

func ListItem_SetStateIndex(obj uintptr, value int32) {
   listItem_SetStateIndex.Call(obj, uintptr(value))
}

func ListItem_GetSubItems(obj uintptr) uintptr {
    ret, _, _ := listItem_GetSubItems.Call(obj)
    return ret
}

func ListItem_SetSubItems(obj uintptr, value uintptr) {
   listItem_SetSubItems.Call(obj, value)
}

func ListItem_GetTop(obj uintptr) int32 {
    ret, _, _ := listItem_GetTop.Call(obj)
    return int32(ret)
}

func ListItem_SetTop(obj uintptr, value int32) {
   listItem_SetTop.Call(obj, uintptr(value))
}

func ListItem_GetSubItemImages(obj uintptr, Index int32) int32 {
    ret, _, _ := listItem_GetSubItemImages.Call(obj, uintptr(Index))
    return int32(ret)
}

func ListItem_SetSubItemImages(obj uintptr, Index int32, value int32) {
   listItem_SetSubItemImages.Call(obj, uintptr(Index), uintptr(value))
}


//--------------------------- TTreeNode ---------------------------

func TreeNode_Create() uintptr {
    ret, _, _ := treeNode_Create.Call()
    return ret
}

func TreeNode_Free(obj uintptr) {
    treeNode_Free.Call(obj)
}

func TreeNode_AlphaSort(obj uintptr, ARecurse bool) bool {
    ret, _, _ := treeNode_AlphaSort.Call(obj, GoBoolToDBool(ARecurse) )
    return DBoolToGoBool(ret)
}

func TreeNode_Assign(obj uintptr, Source uintptr)  {
    treeNode_Assign.Call(obj, Source )
}

func TreeNode_Collapse(obj uintptr, Recurse bool)  {
    treeNode_Collapse.Call(obj, GoBoolToDBool(Recurse) )
}

func TreeNode_Delete(obj uintptr)  {
    treeNode_Delete.Call(obj)
}

func TreeNode_Expand(obj uintptr, Recurse bool)  {
    treeNode_Expand.Call(obj, GoBoolToDBool(Recurse) )
}

func TreeNode_IndexOf(obj uintptr, Value uintptr) int32 {
    ret, _, _ := treeNode_IndexOf.Call(obj, Value )
    return int32(ret)
}

func TreeNode_MakeVisible(obj uintptr)  {
    treeNode_MakeVisible.Call(obj)
}

func TreeNode_MoveTo(obj uintptr, Destination uintptr, Mode TNodeAttachMode)  {
    treeNode_MoveTo.Call(obj, Destination , uintptr(Mode) )
}

func TreeNode_GetNamePath(obj uintptr) string {
    ret, _, _ := treeNode_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func TreeNode_ClassName(obj uintptr) string {
    ret, _, _ := treeNode_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func TreeNode_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := treeNode_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func TreeNode_GetHashCode(obj uintptr) int32 {
    ret, _, _ := treeNode_GetHashCode.Call(obj)
    return int32(ret)
}

func TreeNode_ToString(obj uintptr) string {
    ret, _, _ := treeNode_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func TreeNode_GetAbsoluteIndex(obj uintptr) int32 {
    ret, _, _ := treeNode_GetAbsoluteIndex.Call(obj)
    return int32(ret)
}

func TreeNode_GetCut(obj uintptr) bool {
    ret, _, _ := treeNode_GetCut.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeNode_SetCut(obj uintptr, value bool) {
   treeNode_SetCut.Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetData(obj uintptr) uintptr {
    ret, _, _ := treeNode_GetData.Call(obj)
    return ret
}

func TreeNode_SetData(obj uintptr, value uintptr) {
   treeNode_SetData.Call(obj, value)
}

func TreeNode_GetDeleting(obj uintptr) bool {
    ret, _, _ := treeNode_GetDeleting.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeNode_GetFocused(obj uintptr) bool {
    ret, _, _ := treeNode_GetFocused.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeNode_SetFocused(obj uintptr, value bool) {
   treeNode_SetFocused.Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetDropTarget(obj uintptr) bool {
    ret, _, _ := treeNode_GetDropTarget.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeNode_SetDropTarget(obj uintptr, value bool) {
   treeNode_SetDropTarget.Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetSelected(obj uintptr) bool {
    ret, _, _ := treeNode_GetSelected.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeNode_SetSelected(obj uintptr, value bool) {
   treeNode_SetSelected.Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetExpanded(obj uintptr) bool {
    ret, _, _ := treeNode_GetExpanded.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeNode_SetExpanded(obj uintptr, value bool) {
   treeNode_SetExpanded.Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetExpandedImageIndex(obj uintptr) int32 {
    ret, _, _ := treeNode_GetExpandedImageIndex.Call(obj)
    return int32(ret)
}

func TreeNode_SetExpandedImageIndex(obj uintptr, value int32) {
   treeNode_SetExpandedImageIndex.Call(obj, uintptr(value))
}

func TreeNode_GetHandle(obj uintptr) HWND {
    ret, _, _ := treeNode_GetHandle.Call(obj)
    return HWND(ret)
}

func TreeNode_GetHasChildren(obj uintptr) bool {
    ret, _, _ := treeNode_GetHasChildren.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeNode_SetHasChildren(obj uintptr, value bool) {
   treeNode_SetHasChildren.Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetImageIndex(obj uintptr) int32 {
    ret, _, _ := treeNode_GetImageIndex.Call(obj)
    return int32(ret)
}

func TreeNode_SetImageIndex(obj uintptr, value int32) {
   treeNode_SetImageIndex.Call(obj, uintptr(value))
}

func TreeNode_GetIndex(obj uintptr) int32 {
    ret, _, _ := treeNode_GetIndex.Call(obj)
    return int32(ret)
}

func TreeNode_GetIsVisible(obj uintptr) bool {
    ret, _, _ := treeNode_GetIsVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeNode_GetItemId(obj uintptr) uintptr {
    ret, _, _ := treeNode_GetItemId.Call(obj)
    return ret
}

func TreeNode_GetLevel(obj uintptr) int32 {
    ret, _, _ := treeNode_GetLevel.Call(obj)
    return int32(ret)
}

func TreeNode_GetOverlayIndex(obj uintptr) int32 {
    ret, _, _ := treeNode_GetOverlayIndex.Call(obj)
    return int32(ret)
}

func TreeNode_SetOverlayIndex(obj uintptr, value int32) {
   treeNode_SetOverlayIndex.Call(obj, uintptr(value))
}

func TreeNode_GetOwner(obj uintptr) uintptr {
    ret, _, _ := treeNode_GetOwner.Call(obj)
    return ret
}

func TreeNode_GetParent(obj uintptr) uintptr {
    ret, _, _ := treeNode_GetParent.Call(obj)
    return ret
}

func TreeNode_GetSelectedIndex(obj uintptr) int32 {
    ret, _, _ := treeNode_GetSelectedIndex.Call(obj)
    return int32(ret)
}

func TreeNode_SetSelectedIndex(obj uintptr, value int32) {
   treeNode_SetSelectedIndex.Call(obj, uintptr(value))
}

func TreeNode_GetEnabled(obj uintptr) bool {
    ret, _, _ := treeNode_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func TreeNode_SetEnabled(obj uintptr, value bool) {
   treeNode_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func TreeNode_GetStateIndex(obj uintptr) int32 {
    ret, _, _ := treeNode_GetStateIndex.Call(obj)
    return int32(ret)
}

func TreeNode_SetStateIndex(obj uintptr, value int32) {
   treeNode_SetStateIndex.Call(obj, uintptr(value))
}

func TreeNode_GetText(obj uintptr) string {
    ret, _, _ := treeNode_GetText.Call(obj)
    return DStrToGoStr(ret)
}

func TreeNode_SetText(obj uintptr, value string) {
   treeNode_SetText.Call(obj, GoStrToDStr(value))
}

func TreeNode_GetItem(obj uintptr, Index int32) uintptr {
    ret, _, _ := treeNode_GetItem.Call(obj, uintptr(Index))
    return ret
}

func TreeNode_SetItem(obj uintptr, Index int32, value uintptr) {
   treeNode_SetItem.Call(obj, uintptr(Index), value)
}


//--------------------------- TPageControl ---------------------------

func PageControl_Create(obj uintptr) uintptr {
    ret, _, _ := pageControl_Create.Call(obj)
    return ret
}

func PageControl_Free(obj uintptr) {
    pageControl_Free.Call(obj)
}

func PageControl_SelectNextPage(obj uintptr, GoForward bool, CheckTabVisible bool)  {
    pageControl_SelectNextPage.Call(obj, GoBoolToDBool(GoForward) , GoBoolToDBool(CheckTabVisible) )
}

func PageControl_RowCount(obj uintptr) int32 {
    ret, _, _ := pageControl_RowCount.Call(obj)
    return int32(ret)
}

func PageControl_CanFocus(obj uintptr) bool {
    ret, _, _ := pageControl_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_FlipChildren(obj uintptr, AllLevels bool)  {
    pageControl_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func PageControl_Focused(obj uintptr) bool {
    ret, _, _ := pageControl_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_HandleAllocated(obj uintptr) bool {
    ret, _, _ := pageControl_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_Invalidate(obj uintptr)  {
    pageControl_Invalidate.Call(obj)
}

func PageControl_Realign(obj uintptr)  {
    pageControl_Realign.Call(obj)
}

func PageControl_Repaint(obj uintptr)  {
    pageControl_Repaint.Call(obj)
}

func PageControl_ScaleBy(obj uintptr, M int32, D int32)  {
    pageControl_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func PageControl_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    pageControl_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func PageControl_SetFocus(obj uintptr)  {
    pageControl_SetFocus.Call(obj)
}

func PageControl_Update(obj uintptr)  {
    pageControl_Update.Call(obj)
}

func PageControl_BringToFront(obj uintptr)  {
    pageControl_BringToFront.Call(obj)
}

func PageControl_HasParent(obj uintptr) bool {
    ret, _, _ := pageControl_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_Hide(obj uintptr)  {
    pageControl_Hide.Call(obj)
}

func PageControl_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := pageControl_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func PageControl_Refresh(obj uintptr)  {
    pageControl_Refresh.Call(obj)
}

func PageControl_SendToBack(obj uintptr)  {
    pageControl_SendToBack.Call(obj)
}

func PageControl_Show(obj uintptr)  {
    pageControl_Show.Call(obj)
}

func PageControl_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := pageControl_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func PageControl_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := pageControl_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func PageControl_GetNamePath(obj uintptr) string {
    ret, _, _ := pageControl_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func PageControl_Assign(obj uintptr, Source uintptr)  {
    pageControl_Assign.Call(obj, Source )
}

func PageControl_ClassName(obj uintptr) string {
    ret, _, _ := pageControl_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func PageControl_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := pageControl_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func PageControl_GetHashCode(obj uintptr) int32 {
    ret, _, _ := pageControl_GetHashCode.Call(obj)
    return int32(ret)
}

func PageControl_ToString(obj uintptr) string {
    ret, _, _ := pageControl_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func PageControl_GetActivePageIndex(obj uintptr) int32 {
    ret, _, _ := pageControl_GetActivePageIndex.Call(obj)
    return int32(ret)
}

func PageControl_SetActivePageIndex(obj uintptr, value int32) {
   pageControl_SetActivePageIndex.Call(obj, uintptr(value))
}

func PageControl_GetPageCount(obj uintptr) int32 {
    ret, _, _ := pageControl_GetPageCount.Call(obj)
    return int32(ret)
}

func PageControl_GetAlign(obj uintptr) TAlign {
    ret, _, _ := pageControl_GetAlign.Call(obj)
    return TAlign(ret)
}

func PageControl_SetAlign(obj uintptr, value TAlign) {
   pageControl_SetAlign.Call(obj, uintptr(value))
}

func PageControl_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := pageControl_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func PageControl_SetAnchors(obj uintptr, value TAnchors) {
   pageControl_SetAnchors.Call(obj, uintptr(value))
}

func PageControl_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := pageControl_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func PageControl_SetBiDiMode(obj uintptr, value TBiDiMode) {
   pageControl_SetBiDiMode.Call(obj, uintptr(value))
}

func PageControl_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := pageControl_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetDoubleBuffered(obj uintptr, value bool) {
   pageControl_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetEnabled(obj uintptr) bool {
    ret, _, _ := pageControl_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetEnabled(obj uintptr, value bool) {
   pageControl_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetFont(obj uintptr) uintptr {
    ret, _, _ := pageControl_GetFont.Call(obj)
    return ret
}

func PageControl_SetFont(obj uintptr, value uintptr) {
   pageControl_SetFont.Call(obj, value)
}

func PageControl_GetHotTrack(obj uintptr) bool {
    ret, _, _ := pageControl_GetHotTrack.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetHotTrack(obj uintptr, value bool) {
   pageControl_SetHotTrack.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetImages(obj uintptr) uintptr {
    ret, _, _ := pageControl_GetImages.Call(obj)
    return ret
}

func PageControl_SetImages(obj uintptr, value uintptr) {
   pageControl_SetImages.Call(obj, value)
}

func PageControl_GetMultiLine(obj uintptr) bool {
    ret, _, _ := pageControl_GetMultiLine.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetMultiLine(obj uintptr, value bool) {
   pageControl_SetMultiLine.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := pageControl_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetParentDoubleBuffered(obj uintptr, value bool) {
   pageControl_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetParentFont(obj uintptr) bool {
    ret, _, _ := pageControl_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetParentFont(obj uintptr, value bool) {
   pageControl_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := pageControl_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetParentShowHint(obj uintptr, value bool) {
   pageControl_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := pageControl_GetPopupMenu.Call(obj)
    return ret
}

func PageControl_SetPopupMenu(obj uintptr, value uintptr) {
   pageControl_SetPopupMenu.Call(obj, value)
}

func PageControl_GetShowHint(obj uintptr) bool {
    ret, _, _ := pageControl_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetShowHint(obj uintptr, value bool) {
   pageControl_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetStyle(obj uintptr) TTabStyle {
    ret, _, _ := pageControl_GetStyle.Call(obj)
    return TTabStyle(ret)
}

func PageControl_SetStyle(obj uintptr, value TTabStyle) {
   pageControl_SetStyle.Call(obj, uintptr(value))
}

func PageControl_GetTabHeight(obj uintptr) int16 {
    ret, _, _ := pageControl_GetTabHeight.Call(obj)
    return int16(ret)
}

func PageControl_SetTabHeight(obj uintptr, value int16) {
   pageControl_SetTabHeight.Call(obj, uintptr(value))
}

func PageControl_GetTabIndex(obj uintptr) int32 {
    ret, _, _ := pageControl_GetTabIndex.Call(obj)
    return int32(ret)
}

func PageControl_SetTabIndex(obj uintptr, value int32) {
   pageControl_SetTabIndex.Call(obj, uintptr(value))
}

func PageControl_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := pageControl_GetTabOrder.Call(obj)
    return uint16(ret)
}

func PageControl_SetTabOrder(obj uintptr, value uint16) {
   pageControl_SetTabOrder.Call(obj, uintptr(value))
}

func PageControl_GetTabPosition(obj uintptr) TTabPosition {
    ret, _, _ := pageControl_GetTabPosition.Call(obj)
    return TTabPosition(ret)
}

func PageControl_SetTabPosition(obj uintptr, value TTabPosition) {
   pageControl_SetTabPosition.Call(obj, uintptr(value))
}

func PageControl_GetTabStop(obj uintptr) bool {
    ret, _, _ := pageControl_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetTabStop(obj uintptr, value bool) {
   pageControl_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetTabWidth(obj uintptr) int16 {
    ret, _, _ := pageControl_GetTabWidth.Call(obj)
    return int16(ret)
}

func PageControl_SetTabWidth(obj uintptr, value int16) {
   pageControl_SetTabWidth.Call(obj, uintptr(value))
}

func PageControl_GetVisible(obj uintptr) bool {
    ret, _, _ := pageControl_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetVisible(obj uintptr, value bool) {
   pageControl_SetVisible.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := pageControl_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func PageControl_SetStyleElements(obj uintptr, value TStyleElements) {
   pageControl_SetStyleElements.Call(obj, uintptr(value))
}

func PageControl_SetOnChange(obj uintptr, fn interface{}) {
    pageControl_SetOnChange.Call(obj, addEventToMap(fn))
}

func PageControl_SetOnEnter(obj uintptr, fn interface{}) {
    pageControl_SetOnEnter.Call(obj, addEventToMap(fn))
}

func PageControl_SetOnExit(obj uintptr, fn interface{}) {
    pageControl_SetOnExit.Call(obj, addEventToMap(fn))
}

func PageControl_SetOnGetImageIndex(obj uintptr, fn interface{}) {
    pageControl_SetOnGetImageIndex.Call(obj, addEventToMap(fn))
}

func PageControl_SetOnMouseDown(obj uintptr, fn interface{}) {
    pageControl_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func PageControl_SetOnMouseEnter(obj uintptr, fn interface{}) {
    pageControl_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func PageControl_SetOnMouseLeave(obj uintptr, fn interface{}) {
    pageControl_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func PageControl_SetOnMouseMove(obj uintptr, fn interface{}) {
    pageControl_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func PageControl_SetOnMouseUp(obj uintptr, fn interface{}) {
    pageControl_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func PageControl_SetOnResize(obj uintptr, fn interface{}) {
    pageControl_SetOnResize.Call(obj, addEventToMap(fn))
}

func PageControl_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := pageControl_GetCanvas.Call(obj)
    return ret
}

func PageControl_GetBrush(obj uintptr) uintptr {
    ret, _, _ := pageControl_GetBrush.Call(obj)
    return ret
}

func PageControl_GetControlCount(obj uintptr) int32 {
    ret, _, _ := pageControl_GetControlCount.Call(obj)
    return int32(ret)
}

func PageControl_GetHandle(obj uintptr) HWND {
    ret, _, _ := pageControl_GetHandle.Call(obj)
    return HWND(ret)
}

func PageControl_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := pageControl_GetParentWindow.Call(obj)
    return HWND(ret)
}

func PageControl_SetParentWindow(obj uintptr, value HWND) {
   pageControl_SetParentWindow.Call(obj, uintptr(value))
}

func PageControl_GetAction(obj uintptr) uintptr {
    ret, _, _ := pageControl_GetAction.Call(obj)
    return ret
}

func PageControl_SetAction(obj uintptr, value uintptr) {
   pageControl_SetAction.Call(obj, value)
}

func PageControl_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    pageControl_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func PageControl_SetBoundsRect(obj uintptr, value TRect) {
   pageControl_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func PageControl_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := pageControl_GetClientHeight.Call(obj)
    return int32(ret)
}

func PageControl_SetClientHeight(obj uintptr, value int32) {
   pageControl_SetClientHeight.Call(obj, uintptr(value))
}

func PageControl_GetClientRect(obj uintptr) TRect {
    var ret TRect
    pageControl_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func PageControl_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := pageControl_GetClientWidth.Call(obj)
    return int32(ret)
}

func PageControl_SetClientWidth(obj uintptr, value int32) {
   pageControl_SetClientWidth.Call(obj, uintptr(value))
}

func PageControl_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := pageControl_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func PageControl_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := pageControl_GetExplicitTop.Call(obj)
    return int32(ret)
}

func PageControl_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := pageControl_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func PageControl_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := pageControl_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func PageControl_GetParent(obj uintptr) uintptr {
    ret, _, _ := pageControl_GetParent.Call(obj)
    return ret
}

func PageControl_SetParent(obj uintptr, value uintptr) {
   pageControl_SetParent.Call(obj, value)
}

func PageControl_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := pageControl_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func PageControl_SetAlignWithMargins(obj uintptr, value bool) {
   pageControl_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func PageControl_GetLeft(obj uintptr) int32 {
    ret, _, _ := pageControl_GetLeft.Call(obj)
    return int32(ret)
}

func PageControl_SetLeft(obj uintptr, value int32) {
   pageControl_SetLeft.Call(obj, uintptr(value))
}

func PageControl_GetTop(obj uintptr) int32 {
    ret, _, _ := pageControl_GetTop.Call(obj)
    return int32(ret)
}

func PageControl_SetTop(obj uintptr, value int32) {
   pageControl_SetTop.Call(obj, uintptr(value))
}

func PageControl_GetWidth(obj uintptr) int32 {
    ret, _, _ := pageControl_GetWidth.Call(obj)
    return int32(ret)
}

func PageControl_SetWidth(obj uintptr, value int32) {
   pageControl_SetWidth.Call(obj, uintptr(value))
}

func PageControl_GetHeight(obj uintptr) int32 {
    ret, _, _ := pageControl_GetHeight.Call(obj)
    return int32(ret)
}

func PageControl_SetHeight(obj uintptr, value int32) {
   pageControl_SetHeight.Call(obj, uintptr(value))
}

func PageControl_GetCursor(obj uintptr) TCursor {
    ret, _, _ := pageControl_GetCursor.Call(obj)
    return TCursor(ret)
}

func PageControl_SetCursor(obj uintptr, value TCursor) {
   pageControl_SetCursor.Call(obj, uintptr(value))
}

func PageControl_GetHint(obj uintptr) string {
    ret, _, _ := pageControl_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func PageControl_SetHint(obj uintptr, value string) {
   pageControl_SetHint.Call(obj, GoStrToDStr(value))
}

func PageControl_GetMargins(obj uintptr) uintptr {
    ret, _, _ := pageControl_GetMargins.Call(obj)
    return ret
}

func PageControl_SetMargins(obj uintptr, value uintptr) {
   pageControl_SetMargins.Call(obj, value)
}

func PageControl_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := pageControl_GetComponentCount.Call(obj)
    return int32(ret)
}

func PageControl_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := pageControl_GetComponentIndex.Call(obj)
    return int32(ret)
}

func PageControl_SetComponentIndex(obj uintptr, value int32) {
   pageControl_SetComponentIndex.Call(obj, uintptr(value))
}

func PageControl_GetOwner(obj uintptr) uintptr {
    ret, _, _ := pageControl_GetOwner.Call(obj)
    return ret
}

func PageControl_GetName(obj uintptr) string {
    ret, _, _ := pageControl_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func PageControl_SetName(obj uintptr, value string) {
   pageControl_SetName.Call(obj, GoStrToDStr(value))
}

func PageControl_GetTag(obj uintptr) int {
    ret, _, _ := pageControl_GetTag.Call(obj)
    return int(ret)
}

func PageControl_SetTag(obj uintptr, value int) {
   pageControl_SetTag.Call(obj, uintptr(value))
}

func PageControl_GetPages(obj uintptr, Index int32) uintptr {
    ret, _, _ := pageControl_GetPages.Call(obj, uintptr(Index))
    return ret
}

func PageControl_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := pageControl_GetControls.Call(obj, uintptr(Index))
    return ret
}

func PageControl_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := pageControl_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TTabSheet ---------------------------

func TabSheet_Create(obj uintptr) uintptr {
    ret, _, _ := tabSheet_Create.Call(obj)
    return ret
}

func TabSheet_Free(obj uintptr) {
    tabSheet_Free.Call(obj)
}

func TabSheet_CanFocus(obj uintptr) bool {
    ret, _, _ := tabSheet_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_FlipChildren(obj uintptr, AllLevels bool)  {
    tabSheet_FlipChildren.Call(obj, GoBoolToDBool(AllLevels) )
}

func TabSheet_Focused(obj uintptr) bool {
    ret, _, _ := tabSheet_Focused.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_HandleAllocated(obj uintptr) bool {
    ret, _, _ := tabSheet_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_Invalidate(obj uintptr)  {
    tabSheet_Invalidate.Call(obj)
}

func TabSheet_Realign(obj uintptr)  {
    tabSheet_Realign.Call(obj)
}

func TabSheet_Repaint(obj uintptr)  {
    tabSheet_Repaint.Call(obj)
}

func TabSheet_ScaleBy(obj uintptr, M int32, D int32)  {
    tabSheet_ScaleBy.Call(obj, uintptr(M) , uintptr(D) )
}

func TabSheet_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    tabSheet_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func TabSheet_SetFocus(obj uintptr)  {
    tabSheet_SetFocus.Call(obj)
}

func TabSheet_Update(obj uintptr)  {
    tabSheet_Update.Call(obj)
}

func TabSheet_BringToFront(obj uintptr)  {
    tabSheet_BringToFront.Call(obj)
}

func TabSheet_HasParent(obj uintptr) bool {
    ret, _, _ := tabSheet_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_Hide(obj uintptr)  {
    tabSheet_Hide.Call(obj)
}

func TabSheet_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := tabSheet_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func TabSheet_Refresh(obj uintptr)  {
    tabSheet_Refresh.Call(obj)
}

func TabSheet_SendToBack(obj uintptr)  {
    tabSheet_SendToBack.Call(obj)
}

func TabSheet_Show(obj uintptr)  {
    tabSheet_Show.Call(obj)
}

func TabSheet_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := tabSheet_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func TabSheet_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := tabSheet_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func TabSheet_GetNamePath(obj uintptr) string {
    ret, _, _ := tabSheet_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func TabSheet_Assign(obj uintptr, Source uintptr)  {
    tabSheet_Assign.Call(obj, Source )
}

func TabSheet_ClassName(obj uintptr) string {
    ret, _, _ := tabSheet_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func TabSheet_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := tabSheet_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func TabSheet_GetHashCode(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetHashCode.Call(obj)
    return int32(ret)
}

func TabSheet_ToString(obj uintptr) string {
    ret, _, _ := tabSheet_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func TabSheet_GetPageControl(obj uintptr) uintptr {
    ret, _, _ := tabSheet_GetPageControl.Call(obj)
    return ret
}

func TabSheet_SetPageControl(obj uintptr, value uintptr) {
   tabSheet_SetPageControl.Call(obj, value)
}

func TabSheet_GetTabIndex(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetTabIndex.Call(obj)
    return int32(ret)
}

func TabSheet_GetBorderWidth(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetBorderWidth.Call(obj)
    return int32(ret)
}

func TabSheet_SetBorderWidth(obj uintptr, value int32) {
   tabSheet_SetBorderWidth.Call(obj, uintptr(value))
}

func TabSheet_GetCaption(obj uintptr) string {
    ret, _, _ := tabSheet_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func TabSheet_SetCaption(obj uintptr, value string) {
   tabSheet_SetCaption.Call(obj, GoStrToDStr(value))
}

func TabSheet_GetDoubleBuffered(obj uintptr) bool {
    ret, _, _ := tabSheet_GetDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetDoubleBuffered(obj uintptr, value bool) {
   tabSheet_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetEnabled(obj uintptr) bool {
    ret, _, _ := tabSheet_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetEnabled(obj uintptr, value bool) {
   tabSheet_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetFont(obj uintptr) uintptr {
    ret, _, _ := tabSheet_GetFont.Call(obj)
    return ret
}

func TabSheet_SetFont(obj uintptr, value uintptr) {
   tabSheet_SetFont.Call(obj, value)
}

func TabSheet_GetHeight(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetHeight.Call(obj)
    return int32(ret)
}

func TabSheet_SetHeight(obj uintptr, value int32) {
   tabSheet_SetHeight.Call(obj, uintptr(value))
}

func TabSheet_GetHighlighted(obj uintptr) bool {
    ret, _, _ := tabSheet_GetHighlighted.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetHighlighted(obj uintptr, value bool) {
   tabSheet_SetHighlighted.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetImageIndex(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetImageIndex.Call(obj)
    return int32(ret)
}

func TabSheet_SetImageIndex(obj uintptr, value int32) {
   tabSheet_SetImageIndex.Call(obj, uintptr(value))
}

func TabSheet_GetLeft(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetLeft.Call(obj)
    return int32(ret)
}

func TabSheet_SetLeft(obj uintptr, value int32) {
   tabSheet_SetLeft.Call(obj, uintptr(value))
}

func TabSheet_GetPageIndex(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetPageIndex.Call(obj)
    return int32(ret)
}

func TabSheet_SetPageIndex(obj uintptr, value int32) {
   tabSheet_SetPageIndex.Call(obj, uintptr(value))
}

func TabSheet_GetParentDoubleBuffered(obj uintptr) bool {
    ret, _, _ := tabSheet_GetParentDoubleBuffered.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetParentDoubleBuffered(obj uintptr, value bool) {
   tabSheet_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetParentFont(obj uintptr) bool {
    ret, _, _ := tabSheet_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetParentFont(obj uintptr, value bool) {
   tabSheet_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := tabSheet_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetParentShowHint(obj uintptr, value bool) {
   tabSheet_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := tabSheet_GetPopupMenu.Call(obj)
    return ret
}

func TabSheet_SetPopupMenu(obj uintptr, value uintptr) {
   tabSheet_SetPopupMenu.Call(obj, value)
}

func TabSheet_GetShowHint(obj uintptr) bool {
    ret, _, _ := tabSheet_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetShowHint(obj uintptr, value bool) {
   tabSheet_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetTabVisible(obj uintptr) bool {
    ret, _, _ := tabSheet_GetTabVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetTabVisible(obj uintptr, value bool) {
   tabSheet_SetTabVisible.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetTop(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetTop.Call(obj)
    return int32(ret)
}

func TabSheet_SetTop(obj uintptr, value int32) {
   tabSheet_SetTop.Call(obj, uintptr(value))
}

func TabSheet_GetVisible(obj uintptr) bool {
    ret, _, _ := tabSheet_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetVisible(obj uintptr, value bool) {
   tabSheet_SetVisible.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetWidth(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetWidth.Call(obj)
    return int32(ret)
}

func TabSheet_SetWidth(obj uintptr, value int32) {
   tabSheet_SetWidth.Call(obj, uintptr(value))
}

func TabSheet_SetOnEnter(obj uintptr, fn interface{}) {
    tabSheet_SetOnEnter.Call(obj, addEventToMap(fn))
}

func TabSheet_SetOnExit(obj uintptr, fn interface{}) {
    tabSheet_SetOnExit.Call(obj, addEventToMap(fn))
}

func TabSheet_SetOnHide(obj uintptr, fn interface{}) {
    tabSheet_SetOnHide.Call(obj, addEventToMap(fn))
}

func TabSheet_SetOnMouseDown(obj uintptr, fn interface{}) {
    tabSheet_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func TabSheet_SetOnMouseEnter(obj uintptr, fn interface{}) {
    tabSheet_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func TabSheet_SetOnMouseLeave(obj uintptr, fn interface{}) {
    tabSheet_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func TabSheet_SetOnMouseMove(obj uintptr, fn interface{}) {
    tabSheet_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func TabSheet_SetOnMouseUp(obj uintptr, fn interface{}) {
    tabSheet_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func TabSheet_SetOnResize(obj uintptr, fn interface{}) {
    tabSheet_SetOnResize.Call(obj, addEventToMap(fn))
}

func TabSheet_SetOnShow(obj uintptr, fn interface{}) {
    tabSheet_SetOnShow.Call(obj, addEventToMap(fn))
}

func TabSheet_GetBrush(obj uintptr) uintptr {
    ret, _, _ := tabSheet_GetBrush.Call(obj)
    return ret
}

func TabSheet_GetControlCount(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetControlCount.Call(obj)
    return int32(ret)
}

func TabSheet_GetHandle(obj uintptr) HWND {
    ret, _, _ := tabSheet_GetHandle.Call(obj)
    return HWND(ret)
}

func TabSheet_GetParentWindow(obj uintptr) HWND {
    ret, _, _ := tabSheet_GetParentWindow.Call(obj)
    return HWND(ret)
}

func TabSheet_SetParentWindow(obj uintptr, value HWND) {
   tabSheet_SetParentWindow.Call(obj, uintptr(value))
}

func TabSheet_GetTabOrder(obj uintptr) uint16 {
    ret, _, _ := tabSheet_GetTabOrder.Call(obj)
    return uint16(ret)
}

func TabSheet_SetTabOrder(obj uintptr, value uint16) {
   tabSheet_SetTabOrder.Call(obj, uintptr(value))
}

func TabSheet_GetTabStop(obj uintptr) bool {
    ret, _, _ := tabSheet_GetTabStop.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetTabStop(obj uintptr, value bool) {
   tabSheet_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetAction(obj uintptr) uintptr {
    ret, _, _ := tabSheet_GetAction.Call(obj)
    return ret
}

func TabSheet_SetAction(obj uintptr, value uintptr) {
   tabSheet_SetAction.Call(obj, value)
}

func TabSheet_GetAlign(obj uintptr) TAlign {
    ret, _, _ := tabSheet_GetAlign.Call(obj)
    return TAlign(ret)
}

func TabSheet_SetAlign(obj uintptr, value TAlign) {
   tabSheet_SetAlign.Call(obj, uintptr(value))
}

func TabSheet_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := tabSheet_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func TabSheet_SetAnchors(obj uintptr, value TAnchors) {
   tabSheet_SetAnchors.Call(obj, uintptr(value))
}

func TabSheet_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := tabSheet_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func TabSheet_SetBiDiMode(obj uintptr, value TBiDiMode) {
   tabSheet_SetBiDiMode.Call(obj, uintptr(value))
}

func TabSheet_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    tabSheet_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func TabSheet_SetBoundsRect(obj uintptr, value TRect) {
   tabSheet_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func TabSheet_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetClientHeight.Call(obj)
    return int32(ret)
}

func TabSheet_SetClientHeight(obj uintptr, value int32) {
   tabSheet_SetClientHeight.Call(obj, uintptr(value))
}

func TabSheet_GetClientRect(obj uintptr) TRect {
    var ret TRect
    tabSheet_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func TabSheet_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetClientWidth.Call(obj)
    return int32(ret)
}

func TabSheet_SetClientWidth(obj uintptr, value int32) {
   tabSheet_SetClientWidth.Call(obj, uintptr(value))
}

func TabSheet_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func TabSheet_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetExplicitTop.Call(obj)
    return int32(ret)
}

func TabSheet_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func TabSheet_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func TabSheet_GetParent(obj uintptr) uintptr {
    ret, _, _ := tabSheet_GetParent.Call(obj)
    return ret
}

func TabSheet_SetParent(obj uintptr, value uintptr) {
   tabSheet_SetParent.Call(obj, value)
}

func TabSheet_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := tabSheet_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func TabSheet_SetStyleElements(obj uintptr, value TStyleElements) {
   tabSheet_SetStyleElements.Call(obj, uintptr(value))
}

func TabSheet_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := tabSheet_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func TabSheet_SetAlignWithMargins(obj uintptr, value bool) {
   tabSheet_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func TabSheet_GetCursor(obj uintptr) TCursor {
    ret, _, _ := tabSheet_GetCursor.Call(obj)
    return TCursor(ret)
}

func TabSheet_SetCursor(obj uintptr, value TCursor) {
   tabSheet_SetCursor.Call(obj, uintptr(value))
}

func TabSheet_GetHint(obj uintptr) string {
    ret, _, _ := tabSheet_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func TabSheet_SetHint(obj uintptr, value string) {
   tabSheet_SetHint.Call(obj, GoStrToDStr(value))
}

func TabSheet_GetMargins(obj uintptr) uintptr {
    ret, _, _ := tabSheet_GetMargins.Call(obj)
    return ret
}

func TabSheet_SetMargins(obj uintptr, value uintptr) {
   tabSheet_SetMargins.Call(obj, value)
}

func TabSheet_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetComponentCount.Call(obj)
    return int32(ret)
}

func TabSheet_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := tabSheet_GetComponentIndex.Call(obj)
    return int32(ret)
}

func TabSheet_SetComponentIndex(obj uintptr, value int32) {
   tabSheet_SetComponentIndex.Call(obj, uintptr(value))
}

func TabSheet_GetOwner(obj uintptr) uintptr {
    ret, _, _ := tabSheet_GetOwner.Call(obj)
    return ret
}

func TabSheet_GetName(obj uintptr) string {
    ret, _, _ := tabSheet_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func TabSheet_SetName(obj uintptr, value string) {
   tabSheet_SetName.Call(obj, GoStrToDStr(value))
}

func TabSheet_GetTag(obj uintptr) int {
    ret, _, _ := tabSheet_GetTag.Call(obj)
    return int(ret)
}

func TabSheet_SetTag(obj uintptr, value int) {
   tabSheet_SetTag.Call(obj, uintptr(value))
}

func TabSheet_GetControls(obj uintptr, Index int32) uintptr {
    ret, _, _ := tabSheet_GetControls.Call(obj, uintptr(Index))
    return ret
}

func TabSheet_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := tabSheet_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TControl ---------------------------

func Control_Create(obj uintptr) uintptr {
    ret, _, _ := control_Create.Call(obj)
    return ret
}

func Control_Free(obj uintptr) {
    control_Free.Call(obj)
}

func Control_BringToFront(obj uintptr)  {
    control_BringToFront.Call(obj)
}

func Control_HasParent(obj uintptr) bool {
    ret, _, _ := control_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Control_Hide(obj uintptr)  {
    control_Hide.Call(obj)
}

func Control_Invalidate(obj uintptr)  {
    control_Invalidate.Call(obj)
}

func Control_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := control_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func Control_Refresh(obj uintptr)  {
    control_Refresh.Call(obj)
}

func Control_Repaint(obj uintptr)  {
    control_Repaint.Call(obj)
}

func Control_SendToBack(obj uintptr)  {
    control_SendToBack.Call(obj)
}

func Control_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    control_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func Control_Show(obj uintptr)  {
    control_Show.Call(obj)
}

func Control_Update(obj uintptr)  {
    control_Update.Call(obj)
}

func Control_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := control_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Control_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := control_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Control_GetNamePath(obj uintptr) string {
    ret, _, _ := control_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Control_Assign(obj uintptr, Source uintptr)  {
    control_Assign.Call(obj, Source )
}

func Control_ClassName(obj uintptr) string {
    ret, _, _ := control_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Control_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := control_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Control_GetHashCode(obj uintptr) int32 {
    ret, _, _ := control_GetHashCode.Call(obj)
    return int32(ret)
}

func Control_ToString(obj uintptr) string {
    ret, _, _ := control_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Control_GetEnabled(obj uintptr) bool {
    ret, _, _ := control_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Control_SetEnabled(obj uintptr, value bool) {
   control_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Control_GetAction(obj uintptr) uintptr {
    ret, _, _ := control_GetAction.Call(obj)
    return ret
}

func Control_SetAction(obj uintptr, value uintptr) {
   control_SetAction.Call(obj, value)
}

func Control_GetAlign(obj uintptr) TAlign {
    ret, _, _ := control_GetAlign.Call(obj)
    return TAlign(ret)
}

func Control_SetAlign(obj uintptr, value TAlign) {
   control_SetAlign.Call(obj, uintptr(value))
}

func Control_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := control_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func Control_SetAnchors(obj uintptr, value TAnchors) {
   control_SetAnchors.Call(obj, uintptr(value))
}

func Control_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := control_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func Control_SetBiDiMode(obj uintptr, value TBiDiMode) {
   control_SetBiDiMode.Call(obj, uintptr(value))
}

func Control_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    control_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Control_SetBoundsRect(obj uintptr, value TRect) {
   control_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Control_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := control_GetClientHeight.Call(obj)
    return int32(ret)
}

func Control_SetClientHeight(obj uintptr, value int32) {
   control_SetClientHeight.Call(obj, uintptr(value))
}

func Control_GetClientRect(obj uintptr) TRect {
    var ret TRect
    control_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Control_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := control_GetClientWidth.Call(obj)
    return int32(ret)
}

func Control_SetClientWidth(obj uintptr, value int32) {
   control_SetClientWidth.Call(obj, uintptr(value))
}

func Control_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := control_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Control_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := control_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Control_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := control_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Control_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := control_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Control_GetShowHint(obj uintptr) bool {
    ret, _, _ := control_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func Control_SetShowHint(obj uintptr, value bool) {
   control_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func Control_GetVisible(obj uintptr) bool {
    ret, _, _ := control_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Control_SetVisible(obj uintptr, value bool) {
   control_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Control_GetParent(obj uintptr) uintptr {
    ret, _, _ := control_GetParent.Call(obj)
    return ret
}

func Control_SetParent(obj uintptr, value uintptr) {
   control_SetParent.Call(obj, value)
}

func Control_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := control_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func Control_SetStyleElements(obj uintptr, value TStyleElements) {
   control_SetStyleElements.Call(obj, uintptr(value))
}

func Control_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := control_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func Control_SetAlignWithMargins(obj uintptr, value bool) {
   control_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func Control_GetLeft(obj uintptr) int32 {
    ret, _, _ := control_GetLeft.Call(obj)
    return int32(ret)
}

func Control_SetLeft(obj uintptr, value int32) {
   control_SetLeft.Call(obj, uintptr(value))
}

func Control_GetTop(obj uintptr) int32 {
    ret, _, _ := control_GetTop.Call(obj)
    return int32(ret)
}

func Control_SetTop(obj uintptr, value int32) {
   control_SetTop.Call(obj, uintptr(value))
}

func Control_GetWidth(obj uintptr) int32 {
    ret, _, _ := control_GetWidth.Call(obj)
    return int32(ret)
}

func Control_SetWidth(obj uintptr, value int32) {
   control_SetWidth.Call(obj, uintptr(value))
}

func Control_GetHeight(obj uintptr) int32 {
    ret, _, _ := control_GetHeight.Call(obj)
    return int32(ret)
}

func Control_SetHeight(obj uintptr, value int32) {
   control_SetHeight.Call(obj, uintptr(value))
}

func Control_GetCursor(obj uintptr) TCursor {
    ret, _, _ := control_GetCursor.Call(obj)
    return TCursor(ret)
}

func Control_SetCursor(obj uintptr, value TCursor) {
   control_SetCursor.Call(obj, uintptr(value))
}

func Control_GetHint(obj uintptr) string {
    ret, _, _ := control_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Control_SetHint(obj uintptr, value string) {
   control_SetHint.Call(obj, GoStrToDStr(value))
}

func Control_GetMargins(obj uintptr) uintptr {
    ret, _, _ := control_GetMargins.Call(obj)
    return ret
}

func Control_SetMargins(obj uintptr, value uintptr) {
   control_SetMargins.Call(obj, value)
}

func Control_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := control_GetComponentCount.Call(obj)
    return int32(ret)
}

func Control_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := control_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Control_SetComponentIndex(obj uintptr, value int32) {
   control_SetComponentIndex.Call(obj, uintptr(value))
}

func Control_GetOwner(obj uintptr) uintptr {
    ret, _, _ := control_GetOwner.Call(obj)
    return ret
}

func Control_GetName(obj uintptr) string {
    ret, _, _ := control_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Control_SetName(obj uintptr, value string) {
   control_SetName.Call(obj, GoStrToDStr(value))
}

func Control_GetTag(obj uintptr) int {
    ret, _, _ := control_GetTag.Call(obj)
    return int(ret)
}

func Control_SetTag(obj uintptr, value int) {
   control_SetTag.Call(obj, uintptr(value))
}

func Control_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := control_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TScreen ---------------------------

func Screen_Create(obj uintptr) uintptr {
    ret, _, _ := screen_Create.Call(obj)
    return ret
}

func Screen_Free(obj uintptr) {
    screen_Free.Call(obj)
}

func Screen_Realign(obj uintptr)  {
    screen_Realign.Call(obj)
}

func Screen_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := screen_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Screen_GetNamePath(obj uintptr) string {
    ret, _, _ := screen_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Screen_HasParent(obj uintptr) bool {
    ret, _, _ := screen_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Screen_Assign(obj uintptr, Source uintptr)  {
    screen_Assign.Call(obj, Source )
}

func Screen_ClassName(obj uintptr) string {
    ret, _, _ := screen_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Screen_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := screen_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Screen_GetHashCode(obj uintptr) int32 {
    ret, _, _ := screen_GetHashCode.Call(obj)
    return int32(ret)
}

func Screen_ToString(obj uintptr) string {
    ret, _, _ := screen_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Screen_GetActiveForm(obj uintptr) uintptr {
    ret, _, _ := screen_GetActiveForm.Call(obj)
    return ret
}

func Screen_GetCustomFormCount(obj uintptr) int32 {
    ret, _, _ := screen_GetCustomFormCount.Call(obj)
    return int32(ret)
}

func Screen_GetCursorCount(obj uintptr) int32 {
    ret, _, _ := screen_GetCursorCount.Call(obj)
    return int32(ret)
}

func Screen_GetCursor(obj uintptr) TCursor {
    ret, _, _ := screen_GetCursor.Call(obj)
    return TCursor(ret)
}

func Screen_SetCursor(obj uintptr, value TCursor) {
   screen_SetCursor.Call(obj, uintptr(value))
}

func Screen_GetFocusedForm(obj uintptr) uintptr {
    ret, _, _ := screen_GetFocusedForm.Call(obj)
    return ret
}

func Screen_SetFocusedForm(obj uintptr, value uintptr) {
   screen_SetFocusedForm.Call(obj, value)
}

func Screen_GetMonitorCount(obj uintptr) int32 {
    ret, _, _ := screen_GetMonitorCount.Call(obj)
    return int32(ret)
}

func Screen_GetDesktopRect(obj uintptr) TRect {
    var ret TRect
    screen_GetDesktopRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Screen_GetDesktopHeight(obj uintptr) int32 {
    ret, _, _ := screen_GetDesktopHeight.Call(obj)
    return int32(ret)
}

func Screen_GetDesktopLeft(obj uintptr) int32 {
    ret, _, _ := screen_GetDesktopLeft.Call(obj)
    return int32(ret)
}

func Screen_GetDesktopTop(obj uintptr) int32 {
    ret, _, _ := screen_GetDesktopTop.Call(obj)
    return int32(ret)
}

func Screen_GetDesktopWidth(obj uintptr) int32 {
    ret, _, _ := screen_GetDesktopWidth.Call(obj)
    return int32(ret)
}

func Screen_GetWorkAreaRect(obj uintptr) TRect {
    var ret TRect
    screen_GetWorkAreaRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Screen_GetWorkAreaHeight(obj uintptr) int32 {
    ret, _, _ := screen_GetWorkAreaHeight.Call(obj)
    return int32(ret)
}

func Screen_GetWorkAreaLeft(obj uintptr) int32 {
    ret, _, _ := screen_GetWorkAreaLeft.Call(obj)
    return int32(ret)
}

func Screen_GetWorkAreaTop(obj uintptr) int32 {
    ret, _, _ := screen_GetWorkAreaTop.Call(obj)
    return int32(ret)
}

func Screen_GetWorkAreaWidth(obj uintptr) int32 {
    ret, _, _ := screen_GetWorkAreaWidth.Call(obj)
    return int32(ret)
}

func Screen_GetFonts(obj uintptr) uintptr {
    ret, _, _ := screen_GetFonts.Call(obj)
    return ret
}

func Screen_GetFormCount(obj uintptr) int32 {
    ret, _, _ := screen_GetFormCount.Call(obj)
    return int32(ret)
}

func Screen_GetImes(obj uintptr) uintptr {
    ret, _, _ := screen_GetImes.Call(obj)
    return ret
}

func Screen_GetDefaultIme(obj uintptr) string {
    ret, _, _ := screen_GetDefaultIme.Call(obj)
    return DStrToGoStr(ret)
}

func Screen_GetHeight(obj uintptr) int32 {
    ret, _, _ := screen_GetHeight.Call(obj)
    return int32(ret)
}

func Screen_GetPixelsPerInch(obj uintptr) int32 {
    ret, _, _ := screen_GetPixelsPerInch.Call(obj)
    return int32(ret)
}

func Screen_GetPrimaryMonitor(obj uintptr) uintptr {
    ret, _, _ := screen_GetPrimaryMonitor.Call(obj)
    return ret
}

func Screen_GetWidth(obj uintptr) int32 {
    ret, _, _ := screen_GetWidth.Call(obj)
    return int32(ret)
}

func Screen_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := screen_GetComponentCount.Call(obj)
    return int32(ret)
}

func Screen_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := screen_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Screen_SetComponentIndex(obj uintptr, value int32) {
   screen_SetComponentIndex.Call(obj, uintptr(value))
}

func Screen_GetOwner(obj uintptr) uintptr {
    ret, _, _ := screen_GetOwner.Call(obj)
    return ret
}

func Screen_GetName(obj uintptr) string {
    ret, _, _ := screen_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Screen_SetName(obj uintptr, value string) {
   screen_SetName.Call(obj, GoStrToDStr(value))
}

func Screen_GetTag(obj uintptr) int {
    ret, _, _ := screen_GetTag.Call(obj)
    return int(ret)
}

func Screen_SetTag(obj uintptr, value int) {
   screen_SetTag.Call(obj, uintptr(value))
}

func Screen_GetCursors(obj uintptr, Index int32) HICON {
    ret, _, _ := screen_GetCursors.Call(obj, uintptr(Index))
    return HICON(ret)
}

func Screen_SetCursors(obj uintptr, Index int32, value HICON) {
   screen_SetCursors.Call(obj, uintptr(Index), uintptr(value))
}

func Screen_GetMonitors(obj uintptr, Index int32) uintptr {
    ret, _, _ := screen_GetMonitors.Call(obj, uintptr(Index))
    return ret
}

func Screen_GetForms(obj uintptr, Index int32) uintptr {
    ret, _, _ := screen_GetForms.Call(obj, uintptr(Index))
    return ret
}

func Screen_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := screen_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TMouse ---------------------------

func Mouse_Create() uintptr {
    ret, _, _ := mouse_Create.Call()
    return ret
}

func Mouse_Free(obj uintptr) {
    mouse_Free.Call(obj)
}

func Mouse_ClassName(obj uintptr) string {
    ret, _, _ := mouse_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Mouse_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := mouse_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Mouse_GetHashCode(obj uintptr) int32 {
    ret, _, _ := mouse_GetHashCode.Call(obj)
    return int32(ret)
}

func Mouse_ToString(obj uintptr) string {
    ret, _, _ := mouse_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Mouse_GetCapture(obj uintptr) HWND {
    ret, _, _ := mouse_GetCapture.Call(obj)
    return HWND(ret)
}

func Mouse_SetCapture(obj uintptr, value HWND) {
   mouse_SetCapture.Call(obj, uintptr(value))
}

func Mouse_GetCursorPos(obj uintptr) TPoint {
    var ret TPoint
    mouse_GetCursorPos.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Mouse_SetCursorPos(obj uintptr, value TPoint) {
   mouse_SetCursorPos.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func Mouse_GetIsDragging(obj uintptr) bool {
    ret, _, _ := mouse_GetIsDragging.Call(obj)
    return DBoolToGoBool(ret)
}

func Mouse_GetIsPanning(obj uintptr) bool {
    ret, _, _ := mouse_GetIsPanning.Call(obj)
    return DBoolToGoBool(ret)
}

func Mouse_GetWheelPresent(obj uintptr) bool {
    ret, _, _ := mouse_GetWheelPresent.Call(obj)
    return DBoolToGoBool(ret)
}

func Mouse_GetWheelScrollLines(obj uintptr) int32 {
    ret, _, _ := mouse_GetWheelScrollLines.Call(obj)
    return int32(ret)
}


//--------------------------- TListGroup ---------------------------

func ListGroup_Create() uintptr {
    ret, _, _ := listGroup_Create.Call()
    return ret
}

func ListGroup_Free(obj uintptr) {
    listGroup_Free.Call(obj)
}

func ListGroup_Assign(obj uintptr, Source uintptr)  {
    listGroup_Assign.Call(obj, Source )
}

func ListGroup_GetNamePath(obj uintptr) string {
    ret, _, _ := listGroup_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ListGroup_ClassName(obj uintptr) string {
    ret, _, _ := listGroup_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ListGroup_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := listGroup_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ListGroup_GetHashCode(obj uintptr) int32 {
    ret, _, _ := listGroup_GetHashCode.Call(obj)
    return int32(ret)
}

func ListGroup_ToString(obj uintptr) string {
    ret, _, _ := listGroup_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ListGroup_GetHeader(obj uintptr) string {
    ret, _, _ := listGroup_GetHeader.Call(obj)
    return DStrToGoStr(ret)
}

func ListGroup_SetHeader(obj uintptr, value string) {
   listGroup_SetHeader.Call(obj, GoStrToDStr(value))
}

func ListGroup_GetFooter(obj uintptr) string {
    ret, _, _ := listGroup_GetFooter.Call(obj)
    return DStrToGoStr(ret)
}

func ListGroup_SetFooter(obj uintptr, value string) {
   listGroup_SetFooter.Call(obj, GoStrToDStr(value))
}

func ListGroup_GetGroupID(obj uintptr) int32 {
    ret, _, _ := listGroup_GetGroupID.Call(obj)
    return int32(ret)
}

func ListGroup_SetGroupID(obj uintptr, value int32) {
   listGroup_SetGroupID.Call(obj, uintptr(value))
}

func ListGroup_GetState(obj uintptr) TListGroupStateSet {
    ret, _, _ := listGroup_GetState.Call(obj)
    return TListGroupStateSet(ret)
}

func ListGroup_SetState(obj uintptr, value TListGroupStateSet) {
   listGroup_SetState.Call(obj, uintptr(value))
}

func ListGroup_GetHeaderAlign(obj uintptr) TAlignment {
    ret, _, _ := listGroup_GetHeaderAlign.Call(obj)
    return TAlignment(ret)
}

func ListGroup_SetHeaderAlign(obj uintptr, value TAlignment) {
   listGroup_SetHeaderAlign.Call(obj, uintptr(value))
}

func ListGroup_GetFooterAlign(obj uintptr) TAlignment {
    ret, _, _ := listGroup_GetFooterAlign.Call(obj)
    return TAlignment(ret)
}

func ListGroup_SetFooterAlign(obj uintptr, value TAlignment) {
   listGroup_SetFooterAlign.Call(obj, uintptr(value))
}

func ListGroup_GetSubtitle(obj uintptr) string {
    ret, _, _ := listGroup_GetSubtitle.Call(obj)
    return DStrToGoStr(ret)
}

func ListGroup_SetSubtitle(obj uintptr, value string) {
   listGroup_SetSubtitle.Call(obj, GoStrToDStr(value))
}

func ListGroup_GetTitleImage(obj uintptr) int32 {
    ret, _, _ := listGroup_GetTitleImage.Call(obj)
    return int32(ret)
}

func ListGroup_SetTitleImage(obj uintptr, value int32) {
   listGroup_SetTitleImage.Call(obj, uintptr(value))
}

func ListGroup_GetIndex(obj uintptr) int32 {
    ret, _, _ := listGroup_GetIndex.Call(obj)
    return int32(ret)
}

func ListGroup_SetIndex(obj uintptr, value int32) {
   listGroup_SetIndex.Call(obj, uintptr(value))
}


//--------------------------- TListColumn ---------------------------

func ListColumn_Create() uintptr {
    ret, _, _ := listColumn_Create.Call()
    return ret
}

func ListColumn_Free(obj uintptr) {
    listColumn_Free.Call(obj)
}

func ListColumn_Assign(obj uintptr, Source uintptr)  {
    listColumn_Assign.Call(obj, Source )
}

func ListColumn_GetNamePath(obj uintptr) string {
    ret, _, _ := listColumn_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ListColumn_ClassName(obj uintptr) string {
    ret, _, _ := listColumn_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ListColumn_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := listColumn_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ListColumn_GetHashCode(obj uintptr) int32 {
    ret, _, _ := listColumn_GetHashCode.Call(obj)
    return int32(ret)
}

func ListColumn_ToString(obj uintptr) string {
    ret, _, _ := listColumn_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ListColumn_GetAlignment(obj uintptr) TAlignment {
    ret, _, _ := listColumn_GetAlignment.Call(obj)
    return TAlignment(ret)
}

func ListColumn_SetAlignment(obj uintptr, value TAlignment) {
   listColumn_SetAlignment.Call(obj, uintptr(value))
}

func ListColumn_GetAutoSize(obj uintptr) bool {
    ret, _, _ := listColumn_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func ListColumn_SetAutoSize(obj uintptr, value bool) {
   listColumn_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func ListColumn_GetCaption(obj uintptr) string {
    ret, _, _ := listColumn_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func ListColumn_SetCaption(obj uintptr, value string) {
   listColumn_SetCaption.Call(obj, GoStrToDStr(value))
}

func ListColumn_GetImageIndex(obj uintptr) int32 {
    ret, _, _ := listColumn_GetImageIndex.Call(obj)
    return int32(ret)
}

func ListColumn_SetImageIndex(obj uintptr, value int32) {
   listColumn_SetImageIndex.Call(obj, uintptr(value))
}

func ListColumn_GetTag(obj uintptr) int32 {
    ret, _, _ := listColumn_GetTag.Call(obj)
    return int32(ret)
}

func ListColumn_SetTag(obj uintptr, value int32) {
   listColumn_SetTag.Call(obj, uintptr(value))
}

func ListColumn_GetWidth(obj uintptr) int32 {
    ret, _, _ := listColumn_GetWidth.Call(obj)
    return int32(ret)
}

func ListColumn_SetWidth(obj uintptr, value int32) {
   listColumn_SetWidth.Call(obj, uintptr(value))
}

func ListColumn_GetIndex(obj uintptr) int32 {
    ret, _, _ := listColumn_GetIndex.Call(obj)
    return int32(ret)
}

func ListColumn_SetIndex(obj uintptr, value int32) {
   listColumn_SetIndex.Call(obj, uintptr(value))
}


//--------------------------- TCollectionItem ---------------------------

func CollectionItem_Create() uintptr {
    ret, _, _ := collectionItem_Create.Call()
    return ret
}

func CollectionItem_Free(obj uintptr) {
    collectionItem_Free.Call(obj)
}

func CollectionItem_GetNamePath(obj uintptr) string {
    ret, _, _ := collectionItem_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func CollectionItem_Assign(obj uintptr, Source uintptr)  {
    collectionItem_Assign.Call(obj, Source )
}

func CollectionItem_ClassName(obj uintptr) string {
    ret, _, _ := collectionItem_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func CollectionItem_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := collectionItem_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func CollectionItem_GetHashCode(obj uintptr) int32 {
    ret, _, _ := collectionItem_GetHashCode.Call(obj)
    return int32(ret)
}

func CollectionItem_ToString(obj uintptr) string {
    ret, _, _ := collectionItem_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func CollectionItem_GetIndex(obj uintptr) int32 {
    ret, _, _ := collectionItem_GetIndex.Call(obj)
    return int32(ret)
}

func CollectionItem_SetIndex(obj uintptr, value int32) {
   collectionItem_SetIndex.Call(obj, uintptr(value))
}


//--------------------------- TStatusPanels ---------------------------

func StatusPanels_Create() uintptr {
    ret, _, _ := statusPanels_Create.Call()
    return ret
}

func StatusPanels_Free(obj uintptr) {
    statusPanels_Free.Call(obj)
}

func StatusPanels_Add(obj uintptr) uintptr {
    ret, _, _ := statusPanels_Add.Call(obj)
    return ret
}

func StatusPanels_AddItem(obj uintptr, Item uintptr, Index int32) uintptr {
    ret, _, _ := statusPanels_AddItem.Call(obj, Item , uintptr(Index) )
    return ret
}

func StatusPanels_Insert(obj uintptr, Index int32) uintptr {
    ret, _, _ := statusPanels_Insert.Call(obj, uintptr(Index) )
    return ret
}

func StatusPanels_Owner(obj uintptr) uintptr {
    ret, _, _ := statusPanels_Owner.Call(obj)
    return ret
}

func StatusPanels_Assign(obj uintptr, Source uintptr)  {
    statusPanels_Assign.Call(obj, Source )
}

func StatusPanels_BeginUpdate(obj uintptr)  {
    statusPanels_BeginUpdate.Call(obj)
}

func StatusPanels_Clear(obj uintptr)  {
    statusPanels_Clear.Call(obj)
}

func StatusPanels_Delete(obj uintptr, Index int32)  {
    statusPanels_Delete.Call(obj, uintptr(Index) )
}

func StatusPanels_EndUpdate(obj uintptr)  {
    statusPanels_EndUpdate.Call(obj)
}

func StatusPanels_GetNamePath(obj uintptr) string {
    ret, _, _ := statusPanels_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func StatusPanels_ClassName(obj uintptr) string {
    ret, _, _ := statusPanels_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func StatusPanels_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := statusPanels_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func StatusPanels_GetHashCode(obj uintptr) int32 {
    ret, _, _ := statusPanels_GetHashCode.Call(obj)
    return int32(ret)
}

func StatusPanels_ToString(obj uintptr) string {
    ret, _, _ := statusPanels_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func StatusPanels_GetItems(obj uintptr, Index int32) uintptr {
    ret, _, _ := statusPanels_GetItems.Call(obj, uintptr(Index))
    return ret
}

func StatusPanels_SetItems(obj uintptr, Index int32, value uintptr) {
   statusPanels_SetItems.Call(obj, uintptr(Index), value)
}


//--------------------------- TStatusPanel ---------------------------

func StatusPanel_Create() uintptr {
    ret, _, _ := statusPanel_Create.Call()
    return ret
}

func StatusPanel_Free(obj uintptr) {
    statusPanel_Free.Call(obj)
}

func StatusPanel_Assign(obj uintptr, Source uintptr)  {
    statusPanel_Assign.Call(obj, Source )
}

func StatusPanel_GetNamePath(obj uintptr) string {
    ret, _, _ := statusPanel_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func StatusPanel_ClassName(obj uintptr) string {
    ret, _, _ := statusPanel_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func StatusPanel_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := statusPanel_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func StatusPanel_GetHashCode(obj uintptr) int32 {
    ret, _, _ := statusPanel_GetHashCode.Call(obj)
    return int32(ret)
}

func StatusPanel_ToString(obj uintptr) string {
    ret, _, _ := statusPanel_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func StatusPanel_GetAlignment(obj uintptr) TAlignment {
    ret, _, _ := statusPanel_GetAlignment.Call(obj)
    return TAlignment(ret)
}

func StatusPanel_SetAlignment(obj uintptr, value TAlignment) {
   statusPanel_SetAlignment.Call(obj, uintptr(value))
}

func StatusPanel_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := statusPanel_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func StatusPanel_SetBiDiMode(obj uintptr, value TBiDiMode) {
   statusPanel_SetBiDiMode.Call(obj, uintptr(value))
}

func StatusPanel_GetStyle(obj uintptr) TStatusPanelStyle {
    ret, _, _ := statusPanel_GetStyle.Call(obj)
    return TStatusPanelStyle(ret)
}

func StatusPanel_SetStyle(obj uintptr, value TStatusPanelStyle) {
   statusPanel_SetStyle.Call(obj, uintptr(value))
}

func StatusPanel_GetText(obj uintptr) string {
    ret, _, _ := statusPanel_GetText.Call(obj)
    return DStrToGoStr(ret)
}

func StatusPanel_SetText(obj uintptr, value string) {
   statusPanel_SetText.Call(obj, GoStrToDStr(value))
}

func StatusPanel_GetWidth(obj uintptr) int32 {
    ret, _, _ := statusPanel_GetWidth.Call(obj)
    return int32(ret)
}

func StatusPanel_SetWidth(obj uintptr, value int32) {
   statusPanel_SetWidth.Call(obj, uintptr(value))
}

func StatusPanel_GetIndex(obj uintptr) int32 {
    ret, _, _ := statusPanel_GetIndex.Call(obj)
    return int32(ret)
}

func StatusPanel_SetIndex(obj uintptr, value int32) {
   statusPanel_SetIndex.Call(obj, uintptr(value))
}


//--------------------------- TCanvas ---------------------------

func Canvas_Create() uintptr {
    ret, _, _ := canvas_Create.Call()
    return ret
}

func Canvas_Free(obj uintptr) {
    canvas_Free.Call(obj)
}

func Canvas_Arc(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32)  {
    canvas_Arc.Call(obj, uintptr(X1) , uintptr(Y1) , uintptr(X2) , uintptr(Y2) , uintptr(X3) , uintptr(Y3) , uintptr(X4) , uintptr(Y4) )
}

func Canvas_ArcTo(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32)  {
    canvas_ArcTo.Call(obj, uintptr(X1) , uintptr(Y1) , uintptr(X2) , uintptr(Y2) , uintptr(X3) , uintptr(Y3) , uintptr(X4) , uintptr(Y4) )
}

func Canvas_AngleArc(obj uintptr, X int32, Y int32, Radius uint32, StartAngle float32, SweepAngle float32)  {
    canvas_AngleArc.Call(obj, uintptr(X) , uintptr(Y) , uintptr(Radius) , uintptr(unsafe.Pointer(&StartAngle)), uintptr(unsafe.Pointer(&SweepAngle)))
}

func Canvas_Chord(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32)  {
    canvas_Chord.Call(obj, uintptr(X1) , uintptr(Y1) , uintptr(X2) , uintptr(Y2) , uintptr(X3) , uintptr(Y3) , uintptr(X4) , uintptr(Y4) )
}

func Canvas_Ellipse(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32)  {
    canvas_Ellipse.Call(obj, uintptr(X1) , uintptr(Y1) , uintptr(X2) , uintptr(Y2) )
}

func Canvas_FloodFill(obj uintptr, X int32, Y int32, Color TColor, FillStyle TFillStyle)  {
    canvas_FloodFill.Call(obj, uintptr(X) , uintptr(Y) , uintptr(Color) , uintptr(FillStyle) )
}

func Canvas_HandleAllocated(obj uintptr) bool {
    ret, _, _ := canvas_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func Canvas_LineTo(obj uintptr, X int32, Y int32)  {
    canvas_LineTo.Call(obj, uintptr(X) , uintptr(Y) )
}

func Canvas_MoveTo(obj uintptr, X int32, Y int32)  {
    canvas_MoveTo.Call(obj, uintptr(X) , uintptr(Y) )
}

func Canvas_Pie(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32, X4 int32, Y4 int32)  {
    canvas_Pie.Call(obj, uintptr(X1) , uintptr(Y1) , uintptr(X2) , uintptr(Y2) , uintptr(X3) , uintptr(Y3) , uintptr(X4) , uintptr(Y4) )
}

func Canvas_Rectangle(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32)  {
    canvas_Rectangle.Call(obj, uintptr(X1) , uintptr(Y1) , uintptr(X2) , uintptr(Y2) )
}

func Canvas_Refresh(obj uintptr)  {
    canvas_Refresh.Call(obj)
}

func Canvas_RoundRect(obj uintptr, X1 int32, Y1 int32, X2 int32, Y2 int32, X3 int32, Y3 int32)  {
    canvas_RoundRect.Call(obj, uintptr(X1) , uintptr(Y1) , uintptr(X2) , uintptr(Y2) , uintptr(X3) , uintptr(Y3) )
}

func Canvas_TextExtent(obj uintptr, Text string) TSize {
    var ret TSize
    canvas_TextExtent.Call(obj, GoStrToDStr(Text) , uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Canvas_TextOut(obj uintptr, X int32, Y int32, Text string)  {
    canvas_TextOut.Call(obj, uintptr(X) , uintptr(Y) , GoStrToDStr(Text) )
}

func Canvas_Lock(obj uintptr)  {
    canvas_Lock.Call(obj)
}

func Canvas_TextHeight(obj uintptr, Text string) int32 {
    ret, _, _ := canvas_TextHeight.Call(obj, GoStrToDStr(Text) )
    return int32(ret)
}

func Canvas_TextWidth(obj uintptr, Text string) int32 {
    ret, _, _ := canvas_TextWidth.Call(obj, GoStrToDStr(Text) )
    return int32(ret)
}

func Canvas_Assign(obj uintptr, Source uintptr)  {
    canvas_Assign.Call(obj, Source )
}

func Canvas_GetNamePath(obj uintptr) string {
    ret, _, _ := canvas_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Canvas_ClassName(obj uintptr) string {
    ret, _, _ := canvas_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Canvas_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := canvas_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Canvas_GetHashCode(obj uintptr) int32 {
    ret, _, _ := canvas_GetHashCode.Call(obj)
    return int32(ret)
}

func Canvas_ToString(obj uintptr) string {
    ret, _, _ := canvas_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Canvas_GetHandle(obj uintptr) HDC {
    ret, _, _ := canvas_GetHandle.Call(obj)
    return HDC(ret)
}

func Canvas_SetHandle(obj uintptr, value HDC) {
   canvas_SetHandle.Call(obj, uintptr(value))
}

func Canvas_GetBrush(obj uintptr) uintptr {
    ret, _, _ := canvas_GetBrush.Call(obj)
    return ret
}

func Canvas_SetBrush(obj uintptr, value uintptr) {
   canvas_SetBrush.Call(obj, value)
}

func Canvas_GetCopyMode(obj uintptr) int32 {
    ret, _, _ := canvas_GetCopyMode.Call(obj)
    return int32(ret)
}

func Canvas_SetCopyMode(obj uintptr, value int32) {
   canvas_SetCopyMode.Call(obj, uintptr(value))
}

func Canvas_GetFont(obj uintptr) uintptr {
    ret, _, _ := canvas_GetFont.Call(obj)
    return ret
}

func Canvas_SetFont(obj uintptr, value uintptr) {
   canvas_SetFont.Call(obj, value)
}

func Canvas_GetPen(obj uintptr) uintptr {
    ret, _, _ := canvas_GetPen.Call(obj)
    return ret
}

func Canvas_SetPen(obj uintptr, value uintptr) {
   canvas_SetPen.Call(obj, value)
}

func Canvas_SetOnChange(obj uintptr, fn interface{}) {
    canvas_SetOnChange.Call(obj, addEventToMap(fn))
}


//--------------------------- TObject ---------------------------

func Object_Create() uintptr {
    ret, _, _ := object_Create.Call()
    return ret
}

func Object_Free(obj uintptr) {
    object_Free.Call(obj)
}

func Object_ClassName(obj uintptr) string {
    ret, _, _ := object_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Object_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := object_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Object_GetHashCode(obj uintptr) int32 {
    ret, _, _ := object_GetHashCode.Call(obj)
    return int32(ret)
}

func Object_ToString(obj uintptr) string {
    ret, _, _ := object_ToString.Call(obj)
    return DStrToGoStr(ret)
}


//--------------------------- TPngImage ---------------------------

func PngImage_Create() uintptr {
    ret, _, _ := pngImage_Create.Call()
    return ret
}

func PngImage_Free(obj uintptr) {
    pngImage_Free.Call(obj)
}

func PngImage_Assign(obj uintptr, Source uintptr)  {
    pngImage_Assign.Call(obj, Source )
}

func PngImage_LoadFromStream(obj uintptr, Stream uintptr)  {
    pngImage_LoadFromStream.Call(obj, Stream )
}

func PngImage_SaveToStream(obj uintptr, Stream uintptr)  {
    pngImage_SaveToStream.Call(obj, Stream )
}

func PngImage_LoadFromResourceName(obj uintptr, Instance uintptr, Name string)  {
    pngImage_LoadFromResourceName.Call(obj, Instance , GoStrToDStr(Name) )
}

func PngImage_LoadFromResourceID(obj uintptr, Instance uintptr, ResID int32)  {
    pngImage_LoadFromResourceID.Call(obj, Instance , uintptr(ResID) )
}

func PngImage_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := pngImage_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func PngImage_LoadFromFile(obj uintptr, Filename string)  {
    pngImage_LoadFromFile.Call(obj, GoStrToDStr(Filename) )
}

func PngImage_SaveToFile(obj uintptr, Filename string)  {
    pngImage_SaveToFile.Call(obj, GoStrToDStr(Filename) )
}

func PngImage_SetSize(obj uintptr, AWidth int32, AHeight int32)  {
    pngImage_SetSize.Call(obj, uintptr(AWidth) , uintptr(AHeight) )
}

func PngImage_GetNamePath(obj uintptr) string {
    ret, _, _ := pngImage_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func PngImage_ClassName(obj uintptr) string {
    ret, _, _ := pngImage_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func PngImage_GetHashCode(obj uintptr) int32 {
    ret, _, _ := pngImage_GetHashCode.Call(obj)
    return int32(ret)
}

func PngImage_ToString(obj uintptr) string {
    ret, _, _ := pngImage_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func PngImage_GetTransparentColor(obj uintptr) TColor {
    ret, _, _ := pngImage_GetTransparentColor.Call(obj)
    return TColor(ret)
}

func PngImage_SetTransparentColor(obj uintptr, value TColor) {
   pngImage_SetTransparentColor.Call(obj, uintptr(value))
}

func PngImage_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := pngImage_GetCanvas.Call(obj)
    return ret
}

func PngImage_GetWidth(obj uintptr) int32 {
    ret, _, _ := pngImage_GetWidth.Call(obj)
    return int32(ret)
}

func PngImage_GetHeight(obj uintptr) int32 {
    ret, _, _ := pngImage_GetHeight.Call(obj)
    return int32(ret)
}

func PngImage_GetMaxIdatSize(obj uintptr) int32 {
    ret, _, _ := pngImage_GetMaxIdatSize.Call(obj)
    return int32(ret)
}

func PngImage_SetMaxIdatSize(obj uintptr, value int32) {
   pngImage_SetMaxIdatSize.Call(obj, uintptr(value))
}

func PngImage_GetEmpty(obj uintptr) bool {
    ret, _, _ := pngImage_GetEmpty.Call(obj)
    return DBoolToGoBool(ret)
}

func PngImage_GetCompressionLevel(obj uintptr) TCompressionLevel {
    ret, _, _ := pngImage_GetCompressionLevel.Call(obj)
    return TCompressionLevel(ret)
}

func PngImage_SetCompressionLevel(obj uintptr, value TCompressionLevel) {
   pngImage_SetCompressionLevel.Call(obj, uintptr(value))
}

func PngImage_GetVersion(obj uintptr) string {
    ret, _, _ := pngImage_GetVersion.Call(obj)
    return DStrToGoStr(ret)
}

func PngImage_GetModified(obj uintptr) bool {
    ret, _, _ := pngImage_GetModified.Call(obj)
    return DBoolToGoBool(ret)
}

func PngImage_SetModified(obj uintptr, value bool) {
   pngImage_SetModified.Call(obj, GoBoolToDBool(value))
}

func PngImage_GetPaletteModified(obj uintptr) bool {
    ret, _, _ := pngImage_GetPaletteModified.Call(obj)
    return DBoolToGoBool(ret)
}

func PngImage_SetPaletteModified(obj uintptr, value bool) {
   pngImage_SetPaletteModified.Call(obj, GoBoolToDBool(value))
}

func PngImage_GetTransparent(obj uintptr) bool {
    ret, _, _ := pngImage_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func PngImage_SetTransparent(obj uintptr, value bool) {
   pngImage_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func PngImage_SetOnChange(obj uintptr, fn interface{}) {
    pngImage_SetOnChange.Call(obj, addEventToMap(fn))
}


//--------------------------- TJPEGImage ---------------------------

func JPEGImage_Create() uintptr {
    ret, _, _ := jPEGImage_Create.Call()
    return ret
}

func JPEGImage_Free(obj uintptr) {
    jPEGImage_Free.Call(obj)
}

func JPEGImage_Assign(obj uintptr, Source uintptr)  {
    jPEGImage_Assign.Call(obj, Source )
}

func JPEGImage_LoadFromStream(obj uintptr, Stream uintptr)  {
    jPEGImage_LoadFromStream.Call(obj, Stream )
}

func JPEGImage_SaveToStream(obj uintptr, Stream uintptr)  {
    jPEGImage_SaveToStream.Call(obj, Stream )
}

func JPEGImage_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := jPEGImage_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func JPEGImage_LoadFromFile(obj uintptr, Filename string)  {
    jPEGImage_LoadFromFile.Call(obj, GoStrToDStr(Filename) )
}

func JPEGImage_SaveToFile(obj uintptr, Filename string)  {
    jPEGImage_SaveToFile.Call(obj, GoStrToDStr(Filename) )
}

func JPEGImage_SetSize(obj uintptr, AWidth int32, AHeight int32)  {
    jPEGImage_SetSize.Call(obj, uintptr(AWidth) , uintptr(AHeight) )
}

func JPEGImage_GetNamePath(obj uintptr) string {
    ret, _, _ := jPEGImage_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func JPEGImage_ClassName(obj uintptr) string {
    ret, _, _ := jPEGImage_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func JPEGImage_GetHashCode(obj uintptr) int32 {
    ret, _, _ := jPEGImage_GetHashCode.Call(obj)
    return int32(ret)
}

func JPEGImage_ToString(obj uintptr) string {
    ret, _, _ := jPEGImage_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func JPEGImage_GetPixelFormat(obj uintptr) TJPEGPixelFormat {
    ret, _, _ := jPEGImage_GetPixelFormat.Call(obj)
    return TJPEGPixelFormat(ret)
}

func JPEGImage_SetPixelFormat(obj uintptr, value TJPEGPixelFormat) {
   jPEGImage_SetPixelFormat.Call(obj, uintptr(value))
}

func JPEGImage_GetProgressiveDisplay(obj uintptr) bool {
    ret, _, _ := jPEGImage_GetProgressiveDisplay.Call(obj)
    return DBoolToGoBool(ret)
}

func JPEGImage_SetProgressiveDisplay(obj uintptr, value bool) {
   jPEGImage_SetProgressiveDisplay.Call(obj, GoBoolToDBool(value))
}

func JPEGImage_GetPerformance(obj uintptr) TJPEGPerformance {
    ret, _, _ := jPEGImage_GetPerformance.Call(obj)
    return TJPEGPerformance(ret)
}

func JPEGImage_SetPerformance(obj uintptr, value TJPEGPerformance) {
   jPEGImage_SetPerformance.Call(obj, uintptr(value))
}

func JPEGImage_GetScale(obj uintptr) TJPEGScale {
    ret, _, _ := jPEGImage_GetScale.Call(obj)
    return TJPEGScale(ret)
}

func JPEGImage_SetScale(obj uintptr, value TJPEGScale) {
   jPEGImage_SetScale.Call(obj, uintptr(value))
}

func JPEGImage_GetSmoothing(obj uintptr) bool {
    ret, _, _ := jPEGImage_GetSmoothing.Call(obj)
    return DBoolToGoBool(ret)
}

func JPEGImage_SetSmoothing(obj uintptr, value bool) {
   jPEGImage_SetSmoothing.Call(obj, GoBoolToDBool(value))
}

func JPEGImage_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := jPEGImage_GetCanvas.Call(obj)
    return ret
}

func JPEGImage_GetEmpty(obj uintptr) bool {
    ret, _, _ := jPEGImage_GetEmpty.Call(obj)
    return DBoolToGoBool(ret)
}

func JPEGImage_GetHeight(obj uintptr) int32 {
    ret, _, _ := jPEGImage_GetHeight.Call(obj)
    return int32(ret)
}

func JPEGImage_SetHeight(obj uintptr, value int32) {
   jPEGImage_SetHeight.Call(obj, uintptr(value))
}

func JPEGImage_GetModified(obj uintptr) bool {
    ret, _, _ := jPEGImage_GetModified.Call(obj)
    return DBoolToGoBool(ret)
}

func JPEGImage_SetModified(obj uintptr, value bool) {
   jPEGImage_SetModified.Call(obj, GoBoolToDBool(value))
}

func JPEGImage_GetPaletteModified(obj uintptr) bool {
    ret, _, _ := jPEGImage_GetPaletteModified.Call(obj)
    return DBoolToGoBool(ret)
}

func JPEGImage_SetPaletteModified(obj uintptr, value bool) {
   jPEGImage_SetPaletteModified.Call(obj, GoBoolToDBool(value))
}

func JPEGImage_GetTransparent(obj uintptr) bool {
    ret, _, _ := jPEGImage_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func JPEGImage_SetTransparent(obj uintptr, value bool) {
   jPEGImage_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func JPEGImage_GetWidth(obj uintptr) int32 {
    ret, _, _ := jPEGImage_GetWidth.Call(obj)
    return int32(ret)
}

func JPEGImage_SetWidth(obj uintptr, value int32) {
   jPEGImage_SetWidth.Call(obj, uintptr(value))
}

func JPEGImage_SetOnChange(obj uintptr, fn interface{}) {
    jPEGImage_SetOnChange.Call(obj, addEventToMap(fn))
}


//--------------------------- TGIFImage ---------------------------

func GIFImage_Create() uintptr {
    ret, _, _ := gIFImage_Create.Call()
    return ret
}

func GIFImage_Free(obj uintptr) {
    gIFImage_Free.Call(obj)
}

func GIFImage_SaveToStream(obj uintptr, Stream uintptr)  {
    gIFImage_SaveToStream.Call(obj, Stream )
}

func GIFImage_LoadFromStream(obj uintptr, Stream uintptr)  {
    gIFImage_LoadFromStream.Call(obj, Stream )
}

func GIFImage_Add(obj uintptr, Source uintptr) uintptr {
    ret, _, _ := gIFImage_Add.Call(obj, Source )
    return ret
}

func GIFImage_Clear(obj uintptr)  {
    gIFImage_Clear.Call(obj)
}

func GIFImage_Assign(obj uintptr, Source uintptr)  {
    gIFImage_Assign.Call(obj, Source )
}

func GIFImage_StopDraw(obj uintptr)  {
    gIFImage_StopDraw.Call(obj)
}

func GIFImage_SuspendDraw(obj uintptr)  {
    gIFImage_SuspendDraw.Call(obj)
}

func GIFImage_ResumeDraw(obj uintptr)  {
    gIFImage_ResumeDraw.Call(obj)
}

func GIFImage_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := gIFImage_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func GIFImage_LoadFromFile(obj uintptr, Filename string)  {
    gIFImage_LoadFromFile.Call(obj, GoStrToDStr(Filename) )
}

func GIFImage_SaveToFile(obj uintptr, Filename string)  {
    gIFImage_SaveToFile.Call(obj, GoStrToDStr(Filename) )
}

func GIFImage_SetSize(obj uintptr, AWidth int32, AHeight int32)  {
    gIFImage_SetSize.Call(obj, uintptr(AWidth) , uintptr(AHeight) )
}

func GIFImage_GetNamePath(obj uintptr) string {
    ret, _, _ := gIFImage_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func GIFImage_ClassName(obj uintptr) string {
    ret, _, _ := gIFImage_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func GIFImage_GetHashCode(obj uintptr) int32 {
    ret, _, _ := gIFImage_GetHashCode.Call(obj)
    return int32(ret)
}

func GIFImage_ToString(obj uintptr) string {
    ret, _, _ := gIFImage_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func GIFImage_GetVersion(obj uintptr) TGIFVersion {
    ret, _, _ := gIFImage_GetVersion.Call(obj)
    return TGIFVersion(ret)
}

func GIFImage_GetBitsPerPixel(obj uintptr) int32 {
    ret, _, _ := gIFImage_GetBitsPerPixel.Call(obj)
    return int32(ret)
}

func GIFImage_GetBackgroundColor(obj uintptr) TColor {
    ret, _, _ := gIFImage_GetBackgroundColor.Call(obj)
    return TColor(ret)
}

func GIFImage_SetBackgroundColor(obj uintptr, value TColor) {
   gIFImage_SetBackgroundColor.Call(obj, uintptr(value))
}

func GIFImage_GetAspectRatio(obj uintptr) uint8 {
    ret, _, _ := gIFImage_GetAspectRatio.Call(obj)
    return uint8(ret)
}

func GIFImage_SetAspectRatio(obj uintptr, value uint8) {
   gIFImage_SetAspectRatio.Call(obj, uintptr(value))
}

func GIFImage_GetIsTransparent(obj uintptr) bool {
    ret, _, _ := gIFImage_GetIsTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func GIFImage_GetAnimationSpeed(obj uintptr) int32 {
    ret, _, _ := gIFImage_GetAnimationSpeed.Call(obj)
    return int32(ret)
}

func GIFImage_SetAnimationSpeed(obj uintptr, value int32) {
   gIFImage_SetAnimationSpeed.Call(obj, uintptr(value))
}

func GIFImage_GetBitmap(obj uintptr) uintptr {
    ret, _, _ := gIFImage_GetBitmap.Call(obj)
    return ret
}

func GIFImage_SetOnPaint(obj uintptr, fn interface{}) {
    gIFImage_SetOnPaint.Call(obj, addEventToMap(fn))
}

func GIFImage_GetAnimate(obj uintptr) bool {
    ret, _, _ := gIFImage_GetAnimate.Call(obj)
    return DBoolToGoBool(ret)
}

func GIFImage_SetAnimate(obj uintptr, value bool) {
   gIFImage_SetAnimate.Call(obj, GoBoolToDBool(value))
}

func GIFImage_GetAnimateLoop(obj uintptr) TGIFAnimationLoop {
    ret, _, _ := gIFImage_GetAnimateLoop.Call(obj)
    return TGIFAnimationLoop(ret)
}

func GIFImage_SetAnimateLoop(obj uintptr, value TGIFAnimationLoop) {
   gIFImage_SetAnimateLoop.Call(obj, uintptr(value))
}

func GIFImage_GetShouldDither(obj uintptr) bool {
    ret, _, _ := gIFImage_GetShouldDither.Call(obj)
    return DBoolToGoBool(ret)
}

func GIFImage_GetEmpty(obj uintptr) bool {
    ret, _, _ := gIFImage_GetEmpty.Call(obj)
    return DBoolToGoBool(ret)
}

func GIFImage_GetHeight(obj uintptr) int32 {
    ret, _, _ := gIFImage_GetHeight.Call(obj)
    return int32(ret)
}

func GIFImage_SetHeight(obj uintptr, value int32) {
   gIFImage_SetHeight.Call(obj, uintptr(value))
}

func GIFImage_GetModified(obj uintptr) bool {
    ret, _, _ := gIFImage_GetModified.Call(obj)
    return DBoolToGoBool(ret)
}

func GIFImage_SetModified(obj uintptr, value bool) {
   gIFImage_SetModified.Call(obj, GoBoolToDBool(value))
}

func GIFImage_GetPaletteModified(obj uintptr) bool {
    ret, _, _ := gIFImage_GetPaletteModified.Call(obj)
    return DBoolToGoBool(ret)
}

func GIFImage_SetPaletteModified(obj uintptr, value bool) {
   gIFImage_SetPaletteModified.Call(obj, GoBoolToDBool(value))
}

func GIFImage_GetTransparent(obj uintptr) bool {
    ret, _, _ := gIFImage_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func GIFImage_SetTransparent(obj uintptr, value bool) {
   gIFImage_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func GIFImage_GetWidth(obj uintptr) int32 {
    ret, _, _ := gIFImage_GetWidth.Call(obj)
    return int32(ret)
}

func GIFImage_SetWidth(obj uintptr, value int32) {
   gIFImage_SetWidth.Call(obj, uintptr(value))
}

func GIFImage_SetOnChange(obj uintptr, fn interface{}) {
    gIFImage_SetOnChange.Call(obj, addEventToMap(fn))
}


//--------------------------- TGIFFrame ---------------------------

func GIFFrame_Create() uintptr {
    ret, _, _ := gIFFrame_Create.Call()
    return ret
}

func GIFFrame_Free(obj uintptr) {
    gIFFrame_Free.Call(obj)
}

func GIFFrame_Clear(obj uintptr)  {
    gIFFrame_Clear.Call(obj)
}

func GIFFrame_SaveToStream(obj uintptr, Stream uintptr)  {
    gIFFrame_SaveToStream.Call(obj, Stream )
}

func GIFFrame_LoadFromStream(obj uintptr, Stream uintptr)  {
    gIFFrame_LoadFromStream.Call(obj, Stream )
}

func GIFFrame_Assign(obj uintptr, Source uintptr)  {
    gIFFrame_Assign.Call(obj, Source )
}

func GIFFrame_SaveToFile(obj uintptr, Filename string)  {
    gIFFrame_SaveToFile.Call(obj, GoStrToDStr(Filename) )
}

func GIFFrame_LoadFromFile(obj uintptr, Filename string)  {
    gIFFrame_LoadFromFile.Call(obj, GoStrToDStr(Filename) )
}

func GIFFrame_GetNamePath(obj uintptr) string {
    ret, _, _ := gIFFrame_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func GIFFrame_ClassName(obj uintptr) string {
    ret, _, _ := gIFFrame_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func GIFFrame_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := gIFFrame_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func GIFFrame_GetHashCode(obj uintptr) int32 {
    ret, _, _ := gIFFrame_GetHashCode.Call(obj)
    return int32(ret)
}

func GIFFrame_ToString(obj uintptr) string {
    ret, _, _ := gIFFrame_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func GIFFrame_GetHasBitmap(obj uintptr) bool {
    ret, _, _ := gIFFrame_GetHasBitmap.Call(obj)
    return DBoolToGoBool(ret)
}

func GIFFrame_SetHasBitmap(obj uintptr, value bool) {
   gIFFrame_SetHasBitmap.Call(obj, GoBoolToDBool(value))
}

func GIFFrame_GetLeft(obj uintptr) uint16 {
    ret, _, _ := gIFFrame_GetLeft.Call(obj)
    return uint16(ret)
}

func GIFFrame_SetLeft(obj uintptr, value uint16) {
   gIFFrame_SetLeft.Call(obj, uintptr(value))
}

func GIFFrame_GetTop(obj uintptr) uint16 {
    ret, _, _ := gIFFrame_GetTop.Call(obj)
    return uint16(ret)
}

func GIFFrame_SetTop(obj uintptr, value uint16) {
   gIFFrame_SetTop.Call(obj, uintptr(value))
}

func GIFFrame_GetWidth(obj uintptr) uint16 {
    ret, _, _ := gIFFrame_GetWidth.Call(obj)
    return uint16(ret)
}

func GIFFrame_SetWidth(obj uintptr, value uint16) {
   gIFFrame_SetWidth.Call(obj, uintptr(value))
}

func GIFFrame_GetHeight(obj uintptr) uint16 {
    ret, _, _ := gIFFrame_GetHeight.Call(obj)
    return uint16(ret)
}

func GIFFrame_SetHeight(obj uintptr, value uint16) {
   gIFFrame_SetHeight.Call(obj, uintptr(value))
}

func GIFFrame_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    gIFFrame_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func GIFFrame_SetBoundsRect(obj uintptr, value TRect) {
   gIFFrame_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func GIFFrame_GetClientRect(obj uintptr) TRect {
    var ret TRect
    gIFFrame_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func GIFFrame_GetData(obj uintptr) uintptr {
    ret, _, _ := gIFFrame_GetData.Call(obj)
    return ret
}

func GIFFrame_GetDataSize(obj uintptr) int32 {
    ret, _, _ := gIFFrame_GetDataSize.Call(obj)
    return int32(ret)
}

func GIFFrame_GetVersion(obj uintptr) TGIFVersion {
    ret, _, _ := gIFFrame_GetVersion.Call(obj)
    return TGIFVersion(ret)
}

func GIFFrame_GetBitsPerPixel(obj uintptr) int32 {
    ret, _, _ := gIFFrame_GetBitsPerPixel.Call(obj)
    return int32(ret)
}

func GIFFrame_GetBitmap(obj uintptr) uintptr {
    ret, _, _ := gIFFrame_GetBitmap.Call(obj)
    return ret
}

func GIFFrame_SetBitmap(obj uintptr, value uintptr) {
   gIFFrame_SetBitmap.Call(obj, value)
}

func GIFFrame_GetEmpty(obj uintptr) bool {
    ret, _, _ := gIFFrame_GetEmpty.Call(obj)
    return DBoolToGoBool(ret)
}

func GIFFrame_GetTransparent(obj uintptr) bool {
    ret, _, _ := gIFFrame_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}


//--------------------------- TActionList ---------------------------

func ActionList_Create(obj uintptr) uintptr {
    ret, _, _ := actionList_Create.Call(obj)
    return ret
}

func ActionList_Free(obj uintptr) {
    actionList_Free.Call(obj)
}

func ActionList_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := actionList_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ActionList_GetNamePath(obj uintptr) string {
    ret, _, _ := actionList_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ActionList_HasParent(obj uintptr) bool {
    ret, _, _ := actionList_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ActionList_Assign(obj uintptr, Source uintptr)  {
    actionList_Assign.Call(obj, Source )
}

func ActionList_ClassName(obj uintptr) string {
    ret, _, _ := actionList_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ActionList_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := actionList_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ActionList_GetHashCode(obj uintptr) int32 {
    ret, _, _ := actionList_GetHashCode.Call(obj)
    return int32(ret)
}

func ActionList_ToString(obj uintptr) string {
    ret, _, _ := actionList_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ActionList_GetImages(obj uintptr) uintptr {
    ret, _, _ := actionList_GetImages.Call(obj)
    return ret
}

func ActionList_SetImages(obj uintptr, value uintptr) {
   actionList_SetImages.Call(obj, value)
}

func ActionList_GetState(obj uintptr) TActionListState {
    ret, _, _ := actionList_GetState.Call(obj)
    return TActionListState(ret)
}

func ActionList_SetState(obj uintptr, value TActionListState) {
   actionList_SetState.Call(obj, uintptr(value))
}

func ActionList_SetOnChange(obj uintptr, fn interface{}) {
    actionList_SetOnChange.Call(obj, addEventToMap(fn))
}

func ActionList_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := actionList_GetComponentCount.Call(obj)
    return int32(ret)
}

func ActionList_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := actionList_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ActionList_SetComponentIndex(obj uintptr, value int32) {
   actionList_SetComponentIndex.Call(obj, uintptr(value))
}

func ActionList_GetOwner(obj uintptr) uintptr {
    ret, _, _ := actionList_GetOwner.Call(obj)
    return ret
}

func ActionList_GetName(obj uintptr) string {
    ret, _, _ := actionList_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ActionList_SetName(obj uintptr, value string) {
   actionList_SetName.Call(obj, GoStrToDStr(value))
}

func ActionList_GetTag(obj uintptr) int {
    ret, _, _ := actionList_GetTag.Call(obj)
    return int(ret)
}

func ActionList_SetTag(obj uintptr, value int) {
   actionList_SetTag.Call(obj, uintptr(value))
}

func ActionList_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := actionList_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TAction ---------------------------

func Action_Create(obj uintptr) uintptr {
    ret, _, _ := action_Create.Call(obj)
    return ret
}

func Action_Free(obj uintptr) {
    action_Free.Call(obj)
}

func Action_Execute(obj uintptr) bool {
    ret, _, _ := action_Execute.Call(obj)
    return DBoolToGoBool(ret)
}

func Action_Update(obj uintptr) bool {
    ret, _, _ := action_Update.Call(obj)
    return DBoolToGoBool(ret)
}

func Action_HasParent(obj uintptr) bool {
    ret, _, _ := action_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Action_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := action_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Action_GetNamePath(obj uintptr) string {
    ret, _, _ := action_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Action_Assign(obj uintptr, Source uintptr)  {
    action_Assign.Call(obj, Source )
}

func Action_ClassName(obj uintptr) string {
    ret, _, _ := action_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Action_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := action_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Action_GetHashCode(obj uintptr) int32 {
    ret, _, _ := action_GetHashCode.Call(obj)
    return int32(ret)
}

func Action_ToString(obj uintptr) string {
    ret, _, _ := action_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Action_GetCaption(obj uintptr) string {
    ret, _, _ := action_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func Action_SetCaption(obj uintptr, value string) {
   action_SetCaption.Call(obj, GoStrToDStr(value))
}

func Action_GetChecked(obj uintptr) bool {
    ret, _, _ := action_GetChecked.Call(obj)
    return DBoolToGoBool(ret)
}

func Action_SetChecked(obj uintptr, value bool) {
   action_SetChecked.Call(obj, GoBoolToDBool(value))
}

func Action_GetEnabled(obj uintptr) bool {
    ret, _, _ := action_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Action_SetEnabled(obj uintptr, value bool) {
   action_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Action_GetGroupIndex(obj uintptr) int32 {
    ret, _, _ := action_GetGroupIndex.Call(obj)
    return int32(ret)
}

func Action_SetGroupIndex(obj uintptr, value int32) {
   action_SetGroupIndex.Call(obj, uintptr(value))
}

func Action_GetHint(obj uintptr) string {
    ret, _, _ := action_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func Action_SetHint(obj uintptr, value string) {
   action_SetHint.Call(obj, GoStrToDStr(value))
}

func Action_GetImageIndex(obj uintptr) int32 {
    ret, _, _ := action_GetImageIndex.Call(obj)
    return int32(ret)
}

func Action_SetImageIndex(obj uintptr, value int32) {
   action_SetImageIndex.Call(obj, uintptr(value))
}

func Action_GetShortCut(obj uintptr) TShortCut {
    ret, _, _ := action_GetShortCut.Call(obj)
    return TShortCut(ret)
}

func Action_SetShortCut(obj uintptr, value TShortCut) {
   action_SetShortCut.Call(obj, uintptr(value))
}

func Action_GetVisible(obj uintptr) bool {
    ret, _, _ := action_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func Action_SetVisible(obj uintptr, value bool) {
   action_SetVisible.Call(obj, GoBoolToDBool(value))
}

func Action_SetOnExecute(obj uintptr, fn interface{}) {
    action_SetOnExecute.Call(obj, addEventToMap(fn))
}

func Action_SetOnUpdate(obj uintptr, fn interface{}) {
    action_SetOnUpdate.Call(obj, addEventToMap(fn))
}

func Action_GetImages(obj uintptr) uintptr {
    ret, _, _ := action_GetImages.Call(obj)
    return ret
}

func Action_GetIndex(obj uintptr) int32 {
    ret, _, _ := action_GetIndex.Call(obj)
    return int32(ret)
}

func Action_SetIndex(obj uintptr, value int32) {
   action_SetIndex.Call(obj, uintptr(value))
}

func Action_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := action_GetComponentCount.Call(obj)
    return int32(ret)
}

func Action_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := action_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Action_SetComponentIndex(obj uintptr, value int32) {
   action_SetComponentIndex.Call(obj, uintptr(value))
}

func Action_GetOwner(obj uintptr) uintptr {
    ret, _, _ := action_GetOwner.Call(obj)
    return ret
}

func Action_GetName(obj uintptr) string {
    ret, _, _ := action_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Action_SetName(obj uintptr, value string) {
   action_SetName.Call(obj, GoStrToDStr(value))
}

func Action_GetTag(obj uintptr) int {
    ret, _, _ := action_GetTag.Call(obj)
    return int(ret)
}

func Action_SetTag(obj uintptr, value int) {
   action_SetTag.Call(obj, uintptr(value))
}

func Action_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := action_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TToolButton ---------------------------

func ToolButton_Create(obj uintptr) uintptr {
    ret, _, _ := toolButton_Create.Call(obj)
    return ret
}

func ToolButton_Free(obj uintptr) {
    toolButton_Free.Call(obj)
}

func ToolButton_CheckMenuDropdown(obj uintptr) bool {
    ret, _, _ := toolButton_CheckMenuDropdown.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_Click(obj uintptr)  {
    toolButton_Click.Call(obj)
}

func ToolButton_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    toolButton_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func ToolButton_BringToFront(obj uintptr)  {
    toolButton_BringToFront.Call(obj)
}

func ToolButton_HasParent(obj uintptr) bool {
    ret, _, _ := toolButton_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_Hide(obj uintptr)  {
    toolButton_Hide.Call(obj)
}

func ToolButton_Invalidate(obj uintptr)  {
    toolButton_Invalidate.Call(obj)
}

func ToolButton_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := toolButton_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func ToolButton_Refresh(obj uintptr)  {
    toolButton_Refresh.Call(obj)
}

func ToolButton_Repaint(obj uintptr)  {
    toolButton_Repaint.Call(obj)
}

func ToolButton_SendToBack(obj uintptr)  {
    toolButton_SendToBack.Call(obj)
}

func ToolButton_Show(obj uintptr)  {
    toolButton_Show.Call(obj)
}

func ToolButton_Update(obj uintptr)  {
    toolButton_Update.Call(obj)
}

func ToolButton_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := toolButton_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func ToolButton_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := toolButton_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func ToolButton_GetNamePath(obj uintptr) string {
    ret, _, _ := toolButton_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ToolButton_Assign(obj uintptr, Source uintptr)  {
    toolButton_Assign.Call(obj, Source )
}

func ToolButton_ClassName(obj uintptr) string {
    ret, _, _ := toolButton_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ToolButton_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := toolButton_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ToolButton_GetHashCode(obj uintptr) int32 {
    ret, _, _ := toolButton_GetHashCode.Call(obj)
    return int32(ret)
}

func ToolButton_ToString(obj uintptr) string {
    ret, _, _ := toolButton_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ToolButton_GetIndex(obj uintptr) int32 {
    ret, _, _ := toolButton_GetIndex.Call(obj)
    return int32(ret)
}

func ToolButton_GetAction(obj uintptr) uintptr {
    ret, _, _ := toolButton_GetAction.Call(obj)
    return ret
}

func ToolButton_SetAction(obj uintptr, value uintptr) {
   toolButton_SetAction.Call(obj, value)
}

func ToolButton_GetAllowAllUp(obj uintptr) bool {
    ret, _, _ := toolButton_GetAllowAllUp.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetAllowAllUp(obj uintptr, value bool) {
   toolButton_SetAllowAllUp.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetAutoSize(obj uintptr) bool {
    ret, _, _ := toolButton_GetAutoSize.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetAutoSize(obj uintptr, value bool) {
   toolButton_SetAutoSize.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetCaption(obj uintptr) string {
    ret, _, _ := toolButton_GetCaption.Call(obj)
    return DStrToGoStr(ret)
}

func ToolButton_SetCaption(obj uintptr, value string) {
   toolButton_SetCaption.Call(obj, GoStrToDStr(value))
}

func ToolButton_GetDown(obj uintptr) bool {
    ret, _, _ := toolButton_GetDown.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetDown(obj uintptr, value bool) {
   toolButton_SetDown.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetDropdownMenu(obj uintptr) uintptr {
    ret, _, _ := toolButton_GetDropdownMenu.Call(obj)
    return ret
}

func ToolButton_SetDropdownMenu(obj uintptr, value uintptr) {
   toolButton_SetDropdownMenu.Call(obj, value)
}

func ToolButton_GetEnabled(obj uintptr) bool {
    ret, _, _ := toolButton_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetEnabled(obj uintptr, value bool) {
   toolButton_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetEnableDropdown(obj uintptr) bool {
    ret, _, _ := toolButton_GetEnableDropdown.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetEnableDropdown(obj uintptr, value bool) {
   toolButton_SetEnableDropdown.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetGrouped(obj uintptr) bool {
    ret, _, _ := toolButton_GetGrouped.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetGrouped(obj uintptr, value bool) {
   toolButton_SetGrouped.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetHeight(obj uintptr) int32 {
    ret, _, _ := toolButton_GetHeight.Call(obj)
    return int32(ret)
}

func ToolButton_SetHeight(obj uintptr, value int32) {
   toolButton_SetHeight.Call(obj, uintptr(value))
}

func ToolButton_GetImageIndex(obj uintptr) int32 {
    ret, _, _ := toolButton_GetImageIndex.Call(obj)
    return int32(ret)
}

func ToolButton_SetImageIndex(obj uintptr, value int32) {
   toolButton_SetImageIndex.Call(obj, uintptr(value))
}

func ToolButton_GetIndeterminate(obj uintptr) bool {
    ret, _, _ := toolButton_GetIndeterminate.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetIndeterminate(obj uintptr, value bool) {
   toolButton_SetIndeterminate.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetMarked(obj uintptr) bool {
    ret, _, _ := toolButton_GetMarked.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetMarked(obj uintptr, value bool) {
   toolButton_SetMarked.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := toolButton_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetParentShowHint(obj uintptr, value bool) {
   toolButton_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := toolButton_GetPopupMenu.Call(obj)
    return ret
}

func ToolButton_SetPopupMenu(obj uintptr, value uintptr) {
   toolButton_SetPopupMenu.Call(obj, value)
}

func ToolButton_GetWrap(obj uintptr) bool {
    ret, _, _ := toolButton_GetWrap.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetWrap(obj uintptr, value bool) {
   toolButton_SetWrap.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetShowHint(obj uintptr) bool {
    ret, _, _ := toolButton_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetShowHint(obj uintptr, value bool) {
   toolButton_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetStyle(obj uintptr) TToolButtonStyle {
    ret, _, _ := toolButton_GetStyle.Call(obj)
    return TToolButtonStyle(ret)
}

func ToolButton_SetStyle(obj uintptr, value TToolButtonStyle) {
   toolButton_SetStyle.Call(obj, uintptr(value))
}

func ToolButton_GetVisible(obj uintptr) bool {
    ret, _, _ := toolButton_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetVisible(obj uintptr, value bool) {
   toolButton_SetVisible.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetWidth(obj uintptr) int32 {
    ret, _, _ := toolButton_GetWidth.Call(obj)
    return int32(ret)
}

func ToolButton_SetWidth(obj uintptr, value int32) {
   toolButton_SetWidth.Call(obj, uintptr(value))
}

func ToolButton_SetOnClick(obj uintptr, fn interface{}) {
    toolButton_SetOnClick.Call(obj, addEventToMap(fn))
}

func ToolButton_SetOnMouseDown(obj uintptr, fn interface{}) {
    toolButton_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func ToolButton_SetOnMouseEnter(obj uintptr, fn interface{}) {
    toolButton_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func ToolButton_SetOnMouseLeave(obj uintptr, fn interface{}) {
    toolButton_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func ToolButton_SetOnMouseMove(obj uintptr, fn interface{}) {
    toolButton_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func ToolButton_SetOnMouseUp(obj uintptr, fn interface{}) {
    toolButton_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func ToolButton_GetAlign(obj uintptr) TAlign {
    ret, _, _ := toolButton_GetAlign.Call(obj)
    return TAlign(ret)
}

func ToolButton_SetAlign(obj uintptr, value TAlign) {
   toolButton_SetAlign.Call(obj, uintptr(value))
}

func ToolButton_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := toolButton_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func ToolButton_SetAnchors(obj uintptr, value TAnchors) {
   toolButton_SetAnchors.Call(obj, uintptr(value))
}

func ToolButton_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := toolButton_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func ToolButton_SetBiDiMode(obj uintptr, value TBiDiMode) {
   toolButton_SetBiDiMode.Call(obj, uintptr(value))
}

func ToolButton_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    toolButton_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ToolButton_SetBoundsRect(obj uintptr, value TRect) {
   toolButton_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func ToolButton_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := toolButton_GetClientHeight.Call(obj)
    return int32(ret)
}

func ToolButton_SetClientHeight(obj uintptr, value int32) {
   toolButton_SetClientHeight.Call(obj, uintptr(value))
}

func ToolButton_GetClientRect(obj uintptr) TRect {
    var ret TRect
    toolButton_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func ToolButton_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := toolButton_GetClientWidth.Call(obj)
    return int32(ret)
}

func ToolButton_SetClientWidth(obj uintptr, value int32) {
   toolButton_SetClientWidth.Call(obj, uintptr(value))
}

func ToolButton_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := toolButton_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func ToolButton_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := toolButton_GetExplicitTop.Call(obj)
    return int32(ret)
}

func ToolButton_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := toolButton_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func ToolButton_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := toolButton_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func ToolButton_GetParent(obj uintptr) uintptr {
    ret, _, _ := toolButton_GetParent.Call(obj)
    return ret
}

func ToolButton_SetParent(obj uintptr, value uintptr) {
   toolButton_SetParent.Call(obj, value)
}

func ToolButton_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := toolButton_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func ToolButton_SetStyleElements(obj uintptr, value TStyleElements) {
   toolButton_SetStyleElements.Call(obj, uintptr(value))
}

func ToolButton_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := toolButton_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func ToolButton_SetAlignWithMargins(obj uintptr, value bool) {
   toolButton_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func ToolButton_GetLeft(obj uintptr) int32 {
    ret, _, _ := toolButton_GetLeft.Call(obj)
    return int32(ret)
}

func ToolButton_SetLeft(obj uintptr, value int32) {
   toolButton_SetLeft.Call(obj, uintptr(value))
}

func ToolButton_GetTop(obj uintptr) int32 {
    ret, _, _ := toolButton_GetTop.Call(obj)
    return int32(ret)
}

func ToolButton_SetTop(obj uintptr, value int32) {
   toolButton_SetTop.Call(obj, uintptr(value))
}

func ToolButton_GetCursor(obj uintptr) TCursor {
    ret, _, _ := toolButton_GetCursor.Call(obj)
    return TCursor(ret)
}

func ToolButton_SetCursor(obj uintptr, value TCursor) {
   toolButton_SetCursor.Call(obj, uintptr(value))
}

func ToolButton_GetHint(obj uintptr) string {
    ret, _, _ := toolButton_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func ToolButton_SetHint(obj uintptr, value string) {
   toolButton_SetHint.Call(obj, GoStrToDStr(value))
}

func ToolButton_GetMargins(obj uintptr) uintptr {
    ret, _, _ := toolButton_GetMargins.Call(obj)
    return ret
}

func ToolButton_SetMargins(obj uintptr, value uintptr) {
   toolButton_SetMargins.Call(obj, value)
}

func ToolButton_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := toolButton_GetComponentCount.Call(obj)
    return int32(ret)
}

func ToolButton_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := toolButton_GetComponentIndex.Call(obj)
    return int32(ret)
}

func ToolButton_SetComponentIndex(obj uintptr, value int32) {
   toolButton_SetComponentIndex.Call(obj, uintptr(value))
}

func ToolButton_GetOwner(obj uintptr) uintptr {
    ret, _, _ := toolButton_GetOwner.Call(obj)
    return ret
}

func ToolButton_GetName(obj uintptr) string {
    ret, _, _ := toolButton_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func ToolButton_SetName(obj uintptr, value string) {
   toolButton_SetName.Call(obj, GoStrToDStr(value))
}

func ToolButton_GetTag(obj uintptr) int {
    ret, _, _ := toolButton_GetTag.Call(obj)
    return int(ret)
}

func ToolButton_SetTag(obj uintptr, value int) {
   toolButton_SetTag.Call(obj, uintptr(value))
}

func ToolButton_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := toolButton_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TIniFile ---------------------------

func IniFile_Create(filename string) uintptr {
    ret, _, _ := iniFile_Create.Call(GoStrToDStr(filename))
    return ret
}

func IniFile_Free(obj uintptr) {
    iniFile_Free.Call(obj)
}

func IniFile_ReadString(obj uintptr, Section string, Ident string, Default string) string {
    ret, _, _ := iniFile_ReadString.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Ident) , GoStrToDStr(Default) )
    return DStrToGoStr(ret)
}

func IniFile_WriteString(obj uintptr, Section string, Ident string, Value string)  {
    iniFile_WriteString.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Ident) , GoStrToDStr(Value) )
}

func IniFile_ReadSections(obj uintptr, Strings uintptr)  {
    iniFile_ReadSections.Call(obj, Strings )
}

func IniFile_ReadSectionValues(obj uintptr, Section string, Strings uintptr)  {
    iniFile_ReadSectionValues.Call(obj, GoStrToDStr(Section) , Strings )
}

func IniFile_EraseSection(obj uintptr, Section string)  {
    iniFile_EraseSection.Call(obj, GoStrToDStr(Section) )
}

func IniFile_DeleteKey(obj uintptr, Section string, Ident string)  {
    iniFile_DeleteKey.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Ident) )
}

func IniFile_UpdateFile(obj uintptr)  {
    iniFile_UpdateFile.Call(obj)
}

func IniFile_SectionExists(obj uintptr, Section string) bool {
    ret, _, _ := iniFile_SectionExists.Call(obj, GoStrToDStr(Section) )
    return DBoolToGoBool(ret)
}

func IniFile_ReadInteger(obj uintptr, Section string, Ident string, Default int32) int32 {
    ret, _, _ := iniFile_ReadInteger.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Ident) , uintptr(Default) )
    return int32(ret)
}

func IniFile_WriteInteger(obj uintptr, Section string, Ident string, Value int32)  {
    iniFile_WriteInteger.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Ident) , uintptr(Value) )
}

func IniFile_ReadBool(obj uintptr, Section string, Ident string, Default bool) bool {
    ret, _, _ := iniFile_ReadBool.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Ident) , GoBoolToDBool(Default) )
    return DBoolToGoBool(ret)
}

func IniFile_WriteBool(obj uintptr, Section string, Ident string, Value bool)  {
    iniFile_WriteBool.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Ident) , GoBoolToDBool(Value) )
}

func IniFile_ReadDate(obj uintptr, Section string, Name string, Default time.Time) time.Time {
    ret, _, _ := iniFile_ReadDate.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Name) , uintptr(Default.Unix()))
    return time.Unix(int64(ret), 0)
}

func IniFile_ReadDateTime(obj uintptr, Section string, Name string, Default time.Time) time.Time {
    ret, _, _ := iniFile_ReadDateTime.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Name) , uintptr(Default.Unix()))
    return time.Unix(int64(ret), 0)
}

func IniFile_ReadFloat(obj uintptr, Section string, Name string, Default float64) float64 {
    var ret float64
    iniFile_ReadFloat.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Name) , uintptr(unsafe.Pointer(&Default)), uintptr(unsafe.Pointer(&ret)))
    return ret
}

func IniFile_ReadTime(obj uintptr, Section string, Name string, Default time.Time) time.Time {
    ret, _, _ := iniFile_ReadTime.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Name) , uintptr(Default.Unix()))
    return time.Unix(int64(ret), 0)
}

func IniFile_WriteDate(obj uintptr, Section string, Name string, Value time.Time)  {
    iniFile_WriteDate.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Name) , uintptr(Value.Unix()))
}

func IniFile_WriteDateTime(obj uintptr, Section string, Name string, Value time.Time)  {
    iniFile_WriteDateTime.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Name) , uintptr(Value.Unix()))
}

func IniFile_WriteFloat(obj uintptr, Section string, Name string, Value float64)  {
    iniFile_WriteFloat.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Name) , uintptr(unsafe.Pointer(&Value)))
}

func IniFile_WriteTime(obj uintptr, Section string, Name string, Value time.Time)  {
    iniFile_WriteTime.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Name) , uintptr(Value.Unix()))
}

func IniFile_ReadSubSections(obj uintptr, Section string, Strings uintptr, Recurse bool)  {
    iniFile_ReadSubSections.Call(obj, GoStrToDStr(Section) , Strings , GoBoolToDBool(Recurse) )
}

func IniFile_ValueExists(obj uintptr, Section string, Ident string) bool {
    ret, _, _ := iniFile_ValueExists.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Ident) )
    return DBoolToGoBool(ret)
}

func IniFile_ClassName(obj uintptr) string {
    ret, _, _ := iniFile_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func IniFile_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := iniFile_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func IniFile_GetHashCode(obj uintptr) int32 {
    ret, _, _ := iniFile_GetHashCode.Call(obj)
    return int32(ret)
}

func IniFile_ToString(obj uintptr) string {
    ret, _, _ := iniFile_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func IniFile_GetFileName(obj uintptr) string {
    ret, _, _ := iniFile_GetFileName.Call(obj)
    return DStrToGoStr(ret)
}


//--------------------------- TRegistry ---------------------------

func Registry_Create(aAccess uint32) uintptr {
    ret, _, _ := registry_Create.Call(uintptr(aAccess))
    return ret
}

func Registry_Free(obj uintptr) {
    registry_Free.Call(obj)
}

func Registry_CloseKey(obj uintptr)  {
    registry_CloseKey.Call(obj)
}

func Registry_CreateKey(obj uintptr, Key string) bool {
    ret, _, _ := registry_CreateKey.Call(obj, GoStrToDStr(Key) )
    return DBoolToGoBool(ret)
}

func Registry_DeleteKey(obj uintptr, Key string) bool {
    ret, _, _ := registry_DeleteKey.Call(obj, GoStrToDStr(Key) )
    return DBoolToGoBool(ret)
}

func Registry_DeleteValue(obj uintptr, Name string) bool {
    ret, _, _ := registry_DeleteValue.Call(obj, GoStrToDStr(Name) )
    return DBoolToGoBool(ret)
}

func Registry_HasSubKeys(obj uintptr) bool {
    ret, _, _ := registry_HasSubKeys.Call(obj)
    return DBoolToGoBool(ret)
}

func Registry_KeyExists(obj uintptr, Key string) bool {
    ret, _, _ := registry_KeyExists.Call(obj, GoStrToDStr(Key) )
    return DBoolToGoBool(ret)
}

func Registry_LoadKey(obj uintptr, Key string, FileName string) bool {
    ret, _, _ := registry_LoadKey.Call(obj, GoStrToDStr(Key) , GoStrToDStr(FileName) )
    return DBoolToGoBool(ret)
}

func Registry_MoveKey(obj uintptr, OldName string, NewName string, Delete bool)  {
    registry_MoveKey.Call(obj, GoStrToDStr(OldName) , GoStrToDStr(NewName) , GoBoolToDBool(Delete) )
}

func Registry_OpenKey(obj uintptr, Key string, CanCreate bool) bool {
    ret, _, _ := registry_OpenKey.Call(obj, GoStrToDStr(Key) , GoBoolToDBool(CanCreate) )
    return DBoolToGoBool(ret)
}

func Registry_OpenKeyReadOnly(obj uintptr, Key string) bool {
    ret, _, _ := registry_OpenKeyReadOnly.Call(obj, GoStrToDStr(Key) )
    return DBoolToGoBool(ret)
}

func Registry_ReadBool(obj uintptr, Name string) bool {
    ret, _, _ := registry_ReadBool.Call(obj, GoStrToDStr(Name) )
    return DBoolToGoBool(ret)
}

func Registry_ReadDate(obj uintptr, Name string) time.Time {
    ret, _, _ := registry_ReadDate.Call(obj, GoStrToDStr(Name) )
    return time.Unix(int64(ret), 0)
}

func Registry_ReadDateTime(obj uintptr, Name string) time.Time {
    ret, _, _ := registry_ReadDateTime.Call(obj, GoStrToDStr(Name) )
    return time.Unix(int64(ret), 0)
}

func Registry_ReadFloat(obj uintptr, Name string) float64 {
    var ret float64
    registry_ReadFloat.Call(obj, GoStrToDStr(Name) , uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Registry_ReadInteger(obj uintptr, Name string) int32 {
    ret, _, _ := registry_ReadInteger.Call(obj, GoStrToDStr(Name) )
    return int32(ret)
}

func Registry_ReadString(obj uintptr, Name string) string {
    ret, _, _ := registry_ReadString.Call(obj, GoStrToDStr(Name) )
    return DStrToGoStr(ret)
}

func Registry_ReadTime(obj uintptr, Name string) time.Time {
    ret, _, _ := registry_ReadTime.Call(obj, GoStrToDStr(Name) )
    return time.Unix(int64(ret), 0)
}

func Registry_RegistryConnect(obj uintptr, UNCName string) bool {
    ret, _, _ := registry_RegistryConnect.Call(obj, GoStrToDStr(UNCName) )
    return DBoolToGoBool(ret)
}

func Registry_RenameValue(obj uintptr, OldName string, NewName string)  {
    registry_RenameValue.Call(obj, GoStrToDStr(OldName) , GoStrToDStr(NewName) )
}

func Registry_ReplaceKey(obj uintptr, Key string, FileName string, BackUpFileName string) bool {
    ret, _, _ := registry_ReplaceKey.Call(obj, GoStrToDStr(Key) , GoStrToDStr(FileName) , GoStrToDStr(BackUpFileName) )
    return DBoolToGoBool(ret)
}

func Registry_RestoreKey(obj uintptr, Key string, FileName string) bool {
    ret, _, _ := registry_RestoreKey.Call(obj, GoStrToDStr(Key) , GoStrToDStr(FileName) )
    return DBoolToGoBool(ret)
}

func Registry_SaveKey(obj uintptr, Key string, FileName string) bool {
    ret, _, _ := registry_SaveKey.Call(obj, GoStrToDStr(Key) , GoStrToDStr(FileName) )
    return DBoolToGoBool(ret)
}

func Registry_UnLoadKey(obj uintptr, Key string) bool {
    ret, _, _ := registry_UnLoadKey.Call(obj, GoStrToDStr(Key) )
    return DBoolToGoBool(ret)
}

func Registry_ValueExists(obj uintptr, Name string) bool {
    ret, _, _ := registry_ValueExists.Call(obj, GoStrToDStr(Name) )
    return DBoolToGoBool(ret)
}

func Registry_WriteBool(obj uintptr, Name string, Value bool)  {
    registry_WriteBool.Call(obj, GoStrToDStr(Name) , GoBoolToDBool(Value) )
}

func Registry_WriteDate(obj uintptr, Name string, Value time.Time)  {
    registry_WriteDate.Call(obj, GoStrToDStr(Name) , uintptr(Value.Unix()))
}

func Registry_WriteDateTime(obj uintptr, Name string, Value time.Time)  {
    registry_WriteDateTime.Call(obj, GoStrToDStr(Name) , uintptr(Value.Unix()))
}

func Registry_WriteFloat(obj uintptr, Name string, Value float64)  {
    registry_WriteFloat.Call(obj, GoStrToDStr(Name) , uintptr(unsafe.Pointer(&Value)))
}

func Registry_WriteInteger(obj uintptr, Name string, Value int32)  {
    registry_WriteInteger.Call(obj, GoStrToDStr(Name) , uintptr(Value) )
}

func Registry_WriteString(obj uintptr, Name string, Value string)  {
    registry_WriteString.Call(obj, GoStrToDStr(Name) , GoStrToDStr(Value) )
}

func Registry_WriteExpandString(obj uintptr, Name string, Value string)  {
    registry_WriteExpandString.Call(obj, GoStrToDStr(Name) , GoStrToDStr(Value) )
}

func Registry_WriteTime(obj uintptr, Name string, Value time.Time)  {
    registry_WriteTime.Call(obj, GoStrToDStr(Name) , uintptr(Value.Unix()))
}

func Registry_ClassName(obj uintptr) string {
    ret, _, _ := registry_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Registry_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := registry_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Registry_GetHashCode(obj uintptr) int32 {
    ret, _, _ := registry_GetHashCode.Call(obj)
    return int32(ret)
}

func Registry_ToString(obj uintptr) string {
    ret, _, _ := registry_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Registry_GetCurrentKey(obj uintptr) HKEY {
    ret, _, _ := registry_GetCurrentKey.Call(obj)
    return HKEY(ret)
}

func Registry_GetCurrentPath(obj uintptr) string {
    ret, _, _ := registry_GetCurrentPath.Call(obj)
    return DStrToGoStr(ret)
}

func Registry_GetLazyWrite(obj uintptr) bool {
    ret, _, _ := registry_GetLazyWrite.Call(obj)
    return DBoolToGoBool(ret)
}

func Registry_SetLazyWrite(obj uintptr, value bool) {
   registry_SetLazyWrite.Call(obj, GoBoolToDBool(value))
}

func Registry_GetLastError(obj uintptr) int32 {
    ret, _, _ := registry_GetLastError.Call(obj)
    return int32(ret)
}

func Registry_GetLastErrorMsg(obj uintptr) string {
    ret, _, _ := registry_GetLastErrorMsg.Call(obj)
    return DStrToGoStr(ret)
}

func Registry_GetRootKey(obj uintptr) HKEY {
    ret, _, _ := registry_GetRootKey.Call(obj)
    return HKEY(ret)
}

func Registry_SetRootKey(obj uintptr, value HKEY) {
   registry_SetRootKey.Call(obj, uintptr(value))
}

func Registry_GetRootKeyName(obj uintptr) string {
    ret, _, _ := registry_GetRootKeyName.Call(obj)
    return DStrToGoStr(ret)
}

func Registry_GetAccess(obj uintptr) uint32 {
    ret, _, _ := registry_GetAccess.Call(obj)
    return uint32(ret)
}

func Registry_SetAccess(obj uintptr, value uint32) {
   registry_SetAccess.Call(obj, uintptr(value))
}


//--------------------------- TClipboard ---------------------------

func Clipboard_Create() uintptr {
    ret, _, _ := clipboard_Create.Call()
    return ret
}

func Clipboard_Free(obj uintptr) {
    clipboard_Free.Call(obj)
}

func Clipboard_Assign(obj uintptr, Source uintptr)  {
    clipboard_Assign.Call(obj, Source )
}

func Clipboard_Clear(obj uintptr)  {
    clipboard_Clear.Call(obj)
}

func Clipboard_Close(obj uintptr)  {
    clipboard_Close.Call(obj)
}

func Clipboard_GetAsHandle(obj uintptr, Format uint16) uintptr {
    ret, _, _ := clipboard_GetAsHandle.Call(obj, uintptr(Format) )
    return ret
}

func Clipboard_HasFormat(obj uintptr, Format uint16) bool {
    ret, _, _ := clipboard_HasFormat.Call(obj, uintptr(Format) )
    return DBoolToGoBool(ret)
}

func Clipboard_Open(obj uintptr)  {
    clipboard_Open.Call(obj)
}

func Clipboard_SetAsHandle(obj uintptr, Format uint16, Value uintptr)  {
    clipboard_SetAsHandle.Call(obj, uintptr(Format) , Value )
}

func Clipboard_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := clipboard_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func Clipboard_GetNamePath(obj uintptr) string {
    ret, _, _ := clipboard_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Clipboard_ClassName(obj uintptr) string {
    ret, _, _ := clipboard_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Clipboard_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := clipboard_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Clipboard_GetHashCode(obj uintptr) int32 {
    ret, _, _ := clipboard_GetHashCode.Call(obj)
    return int32(ret)
}

func Clipboard_ToString(obj uintptr) string {
    ret, _, _ := clipboard_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Clipboard_GetAsText(obj uintptr) string {
    ret, _, _ := clipboard_GetAsText.Call(obj)
    return DStrToGoStr(ret)
}

func Clipboard_SetAsText(obj uintptr, value string) {
   clipboard_SetAsText.Call(obj, GoStrToDStr(value))
}

func Clipboard_GetFormatCount(obj uintptr) int32 {
    ret, _, _ := clipboard_GetFormatCount.Call(obj)
    return int32(ret)
}

func Clipboard_GetFormats(obj uintptr, Index int32) uint16 {
    ret, _, _ := clipboard_GetFormats.Call(obj, uintptr(Index))
    return uint16(ret)
}


//--------------------------- TMonitor ---------------------------

func Monitor_Create() uintptr {
    ret, _, _ := monitor_Create.Call()
    return ret
}

func Monitor_Free(obj uintptr) {
    monitor_Free.Call(obj)
}

func Monitor_ClassName(obj uintptr) string {
    ret, _, _ := monitor_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Monitor_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := monitor_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Monitor_GetHashCode(obj uintptr) int32 {
    ret, _, _ := monitor_GetHashCode.Call(obj)
    return int32(ret)
}

func Monitor_ToString(obj uintptr) string {
    ret, _, _ := monitor_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Monitor_GetHandle(obj uintptr) HMONITOR {
    ret, _, _ := monitor_GetHandle.Call(obj)
    return HMONITOR(ret)
}

func Monitor_GetMonitorNum(obj uintptr) int32 {
    ret, _, _ := monitor_GetMonitorNum.Call(obj)
    return int32(ret)
}

func Monitor_GetLeft(obj uintptr) int32 {
    ret, _, _ := monitor_GetLeft.Call(obj)
    return int32(ret)
}

func Monitor_GetHeight(obj uintptr) int32 {
    ret, _, _ := monitor_GetHeight.Call(obj)
    return int32(ret)
}

func Monitor_GetTop(obj uintptr) int32 {
    ret, _, _ := monitor_GetTop.Call(obj)
    return int32(ret)
}

func Monitor_GetWidth(obj uintptr) int32 {
    ret, _, _ := monitor_GetWidth.Call(obj)
    return int32(ret)
}

func Monitor_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    monitor_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Monitor_GetWorkareaRect(obj uintptr) TRect {
    var ret TRect
    monitor_GetWorkareaRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Monitor_GetPrimary(obj uintptr) bool {
    ret, _, _ := monitor_GetPrimary.Call(obj)
    return DBoolToGoBool(ret)
}

func Monitor_GetPixelsPerInch(obj uintptr) int32 {
    ret, _, _ := monitor_GetPixelsPerInch.Call(obj)
    return int32(ret)
}


//--------------------------- TMargins ---------------------------

func Margins_Create() uintptr {
    ret, _, _ := margins_Create.Call()
    return ret
}

func Margins_Free(obj uintptr) {
    margins_Free.Call(obj)
}

func Margins_SetBounds(obj uintptr, ALeft int32, ATop int32, ARight int32, ABottom int32)  {
    margins_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(ARight) , uintptr(ABottom) )
}

func Margins_Assign(obj uintptr, Source uintptr)  {
    margins_Assign.Call(obj, Source )
}

func Margins_GetNamePath(obj uintptr) string {
    ret, _, _ := margins_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Margins_ClassName(obj uintptr) string {
    ret, _, _ := margins_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Margins_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := margins_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Margins_GetHashCode(obj uintptr) int32 {
    ret, _, _ := margins_GetHashCode.Call(obj)
    return int32(ret)
}

func Margins_ToString(obj uintptr) string {
    ret, _, _ := margins_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Margins_GetControlLeft(obj uintptr) int32 {
    ret, _, _ := margins_GetControlLeft.Call(obj)
    return int32(ret)
}

func Margins_GetControlTop(obj uintptr) int32 {
    ret, _, _ := margins_GetControlTop.Call(obj)
    return int32(ret)
}

func Margins_GetControlWidth(obj uintptr) int32 {
    ret, _, _ := margins_GetControlWidth.Call(obj)
    return int32(ret)
}

func Margins_GetControlHeight(obj uintptr) int32 {
    ret, _, _ := margins_GetControlHeight.Call(obj)
    return int32(ret)
}

func Margins_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := margins_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func Margins_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := margins_GetExplicitTop.Call(obj)
    return int32(ret)
}

func Margins_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := margins_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func Margins_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := margins_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func Margins_SetOnChange(obj uintptr, fn interface{}) {
    margins_SetOnChange.Call(obj, addEventToMap(fn))
}

func Margins_GetLeft(obj uintptr) int32 {
    ret, _, _ := margins_GetLeft.Call(obj)
    return int32(ret)
}

func Margins_SetLeft(obj uintptr, value int32) {
   margins_SetLeft.Call(obj, uintptr(value))
}

func Margins_GetTop(obj uintptr) int32 {
    ret, _, _ := margins_GetTop.Call(obj)
    return int32(ret)
}

func Margins_SetTop(obj uintptr, value int32) {
   margins_SetTop.Call(obj, uintptr(value))
}

func Margins_GetRight(obj uintptr) int32 {
    ret, _, _ := margins_GetRight.Call(obj)
    return int32(ret)
}

func Margins_SetRight(obj uintptr, value int32) {
   margins_SetRight.Call(obj, uintptr(value))
}

func Margins_GetBottom(obj uintptr) int32 {
    ret, _, _ := margins_GetBottom.Call(obj)
    return int32(ret)
}

func Margins_SetBottom(obj uintptr, value int32) {
   margins_SetBottom.Call(obj, uintptr(value))
}


//--------------------------- TPaintBox ---------------------------

func PaintBox_Create(obj uintptr) uintptr {
    ret, _, _ := paintBox_Create.Call(obj)
    return ret
}

func PaintBox_Free(obj uintptr) {
    paintBox_Free.Call(obj)
}

func PaintBox_BringToFront(obj uintptr)  {
    paintBox_BringToFront.Call(obj)
}

func PaintBox_HasParent(obj uintptr) bool {
    ret, _, _ := paintBox_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func PaintBox_Hide(obj uintptr)  {
    paintBox_Hide.Call(obj)
}

func PaintBox_Invalidate(obj uintptr)  {
    paintBox_Invalidate.Call(obj)
}

func PaintBox_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
    ret, _, _ := paintBox_Perform.Call(obj, uintptr(Msg) , WParam , uintptr(LParam) )
    return int(ret)
}

func PaintBox_Refresh(obj uintptr)  {
    paintBox_Refresh.Call(obj)
}

func PaintBox_Repaint(obj uintptr)  {
    paintBox_Repaint.Call(obj)
}

func PaintBox_SendToBack(obj uintptr)  {
    paintBox_SendToBack.Call(obj)
}

func PaintBox_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    paintBox_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func PaintBox_Show(obj uintptr)  {
    paintBox_Show.Call(obj)
}

func PaintBox_Update(obj uintptr)  {
    paintBox_Update.Call(obj)
}

func PaintBox_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
    ret, _, _ := paintBox_GetTextBuf.Call(obj, GoStrToDStr(Buffer) , uintptr(BufSize) )
    return int32(ret)
}

func PaintBox_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := paintBox_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func PaintBox_GetNamePath(obj uintptr) string {
    ret, _, _ := paintBox_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func PaintBox_Assign(obj uintptr, Source uintptr)  {
    paintBox_Assign.Call(obj, Source )
}

func PaintBox_ClassName(obj uintptr) string {
    ret, _, _ := paintBox_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func PaintBox_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := paintBox_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func PaintBox_GetHashCode(obj uintptr) int32 {
    ret, _, _ := paintBox_GetHashCode.Call(obj)
    return int32(ret)
}

func PaintBox_ToString(obj uintptr) string {
    ret, _, _ := paintBox_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func PaintBox_GetCanvas(obj uintptr) uintptr {
    ret, _, _ := paintBox_GetCanvas.Call(obj)
    return ret
}

func PaintBox_GetAlign(obj uintptr) TAlign {
    ret, _, _ := paintBox_GetAlign.Call(obj)
    return TAlign(ret)
}

func PaintBox_SetAlign(obj uintptr, value TAlign) {
   paintBox_SetAlign.Call(obj, uintptr(value))
}

func PaintBox_GetAnchors(obj uintptr) TAnchors {
    ret, _, _ := paintBox_GetAnchors.Call(obj)
    return TAnchors(ret)
}

func PaintBox_SetAnchors(obj uintptr, value TAnchors) {
   paintBox_SetAnchors.Call(obj, uintptr(value))
}

func PaintBox_GetColor(obj uintptr) TColor {
    ret, _, _ := paintBox_GetColor.Call(obj)
    return TColor(ret)
}

func PaintBox_SetColor(obj uintptr, value TColor) {
   paintBox_SetColor.Call(obj, uintptr(value))
}

func PaintBox_GetEnabled(obj uintptr) bool {
    ret, _, _ := paintBox_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func PaintBox_SetEnabled(obj uintptr, value bool) {
   paintBox_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func PaintBox_GetFont(obj uintptr) uintptr {
    ret, _, _ := paintBox_GetFont.Call(obj)
    return ret
}

func PaintBox_SetFont(obj uintptr, value uintptr) {
   paintBox_SetFont.Call(obj, value)
}

func PaintBox_GetParentColor(obj uintptr) bool {
    ret, _, _ := paintBox_GetParentColor.Call(obj)
    return DBoolToGoBool(ret)
}

func PaintBox_SetParentColor(obj uintptr, value bool) {
   paintBox_SetParentColor.Call(obj, GoBoolToDBool(value))
}

func PaintBox_GetParentFont(obj uintptr) bool {
    ret, _, _ := paintBox_GetParentFont.Call(obj)
    return DBoolToGoBool(ret)
}

func PaintBox_SetParentFont(obj uintptr, value bool) {
   paintBox_SetParentFont.Call(obj, GoBoolToDBool(value))
}

func PaintBox_GetParentShowHint(obj uintptr) bool {
    ret, _, _ := paintBox_GetParentShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func PaintBox_SetParentShowHint(obj uintptr, value bool) {
   paintBox_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func PaintBox_GetPopupMenu(obj uintptr) uintptr {
    ret, _, _ := paintBox_GetPopupMenu.Call(obj)
    return ret
}

func PaintBox_SetPopupMenu(obj uintptr, value uintptr) {
   paintBox_SetPopupMenu.Call(obj, value)
}

func PaintBox_GetShowHint(obj uintptr) bool {
    ret, _, _ := paintBox_GetShowHint.Call(obj)
    return DBoolToGoBool(ret)
}

func PaintBox_SetShowHint(obj uintptr, value bool) {
   paintBox_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func PaintBox_GetVisible(obj uintptr) bool {
    ret, _, _ := paintBox_GetVisible.Call(obj)
    return DBoolToGoBool(ret)
}

func PaintBox_SetVisible(obj uintptr, value bool) {
   paintBox_SetVisible.Call(obj, GoBoolToDBool(value))
}

func PaintBox_SetOnClick(obj uintptr, fn interface{}) {
    paintBox_SetOnClick.Call(obj, addEventToMap(fn))
}

func PaintBox_SetOnDblClick(obj uintptr, fn interface{}) {
    paintBox_SetOnDblClick.Call(obj, addEventToMap(fn))
}

func PaintBox_SetOnMouseDown(obj uintptr, fn interface{}) {
    paintBox_SetOnMouseDown.Call(obj, addEventToMap(fn))
}

func PaintBox_SetOnMouseEnter(obj uintptr, fn interface{}) {
    paintBox_SetOnMouseEnter.Call(obj, addEventToMap(fn))
}

func PaintBox_SetOnMouseLeave(obj uintptr, fn interface{}) {
    paintBox_SetOnMouseLeave.Call(obj, addEventToMap(fn))
}

func PaintBox_SetOnMouseMove(obj uintptr, fn interface{}) {
    paintBox_SetOnMouseMove.Call(obj, addEventToMap(fn))
}

func PaintBox_SetOnMouseUp(obj uintptr, fn interface{}) {
    paintBox_SetOnMouseUp.Call(obj, addEventToMap(fn))
}

func PaintBox_SetOnPaint(obj uintptr, fn interface{}) {
    paintBox_SetOnPaint.Call(obj, addEventToMap(fn))
}

func PaintBox_GetAction(obj uintptr) uintptr {
    ret, _, _ := paintBox_GetAction.Call(obj)
    return ret
}

func PaintBox_SetAction(obj uintptr, value uintptr) {
   paintBox_SetAction.Call(obj, value)
}

func PaintBox_GetBiDiMode(obj uintptr) TBiDiMode {
    ret, _, _ := paintBox_GetBiDiMode.Call(obj)
    return TBiDiMode(ret)
}

func PaintBox_SetBiDiMode(obj uintptr, value TBiDiMode) {
   paintBox_SetBiDiMode.Call(obj, uintptr(value))
}

func PaintBox_GetBoundsRect(obj uintptr) TRect {
    var ret TRect
    paintBox_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func PaintBox_SetBoundsRect(obj uintptr, value TRect) {
   paintBox_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func PaintBox_GetClientHeight(obj uintptr) int32 {
    ret, _, _ := paintBox_GetClientHeight.Call(obj)
    return int32(ret)
}

func PaintBox_SetClientHeight(obj uintptr, value int32) {
   paintBox_SetClientHeight.Call(obj, uintptr(value))
}

func PaintBox_GetClientRect(obj uintptr) TRect {
    var ret TRect
    paintBox_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
    return ret
}

func PaintBox_GetClientWidth(obj uintptr) int32 {
    ret, _, _ := paintBox_GetClientWidth.Call(obj)
    return int32(ret)
}

func PaintBox_SetClientWidth(obj uintptr, value int32) {
   paintBox_SetClientWidth.Call(obj, uintptr(value))
}

func PaintBox_GetExplicitLeft(obj uintptr) int32 {
    ret, _, _ := paintBox_GetExplicitLeft.Call(obj)
    return int32(ret)
}

func PaintBox_GetExplicitTop(obj uintptr) int32 {
    ret, _, _ := paintBox_GetExplicitTop.Call(obj)
    return int32(ret)
}

func PaintBox_GetExplicitWidth(obj uintptr) int32 {
    ret, _, _ := paintBox_GetExplicitWidth.Call(obj)
    return int32(ret)
}

func PaintBox_GetExplicitHeight(obj uintptr) int32 {
    ret, _, _ := paintBox_GetExplicitHeight.Call(obj)
    return int32(ret)
}

func PaintBox_GetParent(obj uintptr) uintptr {
    ret, _, _ := paintBox_GetParent.Call(obj)
    return ret
}

func PaintBox_SetParent(obj uintptr, value uintptr) {
   paintBox_SetParent.Call(obj, value)
}

func PaintBox_GetStyleElements(obj uintptr) TStyleElements {
    ret, _, _ := paintBox_GetStyleElements.Call(obj)
    return TStyleElements(ret)
}

func PaintBox_SetStyleElements(obj uintptr, value TStyleElements) {
   paintBox_SetStyleElements.Call(obj, uintptr(value))
}

func PaintBox_GetAlignWithMargins(obj uintptr) bool {
    ret, _, _ := paintBox_GetAlignWithMargins.Call(obj)
    return DBoolToGoBool(ret)
}

func PaintBox_SetAlignWithMargins(obj uintptr, value bool) {
   paintBox_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func PaintBox_GetLeft(obj uintptr) int32 {
    ret, _, _ := paintBox_GetLeft.Call(obj)
    return int32(ret)
}

func PaintBox_SetLeft(obj uintptr, value int32) {
   paintBox_SetLeft.Call(obj, uintptr(value))
}

func PaintBox_GetTop(obj uintptr) int32 {
    ret, _, _ := paintBox_GetTop.Call(obj)
    return int32(ret)
}

func PaintBox_SetTop(obj uintptr, value int32) {
   paintBox_SetTop.Call(obj, uintptr(value))
}

func PaintBox_GetWidth(obj uintptr) int32 {
    ret, _, _ := paintBox_GetWidth.Call(obj)
    return int32(ret)
}

func PaintBox_SetWidth(obj uintptr, value int32) {
   paintBox_SetWidth.Call(obj, uintptr(value))
}

func PaintBox_GetHeight(obj uintptr) int32 {
    ret, _, _ := paintBox_GetHeight.Call(obj)
    return int32(ret)
}

func PaintBox_SetHeight(obj uintptr, value int32) {
   paintBox_SetHeight.Call(obj, uintptr(value))
}

func PaintBox_GetCursor(obj uintptr) TCursor {
    ret, _, _ := paintBox_GetCursor.Call(obj)
    return TCursor(ret)
}

func PaintBox_SetCursor(obj uintptr, value TCursor) {
   paintBox_SetCursor.Call(obj, uintptr(value))
}

func PaintBox_GetHint(obj uintptr) string {
    ret, _, _ := paintBox_GetHint.Call(obj)
    return DStrToGoStr(ret)
}

func PaintBox_SetHint(obj uintptr, value string) {
   paintBox_SetHint.Call(obj, GoStrToDStr(value))
}

func PaintBox_GetMargins(obj uintptr) uintptr {
    ret, _, _ := paintBox_GetMargins.Call(obj)
    return ret
}

func PaintBox_SetMargins(obj uintptr, value uintptr) {
   paintBox_SetMargins.Call(obj, value)
}

func PaintBox_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := paintBox_GetComponentCount.Call(obj)
    return int32(ret)
}

func PaintBox_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := paintBox_GetComponentIndex.Call(obj)
    return int32(ret)
}

func PaintBox_SetComponentIndex(obj uintptr, value int32) {
   paintBox_SetComponentIndex.Call(obj, uintptr(value))
}

func PaintBox_GetOwner(obj uintptr) uintptr {
    ret, _, _ := paintBox_GetOwner.Call(obj)
    return ret
}

func PaintBox_GetName(obj uintptr) string {
    ret, _, _ := paintBox_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func PaintBox_SetName(obj uintptr, value string) {
   paintBox_SetName.Call(obj, GoStrToDStr(value))
}

func PaintBox_GetTag(obj uintptr) int {
    ret, _, _ := paintBox_GetTag.Call(obj)
    return int(ret)
}

func PaintBox_SetTag(obj uintptr, value int) {
   paintBox_SetTag.Call(obj, uintptr(value))
}

func PaintBox_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := paintBox_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TTimer ---------------------------

func Timer_Create(obj uintptr) uintptr {
    ret, _, _ := timer_Create.Call(obj)
    return ret
}

func Timer_Free(obj uintptr) {
    timer_Free.Call(obj)
}

func Timer_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := timer_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Timer_GetNamePath(obj uintptr) string {
    ret, _, _ := timer_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Timer_HasParent(obj uintptr) bool {
    ret, _, _ := timer_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Timer_Assign(obj uintptr, Source uintptr)  {
    timer_Assign.Call(obj, Source )
}

func Timer_ClassName(obj uintptr) string {
    ret, _, _ := timer_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Timer_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := timer_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Timer_GetHashCode(obj uintptr) int32 {
    ret, _, _ := timer_GetHashCode.Call(obj)
    return int32(ret)
}

func Timer_ToString(obj uintptr) string {
    ret, _, _ := timer_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Timer_GetEnabled(obj uintptr) bool {
    ret, _, _ := timer_GetEnabled.Call(obj)
    return DBoolToGoBool(ret)
}

func Timer_SetEnabled(obj uintptr, value bool) {
   timer_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func Timer_GetInterval(obj uintptr) uint32 {
    ret, _, _ := timer_GetInterval.Call(obj)
    return uint32(ret)
}

func Timer_SetInterval(obj uintptr, value uint32) {
   timer_SetInterval.Call(obj, uintptr(value))
}

func Timer_SetOnTimer(obj uintptr, fn interface{}) {
    timer_SetOnTimer.Call(obj, addEventToMap(fn))
}

func Timer_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := timer_GetComponentCount.Call(obj)
    return int32(ret)
}

func Timer_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := timer_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Timer_SetComponentIndex(obj uintptr, value int32) {
   timer_SetComponentIndex.Call(obj, uintptr(value))
}

func Timer_GetOwner(obj uintptr) uintptr {
    ret, _, _ := timer_GetOwner.Call(obj)
    return ret
}

func Timer_GetName(obj uintptr) string {
    ret, _, _ := timer_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Timer_SetName(obj uintptr, value string) {
   timer_SetName.Call(obj, GoStrToDStr(value))
}

func Timer_GetTag(obj uintptr) int {
    ret, _, _ := timer_GetTag.Call(obj)
    return int(ret)
}

func Timer_SetTag(obj uintptr, value int) {
   timer_SetTag.Call(obj, uintptr(value))
}

func Timer_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := timer_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TList ---------------------------

func List_Create() uintptr {
    ret, _, _ := list_Create.Call()
    return ret
}

func List_Free(obj uintptr) {
    list_Free.Call(obj)
}

func List_Add(obj uintptr, Item uintptr) int32 {
    ret, _, _ := list_Add.Call(obj, Item )
    return int32(ret)
}

func List_Clear(obj uintptr)  {
    list_Clear.Call(obj)
}

func List_Delete(obj uintptr, Index int32)  {
    list_Delete.Call(obj, uintptr(Index) )
}

func List_Expand(obj uintptr) uintptr {
    ret, _, _ := list_Expand.Call(obj)
    return ret
}

func List_IndexOf(obj uintptr, Item uintptr) int32 {
    ret, _, _ := list_IndexOf.Call(obj, Item )
    return int32(ret)
}

func List_Insert(obj uintptr, Index int32, Item uintptr)  {
    list_Insert.Call(obj, uintptr(Index) , Item )
}

func List_Move(obj uintptr, CurIndex int32, NewIndex int32)  {
    list_Move.Call(obj, uintptr(CurIndex) , uintptr(NewIndex) )
}

func List_ClassName(obj uintptr) string {
    ret, _, _ := list_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func List_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := list_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func List_GetHashCode(obj uintptr) int32 {
    ret, _, _ := list_GetHashCode.Call(obj)
    return int32(ret)
}

func List_ToString(obj uintptr) string {
    ret, _, _ := list_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func List_GetList(obj uintptr) uintptr {
    ret, _, _ := list_GetList.Call(obj)
    return ret
}

func List_GetItems(obj uintptr, Index int32) uintptr {
    ret, _, _ := list_GetItems.Call(obj, uintptr(Index))
    return ret
}

func List_SetItems(obj uintptr, Index int32, value uintptr) {
   list_SetItems.Call(obj, uintptr(Index), value)
}


//--------------------------- TGraphic ---------------------------

func Graphic_Create() uintptr {
    ret, _, _ := graphic_Create.Call()
    return ret
}

func Graphic_Free(obj uintptr) {
    graphic_Free.Call(obj)
}

func Graphic_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := graphic_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Graphic_LoadFromFile(obj uintptr, Filename string)  {
    graphic_LoadFromFile.Call(obj, GoStrToDStr(Filename) )
}

func Graphic_SaveToFile(obj uintptr, Filename string)  {
    graphic_SaveToFile.Call(obj, GoStrToDStr(Filename) )
}

func Graphic_LoadFromStream(obj uintptr, Stream uintptr)  {
    graphic_LoadFromStream.Call(obj, Stream )
}

func Graphic_SaveToStream(obj uintptr, Stream uintptr)  {
    graphic_SaveToStream.Call(obj, Stream )
}

func Graphic_SetSize(obj uintptr, AWidth int32, AHeight int32)  {
    graphic_SetSize.Call(obj, uintptr(AWidth) , uintptr(AHeight) )
}

func Graphic_Assign(obj uintptr, Source uintptr)  {
    graphic_Assign.Call(obj, Source )
}

func Graphic_GetNamePath(obj uintptr) string {
    ret, _, _ := graphic_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Graphic_ClassName(obj uintptr) string {
    ret, _, _ := graphic_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Graphic_GetHashCode(obj uintptr) int32 {
    ret, _, _ := graphic_GetHashCode.Call(obj)
    return int32(ret)
}

func Graphic_ToString(obj uintptr) string {
    ret, _, _ := graphic_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Graphic_GetEmpty(obj uintptr) bool {
    ret, _, _ := graphic_GetEmpty.Call(obj)
    return DBoolToGoBool(ret)
}

func Graphic_GetHeight(obj uintptr) int32 {
    ret, _, _ := graphic_GetHeight.Call(obj)
    return int32(ret)
}

func Graphic_SetHeight(obj uintptr, value int32) {
   graphic_SetHeight.Call(obj, uintptr(value))
}

func Graphic_GetModified(obj uintptr) bool {
    ret, _, _ := graphic_GetModified.Call(obj)
    return DBoolToGoBool(ret)
}

func Graphic_SetModified(obj uintptr, value bool) {
   graphic_SetModified.Call(obj, GoBoolToDBool(value))
}

func Graphic_GetPaletteModified(obj uintptr) bool {
    ret, _, _ := graphic_GetPaletteModified.Call(obj)
    return DBoolToGoBool(ret)
}

func Graphic_SetPaletteModified(obj uintptr, value bool) {
   graphic_SetPaletteModified.Call(obj, GoBoolToDBool(value))
}

func Graphic_GetTransparent(obj uintptr) bool {
    ret, _, _ := graphic_GetTransparent.Call(obj)
    return DBoolToGoBool(ret)
}

func Graphic_SetTransparent(obj uintptr, value bool) {
   graphic_SetTransparent.Call(obj, GoBoolToDBool(value))
}

func Graphic_GetWidth(obj uintptr) int32 {
    ret, _, _ := graphic_GetWidth.Call(obj)
    return int32(ret)
}

func Graphic_SetWidth(obj uintptr, value int32) {
   graphic_SetWidth.Call(obj, uintptr(value))
}

func Graphic_SetOnChange(obj uintptr, fn interface{}) {
    graphic_SetOnChange.Call(obj, addEventToMap(fn))
}


//--------------------------- TComponent ---------------------------

func Component_Create(obj uintptr) uintptr {
    ret, _, _ := component_Create.Call(obj)
    return ret
}

func Component_Free(obj uintptr) {
    component_Free.Call(obj)
}

func Component_FindComponent(obj uintptr, AName string) uintptr {
    ret, _, _ := component_FindComponent.Call(obj, GoStrToDStr(AName) )
    return ret
}

func Component_GetNamePath(obj uintptr) string {
    ret, _, _ := component_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Component_HasParent(obj uintptr) bool {
    ret, _, _ := component_HasParent.Call(obj)
    return DBoolToGoBool(ret)
}

func Component_Assign(obj uintptr, Source uintptr)  {
    component_Assign.Call(obj, Source )
}

func Component_ClassName(obj uintptr) string {
    ret, _, _ := component_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Component_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := component_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Component_GetHashCode(obj uintptr) int32 {
    ret, _, _ := component_GetHashCode.Call(obj)
    return int32(ret)
}

func Component_ToString(obj uintptr) string {
    ret, _, _ := component_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Component_GetComponentCount(obj uintptr) int32 {
    ret, _, _ := component_GetComponentCount.Call(obj)
    return int32(ret)
}

func Component_GetComponentIndex(obj uintptr) int32 {
    ret, _, _ := component_GetComponentIndex.Call(obj)
    return int32(ret)
}

func Component_SetComponentIndex(obj uintptr, value int32) {
   component_SetComponentIndex.Call(obj, uintptr(value))
}

func Component_GetOwner(obj uintptr) uintptr {
    ret, _, _ := component_GetOwner.Call(obj)
    return ret
}

func Component_GetName(obj uintptr) string {
    ret, _, _ := component_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func Component_SetName(obj uintptr, value string) {
   component_SetName.Call(obj, GoStrToDStr(value))
}

func Component_GetTag(obj uintptr) int {
    ret, _, _ := component_GetTag.Call(obj)
    return int(ret)
}

func Component_SetTag(obj uintptr, value int) {
   component_SetTag.Call(obj, uintptr(value))
}

func Component_GetComponents(obj uintptr, AIndex int32) uintptr {
    ret, _, _ := component_GetComponents.Call(obj, uintptr(AIndex))
    return ret
}


//--------------------------- TMonthCalColors ---------------------------

func MonthCalColors_Create() uintptr {
    ret, _, _ := monthCalColors_Create.Call()
    return ret
}

func MonthCalColors_Free(obj uintptr) {
    monthCalColors_Free.Call(obj)
}

func MonthCalColors_Assign(obj uintptr, Source uintptr)  {
    monthCalColors_Assign.Call(obj, Source )
}

func MonthCalColors_GetNamePath(obj uintptr) string {
    ret, _, _ := monthCalColors_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func MonthCalColors_ClassName(obj uintptr) string {
    ret, _, _ := monthCalColors_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func MonthCalColors_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := monthCalColors_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func MonthCalColors_GetHashCode(obj uintptr) int32 {
    ret, _, _ := monthCalColors_GetHashCode.Call(obj)
    return int32(ret)
}

func MonthCalColors_ToString(obj uintptr) string {
    ret, _, _ := monthCalColors_ToString.Call(obj)
    return DStrToGoStr(ret)
}


//--------------------------- TParaAttributes ---------------------------

func ParaAttributes_Assign(obj uintptr, Source uintptr)  {
    paraAttributes_Assign.Call(obj, Source )
}

func ParaAttributes_GetNamePath(obj uintptr) string {
    ret, _, _ := paraAttributes_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ParaAttributes_ClassName(obj uintptr) string {
    ret, _, _ := paraAttributes_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ParaAttributes_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := paraAttributes_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ParaAttributes_GetHashCode(obj uintptr) int32 {
    ret, _, _ := paraAttributes_GetHashCode.Call(obj)
    return int32(ret)
}

func ParaAttributes_ToString(obj uintptr) string {
    ret, _, _ := paraAttributes_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ParaAttributes_GetAlignment(obj uintptr) TAlignment {
    ret, _, _ := paraAttributes_GetAlignment.Call(obj)
    return TAlignment(ret)
}

func ParaAttributes_SetAlignment(obj uintptr, value TAlignment) {
   paraAttributes_SetAlignment.Call(obj, uintptr(value))
}

func ParaAttributes_GetFirstIndent(obj uintptr) int32 {
    ret, _, _ := paraAttributes_GetFirstIndent.Call(obj)
    return int32(ret)
}

func ParaAttributes_SetFirstIndent(obj uintptr, value int32) {
   paraAttributes_SetFirstIndent.Call(obj, uintptr(value))
}

func ParaAttributes_GetLeftIndent(obj uintptr) int32 {
    ret, _, _ := paraAttributes_GetLeftIndent.Call(obj)
    return int32(ret)
}

func ParaAttributes_SetLeftIndent(obj uintptr, value int32) {
   paraAttributes_SetLeftIndent.Call(obj, uintptr(value))
}

func ParaAttributes_GetNumbering(obj uintptr) TNumberingStyle {
    ret, _, _ := paraAttributes_GetNumbering.Call(obj)
    return TNumberingStyle(ret)
}

func ParaAttributes_SetNumbering(obj uintptr, value TNumberingStyle) {
   paraAttributes_SetNumbering.Call(obj, uintptr(value))
}

func ParaAttributes_GetRightIndent(obj uintptr) int32 {
    ret, _, _ := paraAttributes_GetRightIndent.Call(obj)
    return int32(ret)
}

func ParaAttributes_SetRightIndent(obj uintptr, value int32) {
   paraAttributes_SetRightIndent.Call(obj, uintptr(value))
}

func ParaAttributes_GetTabCount(obj uintptr) int32 {
    ret, _, _ := paraAttributes_GetTabCount.Call(obj)
    return int32(ret)
}

func ParaAttributes_SetTabCount(obj uintptr, value int32) {
   paraAttributes_SetTabCount.Call(obj, uintptr(value))
}

func ParaAttributes_GetTab(obj uintptr, Index uint8) int32 {
    ret, _, _ := paraAttributes_GetTab.Call(obj, uintptr(Index))
    return int32(ret)
}

func ParaAttributes_SetTab(obj uintptr, Index uint8, value int32) {
   paraAttributes_SetTab.Call(obj, uintptr(Index), uintptr(value))
}


//--------------------------- TTextAttributes ---------------------------

func TextAttributes_Assign(obj uintptr, Source uintptr)  {
    textAttributes_Assign.Call(obj, Source )
}

func TextAttributes_GetNamePath(obj uintptr) string {
    ret, _, _ := textAttributes_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func TextAttributes_ClassName(obj uintptr) string {
    ret, _, _ := textAttributes_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func TextAttributes_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := textAttributes_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func TextAttributes_GetHashCode(obj uintptr) int32 {
    ret, _, _ := textAttributes_GetHashCode.Call(obj)
    return int32(ret)
}

func TextAttributes_ToString(obj uintptr) string {
    ret, _, _ := textAttributes_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func TextAttributes_GetCharset(obj uintptr) TFontCharset {
    ret, _, _ := textAttributes_GetCharset.Call(obj)
    return TFontCharset(ret)
}

func TextAttributes_SetCharset(obj uintptr, value TFontCharset) {
   textAttributes_SetCharset.Call(obj, uintptr(value))
}

func TextAttributes_GetColor(obj uintptr) TColor {
    ret, _, _ := textAttributes_GetColor.Call(obj)
    return TColor(ret)
}

func TextAttributes_SetColor(obj uintptr, value TColor) {
   textAttributes_SetColor.Call(obj, uintptr(value))
}

func TextAttributes_GetConsistentAttributes(obj uintptr) TConsistentAttributes {
    ret, _, _ := textAttributes_GetConsistentAttributes.Call(obj)
    return TConsistentAttributes(ret)
}

func TextAttributes_GetName(obj uintptr) string {
    ret, _, _ := textAttributes_GetName.Call(obj)
    return DStrToGoStr(ret)
}

func TextAttributes_SetName(obj uintptr, value string) {
   textAttributes_SetName.Call(obj, GoStrToDStr(value))
}

func TextAttributes_GetPitch(obj uintptr) TFontPitch {
    ret, _, _ := textAttributes_GetPitch.Call(obj)
    return TFontPitch(ret)
}

func TextAttributes_SetPitch(obj uintptr, value TFontPitch) {
   textAttributes_SetPitch.Call(obj, uintptr(value))
}

func TextAttributes_GetProtected(obj uintptr) bool {
    ret, _, _ := textAttributes_GetProtected.Call(obj)
    return DBoolToGoBool(ret)
}

func TextAttributes_SetProtected(obj uintptr, value bool) {
   textAttributes_SetProtected.Call(obj, GoBoolToDBool(value))
}

func TextAttributes_GetSize(obj uintptr) int32 {
    ret, _, _ := textAttributes_GetSize.Call(obj)
    return int32(ret)
}

func TextAttributes_SetSize(obj uintptr, value int32) {
   textAttributes_SetSize.Call(obj, uintptr(value))
}

func TextAttributes_GetStyle(obj uintptr) TFontStyles {
    ret, _, _ := textAttributes_GetStyle.Call(obj)
    return TFontStyles(ret)
}

func TextAttributes_SetStyle(obj uintptr, value TFontStyles) {
   textAttributes_SetStyle.Call(obj, uintptr(value))
}

func TextAttributes_GetHeight(obj uintptr) int32 {
    ret, _, _ := textAttributes_GetHeight.Call(obj)
    return int32(ret)
}

func TextAttributes_SetHeight(obj uintptr, value int32) {
   textAttributes_SetHeight.Call(obj, uintptr(value))
}


//--------------------------- TIconOptions ---------------------------

func IconOptions_Assign(obj uintptr, Source uintptr)  {
    iconOptions_Assign.Call(obj, Source )
}

func IconOptions_GetNamePath(obj uintptr) string {
    ret, _, _ := iconOptions_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func IconOptions_ClassName(obj uintptr) string {
    ret, _, _ := iconOptions_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func IconOptions_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := iconOptions_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func IconOptions_GetHashCode(obj uintptr) int32 {
    ret, _, _ := iconOptions_GetHashCode.Call(obj)
    return int32(ret)
}

func IconOptions_ToString(obj uintptr) string {
    ret, _, _ := iconOptions_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func IconOptions_GetArrangement(obj uintptr) TIconArrangement {
    ret, _, _ := iconOptions_GetArrangement.Call(obj)
    return TIconArrangement(ret)
}

func IconOptions_SetArrangement(obj uintptr, value TIconArrangement) {
   iconOptions_SetArrangement.Call(obj, uintptr(value))
}

func IconOptions_GetAutoArrange(obj uintptr) bool {
    ret, _, _ := iconOptions_GetAutoArrange.Call(obj)
    return DBoolToGoBool(ret)
}

func IconOptions_SetAutoArrange(obj uintptr, value bool) {
   iconOptions_SetAutoArrange.Call(obj, GoBoolToDBool(value))
}


//--------------------------- Exception ---------------------------

func Exception_ToString(obj uintptr) string {
    ret, _, _ := exception_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Exception_ClassName(obj uintptr) string {
    ret, _, _ := exception_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Exception_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := exception_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Exception_GetHashCode(obj uintptr) int32 {
    ret, _, _ := exception_GetHashCode.Call(obj)
    return int32(ret)
}

func Exception_GetBaseException(obj uintptr) uintptr {
    ret, _, _ := exception_GetBaseException.Call(obj)
    return ret
}

func Exception_GetInnerException(obj uintptr) uintptr {
    ret, _, _ := exception_GetInnerException.Call(obj)
    return ret
}

func Exception_GetMessage(obj uintptr) string {
    ret, _, _ := exception_GetMessage.Call(obj)
    return DStrToGoStr(ret)
}

func Exception_SetMessage(obj uintptr, value string) {
   exception_SetMessage.Call(obj, GoStrToDStr(value))
}

func Exception_GetStackTrace(obj uintptr) string {
    ret, _, _ := exception_GetStackTrace.Call(obj)
    return DStrToGoStr(ret)
}

func Exception_GetStackInfo(obj uintptr) uintptr {
    ret, _, _ := exception_GetStackInfo.Call(obj)
    return ret
}

