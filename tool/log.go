package tool

import (
	"fmt"

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
