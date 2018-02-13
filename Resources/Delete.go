package Resources

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	Request := new(UserTableScheme)

	if err := json.NewDecoder(r.Body).Decode(Request); err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}

	db, err := sql.Open("mysql", "db:db@/db")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = db.Exec("DELETE 	FROM users WHERE id=(?)",
		Request.ID)
}
