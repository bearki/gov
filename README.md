# Gov - Go multi-version management tool
This is a small and flexible Golang SDK multi-version management tool  

[![License Apache-2.0](https://img.shields.io/badge/license-Apache--2.0-yellow)](https://www.apache.org/licenses/LICENSE-2.0)
[![Golang Version](https://img.shields.io/github/go-mod/go-version/bearki/gov?filename=go.mod)](https://go.dev/dl)
![](https://img.shields.io/badge/platform-windows%20%7C%20macos%20%7C%20linux-orange)
[![Github Release](https://img.shields.io/github/v/release/bearki/gov)](https://github.com/bearki/gov/releases)
[![Github Last Commit](https://img.shields.io/github/last-commit/bearki/gov)](https://github.com/bearki/gov/commits/main)  

![Gov LOGO](https://qiniu.github.bearki.cn/gov/gov-log.png)

## 一、提示
> 经过一系列优化，该工具可在无任何配置状态下直接使用，是的，非常爽；在Windows、MacOS、Linux上你都能得到极致体验，多线程下载功能也提上日程，敬请期待。

## 二、环境变量解释
> 有两个建议的环境变量【 GOSDKPATH，GOROOT 】，以及两个可选的环境变量【 GOSDKVERURL，GOSDKDOWNURL 】需要配置,下面详细解释一下这几个环境变量的作用。

### GOSDKPATH
> Gov在启动时会去操作系统中获取该环境变量的值，该值应该是一个文件夹，用于储存Gov工具所需要的依赖文件以及Golang SDK的各个版本文件，所以该环境变量至关重要，在不配置该环境变量时，Gov将会自动选择合适和目录作为GOSDKPATH的值（`Windows: %LOCALAPPDATA%\Gov` `Unix: $HOME/Gov`），由于工作目录容易变化，因此我们强烈建议在操作系统中为该环境变量赋值。

### GOROOT
> 这个是Golang SDK的环境变量，表示Golang SDK的安装位置，在不配置该环境变量时，Gov将会自动选择合适和目录作为GOROOT的值（`Windows: %LOCALAPPDATA%\Go` `Unix: $HOME/Go`），；需要注意的是GOSDKPATH与GOROOT的路径不要存在嵌套关系，否则会引起致命错误，同时不建议将GOROOT配置到受保护的目录，否则每次都需要以管理员权限启动终端才能切换版本。

### GOSDKVERURL
> 当你使用Gov获取版本列表失败时，那么你应该需要配置该环境变量了，该环境变量可用值如下，后续会陆续更新到此处(我们将会在该URL之后追加参数获取版本列表，请确保能正常访问，例如：https://golang.google.cn/dl/?mode=json&include=all)：
> * https://qiniu.github.bearki.cn/gov/version/version.json （default）
> * https://golang.google.cn/dl/
> * https://go.dev/dl/

### GOSDKDOWNURL
> 当你使用Gov下载Go SDK失败时，那么你应该需要配置该环境变量了，该环境变量可用值如下，后续会陆续更新到此处(我们将会在该URL之后追加文件名称直接下载，请确保能正常访问，例如：https://golang.google.cn/dl/go1.10.1.windows-386.zip)：
> * https://mirrors.ustc.edu.cn/golang （default）
> * https://golang.google.cn/dl
> * https://go.dev/dl

## 三、配置环境变量
### Windows环境 
1. 打开环境变量配置页面   
![](https://qiniu.github.bearki.cn/gov/gov-windows-env-1.png)  
2. 新增两个至关重要的环境变量  
![](https://qiniu.github.bearki.cn/gov/gov-windows-env-2.png)  
3. 将GO SDK可执行文件目录添加到PATH，使其在任意位置可以访问  
![](https://qiniu.github.bearki.cn/gov/gov-windows-env-3.png)  
4. 然后一直确认，最后关闭窗口即可  
  
### Unix
* 顺手把GOPATH也配置了
```shell
export GOSDKVERURL="https://golang.google.cn/dl"
export GOSDKDOWNURL="https://golang.google.cn/dl"
export GOSDKPATH="$HOME/Gov"
export GOROOT="$HOME/Go"

export GOPATH="$HOME/gopath"
export PATH="$GOROOT/bin:$GOPATH/bin:$PATH"
```

## 四、使用方式
使用help命令你将会得到详细的使用介绍及示例
```shell
gov help
```

## 五、更新Gov
1. 使用Go来更新Gov
```shell
go install github.com/bearki/gov@latest
```
2. 前往发布页下载二进制包  
[最新发布页](https://github.com/bearki/gov/releases)
