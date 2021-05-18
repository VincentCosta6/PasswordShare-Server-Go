package models

type User struct {
	ID       string `json:"_id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

type UserRepository interface {
	FindByID(id string) (*User, error)
	FindByUsername(username string) (*User, error)
	Save(user *User) error
}
