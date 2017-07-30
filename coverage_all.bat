@echo off
setlocal
if exist coverage.bat goto ok
echo coverage.bat must be run from its folder
goto end
: ok

call env.bat

if not exist test_temp mkdir test_temp

if exist .\test_temp\coverage.out  del .\test_temp\coverage.out
if exist .\test_temp\coverage.html del .\test_temp\coverage.html

echo mode: set >> .\test_temp\coverage.out
for /f %%d in ('go list ./...^|findstr -v main^|findstr -v github') do (
    go test -coverprofile=./test_temp/coverage1.out %%d
    findstr -v mode: .\test_temp\coverage1.out >> .\test_temp\coverage2.out
    cat .\test_temp\coverage2.out >> .\test_temp\coverage.out
    del .\test_temp\coverage1.out
    del .\test_temp\coverage2.out
)

go tool cover -html=./test_temp/coverage.out -o ./test_temp/coverage.html
.\test_temp\coverage.html

:end
echo finished