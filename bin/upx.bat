@echo off
echo ѹ��win32
upx.exe ".\win32\libvcl.dll" -9
upx.exe ".\win32\libvcl.dll" -t

echo ѹ��win64
upx.exe ".\win64\libvclx64.dll" -9
upx.exe ".\win32\libvcl.dll" -t

pause