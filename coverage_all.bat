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

	
for /f %%d in ('go list ./...^|findstr -v "github"') do (
	if exist .\test_temp\coverage.out (
		go test -coverprofile=./test_temp/coverage1.out %%d
		if exist .\test_temp\coverage1.out (
			findstr -v "mode": .\test_temp\coverage1.out >> .\test_temp\coverage.out
			@echo off
			del .\test_temp\coverage1.out
		)
	) else (
		go test -coverprofile=./test_temp/coverage.out %%d
	)
)

go tool cover -func=./test_temp/coverage.out -o ./test_temp/coverage.txt
findstr "total" .\test_temp\coverage.txt >> .\test_temp\coverage2.txt
del .\test_temp\coverage.txt

for /f "tokens=1,2,3 delims=	" %%a in (.\test_temp\coverage2.txt) do (
    echo %%a %%c of statements
)
del .\test_temp\coverage2.txt


go tool cover -html=./test_temp/coverage.out -o ./test_temp/coverage.html
if exist .\test_temp\coverage.html (
	.\test_temp\coverage.html
)

:end
echo finished