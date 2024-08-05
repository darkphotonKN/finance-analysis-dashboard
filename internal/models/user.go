package models

type User struct {
	ID       int    `db:"id"`
	Email    string `db:"email"` // username
	Password string `db:"password"`
	Role     string `db:"role"`
}
