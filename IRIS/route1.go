package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	//handle method
	app.Handle("GET", "/contact", func(ctx iris.Context) {
		ctx.HTML("<h1> Hello from /contact </h1>")
	})

	//post method
	app.Post("/post", func(ctx iris.Context) {
		ctx.HTML("<h1>Post path</h1>")
	})

	app.Run(iris.Addr(":8080"))
}
