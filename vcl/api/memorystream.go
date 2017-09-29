
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

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
    ret, _, _ := memoryStream_Seek.Call(obj, uintptr(Offset) , uintptr(Origin) )
    return int64(ret)
}

func MemoryStream_SaveToStream(obj uintptr, Stream uintptr)  {
    memoryStream_SaveToStream.Call(obj, Stream )
}

func MemoryStream_SaveToFile(obj uintptr, FileName string)  {
    memoryStream_SaveToFile.Call(obj, GoStrToDStr(FileName) )
}

func MemoryStream_CopyFrom(obj uintptr, Source uintptr, Count int64) int64 {
    ret, _, _ := memoryStream_CopyFrom.Call(obj, Source , uintptr(Count) )
    return int64(ret)
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
    ret, _, _ := memoryStream_GetPosition.Call(obj)
    return int64(ret)
}

func MemoryStream_SetPosition(obj uintptr, value int64) {
   memoryStream_SetPosition.Call(obj, uintptr(value))
}

func MemoryStream_GetSize(obj uintptr) int64 {
    ret, _, _ := memoryStream_GetSize.Call(obj)
    return int64(ret)
}

func MemoryStream_SetSize(obj uintptr, value int64) {
   memoryStream_SetSize.Call(obj, uintptr(value))
}

