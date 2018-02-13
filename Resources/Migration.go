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

	_, err = db.Query("CREATE TABLE IF NOT EXISTS users(id INT NOT NULL AUTO_INCREMENT,	firstName VARCHAR(255) NOT NULL, lastName VARCHAR(255),	number VARCHAR(255),PRIMARY KEY (id))")
	if err != nil {
		fmt.Println(err)
	}
}
