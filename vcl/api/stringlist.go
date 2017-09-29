
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func StringList_Create() uintptr {
    ret, _, _ := stringList_Create.Call()
    return ret
}

func StringList_Free(obj uintptr) {
    stringList_Free.Call(obj)
}

func StringList_Add(obj uintptr, S string) int32 {
    ret, _, _ := stringList_Add.Call(obj, GoStrToDStr(S) )
    return int32(ret)
}

func StringList_AddObject(obj uintptr, S string, AObject uintptr) int32 {
    ret, _, _ := stringList_AddObject.Call(obj, GoStrToDStr(S) , AObject )
    return int32(ret)
}

func StringList_Assign(obj uintptr, Source uintptr)  {
    stringList_Assign.Call(obj, Source )
}

func StringList_Clear(obj uintptr)  {
    stringList_Clear.Call(obj)
}

func StringList_Delete(obj uintptr, Index int32)  {
    stringList_Delete.Call(obj, uintptr(Index) )
}

func StringList_IndexOf(obj uintptr, S string) int32 {
    ret, _, _ := stringList_IndexOf.Call(obj, GoStrToDStr(S) )
    return int32(ret)
}

func StringList_Insert(obj uintptr, Index int32, S string)  {
    stringList_Insert.Call(obj, uintptr(Index) , GoStrToDStr(S) )
}

func StringList_InsertObject(obj uintptr, Index int32, S string, AObject uintptr)  {
    stringList_InsertObject.Call(obj, uintptr(Index) , GoStrToDStr(S) , AObject )
}

func StringList_Append(obj uintptr, S string)  {
    stringList_Append.Call(obj, GoStrToDStr(S) )
}

func StringList_BeginUpdate(obj uintptr)  {
    stringList_BeginUpdate.Call(obj)
}

func StringList_EndUpdate(obj uintptr)  {
    stringList_EndUpdate.Call(obj)
}

func StringList_Equals(obj uintptr, Strings uintptr) bool {
    ret, _, _ := stringList_Equals.Call(obj, Strings )
    return DBoolToGoBool(ret)
}

func StringList_IndexOfName(obj uintptr, Name string) int32 {
    ret, _, _ := stringList_IndexOfName.Call(obj, GoStrToDStr(Name) )
    return int32(ret)
}

func StringList_IndexOfObject(obj uintptr, AObject uintptr) int32 {
    ret, _, _ := stringList_IndexOfObject.Call(obj, AObject )
    return int32(ret)
}

func StringList_LoadFromFile(obj uintptr, FileName string)  {
    stringList_LoadFromFile.Call(obj, GoStrToDStr(FileName) )
}

func StringList_LoadFromStream(obj uintptr, Stream uintptr)  {
    stringList_LoadFromStream.Call(obj, Stream )
}

func StringList_Move(obj uintptr, CurIndex int32, NewIndex int32)  {
    stringList_Move.Call(obj, uintptr(CurIndex) , uintptr(NewIndex) )
}

func StringList_SaveToFile(obj uintptr, FileName string)  {
    stringList_SaveToFile.Call(obj, GoStrToDStr(FileName) )
}

func StringList_SaveToStream(obj uintptr, Stream uintptr)  {
    stringList_SaveToStream.Call(obj, Stream )
}

func StringList_GetNamePath(obj uintptr) string {
    ret, _, _ := stringList_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func StringList_ClassName(obj uintptr) string {
    ret, _, _ := stringList_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func StringList_GetHashCode(obj uintptr) int32 {
    ret, _, _ := stringList_GetHashCode.Call(obj)
    return int32(ret)
}

func StringList_ToString(obj uintptr) string {
    ret, _, _ := stringList_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func StringList_GetSorted(obj uintptr) bool {
    ret, _, _ := stringList_GetSorted.Call(obj)
    return DBoolToGoBool(ret)
}

func StringList_SetSorted(obj uintptr, value bool) {
   stringList_SetSorted.Call(obj, GoBoolToDBool(value))
}

func StringList_SetOnChange(obj uintptr, fn interface{}) {
    stringList_SetOnChange.Call(obj, addEventToMap(fn))
}

func StringList_GetCommaText(obj uintptr) string {
    ret, _, _ := stringList_GetCommaText.Call(obj)
    return DStrToGoStr(ret)
}

func StringList_SetCommaText(obj uintptr, value string) {
   stringList_SetCommaText.Call(obj, GoStrToDStr(value))
}

func StringList_GetDelimiter(obj uintptr) uint16 {
    ret, _, _ := stringList_GetDelimiter.Call(obj)
    return uint16(ret)
}

func StringList_SetDelimiter(obj uintptr, value uint16) {
   stringList_SetDelimiter.Call(obj, uintptr(value))
}

func StringList_GetText(obj uintptr) string {
    ret, _, _ := stringList_GetText.Call(obj)
    return DStrToGoStr(ret)
}

func StringList_SetText(obj uintptr, value string) {
   stringList_SetText.Call(obj, GoStrToDStr(value))
}

func StringList_GetWriteBOM(obj uintptr) bool {
    ret, _, _ := stringList_GetWriteBOM.Call(obj)
    return DBoolToGoBool(ret)
}

func StringList_SetWriteBOM(obj uintptr, value bool) {
   stringList_SetWriteBOM.Call(obj, GoBoolToDBool(value))
}

func StringList_GetOptions(obj uintptr) TStringsOptions {
    ret, _, _ := stringList_GetOptions.Call(obj)
    return TStringsOptions(ret)
}

func StringList_SetOptions(obj uintptr, value TStringsOptions) {
   stringList_SetOptions.Call(obj, uintptr(value))
}

func StringList_GetValues(obj uintptr, Name string) string {
    ret, _, _ := stringList_GetValues.Call(obj, GoStrToDStr(Name))
    return DStrToGoStr(ret)
}

func StringList_SetValues(obj uintptr, Name string, value string) {
   stringList_SetValues.Call(obj, GoStrToDStr(Name), GoStrToDStr(value))
}

func StringList_GetValueFromIndex(obj uintptr, Index int32) string {
    ret, _, _ := stringList_GetValueFromIndex.Call(obj, uintptr(Index))
    return DStrToGoStr(ret)
}

func StringList_SetValueFromIndex(obj uintptr, Index int32, value string) {
   stringList_SetValueFromIndex.Call(obj, uintptr(Index), GoStrToDStr(value))
}

func StringList_GetStrings(obj uintptr, Index int32) string {
    ret, _, _ := stringList_GetStrings.Call(obj, uintptr(Index))
    return DStrToGoStr(ret)
}

func StringList_SetStrings(obj uintptr, Index int32, value string) {
   stringList_SetStrings.Call(obj, uintptr(Index), GoStrToDStr(value))
}

