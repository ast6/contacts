package Resources

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type AddContactRequest struct {
	FirstName string `db:"FirstName"`
	LastName  string `db:"LastName"`
	Number    string `db:"Number"`
}
