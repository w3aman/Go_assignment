package router

import (
	"root3/controller"

	"github.com/gorilla/mux"
)

var Mux = mux.NewRouter()

func Routers() {
	Mux.HandleFunc("/Create", controller.Create).Methods("GET")
	Mux.HandleFunc("/Update", controller.Update)
	Mux.HandleFunc("/GetById", controller.GetById)
	Mux.HandleFunc("/GetAll", controller.GetAll)
	Mux.HandleFunc("/Delete", controller.Delete)
}
