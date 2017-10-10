
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function TrayIcon_Create(AOwner: TComponent): TTrayIcon; cdecl;
begin
  Result :=  TTrayIcon.Create(AOwner);
end;

procedure TrayIcon_Free(AObj: TTrayIcon); cdecl;
begin
  AObj.Free;
end;

procedure TrayIcon_Refresh(AObj: TTrayIcon); cdecl;
begin
  AObj.Refresh;
end;

procedure TrayIcon_SetDefaultIcon(AObj: TTrayIcon); cdecl;
begin
  AObj.SetDefaultIcon;
end;

procedure TrayIcon_ShowBalloonHint(AObj: TTrayIcon); cdecl;
begin
  AObj.ShowBalloonHint;
end;

function TrayIcon_FindComponent(AObj: TTrayIcon; AName: PChar): TComponent; cdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function TrayIcon_GetNamePath(AObj: TTrayIcon): PChar; cdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function TrayIcon_HasParent(AObj: TTrayIcon): LongBool; cdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure TrayIcon_Assign(AObj: TTrayIcon; Source: TPersistent); cdecl;
begin
  AObj.Assign(Source);
end;

function TrayIcon_ClassName(AObj: TTrayIcon): PChar; cdecl;
begin
  Result :=  PChar(AObj.ClassName);
end;

function TrayIcon_Equals(AObj: TTrayIcon; Obj: TObject): LongBool; cdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function TrayIcon_GetHashCode(AObj: TTrayIcon): Integer; cdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function TrayIcon_ToString(AObj: TTrayIcon): PChar; cdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function TrayIcon_GetAnimate(AObj: TTrayIcon): LongBool; cdecl;
begin
  Result :=  AObj.Animate;
end;

procedure TrayIcon_SetAnimate(AObj: TTrayIcon; AValue: LongBool); cdecl;
begin
  AObj.Animate := AValue;
end;

function TrayIcon_GetAnimateInterval(AObj: TTrayIcon): Cardinal; cdecl;
begin
  Result :=  AObj.AnimateInterval;
end;

procedure TrayIcon_SetAnimateInterval(AObj: TTrayIcon; AValue: Cardinal); cdecl;
begin
  AObj.AnimateInterval := AValue;
end;

function TrayIcon_GetHint(AObj: TTrayIcon): PChar; cdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure TrayIcon_SetHint(AObj: TTrayIcon; AValue: PChar); cdecl;
begin
  AObj.Hint := AValue;
end;

function TrayIcon_GetBalloonHint(AObj: TTrayIcon): PChar; cdecl;
begin
  Result :=  PChar(AObj.BalloonHint);
end;

procedure TrayIcon_SetBalloonHint(AObj: TTrayIcon; AValue: PChar); cdecl;
begin
  AObj.BalloonHint := AValue;
end;

function TrayIcon_GetBalloonTitle(AObj: TTrayIcon): PChar; cdecl;
begin
  Result :=  PChar(AObj.BalloonTitle);
end;

procedure TrayIcon_SetBalloonTitle(AObj: TTrayIcon; AValue: PChar); cdecl;
begin
  AObj.BalloonTitle := AValue;
end;

function TrayIcon_GetBalloonTimeout(AObj: TTrayIcon): Integer; cdecl;
begin
  Result :=  AObj.BalloonTimeout;
end;

procedure TrayIcon_SetBalloonTimeout(AObj: TTrayIcon; AValue: Integer); cdecl;
begin
  AObj.BalloonTimeout := AValue;
end;

function TrayIcon_GetBalloonFlags(AObj: TTrayIcon): TBalloonFlags; cdecl;
begin
  Result :=  AObj.BalloonFlags;
end;

procedure TrayIcon_SetBalloonFlags(AObj: TTrayIcon; AValue: TBalloonFlags); cdecl;
begin
  AObj.BalloonFlags := AValue;
end;

function TrayIcon_GetIcon(AObj: TTrayIcon): TIcon; cdecl;
begin
  Result :=  AObj.Icon;
end;

procedure TrayIcon_SetIcon(AObj: TTrayIcon; AValue: TIcon); cdecl;
begin
  AObj.Icon := AValue;
end;

function TrayIcon_GetIconIndex(AObj: TTrayIcon): Integer; cdecl;
begin
  Result :=  AObj.IconIndex;
end;

procedure TrayIcon_SetIconIndex(AObj: TTrayIcon; AValue: Integer); cdecl;
begin
  AObj.IconIndex := AValue;
end;

function TrayIcon_GetPopupMenu(AObj: TTrayIcon): TPopupMenu; cdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure TrayIcon_SetPopupMenu(AObj: TTrayIcon; AValue: TPopupMenu); cdecl;
begin
  AObj.PopupMenu := AValue;
end;

function TrayIcon_GetVisible(AObj: TTrayIcon): LongBool; cdecl;
begin
  Result :=  AObj.Visible;
end;

procedure TrayIcon_SetVisible(AObj: TTrayIcon; AValue: LongBool); cdecl;
begin
  AObj.Visible := AValue;
end;

procedure TrayIcon_SetOnBalloonClick(AObj: TTrayIcon; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnBalloonClick := nil;
    TEventClass.Remove(AObj, geBalloonClick);
    Exit;
  end;
  AObj.OnBalloonClick := TEventClass.OnBalloonClick;
  TEventClass.Add(AObj, geBalloonClick, AEventId);
end;

procedure TrayIcon_SetOnClick(AObj: TTrayIcon; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnClick := nil;
    TEventClass.Remove(AObj, geClick);
    Exit;
  end;
  AObj.OnClick := TEventClass.OnClick;
  TEventClass.Add(AObj, geClick, AEventId);
end;

procedure TrayIcon_SetOnDblClick(AObj: TTrayIcon; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnDblClick := nil;
    TEventClass.Remove(AObj, geDblClick);
    Exit;
  end;
  AObj.OnDblClick := TEventClass.OnDblClick;
  TEventClass.Add(AObj, geDblClick, AEventId);
end;

procedure TrayIcon_SetOnMouseMove(AObj: TTrayIcon; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseMove := nil;
    TEventClass.Remove(AObj, geMouseMove);
    Exit;
  end;
  AObj.OnMouseMove := TEventClass.OnMouseMove;
  TEventClass.Add(AObj, geMouseMove, AEventId);
end;

procedure TrayIcon_SetOnMouseUp(AObj: TTrayIcon; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseUp := nil;
    TEventClass.Remove(AObj, geMouseUp);
    Exit;
  end;
  AObj.OnMouseUp := TEventClass.OnMouseUp;
  TEventClass.Add(AObj, geMouseUp, AEventId);
end;

procedure TrayIcon_SetOnMouseDown(AObj: TTrayIcon; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseDown := nil;
    TEventClass.Remove(AObj, geMouseDown);
    Exit;
  end;
  AObj.OnMouseDown := TEventClass.OnMouseDown;
  TEventClass.Add(AObj, geMouseDown, AEventId);
end;

function TrayIcon_GetComponentCount(AObj: TTrayIcon): Integer; cdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function TrayIcon_GetComponentIndex(AObj: TTrayIcon): Integer; cdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure TrayIcon_SetComponentIndex(AObj: TTrayIcon; AValue: Integer); cdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function TrayIcon_GetOwner(AObj: TTrayIcon): TComponent; cdecl;
begin
  Result :=  AObj.Owner;
end;

function TrayIcon_GetName(AObj: TTrayIcon): PChar; cdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure TrayIcon_SetName(AObj: TTrayIcon; AValue: PChar); cdecl;
begin
  AObj.Name := AValue;
end;

function TrayIcon_GetTag(AObj: TTrayIcon): NativeInt; cdecl;
begin
  Result :=  AObj.Tag;
end;

procedure TrayIcon_SetTag(AObj: TTrayIcon; AValue: NativeInt); cdecl;
begin
  AObj.Tag := AValue;
end;

function TrayIcon_GetComponents(AObj: TTrayIcon; AIndex: Integer): TComponent; cdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  TrayIcon_Create {$IFNDEF MSWINDOWS}name '_TrayIcon_Create'{$ENDIF},
  TrayIcon_Free {$IFNDEF MSWINDOWS}name '_TrayIcon_Free'{$ENDIF},
  TrayIcon_Refresh {$IFNDEF MSWINDOWS}name '_TrayIcon_Refresh'{$ENDIF},
  TrayIcon_SetDefaultIcon {$IFNDEF MSWINDOWS}name '_TrayIcon_SetDefaultIcon'{$ENDIF},
  TrayIcon_ShowBalloonHint {$IFNDEF MSWINDOWS}name '_TrayIcon_ShowBalloonHint'{$ENDIF},
  TrayIcon_FindComponent {$IFNDEF MSWINDOWS}name '_TrayIcon_FindComponent'{$ENDIF},
  TrayIcon_GetNamePath {$IFNDEF MSWINDOWS}name '_TrayIcon_GetNamePath'{$ENDIF},
  TrayIcon_HasParent {$IFNDEF MSWINDOWS}name '_TrayIcon_HasParent'{$ENDIF},
  TrayIcon_Assign {$IFNDEF MSWINDOWS}name '_TrayIcon_Assign'{$ENDIF},
  TrayIcon_ClassName {$IFNDEF MSWINDOWS}name '_TrayIcon_ClassName'{$ENDIF},
  TrayIcon_Equals {$IFNDEF MSWINDOWS}name '_TrayIcon_Equals'{$ENDIF},
  TrayIcon_GetHashCode {$IFNDEF MSWINDOWS}name '_TrayIcon_GetHashCode'{$ENDIF},
  TrayIcon_ToString {$IFNDEF MSWINDOWS}name '_TrayIcon_ToString'{$ENDIF},
  TrayIcon_GetAnimate {$IFNDEF MSWINDOWS}name '_TrayIcon_GetAnimate'{$ENDIF},
  TrayIcon_SetAnimate {$IFNDEF MSWINDOWS}name '_TrayIcon_SetAnimate'{$ENDIF},
  TrayIcon_GetAnimateInterval {$IFNDEF MSWINDOWS}name '_TrayIcon_GetAnimateInterval'{$ENDIF},
  TrayIcon_SetAnimateInterval {$IFNDEF MSWINDOWS}name '_TrayIcon_SetAnimateInterval'{$ENDIF},
  TrayIcon_GetHint {$IFNDEF MSWINDOWS}name '_TrayIcon_GetHint'{$ENDIF},
  TrayIcon_SetHint {$IFNDEF MSWINDOWS}name '_TrayIcon_SetHint'{$ENDIF},
  TrayIcon_GetBalloonHint {$IFNDEF MSWINDOWS}name '_TrayIcon_GetBalloonHint'{$ENDIF},
  TrayIcon_SetBalloonHint {$IFNDEF MSWINDOWS}name '_TrayIcon_SetBalloonHint'{$ENDIF},
  TrayIcon_GetBalloonTitle {$IFNDEF MSWINDOWS}name '_TrayIcon_GetBalloonTitle'{$ENDIF},
  TrayIcon_SetBalloonTitle {$IFNDEF MSWINDOWS}name '_TrayIcon_SetBalloonTitle'{$ENDIF},
  TrayIcon_GetBalloonTimeout {$IFNDEF MSWINDOWS}name '_TrayIcon_GetBalloonTimeout'{$ENDIF},
  TrayIcon_SetBalloonTimeout {$IFNDEF MSWINDOWS}name '_TrayIcon_SetBalloonTimeout'{$ENDIF},
  TrayIcon_GetBalloonFlags {$IFNDEF MSWINDOWS}name '_TrayIcon_GetBalloonFlags'{$ENDIF},
  TrayIcon_SetBalloonFlags {$IFNDEF MSWINDOWS}name '_TrayIcon_SetBalloonFlags'{$ENDIF},
  TrayIcon_GetIcon {$IFNDEF MSWINDOWS}name '_TrayIcon_GetIcon'{$ENDIF},
  TrayIcon_SetIcon {$IFNDEF MSWINDOWS}name '_TrayIcon_SetIcon'{$ENDIF},
  TrayIcon_GetIconIndex {$IFNDEF MSWINDOWS}name '_TrayIcon_GetIconIndex'{$ENDIF},
  TrayIcon_SetIconIndex {$IFNDEF MSWINDOWS}name '_TrayIcon_SetIconIndex'{$ENDIF},
  TrayIcon_GetPopupMenu {$IFNDEF MSWINDOWS}name '_TrayIcon_GetPopupMenu'{$ENDIF},
  TrayIcon_SetPopupMenu {$IFNDEF MSWINDOWS}name '_TrayIcon_SetPopupMenu'{$ENDIF},
  TrayIcon_GetVisible {$IFNDEF MSWINDOWS}name '_TrayIcon_GetVisible'{$ENDIF},
  TrayIcon_SetVisible {$IFNDEF MSWINDOWS}name '_TrayIcon_SetVisible'{$ENDIF},
  TrayIcon_SetOnBalloonClick {$IFNDEF MSWINDOWS}name '_TrayIcon_SetOnBalloonClick'{$ENDIF},
  TrayIcon_SetOnClick {$IFNDEF MSWINDOWS}name '_TrayIcon_SetOnClick'{$ENDIF},
  TrayIcon_SetOnDblClick {$IFNDEF MSWINDOWS}name '_TrayIcon_SetOnDblClick'{$ENDIF},
  TrayIcon_SetOnMouseMove {$IFNDEF MSWINDOWS}name '_TrayIcon_SetOnMouseMove'{$ENDIF},
  TrayIcon_SetOnMouseUp {$IFNDEF MSWINDOWS}name '_TrayIcon_SetOnMouseUp'{$ENDIF},
  TrayIcon_SetOnMouseDown {$IFNDEF MSWINDOWS}name '_TrayIcon_SetOnMouseDown'{$ENDIF},
  TrayIcon_GetComponentCount {$IFNDEF MSWINDOWS}name '_TrayIcon_GetComponentCount'{$ENDIF},
  TrayIcon_GetComponentIndex {$IFNDEF MSWINDOWS}name '_TrayIcon_GetComponentIndex'{$ENDIF},
  TrayIcon_SetComponentIndex {$IFNDEF MSWINDOWS}name '_TrayIcon_SetComponentIndex'{$ENDIF},
  TrayIcon_GetOwner {$IFNDEF MSWINDOWS}name '_TrayIcon_GetOwner'{$ENDIF},
  TrayIcon_GetName {$IFNDEF MSWINDOWS}name '_TrayIcon_GetName'{$ENDIF},
  TrayIcon_SetName {$IFNDEF MSWINDOWS}name '_TrayIcon_SetName'{$ENDIF},
  TrayIcon_GetTag {$IFNDEF MSWINDOWS}name '_TrayIcon_GetTag'{$ENDIF},
  TrayIcon_SetTag {$IFNDEF MSWINDOWS}name '_TrayIcon_SetTag'{$ENDIF},
  TrayIcon_GetComponents {$IFNDEF MSWINDOWS}name '_TrayIcon_GetComponents'{$ENDIF};

