
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

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

