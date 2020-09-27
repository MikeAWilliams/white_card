package auth

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Database interface {
	GetUser(email string) (User, error)
	AddUser(user User) error
	UserExists(email string) (bool, error)
}
