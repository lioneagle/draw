@echo off
setlocal
if exist install.bat goto ok
echo install.bat must be run from its folder
goto end
: ok
call env.bat
gofmt -w src

if %GOARCH% == amd64 (
go install %1
copy .\bin\%1.exe  .\bin\%164.exe
del .\bin\%1.exe

set GOARCH=386
go build -o=.\bin\%132.exe %1
set GOARCH=amd64

) else (

go install %1
copy .\bin\%1.exe  .\bin\%132.exe
del .\bin\%1.exe

)

:end
echo finished