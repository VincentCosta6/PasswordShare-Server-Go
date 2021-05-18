package repositories

import (
	"password-share-server-golang/src/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	db *mongo.Client
}

func NewUserRepo(db *mongo.Client) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) FindByID(id string) (*models.User, error) {
	return &models.User{}, nil
}

func (r *UserRepo) FindByUsername(username string) (*models.User, error) {
	return &models.User{}, nil
}

func (r *UserRepo) Save(user *models.User) error {
	return nil
}
