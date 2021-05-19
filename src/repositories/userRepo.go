package repositories

import (
	"context"
	"password-share-server-golang/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	users *mongo.Collection
}

func NewUserRepo(db *mongo.Database) *UserRepo {
	return &UserRepo{
		users: db.Collection("users"),
	}
}

func (r *UserRepo) FindByID(id string) (*models.User, error) {
	return &models.User{}, nil
}

func (r *UserRepo) FindByUsername(username string) (models.User, error) {
	var foundUser models.User

	err := r.users.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&foundUser)

	return foundUser, err
}

func (r *UserRepo) CreateUser(username string, hashedPassword string) (*models.User, error) {
	newUser := models.User{primitive.NewObjectID(), username, hashedPassword}

	_, err := r.users.InsertOne(context.TODO(), newUser)

	return &newUser, err
}

func (r *UserRepo) Save(user *models.User) error {
	return nil
}
