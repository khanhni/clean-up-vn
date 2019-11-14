package main

import (
	"fmt"
	"net/http"

	"github.com/Lycheeeeeee/clean-up-vn/controllers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/user/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/getone/{id}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/api/user/getall", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user/update/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/project/create", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/api/project/getone/{id}", controllers.GetProjectByID).Methods("GET")
	router.HandleFunc("/api/project/getall", controllers.GetAllProjects).Methods("GET")
	router.HandleFunc("/api/project/inputresult/{id}", controllers.InputResult).Methods("PUT")
	router.HandleFunc("/api/userproject/adduser", controllers.CreateUserProject).Methods("POST")
	//port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	//if port == "" {
	//port = 8000 //localhost
	//}

	fmt.Println("8000")

	err := http.ListenAndServe(":8000", router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}