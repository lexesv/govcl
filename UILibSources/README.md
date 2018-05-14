### libvcl

libvcl 基于Delphi 10.2.1版本


编译步骤：  


* 1、安装好Delphi 10.2.1。  
* 2、双击vcl.dproj或者vcl.dpr在IDE中打开工程。  
* 3、在右边Project Manager的File列表中选择Build Configurations为Release。
* 4、在右边Project Manager的File列表中选择Target Platforms (32-bit Windows或者64-bit Windows)。  
* 5、按下Ctrl+F9（或直接在Release配置右键，选择Make或者Build）。  

因为配置中设置了相关编译后的操作，编译完后的二进制可以在以下目录查看  

> %GOPATH%\bin\libvcl.dll  
> %GOPATH%\bin\libvclx64.dll    
> %GOPATH%\src\gitee.com\ying32\bin\win32\libvcl.dll  
> %GOPATH%\src\gitee.com\ying32\bin\win64\libvclx64.dll    


**注：源码无第三方依赖库，安装好相关版本Delphi后直接编译即可。**


----

### liblcl 

liblcl 基于Lazarus 1.8.0版本 FPC 3.0.4


编译步骤：  

* 1、安装好Lazarus 1.8.0 64bit版本及i386扩展包   
* 2、双击lcl.lpi  
* 3、菜单->Project->Project Options -> Compiler Options -> Build modes 切换相关编环模式，当前有效模式为以下4种：   
  * ReleaseWindows64  
  * ReleaseWindows32  
  * ReleaseLinux64  
  * ReleaseMacOS32  
* 4、菜单->Run->Compile(或者Build)

生成的文件位于：  

> Windows: `"..\..\..\..\..\bin\liblcl"`     
> Linux: `"../bin/liblcl"`  
> MacOS: `"../../../../../bin/liblcl"`

**注：源码无第三方依赖库，安装好相关版本Lazarus后直接编译即可。**