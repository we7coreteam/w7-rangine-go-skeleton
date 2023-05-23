package main

import (
	"github.com/gin-gonic/gin"
	"github.com/we7coreteam/w7-rangine-go-skeleton/app/home"
	app "github.com/we7coreteam/w7-rangine-go/src"
	"github.com/we7coreteam/w7-rangine-go/src/http"
	"github.com/we7coreteam/w7-rangine-go/src/http/middleware"
	"io"
	http2 "net/http"
)

func main() {
	app := app.NewApp()

	// 业务中需要使用 http server，这里需要先实例化
	httpServer := new(http.Provider).Register(app.GetConfig(), app.GetConsole(), app.GetServerManager()).Export()
	// 注册一些全局中间件，路由或是其它一些全局操作
	httpServer.Use(middleware.GetPanicHandlerMiddleware(nil))
	httpServer.RegisterRouters(func(engine *gin.Engine) {
		engine.GET("favicon.ico", func(context *gin.Context) {
			response, _ := http2.Get("https://g.csdnimg.cn/static/logo/favicon32.ico")
			defer response.Body.Close()
			data, _ := io.ReadAll(response.Body)
			context.Data(200, "image/x-icon", data)
		})
		engine.NoRoute(func(context *gin.Context) {
			context.String(404, "404 Not Found")
		})
	})

	// 注册业务 provider，此模块中需要使用 http server 和 console
	new(home.Provider).Register(httpServer, app.GetConsole())
	app.RunConsole()
}
