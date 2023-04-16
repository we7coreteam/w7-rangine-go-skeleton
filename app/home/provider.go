package home

import (
	"github.com/gin-gonic/gin"
	"github.com/we7coreteam/w7-rangine-go/src/core/provider"
	http_server "github.com/we7coreteam/w7-rangine-go/src/http/server"
	"github.com/we7coreteam/w7-rangine-go-skeleton/app/home/command"
	"github.com/we7coreteam/w7-rangine-go-skeleton/app/home/http/controller"
	"github.com/we7coreteam/w7-rangine-go-skeleton/app/home/http/middleware"
)

type Provider struct {
	provider.Abstract
}

func (provider *Provider) Register() {
	provider.GetConsole().RegisterCommand(new(command.Test))

	http_server.RegisterRouters(func(engine *gin.Engine) {
		engine.GET("/home/index", middleware.Home{}.Process, controller.Home{}.Index)
	})
}
