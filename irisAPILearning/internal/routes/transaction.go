package routes

import (
	"fmt"
	"restapi/utils"
)

func Transaction() {
	db := utils.ConnectToDatabase()
	defer db.Close()

	tx,err := db.Begin()

	if err != nil {
		fmt.Println(err)
	}


	_,err = tx.Exec("CREATE DATABATE test")
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
	}

	_,err = tx.Exec("CREATE TABLE users (login varchar(20),password varchar(20))")
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
	}

	if err := tx.Commit(); err != nil {
		fmt.Println(err)
    }
}