package vcl

import (
	//	"reflect"
	"unsafe"

	. "gitee.com/ying32/govcl/vcl/api"
	. "gitee.com/ying32/govcl/vcl/types"
)

//const (
//	// TVarRec.vType
//	vtInteger       = 0 //iota + 0
//	vtBoolean       = 1
//	vtChar          = 2
//	vtExtended      = 3
//	vtString        = 4
//	vtPointer       = 5
//	vtPChar         = 6
//	vtObject        = 7
//	vtClass         = 8
//	vtWideChar      = 9
//	vtPWideChar     = 10
//	vtAnsiString    = 11
//	vtCurrency      = 12
//	vtVariant       = 13
//	vtInterface     = 14
//	vtWideString    = 15
//	vtInt64         = 16
//	vtUnicodeString = 17
//)

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

		// func(sender IObject, item IObject, change int32)
		case TLVChangeEvent:
			v.(TLVChangeEvent)(ObjectFromInst(getVal(0)), ObjectFromInst(getVal(1)), int32(getVal(2)))

		// func(sender IObject, action *TCloseAction) // Action *uintptr
		case TCloseEvent:
			v.(TCloseEvent)(ObjectFromInst(getVal(0)), (*TCloseAction)(unsafe.Pointer(getVal(1))))

		// func(sender IObject, canClose *bool) //CanClose *uintptr
		case TCloseQueryEvent:
			v.(TCloseQueryEvent)(ObjectFromInst(getVal(0)), (*bool)(unsafe.Pointer(getVal(1))))

		// func(sender IObject, source IObject, rebuild bool)
		case TMenuChangeEvent:
			v.(TMenuChangeEvent)(ObjectFromInst(getVal(0)), ObjectFromInst(getVal(1)), DBoolToGoBool(getVal(2)))

		// func(sender IObject, node IObject)
		case TTVChangedEvent:
			v.(TTVChangedEvent)(ObjectFromInst(getVal(0)), ObjectFromInst(getVal(1)))

		// func(sender IObject, link string, linkType TSysLinkType) // TSysLinkType
		case TSysLinkEvent:
			v.(TSysLinkEvent)(ObjectFromInst(getVal(0)), DStrToGoStr(getVal(1)), TSysLinkType(getVal(2)))

		// func(sender, e IObject)
		case TExceptionEvent:
			v.(TExceptionEvent)(ObjectFromInst(getVal(0)), ObjectFromInst(getVal(1)))

		// func(sender IObject, key *Char, shift TShiftState)
		case TKeyEvent:
			v.(TKeyEvent)(ObjectFromInst(getVal(0)), (*Char)(unsafe.Pointer(getVal(1))), TShiftState(getVal(2)))

		// func(sender IObject, key *Char)
		case TKeyPressEvent:
			v.(TKeyPressEvent)(ObjectFromInst(getVal(0)), (*Char)(unsafe.Pointer(getVal(1))))

		// func(sender IObject, button TMouseButton, shift TShiftState, x, y int32)
		case TMouseEvent:
			v.(TMouseEvent)(ObjectFromInst(getVal(0)), TMouseButton(getVal(1)), TShiftState(getVal(2)), int32(getVal(3)), int32(getVal(4)))

		// func(sender IObject, shift TShiftState, x, y int32)
		case TMouseMoveEvent:
			v.(TMouseMoveEvent)(ObjectFromInst(getVal(0)), TShiftState(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		// func(sender IObject, shift TShiftState, wheelDelta, x, y int32, handled *bool)
		case TMouseWheelEvent:
			v.(TMouseWheelEvent)(ObjectFromInst(getVal(0)), TShiftState(getVal(1)), int32(getVal(2)), int32(getVal(3)), int32(getVal(4)),
				(*bool)(unsafe.Pointer(getVal(5))))
		default:
		}

		//
		//		params := make([]reflect.Value, argcount)
		//		for i := 0; i < argcount; i++ {
		//			var p TGoParam = DGetParam(i, args)

		//			switch p.Type {
		//			case vtInteger:
		//				params[i] = reflect.ValueOf(int32(p.Value))
		//			case vtBoolean:
		//				params[i] = reflect.ValueOf(DBoolToGoBool(p.Value))
		//			case vtChar:
		//				params[i] = reflect.ValueOf(uint8(p.Value))
		//			case vtExtended:
		//				params[i] = reflect.ValueOf(*(*float64)(unsafe.Pointer(p.Value)))
		//			case vtPointer:
		//				params[i] = reflect.ValueOf(p.Value)
		//			case vtString, vtPChar, vtPWideChar, vtWideString, vtUnicodeString:
		//				params[i] = reflect.ValueOf(DStrToGoStr(p.Value))
		//			case vtObject, vtClass, vtInterface:
		//				params[i] = reflect.ValueOf(ObjectFromInst(p.Value))
		//			case vtWideChar:
		//				params[i] = reflect.ValueOf(uint16(p.Value))
		//				//			case vtAnsiString:
		//				//			case vtCurrency:
		//				//			case vtVariant:
		//			//case :
		//			//	params[i] = reflect.ValueOf(p.Value)
		//			case vtInt64:
		//				params[i] = reflect.ValueOf(*(*int64)(unsafe.Pointer(p.Value)))

		//			default:
		//				params[i] = reflect.ValueOf(p.Value)
		//			}
		//		}
		//		rv := reflect.ValueOf(v)
		//		rv.Call(params)
	}
	return 0
}
