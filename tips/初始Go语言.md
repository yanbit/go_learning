# 初始Go语言



踩坑记录

1.指定国内代理，加快下载以来

国内需要指定下载代理（Go 1.11以上版本支持），否则经常会TimeOut，详细查看go get失败安装日志
指定
go env -w GOPROXY=https://goproxy.cn,direct
查看是否生效
go env ｜ grep GOPROXY
https://goproxy.cn,direct

2.使用go mod管理项目以来
项目文件夹下执行
go mod init {项目名称}
go mod tidy

3. go test没有中文显示
Code -》 Perfernce ——》settings -》go test 
go testflags edit setting.json
    
    "go.testFlags": [
        "-v"
    ],
