@echo off

SETLOCAL ENABLEEXTENSIONS
set $path=%~dp0

md dist 1>nul 2>nul

SET GOOS=windows
SET GOARCH=amd64
go build -ldflags "-s -w" -o "%$path%dist\%GOOS%\%GOARCH%\avwx.exe" "%$path%main.go"

SET GOOS=windows
SET GOARCH=386
go build -ldflags "-s -w" -o "%$path%dist\%GOOS%\%GOARCH%\avwx.exe" "%$path%main.go"

ENDLOCAL
