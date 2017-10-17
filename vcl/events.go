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
type TExceptionEvent = func(sender, e IObject)

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
