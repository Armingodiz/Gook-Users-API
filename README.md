# Gook-Users-API


This is one of the [Gook](https://github.com/Armingodiz/Gook) services which provide DB for saving users and intractions with oauth api for login and creating access token which follow mvc design pattern .

## Features 

* Signup user 
* Find user with specefic id 
* Update users information
* Search for users
* Delete user
* An end point for oauth api to login user

## Dependencies
name     | repo
------------- | -------------
  gin-gonic   | https://github.com/gin-gonic/gin
  mysql driver| github.com/go-sql-driver/mysql
  crypto/md5  | https://golang.org/pkg/crypto/md5/
  

## Installation

First make sure you have installed all dependencies ,
create a scheme in your mysql connection and a table for users then go to `Gook-Users-API/datasources/mysql/users_db/users_db.go`
and set your mysql configurations . 
Then just simply clone this repository and start service with `go run main.go` (your service will be running on `localhost:1111`)


## EndPoints 

	POST == > /users (create or signup the user with given informations as json)
	GET ==> /users/:user_id (return user with given id)
	PUT == > /users/:user_id  (update)
	PATCH ==> /users/:user_id (partial update)
	DELETE ==> /users/:user_id (Delete user with given id)
	GET == > /internal/users/search (Search for users)
	POST ==> /users/login (oauth API use this endpoint for Login)









