
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

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

