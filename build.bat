chcp 65001

SET GOOS=windows
SET GOARCH=386
go build -ldflags "-s -w" -o .\build\gov0.0.5.%GOOS%-%GOARCH%.exe .
upx -9 .\build\gov0.0.5.%GOOS%-%GOARCH%.exe

SET GOOS=windows
SET GOARCH=amd64
go build -ldflags "-s -w" -o .\build\gov0.0.5.%GOOS%-%GOARCH%.exe .
upx -9 .\build\gov0.0.5.%GOOS%-%GOARCH%.exe