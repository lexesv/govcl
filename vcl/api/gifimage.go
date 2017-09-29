
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

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

