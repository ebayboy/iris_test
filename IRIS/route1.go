package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	//handle method
	app.Handle("GET", "/name/{name:string}", func(ctx iris.Context) {
		ctx.Application().Logger().Info("ctx.Params:", ctx.Params())

		//TODO: ctx.Values
		ctx.Application().Logger().Info("ctx.Values:", ctx.Values())
		ctx.HTML("<h1> Hello from /contact </h1>")
	})

	app.Run(iris.Addr(":8080"))
}
