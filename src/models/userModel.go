package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Username string             `json:"username"`
	Password string             `json:"-"`
}

type UserRepository interface {
	FindByID(id string) (*User, error)
	FindByUsername(username string) (User, error)
	CreateUser(username string, hashedPassword string) (*User, error)
}
