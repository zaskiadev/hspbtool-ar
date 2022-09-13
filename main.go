package main

import (
	"github.com/julienschmidt/httprouter"
	"hpbtool-ar/controllers"
	"log"
	"net/http"
)

type Response struct {
	Status  string
	Code    string
	Total   int
	Persons []Persons `json:"data"`
}

type Persons struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Birthday  string
	Gender    string
	Address   Address
	Website   string
	Image     string
}

type Address struct {
	Street         string
	StreetName     string
	BuildingNumber int
	City           string
	Zipcode        int
	Country        string
	CountyCode     string `json:"county_code"`
	Latitude       float64
	Longitude      float64
}

func main() {
	webController := &controllers.WebControllers{}

	router := httprouter.New()
	router.GET("/", webController.Login)
	router.POST("/", webController.Login)
	router.GET("/home", webController.Home)
	router.GET("/add_task", webController.AddTask)
	router.POST("/add_task", webController.AddTask)
	router.GET("/edit_task/:codetask", webController.EditTask)
	router.POST("/edit_task/:codetask", webController.UpdateTask)

	router.GET("/add_comment_task/", webController.AddCommentTask)
	router.POST("/add_comment_task", webController.AddCommentTask)

	router.GET("/show_all_comment_task", webController.ShowAllCommentTask)

	router.POST("/done/:codetask", webController.DoneTask)
	fs := http.FileServer(http.Dir("http/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	log.Fatal(http.ListenAndServe(":6490", router))

}
