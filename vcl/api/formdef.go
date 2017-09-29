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
