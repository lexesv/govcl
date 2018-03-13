package vcl

import (
	. "gitee.com/ying32/govcl/vcl/types"
)

// TNotifyEvent
type TNotifyEvent func(sender IObject)

// TUDClickEvent TUpDown  TUDBtnType
type TUDClickEvent func(sender IObject, button TUDBtnType)

// TCloseEvent Form
type TCloseEvent func(sender IObject, action *TCloseAction) // Action *uintptr

// TCloseQueryEvent Form
type TCloseQueryEvent func(sender IObject, canClose *bool) //CanClose *uintptr

// TMenuChangeEvent Menu
type TMenuChangeEvent func(sender IObject, source *TMenuItem, rebuild bool)

// TSysLinkEvent LinkLabel
type TSysLinkEvent func(sender IObject, link string, linkType TSysLinkType) // TSysLinkType

// TExceptionEvent TApplication
type TExceptionEvent func(sender IObject, e *Exception)

// TKeyEvent = procedure(Sender: TObject; var Key: Word; Shift: TShiftState)
type TKeyEvent func(sender IObject, key *Char, shift TShiftState)

// TKeyPressEvent = procedure(Sender: TObject; var Key: Char) of object;
type TKeyPressEvent func(sender IObject, key *Char)

// TMouseEvent = procedure(Sender: TObject; Button: TMouseButton;
//     Shift: TShiftState; X, Y: Integer) of object;
type TMouseEvent func(sender IObject, button TMouseButton, shift TShiftState, x, y int32)

// TMouseMoveEvent = procedure(Sender: TObject; Shift: TShiftState;
//    X, Y: Integer) of object;
type TMouseMoveEvent func(sender IObject, shift TShiftState, x, y int32)

// TMouseWheelEvent = procedure(Sender: TObject; Shift: TShiftState;
//    WheelDelta: Integer; MousePos: TPoint; var Handled: Boolean) of object;
type TMouseWheelEvent func(sender IObject, shift TShiftState, wheelDelta, x, y int32, handled *bool)

//  TDrawItemEvent = procedure(Control: TWinControl; Index: Integer;
//    Rect: TRect; State: TOwnerDrawState) of object;
type TDrawItemEvent func(control IControl, index int32, aRect TRect, state TOwnerDrawState)

//  TMenuDrawItemEvent = procedure (Sender: TObject; ACanvas: TCanvas;
//    ARect: TRect; Selected: Boolean) of object;
type TMenuDrawItemEvent func(sender IObject, aCanvas *TCanvas, aRect TRect, selected bool)

// ---------------TListView
// TLVColumnClickEvent = procedure(Sender: TObject; Column: TListColumn) of object;
type TLVColumnClickEvent func(sender IObject, column *TListColumn)

// TLVColumnRClickEvent = procedure(Sender: TObject; Column: TListColumn; Point: TPoint) of object;
type TLVColumnRClickEvent func(sender IObject, column *TListColumn, point TPoint)

// TLVSelectItemEvent = procedure(Sender: TObject; Item: TListItem;  Selected: Boolean) of object;
type TLVSelectItemEvent func(sender IObject, item *TListItem, selected bool)

// TLVCheckedItemEvent = procedure(Sender: TObject; Item: TListItem) of object;
type TLVCheckedItemEvent func(sender IObject, item *TListItem)

// TLVCompareEvent = procedure(Sender: TObject; Item1, Item2: TListItem;
// 	Data: Integer; var Compare: Integer) of object;
type TLVCompareEvent func(sender IObject, item1, item2 *TListItem, data int32, compare *int32)

// TLVChangeEvent TListView TTItemChange
type TLVChangeEvent func(sender IObject, item *TListItem, change TItemChange)

// TLVNotifyEvent = procedure(Sender: TObject; Item: TListItem) of object;
type TLVNotifyEvent func(sender IObject, item *TListItem)

//TLVAdvancedCustomDrawEvent = procedure(Sender: TCustomListView; const ARect: TRect;
//Stage: TCustomDrawStage; var DefaultDraw: Boolean) of object;
type TLVAdvancedCustomDrawEvent func(sender *TListView, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)

//TLVAdvancedCustomDrawItemEvent = procedure(Sender: TCustomListView; Item: TListItem;
//State: TCustomDrawState; Stage: TCustomDrawStage; var DefaultDraw: Boolean) of object;
type TLVAdvancedCustomDrawItemEvent func(sender *TListView, item *TListItem, state TCustomDrawState, Stage TCustomDrawStage, defaultDraw *bool)

//TLVAdvancedCustomDrawSubItemEvent = procedure(Sender: TCustomListView; Item: TListItem;
//SubItem: Integer; State: TCustomDrawState; Stage: TCustomDrawStage;
//var DefaultDraw: Boolean) of object;
type TLVAdvancedCustomDrawSubItemEvent func(sender *TListView, item *TListItem, subItem int32, state TCustomDrawState, stage TCustomDrawStage, defaultDraw *bool)

// ----------------TTreeView
// TTVCompareEvent = procedure(Sender: TObject; Node1, Node2: TTreeNode;
// 	Data: Integer; var Compare: Integer) of object;
type TTVCompareEvent func(sender IObject, node1, node2 *TTreeNode, data int32, compare *int32)

// TTVExpandedEvent = procedure(Sender: TObject; Node: TTreeNode) of object;
type TTVExpandedEvent func(sender IObject, node *TTreeNode)

// TTVChangedEvent TTreeView
type TTVChangedEvent func(sender IObject, node *TTreeNode)

//TTVAdvancedCustomDrawEvent = procedure(Sender: TCustomTreeView; const ARect: TRect;
//  Stage: TCustomDrawStage; var DefaultDraw: Boolean) of object;
type TTVAdvancedCustomDrawEvent func(sender *TTreeView, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)

//TTVAdvancedCustomDrawItemEvent = procedure(Sender: TCustomTreeView; Node: TTreeNode;
//  State: TCustomDrawState; Stage: TCustomDrawStage; var PaintImages,
//DefaultDraw: Boolean) of object;
type TTVAdvancedCustomDrawItemEvent func(sender *TTreeView, node *TTreeNode, state TCustomDrawState, stage TCustomDrawStage, paintImages, defaultDraw *bool)

// ----------------TPageControl
// TTabGetImageEvent = procedure(Sender: TObject; TabIndex: Integer; var ImageIndex: Integer) of object;
type TTabGetImageEvent func(sender IObject, tabIndex int32, imageIndex *int32)

// -------------------------TToolBar
//TTBAdvancedCustomDrawEvent = procedure(Sender: TToolBar; const ARect: TRect;
//  Stage: TCustomDrawStage; var DefaultDraw: Boolean) of object;
type TTBAdvancedCustomDrawEvent func(sender *TToolBar, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)

//TTBAdvancedCustomDrawBtnEvent = procedure(Sender: TToolBar; Button: TToolButton;
//  State: TCustomDrawState; Stage: TCustomDrawStage;
//  var Flags: TTBCustomDrawFlags; var DefaultDraw: Boolean) of object;
type TTBAdvancedCustomDrawBtnEvent func(sender *TToolBar, button *TToolButton, state TCustomDrawState, stage TCustomDrawStage, flags *TTBCustomDrawFlags, defaultDraw *bool)

// TThreadProc
type TThreadProc func()
