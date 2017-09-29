package vcl

type TNotifyEvent func(sender IObject)

// TUpDown  TUDBtnType
type TUDClickEvent func(Sender IObject, Button int32)

// TListView TTItemChange
type TLVChangeEvent func(Sender IObject, Item IObject, Change int32)

// Form
type TCloseEvent func(Sender IObject, Action uintptr)        // Action *uintptr
type TCloseQueryEvent func(Sender IObject, CanClose uintptr) //CanClose *uintptr

// Menu
type TMenuChangeEvent func(Sender IObject, Source IObject, Rebuild bool)

// TTreeView
type TTVChangedEvent func(Sender IObject, Node IObject)

// TLinkLabel
type TSysLinkEvent = func(Sender IObject, Link string, LinkType int32) // TSysLinkType

// TApplication
type TExceptionEvent = func(Sender, E IObject)

// procedure(Sender: TObject; var Key: Word; Shift: TShiftState)
type TKeyEvent = func(Sender IObject, Key uintptr, Shift int32)

// TKeyPressEvent = procedure(Sender: TObject; var Key: Char) of object;
type TKeyPressEvent = func(Sender IObject, Key uintptr)

// TMouseEvent = procedure(Sender: TObject; Button: TMouseButton;
//     Shift: TShiftState; X, Y: Integer) of object;
type TMouseEvent = func(Sender IObject, Button, Shift, X, Y int32)

// TMouseMoveEvent = procedure(Sender: TObject; Shift: TShiftState;
//    X, Y: Integer) of object;
type TMouseMoveEvent = func(Sender IObject, Shift, X, Y int32)

// TMouseWheelEvent = procedure(Sender: TObject; Shift: TShiftState;
//    WheelDelta: Integer; MousePos: TPoint; var Handled: Boolean) of object;
type TMouseWheelEvent = func(Sender IObject, Shift, WheelDelta, X, Y int32, Handled uintptr)
