# govcl

#### 目录
* [项目介绍](#项目介绍)
* [重要说明](#重要说明)
* [使用方法](#使用方法)
* [项目中的包说明](#项目中的包说明)
* [实例类说明](#实例类说明)
* [支持的组件列表](#支持的组件列表)
* [截图](#截图)
* [备注](#备注)
* [作者信息](#作者信息)

#### 项目介绍

> 1、由于现有第三方的Go UI库不是太宠大就是用的不习惯，或者组件太少。就萌生了自己写一个UI库的相法  
> Delphi有些许多优秀的VCL组件，不拿来使用太可惜了。所以就索性做了一套。目前支持Win32跟Win64，  
> 只需要带上一个libvcl.dll即可。  

> 2、项目现在支持VCL标准控件中的大部分，足以满足日常操作了，具体见[支持的组件列表](#支持的组件列表)。  
> 事件方面也支持部分，如下：  
```delphi
 TGoEvent = (geClick, geClose, geFormClose, geFormCloseQuery, geChange,
              geUpDownClick, geTreeViewChange, geListViewChange, geDblClick, gePaint,
              geResize, geShow, geMenuChange, geEnter, geExit, gePopup, geBalloonClick,
              geLinkClick, geExecute, geUpdate, geException, geTimer, geMinimize,
              geRestore, geHide, geKeyDown, geKeyPress, geKeyUp, geMouseDown,
              geMouseEnter, geMouseLeave, geMouseMove, geMouseUp, geMouseWheel);
```

> 3、关于不能跨平台问题，原本是有打算借助CrossVcl这个库做到跨三个平台的，但在尝试中发现在共享  
> 库中创建会报错，调试了下，没有找到原因，所以暂时就放弃这块了，后期视情况看能不能解决这个问题。  

#### 重要说明
**所有的代码只会存储在OSC的[码云](https://gitee.com/ying32/govcl)中，原因在于go包路径的问题。**  
**至于github上会建一个同名的项目[govcl](https://github.com/ying32/govcl)，但不会提交任何代码**  


#### 使用方法
> go get gitee.com/ying32/govcl  

```golang
package main

import (
   "gitee.com/ying32/govcl/vcl"
   "gitee.com/ying32/govcl/api" // 这个视情况导入
)

var (
   mainForm *vcl.TForm
)

func main() {
    vcl.Application.Initialize()
    mainForm = vcl.Application.CreateForm()
    mainForm.SetCaption("Hello")
    mainForm.EnabledMaximize(false)
    mainForm.SetDoubleBuffered(true)
    mainForm.ScreenCenter()
    vcl.Application.Run()
}

```  

#### 项目中的包说明

* api  
  DLL函数申明与重新包装  
* dylib  
  仅针对Linux及MacOS，模拟windows下动态调用，需要用到cgo  
* rtl  
  包含Delphi中Set类型操作、内存操作等其它函数  
* win  
  包含windows下的常量、函数、类型定义  
* xui  
  包含一个使用xml创建UI的类  
* types  
  包含 类型定义、枚举定义、常量


#### 实例类说明

> 按照Delphi中的Application、 Screen、 Mouse、Clipboard四个类实例是可以直接访问的，不需要释放  
其实组件带有Owner参数的一般指定当前组件对应的TForm就好了，这样就不需要手动释放，反之Owner填   
写nil则需要手动调用Free，就像其它非组件类的。  

#### 支持的组件列表
> 现支持组件和非组件类列表：  
>
TApplication    
TForm    
TButton    
TEdit    
TMainMenu    
TPopupMenu    
TMemo    
TCheckBox    
TRadioButton    
TGroupBox    
TLabel    
TListBox    
TComboBox    
TPanel    
TImage    
TLinkLabel    
TSpeedButton   
TSplitter    
TRadioGroup    
TStaticText    
TColorBox    
TColorListBox    
TTrayIcon    
TBalloonHint    
TCategoryPanelGroup    
TOpenDialog    
TSaveDialog    
TColorDialog    
TFontDialog    
TPrintDialog    
TOpenPictureDialog    
TSavePictureDialog    
TSaveTextFileDialog    
TOpenTextFileDialog    
TRichEdit    
TTrackBar    
TImageList    
TUpDown    
TProgressBar    
THotKey    
TDateTimePicker    
TMonthCalendar    
TListView    
TTreeView    
TStatusBar    
TToolBar    
TPageControl  
TTabSheet    
TControl 
TActionList  
TToolButton     
TPaintBox  
TTimer  
> 
TIcon    
TBitmap    
TMemoryStream    
TFont    
TStrings    
TStringList    
TBrush    
TPen    
TMenuItem    
TListGroups    
TPicture    
TListColumns    
TListItems    
TTreeNodes    
TListItem    
TTreeNode      
TScreen    
TMouse    
TListGroup    
TListColumn    
TCollectionItem    
TStatusPanels    
TStatusPanel    
TCanvas    
TObject    
TPngImage    
TJPEGImage    
TGIFImage    
TGIFFrame    
TIniFile  
TRegistry  
TClipboard  
TMonitor  
TMargins  
TList  
TGraphic  
TComponent  

#### 截图

![截图1](https://gitee.com/ying32/govcl/raw/master/Screenshot/1.png)   
![截图2](https://gitee.com/ying32/govcl/raw/master/Screenshot/2.png)    

#### 备注
**文件名后面带有def的为手动编写**   

#### 作者信息
by: ying32