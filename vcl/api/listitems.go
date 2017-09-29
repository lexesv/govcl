
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

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

