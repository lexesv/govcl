
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function TrackBar_Create(AOwner: TComponent): TTrackBar; cdecl;
begin
  Result :=  TTrackBar.Create(AOwner);
end;

procedure TrackBar_Free(AObj: TTrackBar); cdecl;
begin
  AObj.Free;
end;

procedure TrackBar_SetTick(AObj: TTrackBar; Value: Integer); cdecl;
begin
  AObj.SetTick(Value);
end;

function TrackBar_CanFocus(AObj: TTrackBar): LongBool; cdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure TrackBar_FlipChildren(AObj: TTrackBar; AllLevels: LongBool); cdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function TrackBar_Focused(AObj: TTrackBar): LongBool; cdecl;
begin
  Result :=  AObj.Focused;
end;

function TrackBar_HandleAllocated(AObj: TTrackBar): LongBool; cdecl;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure TrackBar_Invalidate(AObj: TTrackBar); cdecl;
begin
  AObj.Invalidate;
end;

procedure TrackBar_Realign(AObj: TTrackBar); cdecl;
begin
  AObj.Realign;
end;

procedure TrackBar_Repaint(AObj: TTrackBar); cdecl;
begin
  AObj.Repaint;
end;

procedure TrackBar_ScaleBy(AObj: TTrackBar; M: Integer; D: Integer); cdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure TrackBar_SetBounds(AObj: TTrackBar; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); cdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure TrackBar_SetFocus(AObj: TTrackBar); cdecl;
begin
  AObj.SetFocus;
end;

procedure TrackBar_Update(AObj: TTrackBar); cdecl;
begin
  AObj.Update;
end;

procedure TrackBar_BringToFront(AObj: TTrackBar); cdecl;
begin
  AObj.BringToFront;
end;

function TrackBar_HasParent(AObj: TTrackBar): LongBool; cdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure TrackBar_Hide(AObj: TTrackBar); cdecl;
begin
  AObj.Hide;
end;

function TrackBar_Perform(AObj: TTrackBar; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; cdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure TrackBar_Refresh(AObj: TTrackBar); cdecl;
begin
  AObj.Refresh;
end;

procedure TrackBar_SendToBack(AObj: TTrackBar); cdecl;
begin
  AObj.SendToBack;
end;

procedure TrackBar_Show(AObj: TTrackBar); cdecl;
begin
  AObj.Show;
end;

function TrackBar_GetTextBuf(AObj: TTrackBar; Buffer: PWideChar; BufSize: Integer): Integer; cdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function TrackBar_FindComponent(AObj: TTrackBar; AName: PChar): TComponent; cdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function TrackBar_GetNamePath(AObj: TTrackBar): PChar; cdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure TrackBar_Assign(AObj: TTrackBar; Source: TPersistent); cdecl;
begin
  AObj.Assign(Source);
end;

function TrackBar_ClassName(AObj: TTrackBar): PChar; cdecl;
begin
  Result :=  PChar(AObj.ClassName);
end;

function TrackBar_Equals(AObj: TTrackBar; Obj: TObject): LongBool; cdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function TrackBar_GetHashCode(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function TrackBar_ToString(AObj: TTrackBar): PChar; cdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function TrackBar_GetAlign(AObj: TTrackBar): TAlign; cdecl;
begin
  Result :=  AObj.Align;
end;

procedure TrackBar_SetAlign(AObj: TTrackBar; AValue: TAlign); cdecl;
begin
  AObj.Align := AValue;
end;

function TrackBar_GetAnchors(AObj: TTrackBar): TAnchors; cdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure TrackBar_SetAnchors(AObj: TTrackBar; AValue: TAnchors); cdecl;
begin
  AObj.Anchors := AValue;
end;

function TrackBar_GetBorderWidth(AObj: TTrackBar): TBorderWidth; cdecl;
begin
  Result :=  AObj.BorderWidth;
end;

procedure TrackBar_SetBorderWidth(AObj: TTrackBar; AValue: TBorderWidth); cdecl;
begin
  AObj.BorderWidth := AValue;
end;

function TrackBar_GetDoubleBuffered(AObj: TTrackBar): LongBool; cdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure TrackBar_SetDoubleBuffered(AObj: TTrackBar; AValue: LongBool); cdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function TrackBar_GetEnabled(AObj: TTrackBar): LongBool; cdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure TrackBar_SetEnabled(AObj: TTrackBar; AValue: LongBool); cdecl;
begin
  AObj.Enabled := AValue;
end;

function TrackBar_GetLineSize(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.LineSize;
end;

procedure TrackBar_SetLineSize(AObj: TTrackBar; AValue: Integer); cdecl;
begin
  AObj.LineSize := AValue;
end;

function TrackBar_GetMax(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.Max;
end;

procedure TrackBar_SetMax(AObj: TTrackBar; AValue: Integer); cdecl;
begin
  AObj.Max := AValue;
end;

function TrackBar_GetMin(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.Min;
end;

procedure TrackBar_SetMin(AObj: TTrackBar; AValue: Integer); cdecl;
begin
  AObj.Min := AValue;
end;

function TrackBar_GetOrientation(AObj: TTrackBar): TTrackBarOrientation; cdecl;
begin
  Result :=  AObj.Orientation;
end;

procedure TrackBar_SetOrientation(AObj: TTrackBar; AValue: TTrackBarOrientation); cdecl;
begin
  AObj.Orientation := AValue;
end;

function TrackBar_GetParentCtl3D(AObj: TTrackBar): LongBool; cdecl;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure TrackBar_SetParentCtl3D(AObj: TTrackBar; AValue: LongBool); cdecl;
begin
  AObj.ParentCtl3D := AValue;
end;

function TrackBar_GetParentDoubleBuffered(AObj: TTrackBar): LongBool; cdecl;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure TrackBar_SetParentDoubleBuffered(AObj: TTrackBar; AValue: LongBool); cdecl;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function TrackBar_GetParentShowHint(AObj: TTrackBar): LongBool; cdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure TrackBar_SetParentShowHint(AObj: TTrackBar; AValue: LongBool); cdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function TrackBar_GetPageSize(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.PageSize;
end;

procedure TrackBar_SetPageSize(AObj: TTrackBar; AValue: Integer); cdecl;
begin
  AObj.PageSize := AValue;
end;

function TrackBar_GetPopupMenu(AObj: TTrackBar): TPopupMenu; cdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure TrackBar_SetPopupMenu(AObj: TTrackBar; AValue: TPopupMenu); cdecl;
begin
  AObj.PopupMenu := AValue;
end;

function TrackBar_GetFrequency(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.Frequency;
end;

procedure TrackBar_SetFrequency(AObj: TTrackBar; AValue: Integer); cdecl;
begin
  AObj.Frequency := AValue;
end;

function TrackBar_GetPosition(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.Position;
end;

procedure TrackBar_SetPosition(AObj: TTrackBar; AValue: Integer); cdecl;
begin
  AObj.Position := AValue;
end;

function TrackBar_GetPositionToolTip(AObj: TTrackBar): TPositionToolTip; cdecl;
begin
  Result :=  AObj.PositionToolTip;
end;

procedure TrackBar_SetPositionToolTip(AObj: TTrackBar; AValue: TPositionToolTip); cdecl;
begin
  AObj.PositionToolTip := AValue;
end;

function TrackBar_GetSliderVisible(AObj: TTrackBar): LongBool; cdecl;
begin
  Result :=  AObj.SliderVisible;
end;

procedure TrackBar_SetSliderVisible(AObj: TTrackBar; AValue: LongBool); cdecl;
begin
  AObj.SliderVisible := AValue;
end;

function TrackBar_GetSelEnd(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.SelEnd;
end;

procedure TrackBar_SetSelEnd(AObj: TTrackBar; AValue: Integer); cdecl;
begin
  AObj.SelEnd := AValue;
end;

function TrackBar_GetSelStart(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.SelStart;
end;

procedure TrackBar_SetSelStart(AObj: TTrackBar; AValue: Integer); cdecl;
begin
  AObj.SelStart := AValue;
end;

function TrackBar_GetShowHint(AObj: TTrackBar): LongBool; cdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure TrackBar_SetShowHint(AObj: TTrackBar; AValue: LongBool); cdecl;
begin
  AObj.ShowHint := AValue;
end;

function TrackBar_GetShowSelRange(AObj: TTrackBar): LongBool; cdecl;
begin
  Result :=  AObj.ShowSelRange;
end;

procedure TrackBar_SetShowSelRange(AObj: TTrackBar; AValue: LongBool); cdecl;
begin
  AObj.ShowSelRange := AValue;
end;

function TrackBar_GetTabOrder(AObj: TTrackBar): TTabOrder; cdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure TrackBar_SetTabOrder(AObj: TTrackBar; AValue: TTabOrder); cdecl;
begin
  AObj.TabOrder := AValue;
end;

function TrackBar_GetTabStop(AObj: TTrackBar): LongBool; cdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure TrackBar_SetTabStop(AObj: TTrackBar; AValue: LongBool); cdecl;
begin
  AObj.TabStop := AValue;
end;

function TrackBar_GetThumbLength(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.ThumbLength;
end;

procedure TrackBar_SetThumbLength(AObj: TTrackBar; AValue: Integer); cdecl;
begin
  AObj.ThumbLength := AValue;
end;

function TrackBar_GetTickMarks(AObj: TTrackBar): TTickMark; cdecl;
begin
  Result :=  AObj.TickMarks;
end;

procedure TrackBar_SetTickMarks(AObj: TTrackBar; AValue: TTickMark); cdecl;
begin
  AObj.TickMarks := AValue;
end;

function TrackBar_GetTickStyle(AObj: TTrackBar): TTickStyle; cdecl;
begin
  Result :=  AObj.TickStyle;
end;

procedure TrackBar_SetTickStyle(AObj: TTrackBar; AValue: TTickStyle); cdecl;
begin
  AObj.TickStyle := AValue;
end;

function TrackBar_GetVisible(AObj: TTrackBar): LongBool; cdecl;
begin
  Result :=  AObj.Visible;
end;

procedure TrackBar_SetVisible(AObj: TTrackBar; AValue: LongBool); cdecl;
begin
  AObj.Visible := AValue;
end;

function TrackBar_GetStyleElements(AObj: TTrackBar): TStyleElements; cdecl;
begin
  Result :=  AObj.StyleElements;
end;

procedure TrackBar_SetStyleElements(AObj: TTrackBar; AValue: TStyleElements); cdecl;
begin
  AObj.StyleElements := AValue;
end;

procedure TrackBar_SetOnChange(AObj: TTrackBar; AEventId: NativeUInt); stdcall
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

procedure TrackBar_SetOnEnter(AObj: TTrackBar; AEventId: NativeUInt); stdcall
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

procedure TrackBar_SetOnExit(AObj: TTrackBar; AEventId: NativeUInt); stdcall
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

procedure TrackBar_SetOnKeyDown(AObj: TTrackBar; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyDown := nil;
    TEventClass.Remove(AObj, geKeyDown);
    Exit;
  end;
  AObj.OnKeyDown := TEventClass.OnKeyDown;
  TEventClass.Add(AObj, geKeyDown, AEventId);
end;

procedure TrackBar_SetOnKeyPress(AObj: TTrackBar; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyPress := nil;
    TEventClass.Remove(AObj, geKeyPress);
    Exit;
  end;
  AObj.OnKeyPress := TEventClass.OnKeyPress;
  TEventClass.Add(AObj, geKeyPress, AEventId);
end;

procedure TrackBar_SetOnKeyUp(AObj: TTrackBar; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnKeyUp := nil;
    TEventClass.Remove(AObj, geKeyUp);
    Exit;
  end;
  AObj.OnKeyUp := TEventClass.OnKeyUp;
  TEventClass.Add(AObj, geKeyUp, AEventId);
end;

function TrackBar_GetBrush(AObj: TTrackBar): TBrush; cdecl;
begin
  Result :=  AObj.Brush;
end;

function TrackBar_GetControlCount(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.ControlCount;
end;

function TrackBar_GetHandle(AObj: TTrackBar): HWND; cdecl;
begin
  Result :=  AObj.Handle;
end;

function TrackBar_GetParentWindow(AObj: TTrackBar): HWND; cdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure TrackBar_SetParentWindow(AObj: TTrackBar; AValue: HWND); cdecl;
begin
  AObj.ParentWindow := AValue;
end;

function TrackBar_GetAction(AObj: TTrackBar): TBasicAction; cdecl;
begin
  Result :=  AObj.Action;
end;

procedure TrackBar_SetAction(AObj: TTrackBar; AValue: TBasicAction); cdecl;
begin
  AObj.Action := AValue;
end;

function TrackBar_GetBiDiMode(AObj: TTrackBar): TBiDiMode; cdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure TrackBar_SetBiDiMode(AObj: TTrackBar; AValue: TBiDiMode); cdecl;
begin
  AObj.BiDiMode := AValue;
end;

procedure TrackBar_GetBoundsRect(AObj: TTrackBar; var Result: TRect); cdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure TrackBar_SetBoundsRect(AObj: TTrackBar; var AValue: TRect); cdecl;
begin
  AObj.BoundsRect := AValue;
end;

function TrackBar_GetClientHeight(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure TrackBar_SetClientHeight(AObj: TTrackBar; AValue: Integer); cdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure TrackBar_GetClientRect(AObj: TTrackBar; var Result: TRect); cdecl;
begin
  Result :=  AObj.ClientRect;
end;

function TrackBar_GetClientWidth(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure TrackBar_SetClientWidth(AObj: TTrackBar; AValue: Integer); cdecl;
begin
  AObj.ClientWidth := AValue;
end;

function TrackBar_GetExplicitLeft(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.ExplicitLeft;
end;

function TrackBar_GetExplicitTop(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.ExplicitTop;
end;

function TrackBar_GetExplicitWidth(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.ExplicitWidth;
end;

function TrackBar_GetExplicitHeight(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.ExplicitHeight;
end;

function TrackBar_GetParent(AObj: TTrackBar): TWinControl; cdecl;
begin
  Result :=  AObj.Parent;
end;

procedure TrackBar_SetParent(AObj: TTrackBar; AValue: TWinControl); cdecl;
begin
  AObj.Parent := AValue;
end;

function TrackBar_GetAlignWithMargins(AObj: TTrackBar): LongBool; cdecl;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure TrackBar_SetAlignWithMargins(AObj: TTrackBar; AValue: LongBool); cdecl;
begin
  AObj.AlignWithMargins := AValue;
end;

function TrackBar_GetLeft(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.Left;
end;

procedure TrackBar_SetLeft(AObj: TTrackBar; AValue: Integer); cdecl;
begin
  AObj.Left := AValue;
end;

function TrackBar_GetTop(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.Top;
end;

procedure TrackBar_SetTop(AObj: TTrackBar; AValue: Integer); cdecl;
begin
  AObj.Top := AValue;
end;

function TrackBar_GetWidth(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.Width;
end;

procedure TrackBar_SetWidth(AObj: TTrackBar; AValue: Integer); cdecl;
begin
  AObj.Width := AValue;
end;

function TrackBar_GetHeight(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.Height;
end;

procedure TrackBar_SetHeight(AObj: TTrackBar; AValue: Integer); cdecl;
begin
  AObj.Height := AValue;
end;

function TrackBar_GetCursor(AObj: TTrackBar): TCursor; cdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure TrackBar_SetCursor(AObj: TTrackBar; AValue: TCursor); cdecl;
begin
  AObj.Cursor := AValue;
end;

function TrackBar_GetHint(AObj: TTrackBar): PChar; cdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure TrackBar_SetHint(AObj: TTrackBar; AValue: PChar); cdecl;
begin
  AObj.Hint := AValue;
end;

function TrackBar_GetMargins(AObj: TTrackBar): TMargins; cdecl;
begin
  Result :=  AObj.Margins;
end;

procedure TrackBar_SetMargins(AObj: TTrackBar; AValue: TMargins); cdecl;
begin
  AObj.Margins := AValue;
end;

function TrackBar_GetComponentCount(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function TrackBar_GetComponentIndex(AObj: TTrackBar): Integer; cdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure TrackBar_SetComponentIndex(AObj: TTrackBar; AValue: Integer); cdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function TrackBar_GetOwner(AObj: TTrackBar): TComponent; cdecl;
begin
  Result :=  AObj.Owner;
end;

function TrackBar_GetName(AObj: TTrackBar): PChar; cdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure TrackBar_SetName(AObj: TTrackBar; AValue: PChar); cdecl;
begin
  AObj.Name := AValue;
end;

function TrackBar_GetTag(AObj: TTrackBar): NativeInt; cdecl;
begin
  Result :=  AObj.Tag;
end;

procedure TrackBar_SetTag(AObj: TTrackBar; AValue: NativeInt); cdecl;
begin
  AObj.Tag := AValue;
end;

function TrackBar_GetControls(AObj: TTrackBar; Index: Integer): TControl; cdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function TrackBar_GetComponents(AObj: TTrackBar; AIndex: Integer): TComponent; cdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  TrackBar_Create {$IFNDEF MSWINDOWS}name '_TrackBar_Create'{$ENDIF},
  TrackBar_Free {$IFNDEF MSWINDOWS}name '_TrackBar_Free'{$ENDIF},
  TrackBar_SetTick {$IFNDEF MSWINDOWS}name '_TrackBar_SetTick'{$ENDIF},
  TrackBar_CanFocus {$IFNDEF MSWINDOWS}name '_TrackBar_CanFocus'{$ENDIF},
  TrackBar_FlipChildren {$IFNDEF MSWINDOWS}name '_TrackBar_FlipChildren'{$ENDIF},
  TrackBar_Focused {$IFNDEF MSWINDOWS}name '_TrackBar_Focused'{$ENDIF},
  TrackBar_HandleAllocated {$IFNDEF MSWINDOWS}name '_TrackBar_HandleAllocated'{$ENDIF},
  TrackBar_Invalidate {$IFNDEF MSWINDOWS}name '_TrackBar_Invalidate'{$ENDIF},
  TrackBar_Realign {$IFNDEF MSWINDOWS}name '_TrackBar_Realign'{$ENDIF},
  TrackBar_Repaint {$IFNDEF MSWINDOWS}name '_TrackBar_Repaint'{$ENDIF},
  TrackBar_ScaleBy {$IFNDEF MSWINDOWS}name '_TrackBar_ScaleBy'{$ENDIF},
  TrackBar_SetBounds {$IFNDEF MSWINDOWS}name '_TrackBar_SetBounds'{$ENDIF},
  TrackBar_SetFocus {$IFNDEF MSWINDOWS}name '_TrackBar_SetFocus'{$ENDIF},
  TrackBar_Update {$IFNDEF MSWINDOWS}name '_TrackBar_Update'{$ENDIF},
  TrackBar_BringToFront {$IFNDEF MSWINDOWS}name '_TrackBar_BringToFront'{$ENDIF},
  TrackBar_HasParent {$IFNDEF MSWINDOWS}name '_TrackBar_HasParent'{$ENDIF},
  TrackBar_Hide {$IFNDEF MSWINDOWS}name '_TrackBar_Hide'{$ENDIF},
  TrackBar_Perform {$IFNDEF MSWINDOWS}name '_TrackBar_Perform'{$ENDIF},
  TrackBar_Refresh {$IFNDEF MSWINDOWS}name '_TrackBar_Refresh'{$ENDIF},
  TrackBar_SendToBack {$IFNDEF MSWINDOWS}name '_TrackBar_SendToBack'{$ENDIF},
  TrackBar_Show {$IFNDEF MSWINDOWS}name '_TrackBar_Show'{$ENDIF},
  TrackBar_GetTextBuf {$IFNDEF MSWINDOWS}name '_TrackBar_GetTextBuf'{$ENDIF},
  TrackBar_FindComponent {$IFNDEF MSWINDOWS}name '_TrackBar_FindComponent'{$ENDIF},
  TrackBar_GetNamePath {$IFNDEF MSWINDOWS}name '_TrackBar_GetNamePath'{$ENDIF},
  TrackBar_Assign {$IFNDEF MSWINDOWS}name '_TrackBar_Assign'{$ENDIF},
  TrackBar_ClassName {$IFNDEF MSWINDOWS}name '_TrackBar_ClassName'{$ENDIF},
  TrackBar_Equals {$IFNDEF MSWINDOWS}name '_TrackBar_Equals'{$ENDIF},
  TrackBar_GetHashCode {$IFNDEF MSWINDOWS}name '_TrackBar_GetHashCode'{$ENDIF},
  TrackBar_ToString {$IFNDEF MSWINDOWS}name '_TrackBar_ToString'{$ENDIF},
  TrackBar_GetAlign {$IFNDEF MSWINDOWS}name '_TrackBar_GetAlign'{$ENDIF},
  TrackBar_SetAlign {$IFNDEF MSWINDOWS}name '_TrackBar_SetAlign'{$ENDIF},
  TrackBar_GetAnchors {$IFNDEF MSWINDOWS}name '_TrackBar_GetAnchors'{$ENDIF},
  TrackBar_SetAnchors {$IFNDEF MSWINDOWS}name '_TrackBar_SetAnchors'{$ENDIF},
  TrackBar_GetBorderWidth {$IFNDEF MSWINDOWS}name '_TrackBar_GetBorderWidth'{$ENDIF},
  TrackBar_SetBorderWidth {$IFNDEF MSWINDOWS}name '_TrackBar_SetBorderWidth'{$ENDIF},
  TrackBar_GetDoubleBuffered {$IFNDEF MSWINDOWS}name '_TrackBar_GetDoubleBuffered'{$ENDIF},
  TrackBar_SetDoubleBuffered {$IFNDEF MSWINDOWS}name '_TrackBar_SetDoubleBuffered'{$ENDIF},
  TrackBar_GetEnabled {$IFNDEF MSWINDOWS}name '_TrackBar_GetEnabled'{$ENDIF},
  TrackBar_SetEnabled {$IFNDEF MSWINDOWS}name '_TrackBar_SetEnabled'{$ENDIF},
  TrackBar_GetLineSize {$IFNDEF MSWINDOWS}name '_TrackBar_GetLineSize'{$ENDIF},
  TrackBar_SetLineSize {$IFNDEF MSWINDOWS}name '_TrackBar_SetLineSize'{$ENDIF},
  TrackBar_GetMax {$IFNDEF MSWINDOWS}name '_TrackBar_GetMax'{$ENDIF},
  TrackBar_SetMax {$IFNDEF MSWINDOWS}name '_TrackBar_SetMax'{$ENDIF},
  TrackBar_GetMin {$IFNDEF MSWINDOWS}name '_TrackBar_GetMin'{$ENDIF},
  TrackBar_SetMin {$IFNDEF MSWINDOWS}name '_TrackBar_SetMin'{$ENDIF},
  TrackBar_GetOrientation {$IFNDEF MSWINDOWS}name '_TrackBar_GetOrientation'{$ENDIF},
  TrackBar_SetOrientation {$IFNDEF MSWINDOWS}name '_TrackBar_SetOrientation'{$ENDIF},
  TrackBar_GetParentCtl3D {$IFNDEF MSWINDOWS}name '_TrackBar_GetParentCtl3D'{$ENDIF},
  TrackBar_SetParentCtl3D {$IFNDEF MSWINDOWS}name '_TrackBar_SetParentCtl3D'{$ENDIF},
  TrackBar_GetParentDoubleBuffered {$IFNDEF MSWINDOWS}name '_TrackBar_GetParentDoubleBuffered'{$ENDIF},
  TrackBar_SetParentDoubleBuffered {$IFNDEF MSWINDOWS}name '_TrackBar_SetParentDoubleBuffered'{$ENDIF},
  TrackBar_GetParentShowHint {$IFNDEF MSWINDOWS}name '_TrackBar_GetParentShowHint'{$ENDIF},
  TrackBar_SetParentShowHint {$IFNDEF MSWINDOWS}name '_TrackBar_SetParentShowHint'{$ENDIF},
  TrackBar_GetPageSize {$IFNDEF MSWINDOWS}name '_TrackBar_GetPageSize'{$ENDIF},
  TrackBar_SetPageSize {$IFNDEF MSWINDOWS}name '_TrackBar_SetPageSize'{$ENDIF},
  TrackBar_GetPopupMenu {$IFNDEF MSWINDOWS}name '_TrackBar_GetPopupMenu'{$ENDIF},
  TrackBar_SetPopupMenu {$IFNDEF MSWINDOWS}name '_TrackBar_SetPopupMenu'{$ENDIF},
  TrackBar_GetFrequency {$IFNDEF MSWINDOWS}name '_TrackBar_GetFrequency'{$ENDIF},
  TrackBar_SetFrequency {$IFNDEF MSWINDOWS}name '_TrackBar_SetFrequency'{$ENDIF},
  TrackBar_GetPosition {$IFNDEF MSWINDOWS}name '_TrackBar_GetPosition'{$ENDIF},
  TrackBar_SetPosition {$IFNDEF MSWINDOWS}name '_TrackBar_SetPosition'{$ENDIF},
  TrackBar_GetPositionToolTip {$IFNDEF MSWINDOWS}name '_TrackBar_GetPositionToolTip'{$ENDIF},
  TrackBar_SetPositionToolTip {$IFNDEF MSWINDOWS}name '_TrackBar_SetPositionToolTip'{$ENDIF},
  TrackBar_GetSliderVisible {$IFNDEF MSWINDOWS}name '_TrackBar_GetSliderVisible'{$ENDIF},
  TrackBar_SetSliderVisible {$IFNDEF MSWINDOWS}name '_TrackBar_SetSliderVisible'{$ENDIF},
  TrackBar_GetSelEnd {$IFNDEF MSWINDOWS}name '_TrackBar_GetSelEnd'{$ENDIF},
  TrackBar_SetSelEnd {$IFNDEF MSWINDOWS}name '_TrackBar_SetSelEnd'{$ENDIF},
  TrackBar_GetSelStart {$IFNDEF MSWINDOWS}name '_TrackBar_GetSelStart'{$ENDIF},
  TrackBar_SetSelStart {$IFNDEF MSWINDOWS}name '_TrackBar_SetSelStart'{$ENDIF},
  TrackBar_GetShowHint {$IFNDEF MSWINDOWS}name '_TrackBar_GetShowHint'{$ENDIF},
  TrackBar_SetShowHint {$IFNDEF MSWINDOWS}name '_TrackBar_SetShowHint'{$ENDIF},
  TrackBar_GetShowSelRange {$IFNDEF MSWINDOWS}name '_TrackBar_GetShowSelRange'{$ENDIF},
  TrackBar_SetShowSelRange {$IFNDEF MSWINDOWS}name '_TrackBar_SetShowSelRange'{$ENDIF},
  TrackBar_GetTabOrder {$IFNDEF MSWINDOWS}name '_TrackBar_GetTabOrder'{$ENDIF},
  TrackBar_SetTabOrder {$IFNDEF MSWINDOWS}name '_TrackBar_SetTabOrder'{$ENDIF},
  TrackBar_GetTabStop {$IFNDEF MSWINDOWS}name '_TrackBar_GetTabStop'{$ENDIF},
  TrackBar_SetTabStop {$IFNDEF MSWINDOWS}name '_TrackBar_SetTabStop'{$ENDIF},
  TrackBar_GetThumbLength {$IFNDEF MSWINDOWS}name '_TrackBar_GetThumbLength'{$ENDIF},
  TrackBar_SetThumbLength {$IFNDEF MSWINDOWS}name '_TrackBar_SetThumbLength'{$ENDIF},
  TrackBar_GetTickMarks {$IFNDEF MSWINDOWS}name '_TrackBar_GetTickMarks'{$ENDIF},
  TrackBar_SetTickMarks {$IFNDEF MSWINDOWS}name '_TrackBar_SetTickMarks'{$ENDIF},
  TrackBar_GetTickStyle {$IFNDEF MSWINDOWS}name '_TrackBar_GetTickStyle'{$ENDIF},
  TrackBar_SetTickStyle {$IFNDEF MSWINDOWS}name '_TrackBar_SetTickStyle'{$ENDIF},
  TrackBar_GetVisible {$IFNDEF MSWINDOWS}name '_TrackBar_GetVisible'{$ENDIF},
  TrackBar_SetVisible {$IFNDEF MSWINDOWS}name '_TrackBar_SetVisible'{$ENDIF},
  TrackBar_GetStyleElements {$IFNDEF MSWINDOWS}name '_TrackBar_GetStyleElements'{$ENDIF},
  TrackBar_SetStyleElements {$IFNDEF MSWINDOWS}name '_TrackBar_SetStyleElements'{$ENDIF},
  TrackBar_SetOnChange {$IFNDEF MSWINDOWS}name '_TrackBar_SetOnChange'{$ENDIF},
  TrackBar_SetOnEnter {$IFNDEF MSWINDOWS}name '_TrackBar_SetOnEnter'{$ENDIF},
  TrackBar_SetOnExit {$IFNDEF MSWINDOWS}name '_TrackBar_SetOnExit'{$ENDIF},
  TrackBar_SetOnKeyDown {$IFNDEF MSWINDOWS}name '_TrackBar_SetOnKeyDown'{$ENDIF},
  TrackBar_SetOnKeyPress {$IFNDEF MSWINDOWS}name '_TrackBar_SetOnKeyPress'{$ENDIF},
  TrackBar_SetOnKeyUp {$IFNDEF MSWINDOWS}name '_TrackBar_SetOnKeyUp'{$ENDIF},
  TrackBar_GetBrush {$IFNDEF MSWINDOWS}name '_TrackBar_GetBrush'{$ENDIF},
  TrackBar_GetControlCount {$IFNDEF MSWINDOWS}name '_TrackBar_GetControlCount'{$ENDIF},
  TrackBar_GetHandle {$IFNDEF MSWINDOWS}name '_TrackBar_GetHandle'{$ENDIF},
  TrackBar_GetParentWindow {$IFNDEF MSWINDOWS}name '_TrackBar_GetParentWindow'{$ENDIF},
  TrackBar_SetParentWindow {$IFNDEF MSWINDOWS}name '_TrackBar_SetParentWindow'{$ENDIF},
  TrackBar_GetAction {$IFNDEF MSWINDOWS}name '_TrackBar_GetAction'{$ENDIF},
  TrackBar_SetAction {$IFNDEF MSWINDOWS}name '_TrackBar_SetAction'{$ENDIF},
  TrackBar_GetBiDiMode {$IFNDEF MSWINDOWS}name '_TrackBar_GetBiDiMode'{$ENDIF},
  TrackBar_SetBiDiMode {$IFNDEF MSWINDOWS}name '_TrackBar_SetBiDiMode'{$ENDIF},
  TrackBar_GetBoundsRect {$IFNDEF MSWINDOWS}name '_TrackBar_GetBoundsRect'{$ENDIF},
  TrackBar_SetBoundsRect {$IFNDEF MSWINDOWS}name '_TrackBar_SetBoundsRect'{$ENDIF},
  TrackBar_GetClientHeight {$IFNDEF MSWINDOWS}name '_TrackBar_GetClientHeight'{$ENDIF},
  TrackBar_SetClientHeight {$IFNDEF MSWINDOWS}name '_TrackBar_SetClientHeight'{$ENDIF},
  TrackBar_GetClientRect {$IFNDEF MSWINDOWS}name '_TrackBar_GetClientRect'{$ENDIF},
  TrackBar_GetClientWidth {$IFNDEF MSWINDOWS}name '_TrackBar_GetClientWidth'{$ENDIF},
  TrackBar_SetClientWidth {$IFNDEF MSWINDOWS}name '_TrackBar_SetClientWidth'{$ENDIF},
  TrackBar_GetExplicitLeft {$IFNDEF MSWINDOWS}name '_TrackBar_GetExplicitLeft'{$ENDIF},
  TrackBar_GetExplicitTop {$IFNDEF MSWINDOWS}name '_TrackBar_GetExplicitTop'{$ENDIF},
  TrackBar_GetExplicitWidth {$IFNDEF MSWINDOWS}name '_TrackBar_GetExplicitWidth'{$ENDIF},
  TrackBar_GetExplicitHeight {$IFNDEF MSWINDOWS}name '_TrackBar_GetExplicitHeight'{$ENDIF},
  TrackBar_GetParent {$IFNDEF MSWINDOWS}name '_TrackBar_GetParent'{$ENDIF},
  TrackBar_SetParent {$IFNDEF MSWINDOWS}name '_TrackBar_SetParent'{$ENDIF},
  TrackBar_GetAlignWithMargins {$IFNDEF MSWINDOWS}name '_TrackBar_GetAlignWithMargins'{$ENDIF},
  TrackBar_SetAlignWithMargins {$IFNDEF MSWINDOWS}name '_TrackBar_SetAlignWithMargins'{$ENDIF},
  TrackBar_GetLeft {$IFNDEF MSWINDOWS}name '_TrackBar_GetLeft'{$ENDIF},
  TrackBar_SetLeft {$IFNDEF MSWINDOWS}name '_TrackBar_SetLeft'{$ENDIF},
  TrackBar_GetTop {$IFNDEF MSWINDOWS}name '_TrackBar_GetTop'{$ENDIF},
  TrackBar_SetTop {$IFNDEF MSWINDOWS}name '_TrackBar_SetTop'{$ENDIF},
  TrackBar_GetWidth {$IFNDEF MSWINDOWS}name '_TrackBar_GetWidth'{$ENDIF},
  TrackBar_SetWidth {$IFNDEF MSWINDOWS}name '_TrackBar_SetWidth'{$ENDIF},
  TrackBar_GetHeight {$IFNDEF MSWINDOWS}name '_TrackBar_GetHeight'{$ENDIF},
  TrackBar_SetHeight {$IFNDEF MSWINDOWS}name '_TrackBar_SetHeight'{$ENDIF},
  TrackBar_GetCursor {$IFNDEF MSWINDOWS}name '_TrackBar_GetCursor'{$ENDIF},
  TrackBar_SetCursor {$IFNDEF MSWINDOWS}name '_TrackBar_SetCursor'{$ENDIF},
  TrackBar_GetHint {$IFNDEF MSWINDOWS}name '_TrackBar_GetHint'{$ENDIF},
  TrackBar_SetHint {$IFNDEF MSWINDOWS}name '_TrackBar_SetHint'{$ENDIF},
  TrackBar_GetMargins {$IFNDEF MSWINDOWS}name '_TrackBar_GetMargins'{$ENDIF},
  TrackBar_SetMargins {$IFNDEF MSWINDOWS}name '_TrackBar_SetMargins'{$ENDIF},
  TrackBar_GetComponentCount {$IFNDEF MSWINDOWS}name '_TrackBar_GetComponentCount'{$ENDIF},
  TrackBar_GetComponentIndex {$IFNDEF MSWINDOWS}name '_TrackBar_GetComponentIndex'{$ENDIF},
  TrackBar_SetComponentIndex {$IFNDEF MSWINDOWS}name '_TrackBar_SetComponentIndex'{$ENDIF},
  TrackBar_GetOwner {$IFNDEF MSWINDOWS}name '_TrackBar_GetOwner'{$ENDIF},
  TrackBar_GetName {$IFNDEF MSWINDOWS}name '_TrackBar_GetName'{$ENDIF},
  TrackBar_SetName {$IFNDEF MSWINDOWS}name '_TrackBar_SetName'{$ENDIF},
  TrackBar_GetTag {$IFNDEF MSWINDOWS}name '_TrackBar_GetTag'{$ENDIF},
  TrackBar_SetTag {$IFNDEF MSWINDOWS}name '_TrackBar_SetTag'{$ENDIF},
  TrackBar_GetControls {$IFNDEF MSWINDOWS}name '_TrackBar_GetControls'{$ENDIF},
  TrackBar_GetComponents {$IFNDEF MSWINDOWS}name '_TrackBar_GetComponents'{$ENDIF};

