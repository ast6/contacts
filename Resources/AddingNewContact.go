package Resources

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Add(w http.ResponseWriter, r *http.Request) {
	request := new(AddContactRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}
	fmt.Println(AddContactRequest{})
}
