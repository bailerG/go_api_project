package routes

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go_api_project/controllers"
	middleware "go_api_project/middlewares"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalities", controllers.GetPersonalities).Methods("Get")
	r.HandleFunc("/personalities", controllers.CreateNewPersonality).Methods("Post")
	r.HandleFunc("/personalities/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/personalities/{id}", controllers.GetSinglePersonality).Methods("Get")
	r.HandleFunc("/personalities/{id}", controllers.EditPersonality).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
