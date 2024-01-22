package main

import (
	"github.com/julienschmidt/httprouter",
	"net/http",
	"gopkg.in/mgo.v2"

)


func main(){

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id",uc.GetUser)

	http.ListenAndServe("localhost:9000",r)
}

func getSession() *mgo.Session{
	s, err := mgo.Dail("mongodb://localhost:27107")
	if err != nil{
		panic(err)
	}
	return s
}