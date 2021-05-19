package repositories

import (
	"context"
	"password-share-server-golang/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SecurePassworRepo struct {
	securePasswords *mongo.Collection
}

func NewSecurePasswordRepo(db *mongo.Database) *SecurePassworRepo {
	return &SecurePassworRepo{
		securePasswords: db.Collection("secure-passwords"),
	}
}

func (r *SecurePassworRepo) FindByID(id string) (*models.SecurePassword, error) {
	var foundPassword models.SecurePassword

	err := r.securePasswords.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&foundPassword)

	return &foundPassword, err
}

func (r *SecurePassworRepo) FindAllByUserId(userId string) (*[]models.SecurePassword, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var foundSecurePasswords []models.SecurePassword

	cursor, err := r.securePasswords.Find(context.TODO(), bson.D{primitive.E{Key: "user_id", Value: userId}})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result models.SecurePassword
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		foundSecurePasswords = append(foundSecurePasswords, result)
	}

	return &foundSecurePasswords, err
}

func (r *SecurePassworRepo) CreateSecurePassword(userId string, encryptedData string) (*models.SecurePassword, error) {
	convertedUserId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return nil, err
	}

	newSecurePassword := models.SecurePassword{ID: primitive.NewObjectID(), UserId: convertedUserId, EncryptedData: encryptedData}

	_, err = r.securePasswords.InsertOne(context.TODO(), newSecurePassword)

	return &newSecurePassword, err
}

func (r *SecurePassworRepo) SetEncryptedData(passwordId string, newEncryptedData string) error {
	convertedId, err := primitive.ObjectIDFromHex(passwordId)

	if err != nil {
		return err
	}

	_, err = r.securePasswords.UpdateOne(
		context.TODO(),
		bson.M{"_id": convertedId},
		bson.D{
			primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "encrypted_data", Value: newEncryptedData}}},
		})

	return err
}
