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
		if pathList[0] == "go" {
			continue
		}
		// 文件或文件夹的目标路径
		dstFile := filepath.Join(pathList...)
		dstFile = filepath.Join(dstPath, dstFile)
		// 判断文件或文件夹是否已存在
		_, err = os.Stat(dstFile)
		if err == nil {
			// 跳过该文件
			continue
		}
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

// DeCompressGzip Gzip解压文件
// @params gzipFile string 压缩文件路径
// @params dstPath  string 解压后的文件夹路径
// @return          error  错误信息
func DeCompressGzip(gzipFile string, dstPath string) error {
	// 不管三七二十一，先创建目标文件夹
	err := os.MkdirAll(dstPath, 0755)
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
			return nil
		}
		if err != nil {
			return err
		}
		// 去掉第一层文件夹
		newFileName := strings.ReplaceAll(filepath.Join(h.Name), "\\", "/")
		pathList := strings.Split(newFileName, "/")
		if len(pathList) >= 2 {
			pathList = pathList[1:]
		}
		if pathList[0] == "go" {
			continue
		}
		// 文件或文件夹的目标路径
		dstFile := filepath.Join(pathList...)
		dstFile = filepath.Join(dstPath, dstFile)
		// 判断文件或文件夹是否已存在
		_, err = os.Stat(dstFile)
		if err == nil {
			// 跳过该文件
			continue
		}
		// 判断是文件还是文件夹
		if h.FileInfo().IsDir() {
			// 直接创建文件夹即可
			err = os.MkdirAll(dstFile, 0755)
			if err != nil {
				return err
			}
			continue
		}
		// 在函数中使用defer
		err = func() error {
			// 创建或打开这个文件
			fw, e := os.OpenFile(
				dstFile,
				os.O_CREATE|os.O_WRONLY,
				os.FileMode(h.Mode),
			)
			if e != nil {
				return e
			}
			defer fw.Close()
			// 拷贝整个数据
			_, e = io.Copy(fw, tr)
			if e != nil {
				return e
			}
			return nil
		}()
		if err != nil {
			return err
		}
	}
}
