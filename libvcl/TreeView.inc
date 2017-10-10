
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function TreeView_Create(AOwner: TComponent): TTreeView; cdecl;
begin
  Result :=  TTreeView.Create(AOwner);
end;

procedure TreeView_Free(AObj: TTreeView); cdecl;
begin
  AObj.Free;
end;

function TreeView_AlphaSort(AObj: TTreeView; ARecurse: LongBool): LongBool; cdecl;
begin
  Result :=  AObj.AlphaSort(ARecurse);
end;

procedure TreeView_FullCollapse(AObj: TTreeView); cdecl;
begin
  AObj.FullCollapse;
end;

procedure TreeView_FullExpand(AObj: TTreeView); cdecl;
begin
  AObj.FullExpand;
end;

function TreeView_GetNodeAt(AObj: TTreeView; X: Integer; Y: Integer): TTreeNode; cdecl;
begin
  Result :=  AObj.GetNodeAt(X, Y);
end;

function TreeView_IsEditing(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.IsEditing;
end;

procedure TreeView_LoadFromFile(AObj: TTreeView; FileName: PChar); cdecl;
begin
  AObj.LoadFromFile(FileName);
end;

procedure TreeView_LoadFromStream(AObj: TTreeView; Stream: TStream); cdecl;
begin
  AObj.LoadFromStream(Stream);
end;

procedure TreeView_SaveToFile(AObj: TTreeView; FileName: PChar); cdecl;
begin
  AObj.SaveToFile(FileName);
end;

procedure TreeView_SaveToStream(AObj: TTreeView; Stream: TStream); cdecl;
begin
  AObj.SaveToStream(Stream);
end;

procedure TreeView_Deselect(AObj: TTreeView; Node: TTreeNode); cdecl;
begin
  AObj.Deselect(Node);
end;

procedure TreeView_Subselect(AObj: TTreeView; Node: TTreeNode; Validate: LongBool); cdecl;
begin
  AObj.Subselect(Node, Validate);
end;

procedure TreeView_ClearSelection(AObj: TTreeView; KeepPrimary: LongBool); cdecl;
begin
  AObj.ClearSelection(KeepPrimary);
end;

function TreeView_FindNextToSelect(AObj: TTreeView): TTreeNode; cdecl;
begin
  Result :=  AObj.FindNextToSelect;
end;

function TreeView_CanFocus(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.CanFocus;
end;

procedure TreeView_FlipChildren(AObj: TTreeView; AllLevels: LongBool); cdecl;
begin
  AObj.FlipChildren(AllLevels);
end;

function TreeView_Focused(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.Focused;
end;

function TreeView_HandleAllocated(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.HandleAllocated;
end;

procedure TreeView_Invalidate(AObj: TTreeView); cdecl;
begin
  AObj.Invalidate;
end;

procedure TreeView_Realign(AObj: TTreeView); cdecl;
begin
  AObj.Realign;
end;

procedure TreeView_Repaint(AObj: TTreeView); cdecl;
begin
  AObj.Repaint;
end;

procedure TreeView_ScaleBy(AObj: TTreeView; M: Integer; D: Integer); cdecl;
begin
  AObj.ScaleBy(M, D);
end;

procedure TreeView_SetBounds(AObj: TTreeView; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); cdecl;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

procedure TreeView_SetFocus(AObj: TTreeView); cdecl;
begin
  AObj.SetFocus;
end;

procedure TreeView_Update(AObj: TTreeView); cdecl;
begin
  AObj.Update;
end;

procedure TreeView_BringToFront(AObj: TTreeView); cdecl;
begin
  AObj.BringToFront;
end;

function TreeView_HasParent(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.HasParent;
end;

procedure TreeView_Hide(AObj: TTreeView); cdecl;
begin
  AObj.Hide;
end;

function TreeView_Perform(AObj: TTreeView; Msg: Cardinal; WParam: NativeUInt; LParam: NativeInt): NativeInt; cdecl;
begin
  Result :=  AObj.Perform(Msg, WParam, LParam);
end;

procedure TreeView_Refresh(AObj: TTreeView); cdecl;
begin
  AObj.Refresh;
end;

procedure TreeView_SendToBack(AObj: TTreeView); cdecl;
begin
  AObj.SendToBack;
end;

procedure TreeView_Show(AObj: TTreeView); cdecl;
begin
  AObj.Show;
end;

function TreeView_GetTextBuf(AObj: TTreeView; Buffer: PWideChar; BufSize: Integer): Integer; cdecl;
begin
  Result :=  AObj.GetTextBuf(Buffer, BufSize);
end;

function TreeView_FindComponent(AObj: TTreeView; AName: PChar): TComponent; cdecl;
begin
  Result :=  AObj.FindComponent(AName);
end;

function TreeView_GetNamePath(AObj: TTreeView): PChar; cdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

procedure TreeView_Assign(AObj: TTreeView; Source: TPersistent); cdecl;
begin
  AObj.Assign(Source);
end;

function TreeView_ClassName(AObj: TTreeView): PChar; cdecl;
begin
  Result :=  PChar(AObj.ClassName);
end;

function TreeView_Equals(AObj: TTreeView; Obj: TObject): LongBool; cdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function TreeView_GetHashCode(AObj: TTreeView): Integer; cdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function TreeView_ToString(AObj: TTreeView): PChar; cdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function TreeView_GetAlign(AObj: TTreeView): TAlign; cdecl;
begin
  Result :=  AObj.Align;
end;

procedure TreeView_SetAlign(AObj: TTreeView; AValue: TAlign); cdecl;
begin
  AObj.Align := AValue;
end;

function TreeView_GetAnchors(AObj: TTreeView): TAnchors; cdecl;
begin
  Result :=  AObj.Anchors;
end;

procedure TreeView_SetAnchors(AObj: TTreeView; AValue: TAnchors); cdecl;
begin
  AObj.Anchors := AValue;
end;

function TreeView_GetAutoExpand(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.AutoExpand;
end;

procedure TreeView_SetAutoExpand(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.AutoExpand := AValue;
end;

function TreeView_GetBevelEdges(AObj: TTreeView): TBevelEdges; cdecl;
begin
  Result :=  AObj.BevelEdges;
end;

procedure TreeView_SetBevelEdges(AObj: TTreeView; AValue: TBevelEdges); cdecl;
begin
  AObj.BevelEdges := AValue;
end;

function TreeView_GetBevelInner(AObj: TTreeView): TBevelCut; cdecl;
begin
  Result :=  AObj.BevelInner;
end;

procedure TreeView_SetBevelInner(AObj: TTreeView; AValue: TBevelCut); cdecl;
begin
  AObj.BevelInner := AValue;
end;

function TreeView_GetBevelOuter(AObj: TTreeView): TBevelCut; cdecl;
begin
  Result :=  AObj.BevelOuter;
end;

procedure TreeView_SetBevelOuter(AObj: TTreeView; AValue: TBevelCut); cdecl;
begin
  AObj.BevelOuter := AValue;
end;

function TreeView_GetBevelKind(AObj: TTreeView): TBevelKind; cdecl;
begin
  Result :=  AObj.BevelKind;
end;

procedure TreeView_SetBevelKind(AObj: TTreeView; AValue: TBevelKind); cdecl;
begin
  AObj.BevelKind := AValue;
end;

function TreeView_GetBiDiMode(AObj: TTreeView): TBiDiMode; cdecl;
begin
  Result :=  AObj.BiDiMode;
end;

procedure TreeView_SetBiDiMode(AObj: TTreeView; AValue: TBiDiMode); cdecl;
begin
  AObj.BiDiMode := AValue;
end;

function TreeView_GetBorderStyle(AObj: TTreeView): TBorderStyle; cdecl;
begin
  Result :=  AObj.BorderStyle;
end;

procedure TreeView_SetBorderStyle(AObj: TTreeView; AValue: TBorderStyle); cdecl;
begin
  AObj.BorderStyle := AValue;
end;

function TreeView_GetBorderWidth(AObj: TTreeView): TBorderWidth; cdecl;
begin
  Result :=  AObj.BorderWidth;
end;

procedure TreeView_SetBorderWidth(AObj: TTreeView; AValue: TBorderWidth); cdecl;
begin
  AObj.BorderWidth := AValue;
end;

function TreeView_GetChangeDelay(AObj: TTreeView): Integer; cdecl;
begin
  Result :=  AObj.ChangeDelay;
end;

procedure TreeView_SetChangeDelay(AObj: TTreeView; AValue: Integer); cdecl;
begin
  AObj.ChangeDelay := AValue;
end;

function TreeView_GetColor(AObj: TTreeView): TColor; cdecl;
begin
  Result :=  AObj.Color;
end;

procedure TreeView_SetColor(AObj: TTreeView; AValue: TColor); cdecl;
begin
  AObj.Color := AValue;
end;

function TreeView_GetDoubleBuffered(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.DoubleBuffered;
end;

procedure TreeView_SetDoubleBuffered(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.DoubleBuffered := AValue;
end;

function TreeView_GetEnabled(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure TreeView_SetEnabled(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.Enabled := AValue;
end;

function TreeView_GetFont(AObj: TTreeView): TFont; cdecl;
begin
  Result :=  AObj.Font;
end;

procedure TreeView_SetFont(AObj: TTreeView; AValue: TFont); cdecl;
begin
  AObj.Font := AValue;
end;

function TreeView_GetHideSelection(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.HideSelection;
end;

procedure TreeView_SetHideSelection(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.HideSelection := AValue;
end;

function TreeView_GetHotTrack(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.HotTrack;
end;

procedure TreeView_SetHotTrack(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.HotTrack := AValue;
end;

function TreeView_GetImages(AObj: TTreeView): TCustomImageList; cdecl;
begin
  Result :=  AObj.Images;
end;

procedure TreeView_SetImages(AObj: TTreeView; AValue: TCustomImageList); cdecl;
begin
  AObj.Images := AValue;
end;

function TreeView_GetIndent(AObj: TTreeView): Integer; cdecl;
begin
  Result :=  AObj.Indent;
end;

procedure TreeView_SetIndent(AObj: TTreeView; AValue: Integer); cdecl;
begin
  AObj.Indent := AValue;
end;

function TreeView_GetMultiSelect(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.MultiSelect;
end;

procedure TreeView_SetMultiSelect(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.MultiSelect := AValue;
end;

function TreeView_GetMultiSelectStyle(AObj: TTreeView): TMultiSelectStyle; cdecl;
begin
  Result :=  AObj.MultiSelectStyle;
end;

procedure TreeView_SetMultiSelectStyle(AObj: TTreeView; AValue: TMultiSelectStyle); cdecl;
begin
  AObj.MultiSelectStyle := AValue;
end;

function TreeView_GetParentColor(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.ParentColor;
end;

procedure TreeView_SetParentColor(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.ParentColor := AValue;
end;

function TreeView_GetParentCtl3D(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.ParentCtl3D;
end;

procedure TreeView_SetParentCtl3D(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.ParentCtl3D := AValue;
end;

function TreeView_GetParentDoubleBuffered(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.ParentDoubleBuffered;
end;

procedure TreeView_SetParentDoubleBuffered(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.ParentDoubleBuffered := AValue;
end;

function TreeView_GetParentFont(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.ParentFont;
end;

procedure TreeView_SetParentFont(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.ParentFont := AValue;
end;

function TreeView_GetParentShowHint(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.ParentShowHint;
end;

procedure TreeView_SetParentShowHint(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.ParentShowHint := AValue;
end;

function TreeView_GetPopupMenu(AObj: TTreeView): TPopupMenu; cdecl;
begin
  Result :=  AObj.PopupMenu;
end;

procedure TreeView_SetPopupMenu(AObj: TTreeView; AValue: TPopupMenu); cdecl;
begin
  AObj.PopupMenu := AValue;
end;

function TreeView_GetReadOnly(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.ReadOnly;
end;

procedure TreeView_SetReadOnly(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.ReadOnly := AValue;
end;

function TreeView_GetRightClickSelect(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.RightClickSelect;
end;

procedure TreeView_SetRightClickSelect(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.RightClickSelect := AValue;
end;

function TreeView_GetRowSelect(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.RowSelect;
end;

procedure TreeView_SetRowSelect(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.RowSelect := AValue;
end;

function TreeView_GetShowButtons(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.ShowButtons;
end;

procedure TreeView_SetShowButtons(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.ShowButtons := AValue;
end;

function TreeView_GetShowHint(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.ShowHint;
end;

procedure TreeView_SetShowHint(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.ShowHint := AValue;
end;

function TreeView_GetShowLines(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.ShowLines;
end;

procedure TreeView_SetShowLines(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.ShowLines := AValue;
end;

function TreeView_GetShowRoot(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.ShowRoot;
end;

procedure TreeView_SetShowRoot(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.ShowRoot := AValue;
end;

function TreeView_GetSortType(AObj: TTreeView): TSortType; cdecl;
begin
  Result :=  AObj.SortType;
end;

procedure TreeView_SetSortType(AObj: TTreeView; AValue: TSortType); cdecl;
begin
  AObj.SortType := AValue;
end;

function TreeView_GetStateImages(AObj: TTreeView): TCustomImageList; cdecl;
begin
  Result :=  AObj.StateImages;
end;

procedure TreeView_SetStateImages(AObj: TTreeView; AValue: TCustomImageList); cdecl;
begin
  AObj.StateImages := AValue;
end;

function TreeView_GetTabOrder(AObj: TTreeView): TTabOrder; cdecl;
begin
  Result :=  AObj.TabOrder;
end;

procedure TreeView_SetTabOrder(AObj: TTreeView; AValue: TTabOrder); cdecl;
begin
  AObj.TabOrder := AValue;
end;

function TreeView_GetTabStop(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.TabStop;
end;

procedure TreeView_SetTabStop(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.TabStop := AValue;
end;

function TreeView_GetToolTips(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.ToolTips;
end;

procedure TreeView_SetToolTips(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.ToolTips := AValue;
end;

function TreeView_GetVisible(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.Visible;
end;

procedure TreeView_SetVisible(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.Visible := AValue;
end;

function TreeView_GetStyleElements(AObj: TTreeView): TStyleElements; cdecl;
begin
  Result :=  AObj.StyleElements;
end;

procedure TreeView_SetStyleElements(AObj: TTreeView; AValue: TStyleElements); cdecl;
begin
  AObj.StyleElements := AValue;
end;

procedure TreeView_SetOnChange(AObj: TTreeView; AEventId: NativeUInt); stdcall
begin
  if AEventId = 0 then
  begin
    AObj.OnChange := nil;
    TEventClass.Remove(AObj, geTreeViewChange);
    Exit;
  end;
  AObj.OnChange := TEventClass.TreeViewOnChange;
  TEventClass.Add(AObj, geTreeViewChange, AEventId);
end;

procedure TreeView_SetOnClick(AObj: TTreeView; AEventId: NativeUInt); stdcall
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

procedure TreeView_SetOnDblClick(AObj: TTreeView; AEventId: NativeUInt); stdcall
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

procedure TreeView_SetOnEnter(AObj: TTreeView; AEventId: NativeUInt); stdcall
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

procedure TreeView_SetOnExit(AObj: TTreeView; AEventId: NativeUInt); stdcall
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

procedure TreeView_SetOnKeyDown(AObj: TTreeView; AEventId: NativeUInt); stdcall
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

procedure TreeView_SetOnKeyPress(AObj: TTreeView; AEventId: NativeUInt); stdcall
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

procedure TreeView_SetOnKeyUp(AObj: TTreeView; AEventId: NativeUInt); stdcall
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

procedure TreeView_SetOnMouseDown(AObj: TTreeView; AEventId: NativeUInt); stdcall
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

procedure TreeView_SetOnMouseEnter(AObj: TTreeView; AEventId: NativeUInt); stdcall
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

procedure TreeView_SetOnMouseLeave(AObj: TTreeView; AEventId: NativeUInt); stdcall
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

procedure TreeView_SetOnMouseMove(AObj: TTreeView; AEventId: NativeUInt); stdcall
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

procedure TreeView_SetOnMouseUp(AObj: TTreeView; AEventId: NativeUInt); stdcall
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

function TreeView_GetItems(AObj: TTreeView): TTreeNodes; cdecl;
begin
  Result :=  AObj.Items;
end;

procedure TreeView_SetItems(AObj: TTreeView; AValue: TTreeNodes); cdecl;
begin
  AObj.Items := AValue;
end;

function TreeView_GetCanvas(AObj: TTreeView): TCanvas; cdecl;
begin
  Result :=  AObj.Canvas;
end;

function TreeView_GetDropTarget(AObj: TTreeView): TTreeNode; cdecl;
begin
  Result :=  AObj.DropTarget;
end;

procedure TreeView_SetDropTarget(AObj: TTreeView; AValue: TTreeNode); cdecl;
begin
  AObj.DropTarget := AValue;
end;

function TreeView_GetSelected(AObj: TTreeView): TTreeNode; cdecl;
begin
  Result :=  AObj.Selected;
end;

procedure TreeView_SetSelected(AObj: TTreeView; AValue: TTreeNode); cdecl;
begin
  AObj.Selected := AValue;
end;

function TreeView_GetTopItem(AObj: TTreeView): TTreeNode; cdecl;
begin
  Result :=  AObj.TopItem;
end;

procedure TreeView_SetTopItem(AObj: TTreeView; AValue: TTreeNode); cdecl;
begin
  AObj.TopItem := AValue;
end;

function TreeView_GetSelectionCount(AObj: TTreeView): Cardinal; cdecl;
begin
  Result :=  AObj.SelectionCount;
end;

function TreeView_GetBrush(AObj: TTreeView): TBrush; cdecl;
begin
  Result :=  AObj.Brush;
end;

function TreeView_GetControlCount(AObj: TTreeView): Integer; cdecl;
begin
  Result :=  AObj.ControlCount;
end;

function TreeView_GetHandle(AObj: TTreeView): HWND; cdecl;
begin
  Result :=  AObj.Handle;
end;

function TreeView_GetParentWindow(AObj: TTreeView): HWND; cdecl;
begin
  Result :=  AObj.ParentWindow;
end;

procedure TreeView_SetParentWindow(AObj: TTreeView; AValue: HWND); cdecl;
begin
  AObj.ParentWindow := AValue;
end;

function TreeView_GetAction(AObj: TTreeView): TBasicAction; cdecl;
begin
  Result :=  AObj.Action;
end;

procedure TreeView_SetAction(AObj: TTreeView; AValue: TBasicAction); cdecl;
begin
  AObj.Action := AValue;
end;

procedure TreeView_GetBoundsRect(AObj: TTreeView; var Result: TRect); cdecl;
begin
  Result :=  AObj.BoundsRect;
end;

procedure TreeView_SetBoundsRect(AObj: TTreeView; var AValue: TRect); cdecl;
begin
  AObj.BoundsRect := AValue;
end;

function TreeView_GetClientHeight(AObj: TTreeView): Integer; cdecl;
begin
  Result :=  AObj.ClientHeight;
end;

procedure TreeView_SetClientHeight(AObj: TTreeView; AValue: Integer); cdecl;
begin
  AObj.ClientHeight := AValue;
end;

procedure TreeView_GetClientRect(AObj: TTreeView; var Result: TRect); cdecl;
begin
  Result :=  AObj.ClientRect;
end;

function TreeView_GetClientWidth(AObj: TTreeView): Integer; cdecl;
begin
  Result :=  AObj.ClientWidth;
end;

procedure TreeView_SetClientWidth(AObj: TTreeView; AValue: Integer); cdecl;
begin
  AObj.ClientWidth := AValue;
end;

function TreeView_GetExplicitLeft(AObj: TTreeView): Integer; cdecl;
begin
  Result :=  AObj.ExplicitLeft;
end;

function TreeView_GetExplicitTop(AObj: TTreeView): Integer; cdecl;
begin
  Result :=  AObj.ExplicitTop;
end;

function TreeView_GetExplicitWidth(AObj: TTreeView): Integer; cdecl;
begin
  Result :=  AObj.ExplicitWidth;
end;

function TreeView_GetExplicitHeight(AObj: TTreeView): Integer; cdecl;
begin
  Result :=  AObj.ExplicitHeight;
end;

function TreeView_GetParent(AObj: TTreeView): TWinControl; cdecl;
begin
  Result :=  AObj.Parent;
end;

procedure TreeView_SetParent(AObj: TTreeView; AValue: TWinControl); cdecl;
begin
  AObj.Parent := AValue;
end;

function TreeView_GetAlignWithMargins(AObj: TTreeView): LongBool; cdecl;
begin
  Result :=  AObj.AlignWithMargins;
end;

procedure TreeView_SetAlignWithMargins(AObj: TTreeView; AValue: LongBool); cdecl;
begin
  AObj.AlignWithMargins := AValue;
end;

function TreeView_GetLeft(AObj: TTreeView): Integer; cdecl;
begin
  Result :=  AObj.Left;
end;

procedure TreeView_SetLeft(AObj: TTreeView; AValue: Integer); cdecl;
begin
  AObj.Left := AValue;
end;

function TreeView_GetTop(AObj: TTreeView): Integer; cdecl;
begin
  Result :=  AObj.Top;
end;

procedure TreeView_SetTop(AObj: TTreeView; AValue: Integer); cdecl;
begin
  AObj.Top := AValue;
end;

function TreeView_GetWidth(AObj: TTreeView): Integer; cdecl;
begin
  Result :=  AObj.Width;
end;

procedure TreeView_SetWidth(AObj: TTreeView; AValue: Integer); cdecl;
begin
  AObj.Width := AValue;
end;

function TreeView_GetHeight(AObj: TTreeView): Integer; cdecl;
begin
  Result :=  AObj.Height;
end;

procedure TreeView_SetHeight(AObj: TTreeView; AValue: Integer); cdecl;
begin
  AObj.Height := AValue;
end;

function TreeView_GetCursor(AObj: TTreeView): TCursor; cdecl;
begin
  Result :=  AObj.Cursor;
end;

procedure TreeView_SetCursor(AObj: TTreeView; AValue: TCursor); cdecl;
begin
  AObj.Cursor := AValue;
end;

function TreeView_GetHint(AObj: TTreeView): PChar; cdecl;
begin
  Result :=  PChar(AObj.Hint);
end;

procedure TreeView_SetHint(AObj: TTreeView; AValue: PChar); cdecl;
begin
  AObj.Hint := AValue;
end;

function TreeView_GetMargins(AObj: TTreeView): TMargins; cdecl;
begin
  Result :=  AObj.Margins;
end;

procedure TreeView_SetMargins(AObj: TTreeView; AValue: TMargins); cdecl;
begin
  AObj.Margins := AValue;
end;

function TreeView_GetComponentCount(AObj: TTreeView): Integer; cdecl;
begin
  Result :=  AObj.ComponentCount;
end;

function TreeView_GetComponentIndex(AObj: TTreeView): Integer; cdecl;
begin
  Result :=  AObj.ComponentIndex;
end;

procedure TreeView_SetComponentIndex(AObj: TTreeView; AValue: Integer); cdecl;
begin
  AObj.ComponentIndex := AValue;
end;

function TreeView_GetOwner(AObj: TTreeView): TComponent; cdecl;
begin
  Result :=  AObj.Owner;
end;

function TreeView_GetName(AObj: TTreeView): PChar; cdecl;
begin
  Result :=  PChar(AObj.Name);
end;

procedure TreeView_SetName(AObj: TTreeView; AValue: PChar); cdecl;
begin
  AObj.Name := AValue;
end;

function TreeView_GetTag(AObj: TTreeView): NativeInt; cdecl;
begin
  Result :=  AObj.Tag;
end;

procedure TreeView_SetTag(AObj: TTreeView; AValue: NativeInt); cdecl;
begin
  AObj.Tag := AValue;
end;

function TreeView_GetSelections(AObj: TTreeView; Index: Integer): TTreeNode; cdecl;
begin
  Result :=  AObj.Selections[Index];
end;

function TreeView_GetControls(AObj: TTreeView; Index: Integer): TControl; cdecl;
begin
  Result :=  AObj.Controls[Index];
end;

function TreeView_GetComponents(AObj: TTreeView; AIndex: Integer): TComponent; cdecl;
begin
  Result :=  AObj.Components[AIndex];
end;


exports
  TreeView_Create {$IFNDEF MSWINDOWS}name '_TreeView_Create'{$ENDIF},
  TreeView_Free {$IFNDEF MSWINDOWS}name '_TreeView_Free'{$ENDIF},
  TreeView_AlphaSort {$IFNDEF MSWINDOWS}name '_TreeView_AlphaSort'{$ENDIF},
  TreeView_FullCollapse {$IFNDEF MSWINDOWS}name '_TreeView_FullCollapse'{$ENDIF},
  TreeView_FullExpand {$IFNDEF MSWINDOWS}name '_TreeView_FullExpand'{$ENDIF},
  TreeView_GetNodeAt {$IFNDEF MSWINDOWS}name '_TreeView_GetNodeAt'{$ENDIF},
  TreeView_IsEditing {$IFNDEF MSWINDOWS}name '_TreeView_IsEditing'{$ENDIF},
  TreeView_LoadFromFile {$IFNDEF MSWINDOWS}name '_TreeView_LoadFromFile'{$ENDIF},
  TreeView_LoadFromStream {$IFNDEF MSWINDOWS}name '_TreeView_LoadFromStream'{$ENDIF},
  TreeView_SaveToFile {$IFNDEF MSWINDOWS}name '_TreeView_SaveToFile'{$ENDIF},
  TreeView_SaveToStream {$IFNDEF MSWINDOWS}name '_TreeView_SaveToStream'{$ENDIF},
  TreeView_Deselect {$IFNDEF MSWINDOWS}name '_TreeView_Deselect'{$ENDIF},
  TreeView_Subselect {$IFNDEF MSWINDOWS}name '_TreeView_Subselect'{$ENDIF},
  TreeView_ClearSelection {$IFNDEF MSWINDOWS}name '_TreeView_ClearSelection'{$ENDIF},
  TreeView_FindNextToSelect {$IFNDEF MSWINDOWS}name '_TreeView_FindNextToSelect'{$ENDIF},
  TreeView_CanFocus {$IFNDEF MSWINDOWS}name '_TreeView_CanFocus'{$ENDIF},
  TreeView_FlipChildren {$IFNDEF MSWINDOWS}name '_TreeView_FlipChildren'{$ENDIF},
  TreeView_Focused {$IFNDEF MSWINDOWS}name '_TreeView_Focused'{$ENDIF},
  TreeView_HandleAllocated {$IFNDEF MSWINDOWS}name '_TreeView_HandleAllocated'{$ENDIF},
  TreeView_Invalidate {$IFNDEF MSWINDOWS}name '_TreeView_Invalidate'{$ENDIF},
  TreeView_Realign {$IFNDEF MSWINDOWS}name '_TreeView_Realign'{$ENDIF},
  TreeView_Repaint {$IFNDEF MSWINDOWS}name '_TreeView_Repaint'{$ENDIF},
  TreeView_ScaleBy {$IFNDEF MSWINDOWS}name '_TreeView_ScaleBy'{$ENDIF},
  TreeView_SetBounds {$IFNDEF MSWINDOWS}name '_TreeView_SetBounds'{$ENDIF},
  TreeView_SetFocus {$IFNDEF MSWINDOWS}name '_TreeView_SetFocus'{$ENDIF},
  TreeView_Update {$IFNDEF MSWINDOWS}name '_TreeView_Update'{$ENDIF},
  TreeView_BringToFront {$IFNDEF MSWINDOWS}name '_TreeView_BringToFront'{$ENDIF},
  TreeView_HasParent {$IFNDEF MSWINDOWS}name '_TreeView_HasParent'{$ENDIF},
  TreeView_Hide {$IFNDEF MSWINDOWS}name '_TreeView_Hide'{$ENDIF},
  TreeView_Perform {$IFNDEF MSWINDOWS}name '_TreeView_Perform'{$ENDIF},
  TreeView_Refresh {$IFNDEF MSWINDOWS}name '_TreeView_Refresh'{$ENDIF},
  TreeView_SendToBack {$IFNDEF MSWINDOWS}name '_TreeView_SendToBack'{$ENDIF},
  TreeView_Show {$IFNDEF MSWINDOWS}name '_TreeView_Show'{$ENDIF},
  TreeView_GetTextBuf {$IFNDEF MSWINDOWS}name '_TreeView_GetTextBuf'{$ENDIF},
  TreeView_FindComponent {$IFNDEF MSWINDOWS}name '_TreeView_FindComponent'{$ENDIF},
  TreeView_GetNamePath {$IFNDEF MSWINDOWS}name '_TreeView_GetNamePath'{$ENDIF},
  TreeView_Assign {$IFNDEF MSWINDOWS}name '_TreeView_Assign'{$ENDIF},
  TreeView_ClassName {$IFNDEF MSWINDOWS}name '_TreeView_ClassName'{$ENDIF},
  TreeView_Equals {$IFNDEF MSWINDOWS}name '_TreeView_Equals'{$ENDIF},
  TreeView_GetHashCode {$IFNDEF MSWINDOWS}name '_TreeView_GetHashCode'{$ENDIF},
  TreeView_ToString {$IFNDEF MSWINDOWS}name '_TreeView_ToString'{$ENDIF},
  TreeView_GetAlign {$IFNDEF MSWINDOWS}name '_TreeView_GetAlign'{$ENDIF},
  TreeView_SetAlign {$IFNDEF MSWINDOWS}name '_TreeView_SetAlign'{$ENDIF},
  TreeView_GetAnchors {$IFNDEF MSWINDOWS}name '_TreeView_GetAnchors'{$ENDIF},
  TreeView_SetAnchors {$IFNDEF MSWINDOWS}name '_TreeView_SetAnchors'{$ENDIF},
  TreeView_GetAutoExpand {$IFNDEF MSWINDOWS}name '_TreeView_GetAutoExpand'{$ENDIF},
  TreeView_SetAutoExpand {$IFNDEF MSWINDOWS}name '_TreeView_SetAutoExpand'{$ENDIF},
  TreeView_GetBevelEdges {$IFNDEF MSWINDOWS}name '_TreeView_GetBevelEdges'{$ENDIF},
  TreeView_SetBevelEdges {$IFNDEF MSWINDOWS}name '_TreeView_SetBevelEdges'{$ENDIF},
  TreeView_GetBevelInner {$IFNDEF MSWINDOWS}name '_TreeView_GetBevelInner'{$ENDIF},
  TreeView_SetBevelInner {$IFNDEF MSWINDOWS}name '_TreeView_SetBevelInner'{$ENDIF},
  TreeView_GetBevelOuter {$IFNDEF MSWINDOWS}name '_TreeView_GetBevelOuter'{$ENDIF},
  TreeView_SetBevelOuter {$IFNDEF MSWINDOWS}name '_TreeView_SetBevelOuter'{$ENDIF},
  TreeView_GetBevelKind {$IFNDEF MSWINDOWS}name '_TreeView_GetBevelKind'{$ENDIF},
  TreeView_SetBevelKind {$IFNDEF MSWINDOWS}name '_TreeView_SetBevelKind'{$ENDIF},
  TreeView_GetBiDiMode {$IFNDEF MSWINDOWS}name '_TreeView_GetBiDiMode'{$ENDIF},
  TreeView_SetBiDiMode {$IFNDEF MSWINDOWS}name '_TreeView_SetBiDiMode'{$ENDIF},
  TreeView_GetBorderStyle {$IFNDEF MSWINDOWS}name '_TreeView_GetBorderStyle'{$ENDIF},
  TreeView_SetBorderStyle {$IFNDEF MSWINDOWS}name '_TreeView_SetBorderStyle'{$ENDIF},
  TreeView_GetBorderWidth {$IFNDEF MSWINDOWS}name '_TreeView_GetBorderWidth'{$ENDIF},
  TreeView_SetBorderWidth {$IFNDEF MSWINDOWS}name '_TreeView_SetBorderWidth'{$ENDIF},
  TreeView_GetChangeDelay {$IFNDEF MSWINDOWS}name '_TreeView_GetChangeDelay'{$ENDIF},
  TreeView_SetChangeDelay {$IFNDEF MSWINDOWS}name '_TreeView_SetChangeDelay'{$ENDIF},
  TreeView_GetColor {$IFNDEF MSWINDOWS}name '_TreeView_GetColor'{$ENDIF},
  TreeView_SetColor {$IFNDEF MSWINDOWS}name '_TreeView_SetColor'{$ENDIF},
  TreeView_GetDoubleBuffered {$IFNDEF MSWINDOWS}name '_TreeView_GetDoubleBuffered'{$ENDIF},
  TreeView_SetDoubleBuffered {$IFNDEF MSWINDOWS}name '_TreeView_SetDoubleBuffered'{$ENDIF},
  TreeView_GetEnabled {$IFNDEF MSWINDOWS}name '_TreeView_GetEnabled'{$ENDIF},
  TreeView_SetEnabled {$IFNDEF MSWINDOWS}name '_TreeView_SetEnabled'{$ENDIF},
  TreeView_GetFont {$IFNDEF MSWINDOWS}name '_TreeView_GetFont'{$ENDIF},
  TreeView_SetFont {$IFNDEF MSWINDOWS}name '_TreeView_SetFont'{$ENDIF},
  TreeView_GetHideSelection {$IFNDEF MSWINDOWS}name '_TreeView_GetHideSelection'{$ENDIF},
  TreeView_SetHideSelection {$IFNDEF MSWINDOWS}name '_TreeView_SetHideSelection'{$ENDIF},
  TreeView_GetHotTrack {$IFNDEF MSWINDOWS}name '_TreeView_GetHotTrack'{$ENDIF},
  TreeView_SetHotTrack {$IFNDEF MSWINDOWS}name '_TreeView_SetHotTrack'{$ENDIF},
  TreeView_GetImages {$IFNDEF MSWINDOWS}name '_TreeView_GetImages'{$ENDIF},
  TreeView_SetImages {$IFNDEF MSWINDOWS}name '_TreeView_SetImages'{$ENDIF},
  TreeView_GetIndent {$IFNDEF MSWINDOWS}name '_TreeView_GetIndent'{$ENDIF},
  TreeView_SetIndent {$IFNDEF MSWINDOWS}name '_TreeView_SetIndent'{$ENDIF},
  TreeView_GetMultiSelect {$IFNDEF MSWINDOWS}name '_TreeView_GetMultiSelect'{$ENDIF},
  TreeView_SetMultiSelect {$IFNDEF MSWINDOWS}name '_TreeView_SetMultiSelect'{$ENDIF},
  TreeView_GetMultiSelectStyle {$IFNDEF MSWINDOWS}name '_TreeView_GetMultiSelectStyle'{$ENDIF},
  TreeView_SetMultiSelectStyle {$IFNDEF MSWINDOWS}name '_TreeView_SetMultiSelectStyle'{$ENDIF},
  TreeView_GetParentColor {$IFNDEF MSWINDOWS}name '_TreeView_GetParentColor'{$ENDIF},
  TreeView_SetParentColor {$IFNDEF MSWINDOWS}name '_TreeView_SetParentColor'{$ENDIF},
  TreeView_GetParentCtl3D {$IFNDEF MSWINDOWS}name '_TreeView_GetParentCtl3D'{$ENDIF},
  TreeView_SetParentCtl3D {$IFNDEF MSWINDOWS}name '_TreeView_SetParentCtl3D'{$ENDIF},
  TreeView_GetParentDoubleBuffered {$IFNDEF MSWINDOWS}name '_TreeView_GetParentDoubleBuffered'{$ENDIF},
  TreeView_SetParentDoubleBuffered {$IFNDEF MSWINDOWS}name '_TreeView_SetParentDoubleBuffered'{$ENDIF},
  TreeView_GetParentFont {$IFNDEF MSWINDOWS}name '_TreeView_GetParentFont'{$ENDIF},
  TreeView_SetParentFont {$IFNDEF MSWINDOWS}name '_TreeView_SetParentFont'{$ENDIF},
  TreeView_GetParentShowHint {$IFNDEF MSWINDOWS}name '_TreeView_GetParentShowHint'{$ENDIF},
  TreeView_SetParentShowHint {$IFNDEF MSWINDOWS}name '_TreeView_SetParentShowHint'{$ENDIF},
  TreeView_GetPopupMenu {$IFNDEF MSWINDOWS}name '_TreeView_GetPopupMenu'{$ENDIF},
  TreeView_SetPopupMenu {$IFNDEF MSWINDOWS}name '_TreeView_SetPopupMenu'{$ENDIF},
  TreeView_GetReadOnly {$IFNDEF MSWINDOWS}name '_TreeView_GetReadOnly'{$ENDIF},
  TreeView_SetReadOnly {$IFNDEF MSWINDOWS}name '_TreeView_SetReadOnly'{$ENDIF},
  TreeView_GetRightClickSelect {$IFNDEF MSWINDOWS}name '_TreeView_GetRightClickSelect'{$ENDIF},
  TreeView_SetRightClickSelect {$IFNDEF MSWINDOWS}name '_TreeView_SetRightClickSelect'{$ENDIF},
  TreeView_GetRowSelect {$IFNDEF MSWINDOWS}name '_TreeView_GetRowSelect'{$ENDIF},
  TreeView_SetRowSelect {$IFNDEF MSWINDOWS}name '_TreeView_SetRowSelect'{$ENDIF},
  TreeView_GetShowButtons {$IFNDEF MSWINDOWS}name '_TreeView_GetShowButtons'{$ENDIF},
  TreeView_SetShowButtons {$IFNDEF MSWINDOWS}name '_TreeView_SetShowButtons'{$ENDIF},
  TreeView_GetShowHint {$IFNDEF MSWINDOWS}name '_TreeView_GetShowHint'{$ENDIF},
  TreeView_SetShowHint {$IFNDEF MSWINDOWS}name '_TreeView_SetShowHint'{$ENDIF},
  TreeView_GetShowLines {$IFNDEF MSWINDOWS}name '_TreeView_GetShowLines'{$ENDIF},
  TreeView_SetShowLines {$IFNDEF MSWINDOWS}name '_TreeView_SetShowLines'{$ENDIF},
  TreeView_GetShowRoot {$IFNDEF MSWINDOWS}name '_TreeView_GetShowRoot'{$ENDIF},
  TreeView_SetShowRoot {$IFNDEF MSWINDOWS}name '_TreeView_SetShowRoot'{$ENDIF},
  TreeView_GetSortType {$IFNDEF MSWINDOWS}name '_TreeView_GetSortType'{$ENDIF},
  TreeView_SetSortType {$IFNDEF MSWINDOWS}name '_TreeView_SetSortType'{$ENDIF},
  TreeView_GetStateImages {$IFNDEF MSWINDOWS}name '_TreeView_GetStateImages'{$ENDIF},
  TreeView_SetStateImages {$IFNDEF MSWINDOWS}name '_TreeView_SetStateImages'{$ENDIF},
  TreeView_GetTabOrder {$IFNDEF MSWINDOWS}name '_TreeView_GetTabOrder'{$ENDIF},
  TreeView_SetTabOrder {$IFNDEF MSWINDOWS}name '_TreeView_SetTabOrder'{$ENDIF},
  TreeView_GetTabStop {$IFNDEF MSWINDOWS}name '_TreeView_GetTabStop'{$ENDIF},
  TreeView_SetTabStop {$IFNDEF MSWINDOWS}name '_TreeView_SetTabStop'{$ENDIF},
  TreeView_GetToolTips {$IFNDEF MSWINDOWS}name '_TreeView_GetToolTips'{$ENDIF},
  TreeView_SetToolTips {$IFNDEF MSWINDOWS}name '_TreeView_SetToolTips'{$ENDIF},
  TreeView_GetVisible {$IFNDEF MSWINDOWS}name '_TreeView_GetVisible'{$ENDIF},
  TreeView_SetVisible {$IFNDEF MSWINDOWS}name '_TreeView_SetVisible'{$ENDIF},
  TreeView_GetStyleElements {$IFNDEF MSWINDOWS}name '_TreeView_GetStyleElements'{$ENDIF},
  TreeView_SetStyleElements {$IFNDEF MSWINDOWS}name '_TreeView_SetStyleElements'{$ENDIF},
  TreeView_SetOnChange {$IFNDEF MSWINDOWS}name '_TreeView_SetOnChange'{$ENDIF},
  TreeView_SetOnClick {$IFNDEF MSWINDOWS}name '_TreeView_SetOnClick'{$ENDIF},
  TreeView_SetOnDblClick {$IFNDEF MSWINDOWS}name '_TreeView_SetOnDblClick'{$ENDIF},
  TreeView_SetOnEnter {$IFNDEF MSWINDOWS}name '_TreeView_SetOnEnter'{$ENDIF},
  TreeView_SetOnExit {$IFNDEF MSWINDOWS}name '_TreeView_SetOnExit'{$ENDIF},
  TreeView_SetOnKeyDown {$IFNDEF MSWINDOWS}name '_TreeView_SetOnKeyDown'{$ENDIF},
  TreeView_SetOnKeyPress {$IFNDEF MSWINDOWS}name '_TreeView_SetOnKeyPress'{$ENDIF},
  TreeView_SetOnKeyUp {$IFNDEF MSWINDOWS}name '_TreeView_SetOnKeyUp'{$ENDIF},
  TreeView_SetOnMouseDown {$IFNDEF MSWINDOWS}name '_TreeView_SetOnMouseDown'{$ENDIF},
  TreeView_SetOnMouseEnter {$IFNDEF MSWINDOWS}name '_TreeView_SetOnMouseEnter'{$ENDIF},
  TreeView_SetOnMouseLeave {$IFNDEF MSWINDOWS}name '_TreeView_SetOnMouseLeave'{$ENDIF},
  TreeView_SetOnMouseMove {$IFNDEF MSWINDOWS}name '_TreeView_SetOnMouseMove'{$ENDIF},
  TreeView_SetOnMouseUp {$IFNDEF MSWINDOWS}name '_TreeView_SetOnMouseUp'{$ENDIF},
  TreeView_GetItems {$IFNDEF MSWINDOWS}name '_TreeView_GetItems'{$ENDIF},
  TreeView_SetItems {$IFNDEF MSWINDOWS}name '_TreeView_SetItems'{$ENDIF},
  TreeView_GetCanvas {$IFNDEF MSWINDOWS}name '_TreeView_GetCanvas'{$ENDIF},
  TreeView_GetDropTarget {$IFNDEF MSWINDOWS}name '_TreeView_GetDropTarget'{$ENDIF},
  TreeView_SetDropTarget {$IFNDEF MSWINDOWS}name '_TreeView_SetDropTarget'{$ENDIF},
  TreeView_GetSelected {$IFNDEF MSWINDOWS}name '_TreeView_GetSelected'{$ENDIF},
  TreeView_SetSelected {$IFNDEF MSWINDOWS}name '_TreeView_SetSelected'{$ENDIF},
  TreeView_GetTopItem {$IFNDEF MSWINDOWS}name '_TreeView_GetTopItem'{$ENDIF},
  TreeView_SetTopItem {$IFNDEF MSWINDOWS}name '_TreeView_SetTopItem'{$ENDIF},
  TreeView_GetSelectionCount {$IFNDEF MSWINDOWS}name '_TreeView_GetSelectionCount'{$ENDIF},
  TreeView_GetBrush {$IFNDEF MSWINDOWS}name '_TreeView_GetBrush'{$ENDIF},
  TreeView_GetControlCount {$IFNDEF MSWINDOWS}name '_TreeView_GetControlCount'{$ENDIF},
  TreeView_GetHandle {$IFNDEF MSWINDOWS}name '_TreeView_GetHandle'{$ENDIF},
  TreeView_GetParentWindow {$IFNDEF MSWINDOWS}name '_TreeView_GetParentWindow'{$ENDIF},
  TreeView_SetParentWindow {$IFNDEF MSWINDOWS}name '_TreeView_SetParentWindow'{$ENDIF},
  TreeView_GetAction {$IFNDEF MSWINDOWS}name '_TreeView_GetAction'{$ENDIF},
  TreeView_SetAction {$IFNDEF MSWINDOWS}name '_TreeView_SetAction'{$ENDIF},
  TreeView_GetBoundsRect {$IFNDEF MSWINDOWS}name '_TreeView_GetBoundsRect'{$ENDIF},
  TreeView_SetBoundsRect {$IFNDEF MSWINDOWS}name '_TreeView_SetBoundsRect'{$ENDIF},
  TreeView_GetClientHeight {$IFNDEF MSWINDOWS}name '_TreeView_GetClientHeight'{$ENDIF},
  TreeView_SetClientHeight {$IFNDEF MSWINDOWS}name '_TreeView_SetClientHeight'{$ENDIF},
  TreeView_GetClientRect {$IFNDEF MSWINDOWS}name '_TreeView_GetClientRect'{$ENDIF},
  TreeView_GetClientWidth {$IFNDEF MSWINDOWS}name '_TreeView_GetClientWidth'{$ENDIF},
  TreeView_SetClientWidth {$IFNDEF MSWINDOWS}name '_TreeView_SetClientWidth'{$ENDIF},
  TreeView_GetExplicitLeft {$IFNDEF MSWINDOWS}name '_TreeView_GetExplicitLeft'{$ENDIF},
  TreeView_GetExplicitTop {$IFNDEF MSWINDOWS}name '_TreeView_GetExplicitTop'{$ENDIF},
  TreeView_GetExplicitWidth {$IFNDEF MSWINDOWS}name '_TreeView_GetExplicitWidth'{$ENDIF},
  TreeView_GetExplicitHeight {$IFNDEF MSWINDOWS}name '_TreeView_GetExplicitHeight'{$ENDIF},
  TreeView_GetParent {$IFNDEF MSWINDOWS}name '_TreeView_GetParent'{$ENDIF},
  TreeView_SetParent {$IFNDEF MSWINDOWS}name '_TreeView_SetParent'{$ENDIF},
  TreeView_GetAlignWithMargins {$IFNDEF MSWINDOWS}name '_TreeView_GetAlignWithMargins'{$ENDIF},
  TreeView_SetAlignWithMargins {$IFNDEF MSWINDOWS}name '_TreeView_SetAlignWithMargins'{$ENDIF},
  TreeView_GetLeft {$IFNDEF MSWINDOWS}name '_TreeView_GetLeft'{$ENDIF},
  TreeView_SetLeft {$IFNDEF MSWINDOWS}name '_TreeView_SetLeft'{$ENDIF},
  TreeView_GetTop {$IFNDEF MSWINDOWS}name '_TreeView_GetTop'{$ENDIF},
  TreeView_SetTop {$IFNDEF MSWINDOWS}name '_TreeView_SetTop'{$ENDIF},
  TreeView_GetWidth {$IFNDEF MSWINDOWS}name '_TreeView_GetWidth'{$ENDIF},
  TreeView_SetWidth {$IFNDEF MSWINDOWS}name '_TreeView_SetWidth'{$ENDIF},
  TreeView_GetHeight {$IFNDEF MSWINDOWS}name '_TreeView_GetHeight'{$ENDIF},
  TreeView_SetHeight {$IFNDEF MSWINDOWS}name '_TreeView_SetHeight'{$ENDIF},
  TreeView_GetCursor {$IFNDEF MSWINDOWS}name '_TreeView_GetCursor'{$ENDIF},
  TreeView_SetCursor {$IFNDEF MSWINDOWS}name '_TreeView_SetCursor'{$ENDIF},
  TreeView_GetHint {$IFNDEF MSWINDOWS}name '_TreeView_GetHint'{$ENDIF},
  TreeView_SetHint {$IFNDEF MSWINDOWS}name '_TreeView_SetHint'{$ENDIF},
  TreeView_GetMargins {$IFNDEF MSWINDOWS}name '_TreeView_GetMargins'{$ENDIF},
  TreeView_SetMargins {$IFNDEF MSWINDOWS}name '_TreeView_SetMargins'{$ENDIF},
  TreeView_GetComponentCount {$IFNDEF MSWINDOWS}name '_TreeView_GetComponentCount'{$ENDIF},
  TreeView_GetComponentIndex {$IFNDEF MSWINDOWS}name '_TreeView_GetComponentIndex'{$ENDIF},
  TreeView_SetComponentIndex {$IFNDEF MSWINDOWS}name '_TreeView_SetComponentIndex'{$ENDIF},
  TreeView_GetOwner {$IFNDEF MSWINDOWS}name '_TreeView_GetOwner'{$ENDIF},
  TreeView_GetName {$IFNDEF MSWINDOWS}name '_TreeView_GetName'{$ENDIF},
  TreeView_SetName {$IFNDEF MSWINDOWS}name '_TreeView_SetName'{$ENDIF},
  TreeView_GetTag {$IFNDEF MSWINDOWS}name '_TreeView_GetTag'{$ENDIF},
  TreeView_SetTag {$IFNDEF MSWINDOWS}name '_TreeView_SetTag'{$ENDIF},
  TreeView_GetSelections {$IFNDEF MSWINDOWS}name '_TreeView_GetSelections'{$ENDIF},
  TreeView_GetControls {$IFNDEF MSWINDOWS}name '_TreeView_GetControls'{$ENDIF},
  TreeView_GetComponents {$IFNDEF MSWINDOWS}name '_TreeView_GetComponents'{$ENDIF};

