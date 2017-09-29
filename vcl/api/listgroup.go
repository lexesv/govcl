
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func ListGroup_Create() uintptr {
    ret, _, _ := listGroup_Create.Call()
    return ret
}

func ListGroup_Free(obj uintptr) {
    listGroup_Free.Call(obj)
}

func ListGroup_Assign(obj uintptr, Source uintptr)  {
    listGroup_Assign.Call(obj, Source )
}

func ListGroup_GetNamePath(obj uintptr) string {
    ret, _, _ := listGroup_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func ListGroup_ClassName(obj uintptr) string {
    ret, _, _ := listGroup_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func ListGroup_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := listGroup_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func ListGroup_GetHashCode(obj uintptr) int32 {
    ret, _, _ := listGroup_GetHashCode.Call(obj)
    return int32(ret)
}

func ListGroup_ToString(obj uintptr) string {
    ret, _, _ := listGroup_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func ListGroup_GetGroupID(obj uintptr) int32 {
    ret, _, _ := listGroup_GetGroupID.Call(obj)
    return int32(ret)
}

func ListGroup_SetGroupID(obj uintptr, value int32) {
   listGroup_SetGroupID.Call(obj, uintptr(value))
}

func ListGroup_GetState(obj uintptr) TListGroupStateSet {
    ret, _, _ := listGroup_GetState.Call(obj)
    return TListGroupStateSet(ret)
}

func ListGroup_SetState(obj uintptr, value TListGroupStateSet) {
   listGroup_SetState.Call(obj, uintptr(value))
}

func ListGroup_GetIndex(obj uintptr) int32 {
    ret, _, _ := listGroup_GetIndex.Call(obj)
    return int32(ret)
}

func ListGroup_SetIndex(obj uintptr, value int32) {
   listGroup_SetIndex.Call(obj, uintptr(value))
}

