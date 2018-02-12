package Resources

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AddContactRequest struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Number    string `json:"number"`
}
