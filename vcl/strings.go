
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
)

type TStrings struct {
    IObject
    instance uintptr
}

func NewStrings() *TStrings {
    s := new(TStrings)
    s.instance = Strings_Create()
    return s
}

func StringsFromInst(inst uintptr) *TStrings {
    s := new(TStrings)
    s.instance = inst
    return s
}

func StringsFromObj(obj IObject) *TStrings {
    s := new(TStrings)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TStrings) Free() {
    if s.instance != 0 {
        Strings_Free(s.instance)
    }
}

func (s *TStrings) Instance() uintptr {
    return s.instance
}

func (s *TStrings) IsValid() bool {
    return s.instance != 0
}

func (s *TStrings) Add(S string) int32 {
    defer exceptionProc()
    return Strings_Add(s.instance, S )
}

func (s *TStrings) AddObject(S string, AObject IObject) int32 {
    defer exceptionProc()
    return Strings_AddObject(s.instance, S , CheckPtr(AObject))
}

func (s *TStrings) Append(S string) {
    defer exceptionProc()
    Strings_Append(s.instance, S )
}

func (s *TStrings) Assign(Source IObject) {
    defer exceptionProc()
    Strings_Assign(s.instance, CheckPtr(Source))
}

func (s *TStrings) BeginUpdate() {
    defer exceptionProc()
    Strings_BeginUpdate(s.instance)
}

func (s *TStrings) Clear() {
    defer exceptionProc()
    Strings_Clear(s.instance)
}

func (s *TStrings) Delete(Index int32) {
    defer exceptionProc()
    Strings_Delete(s.instance, Index )
}

func (s *TStrings) EndUpdate() {
    defer exceptionProc()
    Strings_EndUpdate(s.instance)
}

func (s *TStrings) Equals(Strings IObject) bool {
    defer exceptionProc()
    return Strings_Equals(s.instance, CheckPtr(Strings))
}

func (s *TStrings) IndexOf(S string) int32 {
    defer exceptionProc()
    return Strings_IndexOf(s.instance, S )
}

func (s *TStrings) IndexOfName(Name string) int32 {
    defer exceptionProc()
    return Strings_IndexOfName(s.instance, Name )
}

func (s *TStrings) IndexOfObject(AObject IObject) int32 {
    defer exceptionProc()
    return Strings_IndexOfObject(s.instance, CheckPtr(AObject))
}

func (s *TStrings) Insert(Index int32, S string) {
    defer exceptionProc()
    Strings_Insert(s.instance, Index , S )
}

func (s *TStrings) InsertObject(Index int32, S string, AObject IObject) {
    defer exceptionProc()
    Strings_InsertObject(s.instance, Index , S , CheckPtr(AObject))
}

func (s *TStrings) LoadFromFile(FileName string) {
    defer exceptionProc()
    Strings_LoadFromFile(s.instance, FileName )
}

func (s *TStrings) LoadFromStream(Stream IObject) {
    defer exceptionProc()
    Strings_LoadFromStream(s.instance, CheckPtr(Stream))
}

func (s *TStrings) Move(CurIndex int32, NewIndex int32) {
    defer exceptionProc()
    Strings_Move(s.instance, CurIndex , NewIndex )
}

func (s *TStrings) SaveToFile(FileName string) {
    defer exceptionProc()
    Strings_SaveToFile(s.instance, FileName )
}

func (s *TStrings) SaveToStream(Stream IObject) {
    defer exceptionProc()
    Strings_SaveToStream(s.instance, CheckPtr(Stream))
}

func (s *TStrings) GetNamePath() string {
    defer exceptionProc()
    return Strings_GetNamePath(s.instance)
}

func (s *TStrings) ClassName() string {
    defer exceptionProc()
    return Strings_ClassName(s.instance)
}

func (s *TStrings) GetHashCode() int32 {
    defer exceptionProc()
    return Strings_GetHashCode(s.instance)
}

func (s *TStrings) ToString() string {
    defer exceptionProc()
    return Strings_ToString(s.instance)
}

func (s *TStrings) CommaText() string {
    defer exceptionProc()
    return Strings_GetCommaText(s.instance)
}

func (s *TStrings) SetCommaText(value string) {
    defer exceptionProc()
    Strings_SetCommaText(s.instance, value)
}

func (s *TStrings) Delimiter() uint16 {
    defer exceptionProc()
    return Strings_GetDelimiter(s.instance)
}

func (s *TStrings) SetDelimiter(value uint16) {
    defer exceptionProc()
    Strings_SetDelimiter(s.instance, value)
}

func (s *TStrings) Text() string {
    defer exceptionProc()
    return Strings_GetText(s.instance)
}

func (s *TStrings) SetText(value string) {
    defer exceptionProc()
    Strings_SetText(s.instance, value)
}

func (s *TStrings) WriteBOM() bool {
    defer exceptionProc()
    return Strings_GetWriteBOM(s.instance)
}

func (s *TStrings) SetWriteBOM(value bool) {
    defer exceptionProc()
    Strings_SetWriteBOM(s.instance, value)
}

func (s *TStrings) Options() TStringsOptions {
    defer exceptionProc()
    return Strings_GetOptions(s.instance)
}

func (s *TStrings) SetOptions(value TStringsOptions) {
    defer exceptionProc()
    Strings_SetOptions(s.instance, value)
}

func (s *TStrings) Values(Name string) string {
    defer exceptionProc()
    return Strings_GetValues(s.instance, Name)
}

func (s *TStrings) SetValues(Name string, value string) {
    defer exceptionProc()
    Strings_SetValues(s.instance, Name, value)
}

func (s *TStrings) ValueFromIndex(Index int32) string {
    defer exceptionProc()
    return Strings_GetValueFromIndex(s.instance, Index)
}

func (s *TStrings) SetValueFromIndex(Index int32, value string) {
    defer exceptionProc()
    Strings_SetValueFromIndex(s.instance, Index, value)
}

func (s *TStrings) Strings(Index int32) string {
    defer exceptionProc()
    return Strings_GetStrings(s.instance, Index)
}

func (s *TStrings) SetStrings(Index int32, value string) {
    defer exceptionProc()
    Strings_SetStrings(s.instance, Index, value)
}

