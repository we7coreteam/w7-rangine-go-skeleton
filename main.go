package main

import (
	"github.com/we7coreteam/w7-rangine-go-skeleton/app/home"
	app "github.com/we7coreteam/w7-rangine-go/src"
	"github.com/we7coreteam/w7-rangine-go/src/http"
)

func main() {
	app := app.NewApp()

	// 业务中需要使用 http server，这里需要先实例化
	httpServer := new(http.Provider).Register(app.GetConfig(), app.GetConsole(), app.GetServerManager()).Export()

	// 注册业务 provider，此模块中需要使用 http server 和 console
	new(home.Provider).Register(httpServer, app.GetConsole())

	app.RunConsole()
}
