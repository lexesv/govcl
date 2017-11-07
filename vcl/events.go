package vcl

import (
	. "gitee.com/ying32/govcl/vcl/types"
)

// TNotifyEvent
type TNotifyEvent func(sender IObject)

// TUDClickEvent TUpDown  TUDBtnType
type TUDClickEvent func(sender IObject, button TUDBtnType)

// TLVChangeEvent TListView TTItemChange
type TLVChangeEvent func(sender IObject, item *TListItem, change TItemChange)

// TCloseEvent Form
type TCloseEvent func(sender IObject, action *TCloseAction) // Action *uintptr

// TCloseQueryEvent Form
type TCloseQueryEvent func(sender IObject, canClose *bool) //CanClose *uintptr

// TMenuChangeEvent Menu
type TMenuChangeEvent func(sender IObject, source *TMenuItem, rebuild bool)

// TTVChangedEvent TTreeView
type TTVChangedEvent func(sender IObject, node *TTreeNode)

// TSysLinkEvent LinkLabel
type TSysLinkEvent func(sender IObject, link string, linkType TSysLinkType) // TSysLinkType

// TExceptionEvent TApplication
type TExceptionEvent func(sender, e IObject)

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

// TLVNotifyEvent = procedure(Sender: TObject; Item: TListItem) of object;
type TLVNotifyEvent func(sender IObject, item *TListItem)

// TLVColumnClickEvent = procedure(Sender: TObject; Column: TListColumn) of object;
type TLVColumnClickEvent func(sender IObject, column *TListColumn)

// TLVColumnRClickEvent = procedure(Sender: TObject; Column: TListColumn; Point: TPoint) of object;
type TLVColumnRClickEvent func(sender IObject, column *TListColumn, point TPoint)

// TLVSelectItemEvent = procedure(Sender: TObject; Item: TListItem;  Selected: Boolean) of object;
type TLVSelectItemEvent func(sender IObject, item *TListItem, selected bool)

// TLVCheckedItemEvent = procedure(Sender: TObject; Item: TListItem) of object;
type TLVCheckedItemEvent func(sender IObject, item *TListItem)

// TTabGetImageEvent = procedure(Sender: TObject; TabIndex: Integer; var ImageIndex: Integer) of object;
type TTabGetImageEvent func(sender IObject, tabIndex int32, imageIndex *int32)

// TTVExpandedEvent = procedure(Sender: TObject; Node: TTreeNode) of object;
type TTVExpandedEvent func(sender IObject, node *TTreeNode)
