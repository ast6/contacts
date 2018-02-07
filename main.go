package main

import (
	"github.com/ast6/contacts/Resources"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	Resources.Migration()
	Resources.Add()

}
