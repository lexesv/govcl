@echo off
echo Ñ¹Ëõwin32
upx.exe ".\win32\libvcl.dll" -9
upx.exe ".\win32\libvcl.dll" -t

rem echo Ñ¹Ëõwin64
rem upx.exe ".\win64\libvclx64.dll" -9
rem upx.exe ".\win32\libvcl.dll" -t

pause