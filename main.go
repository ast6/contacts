package main

import "database/sql"
import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "db:db@/db")
	fmt.Println(db, err)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	_, err = db.Query("CREATE TABLE Users(ID int NOT NULL AUTO_INCREMENT,	FirstName varchar(255) NOT NULL, LastName varchar(255),	Age int,PRIMARY KEY (ID))")
	if err != nil {
		fmt.Println(err)
	}

}
