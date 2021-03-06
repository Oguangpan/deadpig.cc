# Golang 自学过程中总结的速查手册
----
## 代码
### 语法
#### 基本结构和基本数据类型
#### 数组于切片
#### Map字典
#### 包

### 资源管理
#### go doc 的使用
如何查看相应package的文档呢？ 例如builtin包，那么执行```go doc builtin``` 如果是http包，那么执行```go doc net/http``` 查看某一个包里面的函数，那么执行```godoc fmt Printf``` 也可以查看相应的代码，执行```godoc -src fmt Printf```

通过命令在命令行执行 ```godoc -http=:端口号``` 比如```godoc -http=:8080```。然后在浏览器中打开127.0.0.1:8080，你将会看到一个golang.org的本地copy版本，通过它你可以查询pkg文档等其它内容。如果你设置了GOPATH，在pkg分类下，不但会列出标准包的文档，还会列出你本地GOPATH中所有项目的相关文档，这对于经常被墙的用户来说是一个不错的选择。
#### 查看当前go环境的信息
```go version```    
查看当前golang的版本     
```go env```    
查看当前所有环境变量
### 格式化代码
```go fmt ```    
命令可格式化代码成为标准格式. 一般现代编辑器软件自带该功能.

## 范式
### 控制结构
### 函数
### 结构与方法
### 接口与反射

## 产出
### 如何编译程序
```go build ```

编译包和包的依赖关系

- 普通包不会生成任何东西
- main包会在当前目录生成文件
- 可指定要便宜的文件名
- 主动忽略 _ 或 . 开头的go文件

```go run```     
编译并运行当前程序

```go clean```     
清除当前源码包里面经过编译生成的文件. 编译完成后使用该命令可以清洁目录.

```go install```    
该命令包含两个方面,首先编译目标,并且将目标生成的可执行文件放到bin,.a链接文件放到pkg中,实现整个项目的部署.
### 库
库的使用更多的是他人的库,而他人的库都放在类似github之类的地方.所以golang自带方便的部署远程库的工具.

#### git
相关资料请查看我的另外一篇文章 [如何使用git]()

### go get 
这个命令是用来动态获取远程代码包的，目前支持的有BitBucket、GitHub、Google Code和Launchpad。这个命令在内部实际上分成了两步操作：第一步是下载源码包，第二步是执行go install。    
```go get https://github.com/xxx/xxxx```

### go list
该命令可以方便的列出本机上已经安装的所有库的列表.

## 工具
### 编辑器选择
- vscode
- liteide
- vim

### 测试项目文件
```go test```
执行这个命令，会自动读取源码目录下面名为*_test.go的文件，生成并运行测试用的可执行文件.