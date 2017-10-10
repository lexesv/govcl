
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function UpDown_Create(AOwner: TComponent): TUpDown; cdecl;
begin
  Result :=  TUpDown.Create(AOwner);
end;

procedure UpDown_Free(AObj: TUpDown); cdecl;
begin
  AObj.Free;
end;

function UpDown_CanFocus(AObj: TUpDown): LongBool; cdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure UpDown_FlipChildren(AObj: TUpDown; AllLevels: LongBool); cdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function UpDown_Focused(AObj: TUpDown): LongBool; cdecl;
begin
  Result :=  AObj.Focused;
end;

function UpDown_HandleAllocated(AObj: TUpDown): LongBool; cdecl;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure UpDown_Invalidate(AObj: TUpDown); cdecl;
begin
  AObj.Invalidate;
end;

procedure UpDown_Realign(AObj: TUpDown); cdecl;
begin
  AObj.Realign;
end;

procedure UpDown_Repaint(AObj: TUpDown); cdecl;
begin
  AObj.Repaint;
end;

procedure UpDown_ScaleBy(AObj: TUpDown; M: Integer; D: Integer); cdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure UpDown_SetBounds(AObj: TUpDown; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); cdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure UpDown_SetFocus(AObj: TUpDown); cdecl;
begin
  AObj.SetFocus;
end;

procedure UpDown_Update(AObj: TUpDown); cdecl;
begin
  AObj.Update;
end;

procedure UpDown_BringToFront(AObj: TUpDown); cdecl;
begin
  AObj.BringToFront;
end;

function UpDown_HasParent(AObj: TUpDown): LongBool; cdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure UpDown_Hide(AObj: TUpDown); cdecl;
begin
  AObj.Hide;
end;

function UpDown_Perform(AObj: TUpDown; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; cdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure UpDown_Refresh(AObj: TUpDown); cdecl;
begin
  AObj.Refresh;
end;

procedure UpDown_SendToBack(AObj: TUpDown); cdecl;
begin
  AObj.SendToBack;
end;

procedure UpDown_Show(AObj: TUpDown); cdecl;
begin
  AObj.Show;
end;

function UpDown_GetTextBuf(AObj: TUpDown; Buffer: PWideChar; BufSize: Integer): Integer; cdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function UpDown_FindComponent(AObj: TUpDown; AName: PChar): TComponent; cdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function UpDown_GetNamePath(AObj: TUpDown): PChar; cdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure UpDown_Assign(AObj: TUpDown; Source: TPersistent); cdecl;
begin
  AObj.Assign(Source);
end;

function UpDown_ClassName(AObj: TUpDown): PChar; cdecl;
begin
  Result :=  PChar(AObj.ClassName);
end;

function UpDown_Equals(AObj: TUpDown; Obj: TObject): LongBool; cdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function UpDown_GetHashCode(AObj: TUpDown): Integer; cdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function UpDown_ToString(AObj: TUpDown): PChar; cdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function UpDown_GetAnchors(AObj: TUpDown): TAnchors; cdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure UpDown_SetAnchors(AObj: TUpDown; AValue: TAnchors); cdecl;
begin
  AObj.Anchors := AValue;
end;

function UpDown_GetDoubleBuffered(AObj: TUpDown): LongBool; cdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure UpDown_SetDoubleBuffered(AObj: TUpDown; AValue: LongBool); cdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function UpDown_GetEnabled(AObj: TUpDown): LongBool; cdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure UpDown_SetEnabled(AObj: TUpDown; AValue: LongBool); cdecl;
begin
  AObj.Enabled := AValue;
end;

function UpDown_GetHint(AObj: TUpDown): PChar; cdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure UpDown_SetHint(AObj: TUpDown; AValue: PChar); cdecl;
begin
  AObj.Hint := AValue;
end;

function UpDown_GetMin(AObj: TUpDown): Integer; cdecl;
begin
  Result :=  AObj.Min;
end;

procedure UpDown_SetMin(AObj: TUpDown; AValue: Integer); cdecl;
begin
  AObj.Min := AValue;
end;

function UpDown_GetMax(AObj: TUpDown): Integer; cdecl;
begin
  Result :=  AObj.Max;
end;

procedure UpDown_SetMax(AObj: TUpDown; AValue: Integer); cdecl;
begin
  AObj.Max := AValue;
end;

function UpDown_GetOrientation(AObj: TUpDown): TUDOrientation; cdecl;
begin
  Result :=  AObj.Orientation;
end;

procedure UpDown_SetOrientation(AObj: TUpDown; AValue: TUDOrientation); cdecl;
begin
  AObj.Orientation := AValue;
end;

function UpDown_GetParentDoubleBuffered(AObj: TUpDown): LongBool; cdecl;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure UpDown_SetParentDoubleBuffered(AObj: TUpDown; AValue: LongBool); cdecl;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function UpDown_GetParentShowHint(AObj: TUpDown): LongBool; cdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure UpDown_SetParentShowHint(AObj: TUpDown; AValue: LongBool); cdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function UpDown_GetPopupMenu(AObj: TUpDown): TPopupMenu; cdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure UpDown_SetPopupMenu(AObj: TUpDown; AValue: TPopupMenu); cdecl;
begin
  AObj.PopupMenu := AValue;
end;

function UpDown_GetPosition(AObj: TUpDown): Integer; cdecl;
begin
  Result :=  AObj.Position;
end;

procedure UpDown_SetPosition(AObj: TUpDown; AValue: Integer); cdecl;
begin
  AObj.Position := AValue;
end;

function UpDown_GetShowHint(AObj: TUpDown): LongBool; cdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure UpDown_SetShowHint(AObj: TUpDown; AValue: LongBool); cdecl;
begin
  AObj.ShowHint := AValue;
end;

function UpDown_GetTabOrder(AObj: TUpDown): TTabOrder; cdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure UpDown_SetTabOrder(AObj: TUpDown; AValue: TTabOrder); cdecl;
begin
  AObj.TabOrder := AValue;
end;

function UpDown_GetTabStop(AObj: TUpDown): LongBool; cdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure UpDown_SetTabStop(AObj: TUpDown; AValue: LongBool); cdecl;
begin
  AObj.TabStop := AValue;
end;

function UpDown_GetVisible(AObj: TUpDown): LongBool; cdecl;
begin
  Result :=  AObj.Visible;
end;

procedure UpDown_SetVisible(AObj: TUpDown; AValue: LongBool); cdecl;
begin
  AObj.Visible := AValue;
end;

function UpDown_GetWrap(AObj: TUpDown): LongBool; cdecl;
begin
  Result :=  AObj.Wrap;
end;

procedure UpDown_SetWrap(AObj: TUpDown; AValue: LongBool); cdecl;
begin
  AObj.Wrap := AValue;
end;

function UpDown_GetStyleElements(AObj: TUpDown): TStyleElements; cdecl;
begin
  Result :=  AObj.StyleElements;
end;

procedure UpDown_SetStyleElements(AObj: TUpDown; AValue: TStyleElements); cdecl;
begin
  AObj.StyleElements := AValue;
end;

procedure UpDown_SetOnClick(AObj: TUpDown; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnClick := nil;
    TEventClass.Remove(AObj, geUpDownClick);
    Exit;
  end;
  AObj.OnClick := TEventClass.UpDownOnClick;
  TEventClass.Add(AObj, geUpDownClick, AEventId);
end;

procedure UpDown_SetOnEnter(AObj: TUpDown; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnEnter := nil;
    TEventClass.Remove(AObj, geEnter);
    Exit;
  end;
  AObj.OnEnter := TEventClass.OnEnter;
  TEventClass.Add(AObj, geEnter, AEventId);
end;

procedure UpDown_SetOnExit(AObj: TUpDown; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnExit := nil;
    TEventClass.Remove(AObj, geExit);
    Exit;
  end;
  AObj.OnExit := TEventClass.OnExit;
  TEventClass.Add(AObj, geExit, AEventId);
end;

procedure UpDown_SetOnMouseDown(AObj: TUpDown; AEventId: NativeUInt); stdcall
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

procedure UpDown_SetOnMouseEnter(AObj: TUpDown; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseEnter := nil;
    TEventClass.Remove(AObj, geMouseEnter);
    Exit;
  end;
  AObj.OnMouseEnter := TEventClass.OnMouseEnter;
  TEventClass.Add(AObj, geMouseEnter, AEventId);
end;

procedure UpDown_SetOnMouseLeave(AObj: TUpDown; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnMouseLeave := nil;
    TEventClass.Remove(AObj, geMouseLeave);
    Exit;
  end;
  AObj.OnMouseLeave := TEventClass.OnMouseLeave;
  TEventClass.Add(AObj, geMouseLeave, AEventId);
end;

procedure UpDown_SetOnMouseMove(AObj: TUpDown; AEventId: NativeUInt); stdcall
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

procedure UpDown_SetOnMouseUp(AObj: TUpDown; AEventId: NativeUInt); stdcall
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

function UpDown_GetBrush(AObj: TUpDown): TBrush; cdecl;
begin
  Result :=  AObj.Brush;
end;

function UpDown_GetControlCount(AObj: TUpDown): Integer; cdecl;
begin
  Result :=  AObj.ControlCount;
end;

function UpDown_GetHandle(AObj: TUpDown): HWND; cdecl;
begin
  Result :=  AObj.Handle;
end;

function UpDown_GetParentWindow(AObj: TUpDown): HWND; cdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure UpDown_SetParentWindow(AObj: TUpDown; AValue: HWND); cdecl;
begin
  AObj.ParentWindow := AValue;
end;

function UpDown_GetAction(AObj: TUpDown): TBasicAction; cdecl;
begin
  Result :=  AObj.Action;
end;

procedure UpDown_SetAction(AObj: TUpDown; AValue: TBasicAction); cdecl;
begin
  AObj.Action := AValue;
end;

function UpDown_GetAlign(AObj: TUpDown): TAlign; cdecl;
begin
  Result :=  AObj.Align;
end;

procedure UpDown_SetAlign(AObj: TUpDown; AValue: TAlign); cdecl;
begin
  AObj.Align := AValue;
end;

function UpDown_GetBiDiMode(AObj: TUpDown): TBiDiMode; cdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure UpDown_SetBiDiMode(AObj: TUpDown; AValue: TBiDiMode); cdecl;
begin
  AObj.BiDiMode := AValue;
end;

procedure UpDown_GetBoundsRect(AObj: TUpDown; var Result: TRect); cdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure UpDown_SetBoundsRect(AObj: TUpDown; var AValue: TRect); cdecl;
begin
  AObj.BoundsRect := AValue;
end;

function UpDown_GetClientHeight(AObj: TUpDown): Integer; cdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure UpDown_SetClientHeight(AObj: TUpDown; AValue: Integer); cdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure UpDown_GetClientRect(AObj: TUpDown; var Result: TRect); cdecl;
begin
  Result :=  AObj.ClientRect;
end;

function UpDown_GetClientWidth(AObj: TUpDown): Integer; cdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure UpDown_SetClientWidth(AObj: TUpDown; AValue: Integer); cdecl;
begin
  AObj.ClientWidth := AValue;
end;

function UpDown_GetExplicitLeft(AObj: TUpDown): Integer; cdecl;
begin
  Result :=  AObj.ExplicitLeft;
end;

function UpDown_GetExplicitTop(AObj: TUpDown): Integer; cdecl;
begin
  Result :=  AObj.ExplicitTop;
end;

function UpDown_GetExplicitWidth(AObj: TUpDown): Integer; cdecl;
begin
  Result :=  AObj.ExplicitWidth;
end;

function UpDown_GetExplicitHeight(AObj: TUpDown): Integer; cdecl;
begin
  Result :=  AObj.ExplicitHeight;
end;

function UpDown_GetParent(AObj: TUpDown): TWinControl; cdecl;
begin
  Result :=  AObj.Parent;
end;

procedure UpDown_SetParent(AObj: TUpDown; AValue: TWinControl); cdecl;
begin
  AObj.Parent := AValue;
end;

function UpDown_GetAlignWithMargins(AObj: TUpDown): LongBool; cdecl;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure UpDown_SetAlignWithMargins(AObj: TUpDown; AValue: LongBool); cdecl;
begin
  AObj.AlignWithMargins := AValue;
end;

function UpDown_GetLeft(AObj: TUpDown): Integer; cdecl;
begin
  Result :=  AObj.Left;
end;

procedure UpDown_SetLeft(AObj: TUpDown; AValue: Integer); cdecl;
begin
  AObj.Left := AValue;
end;

function UpDown_GetTop(AObj: TUpDown): Integer; cdecl;
begin
  Result :=  AObj.Top;
end;

procedure UpDown_SetTop(AObj: TUpDown; AValue: Integer); cdecl;
begin
  AObj.Top := AValue;
end;

function UpDown_GetWidth(AObj: TUpDown): Integer; cdecl;
begin
  Result :=  AObj.Width;
end;

procedure UpDown_SetWidth(AObj: TUpDown; AValue: Integer); cdecl;
begin
  AObj.Width := AValue;
end;

function UpDown_GetHeight(AObj: TUpDown): Integer; cdecl;
begin
  Result :=  AObj.Height;
end;

procedure UpDown_SetHeight(AObj: TUpDown; AValue: Integer); cdecl;
begin
  AObj.Height := AValue;
end;

function UpDown_GetCursor(AObj: TUpDown): TCursor; cdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure UpDown_SetCursor(AObj: TUpDown; AValue: TCursor); cdecl;
begin
  AObj.Cursor := AValue;
end;

function UpDown_GetMargins(AObj: TUpDown): TMargins; cdecl;
begin
  Result :=  AObj.Margins;
end;

procedure UpDown_SetMargins(AObj: TUpDown; AValue: TMargins); cdecl;
begin
  AObj.Margins := AValue;
end;

function UpDown_GetComponentCount(AObj: TUpDown): Integer; cdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function UpDown_GetComponentIndex(AObj: TUpDown): Integer; cdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure UpDown_SetComponentIndex(AObj: TUpDown; AValue: Integer); cdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function UpDown_GetOwner(AObj: TUpDown): TComponent; cdecl;
begin
  Result :=  AObj.Owner;
end;

function UpDown_GetName(AObj: TUpDown): PChar; cdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure UpDown_SetName(AObj: TUpDown; AValue: PChar); cdecl;
begin
  AObj.Name := AValue;
end;

function UpDown_GetTag(AObj: TUpDown): NativeInt; cdecl;
begin
  Result :=  AObj.Tag;
end;

procedure UpDown_SetTag(AObj: TUpDown; AValue: NativeInt); cdecl;
begin
  AObj.Tag := AValue;
end;

function UpDown_GetControls(AObj: TUpDown; Index: Integer): TControl; cdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function UpDown_GetComponents(AObj: TUpDown; AIndex: Integer): TComponent; cdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  UpDown_Create {$IFNDEF MSWINDOWS}name '_UpDown_Create'{$ENDIF},
  UpDown_Free {$IFNDEF MSWINDOWS}name '_UpDown_Free'{$ENDIF},
  UpDown_CanFocus {$IFNDEF MSWINDOWS}name '_UpDown_CanFocus'{$ENDIF},
  UpDown_FlipChildren {$IFNDEF MSWINDOWS}name '_UpDown_FlipChildren'{$ENDIF},
  UpDown_Focused {$IFNDEF MSWINDOWS}name '_UpDown_Focused'{$ENDIF},
  UpDown_HandleAllocated {$IFNDEF MSWINDOWS}name '_UpDown_HandleAllocated'{$ENDIF},
  UpDown_Invalidate {$IFNDEF MSWINDOWS}name '_UpDown_Invalidate'{$ENDIF},
  UpDown_Realign {$IFNDEF MSWINDOWS}name '_UpDown_Realign'{$ENDIF},
  UpDown_Repaint {$IFNDEF MSWINDOWS}name '_UpDown_Repaint'{$ENDIF},
  UpDown_ScaleBy {$IFNDEF MSWINDOWS}name '_UpDown_ScaleBy'{$ENDIF},
  UpDown_SetBounds {$IFNDEF MSWINDOWS}name '_UpDown_SetBounds'{$ENDIF},
  UpDown_SetFocus {$IFNDEF MSWINDOWS}name '_UpDown_SetFocus'{$ENDIF},
  UpDown_Update {$IFNDEF MSWINDOWS}name '_UpDown_Update'{$ENDIF},
  UpDown_BringToFront {$IFNDEF MSWINDOWS}name '_UpDown_BringToFront'{$ENDIF},
  UpDown_HasParent {$IFNDEF MSWINDOWS}name '_UpDown_HasParent'{$ENDIF},
  UpDown_Hide {$IFNDEF MSWINDOWS}name '_UpDown_Hide'{$ENDIF},
  UpDown_Perform {$IFNDEF MSWINDOWS}name '_UpDown_Perform'{$ENDIF},
  UpDown_Refresh {$IFNDEF MSWINDOWS}name '_UpDown_Refresh'{$ENDIF},
  UpDown_SendToBack {$IFNDEF MSWINDOWS}name '_UpDown_SendToBack'{$ENDIF},
  UpDown_Show {$IFNDEF MSWINDOWS}name '_UpDown_Show'{$ENDIF},
  UpDown_GetTextBuf {$IFNDEF MSWINDOWS}name '_UpDown_GetTextBuf'{$ENDIF},
  UpDown_FindComponent {$IFNDEF MSWINDOWS}name '_UpDown_FindComponent'{$ENDIF},
  UpDown_GetNamePath {$IFNDEF MSWINDOWS}name '_UpDown_GetNamePath'{$ENDIF},
  UpDown_Assign {$IFNDEF MSWINDOWS}name '_UpDown_Assign'{$ENDIF},
  UpDown_ClassName {$IFNDEF MSWINDOWS}name '_UpDown_ClassName'{$ENDIF},
  UpDown_Equals {$IFNDEF MSWINDOWS}name '_UpDown_Equals'{$ENDIF},
  UpDown_GetHashCode {$IFNDEF MSWINDOWS}name '_UpDown_GetHashCode'{$ENDIF},
  UpDown_ToString {$IFNDEF MSWINDOWS}name '_UpDown_ToString'{$ENDIF},
  UpDown_GetAnchors {$IFNDEF MSWINDOWS}name '_UpDown_GetAnchors'{$ENDIF},
  UpDown_SetAnchors {$IFNDEF MSWINDOWS}name '_UpDown_SetAnchors'{$ENDIF},
  UpDown_GetDoubleBuffered {$IFNDEF MSWINDOWS}name '_UpDown_GetDoubleBuffered'{$ENDIF},
  UpDown_SetDoubleBuffered {$IFNDEF MSWINDOWS}name '_UpDown_SetDoubleBuffered'{$ENDIF},
  UpDown_GetEnabled {$IFNDEF MSWINDOWS}name '_UpDown_GetEnabled'{$ENDIF},
  UpDown_SetEnabled {$IFNDEF MSWINDOWS}name '_UpDown_SetEnabled'{$ENDIF},
  UpDown_GetHint {$IFNDEF MSWINDOWS}name '_UpDown_GetHint'{$ENDIF},
  UpDown_SetHint {$IFNDEF MSWINDOWS}name '_UpDown_SetHint'{$ENDIF},
  UpDown_GetMin {$IFNDEF MSWINDOWS}name '_UpDown_GetMin'{$ENDIF},
  UpDown_SetMin {$IFNDEF MSWINDOWS}name '_UpDown_SetMin'{$ENDIF},
  UpDown_GetMax {$IFNDEF MSWINDOWS}name '_UpDown_GetMax'{$ENDIF},
  UpDown_SetMax {$IFNDEF MSWINDOWS}name '_UpDown_SetMax'{$ENDIF},
  UpDown_GetOrientation {$IFNDEF MSWINDOWS}name '_UpDown_GetOrientation'{$ENDIF},
  UpDown_SetOrientation {$IFNDEF MSWINDOWS}name '_UpDown_SetOrientation'{$ENDIF},
  UpDown_GetParentDoubleBuffered {$IFNDEF MSWINDOWS}name '_UpDown_GetParentDoubleBuffered'{$ENDIF},
  UpDown_SetParentDoubleBuffered {$IFNDEF MSWINDOWS}name '_UpDown_SetParentDoubleBuffered'{$ENDIF},
  UpDown_GetParentShowHint {$IFNDEF MSWINDOWS}name '_UpDown_GetParentShowHint'{$ENDIF},
  UpDown_SetParentShowHint {$IFNDEF MSWINDOWS}name '_UpDown_SetParentShowHint'{$ENDIF},
  UpDown_GetPopupMenu {$IFNDEF MSWINDOWS}name '_UpDown_GetPopupMenu'{$ENDIF},
  UpDown_SetPopupMenu {$IFNDEF MSWINDOWS}name '_UpDown_SetPopupMenu'{$ENDIF},
  UpDown_GetPosition {$IFNDEF MSWINDOWS}name '_UpDown_GetPosition'{$ENDIF},
  UpDown_SetPosition {$IFNDEF MSWINDOWS}name '_UpDown_SetPosition'{$ENDIF},
  UpDown_GetShowHint {$IFNDEF MSWINDOWS}name '_UpDown_GetShowHint'{$ENDIF},
  UpDown_SetShowHint {$IFNDEF MSWINDOWS}name '_UpDown_SetShowHint'{$ENDIF},
  UpDown_GetTabOrder {$IFNDEF MSWINDOWS}name '_UpDown_GetTabOrder'{$ENDIF},
  UpDown_SetTabOrder {$IFNDEF MSWINDOWS}name '_UpDown_SetTabOrder'{$ENDIF},
  UpDown_GetTabStop {$IFNDEF MSWINDOWS}name '_UpDown_GetTabStop'{$ENDIF},
  UpDown_SetTabStop {$IFNDEF MSWINDOWS}name '_UpDown_SetTabStop'{$ENDIF},
  UpDown_GetVisible {$IFNDEF MSWINDOWS}name '_UpDown_GetVisible'{$ENDIF},
  UpDown_SetVisible {$IFNDEF MSWINDOWS}name '_UpDown_SetVisible'{$ENDIF},
  UpDown_GetWrap {$IFNDEF MSWINDOWS}name '_UpDown_GetWrap'{$ENDIF},
  UpDown_SetWrap {$IFNDEF MSWINDOWS}name '_UpDown_SetWrap'{$ENDIF},
  UpDown_GetStyleElements {$IFNDEF MSWINDOWS}name '_UpDown_GetStyleElements'{$ENDIF},
  UpDown_SetStyleElements {$IFNDEF MSWINDOWS}name '_UpDown_SetStyleElements'{$ENDIF},
  UpDown_SetOnClick {$IFNDEF MSWINDOWS}name '_UpDown_SetOnClick'{$ENDIF},
  UpDown_SetOnEnter {$IFNDEF MSWINDOWS}name '_UpDown_SetOnEnter'{$ENDIF},
  UpDown_SetOnExit {$IFNDEF MSWINDOWS}name '_UpDown_SetOnExit'{$ENDIF},
  UpDown_SetOnMouseDown {$IFNDEF MSWINDOWS}name '_UpDown_SetOnMouseDown'{$ENDIF},
  UpDown_SetOnMouseEnter {$IFNDEF MSWINDOWS}name '_UpDown_SetOnMouseEnter'{$ENDIF},
  UpDown_SetOnMouseLeave {$IFNDEF MSWINDOWS}name '_UpDown_SetOnMouseLeave'{$ENDIF},
  UpDown_SetOnMouseMove {$IFNDEF MSWINDOWS}name '_UpDown_SetOnMouseMove'{$ENDIF},
  UpDown_SetOnMouseUp {$IFNDEF MSWINDOWS}name '_UpDown_SetOnMouseUp'{$ENDIF},
  UpDown_GetBrush {$IFNDEF MSWINDOWS}name '_UpDown_GetBrush'{$ENDIF},
  UpDown_GetControlCount {$IFNDEF MSWINDOWS}name '_UpDown_GetControlCount'{$ENDIF},
  UpDown_GetHandle {$IFNDEF MSWINDOWS}name '_UpDown_GetHandle'{$ENDIF},
  UpDown_GetParentWindow {$IFNDEF MSWINDOWS}name '_UpDown_GetParentWindow'{$ENDIF},
  UpDown_SetParentWindow {$IFNDEF MSWINDOWS}name '_UpDown_SetParentWindow'{$ENDIF},
  UpDown_GetAction {$IFNDEF MSWINDOWS}name '_UpDown_GetAction'{$ENDIF},
  UpDown_SetAction {$IFNDEF MSWINDOWS}name '_UpDown_SetAction'{$ENDIF},
  UpDown_GetAlign {$IFNDEF MSWINDOWS}name '_UpDown_GetAlign'{$ENDIF},
  UpDown_SetAlign {$IFNDEF MSWINDOWS}name '_UpDown_SetAlign'{$ENDIF},
  UpDown_GetBiDiMode {$IFNDEF MSWINDOWS}name '_UpDown_GetBiDiMode'{$ENDIF},
  UpDown_SetBiDiMode {$IFNDEF MSWINDOWS}name '_UpDown_SetBiDiMode'{$ENDIF},
  UpDown_GetBoundsRect {$IFNDEF MSWINDOWS}name '_UpDown_GetBoundsRect'{$ENDIF},
  UpDown_SetBoundsRect {$IFNDEF MSWINDOWS}name '_UpDown_SetBoundsRect'{$ENDIF},
  UpDown_GetClientHeight {$IFNDEF MSWINDOWS}name '_UpDown_GetClientHeight'{$ENDIF},
  UpDown_SetClientHeight {$IFNDEF MSWINDOWS}name '_UpDown_SetClientHeight'{$ENDIF},
  UpDown_GetClientRect {$IFNDEF MSWINDOWS}name '_UpDown_GetClientRect'{$ENDIF},
  UpDown_GetClientWidth {$IFNDEF MSWINDOWS}name '_UpDown_GetClientWidth'{$ENDIF},
  UpDown_SetClientWidth {$IFNDEF MSWINDOWS}name '_UpDown_SetClientWidth'{$ENDIF},
  UpDown_GetExplicitLeft {$IFNDEF MSWINDOWS}name '_UpDown_GetExplicitLeft'{$ENDIF},
  UpDown_GetExplicitTop {$IFNDEF MSWINDOWS}name '_UpDown_GetExplicitTop'{$ENDIF},
  UpDown_GetExplicitWidth {$IFNDEF MSWINDOWS}name '_UpDown_GetExplicitWidth'{$ENDIF},
  UpDown_GetExplicitHeight {$IFNDEF MSWINDOWS}name '_UpDown_GetExplicitHeight'{$ENDIF},
  UpDown_GetParent {$IFNDEF MSWINDOWS}name '_UpDown_GetParent'{$ENDIF},
  UpDown_SetParent {$IFNDEF MSWINDOWS}name '_UpDown_SetParent'{$ENDIF},
  UpDown_GetAlignWithMargins {$IFNDEF MSWINDOWS}name '_UpDown_GetAlignWithMargins'{$ENDIF},
  UpDown_SetAlignWithMargins {$IFNDEF MSWINDOWS}name '_UpDown_SetAlignWithMargins'{$ENDIF},
  UpDown_GetLeft {$IFNDEF MSWINDOWS}name '_UpDown_GetLeft'{$ENDIF},
  UpDown_SetLeft {$IFNDEF MSWINDOWS}name '_UpDown_SetLeft'{$ENDIF},
  UpDown_GetTop {$IFNDEF MSWINDOWS}name '_UpDown_GetTop'{$ENDIF},
  UpDown_SetTop {$IFNDEF MSWINDOWS}name '_UpDown_SetTop'{$ENDIF},
  UpDown_GetWidth {$IFNDEF MSWINDOWS}name '_UpDown_GetWidth'{$ENDIF},
  UpDown_SetWidth {$IFNDEF MSWINDOWS}name '_UpDown_SetWidth'{$ENDIF},
  UpDown_GetHeight {$IFNDEF MSWINDOWS}name '_UpDown_GetHeight'{$ENDIF},
  UpDown_SetHeight {$IFNDEF MSWINDOWS}name '_UpDown_SetHeight'{$ENDIF},
  UpDown_GetCursor {$IFNDEF MSWINDOWS}name '_UpDown_GetCursor'{$ENDIF},
  UpDown_SetCursor {$IFNDEF MSWINDOWS}name '_UpDown_SetCursor'{$ENDIF},
  UpDown_GetMargins {$IFNDEF MSWINDOWS}name '_UpDown_GetMargins'{$ENDIF},
  UpDown_SetMargins {$IFNDEF MSWINDOWS}name '_UpDown_SetMargins'{$ENDIF},
  UpDown_GetComponentCount {$IFNDEF MSWINDOWS}name '_UpDown_GetComponentCount'{$ENDIF},
  UpDown_GetComponentIndex {$IFNDEF MSWINDOWS}name '_UpDown_GetComponentIndex'{$ENDIF},
  UpDown_SetComponentIndex {$IFNDEF MSWINDOWS}name '_UpDown_SetComponentIndex'{$ENDIF},
  UpDown_GetOwner {$IFNDEF MSWINDOWS}name '_UpDown_GetOwner'{$ENDIF},
  UpDown_GetName {$IFNDEF MSWINDOWS}name '_UpDown_GetName'{$ENDIF},
  UpDown_SetName {$IFNDEF MSWINDOWS}name '_UpDown_SetName'{$ENDIF},
  UpDown_GetTag {$IFNDEF MSWINDOWS}name '_UpDown_GetTag'{$ENDIF},
  UpDown_SetTag {$IFNDEF MSWINDOWS}name '_UpDown_SetTag'{$ENDIF},
  UpDown_GetControls {$IFNDEF MSWINDOWS}name '_UpDown_GetControls'{$ENDIF},
  UpDown_GetComponents {$IFNDEF MSWINDOWS}name '_UpDown_GetComponents'{$ENDIF};

