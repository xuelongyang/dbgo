@echo off
set BUILD_DIR=.\release
set PROJECT_DIR=.

rem Check if build directory exists
if exist %BUILD_DIR% (
    rem Delete existing executable files
    del /Q %BUILD_DIR%\linux-x86_64\dbgo
    del /Q %BUILD_DIR%\linux-arm64\dbgo
    del /Q %BUILD_DIR%\windows-amd64\dbgo.exe
) else (
    mkdir %BUILD_DIR%
)

rem Set GOOS to linux and GOARCH to 386
go env -w GOOS=linux
go env -w GOARCH=386
go build -o %BUILD_DIR%\linux-x86_64\dbgo
copy %PROJECT_DIR%\LICENSE %BUILD_DIR%\linux-x86_64\LICENSE
copy %PROJECT_DIR%\README.md %BUILD_DIR%\linux-x86_64\README.md

rem Set GOOS to linux and GOARCH to arm64
go env -w GOOS=linux
go env -w GOARCH=arm64
go build -o %BUILD_DIR%\linux-arm64\dbgo
copy %PROJECT_DIR%\LICENSE %BUILD_DIR%\linux-arm64\LICENSE
copy %PROJECT_DIR%\README.md %BUILD_DIR%\linux-arm64\README.md

rem Set GOOS to windows and GOARCH to amd64
go env -w GOOS=windows
go env -w GOARCH=amd64
go build -o %BUILD_DIR%\windows-amd64\dbgo.exe
copy %PROJECT_DIR%\LICENSE %BUILD_DIR%\windows-amd64\LICENSE
copy %PROJECT_DIR%\README.md %BUILD_DIR%\windows-amd64\README.md

rem Reset GOOS and GOARCH
go env -u GOOS
go env -u GOARCH

echo Build completed.
