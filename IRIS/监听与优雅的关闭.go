package main

import (
	"net"
	"net/http"
	"os"

	"github.com/kataras/iris/v12"
)

func main() {

	app := iris.New()

	//方式1: 推荐 监听`tcp` 的 `0.0.0.0:8080` 网络地址 也就是 `net.Listener` 类型
	app.Run(iris.Addr(":8080"))

	//方式2： 与上面之前一样的，但是使用自定义`HTTP.Server`也可能在其他地方使用达到相同的效果。
	app.Run(iris.Server(&http.Server{Addr: ":8080"}))

	//方式3： 使用自定义的 `net.Listener`
	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		panic(err)
	}
	app.Run(iris.Listener(l))

	//使用文件 安全传输层协议`TLS`
	//mycer.cert 证书： 即是公钥;  mykey.key： 私钥， 解密使用公钥加密后的内容
	app.Run(iris.TLS("127.0.0.1:443", "mycert.cert", "mykey.key"))

	// 自动 安全传输层协议`TLS`
	app.Run(iris.AutoTLS(":443", "example.com", "admin@example.com"))

	// UNIX 套接字
	if errOs := os.Remove(socketFile); errOs != nil && !os.IsNotExist(errOs) {
		app.Logger().Fatal(errOs)
	}

	l, err := net.Listen("unix", socketFile)
	if err != nil {
		app.Logger().Fatal(err)
	}

	if err = os.Chmod(socketFile, mode); err != nil {
		app.Logger().Fatal(err)
	}

	app.Run(iris.Listener(l))

	//使用任何 func（）error，
	//启动监听者的责任取决于你这个方式，
	//为了简单起见，我们将使用
	//`net/http`包的ListenAndServe函数。
	app.Run(iris.Raw(&http.Server{Addr: ":8080"}).ListenAndServe)

}
