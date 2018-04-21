// govcl wke测试头文件，by: ying32
// https://gitee.com/ying32/ying32

package main

import (
	"syscall"

	"fmt"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/api"
	"gitee.com/ying32/govcl/vcl/types"
)

type wkeWindowType int32

const (
	WKE_WINDOW_TYPE_POPUP wkeWindowType = iota + 0
	WKE_WINDOW_TYPE_TRANSPARENT
	WKE_WINDOW_TYPE_CONTROL
)

type wkeNavigationType int32

const (
	WKE_NAVIGATION_TYPE_LINKCLICK wkeNavigationType = iota + 0
	WKE_NAVIGATION_TYPE_FORMSUBMITTE
	WKE_NAVIGATION_TYPE_BACKFORWARD
	WKE_NAVIGATION_TYPE_RELOAD
	WKE_NAVIGATION_TYPE_FORMRESUBMITT
	WKE_NAVIGATION_TYPE_OTHER
)

type wkeLoadingResult int32

const (
	WKE_LOADING_SUCCEEDED wkeLoadingResult = iota + 0
	WKE_LOADING_FAILED
	WKE_LOADING_CANCELED
)

type wkeDocumentReadyInfo struct {
	url              uintptr
	frameJSState     uintptr
	mainFrameJSState uintptr
}

/*
  PwkeNewViewInfo = ^wkeNewViewInfo;
  wkeNewViewInfo = record
    navigationType: wkeNavigationType;
    url: wkeString;
    target: wkeString;

    x: Integer;
    y: Integer;
    width: Integer;
    height: Integer;
    menuBarVisible: Boolean;
    statusBarVisible: bool ;
    toolBarVisible: Boolean;
    locationBarVisible: Boolean;
    scrollbarsVisible: Boolean;
    resizable: Boolean;
    fullscreen: Boolean;
  end;
  TwkeNewViewInfo = wkeNewViewInfo;
*/

var (
	wkedll               = syscall.NewLazyDLL("wke.dll")
	_wkeInitialize       = wkedll.NewProc("wkeInitialize")
	_wkeFinalize         = wkedll.NewProc("wkeFinalize")
	_wkeCreateWebWindow  = wkedll.NewProc("wkeCreateWebWindow")
	_wkeShowWindow       = wkedll.NewProc("wkeShowWindow")
	_wkeDestroyWebWindow = wkedll.NewProc("wkeDestroyWebWindow")
	_wkeMoveWindow       = wkedll.NewProc("wkeMoveWindow")
	_wkeLoadW            = wkedll.NewProc("wkeLoadW")
	_wkeRepaintAllNeeded = wkedll.NewProc("wkeRepaintAllNeeded")
	_wkeReload           = wkedll.NewProc("wkeReload")
	_wkeCanGoBack        = wkedll.NewProc("wkeCanGoBack")
	_wkeGoBack           = wkedll.NewProc("wkeGoBack")
	_wkeCanGoForward     = wkedll.NewProc("wkeCanGoForward")
	_wkeGoForward        = wkedll.NewProc("wkeGoForward")
	_wkeOnTitleChanged   = wkedll.NewProc("wkeOnTitleChanged")
	_wkeOnURLChanged     = wkedll.NewProc("wkeOnURLChanged")
	_wkeGetStringW       = wkedll.NewProc("wkeGetStringW")

	// 事件设置
	//procedure wkeOnPaintUpdated; external wkedll name 'wkeOnPaintUpdated';
	//procedure wkeOnAlertBox; external wkedll name 'wkeOnAlertBox';
	//procedure wkeOnConfirmBox; external wkedll name 'wkeOnConfirmBox';
	//procedure wkeOnPromptBox; external wkedll name 'wkeOnPromptBox';
	//procedure wkeOnConsoleMessage; external wkedll name 'wkeOnConsoleMessage';
	_wkeOnNavigation    = wkedll.NewProc("wkeOnNavigation")
	_wkeOnCreateView    = wkedll.NewProc("wkeOnCreateView")
	_wkeOnDocumentReady = wkedll.NewProc("wkeOnDocumentReady")
	_wkeOnLoadingFinish = wkedll.NewProc("wkeOnLoadingFinish")

	// callback
	_wkeTitleChangedCallback  = syscall.NewCallbackCDecl(wkeTitleChangedCallback)
	_wkeURLChangedCallback    = syscall.NewCallbackCDecl(wkeURLChangedCallback)
	_wkeNavigationCallback    = syscall.NewCallbackCDecl(wkeNavigationCallback)
	_wkeLoadingFinishCallback = syscall.NewCallbackCDecl(wkeLoadingFinishCallback)
	_wkeDocumentReadyCallback = syscall.NewCallbackCDecl(wkeDocumentReadyCallback)
	_wkeCreateViewCallback    = syscall.NewCallbackCDecl(wkeCreateViewCallback)
)

func wkeTitleChangedCallback(webView, param, title uintptr) uintptr {
	fmt.Println("title changed.")
	vcl.FormFromInst(param).SetCaption(wkeGetStringW(title) + " - Wke测试")
	return 1
}

func wkeURLChangedCallback(webView, param, url uintptr) uintptr {
	fmt.Println("url changed.")
	vcl.EditFromInst(param).SetText(wkeGetStringW(url))
	return 1
}

func wkeNavigationCallback(webView, param uintptr, navigationType wkeNavigationType, url uintptr) uintptr {
	fmt.Println("nav:", navigationType, ", url:", wkeGetStringW(url))
	switch navigationType {
	case WKE_NAVIGATION_TYPE_LINKCLICK, WKE_NAVIGATION_TYPE_BACKFORWARD:
		vcl.StatusBarFromInst(param).SetSimpleText("正在跳转：" + wkeGetStringW(url))

	case WKE_NAVIGATION_TYPE_FORMSUBMITTE:
		vcl.StatusBarFromInst(param).SetSimpleText("提交表单中：" + wkeGetStringW(url))

	case WKE_NAVIGATION_TYPE_RELOAD:
		vcl.StatusBarFromInst(param).SetSimpleText("正在重新载入：" + wkeGetStringW(url))

	case WKE_NAVIGATION_TYPE_FORMRESUBMITT:
		vcl.StatusBarFromInst(param).SetSimpleText("重新提交表单：" + wkeGetStringW(url))

	case WKE_NAVIGATION_TYPE_OTHER:
		// 不知道啥类型了
		vcl.StatusBarFromInst(param).SetSimpleText("正在载入：" + wkeGetStringW(url))
	}

	return 1
}

func wkeLoadingFinishCallback(webView, param, url uintptr, result wkeLoadingResult, failedReason uintptr) uintptr {
	fmt.Println("LoadingFinish:", result)
	switch result {
	case WKE_LOADING_SUCCEEDED:
		vcl.StatusBarFromInst(param).SetSimpleText("加载完成.")
	case WKE_LOADING_FAILED:
		vcl.StatusBarFromInst(param).SetSimpleText("加载失败.")
	case WKE_LOADING_CANCELED:
		vcl.StatusBarFromInst(param).SetSimpleText("加载缓存.")
	}
	return 1
}

func wkeDocumentReadyCallback(webView, param, info uintptr) uintptr {
	vcl.StatusBarFromInst(param).SetSimpleText("文档已经准备")
	return 1
}

func wkeCreateViewCallback(webView, param, info uintptr) uintptr {
	// PwkeNewViewInfo
	// 这里直接返回最初创建的webview,param到时候要转过来吧，不让弹出新窗口
	return param
}

//wkeAlertBoxCallback = procedure(webView: wkeWebView; param: Pointer; msg: wkeString); cdecl;
//wkeConfirmBoxCallback = function(webView: wkeWebView; param: Pointer; msg: wkeString): Boolean; cdecl;
//wkePromptBoxCallback = function(webView: wkeWebView; param: Pointer; msg: wkeString; defaultResult: wkeString; result: wkeString): Boolean; cdecl;
//wkeDocumentReadyCallback = procedure(webView: wkeWebView; param: Pointer; info: PwkeDocumentReadyInfo); cdecl;
//wkeWindowClosingCallback = function(webWindow: wkeWebView; param: Pointer): Boolean; cdecl;
//wkeWindowDestroyCallback = procedure(webWindow: wkeWebView; param: Pointer); cdecl;

//--------------------------------------------------------

func wkeInitialize() {
	_wkeInitialize.Call()
}

func wkeFinalize() {
	_wkeFinalize.Call()
}

func wkeCreateWebWindow(aType wkeWindowType, hWd types.HWND, x, y, width, height int32) uintptr {
	r, _, _ := _wkeCreateWebWindow.Call(uintptr(aType), uintptr(hWd), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	return r
}

func wkeShowWindow(webWindow uintptr, show bool) {
	_wkeShowWindow.Call(webWindow, api.GoBoolToDBool(show))
}

func wkeDestroyWebWindow(webWindow uintptr) {
	_wkeDestroyWebWindow.Call(webWindow)
}

func wkeMoveWindow(webWindow uintptr, x, y, width, height int32) {
	_wkeMoveWindow.Call(webWindow, uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}

func wkeLoadW(webWindow uintptr, url string) {
	_wkeLoadW.Call(webWindow, api.GoStrToDStr(url))
}

func wkeRepaintAllNeeded() bool {
	r, _, _ := _wkeRepaintAllNeeded.Call()
	return api.DBoolToGoBool(r)

}

func wkeReload(webWindow uintptr) {
	_wkeReload.Call(webWindow)
}

func wkeCanGoBack(webWindow uintptr) bool {
	r, _, _ := _wkeCanGoBack.Call(webWindow)
	return api.DBoolToGoBool(r)
}

func wkeGoBack(webWindow uintptr) {
	_wkeGoBack.Call(webWindow)
}

func wkeCanGoForward(webWindow uintptr) bool {
	r, _, _ := _wkeCanGoForward.Call(webWindow)
	return api.DBoolToGoBool(r)
}

func wkeGoForward(webWindow uintptr) {
	_wkeGoForward.Call(webWindow)
}

func wkeGetStringW(wkeStr uintptr) string {
	r, _, _ := _wkeGetStringW.Call(wkeStr)
	return api.DStrToGoStr(r)
}

func wkeOnTitleChanged(webView, callback, callbackParam uintptr) {
	_wkeOnTitleChanged.Call(webView, callback, callbackParam)
}

func wkeOnURLChanged(webView, callback, callbackParam uintptr) {
	_wkeOnURLChanged.Call(webView, callback, callbackParam)
}

func wkeOnNavigation(webView, callback, callbackParam uintptr) {
	_wkeOnNavigation.Call(webView, callback, callbackParam)
}

func wkeOnLoadingFinish(webView, callback, callbackParam uintptr) {
	_wkeOnLoadingFinish.Call(webView, callback, callbackParam)
}

func wkeOnDocumentReady(webView, callback, callbackParam uintptr) {
	_wkeOnDocumentReady.Call(webView, callback, callbackParam)
}

func wkeOnCreateView(webView, callback, callbackParam uintptr) {
	_wkeOnCreateView.Call(webView, callback, callbackParam)
}

//procedure wkeOnPaintUpdated(webView: wkeWebView; callback: wkePaintUpdatedCallback; callbackParam: Pointer); cdecl;
//procedure wkeOnAlertBox(webView: wkeWebView; callback: wkeAlertBoxCallback; callbackParam: Pointer); cdecl;
//procedure wkeOnConfirmBox(webView: wkeWebView; callback: wkeConfirmBoxCallback; callbackParam: Pointer); cdecl;
//procedure wkeOnPromptBox(webView: wkeWebView; callback: wkePromptBoxCallback; callbackParam: Pointer); cdecl;
//procedure wkeOnCreateView(webView: wkeWebView; callback: wkeCreateViewCallback; param: Pointer); cdecl;
//procedure wkeOnConsoleMessage(webView: wkeWebView; callback: wkeConsoleMessageCallback; callbackParam: Pointer); cdecl;
