package main

import (
	"github.com/kataras/iris/v12"
)

func main() {

	app := iris.New()

	users := app.Party("/users", myAuthMiddlewareHandler)

	// http://myhost.com/users/42/profile
	users.Get("/{id:int}/profile", userProfileHandler)

	// http://myhost.com/users/messages/1
	users.Get("/messages/{id:int}", userMessageHandler)

	//注册自定义错误页面
	app.RegisterView(iris.HTML("./web/views", ".html"))
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.View("errors/404.html")
	})

	app.Run(iris.Addr(":9999"))
}

func myAuthMiddlewareHandler(ctx iris.Context) {
	ctx.StatusCode(iris.StatusOK) // 默认 200 == iris.StatusOK
	ctx.HTML("<h1>myAuthMiddlewareHandler/</h1>")
	ctx.Next() //如果没有这个， 则进入不到后面的分路由userProfileHandler和userMessageHandler
	//为了执行adminRoutes.Done() 处理程序
}

func userProfileHandler(ctx iris.Context) {
	ctx.HTML("userProfileHandler")
}

func userMessageHandler(ctx iris.Context) {
	ctx.HTML("userMessageHandler")
}
