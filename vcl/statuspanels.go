
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TStatusPanels struct {
    IObject
    instance uintptr
}

func NewStatusPanels() *TStatusPanels {
    s := new(TStatusPanels)
    s.instance = StatusPanels_Create()
    return s
}

func StatusPanelsFromInst(inst uintptr) *TStatusPanels {
    s := new(TStatusPanels)
    s.instance = inst
    return s
}

func StatusPanelsFromObj(obj IObject) *TStatusPanels {
    s := new(TStatusPanels)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TStatusPanels) Free() {
    if s.instance != 0 {
        StatusPanels_Free(s.instance)
    }
}

func (s *TStatusPanels) Instance() uintptr {
    return s.instance
}

func (s *TStatusPanels) IsValid() bool {
    return s.instance != 0
}

func (s *TStatusPanels) Add() *TStatusPanel {
    defer exceptionProc()
    return StatusPanelFromInst(StatusPanels_Add(s.instance))
}

func (s *TStatusPanels) AddItem(Item *TStatusPanel, Index int32) *TStatusPanel {
    defer exceptionProc()
    return StatusPanelFromInst(StatusPanels_AddItem(s.instance, CheckPtr(Item), Index ))
}

func (s *TStatusPanels) Insert(Index int32) *TStatusPanel {
    defer exceptionProc()
    return StatusPanelFromInst(StatusPanels_Insert(s.instance, Index ))
}

func (s *TStatusPanels) Owner() *TObject {
    defer exceptionProc()
    return ObjectFromInst(StatusPanels_Owner(s.instance))
}

func (s *TStatusPanels) Assign(Source IObject) {
    defer exceptionProc()
    StatusPanels_Assign(s.instance, CheckPtr(Source))
}

func (s *TStatusPanels) BeginUpdate() {
    defer exceptionProc()
    StatusPanels_BeginUpdate(s.instance)
}

func (s *TStatusPanels) Clear() {
    defer exceptionProc()
    StatusPanels_Clear(s.instance)
}

func (s *TStatusPanels) Delete(Index int32) {
    defer exceptionProc()
    StatusPanels_Delete(s.instance, Index )
}

func (s *TStatusPanels) EndUpdate() {
    defer exceptionProc()
    StatusPanels_EndUpdate(s.instance)
}

func (s *TStatusPanels) GetNamePath() string {
    defer exceptionProc()
    return StatusPanels_GetNamePath(s.instance)
}

func (s *TStatusPanels) ClassName() string {
    defer exceptionProc()
    return StatusPanels_ClassName(s.instance)
}

func (s *TStatusPanels) Equals(Obj IObject) bool {
    defer exceptionProc()
    return StatusPanels_Equals(s.instance, CheckPtr(Obj))
}

func (s *TStatusPanels) GetHashCode() int32 {
    defer exceptionProc()
    return StatusPanels_GetHashCode(s.instance)
}

func (s *TStatusPanels) ToString() string {
    defer exceptionProc()
    return StatusPanels_ToString(s.instance)
}

func (s *TStatusPanels) Items(Index int32) *TStatusPanel {
    defer exceptionProc()
    return StatusPanelFromInst(StatusPanels_GetItems(s.instance, Index))
}

func (s *TStatusPanels) SetItems(Index int32, value *TStatusPanel) {
    defer exceptionProc()
    StatusPanels_SetItems(s.instance, Index, CheckPtr(value))
}

