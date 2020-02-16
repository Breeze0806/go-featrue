# go1.13新增特性
介绍主要的golang1.13的相关特性改动
## 运行环境
支持Android 10，windows 7是其最小实用版本？
##工具变更
### 环境变量
GO111MODULE=auto表明module-aware mode下,go tool默认不再使用传统GOPATH下或vendor下面的包，
为了兼容原有项目，建议直接GO111MODULE=off
### go doc
go doc的http服务器功能现在要用[golang.org/x/tools/cmd/godoc](https://gihtub.com/golang/tools)替代使用
### 编译工具链
golang采用新的逃逸分析办法，将原有更多的堆分配分配到了栈上，如要使用原先的，在编译时加入go build -gcflags=all=-newescape=false
##运行时调整
+ defer效率提升30%，未压测，之后增加
+ 数组溢出时会报告那个下标溢出，未有样例，之后增加
+ 现在堆缩小后会立刻想系统返还内存，而不是想之前那样保留4到5分钟，
但是由于大多数操作系统用懒汉式方式统计内存，RSS内存不会立刻下降。
### 数字字面量特性
打开[number literal](numberliteral/README.md)的说文档即可查看，numberliteral包中有这些特性的用例
## 包特性变更说明
### TLS 1.3
1.13默认使用tls1.3，如要去掉该特性设置GODEBUG=tls13=0，在1.14会异常该选项，在1.12该选项GODEBUG=tls13=1才默认启用tls1.3
### crypto/ed25519
新增签名加密算法Ed25519
### errors包特性说明
打开[errors](errors/README.md)的说文档即可查看，errors包中有这些特性的用例
### 其他包特性改动
打开[others](others/README.md)的说文档即可查看，others包中有这些特性的用例
### 未涉及到的包特性
#### net
ListenConfig.KeepAlive设置监听连接的保活时间
#### net/http
相关改动设计http2，不做介绍
## 参考文献
[go1.13文档](https://golang.google.cn/doc/go1.13)