
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func ImageList_Create(obj uintptr) uintptr {
    ret, _, _ := imageList_Create.Call(obj)
    return ret
}

func ImageList_Free(obj uintptr) {
    imageList_Free.Call(obj)
}

func ImageList_Assign(obj uintptr, Source uintptr)  {
    imageList_Assign.Call(obj, Source )
}

func ImageList_Add(obj uintptr, Image uintptr, Mask uintptr) int32 {
    ret, _, _ := imageList_Add.Call(obj, Image , Mask )
    return int32(ret)
}

func ImageList_Clear(obj uintptr)  {
    imageList_Clear.Call(obj)
}

func ImageList_Delete(obj uintptr, Index int32)  {
    imageList_Delete.Call(obj, uintptr(Index) )
}

func ImageList_HandleAllocated(obj uintptr) bool {
    ret, _, _ := imageList_HandleAllocated.Call(obj)
    return DBoolToGoBool(ret)
}

func ImageList_Insert(obj uintptr, Index int32, Image uintptr, Mask uintptr)  {
    imageList_Insert.Call(obj, uintptr(Index) , Image , Mask )
}

func ImageList_Move(obj uintptr, CurIndex int32, NewIndex int32)  {
    imageList_Move.Call(obj, uintptr(CurIndex) , uintptr(NewIndex) )
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

