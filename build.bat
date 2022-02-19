echo off
chcp 65001

SET VERSION=0.1.2
SET CGO_ENABLED=0


@REM ----------------------------- OS Windows -----------------------------
SET GOOS=windows

SET GOARCH=386
SET OUTFILE=.\build\gov%VERSION%.%GOOS%-%GOARCH%.exe
go build -ldflags "-s -w" -o %OUTFILE% .
upx -9 %OUTFILE%

SET GOARCH=amd64
SET OUTFILE=.\build\gov%VERSION%.%GOOS%-%GOARCH%.exe
go build -ldflags "-s -w" -o %OUTFILE% .
upx -9 %OUTFILE%


@REM ------------------------------ OS Linux ------------------------------
SET GOOS=linux

SET GOARCH=amd64
SET OUTFILE=.\build\gov%VERSION%.%GOOS%-%GOARCH%
go build -ldflags "-s -w" -o %OUTFILE% .
upx -9 %OUTFILE%

SET GOARCH=386
SET OUTFILE=.\build\gov%VERSION%.%GOOS%-%GOARCH%
go build -ldflags "-s -w" -o %OUTFILE% .
upx -9 %OUTFILE%

SET GOARCH=arm64
SET OUTFILE=.\build\gov%VERSION%.%GOOS%-%GOARCH%
go build -ldflags "-s -w" -o %OUTFILE% .
upx -9 %OUTFILE%


@REM ------------------------------ OS MacOS ------------------------------
SET GOOS=darwin

SET GOARCH=amd64
SET OUTFILE=.\build\gov%VERSION%.%GOOS%-%GOARCH%
go build -ldflags "-s -w" -o %OUTFILE% .
upx -9 %OUTFILE%

SET GOARCH=arm64
SET OUTFILE=.\build\gov%VERSION%.%GOOS%-%GOARCH%
go build -ldflags "-s -w" -o %OUTFILE% .
upx -9 %OUTFILE%