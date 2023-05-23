package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/we7coreteam/w7-rangine-go-skeleton/app/ws/command"
	"github.com/we7coreteam/w7-rangine-go-skeleton/app/ws/http/controller"
	"github.com/we7coreteam/w7-rangine-go-skeleton/app/ws/http/middleware"
	"github.com/we7coreteam/w7-rangine-go-support/src/console"
	http_server "github.com/we7coreteam/w7-rangine-go/src/http/server"
)

type Provider struct {
}

func (provider *Provider) Register(httpServer *http_server.Server, console console.Console) {
	// 注册一个 ws:test 命令
	console.RegisterCommand(new(command.Test))

	// 注册一些路由
	httpServer.RegisterRouters(func(engine *gin.Engine) {
		engine.GET("/ws/index", middleware.Home{}.Process, controller.Home{}.Index)
	})

	httpServer.RegisterRouters(func(engine *gin.Engine) {
		engine.GET("/ws", controller.Chat{}.Index)
	})
}
