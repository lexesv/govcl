package main

import (
	"fmt"
	"time"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/types"
)

var (
	frmLogin       *vcl.TForm
	imgQRCode      *vcl.TImage
	lblMessage     *vcl.TLabel
	exitCheckLogin bool
	isLogin        bool
)

const (
	TB_QRLOGIN   = "https://qrlogin.taobao.com/qrcodelogin"
	TB_GENQRCODE = TB_QRLOGIN + "/generateQRCode4Login.do?from=tb&_ksTS=%d"
	TB_CHECK     = TB_QRLOGIN + "/qrcodeLoginCheck.do?lgToken=%s&defaulturl=http%%3A%%2F%%2Fcainiaoyizhan.com%%2Findex.htm&_ksTS=%d"
)

// TTBGetQRCode 淘宝二维码解析
type TTBGetQRCode struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Url     string `json:"url"`
	LgToken string `json:"lgToken"`
	AdToken string `json:"adToken"`
}

type TTBCheck struct {
	Code    int    `json:"code,string"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Url     string `json:"url"`
}

func initfrmLogin() {
	frmLogin = vcl.Application.CreateForm()
	frmLogin.SetCaption("登录")
	frmLogin.ScreenCenter()
	frmLogin.SetBorderStyle(types.BsSingle)
	frmLogin.EnabledMaximize(false)
	frmLogin.EnabledMinimize(false)
	// SetClientWidth ??? bug，汗竟然出现显示任务栏的bug了
	frmLogin.SetWidth(210)
	frmLogin.SetHeight(250)
	frmLogin.SetDoubleBuffered(true)
	frmLogin.SetOnShow(func(vcl.IObject) {
		frmLogin.SetCaption("登录")
		lblMessage.SetCaption("正在获取二码")
		imgQRCode.Picture().Assign(nil)
		go downloadQRCode()
	})
	frmLogin.SetOnClose(func(vcl.IObject, *types.TCloseAction) {
		exitCheckLogin = true
	})
	frmLogin.SetOnHide(func(vcl.IObject) {
		exitCheckLogin = true
	})

	imgQRCode = vcl.NewImage(frmLogin)
	imgQRCode.SetParent(frmLogin)
	imgQRCode.SetProportional(true)
	imgQRCode.SetAlign(types.AlClient)
	imgQRCode.SetCenter(true)

	lblMessage = vcl.NewLabel(frmLogin)
	lblMessage.SetParent(frmLogin)
	lblMessage.SetAlign(types.AlBottom)
	lblMessage.SetAlignment(types.TaCenter)
	lblMessage.SetLayout(types.TlCenter)
	lblMessage.SetAutoSize(false)
	lblMessage.SetHeight(50)
}

func downloadQRCode() {
	if imgQRCode != nil && imgQRCode.IsValid() {
		var res TTBGetQRCode
		err := httpGetJSON(fmt.Sprintf(TB_GENQRCODE, time.Now().Unix()), &res)
		if err != nil {
			fmt.Println("解析错误：", err)
			return
		}
		if res.Success {
			bs, err := httpGet("https:" + res.Url)
			if err != nil {
				fmt.Println("读取二维码数据错误：", err)
				return
			}
			mem := vcl.NewMemoryStreamFromBytes(bs)
			defer mem.Free()
			mem.SaveToFile(".\\qr.png")
			mem.SetPosition(0)
			imgQRCode.Picture().LoadFromStream(mem)
			lblMessage.SetCaption("请扫描二维码登录")
			exitCheckLogin = false
			defer checkQRCodeLogin(res.LgToken)
		}
	}
}

func checkQRCodeLogin(token string) {
	urlStr := fmt.Sprintf(TB_CHECK, token, time.Now().Unix())
	fmt.Println("checkURL:", urlStr)
	for !exitCheckLogin {
		var res TTBCheck
		err := httpGetJSON(urlStr, &res)
		if err != nil {
			fmt.Println("JSON解析错误:", err)
			break
		}
		switch res.Code {
		case 10000:
			fmt.Println("请扫描二维码登录")
			lblMessage.SetCaption("请扫描二维码登录")
		case 10001:
			fmt.Println("已扫描二维码，请在确认登录")
			lblMessage.SetCaption("已扫描二维码，请在确认登录")
			break
		case 10004:
			fmt.Println("已过期请重新扫描")
			lblMessage.SetCaption("已过期请重新扫描")
			break
		case 10006:
			isLogin = true
			exitCheckLogin = true
			fmt.Println("登录成功")
			lblMessage.SetCaption("登录成功")
			session.Get(res.Url)
			break
		default:
			lblMessage.SetCaption("未知状态")
			break
		}
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Println("已退出检测。")
}
