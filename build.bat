chcp 65001
go build ^
-ldflags "-s -w -X github.com/bearki/gov/conf.Version=0.0.0.1" ^
-o gov.exe ^
.