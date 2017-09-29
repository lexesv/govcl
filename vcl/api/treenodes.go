
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func TreeNodes_Create() uintptr {
    ret, _, _ := treeNodes_Create.Call()
    return ret
}

func TreeNodes_Free(obj uintptr) {
    treeNodes_Free.Call(obj)
}

func TreeNodes_AddChildFirst(obj uintptr, Parent uintptr, S string) uintptr {
    ret, _, _ := treeNodes_AddChildFirst.Call(obj, Parent , GoStrToDStr(S) )
    return ret
}

func TreeNodes_AddChild(obj uintptr, Parent uintptr, S string) uintptr {
    ret, _, _ := treeNodes_AddChild.Call(obj, Parent , GoStrToDStr(S) )
    return ret
}

func TreeNodes_AddChildObjectFirst(obj uintptr, Parent uintptr, S string, Ptr uintptr) uintptr {
    ret, _, _ := treeNodes_AddChildObjectFirst.Call(obj, Parent , GoStrToDStr(S) , Ptr )
    return ret
}

func TreeNodes_AddChildObject(obj uintptr, Parent uintptr, S string, Ptr uintptr) uintptr {
    ret, _, _ := treeNodes_AddChildObject.Call(obj, Parent , GoStrToDStr(S) , Ptr )
    return ret
}

func TreeNodes_AddObjectFirst(obj uintptr, Sibling uintptr, S string, Ptr uintptr) uintptr {
    ret, _, _ := treeNodes_AddObjectFirst.Call(obj, Sibling , GoStrToDStr(S) , Ptr )
    return ret
}

func TreeNodes_AddObject(obj uintptr, Sibling uintptr, S string, Ptr uintptr) uintptr {
    ret, _, _ := treeNodes_AddObject.Call(obj, Sibling , GoStrToDStr(S) , Ptr )
    return ret
}

func TreeNodes_AddNode(obj uintptr, Node uintptr, Relative uintptr, S string, Ptr uintptr, Method TNodeAttachMode) uintptr {
    ret, _, _ := treeNodes_AddNode.Call(obj, Node , Relative , GoStrToDStr(S) , Ptr , uintptr(Method) )
    return ret
}

func TreeNodes_AddFirst(obj uintptr, Sibling uintptr, S string) uintptr {
    ret, _, _ := treeNodes_AddFirst.Call(obj, Sibling , GoStrToDStr(S) )
    return ret
}

func TreeNodes_Add(obj uintptr, Sibling uintptr, S string) uintptr {
    ret, _, _ := treeNodes_Add.Call(obj, Sibling , GoStrToDStr(S) )
    return ret
}

func TreeNodes_AlphaSort(obj uintptr, ARecurse bool) bool {
    ret, _, _ := treeNodes_AlphaSort.Call(obj, GoBoolToDBool(ARecurse) )
    return DBoolToGoBool(ret)
}

func TreeNodes_Assign(obj uintptr, Source uintptr)  {
    treeNodes_Assign.Call(obj, Source )
}

func TreeNodes_BeginUpdate(obj uintptr)  {
    treeNodes_BeginUpdate.Call(obj)
}

func TreeNodes_Clear(obj uintptr)  {
    treeNodes_Clear.Call(obj)
}

func TreeNodes_Delete(obj uintptr, Node uintptr)  {
    treeNodes_Delete.Call(obj, Node )
}

func TreeNodes_EndUpdate(obj uintptr)  {
    treeNodes_EndUpdate.Call(obj)
}

func TreeNodes_GetFirstNode(obj uintptr) uintptr {
    ret, _, _ := treeNodes_GetFirstNode.Call(obj)
    return ret
}

func TreeNodes_GetNode(obj uintptr, ItemId uintptr) uintptr {
    ret, _, _ := treeNodes_GetNode.Call(obj, ItemId )
    return ret
}

func TreeNodes_Insert(obj uintptr, Sibling uintptr, S string) uintptr {
    ret, _, _ := treeNodes_Insert.Call(obj, Sibling , GoStrToDStr(S) )
    return ret
}

func TreeNodes_InsertObject(obj uintptr, Sibling uintptr, S string, Ptr uintptr) uintptr {
    ret, _, _ := treeNodes_InsertObject.Call(obj, Sibling , GoStrToDStr(S) , Ptr )
    return ret
}

func TreeNodes_GetNamePath(obj uintptr) string {
    ret, _, _ := treeNodes_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func TreeNodes_ClassName(obj uintptr) string {
    ret, _, _ := treeNodes_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func TreeNodes_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := treeNodes_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func TreeNodes_GetHashCode(obj uintptr) int32 {
    ret, _, _ := treeNodes_GetHashCode.Call(obj)
    return int32(ret)
}

func TreeNodes_ToString(obj uintptr) string {
    ret, _, _ := treeNodes_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func TreeNodes_GetHandle(obj uintptr) HWND {
    ret, _, _ := treeNodes_GetHandle.Call(obj)
    return HWND(ret)
}

func TreeNodes_GetOwner(obj uintptr) uintptr {
    ret, _, _ := treeNodes_GetOwner.Call(obj)
    return ret
}

func TreeNodes_GetItem(obj uintptr, Index int32) uintptr {
    ret, _, _ := treeNodes_GetItem.Call(obj, uintptr(Index))
    return ret
}

