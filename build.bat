chcp 65001

SET VERSION=0.0.6

SET GOOS=windows
SET GOARCH=386
SET OUTFILE=.\build\gov%VERSION%.%GOOS%-%GOARCH%.exe
go build -ldflags "-s -w" -o %OUTFILE% .
upx -9 %OUTFILE%

SET GOOS=windows
SET GOARCH=amd64
SET OUTFILE=.\build\gov%VERSION%.%GOOS%-%GOARCH%.exe
go build -ldflags "-s -w" -o %OUTFILE% .
upx -9 %OUTFILE%