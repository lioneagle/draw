@echo off
setlocal
if exist install.bat goto ok
echo install.bat must be run from its folder
goto end
: ok
call env.bat
gofmt -w src
go install %1
if %GOARCH% == amd64 goto build_32
goto end
:build_32
set GOARCH=386
go build -o=%132.exe %1
copy %132.exe  .\bin\%132.exe
del %132.exe
set GOARCH=amd64
:end
echo finished