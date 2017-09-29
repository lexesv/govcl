package vcl

import (
	"reflect"
	"unsafe"

	. "gitee.com/ying32/govcl/vcl/api"
)

const (
	// TVarRec.vType
	vtInteger       = 0 //iota + 0
	vtBoolean       = 1
	vtChar          = 2
	vtExtended      = 3
	vtString        = 4
	vtPointer       = 5
	vtPChar         = 6
	vtObject        = 7
	vtClass         = 8
	vtWideChar      = 9
	vtPWideChar     = 10
	vtAnsiString    = 11
	vtCurrency      = 12
	vtVariant       = 13
	vtInterface     = 14
	vtWideString    = 15
	vtInt64         = 16
	vtUnicodeString = 17
)

// 回调过程
func callbackProc(f uintptr, args uintptr, argcount int) uintptr {
	v, ok := CallbackMap[f]
	if ok {
		params := make([]reflect.Value, argcount)
		for i := 0; i < argcount; i++ {
			var p TGoParam = DGetParam(i, args)

			switch p.Type {
			case vtInteger:
				params[i] = reflect.ValueOf(int32(p.Value))
			case vtBoolean:
				params[i] = reflect.ValueOf(DBoolToGoBool(p.Value))
			case vtChar:
				params[i] = reflect.ValueOf(uint8(p.Value))
			case vtExtended:
				params[i] = reflect.ValueOf(*(*float64)(unsafe.Pointer(p.Value)))
			case vtPointer:
				params[i] = reflect.ValueOf(p.Value)
			case vtString, vtPChar, vtPWideChar, vtWideString, vtUnicodeString:
				params[i] = reflect.ValueOf(DStrToGoStr(p.Value))
			case vtObject, vtClass, vtInterface:
				params[i] = reflect.ValueOf(ObjectFromInst(p.Value))
			case vtWideChar:
				params[i] = reflect.ValueOf(uint16(p.Value))
				//			case vtAnsiString:
				//			case vtCurrency:
				//			case vtVariant:
			//case :
			//	params[i] = reflect.ValueOf(p.Value)
			case vtInt64:
				params[i] = reflect.ValueOf(*(*int64)(unsafe.Pointer(p.Value)))

			default:
				params[i] = reflect.ValueOf(p.Value)
			}
		}
		rv := reflect.ValueOf(v)
		rv.Call(params)
	}
	return 0
}
