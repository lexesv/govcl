
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright ? ying32. All Rights Reserved.
//
//----------------------------------------

function TreeNodes_Create(AOwner: TCustomTreeView): TTreeNodes; cdecl;
begin
  Result :=  TTreeNodes.Create(AOwner);
end;

procedure TreeNodes_Free(AObj: TTreeNodes); cdecl;
begin
  AObj.Free;
end;

function TreeNodes_AddChildFirst(AObj: TTreeNodes; Parent: TTreeNode; S: PChar): TTreeNode; cdecl;
begin
  Result :=  AObj.AddChildFirst(Parent, S);
end;

function TreeNodes_AddChild(AObj: TTreeNodes; Parent: TTreeNode; S: PChar): TTreeNode; cdecl;
begin
  Result :=  AObj.AddChild(Parent, S);
end;

function TreeNodes_AddChildObjectFirst(AObj: TTreeNodes; Parent: TTreeNode; S: PChar; Ptr: Pointer): TTreeNode; cdecl;
begin
  Result :=  AObj.AddChildObjectFirst(Parent, S, Ptr);
end;

function TreeNodes_AddChildObject(AObj: TTreeNodes; Parent: TTreeNode; S: PChar; Ptr: Pointer): TTreeNode; cdecl;
begin
  Result :=  AObj.AddChildObject(Parent, S, Ptr);
end;

function TreeNodes_AddObjectFirst(AObj: TTreeNodes; Sibling: TTreeNode; S: PChar; Ptr: Pointer): TTreeNode; cdecl;
begin
  Result :=  AObj.AddObjectFirst(Sibling, S, Ptr);
end;

function TreeNodes_AddObject(AObj: TTreeNodes; Sibling: TTreeNode; S: PChar; Ptr: Pointer): TTreeNode; cdecl;
begin
  Result :=  AObj.AddObject(Sibling, S, Ptr);
end;

function TreeNodes_AddNode(AObj: TTreeNodes; Node: TTreeNode; Relative: TTreeNode; S: PChar; Ptr: Pointer; Method: TNodeAttachMode): TTreeNode; cdecl;
begin
  Result :=  AObj.AddNode(Node, Relative, S, Ptr, Method);
end;

function TreeNodes_AddFirst(AObj: TTreeNodes; Sibling: TTreeNode; S: PChar): TTreeNode; cdecl;
begin
  Result :=  AObj.AddFirst(Sibling, S);
end;

function TreeNodes_Add(AObj: TTreeNodes; Sibling: TTreeNode; S: PChar): TTreeNode; cdecl;
begin
  Result :=  AObj.Add(Sibling, S);
end;

function TreeNodes_AlphaSort(AObj: TTreeNodes; ARecurse: LongBool): LongBool; cdecl;
begin
  Result :=  AObj.AlphaSort(ARecurse);
end;

procedure TreeNodes_Assign(AObj: TTreeNodes; Source: TPersistent); cdecl;
begin
  AObj.Assign(Source);
end;

procedure TreeNodes_BeginUpdate(AObj: TTreeNodes); cdecl;
begin
  AObj.BeginUpdate;
end;

procedure TreeNodes_Clear(AObj: TTreeNodes); cdecl;
begin
  AObj.Clear;
end;

procedure TreeNodes_Delete(AObj: TTreeNodes; Node: TTreeNode); cdecl;
begin
  AObj.Delete(Node);
end;

procedure TreeNodes_EndUpdate(AObj: TTreeNodes); cdecl;
begin
  AObj.EndUpdate;
end;

function TreeNodes_GetFirstNode(AObj: TTreeNodes): TTreeNode; cdecl;
begin
  Result :=  AObj.GetFirstNode;
end;

function TreeNodes_GetNode(AObj: TTreeNodes; ItemId: HTREEITEM): TTreeNode; cdecl;
begin
  Result :=  AObj.GetNode(ItemId);
end;

function TreeNodes_Insert(AObj: TTreeNodes; Sibling: TTreeNode; S: PChar): TTreeNode; cdecl;
begin
  Result :=  AObj.Insert(Sibling, S);
end;

function TreeNodes_InsertObject(AObj: TTreeNodes; Sibling: TTreeNode; S: PChar; Ptr: Pointer): TTreeNode; cdecl;
begin
  Result :=  AObj.InsertObject(Sibling, S, Ptr);
end;

function TreeNodes_GetNamePath(AObj: TTreeNodes): PChar; cdecl;
begin
  Result :=  PChar(AObj.GetNamePath);
end;

function TreeNodes_ClassName(AObj: TTreeNodes): PChar; cdecl;
begin
  Result :=  PChar(AObj.ClassName);
end;

function TreeNodes_Equals(AObj: TTreeNodes; Obj: TObject): LongBool; cdecl;
begin
  Result :=  AObj.Equals(Obj);
end;

function TreeNodes_GetHashCode(AObj: TTreeNodes): Integer; cdecl;
begin
  Result :=  AObj.GetHashCode;
end;

function TreeNodes_ToString(AObj: TTreeNodes): PChar; cdecl;
begin
  Result :=  PChar(AObj.ToString);
end;

function TreeNodes_GetHandle(AObj: TTreeNodes): HWND; cdecl;
begin
  Result :=  AObj.Handle;
end;

function TreeNodes_GetOwner(AObj: TTreeNodes): TCustomTreeView; cdecl;
begin
  Result :=  AObj.Owner;
end;

function TreeNodes_GetItem(AObj: TTreeNodes; Index: Integer): TTreeNode; cdecl;
begin
  Result :=  AObj.Item[Index];
end;


exports
  TreeNodes_Create {$IFNDEF MSWINDOWS}name '_TreeNodes_Create'{$ENDIF},
  TreeNodes_Free {$IFNDEF MSWINDOWS}name '_TreeNodes_Free'{$ENDIF},
  TreeNodes_AddChildFirst {$IFNDEF MSWINDOWS}name '_TreeNodes_AddChildFirst'{$ENDIF},
  TreeNodes_AddChild {$IFNDEF MSWINDOWS}name '_TreeNodes_AddChild'{$ENDIF},
  TreeNodes_AddChildObjectFirst {$IFNDEF MSWINDOWS}name '_TreeNodes_AddChildObjectFirst'{$ENDIF},
  TreeNodes_AddChildObject {$IFNDEF MSWINDOWS}name '_TreeNodes_AddChildObject'{$ENDIF},
  TreeNodes_AddObjectFirst {$IFNDEF MSWINDOWS}name '_TreeNodes_AddObjectFirst'{$ENDIF},
  TreeNodes_AddObject {$IFNDEF MSWINDOWS}name '_TreeNodes_AddObject'{$ENDIF},
  TreeNodes_AddNode {$IFNDEF MSWINDOWS}name '_TreeNodes_AddNode'{$ENDIF},
  TreeNodes_AddFirst {$IFNDEF MSWINDOWS}name '_TreeNodes_AddFirst'{$ENDIF},
  TreeNodes_Add {$IFNDEF MSWINDOWS}name '_TreeNodes_Add'{$ENDIF},
  TreeNodes_AlphaSort {$IFNDEF MSWINDOWS}name '_TreeNodes_AlphaSort'{$ENDIF},
  TreeNodes_Assign {$IFNDEF MSWINDOWS}name '_TreeNodes_Assign'{$ENDIF},
  TreeNodes_BeginUpdate {$IFNDEF MSWINDOWS}name '_TreeNodes_BeginUpdate'{$ENDIF},
  TreeNodes_Clear {$IFNDEF MSWINDOWS}name '_TreeNodes_Clear'{$ENDIF},
  TreeNodes_Delete {$IFNDEF MSWINDOWS}name '_TreeNodes_Delete'{$ENDIF},
  TreeNodes_EndUpdate {$IFNDEF MSWINDOWS}name '_TreeNodes_EndUpdate'{$ENDIF},
  TreeNodes_GetFirstNode {$IFNDEF MSWINDOWS}name '_TreeNodes_GetFirstNode'{$ENDIF},
  TreeNodes_GetNode {$IFNDEF MSWINDOWS}name '_TreeNodes_GetNode'{$ENDIF},
  TreeNodes_Insert {$IFNDEF MSWINDOWS}name '_TreeNodes_Insert'{$ENDIF},
  TreeNodes_InsertObject {$IFNDEF MSWINDOWS}name '_TreeNodes_InsertObject'{$ENDIF},
  TreeNodes_GetNamePath {$IFNDEF MSWINDOWS}name '_TreeNodes_GetNamePath'{$ENDIF},
  TreeNodes_ClassName {$IFNDEF MSWINDOWS}name '_TreeNodes_ClassName'{$ENDIF},
  TreeNodes_Equals {$IFNDEF MSWINDOWS}name '_TreeNodes_Equals'{$ENDIF},
  TreeNodes_GetHashCode {$IFNDEF MSWINDOWS}name '_TreeNodes_GetHashCode'{$ENDIF},
  TreeNodes_ToString {$IFNDEF MSWINDOWS}name '_TreeNodes_ToString'{$ENDIF},
  TreeNodes_GetHandle {$IFNDEF MSWINDOWS}name '_TreeNodes_GetHandle'{$ENDIF},
  TreeNodes_GetOwner {$IFNDEF MSWINDOWS}name '_TreeNodes_GetOwner'{$ENDIF},
  TreeNodes_GetItem {$IFNDEF MSWINDOWS}name '_TreeNodes_GetItem'{$ENDIF};

