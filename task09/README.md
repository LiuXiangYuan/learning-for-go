#### 环境变量
输入`go env`查看当前配置
```
set GO111MODULE=on
set GOARCH=amd64
set GOBIN=
set GOCACHE=C:\Users\dell\AppData\Local\go-build
set GOENV=C:\Users\dell\AppData\Roaming\go\env  
set GOEXE=.exe
set GOFLAGS=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOINSECURE=
set GOMODCACHE=C:\Users\dell\go\pkg\mod
set GONOPROXY=
set GONOSUMDB=
set GOOS=windows
set GOPATH=C:\Users\dell\go
set GOPRIVATE=
set GOPROXY=https://goproxy.io,direct
set GOROOT=C:\Go
set GOSUMDB=sum.golang.org
set GOTMPDIR=
set GOTOOLDIR=C:\Go\pkg\tool\windows_amd64
set GCCGO=gccgo
set AR=ar
set CC=gcc
set CXX=g++
set CGO_ENABLED=1
set GOMOD=NUL
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
set GOGCCFLAGS=-m64 -mthreads -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=C:\Users\dell\AppData\Local\Temp\go-build001990230=/tmp/go-build -gno-record-gcc-switches
```

##### GO111MODULE
- auto：只要项目包含了 go.mod 文件的话启用 Go modules，目前在 Go1.11 至 Go1.14 中仍然是默认值。
- on：启用 Go modules，推荐设置，将会是未来版本中的默认值。
- off：禁用 Go modules，不推荐设置。

##### GOPROXY
此环境变量主要用于设计Go Module的代理

##### GOSUMDB
此环境变量用于在拉取模块的时候保证模块版本数据的一致性

#### go get
`go get`用于拉取新的依赖，以下为go get命令具体用法
| `go get` | **拉取依赖，会进行指定性拉取（更新），并不会更新所依赖的其它模块** |
| :----: | :----: |
| `go get -u` | 更新现有的依赖，会强制更新它所依赖的其它全部模块，不包括自身 |
| `go get -u -t ./...` | 更新所有直接依赖和间接依赖的模块版本，包括单元测试中用到的 |

其他参数
```
-d 只下载不安装
-f 只有在你包含了 -u 参数的时候才有效，不让 -u 去验证 import 中的每一个都已经获取了，这对于本地 fork 的包特别有用
-fix 在获取源码之后先运行 fix，然后再去做其他的事情
-t 同时也下载需要为运行测试所需要的包
-u 强制使用网络去更新包和它的依赖包
-v 显示执行的命令
```

常用命令
```
go mod init  // 初始化go.mod
go mod tidy  // 更新依赖文件
go mod download  // 下载依赖文件
go mod vendor  // 将依赖转移至本地的vendor文件
go mod edit  // 手动修改依赖文件
go mod graph  // 查看现有的依赖结构
go mod verify  // 校验依赖
```