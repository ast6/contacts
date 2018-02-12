package Resources

import (
	"database/sql"
	"fmt"
)

func Migration() {
	db, err := sql.Open("mysql", "db:db@/db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	_, err = db.Query("CREATE TABLE IF NOT EXISTS Users(ID INT NOT NULL AUTO_INCREMENT,	FirstName VARCHAR(255) NOT NULL, LastName VARCHAR(255),	Number VARCHAR(255),PRIMARY KEY (ID))")
	if err != nil {
		fmt.Println(err)
	}
}
