package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	//handle method
	app.Handle("GET", "/contact", func(ctx iris.Context) {
		ctx.HTML("<h1> Hello from /contact </h1>")
	})

}
