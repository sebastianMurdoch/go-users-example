package domain

type User struct {
	Id       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

type UsersRepository interface {
	FindAll() []User
	Save(user User) error
}