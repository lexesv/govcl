
// 枚举都按4字节对齐
{$Z4+}

type
  // 重定定义
  TUnixDateTime = type LongInt; // int64改为LongInt。只精确到秒

function ToUnixTime(ADateTime: TDateTime): TUnixDateTime; inline;
begin
  Result := DateTimeToUnix(ADateTime, False);
end;

function UnixToTime(ADateTime: TUnixDateTime): TDateTime; inline;
begin
  Result := UnixToDateTime(ADateTime, False);
end;

// TApplication
function Application_Instance: TApplication; stdcall;
begin
  Result := Application;
end;

function Application_CreateForm(App: TApplication): TGoForm; stdcall;
begin
  App.CreateForm(TGoForm, Result);
end;

// TForm

procedure Form_EnabledMaximize(AObj: TGoForm; AValue: Boolean); stdcall;
begin
  if AValue then
  begin
    if not(biMaximize in AObj.BorderIcons) then
      AObj.BorderIcons := AObj.BorderIcons + [biMaximize]
  end else
  begin
    if biMaximize in AObj.BorderIcons then
      AObj.BorderIcons := AObj.BorderIcons - [biMaximize]
  end;
end;

procedure Form_EnabledMinimize(AObj: TGoForm; AValue: Boolean); stdcall;
begin
  if AValue then
  begin
    if not(biMinimize in AObj.BorderIcons) then
      AObj.BorderIcons := AObj.BorderIcons + [biMinimize]
  end else
  begin
    if biMinimize in AObj.BorderIcons then
      AObj.BorderIcons := AObj.BorderIcons - [biMinimize]
  end;
end;

procedure Form_EnabledSystemMenu(AObj: TGoForm; AValue: Boolean); stdcall;
begin
  if AValue then
  begin
    if not(biSystemMenu in AObj.BorderIcons) then
      AObj.BorderIcons := AObj.BorderIcons + [biSystemMenu]
  end else
  begin
    if biSystemMenu in AObj.BorderIcons then
      AObj.BorderIcons := AObj.BorderIcons - [biSystemMenu]
  end;
end;

// mouse
function Mouse_Instance: TMouse; stdcall;
begin
  Result := Mouse;
end;

// screen
function Screen_Instance: TScreen; stdcall;
begin
  Result := Screen;
end;

function DTextToShortCut(AText: PChar): TShortCut; stdcall;
begin
  Result := TextToShortCut(AText);
end;

function DShortCutToText(AVal: TShortCut): PChar; stdcall;
begin
  Result := PChar(ShortCutToText(AVal));
end;

// TClipboard
function Clipboard_Instance: TClipboard; stdcall;
begin
  Result := Clipboard;
end;

function Clipboard_SetClipboard(NewClipboard: TClipboard): TClipboard; stdcall;
begin
  Result := SetClipboard(NewClipboard);
end;

// TMemoryStream
// function Write(const Buffer; Count: Longint): Longint;
function MemoryStream_Write(AObj: TMemoryStream; Buffer: Pointer; Count: Longint): Longint; stdcall;
begin
  Result := AObj.Write(Buffer^, Count);
end;

// function Read(var Buffer; Count: Longint): Longint;
function MemoryStream_Read(AObj: TMemoryStream; Buffer: Pointer; Count: Longint): Longint; stdcall;
begin
  Result := AObj.Read(Buffer^, Count);
end;

// TCanvas
// Canvas_BrushCopy
// Canvas_CopyRect
// Canvas_Draw1
// Canvas_Draw2
// Canvas_DrawFocusRect
// Canvas_FillRect
// Canvas_FrameRect
// Canvas_StretchDraw
// Canvas_TextRect1
// Canvas_TextRect2

procedure Canvas_BrushCopy(AObj: TCanvas; var Dest: TRect; Bitmap: TBitmap;
      var Source: TRect; Color: TColor); stdcall;
begin
  AObj.BrushCopy(Dest, Bitmap, Source, Color);
end;

procedure Canvas_CopyRect(AObj: TCanvas; var Dest: TRect; Canvas: TCanvas;
  var Source: TRect); stdcall;
begin
  AObj.CopyRect(Dest, Canvas, Source);
end;

procedure Canvas_Draw1(AObj: TCanvas; X, Y: Integer; Graphic: TGraphic); stdcall;
begin
  AObj.Draw(X, Y, Graphic);
end;

procedure Canvas_Draw2(AObj: TCanvas; X, Y: Integer; Graphic: TGraphic; Opacity: Byte); stdcall;
begin
  AObj.Draw(X, Y, Graphic, Opacity);
end;

procedure Canvas_DrawFocusRect(AObj: TCanvas; var ARect: TRect); stdcall;
begin
  AObj.DrawFocusRect(ARect);
end;

procedure Canvas_FillRect(AObj: TCanvas; var Rect: TRect); stdcall;
begin
  AObj.FillRect(Rect);
end;

procedure Canvas_FrameRect(AObj: TCanvas; var Rect: TRect); stdcall;
begin
  AObj.FrameRect(Rect);
end;

procedure Canvas_StretchDraw(AObj: TCanvas; var Rect: TRect; Graphic: TGraphic); stdcall;
begin
  AObj.StretchDraw(Rect, Graphic);
end;

procedure Canvas_TextRect1(AObj: TCanvas; var Rect: TRect; X, Y: Integer; const Text: PChar); stdcall;
begin
  AObj.TextRect(Rect, X, Y, Text);
end;

function Canvas_TextRect2(AObj: TCanvas; var Rect: TRect; Text: PChar; var AOutStr: PChar; TextFormat: TTextFormat): Integer; stdcall;
var
  S: string;
begin
  Result := 0;
  S := Text;
  AObj.TextRect(Rect, S, TextFormat);
  if tfModifyString in TextFormat then
  begin
    AOutStr := PChar(S);
    Result := 1;
  end;
end;

// TImageList
// ImageList_Draw1
// ImageList_Draw2
// ImageList_DrawOverlay1
// ImageList_DrawOverlay2
// ImageList_GetIcon1
// ImageList_GetIcon2

procedure ImageList_Draw1(AObj: TImageList; Canvas: TCanvas; X, Y, Index: Integer; Enabled: LongBool); stdcall;
begin
  AObj.Draw(Canvas, X, Y, Index, Enabled);
end;

procedure ImageList_Draw2(AObj: TImageList; Canvas: TCanvas; X, Y, Index: Integer;
  ADrawingStyle: TDrawingStyle; AImageType: TImageType;
  Enabled: LongBool); stdcall;
begin
  AObj.Draw(Canvas, X, Y, Index, ADrawingStyle, AImageType, Enabled);
end;
  
procedure ImageList_DrawOverlay1(AObj: TImageList; Canvas: TCanvas; X, Y: Integer;
  ImageIndex: Integer; Overlay: TOverlay; Enabled: LongBool); stdcall;
begin
  AObj.DrawOverlay(Canvas, X, Y, ImageIndex, Overlay, Enabled);
end;
  
procedure ImageList_DrawOverlay2(AObj: TImageList; Canvas: TCanvas; X, Y: Integer;
  ImageIndex: Integer; Overlay: TOverlay; ADrawingStyle: TDrawingStyle;
  AImageType: TImageType; Enabled: LongBool); stdcall;
begin
  AObj.DrawOverlay(Canvas, X, Y, ImageIndex, Overlay, ADrawingStyle, AImageType, Enabled);
end;
  
procedure ImageList_GetIcon1(AObj: TImageList; Index: Integer; Image: TIcon); stdcall;
begin
  AObj.GetIcon(Index, Image);
end;

procedure ImageList_GetIcon2(AObj: TImageList; Index: Integer; Image: TIcon; ADrawingStyle: TDrawingStyle; AImageType: TImageType); stdcall;
begin
  AObj.GetIcon(Index, Image, ADrawingStyle, AImageType);
end;
	
 
// Other

procedure SetEventCallback(APtr: Pointer); stdcall;
begin
  GCallbackPtr := APtr;
end;

procedure DGetParam(Index: NativeInt; AArr: Pointer; out P: TGoParam); stdcall;
begin
  P := PGoParam(PByte(AArr) + Index * SizeOf(TGoParam))^;
end;

function DStrLen(p: PChar): NativeInt; stdcall;
begin
  Result := Length(p);
end;

procedure DMove(Src, Dest: Pointer; Len: NativeInt); stdcall;
begin
  Move(Src^, Dest^, Len);
end;

function DGetClassName(Obj: TObject): PChar; stdcall;
begin
  Result := PChar(Obj.ClassName);
end;

procedure DShowMessage(AMsg: PChar); stdcall;
begin
  ShowMessage(AMsg);
end;

function DGetMainInstance: HINST; stdcall;
begin
  Result := GetModuleHandle(nil);
end;

function DMessageDlg(const Msg: PChar; DlgType: TMsgDlgType;
  Buttons: TMsgDlgButtons; HelpCtx: Longint): Integer; stdcall;
begin
  Result := MessageDlg(Msg, DlgType, Buttons, HelpCtx);
end;

procedure DSetReportMemoryLeaksOnShutdown(Value: Boolean); stdcall;
begin
{$IFDEF MSWINDOWS}
  ReportMemoryLeaksOnShutdown := Value;
{$ENDIF MSWINDOWS}
end;

procedure DSysOpen(FileName: PChar); stdcall;
begin
  ShellExecute(0, nil, FileName, nil, nil, SW_SHOW);
end;

// 不想导入包，所以干脆用delphi的得了
function DExtractFilePath(AFileName: PChar): PChar; stdcall;
begin
  Result := PChar(ExtractFilePath(AFileName));
end;

function DFileExists(AFileName: PChar): LongBool; stdcall;
begin
  Result := FileExists(AFileName);
end;

// TStyleManager
// StyleManager_Initialize
// StyleManager_UnInitialize
// StyleManager_IsValidStyle
// StyleManager_LoadFromFile
// StyleManager_CheckSysClassName
// StyleManager_TrySetStyle
// StyleManager_SetStyle1
// StyleManager_SetStyle2
// StyleManager_TryLoadFromResource

procedure StyleManager_Initialize; cdecl
begin
  TStyleManager.Initialize;
end;

procedure StyleManager_UnInitialize; cdecl
begin
  TStyleManager.UnInitialize;
end;

function StyleManager_IsValidStyle(AFileName: PChar): LongBool; stdcall;
begin
  Result := TStyleManager.IsValidStyle(AFileName);
end;

function StyleManager_LoadFromFile(AFileName: PChar): TStyleManager.TStyleServicesHandle; stdcall;
begin
  Result := TStyleManager.LoadFromFile(AFileName);
end;

function StyleManager_CheckSysClassName(AClassName: PChar): LongBool; stdcall;
begin
  Result := TStyleManager.CheckSysClassName(AClassName);
end;

function StyleManager_TrySetStyle(AName: PChar; ShowErrorDialog: LongBool): LongBool; stdcall;
begin
  Result := TStyleManager.TrySetStyle(AName, ShowErrorDialog);
end;

procedure StyleManager_SetStyle1(Handle: TStyleManager.TStyleServicesHandle); stdcall;
begin
  TStyleManager.SetStyle(Handle);
end;

procedure StyleManager_SetStyle2(AName: PChar); stdcall;
begin
  TStyleManager.SetStyle(string(AName));
end;

function StyleManager_TryLoadFromResource(Instance: HINST; ResourceName: PChar;
   ResourceType: PChar; var Handle: TStyleManager.TStyleServicesHandle): LongBool; stdcall;
begin
  Result := TStyleManager.TryLoadFromResource(Instance, ResourceName, ResourceType, Handle);
end;

// StyleManager_ActiveStyle
// StyleManager_SystemStyle
// StyleManager_Enabled
// StyleManager_IsCustomStyleActive
// StyleManager_UnRegisterStyle
// StyleManager_RegisterStyle
// StyleManager_Style
// StyleManager_StyleDescriptor
function StyleManager_ActiveStyle: TCustomStyleServices; stdcall;
begin
  Result := TStyleManager.ActiveStyle;
end;

function StyleManager_SystemStyle: TCustomStyleServices; stdcall;
begin
  Result := TStyleManager.SystemStyle;
end;

function StyleManager_Enabled: LongBool; stdcall;
begin
  Result := TStyleManager.Enabled;
end;

function StyleManager_IsCustomStyleActive: LongBool; stdcall;
begin
  Result := TStyleManager.IsCustomStyleActive;
end;

procedure StyleManager_UnRegisterStyle(Style: TCustomStyleServices); stdcall;
begin
  TStyleManager.UnRegisterStyle(Style);
end;

procedure StyleManager_RegisterStyle(Style: TCustomStyleServices); stdcall;
begin
  TStyleManager.RegisterStyle(Style);
end;

function StyleManager_Style(name: PChar): TCustomStyleServices; stdcall;
begin
  Result := TStyleManager.Style[name];
end;

function StyleManager_StyleDescriptor(StyleName: PChar): TStyleManager.TStyleClassDescriptor; stdcall;
begin
  Result := TStyleManager.StyleDescriptor[StyleName];
end;



exports
  Application_Instance,
  Application_CreateForm,

  Form_EnabledMaximize,
  Form_EnabledMinimize,
  Form_EnabledSystemMenu,


  DSysOpen,


  Mouse_Instance,
  Screen_Instance,

  SetEventCallback,
  DGetParam,
  DStrLen,
  DMove,
  DGetClassName,
  DShowMessage,
  DGetMainInstance,


  MemoryStream_Write,
  MemoryStream_Read,

  DTextToShortCut,
  DShortCutToText,

  Clipboard_Instance,
  Clipboard_SetClipboard,

  DMessageDlg,

  Canvas_BrushCopy,
  Canvas_CopyRect,
  Canvas_Draw1,
  Canvas_Draw2,
  Canvas_DrawFocusRect,
  Canvas_FillRect,
  Canvas_FrameRect,
  Canvas_StretchDraw,
  Canvas_TextRect1,
  Canvas_TextRect2,

  // TImageList
  ImageList_Draw1,
  ImageList_Draw2,
  ImageList_DrawOverlay1,
  ImageList_DrawOverlay2,
  ImageList_GetIcon1,
  ImageList_GetIcon2,

// TStyleManager
   StyleManager_Initialize,
   StyleManager_UnInitialize,
   StyleManager_IsValidStyle,
   StyleManager_LoadFromFile,
   StyleManager_CheckSysClassName,
   StyleManager_TrySetStyle,
   StyleManager_SetStyle1,
   StyleManager_SetStyle2,
   StyleManager_TryLoadFromResource,

   StyleManager_ActiveStyle,
   StyleManager_SystemStyle,
   StyleManager_Enabled,
   StyleManager_IsCustomStyleActive,
   StyleManager_UnRegisterStyle,
   StyleManager_RegisterStyle,
   StyleManager_Style,
   StyleManager_StyleDescriptor,

  DExtractFilePath,
  DFileExists,

  DSetReportMemoryLeaksOnShutdown
  ;
