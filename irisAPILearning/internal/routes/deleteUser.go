package routes

import (
	"restapi/utils"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)

func DeleteUser(ctx iris.Context) {
	 
	db := utils.ConnectToDatabase()
	defer db.Close()

	tokenClaims := jwt.Get(ctx).(*utils.UserClaims)

	_, err :=  db.Exec("DELETE FROM users WHERE login = ?", tokenClaims.User_Login)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"errorMessage" : err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message" : "Account Deleted"})
	ctx.Logout()
}