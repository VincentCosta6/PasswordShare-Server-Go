package main

import (
	"context"
	"log"
	"os"

	"password-share-server-golang/src/controllers"
	"password-share-server-golang/src/repositories"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
		return
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatal("$SERVER_PORT must be set")
		return
	}

	if os.Getenv("MONGO_CONNECTION_URI") == "" {
		log.Fatal("$MONGO_CONNECTION_URI must be set")
		return
	}

	if os.Getenv("MONGO_JUST_URI") == "" {
		log.Fatal("$MONGO_JUST_URI must be set")
		return
	} else if os.Getenv("MONGO_JUST_URI") == "false" {
		if os.Getenv("MONGO_USERNAME") == "" || os.Getenv("MONGO_PASSWORD") == "" {
			log.Fatal("$MONGO_JUST_URI is false but $MONGO_USERNAME or $MONGO_PASSWORD was missing")
			return
		}
	}

	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "success", "message": "Successfully retrieved /"})
	})

	client := initMongoDBConnection()

	userRepo := repositories.NewUserRepo(client)
	userController := controllers.NewUserController(userRepo)

	router.GET("/test", userController.SuccessRoute)

	router.Run(":" + port)
}

func initMongoDBConnection() *mongo.Client {
	mongoConnectionURI := os.Getenv("MONGO_CONNECTION_URI")

	clientOptions := options.Client().ApplyURI(mongoConnectionURI)
	if os.Getenv("MONGO_JUST_URI") == "false" {
		credential := options.Credential{
			Username: os.Getenv("MONGO_USERNAME"),
			Password: os.Getenv("MONGO_PASSWORD"),
		}
		clientOptions = options.Client().ApplyURI(mongoConnectionURI).SetAuth(credential)
	}

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return client
}
