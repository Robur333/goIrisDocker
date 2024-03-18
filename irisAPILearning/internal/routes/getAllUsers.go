package routes

import (
	"fmt"
	"restapi/utils"

	"github.com/kataras/iris/v12"
)

type User struct {
	Username string 
    Password string 
}

func GetAllUsers(ctx iris.Context) {

	db := utils.ConnectToDatabase()

	

	rows,err := db.Query("SELECT login, password FROM users")

	if err != nil { 
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"errorMessage": err.Error()})
		return
	}

	defer rows.Close() 

	users := []User{}
	for rows.Next() {
		var  dbUser, dbPassword string
		if err := rows.Scan(&dbUser, &dbPassword); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"errorMessage": err.Error()})
			return
		}

		users = append(users,User{Username: dbUser, Password: dbPassword} )

		fmt.Println(users)
	}
	
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(users)
}