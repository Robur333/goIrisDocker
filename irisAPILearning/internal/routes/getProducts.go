package routes

import "github.com/kataras/iris/v12"

func GetProducts(ctx iris.Context) {
	ctx.JSON(iris.Map{"fajne" : "produkty"})
}