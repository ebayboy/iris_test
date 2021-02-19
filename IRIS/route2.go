package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	//这个例子完美的解释了，上面的定义
	app := iris.New()

	// Method: "GET"
	app.Get("/", handler)

	// Method: "POST"
	app.Post("/", handler)

	// 处理除GET和POST之外的方法
	app.Any("/", handlerAny)

	app.Run(iris.Addr(":8080"))
}

func writeResp(ctx iris.Context) {
	for _, v := range ctx.Request().Cookies() {
		ctx.Writef("Name:%s Value:%s\n", v.Name, v.Value)
	}
}

func handler(ctx iris.Context) {
	ctx.Writef("handler Hello from method: %s and path: %s\n", ctx.Method(), ctx.Path())
	writeResp(ctx)
}

func handlerAny(ctx iris.Context) {
	ctx.Writef("handlerAny Hello from method: %s and path: %s\n", ctx.Method(), ctx.Path())
	writeResp(ctx)
}
