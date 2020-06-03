### gin demo

* export GOPROXY=https://goproxy.cn
* export GO111MODULE=on

### go mod 使用
* go mod init # 初始化当前目录为模块根目录，生成go.mod, go.sum文件
* go mod download # 下载依赖包
* go mod tidy #整理检查依赖，如果缺失包会下载或者引用的不需要的包会删除
* go mod vendor #复制依赖到vendor目录下面
* go mod 可看完整所有的命令


