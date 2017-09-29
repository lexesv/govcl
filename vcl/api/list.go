
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

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

