
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function TreeNode_Create(AOwner: TTreeNodes): TTreeNode; cdecl;
begin
  Result :=  TTreeNode.Create(AOwner);
end;

procedure TreeNode_Free(AObj: TTreeNode); cdecl;
begin
  AObj.Free;
end;

function TreeNode_AlphaSort(AObj: TTreeNode; ARecurse: LongBool): LongBool; cdecl;
begin
  Result :=  AObj.AlphaSort(ARecurse);
end;

procedure TreeNode_Assign(AObj: TTreeNode; Source: TPersistent); cdecl;
begin
  AObj.Assign(Source);
end;

procedure TreeNode_Collapse(AObj: TTreeNode; Recurse: LongBool); cdecl;
begin
  AObj.Collapse(Recurse);
end;

procedure TreeNode_Delete(AObj: TTreeNode); cdecl;
begin
  AObj.Delete;
end;

procedure TreeNode_Expand(AObj: TTreeNode; Recurse: LongBool); cdecl;
begin
  AObj.Expand(Recurse);
end;

function TreeNode_IndexOf(AObj: TTreeNode; Value: TTreeNode): Integer; cdecl;
begin
  Result :=  AObj.IndexOf(Value);
end;

procedure TreeNode_MakeVisible(AObj: TTreeNode); cdecl;
begin
  AObj.MakeVisible;
end;

procedure TreeNode_MoveTo(AObj: TTreeNode; Destination: TTreeNode; Mode: TNodeAttachMode); cdecl;
begin
  AObj.MoveTo(Destination, Mode);
end;

function TreeNode_GetNamePath(AObj: TTreeNode): PChar; cdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function TreeNode_ClassName(AObj: TTreeNode): PChar; cdecl;
begin
  Result :=  PChar(AObj.ClassName);
end;

function TreeNode_Equals(AObj: TTreeNode; Obj: TObject): LongBool; cdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function TreeNode_GetHashCode(AObj: TTreeNode): Integer; cdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function TreeNode_ToString(AObj: TTreeNode): PChar; cdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function TreeNode_GetAbsoluteIndex(AObj: TTreeNode): Integer; cdecl;
begin
  Result :=  AObj.AbsoluteIndex;
end;

function TreeNode_GetCut(AObj: TTreeNode): LongBool; cdecl;
begin
  Result :=  AObj.Cut;
end;

procedure TreeNode_SetCut(AObj: TTreeNode; AValue: LongBool); cdecl;
begin
  AObj.Cut := AValue;
end;

function TreeNode_GetData(AObj: TTreeNode): Pointer; cdecl;
begin
  Result :=  AObj.Data;
end;

procedure TreeNode_SetData(AObj: TTreeNode; AValue: Pointer); cdecl;
begin
  AObj.Data := AValue;
end;

function TreeNode_GetDeleting(AObj: TTreeNode): LongBool; cdecl;
begin
  Result :=  AObj.Deleting;
end;

function TreeNode_GetFocused(AObj: TTreeNode): LongBool; cdecl;
begin
  Result :=  AObj.Focused;
end;

procedure TreeNode_SetFocused(AObj: TTreeNode; AValue: LongBool); cdecl;
begin
  AObj.Focused := AValue;
end;

function TreeNode_GetDropTarget(AObj: TTreeNode): LongBool; cdecl;
begin
  Result :=  AObj.DropTarget;
end;

procedure TreeNode_SetDropTarget(AObj: TTreeNode; AValue: LongBool); cdecl;
begin
  AObj.DropTarget := AValue;
end;

function TreeNode_GetSelected(AObj: TTreeNode): LongBool; cdecl;
begin
  Result :=  AObj.Selected;
end;

procedure TreeNode_SetSelected(AObj: TTreeNode; AValue: LongBool); cdecl;
begin
  AObj.Selected := AValue;
end;

function TreeNode_GetExpanded(AObj: TTreeNode): LongBool; cdecl;
begin
  Result :=  AObj.Expanded;
end;

procedure TreeNode_SetExpanded(AObj: TTreeNode; AValue: LongBool); cdecl;
begin
  AObj.Expanded := AValue;
end;

function TreeNode_GetExpandedImageIndex(AObj: TTreeNode): TImageIndex; cdecl;
begin
  Result :=  AObj.ExpandedImageIndex;
end;

procedure TreeNode_SetExpandedImageIndex(AObj: TTreeNode; AValue: TImageIndex); cdecl;
begin
  AObj.ExpandedImageIndex := AValue;
end;

function TreeNode_GetHandle(AObj: TTreeNode): HWND; cdecl;
begin
  Result :=  AObj.Handle;
end;

function TreeNode_GetHasChildren(AObj: TTreeNode): LongBool; cdecl;
begin
  Result :=  AObj.HasChildren;
end;

procedure TreeNode_SetHasChildren(AObj: TTreeNode; AValue: LongBool); cdecl;
begin
  AObj.HasChildren := AValue;
end;

function TreeNode_GetImageIndex(AObj: TTreeNode): TImageIndex; cdecl;
begin
  Result :=  AObj.ImageIndex;
end;

procedure TreeNode_SetImageIndex(AObj: TTreeNode; AValue: TImageIndex); cdecl;
begin
  AObj.ImageIndex := AValue;
end;

function TreeNode_GetIndex(AObj: TTreeNode): Integer; cdecl;
begin
  Result :=  AObj.Index;
end;

function TreeNode_GetIsVisible(AObj: TTreeNode): LongBool; cdecl;
begin
  Result :=  AObj.IsVisible;
end;

function TreeNode_GetItemId(AObj: TTreeNode): HTREEITEM; cdecl;
begin
  Result :=  AObj.ItemId;
end;

function TreeNode_GetLevel(AObj: TTreeNode): Integer; cdecl;
begin
  Result :=  AObj.Level;
end;

function TreeNode_GetOverlayIndex(AObj: TTreeNode): Integer; cdecl;
begin
  Result :=  AObj.OverlayIndex;
end;

procedure TreeNode_SetOverlayIndex(AObj: TTreeNode; AValue: Integer); cdecl;
begin
  AObj.OverlayIndex := AValue;
end;

function TreeNode_GetOwner(AObj: TTreeNode): TTreeNodes; cdecl;
begin
  Result :=  AObj.Owner;
end;

function TreeNode_GetParent(AObj: TTreeNode): TTreeNode; cdecl;
begin
  Result :=  AObj.Parent;
end;

function TreeNode_GetSelectedIndex(AObj: TTreeNode): Integer; cdecl;
begin
  Result :=  AObj.SelectedIndex;
end;

procedure TreeNode_SetSelectedIndex(AObj: TTreeNode; AValue: Integer); cdecl;
begin
  AObj.SelectedIndex := AValue;
end;

function TreeNode_GetEnabled(AObj: TTreeNode): LongBool; cdecl;
begin
  Result :=  AObj.Enabled;
end;

procedure TreeNode_SetEnabled(AObj: TTreeNode; AValue: LongBool); cdecl;
begin
  AObj.Enabled := AValue;
end;

function TreeNode_GetStateIndex(AObj: TTreeNode): Integer; cdecl;
begin
  Result :=  AObj.StateIndex;
end;

procedure TreeNode_SetStateIndex(AObj: TTreeNode; AValue: Integer); cdecl;
begin
  AObj.StateIndex := AValue;
end;

function TreeNode_GetText(AObj: TTreeNode): PChar; cdecl;
begin
  Result :=  PChar(AObj.Text);
end;

procedure TreeNode_SetText(AObj: TTreeNode; AValue: PChar); cdecl;
begin
  AObj.Text := AValue;
end;

function TreeNode_GetItem(AObj: TTreeNode; Index: Integer): TTreeNode; cdecl;
begin
  Result :=  AObj.Item[Index];
end;

procedure TreeNode_SetItem(AObj: TTreeNode; Index: Integer; AValue: TTreeNode); cdecl;
begin
  AObj.Item[Index] := AValue;
end;


exports
  TreeNode_Create {$IFNDEF MSWINDOWS}name '_TreeNode_Create'{$ENDIF},
  TreeNode_Free {$IFNDEF MSWINDOWS}name '_TreeNode_Free'{$ENDIF},
  TreeNode_AlphaSort {$IFNDEF MSWINDOWS}name '_TreeNode_AlphaSort'{$ENDIF},
  TreeNode_Assign {$IFNDEF MSWINDOWS}name '_TreeNode_Assign'{$ENDIF},
  TreeNode_Collapse {$IFNDEF MSWINDOWS}name '_TreeNode_Collapse'{$ENDIF},
  TreeNode_Delete {$IFNDEF MSWINDOWS}name '_TreeNode_Delete'{$ENDIF},
  TreeNode_Expand {$IFNDEF MSWINDOWS}name '_TreeNode_Expand'{$ENDIF},
  TreeNode_IndexOf {$IFNDEF MSWINDOWS}name '_TreeNode_IndexOf'{$ENDIF},
  TreeNode_MakeVisible {$IFNDEF MSWINDOWS}name '_TreeNode_MakeVisible'{$ENDIF},
  TreeNode_MoveTo {$IFNDEF MSWINDOWS}name '_TreeNode_MoveTo'{$ENDIF},
  TreeNode_GetNamePath {$IFNDEF MSWINDOWS}name '_TreeNode_GetNamePath'{$ENDIF},
  TreeNode_ClassName {$IFNDEF MSWINDOWS}name '_TreeNode_ClassName'{$ENDIF},
  TreeNode_Equals {$IFNDEF MSWINDOWS}name '_TreeNode_Equals'{$ENDIF},
  TreeNode_GetHashCode {$IFNDEF MSWINDOWS}name '_TreeNode_GetHashCode'{$ENDIF},
  TreeNode_ToString {$IFNDEF MSWINDOWS}name '_TreeNode_ToString'{$ENDIF},
  TreeNode_GetAbsoluteIndex {$IFNDEF MSWINDOWS}name '_TreeNode_GetAbsoluteIndex'{$ENDIF},
  TreeNode_GetCut {$IFNDEF MSWINDOWS}name '_TreeNode_GetCut'{$ENDIF},
  TreeNode_SetCut {$IFNDEF MSWINDOWS}name '_TreeNode_SetCut'{$ENDIF},
  TreeNode_GetData {$IFNDEF MSWINDOWS}name '_TreeNode_GetData'{$ENDIF},
  TreeNode_SetData {$IFNDEF MSWINDOWS}name '_TreeNode_SetData'{$ENDIF},
  TreeNode_GetDeleting {$IFNDEF MSWINDOWS}name '_TreeNode_GetDeleting'{$ENDIF},
  TreeNode_GetFocused {$IFNDEF MSWINDOWS}name '_TreeNode_GetFocused'{$ENDIF},
  TreeNode_SetFocused {$IFNDEF MSWINDOWS}name '_TreeNode_SetFocused'{$ENDIF},
  TreeNode_GetDropTarget {$IFNDEF MSWINDOWS}name '_TreeNode_GetDropTarget'{$ENDIF},
  TreeNode_SetDropTarget {$IFNDEF MSWINDOWS}name '_TreeNode_SetDropTarget'{$ENDIF},
  TreeNode_GetSelected {$IFNDEF MSWINDOWS}name '_TreeNode_GetSelected'{$ENDIF},
  TreeNode_SetSelected {$IFNDEF MSWINDOWS}name '_TreeNode_SetSelected'{$ENDIF},
  TreeNode_GetExpanded {$IFNDEF MSWINDOWS}name '_TreeNode_GetExpanded'{$ENDIF},
  TreeNode_SetExpanded {$IFNDEF MSWINDOWS}name '_TreeNode_SetExpanded'{$ENDIF},
  TreeNode_GetExpandedImageIndex {$IFNDEF MSWINDOWS}name '_TreeNode_GetExpandedImageIndex'{$ENDIF},
  TreeNode_SetExpandedImageIndex {$IFNDEF MSWINDOWS}name '_TreeNode_SetExpandedImageIndex'{$ENDIF},
  TreeNode_GetHandle {$IFNDEF MSWINDOWS}name '_TreeNode_GetHandle'{$ENDIF},
  TreeNode_GetHasChildren {$IFNDEF MSWINDOWS}name '_TreeNode_GetHasChildren'{$ENDIF},
  TreeNode_SetHasChildren {$IFNDEF MSWINDOWS}name '_TreeNode_SetHasChildren'{$ENDIF},
  TreeNode_GetImageIndex {$IFNDEF MSWINDOWS}name '_TreeNode_GetImageIndex'{$ENDIF},
  TreeNode_SetImageIndex {$IFNDEF MSWINDOWS}name '_TreeNode_SetImageIndex'{$ENDIF},
  TreeNode_GetIndex {$IFNDEF MSWINDOWS}name '_TreeNode_GetIndex'{$ENDIF},
  TreeNode_GetIsVisible {$IFNDEF MSWINDOWS}name '_TreeNode_GetIsVisible'{$ENDIF},
  TreeNode_GetItemId {$IFNDEF MSWINDOWS}name '_TreeNode_GetItemId'{$ENDIF},
  TreeNode_GetLevel {$IFNDEF MSWINDOWS}name '_TreeNode_GetLevel'{$ENDIF},
  TreeNode_GetOverlayIndex {$IFNDEF MSWINDOWS}name '_TreeNode_GetOverlayIndex'{$ENDIF},
  TreeNode_SetOverlayIndex {$IFNDEF MSWINDOWS}name '_TreeNode_SetOverlayIndex'{$ENDIF},
  TreeNode_GetOwner {$IFNDEF MSWINDOWS}name '_TreeNode_GetOwner'{$ENDIF},
  TreeNode_GetParent {$IFNDEF MSWINDOWS}name '_TreeNode_GetParent'{$ENDIF},
  TreeNode_GetSelectedIndex {$IFNDEF MSWINDOWS}name '_TreeNode_GetSelectedIndex'{$ENDIF},
  TreeNode_SetSelectedIndex {$IFNDEF MSWINDOWS}name '_TreeNode_SetSelectedIndex'{$ENDIF},
  TreeNode_GetEnabled {$IFNDEF MSWINDOWS}name '_TreeNode_GetEnabled'{$ENDIF},
  TreeNode_SetEnabled {$IFNDEF MSWINDOWS}name '_TreeNode_SetEnabled'{$ENDIF},
  TreeNode_GetStateIndex {$IFNDEF MSWINDOWS}name '_TreeNode_GetStateIndex'{$ENDIF},
  TreeNode_SetStateIndex {$IFNDEF MSWINDOWS}name '_TreeNode_SetStateIndex'{$ENDIF},
  TreeNode_GetText {$IFNDEF MSWINDOWS}name '_TreeNode_GetText'{$ENDIF},
  TreeNode_SetText {$IFNDEF MSWINDOWS}name '_TreeNode_SetText'{$ENDIF},
  TreeNode_GetItem {$IFNDEF MSWINDOWS}name '_TreeNode_GetItem'{$ENDIF},
  TreeNode_SetItem {$IFNDEF MSWINDOWS}name '_TreeNode_SetItem'{$ENDIF};

