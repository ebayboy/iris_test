package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	//checkAgeRange with int type
	app.Macros().Get("int").RegisterFunc("checkAgeRange", func(minAge, maxAge int) func(int) bool {
		return func(age int) bool {
			fmt.Println("age:", age)
			return age >= minAge && age <= maxAge
		}
	})
	app.Get("/age/{age:int checkAgeRange(3,5)}", func(ctx iris.Context) {
		ctx.Writef("checkAgeRange ok!")
	}) //缺少类型= {name：string}

	//checkUserNameLen with string type
	app.Macros().Get("string").RegisterFunc("checkUserNameLen", func(minLength, maxLength int) func(string) bool {
		return func(paramValue string) bool {
			return len(paramValue) >= minLength && len(paramValue) <= maxLength
		}
	})

	//您可以使用“string”类型，该类型对于可以是任何内容的单个路径参数有效。
	app.Get("/username/{name:string checkUserNameLen(5,10)}", func(ctx iris.Context) {
		ctx.Writef("Hello %s", ctx.Params().Get("name"))
	}) //缺少类型= {name：string}

	//check min(id)
	app.Macros().Get("int").RegisterFunc("min", func(minValue int) func(int) bool {
		return func(id int) bool {
			return id >= minValue
		}
	})
	// http://localhost:8080/profile/id>=1
	//这将抛出404，如果请求路由：/ profile/0，/profile/blabla，/profile/-1
	//宏参数函数当然是可选的。
	app.Get("/profile/{id:int min(1)}", func(ctx iris.Context) {
		//第二个参数是错误，但它总是为 nil，因为我们使用宏，
		//验证已经发生。
		id, _ := ctx.Params().GetInt("id")
		ctx.Writef("Hello id: %d", id)
	})

	//更改每个路由的宏响应的错误代码程序：
	//验证错误 则由默认的抛出404 改成 403
	app.Get("/profile/{id:int min(1)}/friends/{friendid:int min(1) else 403}", func(ctx iris.Context) {
		id, _ := ctx.Params().GetInt("id")
		friendid, _ := ctx.Params().GetInt("friendid")
		ctx.Writef("Hello id: %d looking for friend id: %d", id, friendid)
	}) //如果没有传递所有路由的宏，这将抛出e 504错误代码而不是404。

	// http://localhost:8080/game/a-zA-Z/level/0-9
	//http://localhost:8080/game/aA/level/12
	// 记住，alphabetical 只是小写或大写字母。
	app.Get("/game/{name:alphabetical}/level/{level:int}", func(ctx iris.Context) {
		ctx.Writef("name: %s | level: %s", ctx.Params().Get("name"), ctx.Params().Get("level"))
	})

	//让我们使用一个简单的自定义regexp来验证单个路径参数
	//它的值只是小写字母。
	// http://localhost:8080/lowercase/anylowercase
	app.Get("/lowercase/{name:string regexp(^[a-z]+)}", func(ctx iris.Context) {
		ctx.Writef("name should be only lowercase, otherwise this handler will never executed: %s", ctx.Params().Get("name"))
	})

	// http://localhost:8080/single_file/app.js
	app.Get("/single_file/{myfile:file}", func(ctx iris.Context) {
		ctx.Writef("file type validates if the parameter value has a form of a file name, got: %s", ctx.Params().Get("myfile"))
	})

	// http://localhost:8080/myfiles/any/directory/here/
	//这是唯一接受任意数量路径段的宏类型。
	app.Get("/myfiles/{directory:path}", func(ctx iris.Context) {
		//path type accepts any number of path segments, path after /myfiles/ is: any/directory/here
		ctx.Writef("path type accepts any number of path segments, path after /myfiles/ is: %s", ctx.Params().Get("directory"))
	})
	app.Run(iris.Addr(":8080"))
}
