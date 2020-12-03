package router

import (
	"github.com/go-starter/controller"
	"github.com/gorilla/mux"
)

func InitializeRoutes(ctl *controller.CustomerController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/customers", ctl.GetAll).Methods("GET")
	r.HandleFunc("/api/customer/{id}", ctl.Get).Methods("GET")
	r.HandleFunc("/api/customer", ctl.Post).Methods("POST")
	r.HandleFunc("/api/customer/{id}", ctl.Put).Methods("PUT")
	r.HandleFunc("/api/customer/{id}", ctl.Delete).Methods("DELETE")
	return r
}

