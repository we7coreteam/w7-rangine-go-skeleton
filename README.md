# w7-rangine-go-empty

go版软擎骨架项目

文档地址： https://wiki.w7.com/document/2315/7788

## 创建项目

#### 安装 cli 工具
```
export GOPATH=/tmp && go install -x github.com/we7coreteam/w7-rangine-go-skeleton@v1.0.5
```

#### 创建项目
```
/tmp/w7-rangine-go-skeleton make:project /home/root/my-go-project

# 进入目录，初始化，编译
cd /home/root/my-go-project
go get -u && go mod tidy
go build -o ./bin/rangine .
```

## 辅助命令

### 生成一个 Provider

```shell
./bin/rangine make:module [flags]
```

- \-\-name 要创建的模块名称

### 启动服务

```shell
./bin/rangine server:start
```

### 查看所有注册路由

```shell
./bin/rangine route:list
```

### 查看帮助信息

```shell
./bin/rangine -h
```

### 查看框架版本

```shell
./bin/rangine version
```

### 热启动

```shell
./reload.sh
```

### 编译

```shell
# osx
make build-osx 

# linux
make build

# windows
make build-windows
```
