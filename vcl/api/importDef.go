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
	memoryStream_Read  = libvcl.NewProc("MemoryStream_Read")
	memoryStream_Write = libvcl.NewProc("MemoryStream_Write")

	// TCanvas
	canvas_BrushCopy     = libvcl.NewProc("Canvas_BrushCopy")
	canvas_CopyRect      = libvcl.NewProc("Canvas_CopyRect")
	canvas_Draw1         = libvcl.NewProc("Canvas_Draw1")
	canvas_Draw2         = libvcl.NewProc("Canvas_Draw2")
	canvas_DrawFocusRect = libvcl.NewProc("Canvas_DrawFocusRect")
	canvas_FillRect      = libvcl.NewProc("Canvas_FillRect")
	canvas_FrameRect     = libvcl.NewProc("Canvas_FrameRect")
	canvas_StretchDraw   = libvcl.NewProc("Canvas_StretchDraw")
	canvas_TextRect1     = libvcl.NewProc("Canvas_TextRect1")
	canvas_TextRect2     = libvcl.NewProc("Canvas_TextRect2")
)
