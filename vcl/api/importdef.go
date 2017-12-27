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

	// TImageList
	imageList_Draw1        = libvcl.NewProc("ImageList_Draw1")
	imageList_Draw2        = libvcl.NewProc("ImageList_Draw2")
	imageList_DrawOverlay1 = libvcl.NewProc("ImageList_DrawOverlay1")
	imageList_DrawOverlay2 = libvcl.NewProc("ImageList_DrawOverlay2")
	imageList_GetIcon1     = libvcl.NewProc("ImageList_GetIcon1")
	imageList_GetIcon2     = libvcl.NewProc("ImageList_GetIcon2")

	dExtractFilePath = libvcl.NewProc("DExtractFilePath")
	dFileExists      = libvcl.NewProc("DFileExists")

	dInheritsFromControl    = libvcl.NewProc("DInheritsFromControl")
	dInheritsFromWinControl = libvcl.NewProc("DInheritsFromWinControl")
	dInheritsFromComponent  = libvcl.NewProc("DInheritsFromComponent")

	// TForm相关设置
	setGlobalFormScaled      = libvcl.NewProc("SetGlobalFormScaled")
	form_ScaleForPPI         = libvcl.NewProc("Form_ScaleForPPI")
	form_ScaleControlsForDpi = libvcl.NewProc("Form_ScaleControlsForDpi")

	// TStyleManager
	styleManager_IsValidStyle        = libvcl.NewProc("StyleManager_IsValidStyle")
	styleManager_LoadFromFile        = libvcl.NewProc("StyleManager_LoadFromFile")
	styleManager_CheckSysClassName   = libvcl.NewProc("StyleManager_CheckSysClassName")
	styleManager_TrySetStyle         = libvcl.NewProc("StyleManager_TrySetStyle")
	styleManager_SetStyle1           = libvcl.NewProc("StyleManager_SetStyle1")
	styleManager_SetStyle2           = libvcl.NewProc("StyleManager_SetStyle2")
	styleManager_TryLoadFromResource = libvcl.NewProc("StyleManager_TryLoadFromResource")

	styleManager_ActiveStyle         = libvcl.NewProc("StyleManager_ActiveStyle")
	styleManager_SystemStyle         = libvcl.NewProc("StyleManager_SystemStyle")
	styleManager_Enabled             = libvcl.NewProc("StyleManager_Enabled")
	styleManager_IsCustomStyleActive = libvcl.NewProc("StyleManager_IsCustomStyleActive")
	styleManager_UnRegisterStyle     = libvcl.NewProc("StyleManager_UnRegisterStyle")
	styleManager_RegisterStyle       = libvcl.NewProc("StyleManager_RegisterStyle")
	styleManager_Style               = libvcl.NewProc("StyleManager_Style")
	styleManager_StyleDescriptor     = libvcl.NewProc("StyleManager_StyleDescriptor")
)
