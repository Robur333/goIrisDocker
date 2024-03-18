package routes

import (
	"fmt"
	"restapi/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
)


func Login(ctx iris.Context) {
	db := utils.ConnectToDatabase()

	fmt.Println(ctx.URLParam("login"))
	userLogin := ctx.URLParam("login")
	userPassword := ctx.URLParam("password")


	var dbPassword string
	rows := db.QueryRow("SELECT password FROM USERS where login=? ",userLogin).Scan(&dbPassword)
	
	if err := rows; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"errorMessage" : err.Error()})
		return
	}

	if (userPassword != dbPassword){
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"error": "invalid credentials"})
		return
	}

	signer := jwt.NewSigner(jwt.HS256, utils.SignKey, 10*time.Minute)
	ctx.Write(utils.GenerateToken(signer, userLogin))
	ctx.StatusCode(iris.StatusOK)
	ctx.Header("Content-Type", "text/plain")

}


