package app

import "github.com/akramulfata10/gotoko/app/controllers"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
