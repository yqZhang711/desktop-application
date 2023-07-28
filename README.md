# desktop-application


### 运行环境
```azure
  golang 1.19+
```

### MacOs 打包程序
MacOs系统不支持11一下的版本
```azure
   go mod tidy
   
   ./macBuild.sh
```

### Windows 打包程序
下载安装gcc: https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download-section
使用cmd打开命令行窗口
```azure
   go mod tidy
   
   set CGO_CPPFLAGS="-I%cd%\libs\webview2\build\native\include"

   go build -ldflags="-H windowsgui"
```