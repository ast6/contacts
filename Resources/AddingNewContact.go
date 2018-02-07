package Resources

import (
	"encoding/json"
	"net/http"
)

func Add(w http.ResponseWriter, r *http.Request) {
	request := http.Post()
	err := json.NewDecoder(r.Body).Decode(request)

}
