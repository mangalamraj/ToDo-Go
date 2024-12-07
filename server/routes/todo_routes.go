package routes

import (
	"todo/controller"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/addtodo", controller.AddTodo).Methods("POST")
	return router
}
