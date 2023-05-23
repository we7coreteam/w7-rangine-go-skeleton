package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/we7coreteam/w7-rangine-go/src/core/err_handler"
	"github.com/we7coreteam/w7-rangine-go/src/http/controller"
)

type Chat struct {
	controller.Abstract
}

func (chat Chat) Index(ctx *gin.Context) {
	ws, err := new(websocket.Upgrader).Upgrade(ctx.Writer, ctx.Request, nil)
	defer ws.Close()

	if err_handler.Found(err) {
		chat.JsonResponseWithoutError(ctx, "hello world!")
		return
	}

}
