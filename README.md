# PasswordShare Server - Golang
This repository is the server used by PasswordShare, an open source alternative to LastPass. It is written in Golang and deployed to Heroku. It uses gin-gonic as its REST server framework. MongoDB is used as its primary database.

# Live Server
PasswordShare Server is deployed to Heroku and is available at https://password-share-server-golang.herokuapp.com/ and uses a free tier of MongoDB Atlas

# Development Server
In order to run this repository locally for development, make sure to create a .env file with the following options:
```
MONGO_JUST_URI = <boolean>
MONGO_CONNECTION_URI = <mongo server URI>
MONGO_USERNAME = <mongo_username>
MONGO_PASSWORD = <mongo_password>
SERVER_PORT = <port> (defaults to 8000)
```
* NOTE: if only using URI for connection set $MONGO_JUST_URI to true and ignore $MONGO_USERNAME and $MONGO_PASSWORD. If using username/password seperately then set this to false and include the $MONGO_USERNAME and $MONGO_PASSWORD environment variables

# Development Dependencies
This repository uses govendor as its dependency manager.

## Adding a dependency
Use the command: 
```
govendor fetch <package>
```
to add new dependencies