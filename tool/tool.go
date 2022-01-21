package tool

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
)

const StartLine = `
-------------------------------- Gov Start --------------------------------
`
const EndLine = `
--------------------------------- Gov End ---------------------------------
`

// 日志对象
type Log struct {
	newLine string
}

// 实例化日志
var L Log = Log{
	newLine: "\r\n",
}

// print log
func (l *Log) print(newColor color.Color, format string, val ...interface{}) {
	if len(val) > 0 {
		fmt.Println(newColor.Sprintf(format, val...))
		return
	}
	fmt.Println(newColor.Sprintf(format))
}

// error log
func (l *Log) Error(format string, val ...interface{}) {
	l.print(color.Red, format, val...)
}

// Warning log
func (l *Log) Warn(format string, val ...interface{}) {
	l.print(color.Yellow, format, val...)
}

// Success log
func (l *Log) Success(format string, val ...interface{}) {
	l.print(color.Green, format, val...)
}

// Trace log
func (l *Log) Trace(format string, val ...interface{}) {
	l.print(color.Blue, format, val...)
}

// Info log
func (l *Log) Info(format string, val ...interface{}) {
	l.print(color.White, format, val...)
}

// Math data sha356 value
func MathSha256(data []byte) string {
	str := sha256.Sum256(data)
	return hex.EncodeToString(str[:])
}

// DeCompressZip Zip解压文件
// @params zipFile  string 压缩文件路径
// @params dstPath  string 解压后的文件夹路径
// @return          error  错误信息
func DeCompressZip(zipFile string, dstPath string) error {
	// 不管三七二十一，先创建目标文件夹
	err := os.MkdirAll(dstPath, 0755)
	if err != nil {
		return err
	}
	// 使用zip打开压缩包
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	// 遍历压缩包的内容
	// 注意：reader.File获取到的是压缩包内的所有文件，包括子文件夹下的文件
	for _, file := range reader.File {
		// 去掉第一层文件夹
		newFileName := strings.ReplaceAll(filepath.Join(file.Name), "\\", "/")
		pathList := strings.Split(newFileName, "/")
		if len(pathList) >= 2 {
			pathList = pathList[1:]
		}
		// 文件或文件夹的目标路径
		dstFile := filepath.Join(pathList...)
		dstFile = filepath.Join(dstPath, dstFile)
		// 文件夹就不解压出来了
		if file.FileInfo().IsDir() { // 文件夹
			// 不管三七二十一，先创建目标文件夹
			err := os.MkdirAll(dstFile, 0755)
			if err != nil {
				return err
			}
			// 文件夹创建完毕就跳过吧
			continue
		} else { // 文件
			// 判断文件是否已存在
			_, err = os.Stat(dstFile)
			if err == nil {
				// 跳过该文件
				continue
			}
			// 打开压缩包内的文件
			srcFile, err := file.Open()
			if err != nil {
				return err
			}
			defer srcFile.Close()
			// 在文件夹内创建这个文件
			destFile, err := os.Create(dstFile)
			if err != nil {
				return err
			}
			defer destFile.Close()
			// 执行文件拷贝
			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				return err
			}
		}
	}
	// 全部解压完毕
	return nil
}
