package api

var (
	application_Instance   = libvcl.NewProc("Application_Instance")
	application_CreateForm = libvcl.NewProc("Application_CreateForm")

	form_EnabledMaximize   = libvcl.NewProc("Form_EnabledMaximize")
	form_EnabledMinimize   = libvcl.NewProc("Form_EnabledMinimize")
	form_EnabledSystemMenu = libvcl.NewProc("Form_EnabledSystemMenu")

	setEventCallback = libvcl.NewProc("SetEventCallback")
	dGetParam        = libvcl.NewProc("DGetParam")
	dStrLen          = libvcl.NewProc("DStrLen")
	dMove            = libvcl.NewProc("DMove")

	dShowMessage     = libvcl.NewProc("DShowMessage")
	dGetMainInstance = libvcl.NewProc("DGetMainInstance")
	dMessageDlg      = libvcl.NewProc("DMessageDlg")

	mouse_Instance  = libvcl.NewProc("Mouse_Instance")
	screen_Instance = libvcl.NewProc("Screen_Instance")

	dSetReportMemoryLeaksOnShutdown = libvcl.NewProc("DSetReportMemoryLeaksOnShutdown")

	// TMenuItem
	dTextToShortCut = libvcl.NewProc("DTextToShortCut")
	dShortCutToText = libvcl.NewProc("DShortCutToText")

	// TClipboard
	clipboard_Instance     = libvcl.NewProc("Clipboard_Instance")
	clipboard_SetClipboard = libvcl.NewProc("Clipboard_SetClipboard")

	// DSysOpen
	dSysOpen = libvcl.NewProc("DSysOpen")

	// TMemoryStream
	memoryStream_Read     = libvcl.NewProc("MemoryStream_Read")
	memoryStream_Write    = libvcl.NewProc("MemoryStream_Write")
)
