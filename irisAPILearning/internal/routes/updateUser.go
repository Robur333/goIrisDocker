package routes

import (
	"restapi/utils"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func UpdateUser(ctx iris.Context) {

	db := utils.ConnectToDatabase()
	defer db.Close()

	tokenClaims := jwt.Get(ctx).(*utils.UserClaims)
	
	userNewLogin := ctx.URLParam("userNewLogin")

	_,err := db.Exec("UPDATE users SET login = ? WHERE login = ?", userNewLogin, tokenClaims.User_Login)
	
	if(err != nil) {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"errorMessage" : err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "Login updated"})

	
}