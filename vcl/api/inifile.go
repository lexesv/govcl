
//----------------------------------------
// 代码由Genvcllib工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package api

import (
    "unsafe"
)

func IniFile_Create(filename string) uintptr {
    ret, _, _ := iniFile_Create.Call(GoStrToDStr(filename))
    return ret
}

func IniFile_Free(obj uintptr) {
    iniFile_Free.Call(obj)
}

func IniFile_ReadString(obj uintptr, Section string, Ident string, Default string) string {
    ret, _, _ := iniFile_ReadString.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Ident) , GoStrToDStr(Default) )
    return DStrToGoStr(ret)
}

func IniFile_WriteString(obj uintptr, Section string, Ident string, Value string)  {
    iniFile_WriteString.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Ident) , GoStrToDStr(Value) )
}

func IniFile_ReadSections(obj uintptr, Strings uintptr)  {
    iniFile_ReadSections.Call(obj, Strings )
}

func IniFile_ReadSectionValues(obj uintptr, Section string, Strings uintptr)  {
    iniFile_ReadSectionValues.Call(obj, GoStrToDStr(Section) , Strings )
}

func IniFile_EraseSection(obj uintptr, Section string)  {
    iniFile_EraseSection.Call(obj, GoStrToDStr(Section) )
}

func IniFile_DeleteKey(obj uintptr, Section string, Ident string)  {
    iniFile_DeleteKey.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Ident) )
}

func IniFile_UpdateFile(obj uintptr)  {
    iniFile_UpdateFile.Call(obj)
}

func IniFile_SectionExists(obj uintptr, Section string) bool {
    ret, _, _ := iniFile_SectionExists.Call(obj, GoStrToDStr(Section) )
    return DBoolToGoBool(ret)
}

func IniFile_ReadInteger(obj uintptr, Section string, Ident string, Default int32) int32 {
    ret, _, _ := iniFile_ReadInteger.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Ident) , uintptr(Default) )
    return int32(ret)
}

func IniFile_WriteInteger(obj uintptr, Section string, Ident string, Value int32)  {
    iniFile_WriteInteger.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Ident) , uintptr(Value) )
}

func IniFile_ReadBool(obj uintptr, Section string, Ident string, Default bool) bool {
    ret, _, _ := iniFile_ReadBool.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Ident) , GoBoolToDBool(Default) )
    return DBoolToGoBool(ret)
}

func IniFile_WriteBool(obj uintptr, Section string, Ident string, Value bool)  {
    iniFile_WriteBool.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Ident) , GoBoolToDBool(Value) )
}

func IniFile_ReadDate(obj uintptr, Section string, Name string, Default TDateTime) TDateTime {
    var ret TDateTime
    iniFile_ReadDate.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Name) , uintptr(unsafe.Pointer(&Default)), uintptr(unsafe.Pointer(&ret)))
    return ret
}

func IniFile_ReadDateTime(obj uintptr, Section string, Name string, Default TDateTime) TDateTime {
    var ret TDateTime
    iniFile_ReadDateTime.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Name) , uintptr(unsafe.Pointer(&Default)), uintptr(unsafe.Pointer(&ret)))
    return ret
}

func IniFile_ReadFloat(obj uintptr, Section string, Name string, Default float64) float64 {
    var ret float64
    iniFile_ReadFloat.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Name) , uintptr(unsafe.Pointer(&Default)), uintptr(unsafe.Pointer(&ret)))
    return ret
}

func IniFile_ReadTime(obj uintptr, Section string, Name string, Default TDateTime) TDateTime {
    var ret TDateTime
    iniFile_ReadTime.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Name) , uintptr(unsafe.Pointer(&Default)), uintptr(unsafe.Pointer(&ret)))
    return ret
}

func IniFile_WriteDate(obj uintptr, Section string, Name string, Value TDateTime)  {
    iniFile_WriteDate.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Name) , uintptr(unsafe.Pointer(&Value)))
}

func IniFile_WriteDateTime(obj uintptr, Section string, Name string, Value TDateTime)  {
    iniFile_WriteDateTime.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Name) , uintptr(unsafe.Pointer(&Value)))
}

func IniFile_WriteFloat(obj uintptr, Section string, Name string, Value float64)  {
    iniFile_WriteFloat.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Name) , uintptr(unsafe.Pointer(&Value)))
}

func IniFile_WriteTime(obj uintptr, Section string, Name string, Value TDateTime)  {
    iniFile_WriteTime.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Name) , uintptr(unsafe.Pointer(&Value)))
}

func IniFile_ReadSubSections(obj uintptr, Section string, Strings uintptr, Recurse bool)  {
    iniFile_ReadSubSections.Call(obj, GoStrToDStr(Section) , Strings , GoBoolToDBool(Recurse) )
}

func IniFile_ValueExists(obj uintptr, Section string, Ident string) bool {
    ret, _, _ := iniFile_ValueExists.Call(obj, GoStrToDStr(Section) , GoStrToDStr(Ident) )
    return DBoolToGoBool(ret)
}

func IniFile_ClassName(obj uintptr) string {
    ret, _, _ := iniFile_ClassName.Call(obj)
    return DStrToGoStr(ret)
}

func IniFile_Equals(obj uintptr, Obj uintptr) bool {
    ret, _, _ := iniFile_Equals.Call(obj, Obj )
    return DBoolToGoBool(ret)
}

func IniFile_GetHashCode(obj uintptr) int32 {
    ret, _, _ := iniFile_GetHashCode.Call(obj)
    return int32(ret)
}

func IniFile_ToString(obj uintptr) string {
    ret, _, _ := iniFile_ToString.Call(obj)
    return DStrToGoStr(ret)
}

func IniFile_GetFileName(obj uintptr) string {
    ret, _, _ := iniFile_GetFileName.Call(obj)
    return DStrToGoStr(ret)
}

