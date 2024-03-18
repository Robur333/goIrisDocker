package utils

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/kataras/iris/v12/middleware/jwt"
)

func GenerateToken(signer *jwt.Signer, userLogin string) []byte {

	claims := UserClaims{User_Login: userLogin }

	token, err := signer.Sign(claims)
	if err != nil {
		fmt.Println("error")
	}

	return token
}


func ConnectToDatabase() *sql.DB {

db, err := sql.Open("mysql", os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@" + os.Getenv("DB_HOST"))
 
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime( time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db 
}