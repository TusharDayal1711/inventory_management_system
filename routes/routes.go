package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetRoutes() *mux.Router {
	mainRouter := mux.NewRouter()

	mainRouter.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "connection established...")
	}).Methods("GET")

	//
	return mainRouter
}
