package app

import (
	"github.com/akramulfata10/gotoko/app/controllers"
	"github.com/gorilla/mux"

)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
