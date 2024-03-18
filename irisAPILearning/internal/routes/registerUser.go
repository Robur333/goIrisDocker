package routes

import (
	"database/sql"
	"restapi/utils"

	"github.com/kataras/iris/v12"
)

func Register(ctx iris.Context) {
    db := utils.ConnectToDatabase()
  
    userLogin := ctx.URLParam("login")
    userPassword := ctx.URLParam("password")
    var dbLogin string
    err := db.QueryRow("SELECT login FROM users WHERE login = ?", userLogin).Scan(&dbLogin)

    if err != nil {
        if err == sql.ErrNoRows {
            _, err := db.Exec("INSERT INTO users (login, password) VALUES (?, ?)", userLogin, userPassword)
            if err != nil {
                ctx.StatusCode(iris.StatusInternalServerError)
                ctx.JSON(iris.Map{"message": "Failed to create account"})
                return
            }
            ctx.StatusCode(iris.StatusOK)
            ctx.JSON(iris.Map{"message": "Account created"})
            return
        } else {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{"message": err.Error()})
            return
        }
    }

    if dbLogin == userLogin {
        ctx.StatusCode(iris.StatusBadRequest)
        ctx.JSON(iris.Map{"errorMessage": "Login is already taken"})
    }
}