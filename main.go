package main

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main()  {

	db, err := sql.Open("mysql", "db:db@/db")
	fmt.Println(db,err)
}


