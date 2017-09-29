
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TStringList struct {
    IObject
    instance uintptr
}

func NewStringList() *TStringList {
    s := new(TStringList)
    s.instance = StringList_Create()
    return s
}

func StringListFromInst(inst uintptr) *TStringList {
    s := new(TStringList)
    s.instance = inst
    return s
}

func StringListFromObj(obj IObject) *TStringList {
    s := new(TStringList)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TStringList) Free() {
    if s.instance != 0 {
        StringList_Free(s.instance)
    }
}

func (s *TStringList) Instance() uintptr {
    return s.instance
}

func (s *TStringList) IsValid() bool {
    return s.instance != 0
}

func (s *TStringList) Add(S string) int32 {
    defer exceptionProc()
    return StringList_Add(s.instance, S )
}

func (s *TStringList) AddObject(S string, AObject IObject) int32 {
    defer exceptionProc()
    return StringList_AddObject(s.instance, S , CheckPtr(AObject))
}

func (s *TStringList) Assign(Source IObject) {
    defer exceptionProc()
    StringList_Assign(s.instance, CheckPtr(Source))
}

func (s *TStringList) Clear() {
    defer exceptionProc()
    StringList_Clear(s.instance)
}

func (s *TStringList) Delete(Index int32) {
    defer exceptionProc()
    StringList_Delete(s.instance, Index )
}

func (s *TStringList) IndexOf(S string) int32 {
    defer exceptionProc()
    return StringList_IndexOf(s.instance, S )
}

func (s *TStringList) Insert(Index int32, S string) {
    defer exceptionProc()
    StringList_Insert(s.instance, Index , S )
}

func (s *TStringList) InsertObject(Index int32, S string, AObject IObject) {
    defer exceptionProc()
    StringList_InsertObject(s.instance, Index , S , CheckPtr(AObject))
}

func (s *TStringList) Append(S string) {
    defer exceptionProc()
    StringList_Append(s.instance, S )
}

func (s *TStringList) BeginUpdate() {
    defer exceptionProc()
    StringList_BeginUpdate(s.instance)
}

func (s *TStringList) EndUpdate() {
    defer exceptionProc()
    StringList_EndUpdate(s.instance)
}

func (s *TStringList) Equals(Strings IObject) bool {
    defer exceptionProc()
    return StringList_Equals(s.instance, CheckPtr(Strings))
}

func (s *TStringList) IndexOfName(Name string) int32 {
    defer exceptionProc()
    return StringList_IndexOfName(s.instance, Name )
}

func (s *TStringList) IndexOfObject(AObject IObject) int32 {
    defer exceptionProc()
    return StringList_IndexOfObject(s.instance, CheckPtr(AObject))
}

func (s *TStringList) LoadFromFile(FileName string) {
    defer exceptionProc()
    StringList_LoadFromFile(s.instance, FileName )
}

func (s *TStringList) LoadFromStream(Stream IObject) {
    defer exceptionProc()
    StringList_LoadFromStream(s.instance, CheckPtr(Stream))
}

func (s *TStringList) Move(CurIndex int32, NewIndex int32) {
    defer exceptionProc()
    StringList_Move(s.instance, CurIndex , NewIndex )
}

func (s *TStringList) SaveToFile(FileName string) {
    defer exceptionProc()
    StringList_SaveToFile(s.instance, FileName )
}

func (s *TStringList) SaveToStream(Stream IObject) {
    defer exceptionProc()
    StringList_SaveToStream(s.instance, CheckPtr(Stream))
}

func (s *TStringList) GetNamePath() string {
    defer exceptionProc()
    return StringList_GetNamePath(s.instance)
}

func (s *TStringList) ClassName() string {
    defer exceptionProc()
    return StringList_ClassName(s.instance)
}

func (s *TStringList) GetHashCode() int32 {
    defer exceptionProc()
    return StringList_GetHashCode(s.instance)
}

func (s *TStringList) ToString() string {
    defer exceptionProc()
    return StringList_ToString(s.instance)
}

func (s *TStringList) Sorted() bool {
    defer exceptionProc()
    return StringList_GetSorted(s.instance)
}

func (s *TStringList) SetSorted(value bool) {
    defer exceptionProc()
    StringList_SetSorted(s.instance, value)
}

func (s *TStringList) SetOnChange(fn TNotifyEvent) {
    StringList_SetOnChange(s.instance, fn)
}

func (s *TStringList) CommaText() string {
    defer exceptionProc()
    return StringList_GetCommaText(s.instance)
}

func (s *TStringList) SetCommaText(value string) {
    defer exceptionProc()
    StringList_SetCommaText(s.instance, value)
}

func (s *TStringList) Delimiter() uint16 {
    defer exceptionProc()
    return StringList_GetDelimiter(s.instance)
}

func (s *TStringList) SetDelimiter(value uint16) {
    defer exceptionProc()
    StringList_SetDelimiter(s.instance, value)
}

func (s *TStringList) Text() string {
    defer exceptionProc()
    return StringList_GetText(s.instance)
}

func (s *TStringList) SetText(value string) {
    defer exceptionProc()
    StringList_SetText(s.instance, value)
}

func (s *TStringList) WriteBOM() bool {
    defer exceptionProc()
    return StringList_GetWriteBOM(s.instance)
}

func (s *TStringList) SetWriteBOM(value bool) {
    defer exceptionProc()
    StringList_SetWriteBOM(s.instance, value)
}

func (s *TStringList) Options() TStringsOptions {
    defer exceptionProc()
    return StringList_GetOptions(s.instance)
}

func (s *TStringList) SetOptions(value TStringsOptions) {
    defer exceptionProc()
    StringList_SetOptions(s.instance, value)
}

func (s *TStringList) Values(Name string) string {
    defer exceptionProc()
    return StringList_GetValues(s.instance, Name)
}

func (s *TStringList) SetValues(Name string, value string) {
    defer exceptionProc()
    StringList_SetValues(s.instance, Name, value)
}

func (s *TStringList) ValueFromIndex(Index int32) string {
    defer exceptionProc()
    return StringList_GetValueFromIndex(s.instance, Index)
}

func (s *TStringList) SetValueFromIndex(Index int32, value string) {
    defer exceptionProc()
    StringList_SetValueFromIndex(s.instance, Index, value)
}

func (s *TStringList) Strings(Index int32) string {
    defer exceptionProc()
    return StringList_GetStrings(s.instance, Index)
}

func (s *TStringList) SetStrings(Index int32, value string) {
    defer exceptionProc()
    StringList_SetStrings(s.instance, Index, value)
}

