
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

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

