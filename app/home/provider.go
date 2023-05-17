package home

import (
	"github.com/gin-gonic/gin"
	"github.com/we7coreteam/w7-rangine-go-skeleton/app/home/command"
	"github.com/we7coreteam/w7-rangine-go-skeleton/app/home/http/controller"
	"github.com/we7coreteam/w7-rangine-go-skeleton/app/home/http/middleware"
	"github.com/we7coreteam/w7-rangine-go-support/src/console"
	http_server "github.com/we7coreteam/w7-rangine-go/src/http/server"
)

type Provider struct {
}

func (provider *Provider) Register(httpServer *http_server.Server, console console.Console) {
	// 注册一个 home:test 命令
	console.RegisterCommand(new(command.Test))

	// 注册一些路由
	httpServer.RegisterRouters(func(engine *gin.Engine) {
		engine.GET("/home/index", middleware.Home{}.Process, controller.Home{}.Index)
	})
}
