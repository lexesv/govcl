package api

func Form_EnabledMaximize(obj uintptr, val bool) {
	form_EnabledMaximize.Call(obj, GoBoolToDBool(val))
}

func Form_EnabledMinimize(obj uintptr, val bool) {
	form_EnabledMinimize.Call(obj, GoBoolToDBool(val))
}

func Form_EnabledSystemMenu(obj uintptr, val bool) {
	form_EnabledSystemMenu.Call(obj, GoBoolToDBool(val))
}

func Form_ScaleForPPI(obj uintptr, val int32) {
	form_ScaleForPPI.Call(obj, uintptr(val))
}

func Form_ScaleControlsForDpi(obj uintptr, val int32) {
	form_ScaleControlsForDpi.Call(obj, uintptr(val))
}


// SetGlobalFormScaled 下面两个函数放在Application下面吧，直接调用实例类
func SetGlobalFormScaled(val bool) {
	setGlobalFormScaled.Call(GoBoolToDBool(val))
}