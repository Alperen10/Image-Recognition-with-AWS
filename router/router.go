package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Alperen10/Image-Recognition/controller"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	router *mux.Router
)

func CreateRouter() {
	router = mux.NewRouter()
}

func InitializeRoute() {
	router.HandleFunc("/api/image-recognition", controller.ImageController).Methods("POST")
}

func ServerStarter() {
	fmt.Println("Server is running on port 3000")
	err := http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{
		"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization",
	}),
		handlers.AllowedMethods([]string{
			"POST", "GET", "PUT", "DELETE",
		}),
		handlers.AllowedOrigins([]string{("*")}),
	)(router))
	if err != nil {
		log.Fatal(err)
	}
}
