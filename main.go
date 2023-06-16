package main

import (
	"04-go-mongo-restAPI/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	// Initiating router and assigning it to a variable
	r := httprouter.New()

	// Call and store controllers with session
	uc := controllers.NewUserController(getSession())

	// Declare routes
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	// Creates golang server
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	// Calls mongo db session
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
