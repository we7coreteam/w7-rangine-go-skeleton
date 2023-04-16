# w7-rangine-go-empty
go版软擎骨架项目

## 创建项目

```
export GOPATH=/tmp && go install github.com/we7coreteam/w7-rangine-go-skeleton@v1.0.2
/tmp/w7-rangine-go-skeleton make:project --dir=/home/
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
