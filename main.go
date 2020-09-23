package main

import (
	"fmt"
	"irisORM/configs"
	"irisORM/extra"
	"strconv"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func main() {
	fmt.Printf("get config %v ,%v\n", configs.Config.Common.Port, configs.Config.Db.Connstr)
	// 注册模板在work目录的views文件夹
	app := iris.New()
	app.RegisterView(iris.HTML("./view", ".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello world!")
		ctx.View("old.html")
	})

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Authorization"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})
	app.Get("/ping", crs, func(ctx iris.Context) {
		res := ctx.Request().URL.Query()
		fmt.Println("account", res.Get("account"))
		fmt.Println("pwd", ctx.URLParam("pwd"))
		ctx.JSON(iris.Map{
			"query": res,
		})
	})
	app.Get("/select", crs, func(ctx iris.Context) {
		i, _ := strconv.Atoi(ctx.URLParam("id"))
		ctx.JSON(extra.FindOne([]string{"account", "pwd", "name"}, "account", i))
	})
	app.Run(iris.Addr(configs.Config.Common.Port))
}
