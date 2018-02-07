package Resources

import (
	"database/sql"
	"fmt"
)

func Migration() {
	db, err := sql.Open("mysql", "db:db@/db")
	fmt.Println(db, err)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	_, err = db.Query("CREATE TABLE IF NOT EXISTS Users(ID int NOT NULL AUTO_INCREMENT,	FirstName varchar(255) NOT NULL, LastName varchar(255),	Number LINESTRING,PRIMARY KEY (ID))")
	if err != nil {
		fmt.Println(err)
	}
}
