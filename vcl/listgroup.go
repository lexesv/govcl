
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
    . "gitee.com/ying32/govcl/vcl/types"
)

type TListGroup struct {
    IObject
    instance uintptr
}

func NewListGroup() *TListGroup {
    l := new(TListGroup)
    l.instance = ListGroup_Create()
    return l
}

func ListGroupFromInst(inst uintptr) *TListGroup {
    l := new(TListGroup)
    l.instance = inst
    return l
}

func ListGroupFromObj(obj IObject) *TListGroup {
    l := new(TListGroup)
    l.instance = CheckPtr(obj)
    return l
}

func (l *TListGroup) Free() {
    if l.instance != 0 {
        ListGroup_Free(l.instance)
    }
}

func (l *TListGroup) Instance() uintptr {
    return l.instance
}

func (l *TListGroup) IsValid() bool {
    return l.instance != 0
}

func (l *TListGroup) Assign(Source IObject) {
    ListGroup_Assign(l.instance, CheckPtr(Source))
}

func (l *TListGroup) GetNamePath() string {
    return ListGroup_GetNamePath(l.instance)
}

func (l *TListGroup) ClassName() string {
    return ListGroup_ClassName(l.instance)
}

func (l *TListGroup) Equals(Obj IObject) bool {
    return ListGroup_Equals(l.instance, CheckPtr(Obj))
}

func (l *TListGroup) GetHashCode() int32 {
    return ListGroup_GetHashCode(l.instance)
}

func (l *TListGroup) ToString() string {
    return ListGroup_ToString(l.instance)
}

func (l *TListGroup) GroupID() int32 {
    return ListGroup_GetGroupID(l.instance)
}

func (l *TListGroup) SetGroupID(value int32) {
    ListGroup_SetGroupID(l.instance, value)
}

func (l *TListGroup) State() TListGroupStateSet {
    return ListGroup_GetState(l.instance)
}

func (l *TListGroup) SetState(value TListGroupStateSet) {
    ListGroup_SetState(l.instance, value)
}

func (l *TListGroup) Index() int32 {
    return ListGroup_GetIndex(l.instance)
}

func (l *TListGroup) SetIndex(value int32) {
    ListGroup_SetIndex(l.instance, value)
}

