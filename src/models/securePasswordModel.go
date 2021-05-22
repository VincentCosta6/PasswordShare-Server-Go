package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SecurePassword struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId        primitive.ObjectID `json:"user_id" bson:"user_id"`
	EncryptedData string             `json:"encrypted_data"`
}

type SecurePasswordRepository interface {
	FindByID(id string) (*SecurePassword, error)
	FindAllByUserId(userId primitive.ObjectID) (securePassword *[]SecurePassword, err error)
	CreateSecurePassword(userId primitive.ObjectID, encryptedData string) (*SecurePassword, error)
	SetEncryptedData(passwordId string, newEncryptedData string) error
}
