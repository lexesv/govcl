
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

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

