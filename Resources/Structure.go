package Resources

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
