chcp 65001

SET GOOS=windows
SET GOARCH=386
go build -ldflags "-s -w" -o gov0.0.4.%GOOS%-%GOARCH%.exe .
upx gov0.0.4.%GOOS%-%GOARCH%.exe

SET GOOS=windows
SET GOARCH=adm64
go build -ldflags "-s -w" -o gov0.0.4.%GOOS%-%GOARCH%.exe .
upx gov0.0.4.%GOOS%-%GOARCH%.exe