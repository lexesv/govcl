
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func Registry_Create() uintptr {
    ret, _, _ := registry_Create.Call()
    return ret
}

func Registry_Free(obj uintptr) {
    registry_Free.Call(obj)
}

func Registry_CloseKey(obj uintptr)  {
    registry_CloseKey.Call(obj)
}

func Registry_CreateKey(obj uintptr, Key string) bool {
    ret, _, _ := registry_CreateKey.Call(obj, GoStrToDStr(Key) )
    return DBoolToGoBool(ret)
}

func Registry_DeleteKey(obj uintptr, Key string) bool {
    ret, _, _ := registry_DeleteKey.Call(obj, GoStrToDStr(Key) )
    return DBoolToGoBool(ret)
}

func Registry_DeleteValue(obj uintptr, Name string) bool {
    ret, _, _ := registry_DeleteValue.Call(obj, GoStrToDStr(Name) )
    return DBoolToGoBool(ret)
}

func Registry_HasSubKeys(obj uintptr) bool {
    ret, _, _ := registry_HasSubKeys.Call(obj)
    return DBoolToGoBool(ret)
}

func Registry_KeyExists(obj uintptr, Key string) bool {
    ret, _, _ := registry_KeyExists.Call(obj, GoStrToDStr(Key) )
    return DBoolToGoBool(ret)
}

func Registry_LoadKey(obj uintptr, Key string, FileName string) bool {
    ret, _, _ := registry_LoadKey.Call(obj, GoStrToDStr(Key) , GoStrToDStr(FileName) )
    return DBoolToGoBool(ret)
}

func Registry_MoveKey(obj uintptr, OldName string, NewName string, Delete bool)  {
    registry_MoveKey.Call(obj, GoStrToDStr(OldName) , GoStrToDStr(NewName) , GoBoolToDBool(Delete) )
}

func Registry_OpenKey(obj uintptr, Key string, CanCreate bool) bool {
    ret, _, _ := registry_OpenKey.Call(obj, GoStrToDStr(Key) , GoBoolToDBool(CanCreate) )
    return DBoolToGoBool(ret)
}

func Registry_OpenKeyReadOnly(obj uintptr, Key string) bool {
    ret, _, _ := registry_OpenKeyReadOnly.Call(obj, GoStrToDStr(Key) )
    return DBoolToGoBool(ret)
}

func Registry_ReadBool(obj uintptr, Name string) bool {
    ret, _, _ := registry_ReadBool.Call(obj, GoStrToDStr(Name) )
    return DBoolToGoBool(ret)
}

func Registry_ReadDate(obj uintptr, Name string) TDateTime {
    var ret TDateTime
    registry_ReadDate.Call(obj, GoStrToDStr(Name) , uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Registry_ReadDateTime(obj uintptr, Name string) TDateTime {
    var ret TDateTime
    registry_ReadDateTime.Call(obj, GoStrToDStr(Name) , uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Registry_ReadFloat(obj uintptr, Name string) float64 {
    var ret float64
    registry_ReadFloat.Call(obj, GoStrToDStr(Name) , uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Registry_ReadInteger(obj uintptr, Name string) int32 {
    ret, _, _ := registry_ReadInteger.Call(obj, GoStrToDStr(Name) )
    return int32(ret)
}

func Registry_ReadString(obj uintptr, Name string) string {
    ret, _, _ := registry_ReadString.Call(obj, GoStrToDStr(Name) )
    return DStrToGoStr(ret)
}

func Registry_ReadTime(obj uintptr, Name string) TDateTime {
    var ret TDateTime
    registry_ReadTime.Call(obj, GoStrToDStr(Name) , uintptr(unsafe.Pointer(&ret)))
    return ret
}

func Registry_RegistryConnect(obj uintptr, UNCName string) bool {
    ret, _, _ := registry_RegistryConnect.Call(obj, GoStrToDStr(UNCName) )
    return DBoolToGoBool(ret)
}

func Registry_RenameValue(obj uintptr, OldName string, NewName string)  {
    registry_RenameValue.Call(obj, GoStrToDStr(OldName) , GoStrToDStr(NewName) )
}

func Registry_ReplaceKey(obj uintptr, Key string, FileName string, BackUpFileName string) bool {
    ret, _, _ := registry_ReplaceKey.Call(obj, GoStrToDStr(Key) , GoStrToDStr(FileName) , GoStrToDStr(BackUpFileName) )
    return DBoolToGoBool(ret)
}

func Registry_RestoreKey(obj uintptr, Key string, FileName string) bool {
    ret, _, _ := registry_RestoreKey.Call(obj, GoStrToDStr(Key) , GoStrToDStr(FileName) )
    return DBoolToGoBool(ret)
}

func Registry_SaveKey(obj uintptr, Key string, FileName string) bool {
    ret, _, _ := registry_SaveKey.Call(obj, GoStrToDStr(Key) , GoStrToDStr(FileName) )
    return DBoolToGoBool(ret)
}

func Registry_UnLoadKey(obj uintptr, Key string) bool {
    ret, _, _ := registry_UnLoadKey.Call(obj, GoStrToDStr(Key) )
    return DBoolToGoBool(ret)
}

func Registry_ValueExists(obj uintptr, Name string) bool {
    ret, _, _ := registry_ValueExists.Call(obj, GoStrToDStr(Name) )
    return DBoolToGoBool(ret)
}

func Registry_WriteBool(obj uintptr, Name string, Value bool)  {
    registry_WriteBool.Call(obj, GoStrToDStr(Name) , GoBoolToDBool(Value) )
}

func Registry_WriteDate(obj uintptr, Name string, Value TDateTime)  {
    registry_WriteDate.Call(obj, GoStrToDStr(Name) , uintptr(unsafe.Pointer(&Value)))
}

func Registry_WriteDateTime(obj uintptr, Name string, Value TDateTime)  {
    registry_WriteDateTime.Call(obj, GoStrToDStr(Name) , uintptr(unsafe.Pointer(&Value)))
}

func Registry_WriteFloat(obj uintptr, Name string, Value float64)  {
    registry_WriteFloat.Call(obj, GoStrToDStr(Name) , uintptr(unsafe.Pointer(&Value)))
}

func Registry_WriteInteger(obj uintptr, Name string, Value int32)  {
    registry_WriteInteger.Call(obj, GoStrToDStr(Name) , uintptr(Value) )
}

func Registry_WriteString(obj uintptr, Name string, Value string)  {
    registry_WriteString.Call(obj, GoStrToDStr(Name) , GoStrToDStr(Value) )
}

func Registry_WriteExpandString(obj uintptr, Name string, Value string)  {
    registry_WriteExpandString.Call(obj, GoStrToDStr(Name) , GoStrToDStr(Value) )
}

func Registry_WriteTime(obj uintptr, Name string, Value TDateTime)  {
    registry_WriteTime.Call(obj, GoStrToDStr(Name) , uintptr(unsafe.Pointer(&Value)))
}

func Registry_ClassName(obj uintptr) string {
    ret, _, _ := registry_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func Registry_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := registry_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func Registry_GetHashCode(obj uintptr) int32 {
    ret, _, _ := registry_GetHashCode.Call(obj)
    return int32(ret)
}

func Registry_ToString(obj uintptr) string {
    ret, _, _ := registry_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func Registry_GetCurrentKey(obj uintptr) HKEY {
    ret, _, _ := registry_GetCurrentKey.Call(obj)
    return HKEY(ret)
}

func Registry_GetCurrentPath(obj uintptr) string {
    ret, _, _ := registry_GetCurrentPath.Call(obj)
    return DStrToGoStr(ret)
}

func Registry_GetLazyWrite(obj uintptr) bool {
    ret, _, _ := registry_GetLazyWrite.Call(obj)
    return DBoolToGoBool(ret)
}

func Registry_SetLazyWrite(obj uintptr, value bool) {
   registry_SetLazyWrite.Call(obj, GoBoolToDBool(value))
}

func Registry_GetLastError(obj uintptr) int32 {
    ret, _, _ := registry_GetLastError.Call(obj)
    return int32(ret)
}

func Registry_GetLastErrorMsg(obj uintptr) string {
    ret, _, _ := registry_GetLastErrorMsg.Call(obj)
    return DStrToGoStr(ret)
}

func Registry_GetRootKey(obj uintptr) HKEY {
    ret, _, _ := registry_GetRootKey.Call(obj)
    return HKEY(ret)
}

func Registry_SetRootKey(obj uintptr, value HKEY) {
   registry_SetRootKey.Call(obj, uintptr(value))
}

func Registry_GetRootKeyName(obj uintptr) string {
    ret, _, _ := registry_GetRootKeyName.Call(obj)
    return DStrToGoStr(ret)
}

func Registry_GetAccess(obj uintptr) uint32 {
    ret, _, _ := registry_GetAccess.Call(obj)
    return uint32(ret)
}

func Registry_SetAccess(obj uintptr, value uint32) {
   registry_SetAccess.Call(obj, uintptr(value))
}

