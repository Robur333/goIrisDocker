package internal

import (
	"restapi/internal/routes"
	"restapi/utils"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func IrisMainApp() *iris.Application {
	
	app := iris.Default()

	verifier := jwt.NewVerifier(jwt.HS256, utils.SignKey)
	verifier.WithDefaultBlocklist()
	verifyMiddleware := verifier.Verify(func() interface{} {
		return new(utils.UserClaims)
	})

	app.Get("/user/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef("Hello %s", name)
	})

	app.Get(utils.ApiPaths.GetAllUsers, routes.GetAllUsers)

	app.Post(utils.ApiPaths.Login, routes.Login)

	app.Post(utils.ApiPaths.Register, routes.Register)

	app.Get(utils.ApiPaths.Products, routes.GetProducts)

	protectedAPI := app.Party(utils.ApiPaths.Protected)

	protectedAPI.Use(verifyMiddleware)

	protectedAPI.Put(utils.ApiPaths.UpdateUser, routes.UpdateUser)

	protectedAPI.Delete(utils.ApiPaths.DeleteUser, routes.DeleteUser)

	protectedAPI.Get("/test", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"test": "dobra zabezpieczona"})
	})
	return app
}