
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

func Strings_Create() uintptr {
    ret, _, _ := strings_Create.Call()
    return ret
}

func Strings_Free(obj uintptr) {
    strings_Free.Call(obj)
}

func Strings_Add(obj uintptr, S string) int32 {
    ret, _, _ := strings_Add.Call(obj, GoStrToDStr(S) )
    return int32(ret)
}

func Strings_AddObject(obj uintptr, S string, AObject uintptr) int32 {
    ret, _, _ := strings_AddObject.Call(obj, GoStrToDStr(S) , AObject )
    return int32(ret)
}

func Strings_Append(obj uintptr, S string)  {
    strings_Append.Call(obj, GoStrToDStr(S) )
}

func Strings_Assign(obj uintptr, Source uintptr)  {
    strings_Assign.Call(obj, Source )
}

func Strings_BeginUpdate(obj uintptr)  {
    strings_BeginUpdate.Call(obj)
}

func Strings_Clear(obj uintptr)  {
    strings_Clear.Call(obj)
}

func Strings_Delete(obj uintptr, Index int32)  {
    strings_Delete.Call(obj, uintptr(Index) )
}

func Strings_EndUpdate(obj uintptr)  {
    strings_EndUpdate.Call(obj)
}

func Strings_Equals(obj uintptr, Strings uintptr) bool {
    ret, _, _ := strings_Equals.Call(obj, Strings )
    return DBoolToGoBool(ret)
}

func Strings_IndexOf(obj uintptr, S string) int32 {
    ret, _, _ := strings_IndexOf.Call(obj, GoStrToDStr(S) )
    return int32(ret)
}

func Strings_IndexOfName(obj uintptr, Name string) int32 {
    ret, _, _ := strings_IndexOfName.Call(obj, GoStrToDStr(Name) )
    return int32(ret)
}

func Strings_IndexOfObject(obj uintptr, AObject uintptr) int32 {
    ret, _, _ := strings_IndexOfObject.Call(obj, AObject )
    return int32(ret)
}

func Strings_Insert(obj uintptr, Index int32, S string)  {
    strings_Insert.Call(obj, uintptr(Index) , GoStrToDStr(S) )
}

func Strings_InsertObject(obj uintptr, Index int32, S string, AObject uintptr)  {
    strings_InsertObject.Call(obj, uintptr(Index) , GoStrToDStr(S) , AObject )
}

func Strings_LoadFromFile(obj uintptr, FileName string)  {
    strings_LoadFromFile.Call(obj, GoStrToDStr(FileName) )
}

func Strings_LoadFromStream(obj uintptr, Stream uintptr)  {
    strings_LoadFromStream.Call(obj, Stream )
}

func Strings_Move(obj uintptr, CurIndex int32, NewIndex int32)  {
    strings_Move.Call(obj, uintptr(CurIndex) , uintptr(NewIndex) )
}

func Strings_SaveToFile(obj uintptr, FileName string)  {
    strings_SaveToFile.Call(obj, GoStrToDStr(FileName) )
}

func Strings_SaveToStream(obj uintptr, Stream uintptr)  {
    strings_SaveToStream.Call(obj, Stream )
}

func Strings_GetNamePath(obj uintptr) string {
    ret, _, _ := strings_GetNamePath.Call(obj)
    return DStrToGoStr(ret)
}

func Strings_ClassName(obj uintptr) string {
    ret, _, _ := strings_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Strings_GetHashCode(obj uintptr) int32 {
    ret, _, _ := strings_GetHashCode.Call(obj)
    return int32(ret)
}

func Strings_ToString(obj uintptr) string {
    ret, _, _ := strings_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Strings_GetCommaText(obj uintptr) string {
    ret, _, _ := strings_GetCommaText.Call(obj)
    return DStrToGoStr(ret)
}

func Strings_SetCommaText(obj uintptr, value string) {
   strings_SetCommaText.Call(obj, GoStrToDStr(value))
}

func Strings_GetDelimiter(obj uintptr) uint16 {
    ret, _, _ := strings_GetDelimiter.Call(obj)
    return uint16(ret)
}

func Strings_SetDelimiter(obj uintptr, value uint16) {
   strings_SetDelimiter.Call(obj, uintptr(value))
}

func Strings_GetText(obj uintptr) string {
    ret, _, _ := strings_GetText.Call(obj)
    return DStrToGoStr(ret)
}

func Strings_SetText(obj uintptr, value string) {
   strings_SetText.Call(obj, GoStrToDStr(value))
}

func Strings_GetWriteBOM(obj uintptr) bool {
    ret, _, _ := strings_GetWriteBOM.Call(obj)
    return DBoolToGoBool(ret)
}

func Strings_SetWriteBOM(obj uintptr, value bool) {
   strings_SetWriteBOM.Call(obj, GoBoolToDBool(value))
}

func Strings_GetOptions(obj uintptr) TStringsOptions {
    ret, _, _ := strings_GetOptions.Call(obj)
    return TStringsOptions(ret)
}

func Strings_SetOptions(obj uintptr, value TStringsOptions) {
   strings_SetOptions.Call(obj, uintptr(value))
}

func Strings_GetValues(obj uintptr, Name string) string {
    ret, _, _ := strings_GetValues.Call(obj, GoStrToDStr(Name))
    return DStrToGoStr(ret)
}

func Strings_SetValues(obj uintptr, Name string, value string) {
   strings_SetValues.Call(obj, GoStrToDStr(Name), GoStrToDStr(value))
}

func Strings_GetValueFromIndex(obj uintptr, Index int32) string {
    ret, _, _ := strings_GetValueFromIndex.Call(obj, uintptr(Index))
    return DStrToGoStr(ret)
}

func Strings_SetValueFromIndex(obj uintptr, Index int32, value string) {
   strings_SetValueFromIndex.Call(obj, uintptr(Index), GoStrToDStr(value))
}

func Strings_GetStrings(obj uintptr, Index int32) string {
    ret, _, _ := strings_GetStrings.Call(obj, uintptr(Index))
    return DStrToGoStr(ret)
}

func Strings_SetStrings(obj uintptr, Index int32, value string) {
   strings_SetStrings.Call(obj, uintptr(Index), GoStrToDStr(value))
}

