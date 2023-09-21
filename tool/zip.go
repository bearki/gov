package tool

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// DeCompressZip Zip解压文件
//
//	@var	zipFile	压缩文件路径
//	@var	dstDir	解压后的文件夹路径
//	@return	错误信息
func DeCompressZip(zipFile string, dstDir string) error {
	// 不管三七二十一，先创建目标文件夹
	err := os.MkdirAll(dstDir, 0755)
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
	for _, f := range reader.File {
		// 转换为目标文件路径
		dstFileName := strings.ReplaceAll(filepath.Join(f.Name), "\\", "/")
		// 移除前置go目录
		dstFileName = strings.TrimPrefix(dstFileName, "go/")
		// 文件或文件夹的目标路径
		dstFilePath := filepath.Join(dstDir, dstFileName)
		// 判断文件或文件夹是否已存在
		_, err = os.Stat(dstFilePath)
		if err == nil {
			// 跳过该文件
			continue
		}
		// 判断是文件还是文件夹
		if f.FileInfo().IsDir() {
			continue
		}
		// 直接创建文件夹即可
		err = os.MkdirAll(filepath.Dir(dstFilePath), 0755)
		if err != nil {
			return err
		}
		// 在函数中使用defer
		err = func() error {
			// 打开源文件
			srcFile, e := f.Open()
			if e != nil {
				return e
			}
			defer srcFile.Close()
			// 创建目标文件
			dstFile, e := os.Create(dstFilePath)
			if e != nil {
				return e
			}
			defer dstFile.Close()
			// 拷贝整个数据
			_, e = io.Copy(dstFile, srcFile)
			if e != nil {
				return e
			}
			// OK
			return nil
		}()
		if err != nil {
			return err
		}
	}

	// 全部解压完毕
	return nil
}

// DeCompressGzip Gzip解压文件
//
//	@var	gzipFile	压缩文件路径
//	@var	dstDir		解压后的文件夹路径
//	@return	错误信息
func DeCompressGzip(gzipFile string, dstDir string) error {
	// 不管三七二十一，先创建目标文件夹
	err := os.MkdirAll(dstDir, 0755)
	if err != nil {
		return err
	}
	// 打开文件流
	reader, err := os.Open(gzipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	// 使用gzip读取文件流
	gr, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer gr.Close()
	// 使用tar展开文件
	tr := tar.NewReader(gr)
	// 遍历压缩包的内容
	for {
		// 下一跳
		h, err := tr.Next()
		// 是否读取结束
		if err == io.EOF {
			// OK
			return nil
		}
		if err != nil {
			return err
		}
		// 转换为目标文件路径
		dstFileName := strings.ReplaceAll(filepath.Join(h.Name), "\\", "/")
		// 移除前置go目录
		dstFileName = strings.TrimPrefix(dstFileName, "go/")
		// 文件或文件夹的目标路径
		dstFilePath := filepath.Join(dstDir, dstFileName)
		// 判断文件或文件夹是否已存在
		_, err = os.Stat(dstFilePath)
		if err == nil {
			// 跳过该文件
			continue
		}
		// 判断是文件还是文件夹
		if h.FileInfo().IsDir() {
			continue
		}
		// 直接创建文件夹即可
		err = os.MkdirAll(filepath.Dir(dstFilePath), 0755)
		if err != nil {
			return err
		}
		// 在函数中使用defer
		err = func() error {
			// 创建目标文件
			dstFile, e := os.Create(dstFilePath)
			if e != nil {
				return e
			}
			defer dstFile.Close()
			// 拷贝整个数据
			_, e = io.Copy(dstFile, tr)
			if e != nil {
				return e
			}
			// OK
			return nil
		}()
		// 检查
		if err != nil {
			return err
		}
	}
}
