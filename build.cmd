@echo off
SETLOCAL ENABLEEXTENSIONS

md dist 1>nul 2>nul

SET GOOS=windows
SET GOARCH=amd64
go build -ldflags "-s -w" -o dist/%GOOS%/%GOARCH%/avwx.exe ./cmd/avwx/main.go

SET GOOS=windows
SET GOARCH=386
go build -ldflags "-s -w" -o dist/%GOOS%/%GOARCH%/avwx.exe ./cmd/avwx/main.go

ENDLOCAL
