package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"inventory_management_system/handler"
	"net/http"
)

func GetRoutes() *mux.Router {
	mainRouter := mux.NewRouter()

	mainRouter.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "connection established...")
	}).Methods("GET")

	publicRoutes := mainRouter.PathPrefix("/api").Subrouter()
	
	//public routes
	publicRoutes.HandleFunc("/user/register", handler.PublicRegister).Methods("POST")

	return mainRouter
}
