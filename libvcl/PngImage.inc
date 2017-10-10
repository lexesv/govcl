
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function PngImage_Create: TPngImage; cdecl;
begin
  Result :=  TPngImage.Create;
end;

procedure PngImage_Free(AObj: TPngImage); cdecl;
begin
  AObj.Free;
end;

procedure PngImage_Assign(AObj: TPngImage; Source: TPersistent); cdecl;
begin
  AObj.Assign(Source);
end;

procedure PngImage_LoadFromStream(AObj: TPngImage; Stream: TStream); cdecl;
begin
  AObj.LoadFromStream(Stream);
end;

procedure PngImage_SaveToStream(AObj: TPngImage; Stream: TStream); cdecl;
begin
  AObj.SaveToStream(Stream);
end;

procedure PngImage_LoadFromResourceName(AObj: TPngImage; Instance: NativeUInt; Name: PChar); cdecl;
begin
  AObj.LoadFromResourceName(Instance, Name);
end;

procedure PngImage_LoadFromResourceID(AObj: TPngImage; Instance: NativeUInt; ResID: Integer); cdecl;
begin
{$IFDEF MSWINDOWS}
  AObj.LoadFromResourceID(Instance, ResID);
{$ENDIF MSWINDOWS}
end;

function PngImage_Equals(AObj: TPngImage; Obj: TObject): LongBool; cdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

procedure PngImage_LoadFromFile(AObj: TPngImage; Filename: PChar); cdecl;
begin
  AObj.LoadFromFile(Filename);
end;

procedure PngImage_SaveToFile(AObj: TPngImage; Filename: PChar); cdecl;
begin
  AObj.SaveToFile(Filename);
end;

procedure PngImage_SetSize(AObj: TPngImage; AWidth: Integer; AHeight: Integer); cdecl;
begin
  AObj.SetSize(AWidth, AHeight);
end;

function PngImage_GetNamePath(AObj: TPngImage): PChar; cdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function PngImage_ClassName(AObj: TPngImage): PChar; cdecl;
begin
  Result :=  PChar(AObj.ClassName);
end;

function PngImage_GetHashCode(AObj: TPngImage): Integer; cdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function PngImage_ToString(AObj: TPngImage): PChar; cdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function PngImage_GetTransparentColor(AObj: TPngImage): TColor; cdecl;
begin
  Result :=  AObj.TransparentColor;
end;

procedure PngImage_SetTransparentColor(AObj: TPngImage; AValue: TColor); cdecl;
begin
  AObj.TransparentColor := AValue;
end;

function PngImage_GetCanvas(AObj: TPngImage): TCanvas; cdecl;
begin
  Result :=  AObj.Canvas;
end;

function PngImage_GetWidth(AObj: TPngImage): Integer; cdecl;
begin
  Result :=  AObj.Width;
end;

function PngImage_GetHeight(AObj: TPngImage): Integer; cdecl;
begin
  Result :=  AObj.Height;
end;

function PngImage_GetMaxIdatSize(AObj: TPngImage): Integer; cdecl;
begin
  Result :=  AObj.MaxIdatSize;
end;

procedure PngImage_SetMaxIdatSize(AObj: TPngImage; AValue: Integer); cdecl;
begin
  AObj.MaxIdatSize := AValue;
end;

function PngImage_GetEmpty(AObj: TPngImage): LongBool; cdecl;
begin
  Result :=  AObj.Empty;
end;

function PngImage_GetCompressionLevel(AObj: TPngImage): TCompressionLevel; cdecl;
begin
  Result :=  AObj.CompressionLevel;
end;

procedure PngImage_SetCompressionLevel(AObj: TPngImage; AValue: TCompressionLevel); cdecl;
begin
  AObj.CompressionLevel := AValue;
end;

function PngImage_GetVersion(AObj: TPngImage): PChar; cdecl;
begin
  Result :=  PChar(AObj.Version);
end;

function PngImage_GetModified(AObj: TPngImage): LongBool; cdecl;
begin
  Result :=  AObj.Modified;
end;

procedure PngImage_SetModified(AObj: TPngImage; AValue: LongBool); cdecl;
begin
  AObj.Modified := AValue;
end;

function PngImage_GetPaletteModified(AObj: TPngImage): LongBool; cdecl;
begin
  Result :=  AObj.PaletteModified;
end;

procedure PngImage_SetPaletteModified(AObj: TPngImage; AValue: LongBool); cdecl;
begin
  AObj.PaletteModified := AValue;
end;

function PngImage_GetTransparent(AObj: TPngImage): LongBool; cdecl;
begin
  Result :=  AObj.Transparent;
end;

procedure PngImage_SetTransparent(AObj: TPngImage; AValue: LongBool); cdecl;
begin
  AObj.Transparent := AValue;
end;

procedure PngImage_SetOnChange(AObj: TPngImage; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    TEventClass.Remove(AObj, geChange);
    Exit;
  end;
  AObj.OnChange := TEventClass.OnChange;
  TEventClass.Add(AObj, geChange, AEventId);
end;


exports
  PngImage_Create {$IFNDEF MSWINDOWS}name '_PngImage_Create'{$ENDIF},
  PngImage_Free {$IFNDEF MSWINDOWS}name '_PngImage_Free'{$ENDIF},
  PngImage_Assign {$IFNDEF MSWINDOWS}name '_PngImage_Assign'{$ENDIF},
  PngImage_LoadFromStream {$IFNDEF MSWINDOWS}name '_PngImage_LoadFromStream'{$ENDIF},
  PngImage_SaveToStream {$IFNDEF MSWINDOWS}name '_PngImage_SaveToStream'{$ENDIF},
  PngImage_LoadFromResourceName {$IFNDEF MSWINDOWS}name '_PngImage_LoadFromResourceName'{$ENDIF},
  PngImage_LoadFromResourceID {$IFNDEF MSWINDOWS}name '_PngImage_LoadFromResourceID'{$ENDIF},
  PngImage_Equals {$IFNDEF MSWINDOWS}name '_PngImage_Equals'{$ENDIF},
  PngImage_LoadFromFile {$IFNDEF MSWINDOWS}name '_PngImage_LoadFromFile'{$ENDIF},
  PngImage_SaveToFile {$IFNDEF MSWINDOWS}name '_PngImage_SaveToFile'{$ENDIF},
  PngImage_SetSize {$IFNDEF MSWINDOWS}name '_PngImage_SetSize'{$ENDIF},
  PngImage_GetNamePath {$IFNDEF MSWINDOWS}name '_PngImage_GetNamePath'{$ENDIF},
  PngImage_ClassName {$IFNDEF MSWINDOWS}name '_PngImage_ClassName'{$ENDIF},
  PngImage_GetHashCode {$IFNDEF MSWINDOWS}name '_PngImage_GetHashCode'{$ENDIF},
  PngImage_ToString {$IFNDEF MSWINDOWS}name '_PngImage_ToString'{$ENDIF},
  PngImage_GetTransparentColor {$IFNDEF MSWINDOWS}name '_PngImage_GetTransparentColor'{$ENDIF},
  PngImage_SetTransparentColor {$IFNDEF MSWINDOWS}name '_PngImage_SetTransparentColor'{$ENDIF},
  PngImage_GetCanvas {$IFNDEF MSWINDOWS}name '_PngImage_GetCanvas'{$ENDIF},
  PngImage_GetWidth {$IFNDEF MSWINDOWS}name '_PngImage_GetWidth'{$ENDIF},
  PngImage_GetHeight {$IFNDEF MSWINDOWS}name '_PngImage_GetHeight'{$ENDIF},
  PngImage_GetMaxIdatSize {$IFNDEF MSWINDOWS}name '_PngImage_GetMaxIdatSize'{$ENDIF},
  PngImage_SetMaxIdatSize {$IFNDEF MSWINDOWS}name '_PngImage_SetMaxIdatSize'{$ENDIF},
  PngImage_GetEmpty {$IFNDEF MSWINDOWS}name '_PngImage_GetEmpty'{$ENDIF},
  PngImage_GetCompressionLevel {$IFNDEF MSWINDOWS}name '_PngImage_GetCompressionLevel'{$ENDIF},
  PngImage_SetCompressionLevel {$IFNDEF MSWINDOWS}name '_PngImage_SetCompressionLevel'{$ENDIF},
  PngImage_GetVersion {$IFNDEF MSWINDOWS}name '_PngImage_GetVersion'{$ENDIF},
  PngImage_GetModified {$IFNDEF MSWINDOWS}name '_PngImage_GetModified'{$ENDIF},
  PngImage_SetModified {$IFNDEF MSWINDOWS}name '_PngImage_SetModified'{$ENDIF},
  PngImage_GetPaletteModified {$IFNDEF MSWINDOWS}name '_PngImage_GetPaletteModified'{$ENDIF},
  PngImage_SetPaletteModified {$IFNDEF MSWINDOWS}name '_PngImage_SetPaletteModified'{$ENDIF},
  PngImage_GetTransparent {$IFNDEF MSWINDOWS}name '_PngImage_GetTransparent'{$ENDIF},
  PngImage_SetTransparent {$IFNDEF MSWINDOWS}name '_PngImage_SetTransparent'{$ENDIF},
  PngImage_SetOnChange {$IFNDEF MSWINDOWS}name '_PngImage_SetOnChange'{$ENDIF};

