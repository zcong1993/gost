package main

import "github.com/kataras/iris"

func main() {
	app := iris.Default()

	app.Run(iris.Addr("localhost:9393"))
}
