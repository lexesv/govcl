# govcl

#### 目录
* [项目介绍](#项目介绍)
* [重要说明](#重要说明)
* [WIKI](#WIKI)
* [使用方法](#使用方法)
* [图标及manifest文件集成](#图标及manifest文件集成)
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

#### 重要说明
**所有的代码只会存储在OSC的[码云](https://gitee.com/ying32/govcl)中，原因在于go包路径的问题。**  
**至于github上会建一个同名的项目[govcl](https://github.com/ying32/govcl)，但不会提交任何代码**  
**xui包目前还未完成，但不影响正常使用**  


#### WIKI 

后面视情况慢慢补充相关VCL组件的知识，详情可以[WIKI](https://gitee.com/ying32/govcl/wikis/Home)  

#### 使用方法
> go get gitee.com/ying32/govcl  

```golang
package main

import (
   "gitee.com/ying32/govcl/vcl"
)

var (
   mainForm *vcl.TForm
)

func main() {
    vcl.Application.Initialize()
    mainForm = vcl.Application.CreateForm()
    mainForm.SetCaption("Hello")
    mainForm.EnabledMaximize(false)
    mainForm.ScreenCenter()
    vcl.Application.Run()
}

```  

复制"bin\win32\libvcl.dll"或者"bin\win64\libvclx64.dll"到当前exe目录或系统环境路径下  

#### 图标及manifest文件集成

需要使用rsrc工具生成syso文件    
> go get github.com/akavel/rsrc 

```bat
rsrc -ico="your.ico" -o="your.manifest"
```

然后将生成的文件复制到你的工程目录下  

> manifest文件内容  

```xml  

<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0" xmlns:asmv3="urn:schemas-microsoft-com:asm.v3">
	<assemblyIdentity version="1.0.0.0" processorArchitecture="*" name="SomeFunkyNameHere" type="win32"/>
	<dependency>
		<dependentAssembly>
			<assemblyIdentity type="win32" name="Microsoft.Windows.Common-Controls" version="6.0.0.0" processorArchitecture="*" publicKeyToken="6595b64144ccf1df" language="*"/>
		</dependentAssembly>
	</dependency>
	<asmv3:application>
		<asmv3:windowsSettings xmlns="http://schemas.microsoft.com/SMI/2005/WindowsSettings">
			<dpiAware>true</dpiAware>
		</asmv3:windowsSettings>
	</asmv3:application>
</assembly>

```  

#### 项目中的包说明

* vcl  
  包含Delphi标准组件中的大部分    
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

[现支持组件和非组件类列表](https://gitee.com/ying32/govcl/wikis/%E6%94%AF%E6%8C%81%E7%9A%84%E7%BB%84%E4%BB%B6%E5%88%97%E8%A1%A8)  

#### 截图

![截图1](https://gitee.com/ying32/govcl/raw/master/Screenshot/1.png)   
![截图2](https://gitee.com/ying32/govcl/raw/master/Screenshot/2.png)      
![绘图](https://gitee.com/ying32/govcl/raw/master/Screenshot/draw.png)  
![ListView](https://gitee.com/ying32/govcl/raw/master/Screenshot/listview.png)  
![RichEdit](https://gitee.com/ying32/govcl/raw/master/Screenshot/richedit.png)  
![标准控件](https://gitee.com/ying32/govcl/raw/master/Screenshot/std.png)  
![样式](https://gitee.com/ying32/govcl/raw/master/Screenshot/style.png)  
#### 备注
**文件名后面带有def的为手动编写**   

#### 作者信息
by: ying32