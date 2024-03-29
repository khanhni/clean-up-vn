package main

import (
	"fmt"
	"net/http"

	"github.com/Lycheeeeeee/clean-up-vn/controllers"
	"github.com/gorilla/mux"

	// "github.com/gorilla/handlers"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()
	// router.Use(app.JwtAuthentication)
	// cors := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	//   })

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})
	handler := c.Handler(router)
	//   router.Use(cors.Handler
	//attach JWT auth middleware
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/user/sociallogin/{id}", controllers.Socialauthenticate).Methods("GET")
	router.HandleFunc("/api/user/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/getone/{id}", controllers.GetUserByID).Methods("GET")
	router.HandleFunc("/api/user/getall", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user/update/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/project/create/{time}", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/api/project/getone/{id}", controllers.GetProjectByID).Methods("GET")
	router.HandleFunc("/api/project/getallbyowner/{id}", controllers.GetAllProjectsByOwner).Methods("GET")
	router.HandleFunc("/api/project/getallbyuser/{id}", controllers.GetAllProjectsByUser).Methods("GET")
	router.HandleFunc("/api/project/getall", controllers.GetAllProjects).Methods("GET")
	router.HandleFunc("/api/project/inputresult/{id}", controllers.InputResult).Methods("POST")
	router.HandleFunc("/api/userproject/adduser", controllers.CreateUserProject).Methods("POST")
	router.HandleFunc("/api/userproject/leaveproject", controllers.DeleteUserProject).Methods("POST")
	router.HandleFunc("/api/project/downloadvolunteerlist/{userid}", controllers.DownloadFile).Methods("POST")
	router.HandleFunc("/api/project/testing", controllers.Testing).Methods("GET")
	router.HandleFunc("/api/report/runreport", controllers.Report).Methods("GET")
	// router.HandleFunc("/api/project/numberofvolunteer/{id}", controllers.GetVolunteer).Methods("GET")

	//port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	//if port == "" {
	//port = 8000 //localhost
	//}

	fmt.Println("8000")
	// _ = cors.Default().Handler(router)
	err := http.ListenAndServe(":8000", handler) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
