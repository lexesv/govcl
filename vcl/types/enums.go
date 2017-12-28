package types

/*
  注意：Delphi中所有集合这里全部使用uint32表示，也就是说最多32个元素

*/

// TAlign = (alNone, alTop, alBottom, alLeft, alRight, alClient, alCustom);
type TAlign int32

const (
	AlNone = iota + 0
	AlTop
	AlBottom
	AlLeft
	AlRight
	AlClient
	AlCustom
)

//  TFormBorderStyle = (bsNone, bsSingle, bsSizeable, bsDialog, bsToolWindow,
//    bsSizeToolWin);
//  TBorderStyle = bsNone..bsSingle;
type TBorderStyle int32

const (
	BsNone = iota + 0
	BsSingle
	BsSizeable
	BsDialog
	BsToolWindow
	BsSizeToolWin
)

type TFormBorderStyle TBorderStyle

//  TFormStyle = (fsNormal, fsMDIChild, fsMDIForm, fsStayOnTop);
type TFormStyle int32

const (
	FsNormal = iota + 0
	FsMDIChild
	FsMDIForm
	FsStayOnTop
)

//  TPosition = (poDesigned, poDefault, poDefaultPosOnly, poDefaultSizeOnly,
//    poScreenCenter, poDesktopCenter, poMainFormCenter, poOwnerFormCenter);
type TPosition int32

const (
	PoDesigned = iota + 0
	PoDefault
	PoDefaultPosOnly
	PoDefaultSizeOnly
	PoScreenCenter
	PoDesktopCenter
	PoMainFormCenter
	PoOwnerFormCenter
)

//  TCursor = -32768..32767;
type TCursor int16

const (
	CrDefault   = 0
	CrNone      = -1
	CrArrow     = -2
	CrCross     = -3
	CrIBeam     = -4
	CrSize      = -22
	CrSizeNESW  = -6
	CrSizeNS    = -7
	CrSizeNWSE  = -8
	CrSizeWE    = -9
	CrUpArrow   = -10
	CrHourGlass = -11
	CrDrag      = -12
	CrNoDrop    = -13
	CrHSplit    = -14
	CrVSplit    = -15
	CrMultiDrag = -16
	CrSQLWait   = -17
	CrNo        = -18
	CrAppStart  = -19
	CrHelp      = -20
	CrHandPoint = -21
	CrSizeAll   = -22
)

// TSeekOrigin = (soBeginning, soCurrent, soEnd);
type TSeekOrigin int32

const (
	SoBeginning = iota + 0
	SoCurrent
	SoEnd
)

// TPixelFormat = (pfDevice, pf1bit, pf4bit, pf8bit, pf15bit, pf16bit, pf24bit, pf32bit, pfCustom);
type TPixelFormat int32

const (
	PfDevice = iota + 0
	Pf1bit
	Pf4bit
	Pf8bit
	Pf15bit
	Pf16bit
	Pf24bit
	Pf32bit
	PfCustom
)

// TBalloonHintStyle = (bhsStandard, bhsBalloon);
type TBalloonHintStyle int32

const (
	bhsStandard = iota + 0
	bhsBalloon
)

//TAlignment = (taLeftJustify, taRightJustify, taCenter);
type TAlignment int32

const (
	TaLeftJustify = iota + 0
	TaRightJustify
	TaCenter
)

//  TLeftRight = TAlignment.taLeftJustify..TAlignment.taRightJustify;
type TLeftRight int32

//  TBiDiMode = (bdLeftToRight, bdRightToLeft, bdRightToLeftNoAlign,
//    bdRightToLeftReadingOnly);
type TBiDiMode int32

const (
	BdLeftToRight = iota + 0
	BdRightToLeft
	BdRightToLeftNoAlign
	BdRightToLeftReadingOnly
)

//  TVerticalAlignment = (taAlignTop, taAlignBottom, taVerticalCenter);
type TVerticalAlignment int32

const (
	TaAlignTop = iota + 0
	TaAlignBottom
	TaVerticalCenter
)

// TButtonStyle = (bsPushButton, bsCommandLink, bsSplitButton);
type TButtonStyle int32

const (
	BsPushButton = iota + 0
	BsCommandLink
	BsSplitButton
)

type TColorBoxStyle int32

//  TComboBoxStyle = (csDropDown, csSimple, csDropDownList, csOwnerDrawFixed,
//    csOwnerDrawVariable);
type TComboBoxStyle int32

const (
	CsDropDown = iota + 0
	CsSimple
	CsDropDownList
	CsOwnerDrawFixed
	CsOwnerDrawVariable
)

//  TWindowState = (wsNormal, wsMinimized, wsMaximized);
type TWindowState int32

const (
	WsNormal = iota + 0
	WsMinimized
	WsMaximized
)

//  TTextLayout = (tlTop, tlCenter, tlBottom);
type TTextLayout int32

const (
	TlTop = iota + 0
	TlCenter
	TlBottom
)

//  TEllipsisPosition = (epNone, epPathEllipsis, epEndEllipsis, epWordEllipsis);
type TEllipsisPosition int32

const (
	EpNone = iota + 0
	EpPathEllipsis
	EpEndEllipsis
	EpWordEllipsis
)

type TLinkAlignment TAlignment

//  TListBoxStyle = (lbStandard, lbOwnerDrawFixed, lbOwnerDrawVariable,
//    lbVirtual, lbVirtualOwnerDraw);
type TListBoxStyle int32

const (
	LbStandard = iota + 0
	LbOwnerDrawFixed
	LbOwnerDrawVariable
	LbVirtual
	LbVirtualOwnerDraw
)

//TMenuItemAutoFlag = (maAutomatic, maManual, maParent);
type TMenuItemAutoFlag int32

const (
	MaAutomatic = iota + 0
	MaManual
	MaParent
)

//  TMenuAutoFlag = maAutomatic..maManual;
type TMenuAutoFlag TMenuItemAutoFlag

//TPopupAlignment = (paLeft, paRight, paCenter);
type TPopupAlignment int32

const (
	PaLeft = iota + 0
	PaRight
	PaCenter
)

//  TTrackButton = (tbRightButton, tbLeftButton);
type TTrackButton int32

const (
	TbRightButton = iota + 0
	TbLeftButton
)

//  TProgressBarOrientation = (pbHorizontal, pbVertical);
type TProgressBarOrientation int32

const (
	PbHorizontal = iota + 0
	PbVertical
)

//  TProgressBarStyle = (pbstNormal, pbstMarquee);
type TProgressBarStyle int32

const (
	PbstNormal = iota + 0
	PbstMarquee
)

//  TProgressBarState = (pbsNormal, pbsError, pbsPaused);
type TProgressBarState int32

const (
	PbsNormal = iota + 0
	PbsError
	PbsPaused
)

//TButtonLayout = (blGlyphLeft, blGlyphRight, blGlyphTop, blGlyphBottom);
type TButtonLayout int32

const (
	BlGlyphLeft = iota + 0
	BlGlyphRight
	BlGlyphTop
	BlGlyphBottom
)

//  TButtonState = (bsUp, bsDisabled, bsDown, bsExclusive);
type TButtonState int32

const (
	BsUp = iota + 0
	BsDisabled
	BsDown
	BsExclusive
)

//  TButtonStyle = (bsAutoDetect, bsWin31, bsNew);
//type TButtonStyle int32

const (
	BsAutoDetect = iota + 0
	BsWin31
	BsNew
)

//  TNumGlyphs = 1..4;
type TNumGlyphs int32

// TStaticBorderStyle = (sbsNone, sbsSingle, sbsSunken);
type TStaticBorderStyle int32

const (
	SbsNone = iota + 0
	sbsSingle
	sbsSunken
)

// TFontStyle = (fsBold, fsItalic, fsUnderline, fsStrikeOut);
type TFontStyle int32

const (
	FsBold = iota + 0
	FsItalic
	FsUnderline
	FsStrikeOut
)

type TFontStyles uint32

// TScrollStyle = (ssNone, ssHorizontal, ssVertical, ssBoth);
type TScrollStyle int32

const (
	SsNone = iota + 0
	SsHorizontal
	SsVertical
	SsBoth
)

//TSortType = (stNone, stData, stText, stBoth);
type TSortType int32

const (
	StNone = iota + 0
	StData
	StText
	StBoth
)

//  TMultiSelectStyles = (msControlSelect, msShiftSelect,
//                        msVisibleOnly, msSiblingOnly);
type TMultiSelectStyles int32

const (
	MsControlSelect = iota + 0
	MsShiftSelect
	MsVisibleOnly
	MsSiblingOnly
)

//  TListArrangement = (arAlignBottom, arAlignLeft, arAlignRight,
//    arAlignTop, arDefault, arSnapToGrid);
type TListArrangement int32

const (
	ArAlignBottom = iota + 0
	ArAlignLeft
	ArAlignRight
	ArAlignTop
	ArDefault
	ArSnapToGrid
)

//  TViewStyle = (vsIcon, vsSmallIcon, vsList, vsReport);
type TViewStyle int32

const (
	VsIcon = iota + 0
	VsSmallIcon
	VsList
	VsReport
)

//  TItemState = (isNone, isCut, isDropHilited, isFocused, isSelected, isActivating);
type TItemState int32

const (
	IsNone = iota + 0
	IsCut
	IsDropHilited
	IsFocused
	IsSelected
	IsActivating
)

//  TItemStates = set of TItemState;
type TItemStates uint32

//  TItemChange = (ctText, ctImage, ctState);
type TItemChange int32

const (
	CtText = iota + 0
	CtImage
	CtState
)

//  TItemFind = (ifData, ifPartialString, ifExactString, ifNearest);
type TItemFind int32

const (
	IfData = iota + 0
	IfPartialString
	IfExactString
	IfNearest
)

//  TSearchDirection = (sdLeft, sdRight, sdAbove, sdBelow, sdAll);
type TSearchDirection int32

const (
	SdLeft = iota + 0
	SdRight
	SdAbove
	SdBelow
	SdAll
)

//  TListHotTrackStyle = (htHandPoint, htUnderlineCold, htUnderlineHot);
type TListHotTrackStyle int32

const (
	HtHandPoint = iota + 0
	HtUnderlineCold
	HtUnderlineHot
)

//  TListHotTrackStyles = set of TListHotTrackStyle;
type TListHotTrackStyles uint32

//  TItemRequests = (irText, irImage, irParam, irState, irIndent);
type TItemRequests int32

const (
	IrText = iota + 0
	IrImage
	IrParam
	IrState
	IrIndent
)

//TBrushStyle = (bsSolid, bsClear, bsHorizontal, bsVertical,
//    bsFDiagonal, bsBDiagonal, bsCross, bsDiagCross);
type TBrushStyle int32

const (
	BsSolid = iota + 0
	BsClear
	BsHorizontal
	BsVertical
	BsFDiagonal
	BsBDiagonal
	BsCross
	BsDiagCross
)

//TPenStyle = (psSolid, psDash, psDot, psDashDot, psDashDotDot, psClear,
//    psInsideFrame, psUserStyle, psAlternate);
type TPenStyle int32

const (
	PsSolid = iota + 0
	PsDash
	PsDot
	PsDashDot
	PsDashDotDot
	PsClear
	PsInsideFrame
	PsUserStyle
	PsAlternate
)

// TUDBtnType = (btNext, btPrev);
type TUDBtnType int32

const (
	btNext = iota + 0
	btPrev
)

//  TTabPosition = (tpTop, tpBottom, tpLeft, tpRight);
type TTabPosition int32

const (
	TpTop = iota + 0
	TpBottom
	TpLeft
	TpRight
)

//  TTabStyle = (tsTabs, tsButtons, tsFlatButtons);
type TTabStyle int32

const (
	TsTabs = iota + 0
	TsButtons
	TsFlatButtons
)

// TFontPitch = (fpDefault, fpVariable, fpFixed);
type TFontPitch int32

const (
	FpDefault = iota + 0
	FpVariable
	FpFixed
)

// TPenMode = (pmBlack, pmWhite, pmNop, pmNot, pmCopy, pmNotCopy,
//    pmMergePenNot, pmMaskPenNot, pmMergeNotPen, pmMaskNotPen, pmMerge,
//    pmNotMerge, pmMask, pmNotMask, pmXor, pmNotXor);
type TPenMode int32

const (
	PmBlack = iota + 0
	PmWhite
	PmNop
	PmNot
	PmCopy
	PmNotCopy
	PmMergePenNot
	PmMaskPenNot
	PmMergeNotPen
	PmMaskNotPen
	PmMerge
	PmNotMerge
	PmMask
	PmNotMask
	PmXor
	PmNotXor
)

// TTrackBarOrientation = (trHorizontal, trVertical);
type TTrackBarOrientation int32

const (
	TrHorizontal = iota + 0
	TrVertical
)

// TUDOrientation = (udHorizontal, udVertical);
type TUDOrientation int32

const (
	UdHorizontal = iota + 0
	UdVertical
)

//  TFontQuality = (fqDefault, fqDraft, fqProof, fqNonAntialiased, fqAntialiased,
//    fqClearType, fqClearTypeNatural);
type TFontQuality int32

const (
	FqDefault = iota + 0
	FqDraft
	FqProof
	FqNonAntialiased
	FqAntialiased
	FqClearType
	FqClearTypeNatural
)

// TCloseAction = (caNone, caHide, caFree, caMinimize);
type TCloseAction int32

const (
	CaNone = iota + 0
	CaHide
	CaFree
	CaMinimize
)

//  TBalloonFlags = (bfNone = NIIF_NONE, bfInfo = NIIF_INFO,
//    bfWarning = NIIF_WARNING, bfError = NIIF_ERROR);
type TBalloonFlags int32

const (
	BfNone    = 0
	BfInfo    = 1
	BfWarning = 2
	BfError   = 3
)

//  TMsgDlgType = (mtWarning, mtError, mtInformation, mtConfirmation, mtCustom);
type TMsgDlgType int32

const (
	MtWarning = iota + 0
	MtError
	MtInformation
	MtConfirmation
	MtCustom
)

//  TMsgDlgBtn = (mbYes, mbNo, mbOK, mbCancel, mbAbort, mbRetry, mbIgnore,
//    mbAll, mbNoToAll, mbYesToAll, mbHelp, mbClose);
type TMsgDlgBtn int32

const (
	MbYes = iota + 0
	MbNo
	MbOK
	MbCancel
	MbAbort
	MbRetry
	MbIgnore
	MbAll
	MbNoToAll
	MbYesToAll
	MbHelp
	MbClose
)

//  TMsgDlgButtons = set of TMsgDlgBtn;
type TMsgDlgButtons int32

// TSysLinkType = (sltURL, sltID);
type TSysLinkType int32

const (
	SltURL = iota + 0
	SltID
)

//  TStatusPanelStyle = (psText, psOwnerDraw);
type TStatusPanelStyle int32

const (
	PsText = iota + 0
	PsOwnerDraw
)

//  TStatusPanelBevel = (pbNone, pbLowered, pbRaised);
type TStatusPanelBevel int32

const (
	PbNone = iota + 0
	PbLowered
	PbRaised
)

//  TJPEGQualityRange = 1..100;   // 100 = best quality, 25 = pretty awful
type TJPEGQualityRange uint32

//  TJPEGPerformance = (jpBestQuality, jpBestSpeed);
type TJPEGPerformance int32

const (
	JpBestQuality = iota + 0
	JpBestSpeed
)

//  TJPEGScale = (jsFullSize, jsHalf, jsQuarter, jsEighth);
type TJPEGScale int32

const (
	JsFullSize = iota + 0
	JsHalf
	JsQuarter
	JsEighth
)

//  TJPEGPixelFormat = (jf24Bit, jf8Bit);
type TJPEGPixelFormat int32

const (
	Jf24Bit = iota + 0
	Jf8Bit
)

//  TGIFVersion = (gvUnknown, gv87a, gv89a);
type TGIFVersion int32

const (
	GvUnknown = iota + 0
	Gv87a
	Gv89a
)

// Animation loop behaviour
//TGIFAnimationLoop = (
//  glDisabled,                 // Never loop
//  glEnabled,                  // Loop is specified in GIF
//  glContinously               // Loop continously regardless of GIF
//  );
type TGIFAnimationLoop int32

const (
	GlDisabled = iota + 0
	GlEnabled
	GlContinously
)

// Auto dithering of GIF output to Netscape 216 color palette
//TGIFDithering = (
//  gdDisabled,                 // Never dither
//  gdEnabled,                  // Always dither
//  gdAuto                      // Dither if Desktop DC supports <= 256 colors.
//  );
type TGIFDithering int32

const (
	GdDisabled = iota + 0
	GdEnabled
	GdAuto
)

// TCompressionLevel = 0..9;
type TCompressionLevel uint32

type TShortCut uint16

//  TNodeState = (nsCut, nsDropHilited, nsFocused, nsSelected, nsExpanded);
type TNodeState int32

const (
	NsCut = iota + 0
	NsDropHilited
	NsFocused
	NsSelected
	NsExpanded
)

//  TNodeAttachMode = (naAdd, naAddFirst, naAddChild, naAddChildFirst, naInsert);
type TNodeAttachMode int32

const (
	NaAdd = iota + 0
	NaAddFirst
	NaAddChild
	NaAddChildFirst
	naInsert
)

//  TAddMode = (taAddFirst, taAdd, taInsert);
type TAddMode int32

const (
	TaAddFirst = iota + 0
	TaAdd
	TaInsert
)

type TMultiSelectStyle int32

// TActionListState = (asNormal, asSuspended, asSuspendedEnabled);
type TActionListState int32

const (
	AsNormal = iota + 0
	AsSuspended
	AsSuspendedEnabled
)

// TGradientDirection = (gdHorizontal, gdVertical);
type TGradientDirection int32

const (
	GdHorizontal = iota + 0
	GdVertical
)

// TDrawingStyle = (dsFocus, dsSelected, dsNormal, dsTransparent);
type TDrawingStyle int32

const (
	DSFocus = iota + 0
	DSSelected
	DSNormal
	DSTransparent
)

// TImageType = (itImage, itMask);
type TImageType int32

const (
	ItImage = iota + 0
	ItMask
)

// TResType = (rtBitmap, rtCursor, rtIcon);
type TResType int32

const (
	RtBitmap = iota + 0
	RtCursor
	RtIcon
)

// TOverlay = 0..14;

// TLoadResource = (lrDefaultColor, lrDefaultSize, lrFromFile, lrMap3DColors, lrTransparent, lrMonoChrome);
type TLoadResource int32

const (
	LrDefaultColor = iota + 0
	LrDefaultSize
	LrFromFile
	LrMap3DColors
	LrTransparent
	LrMonoChrome
)

//   TLoadResources = set of TLoadResource;
type TLoadResources uint32

// TColorDepth = (cdDefault, cdDeviceDependent, cd4Bit, cd8Bit, cd16Bit, cd24Bit, cd32Bit);
type TColorDepth int32

const (
	CdDefault = iota + 0
	CdDeviceDependent
	Cd4Bit
	Cd8Bit
	Cd16Bit
	Cd24Bit
	Cd32Bit
)

// TCheckBoxState = (cbUnchecked, cbChecked, cbGrayed);
type TCheckBoxState int32

const (
	CbUnchecked = iota + 0
	CbChecked
	CbGrayed
)

//  TListGroupState = (
//    lgsNormal,
//    lgsHidden,
//    lgsCollapsed,
//    lgsNoHeader,
//    lgsCollapsible,
//    lgsFocused,
//    lgsSelected,
//    lgsSubseted,
//    lgsSubSetLinkFocused
//  );
type TListGroupState int32

const (
	LgsNormal = iota + 0
	LgsHidden
	LgsCollapsed
	LgsNoHeader
	LgsCollapsible
	LgsFocused
	LgsSelected
	LgsSubseted
	LgsSubSetLinkFocused
)

//  TListGroupStateSet = set of TListGroupState;
type TListGroupStateSet uint32

// TTBDrawingStyle = (dsNormal, dsGradient);
type TTBDrawingStyle int32

const (
	DsNormal = iota + 0
	DsGradient
)

//TToolButtonStyle = (tbsButton, tbsCheck, tbsDropDown, tbsSeparator,
//    tbsDivider, tbsTextButton);
type TToolButtonStyle int32

const (
	TbsButton = iota + 0
	TbsCheck
	TbsDropDown
	TbsSeparator
	TbsDivider
	TbsTextButton
)

//  TToolButtonState = (tbsChecked, tbsPressed, tbsEnabled, tbsHidden,
//    tbsIndeterminate, tbsWrap, tbsEllipses, tbsMarked);
type TToolButtonState int32

const (
	TbsChecked = iota + 0
	TbsPressed
	TbsEnabled
	TbsHidden
	TbsIndeterminate
	TbsWrap
	TbsEllipses
	TbsMarked
)

// TTBGradientDrawingOption = (gdoHotTrack, gdoGradient)
type TTBGradientDrawingOption int32

const (
	GdoHotTrack = iota + 0
	GdoGradient
)

// TTBGradientDrawingOptions = set of (gdoHotTrack, gdoGradient);
type TTBGradientDrawingOptions uint32

// TColorDialogOption = (cdFullOpen, cdPreventFullOpen, cdShowHelp,
// cdSolidColor, cdAnyColor);
type TColorDialogOption int32

const (
	CdFullOpen = iota + 0
	CdPreventFullOpen
	CdShowHelp
	CdSolidColor
	CdAnyColor
)

type TColorDialogOptions uint32

// TBorderIcon = (biSystemMenu, biMinimize, biMaximize, biHelp);
type TBorderIcon int32

const (
	BiSystemMenu = iota + 0
	BiMinimize
	BiMaximize
	BiHelp
)

// TBorderIcons = set of TBorderIcon;
type TBorderIcons uint32

// TFontDialogOption = (fdAnsiOnly, fdTrueTypeOnly, fdEffects,
//     fdFixedPitchOnly, fdForceFontExist, fdNoFaceSel, fdNoOEMFonts,
//     fdNoSimulations, fdNoSizeSel, fdNoStyleSel,  fdNoVectorFonts,
//     fdShowHelp, fdWysiwyg, fdLimitSize, fdScalableOnly, fdApplyButton);
type TFontDialogOption int32

const (
	FdAnsiOnly = iota + 0
	FdTrueTypeOnly
	FdEffects
	FdFixedPitchOnly
	FdForceFontExist
	FdNoFaceSel
	FdNoOEMFonts
	FdNoSimulations
	FdNoSizeSel
	FdNoStyleSel
	FdNoVectorFonts
	FdShowHelp
	FdWysiwyg
	FdLimitSize
	FdScalableOnly
	FdApplyButton
)

//   TFontDialogOptions = set of TFontDialogOption;
type TFontDialogOptions uint32

// { TOpenDialog }
// TOpenOption = (ofReadOnly, ofOverwritePrompt, ofHideReadOnly,
//   ofNoChangeDir, ofShowHelp, ofNoValidate, ofAllowMultiSelect,
//   ofExtensionDifferent, ofPathMustExist, ofFileMustExist, ofCreatePrompt,
//   ofShareAware, ofNoReadOnlyReturn, ofNoTestFileCreate, ofNoNetworkButton,
//   ofNoLongNames, ofOldStyleDialog, ofNoDereferenceLinks, ofEnableIncludeNotify,
//   ofEnableSizing, ofDontAddToRecent, ofForceShowHidden);
type TOpenOption int32

const (
	OfReadOnly = iota + 0
	OfOverwritePrompt
	OfHideReadOnly
	OfNoChangeDir
	OfShowHelp
	OfNoValidate
	OfAllowMultiSelect
	OfExtensionDifferent
	OfPathMustExist
	OfFileMustExist
	OfCreatePrompt
	OfShareAware
	OfNoReadOnlyReturn
	OfNoTestFileCreate
	OfNoNetworkButton
	OfNoLongNames
	OfOldStyleDialog
	OfNoDereferenceLinks
	OfEnableIncludeNotify
	OfEnableSizing
	OfDontAddToRecent
	OfForceShowHidden
)

// TOpenOptions = set of TOpenOption;
type TOpenOptions uint32

// TOpenOptionEx = (ofExNoPlacesBar);
type TOpenOptionEx int32

const (
	OfExNoPlacesBar = iota + 0
)

// TOpenOptionsEx = set of TOpenOptionEx;
type TOpenOptionsEx uint32

// { TPrintDialog }
// TPrintRange = (prAllPages, prSelection, prPageNums);
type TPrintRange int32

const (
	PrAllPages = iota + 0
	PrSelection
	PrPageNums
)

// TPrintDialogOption = (poPrintToFile, poPageNums, poSelection, poWarning,
//   poHelp, poDisablePrintToFile);
type TPrintDialogOption int32

const (
	PoPrintToFile = iota + 0
	PoPageNums
	PoSelection
	PoWarning
	PoHelp
	PoDisablePrintToFile
)

// TPrintDialogOptions = set of TPrintDialogOption;
type TPrintDialogOptions uint32

// TPageSetupDialogOption = (psoDefaultMinMargins, psoDisableMargins,
//   psoDisableOrientation, psoDisablePagePainting, psoDisablePaper, psoDisablePrinter,
//   psoMargins, psoMinMargins, psoShowHelp, psoWarning, psoNoNetworkButton);
type TPageSetupDialogOption int32

const (
	PsoDefaultMinMargins = iota + 0
	PsoDisableMargins
	PsoDisableOrientation
	PsoDisablePagePainting
	PsoDisablePaper
	PsoDisablePrinter
	PsoMargins
	PsoMinMargins
	PsoShowHelp
	PsoWarning
	PsoNoNetworkButton
)

// TPageSetupDialogOptions = set of TPageSetupDialogOption;
type TPageSetupDialogOptions uint32

// TPrinterKind = (pkDotMatrix, pkHPPCL);
type TPrinterKind int32

const (
	PkDotMatrix = iota + 0
	PkHPPCL
)

// TPageType = (ptEnvelope, ptPaper);
type TPageType int32

const (
	PtEnvelope = iota + 0
	PtPaper
)

// TPageMeasureUnits = (pmDefault, pmMillimeters, pmInches);
type TPageMeasureUnits int32

const (
	PmDefault = iota + 0
	PmMillimeters
	PmInches
)

//   TStringsOption = (soStrictDelimiter, soWriteBOM, soTrailingLineBreak,
//     soUseLocale);
type TStringsOption uint32

const (
	SoStrictDelimiter = iota + 0
	SoWriteBOM
	SoTrailingLineBreak
	SoUseLocale
)

//   TStringsOptions = set of TStringsOption;
type TStringsOptions uint32

// TShiftState = set of (ssShift, ssAlt, ssCtrl,
//     ssLeft, ssRight, ssMiddle, ssDouble, ssTouch, ssPen, ssCommand, ssHorizontal);
type TShiftState uint32

const (
	SsShift = iota + 0
	SsAlt
	SsCtrl
	SsLeft
	SsRight
	SsMiddle
	SsDouble
	SsTouch
	SsPen
	SsCommand
	SssHorizontal // 有冲突，所以加了个s
)

// TMouseButton = (mbLeft, mbRight, mbMiddle);
type TMouseButton int32

const (
	MbLeft = iota + 0
	MbRight
	MbMiddle
)

// TFillStyle = (fsSurface, fsBorder);
type TFillStyle int32

const (
	FsSurface = iota + 0
	FsBorder
)

// TFillMode = (fmAlternate, fmWinding);
type TFillMode int32

const (
	FmAlternate = iota + 0
	FmWinding
)

// TCopyMode = Longint;

// TCanvasStates = (csHandleValid, csFontValid, csPenValid, csBrushValid);
type TCanvasStates int32

const (
	CsHandleValid = iota + 0
	CsFontValid
	CsPenValid
	CsBrushValid
)

// TCanvasState = set of TCanvasStates;
type TCanvasState uint32

// TCanvasOrientation = (coLeftToRight, coRightToLeft);
type TCanvasOrientation int32

const (
	CoLeftToRight = iota + 0
	CoRightToLeft
)

// // Note: tfComposited only supported by ThemeServices.DrawText
// TTextFormats = (tfBottom, tfCalcRect, tfCenter, tfEditControl, tfEndEllipsis,
//   tfPathEllipsis, tfExpandTabs, tfExternalLeading, tfLeft, tfModifyString,
//   tfNoClip, tfNoPrefix, tfRight, tfRtlReading, tfSingleLine, tfTop,
//   tfVerticalCenter, tfWordBreak, tfHidePrefix, tfNoFullWidthCharBreak,
//   tfPrefixOnly, tfTabStop, tfWordEllipsis, tfComposited);
type TTextFormats int32

const (
	TfBottom = iota + 0
	TfCalcRect
	TfCenter
	TfEditControl
	TfEndEllipsis
	TfPathEllipsis
	TfExpandTabs
	TfExternalLeading
	TfLeft
	TfModifyString
	TfNoClip
	TfNoPrefix
	TfRight
	TfRtlReading
	TfSingleLine
	TfTop
	TfVerticalCenter
	TfWordBreak
	TfHidePrefix
	TfNoFullWidthCharBreak
	TfPrefixOnly
	TfTabStop
	TfWordEllipsis
	TfComposited
)

// TTextFormat = set of TTextFormats;
type TTextFormat uint32

// TStyleElements = set of (seFont, seClient, seBorder);
type TStyleElements uint32

const (
	SeFont = iota + 0
	SeClient
	SeBorder
)

// TBevelCut = (bvNone, bvLowered, bvRaised, bvSpace);
type TBevelCut int32

const (
	BvNone = iota + 0
	BvLowered
	BvRaised
	BvSpace
)

// TBevelEdge = (beLeft, beTop, beRight, beBottom);
type TBevelEdge int32

const (
	BeLeft = iota + 0
	BeTop
	BeRight
	BeBottom
)

// TBevelEdges = set of TBevelEdge;
type TBevelEdges uint32

// TBevelKind = (bkNone, bkTile, bkSoft, bkFlat);
type TBevelKind int32

const (
	BkNone = iota + 0
	BkTile
	BkSoft
	BkFlat
)

// TTickMark = (tmBottomRight, tmTopLeft, tmBoth);
type TTickMark int32

const (
	TmBottomRight = iota + 0
	TmTopLeft
	TmBoth
)

// TTickStyle = (tsNone, tsAuto, tsManual);
type TTickStyle int32

const (
	TsNone = iota + 0
	TsAuto
	TsManual
)

// TPositionToolTip = (ptNone, ptTop, ptLeft, ptBottom, ptRight);
type TPositionToolTip int32

const (
	PtNone = iota + 0
	PtTop
	PtLeft
	PtBottom
	PtRight
)

// TDateTimeKind = (dtkDate, dtkTime);
type TDateTimeKind int32

const (
	DtkDate = iota + 0
	DtkTime
)

// TDTDateMode = (dmComboBox, dmUpDown);
type TDTDateMode int32

const (
	DmComboBox = iota + 0
	DmUpDown
)

// TDTDateFormat = (dfShort, dfLong);
type TDTDateFormat int32

const (
	DfShort = iota + 0
	DfLong
)

// TDTCalAlignment = (dtaLeft, dtaRight);
type TDTCalAlignment int32

const (
	DtaLeft = iota + 0
	DtaRight
)

// { Calendar common control support }
// TCalDayOfWeek = (dowMonday, dowTuesday, dowWednesday, dowThursday,
//   dowFriday, dowSaturday, dowSunday, dowLocaleDefault);
type TCalDayOfWeek int32

const (
	dowMonday = iota + 0
	dowTuesday
	dowWednesday
	dowThursday
	dowFriday
	dowSaturday
	dowSunday
	dowLocaleDefault
)

// TSearchType = (stWholeWord, stMatchCase);
type TSearchType int32

const (
	StWholeWord = iota + 0
	StMatchCase
)

// TSearchTypes = set of TSearchType;

type TSearchTypes uint32

// TNumberingStyle = (nsNone, nsBullet);
type TNumberingStyle int32

const (
	NsNone = iota + 0
	NsBullte
)

// TAttributeType = (atSelected, atDefaultText);
type TAttributeType int32

const (
	AtSelected = iota + 0
	AtDefaultText
)

// TConsistentAttribute = (caBold, caColor, caFace, caItalic,
//   caSize, caStrikeOut, caUnderline, caProtected);
type TConsistentAttribute int32

const (
	CaBold = iota + 0
	CaColor
	CaFace
	CaItalic
	CaSize
	CaStrikeOut
	CaUnderline
	CaProtected
)

// TConsistentAttributes = set of TConsistentAttribute;
type TConsistentAttributes uint32

// TIconArrangement = (iaTop, iaLeft);
type TIconArrangement int32

const (
	IaTop = iota + 0
	IaLeft
)

// THeaderStyle = (hsGradient, hsImage, hsThemed);
type THeaderStyle int32

const (
	HsGradient = iota + 0
	HsImage
	HsThemed
)

// TImageAlignment = (iaLeft, iaRight, iaTop, iaBottom, iaCenter);
type TImageAlignment int32

const (
	IiaLeft = iota + 0 // IaTop有冲突，所以增加一个i
	IiaRight
	IiaTop
	IiaBottom
	IiaCenter
)

//  TAnchorKind = (akLeft, akTop, akRight, akBottom);
type TAnchorKind int32

const (
	AkLeft = iota + 0
	AkTop
	AkRight
	AkBottom
)

//  TAnchors = set of TAnchorKind;
type TAnchors uint32

// TOwnerDrawState = set of (odSelected, odGrayed, odDisabled, odChecked,
//    odFocused, odDefault, odHotLight, odInactive, odNoAccel, odNoFocusRect,
//    odReserved1, odReserved2, odComboBoxEdit);
type TOwnerDrawState uint32

const (
	OdSelected = iota + 0
	OdGrayed
	OdDisabled
	OdChecked
	OdFocused
	OdDefault
	OdHotLight
	OdInactive
	OdNoAccel
	OdNoFocusRect
	OdReserved1
	OdReserved2
	OdComboBoxEdit
)

//   TBitBtnKind = (bkCustom, bkOK, bkCancel, bkHelp, bkYes, bkNo, bkClose,
// bkAbort, bkRetry, bkIgnore, bkAll);
type TBitBtnKind int32

const (
	BkCustom = iota + 0
	BkOK
	BkCancel
	BkHelp
	BkYes
	BkNo
	BkClose
	BkAbort
	BkRetry
	BkIgnore
	BkAll
)
