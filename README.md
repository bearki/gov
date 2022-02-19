## Gov - Go multi-version management tool
This is a small and flexible Golang SDK multi-version management tool

## 一、提示
> 经过一系列优化，该工具可在无任何配置状态下直接使用，是的，非常爽，但是目前在windows环境下有缺陷，首次使用gov安装golang环境后会出现go命令找不到的情况，这是由于windows全局环境变量未刷新造成的，注销一下用户或重启计算机即可使用，如有解决方案请随时通知我。

## 二、环境变量解释
> 有两个强烈建议的环境变量【 GOSDKPATH，GOROOT 】，以及两个可选的环境变量【 GOSDKVERURL，GOSDKDOWNURL 】需要配置,下面详细解释一下这几个环境变量的作用。

### GOSDKPATH
> Gov在启动时会去操作系统中获取该环境变量的值，该值应该是一个文件夹，用于储存Gov工具所需要的依赖文件以及Golang SDK的各个版本文件，所以该环境变量至关重要，在不配置该环境变量时，Gov将会自动选择合适和目录作为GOSDKPATH的值（`windows: %LOCALAPPDATA%\Gov` `linux: $HOME/Gov`），由于工作目录容易变化，因此我们强烈建议在操作系统中为该环境变量赋值。

### GOROOT
> 这个是Golang SDK的环境变量，表示Golang SDK的安装位置，在不配置该环境变量时，Gov将会自动选择合适和目录作为GOROOT的值（`windows: %LOCALAPPDATA%\Go` `linux: $HOME/Go`），；需要注意的是GOSDKPATH与GOROOT的路径不要存在嵌套关系，否则会引起致命错误，同时不建议将GOROOT配置到受保护的目录，否则每次都需要以管理员权限启动终端才能切换版本。

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
![avatar](https://qiniu.github.bearki.cn/gov/gov-windows-env-1.png)  
2. 新增两个至关重要的环境变量  
![avatar](https://qiniu.github.bearki.cn/gov/gov-windows-env-2.png)  
3. 将GO SDK可执行文件目录添加到PATH，使其在任意位置可以访问  
![avatar](https://qiniu.github.bearki.cn/gov/gov-windows-env-3.png)  
4. 然后一直确认，最后关闭窗口即可  
  
### Linux
* 顺手把GOPATH也配置了
```shell
export GOSDKVERURL="https://golang.google.cn/dl"
export GOSDKDOWNURL="https://golang.google.cn/dl"
export GOSDKPATH="$HOME/gov"
export GOROOT="$HOME/go"
export GOPATH="$HOME/gopath"
export PATH="$GOROOT/bin:$GOPATH/bin:$PATH"
```

## 四、使用方式
使用help命令你将会得到详细的使用介绍及示例
```shell
# windows环境下一般为单账户，故可直接使用gov命令
gov help

# linux环境或类linux环境一般至少会存在两个账户，并且会经常切换账户，故推荐gov命令前加sudo
sudo gov help
```

## 五、更新Gov
1. 使用Go来更新Gov
```shell
go install github.com/bearki/gov@latest
```
2. 前往发布页下载二进制包https://github.com/bearki/gov/releases
