
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function MonthCalColors_Create(AOwner: TCommonCalendar): TMonthCalColors; stdcall;
begin
  Result :=  TMonthCalColors.Create(AOwner);
end;

procedure MonthCalColors_Free(AObj: TMonthCalColors); stdcall;
begin
  AObj.Free;
end;

procedure MonthCalColors_Assign(AObj: TMonthCalColors; Source: TPersistent); stdcall;
begin
  AObj.Assign(Source);
end;

function MonthCalColors_GetNamePath(AObj: TMonthCalColors): PChar; stdcall;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function MonthCalColors_ClassName(AObj: TMonthCalColors): PChar; stdcall;
begin
  Result :=  PChar(AObj.ClassName);
end;

function MonthCalColors_Equals(AObj: TMonthCalColors; Obj: TObject): LongBool; stdcall;
begin
  Result :=  AObj.Equals(Obj);
end;

function MonthCalColors_GetHashCode(AObj: TMonthCalColors): Integer; stdcall;
begin
  Result :=  AObj.GetHashCode;
end;

function MonthCalColors_ToString(AObj: TMonthCalColors): PChar; stdcall;
begin
  Result :=  PChar(AObj.ToString);
end;


exports
  MonthCalColors_Create,
  MonthCalColors_Free,
  MonthCalColors_Assign,
  MonthCalColors_GetNamePath,
  MonthCalColors_ClassName,
  MonthCalColors_Equals,
  MonthCalColors_GetHashCode,
  MonthCalColors_ToString;

