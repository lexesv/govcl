package vcl

import (
	"unsafe"

	. "gitee.com/ying32/govcl/vcl/api"
	. "gitee.com/ying32/govcl/vcl/types"
)

// 回调过程
func callbackProc(f uintptr, args uintptr, argcount int) uintptr {
	v, ok := CallbackMap[f]
	if ok {

		getVal := func(i int) uintptr {
			return DGetParam(i, args).Value
		}

		switch v.(type) {
		// func(sender IObject)
		case TNotifyEvent:
			v.(TNotifyEvent)(ObjectFromInst(getVal(0)))

		// func(sender IObject, button TUDBtnType)
		case TUDClickEvent:
			v.(TUDClickEvent)(ObjectFromInst(getVal(0)), TUDBtnType(getVal(1)))

		// func(sender IObject, item *TListItem, change int32)
		case TLVChangeEvent:
			v.(TLVChangeEvent)(ObjectFromInst(getVal(0)),
				ListItemFromInst(getVal(1)), TItemChange(getVal(2)))

		// func(sender IObject, action *TCloseAction) // Action *uintptr
		case TCloseEvent:
			v.(TCloseEvent)(ObjectFromInst(getVal(0)),
				(*TCloseAction)(unsafe.Pointer(getVal(1))))

		// func(sender IObject, canClose *bool) //CanClose *uintptr
		case TCloseQueryEvent:
			v.(TCloseQueryEvent)(ObjectFromInst(getVal(0)), (*bool)(unsafe.Pointer(getVal(1))))

		// func(sender IObject, source *TMenuItem, rebuild bool)
		case TMenuChangeEvent:
			v.(TMenuChangeEvent)(ObjectFromInst(getVal(0)),
				MenuItemFromInst(getVal(1)), DBoolToGoBool(getVal(2)))

		// func(sender IObject, node *TreeNode)
		case TTVChangedEvent:
			v.(TTVChangedEvent)(ObjectFromInst(getVal(0)), TreeNodeFromInst(getVal(1)))

		// func(sender IObject, link string, linkType TSysLinkType) // TSysLinkType
		case TSysLinkEvent:
			v.(TSysLinkEvent)(ObjectFromInst(getVal(0)), DStrToGoStr(getVal(1)),
				TSysLinkType(getVal(2)))

		// func(sender, e IObject)
		case TExceptionEvent:
			v.(TExceptionEvent)(ObjectFromInst(getVal(0)), ObjectFromInst(getVal(1)))

		// func(sender IObject, key *Char, shift TShiftState)
		case TKeyEvent:
			v.(TKeyEvent)(ObjectFromInst(getVal(0)),
				(*Char)(unsafe.Pointer(getVal(1))), TShiftState(getVal(2)))

		// func(sender IObject, key *Char)
		case TKeyPressEvent:
			v.(TKeyPressEvent)(ObjectFromInst(getVal(0)),
				(*Char)(unsafe.Pointer(getVal(1))))

		// func(sender IObject, button TMouseButton, shift TShiftState, x, y int32)
		case TMouseEvent:
			v.(TMouseEvent)(ObjectFromInst(getVal(0)), TMouseButton(getVal(1)),
				TShiftState(getVal(2)), int32(getVal(3)), int32(getVal(4)))

		// func(sender IObject, shift TShiftState, x, y int32)
		case TMouseMoveEvent:
			v.(TMouseMoveEvent)(ObjectFromInst(getVal(0)), TShiftState(getVal(1)),
				int32(getVal(2)), int32(getVal(3)))

		// func(sender IObject, shift TShiftState, wheelDelta, x, y int32, handled *bool)
		case TMouseWheelEvent:
			v.(TMouseWheelEvent)(ObjectFromInst(getVal(0)), TShiftState(getVal(1)),
				int32(getVal(2)), int32(getVal(3)), int32(getVal(4)),
				(*bool)(unsafe.Pointer(getVal(5))))

			// func(control IControl, index int32, aRect TRect, state TOwnerDrawState)
		case TDrawItemEvent:
			v.(TDrawItemEvent)(ControlFromInst(getVal(0)),
				int32(getVal(1)), *(*TRect)(unsafe.Pointer(getVal(2))),
				TOwnerDrawState(getVal(3)))

			// func(sender IObject, aCanvas *TCanvas, aRect TRect, selected bool)
		case TMenuDrawItemEvent:
			v.(TMenuDrawItemEvent)(ObjectFromInst(getVal(0)),
				CanvasFromInst(getVal(1)), *(*TRect)(unsafe.Pointer(getVal(2))),
				DBoolToGoBool(getVal(3)))

			// type TLVNotifyEvent func(sender IObject, item *TListItem)
		case TLVNotifyEvent:
			v.(TLVNotifyEvent)(ObjectFromInst(getVal(0)), ListItemFromInst(getVal(1)))

			// type TLVColumnClickEvent func(sender IObject, column *TListColumn)
		case TLVColumnClickEvent:
			v.(TLVColumnClickEvent)(ObjectFromInst(getVal(0)), ListColumnFromInst(getVal(1)))

			// type TLVColumnRClickEvent func(sender IObject, column *TListColumn, point TPoint)
		case TLVColumnRClickEvent:
			v.(TLVColumnRClickEvent)(ObjectFromInst(getVal(0)), ListColumnFromInst(getVal(1)), TPoint{X: int32(getVal(2)), Y: int32(getVal(3))})

			// type TLVSelectItemEvent func(sender IObject, item *TListItem, selected bool)
		case TLVSelectItemEvent:
			v.(TLVSelectItemEvent)(ObjectFromInst(getVal(0)), ListItemFromInst(getVal(1)), DBoolToGoBool(getVal(2)))

			// type TLVCheckedItemEvent func(sender IObject, item *TListItem)
		case TLVCheckedItemEvent:
			v.(TLVCheckedItemEvent)(ObjectFromInst(getVal(0)), ListItemFromInst(getVal(1)))

			// type TTabGetImageEvent func(sender IObject, tabIndex int32, imageIndex *int32)
		case TTabGetImageEvent:
			v.(TTabGetImageEvent)(ObjectFromInst(getVal(0)), int32(getVal(1)), (*int32)(unsafe.Pointer(getVal(2))))

			// type TTVExpandedEvent func(sender IObject, node *TTreeNode)
		case TTVExpandedEvent:
			v.(TTVExpandedEvent)(ObjectFromInst(getVal(0)), TreeNodeFromInst(getVal(1)))

		//type TLVCompareEvent func(sender IObject, item1, item2 *TListItem, data int32, compare *int32)
		case TLVCompareEvent:
			v.(TLVCompareEvent)(ObjectFromInst(getVal(0)), ListItemFromInst(getVal(1)), ListItemFromInst(getVal(2)),
				int32(getVal(3)), (*int32)(unsafe.Pointer(getVal(4))))

		//type TTVCompareEvent func(sender IObject, node1, node2 *TTreeNode, data int32, compare *int32)
		case TTVCompareEvent:
			v.(TTVCompareEvent)(ObjectFromInst(getVal(0)), TreeNodeFromInst(getVal(1)), TreeNodeFromInst(getVal(2)),
				int32(getVal(3)), (*int32)(unsafe.Pointer(getVal(4))))

		default:
		}
	}
	return 0
}
