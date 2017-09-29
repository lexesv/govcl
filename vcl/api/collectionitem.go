
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

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

