package main

import (
	app "github.com/we7coreteam/w7-rangine-go/src"
	"github.com/we7coreteam/w7-rangine-go/src/http/middleware"
	httpserver "github.com/we7coreteam/w7-rangine-go/src/http/server"
	"github.com/we7coreteam/w7-rangine-go-skeleton/app/home"
)

func main() {
	app := app.NewApp()

	httpserver.Use(middleware.PanicHandlerMiddleware{}.GetProcess())
	httpserver.Use(middleware.NewSessionMiddleware(httpserver.GetSession()).Process)

	// 此处注册 provider
	app.GetProviderManager().RegisterProvider(new(home.Provider)).Register()

	app.RunConsole()
}
