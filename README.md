# govcl

#### 目录
* [项目介绍](#项目介绍)
* [重要说明](#重要说明)
* [UI设计器](https://gitee.com/ying32/govcl/wikis/UI%E8%AE%BE%E8%AE%A1%E5%99%A8)
* [WIKI](https://gitee.com/ying32/govcl/wikis/Home)
* [使用方法](#使用方法)
* [icon及manifest文件集成](https://gitee.com/ying32/govcl/wikis/Icon%E5%8F%8Amanifest%E6%96%87%E4%BB%B6%E9%9B%86%E6%88%90)
* [关于跨平台问题](https://gitee.com/ying32/govcl/wikis/%E5%85%B3%E4%BA%8EUI%E5%BA%93%E8%B7%A8%E5%B9%B3%E5%8F%B0%E9%97%AE%E9%A2%98?parent=FQA)  
* [项目中的包说明](#项目中的包说明)
* [实例类说明](#实例类说明)
* [支持的组件列表](https://gitee.com/ying32/govcl/wikis/%E6%94%AF%E6%8C%81%E7%9A%84%E7%BB%84%E4%BB%B6%E5%88%97%E8%A1%A8)
* [截图](#截图)
* [备注](#备注)
* [点击链接加入群【govcl交流】](https://jq.qq.com/?_wv=1027&k=5Sv7Qiq) 群号：263106281
* [作者信息](#作者信息)

#### 项目介绍

> 1、由于现有第三方的Go UI库不是太宠大就是用的不习惯，或者组件太少。就萌生了自己写一个UI库的相法  
> Delphi有些许多优秀的VCL组件，不拿来使用太可惜了。所以就索性做了一套。目前支持Win32跟Win64，  
> 只需要带上一个libvcl.dll即可。  

> 2、项目现在支持VCL标准控件中的大部分，足以满足日常操作了，具体见[支持的组件列表](https://gitee.com/ying32/govcl/wikis/%E6%94%AF%E6%8C%81%E7%9A%84%E7%BB%84%E4%BB%B6%E5%88%97%E8%A1%A8)。  
> 事件方面也支持部分，参见：[支持的事件](https://gitee.com/ying32/govcl/wikis/%E6%94%AF%E6%8C%81%E7%9A%84%E4%BA%8B%E4%BB%B6)：  
 

#### 重要说明
**所有的代码只会存储在OSC的[码云](https://gitee.com/ying32/govcl)中，原因在于go包路径的问题。**  
**至于github上会建一个同名的项目[govcl](https://github.com/ying32/govcl)，但不会提交任何代码**  
**xui包目前还未完成，但不影响正常使用**   

> **govcl是一个不需要经常更新的项目，因为原本vcl就是个成熟的产品，这里也只是做了相关的桥接，相对是稳定的，但也不排除某些环境下的问题。**  

**govcl是以go 1.9作为基础版本开发，目前已知的是在go 1.7也可以编译，但作为以后的考虑会使用1.9中的特性，所以大家在使用的时候尽量使用go 1.9+版本**  

**希望大家有问题的话通过Issues来进行反馈，反馈错误的话最好能带有相关错误的截图之类的， 而不是通过评论来提问。wiki也可关注下，有些问题在会里面作解答。**  


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
    vcl.Application.SetIconResId(3)
    vcl.Application.Initialize()
    mainForm = vcl.Application.CreateForm()
    mainForm.SetCaption("Hello")
    mainForm.EnabledMaximize(false)
    mainForm.ScreenCenter()
    vcl.Application.Run()
}

```  

复制"bin\win32\libvcl.dll"或者"bin\win64\libvclx64.dll"到当前exe目录或系统环境路径下  

#### 项目中的包说明

* vcl  
  包含Delphi标准组件中的大部分    
* api  
  DLL函数申明与重新包装  
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